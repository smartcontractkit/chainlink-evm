package registry

import (
	"fmt"
	"math"
	"math/big"

	wt_msg "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/platform"
	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/report/platform"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/por"
)

func DecodeAsFeedUpdated(m *wt_msg.WriteConfirmed, ccip bool) ([]*FeedUpdated, error) {
	// Decode the confirmed report (WT -> DF contract event)
	r, err := platform.Decode(m.Report)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	// Decode the underlying Data Feeds reports
	reports, err := datafeeds.Decode(r.Data, ccip)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Data Feeds report: %w", err)
	}

	// Allocate space for the messages (event per updated feed)
	msgs := make([]*FeedUpdated, 0, len(*reports))

	// Iterate over the underlying Mercury reports
	for _, rf := range *reports {
		feedID := datafeeds.FeedID(rf.FeedID)

		// Notice: this encoding of a DF report doesn't provide a raw underlying report
		msgs = append(msgs, NewFeedUpdated(m, feedID, rf.Timestamp, rf.Price, []byte{}, []byte{}, true))
	}

	return msgs, nil
}

func DecodePORAsFeedUpdated(m *wt_msg.WriteConfirmed) ([]*FeedUpdated, error) {
	// Decode the confirmed report (WT -> DF contract event)
	r, err := platform.Decode(m.Report)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	// Decode the underlying Data Feeds reports
	reports, err := por.Decode(r.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Data Feeds report: %w", err)
	}

	// Allocate space for the messages (event per updated feed)
	msgs := make([]*FeedUpdated, 0, len(*reports))

	// Iterate over the underlying Mercury reports
	for _, rf := range *reports {
		feedID := datafeeds.FeedID(rf.DataId)

		// Notice: uses a placeholder for the benchmark price
		msgs = append(msgs, NewFeedUpdated(m, feedID, rf.Timestamp, big.NewInt(0), rf.Bundle, []byte{}, true))
	}

	return msgs, nil
}

// newFeedUpdated creates a FeedUpdated from the given common parameters.
// If includeTxInfo is true, TxSender and TxReceiver are set.
func NewFeedUpdated(
	m *wt_msg.WriteConfirmed,
	feedID datafeeds.FeedID,
	observationsTimestamp uint32,
	benchmarkPrice *big.Int,
	bundle []byte,
	report []byte,
	includeTxInfo bool,
) *FeedUpdated {
	fu := &FeedUpdated{
		FeedId:                feedID.String(),
		ObservationsTimestamp: observationsTimestamp,
		Benchmark:             benchmarkPrice.Bytes(),
		Bundle:                bundle,
		Report:                report,
		BenchmarkVal:          toBenchmarkVal(feedID, benchmarkPrice),

		// Head data - when was the event produced on-chain
		BlockHash:      m.BlockHash,
		BlockHeight:    m.BlockHeight,
		BlockTimestamp: m.BlockTimestamp,

		// Execution Context - Source
		MetaSourceId: m.ExecutionContext.MetaSourceId,

		// Execution Context - Chain
		MetaChainFamilyName: m.ExecutionContext.MetaChainFamilyName,
		MetaChainId:         m.ExecutionContext.MetaChainId,
		MetaNetworkName:     m.ExecutionContext.MetaNetworkName,
		MetaNetworkNameFull: m.ExecutionContext.MetaNetworkNameFull,

		// Execution Context - Workflow (capabilities.RequestMetadata)
		MetaWorkflowId:               m.ExecutionContext.MetaWorkflowId,
		MetaWorkflowOwner:            m.ExecutionContext.MetaWorkflowOwner,
		MetaWorkflowExecutionId:      m.ExecutionContext.MetaWorkflowExecutionId,
		MetaWorkflowName:             m.ExecutionContext.MetaWorkflowName,
		MetaWorkflowDonId:            m.ExecutionContext.MetaWorkflowDonId,
		MetaWorkflowDonConfigVersion: m.ExecutionContext.MetaWorkflowDonConfigVersion,
		MetaReferenceId:              m.ExecutionContext.MetaReferenceId,

		// Execution Context - Capability
		MetaCapabilityType:           m.ExecutionContext.MetaCapabilityType,
		MetaCapabilityId:             m.ExecutionContext.MetaCapabilityId,
		MetaCapabilityTimestampStart: m.ExecutionContext.MetaCapabilityTimestampStart,
		MetaCapabilityTimestampEmit:  m.ExecutionContext.MetaCapabilityTimestampEmit,
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
func toBenchmarkVal(feedID datafeeds.FeedID, val *big.Int) float64 {
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
