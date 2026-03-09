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
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
)

func TestMultiplexCallBest_SelectsBestAndReturnsAllSuccessful(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	m.On("CallContextAll", mock.Anything, "eth_getTransactionCount", "0xabc", "latest").Return([]client.CallContextAllResult{
		{NodeName: "a", Result: []byte(`"0x2"`)},
		{NodeName: "b", Err: errors.New("rpc failed")},
		{NodeName: "c", Result: []byte(`"0x9"`)},
		{NodeName: "d", Result: []byte(`"not-a-nonce"`)},
	}, nil).Once()

	decode := func(raw json.RawMessage) (uint64, error) {
		var nonce hexutil.Uint64
		if err := json.Unmarshal(raw, &nonce); err != nil {
			return 0, err
		}

		return uint64(nonce), nil
	}

	best, all, duration, err := multiplexCallBest(context.Background(), m, "eth_getTransactionCount", []interface{}{"0xabc", "latest"}, decode, func(candidate, current uint64) bool {
		return candidate > current
	})
	require.NoError(t, err)
	require.Equal(t, uint64(9), best)
	require.Equal(t, []uint64{2, 9}, all)
	require.GreaterOrEqual(t, duration.Nanoseconds(), int64(0))
}

func TestMultiplexCallBest_ErrorsWhenNoSuccessfulDecodedResults(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	m.On("CallContextAll", mock.Anything, "eth_getTransactionCount", "0xabc", "latest").Return([]client.CallContextAllResult{
		{NodeName: "a", Err: errors.New("rpc failed")},
		{NodeName: "b", Result: []byte(`"not-a-nonce"`)},
	}, nil).Once()

	decode := func(raw json.RawMessage) (uint64, error) {
		var nonce hexutil.Uint64
		if err := json.Unmarshal(raw, &nonce); err != nil {
			return 0, err
		}

		return uint64(nonce), nil
	}

	_, _, _, err := multiplexCallBest(context.Background(), m, "eth_getTransactionCount", []interface{}{"0xabc", "latest"}, decode, func(candidate, current uint64) bool {
		return candidate > current
	})
	require.ErrorContains(t, err, "no successful results")
}

func TestGetTransactionCountMultiplexed_ReturnsHighestNonce(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")

	m.On("CallContextAll", mock.Anything, "eth_getTransactionCount", address, "latest").Return([]client.CallContextAllResult{
		{NodeName: "a", Result: []byte(`"0x1"`)},
		{NodeName: "b", Result: []byte(`"0xa"`)},
		{NodeName: "c", Err: errors.New("rpc failed")},
	}, nil).Once()

	nonce, err := GetTransactionCountMultiplexed(context.Background(), m, logger.Sugared(logger.Test(t)), address, "latest")
	require.NoError(t, err)
	require.Equal(t, uint64(10), nonce)
}

func TestGetTransactionCountMultiplexed_ErrorsWhenNoSuccessfulResults(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClient(t)
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")

	m.On("CallContextAll", mock.Anything, "eth_getTransactionCount", address, "pending").Return([]client.CallContextAllResult{
		{NodeName: "a", Err: errors.New("rpc failed")},
		{NodeName: "b", Result: []byte(`"bad"`)},
	}, nil).Once()

	_, err := GetTransactionCountMultiplexed(context.Background(), m, logger.Sugared(logger.Test(t)), address, "pending")
	require.ErrorContains(t, err, "no successful results")
}
