package por_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/por"
)

func TestDecodePORReport(t *testing.T) {
	// Create some sample records.
	original := []por.Report{
		{
			// Example feedID: the first byte is 0x01 and the remainder are zeros.
			DataId:    [32]byte{0x01},
			Timestamp: 1620000000,
			Bundle:    []byte{0x01, 0x02, 0x03},
		},
		{
			DataId:    [32]byte{0xAA, 0xBB, 0xCC},
			Timestamp: 1630000000,
			Bundle:    []byte{0x02, 0x03, 0x04},
		},
	}

	// Get the ABI schema from our constructor.
	schema := por.GetSchema()

	// Pack the original data using the ABI schema.
	encoded, err := schema.Pack(original)
	require.NoError(t, err)

	// Decode the data using our Decode function.
	decoded, err := por.Decode(encoded)
	require.NoError(t, err)

	require.Equal(t, len(original), len(*decoded))

	// Compare each record field by field.
	for i := range original {
		origRecord := original[i]
		decRecord := (*decoded)[i]

		require.Equal(t, origRecord.DataId, decRecord.DataId)
		require.Equal(t, origRecord.Timestamp, decRecord.Timestamp)
		require.Equal(t, origRecord.Bundle, decRecord.Bundle)
	}
}
