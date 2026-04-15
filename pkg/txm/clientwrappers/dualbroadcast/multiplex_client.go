package dualbroadcast

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

// multiplexClient sends a transaction to the primary OFA client and also fires a fire-and-forget request to the secondary OFA client.
// It delegates nonce queries to the primary, ignoring the secondary.
type multiplexClient struct {
	lggr      logger.SugaredLogger
	primary   txm.Client
	secondary txm.Client
}

var _ txm.Client = (*multiplexClient)(nil)

func newMultiplexClient(lggr logger.Logger, primary txm.Client, secondary txm.Client) *multiplexClient {
	return &multiplexClient{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.MultiplexClient")),
		primary:   primary,
		secondary: secondary,
	}
}

func (m *multiplexClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	go func() {
		// Use a background context so the secondary isn't cancelled when the primary returns.
		// Each client is responsible for its own timeout.
		if err := m.secondary.SendTransaction(context.Background(), tx, attempt); err != nil {
			m.lggr.Errorw("Secondary backend send failed", "err", err, "transactionLifecycleID", tx.GetTransactionLifecycleID(m.lggr))
		}
	}()
	return m.primary.SendTransaction(ctx, tx, attempt)
}

func (m *multiplexClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return m.primary.PendingNonceAt(ctx, address)
}

func (m *multiplexClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return m.primary.NonceAt(ctx, address, blockNumber)
}
