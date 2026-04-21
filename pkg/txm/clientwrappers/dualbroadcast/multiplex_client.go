package dualbroadcast

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const rpcTimeout = 10 * time.Second

// multiplexClient implements txm.Client by sending each broadcast to multiple OFA backends.
// The primary decides success or failure for TXM; additional backends run in parallel (best-effort).
// Nonce reads use the primary so nonce state stays aligned with TXM.
type multiplexClient struct {
	lggr                 logger.SugaredLogger
	primary              txm.Client
	secondaries          []txm.Client
	secondarySendTimeout time.Duration
}

var _ txm.Client = (*multiplexClient)(nil)

// newMultiplexClient multiplexes sends to primary (outcome authority) and one or more secondaries.
func newMultiplexClient(lggr logger.Logger, primary txm.Client, secondaries ...txm.Client) *multiplexClient {
	return &multiplexClient{
		lggr:                 logger.Sugared(logger.Named(lggr, "Txm.MultiplexClient")),
		primary:              primary,
		secondaries:          secondaries,
		secondarySendTimeout: rpcTimeout,
	}
}

func (m *multiplexClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	for _, secondary := range m.secondaries {
		sec := secondary
		go func() {
			// Derive timeout from background so completing the primary path does not cancel secondary work early.
			secondaryCtx, cancel := context.WithTimeout(ctx, m.secondarySendTimeout)
			defer cancel()

			if err := sec.SendTransaction(secondaryCtx, tx, attempt); err != nil {
				m.lggr.Errorw("Secondary backend send failed", "err", err, "transactionLifecycleID", tx.GetTransactionLifecycleID(m.lggr))
			}
		}()
	}

	primaryCtx, cancel := context.WithTimeout(ctx, rpcTimeout)
	defer cancel()
	return m.primary.SendTransaction(primaryCtx, tx, attempt)
}

func (m *multiplexClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return m.primary.PendingNonceAt(ctx, address)
}

func (m *multiplexClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return m.primary.NonceAt(ctx, address, blockNumber)
}
