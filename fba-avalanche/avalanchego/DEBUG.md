# Debugging HOWTO

## VSCode

### Prerequisites
- VSCode
- Go 1.15.5 (use gvm for ease of Go version setting). You can grab the latest version of Go and install it first. Then install gvm. Then use `gvm use go1.15.5` to set version. You may also need to point your shell to find gvm after installing by calling `source ~/.gvm/scripts/gvm`.
- [VSCode Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
  - Follow plugin install instructions
- [Delve Debugger](https://github.com/go-delve/delve)
  - This should automatically be installed
- [Delve native DAP implementation](https://github.com/golang/vscode-go/blob/master/docs/dlv-dap.md#getting-started)
  - The setting to auto-load the plugin is already included in workspace settings

### Running
- Open ./main/main.go
- Set breakpoints
- Launch `Debug AvalancheGo` debug configuration from debug tool

### Caveats
- Note that Avalanche Go starts the C-Chain EVM plugin automatically. The EVM does not start under the debugger. If you want to debug the EVM, you will need to alter how the plugin starts, so that it can be debugged. Then, you'd attach the VSCode debugger for remote debugging. There would also be a way to start the EVM plugin directly for debugging under VSCode with a launch configuration, but this has not be done at this time. This might be as simple as attaching to the evm process from another vscode debugger session, but it would need a configuration to know where the go source code for the evm was.
- Note that Ctrl-C handler from the debug console does not forward to the Avalanche Go process. You can terminate the debugging session but note that the EVM plugin does not terminate. You will need to kill it manually.