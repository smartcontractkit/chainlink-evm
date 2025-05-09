package client

import (
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-framework/metrics"
	"github.com/smartcontractkit/chainlink-framework/multinode"

	evmconfig "github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
)

const QueryTimeout = 10 * time.Second

func NewEvmClient(cfg evmconfig.NodePool, chainCfg multinode.ChainConfig, clientErrors evmconfig.ClientErrors, lggr logger.Logger, chainID *big.Int, nodes []*toml.Node, chainType chaintype.ChainType) (Client, error) {
	var primaries []multinode.Node[*big.Int, *RPCClient]
	var sendonlys []multinode.SendOnlyNode[*big.Int, *RPCClient]
	largePayloadRPCTimeout, defaultRPCTimeout := getRPCTimeouts(chainType)

	multiNodeMetrics, err := metrics.NewGenericMultiNodeMetrics(metrics.EVM, chainID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize metrics: %w", err)
	}

	for i, node := range nodes {
		if node.SendOnly != nil && *node.SendOnly {
			rpc := NewRPCClient(cfg, lggr, nil, node.HTTPURL.URL(), *node.Name, i, chainID,
				multinode.Secondary, largePayloadRPCTimeout, defaultRPCTimeout, chainType)
			sendonly := multinode.NewSendOnlyNode(lggr, multiNodeMetrics, (url.URL)(*node.HTTPURL),
				*node.Name, chainID, rpc)
			sendonlys = append(sendonlys, sendonly)
		} else {
			rpc := NewRPCClient(cfg, lggr, node.WSURL.URL(), node.HTTPURL.URL(), *node.Name, i,
				chainID, multinode.Primary, largePayloadRPCTimeout, defaultRPCTimeout, chainType)

			primaryNode := multinode.NewNode(cfg, chainCfg,
				lggr, multiNodeMetrics, node.WSURL.URL(), node.HTTPURL.URL(), *node.Name, i, chainID, *node.Order,
				rpc, "EVM")
			primaries = append(primaries, primaryNode)
		}
	}

	return NewChainClient(lggr, multiNodeMetrics, cfg.SelectionMode(), cfg.LeaseDuration(),
		primaries, sendonlys, chainID, clientErrors, cfg.DeathDeclarationDelay(), chainType), nil
}

func getRPCTimeouts(chainType chaintype.ChainType) (largePayload, defaultTimeout time.Duration) {
	if chaintype.ChainHedera == chainType {
		return 30 * time.Second, QueryTimeout
	}

	return QueryTimeout, QueryTimeout
}
