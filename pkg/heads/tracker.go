package heads

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/mailbox"
	"github.com/smartcontractkit/chainlink-framework/chains/heads"

	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
)

func NewTracker(
	lggr logger.Logger,
	ethClient Client,
	config heads.ChainConfig,
	htConfig heads.TrackerConfig,
	headBroadcaster Broadcaster,
	headSaver HeadSaver,
	mailMon *mailbox.Monitor,
) Tracker {
	return heads.NewTracker[*evmtypes.Head, ethereum.Subscription](
		lggr,
		ethClient,
		config,
		htConfig,
		headBroadcaster,
		headSaver,
		mailMon,
		func() *evmtypes.Head { return nil },
	)
}

var NullTracker Tracker = &nullTracker{}

type nullTracker struct{}

func (*nullTracker) Start(context.Context) error    { return nil }
func (*nullTracker) Close() error                   { return nil }
func (*nullTracker) Ready() error                   { return nil }
func (*nullTracker) HealthReport() map[string]error { return map[string]error{} }
func (*nullTracker) Name() string                   { return "" }
func (*nullTracker) SetLogLevel(zapcore.Level)      {}
func (*nullTracker) Backfill(ctx context.Context, headWithChain *evmtypes.Head, prevHeadWithChain *evmtypes.Head) (err error) {
	return nil
}
func (*nullTracker) LatestChain() *evmtypes.Head { return nil }
func (*nullTracker) LatestAndFinalizedBlock(ctx context.Context) (latest, finalized *evmtypes.Head, err error) {
	return nil, nil, nil
}
