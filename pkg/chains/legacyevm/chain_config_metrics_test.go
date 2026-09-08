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

func TestChainConfigAttributes_exactWhitelist(t *testing.T) {
	t.Parallel()

	attrs := chainConfigAttributes("1", true, false)

	got := map[attribute.Key]attribute.Value{}
	for _, kv := range attrs {
		got[kv.Key] = kv.Value
	}

	// Exactly the three whitelisted keys - nothing else can leak.
	require.Len(t, attrs, 3)
	assert.Equal(t, "1", got["chain_id"].AsString())
	assert.True(t, got["transaction_v2_enabled"].AsBool())
	assert.False(t, got["dual_broadcast"].AsBool())
	assert.NotContains(t, got, attribute.Key("custom_url"))
	assert.NotContains(t, got, attribute.Key("custom_urls"))
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

	metrics.recordConfigInfo(t.Context(), "1", true, false)

	dp := collectChainConfigInfo(t, reader)
	assert.Equal(t, int64(1), dp.Value)
	assert.Equal(t, map[string]string{
		"chain_id":               "1",
		"transaction_v2_enabled": "true",
		"dual_broadcast":         "false",
	}, attrsToStrings(dp.Attributes))
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
	// The configured OFA URL carries a secret and must never reach the metric.
	assert.Equal(t, map[string]string{
		"chain_id":               "42161",
		"transaction_v2_enabled": "true",
		"dual_broadcast":         "false",
	}, attrsToStrings(dp.Attributes))
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
