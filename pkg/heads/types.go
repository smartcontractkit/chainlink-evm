package heads

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"

	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
	"github.com/smartcontractkit/chainlink-framework/chains/heads"
)

// HeadSaver maintains chains persisted in DB. All methods are thread-safe.
type HeadSaver interface {
	heads.Saver[*evmtypes.Head, common.Hash]
	// LatestHeadFromDB returns the highest seen head from DB.
	LatestHeadFromDB(ctx context.Context) (*evmtypes.Head, error)
}

// Type Alias for EVM Head Tracker Components
type (
	Tracker     = heads.Tracker[*evmtypes.Head, common.Hash]
	Trackable   = heads.Trackable[*evmtypes.Head, common.Hash]
	Listener    = heads.Listener[*evmtypes.Head, common.Hash]
	Broadcaster = heads.Broadcaster[*evmtypes.Head, common.Hash]
	Client      = heads.Client[*evmtypes.Head, ethereum.Subscription, *big.Int, common.Hash]
)
