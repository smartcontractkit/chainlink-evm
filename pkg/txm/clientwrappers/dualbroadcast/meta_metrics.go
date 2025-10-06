package dualbroadcast

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/metrics"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

// MetaMetrics handles all Meta-related metrics via OTEL
type MetaMetrics struct {
	metrics.Labeler
	chainID           *big.Int
	statusCodeCounter metric.Int64Counter
	latencyHistogram  metric.Float64Histogram
}

// NewMetaMetrics creates a new MetaMetrics instance
func NewMetaMetrics(chainID *big.Int) (*MetaMetrics, error) {
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
		Labeler:           metrics.NewLabeler().With("chainID", chainID.String()),
		statusCodeCounter: statusCodeCounter,
		latencyHistogram:  latencyHistogram,
	}, nil
}

// RecordStatusCode records the HTTP status code from Meta endpoint
func (m *MetaMetrics) RecordStatusCode(ctx context.Context, statusCode int) {
	m.statusCodeCounter.Add(ctx, 1)
}

// RecordLatency records the latency of Meta endpoint requests
func (m *MetaMetrics) RecordLatency(ctx context.Context, duration time.Duration) {
	m.latencyHistogram.Record(ctx, float64(duration.Milliseconds()))
}

// MetaRequestEvent represents a complete Meta endpoint request for analysis
type MetaRequestEvent struct {
	// Request Identification
	ChainID     string `json:"chainId"`
	TxID        uint64 `json:"txId"`
	RequestTime int64  `json:"requestTime"` // Unix milliseconds

	// Request Info
	RequestPayload string `json:"requestPayload,omitempty"` // JSON string of request params

	// Response Info
	StatusCode   int    `json:"statusCode"`
	LatencyMs    int64  `json:"latencyMs"`
	ErrorMessage string `json:"errorMessage,omitempty"`

	// Solvers Info to be analyzed
	Solvers []SolverInfo `json:"solvers,omitempty"`
}

type SolverInfo struct {
	Address   string `json:"address"`
	BidAmount string `json:"bidAmount"`
	BidToken  string `json:"bidToken"`
}

// EmitMetaRequestEvent sends a structured event to beholder for analysis
func (m *MetaMetrics) EmitMetaRequestEvent(ctx context.Context,
	tx *types.Transaction,
	attempt *types.Attempt,
	requestPayload []byte,
	statusCode int,
	latency time.Duration,
	errorMsg string,
	solverOps []*SO) error {

	// Include ALL important solver data as returned by Meta endpoint
	var solvers []SolverInfo
	for _, so := range solverOps {
		bidAmount := "0"
		if so.BidAmount != nil {
			bidAmount = so.BidAmount.String()
		}
		solvers = append(solvers, SolverInfo{
			Address:   so.Solver.Hex(),
			BidAmount: bidAmount,
			BidToken:  so.BidToken.Hex(),
		})
	}

	// Create event with raw data only
	event := MetaRequestEvent{
		ChainID:        m.chainID.String(),
		TxID:           tx.ID,
		RequestTime:    time.Now().UnixMilli(),
		RequestPayload: string(requestPayload),
		StatusCode:     statusCode,
		LatencyMs:      latency.Milliseconds(),
		ErrorMessage:   errorMsg,
		Solvers:        solvers,
	}

	// Marshal to JSON for beholder
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	// Emit to beholder with metadata
	return beholder.GetEmitter().Emit(
		ctx,
		eventBytes,
		"beholder_domain", "svr",
		"beholder_entity", "svr.v1.RequestEvent",
		"beholder_data_schema", "/beholder-request-event/versions/1",
	)
}
