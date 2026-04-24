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
	promFeeHistoryEstimatorGasPrice = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gas_price_updater",
		Help: "Sets latest gas price (in Wei)",
	}, []string{"evmChainID"})
	promFeeHistoryEstimatorBaseFee = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "base_fee_updater",
		Help: "Sets latest BaseFee (in Wei)",
	}, []string{"evmChainID"})
	promFeeHistoryEstimatorMaxPriorityFeePerGas = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "max_priority_fee_per_gas_updater",
		Help: "Sets latest MaxPriorityFeePerGas (in Wei)",
	}, []string{"evmChainID"})
	promFeeHistoryEstimatorMaxFeePerGas = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "max_fee_per_gas_updater",
		Help: "Sets latest MaxFeePerGas (in Wei)",
	}, []string{"evmChainID"})
)

// feeHistoryEstimatorMetrics dual-writes to Prometheus and optional Beholder OTel instruments.
// All Record* methods are safe to call when Beholder instruments are nil (no-op for OTel side).
type feeHistoryEstimatorMetrics struct {
	chainID string

	gasPriceGauge             metric.Float64Gauge
	baseFeeGauge              metric.Float64Gauge
	maxPriorityFeePerGasGauge metric.Float64Gauge
	maxFeePerGasGauge         metric.Float64Gauge
}

func newFeeHistoryEstimatorMetrics(lggr logger.SugaredLogger, chainID *big.Int) *feeHistoryEstimatorMetrics {
	m := &feeHistoryEstimatorMetrics{chainID: chainID.String()}
	if g, err := beholder.GetMeter().Float64Gauge("gas_price_updater"); err != nil {
		lggr.Errorw("Failed to register Beholder gas_price_updater gauge", "err", err)
	} else {
		m.gasPriceGauge = g
	}
	if g, err := beholder.GetMeter().Float64Gauge("base_fee_updater"); err != nil {
		lggr.Errorw("Failed to register Beholder base_fee_updater gauge", "err", err)
	} else {
		m.baseFeeGauge = g
	}
	if g, err := beholder.GetMeter().Float64Gauge("max_priority_fee_per_gas_updater"); err != nil {
		lggr.Errorw("Failed to register Beholder max_priority_fee_per_gas_updater gauge", "err", err)
	} else {
		m.maxPriorityFeePerGasGauge = g
	}
	if g, err := beholder.GetMeter().Float64Gauge("max_fee_per_gas_updater"); err != nil {
		lggr.Errorw("Failed to register Beholder max_fee_per_gas_updater gauge", "err", err)
	} else {
		m.maxFeePerGasGauge = g
	}
	return m
}

func (m *feeHistoryEstimatorMetrics) RecordGasPrice(ctx context.Context, value float64) {
	promFeeHistoryEstimatorGasPrice.WithLabelValues(m.chainID).Set(value)
	if m.gasPriceGauge != nil {
		m.gasPriceGauge.Record(ctx, value, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
		))
	}
}

func (m *feeHistoryEstimatorMetrics) RecordBaseFee(ctx context.Context, value float64) {
	promFeeHistoryEstimatorBaseFee.WithLabelValues(m.chainID).Set(value)
	if m.baseFeeGauge != nil {
		m.baseFeeGauge.Record(ctx, value, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
		))
	}
}

func (m *feeHistoryEstimatorMetrics) RecordMaxPriorityFeePerGas(ctx context.Context, value float64) {
	promFeeHistoryEstimatorMaxPriorityFeePerGas.WithLabelValues(m.chainID).Set(value)
	if m.maxPriorityFeePerGasGauge != nil {
		m.maxPriorityFeePerGasGauge.Record(ctx, value, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
		))
	}
}

func (m *feeHistoryEstimatorMetrics) RecordMaxFeePerGas(ctx context.Context, value float64) {
	promFeeHistoryEstimatorMaxFeePerGas.WithLabelValues(m.chainID).Set(value)
	if m.maxFeePerGasGauge != nil {
		m.maxFeePerGasGauge.Record(ctx, value, metric.WithAttributes(
			attribute.String("chainID", m.chainID),
		))
	}
}
