package processor

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-evm/pkg/monitoring/pb/data-feeds/on-chain/registry"

	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/beholder/monitor"
	wt "github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/platform"
)

// EVM POR specific processor decodes writes as 'data-feeds.registry.FeedUpdated' messages + metrics
type porFeedsProcessor struct {
	emitter monitor.ProtoEmitter
	metrics *registry.Metrics
}

func NewPORFeedsProcessor(metrics *registry.Metrics, emitter monitor.ProtoEmitter) *porFeedsProcessor {
	return &porFeedsProcessor{
		metrics: metrics,
		emitter: emitter,
	}
}

func (p *porFeedsProcessor) Process(ctx context.Context, m proto.Message, attrKVs ...any) error {
	// Switch on the type of the proto.Message
	switch msg := m.(type) {
	case *wt.WriteConfirmed:
		updates, err := registry.DecodePORAsFeedUpdated(msg)
		if err != nil {
			return fmt.Errorf("failed to decode as 'data-feeds.registry.FeedUpdated': %w", err)
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

func (p *porFeedsProcessor) SetEmitter(e monitor.ProtoEmitter) {
	p.emitter = e
}

func (p *porFeedsProcessor) Name() string {
	return "evm-por-feeds"
}
