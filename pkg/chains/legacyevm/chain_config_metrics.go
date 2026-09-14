package legacyevm

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// evmChainConfigInfoMetricName is an info-style gauge: its value is always 1 and all
// state is carried in the labels.
const evmChainConfigInfoMetricName = "evm_chain_config_info"

// chainConfigMetrics holds the instruments for reporting a chain's whitelisted
// configuration state.
type chainConfigMetrics struct {
	configInfo metric.Int64Gauge
}

func newChainConfigMetrics(meter metric.Meter) (*chainConfigMetrics, error) {
	configInfo, err := meter.Int64Gauge(
		evmChainConfigInfoMetricName,
		metric.WithDescription("Whitelisted EVM chain configuration; value is always 1, state is in the labels"),
		metric.WithUnit("{info}"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create %s gauge: %w", evmChainConfigInfoMetricName, err)
	}

	return &chainConfigMetrics{configInfo: configInfo}, nil
}

// recordConfigInfo records the config gauge for a single chain. A synchronous
// gauge re-exports its last recorded value on every reader interval, so a single
// record at startup keeps the series alive.
func (m *chainConfigMetrics) recordConfigInfo(ctx context.Context, chainID string, txV2Enabled, dualBroadcast bool) {
	m.configInfo.Record(ctx, 1, metric.WithAttributes(chainConfigAttributes(chainID, txV2Enabled, dualBroadcast)...))
}

// chainConfigAttributes returns the exhaustive, whitelisted label set for the
// evm_chain_config_info metric for one EVM chain.
//
// The whitelist is the security boundary of this metric: it must stay limited to
// low-cardinality, non-sensitive values. In particular it must never carry an
// RPC or OFA URL (TransactionManagerV2.CustomURL/CustomURLs), because those can
// embed credentials. See docs on OEV-1648 / INCIDENT-2541.
func chainConfigAttributes(chainID string, txV2Enabled, dualBroadcast bool) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("chain_id", chainID),
		attribute.Bool("transaction_v2_enabled", txV2Enabled),
		attribute.Bool("dual_broadcast", dualBroadcast),
	}
}

// isTrue reads an optional config bool, treating an unset value as false.
func isTrue(b *bool) bool { return b != nil && *b }
