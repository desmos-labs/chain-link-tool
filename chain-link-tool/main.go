package main

import "github.com/desmos-labs/chain-link-tool/cmd"

func main() {
	executor := cmd.NewRootCmd()
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
