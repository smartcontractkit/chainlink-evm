package ocr2transmit

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
)

var (
	promTransmitConfirmed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ocr2_transmit_tx_confirmed_total",
		Help: "OCR2 aggregator transmit transactions that received a successful on-chain receipt.",
	}, []string{"chain_id", "contract_address", "from_address"})
	promTransmitReverted = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ocr2_transmit_tx_reverted_total",
		Help: "OCR2 aggregator transmit transactions that were included but reverted.",
	}, []string{"chain_id", "contract_address", "from_address"})
	promTransmitFatal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ocr2_transmit_tx_fatal_total",
		Help: "OCR2 aggregator transmit transactions marked fatally errored by TXM (e.g. could not get receipt).",
	}, []string{"chain_id", "contract_address", "from_address"})
	promTransmitUnconfirmed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ocr2_transmit_unconfirmed_tx_count",
		Help: "Count of unconfirmed in-memory transactions whose calldata is an OCR2 transmit call (TXM v2 per from_address).",
	}, []string{"chain_id", "from_address"})
)

// OCR2TransmitMetrics records OCR2 transmit outcomes to Prometheus and OpenTelemetry (Beholder), following the same pattern as txmMetrics (package-level Prom vecs, OTel instruments on the metrics type).
type OCR2TransmitMetrics struct {
	promConfirmed   *prometheus.CounterVec
	promReverted    *prometheus.CounterVec
	promFatal       *prometheus.CounterVec
	promUnconfirmed *prometheus.GaugeVec

	otelConfirmed   metric.Int64Counter
	otelReverted    metric.Int64Counter
	otelFatal       metric.Int64Counter
	otelUnconfirmed metric.Int64Gauge
}

var (
	ocr2TransmitMetrics     *OCR2TransmitMetrics
	ocr2TransmitMetricsOnce sync.Once
)

func newOCR2TransmitMetrics() *OCR2TransmitMetrics {
	m := &OCR2TransmitMetrics{
		promConfirmed:   promTransmitConfirmed,
		promReverted:    promTransmitReverted,
		promFatal:       promTransmitFatal,
		promUnconfirmed: promTransmitUnconfirmed,
	}
	meter := beholder.GetMeter()
	if c, err := meter.Int64Counter("ocr2_transmit_tx_confirmed_total"); err == nil {
		m.otelConfirmed = c
	}
	if c, err := meter.Int64Counter("ocr2_transmit_tx_reverted_total"); err == nil {
		m.otelReverted = c
	}
	if c, err := meter.Int64Counter("ocr2_transmit_tx_fatal_total"); err == nil {
		m.otelFatal = c
	}
	if g, err := meter.Int64Gauge("ocr2_transmit_unconfirmed_tx_count"); err == nil {
		m.otelUnconfirmed = g
	}
	return m
}

func ocr2TransmitMetricsInstance() *OCR2TransmitMetrics {
	ocr2TransmitMetricsOnce.Do(func() {
		ocr2TransmitMetrics = newOCR2TransmitMetrics()
	})
	return ocr2TransmitMetrics
}

func transmitAttrs(chainID *big.Int, contract, from string) metric.MeasurementOption {
	return metric.WithAttributes(
		attribute.String("chain_id", chainID.String()),
		attribute.String("contract_address", contract),
		attribute.String("from_address", from),
	)
}

func unconfirmedAttrs(chainID *big.Int, from string) metric.MeasurementOption {
	return metric.WithAttributes(
		attribute.String("chain_id", chainID.String()),
		attribute.String("from_address", from),
	)
}

// RecordOutcome increments confirmed / reverted / fatal counters when calldata matches OCR2 transmit.
func (m *OCR2TransmitMetrics) RecordOutcome(ctx context.Context, chainID *big.Int, from, to common.Address, encodedPayload []byte, fwdrDest *common.Address, outcome string) {
	if chainID == nil || !IsTransmitCalldata(encodedPayload) {
		return
	}
	contract := ContractLabel(to, fwdrDest)
	labels := []string{chainID.String(), contract, from.Hex()}
	opts := transmitAttrs(chainID, contract, from.Hex())
	switch outcome {
	case "confirmed":
		m.promConfirmed.WithLabelValues(labels...).Inc()
		if m.otelConfirmed != nil {
			m.otelConfirmed.Add(ctx, 1, opts)
		}
	case "reverted":
		m.promReverted.WithLabelValues(labels...).Inc()
		if m.otelReverted != nil {
			m.otelReverted.Add(ctx, 1, opts)
		}
	case "fatal":
		m.promFatal.WithLabelValues(labels...).Inc()
		if m.otelFatal != nil {
			m.otelFatal.Add(ctx, 1, opts)
		}
	default:
		return
	}
}

// SetUnconfirmedGauge sets the gauge for OCR2-shaped unconfirmed txs for TXM v2 (optional / phase 3).
func (m *OCR2TransmitMetrics) SetUnconfirmedGauge(ctx context.Context, chainID *big.Int, from common.Address, n int) {
	if chainID == nil {
		return
	}
	m.promUnconfirmed.WithLabelValues(chainID.String(), from.Hex()).Set(float64(n))
	if m.otelUnconfirmed != nil {
		m.otelUnconfirmed.Record(ctx, int64(n), unconfirmedAttrs(chainID, from.Hex()))
	}
}

// RecordOutcome increments confirmed / reverted / fatal counters when calldata matches OCR2 transmit.
func RecordOutcome(ctx context.Context, chainID *big.Int, from, to common.Address, encodedPayload []byte, fwdrDest *common.Address, outcome string) {
	ocr2TransmitMetricsInstance().RecordOutcome(ctx, chainID, from, to, encodedPayload, fwdrDest, outcome)
}

// SetUnconfirmedGauge sets the gauge for OCR2-shaped unconfirmed txs for TXM v2 (optional / phase 3).
func SetUnconfirmedGauge(ctx context.Context, chainID *big.Int, from common.Address, n int) {
	ocr2TransmitMetricsInstance().SetUnconfirmedGauge(ctx, chainID, from, n)
}
