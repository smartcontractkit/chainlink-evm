package common

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeReport(t *testing.T) {
	// Hex-encoded Mercury report data (example)
	// Test case sourced from: contracts/data-feeds/sources/registry.move#L661
	encoded := "0003fbba4fce42f65d6032b18aee53efdf526cc734ad296cb57565979d883bdd0000000000000000000000000000000000000000000000000000000066ed173e0000000000000000000000000000000000000000000000000000000066ed174200000000000000007fffffffffffffffffffffffffffffffffffffffffffffff00000000000000007fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000066ee68c2000000000000000000000000000000000000000000000d808cc35e6ed670bd00000000000000000000000000000000000000000000000d808590c35425347980000000000000000000000000000000000000000000000d8093f5f989878e7c00"

	// Decode the hex data
	decoded, err := hex.DecodeString(encoded)
	require.NoError(t, err)

	m, err := Decode(decoded)
	require.NoError(t, err)

	require.Equal(t, uint16(3), GetReportType(m.FeedID))
}
