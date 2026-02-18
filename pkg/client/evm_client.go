package client

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-framework/metrics"
	"github.com/smartcontractkit/chainlink-framework/multinode"

	evmconfig "github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
)

const QueryTimeout = 10 * time.Second

type ChainConfig interface {
	multinode.ChainConfig
	SafeDepth() uint32
}

func NewEvmClient(cfg evmconfig.NodePool, chainCfg ChainConfig, clientErrors evmconfig.ClientErrors, lggr logger.Logger, chainID *big.Int, nodes []*toml.Node, chainType chaintype.ChainType) (Client, error) {
	var primaries []multinode.Node[*big.Int, *RPCClient]
	var sendonlys []multinode.SendOnlyNode[*big.Int, *RPCClient]
	largePayloadRPCTimeout, defaultRPCTimeout := getRPCTimeouts(chainType)

	multiNodeMetrics, err := metrics.NewGenericMultiNodeMetrics(metrics.EVM, chainID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to initialize metrics: %w", err)
	}

	for i, node := range nodes {
		if node.Archive != nil && *node.Archive {
			// Archive nodes are excluded from the primary pool; they are only used via
			// NewArchiveFilterClient for log backfill fallback.
			continue
		}
		if node.SendOnly != nil && *node.SendOnly {
			rpc := NewRPCClient(cfg, lggr, nil, node.HTTPURL.URL(), *node.Name, i, chainID,
				multinode.Secondary, largePayloadRPCTimeout, defaultRPCTimeout, chainType,
				chainCfg.FinalityTagEnabled(), chainCfg.FinalityDepth(), chainCfg.SafeDepth(), cfg.ExternalRequestMaxResponseSize())
			sendonly := multinode.NewSendOnlyNode(lggr, multiNodeMetrics, (url.URL)(*node.HTTPURL),
				*node.Name, chainID, rpc)
			sendonlys = append(sendonlys, sendonly)
		} else {
			rpc := NewRPCClient(cfg, lggr, node.WSURL.URL(), node.HTTPURL.URL(), *node.Name, i,
				chainID, multinode.Primary, largePayloadRPCTimeout, defaultRPCTimeout, chainType,
				chainCfg.FinalityTagEnabled(), chainCfg.FinalityDepth(), chainCfg.SafeDepth(), cfg.ExternalRequestMaxResponseSize())

			primaryNode := multinode.NewNode(cfg, chainCfg,
				lggr, multiNodeMetrics, node.WSURL.URL(), node.HTTPURL.URL(), *node.Name, i, chainID, *node.Order,
				rpc, "EVM", *node.IsLoadBalancedRPC)
			primaries = append(primaries, primaryNode)
		}
	}

	return NewChainClient(lggr, multiNodeMetrics, cfg.SelectionMode(), cfg.LeaseDuration(),
		primaries, sendonlys, chainID, clientErrors, cfg.DeathDeclarationDelay(), chainType), nil
}

// archiveFilterClient implements logpoller.ArchiveClient using a set of HTTP-only RPCClients
// dedicated to archive reads. Nodes are tried in order; the first successful response is returned.
type archiveFilterClient struct {
	lggr    logger.SugaredLogger
	clients []*RPCClient
}

func (a *archiveFilterClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	var lastErr error
	for _, rpc := range a.clients {
		logs, err := rpc.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		a.lggr.Debugw("Archive node FilterLogs failed, trying next", "err", err)
		lastErr = err
	}
	return nil, lastErr
}

// NewArchiveFilterClient constructs a client from all nodes that have Archive=true.
// Returns nil if no archive nodes are configured. The returned client satisfies the
// logpoller.ArchiveClient interface and can be passed via logpoller.Opts.ArchiveClient.
func NewArchiveFilterClient(cfg evmconfig.NodePool, lggr logger.Logger, chainID *big.Int, nodes []*toml.Node, chainType chaintype.ChainType) (*archiveFilterClient, error) {
	_, defaultRPCTimeout := getRPCTimeouts(chainType)
	var archiveRPCs []*RPCClient
	for i, node := range nodes {
		if node.Archive == nil || !*node.Archive {
			continue
		}
		rpc := NewRPCClient(cfg, lggr, nil, node.HTTPURL.URL(), *node.Name, i, chainID,
			multinode.Secondary, defaultRPCTimeout, defaultRPCTimeout, chainType,
			false, 0, 0, cfg.ExternalRequestMaxResponseSize())
		if err := rpc.DialHTTP(context.Background()); err != nil {
			return nil, fmt.Errorf("failed to dial archive node %q: %w", *node.Name, err)
		}
		archiveRPCs = append(archiveRPCs, rpc)
	}
	if len(archiveRPCs) == 0 {
		return nil, nil
	}
	return &archiveFilterClient{
		lggr:    logger.Sugared(lggr),
		clients: archiveRPCs,
	}, nil
}

func getRPCTimeouts(chainType chaintype.ChainType) (largePayload, defaultTimeout time.Duration) {
	if chaintype.ChainHedera == chainType {
		return 30 * time.Second, QueryTimeout
	}

	return QueryTimeout, QueryTimeout
}
