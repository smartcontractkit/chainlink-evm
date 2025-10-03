package dualbroadcast

import (
	"context"
	"encoding/json"
	"math/big"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/metrics"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

var (
	// Prometheus metrics for Meta endpoint
	promMetaStatusCode = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "meta_endpoint_status_codes_total",
		Help: "Total number of Meta endpoint requests by status code",
	}, []string{"chainID", "status_code"})

	promMetaLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "meta_endpoint_latency_seconds",
		Help:    "Latency of Meta endpoint requests in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"chainID"})

	promMetaBidsReceived = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "meta_bids_received_total",
		Help: "Total number of bids received from Meta endpoint",
	}, []string{"chainID"})

	promMetaEventsProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "meta_events_processed_total",
		Help: "Total number of Meta events processed",
	}, []string{"chainID"})

	promMetaMaxBidAmount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meta_max_bid_amount",
		Help: "Maximum bid amount received in the last Meta auction",
	}, []string{"chainID", "bid_token"})

	promMetaSecondBestBidAmount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meta_second_best_bid_amount",
		Help: "Second best bid amount received in the last Meta auction",
	}, []string{"chainID", "bid_token"})

	promMetaSecondBestBidPercentage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "meta_second_best_bid_percentage",
		Help: "Second best bid as percentage of max bid in the last Meta auction",
	}, []string{"chainID", "bid_token"})
)

// MetaMetrics handles all Meta-related metrics
type MetaMetrics struct {
	metrics.Labeler
	chainID                   *big.Int
	statusCodeCounter         metric.Int64Counter
	latencyHistogram          metric.Float64Histogram
	bidsReceivedCounter       metric.Int64Counter
	eventsProcessedCounter    metric.Int64Counter
	maxBidAmountGauge         metric.Float64Gauge
	secondBestBidAmountGauge  metric.Float64Gauge
	secondBestBidPercentGauge metric.Float64Gauge
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

	bidsReceivedCounter, err := beholder.GetMeter().Int64Counter("meta_bids_received")
	if err != nil {
		return nil, err
	}

	eventsProcessedCounter, err := beholder.GetMeter().Int64Counter("meta_events_processed")
	if err != nil {
		return nil, err
	}

	maxBidAmountGauge, err := beholder.GetMeter().Float64Gauge("meta_max_bid_amount")
	if err != nil {
		return nil, err
	}

	secondBestBidAmountGauge, err := beholder.GetMeter().Float64Gauge("meta_second_best_bid_amount")
	if err != nil {
		return nil, err
	}

	secondBestBidPercentGauge, err := beholder.GetMeter().Float64Gauge("meta_second_best_bid_percentage")
	if err != nil {
		return nil, err
	}

	return &MetaMetrics{
		chainID:                   chainID,
		Labeler:                   metrics.NewLabeler().With("chainID", chainID.String()),
		statusCodeCounter:         statusCodeCounter,
		latencyHistogram:          latencyHistogram,
		bidsReceivedCounter:       bidsReceivedCounter,
		eventsProcessedCounter:    eventsProcessedCounter,
		maxBidAmountGauge:         maxBidAmountGauge,
		secondBestBidAmountGauge:  secondBestBidAmountGauge,
		secondBestBidPercentGauge: secondBestBidPercentGauge,
	}, nil
}

// RecordStatusCode records the HTTP status code from Meta endpoint
func (m *MetaMetrics) RecordStatusCode(ctx context.Context, statusCode int) {
	statusStr := strconv.Itoa(statusCode)
	promMetaStatusCode.WithLabelValues(m.chainID.String(), statusStr).Inc()
	m.statusCodeCounter.Add(ctx, 1)
}

// RecordLatency records the latency of Meta endpoint requests
func (m *MetaMetrics) RecordLatency(ctx context.Context, duration time.Duration) {
	durationSeconds := duration.Seconds()
	promMetaLatency.WithLabelValues(m.chainID.String()).Observe(durationSeconds)
	m.latencyHistogram.Record(ctx, durationSeconds)
}

// RecordBidsReceived records the number of bids received
func (m *MetaMetrics) RecordBidsReceived(ctx context.Context, count int) {
	promMetaBidsReceived.WithLabelValues(m.chainID.String()).Add(float64(count))
	m.bidsReceivedCounter.Add(ctx, int64(count))
}

// RecordEventProcessed records when a Meta event is processed
func (m *MetaMetrics) RecordEventProcessed(ctx context.Context) {
	promMetaEventsProcessed.WithLabelValues(m.chainID.String()).Inc()
	m.eventsProcessedCounter.Add(ctx, 1)
}

// MetaRequestEvent represents a complete Meta endpoint request for analysis
type MetaRequestEvent struct {
	// Request Identification
	ChainID         string `json:"chainId"`
	TxID            uint64 `json:"txId"`
	RequestTime     int64  `json:"requestTime"`     // Unix microseconds
	
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
		RequestTime:     time.Now().UnixMicro(),
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

func (m *MetaMetrics) RecordBidAnalysis(ctx context.Context, solverOps []*SO) {
	// Analyze bids inline
	if len(solverOps) == 0 {
		return
	}

	// Collect all valid bids
	var bids []*big.Int
	var bidToken string
	
	for _, so := range solverOps {
		if so.BidAmount != nil && so.BidAmount.ToInt().Cmp(big.NewInt(0)) > 0 {
			bids = append(bids, so.BidAmount.ToInt())
			if bidToken == "" {
				bidToken = so.BidToken.Hex()
			}
		}
	}
	
	if len(bids) == 0 {
		return
	}

	// First bid is max, second bid (if exists) is second best
	maxBid := bids[0]
	var secondBestBid *big.Int
	if len(bids) > 1 {
		secondBestBid = bids[1]
	} else {
		secondBestBid = big.NewInt(0)
	}

	// Calculate percentage
	var secondBestBidPct float64
	if maxBid.Cmp(big.NewInt(0)) > 0 && secondBestBid.Cmp(big.NewInt(0)) > 0 {
		maxFloat, _ := maxBid.Float64()
		secondFloat, _ := secondBestBid.Float64()
		if maxFloat > 0 {
			secondBestBidPct = (secondFloat / maxFloat) * 100
		}
	}

	// Record metrics
	maxFloat, _ := maxBid.Float64()
	promMetaMaxBidAmount.WithLabelValues(m.chainID.String(), bidToken).Set(maxFloat)
	m.maxBidAmountGauge.Record(ctx, maxFloat)

	secondFloat, _ := secondBestBid.Float64()
	promMetaSecondBestBidAmount.WithLabelValues(m.chainID.String(), bidToken).Set(secondFloat)
	m.secondBestBidAmountGauge.Record(ctx, secondFloat)

	promMetaSecondBestBidPercentage.WithLabelValues(m.chainID.String(), bidToken).Set(secondBestBidPct)
	m.secondBestBidPercentGauge.Record(ctx, secondBestBidPct)
}
