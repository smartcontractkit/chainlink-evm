package dualbroadcast

import (
	"context"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
)

// MetaMetrics handles all Meta-related metrics via OTEL
type MetaMetrics struct {
	chainID           string
	statusCodeCounter metric.Int64Counter
	latencyHistogram  metric.Float64Histogram
}

// NewMetaMetrics creates a new MetaMetrics instance
func NewMetaMetrics(chainID string) (*MetaMetrics, error) {
	statusCodeCounter, err := beholder.GetMeter().Int64Counter("meta_endpoint_status_codes")
	if err != nil {
		return nil, err
	}

	latencyHistogram, err := beholder.GetMeter().Float64Histogram("meta_endpoint_latency")
	if err != nil {
		return nil, err
	}

	return &MetaMetrics{
		chainID:           chainID,
		statusCodeCounter: statusCodeCounter,
		latencyHistogram:  latencyHistogram,
	}, nil
}

// RecordStatusCode records the HTTP status code from Meta endpoint
func (m *MetaMetrics) RecordStatusCode(ctx context.Context, statusCode int) {
	m.statusCodeCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("statusCode", strconv.Itoa(statusCode)),
		),
	)
}

// RecordLatency records the latency of Meta endpoint requests
func (m *MetaMetrics) RecordLatency(ctx context.Context, duration time.Duration) {
	m.latencyHistogram.Record(ctx, float64(duration.Milliseconds()),
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
		),
	)
}
