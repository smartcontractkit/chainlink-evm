package client_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
)

func TestToBackwardCompatibleCallArgWithChainTypeSupport(t *testing.T) {
	testCases := []struct {
		name      string
		msg       ethereum.CallMsg
		chainType chaintype.ChainType
		expected  map[string]interface{}
	}{
		{
			name: "Ethereum chain with data",
			msg: ethereum.CallMsg{
				From: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				To:   &common.Address{0x2},
				Data: []byte{0x1, 0x2, 0x3},
				Gas:  100000,
			},
			chainType: chaintype.ChainType(""),
			expected: map[string]interface{}{
				"from":  common.HexToAddress("0x1234567890123456789012345678901234567890"),
				"to":    &common.Address{0x2},
				"input": hexutil.Bytes{0x1, 0x2, 0x3},
				"data":  hexutil.Bytes{0x1, 0x2, 0x3},
				"gas":   hexutil.Uint64(100000),
			},
		},
		{
			name: "Tron chain with data should omit input field",
			msg: ethereum.CallMsg{
				From: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				To:   &common.Address{0x2},
				Data: []byte{0x1, 0x2, 0x3},
				Gas:  100000,
			},
			chainType: chaintype.ChainTron,
			expected: map[string]interface{}{
				"from": common.HexToAddress("0x1234567890123456789012345678901234567890"),
				"to":   &common.Address{0x2},
				"data": hexutil.Bytes{0x1, 0x2, 0x3},
				"gas":  hexutil.Uint64(100000),
			},
		},
		{
			name: "Other chain types should behave like Ethereum unless explicitly stated",
			msg: ethereum.CallMsg{
				From: common.HexToAddress("0x1234567890123456789012345678901234567890"),
				To:   &common.Address{0x2},
				Data: []byte{0x1, 0x2, 0x3},
				Gas:  100000,
			},
			chainType: chaintype.ChainArbitrum,
			expected: map[string]interface{}{
				"from":  common.HexToAddress("0x1234567890123456789012345678901234567890"),
				"to":    &common.Address{0x2},
				"input": hexutil.Bytes{0x1, 0x2, 0x3},
				"data":  hexutil.Bytes{0x1, 0x2, 0x3},
				"gas":   hexutil.Uint64(100000),
			},
		},
		{
			name: "Call with value, gasPrice and fee caps",
			msg: ethereum.CallMsg{
				From:      common.HexToAddress("0x1234567890123456789012345678901234567890"),
				To:        &common.Address{0x2},
				Data:      []byte{0x1, 0x2, 0x3},
				Gas:       100000,
				Value:     big.NewInt(1000),
				GasPrice:  big.NewInt(20),
				GasFeeCap: big.NewInt(30),
				GasTipCap: big.NewInt(5),
			},
			chainType: chaintype.ChainTron,
			expected: map[string]interface{}{
				"from":                 common.HexToAddress("0x1234567890123456789012345678901234567890"),
				"to":                   &common.Address{0x2},
				"data":                 hexutil.Bytes{0x1, 0x2, 0x3},
				"gas":                  hexutil.Uint64(100000),
				"value":                (*hexutil.Big)(big.NewInt(1000)),
				"gasPrice":             (*hexutil.Big)(big.NewInt(20)),
				"maxFeePerGas":         (*hexutil.Big)(big.NewInt(30)),
				"maxPriorityFeePerGas": (*hexutil.Big)(big.NewInt(5)),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := client.ToBackwardCompatibleCallArgWithChainTypeSupport(tc.msg, tc.chainType)

			// Convert result to map for comparison
			resultMap, ok := result.(map[string]interface{})
			assert.True(t, ok, "Result should be a map")

			// Check each expected key/value
			for k, v := range tc.expected {
				assert.Contains(t, resultMap, k, "Result should contain key: %s", k)

				// For byte slices (input and data fields), we need special handling
				if k == "input" || k == "data" {
					expectedBytes, _ := v.(hexutil.Bytes)
					resultBytes, _ := resultMap[k].(hexutil.Bytes)
					assert.Equal(t, expectedBytes, resultBytes, "Bytes for key %s should match", k)
				} else {
					assert.Equal(t, v, resultMap[k], "Value for key %s should match", k)
				}
			}

			// For Tron, ensure 'input' is not present
			if tc.chainType == chaintype.ChainTron {
				_, hasInput := resultMap["input"]
				assert.False(t, hasInput, "Tron chain should not have 'input' field")
			}
		})
	}
}
