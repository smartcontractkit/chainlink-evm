package v3

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

// Helper function to create a new big.Int from a string
func mustBigNewString(s string, base int) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(s, base)
	return n
}

func TestDecodeReport(t *testing.T) {
	// Hex-encoded Mercury report data (example)
	// Test case sourced from: contracts/data-feeds/sources/registry.move#L661
	encoded := "0003fbba4fce42f65d6032b18aee53efdf526cc734ad296cb57565979d883bdd0000000000000000000000000000000000000000000000000000000066ed173e0000000000000000000000000000000000000000000000000000000066ed174200000000000000007fffffffffffffffffffffffffffffffffffffffffffffff00000000000000007fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000000000000066ee68c2000000000000000000000000000000000000000000000d808cc35e6ed670bd00000000000000000000000000000000000000000000000d808590c35425347980000000000000000000000000000000000000000000000d8093f5f989878e7c00"

	// Decode the hex data
	decoded, err := hex.DecodeString(encoded)
	require.NoError(t, err)

	expectedData := Report{
		FeedID:                [32]uint8{0x0, 0x3, 0xfb, 0xba, 0x4f, 0xce, 0x42, 0xf6, 0x5d, 0x60, 0x32, 0xb1, 0x8a, 0xee, 0x53, 0xef, 0xdf, 0x52, 0x6c, 0xc7, 0x34, 0xad, 0x29, 0x6c, 0xb5, 0x75, 0x65, 0x97, 0x9d, 0x88, 0x3b, 0xdd},
		ObservationsTimestamp: 0x66ed1742,
		BenchmarkPrice:        mustBigNewString("000d808cc35e6ed670bd00", 16),
		Bid:                   mustBigNewString("63761571925910070000000", 10),
		Ask:                   mustBigNewString("63762609220802160000000", 10),
		ValidFromTimestamp:    0x66ed173e,
		ExpiresAt:             0x66ee68c2,
		LinkFee:               mustBigNewString("3138550867693340381917894711603833208051177722232017256447", 10),
		NativeFee:             mustBigNewString("3138550867693340381917894711603833208051177722232017256447", 10),
	}

	m, err := Decode(decoded)
	require.NoError(t, err)
	require.Equal(t, expectedData, *m)
}
