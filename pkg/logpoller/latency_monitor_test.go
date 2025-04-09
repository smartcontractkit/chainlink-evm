package logpoller_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/logpoller"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

type mockClient struct {
	latency time.Duration
}

func (c *mockClient) HeadByNumber(_ context.Context, _ *big.Int) (*types.Head, error) {
	time.Sleep(c.latency)
	return nil, nil
}

func (c *mockClient) HeadByHash(_ context.Context, _ common.Hash) (*types.Head, error) {
	time.Sleep(c.latency)
	return nil, nil
}

func (c *mockClient) FilterLogs(_ context.Context, q ethereum.FilterQuery) ([]ethTypes.Log, error) {
	time.Sleep(c.latency)
	return nil, nil
}

type mockTracker struct {
	latency time.Duration
}

func (t *mockTracker) LatestAndFinalizedBlock(_ context.Context) (*types.Head, *types.Head, error) {
	time.Sleep(t.latency)
	return nil, nil, nil
}

func TestLatencyMonitor(t *testing.T) {
	lggr, logs := logger.TestObserved(t, zapcore.DebugLevel)
	blockProductionRate := 250 * time.Millisecond
	client := &mockClient{}
	tracker := &mockTracker{}

	lm := logpoller.NewLatencyMonitor(client, tracker, lggr, blockProductionRate)

	// Slow client with latency 80% of block production rate
	client.latency = time.Duration(0.8 * float64(blockProductionRate))
	tracker.latency = time.Duration(0.8 * float64(blockProductionRate))
	_, _ = lm.HeadByNumber(t.Context(), nil)
	_, _ = lm.HeadByHash(t.Context(), common.Hash{})
	_, _, _ = lm.LatestAndFinalizedBlock(t.Context())

	// Should not track latency on block range
	filter := ethereum.FilterQuery{FromBlock: big.NewInt(123), ToBlock: big.NewInt(456), BlockHash: nil}
	_, _ = lm.FilterLogs(t.Context(), filter)
	require.Equal(t, 3, logs.Len())

	// Should track latency for a specific block hash call
	filter = ethereum.FilterQuery{FromBlock: nil, ToBlock: nil, BlockHash: &common.Hash{}}
	_, _ = lm.FilterLogs(t.Context(), filter)
	require.Equal(t, 4, logs.Len())
	_ = logs.TakeAll()

	// fast client should not log warnings
	client.latency = 0
	tracker.latency = 0
	_, _ = lm.HeadByNumber(t.Context(), nil)
	_, _ = lm.HeadByHash(t.Context(), common.Hash{})
	_, _, _ = lm.LatestAndFinalizedBlock(t.Context())
	require.Equal(t, 0, logs.Len())
}
