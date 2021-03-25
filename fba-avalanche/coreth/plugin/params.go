package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ava-labs/coreth/plugin/evm"
)

var (
	cliConfig evm.CommandLineConfig
)

func init() {
	fs := flag.NewFlagSet("coreth", flag.ContinueOnError)

	config := fs.String("config", "default", "Pass in CLI Config to set runtime attributes for Coreth")

	if err := fs.Parse(os.Args[1:]); err != nil {
		cliConfig.ParsingError = err
		return
	}

	// if *config == "default" {
	// 	cliConfig.EthAPIEnabled = true
	// 	cliConfig.PersonalAPIEnabled = true
	// 	cliConfig.TxPoolAPIEnabled = true
	// 	cliConfig.NetAPIEnabled = true
	// 	cliConfig.Web3APIEnabled = true
	// 	cliConfig.RPCGasCap = 2500000000 // 25000000 x 100
	// 	cliConfig.RPCTxFeeCap = 100      // 100 AVAX
	// } else {
	// 	// TODO only overwrite values that were explicitly set
	// 	cliConfig.ParsingError = json.Unmarshal([]byte(*config), &cliConfig)
	// }

	cliConfig.EthAPIEnabled = true
	cliConfig.PersonalAPIEnabled = true
	cliConfig.TxPoolAPIEnabled = true
	cliConfig.NetAPIEnabled = true
	cliConfig.Web3APIEnabled = true
	cliConfig.DebugAPIEnabled = true
	cliConfig.RPCGasCap = 2500000000 // 25000000 x 100
	cliConfig.RPCTxFeeCap = 100      // 100 AVAX

	if *config != "default" {
		for _, value := range strings.Split(*config, " ") {
			if value != "" {
				cliConfig.StateConnectorConfig = append(cliConfig.StateConnectorConfig, value)
			} else {
				cliConfig.ParsingError = fmt.Errorf("StateConnectorConfig contains empty string")
				return
			}
		}
	}

}
