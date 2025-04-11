package logpoller

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

const (
	latencyWarningThreshold = 0.7 // Warn if latency exceeds 70% of block production rate
	latencyWarning          = "RPC latency warning: consider using faster endpoints or reviewing network conditions"
)

type LatencyMonitorClient interface {
	HeadByNumber(ctx context.Context, n *big.Int) (*types.Head, error)
	HeadByHash(ctx context.Context, n common.Hash) (*types.Head, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]ethTypes.Log, error)
}

// LatencyMonitor wraps RPC calls with a warning if the latency exceeds the set threshold of block production rate
type LatencyMonitor struct {
	c                   LatencyMonitorClient
	lggr                logger.Logger
	blockProductionRate time.Duration
}

func NewLatencyMonitor(c LatencyMonitorClient, lggr logger.Logger, blockProductionRate time.Duration) LatencyMonitor {
	return LatencyMonitor{
		c:                   c,
		lggr:                lggr,
		blockProductionRate: blockProductionRate,
	}
}

// latencyMonitoredCall wraps any function and logs a warning if it exceeds the threshold of block production rate
func latencyMonitoredCall[T any](lm *LatencyMonitor, name string, fn func() (T, error)) (T, error) {
	start := time.Now()
	result, err := fn()
	latency := time.Since(start)

	threshold := time.Duration(float64(lm.blockProductionRate) * latencyWarningThreshold)
	if latency > threshold {
		lm.lggr.Warnf(
			"%s - %s latency of %s exceeded threshold of %s (%.0f%% of block production time %s)",
			latencyWarning, name, latency, threshold, latencyWarningThreshold*100, lm.blockProductionRate,
		)
	}

	return result, err
}

func (lm *LatencyMonitor) HeadByNumber(ctx context.Context, n *big.Int) (*types.Head, error) {
	return latencyMonitoredCall(lm, "HeadByNumber", func() (*types.Head, error) {
		return lm.c.HeadByNumber(ctx, n)
	})
}

func (lm *LatencyMonitor) HeadByHash(ctx context.Context, n common.Hash) (*types.Head, error) {
	return latencyMonitoredCall(lm, "HeadByHash", func() (*types.Head, error) {
		return lm.c.HeadByHash(ctx, n)
	})
}

func (lm *LatencyMonitor) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]ethTypes.Log, error) {
	// Track latency only for single block hash
	if q.FromBlock == nil && q.ToBlock == nil && q.BlockHash != nil {
		return latencyMonitoredCall(lm, "FilterLogs", func() ([]ethTypes.Log, error) {
			return lm.c.FilterLogs(ctx, q)
		})
	}
	return lm.c.FilterLogs(ctx, q)
}
