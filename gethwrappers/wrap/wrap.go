package main

import (
	"fmt"
	"log"
	"os"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers/generate/wrap"
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

	projectRoot := "../../contracts/solc/" + project
	wrap.GenWrapper(projectRoot, contract, pkgName, outDirSuffix, abiGenPath)
}
