package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {
}

// the AVA must be compiled with
// export GENESIS_FILE=genesis_scdev_160k.go
func Test_Genesis_Existing_Balance(t *testing.T) {
	fmt.Println("Connecting to validator")
	//
	client, err := ethclient.Dial("http://127.0.0.1:9660/ext/bc/C/rpc")
	if err != nil {
		log.Fatal(err)
	}

	// use last address from genesis_scdev_160k
	account := common.HexToAddress("d74678d4980c8a0777a230772d8fd9ec515d2c6c")

	// Pass nil in third parameter to get balance at current block.
	// context.Background() creates an empty context for communicating signals across go routine
	// boundaries...and in this case, you want a fakey context, because you don't want to interrupt the
	// the API call for the purpose of the test.
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	// This balance will be in wei
	fmt.Println(balance)

	if balance.Cmp(big.NewInt(0)) == 0 {
		log.Fatal("address not found (check if validator is compiled with GENESIS_FILE=genesis_scdev_160k.go)")
	}
}

func Test_Genesis_NonExisting_Balance(t *testing.T) {
	fmt.Println("Connecting to validator")
	//
	client, err := ethclient.Dial("http://127.0.0.1:9660/ext/bc/C/rpc")
	if err != nil {
		log.Fatal(err)
	}

	// use last address from genesis_scdev_160k
	account := common.HexToAddress("xxx")

	// Pass nil in third parameter to get balance at current block.
	// context.Background() creates an empty context for communicating signals across go routine
	// boundaries...and in this case, you want a fakey context, because you don't want to interrupt the
	// the API call for the purpose of the test.
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	// This balance will be in wei
	fmt.Println(balance)

	if balance.Cmp(big.NewInt(0)) != 0 {
		log.Fatal("address should not be found")
	}
}
