package txm

import (
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/mocks"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys/keystest"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
	"github.com/smartcontractkit/chainlink-framework/chains/fees"
)

func TestAttemptBuilder_newLegacyAttempt(t *testing.T) {
	ab := NewAttemptBuilder(nil, nil, keystest.TxSigner(nil), 100)
	address := testutils.NewAddress()
	lggr := logger.Test(t)
	var gasLimit uint64 = 100

	t.Run("fails if GasPrice is nil", func(t *testing.T) {
		tx := &types.Transaction{ID: 10, FromAddress: address}
		_, err := ab.newCustomAttempt(t.Context(), tx, gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(1), GasFeeCap: assets.NewWeiI(2)}}, gasLimit, evmtypes.LegacyTxType, lggr)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "estimator did not return legacy fee")
	})

	t.Run("fails if tx doesn't have a nonce", func(t *testing.T) {
		tx := &types.Transaction{ID: 10, FromAddress: address}
		_, err := ab.newCustomAttempt(t.Context(), tx, gas.EvmFee{GasPrice: assets.NewWeiI(25)}, gasLimit, evmtypes.LegacyTxType, lggr)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "nonce empty")
	})

	t.Run("creates attempt with fields", func(t *testing.T) {
		var nonce uint64 = 77
		tx := &types.Transaction{ID: 10, FromAddress: address, Nonce: &nonce}
		a, err := ab.newCustomAttempt(t.Context(), tx, gas.EvmFee{GasPrice: assets.NewWeiI(25)}, gasLimit, evmtypes.LegacyTxType, lggr)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, a.TxID)
		assert.Equal(t, evmtypes.LegacyTxType, int(a.Type))
		assert.NotNil(t, a.Fee.GasPrice)
		assert.Equal(t, "25 wei", a.Fee.GasPrice.String())
		assert.Nil(t, a.Fee.GasTipCap)
		assert.Nil(t, a.Fee.GasFeeCap)
		assert.Equal(t, gasLimit, a.GasLimit)
	})
}

func TestAttemptBuilder_newDynamicFeeAttempt(t *testing.T) {
	ab := NewAttemptBuilder(nil, nil, keystest.TxSigner(nil), 100)
	address := testutils.NewAddress()

	lggr := logger.Test(t)
	var gasLimit uint64 = 100

	t.Run("fails if DynamicFee is invalid", func(t *testing.T) {
		tx := &types.Transaction{ID: 10, FromAddress: address}
		_, err := ab.newCustomAttempt(t.Context(), tx, gas.EvmFee{GasPrice: assets.NewWeiI(1)}, gasLimit, evmtypes.DynamicFeeTxType, lggr)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "estimator did not return dynamic fee")
	})

	t.Run("fails if tx doesn't have a nonce", func(t *testing.T) {
		tx := &types.Transaction{ID: 10, FromAddress: address}
		_, err := ab.newCustomAttempt(t.Context(), tx, gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(1), GasFeeCap: assets.NewWeiI(2)}}, gasLimit, evmtypes.DynamicFeeTxType, lggr)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "nonce empty")
	})

	t.Run("creates attempt with fields", func(t *testing.T) {
		var nonce uint64 = 77
		tx := &types.Transaction{ID: 10, FromAddress: address, Nonce: &nonce}

		a, err := ab.newCustomAttempt(t.Context(), tx, gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(1), GasFeeCap: assets.NewWeiI(2)}}, gasLimit, evmtypes.DynamicFeeTxType, lggr)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, a.TxID)
		assert.Equal(t, evmtypes.DynamicFeeTxType, int(a.Type))
		assert.Equal(t, "1 wei", a.Fee.DynamicFee.GasTipCap.String())
		assert.Equal(t, "2 wei", a.Fee.DynamicFee.GasFeeCap.String())
		assert.Nil(t, a.Fee.GasPrice)
		assert.Equal(t, gasLimit, a.GasLimit)
	})
}

func TestAttemptBuilder_NewAttempt(t *testing.T) {
	mockEstimator := mocks.NewEvmFeeEstimator(t)
	priceMaxKey := func(addr common.Address) *assets.Wei {
		return assets.NewWeiI(1000)
	}
	var nonce uint64 = 1
	var specifiedGasLimit uint64 = 200
	var emptyGasLimit uint64 = 100
	ab := NewAttemptBuilder(priceMaxKey, mockEstimator, keystest.TxSigner(nil), emptyGasLimit)
	address := testutils.NewAddress()
	lggr := logger.Test(t)

	t.Run("creates legacy attempt with fields", func(t *testing.T) {
		tx := &types.Transaction{ID: 10, FromAddress: address, Nonce: &nonce, SpecifiedGasLimit: specifiedGasLimit}
		mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(gas.EvmFee{GasPrice: assets.NewWeiI(100)}, specifiedGasLimit, nil).Once()
		a, err := ab.NewAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, a.TxID)
		assert.Equal(t, evmtypes.LegacyTxType, int(a.Type))
		assert.NotNil(t, a.Fee.GasPrice)
		assert.Equal(t, "100 wei", a.Fee.GasPrice.String())
		assert.Nil(t, a.Fee.GasTipCap)
		assert.Nil(t, a.Fee.GasFeeCap)
		assert.Equal(t, specifiedGasLimit, a.GasLimit)
	})

	t.Run("creates dynamic fee attempt with fields", func(t *testing.T) {
		tx := &types.Transaction{ID: 10, FromAddress: address, Nonce: &nonce, SpecifiedGasLimit: specifiedGasLimit}
		mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(1), GasFeeCap: assets.NewWeiI(2)}}, specifiedGasLimit, nil).Once()
		a, err := ab.NewAttempt(t.Context(), lggr, tx, true)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, a.TxID)
		assert.Equal(t, evmtypes.DynamicFeeTxType, int(a.Type))
	})

	t.Run("creates purgeable attempt with fields", func(t *testing.T) {
		tx := &types.Transaction{ID: 10, FromAddress: address, IsPurgeable: true, Nonce: &nonce, SpecifiedGasLimit: specifiedGasLimit}
		mockEstimator.On("GetMaxFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(gas.EvmFee{GasPrice: assets.NewWeiI(100)}, emptyGasLimit, nil).Once()
		a, err := ab.NewAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, a.TxID)
		assert.Equal(t, evmtypes.LegacyTxType, int(a.Type))
		assert.Equal(t, emptyGasLimit, a.GasLimit)
	})

	t.Run("creates dynamic fee purgeable attempt with fields", func(t *testing.T) {
		tx := &types.Transaction{ID: 10, FromAddress: address, IsPurgeable: true, Nonce: &nonce, SpecifiedGasLimit: specifiedGasLimit}
		mockEstimator.On("GetMaxFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(1), GasFeeCap: assets.NewWeiI(2)}}, emptyGasLimit, nil).Once()
		a, err := ab.NewAttempt(t.Context(), lggr, tx, true)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, a.TxID)
		assert.Equal(t, evmtypes.DynamicFeeTxType, int(a.Type))
		assert.Equal(t, emptyGasLimit, a.GasLimit)
	})

	t.Run("fails if estimator returns error", func(t *testing.T) {
		tx := &types.Transaction{ID: 10, FromAddress: address}
		mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(gas.EvmFee{}, uint64(0), errors.New("estimator error")).Once()
		_, err := ab.NewAttempt(t.Context(), lggr, tx, false)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "estimator error")
		mockEstimator.AssertExpectations(t)
	})
}

func TestAttemptBuilder_NewAgnosticBumpAttempt(t *testing.T) {
	address := testutils.NewAddress()
	lggr := logger.Test(t)
	var nonce uint64 = 77
	priceMaxKey := func(addr common.Address) *assets.Wei {
		return assets.NewWeiI(1000)
	}

	t.Run("returns original attempt when AttemptCount is 0", func(t *testing.T) {
		mockEstimator := mocks.NewEvmFeeEstimator(t)
		ab := NewAttemptBuilder(priceMaxKey, mockEstimator, keystest.TxSigner(nil), 100)

		tx := &types.Transaction{
			ID:           10,
			FromAddress:  address,
			Nonce:        &nonce,
			AttemptCount: 0,
		}

		gasPrice := assets.NewWeiI(100)
		initialFee := gas.EvmFee{GasPrice: gasPrice}
		mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(initialFee, uint64(21000), nil).Once()

		attempt, err := ab.NewAgnosticBumpAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, attempt.TxID)
		assert.Equal(t, gasPrice.String(), attempt.Fee.GasPrice.String())
		assert.Equal(t, evmtypes.LegacyTxType, int(attempt.Type))
		mockEstimator.AssertExpectations(t)
	})

	t.Run("bumps once when AttemptCount is 1", func(t *testing.T) {
		mockEstimator := mocks.NewEvmFeeEstimator(t)
		ab := NewAttemptBuilder(priceMaxKey, mockEstimator, keystest.TxSigner(nil), 100)

		tx := &types.Transaction{
			ID:           10,
			FromAddress:  address,
			Nonce:        &nonce,
			AttemptCount: 1,
		}

		gasPrice := assets.NewWeiI(100)
		initialFee := gas.EvmFee{GasPrice: gasPrice}
		bumpedFee := gas.EvmFee{GasPrice: gasPrice.Add(assets.NewWeiI(20))}
		mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(initialFee, uint64(21000), nil).Once()
		mockEstimator.On("BumpFee", mock.Anything, initialFee, mock.Anything, mock.Anything, mock.Anything).
			Return(bumpedFee, uint64(21000), nil).Once()

		attempt, err := ab.NewAgnosticBumpAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, attempt.TxID)
		assert.Equal(t, bumpedFee.GasPrice.String(), attempt.Fee.GasPrice.String())
		mockEstimator.AssertExpectations(t)
	})

	t.Run("bumps N times when AttemptCount is N", func(t *testing.T) {
		mockEstimator := mocks.NewEvmFeeEstimator(t)
		ab := NewAttemptBuilder(priceMaxKey, mockEstimator, keystest.TxSigner(nil), 100)

		tx := &types.Transaction{
			ID:           10,
			FromAddress:  address,
			Nonce:        &nonce,
			AttemptCount: 3,
		}

		initialFee := gas.EvmFee{GasPrice: assets.NewWeiI(100)}
		firstBump := gas.EvmFee{GasPrice: assets.NewWeiI(110)}
		secondBump := gas.EvmFee{GasPrice: assets.NewWeiI(121)}
		thirdBump := gas.EvmFee{GasPrice: assets.NewWeiI(133)}
		mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(initialFee, uint64(21000), nil).Once()
		mockEstimator.On("BumpFee", mock.Anything, initialFee, mock.Anything, mock.Anything, mock.Anything).
			Return(firstBump, uint64(21000), nil).Once()
		mockEstimator.On("BumpFee", mock.Anything, firstBump, mock.Anything, mock.Anything, mock.Anything).
			Return(secondBump, uint64(21000), nil).Once()
		mockEstimator.On("BumpFee", mock.Anything, secondBump, mock.Anything, mock.Anything, mock.Anything).
			Return(thirdBump, uint64(21000), nil).Once()

		attempt, err := ab.NewAgnosticBumpAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, attempt.TxID)
		assert.Equal(t, thirdBump.GasPrice.String(), attempt.Fee.GasPrice.String())
		mockEstimator.AssertExpectations(t)
	})

	t.Run("returns last valid attempt when BumpFee fails", func(t *testing.T) {
		mockEstimator := mocks.NewEvmFeeEstimator(t)
		ab := NewAttemptBuilder(priceMaxKey, mockEstimator, keystest.TxSigner(nil), 100)

		tx := &types.Transaction{
			ID:           10,
			FromAddress:  address,
			Nonce:        &nonce,
			AttemptCount: 3,
		}

		gasPrice := assets.NewWeiI(100)
		initialFee := gas.EvmFee{GasPrice: gasPrice}
		firstBump := gas.EvmFee{GasPrice: gasPrice.Add(assets.NewWeiI(20))}
		mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(initialFee, uint64(21000), nil).Once()
		mockEstimator.On("BumpFee", mock.Anything, initialFee, mock.Anything, mock.Anything, mock.Anything).
			Return(firstBump, uint64(21000), nil).Once()
		mockEstimator.On("BumpFee", mock.Anything, firstBump, mock.Anything, mock.Anything, mock.Anything).
			Return(gas.EvmFee{}, uint64(0), fees.ErrConnectivity).Once()

		attempt, err := ab.NewAgnosticBumpAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, attempt.TxID)
		// Should return the last valid bumped attempt
		assert.Equal(t, firstBump.GasPrice.String(), attempt.Fee.GasPrice.String())
		mockEstimator.AssertExpectations(t)
	})

	t.Run("caps bumps at maxBumpThreshold", func(t *testing.T) {
		mockEstimator := mocks.NewEvmFeeEstimator(t)
		ab := NewAttemptBuilder(priceMaxKey, mockEstimator, keystest.TxSigner(nil), 100)

		tx := &types.Transaction{
			ID:           10,
			FromAddress:  address,
			Nonce:        &nonce,
			AttemptCount: 10, // More than maxBumpThreshold (5)
		}

		initialFee := gas.EvmFee{GasPrice: assets.NewWeiI(100)}
		bumpedFee := gas.EvmFee{GasPrice: initialFee.GasPrice.Add(assets.NewWeiI(20))}
		mockEstimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(initialFee, uint64(21000), nil).Once()
		// Should only bump 5 times (maxBumpThreshold)
		mockEstimator.On("BumpFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(bumpedFee, uint64(21000), nil).Times(5)

		attempt, err := ab.NewAgnosticBumpAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, attempt.TxID)
		mockEstimator.AssertExpectations(t)
	})

	t.Run("returns max percentile attempt when transaction is purgeable", func(t *testing.T) {
		mockEstimator := mocks.NewEvmFeeEstimator(t)
		ab := NewAttemptBuilder(priceMaxKey, mockEstimator, keystest.TxSigner(nil), 100)

		tx := &types.Transaction{
			ID:           10,
			FromAddress:  address,
			IsPurgeable:  true,
			Nonce:        &nonce,
			AttemptCount: 10,
		}

		maxFee := gas.EvmFee{GasPrice: assets.NewWeiI(300)}
		mockEstimator.On("GetMaxFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(maxFee, uint64(21000), nil).Once()

		attempt, err := ab.NewAgnosticBumpAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, tx.ID, attempt.TxID)
		assert.Equal(t, maxFee.GasPrice.String(), attempt.Fee.GasPrice.String())
		mockEstimator.AssertExpectations(t)
	})
}
