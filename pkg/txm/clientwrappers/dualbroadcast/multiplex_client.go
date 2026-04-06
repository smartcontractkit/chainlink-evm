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

// TODO(gg): make this consistent with the nova client timeout
const secondarySendTimeout = 5 * time.Second

// TODO(gg): unexport?
type MultiplexClient struct {
	lggr      logger.SugaredLogger
	primary   txm.Client
	secondary txm.Client
}

var _ txm.Client = (*MultiplexClient)(nil)

func NewMultiplexClient(lggr logger.Logger, primary txm.Client, secondary txm.Client) *MultiplexClient {
	return &MultiplexClient{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.MultiplexClient")),
		primary:   primary,
		secondary: secondary,
	}
}

func (m *MultiplexClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	go func() {
		// Use a background-derived context so the secondary isn't cancelled when the
		// parent returns. The timeout bounds the goroutine lifetime.
		sCtx, cancel := context.WithTimeout(context.Background(), secondarySendTimeout)
		defer cancel()
		if err := m.secondary.SendTransaction(sCtx, tx, attempt); err != nil {
			m.lggr.Errorw("Secondary backend send failed", "err", err, "transactionLifecycleID", tx.GetTransactionLifecycleID(m.lggr))
		}
	}()
	return m.primary.SendTransaction(ctx, tx, attempt)
}

func (m *MultiplexClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return m.primary.PendingNonceAt(ctx, address)
}

func (m *MultiplexClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return m.primary.NonceAt(ctx, address, blockNumber)
}
