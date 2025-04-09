package registry

import (
	"math"
	"math/big"

	wt_msg "github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
)

// newFeedUpdated creates a FeedUpdated from the given common parameters.
// If includeTxInfo is true, TxSender and TxReceiver are set.
func NewFeedUpdated(
	m *wt_msg.WriteConfirmed,
	feedID datafeeds.FeedID,
	observationsTimestamp uint32,
	benchmarkPrice *big.Int,
	report []byte,
	includeTxInfo bool,
) *FeedUpdated {
	fu := &FeedUpdated{
		FeedId:                feedID.String(),
		ObservationsTimestamp: observationsTimestamp,
		Benchmark:             benchmarkPrice.Bytes(),
		Report:                report,
		BenchmarkVal:          ToBenchmarkVal(feedID, benchmarkPrice),

		// Head data - when was the event produced on-chain
		BlockHash:      m.BlockHash,
		BlockHeight:    m.BlockHeight,
		BlockTimestamp: m.BlockTimestamp,

		// Execution Context - Source
		MetaSourceId: m.MetaSourceId,

		// Execution Context - Chain
		MetaChainFamilyName: m.MetaChainFamilyName,
		MetaChainId:         m.MetaChainId,
		MetaNetworkName:     m.MetaNetworkName,
		MetaNetworkNameFull: m.MetaNetworkNameFull,

		// Execution Context - Workflow (capabilities.RequestMetadata)
		MetaWorkflowId:               m.MetaWorkflowId,
		MetaWorkflowOwner:            m.MetaWorkflowOwner,
		MetaWorkflowExecutionId:      m.MetaWorkflowExecutionId,
		MetaWorkflowName:             m.MetaWorkflowName,
		MetaWorkflowDonId:            m.MetaWorkflowDonId,
		MetaWorkflowDonConfigVersion: m.MetaWorkflowDonConfigVersion,
		MetaReferenceId:              m.MetaReferenceId,

		// Execution Context - Capability
		MetaCapabilityType:           m.MetaCapabilityType,
		MetaCapabilityId:             m.MetaCapabilityId,
		MetaCapabilityTimestampStart: m.MetaCapabilityTimestampStart,
		MetaCapabilityTimestampEmit:  m.MetaCapabilityTimestampEmit,
	}

	if includeTxInfo {
		fu.TxSender = m.Transmitter
		fu.TxReceiver = m.Forwarder
	}

	return fu
}

// toBenchmarkVal returns the benchmark i192 on-chain value decoded as an double (float64), scaled by number of decimals (e.g., 1e-18)
// Where the number of decimals is extracted from the feed ID.
//
// This is the largest type Prometheus supports, and this conversion can overflow but so far was sufficient
// for most use-cases. For big numbers, benchmark bytes should be used instead.
//
// Returns `math.NaN()` if report data type not a number, or `+/-Inf` if number doesn't fit in double.
func ToBenchmarkVal(feedID datafeeds.FeedID, val *big.Int) float64 {
	// Return NaN if the value is nil
	if val == nil {
		return math.NaN()
	}

	// Get the number of decimals from the feed ID
	t := feedID.GetDataType()
	decimals, isNumber := datafeeds.GetDecimals(t)

	// Return NaN if the value is not a number
	if !isNumber {
		return math.NaN()
	}

	// Convert the i192 to a big Float, scaled by the number of decimals
	valF := new(big.Float).SetInt(val)

	if decimals > 0 {
		denominator := big.NewFloat(math.Pow10(int(decimals)))
		valF = new(big.Float).Quo(valF, denominator)
	}

	// Notice: this can overflow, but so far was sufficient for most use-cases
	// On overflow, returns +/-Inf (valid Prometheus value)
	valRes, _ := valF.Float64()
	return valRes
}
