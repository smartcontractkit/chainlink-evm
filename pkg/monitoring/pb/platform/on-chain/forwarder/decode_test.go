//nolint:govet // disable govet
package forwarder

import (
	"testing"

	"github.com/stretchr/testify/require"

	wt_msg "github.com/smartcontractkit/chainlink-evm/pkg/monitoring/pb/platform"
)

func TestDecodeAsReportProcessed(t *testing.T) {
	encoded := NewTestReport(t)

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
