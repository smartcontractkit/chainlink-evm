package dualbroadcast

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
)

// OFAMetrics provides unified OFA (Order Flow Auction) metrics for all dual-broadcast
// backends. The "backend" label differentiates between providers (e.g. "flashbots", "nova").
type OFAMetrics struct { // TODO(gg): unexport?
	chainID       string
	backend       string
	sendTxStatus  metric.Int64Counter
	sendTxLatency metric.Int64Histogram
}

func NewOFAMetrics(chainID, backend string) (*OFAMetrics, error) {
	sendTxStatus, err := beholder.GetMeter().Int64Counter("ofa_send_tx_status")
	if err != nil {
		return nil, err
	}

	sendTxLatency, err := beholder.GetMeter().Int64Histogram("ofa_send_tx_latency",
		metric.WithUnit("ms"),
		metric.WithDescription("Latency of OFA send transaction requests"),
		metric.WithExplicitBucketBoundaries(100, 250, 500, 1000, 2000, 3000, 5000, 7500, 10000), //	TODO(gg): double-check buckets, could we use the default instead?
	)
	if err != nil {
		return nil, err
	}

	return &OFAMetrics{
		chainID:       chainID,
		backend:       backend,
		sendTxStatus:  sendTxStatus,
		sendTxLatency: sendTxLatency,
	}, nil
}

func (m *OFAMetrics) RecordSendTx(ctx context.Context, duration time.Duration, err error) {
	status := "success"
	if err != nil {
		status = "error"
	}
	attrs := metric.WithAttributes(
		attribute.String("chainID", m.chainID),
		attribute.String("backend", m.backend),
		attribute.String("status", status),
	)
	m.sendTxStatus.Add(ctx, 1, attrs)

	latencyAttrs := metric.WithAttributes( // TODO(gg): have separate success/error latency histograms
		attribute.String("chainID", m.chainID),
		attribute.String("backend", m.backend),
	)
	m.sendTxLatency.Record(ctx, duration.Milliseconds(), latencyAttrs)
}
