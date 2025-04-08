//nolint:gosec, revive // disable G115, revive
package registry

import (
	"context"
	"fmt"
	"strconv"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	beholdercommon "github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/beholder"
)

// ns returns a namespaced metric name
func ns(name string) string {
	return fmt.Sprintf("datafeeds_on_chain_registry_%s", name)
}

// Define metrics configuration
var (
	feedUpdated = struct {
		basic beholder.MetricsInfoCapBasic
		// specific to FeedUpdated
		observationsTimestamp beholder.MetricInfo
		duration              beholder.MetricInfo // ts.emit - ts.observation
		benchmark             beholder.MetricInfo
		blockTimestamp        beholder.MetricInfo
		blockNumber           beholder.MetricInfo
	}{
		basic: beholder.NewMetricsInfoCapBasic(ns("feed_updated"), "data-feeds.on-chain.registry.FeedUpdated"),
		observationsTimestamp: beholder.MetricInfo{
			Name:        ns("feed_updated_observations_timestamp"),
			Unit:        "ms",
			Description: "The observations timestamp for the latest confirmed update (as reported)",
		},
		duration: beholder.MetricInfo{
			Name:        ns("feed_updated_duration"),
			Unit:        "ms",
			Description: "The duration (local) since observation to message: 'data-feeds.on-chain.registry.FeedUpdated' emit",
		},
		benchmark: beholder.MetricInfo{
			Name:        ns("feed_updated_benchmark"),
			Unit:        "",
			Description: "The benchmark value for the latest confirmed update (as reported)",
		},
		blockTimestamp: beholder.MetricInfo{
			Name:        ns("feed_updated_block_timestamp"),
			Unit:        "ms",
			Description: "The block timestamp at the latest confirmed update (as observed)",
		},
		blockNumber: beholder.MetricInfo{
			Name:        ns("feed_updated_block_number"),
			Unit:        "",
			Description: "The block number at the latest confirmed update (as observed)",
		},
	}
)

// Define a new struct for metrics
type Metrics struct {
	// Define on FeedUpdated metrics
	feedUpdated struct {
		basic beholder.MetricsCapBasic
		// specific to FeedUpdated
		observationsTimestamp metric.Int64Gauge
		duration              metric.Int64Gauge // ts.emit - ts.observation
		benchmark             metric.Float64Gauge
		blockTimestamp        metric.Int64Gauge
		blockNumber           metric.Int64Gauge
	}
}

func NewMetrics() (*Metrics, error) {
	// Define new metrics
	m := &Metrics{}

	meter := beholdercommon.GetMeter()

	// Create new metrics
	var err error

	m.feedUpdated.basic, err = beholder.NewMetricsCapBasic(feedUpdated.basic)
	if err != nil {
		return nil, fmt.Errorf("failed to create new basic metrics: %w", err)
	}

	m.feedUpdated.observationsTimestamp, err = feedUpdated.observationsTimestamp.NewInt64Gauge(meter)
	if err != nil {
		return nil, fmt.Errorf("failed to create new gauge: %w", err)
	}

	m.feedUpdated.duration, err = feedUpdated.duration.NewInt64Gauge(meter)
	if err != nil {
		return nil, fmt.Errorf("failed to create new gauge: %w", err)
	}

	m.feedUpdated.benchmark, err = feedUpdated.benchmark.NewFloat64Gauge(meter)
	if err != nil {
		return nil, fmt.Errorf("failed to create new gauge: %w", err)
	}

	m.feedUpdated.blockTimestamp, err = feedUpdated.blockTimestamp.NewInt64Gauge(meter)
	if err != nil {
		return nil, fmt.Errorf("failed to create new gauge: %w", err)
	}

	m.feedUpdated.blockNumber, err = feedUpdated.blockNumber.NewInt64Gauge(meter)
	if err != nil {
		return nil, fmt.Errorf("failed to create new gauge: %w", err)
	}

	return m, nil
}

func (m *Metrics) OnFeedUpdated(ctx context.Context, msg *FeedUpdated, attrKVs ...any) error {
	// Define attributes
	attrs := metric.WithAttributes(msg.Attributes()...)

	// Emit basic metrics (count, timestamps)
	start, emit := msg.MetaCapabilityTimestampStart, msg.MetaCapabilityTimestampEmit
	m.feedUpdated.basic.RecordEmit(ctx, start, emit, msg.Attributes()...)

	// Timestamp e2e observation update
	m.feedUpdated.observationsTimestamp.Record(ctx, int64(msg.ObservationsTimestamp), attrs)
	observation := uint64(msg.ObservationsTimestamp) * 1000 // convert to milliseconds
	m.feedUpdated.duration.Record(ctx, int64(emit-observation), attrs)

	// Benchmark
	m.feedUpdated.benchmark.Record(ctx, msg.BenchmarkVal, attrs)

	// Block timestamp
	m.feedUpdated.blockTimestamp.Record(ctx, int64(msg.BlockTimestamp), attrs)

	// Block number
	blockHeightVal, err := strconv.ParseInt(msg.BlockHeight, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse block height: %w", err)
	}
	m.feedUpdated.blockNumber.Record(ctx, blockHeightVal, attrs)

	return nil
}

// Attributes returns the attributes for the FeedUpdated message to be used in metrics
func (m *FeedUpdated) Attributes() []attribute.KeyValue {
	context := beholder.ExecutionMetadata{
		// Execution Context - Source
		SourceID: m.MetaSourceId,
		// Execution Context - Chain
		ChainFamilyName: m.MetaChainFamilyName,
		ChainID:         m.MetaChainId,
		NetworkName:     m.MetaNetworkName,
		NetworkNameFull: m.MetaNetworkNameFull,
		// Execution Context - Workflow (capabilities.RequestMetadata)
		WorkflowID:               m.MetaWorkflowId,
		WorkflowOwner:            m.MetaWorkflowOwner,
		WorkflowExecutionID:      m.MetaWorkflowExecutionId,
		WorkflowName:             m.MetaWorkflowName,
		WorkflowDonID:            m.MetaWorkflowDonId,
		WorkflowDonConfigVersion: m.MetaWorkflowDonConfigVersion,
		ReferenceID:              m.MetaReferenceId,
		// Execution Context - Capability
		CapabilityType: m.MetaCapabilityType,
		CapabilityID:   m.MetaCapabilityId,
	}

	attrs := []attribute.KeyValue{
		// Transaction Data
		attribute.String("tx_sender", m.TxSender),
		attribute.String("tx_receiver", m.TxReceiver),

		// Event Data
		attribute.String("feed_id", m.FeedId),
		// TODO: do we need these attributes? (available in WriteConfirmed)
		// attribute.Int64("report_id", int64(m.ReportId)), // uint32 -> int64

		// We mark confrmations by transmitter so we can query for only initial (fast) confirmations
		// with PromQL, and ignore the slower confirmations by other signers for SLA measurements.
		attribute.Bool("observed_by_transmitter", m.TxSender == m.MetaSourceId), // source_id == node account
		// TODO: remove once NOT_SET bug with non-string labels is fixed
		attribute.String("observed_by_transmitter_str", strconv.FormatBool(m.TxSender == m.MetaSourceId)),
	}

	return append(attrs, context.Attributes()...)
}
