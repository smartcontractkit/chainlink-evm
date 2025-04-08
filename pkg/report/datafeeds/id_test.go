package datafeeds

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDecimals(t *testing.T) {
	tests := []struct {
		name     string
		dataType uint8
		expected uint8
		isNumber bool
	}{
		{
			name:     "Decimal0 (Integer)",
			dataType: 0x20,
			expected: 0,
			isNumber: true,
		},
		{
			name:     "Decimal8",
			dataType: 0x28,
			expected: 8,
			isNumber: true,
		},
		{
			name:     "Decimal18",
			dataType: 0x32,
			expected: 18,
			isNumber: true,
		},
		{
			name:     "Decimal64",
			dataType: 0x60,
			expected: 64,
			isNumber: true,
		},
		{
			name:     "Non-numeric data type",
			dataType: 0x01,
			expected: 0,
			isNumber: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decimals, isNumber := GetDecimals(tt.dataType)
			require.Equal(t, tt.expected, decimals)
			require.Equal(t, tt.isNumber, isNumber)
		})
	}
}

// Examples: [DF2.0 | Data ID Final Specification](https://docs.google.com/document/d/13ciwTx8lSUfyz1IdETwpxlIVSn1lwYzGtzOBBTpl5Vg/edit?usp=sharing)
// 0x01 8e16c39e 0000 32 000000000000000000000000000000000000000000000000 = ETH/USD Benchmark Price with 18 decimals
// 0x01 e880c2b3 0000 28 000000000000000000000000000000000000000000000000 = BTC/USD Benchmark Price with 8 decimals
// 0x01 e880c2b3 0001 32 000000000000000000000000000000000000000000000000 = BTC/USD Best Bid Price with 18 decimals
// 0x01 e880c2b3 0008 20 000000000000000000000000000000000000000000000000 = BTC/USD 24-hour global volume as integer
// 0x01 8933b5e4 0010 32 000000000000000000000000000000000000000000000000 = ARK BTC NAV value with 18 decimals
// 0x01 8933b5e4 0011 01 000000000000000000000000000000000000000000000000 = ARK BTC NAV issuer name as a string
// 0x01 8933b5e4 0012 02 000000000000000000000000000000000000000000000000 = ARK BTC NAV registry location as an address
// 0x01 1e22d6bf 0003 32 000000000000000000000000000000000000000000000000 = X/Y Price with 18 decimals
// 0x01 a80ff216 0003 28 000000000000000000000000000000000000000000000000 = X/Y Price with 8 decimals
func TestFeedIDGetDataTypeDecimals(t *testing.T) {
	tests := []struct {
		name     string
		feedID   string
		expected uint8
		isNumber bool
	}{
		{
			name:     "ETH/USD Benchmark Price with 18 decimals",
			feedID:   "018e16c39e000032000000000000000000000000000000000000000000000000",
			expected: 0x32,
			isNumber: true,
		},
		{
			name:     "BTC/USD Benchmark Price with 8 decimals",
			feedID:   "01e880c2b3000028000000000000000000000000000000000000000000000000",
			expected: 0x28,
			isNumber: true,
		},
		{
			name:     "BTC/USD Best Bid Price with 18 decimals",
			feedID:   "01e880c2b3000132000000000000000000000000000000000000000000000000",
			expected: 0x32,
			isNumber: true,
		},
		{
			name:     "BTC/USD 24-hour global volume as integer",
			feedID:   "01e880c2b3000820000000000000000000000000000000000000000000000000",
			expected: 0x20,
			isNumber: true,
		},
		{
			name:     "ARK BTC NAV value with 18 decimals",
			feedID:   "018933b5e4001032000000000000000000000000000000000000000000000000",
			expected: 0x32,
			isNumber: true,
		},
		{
			name:     "ARK BTC NAV issuer name as a string",
			feedID:   "018933b5e4001101000000000000000000000000000000000000000000000000",
			expected: 0x01,
			isNumber: false,
		},
		{
			name:     "ARK BTC NAV registry location as an address",
			feedID:   "018933b5e4001202000000000000000000000000000000000000000000000000",
			expected: 0x02,
			isNumber: false,
		},
		{
			name:     "X/Y Price with 18 decimals",
			feedID:   "011e22d6bf000332000000000000000000000000000000000000000000000000",
			expected: 0x32,
			isNumber: true,
		},
		{
			name:     "X/Y Price with 8 decimals",
			feedID:   "01a80ff216000328000000000000000000000000000000000000000000000000",
			expected: 0x28,
			isNumber: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			feedID, err := NewFeedIDFromHex(tt.feedID)
			require.NoError(t, err)

			dataType := feedID.GetDataType()
			require.Equal(t, tt.expected, dataType)

			decimals, isNumber := GetDecimals(dataType)
			require.Equal(t, tt.isNumber, isNumber)

			if isNumber {
				require.Equal(t, tt.expected-0x20, decimals)
			} else {
				require.Equal(t, uint8(0), decimals)
			}
		})
	}
}
