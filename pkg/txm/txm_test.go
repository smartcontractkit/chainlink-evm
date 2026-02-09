package txm_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/mocks"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys/keystest"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers/dualbroadcast"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/storage"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

func TestLifecycle(t *testing.T) {
	t.Parallel()

	client := txm.NewMockClient(t)
	ab := txm.NewMockAttemptBuilder(t)
	address1 := testutils.NewAddress()
	address2 := testutils.NewAddress()
	assert.NotEqual(t, address1, address2)
	addresses := []common.Address{address1, address2}

	t.Run("retries if initial pending nonce call fails", func(t *testing.T) {
		lggr, observedLogs := logger.TestObserved(t, zap.DebugLevel)
		config := txm.Config{BlockTime: 1 * time.Minute}
		txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
		require.NoError(t, txStore.Add(address1))
		keystore := keystest.Addresses{address1}
		tm := txm.NewTxm(lggr, testutils.FixtureChainID, client, nil, txStore, nil, config, keystore, nil)
		client.On("PendingNonceAt", mock.Anything, address1).Return(uint64(0), errors.New("error")).Once()
		client.On("PendingNonceAt", mock.Anything, address1).Return(uint64(100), nil).Once()
		servicetest.Run(t, tm)
		tests.AssertLogEventually(t, observedLogs, "Error when fetching initial nonce")
		tests.AssertLogEventually(t, observedLogs, fmt.Sprintf("Set initial nonce for address: %v to %d", address1, 100))
	})

	t.Run("tests lifecycle successfully without any transactions", func(t *testing.T) {
		config := txm.Config{BlockTime: 200 * time.Millisecond}
		keystore := keystest.Addresses(addresses)
		lggr, observedLogs := logger.TestObserved(t, zap.DebugLevel)
		txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
		require.NoError(t, txStore.Add(addresses...))
		tx := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, txStore, nil, config, keystore, nil)
		var nonce uint64
		// Start
		client.On("PendingNonceAt", mock.Anything, address1).Return(nonce, nil).Once()
		client.On("PendingNonceAt", mock.Anything, address2).Return(nonce, nil).Once()
		// backfill loop (may or may not be executed multiple times)
		client.On("NonceAt", mock.Anything, address1, mock.Anything).Return(nonce, nil).Maybe()
		client.On("NonceAt", mock.Anything, address2, mock.Anything).Return(nonce, nil).Maybe()

		servicetest.Run(t, tx)
		tests.AssertLogEventually(t, observedLogs, "Backfill time elapsed")
	})
}

func TestTrigger(t *testing.T) {
	t.Parallel()

	address := testutils.NewAddress()

	t.Run("Trigger fails if Txm is unstarted", func(t *testing.T) {
		lggr, observedLogs := logger.TestObserved(t, zap.ErrorLevel)
		txm := txm.NewTxm(lggr, nil, nil, nil, nil, nil, txm.Config{}, keystest.Addresses{}, nil)
		txm.Trigger(address)
		tests.AssertLogEventually(t, observedLogs, "Txm unstarted")
	})

	t.Run("executes Trigger", func(t *testing.T) {
		lggr := logger.Test(t)
		txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
		require.NoError(t, txStore.Add(address))
		client := txm.NewMockClient(t)
		ab := txm.NewMockAttemptBuilder(t)
		config := txm.Config{BlockTime: 1 * time.Minute, RetryBlockThreshold: 10}
		keystore := keystest.Addresses{address}
		tm := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, txStore, nil, config, keystore, nil)
		var nonce uint64
		// Start
		client.On("PendingNonceAt", mock.Anything, address).Return(nonce, nil).Maybe()
		servicetest.Run(t, tm)
		tm.Trigger(address)
	})
}

func TestBroadcastTransaction(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	client := txm.NewMockClient(t)
	ab := txm.NewMockAttemptBuilder(t)
	config := txm.Config{}
	address := testutils.NewAddress()
	keystore := keystest.Addresses{}

	t.Run("fails if FetchUnconfirmedTransactionAtNonceWithCount for unconfirmed transactions fails", func(t *testing.T) {
		mTxStore := txm.NewMockTxStore(t)
		mTxStore.On("FetchUnconfirmedTransactionAtNonceWithCount", mock.Anything, mock.Anything, mock.Anything).Return(nil, 0, errors.New("call failed")).Once()
		tx := txm.NewTxm(logger.Test(t), testutils.FixtureChainID, client, ab, mTxStore, nil, config, keystore, nil)
		bo, err := tx.BroadcastTransaction(ctx, address)
		require.Error(t, err)
		assert.False(t, bo)
		require.ErrorContains(t, err, "call failed")
	})

	t.Run("throws a warning and returns if unconfirmed transactions exceed maxInFlightTransactions", func(t *testing.T) {
		lggr, observedLogs := logger.TestObserved(t, zap.DebugLevel)
		mTxStore := txm.NewMockTxStore(t)
		mTxStore.On("FetchUnconfirmedTransactionAtNonceWithCount", mock.Anything, mock.Anything, mock.Anything).Return(nil, txm.MaxInFlightTransactions+1, nil).Once()
		tx := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, mTxStore, nil, config, keystore, nil)
		bo, err := tx.BroadcastTransaction(ctx, address)
		assert.True(t, bo)
		require.NoError(t, err)
		tests.AssertLogEventually(t, observedLogs, "Reached transaction limit")
	})

	t.Run("checks pending nonce if unconfirmed transactions are equal or more than maxInFlightSubset", func(t *testing.T) {
		lggr, observedLogs := logger.TestObserved(t, zap.DebugLevel)
		mTxStore := txm.NewMockTxStore(t)
		tm := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, mTxStore, nil, config, keystore, nil)
		tm.SetNonce(address, 1)
		mTxStore.On("FetchUnconfirmedTransactionAtNonceWithCount", mock.Anything, mock.Anything, mock.Anything).Return(nil, txm.MaxInFlightSubset, nil).Twice()

		client.On("PendingNonceAt", mock.Anything, address).Return(uint64(0), nil).Once() // LocalNonce: 1, PendingNonce: 0
		bo, err := tm.BroadcastTransaction(ctx, address)
		assert.True(t, bo)
		require.NoError(t, err)

		client.On("PendingNonceAt", mock.Anything, address).Return(uint64(1), nil).Once() // LocalNonce: 1, PendingNonce: 1
		mTxStore.On("UpdateUnstartedTransactionWithNonce", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Once()
		bo, err = tm.BroadcastTransaction(ctx, address)
		assert.False(t, bo)
		require.NoError(t, err)
		tests.AssertLogCountEventually(t, observedLogs, "Reached transaction limit.", 1)
	})

	t.Run("fails if UpdateUnstartedTransactionWithNonce fails", func(t *testing.T) {
		mTxStore := txm.NewMockTxStore(t)
		mTxStore.On("FetchUnconfirmedTransactionAtNonceWithCount", mock.Anything, mock.Anything, mock.Anything).Return(nil, 0, nil).Once()
		tm := txm.NewTxm(logger.Test(t), testutils.FixtureChainID, client, ab, mTxStore, nil, config, keystore, nil)
		mTxStore.On("UpdateUnstartedTransactionWithNonce", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("call failed")).Once()
		bo, err := tm.BroadcastTransaction(ctx, address)
		assert.False(t, bo)
		require.Error(t, err)
		require.ErrorContains(t, err, "call failed")
	})

	t.Run("returns if there are no unstarted transactions", func(t *testing.T) {
		lggr := logger.Test(t)
		txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
		require.NoError(t, txStore.Add(address))
		tm := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, txStore, nil, config, keystore, nil)
		bo, err := tm.BroadcastTransaction(ctx, address)
		require.NoError(t, err)
		assert.False(t, bo)
		assert.Equal(t, uint64(0), tm.GetNonce(address))
	})

	t.Run("picks a new tx and creates a new attempt then sends it and updates the broadcast time", func(t *testing.T) {
		lggr := logger.Test(t)
		txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
		require.NoError(t, txStore.Add(address))
		tm := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, txStore, nil, config, keystore, nil)
		tm.SetNonce(address, 8)
		metrics, err := txm.NewTxmMetrics(testutils.FixtureChainID)
		require.NoError(t, err)
		tm.Metrics = metrics
		IDK := "IDK"
		txRequest := &types.TxRequest{
			Data:              []byte{100, 200},
			IdempotencyKey:    &IDK,
			ChainID:           testutils.FixtureChainID,
			FromAddress:       address,
			ToAddress:         testutils.NewAddress(),
			SpecifiedGasLimit: 22000,
		}
		tx, err := tm.CreateTransaction(t.Context(), txRequest)
		require.NoError(t, err)
		attempt := &types.Attempt{
			TxID:     tx.ID,
			Fee:      gas.EvmFee{GasPrice: assets.NewWeiI(1)},
			GasLimit: 22000,
		}
		ab.On("NewAgnosticBumpAttempt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(attempt, nil).Once()
		client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		bo, err := tm.BroadcastTransaction(ctx, address)
		require.NoError(t, err)
		assert.False(t, bo)
		assert.Equal(t, uint64(9), tm.GetNonce(address))
		tx, err = txStore.FindTxWithIdempotencyKey(t.Context(), IDK)
		require.NoError(t, err)
		assert.Len(t, tx.Attempts, 1)
		var zeroTime time.Time
		assert.Greater(t, *tx.LastBroadcastAt, zeroTime)
		assert.Greater(t, *tx.Attempts[0].BroadcastAt, zeroTime)
		assert.Greater(t, *tx.InitialBroadcastAt, zeroTime)
	})
}

func TestBackfillTransactions(t *testing.T) {
	t.Parallel()

	client := txm.NewMockClient(t)
	txStore := txm.NewMockTxStore(t)
	config := txm.Config{}
	address := testutils.NewAddress()
	keystore := keystest.Addresses{}

	t.Run("fails if latest nonce fetching fails", func(t *testing.T) {
		ab := txm.NewMockAttemptBuilder(t)
		txm := txm.NewTxm(logger.Test(t), testutils.FixtureChainID, client, ab, txStore, nil, config, keystore, nil)
		client.On("NonceAt", mock.Anything, address, mock.Anything).Return(uint64(0), errors.New("latest nonce fail")).Once()
		err := txm.BackfillTransactions(t.Context(), address)
		require.Error(t, err)
		require.ErrorContains(t, err, "latest nonce fail")
	})

	t.Run("fails if MarkConfirmedAndReorgedTransactions fails", func(t *testing.T) {
		ab := txm.NewMockAttemptBuilder(t)
		txm := txm.NewTxm(logger.Test(t), testutils.FixtureChainID, client, ab, txStore, nil, config, keystore, nil)
		client.On("NonceAt", mock.Anything, address, mock.Anything).Return(uint64(0), nil).Once()
		txStore.On("MarkConfirmedAndReorgedTransactions", mock.Anything, mock.Anything, address).
			Return([]*types.Transaction{}, []uint64{}, errors.New("marking transactions confirmed failed")).Once()
		err := txm.BackfillTransactions(t.Context(), address)
		require.Error(t, err)
		require.ErrorContains(t, err, "marking transactions confirmed failed")
	})

	t.Run("fills nonce gap", func(t *testing.T) {
		lggr, observedLogs := logger.TestObserved(t, zap.DebugLevel)
		txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
		require.NoError(t, txStore.Add(address))
		ab := txm.NewMockAttemptBuilder(t)
		c := txm.Config{EIP1559: false, BlockTime: 10 * time.Minute, RetryBlockThreshold: 10, EmptyTxLimitDefault: 22000}
		tm := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, txStore, nil, c, keystore, nil)
		emptyMetrics, err := txm.NewTxmMetrics(testutils.FixtureChainID)
		require.NoError(t, err)
		tm.Metrics = emptyMetrics

		// Add a new transaction that will be assigned with nonce = 1. Nonce = 0 is not being tracked by the txStore. This will trigger a nonce gap.
		txRequest := &types.TxRequest{
			ChainID:     testutils.FixtureChainID,
			FromAddress: address,
			ToAddress:   testutils.NewAddress(),
		}
		_, err = tm.CreateTransaction(t.Context(), txRequest)
		require.NoError(t, err)
		_, err = txStore.UpdateUnstartedTransactionWithNonce(t.Context(), address, 1) // Create nonce gap
		require.NoError(t, err)

		// During backfill we observe nonce has changed. The transaction with nonce = 1 should be marked unconfirmed.
		// For nonce = 0 there are no transactions stored in txStore, which results in a nonce gap.
		// TXM creates a new empty transaction and fills the gap.
		client.On("NonceAt", mock.Anything, address, mock.Anything).Return(uint64(0), nil).Once()
		attempt := &types.Attempt{
			TxID:     1,
			Fee:      gas.EvmFee{GasPrice: assets.NewWeiI(1)},
			GasLimit: 22000,
		}
		ab.On("NewAgnosticBumpAttempt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(attempt, nil).Once()
		client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		err = tm.BackfillTransactions(t.Context(), address)
		require.NoError(t, err)
		tests.AssertLogEventually(t, observedLogs, fmt.Sprintf("Nonce gap at nonce: %d - address: %v. Creating a new transaction", 0, address))
		_, count, err := txStore.FetchUnconfirmedTransactionAtNonceWithCount(t.Context(), 0, address)
		require.NoError(t, err)
		assert.Equal(t, 2, count)
	})

	t.Run("retries attempt after threshold", func(t *testing.T) {
		lggr, observedLogs := logger.TestObserved(t, zap.DebugLevel)
		txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
		require.NoError(t, txStore.Add(address))
		ab := txm.NewMockAttemptBuilder(t)
		c := txm.Config{EIP1559: false, BlockTime: 1 * time.Second, RetryBlockThreshold: 1, EmptyTxLimitDefault: 22000}
		tm := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, txStore, nil, c, keystore, nil)
		emptyMetrics, err := txm.NewTxmMetrics(testutils.FixtureChainID)
		require.NoError(t, err)
		tm.Metrics = emptyMetrics

		IDK := "IDK"
		txRequest := &types.TxRequest{
			Data:              []byte{100, 200},
			IdempotencyKey:    &IDK,
			ChainID:           testutils.FixtureChainID,
			FromAddress:       address,
			ToAddress:         testutils.NewAddress(),
			SpecifiedGasLimit: 22000,
		}
		tx, err := tm.CreateTransaction(t.Context(), txRequest)
		require.NoError(t, err)
		_, err = txStore.UpdateUnstartedTransactionWithNonce(t.Context(), address, 0)
		require.NoError(t, err)

		attempt := &types.Attempt{
			TxID:     tx.ID,
			Fee:      gas.EvmFee{GasPrice: assets.NewWeiI(1)},
			GasLimit: 22000,
		}
		ab.On("NewAgnosticBumpAttempt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(attempt, nil).Once()

		client.On("NonceAt", mock.Anything, address, mock.Anything).Return(uint64(0), nil).Once()
		client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		err = tm.BackfillTransactions(t.Context(), address)
		require.NoError(t, err)
		tests.AssertLogEventually(t, observedLogs, fmt.Sprintf("Rebroadcasting attempt for txID: %d", attempt.TxID))
	})

	t.Run("retries instantly if the attempt is purgeable", func(t *testing.T) {
		lggr, observedLogs := logger.TestObserved(t, zap.DebugLevel)
		txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
		require.NoError(t, txStore.Add(address))
		ab := txm.NewMockAttemptBuilder(t)
		c := txm.Config{EIP1559: false, BlockTime: 1 * time.Second, RetryBlockThreshold: 10, EmptyTxLimitDefault: 22000}
		tm := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, txStore, nil, c, keystore, nil)
		emptyMetrics, err := txm.NewTxmMetrics(testutils.FixtureChainID)
		require.NoError(t, err)
		tm.Metrics = emptyMetrics

		IDK := "IDK"
		txRequest := &types.TxRequest{
			Data:              []byte{100, 200},
			IdempotencyKey:    &IDK,
			ChainID:           testutils.FixtureChainID,
			FromAddress:       address,
			ToAddress:         testutils.NewAddress(),
			SpecifiedGasLimit: 22000,
		}
		_, err = tm.CreateTransaction(t.Context(), txRequest)
		require.NoError(t, err)
		tx, err := txStore.UpdateUnstartedTransactionWithNonce(t.Context(), address, 0)
		require.NoError(t, err)

		attempt := &types.Attempt{
			TxID:     tx.ID,
			Fee:      gas.EvmFee{GasPrice: assets.NewWeiI(1)},
			GasLimit: 22000,
			Hash:     testutils.NewHash(),
		}
		_, err = txStore.AppendAttemptToTransaction(t.Context(), *tx.Nonce, address, attempt)
		require.NoError(t, err)
		require.NoError(t, txStore.UpdateTransactionBroadcast(t.Context(), tx.ID, *tx.Nonce, attempt.Hash, address))
		require.NoError(t, txStore.MarkUnconfirmedTransactionPurgeable(t.Context(), *tx.Nonce, address))

		client.On("NonceAt", mock.Anything, address, mock.Anything).Return(uint64(0), nil).Once()
		ab.On("NewAgnosticBumpAttempt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(attempt, nil).Once()
		client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		err = tm.BackfillTransactions(t.Context(), address)
		require.NoError(t, err)
		tests.AssertLogEventually(t, observedLogs, fmt.Sprintf("Rebroadcasting attempt for txID: %d", attempt.TxID))

		// Broadcasted once an empty transaction but it didn't get confirmed, so we need to broadcast again.
		client.On("NonceAt", mock.Anything, address, mock.Anything).Return(uint64(0), nil).Once()
		ab.On("NewAgnosticBumpAttempt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(attempt, nil).Once()
		client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		err = tm.BackfillTransactions(t.Context(), address)
		require.NoError(t, err)
		tests.AssertLogEventually(t, observedLogs, fmt.Sprintf("Rebroadcasting attempt for txID: %d", attempt.TxID))
	})

	t.Run("fetches the unconfirmed transaction for a given nonce, throws a warning for max limit and retries with a new attempt", func(t *testing.T) {
		ctx := t.Context()
		lggr, observedLogs := logger.TestObserved(t, zap.DebugLevel)
		txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
		require.NoError(t, txStore.Add(address))
		ab := txm.NewMockAttemptBuilder(t)
		tm := txm.NewTxm(lggr, testutils.FixtureChainID, client, ab, txStore, nil, config, keystore, nil)
		var nonce uint64 = 8
		tm.SetNonce(address, nonce)
		metrics, err := txm.NewTxmMetrics(testutils.FixtureChainID)
		require.NoError(t, err)
		tm.Metrics = metrics
		IDK := "IDK"
		txRequest := &types.TxRequest{
			Data:              []byte{100, 200},
			IdempotencyKey:    &IDK,
			ChainID:           testutils.FixtureChainID,
			FromAddress:       address,
			ToAddress:         testutils.NewAddress(),
			SpecifiedGasLimit: 22000,
		}
		tx, err := tm.CreateTransaction(t.Context(), txRequest)
		require.NoError(t, err)
		_, err = txStore.UpdateUnstartedTransactionWithNonce(ctx, address, nonce)
		require.NoError(t, err)

		attempt := &types.Attempt{
			TxID:     tx.ID,
			Fee:      gas.EvmFee{GasPrice: assets.NewWeiI(1)},
			GasLimit: 22000,
		}
		var attemptsTried uint16 = storage.MaxAllowedAttempts + 2
		for range attemptsTried {
			_, err := txStore.AppendAttemptToTransaction(ctx, nonce, address, attempt)
			require.NoError(t, err)
		}
		client.On("NonceAt", mock.Anything, address, mock.Anything).Return(nonce, nil).Once()
		ab.On("NewAgnosticBumpAttempt", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(attempt, nil).Once()
		client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		err = tm.BackfillTransactions(ctx, address)
		require.NoError(t, err)

		tx2, err := txStore.FindTxWithIdempotencyKey(t.Context(), IDK)
		require.NoError(t, err)
		assert.Len(t, tx2.Attempts, storage.MaxAllowedAttempts)
		assert.Equal(t, attemptsTried+1, tx2.AttemptCount) // the initial attempts tried, plus the one during backfill
		var zeroTime time.Time
		assert.Greater(t, *tx2.LastBroadcastAt, zeroTime)
		assert.Greater(t, *tx2.Attempts[0].BroadcastAt, zeroTime)
		assert.Greater(t, *tx2.InitialBroadcastAt, zeroTime)
		assert.Equal(t, uint64(attemptsTried-storage.MaxAllowedAttempts+1), tx2.Attempts[0].ID) // the initial attempts tried, plus the one during backfill
		tests.AssertLogEventually(t, observedLogs, fmt.Sprintf("Reached max attempts threshold for txID: %d", 0))
	})
}

func TestFlow_ResendTransaction(t *testing.T) {
	t.Parallel()

	client := txm.NewMockClient(t)
	txStoreManager := storage.NewInMemoryStoreManager(logger.Test(t), testutils.FixtureChainID)
	address := testutils.NewAddress()
	require.NoError(t, txStoreManager.Add(address))
	config := txm.Config{EIP1559: true, EmptyTxLimitDefault: 22000, RetryBlockThreshold: 1, BlockTime: 2 * time.Second}
	mockEstimator := mocks.NewEvmFeeEstimator(t)
	defaultGasLimit := uint64(100000)
	keystore := &keystest.FakeChainStore{}
	attemptBuilder := txm.NewAttemptBuilder(func(address common.Address) *assets.Wei { return assets.NewWeiI(1) }, mockEstimator, keystore, 22000)
	stuckTxDetector := txm.NewStuckTxDetector(logger.Test(t), "", txm.StuckTxDetectorConfig{BlockTime: config.BlockTime, StuckTxBlockThreshold: uint32(config.RetryBlockThreshold + 1)})
	tm := txm.NewTxm(logger.Test(t), testutils.FixtureChainID, client, attemptBuilder, txStoreManager, stuckTxDetector, config, keystore, nil)
	metrics, err := txm.NewTxmMetrics(testutils.FixtureChainID)
	require.NoError(t, err)
	tm.Metrics = metrics
	initialNonce := uint64(0)
	tm.SetNonce(address, initialNonce)
	IDK := "IDK"

	// Create transaction
	_, err = tm.CreateTransaction(t.Context(), &types.TxRequest{
		IdempotencyKey: &IDK,
		ChainID:        testutils.FixtureChainID,
		FromAddress:    address,
		ToAddress:      testutils.NewAddress(),
	})
	require.NoError(t, err)

	// Broadcast transaction
	mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(5), GasFeeCap: assets.NewWeiI(10)}}, defaultGasLimit, nil).Once()
	client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	_, err = tm.BroadcastTransaction(t.Context(), address)
	require.NoError(t, err)

	// Backfill transaction
	client.On("NonceAt", mock.Anything, mock.Anything, mock.Anything).Return(initialNonce, nil).Maybe() // Transaction was not confirmed
	require.NoError(t, tm.BackfillTransactions(t.Context(), address))

	// Set LastBroadcastAt to a time in the past to trigger retry condition
	txStore := txStoreManager.InMemoryStoreMap[address]
	require.NotNil(t, txStore)
	tx := txStore.UnconfirmedTransactions[initialNonce]
	require.NotNil(t, tx)
	pastTime := time.Now().Add(-(config.BlockTime*time.Duration(config.RetryBlockThreshold) + 1*time.Second))
	tx.LastBroadcastAt = &pastTime

	// Retry with bumped fee
	client.On("NonceAt", mock.Anything, mock.Anything, mock.Anything).Return(initialNonce, nil).Maybe() // Transaction was not confirmed again
	mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(5), GasFeeCap: assets.NewWeiI(10)}}, defaultGasLimit, nil).Once()
	mockEstimator.On("BumpFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(6), GasFeeCap: assets.NewWeiI(12)}}, defaultGasLimit, nil).Once()
	client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	require.NoError(t, tm.BackfillTransactions(t.Context(), address)) // retry

	// Set LastBroadcastAt to a time in the past to trigger purge condition
	pastTime = time.Now().Add(-(config.BlockTime*time.Duration(config.RetryBlockThreshold) + 2*time.Second))
	tx.LastBroadcastAt = &pastTime

	// Purge transaction
	client.On("NonceAt", mock.Anything, mock.Anything, mock.Anything).Return(initialNonce, nil).Maybe() // Transaction was not confirmed again
	mockEstimator.On("GetMaxFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(5), GasFeeCap: assets.NewWeiI(10)}}, defaultGasLimit, nil).Once()
	client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	require.NoError(t, tm.BackfillTransactions(t.Context(), address)) // retry

	// Instant retransmission of purgeable transaction
	client.On("NonceAt", mock.Anything, mock.Anything, mock.Anything).Return(initialNonce, nil).Maybe() // Transaction was not confirmed again
	mockEstimator.On("GetMaxFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(5), GasFeeCap: assets.NewWeiI(10)}}, defaultGasLimit, nil).Once()
	client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	require.NoError(t, tm.BackfillTransactions(t.Context(), address)) // retry
}

func TestFlow_ErrorHandler(t *testing.T) {
	t.Parallel()

	client := txm.NewMockClient(t)
	txStoreManager := storage.NewInMemoryStoreManager(logger.Test(t), testutils.FixtureChainID)
	address := testutils.NewAddress()
	require.NoError(t, txStoreManager.Add(address))
	config := txm.Config{EIP1559: true, EmptyTxLimitDefault: 22000, RetryBlockThreshold: 0, BlockTime: 2 * time.Second}
	mockEstimator := mocks.NewEvmFeeEstimator(t)
	keystore := &keystest.FakeChainStore{}
	attemptBuilder := txm.NewAttemptBuilder(func(address common.Address) *assets.Wei { return assets.NewWeiI(1) }, mockEstimator, keystore, 22000)
	stuckTxDetector := txm.NewStuckTxDetector(logger.Test(t), "", txm.StuckTxDetectorConfig{BlockTime: config.BlockTime, StuckTxBlockThreshold: uint32(config.RetryBlockThreshold + 1)})
	errorHandler := dualbroadcast.NewErrorHandler()
	tm := txm.NewTxm(logger.Test(t), testutils.FixtureChainID, client, attemptBuilder, txStoreManager, stuckTxDetector, config, keystore, errorHandler)
	metrics, err := txm.NewTxmMetrics(testutils.FixtureChainID)
	require.NoError(t, err)
	tm.Metrics = metrics
	initialNonce := uint64(0)
	tm.SetNonce(address, initialNonce)
	defaultGasLimit := uint64(100000)

	// Create transaction
	IDK := "IDK"
	_, err = tm.CreateTransaction(t.Context(), &types.TxRequest{
		IdempotencyKey: &IDK,
		ChainID:        testutils.FixtureChainID,
		FromAddress:    address,
		ToAddress:      testutils.NewAddress(),
	})
	require.NoError(t, err)

	// Broadcast transaction
	mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(5), GasFeeCap: assets.NewWeiI(10)}}, defaultGasLimit, nil).Once()
	client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(dualbroadcast.ErrNoBids).Once()
	_, err = tm.BroadcastTransaction(t.Context(), address)
	require.Error(t, err)
	require.ErrorContains(t, err, "transaction with txID: 0 marked as fatal")

	// Create transaction 2
	IDK2 := "IDK2"
	_, err = tm.CreateTransaction(t.Context(), &types.TxRequest{
		IdempotencyKey: &IDK2,
		ChainID:        testutils.FixtureChainID,
		FromAddress:    address,
		ToAddress:      testutils.NewAddress(),
	})
	require.NoError(t, err)

	// Broadcast transaction successfully. First transaction is marked as fatal and removed from the store. Transaction 2 takes its nonce.
	mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(5), GasFeeCap: assets.NewWeiI(10)}}, defaultGasLimit, nil).Once()
	client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	_, err = tm.BroadcastTransaction(t.Context(), address)
	require.NoError(t, err)
	tx, count, err := txStoreManager.FetchUnconfirmedTransactionAtNonceWithCount(t.Context(), 0, address)
	require.NoError(t, err)
	require.Equal(t, 1, count)
	require.NotNil(t, IDK, tx.IdempotencyKey)

	// Retry but don't mark transaction as fatal if there is already an attempt.
	client.On("NonceAt", mock.Anything, mock.Anything, mock.Anything).Return(initialNonce, nil).Maybe() // Transaction was not confirmed again
	mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(5), GasFeeCap: assets.NewWeiI(10)}}, defaultGasLimit, nil).Once()
	mockEstimator.On("BumpFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(6), GasFeeCap: assets.NewWeiI(12)}}, defaultGasLimit, nil).Once()
	client.On("SendTransaction", mock.Anything, mock.Anything, mock.Anything).Return(dualbroadcast.ErrNoBids).Once()
	err = tm.BackfillTransactions(t.Context(), address) // retry
	require.Error(t, err)
	require.ErrorContains(t, err, dualbroadcast.ErrNoBids.Error())
	tx, count, err = txStoreManager.FetchUnconfirmedTransactionAtNonceWithCount(t.Context(), 0, address) // same transaction is still in the store
	require.NoError(t, err)
	require.Equal(t, 1, count)
	require.NotNil(t, IDK, tx.IdempotencyKey)
	require.Equal(t, uint16(2), tx.AttemptCount)
}
