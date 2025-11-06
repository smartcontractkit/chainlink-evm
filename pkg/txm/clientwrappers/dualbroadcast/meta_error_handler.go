package dualbroadcast

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

type errorHandler struct{}

func NewErrorHandler() *errorHandler {
	return &errorHandler{}
}

func (e *errorHandler) HandleError(ctx context.Context, tx *types.Transaction, txErr error, txStore txm.TxStore, setNonce func(common.Address, uint64), isFromBroadcastMethod bool) error {
	// If this isn't the first broadcast, don't mark the tx as fatal as other txs might be included on-chain.
	if strings.Contains(txErr.Error(), NoBidsError) && tx.AttemptCount == 1 {
		if err := txStore.MarkTxFatal(ctx, tx, tx.FromAddress); err != nil {
			return err
		}
		setNonce(tx.FromAddress, *tx.Nonce)
		return fmt.Errorf("transaction with txID: %d marked as fatal", tx.ID)
	}

	return txErr
}
