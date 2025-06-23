//nolint:govet, testifylint // disable govet, testifylint
package datafeeds_test

import (
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	ocr3types "github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
	df_processor "github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds/processor"
	por_processor "github.com/smartcontractkit/chainlink-evm/pkg/report/por/processor"
	commonpb "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/common"
	df "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/data-feeds/on-chain/registry"
	wt_msg "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/platform"

	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/report/platform"
)

func TestDecodeAsReportProcessed(t *testing.T) {
	reports := &datafeeds.Reports{
		{
			FeedID:    [32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x8, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			Price:     big.NewInt(1234567890123456789),
			Timestamp: 1620000000,
		},
		{
			FeedID:    [32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x22, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			Price:     big.NewInt(300069),
			Timestamp: 1620000001,
		},
	}

	dfProcessor := df_processor.NewDataFeedsProcessor(nil, nil)

	data, err := df_processor.GetDataFeedsSchema().Pack(reports)
	require.NoError(t, err)

	runTests(t, data, dfProcessor, false)
}

func TestPORDecodeAsReportProcessed(t *testing.T) {
	reports := &datafeeds.PORReports{
		{
			DataID:    [32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x8, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			Timestamp: 1620000000,
			Bundle:    []byte{0x01, 0x02, 0x03},
		},
		{
			DataID:    [32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x22, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			Timestamp: 1620000001,
			Bundle:    []byte{0x04, 0x94, 0x25},
		},
	}

	porProcessor := por_processor.NewPORProcessor(nil, nil)

	data, err := por_processor.GetPORSchema().Pack(reports)
	require.NoError(t, err)

	runTests(t, data, porProcessor, true)
}

func runTests(t *testing.T, data []byte, processor *datafeeds.Processor, por bool) {
	report := platform.Report{
		Metadata: ocr3types.Metadata{
			Version:          1,
			ExecutionID:      "1234567890123456789012345678901234567890123456789012345678901234",
			Timestamp:        1620000000,
			DONID:            1,
			DONConfigVersion: 1,
			WorkflowID:       "1234567890123456789012345678901234567890123456789012345678901234",
			WorkflowName:     "12",
			WorkflowOwner:    "1234567890123456789012345678901234567890",
			ReportID:         "1234",
		},
		Data: data,
	}

	rawReport, err := report.Encode()
	require.NoError(t, err)

	expected := []df.FeedUpdated{
		{
			FeedId:                "0x0102030405060708090a0b0c0d0e0f1000000000000000000000000000000000",
			ObservationsTimestamp: 1620000000,
			Benchmark:             []uint8{0x11, 0x22, 0x10, 0xf4, 0x7d, 0xe9, 0x81, 0x15},
			Bundle:                []uint8{},
			Report:                []uint8{},

			BenchmarkVal: math.NaN(),

			BlockData: &commonpb.BlockData{
				BlockHash:      "0xaa",
				BlockHeight:    "17",
				BlockTimestamp: 0x66f5bf69,
			},

			TransactionData: &commonpb.TransactionData{
				TxSender:   "example-transmitter",
				TxReceiver: "example-forwarder",
			},

			ExecutionContext: &commonpb.ExecutionContext{},
		},
		{
			FeedId:                "0x0102030405060722090a0b0c0d0e0f1000000000000000000000000000000000",
			ObservationsTimestamp: 1620000001,
			Benchmark:             []uint8{0x04, 0x94, 0x25},
			Bundle:                []uint8{},
			Report:                []uint8{},
			BenchmarkVal:          3000.69,

			BlockData: &commonpb.BlockData{
				BlockHash:      "0xaa",
				BlockHeight:    "17",
				BlockTimestamp: 0x66f5bf69,
			},

			TransactionData: &commonpb.TransactionData{
				TxSender:   "example-transmitter",
				TxReceiver: "example-forwarder",
			},

			ExecutionContext: &commonpb.ExecutionContext{},
		},
	}

	if por {
		expected[0].Bundle = []uint8{0x01, 0x02, 0x03}
		expected[1].Bundle = []uint8{0x04, 0x94, 0x25}

		expected[0].Benchmark = []uint8{}
		expected[1].Benchmark = []uint8{}

		expected[0].BenchmarkVal = math.NaN()
		expected[1].BenchmarkVal = 0
	}

	// Define test cases
	tests := []struct {
		name     string
		input    wt_msg.WriteConfirmed
		expected []df.FeedUpdated
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

				BlockData: &commonpb.BlockData{
					BlockHash:      "0xaa",
					BlockHeight:    "17",
					BlockTimestamp: 0x66f5bf69,
				},

				ExecutionContext: &commonpb.ExecutionContext{},
			},
			expected: expected,
			wantErr:  false,
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

				ExecutionContext: &commonpb.ExecutionContext{},
			},
			expected: []df.FeedUpdated{
				{ExecutionContext: &commonpb.ExecutionContext{}},
			},
			wantErr: true,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result []*df.FeedUpdated
			var err error

			result, err = processor.DecodeAsFeedUpdated(&tt.input)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, len(tt.expected), len(result))

				for i, m := range tt.expected {
					// Notice: if BenchmarkVal is NaN we can't compare directly
					if math.IsNaN(m.BenchmarkVal) {
						require.True(t, math.IsNaN(result[i].BenchmarkVal))
						// Skip the comparison (nullify the value)
						m.BenchmarkVal = -1
						result[i].BenchmarkVal = -1
					}
					// Finally, compare the values
					require.Equal(t, m, *result[i])
				}
			}
		})
	}
}

func TestToBenchmarkVal(t *testing.T) {
	// Helper function to set a big.Int value (base 10)
	mustSetString := func(s string) *big.Int {
		val, _ := new(big.Int).SetString(s, 10)
		return val
	}

	tests := []struct {
		name             string
		feedID           string
		val              *big.Int
		expected         float64
		expectedDecimals uint8
	}{
		{
			name:             "Number (price value) with 18 decimals",
			feedID:           "018e16c39e000032000000000000000000000000000000000000000000000000",
			val:              big.NewInt(1000000000000000000),
			expected:         1.0,
			expectedDecimals: 18,
		},
		{
			name:             "Number (price value) with 8 decimals",
			feedID:           "01e880c2b3000028000000000000000000000000000000000000000000000000",
			val:              big.NewInt(1000000000000000000),
			expected:         10000000000.0,
			expectedDecimals: 8,
		},
		{
			name:             "Number (price value) with 18 decimals - feed ID #2",
			feedID:           "01e880c2b3000132000000000000000000000000000000000000000000000000",
			val:              big.NewInt(1000000012340000000), // 1 ETH
			expected:         1.00000001234,
			expectedDecimals: 18,
		},
		{
			name:             "Number (24-hour global volume) as integer",
			feedID:           "01e880c2b3000820000000000000000000000000000000000000000000000000",
			val:              big.NewInt(1000000000000000000), // 1 ETH
			expected:         1000000000000000000.0,
			expectedDecimals: 0,
		},
		{
			name:             "Number (price value) with 18 decimals - feed ID #3",
			feedID:           "018933b5e4001032000000000000000000000000000000000000000000000000",
			val:              big.NewInt(1000000000000000087), // 1 ETH
			expected:         1.000000000000000087,
			expectedDecimals: 18,
		},
		{
			name:             "NaN value (NAV issuer name) as a string",
			feedID:           "018933b5e4001101000000000000000000000000000000000000000000000000",
			val:              big.NewInt(1000000000000000000),
			expected:         math.NaN(),
			expectedDecimals: 0,
		},
		{
			name:             "NaN value (NAV registry location) as an address",
			feedID:           "018933b5e4001202000000000000000000000000000000000000000000000000",
			val:              big.NewInt(1000000000000000000),
			expected:         math.NaN(),
			expectedDecimals: 0,
		},
		{
			name:             "Number (price value) with 18 decimals - feed ID #4",
			feedID:           "011e22d6bf000332000000000000000000000000000000000000000000000000",
			val:              mustSetString("9990000000000000009"),
			expected:         9.990000000000000009,
			expectedDecimals: 18,
		},
		{
			name:             "Number (price value) with 8 decimals - feed ID #2",
			feedID:           "01a80ff216000328000000000000000000000000000000000000000000000000",
			val:              mustSetString("9990000000000000009"),
			expected:         99900000000.00000009,
			expectedDecimals: 8,
		},
		{
			name:             "Number (price value) with 8 decimals - feed ID #2 - huge number (overflow)",
			feedID:           "01a80ff216000328000000000000000000000000000000000000000000000000",
			val:              new(big.Int).Exp(big.NewInt(10), big.NewInt(400), nil), // Very large number
			expected:         math.Inf(1),                                            // positive infinity
			expectedDecimals: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feedID, err := datafeeds.NewFeedIDFromHex(tt.feedID)
			require.NoError(t, err)

			decimals, isNumber := datafeeds.GetDecimals(feedID.GetDataType())

			result := datafeeds.ToBenchmarkVal(feedID, tt.val)
			if math.IsNaN(tt.expected) {
				require.False(t, isNumber)
				require.True(t, math.IsNaN(result))
			} else {
				require.True(t, isNumber)
				require.Equal(t, tt.expected, result)
				require.Equal(t, tt.expectedDecimals, decimals)
			}
		})
	}
}
