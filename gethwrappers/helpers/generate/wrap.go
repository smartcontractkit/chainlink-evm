package main

import (
	"os"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate/wrap"
)

func main() {
	solcProjectRoot := os.Args[1]
	abiGenPath := os.Args[2]
	contract := os.Args[3]
	pkgName := os.Args[4]

	outDirSuffix := "latest"

	wrap.GenWrapper(solcProjectRoot, contract, pkgName, outDirSuffix, abiGenPath)
}
