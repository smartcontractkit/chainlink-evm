//nolint:govet // disable govet
package forwarder

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"

	wt_msg "github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform"
)

func TestDecodeAsReportProcessed(t *testing.T) {
	// Base64-encoded report data (example)
	// version | workflow_execution_id | timestamp | don_id | config_version | ... | data
	encoded := "AYFtgPpLuLNQysw6LjlSNrzGuBOwVoth7qC9PmunIY3TZvW/cAAAAAEAAAABvAbzAOeX1ahXVjehSq4T4/hQgAjR/FT0xGEf/xemjLAwMDAwRk9PQkFSAAAAAAAAAAAAAAAAAAAAAAAAAKoAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHAAAMREREREREREQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEgAAMREREREREREQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAZvW/aQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABm9b9pAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAElCUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASUJQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABnBQGpAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAElCUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASUJQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABJQlAAMiIiIiIiIiIgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEgAAMiIiIiIiIiIgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAZvW/aQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABm9b9pAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAElCUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASUJQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABnBQGpAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAElCUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAASUJQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABJQl"

	// Decode the base64 data
	rawReport, err := base64.StdEncoding.DecodeString(encoded)
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
				Report:        rawReport, // Example valid byte slice
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
				WorkflowExecutionId: "816d80fa4bb8b350cacc3a2e395236bcc6b813b0568b61eea0bd3e6ba7218dd3",
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
