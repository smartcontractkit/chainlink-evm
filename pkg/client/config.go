package client

import (
	"context"
	"fmt"
	"math/big"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
)

// Config is the configuration for a multinode-backed [Client] created by [NewClientFromConfig]: the
// settings a caller with a list of RPCs needs, over the full NodePool/Chain TOML the node
// configures. Tagged for chainlink-common's pkg/config/flags, so a binary can bind it as
// flags, env vars and config-file keys directly.
//
// At least one HTTPURL is required. WSURLs are optional; without them the client polls for
// heads rather than subscribing, which is enough for view calls.
type Config struct {
	HTTPURLs  []string `toml:"http-url" usage:"EVM RPC HTTP URL(s); repeat or comma-separate for a multinode pool" validate:"required" example:"['https://rpc.example.com']"`
	WSURLs    []string `toml:"ws-url" usage:"EVM RPC WebSocket URL(s), positionally paired with --evm.http-url; optional" validate:"excluded_without=HTTPURLs"`
	ChainID   string   `usage:"EVM chain ID" validate:"required" example:"'1'"`
	ChainType string   `usage:"EVM chain type (empty for a generic EVM chain)"`

	SelectionMode              string                `usage:"node selection mode"`
	LeaseDuration              commonconfig.Duration `usage:"lease duration for node selection"`
	PollFailureThreshold       uint32                `usage:"poll failure threshold before marking node unhealthy"`
	PollSuccessThreshold       uint32                `usage:"poll success threshold before marking node healthy"`
	PollInterval               commonconfig.Duration `usage:"per-node health poll interval"`
	SyncThreshold              uint32                `usage:"sync threshold for head height lag"`
	NodeIsSyncingEnabled       bool                  `usage:"enable checking eth_syncing state"`
	NoNewHeadsThreshold        commonconfig.Duration `usage:"duration without new heads before declaring out of sync"`
	FinalityTagEnabled         bool                  `usage:"use the finalized block tag instead of a finality depth"`
	FinalityDepth              uint32                `usage:"finality depth, used when --evm.finality-tag-enabled=false"`
	SafeTagSupported           bool                  `usage:"whether safe block tag is supported"`
	FinalizedBlockOffset       uint32                `usage:"offset applied to finalized block tags"`
	EnforceRepeatableRead      bool                  `usage:"enforce repeatable read across nodes"`
	DeathDeclarationDelay      commonconfig.Duration `usage:"delay before declaring a node dead"`
	NoNewFinalizedHeads        commonconfig.Duration `usage:"threshold duration for missing finalized heads"`
	FinalizedBlockPollInterval commonconfig.Duration `usage:"poll interval for finalized block tag"`
	NewHeadsPollInterval       commonconfig.Duration `usage:"poll interval for new heads when WS is disabled"`
	ConfirmationTimeout        commonconfig.Duration `usage:"timeout duration for block confirmation"`
	SafeDepth                  uint32                `usage:"safe depth block offset"`
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

	chainCfg, nodePool, nodes, err := NewClientConfigs(
		&c.SelectionMode,
		c.LeaseDuration.Duration(),
		c.ChainType,
		nodeCfgs,
		&c.PollFailureThreshold,
		&c.PollSuccessThreshold,
		c.PollInterval.Duration(),
		&c.SyncThreshold,
		&c.NodeIsSyncingEnabled,
		c.NoNewHeadsThreshold.Duration(),
		&c.FinalityDepth,
		&c.FinalityTagEnabled,
		&c.SafeTagSupported,
		&c.FinalizedBlockOffset,
		&c.EnforceRepeatableRead,
		c.DeathDeclarationDelay.Duration(),
		c.NoNewFinalizedHeads.Duration(),
		c.FinalizedBlockPollInterval.Duration(),
		c.NewHeadsPollInterval.Duration(),
		c.ConfirmationTimeout.Duration(),
		&c.SafeDepth,
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

	lggr.Infow("EVM client dialed", "chainID", chainID, "nodes", len(nodes), "selectionMode", c.SelectionMode)
	return cl, nil
}
