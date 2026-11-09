package txm

import (
	"errors"
	"math"
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

func TestNewAttemptBuilder_BackwardCompatibility(t *testing.T) {
	for _, boosted := range []bool{false, true} {
		estimator := mocks.NewEvmFeeEstimator(t)
		method := "GetFee"
		if boosted {
			method = "GetMaxFee"
		}
		estimator.On(method, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(gas.EvmFee{GasPrice: assets.NewWeiI(10)}, uint64(100), nil).Once()
		ab := NewAttemptBuilder(func(common.Address) *assets.Wei {
			return assets.NewWeiI(100)
		}, estimator, keystest.TxSigner(nil), 100, boosted)
		nonce := uint64(1)
		attempt, err := ab.NewAttempt(t.Context(), logger.Test(t), &types.Transaction{
			Nonce: &nonce, SpecifiedGasLimit: 100,
		}, false)
		require.NoError(t, err)
		assert.Equal(t, int64(10), attempt.Fee.GasPrice.Int64())
		assert.Equal(t, uint64(100), attempt.GasLimit)
	}
}

func TestAttemptBuilder_FeeBoostTipAboveMaxPrice(t *testing.T) {
	// Regression: estimator returned tip (~150 gwei) > fee cap (20 gwei = key max). Clamping the tip
	// used to drive the fee cap negative (20 + (20 - 150) gwei).
	gwei := func(n int64) *assets.Wei { return assets.GWei(n) }
	estimator := mocks.NewEvmFeeEstimator(t)
	estimator.On("GetMaxFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: gwei(150).Add(assets.NewWeiI(8)), GasFeeCap: gwei(20)}}, uint64(100), nil).Once()
	ab := NewAttemptBuilderWithConfig(func(common.Address) *assets.Wei { return gwei(20) }, estimator, keystest.TxSigner(nil),
		AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: true, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})
	nonce := uint64(1)
	tx := &types.Transaction{ID: 1, Nonce: &nonce, SpecifiedGasLimit: 100}
	attempt, err := ab.NewAttempt(t.Context(), logger.Test(t), tx, true)
	require.NoError(t, err)
	assert.Equal(t, gwei(20), attempt.Fee.GasTipCap)
	assert.Equal(t, gwei(20), attempt.Fee.GasFeeCap)
	assert.Equal(t, 1, attempt.SignedTransaction.GasFeeCap().Sign())
}

func TestAttemptBuilder_NilPriceDefault(t *testing.T) {
	estimator := mocks.NewEvmFeeEstimator(t)
	estimator.On("GetMaxFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(3), GasFeeCap: assets.NewWeiI(10)}}, uint64(100), nil).Once()
	ab := NewAttemptBuilderWithConfig(func(common.Address) *assets.Wei { return assets.NewWeiI(100) }, estimator, keystest.TxSigner(nil),
		AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: true, FeeMultiplier: 1})
	nonce := uint64(1)
	tx := &types.Transaction{ID: 1, Nonce: &nonce, SpecifiedGasLimit: 100}
	attempt, err := ab.NewAttempt(t.Context(), logger.Test(t), tx, true)
	require.NoError(t, err)
	assert.Equal(t, int64(3), attempt.Fee.GasTipCap.Int64())
	assert.Equal(t, int64(10), attempt.Fee.GasFeeCap.Int64())
}

func TestAttemptBuilder_FeeMultiplierPrecision(t *testing.T) {
	// 1.13 is not exactly representable as a binary float; float math produced 112999999999 wei.
	estimator := mocks.NewEvmFeeEstimator(t)
	estimator.On("GetFee", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(gas.EvmFee{GasPrice: assets.GWei(100)}, uint64(100), nil).Once()
	ab := NewAttemptBuilderWithConfig(func(common.Address) *assets.Wei { return assets.GWei(1000) }, estimator, keystest.TxSigner(nil),
		AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeMultiplier: 1.13})
	nonce := uint64(1)
	tx := &types.Transaction{ID: 1, Nonce: &nonce, SpecifiedGasLimit: 100}
	attempt, err := ab.NewAttempt(t.Context(), logger.Test(t), tx, false)
	require.NoError(t, err)
	assert.Equal(t, assets.GWei(113), attempt.Fee.GasPrice)
}

func TestAttemptBuilder_FeeMultiplier(t *testing.T) {
	for _, tc := range []struct {
		name                string
		multiplier          float64
		maxPrice            int64
		wantTip             int64
		wantCap             int64
		priceDefault        int64
		wantBoostTip        int64
		wantBoostCap        int64
		wantDynamicCap      int64
		wantBoostDynamicCap int64
	}{
		{"zero defaults to one", 0, 100, 3, 10, 0, 3, 10, 10, 10},
		{"negative defaults to one", -1, 100, 3, 10, 0, 3, 10, 10, 10},
		{"fractional below one defaults to one", 0.5, 100, 3, 10, 0, 3, 10, 10, 10},
		{"NaN defaults to one", math.NaN(), 100, 3, 10, 0, 3, 10, 10, 10},
		{"Inf defaults to one", math.Inf(1), 100, 3, 10, 0, 3, 10, 10, 10},
		{"zero with boosted floor", 0, 100, 3, 10, 15, 15, 15, 10, 22},
		{"unchanged", 1, 100, 3, 10, 0, 3, 10, 10, 10},
		{"double", 2, 100, 6, 20, 0, 6, 20, 13, 13},
		{"fractional", 1.5, 100, 5, 15, 0, 5, 15, 12, 12}, // 3 * 1.5 = 4.5 rounds to 5
		{"capped fee cap", 2, 12, 6, 12, 0, 6, 12, 12, 12},
		{"capped both", 2, 5, 5, 5, 0, 5, 5, 5, 5},
		{"floor without multiplication", 1, 100, 3, 10, 15, 15, 15, 10, 22},
		{"floor after multiplication", 2, 100, 6, 20, 15, 15, 20, 13, 22},
		{"multiplied fees above floor", 2, 100, 6, 20, 5, 6, 20, 13, 13},
		{"maximum overrides floor", 2, 12, 6, 12, 15, 12, 12, 12, 12},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, dynamic := range []bool{false, true} {
				for _, boosted := range []bool{false, true} {
					estimator := mocks.NewEvmFeeEstimator(t)
					fee := gas.EvmFee{GasPrice: assets.NewWeiI(10)}
					if dynamic {
						fee = gas.EvmFee{DynamicFee: gas.DynamicFee{GasTipCap: assets.NewWeiI(3), GasFeeCap: assets.NewWeiI(10)}}
					}
					method := "GetFee"
					if boosted {
						method = "GetMaxFee"
					}
					estimator.On(method, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
						Return(fee, uint64(100), nil).Once()
					ab := NewAttemptBuilderWithConfig(func(common.Address) *assets.Wei {
						return assets.NewWeiI(tc.maxPrice)
					}, estimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: boosted, FeeMultiplier: tc.multiplier, PriceDefault: assets.NewWeiI(tc.priceDefault)})
					nonce := uint64(1)
					tx := &types.Transaction{ID: 1, Nonce: &nonce, SpecifiedGasLimit: 100}
					attempt, err := ab.NewAttempt(t.Context(), logger.Test(t), tx, dynamic)
					require.NoError(t, err)
					assert.Equal(t, uint64(100), attempt.GasLimit)
					wantTip, wantCap := tc.wantTip, tc.wantCap
					if boosted {
						wantTip, wantCap = tc.wantBoostTip, tc.wantBoostCap
					}
					if dynamic {
						wantCap = tc.wantDynamicCap
						if boosted {
							wantCap = tc.wantBoostDynamicCap
						}
						assert.Equal(t, wantTip, attempt.Fee.GasTipCap.Int64())
						assert.Equal(t, wantCap, attempt.Fee.GasFeeCap.Int64())
						assert.Equal(t, wantTip, attempt.SignedTransaction.GasTipCap().Int64())
						assert.Equal(t, wantCap, attempt.SignedTransaction.GasFeeCap().Int64())
						assert.Equal(t, int64(3), fee.GasTipCap.Int64())
						assert.Equal(t, int64(10), fee.GasFeeCap.Int64())
					} else {
						assert.Equal(t, wantCap, attempt.Fee.GasPrice.Int64())
						assert.Equal(t, wantCap, attempt.SignedTransaction.GasPrice().Int64())
						assert.Equal(t, int64(10), fee.GasPrice.Int64())
					}

					estimator.On("BumpFee", mock.Anything, attempt.Fee, uint64(100), mock.Anything, mock.Anything).
						Return(attempt.Fee, uint64(100), nil).Once()
					bumped, err := ab.NewBumpAttempt(t.Context(), logger.Test(t), tx, *attempt)
					require.NoError(t, err)
					assert.Equal(t, attempt.Fee, bumped.Fee)
				}
			}
		})
	}
}

func TestAttemptBuilder_newLegacyAttempt(t *testing.T) {
	ab := NewAttemptBuilderWithConfig(nil, nil, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: false, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})
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
	ab := NewAttemptBuilderWithConfig(nil, nil, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: false, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})
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
	ab := NewAttemptBuilderWithConfig(priceMaxKey, mockEstimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: emptyGasLimit, FeeBoost: false, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})
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

	t.Run("uses SpecifiedGasLimit when feeBoost is enabled and tx is not purgeable", func(t *testing.T) {
		boostEstimator := mocks.NewEvmFeeEstimator(t)
		boostAb := NewAttemptBuilderWithConfig(priceMaxKey, boostEstimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: emptyGasLimit, FeeBoost: true, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})
		tx := &types.Transaction{ID: 10, FromAddress: address, Nonce: &nonce, SpecifiedGasLimit: specifiedGasLimit}
		boostEstimator.On("GetMaxFee", mock.Anything, mock.Anything, specifiedGasLimit, mock.Anything, mock.Anything, mock.Anything).
			Return(gas.EvmFee{GasPrice: assets.NewWeiI(100)}, specifiedGasLimit, nil).Once()
		a, err := boostAb.NewAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, specifiedGasLimit, a.GasLimit)
		boostEstimator.AssertExpectations(t)
	})

	t.Run("uses emptyTxLimitDefault when feeBoost is enabled and tx is purgeable", func(t *testing.T) {
		boostEstimator := mocks.NewEvmFeeEstimator(t)
		boostAb := NewAttemptBuilderWithConfig(priceMaxKey, boostEstimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: emptyGasLimit, FeeBoost: true, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})
		tx := &types.Transaction{ID: 10, FromAddress: address, IsPurgeable: true, Nonce: &nonce, SpecifiedGasLimit: specifiedGasLimit}
		boostEstimator.On("GetMaxFee", mock.Anything, mock.Anything, emptyGasLimit, mock.Anything, mock.Anything, mock.Anything).
			Return(gas.EvmFee{GasPrice: assets.NewWeiI(100)}, emptyGasLimit, nil).Once()
		a, err := boostAb.NewAttempt(t.Context(), lggr, tx, false)
		require.NoError(t, err)
		assert.Equal(t, emptyGasLimit, a.GasLimit)
		boostEstimator.AssertExpectations(t)
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
		ab := NewAttemptBuilderWithConfig(priceMaxKey, mockEstimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: false, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})

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
		ab := NewAttemptBuilderWithConfig(priceMaxKey, mockEstimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: false, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})

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
		ab := NewAttemptBuilderWithConfig(priceMaxKey, mockEstimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: false, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})

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
		ab := NewAttemptBuilderWithConfig(priceMaxKey, mockEstimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: false, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})

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
		ab := NewAttemptBuilderWithConfig(priceMaxKey, mockEstimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: false, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})

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
		ab := NewAttemptBuilderWithConfig(priceMaxKey, mockEstimator, keystest.TxSigner(nil), AttemptBuilderConfig{EmptyTxLimitDefault: 100, FeeBoost: false, FeeMultiplier: 1, PriceDefault: assets.NewWeiI(0)})

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
