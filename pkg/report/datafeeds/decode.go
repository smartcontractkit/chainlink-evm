package datafeeds

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"

	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/common"
	df "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/data-feeds/on-chain/registry"
	wt "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/platform"
	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/report/platform"
)

// EVM POR specific processor decodes writes as 'data-feeds.registry.FeedUpdated' messages + metrics
type Processor struct {
	emitter      beholder.ProtoEmitter
	metrics      *df.Metrics
	schema       abi.Arguments
	decodeReport func(*wt.WriteConfirmed, []byte, abi.Arguments) ([]*df.FeedUpdated, error)
}

func NewProcessor(metrics *df.Metrics, emitter beholder.ProtoEmitter, schema abi.Arguments, decodeReport func(*wt.WriteConfirmed, []byte, abi.Arguments) ([]*df.FeedUpdated, error)) *Processor {
	return &Processor{
		metrics:      metrics,
		emitter:      emitter,
		schema:       schema,
		decodeReport: decodeReport,
	}
}

func (p *Processor) Process(ctx context.Context, m proto.Message, attrKVs ...any) error {
	// Switch on the type of the proto.Message
	switch msg := m.(type) {
	case *wt.WriteConfirmed:
		updates, err := p.DecodeAsFeedUpdated(msg)
		if err != nil {
			return fmt.Errorf("failed to decode as 'data-feeds.registry.df.FeedUpdated': %w", err)
		}
		for _, update := range updates {
			err = p.emitter.EmitWithLog(ctx, update, attrKVs...)
			if err != nil {
				return fmt.Errorf("failed to emit with log: %w", err)
			}
			// Process emit and derive metrics
			err = p.metrics.OnFeedUpdated(ctx, update, attrKVs...)
			if err != nil {
				return fmt.Errorf("failed to publish feed updated metrics: %w", err)
			}
		}
		return nil
	default:
		return nil // fallthrough
	}
}

func GetSchema(typ string, internalType string, components []abi.ArgumentMarshaling) abi.Arguments {
	result, err := abi.NewType(typ, internalType, components)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error during abi.NewType: %s", err))
	}
	return abi.Arguments([]abi.Argument{
		{
			// This defines the array of tuple records.
			Type: result,
		},
	})
}

func (p *Processor) DecodeAsFeedUpdated(m *wt.WriteConfirmed) ([]*df.FeedUpdated, error) {
	// Decode the confirmed report (WT -> DF contract event)
	r, err := platform.Decode(m.Report)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	msgs, err := p.decodeReport(m, r.Data, p.schema)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Data Feeds report: %w", err)
	}

	return msgs, nil
}

// newdf.FeedUpdated creates a df.FeedUpdated from the given common parameters.
// If includeTxInfo is true, TxSender and TxReceiver are set.
func NewFeedUpdated(
	m *wt.WriteConfirmed,
	feedID FeedID,
	observationsTimestamp uint32,
	benchmarkPrice *big.Int,
	bundle []byte,
	report []byte,
	includeTxInfo bool,
) *df.FeedUpdated {
	fu := &df.FeedUpdated{
		FeedId:                feedID.String(),
		ObservationsTimestamp: observationsTimestamp,
		Benchmark:             benchmarkPrice.Bytes(),
		Bundle:                bundle,
		Report:                report,
		BenchmarkVal:          ToBenchmarkVal(feedID, benchmarkPrice),

		// Head data - when was the event produced on-chain
		BlockData: m.BlockData,

		ExecutionContext: m.ExecutionContext,
	}

	if includeTxInfo {
		fu.TransactionData = &common.TransactionData{
			TxSender:   m.Transmitter,
			TxReceiver: m.Forwarder,
		}
	}

	return fu
}

// ToBenchmarkVal returns the benchmark i192 on-chain value decoded as an double (float64), scaled by number of decimals (e.g., 1e-18)
// Where the number of decimals is extracted from the feed ID.
//
// This is the largest type Prometheus supports, and this conversion can overflow but so far was sufficient
// for most use-cases. For big numbers, benchmark bytes should be used instead.
//
// Returns `math.NaN()` if report data type not a number, or `+/-Inf` if number doesn't fit in double.
func ToBenchmarkVal(feedID FeedID, val *big.Int) float64 {
	// Return NaN if the value is nil
	if val == nil {
		return math.NaN()
	}

	// Get the number of decimals from the feed ID
	t := feedID.GetDataType()
	decimals, isNumber := GetDecimals(t)

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
