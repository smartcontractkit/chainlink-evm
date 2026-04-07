package txm

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/metrics"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
	svrv1 "github.com/smartcontractkit/chainlink-protos/svr/v1"
)

var (
	promNumBroadcastedTxs = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "txm_num_broadcasted_transactions",
		Help: "Total number of successful broadcasted transactions.",
	}, []string{"chainID"})
	promNumConfirmedTxs = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "txm_num_confirmed_transactions",
		Help: "Total number of confirmed transactions. Note that this can happen multiple times per transaction in the case of re-orgs or when filling the nonce for untracked transactions.",
	}, []string{"chainID"})
	promNumNonceGaps = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "txm_num_nonce_gaps",
		Help: "Total number of nonce gaps created that the transaction manager had to fill.",
	}, []string{"chainID"})
	promReachedMaxAttempts = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "txm_reached_max_attempts",
		Help: "A gauge that is treated as boolean; 1 if the condition is true, 0 otherwise. Controls whether the TXM has reached max attempts threshold or not.",
	}, []string{"chainID"})
	promTimeUntilTxConfirmed = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "txm_time_until_tx_confirmed",
		Help: "The amount of time elapsed from a transaction being broadcast to being included in a block.",
	}, []string{"chainID"})
)

// TxmMetrics is the single metrics type for the TXMv2 transaction lifecycle.
// It records both general TXM operational metrics and lifecycle metrics.
type TxmMetrics struct {
	metrics.Labeler
	chainID              *big.Int
	numBroadcastedTxs    metric.Int64Counter
	numConfirmedTxs      metric.Int64Counter
	numNonceGaps         metric.Int64Counter
	reachedMaxAttempts   metric.Int64Gauge
	timeUntilTxConfirmed metric.Float64Histogram
	lifecycleFailure     metric.Int64Counter
}

func NewTxmMetrics(chainID *big.Int) (*TxmMetrics, error) {
	numBroadcastedTxs, err := beholder.GetMeter().Int64Counter("txm_num_broadcasted_transactions")
	if err != nil {
		return nil, fmt.Errorf("failed to register broadcasted txs number: %w", err)
	}

	numConfirmedTxs, err := beholder.GetMeter().Int64Counter("txm_num_confirmed_transactions")
	if err != nil {
		return nil, fmt.Errorf("failed to register confirmed txs number: %w", err)
	}

	numNonceGaps, err := beholder.GetMeter().Int64Counter("txm_num_nonce_gaps")
	if err != nil {
		return nil, fmt.Errorf("failed to register nonce gaps number: %w", err)
	}

	timeUntilTxConfirmed, err := beholder.GetMeter().Float64Histogram("txm_time_until_tx_confirmed")
	if err != nil {
		return nil, fmt.Errorf("failed to register time until tx confirmed: %w", err)
	}

	reachedMaxAttempts, err := beholder.GetMeter().Int64Gauge("txm_reached_max_attempts")
	if err != nil {
		return nil, fmt.Errorf("failed to register max attempts indicator: %w", err)
	}

	lifecycleFailure, err := beholder.GetMeter().Int64Counter("txm_transaction_lifecycle_failure_total")
	if err != nil {
		return nil, fmt.Errorf("failed to register lifecycle failure counter: %w", err)
	}

	return &TxmMetrics{
		chainID:              chainID,
		Labeler:              metrics.NewLabeler().With("chainID", chainID.String()),
		numBroadcastedTxs:    numBroadcastedTxs,
		numConfirmedTxs:      numConfirmedTxs,
		numNonceGaps:         numNonceGaps,
		reachedMaxAttempts:   reachedMaxAttempts,
		timeUntilTxConfirmed: timeUntilTxConfirmed,
		lifecycleFailure:     lifecycleFailure,
	}, nil
}

// LifecycleFailureStage represents a stage in the transaction lifecycle where a failure can occur.
type LifecycleFailureStage string

const (
	StageCreate         LifecycleFailureStage = "create"
	StageInFlightSubset LifecycleFailureStage = "in_flight_subset"
	StageMaxInFlight    LifecycleFailureStage = "max_in_flight"
	StageBroadcast      LifecycleFailureStage = "broadcast"
	StageNonceAt        LifecycleFailureStage = "nonce_at"

	// SVR-specific stages.
	StageCreatePrimary LifecycleFailureStage = "create_primary"
	StageAuction       LifecycleFailureStage = "auction"
)

// IncrementLifecycleFailure increments the lifecycle failure counter for the given stage.
func (m *TxmMetrics) IncrementLifecycleFailure(ctx context.Context, stage LifecycleFailureStage) {
	m.lifecycleFailure.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("chainID", m.chainID.String()),
			attribute.String("stage", string(stage)),
		),
	)
}

func (m *TxmMetrics) IncrementNumBroadcastedTxs(ctx context.Context) {
	promNumBroadcastedTxs.WithLabelValues(m.chainID.String()).Add(float64(1))
	m.numBroadcastedTxs.Add(ctx, 1)
}

func (m *TxmMetrics) IncrementNumConfirmedTxs(ctx context.Context, confirmedTransactions int) {
	promNumConfirmedTxs.WithLabelValues(m.chainID.String()).Add(float64(confirmedTransactions))
	m.numConfirmedTxs.Add(ctx, int64(confirmedTransactions))
}

func (m *TxmMetrics) IncrementNumNonceGaps(ctx context.Context) {
	promNumNonceGaps.WithLabelValues(m.chainID.String()).Add(float64(1))
	m.numNonceGaps.Add(ctx, 1)
}

func (m *TxmMetrics) ReachedMaxAttempts(ctx context.Context, reached bool) {
	var value float64
	if reached {
		value = 1
	}
	promReachedMaxAttempts.WithLabelValues(m.chainID.String()).Set(value)
	m.reachedMaxAttempts.Record(ctx, int64(value))
}

func (m *TxmMetrics) RecordTimeUntilTxConfirmed(ctx context.Context, duration float64) {
	promTimeUntilTxConfirmed.WithLabelValues(m.chainID.String()).Observe(duration)
	m.timeUntilTxConfirmed.Record(ctx, duration)
}

func (m *TxmMetrics) EmitTxMessage(ctx context.Context, txHash common.Hash, fromAddress common.Address, tx *types.Transaction) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return fmt.Errorf("failed to get meta for tx %s: %w", txHash, err)
	}

	var destAddress string
	if meta != nil && meta.FwdrDestAddress != nil {
		destAddress = meta.FwdrDestAddress.String()
	}

	toAddress := common.Address{}
	if !tx.IsPurgeable {
		toAddress = tx.ToAddress
	}

	message := &svrv1.TxMessage{
		Hash:        txHash.String(),
		FromAddress: fromAddress.String(),
		ToAddress:   toAddress.String(),
		Nonce:       strconv.FormatUint(*tx.Nonce, 10),
		CreatedAt:   time.Now().UnixMicro(),
		ChainId:     m.chainID.String(),
		FeedAddress: destAddress,
	}

	messageBytes, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	return beholder.GetEmitter().Emit(
		ctx,
		messageBytes,
		"beholder_domain", "svr",
		"beholder_entity", "svr.v1.TxMessage",
		"beholder_data_schema", "/beholder-tx-message/versions/2",
	)
}
