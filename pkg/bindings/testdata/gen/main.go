package main

import (
	"github.com/smartcontractkit/chainlink-evm/pkg/bindings"
)

func main() {
	if err := bindings.GenerateBindings(
		"./testdata/DataStorage_combined.json",
		"",
		"bindings",
		"",
		"./testdata/bindings.go",
	); err != nil {
		panic(err)
	}

	if err := bindings.GenerateBindings(
		"./testdata/EmptyContract_combined.json",
		"",
		"bindings",
		"",
		"./testdata/emptybindings.go",
	); err != nil {
		panic(err)
	}
}
