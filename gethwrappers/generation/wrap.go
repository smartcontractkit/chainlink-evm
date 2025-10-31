package main

import (
	"os"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/generation/generate/genwrapper"
)

var (
	rootDir = "../../contracts/solc/"
)

func main() {
	project := os.Args[1]
	className := os.Args[2]
	pkgName := os.Args[3]

	var outDirSuffix string
	if len(os.Args) >= 5 {
		outDirSuffix = os.Args[4]
	}

	// Once vrf is moved to its own subfolder we can delete this rootDir override.
	if project == "vrf" || project == "automation" {
		rootDir = "../contracts/solc/"
	}

	// If the project starts with "workflow" we need to adjust the path
	if len(project) >= 8 && project[0:8] == "workflow" || project == "keystone" {
		rootDir = "../../solc/"
	}

	abiPath := rootDir + project + "/" + className + "/" + className + ".sol/" + className + ".abi.json"
	binPath := rootDir + project + "/" + className + "/" + className + ".sol/" + className + ".bin"

	genwrapper.GenWrapper(abiPath, binPath, className, pkgName, outDirSuffix)
}
