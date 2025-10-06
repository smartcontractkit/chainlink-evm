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
	chainID          *big.Int
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
	ChainID         string `json:"chainId"`
	TxID            uint64 `json:"txId"`
	RequestTime     int64  `json:"requestTime"`     // Unix milliseconds
	
	// Request Info
	RequestPayload  string `json:"requestPayload,omitempty"` // JSON string of request params
	
	// Response Info  
	StatusCode      int    `json:"statusCode"`
	LatencyMs       int64  `json:"latencyMs"`
	ErrorMessage    string `json:"errorMessage,omitempty"`
	
	// Auction Results
	BidCount        int    `json:"bidCount"`
	MaxBidAmount    string `json:"maxBidAmount,omitempty"`    // String to avoid precision loss
	SecondBidAmount string `json:"secondBidAmount,omitempty"`
	BidSpreadPct    float64 `json:"bidSpreadPct"`             // Second best as % of max
	BidToken        string `json:"bidToken,omitempty"`
	WinningSolver   string `json:"winningSolver,omitempty"`
	
	// Solver Details (for deeper analysis)
	Solvers         []SolverInfo `json:"solvers,omitempty"`
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
	
	// Analyze bids for summary data
	var maxBid, secondBid string
	var bidSpreadPct float64
	var bidToken, winningSolver string
	var solvers []SolverInfo
	
	if len(solverOps) > 0 {
		// Build solver details for analysis
		for _, so := range solverOps {
			if so.BidAmount != nil && so.BidAmount.ToInt().Cmp(big.NewInt(0)) > 0 {
				solvers = append(solvers, SolverInfo{
					Address:   so.Solver.Hex(),
					BidAmount: so.BidAmount.String(),
					BidToken:  so.BidToken.Hex(),
				})
			}
		}
		
		if len(solvers) > 0 {
			// First solver is winner, get their details
			maxBid = solvers[0].BidAmount
			bidToken = solvers[0].BidToken
			winningSolver = solvers[0].Address
			
			// Second best if available
			if len(solvers) > 1 {
				secondBid = solvers[1].BidAmount
				
				// Calculate spread percentage
				maxBigInt := new(big.Int)
				secondBigInt := new(big.Int)
				maxBigInt.SetString(maxBid, 10)
				secondBigInt.SetString(secondBid, 10)
				
				if maxBigInt.Cmp(big.NewInt(0)) > 0 && secondBigInt.Cmp(big.NewInt(0)) > 0 {
					maxFloat, _ := maxBigInt.Float64()
					secondFloat, _ := secondBigInt.Float64()
					if maxFloat > 0 {
						bidSpreadPct = (secondFloat / maxFloat) * 100
					}
				}
			}
		}
	}
	
	// Create event structure
	event := MetaRequestEvent{
		ChainID:         m.chainID.String(),
		TxID:            tx.ID,
		RequestTime:     time.Now().UnixMilli(),
		RequestPayload:  string(requestPayload),
		StatusCode:      statusCode,
		LatencyMs:       latency.Milliseconds(),
		ErrorMessage:    errorMsg,
		BidCount:        len(solvers),
		MaxBidAmount:    maxBid,
		SecondBidAmount: secondBid,
		BidSpreadPct:    bidSpreadPct,
		BidToken:        bidToken,
		WinningSolver:   winningSolver,
		Solvers:         solvers,
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
