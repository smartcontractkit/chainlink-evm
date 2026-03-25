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

type ocr2TransmitOtelInstruments struct {
	confirmed   metric.Int64Counter
	reverted    metric.Int64Counter
	fatal       metric.Int64Counter
	unconfirmed metric.Int64Gauge
}

var (
	ocr2TransmitOtelInst *ocr2TransmitOtelInstruments
	ocr2TransmitOtelOnce sync.Once
)

func loadOCR2TransmitOtel() *ocr2TransmitOtelInstruments {
	ocr2TransmitOtelOnce.Do(func() {
		m := beholder.GetMeter()
		confirmed, err := m.Int64Counter("ocr2_transmit_tx_confirmed_total")
		if err != nil {
			return
		}
		reverted, err := m.Int64Counter("ocr2_transmit_tx_reverted_total")
		if err != nil {
			return
		}
		fatal, err := m.Int64Counter("ocr2_transmit_tx_fatal_total")
		if err != nil {
			return
		}
		unconfirmed, err := m.Int64Gauge("ocr2_transmit_unconfirmed_tx_count")
		if err != nil {
			return
		}
		ocr2TransmitOtelInst = &ocr2TransmitOtelInstruments{
			confirmed:   confirmed,
			reverted:    reverted,
			fatal:       fatal,
			unconfirmed: unconfirmed,
		}
	})
	return ocr2TransmitOtelInst
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
func RecordOutcome(ctx context.Context, chainID *big.Int, from, to common.Address, encodedPayload []byte, fwdrDest *common.Address, outcome string) {
	if chainID == nil || !IsTransmitCalldata(encodedPayload) {
		return
	}
	contract := ContractLabel(to, fwdrDest)
	labels := []string{chainID.String(), contract, from.Hex()}
	switch outcome {
	case "confirmed":
		promTransmitConfirmed.WithLabelValues(labels...).Inc()
	case "reverted":
		promTransmitReverted.WithLabelValues(labels...).Inc()
	case "fatal":
		promTransmitFatal.WithLabelValues(labels...).Inc()
	default:
		return
	}
	if ot := loadOCR2TransmitOtel(); ot != nil {
		opts := transmitAttrs(chainID, contract, from.Hex())
		switch outcome {
		case "confirmed":
			ot.confirmed.Add(ctx, 1, opts)
		case "reverted":
			ot.reverted.Add(ctx, 1, opts)
		case "fatal":
			ot.fatal.Add(ctx, 1, opts)
		}
	}
}

// SetUnconfirmedGauge sets the gauge for OCR2-shaped unconfirmed txs for TXM v2 (optional / phase 3).
func SetUnconfirmedGauge(ctx context.Context, chainID *big.Int, from common.Address, n int) {
	if chainID == nil {
		return
	}
	promTransmitUnconfirmed.WithLabelValues(chainID.String(), from.Hex()).Set(float64(n))
	if ot := loadOCR2TransmitOtel(); ot != nil {
		ot.unconfirmed.Record(ctx, int64(n), unconfirmedAttrs(chainID, from.Hex()))
	}
}
