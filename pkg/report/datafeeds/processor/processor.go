package processor

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	monitor "github.com/smartcontractkit/chainlink-evm/pkg/monitor/beholder"
	"github.com/smartcontractkit/chainlink-evm/pkg/monitoring/pb/data-feeds/on-chain/registry"
	wt "github.com/smartcontractkit/chainlink-evm/pkg/monitoring/pb/platform"
)

// EVM Data-Feeds specific processor decodes writes as 'data-feeds.registry.FeedUpdated' messages + metrics
type dataFeedsProcessor struct {
	emitter monitor.ProtoEmitter
	metrics *registry.Metrics
}

func NewDataFeedsProcessor(metrics *registry.Metrics) *dataFeedsProcessor {
	return &dataFeedsProcessor{
		metrics: metrics,
	}
}

func (p *dataFeedsProcessor) Process(ctx context.Context, m proto.Message, attrKVs ...any) error {
	// Switch on the type of the proto.Message
	switch msg := m.(type) {
	case *wt.WriteConfirmed:
		// TODO: fallthrough if not a write containing a DF report
		// https://smartcontract-it.atlassian.net/browse/NONEVM-818
		// Notice: we assume all writes are Data-Feeds (static schema) writes for now

		// Decode as an array of 'data-feeds.registry.FeedUpdated' messages
		updates, err := registry.DecodeAsFeedUpdated(msg)
		if err != nil {
			return fmt.Errorf("failed to decode as 'data-feeds.registry.FeedUpdated': %w", err)
		}
		// Emit the 'data-feeds.registry.FeedUpdated' messages
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

func (p *dataFeedsProcessor) SetEmitter(e monitor.ProtoEmitter) {
	p.emitter = e
}

func (p *dataFeedsProcessor) GetName() string {
	return "datafeeds"
}
