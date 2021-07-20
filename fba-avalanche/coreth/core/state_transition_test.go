package core

import (
	"math/big"
	"testing"

	"github.com/ava-labs/coreth/core/vm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Define the default EVM mock and define default mock receiver functions
type StateTransitionEVMMock struct {
	mockEVMCallerData MockEVMCallerData
	log               log.Logger
}

func (e *StateTransitionEVMMock) Call(caller vm.ContractRef, addr common.Address, input []byte, gas uint64, value *big.Int) (ret []byte, leftOverGas uint64, err error) {
	return defautCall(&e.mockEVMCallerData, caller, addr, input, gas, value)
}

func (e *StateTransitionEVMMock) GetBlockNumber() *big.Int {
	return defaultGetBlockNumber(&e.mockEVMCallerData)
}

func (e *StateTransitionEVMMock) GetGasLimit() uint64 {
	return defaultGetGasLimit(&e.mockEVMCallerData)
}

func (e *StateTransitionEVMMock) AddBalance(addr common.Address, amount *big.Int) {
	defaultAddBalance(&e.mockEVMCallerData, addr, amount)
}

func TestKeeperTriggerFiredAndMinted(t *testing.T) {
	mintRequestReturn, _ := new(big.Int).SetString("50000000000000000000000000", 10)
	mockEVMCallerData := &MockEVMCallerData{
		blockNumber:       *big.NewInt(0),
		gasLimit:          0,
		mintRequestReturn: *mintRequestReturn,
	}
	stateTransitionEVMMock := &StateTransitionEVMMock{
		mockEVMCallerData: *mockEVMCallerData,
		log:               log.New(),
	}

	triggerKeeperAndMint(stateTransitionEVMMock)

	if mintRequest.Cmp(mintRequestReturn) != 0 {
		t.Errorf("got %s want %q", mintRequest.Text(10), "50000000000000000000000000")
	}
}
