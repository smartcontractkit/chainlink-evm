package dualbroadcast

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/storage"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

func TestMetaErrorHandler(t *testing.T) {
	errorHandler := NewErrorHandler()
	require.NotNil(t, errorHandler)

	t.Run("handles no bids error for first attempt", func(t *testing.T) {
		nonce := uint64(1)
		address := testutils.NewAddress()
		txRequest := &types.TxRequest{
			ChainID:     testutils.FixtureChainID,
			FromAddress: address,
			ToAddress:   testutils.NewAddress(),
		}
		setNonce := func(address common.Address, nonce uint64) {}
		txStoreManager := storage.NewInMemoryStoreManager(logger.Test(t), testutils.FixtureChainID)
		require.NoError(t, txStoreManager.Add(address))
		txStore := txStoreManager.InMemoryStoreMap[address]
		_ = txStore.CreateTransaction(txRequest)
		tx, err := txStore.UpdateUnstartedTransactionWithNonce(nonce)
		require.NoError(t, err)
		attempt := &types.Attempt{
			TxID:     tx.ID,
			Fee:      gas.EvmFee{GasPrice: assets.NewWeiI(1)},
			GasLimit: 22000,
			Hash:     testutils.NewHash(),
		}
		_, err = txStore.AppendAttemptToTransaction(*tx.Nonce, attempt)
		require.NoError(t, err)
		tx, _ = txStore.FetchUnconfirmedTransactionAtNonceWithCount(nonce)
		err = errorHandler.HandleError(t.Context(), tx, ErrNoBids, txStoreManager, setNonce, false)
		require.Error(t, err)
		require.ErrorContains(t, err, "transaction with txID: 0 marked as fatal")
		_, unconfirmedCount := txStore.FetchUnconfirmedTransactionAtNonceWithCount(nonce)
		assert.Equal(t, 0, unconfirmedCount)
	})

	t.Run("returns txErr if not the first attempt", func(t *testing.T) {
		nonce := uint64(1)
		address := testutils.NewAddress()
		txRequest := &types.TxRequest{
			ChainID:     testutils.FixtureChainID,
			FromAddress: address,
			ToAddress:   testutils.NewAddress(),
		}
		txErr := errors.New("no bids")
		setNonce := func(address common.Address, nonce uint64) {}
		txStoreManager := storage.NewInMemoryStoreManager(logger.Test(t), testutils.FixtureChainID)
		require.NoError(t, txStoreManager.Add(address))
		txStore := txStoreManager.InMemoryStoreMap[address]
		_ = txStore.CreateTransaction(txRequest)
		tx, err := txStore.UpdateUnstartedTransactionWithNonce(nonce)
		require.NoError(t, err)
		attempt := &types.Attempt{
			TxID:     tx.ID,
			Fee:      gas.EvmFee{GasPrice: assets.NewWeiI(1)},
			GasLimit: 22000,
			Hash:     testutils.NewHash(),
		}
		_, err = txStore.AppendAttemptToTransaction(*tx.Nonce, attempt)
		require.NoError(t, err)
		_, err = txStore.AppendAttemptToTransaction(*tx.Nonce, attempt)
		require.NoError(t, err)
		tx, _ = txStore.FetchUnconfirmedTransactionAtNonceWithCount(nonce)
		err = errorHandler.HandleError(t.Context(), tx, txErr, txStoreManager, setNonce, false)
		require.Error(t, err)
		require.ErrorIs(t, err, txErr)
		_, unconfirmedCount := txStore.FetchUnconfirmedTransactionAtNonceWithCount(nonce)
		assert.Equal(t, 1, unconfirmedCount)
	})

	t.Run("handles auction error for first attempt", func(t *testing.T) {
		nonce := uint64(1)
		address := testutils.NewAddress()
		txRequest := &types.TxRequest{
			ChainID:     testutils.FixtureChainID,
			FromAddress: address,
			ToAddress:   testutils.NewAddress(),
		}
		txErr := ErrAuction
		setNonce := func(address common.Address, nonce uint64) {}
		txStoreManager := storage.NewInMemoryStoreManager(logger.Test(t), testutils.FixtureChainID)
		require.NoError(t, txStoreManager.Add(address))
		txStore := txStoreManager.InMemoryStoreMap[address]
		_ = txStore.CreateTransaction(txRequest)
		tx, err := txStore.UpdateUnstartedTransactionWithNonce(nonce)
		require.NoError(t, err)
		attempt := &types.Attempt{
			TxID:     tx.ID,
			Fee:      gas.EvmFee{GasPrice: assets.NewWeiI(1)},
			GasLimit: 22000,
			Hash:     testutils.NewHash(),
		}
		_, err = txStore.AppendAttemptToTransaction(*tx.Nonce, attempt)
		require.NoError(t, err)
		tx, _ = txStore.FetchUnconfirmedTransactionAtNonceWithCount(nonce)
		err = errorHandler.HandleError(t.Context(), tx, txErr, txStoreManager, setNonce, false)
		require.Error(t, err)
		require.ErrorContains(t, err, "transaction with txID: 0 marked as fatal")
		_, unconfirmedCount := txStore.FetchUnconfirmedTransactionAtNonceWithCount(nonce)
		assert.Equal(t, 0, unconfirmedCount)
	})
}
