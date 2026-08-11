package client

import (
	"context"
	"fmt"
	"math/big"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
)

// The values Config does not expose are fixed here, mirroring chainlink's own EVM chain
// defaults for a generic EVM chain, so a binary polling the same RPCs through Config.New
// behaves like the node does. The settings Config does expose are defaulted by whoever binds
// it, not here.
const (
	fixedSelectionMode              = "HighestHead"
	fixedLeaseDuration              = 0
	fixedPollFailureThreshold       = uint32(5)
	fixedPollSuccessThreshold       = uint32(0)
	fixedSyncThreshold              = uint32(5)
	fixedNoNewHeadsThreshold        = 3 * time.Minute
	fixedFinalizedBlockOffset       = uint32(0)
	fixedDeathDeclarationDelay      = 10 * time.Second
	fixedNoNewFinalizedHeads        = 0
	fixedFinalizedBlockPollInterval = 5 * time.Second
	fixedNewHeadsPollInterval       = 0
	fixedConfirmationTimeout        = 60 * time.Second
	fixedSafeDepth                  = uint32(0)
)

// Config is the configuration for a multinode-backed [Client] created by [NewClientFromConfig]: the
// settings a caller with a list of RPCs needs, over the full NodePool/Chain TOML the node
// configures. Tagged for chainlink-common's pkg/config/flags, so a binary can bind it as
// flags, env vars and config-file keys directly.
//
// At least one HTTPURL is required. WSURLs are optional; without them the client polls for
// heads rather than subscribing, which is enough for view calls.
// A field carries a `toml` tag only where its key differs from its own name: the plural URL
// slices are each bound to a singular key, since one flag is repeated to build the pool.
type Config struct {
	HTTPURLs  []string `toml:"http-url" usage:"EVM RPC HTTP URL(s); repeat or comma-separate for a multinode pool" validate:"required" example:"['https://rpc.example.com']"`
	WSURLs    []string `toml:"ws-url" usage:"EVM RPC WebSocket URL(s), positionally paired with --evm.http-url; optional" validate:"excluded_without=HTTPURLs"`
	ChainID   string   `usage:"EVM chain ID" validate:"required" example:"'1'"`
	ChainType string   `usage:"EVM chain type (empty for a generic EVM chain)"`

	FinalityTagEnabled bool                  `usage:"use the finalized block tag instead of a finality depth"`
	FinalityDepth      uint32                `usage:"finality depth, used when --evm.finality-tag-enabled=false"`
	PollInterval       commonconfig.Duration `usage:"per-node health poll interval"`
}

// NewClientFromConfig builds and dials a multinode-backed Client from cfg, one node per HTTPURL
// (paired with the WSURL at the same index, when WSURLs are configured).
//
// The returned Client is dialed and ready to use; the caller owns it and must Close it.
func NewClientFromConfig(ctx context.Context, lggr logger.Logger, c Config) (Client, error) {
	if len(c.WSURLs) > 0 && len(c.WSURLs) != len(c.HTTPURLs) {
		return nil, fmt.Errorf("ws url count (%d) must match http url count (%d) when provided", len(c.WSURLs), len(c.HTTPURLs))
	}

	chainID, ok := new(big.Int).SetString(c.ChainID, 10)
	if !ok {
		return nil, fmt.Errorf("invalid chain id %q", c.ChainID)
	}

	nodeCfgs := make([]NodeConfig, len(c.HTTPURLs))
	for i := range c.HTTPURLs {
		name := fmt.Sprintf("node-%d", i)
		order := int32(1)
		sendOnly := false
		loadBalanced := false
		nodeCfg := NodeConfig{
			Name:              &name,
			HTTPURL:           &c.HTTPURLs[i],
			Order:             &order,
			SendOnly:          &sendOnly,
			IsLoadBalancedRPC: &loadBalanced,
		}
		if len(c.WSURLs) > 0 {
			nodeCfg.WSURL = &c.WSURLs[i]
		}
		nodeCfgs[i] = nodeCfg
	}

	selectionMode := fixedSelectionMode
	pollFailureThreshold := fixedPollFailureThreshold
	pollSuccessThreshold := fixedPollSuccessThreshold
	syncThreshold := fixedSyncThreshold
	nodeIsSyncingEnabled := false
	finalityDepth := c.FinalityDepth
	finalityTagEnabled := c.FinalityTagEnabled
	safeTagSupported := false
	finalizedBlockOffset := fixedFinalizedBlockOffset
	enforceRepeatableRead := true
	safeDepth := fixedSafeDepth

	chainCfg, nodePool, nodes, err := NewClientConfigs(
		&selectionMode,
		fixedLeaseDuration,
		c.ChainType,
		nodeCfgs,
		&pollFailureThreshold,
		&pollSuccessThreshold,
		c.PollInterval.Duration(),
		&syncThreshold,
		&nodeIsSyncingEnabled,
		fixedNoNewHeadsThreshold,
		&finalityDepth,
		&finalityTagEnabled,
		&safeTagSupported,
		&finalizedBlockOffset,
		&enforceRepeatableRead,
		fixedDeathDeclarationDelay,
		fixedNoNewFinalizedHeads,
		fixedFinalizedBlockPollInterval,
		fixedNewHeadsPollInterval,
		fixedConfirmationTimeout,
		&safeDepth,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to build evm client configs: %w", err)
	}

	// clientErrors is nil: it only feeds ClassifySendError on the transaction send path, which
	// a Config-built client (no keys, no txm) does not drive.
	cl, err := NewEvmClient(nodePool, chainCfg, nil, lggr, chainID, nodes, chaintype.ChainType(c.ChainType))
	if err != nil {
		return nil, fmt.Errorf("failed to create evm client: %w", err)
	}

	if err := cl.Dial(ctx); err != nil {
		return nil, fmt.Errorf("failed to dial evm client: %w", err)
	}

	lggr.Infow("EVM client dialed", "chainID", chainID, "nodes", len(nodes), "selectionMode", selectionMode)
	return cl, nil
}
