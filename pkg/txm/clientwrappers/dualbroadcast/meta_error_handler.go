package dualbroadcast

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

type errorHandler struct{}

func NewErrorHandler() *errorHandler {
	return &errorHandler{}
}

func (e *errorHandler) HandleError(ctx context.Context, tx *types.Transaction, txErr error, txStore txm.TxStore, setNonce func(common.Address, uint64), isFromBroadcastMethod bool) error {
	// Mark the tx as fatal only if this is the first broadcast. In any other case, other txs might be included on-chain.
	if (errors.Is(txErr, ErrNoBids) || errors.Is(txErr, ErrAuction)) && tx.AttemptCount == 1 {
		if err := txStore.MarkTxFatal(ctx, tx, tx.FromAddress); err != nil {
			return err
		}
		setNonce(tx.FromAddress, *tx.Nonce)
		return fmt.Errorf("transaction with txID: %d marked as fatal", tx.ID)
	}

	return txErr
}
