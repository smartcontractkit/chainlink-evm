package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers"
	zksyncwrapper "github.com/smartcontractkit/chainlink-evm/gethwrappers/generation/generate/zksync"
)

func main() {
	project := os.Args[1]
	contractName := os.Args[2]
	packageName := os.Args[3]

	fmt.Println("Generating", packageName, "contract wrapper")

	cwd, err := os.Getwd() // gethwrappers/zksync directory
	if err != nil {
		gethwrappers.Exit("could not get working directory", err)
	}

	srcFile := filepath.Join(cwd, "..", "..", "contracts", "zkout", contractName+".sol", contractName+".json")
	bytecode := zksyncwrapper.ReadBytecodeFromForgeJson(srcFile)

	var outPath string
	if project != "automation" {
		outPath = filepath.Join(cwd, "..", project, "generated", packageName, packageName+"_zksync.go")
	} else {
		outPath = filepath.Join(cwd, "..", "generated", packageName, packageName+"_zksync.go")
	}

	zksyncwrapper.WrapZksyncDeploy(bytecode, contractName, packageName, outPath)
}
