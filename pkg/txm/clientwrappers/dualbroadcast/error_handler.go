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

func (e *errorHandler) HandleError(
	ctx context.Context,
	tx *types.Transaction,
	err error,
	txStore txm.TxStore,
	setNonce func(common.Address, uint64),
	isFromBroadcastMethod bool,
) error {
	if strings.Contains(err.Error(), NoBidsError) {
		if err := txStore.MarkTxFatal(ctx, tx, tx.FromAddress); err != nil {
			return err
		}
		setNonce(tx.FromAddress, *tx.Nonce)
		return fmt.Errorf("transaction with txID: %d marked as fatal", tx.ID)
	}

	return nil
}
