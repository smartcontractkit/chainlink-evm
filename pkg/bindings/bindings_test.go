package bindings_test

import (
	"log"
	"testing"

	"github.com/smartcontractkit/chainlink-evm/pkg/bindings"
)

func TestGenBindings(t *testing.T) {
	if err := bindings.GenerateBindings(
		"build/DataStorage_combined.json",
		"",
		"bindings",
		"",
		"build/bindings.go",
	); err != nil {
		log.Fatalf("failed to generate: %v", err)
	}
}
