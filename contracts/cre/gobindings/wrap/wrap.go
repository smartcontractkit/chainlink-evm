package main

import (
	"os"
	"path/filepath"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate/wrap"
	zksyncwrapper "github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/zksync"
)

func main() {
	project := os.Args[1]
	contract := os.Args[2]
	pkgName := os.Args[3]

	var outDirSuffix string
	if len(os.Args) >= 5 {
		outDirSuffix = os.Args[4]
	} else {
		outDirSuffix = "latest"
	}

	abiGenPath := "../../../../tools/bin/abigen"

	if os.Getenv("ZKSYNC") == "true" {
		zksyncBytecodePath := filepath.Join("..", "zkout", contract+".sol", contract+".json")
		zksyncBytecode := zksyncwrapper.ReadBytecodeFromForgeJSON(zksyncBytecodePath)
		outPath := filepath.Join(wrap.GetOutDir(outDirSuffix, pkgName), pkgName+"_zksync.go")
		zksyncwrapper.WrapZksyncDeploy(zksyncBytecode, contract, pkgName, outPath)
	} else {
		projectRoot := "../../solc/" + project
		wrap.GenWrapper(projectRoot, contract, pkgName, outDirSuffix, abiGenPath)
	}
}
