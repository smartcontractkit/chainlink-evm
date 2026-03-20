package clientwrappers

import (
	"encoding/json"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
)

func TestMultiplexCallSequential_ReturnsFirstSuccessfulResult(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	m.On("CallContextAllSequential", mock.Anything, "eth_getTransactionCount", "0xabc", "latest").Return(json.RawMessage(`"0x9"`), 1, nil).Once()

	decode := func(raw json.RawMessage) (uint64, error) {
		var nonce hexutil.Uint64
		if err := json.Unmarshal(raw, &nonce); err != nil {
			return 0, err
		}

		return uint64(nonce), nil
	}

	result, duration, callCount, err := multiCallSequential(t.Context(), m, "eth_getTransactionCount", []interface{}{"0xabc", "latest"}, decode)
	require.NoError(t, err)
	require.Equal(t, uint64(9), result)
	require.Equal(t, 1, callCount)
	require.GreaterOrEqual(t, duration.Nanoseconds(), int64(0))
}

func TestMultiplexCallSequential_ErrorsWhenDecodeFails(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	m.On("CallContextAllSequential", mock.Anything, "eth_getTransactionCount", "0xabc", "latest").Return(json.RawMessage(`"not-a-nonce"`), 1, nil).Once()

	decode := func(raw json.RawMessage) (uint64, error) {
		var nonce hexutil.Uint64
		if err := json.Unmarshal(raw, &nonce); err != nil {
			return 0, err
		}

		return uint64(nonce), nil
	}

	_, _, _, err := multiCallSequential(t.Context(), m, "eth_getTransactionCount", []interface{}{"0xabc", "latest"}, decode)
	require.ErrorContains(t, err, "error decoding")
}

func TestGetTransactionCountMultiplexed_ReturnsNonce(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")

	m.On("CallContextAllSequential", mock.Anything, "eth_getTransactionCount", address, "latest").Return(json.RawMessage(`"0xa"`), 1, nil).Once()

	metrics, err := newChainClientMetrics(big.NewInt(1))
	require.NoError(t, err)

	nonce, err := getTransactionCountMultiCall(t.Context(), m, logger.Sugared(logger.Test(t)), metrics, address, "latest")
	require.NoError(t, err)
	require.Equal(t, uint64(10), nonce)
}

func TestGetTransactionCountMultiplexed_ErrorsWhenNoSuccessfulResults(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	m.On("CallContextAllSequential", mock.Anything, "eth_getTransactionCount", address, "pending").Return(json.RawMessage(nil), 2, errors.New("all nodes failed for method: eth_getTransactionCount")).Once()

	metrics, err := newChainClientMetrics(big.NewInt(1))
	require.NoError(t, err)

	_, err = getTransactionCountMultiCall(t.Context(), m, logger.Sugared(logger.Test(t)), metrics, address, "pending")
	require.ErrorContains(t, err, "all nodes failed")
}
