// evm_feedupdated_test.go
package evm_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds/evm"

	wt_msg "github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform"
)

func TestDecodeAsFeedUpdated(t *testing.T) {
	feedReports := evm.Reports{
		{
			FeedID:    [32]byte{0x01},
			Price:     big.NewInt(1234567890),
			Timestamp: 1000,
		},
	}

	// ABI‑encode the feed reports using the global schema.
	feedReportsEncoded, err := evm.GetSchema().Pack(feedReports)
	require.NoError(t, err)

	metadata := datafeeds.Metadata{
		Version:             1,
		WorkflowExecutionID: [32]byte{0x02},
		Timestamp:           2000,
		DonID:               0,
		DonConfigVersion:    0,
		WorkflowCID:         [32]byte{0x03},
		WorkflowName:        [10]byte{0x04},
		WorkflowOwner:       [20]byte{0x05},
		ReportID:            [2]byte{0x06, 0x07},
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
