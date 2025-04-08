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
	selectionMode := ptr("HighestHead")
	leaseDuration := 0 * time.Second
	pollFailureThreshold := ptr(uint32(5))
	pollInterval := 10 * time.Second
	syncThreshold := ptr(uint32(5))
	nodeIsSyncingEnabled := ptr(false)
	chainTypeStr := ""
	finalizedBlockOffset := ptr[uint32](16)
	enforceRepeatableRead := ptr(true)
	deathDeclarationDelay := time.Second * 3
	noNewFinalizedBlocksThreshold := time.Second * 5
	finalizedBlockPollInterval := time.Second * 4
	newHeadsPollInterval := time.Second * 4
	nodeConfigs := []client.NodeConfig{
		{
			Name:    ptr("foo"),
			WSURL:   ptr("ws://foo.test"),
			HTTPURL: ptr("http://foo.test"),
		},
	}
	finalityDepth := ptr(uint32(10))
	finalityTagEnabled := ptr(true)
	chainCfg, nodePool, nodes, err := client.NewClientConfigs(selectionMode, leaseDuration, chainTypeStr, nodeConfigs,
		pollFailureThreshold, pollInterval, syncThreshold, nodeIsSyncingEnabled, noNewHeadsThreshold, finalityDepth,
		finalityTagEnabled, finalizedBlockOffset, enforceRepeatableRead, deathDeclarationDelay, noNewFinalizedBlocksThreshold,
		finalizedBlockPollInterval, newHeadsPollInterval)
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
			Name:    ptr("BlueEVMPrimaryNode"),
			WSURL:   ptr("ws://no-blue-node"),
			HTTPURL: ptr("http://no-blue-node"),
		},
		{
			Name:    ptr("YellowEVMPrimaryNode"),
			WSURL:   ptr("ws://no-yellow-node"),
			HTTPURL: ptr("http://no-yellow-node"),
		},
	}
	chainCfg, nodePool, nodes, err := client.NewClientConfigs(ptr("HighestHead"), time.Duration(0), "", nodeConfigs,
		ptr[uint32](5), 10*time.Second, ptr[uint32](5), ptr(false), time.Minute, ptr[uint32](5), ptr(false),
		ptr[uint32](5), ptr(false), 10*time.Second, 10*time.Second, 10*time.Second, 10*time.Second)
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
