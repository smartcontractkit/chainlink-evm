package clientwrappers

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
)

type chainClientMetrics struct {
	chainID                    string
	multiCallDurationHistogram metric.Float64Histogram
}

func newChainClientMetrics(chainID *big.Int) (*chainClientMetrics, error) {
	multiCallDurationHistogram, err := beholder.GetMeter().Float64Histogram("txm_multicall_duration_ms")
	if err != nil {
		return nil, fmt.Errorf("failed to register txm multicall duration metric: %w", err)
	}

	chainIDStr := "unknown"
	if chainID != nil {
		chainIDStr = chainID.String()
	}

	return &chainClientMetrics{
		chainID:                    chainIDStr,
		multiCallDurationHistogram: multiCallDurationHistogram,
	}, nil
}

func (m *chainClientMetrics) recordMultiCallDuration(ctx context.Context, method, blockTag string, duration time.Duration, err error) {
	success := err == nil
	timedOut := errors.Is(err, context.DeadlineExceeded)

	m.multiCallDurationHistogram.Record(ctx, float64(duration)/float64(time.Millisecond), metric.WithAttributes(
		attribute.String("chainID", m.chainID),
		attribute.String("method", method),
		attribute.String("blockTag", blockTag),
		attribute.Bool("success", success),
		attribute.Bool("timedOut", timedOut),
	))
}
