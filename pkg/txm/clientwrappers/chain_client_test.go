package clientwrappers

import (
	"context"
	"encoding/json"
	"errors"
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
	m.On("CallContextAllSequential", mock.Anything, "eth_getTransactionCount", "0xabc", "latest").Return(json.RawMessage(`"0x9"`), nil).Once()

	decode := func(raw json.RawMessage) (uint64, error) {
		var nonce hexutil.Uint64
		if err := json.Unmarshal(raw, &nonce); err != nil {
			return 0, err
		}

		return uint64(nonce), nil
	}

	result, duration, err := multiCallSequential(context.Background(), m, "eth_getTransactionCount", []interface{}{"0xabc", "latest"}, decode)
	require.NoError(t, err)
	require.Equal(t, uint64(9), result)
	require.GreaterOrEqual(t, duration.Nanoseconds(), int64(0))
}

func TestMultiplexCallSequential_ErrorsWhenDecodeFails(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	m.On("CallContextAllSequential", mock.Anything, "eth_getTransactionCount", "0xabc", "latest").Return(json.RawMessage(`"not-a-nonce"`), nil).Once()

	decode := func(raw json.RawMessage) (uint64, error) {
		var nonce hexutil.Uint64
		if err := json.Unmarshal(raw, &nonce); err != nil {
			return 0, err
		}

		return uint64(nonce), nil
	}

	_, _, err := multiCallSequential(context.Background(), m, "eth_getTransactionCount", []interface{}{"0xabc", "latest"}, decode)
	require.ErrorContains(t, err, "error decoding")
}

func TestGetTransactionCountMultiplexed_ReturnsNonce(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")

	m.On("CallContextAllSequential", mock.Anything, "eth_getTransactionCount", address, "latest").Return(json.RawMessage(`"0xa"`), nil).Once()

	nonce, err := getTransactionCountMultiCall(context.Background(), m, logger.Sugared(logger.Test(t)), address, "latest")
	require.NoError(t, err)
	require.Equal(t, uint64(10), nonce)
}

func TestGetTransactionCountMultiplexed_ErrorsWhenNoSuccessfulResults(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	m.On("CallContextAllSequential", mock.Anything, "eth_getTransactionCount", address, "pending").Return(json.RawMessage(nil), errors.New("all nodes failed for method: eth_getTransactionCount")).Once()

	_, err := getTransactionCountMultiCall(context.Background(), m, logger.Sugared(logger.Test(t)), address, "pending")
	require.ErrorContains(t, err, "all nodes failed")
}
