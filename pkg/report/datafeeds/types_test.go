package datafeeds_test

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
)

func TestDecodeFeedReport(t *testing.T) {
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
	schema := datafeeds.GetSchema()

	// Pack the original data using the ABI schema.
	encoded, err := schema.Pack(original)
	if err != nil {
		t.Fatalf("failed to pack data: %v", err)
	}

	// Decode the data using our Decode function.
	decoded, err := datafeeds.Decode(encoded)
	if err != nil {
		t.Fatalf("failed to decode data: %v", err)
	}

	// Check that the lengths match.
	if len(*decoded) != len(original) {
		t.Fatalf("expected %d records, got %d", len(original), len(*decoded))
	}

	// Compare each record field by field.
	for i := range original {
		origRecord := original[i]
		decRecord := (*decoded)[i]

		// Compare FeedID.
		if origRecord.FeedID != decRecord.FeedID {
			t.Errorf("record %d: mismatched FeedID: expected %x, got %x", i, origRecord.FeedID, decRecord.FeedID)
		}

		// Compare Price using big.Int.Cmp.
		if origRecord.Price.Cmp(decRecord.Price) != 0 {
			t.Errorf("record %d: mismatched Price: expected %v, got %v", i, origRecord.Price, decRecord.Price)
		}

		// Compare Timestamp.
		if origRecord.Timestamp != decRecord.Timestamp {
			t.Errorf("record %d: mismatched Timestamp: expected %d, got %d", i, origRecord.Timestamp, decRecord.Timestamp)
		}
	}
}
