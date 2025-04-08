package heads

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-framework/chains/heads"

	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
)

type saver struct {
	orm      ORM
	config   heads.ChainConfig
	htConfig heads.TrackerConfig
	logger   logger.Logger
	heads    HeadSet
}

var _ heads.Saver[*evmtypes.Head, common.Hash] = (*saver)(nil)

func NewSaver(lggr logger.Logger, orm ORM, config heads.ChainConfig, htConfig heads.TrackerConfig) HeadSaver {
	return &saver{
		orm:      orm,
		config:   config,
		htConfig: htConfig,
		logger:   logger.Named(lggr, "HeadSaver"),
		heads:    NewHeadSet(),
	}
}

func (hs *saver) Save(ctx context.Context, head *evmtypes.Head) error {
	// adding new head might form a cycle, so it's better to validate cached chain before persisting it
	if err := hs.heads.AddHeads(head); err != nil {
		return err
	}

	return hs.orm.IdempotentInsertHead(ctx, head)
}

func (hs *saver) Load(ctx context.Context, latestFinalized int64) (chain *evmtypes.Head, err error) {
	minBlockNumber := hs.calculateMinBlockToKeep(latestFinalized)
	heads, err := hs.orm.LatestHeads(ctx, minBlockNumber)
	if err != nil {
		return nil, err
	}

	err = hs.heads.AddHeads(heads...)
	if err != nil {
		return nil, fmt.Errorf("failed to populate cache with loaded heads: %w", err)
	}
	return hs.heads.LatestHead(), nil
}

func (hs *saver) calculateMinBlockToKeep(latestFinalized int64) int64 {
	return max(latestFinalized-int64(hs.htConfig.HistoryDepth()), 0)
}

func (hs *saver) LatestHeadFromDB(ctx context.Context) (head *evmtypes.Head, err error) {
	return hs.orm.LatestHead(ctx)
}

func (hs *saver) LatestChain() *evmtypes.Head {
	head := hs.heads.LatestHead()
	if head == nil {
		return nil
	}
	if head.ChainLength() < hs.config.FinalityDepth() {
		hs.logger.Debugw("chain shorter than FinalityDepth", "chainLen", head.ChainLength(), "evmFinalityDepth", hs.config.FinalityDepth())
	}
	return head
}

func (hs *saver) Chain(hash common.Hash) *evmtypes.Head {
	return hs.heads.HeadByHash(hash)
}

func (hs *saver) MarkFinalized(ctx context.Context, finalized *evmtypes.Head) error {
	minBlockToKeep := hs.calculateMinBlockToKeep(finalized.BlockNumber())
	if !hs.heads.MarkFinalized(finalized.BlockHash(), minBlockToKeep) {
		return fmt.Errorf("failed to find %s block in the canonical chain to mark it as finalized", finalized)
	}

	return hs.orm.TrimOldHeads(ctx, minBlockToKeep)
}

var NullSaver HeadSaver = &nullSaver{}

type nullSaver struct{}

func (*nullSaver) Save(ctx context.Context, head *evmtypes.Head) error { return nil }
func (*nullSaver) Load(ctx context.Context, latestFinalized int64) (*evmtypes.Head, error) {
	return nil, nil
}
func (*nullSaver) LatestHeadFromDB(ctx context.Context) (*evmtypes.Head, error) { return nil, nil }
func (*nullSaver) LatestChain() *evmtypes.Head                                  { return nil }
func (*nullSaver) Chain(hash common.Hash) *evmtypes.Head                        { return nil }
func (*nullSaver) MarkFinalized(ctx context.Context, latestFinalized *evmtypes.Head) error {
	return nil
}
