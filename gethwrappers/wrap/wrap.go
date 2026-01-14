package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate/wrap"
	zksyncwrapper "github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/zksync"
)

// main is the entry point for the wrapper generation tool. The abiGenPath is static
// and assumes this is being called from two levels below the project root. This should
// be true for all modern wrapper generation, which is split into project folders.
func main() {
	project := os.Args[1]
	contract := os.Args[2]
	pkgName := os.Args[3]

	var outDirSuffix string
	if len(os.Args) >= 5 {
		outDirSuffix = os.Args[4] + "/latest"
	} else {
		outDirSuffix = "latest"
	}

	abiGenPath := "../../tools/bin/abigen"

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	if os.Getenv("ZKSYNC") == "true" {
		zksyncBytecodePath := filepath.Join("..", "zkout", contract+".sol", contract+".json")
		zksyncBytecode := zksyncwrapper.ReadBytecodeFromForgeJSON(zksyncBytecodePath)
		outPath := filepath.Join(wrap.GetOutDir(outDirSuffix, pkgName), pkgName+"_zksync.go")
		zksyncwrapper.WrapZksyncDeploy(zksyncBytecode, contract, pkgName, outPath)
	} else {
		projectRoot := "../../contracts/solc/" + project
		wrap.GenWrapper(projectRoot, contract, pkgName, outDirSuffix, abiGenPath)
	}
}
