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

// multiplexClient implements txm.Client by sending each broadcast to two OFA backends. The primary
// decides success or failure for TXM; the secondary runs in parallel and only logs errors. Nonce reads
// use the primary so nonce state stays aligned with TXM.
type multiplexClient struct {
	lggr                 logger.SugaredLogger
	primary              txm.Client
	secondary            txm.Client
	secondarySendTimeout time.Duration
}

var _ txm.Client = (*multiplexClient)(nil)

// newMultiplexClient returns a client that multiplexes sends to primary and secondary OFA implementations.
func newMultiplexClient(lggr logger.Logger, primary, secondary txm.Client) *multiplexClient {
	return &multiplexClient{
		lggr:                 logger.Sugared(logger.Named(lggr, "Txm.MultiplexClient")),
		primary:              primary,
		secondary:            secondary,
		secondarySendTimeout: rpcTimeout,
	}
}

func (m *multiplexClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	go func() {
		secondaryCtx, cancel := context.WithTimeout(ctx, m.secondarySendTimeout)
		defer cancel()

		if err := m.secondary.SendTransaction(secondaryCtx, tx, attempt); err != nil {
			// Ignore errors from the secondary backend; it's fire-and-forget.
			m.lggr.Errorw("Secondary backend send failed", "err", err, "transactionLifecycleID", tx.GetTransactionLifecycleID(m.lggr))
		}
	}()

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
