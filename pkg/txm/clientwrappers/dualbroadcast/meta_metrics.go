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
	latencyHistogram  metric.Int64Histogram
	bidHistogram      metric.Int64Histogram
	errorCounter      metric.Int64Counter
}

// NewMetaMetrics creates a new MetaMetrics instance
func NewMetaMetrics(chainID string) (*MetaMetrics, error) {
	statusCodeCounter, err := beholder.GetMeter().Int64Counter("meta_endpoint_status_codes")
	if err != nil {
		return nil, err
	}

	latencyHistogram, err := beholder.GetMeter().Int64Histogram("meta_endpoint_latency")
	if err != nil {
		return nil, err
	}

	bidHistogram, err := beholder.GetMeter().Int64Histogram("meta_bids_per_transaction")
	if err != nil {
		return nil, err
	}

	errorCounter, err := beholder.GetMeter().Int64Counter("meta_errors")
	if err != nil {
		return nil, err
	}

	return &MetaMetrics{
		chainID:           chainID,
		statusCodeCounter: statusCodeCounter,
		latencyHistogram:  latencyHistogram,
		bidHistogram:      bidHistogram,
		errorCounter:      errorCounter,
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
	m.latencyHistogram.Record(ctx, duration.Milliseconds(),
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
		),
	)
}

// RecordBidsReceived records the distribution of bids per transaction
func (m *MetaMetrics) RecordBidsReceived(ctx context.Context, bidCount int) {
	m.bidHistogram.Record(ctx, int64(bidCount),
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
		),
	)
}

// RecordSendRequestError records errors from SendRequest method
func (m *MetaMetrics) RecordSendRequestError(ctx context.Context) {
	m.errorCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("errorType", "send_request"),
		),
	)
}

// RecordSendOperationError records errors from SendOperation method
func (m *MetaMetrics) RecordSendOperationError(ctx context.Context) {
	m.errorCounter.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("errorType", "send_operation"),
		),
	)
}
