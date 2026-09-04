package legacyevm

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
)

// nodeConfigInfoMetricName is an info-style gauge: its value is always 1 and all
// state is carried in the labels.
const nodeConfigInfoMetricName = "node_config_info"

// nodeConfigAttributes returns the exhaustive, whitelisted label set for the
// node_config_info metric for one EVM chain.
//
// The whitelist is the security boundary of this metric: it must stay limited to
// low-cardinality, non-sensitive values. In particular it must never carry an
// RPC or OFA URL (TransactionManagerV2.CustomURL/CustomURLs), because those can
// embed credentials. See docs on OEV-1648 / INCIDENT-2541.
func nodeConfigAttributes(chainID string, txV2Enabled, dualBroadcast bool) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("chain_id", chainID),
		attribute.Bool("transaction_v2_enabled", txV2Enabled),
		attribute.Bool("dual_broadcast", dualBroadcast),
	}
}

// derefBool reads an optional config bool, treating an unset value as false.
func derefBool(b *bool) bool { return b != nil && *b }

// recordNodeConfigInfo records the node_config_info gauge for a single chain.
// A synchronous gauge re-exports its last recorded value on every reader
// interval, so a single record at startup keeps the series alive.
func recordNodeConfigInfo(ctx context.Context, meter metric.Meter, chainID string, txV2Enabled, dualBroadcast bool) error {
	gauge, err := meter.Int64Gauge(
		nodeConfigInfoMetricName,
		metric.WithDescription("SVR-relevant EVM node config; value is always 1, state is in the labels"),
		metric.WithUnit("{info}"),
	)
	if err != nil {
		return fmt.Errorf("failed to create %s gauge: %w", nodeConfigInfoMetricName, err)
	}

	gauge.Record(ctx, 1, metric.WithAttributes(nodeConfigAttributes(chainID, txV2Enabled, dualBroadcast)...))
	return nil
}

// emitNodeConfigInfo reports this chain's SVR-relevant config state. A nil meter
// falls back to the global beholder meter. Failures are non-fatal: a metrics
// hiccup must never block chain startup.
func (c *chain) emitNodeConfigInfo(ctx context.Context, meter metric.Meter) {
	if meter == nil {
		meter = beholder.GetMeter()
	}

	txV2 := c.cfg.EVM().Transactions().TransactionManagerV2()
	if err := recordNodeConfigInfo(ctx, meter, c.id.String(), txV2.Enabled(), derefBool(txV2.DualBroadcast())); err != nil {
		c.logger.Warnw("Failed to record node config info metric", "chainID", c.id, "err", err)
	}
}
