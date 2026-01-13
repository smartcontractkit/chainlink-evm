package main

import (
	"os"
	"path/filepath"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate/wrap"
	zksyncwrapper "github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/zksync"
)

func main() {
	solcProjectRoot := os.Args[1]
	abiGenPath := os.Args[2]
	contract := os.Args[3]
	pkgName := os.Args[4]

	outDirSuffix := "latest"

	if os.Getenv("ZKSYNC") == "true" {
		zksyncBytecodePath := filepath.Join("..", "zkout", contract+".sol", contract+".json")
		zksyncBytecode := zksyncwrapper.ReadBytecodeFromForgeJSON(zksyncBytecodePath)
		outPath := filepath.Join(wrap.GetOutDir(outDirSuffix, pkgName), pkgName+"_zksync.go")
		zksyncwrapper.WrapZksyncDeploy(zksyncBytecode, contract, pkgName, outPath)
	} else {
		wrap.GenWrapper(solcProjectRoot, contract, pkgName, outDirSuffix, abiGenPath)
	}
}
