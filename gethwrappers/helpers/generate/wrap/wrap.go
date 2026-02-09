package wrap

import (
	"fmt"
	"os"
	"path/filepath"

	gethwrappers "github.com/smartcontractkit/chainlink-evm/gethwrappers/helpers"
)

// GenWrapper generates a contract wrapper for the given contract.
//
// solcProjectRoot is the path to the solc artifacts for the project, e.g. `"../../contracts/solc/" + project`
// abiPath is the path to the contract's ABI JSON file.
// binPath is the path to the contract's binary file, typically with .bin extension.
// contract is the name of the generated contract class.
// pkgName is the name of the package the contract will be generated in. Try
// to follow idiomatic Go package naming conventions where possible.
//
// outDirSuffixInput is the directory suffix to generate the wrapper in. If not provided, the
// wrapper will be generated in the default location. The default location is
// <project>/generated/<pkgName>/<pkgName>.go. The suffix will take place after
// the <project>/generated, so the overridden location would be
// <project>/generated/<outDirSuffixInput>/<pkgName>/<pkgName>.go.
func GenWrapper(solcProjectRoot, contract, pkgName, outDirSuffixInput, abiGenPath string) {
	abiPath := filepath.Join(solcProjectRoot, contract, contract+".sol", contract+".abi.json")
	metadataPath := filepath.Join(solcProjectRoot, contract, contract+".sol", contract+".metadata.json")
	binPath := filepath.Join(solcProjectRoot, contract, contract+".sol", contract+".bin")
	buildInfoPath := filepath.Join(solcProjectRoot, contract, "build", "build.json")

	fmt.Println("Generating", pkgName, "contract wrapper")

	outDir := GetOutDir(outDirSuffixInput, pkgName)
	outPath := filepath.Join(outDir, pkgName+".go")
	metadataOutPath := filepath.Join(outDir, pkgName+"_metadata.go")

	gethwrappers.Abigen(gethwrappers.AbigenArgs{
		Bin:          binPath,
		ABI:          abiPath,
		BuildInfo:    buildInfoPath,
		Metadata:     metadataPath,
		Out:          outPath,
		BuildInfoOut: metadataOutPath,
		Type:         contract,
		Pkg:          pkgName,
		AbiGenPath:   abiGenPath,
	})

	// Build succeeded, so update the versions db with the new contract data
	versions, err := gethwrappers.ReadVersionsDB()
	if err != nil {
		gethwrappers.Exit("could not read current versions database", err)
	}
	versions.GethVersion = gethwrappers.GethVersion
	versions.ContractVersions[pkgName] = gethwrappers.ContractVersion{
		Hash:       gethwrappers.VersionHash(abiPath, binPath),
		AbiPath:    abiPath,
		BinaryPath: binPath,
	}
	if err := gethwrappers.WriteVersionsDB(versions); err != nil {
		gethwrappers.Exit("could not save versions db", err)
	}
}

func GetOutDir(outDirSuffixInput, pkgName string) string {
	cwd, err := os.Getwd() // gethwrappers directory
	if err != nil {
		gethwrappers.Exit("could not get working directory", err)
	}
	outDir := filepath.Join(cwd, "generated", outDirSuffixInput, pkgName)
	if mkdErr := os.MkdirAll(outDir, 0700); mkdErr != nil {
		gethwrappers.Exit(
			fmt.Sprintf("failed to create wrapper dir, outDirSuffixInput: %s (could be empty)", outDirSuffixInput),
			mkdErr)
	}

	return outDir
}
