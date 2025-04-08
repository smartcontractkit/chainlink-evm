package writetarget

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/monitor"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/pb/data-feeds/on-chain/registry"
	wt "github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/pb/platform/on-chain/forwarder"
)

const (
	repoCLLCommon = "https://raw.githubusercontent.com/smartcontractkit/chainlink-common"
	// TODO: replace with main when merged
	versionRefsDevelop = "refs/heads/generalized-monitoring-extraction"
	schemaBasePath     = repoCLLCommon + "/" + versionRefsDevelop + "/pkg/capabilities/writetarget/pb"
)

// NewMonitor initializes a Beholder client for the Write Target
//
// The client is initialized as a BeholderClient extension with a custom ProtoEmitter.
// The ProtoEmitter is proxied with additional processing for emitted messages. This processing
// includes decoding messages as specific types and deriving metrics based on the decoded messages.
// TODO: Report decoding uses the same ABI for EVM and Aptos, however, future chains may need a different
// decoding scheme. Generalize this in the future to support different chains and decoding schemes.
func NewMonitor(ctx context.Context, lggr logger.Logger) (*monitor.BeholderClient, error) {
	// Initialize the Beholder client with a local logger a custom Emitter
	client := beholder.GetClient().ForPackage("write_target")

	registryMetrics, err := registry.NewMetrics()
	if err != nil {
		return nil, fmt.Errorf("failed to create new registry metrics: %w", err)
	}

	forwarderMetrics, err := forwarder.NewMetrics()
	if err != nil {
		return nil, fmt.Errorf("failed to create new forwarder metrics: %w", err)
	}

	wtMetrics, err := wt.NewMetrics()
	if err != nil {
		return nil, fmt.Errorf("failed to create new write target metrics: %w", err)
	}

	// Underlying ProtoEmitter
	emitter := monitor.NewProtoEmitter(lggr, &client, schemaBasePath)

	// Proxy ProtoEmitter with additional processing
	protoEmitterProxy := protoEmitter{
		lggr:    lggr,
		emitter: emitter,
		processors: []monitor.ProtoProcessor{
			&wtProcessor{wtMetrics},
			&keystoneProcessor{emitter, forwarderMetrics},
			&dataFeedsProcessor{emitter, registryMetrics},
		},
	}
	return &monitor.BeholderClient{Client: &client, ProtoEmitter: &protoEmitterProxy}, nil
}

// ProtoEmitter proxy specific to the WT
type protoEmitter struct {
	lggr       logger.Logger
	emitter    monitor.ProtoEmitter
	processors []monitor.ProtoProcessor
}

// Emit emits a proto.Message and runs additional processing
func (e *protoEmitter) Emit(ctx context.Context, m proto.Message, attrKVs ...any) error {
	err := e.emitter.Emit(ctx, m, attrKVs...)
	if err != nil {
		return fmt.Errorf("failed to emit: %w", err)
	}

	// Notice: we skip processing errors (and continue) so this will never error
	return e.Process(ctx, m, attrKVs...)
}

// EmitWithLog emits a proto.Message and runs additional processing
func (e *protoEmitter) EmitWithLog(ctx context.Context, m proto.Message, attrKVs ...any) error {
	err := e.emitter.EmitWithLog(ctx, m, attrKVs...)
	if err != nil {
		return fmt.Errorf("failed to emit with log: %w", err)
	}

	// Notice: we skip processing errors (and continue) so this will never error
	return e.Process(ctx, m, attrKVs...)
}

// Process aggregates further processing for emitted messages
func (e *protoEmitter) Process(ctx context.Context, m proto.Message, attrKVs ...any) error {
	// Further processing for emitted messages
	for _, p := range e.processors {
		err := p.Process(ctx, m, attrKVs...)
		if err != nil {
			// Notice: we swallow and log processing errors
			// These should be investigated and fixed, but are not critical to product runtime,
			// and shouldn't block further processing of the emitted message.
			e.lggr.Errorw("failed to process emitted message", "err", err)
			return nil
		}
	}
	return nil
}

// Write-Target specific processor decodes write messages to derive metrics
type wtProcessor struct {
	metrics *wt.Metrics
}

func (p *wtProcessor) Process(ctx context.Context, m proto.Message, attrKVs ...any) error {
	// Switch on the type of the proto.Message
	switch msg := m.(type) {
	case *wt.WriteInitiated:
		err := p.metrics.OnWriteInitiated(ctx, msg, attrKVs...)
		if err != nil {
			return fmt.Errorf("failed to publish write initiated metrics: %w", err)
		}
		return nil
	case *wt.WriteError:
		err := p.metrics.OnWriteError(ctx, msg, attrKVs...)
		if err != nil {
			return fmt.Errorf("failed to publish write error metrics: %w", err)
		}
		return nil
	case *wt.WriteSent:
		err := p.metrics.OnWriteSent(ctx, msg, attrKVs...)
		if err != nil {
			return fmt.Errorf("failed to publish write sent metrics: %w", err)
		}
		return nil
	case *wt.WriteConfirmed:
		err := p.metrics.OnWriteConfirmed(ctx, msg, attrKVs...)
		if err != nil {
			return fmt.Errorf("failed to publish write confirmed metrics: %w", err)
		}
		return nil
	default:
		return nil // fallthrough
	}
}

// Keystone specific processor decodes writes as 'platform.forwarder.ReportProcessed' messages + metrics
type keystoneProcessor struct {
	emitter monitor.ProtoEmitter
	metrics *forwarder.Metrics
}

func (p *keystoneProcessor) Process(ctx context.Context, m proto.Message, attrKVs ...any) error {
	// Switch on the type of the proto.Message
	switch msg := m.(type) {
	case *wt.WriteConfirmed:
		// TODO: detect the type of write payload (support more than one type of write, first multiple Keystone report versions)
		// https://smartcontract-it.atlassian.net/browse/NONEVM-817
		// Q: Will this msg ever contain different (non-Keystone) types of writes? Hmm.
		// Notice: we assume all writes are Keystone (v1) writes for now

		// Decode as a 'platform.forwarder.ReportProcessed' message
		reportProcessed, err := forwarder.DecodeAsReportProcessed(msg)
		if err != nil {
			return fmt.Errorf("failed to decode as 'platform.forwarder.ReportProcessed': %w", err)
		}
		// Emit the 'platform.forwarder.ReportProcessed' message
		err = p.emitter.EmitWithLog(ctx, reportProcessed, attrKVs...)
		if err != nil {
			return fmt.Errorf("failed to emit with log: %w", err)
		}
		// Process emit and derive metrics
		err = p.metrics.OnReportProcessed(ctx, reportProcessed, attrKVs...)
		if err != nil {
			return fmt.Errorf("failed to publish report processed metrics: %w", err)
		}
		return nil
	default:
		return nil // fallthrough
	}
}

// Data-Feeds specific processor decodes writes as 'data-feeds.registry.FeedUpdated' messages + metrics
type dataFeedsProcessor struct {
	emitter monitor.ProtoEmitter
	metrics *registry.Metrics
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
