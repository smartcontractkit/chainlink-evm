package dualbroadcast

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

type errorHandler struct{}

func NewErrorHandler() *errorHandler {
	return &errorHandler{}
}

func (e *errorHandler) HandleError(
	ctx context.Context,
	lggr logger.Logger,
	tx *types.Transaction,
	txErr error,
	txStore txm.TxStore,
	setNonce func(common.Address, uint64),
) (noTransmission bool, err error) {
	// Mark the tx as fatal only if this is the first broadcast. In any other case, other txs might be included on-chain.
	if (errors.Is(txErr, ErrNoBids) || errors.Is(txErr, ErrAuction)) && tx.AttemptCount == 1 {
		if err := txStore.MarkTxFatal(ctx, tx, tx.FromAddress); err != nil {
			return false, err
		}
		setNonce(tx.FromAddress, *tx.Nonce)
		l := logger.Sugared(logger.Named(lggr, "Txm.MetaErrorHandler"))
		l.Infow("transaction marked as fatal", "txID", tx.ID, "transactionLifecycleID", tx.GetTransactionLifecycleID(l))
		return true, nil
	}

	// If the error message is not recognized, we don't know if we didn't transmit so return false and continue with the standard execution path.
	return false, nil
}
