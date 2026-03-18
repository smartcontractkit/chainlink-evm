package gas

import (
	"context"
	"math/big"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

var (
	promBlockHistoryEstimatorAllGasPricePercentiles = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gas_updater_all_gas_price_percentiles",
		Help: "Gas price at given percentile",
	}, []string{"percentile", "evmChainID"})
	promBlockHistoryEstimatorAllTipCapPercentiles = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gas_updater_all_tip_cap_percentiles",
		Help: "Tip cap at given percentile",
	}, []string{"percentile", "evmChainID"})
	promBlockHistoryEstimatorSetGasPrice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gas_updater_set_gas_price",
		Help: "Gas updater set gas price (in Wei)",
	}, []string{"percentile", "evmChainID"})
	promBlockHistoryEstimatorSetTipCap = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gas_updater_set_tip_cap",
		Help: "Gas updater set gas tip cap (in Wei)",
	}, []string{"percentile", "evmChainID"})
	promBlockHistoryEstimatorCurrentBaseFee = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gas_updater_current_base_fee",
		Help: "Gas updater current block base fee in Wei",
	}, []string{"evmChainID"})
	promBlockHistoryEstimatorConnectivityFailureCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "block_history_estimator_connectivity_failure_count",
		Help: "Counter is incremented every time a gas bump is prevented due to a detected network propagation/connectivity issue",
	}, []string{"evmChainID", "mode"})
)

// blockHistoryEstimatorMetrics dual-writes to Prometheus and optional Beholder OTel instruments.
// All Record* methods are safe to call when Beholder instruments are nil (no-op for OTel side).
type blockHistoryEstimatorMetrics struct {
	chainID string

	gasPriceGauge                   metric.Float64Gauge
	tipCapGauge                     metric.Float64Gauge
	allGasPricePercentilesGauge     metric.Float64Gauge
	allTipCapPercentilesGauge       metric.Float64Gauge
	currentBaseFeeGauge             metric.Float64Gauge
	connectivityFailureCountCounter metric.Int64Counter
}

func newBlockHistoryEstimatorMetrics(lggr logger.SugaredLogger, chainID *big.Int) *blockHistoryEstimatorMetrics {
	m := &blockHistoryEstimatorMetrics{chainID: chainID.String()}

	// otel
	if g, err := beholder.GetMeter().Float64Gauge("gas_updater_set_gas_price"); err != nil {
		lggr.Errorw("Failed to register Beholder gas_updater_set_gas_price gauge", "err", err)
	} else {
		m.gasPriceGauge = g
	}
	if g, err := beholder.GetMeter().Float64Gauge("gas_updater_set_tip_cap"); err != nil {
		lggr.Errorw("Failed to register Beholder gas_updater_set_tip_cap gauge", "err", err)
	} else {
		m.tipCapGauge = g
	}
	if g, err := beholder.GetMeter().Float64Gauge("gas_updater_all_gas_price_percentiles"); err != nil {
		lggr.Errorw("Failed to register Beholder gas_updater_all_gas_price_percentiles gauge", "err", err)
	} else {
		m.allGasPricePercentilesGauge = g
	}
	if g, err := beholder.GetMeter().Float64Gauge("gas_updater_all_tip_cap_percentiles"); err != nil {
		lggr.Errorw("Failed to register Beholder gas_updater_all_tip_cap_percentiles gauge", "err", err)
	} else {
		m.allTipCapPercentilesGauge = g
	}
	if g, err := beholder.GetMeter().Float64Gauge("gas_updater_current_base_fee"); err != nil {
		lggr.Errorw("Failed to register Beholder gas_updater_current_base_fee gauge", "err", err)
	} else {
		m.currentBaseFeeGauge = g
	}
	if c, err := beholder.GetMeter().Int64Counter("block_history_estimator_connectivity_failure_count"); err != nil {
		lggr.Errorw("Failed to register Beholder block_history_estimator_connectivity_failure_count counter", "err", err)
	} else {
		m.connectivityFailureCountCounter = c
	}

	return m
}

func (m *blockHistoryEstimatorMetrics) RecordSetGasPrice(ctx context.Context, percentileLabel string, value float64) {
	promBlockHistoryEstimatorSetGasPrice.WithLabelValues(percentileLabel, m.chainID).Set(value)
	if m.gasPriceGauge != nil {
		m.gasPriceGauge.Record(ctx, value, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("percentile", percentileLabel),
		))
	}
}

func (m *blockHistoryEstimatorMetrics) RecordSetTipCap(ctx context.Context, percentileLabel string, value float64) {
	promBlockHistoryEstimatorSetTipCap.WithLabelValues(percentileLabel, m.chainID).Set(value)
	if m.tipCapGauge != nil {
		m.tipCapGauge.Record(ctx, value, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("percentile", percentileLabel),
		))
	}
}

func (m *blockHistoryEstimatorMetrics) RecordAllGasPricePercentile(ctx context.Context, percentileLabel string, value float64) {
	promBlockHistoryEstimatorAllGasPricePercentiles.WithLabelValues(percentileLabel, m.chainID).Set(value)
	if m.allGasPricePercentilesGauge != nil {
		m.allGasPricePercentilesGauge.Record(ctx, value, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("percentile", percentileLabel),
		))
	}
}

func (m *blockHistoryEstimatorMetrics) RecordAllTipCapPercentile(ctx context.Context, percentileLabel string, value float64) {
	promBlockHistoryEstimatorAllTipCapPercentiles.WithLabelValues(percentileLabel, m.chainID).Set(value)
	if m.allTipCapPercentilesGauge != nil {
		m.allTipCapPercentilesGauge.Record(ctx, value, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("percentile", percentileLabel),
		))
	}
}

func (m *blockHistoryEstimatorMetrics) RecordCurrentBaseFee(ctx context.Context, value float64) {
	promBlockHistoryEstimatorCurrentBaseFee.WithLabelValues(m.chainID).Set(value)
	if m.currentBaseFeeGauge != nil {
		m.currentBaseFeeGauge.Record(ctx, value, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
		))
	}
}

func (m *blockHistoryEstimatorMetrics) RecordConnectivityFailure(ctx context.Context, mode string) {
	promBlockHistoryEstimatorConnectivityFailureCount.WithLabelValues(m.chainID, mode).Inc()
	if m.connectivityFailureCountCounter != nil {
		m.connectivityFailureCountCounter.Add(ctx, 1, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
			attribute.String("mode", mode),
		))
	}
}
