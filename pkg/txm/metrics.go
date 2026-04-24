package txm

import (
	"context"
	"errors"
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
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
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
	promRPCNonce = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "txm_rpc_nonce",
		Help: "The latest confirmed nonce reported by the RPC node for a given address.",
	}, []string{"chainID", "address"})
)

// Metrics is the metrics contract for the TXMv2 transaction lifecycle.
type Metrics interface {
	IncrementLifecycleFailure(context.Context, LifecycleFailureStage)
	IncrementNumBroadcastedTxs(context.Context)
	IncrementNumConfirmedTxs(context.Context, int)
	IncrementNumNonceGaps(context.Context)
	ReachedMaxAttempts(context.Context, bool)
	RecordTimeUntilTxConfirmed(context.Context, float64)
	SetRPCNonce(context.Context, common.Address, uint64)
	EmitTxMessage(context.Context, common.Hash, common.Address, *types.Transaction) error
}

// txmMetrics is the default metrics recorder for the TXMv2 transaction lifecycle.
type txmMetrics struct {
	metrics.Labeler
	chainID              *big.Int
	numBroadcastedTxs    metric.Int64Counter
	numConfirmedTxs      metric.Int64Counter
	numNonceGaps         metric.Int64Counter
	reachedMaxAttempts   metric.Int64Gauge
	timeUntilTxConfirmed metric.Float64Histogram
	rpcNonce             metric.Int64Gauge
	lifecycleFailure     metric.Int64Counter
}

func NewTxmMetrics(lggr logger.Logger, chainID *big.Int) Metrics {
	var initErr error
	numBroadcastedTxs, err := beholder.GetMeter().Int64Counter("txm_num_broadcasted_transactions")
	if err != nil {
		initErr = errors.Join(initErr, fmt.Errorf("txm_num_broadcasted_transactions: %w", err))
	}

	numConfirmedTxs, err := beholder.GetMeter().Int64Counter("txm_num_confirmed_transactions")
	if err != nil {
		initErr = errors.Join(initErr, fmt.Errorf("txm_num_confirmed_transactions: %w", err))
	}

	numNonceGaps, err := beholder.GetMeter().Int64Counter("txm_num_nonce_gaps")
	if err != nil {
		initErr = errors.Join(initErr, fmt.Errorf("txm_num_nonce_gaps: %w", err))
	}

	timeUntilTxConfirmed, err := beholder.GetMeter().Float64Histogram("txm_time_until_tx_confirmed")
	if err != nil {
		initErr = errors.Join(initErr, fmt.Errorf("txm_time_until_tx_confirmed: %w", err))
	}

	reachedMaxAttempts, err := beholder.GetMeter().Int64Gauge("txm_reached_max_attempts")
	if err != nil {
		initErr = errors.Join(initErr, fmt.Errorf("txm_reached_max_attempts: %w", err))
	}

	rpcNonce, err := beholder.GetMeter().Int64Gauge("txm_rpc_nonce")
	if err != nil {
		initErr = errors.Join(initErr, fmt.Errorf("txm_rpc_nonce: %w", err))
	}

	lifecycleFailure, err := beholder.GetMeter().Int64Counter("txm_transaction_lifecycle_failure_total")
	if err != nil {
		initErr = errors.Join(initErr, fmt.Errorf("txm_transaction_lifecycle_failure_total: %w", err))
	}

	// It is very unlikely that another metric will be initialized correctly if there is even a single failure,
	// so it's safer if we use a NOOP metric struct for everything and log an error instead.
	if initErr != nil {
		lggr.Errorw("Failed to initialize TXM metrics; using noop metrics", "err", initErr)
		return NewNoopTxmMetrics()
	}

	return &txmMetrics{
		chainID:              chainID,
		Labeler:              metrics.NewLabeler().With("chainID", chainID.String()),
		numBroadcastedTxs:    numBroadcastedTxs,
		numConfirmedTxs:      numConfirmedTxs,
		numNonceGaps:         numNonceGaps,
		reachedMaxAttempts:   reachedMaxAttempts,
		timeUntilTxConfirmed: timeUntilTxConfirmed,
		rpcNonce:             rpcNonce,
		lifecycleFailure:     lifecycleFailure,
	}
}

func NewNoopTxmMetrics() Metrics {
	return noopTxmMetrics{}
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
func (m *txmMetrics) IncrementLifecycleFailure(ctx context.Context, stage LifecycleFailureStage) {
	m.lifecycleFailure.Add(ctx, 1,
		metric.WithAttributes(
			attribute.String("chainID", m.chainID.String()),
			attribute.String("stage", string(stage)),
		),
	)
}

func (m *txmMetrics) IncrementNumBroadcastedTxs(ctx context.Context) {
	promNumBroadcastedTxs.WithLabelValues(m.chainID.String()).Add(float64(1))
	m.numBroadcastedTxs.Add(ctx, 1)
}

func (m *txmMetrics) IncrementNumConfirmedTxs(ctx context.Context, confirmedTransactions int) {
	promNumConfirmedTxs.WithLabelValues(m.chainID.String()).Add(float64(confirmedTransactions))
	m.numConfirmedTxs.Add(ctx, int64(confirmedTransactions))
}

func (m *txmMetrics) IncrementNumNonceGaps(ctx context.Context) {
	promNumNonceGaps.WithLabelValues(m.chainID.String()).Add(float64(1))
	m.numNonceGaps.Add(ctx, 1)
}

func (m *txmMetrics) ReachedMaxAttempts(ctx context.Context, reached bool) {
	var value float64
	if reached {
		value = 1
	}
	promReachedMaxAttempts.WithLabelValues(m.chainID.String()).Set(value)
	m.reachedMaxAttempts.Record(ctx, int64(value))
}

func (m *txmMetrics) RecordTimeUntilTxConfirmed(ctx context.Context, duration float64) {
	promTimeUntilTxConfirmed.WithLabelValues(m.chainID.String()).Observe(duration)
	m.timeUntilTxConfirmed.Record(ctx, duration)
}

func (m *txmMetrics) SetRPCNonce(ctx context.Context, address common.Address, nonce uint64) {
	promRPCNonce.WithLabelValues(m.chainID.String(), address.String()).Set(float64(nonce))
	m.rpcNonce.Record(ctx, int64(nonce))
}

func (m *txmMetrics) EmitTxMessage(ctx context.Context, txHash common.Hash, fromAddress common.Address, tx *types.Transaction) error {
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

// noopTxmMetrics is a no-op implementation of the Metrics interface.
// It allows the app to run without being blocked in case of metrics initialization errors.
type noopTxmMetrics struct{}

func (noopTxmMetrics) IncrementLifecycleFailure(context.Context, LifecycleFailureStage) {}

func (noopTxmMetrics) IncrementNumBroadcastedTxs(context.Context) {}

func (noopTxmMetrics) IncrementNumConfirmedTxs(context.Context, int) {}

func (noopTxmMetrics) IncrementNumNonceGaps(context.Context) {}

func (noopTxmMetrics) ReachedMaxAttempts(context.Context, bool) {}

func (noopTxmMetrics) RecordTimeUntilTxConfirmed(context.Context, float64) {}

func (noopTxmMetrics) SetRPCNonce(context.Context, common.Address, uint64) {}

func (noopTxmMetrics) EmitTxMessage(context.Context, common.Hash, common.Address, *types.Transaction) error {
	return nil
}
