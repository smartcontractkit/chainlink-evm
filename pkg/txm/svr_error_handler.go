package txm

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

var _ ErrorHandler = &SvrErrorHandler{}

type SvrErrorHandler struct{}

// HandleError marks the errored transaction as purgeable in any case of broadcast error. The window to get bids on a particular transaction is small so if an error is encountered it's likely too late.
// Marking the transaction as purgeable allows the nonce to get filled and let other transactions through.
func (s SvrErrorHandler) HandleError(ctx context.Context, tx *types.Transaction, _ error, _ AttemptBuilder, client Client, txStore TxStore, _ func(common.Address, uint64)) (err error) {
	pendingNonce, pErr := client.PendingNonceAt(ctx, tx.FromAddress)
	if pErr != nil {
		return pErr
	}
	// No-op transaction already exists for nonce on-chain
	if pendingNonce > *tx.Nonce {
		return nil
	}
	return txStore.MarkUnconfirmedTransactionPurgeable(ctx, *tx.Nonce, tx.FromAddress)
}
