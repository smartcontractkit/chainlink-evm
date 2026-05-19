package logpoller

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
)

var (
	promLpLastProcessedBlock = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "evm_log_poller_last_processed_block",
		Help: "The last block that the log poller has processed. Main purpose is to signal if the log poller is stuck and not processing new blocks. May be reported with a delay.",
	}, []string{"chainFamily", "chainID"})
)

type Metrics interface {
	RecordLastProcessedBlock(ctx context.Context, lastProcessedBlock int64)
}

var _ Metrics = (*PromBeholderMetrics)(nil)

type PromBeholderMetrics struct {
	chainID            string
	chainFamily        string
	lastProcessedBlock metric.Int64Gauge
}

func NewPromBeholderMetrics(chainID string, chainFamily string) (*PromBeholderMetrics, error) {
	lastProcessedBlock, err := beholder.GetMeter().Int64Gauge("evm_log_poller_last_processed_block")
	if err != nil {
		return nil, fmt.Errorf("failed to register last processed block metric: %w", err)
	}

	return &PromBeholderMetrics{
		chainID:            chainID,
		chainFamily:        chainFamily,
		lastProcessedBlock: lastProcessedBlock,
	}, nil
}

func (m *PromBeholderMetrics) RecordLastProcessedBlock(ctx context.Context, lastProcessedBlock int64) {
	promLpLastProcessedBlock.WithLabelValues(m.chainFamily, m.chainID).Set(float64(lastProcessedBlock))
	m.lastProcessedBlock.Record(ctx, lastProcessedBlock, metric.WithAttributes(
		attribute.String("chainFamily", m.chainFamily),
		attribute.String("chainID", m.chainID)))
}

var _ Metrics = (*noopMetrics)(nil)

var NoopMetrics = &noopMetrics{}

type noopMetrics struct{}

func (m *noopMetrics) RecordLastProcessedBlock(_ context.Context, _ int64) {}
