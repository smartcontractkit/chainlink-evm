package client_test

import (
	"bytes"
	"context"
	"errors"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
)

func TestNewEvmClient(t *testing.T) {
	t.Parallel()

	noNewHeadsThreshold := 3 * time.Minute
	selectionMode := new("HighestHead")
	leaseDuration := 0 * time.Second
	pollFailureThreshold := new(uint32(5))
	pollSuccessThreshold := new(uint32(0))
	pollInterval := 10 * time.Second
	syncThreshold := new(uint32(5))
	nodeIsSyncingEnabled := new(false)
	chainTypeStr := ""
	finalizedBlockOffset := new(uint32(16))
	enforceRepeatableRead := new(true)
	deathDeclarationDelay := time.Second * 3
	noNewFinalizedBlocksThreshold := time.Second * 5
	finalizedBlockPollInterval := time.Second * 4
	newHeadsPollInterval := time.Second * 4
	confirmationTimeout := time.Second * 60
	nodeConfigs := []client.NodeConfig{
		{
			Name:    new("foo"),
			WSURL:   new("ws://foo.test"),
			HTTPURL: new("http://foo.test"),
		},
	}
	finalityDepth := new(uint32(10))
	safeDepth := new(uint32(6))
	finalityTagEnabled := new(true)
	SafeTagSupported := new(true)
	chainCfg, nodePool, nodes, err := client.NewClientConfigs(selectionMode, leaseDuration, chainTypeStr, nodeConfigs,
		pollFailureThreshold, pollSuccessThreshold, pollInterval, syncThreshold, nodeIsSyncingEnabled, noNewHeadsThreshold, finalityDepth,
		finalityTagEnabled, SafeTagSupported, finalizedBlockOffset, enforceRepeatableRead, deathDeclarationDelay, noNewFinalizedBlocksThreshold,
		finalizedBlockPollInterval, newHeadsPollInterval, confirmationTimeout, safeDepth)
	require.NoError(t, err)

	client, err := client.NewEvmClient(nodePool, chainCfg, nil, logger.Test(t), testutils.FixtureChainID, nodes, chaintype.ChainType(chainTypeStr))
	require.NotNil(t, client)
	require.NoError(t, err)
}

func TestChainClientMetrics(t *testing.T) {
	ctx, cancel := context.WithTimeout(tests.Context(t), tests.WaitTimeout(t))
	defer cancel()

	nodeConfigs := []client.NodeConfig{
		{
			Name:    new("BlueEVMPrimaryNode"),
			WSURL:   new("ws://no-blue-node"),
			HTTPURL: new("http://no-blue-node"),
		},
		{
			Name:    new("YellowEVMPrimaryNode"),
			WSURL:   new("ws://no-yellow-node"),
			HTTPURL: new("http://no-yellow-node"),
		},
	}
	chainCfg, nodePool, nodes, err := client.NewClientConfigs(new("HighestHead"), time.Duration(0), "", nodeConfigs,
		new(uint32(5)), new(uint32(0)), 10*time.Second, new(uint32(5)), new(false), time.Minute, new(uint32(5)), new(false), new(false),
		new(uint32(5)), new(false), 10*time.Second, 10*time.Second, 10*time.Second, 10*time.Second, 60*time.Second, new(uint32(10)))
	require.NoError(t, err)

	chainID := big.NewInt(68472)
	evmClient, err := client.NewEvmClient(nodePool, chainCfg, nil, logger.Test(t), chainID, nodes, "")
	require.NoError(t, err)

	err = evmClient.Dial(ctx)
	require.NoError(t, err)
	defer evmClient.Close()

	expectedMetrics := map[string]struct{}{
		`evm_pool_rpc_node_dials_total{evmChainID="68472",nodeName="BlueEVMPrimaryNode"}`:   {},
		`evm_pool_rpc_node_dials_total{evmChainID="68472",nodeName="YellowEVMPrimaryNode"}`: {},
		`multi_node_states{chainId="68472",network="EVM",state="Alive"}`:                    {},
		`multi_node_states{chainId="68472",network="EVM",state="Closed"}`:                   {},
		`multi_node_states{chainId="68472",network="EVM",state="Dialed"}`:                   {},
		`multi_node_states{chainId="68472",network="EVM",state="InvalidChainID"}`:           {},
		`multi_node_states{chainId="68472",network="EVM",state="OutOfSync"}`:                {},
		`multi_node_states{chainId="68472",network="EVM",state="Undialed"}`:                 {},
		`multi_node_states{chainId="68472",network="EVM",state="Unreachable"}`:              {},
		`multi_node_states{chainId="68472",network="EVM",state="Unusable"}`:                 {},
	}

	var latestDump string
	foundAll := assert.Eventually(t, func() bool {
		latestDump, err = dumpMetrics()
		if err != nil {
			t.Logf("failed to dump metrics: %v", err)
			return false
		}
		for m := range expectedMetrics {
			if !strings.Contains(latestDump, m) {
				continue
			}

			delete(expectedMetrics, m)
		}
		return len(expectedMetrics) == 0
	}, tests.WaitTimeout(t), tests.TestInterval)
	if !foundAll {
		t.Errorf("Failed to find following metrics in the dump:%v\nDump:\n%s", expectedMetrics, latestDump)
	}
}

func dumpMetrics() (string, error) {
	var writer bytes.Buffer
	enc := expfmt.NewEncoder(&writer, expfmt.NewFormat(expfmt.TypeTextPlain))
	metrics, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return "", errors.New("failed to gather metrics")
	}

	for _, mf := range metrics {
		err = enc.Encode(mf)
		if err != nil {
			return "", errors.New("failed to encode metric")
		}
	}
	return writer.String(), nil
}
