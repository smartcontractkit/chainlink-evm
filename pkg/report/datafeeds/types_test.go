package datafeeds_test

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
	"github.com/stretchr/testify/require"
)

func TestDecodeFeedReport(t *testing.T) {
	testCases := []struct {
		name string
		ccip bool
	}{
		{
			name: "Normal DF",
			ccip: false,
		},
		{
			name: "CCIP DF",
			ccip: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create some sample records.
			original := []datafeeds.FeedReport{
				{
					// Example feedID: the first byte is 0x01 and the remainder are zeros.
					FeedID:    [32]byte{0x01},
					Price:     big.NewInt(1234567890123456789),
					Timestamp: 1620000000,
				},
				{
					FeedID:    [32]byte{0xAA, 0xBB, 0xCC},
					Price:     big.NewInt(123),
					Timestamp: 1630000000,
				},
			}

			// Get the ABI schema from our constructor.
			schema := datafeeds.GetSchema(true)

			// Pack the original data using the ABI schema.
			encoded, err := schema.Pack(original)
			require.NoError(t, err)

			// Decode the data using our Decode function.
			decoded, err := datafeeds.Decode(encoded, true)
			require.NoError(t, err)

			require.Equal(t, len(original), len(*decoded))

			// Compare each record field by field.
			for i := range original {
				origRecord := original[i]
				decRecord := (*decoded)[i]

				require.Equal(t, origRecord.FeedID, decRecord.FeedID)
				require.Equal(t, origRecord.Price, decRecord.Price)
				require.Equal(t, origRecord.Timestamp, decRecord.Timestamp)
			}
		})
	}
}
