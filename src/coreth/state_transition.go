// (c) 2021, Flare Networks Limited. All rights reserved.
//
// This file is a derived work, based on the avalanchego library whose original
// notices appear below. It is distributed under a license compatible with the
// licensing terms of the original code from which it is derived.
// Please see the file LICENSE_AVALABS for licensing terms of the original work.
// Please see the file LICENSE for licensing terms.
//
// (c) 2019-2020, Ava Labs, Inc.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/core/vm"
	"github.com/ava-labs/coreth/params"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

var emptyCodeHash = crypto.Keccak256Hash(nil)

/*
The State Transitioning Model

A state transition is a change made when a transaction is applied to the current world state
The state transitioning model does all the necessary work to work out a valid new state root.

1) Nonce handling
2) Pre pay gas
3) Create a new state object if the recipient is \0*32
4) Value transfer
== If contract creation ==
  4a) Attempt to run transaction data
  4b) If valid, use result as code for the new state object
== end ==
5) Run Script section
6) Derive new state root
*/
type StateTransition struct {
	gp         *GasPool
	msg        Message
	gas        uint64
	gasPrice   *big.Int
	gasFeeCap  *big.Int
	gasTipCap  *big.Int
	initialGas uint64
	value      *big.Int
	data       []byte
	state      vm.StateDB
	evm        *vm.EVM
}

// Message represents a message sent to a contract.
type Message interface {
	From() common.Address
	To() *common.Address

	GasPrice() *big.Int
	GasFeeCap() *big.Int
	GasTipCap() *big.Int
	Gas() uint64
	Value() *big.Int

	Nonce() uint64
	IsFake() bool
	Data() []byte
	AccessList() types.AccessList
}

// ExecutionResult includes all output after executing given evm
// message no matter the execution itself is successful or not.
type ExecutionResult struct {
	UsedGas    uint64 // Total used gas but include the refunded gas
	Err        error  // Any error encountered during the execution(listed in core/vm/errors.go)
	ReturnData []byte // Returned data from evm(function result or data supplied with revert opcode)
}

// Unwrap returns the internal evm error which allows us for further
// analysis outside.
func (result *ExecutionResult) Unwrap() error {
	return result.Err
}

// Failed returns the indicator whether the execution is successful or not
func (result *ExecutionResult) Failed() bool { return result.Err != nil }

// Return is a helper function to help caller distinguish between revert reason
// and function return. Return returns the data after execution if no error occurs.
func (result *ExecutionResult) Return() []byte {
	if result.Err != nil {
		return nil
	}
	return common.CopyBytes(result.ReturnData)
}

// Implement the EVMCaller interface on the state transition structure; simply delegate the calls
func (st *StateTransition) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	return st.evm.Call(caller, addr, input, gas, value)
}

func (st *StateTransition) GetBlockNumber() *big.Int {
	return st.evm.Context.BlockNumber
}

func (st *StateTransition) GetGasLimit() uint64 {
	return st.evm.Context.GasLimit
}

func (st *StateTransition) AddBalance(addr common.Address, amount *big.Int) {
	st.state.AddBalance(addr, amount)
}

// Revert returns the concrete revert reason if the execution is aborted by `REVERT`
// opcode. Note the reason can be nil if no data supplied with revert opcode.
func (result *ExecutionResult) Revert() []byte {
	if result.Err != vm.ErrExecutionReverted {
		return nil
	}
	return common.CopyBytes(result.ReturnData)
}

// IntrinsicGas computes the 'intrinsic gas' for a message with the given data.
func IntrinsicGas(data []byte, accessList types.AccessList, isContractCreation bool, isHomestead, isEIP2028 bool) (uint64, error) {
	// Set the starting gas for the raw transaction
	var gas uint64
	if isContractCreation && isHomestead {
		gas = params.TxGasContractCreation
	} else {
		gas = params.TxGas
	}
	// Bump the required gas by the amount of transactional data
	if len(data) > 0 {
		// Zero and non-zero bytes are priced differently
		var nz uint64
		for _, byt := range data {
			if byt != 0 {
				nz++
			}
		}
		// Make sure we don't exceed uint64 for all data combinations
		nonZeroGas := params.TxDataNonZeroGasFrontier
		if isEIP2028 {
			nonZeroGas = params.TxDataNonZeroGasEIP2028
		}
		if (math.MaxUint64-gas)/nonZeroGas < nz {
			return 0, ErrGasUintOverflow
		}
		gas += nz * nonZeroGas

		z := uint64(len(data)) - nz
		if (math.MaxUint64-gas)/params.TxDataZeroGas < z {
			return 0, ErrGasUintOverflow
		}
		gas += z * params.TxDataZeroGas
	}
	if accessList != nil {
		gas += uint64(len(accessList)) * params.TxAccessListAddressGas
		gas += uint64(accessList.StorageKeys()) * params.TxAccessListStorageKeyGas
	}
	return gas, nil
}

// NewStateTransition initialises and returns a new state transition object.
func NewStateTransition(evm *vm.EVM, msg Message, gp *GasPool) *StateTransition {
	return &StateTransition{
		gp:        gp,
		evm:       evm,
		msg:       msg,
		gasPrice:  msg.GasPrice(),
		gasFeeCap: msg.GasFeeCap(),
		gasTipCap: msg.GasTipCap(),
		value:     msg.Value(),
		data:      msg.Data(),
		state:     evm.StateDB,
	}
}

// ApplyMessage computes the new state by applying the given message
// against the old state within the environment.
//
// ApplyMessage returns the bytes returned by any EVM execution (if it took place),
// the gas used (which includes gas refunds) and an error if it failed. An error always
// indicates a core error meaning that the message would always fail for that particular
// state and would never be accepted within a block.
func ApplyMessage(evm *vm.EVM, msg Message, gp *GasPool) (*ExecutionResult, error) {
	return NewStateTransition(evm, msg, gp).TransitionDb()
}

// to returns the recipient of the message.
func (st *StateTransition) to() common.Address {
	if st.msg == nil || st.msg.To() == nil /* contract creation */ {
		return common.Address{}
	}
	return *st.msg.To()
}

func (st *StateTransition) buyGas() error {
	mgval := new(big.Int).SetUint64(st.msg.Gas())
	mgval = mgval.Mul(mgval, st.gasPrice)
	balanceCheck := mgval
	if st.gasFeeCap != nil {
		balanceCheck = new(big.Int).SetUint64(st.msg.Gas())
		balanceCheck.Mul(balanceCheck, st.gasFeeCap)
		balanceCheck.Add(balanceCheck, st.value)
	}
	if have, want := st.state.GetBalance(st.msg.From()), balanceCheck; have.Cmp(want) < 0 {
		return fmt.Errorf("%w: address %v have %v want %v", ErrInsufficientFunds, st.msg.From().Hex(), have, want)
	}
	if err := st.gp.SubGas(st.msg.Gas()); err != nil {
		return err
	}
	st.gas += st.msg.Gas()

	st.initialGas = st.msg.Gas()
	st.state.SubBalance(st.msg.From(), mgval)
	return nil
}

func (st *StateTransition) preCheck() error {
	// Only check transactions that are not fake
	if !st.msg.IsFake() {
		// Make sure this transaction's nonce is correct.
		stNonce := st.state.GetNonce(st.msg.From())
		if msgNonce := st.msg.Nonce(); stNonce < msgNonce {
			return fmt.Errorf("%w: address %v, tx: %d state: %d", ErrNonceTooHigh,
				st.msg.From().Hex(), msgNonce, stNonce)
		} else if stNonce > msgNonce {
			return fmt.Errorf("%w: address %v, tx: %d state: %d", ErrNonceTooLow,
				st.msg.From().Hex(), msgNonce, stNonce)
		}
		// Make sure the sender is an EOA
		if codeHash := st.state.GetCodeHash(st.msg.From()); codeHash != emptyCodeHash && codeHash != (common.Hash{}) {
			return fmt.Errorf("%w: address %v, codehash: %s", ErrSenderNoEOA,
				st.msg.From().Hex(), codeHash)
		}
	}
	// Make sure that transaction gasFeeCap is greater than the baseFee (post london)
	if st.evm.ChainConfig().IsApricotPhase3(st.evm.Context.Time) {
		// Skip the checks if gas fields are zero and baseFee was explicitly disabled (eth_call)
		if !st.evm.Config.NoBaseFee || st.gasFeeCap.BitLen() > 0 || st.gasTipCap.BitLen() > 0 {
			if l := st.gasFeeCap.BitLen(); l > 256 {
				return fmt.Errorf("%w: address %v, maxFeePerGas bit length: %d", ErrFeeCapVeryHigh,
					st.msg.From().Hex(), l)
			}
			if l := st.gasTipCap.BitLen(); l > 256 {
				return fmt.Errorf("%w: address %v, maxPriorityFeePerGas bit length: %d", ErrTipVeryHigh,
					st.msg.From().Hex(), l)
			}
			if st.gasFeeCap.Cmp(st.gasTipCap) < 0 {
				return fmt.Errorf("%w: address %v, maxPriorityFeePerGas: %s, maxFeePerGas: %s", ErrTipAboveFeeCap,
					st.msg.From().Hex(), st.gasTipCap, st.gasFeeCap)
			}
			// This will panic if baseFee is nil, but basefee presence is verified
			// as part of header validation.
			if st.gasFeeCap.Cmp(st.evm.Context.BaseFee) < 0 {
				return fmt.Errorf("%w: address %v, maxFeePerGas: %s baseFee: %s", ErrFeeCapTooLow,
					st.msg.From().Hex(), st.gasFeeCap, st.evm.Context.BaseFee)
			}
		}
	}
	return st.buyGas()
}

// TransitionDb will transition the state by applying the current message and
// returning the evm execution result with following fields.
//
// - used gas:
//      total gas used (including gas being refunded)
// - returndata:
//      the returned data from evm
// - concrete execution error:
//      various **EVM** error which aborts the execution,
//      e.g. ErrOutOfGas, ErrExecutionReverted
//
// However if any consensus issue encountered, return the error directly with
// nil evm execution result.
func (st *StateTransition) TransitionDb() (*ExecutionResult, error) {
	// First check this message satisfies all consensus rules before
	// applying the message. The rules include these clauses
	//
	// 1. the nonce of the message caller is correct
	// 2. caller has enough balance to cover transaction fee(gaslimit * gasprice)
	// 3. the amount of gas required is available in the block
	// 4. the purchased gas is enough to cover intrinsic usage
	// 5. there is no overflow when calculating intrinsic gas
	// 6. caller has enough balance to cover asset transfer for **topmost** call

	// Check clauses 1-3, buy gas if everything is correct
	if err := st.preCheck(); err != nil {
		return nil, err
	}
	msg := st.msg
	sender := vm.AccountRef(msg.From())
	homestead := st.evm.ChainConfig().IsHomestead(st.evm.Context.BlockNumber)
	istanbul := st.evm.ChainConfig().IsIstanbul(st.evm.Context.BlockNumber)
	apricotPhase1 := st.evm.ChainConfig().IsApricotPhase1(st.evm.Context.Time)

	contractCreation := msg.To() == nil

	// Check clauses 4-5, subtract intrinsic gas if everything is correct
	gas, err := IntrinsicGas(st.data, st.msg.AccessList(), contractCreation, homestead, istanbul)
	if err != nil {
		return nil, err
	}
	if st.gas < gas {
		return nil, fmt.Errorf("%w: have %d, want %d", ErrIntrinsicGas, st.gas, gas)
	}
	st.gas -= gas

	// Check clause 6
	if msg.Value().Sign() > 0 && !st.evm.Context.CanTransfer(st.state, msg.From(), msg.Value()) {
		return nil, fmt.Errorf("%w: address %v", ErrInsufficientFundsForTransfer, msg.From().Hex())
	}

	// Set up the initial access list.
	if rules := st.evm.ChainConfig().AvalancheRules(st.evm.Context.BlockNumber, st.evm.Context.Time); rules.IsApricotPhase2 {
		st.state.PrepareAccessList(msg.From(), msg.To(), vm.ActivePrecompiles(rules), msg.AccessList())
	}

	var (
		ret                                       []byte
		vmerr                                     error // vm errors do not affect consensus and are therefore not assigned to err
		selectProveDataAvailabilityPeriodFinality bool
		selectProvePaymentFinality                bool
		selectDisprovePaymentFinality             bool
		prioritisedFTSOContract                   bool
	)

	// In Avalanche, every block has a hard-coded coinbase of `0x01...`. This
	// sanity check ensures that we only execute transactions that comply with
	// this constraint. This is important, because we use some coinbase magic
	// later when interacting with the state connector.
	if st.evm.Context.Coinbase != common.HexToAddress("0x0100000000000000000000000000000000000000") {
		return nil, fmt.Errorf("Invalid value for block.coinbase")
	}

	// The first condition is another sanity check, which ensures that we reject
	// transactions from the hard-coded coinbase address of `0x01...`, in case
	// anyone was ever able to derive the private key for the address.
	// We also make sure that no transactions from privileged Flare contracts,
	// such as the state connector on `0x10...01` or the system trigger on
	// `0x10...02`, are accepted. This is a sanity check that is a backstop
	// against exploits on the contracts that would use them to issue
	// transactions to escalate privileges.
	if st.msg.From() == common.HexToAddress("0x0100000000000000000000000000000000000000") ||
		st.msg.From() == common.HexToAddress(GetStateConnectorContractAddr(st.evm.Context.Time)) ||
		st.msg.From() == common.HexToAddress(GetSystemTriggerContractAddr(st.evm.Context.Time)) {
		return nil, fmt.Errorf("Invalid sender")
	}

	// In Avalanche, tokens consumed as gas during a transaction are simply
	// burned by sending them to the hard-coded coinbase address. We do the same
	// in Flare, but due to the coinbase magic we do later, we explicitly refer
	// to the coinbase address as burn address before we manipulate it.
	burnAddress := st.evm.Context.Coinbase

	// If we don't have a contract creation, which is equivalent of having a
	// destination address for the transaction (as checked earlier), we evaluate
	// the destination address to see if we are dealing with a special kind of
	// transaction within the Flare system.
	if !contractCreation {

		// If the transaction goes to the state connector contract address, we
		// determine what kind of transaction it is by checking the first four
		// bytes of the call data, which allows us to identify the type of
		// proof.
		// Otherwise, we check if the transaction goes to a prioritized FSTO
		// contract, which is located at `0x10...03`.
		if *msg.To() == common.HexToAddress(GetStateConnectorContractAddr(st.evm.Context.Time)) {
			selectProveDataAvailabilityPeriodFinality = bytes.Equal(st.data[0:4], GetProveDataAvailabilityPeriodFinalitySelector(st.evm.Context.Time))
			selectProvePaymentFinality = bytes.Equal(st.data[0:4], GetProvePaymentFinalitySelector(st.evm.Context.Time))
			selectDisprovePaymentFinality = bytes.Equal(st.data[0:4], GetDisprovePaymentFinalitySelector(st.evm.Context.Time))
		} else {
			prioritisedFTSOContract = *msg.To() == common.HexToAddress(GetPrioritisedFTSOContract(st.evm.Context.Time))
		}
	}

	// In case the previous check has determined that we are dealing with a
	// state connector transaction, of which we were able to identify the type,
	// we prepare and execute the state connector call.
	// Otherwise, we just go into the original transaction calling logic below.
	if selectProveDataAvailabilityPeriodFinality || selectProvePaymentFinality || selectDisprovePaymentFinality {

		// Transactions that are not contract creations need to increase the
		// nonce on the sending acount. As we are bypassing the normal execution
		// logic here, we have to increase it, too.
		st.state.SetNonce(msg.From(), st.state.GetNonce(sender.Address())+1)

		// The gas given to state connector calls is a third of the gas given to
		// normal transaction calls. As a safe upper limit, a user should thus
		// put the gas limit for state connector transactions at 3x what he
		// would give the transaction if it was executed normally. Please note
		// that this is only the upper limit, and the actual consumption will be
		// significantly lower in practice, as explained below.
		stateConnectorGas := st.gas / GetStateConnectorGasDivisor(st.evm.Context.Time)

		// The code below executes the state connector call, which interacts
		// with the respective chain API, sandwiched between two EVM calls that
		// are executed against the state connector smart contract on the chain.
		// 1) The first EVM call executes the smart contract call with a
		// cainbase of `0x01...`, which makes it a "read-only" transaction from
		// the perspective of the state connector smart contract. However, it
		// checks whether the call would, in theory, be successful.
		// 2) If the first EVM call is successful, we achieved two things: we
		// avoid making expensive state connector calls for transactions that
		// would fail anyway, and we have the return data we need to do the
		// state connector call. We thus proceed to the state connector call,
		// which either retrieves or checks chain data from the remote chain
		// API.
		// 3) Regardless of whether the first EVM call or the state connector
		// call succeed, we execute a second EVM call. However, if the state
		// connector call succeeds, we set the coinbase to the sender's address
		// and thus enable the second EVM call to update the state connector
		// smart contract's storage.
		// NOTE: As the first EVM call is read-only, it will consume
		// significantly less gas than the second EVM call. Additionally, a
		// third of the gas is reserved for the state connector call, but in
		// effect always remains unused from the perspective of the EVM. This
		// means that the effective gas consumption is somewhere between 1-2x
		// what it would be if it was executed as a normal transaction, with
		// 1x being the baseline from the second EVM call, and 0.X-1.0x being
		// added on top for the first read-only EVM call.
		// TODO: There is probably a way to optimize the gas limit for state
		// connector call transactions by properly estimating the actual gas
		// consumed by the first and second EVM calls, and by using a fixed cost
		// for the state connector call.
		// TODO: As the state connector call never really consumes gas, it does
		// not actually carry a cost for the transaction issuer. Nothing stops
		// the user from using a very high gas limit and having the gas refunded
		// afterwards. This is not entirely true, as the refund logic caps the
		// refund at a certain number, to ensure that users can't use crazy gas
		// limits and fill blocks with transactions that do very little. It
		// would still be a good idea, though, to actually implement a fixed
		// cost for the state connector call that is actually consumed, as it is
		// constitutes an expensive operation for the node operator and could be
		// abused for denial of service and similar exploits.
		// If there is no state connector call, we simply go into the normal
		// transaction execution logic, that either creates a contract or
		// executes a simple call against the EVM.
		checkRet, _, checkVmerr := st.evm.Call(sender, st.to(), st.data, stateConnectorGas, st.value)
		if checkVmerr == nil {
			chainConfig := st.evm.ChainConfig()
			if GetStateConnectorActivated(chainConfig.ChainID, st.evm.Context.Time) && binary.BigEndian.Uint32(checkRet[28:32]) < GetMaxAllowedChains(st.evm.Context.Time) {
				if StateConnectorCall(msg.From(), st.evm.Context.Time, st.data[0:4], checkRet) {
					originalCoinbase := st.evm.Context.Coinbase
					defer func() {
						st.evm.Context.Coinbase = originalCoinbase
					}()
					st.evm.Context.Coinbase = st.msg.From()
				}
			}
		}
		ret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, stateConnectorGas, st.value)
	} else {
		if contractCreation {
			ret, _, st.gas, vmerr = st.evm.Create(sender, st.data, st.gas, st.value)
		} else {
			st.state.SetNonce(msg.From(), st.state.GetNonce(sender.Address())+1)
			ret, st.gas, vmerr = st.evm.Call(sender, st.to(), st.data, st.gas, st.value)
		}
	}

	// After the transaction, Avalanche usually refunds the remaining gas, but
	// it is capped at half of the used gas to make sure people use reasonable
	// gas limits in their transactions.
	st.refundGas(apricotPhase1)

	// On top of the normal gas refund for unused gas, Flare applies an
	// additional refund for transactions that are sent to the FTSO contract.
	// This is done so the system remains stable under heavy load, and spam to
	// other accounts can not affect the FTSO system. Basically, the fee for
	// FTSO contract calls is hard-coded to a certain nominal gas use and gas
	// price. If the gas cost for the transaction exceeds this amount, the
	// difference is refunded to the sender.
	if vmerr == nil && prioritisedFTSOContract {
		nominalGasUsed := uint64(21000)
		nominalGasPrice := uint64(225_000_000_000)
		nominalFee := new(big.Int).Mul(new(big.Int).SetUint64(nominalGasUsed), new(big.Int).SetUint64(nominalGasPrice))
		actualGasUsed := st.gasUsed()
		actualGasPrice := st.gasPrice
		actualFee := new(big.Int).Mul(new(big.Int).SetUint64(actualGasUsed), actualGasPrice)
		if actualFee.Cmp(nominalFee) > 0 {
			feeRefund := new(big.Int).Sub(actualFee, nominalFee)
			st.state.AddBalance(st.msg.From(), feeRefund)
			st.state.AddBalance(burnAddress, nominalFee)
		} else {
			st.state.AddBalance(burnAddress, actualFee)
		}
	} else {
		st.state.AddBalance(burnAddress, new(big.Int).Mul(new(big.Int).SetUint64(st.gasUsed()), st.gasPrice))
	}

	// After every successful transaction, we also trigger the keeper contract,
	// so it can take care of some book-keeping tasks for the Flare system. We
	// do so without EVM debugging, and the keeper has the ability to call back
	// into the EVM for a number of reasons with the additional three methods
	// that we expose (`Call`, `GetBlockNumber`, `GetGasLimit`, `AddBalance`).
	// Previously, all keeper calls beyond the first one in a block would
	// `revert`, which would be undesirable, as it creates errors for expected
	// behaviour. This was therefore changed to `return`.
	if vmerr == nil {
		oldDebug := st.evm.Config.Debug
		st.evm.Config.Debug = false
		log := log.Root()
		triggerKeeperAndMint(st, log)
		st.evm.Config.Debug = oldDebug
	}

	return &ExecutionResult{
		UsedGas:    st.gasUsed(),
		Err:        vmerr,
		ReturnData: ret,
	}, nil
}

func (st *StateTransition) refundGas(apricotPhase1 bool) {
	// Inspired by: https://gist.github.com/holiman/460f952716a74eeb9ab358bb1836d821#gistcomment-3642048
	if !apricotPhase1 {
		// Apply refund counter, capped to half of the used gas.
		refund := st.gasUsed() / 2
		if refund > st.state.GetRefund() {
			refund = st.state.GetRefund()
		}
		st.gas += refund
	}

	// Return ETH for remaining gas, exchanged at the original rate.
	remaining := new(big.Int).Mul(new(big.Int).SetUint64(st.gas), st.gasPrice)
	st.state.AddBalance(st.msg.From(), remaining)

	// Also return remaining gas to the block gas counter so it is
	// available for the next transaction.
	st.gp.AddGas(st.gas)
}

// gasUsed returns the amount of gas used up by the state transition.
func (st *StateTransition) gasUsed() uint64 {
	return st.initialGas - st.gas
}
