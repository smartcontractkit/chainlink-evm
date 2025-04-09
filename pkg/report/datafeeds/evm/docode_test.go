// evm_feedupdated_test.go
package evm_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	// Import the evm package under test.
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds/evm"

	// These packages are used by the DecodeAsFeedUpdated function.
	wt_msg "github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform"
	// platform.Decode is used within DecodeAsFeedUpdated.
)

func TestDecodeAsFeedUpdated(t *testing.T) {
	feedReports := evm.Reports{
		{
			FeedID:    [32]byte{0x01},
			Price:     big.NewInt(1234567890), // example price value
			Timestamp: 1000,                   // example timestamp
		},
	}

	// ABI‑encode the feed reports using the global schema.
	feedReportsEncoded, err := evm.GetSchema().Pack(feedReports)
	require.NoError(t, err)

	// Create a Report metadata with dummy values.
	metadata := types.Metadata{
		Version:          1,
		ExecutionID:      "2",
		Timestamp:        2000,
		DONID:            0,
		DONConfigVersion: 0,
		WorkflowID:       "3",
		WorkflowName:     "4",
		WorkflowOwner:    "5",
		ReportID:         "6",
	}

	// Build the full Report with the metadata and the already ABI‑encoded feed reports.
	report := evm.Report{
		Metadata: metadata,
		Data:     feedReportsEncoded,
	}

	// Use the Report.Encode method to produce the final encoded report bytes.
	encodedReport, err := report.Encode()
	require.NoError(t, err)

	// Build a valid WriteConfirmed message.
	validMsg := wt_msg.WriteConfirmed{
		Node:           "test-node",
		Forwarder:      "test-forwarder",
		Receiver:       "test-receiver",
		ReportId:       123,
		ReportContext:  []byte{},
		Report:         encodedReport, // encoded report (metadata || ABI-encoded feed reports)
		SignersNum:     1,
		Transmitter:    "test-transmitter",
		Success:        true,
		BlockHash:      "0xabc",
		BlockHeight:    "100",
		BlockTimestamp: 3000,
	}

	// Decode the WriteConfirmed message.
	feedUpdates, err := evm.DecodeAsFeedUpdated(&validMsg)
	require.NoError(t, err)

	require.Len(t, feedUpdates, len(feedReports))

	first := feedUpdates[0]
	expectedFeedID := "0x" + hex.EncodeToString(feedReports[0].FeedID[:])
	require.Equal(t, expectedFeedID, first.FeedId)

	require.Equal(t, feedReports[0].Timestamp, first.ObservationsTimestamp)

	priceFromFeedUpdated := new(big.Int).SetBytes(first.Benchmark)
	require.Equal(t, 0, feedReports[0].Price.Cmp(priceFromFeedUpdated))

	require.Equal(t, feedReportsEncoded, first.Report)

	invalidMsg := wt_msg.WriteConfirmed{
		Report: []byte{0x01, 0x02, 0x03, 0x04},
	}
	_, err = evm.DecodeAsFeedUpdated(&invalidMsg)
	require.Error(t, err)
}
