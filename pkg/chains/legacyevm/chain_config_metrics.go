package legacyevm

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-evm/pkg/config"
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
func (m *chainConfigMetrics) recordConfigInfo(ctx context.Context, l chainConfigLabels) {
	m.configInfo.Record(ctx, 1, metric.WithAttributes(chainConfigAttributes(l)...))
}

// chainConfigLabels is a plain-value snapshot of one chain's whitelisted config
// state. It holds no config types on purpose, so chainConfigAttributes has no
// way to reach a sensitive field.
type chainConfigLabels struct {
	chainID               string
	chainType             string
	nodePoolSelectionMode string
	nodeCount             int

	txV2Enabled                 bool
	dualBroadcast               bool
	readRequestsToMultipleNodes bool
	bundles                     bool
	feeBoost                    bool
	customURLsCount             int

	transactionsEnabled   bool
	forwardersEnabled     bool
	autoPurgeEnabled      bool
	finalityTagEnabled    bool
	nonceAutoSync         bool
	logBroadcasterEnabled bool
}

// newChainConfigLabels reads the whitelisted fields out of a chain's config.
func newChainConfigLabels(chainID string, cfg *config.ChainScoped) chainConfigLabels {
	evm := cfg.EVM()
	txs := evm.Transactions()
	txV2 := txs.TransactionManagerV2()

	return chainConfigLabels{
		chainID:               chainID,
		chainType:             string(evm.ChainType()),
		nodePoolSelectionMode: evm.NodePool().SelectionMode(),
		nodeCount:             len(cfg.Nodes()),

		txV2Enabled:                 txV2.Enabled(),
		dualBroadcast:               isTrue(txV2.DualBroadcast()),
		readRequestsToMultipleNodes: isTrue(txV2.ReadRequestsToMultipleNodes()),
		bundles:                     isTrue(txV2.Bundles()),
		feeBoost:                    txV2.FeeBoost(),
		// How many custom OFA/RPC endpoints are configured - never the URLs.
		customURLsCount: len(txV2.CustomURLs()),

		transactionsEnabled: txs.Enabled(),
		forwardersEnabled:   txs.ForwardersEnabled(),
		// AutoPurge contributes its Enabled flag only; DetectionApiUrl is a URL
		// and is off-limits for the same reason CustomURLs is.
		autoPurgeEnabled:      txs.AutoPurge().Enabled(),
		finalityTagEnabled:    evm.FinalityTagEnabled(),
		nonceAutoSync:         evm.NonceAutoSync(),
		logBroadcasterEnabled: evm.LogBroadcasterEnabled(),
	}
}

// chainConfigAttributes returns the exhaustive, whitelisted label set for the
// evm_chain_config_info metric for one EVM chain.
//
// The whitelist is the security boundary of this metric: it must stay limited to
// low-cardinality, non-sensitive values. In particular it must never carry an
// RPC or OFA URL (TransactionManagerV2.CustomURL/CustomURLs,
// Transactions.AutoPurge.DetectionApiUrl), because those can embed credentials -
// not even sanitized or hashed. See docs on OEV-1648 / INCIDENT-2541.
//
// Config is static per node, so each node emits exactly one series per chain:
// widening this set with flags and bounded enums does not multiply series.
// Unbounded values (durations, depths, limits) belong in their own gauges.
func chainConfigAttributes(l chainConfigLabels) []attribute.KeyValue {
	return []attribute.KeyValue{
		attribute.String("chain_id", l.chainID),
		attribute.String("chain_type", l.chainType),
		attribute.String("node_pool_selection_mode", l.nodePoolSelectionMode),
		attribute.Int("node_count", l.nodeCount),

		attribute.Bool("transaction_v2_enabled", l.txV2Enabled),
		attribute.Bool("dual_broadcast", l.dualBroadcast),
		attribute.Bool("read_requests_to_multiple_nodes", l.readRequestsToMultipleNodes),
		attribute.Bool("bundles", l.bundles),
		attribute.Bool("fee_boost", l.feeBoost),
		attribute.Int("custom_urls_count", l.customURLsCount),

		attribute.Bool("transactions_enabled", l.transactionsEnabled),
		attribute.Bool("forwarders_enabled", l.forwardersEnabled),
		attribute.Bool("auto_purge_enabled", l.autoPurgeEnabled),
		attribute.Bool("finality_tag_enabled", l.finalityTagEnabled),
		attribute.Bool("nonce_auto_sync", l.nonceAutoSync),
		attribute.Bool("log_broadcaster_enabled", l.logBroadcasterEnabled),
	}
}

// isTrue reads an optional config bool, treating an unset value as false.
func isTrue(b *bool) bool { return b != nil && *b }
