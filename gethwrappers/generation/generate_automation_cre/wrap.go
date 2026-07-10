package main

import (
	"os"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/generation/generate/genwrapper"
)

func main() {
	rootDir := "../contracts/solc/"
	project := "automation-cre"
	inputClassName := os.Args[1]
	outputClassName := os.Args[2]
	pkgName := os.Args[3]

	abiPath := rootDir + project + "/" + inputClassName + "/" + inputClassName + ".sol/" + inputClassName + ".abi.json"
	binPath := rootDir + project + "/" + inputClassName + "/" + inputClassName + ".sol/" + inputClassName + ".bin"

	genwrapper.GenWrapper(abiPath, binPath, outputClassName, pkgName, "")
}
