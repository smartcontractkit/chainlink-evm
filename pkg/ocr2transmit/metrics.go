package ocr2transmit

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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

// RecordOutcome increments confirmed / reverted / fatal counters when calldata matches OCR2 transmit.
func RecordOutcome(chainID *big.Int, from, to common.Address, encodedPayload []byte, fwdrDest *common.Address, outcome string) {
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
	}
}

// SetUnconfirmedGauge sets the gauge for OCR2-shaped unconfirmed txs for TXM v2 (optional / phase 3).
func SetUnconfirmedGauge(chainID *big.Int, from common.Address, n int) {
	if chainID == nil {
		return
	}
	promTransmitUnconfirmed.WithLabelValues(chainID.String(), from.Hex()).Set(float64(n))
}
