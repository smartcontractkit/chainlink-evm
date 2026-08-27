package gas_test

import (
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/mocks"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-framework/chains/fees"
)

func TestFeeHistoryEstimatorLifecycle(t *testing.T) {
	t.Parallel()
	var gasLimit uint64 = 21000
	maxPrice := assets.NewWeiI(100)
	chainID := testutils.FixtureChainID

	t.Run("fails if you fetch gas price before the estimator starts", func(t *testing.T) {
		cfg := gas.FeeHistoryEstimatorConfig{
			BumpPercent:      20,
			RewardPercentile: 60,
			EIP1559:          false,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), nil, cfg, chainID, nil)
		_, _, err := u.GetLegacyGas(t.Context(), nil, gasLimit, maxPrice)
		assert.ErrorContains(t, err, "gas price not set")
	})

	t.Run("fails to start if BumpPercent is lower than the minimum cap", func(t *testing.T) {
		cfg := gas.FeeHistoryEstimatorConfig{BumpPercent: 9}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), nil, cfg, chainID, nil)
		assert.ErrorContains(t, u.Start(t.Context()), "BumpPercent")
	})

	t.Run("fails to start if RewardPercentile is higher than ConnectivityPercentile in EIP-1559", func(t *testing.T) {
		cfg := gas.FeeHistoryEstimatorConfig{
			BumpPercent:      20,
			RewardPercentile: 99,
			EIP1559:          true,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), nil, cfg, chainID, nil)
		assert.ErrorContains(t, u.Start(t.Context()), "RewardPercentile")
	})

	t.Run("fails to start if CacheTimeout is less than MinimumCacheTimeout", func(t *testing.T) {
		cfg := gas.FeeHistoryEstimatorConfig{
			BumpPercent:      20,
			RewardPercentile: 10,
			CacheTimeout:     200 * time.Millisecond,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), nil, cfg, chainID, nil)
		assert.ErrorContains(t, u.Start(t.Context()), "CacheTimeout")
	})

	t.Run("starts if configs are correct", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		client.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(10), nil).Maybe()

		cfg := gas.FeeHistoryEstimatorConfig{
			BumpPercent:      20,
			RewardPercentile: 10,
			CacheTimeout:     10 * time.Second,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		err := u.Start(t.Context())
		assert.NoError(t, err)
		err = u.Close()
		assert.NoError(t, err)
	})
}

func TestFeeHistoryEstimatorGetLegacyGas(t *testing.T) {
	t.Parallel()

	var gasLimit uint64 = 21000
	maxPrice := assets.NewWeiI(100)
	chainID := testutils.FixtureChainID

	t.Run("fetches a new gas price when first called", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		client.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(10), nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		_, err := u.RefreshGasPrice()
		assert.NoError(t, err)
		gasPrice, _, err := u.GetLegacyGas(t.Context(), nil, gasLimit, maxPrice)
		assert.NoError(t, err)
		assert.Equal(t, assets.NewWeiI(10), gasPrice)
	})

	t.Run("will return max price if estimation exceeds it", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		client.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(10), nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{}

		maxPrice := assets.NewWeiI(1)
		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		_, err := u.RefreshGasPrice()
		assert.NoError(t, err)
		gas1, _, err := u.GetLegacyGas(t.Context(), nil, gasLimit, maxPrice)
		assert.NoError(t, err)
		assert.Equal(t, maxPrice, gas1)
	})

	t.Run("fails if gas price has not been set yet", func(t *testing.T) {
		cfg := gas.FeeHistoryEstimatorConfig{}

		maxPrice := assets.NewWeiI(1)
		u := gas.NewFeeHistoryEstimator(logger.Test(t), nil, cfg, chainID, nil)
		_, _, err := u.GetLegacyGas(t.Context(), nil, gasLimit, maxPrice)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "gas price not set")
	})
}

func TestFeeHistoryEstimatorBumpLegacyGas(t *testing.T) {
	t.Parallel()

	var gasLimit uint64 = 21000
	maxPrice := assets.NewWeiI(100)
	chainID := testutils.FixtureChainID

	t.Run("bumps a previous attempt by BumpPercent", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		originalGasPrice := assets.NewWeiI(10)
		client.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(10), nil)

		cfg := gas.FeeHistoryEstimatorConfig{BumpPercent: 50, CacheTimeout: 5 * time.Second}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		servicetest.RunHealthy(t, u)
		gasPrice, _, err := u.BumpLegacyGas(t.Context(), originalGasPrice, gasLimit, maxPrice, nil)
		assert.NoError(t, err)
		assert.Equal(t, assets.NewWeiI(15), gasPrice)
	})

	t.Run("fails if the original attempt is nil, or equal or higher than the max price", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)

		cfg := gas.FeeHistoryEstimatorConfig{}
		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)

		var originalPrice *assets.Wei
		_, _, err := u.BumpLegacyGas(t.Context(), originalPrice, gasLimit, maxPrice, nil)
		assert.Error(t, err)

		originalPrice = assets.NewWeiI(100)
		_, _, err = u.BumpLegacyGas(t.Context(), originalPrice, gasLimit, maxPrice, nil)
		assert.Error(t, err)
	})

	t.Run("returns market gas price if bumped original fee is lower", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		client.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(80), nil).Once()
		originalGasPrice := assets.NewWeiI(10)

		cfg := gas.FeeHistoryEstimatorConfig{}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		gas, _, err := u.BumpLegacyGas(t.Context(), originalGasPrice, gasLimit, maxPrice, nil)
		assert.NoError(t, err)
		assert.Equal(t, assets.NewWeiI(80), gas)
	})

	t.Run("returns max gas price if bumped original fee is higher", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		client.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(1), nil).Once()
		originalGasPrice := assets.NewWeiI(10)

		cfg := gas.FeeHistoryEstimatorConfig{BumpPercent: 50}

		maxPrice := assets.NewWeiI(14)
		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		gas, _, err := u.BumpLegacyGas(t.Context(), originalGasPrice, gasLimit, maxPrice, nil)
		assert.NoError(t, err)
		assert.Equal(t, maxPrice, gas)
	})

	t.Run("returns max gas price if the aggregation of max and original bumped fee is higher", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		client.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(1), nil).Once()
		originalGasPrice := assets.NewWeiI(10)

		cfg := gas.FeeHistoryEstimatorConfig{BumpPercent: 50}

		maxPrice := assets.NewWeiI(14)
		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		gas, _, err := u.BumpLegacyGas(t.Context(), originalGasPrice, gasLimit, maxPrice, nil)
		assert.NoError(t, err)
		assert.Equal(t, maxPrice, gas)
	})

	t.Run("fails if the bumped gas price is lower than the minimum bump percentage", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		client.On("SuggestGasPrice", mock.Anything).Return(big.NewInt(100), nil).Once()
		originalGasPrice := assets.NewWeiI(100)

		cfg := gas.FeeHistoryEstimatorConfig{BumpPercent: 20}

		// Price will be capped by the max price
		maxPrice := assets.NewWeiI(101)
		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		_, _, err := u.BumpLegacyGas(t.Context(), originalGasPrice, gasLimit, maxPrice, nil)
		assert.Error(t, err)
	})
}

func TestFeeHistoryEstimatorGetDynamicFee(t *testing.T) {
	t.Parallel()

	maxPrice := assets.NewWeiI(100)
	chainID := testutils.FixtureChainID

	t.Run("fetches a new dynamic fee when first called", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		baseFee := big.NewInt(5)
		maxPriorityFeePerGas1 := big.NewInt(33)
		maxPriorityFeePerGas2 := big.NewInt(20)

		// first one represents market price and second one connectivity price
		// empty and nil entries are skipped
		reward := [][]*big.Int{
			{maxPriorityFeePerGas1, big.NewInt(5)},
			{maxPriorityFeePerGas2, big.NewInt(5)},
			{},
			{nil, big.NewInt(5)},
			{big.NewInt(5), nil},
		}
		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock:  big.NewInt(1),
			Reward:       reward,
			BaseFee:      []*big.Int{baseFee, baseFee},
			GasUsedRatio: nil,
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

		blockHistoryLength := 2
		cfg := gas.FeeHistoryEstimatorConfig{BlockHistorySize: uint64(blockHistoryLength)}
		avrgPriorityFee := big.NewInt(0)
		avrgPriorityFee.Add(maxPriorityFeePerGas1, maxPriorityFeePerGas2).Div(avrgPriorityFee, big.NewInt(int64(blockHistoryLength)))
		maxFee := (*assets.Wei)(baseFee).AddPercentage(gas.BaseFeeBufferPercentage).Add((*assets.Wei)(avrgPriorityFee))

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		err := u.RefreshDynamicFee()
		assert.NoError(t, err)
		dynamicFee, err := u.GetDynamicFee(t.Context(), maxPrice)
		assert.NoError(t, err)
		assert.Equal(t, maxFee, dynamicFee.GasFeeCap)
		assert.Equal(t, (*assets.Wei)(avrgPriorityFee), dynamicFee.GasTipCap)
	})

	t.Run("fails if dynamic fee has not been set yet", func(t *testing.T) {
		cfg := gas.FeeHistoryEstimatorConfig{}

		maxPrice := assets.NewWeiI(1)
		u := gas.NewFeeHistoryEstimator(logger.Test(t), nil, cfg, chainID, nil)
		_, err := u.GetDynamicFee(t.Context(), maxPrice)
		assert.Error(t, err)
		assert.ErrorContains(t, err, "dynamic fee not set")
	})

	t.Run("will return max price if tip cap or fee cap exceed it", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		baseFee := big.NewInt(1)
		maxPriorityFeePerGas := big.NewInt(3)
		maxPrice := assets.NewWeiI(2)

		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock:  big.NewInt(1),
			Reward:       [][]*big.Int{{maxPriorityFeePerGas, big.NewInt(5)}}, // first one represents market price and second one connectivity price
			BaseFee:      []*big.Int{baseFee},
			GasUsedRatio: nil,
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{BlockHistorySize: 1}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		err := u.RefreshDynamicFee()
		assert.NoError(t, err)
		dynamicFee, err := u.GetDynamicFee(t.Context(), maxPrice)
		assert.NoError(t, err)
		assert.Equal(t, maxPrice, dynamicFee.GasFeeCap)
		assert.Equal(t, maxPrice, dynamicFee.GasTipCap)
	})

	t.Run("applies the NoMempoolBaseFeeBufferPercentage when there is no mempool", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		baseFee := big.NewInt(100)
		maxPrice := assets.NewWeiI(1000)

		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock: big.NewInt(1),
			Reward:      [][]*big.Int{{big.NewInt(0), big.NewInt(0)}},
			BaseFee:     []*big.Int{baseFee},
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

		// BlockHistorySize of 0 signals a chain without a mempool, so buffer is increased.
		cfg := gas.FeeHistoryEstimatorConfig{BlockHistorySize: 0, EIP1559: true}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		require.NoError(t, u.RefreshDynamicFee())

		dynamicFee, err := u.GetDynamicFee(t.Context(), maxPrice)
		require.NoError(t, err)
		assert.Equal(t, assets.NewWei(baseFee).AddPercentage(gas.NoMempoolBaseFeeBufferPercentage), dynamicFee.GasFeeCap)
		assert.Equal(t, assets.NewWeiI(0), dynamicFee.GasTipCap)
		assert.Positive(t, dynamicFee.GasFeeCap.Cmp(assets.NewWei(baseFee).AddPercentage(gas.BaseFeeBufferPercentage)),
			"the no-mempool buffer must be wider than the default one")

		maxFee, err := u.GetMaxDynamicFee(maxPrice)
		require.NoError(t, err)
		assert.Equal(t, dynamicFee.GasFeeCap, maxFee.GasFeeCap)
	})
}

func TestFeeHistoryEstimatorGetMaxDynamicFee(t *testing.T) {
	t.Parallel()

	maxPrice := assets.NewWeiI(100)
	chainID := testutils.FixtureChainID

	t.Run("returns error when priorityFeeThreshold is not set", func(t *testing.T) {
		cfg := gas.FeeHistoryEstimatorConfig{
			BumpPercent:      20,
			RewardPercentile: 60,
			EIP1559:          true,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), nil, cfg, chainID, nil)
		fee, err := u.GetMaxDynamicFee(maxPrice)
		require.Error(t, err)
		assert.ErrorContains(t, err, "priorityFeeThreshold not set")
		assert.Equal(t, gas.DynamicFee{}, fee)
	})

	t.Run("returns max dynamic fee", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		baseFee := big.NewInt(5)
		marketMaxPriorityFeePerGas := big.NewInt(10)
		connectivityMaxPriorityFeePerGas := big.NewInt(20)

		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock:  big.NewInt(1),
			Reward:       [][]*big.Int{{marketMaxPriorityFeePerGas, connectivityMaxPriorityFeePerGas}}, // first one represents market price and second one connectivity price
			BaseFee:      []*big.Int{baseFee, baseFee},
			GasUsedRatio: nil,
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{
			BumpPercent:      20,
			RewardPercentile: 60,
			EIP1559:          true,
			BlockHistorySize: 2,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		err := u.RefreshDynamicFee()
		require.NoError(t, err)

		expectedPriorityFeeThreshold := assets.NewWei(connectivityMaxPriorityFeePerGas)
		expectedNextBaseFee := assets.NewWei(baseFee)
		expectedMaxFeeCap := expectedNextBaseFee.AddPercentage(gas.BaseFeeBufferPercentage).Add(expectedPriorityFeeThreshold)

		fee, err := u.GetMaxDynamicFee(maxPrice)
		require.NoError(t, err)
		assert.Equal(t, expectedMaxFeeCap, fee.GasFeeCap)
		assert.Equal(t, expectedPriorityFeeThreshold, fee.GasTipCap)
	})

	t.Run("caches the next base fee even when the rewards are all zero", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		baseFee := big.NewInt(5)
		zero := big.NewInt(0)
		marketMaxPriorityFeePerGas := big.NewInt(10)
		connectivityMaxPriorityFeePerGas := big.NewInt(20)

		// First round returns no usable priority fees, so only the base fee can be cached.
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ethereum.FeeHistory{
			OldestBlock: big.NewInt(1),
			Reward:      [][]*big.Int{{zero, zero}, {}},
			BaseFee:     []*big.Int{baseFee, baseFee},
		}, nil).Once()
		// Second round returns no base fee at all, which must fall back to the one cached above.
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ethereum.FeeHistory{
			OldestBlock: big.NewInt(2),
			Reward:      [][]*big.Int{{marketMaxPriorityFeePerGas, connectivityMaxPriorityFeePerGas}},
			BaseFee:     []*big.Int{},
		}, nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{
			BumpPercent:      20,
			RewardPercentile: 60,
			EIP1559:          true,
			BlockHistorySize: 2,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		require.NoError(t, u.RefreshDynamicFee())
		// The priority fees are still unset, but the base fee has been cached.
		_, err := u.GetMaxDynamicFee(maxPrice)
		require.ErrorContains(t, err, "priorityFeeThreshold not set")

		require.NoError(t, u.RefreshDynamicFee())
		expectedPriorityFeeThreshold := assets.NewWei(connectivityMaxPriorityFeePerGas)
		expectedMaxFeeCap := assets.NewWei(baseFee).AddPercentage(gas.BaseFeeBufferPercentage).Add(expectedPriorityFeeThreshold)

		fee, err := u.GetMaxDynamicFee(maxPrice)
		require.NoError(t, err)
		assert.Equal(t, expectedMaxFeeCap, fee.GasFeeCap)
		assert.Equal(t, expectedPriorityFeeThreshold, fee.GasTipCap)
	})

	t.Run("keeps the cached priority fees as a pair when a round has no usable average", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		maxPrice := assets.NewWeiI(1000)

		// First round is busy: both the average and the connectivity threshold are usable.
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ethereum.FeeHistory{
			OldestBlock: big.NewInt(1),
			Reward:      [][]*big.Int{{big.NewInt(100), big.NewInt(200)}},
			BaseFee:     []*big.Int{big.NewInt(10), big.NewInt(10)},
		}, nil).Once()
		// Second round is quiet: the RewardPercentile fee is 0 so no average can be derived, while the
		// ConnectivityPercentile one drops to 50. Applying only the latter would leave the cache with
		// maxPriorityFeePerGas(100) > priorityFeeThreshold(50) and halt bumping with a bogus ErrConnectivity.
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ethereum.FeeHistory{
			OldestBlock: big.NewInt(2),
			Reward:      [][]*big.Int{{big.NewInt(0), big.NewInt(50)}},
			BaseFee:     []*big.Int{big.NewInt(20), big.NewInt(20)},
		}, nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{
			BumpPercent:      20,
			RewardPercentile: 60,
			EIP1559:          true,
			BlockHistorySize: 2,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		require.NoError(t, u.RefreshDynamicFee())
		require.NoError(t, u.RefreshDynamicFee())

		// Both priority fees are retained from the first round, so the invariant still holds.
		fee, err := u.GetDynamicFee(t.Context(), maxPrice)
		require.NoError(t, err)
		assert.Equal(t, assets.NewWeiI(100), fee.GasTipCap)
		// The fee cap still tracked the second round's base fee: 20 * 1.4 + 100.
		assert.Equal(t, assets.NewWeiI(128), fee.GasFeeCap)

		maxFee, err := u.GetMaxDynamicFee(maxPrice)
		require.NoError(t, err)
		assert.Equal(t, assets.NewWeiI(200), maxFee.GasTipCap)
		assert.GreaterOrEqual(t, maxFee.GasTipCap.Cmp(fee.GasTipCap), 0, "the max tip cap must never fall below the market one")
	})

	t.Run("raises the threshold when the RPC returns non-monotonic percentiles", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		maxPrice := assets.NewWeiI(1000)

		// ConnectivityPercentile(50) below RewardPercentile(100) is impossible for percentiles of the same
		// sample, so the threshold is raised to the average instead of halting every bump.
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ethereum.FeeHistory{
			OldestBlock: big.NewInt(1),
			Reward:      [][]*big.Int{{big.NewInt(100), big.NewInt(50)}},
			BaseFee:     []*big.Int{big.NewInt(10), big.NewInt(10)},
		}, nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{
			BumpPercent:      20,
			RewardPercentile: 60,
			EIP1559:          true,
			BlockHistorySize: 2,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		require.NoError(t, u.RefreshDynamicFee())

		maxFee, err := u.GetMaxDynamicFee(maxPrice)
		require.NoError(t, err)
		assert.Equal(t, assets.NewWeiI(100), maxFee.GasTipCap)

		// Bumping up to the market tip is still allowed rather than tripping ErrConnectivity.
		originalFee := gas.DynamicFee{GasFeeCap: assets.NewWeiI(20), GasTipCap: assets.NewWeiI(10)}
		bumped, err := u.BumpDynamicFee(t.Context(), originalFee, maxPrice, nil)
		require.NoError(t, err)
		assert.Equal(t, assets.NewWeiI(100), bumped.GasTipCap)
	})
}

func TestFeeHistoryEstimatorBumpDynamicFee(t *testing.T) {
	t.Parallel()

	globalMaxPrice := assets.NewWeiI(100)
	chainID := testutils.FixtureChainID

	t.Run("bumps a previous attempt by BumpPercent", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		originalFee := gas.DynamicFee{
			GasFeeCap: assets.NewWeiI(20),
			GasTipCap: assets.NewWeiI(10),
		}

		// These values will be ignored because they are lower prices than the originalFee
		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock:  big.NewInt(1),
			Reward:       [][]*big.Int{{big.NewInt(5), big.NewInt(50)}}, // first one represents market price and second one connectivity price
			BaseFee:      []*big.Int{big.NewInt(5)},
			GasUsedRatio: nil,
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{
			BlockHistorySize: 2,
			BumpPercent:      50,
		}

		expectedFeeCap := originalFee.GasFeeCap.AddPercentage(cfg.BumpPercent)
		expectedTipCap := originalFee.GasTipCap.AddPercentage(cfg.BumpPercent)

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		err := u.RefreshDynamicFee()
		assert.NoError(t, err)
		dynamicFee, err := u.BumpDynamicFee(t.Context(), originalFee, globalMaxPrice, nil)
		assert.NoError(t, err)
		assert.Equal(t, expectedFeeCap, dynamicFee.GasFeeCap)
		assert.Equal(t, expectedTipCap, dynamicFee.GasTipCap)
	})

	t.Run("fails if the original attempt is invalid", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		maxPrice := assets.NewWeiI(20)
		cfg := gas.FeeHistoryEstimatorConfig{BlockHistorySize: 1}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		// nil original fee
		var originalFee gas.DynamicFee
		_, err := u.BumpDynamicFee(t.Context(), originalFee, maxPrice, nil)
		assert.Error(t, err)

		// tip cap is higher than fee cap
		originalFee = gas.DynamicFee{
			GasFeeCap: assets.NewWeiI(10),
			GasTipCap: assets.NewWeiI(11),
		}
		_, err = u.BumpDynamicFee(t.Context(), originalFee, maxPrice, nil)
		assert.Error(t, err)

		// fee cap is equal or higher to max price
		originalFee = gas.DynamicFee{
			GasFeeCap: assets.NewWeiI(20),
			GasTipCap: assets.NewWeiI(10),
		}
		_, err = u.BumpDynamicFee(t.Context(), originalFee, maxPrice, nil)
		assert.Error(t, err)
	})

	t.Run("returns market prices if bumped original fee is lower", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		originalFee := gas.DynamicFee{
			GasFeeCap: assets.NewWeiI(20),
			GasTipCap: assets.NewWeiI(10),
		}

		// Market fees
		baseFee := big.NewInt(5)
		maxPriorityFeePerGas := big.NewInt(33)
		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock:  big.NewInt(1),
			Reward:       [][]*big.Int{{maxPriorityFeePerGas, big.NewInt(100)}}, // first one represents market price and second one connectivity price
			BaseFee:      []*big.Int{baseFee},
			GasUsedRatio: nil,
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

		maxFee := (*assets.Wei)(baseFee).AddPercentage(gas.BaseFeeBufferPercentage).Add((*assets.Wei)(maxPriorityFeePerGas))

		cfg := gas.FeeHistoryEstimatorConfig{
			BlockHistorySize: 1,
			BumpPercent:      50,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		err := u.RefreshDynamicFee()
		assert.NoError(t, err)
		bumpedFee, err := u.BumpDynamicFee(t.Context(), originalFee, globalMaxPrice, nil)
		assert.NoError(t, err)
		assert.Equal(t, (*assets.Wei)(maxPriorityFeePerGas), bumpedFee.GasTipCap)
		assert.Equal(t, maxFee, bumpedFee.GasFeeCap)
	})

	t.Run("fails if connectivity percentile value is reached", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		// Bumping the original tip cap by BumpPercent lands on 36, above the market's connectivity percentile of 30.
		originalFee := gas.DynamicFee{
			GasFeeCap: assets.NewWeiI(40),
			GasTipCap: assets.NewWeiI(24),
		}

		// Market fees. The two percentiles come from the same sample, so the market price cannot exceed the
		// connectivity one; the threshold is reached by the bumped attempt, not by the market.
		baseFee := big.NewInt(5)
		maxPriorityFeePerGas := big.NewInt(20)
		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock:  big.NewInt(1),
			Reward:       [][]*big.Int{{maxPriorityFeePerGas, big.NewInt(30)}}, // first one represents market price and second one connectivity price
			BaseFee:      []*big.Int{baseFee},
			GasUsedRatio: nil,
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{
			BlockHistorySize: 1,
			BumpPercent:      50,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		err := u.RefreshDynamicFee()
		assert.NoError(t, err)
		_, err = u.BumpDynamicFee(t.Context(), originalFee, globalMaxPrice, nil)
		assert.Error(t, err)
		assert.True(t, fees.IsBumpErr(err))
		assert.ErrorIs(t, err, fees.ErrConnectivity)
	})

	t.Run("returns max price if the aggregation of max and original bumped fee is higher", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		originalFee := gas.DynamicFee{
			GasFeeCap: assets.NewWeiI(20),
			GasTipCap: assets.NewWeiI(18),
		}

		maxPrice := assets.NewWeiI(25)
		// Market fees
		baseFee := big.NewInt(1)
		maxPriorityFeePerGas := big.NewInt(1)
		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock:  big.NewInt(1),
			Reward:       [][]*big.Int{{maxPriorityFeePerGas, big.NewInt(30)}}, // first one represents market price and second one connectivity price
			BaseFee:      []*big.Int{baseFee},
			GasUsedRatio: nil,
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{
			BlockHistorySize: 1,
			BumpPercent:      50,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		err := u.RefreshDynamicFee()
		assert.NoError(t, err)
		bumpedFee, err := u.BumpDynamicFee(t.Context(), originalFee, maxPrice, nil)
		assert.NoError(t, err)
		assert.Equal(t, maxPrice, bumpedFee.GasTipCap)
		assert.Equal(t, maxPrice, bumpedFee.GasFeeCap)
	})

	t.Run("fails if the bumped gas price is lower than the minimum bump percentage", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		originalFee := gas.DynamicFee{
			GasFeeCap: assets.NewWeiI(20),
			GasTipCap: assets.NewWeiI(18),
		}

		maxPrice := assets.NewWeiI(21)
		// Market fees
		baseFee := big.NewInt(1)
		maxPriorityFeePerGas := big.NewInt(1)
		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock:  big.NewInt(1),
			Reward:       [][]*big.Int{{maxPriorityFeePerGas, big.NewInt(30)}}, // first one represents market price and second one connectivity price
			BaseFee:      []*big.Int{baseFee},
			GasUsedRatio: nil,
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil).Once()

		cfg := gas.FeeHistoryEstimatorConfig{
			BlockHistorySize: 1,
			BumpPercent:      50,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		err := u.RefreshDynamicFee()
		assert.NoError(t, err)
		_, err = u.BumpDynamicFee(t.Context(), originalFee, maxPrice, nil)
		assert.Error(t, err)
	})

	t.Run("ignores maxPriorityFeePerGas if there is no mempool and forces refetch", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		originalFee := gas.DynamicFee{
			GasFeeCap: assets.NewWeiI(40),
			GasTipCap: assets.NewWeiI(0),
		}

		// Market fees
		baseFee := big.NewInt(10)
		maxPriorityFeePerGas := big.NewInt(0)
		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock:  big.NewInt(1),
			Reward:       [][]*big.Int{{maxPriorityFeePerGas, big.NewInt(0)}}, // first one represents market price and second one connectivity price
			BaseFee:      []*big.Int{baseFee},
			GasUsedRatio: nil,
		}
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(feeHistoryResult, nil)

		cfg := gas.FeeHistoryEstimatorConfig{
			BlockHistorySize: 0,
			BumpPercent:      20,
			CacheTimeout:     10 * time.Second,
			EIP1559:          true,
		}

		// No mempool means the wider staleness buffer applies.
		maxFeePerGas := assets.NewWei(baseFee).AddPercentage(gas.NoMempoolBaseFeeBufferPercentage)
		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		servicetest.RunHealthy(t, u)
		bumpedFee, err := u.BumpDynamicFee(t.Context(), originalFee, globalMaxPrice, nil)
		assert.NoError(t, err)
		assert.Equal(t, (*assets.Wei)(maxPriorityFeePerGas), assets.NewWeiI(0))
		assert.Equal(t, maxFeePerGas, bumpedFee.GasFeeCap)
	})

	t.Run("fails if there is no mempool and the estimator is not started", func(t *testing.T) {
		cfg := gas.FeeHistoryEstimatorConfig{
			BlockHistorySize: 0,
			BumpPercent:      20,
			CacheTimeout:     10 * time.Second,
			EIP1559:          true,
		}

		// The client is never called because the estimator isn't started.
		u := gas.NewFeeHistoryEstimator(logger.Test(t), mocks.NewFeeHistoryEstimatorClient(t), cfg, chainID, nil)
		_, err := u.BumpDynamicFee(t.Context(), gas.DynamicFee{GasFeeCap: assets.NewWeiI(40), GasTipCap: assets.NewWeiI(0)}, globalMaxPrice, nil)
		require.ErrorContains(t, err, "estimator not started")
	})

	t.Run("propagates the refresh error if there is no mempool", func(t *testing.T) {
		client := mocks.NewFeeHistoryEstimatorClient(t)
		client.On("FeeHistory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("RPC unreachable"))

		cfg := gas.FeeHistoryEstimatorConfig{
			BlockHistorySize: 0,
			BumpPercent:      20,
			CacheTimeout:     10 * time.Second,
			EIP1559:          true,
		}

		u := gas.NewFeeHistoryEstimator(logger.Test(t), client, cfg, chainID, nil)
		servicetest.RunHealthy(t, u)
		_, err := u.BumpDynamicFee(t.Context(), gas.DynamicFee{GasFeeCap: assets.NewWeiI(40), GasTipCap: assets.NewWeiI(0)}, globalMaxPrice, nil)
		require.ErrorContains(t, err, "RPC unreachable")
	})
}
