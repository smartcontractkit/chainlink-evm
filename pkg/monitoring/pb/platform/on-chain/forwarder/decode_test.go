//nolint:govet // disable govet
package forwarder

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	wt_msg "github.com/smartcontractkit/chainlink-evm/pkg/monitoring/pb/platform"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/platform"
)

func TestDecodeAsReportProcessed(t *testing.T) {
	feedReport := datafeeds.FeedReport{
		FeedID:    [32]byte{0x01},
		Price:     big.NewInt(1234567890123456789),
		Timestamp: 1620000000,
	}

	reports := &datafeeds.Reports{
		feedReport,
	}

	data, err := datafeeds.GetSchema().Pack(reports)
	require.NoError(t, err)

	report := platform.Report{
		Metadata: datafeeds.Metadata{
			Version:             1,
			WorkflowExecutionID: [32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			Timestamp:           1620000000,
			DonID:               1,
			DonConfigVersion:    1,
			WorkflowCID:         [32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			WorkflowName:        [10]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a},
			WorkflowOwner:       [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			ReportID:            [2]byte{0x01},
		},
		Data: data,
	}

	encoded, err := report.Encode()
	require.NoError(t, err)

	// Define test cases
	tests := []struct {
		name     string
		input    wt_msg.WriteConfirmed
		expected ReportProcessed
		wantErr  bool
	}{
		{
			name: "Valid input",
			input: wt_msg.WriteConfirmed{
				Node:      "example-node",
				Forwarder: "example-forwarder",
				Receiver:  "example-receiver",

				// Report Info
				ReportId:      123,
				ReportContext: []byte{},
				Report:        encoded, // Example valid byte slice
				SignersNum:    2,

				// Transmission Info
				Transmitter: "example-transmitter",
				Success:     true,

				// Block Info
				BlockHash:      "0xaa",
				BlockHeight:    "17",
				BlockTimestamp: 0x66f5bf69,
			},
			expected: ReportProcessed{
				Receiver:            "example-receiver",
				WorkflowExecutionId: "0102030405060708090a0b0c0d0e0f1000000000000000000000000000000000",
				ReportId:            123,
				Success:             true,

				BlockHash:      "0xaa",
				BlockHeight:    "17",
				BlockTimestamp: 0x66f5bf69,

				TxSender:   "example-transmitter",
				TxReceiver: "example-forwarder",
			},
			wantErr: false,
		},
		{
			name: "Invalid input",
			input: wt_msg.WriteConfirmed{
				Node:      "example-node",
				Forwarder: "example-forwarder",
				Receiver:  "example-receiver",

				// Report Info
				ReportId:      123,
				ReportContext: []byte{},
				Report:        []byte{0x01, 0x02, 0x03, 0x04}, // Example invalid byte slice
				SignersNum:    2,

				// Transmission Info
				Transmitter: "example-transmitter",
				Success:     true,

				// Block Info
				BlockHash:      "0xaa",
				BlockHeight:    "17",
				BlockTimestamp: 0x66f5bf69,
			},
			expected: ReportProcessed{},
			wantErr:  true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DecodeAsReportProcessed(&tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, *result)
			}
		})
	}
}
