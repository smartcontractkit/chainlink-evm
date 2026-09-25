package legacyevm

import (
	stdbig "math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/configtest"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/heads"
	"github.com/smartcontractkit/chainlink-evm/pkg/log"
	"github.com/smartcontractkit/chainlink-evm/pkg/logpoller"
	"github.com/smartcontractkit/chainlink-evm/pkg/txmgr"
)

// testLabels is a fully populated label set. Every field differs from at least
// one neighbour of the same type, so a mixed-up assignment in
// chainConfigAttributes shows up as a wrong value rather than passing silently.
var testLabels = chainConfigLabels{
	chainID:               "1",
	chainType:             "optimismBedrock",
	nodePoolSelectionMode: "HighestHead",
	nodeCount:             3,

	txV2Enabled:                 true,
	dualBroadcast:               false,
	readRequestsToMultipleNodes: true,
	bundles:                     false,
	feeBoost:                    true,
	customURLsCount:             2,

	transactionsEnabled:   true,
	forwardersEnabled:     false,
	autoPurgeEnabled:      true,
	finalityTagEnabled:    false,
	nonceAutoSync:         true,
	logBroadcasterEnabled: false,
}

var testLabelsExpected = map[string]string{
	"chain_id":                        "1",
	"chain_type":                      "optimismBedrock",
	"node_pool_selection_mode":        "HighestHead",
	"node_count":                      "3",
	"transaction_v2_enabled":          "true",
	"dual_broadcast":                  "false",
	"read_requests_to_multiple_nodes": "true",
	"bundles":                         "false",
	"fee_boost":                       "true",
	"custom_urls_count":               "2",
	"transactions_enabled":            "true",
	"forwarders_enabled":              "false",
	"auto_purge_enabled":              "true",
	"finality_tag_enabled":            "false",
	"nonce_auto_sync":                 "true",
	"log_broadcaster_enabled":         "false",
}

func TestChainConfigAttributes_exactWhitelist(t *testing.T) {
	t.Parallel()

	attrs := chainConfigAttributes(testLabels)

	// Exactly the whitelisted keys, no duplicates - nothing else can leak.
	require.Len(t, attrs, len(testLabelsExpected))
	got := attrsToStrings(attribute.NewSet(attrs...))
	assert.Equal(t, testLabelsExpected, got)

	// URL-bearing config fields must never become labels, in any form.
	assert.NotContains(t, got, "custom_url")
	assert.NotContains(t, got, "custom_urls")
	assert.NotContains(t, got, "detection_api_url")
}

func TestIsTrue_nilIsFalse(t *testing.T) {
	t.Parallel()

	assert.False(t, isTrue(nil))
	v := true
	assert.True(t, isTrue(&v))
}

func TestChainConfigMetrics_recordConfigInfo(t *testing.T) {
	t.Parallel()

	reader := sdkmetric.NewManualReader()
	meter := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader)).Meter("test")

	metrics, err := newChainConfigMetrics(meter)
	require.NoError(t, err)

	metrics.recordConfigInfo(t.Context(), testLabels)

	dp := collectChainConfigInfo(t, reader)
	assert.Equal(t, int64(1), dp.Value)
	assert.Equal(t, testLabelsExpected, attrsToStrings(dp.Attributes))
}

func TestNewChainConfigLabels_readsWhitelistedConfig(t *testing.T) {
	t.Parallel()

	got := newChainConfigLabels("42161", txV2ChainConfig(t))

	assert.Equal(t, chainConfigLabels{
		chainID:               "42161",
		chainType:             "arbitrum",
		nodePoolSelectionMode: "HighestHead",
		nodeCount:             0,

		txV2Enabled:                 true,
		dualBroadcast:               false,
		readRequestsToMultipleNodes: false,
		bundles:                     false,
		feeBoost:                    false,
		// One OFA URL is configured; only its count is read.
		customURLsCount: 1,

		transactionsEnabled:   true,
		forwardersEnabled:     false,
		autoPurgeEnabled:      false,
		finalityTagEnabled:    true,
		nonceAutoSync:         true,
		logBroadcasterEnabled: true,
	}, got)
}

func TestChain_Start_emitsChainConfigInfo(t *testing.T) {
	t.Parallel()

	reader := sdkmetric.NewManualReader()
	lggr := logger.Test(t)
	cfg := txV2ChainConfig(t)
	c := &chain{
		id:              stdbig.NewInt(42161),
		cfg:             cfg,
		logger:          lggr,
		client:          client.NewNullClient(stdbig.NewInt(42161), lggr),
		txm:             &txmgr.NullTxManager{ErrMsg: "no txm"},
		headBroadcaster: heads.NewBroadcaster(lggr),
		headTracker:     heads.NullTracker,
		logBroadcaster:  &log.NullBroadcaster{ErrMsg: "no log broadcaster"},
		logPoller:       logpoller.LogPollerDisabled,
	}
	metrics, err := newChainConfigMetrics(sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader)).Meter("test"))
	require.NoError(t, err)
	c.chainConfigMetrics = metrics

	require.NoError(t, c.Start(t.Context()))
	t.Cleanup(func() { assert.NoError(t, c.Close()) })

	dp := collectChainConfigInfo(t, reader)
	assert.Equal(t, int64(1), dp.Value)
	assert.Equal(t, map[string]string{
		"chain_id":                        "42161",
		"chain_type":                      "arbitrum",
		"node_pool_selection_mode":        "HighestHead",
		"node_count":                      "0",
		"transaction_v2_enabled":          "true",
		"dual_broadcast":                  "false",
		"read_requests_to_multiple_nodes": "false",
		"bundles":                         "false",
		"fee_boost":                       "false",
		"custom_urls_count":               "1",
		"transactions_enabled":            "true",
		"forwarders_enabled":              "false",
		"auto_purge_enabled":              "false",
		"finality_tag_enabled":            "true",
		"nonce_auto_sync":                 "true",
		"log_broadcaster_enabled":         "true",
	}, attrsToStrings(dp.Attributes))

	// The configured OFA URL carries a secret and must never reach the metric.
	for _, kv := range dp.Attributes.ToSlice() {
		v := kv.Value.String()
		assert.NotContains(t, v, "hunter2")
		assert.NotContains(t, v, "ofa.example.com", "%s leaked the OFA URL", kv.Key)
	}
}

// txV2ChainConfig is a chain config with TransactionManagerV2 enabled, dual
// broadcast off, and an OFA URL that embeds a secret.
func txV2ChainConfig(t *testing.T) *config.ChainScoped {
	return configtest.NewChainScopedConfig(t, func(c *toml.EVMConfig) {
		c.ChainID = sqlutil.NewI(42161)
		enabled, dualBroadcast := true, false
		c.Transactions.TransactionManagerV2 = toml.TransactionManagerV2Config{
			Enabled:       &enabled,
			DualBroadcast: &dualBroadcast,
			CustomURLs:    []*commonconfig.URL{commonconfig.MustParseURL("https://user:hunter2@ofa.example.com")},
		}
	})
}

func collectChainConfigInfo(t *testing.T, reader sdkmetric.Reader) metricdata.DataPoint[int64] {
	t.Helper()

	var rm metricdata.ResourceMetrics
	require.NoError(t, reader.Collect(t.Context(), &rm))

	require.Len(t, rm.ScopeMetrics, 1)
	require.Len(t, rm.ScopeMetrics[0].Metrics, 1)
	m := rm.ScopeMetrics[0].Metrics[0]
	assert.Equal(t, "evm_chain_config_info", m.Name)

	g, ok := m.Data.(metricdata.Gauge[int64])
	require.True(t, ok, "expected an int64 gauge, got %T", m.Data)
	require.Len(t, g.DataPoints, 1)
	return g.DataPoints[0]
}

func attrsToStrings(set attribute.Set) map[string]string {
	out := map[string]string{}
	for _, kv := range set.ToSlice() {
		out[string(kv.Key)] = kv.Value.String()
	}
	return out
}
