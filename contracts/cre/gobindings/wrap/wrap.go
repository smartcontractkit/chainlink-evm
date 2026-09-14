package main

import (
	"os"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate/wrap"
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

	projectRoot := "../../solc/" + project
	wrap.GenWrapper(projectRoot, contract, pkgName, outDirSuffix, abiGenPath)
}
