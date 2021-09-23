// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func TestPrecompiledContractSpendsGas(t *testing.T) {
	unwrapped := &sha256hash{}

	input := []byte{'J', 'E', 'T', 'S'}
	requiredGas := unwrapped.RequiredGas(input)
	_, remainingGas, err := RunPrecompiledContract(unwrapped, input, requiredGas)
	if err != nil {
		t.Fatalf("Unexpectedly failed to run precompiled contract: %s", err)
	}

	if remainingGas != 0 {
		t.Fatalf("Found more remaining gas than expected: %d", remainingGas)
	}
}

// CanTransfer checks whether there are enough funds in the address' account to make a transfer.
// This does not take the necessary gas in to account to make the transfer valid.
func CanTransfer(db StateDB, addr common.Address, amount *big.Int) bool {
	return db.GetBalance(addr).Cmp(amount) >= 0
}

func CanTransferMC(db StateDB, addr common.Address, to common.Address, coinID *common.Hash, amount *big.Int) bool {
	log.Info("CanTransferMC", "address", addr, "to", to, "coinID", coinID, "amount", amount)
	if coinID == nil {
		return true
	}
	if db.GetBalanceMultiCoin(addr, *coinID).Cmp(amount) >= 0 {
		return true
	}
	// insufficient balance
	return false
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func Transfer(db StateDB, sender, recipient common.Address, amount *big.Int) {
	db.SubBalance(sender, amount)
	db.AddBalance(recipient, amount)
}

// Transfer subtracts amount from sender and adds amount to recipient using the given Db
func TransferMultiCoin(db StateDB, sender, recipient common.Address, coinID *common.Hash, amount *big.Int) {
	if coinID == nil {
		return
	}
	db.SubBalanceMultiCoin(sender, *coinID, amount)
	db.AddBalanceMultiCoin(recipient, *coinID, amount)
}
