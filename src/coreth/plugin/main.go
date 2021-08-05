// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"fmt"
	"os"

	"github.com/ava-labs/coreth/core"

	"github.com/ethereum/go-ethereum/log"
	"github.com/hashicorp/go-plugin"

	"github.com/ava-labs/avalanchego/vms/rpcchainvm"

	"github.com/ava-labs/coreth/plugin/evm"

	"google.golang.org/grpc"
)

// evi1m3: Expanded buffer factory method for GRPC server
func ExpandedBufferGRPCServer(opts []grpc.ServerOption) *grpc.Server {
	// set max receive message size to 20MB
	size := 20 * 1024 * 1024
	opts = append(opts, grpc.MaxRecvMsgSize(size))
	core.GlobalServer = grpc.NewServer(opts...)
	return core.GlobalServer
}

func main() {
	version, err := PrintVersion()
	if err != nil {
		fmt.Printf("couldn't get config: %s", err)
		os.Exit(1)
	}
	if version {
		fmt.Println(evm.Version)
		os.Exit(0)
	}
	// Set the Ethereum logger to debug by default
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, log.StreamHandler(os.Stderr, log.TerminalFormat(false))))
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: rpcchainvm.Handshake,
		Plugins: map[string]plugin.Plugin{
			"vm": rpcchainvm.New(&evm.VM{}),
		},

		// A non-nil value here enables gRPC serving for this plugin...
		// evi1m3: use new initialization function
		GRPCServer: ExpandedBufferGRPCServer,
	})
}
