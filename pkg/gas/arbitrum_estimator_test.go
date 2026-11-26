package gas_test

import (
	"bytes"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/chainlink-framework/chains/fees"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	evmconfig "github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/mocks"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/rollups"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
)

type arbFeeHistoryConfig struct {
	cacheTimeout time.Duration
}

func (a *arbFeeHistoryConfig) CacheTimeout() time.Duration {
	return a.cacheTimeout
}

type arbConfig struct {
	v               uint64
	bumpPercent     uint16
	bumpMin         *assets.Wei
	eip1559         bool
	feeHistory      arbFeeHistoryConfig
	blockHistory    gas.MockBlockHistoryConfig
	limitMultiplier float32
}

func (a *arbConfig) LimitMax() uint64 {
	return a.v
}

func (a *arbConfig) BumpPercent() uint16 {
	return a.bumpPercent
}

func (a *arbConfig) BumpMin() *assets.Wei {
	return a.bumpMin
}

func (a *arbConfig) EIP1559DynamicFees() bool {
	return a.eip1559
}

func (a *arbConfig) FeeHistory() evmconfig.FeeHistory {
	return &a.feeHistory
}

func (a *arbConfig) BlockHistory() evmconfig.BlockHistory {
	return &a.blockHistory
}

func (a *arbConfig) LimitMultiplier() float32 {
	if a.limitMultiplier == 0 {
		return 1.0
	}
	return a.limitMultiplier
}

func (a *arbConfig) EstimateLimit() bool {
	return false
}

func (a *arbConfig) SenderAddress() *evmtypes.EIP55Address {
	return nil
}

func TestArbitrumEstimator(t *testing.T) {
	t.Parallel()

	chainID := testutils.FixtureChainID
	maxGasPrice := assets.NewWeiI(100)
	const maxGasLimit uint64 = 500_000
	calldata := []byte{0x00, 0x00, 0x01, 0x02, 0x03}
	const gasLimit uint64 = 80000
	const gasPriceBufferPercentage = 90
	const bumpPercent = 10
	var bumpMin = assets.NewWei(big.NewInt(1))

	t.Run("calling GetLegacyGas on unstarted estimator returns error", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, feeEstimatorClient, chainID, l1Oracle)
		_, _, err = o.GetLegacyGas(tests.Context(t), calldata, gasLimit, maxGasPrice)
		assert.EqualError(t, err, "estimator is not started")
	})

	var zeros bytes.Buffer
	zeros.Write(common.BigToHash(big.NewInt(0)).Bytes())
	zeros.Write(common.BigToHash(big.NewInt(0)).Bytes())
	zeros.Write(common.BigToHash(big.NewInt(123455)).Bytes())
	t.Run("calling GetLegacyGas on started estimator returns estimates", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		feeEstimatorClient.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(42)
		})
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, rollups.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, rollups.ArbGasInfo_getPricesInArbGas, hex.EncodeToString(callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(zeros.Bytes(), nil)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{v: maxGasLimit, bumpPercent: bumpPercent, bumpMin: bumpMin}, feeEstimatorClient, chainID, l1Oracle)
		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(tests.Context(t), calldata, gasLimit, maxGasPrice)
		require.NoError(t, err)
		// Expected price for a standard l2_suggested_estimator would be 42, but we add a fixed gasPriceBufferPercentage.
		assert.Equal(t, assets.NewWeiI(42).AddPercentage(gasPriceBufferPercentage), gasPrice)
		assert.Equal(t, gasLimit, chainSpecificGasLimit)
	})

	t.Run("gas price is lower than user specified max gas price", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, feeEstimatorClient, chainID, l1Oracle)

		feeEstimatorClient.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(42)
		})
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, rollups.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, rollups.ArbGasInfo_getPricesInArbGas, hex.EncodeToString(callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(zeros.Bytes(), nil)

		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(tests.Context(t), calldata, gasLimit, assets.NewWeiI(40))
		require.Error(t, err)
		assert.EqualError(t, err, "estimated gas price: 42 wei is greater than the maximum gas price configured: 40 wei")
		assert.Nil(t, gasPrice)
		assert.Equal(t, uint64(0), chainSpecificGasLimit)
	})

	t.Run("gas price is lower than global max gas price", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, feeEstimatorClient, chainID, l1Oracle)

		feeEstimatorClient.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(120)
		})
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, rollups.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, rollups.ArbGasInfo_getPricesInArbGas, hex.EncodeToString(callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(zeros.Bytes(), nil)

		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(tests.Context(t), calldata, gasLimit, assets.NewWeiI(110))
		assert.EqualError(t, err, "estimated gas price: 120 wei is greater than the maximum gas price configured: 110 wei")
		assert.Nil(t, gasPrice)
		assert.Equal(t, uint64(0), chainSpecificGasLimit)
	})

	t.Run("calling BumpLegacyGas on unstarted arbitrum estimator returns error", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, feeEstimatorClient, chainID, l1Oracle)
		_, _, err = o.BumpLegacyGas(tests.Context(t), assets.NewWeiI(42), gasLimit, assets.NewWeiI(10), nil)
		assert.EqualError(t, err, "estimator is not started")
	})

	t.Run("calling GetLegacyGas on started estimator if initial call failed returns error", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, feeEstimatorClient, chainID, l1Oracle)

		feeEstimatorClient.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(pkgerrors.New("kaboom"))
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, rollups.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, rollups.ArbGasInfo_getPricesInArbGas, hex.EncodeToString(callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(zeros.Bytes(), nil)

		servicetest.RunHealthy(t, o)

		_, _, err = o.GetLegacyGas(tests.Context(t), calldata, gasLimit, maxGasPrice)
		assert.EqualError(t, err, "failed to estimate gas; gas price not set")
	})

	t.Run("calling GetDynamicFee always returns error", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, feeEstimatorClient, chainID, l1Oracle)
		_, err = o.GetDynamicFee(tests.Context(t), maxGasPrice)
		assert.EqualError(t, err, "dynamic fees are not implemented for this estimator")
	})

	t.Run("calling BumpDynamicFee always returns error", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{}, feeEstimatorClient, chainID, l1Oracle)
		fee := gas.DynamicFee{
			GasFeeCap: assets.NewWeiI(42),
			GasTipCap: assets.NewWeiI(5),
		}
		_, err = o.BumpDynamicFee(tests.Context(t), fee, maxGasPrice, nil)
		assert.EqualError(t, err, "dynamic fees are not implemented for this estimator")
	})

	t.Run("limit computes", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		feeEstimatorClient.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(42)
		})
		const (
			perL2Tx       = 50_000
			perL1Calldata = 10_000
		)
		var expLimit = gasLimit + perL2Tx + perL1Calldata*uint64(len(calldata))

		var b bytes.Buffer
		b.Write(common.BigToHash(big.NewInt(perL2Tx)).Bytes())
		b.Write(common.BigToHash(big.NewInt(perL1Calldata)).Bytes())
		b.Write(common.BigToHash(big.NewInt(123455)).Bytes())
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, rollups.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, rollups.ArbGasInfo_getPricesInArbGas, hex.EncodeToString(callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(b.Bytes(), nil)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{v: maxGasLimit, bumpPercent: bumpPercent, bumpMin: bumpMin}, feeEstimatorClient, chainID, l1Oracle)
		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(tests.Context(t), calldata, gasLimit, maxGasPrice)
		require.NoError(t, err)
		require.NotNil(t, gasPrice)
		// Again, a normal l2_suggested_estimator would return 42, but arbitrum_estimator adds a buffer.
		assert.Equal(t, "79 wei", gasPrice.String())
		assert.Equal(t, expLimit, chainSpecificGasLimit, "expected %d but got %d", expLimit, chainSpecificGasLimit)
	})

	t.Run("limit exceeds max", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		feeEstimatorClient.On("CallContext", mock.Anything, mock.Anything, "eth_gasPrice").Return(nil).Run(func(args mock.Arguments) {
			res := args.Get(1).(*hexutil.Big)
			(*big.Int)(res).SetInt64(42)
		})
		const (
			perL2Tx       = 500_000
			perL1Calldata = 100_000
		)

		var b bytes.Buffer
		b.Write(common.BigToHash(big.NewInt(perL2Tx)).Bytes())
		b.Write(common.BigToHash(big.NewInt(perL1Calldata)).Bytes())
		b.Write(common.BigToHash(big.NewInt(123455)).Bytes())
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), mock.IsType(&big.Int{})).Run(func(args mock.Arguments) {
			callMsg := args.Get(1).(ethereum.CallMsg)
			blockNumber := args.Get(2).(*big.Int)
			assert.Equal(t, rollups.ArbGasInfoAddress, callMsg.To.String())
			assert.Equal(t, rollups.ArbGasInfo_getPricesInArbGas, hex.EncodeToString(callMsg.Data))
			assert.Equal(t, big.NewInt(-1), blockNumber)
		}).Return(b.Bytes(), nil)

		o := gas.NewArbitrumEstimator(logger.Test(t), &arbConfig{v: maxGasLimit, bumpPercent: bumpPercent, bumpMin: bumpMin}, feeEstimatorClient, chainID, l1Oracle)
		servicetest.RunHealthy(t, o)
		gasPrice, chainSpecificGasLimit, err := o.GetLegacyGas(tests.Context(t), calldata, gasLimit, maxGasPrice)
		require.Error(t, err, "expected error but got (%s, %d)", gasPrice, chainSpecificGasLimit)
	})

	t.Run("EIP1559 dynamic fees use FeeHistoryEstimator and apply L1 gas limit", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		const (
			perL2Tx       = 50_000
			perL1Calldata = 10_000
		)
		expLimit := gasLimit + perL2Tx + perL1Calldata*uint64(len(calldata))

		var arbPrices bytes.Buffer
		arbPrices.Write(common.BigToHash(big.NewInt(perL2Tx)).Bytes())
		arbPrices.Write(common.BigToHash(big.NewInt(perL1Calldata)).Bytes())
		arbPrices.Write(common.BigToHash(big.NewInt(0)).Bytes())

		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), (*big.Int)(nil)).
			Return(common.BigToHash(big.NewInt(1)).Bytes(), nil).Maybe()
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), big.NewInt(-1)).
			Return(arbPrices.Bytes(), nil)

		baseFee := big.NewInt(100)
		priorityFee := big.NewInt(10)
		priorityThreshold := big.NewInt(25)
		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock: big.NewInt(1),
			Reward:      [][]*big.Int{{priorityFee, priorityThreshold}},
			BaseFee:     []*big.Int{baseFee},
		}
		feeEstimatorClient.On("FeeHistory", mock.Anything, uint64(5), (*big.Int)(nil), []float64{60, gas.ConnectivityPercentile}).
			Return(feeHistoryResult, nil)

		cfg := &arbConfig{
			v:            maxGasLimit,
			bumpPercent:  bumpPercent,
			bumpMin:      bumpMin,
			eip1559:      true,
			feeHistory:   arbFeeHistoryConfig{cacheTimeout: gas.MinimumCacheTimeout},
			blockHistory: gas.MockBlockHistoryConfig{BlockHistorySizeF: 5, TransactionPercentileF: 60},
		}

		arbEst := gas.NewArbitrumEstimator(logger.Test(t), cfg, feeEstimatorClient, chainID, l1Oracle)
		wrapped := gas.NewEvmFeeEstimator(logger.Test(t), arbEst, true, cfg, feeEstimatorClient)
		servicetest.RunHealthy(t, wrapped)

		maxFeePrice := assets.NewWeiI(1000)
		var fee gas.EvmFee
		var estLimit uint64
		require.Eventually(t, func() bool {
			fee, estLimit, err = wrapped.GetFee(t.Context(), calldata, gasLimit, maxFeePrice, nil, nil)
			return err == nil
		}, 5*time.Second, 20*time.Millisecond)

		expectedFeeCap := assets.NewWei(baseFee).AddPercentage(gas.NoMempoolBaseFeeBufferPercentage).Add(assets.NewWei(priorityFee))
		assert.Equal(t, expectedFeeCap, fee.GasFeeCap)
		assert.Equal(t, assets.NewWei(priorityFee), fee.GasTipCap)
		assert.Nil(t, fee.GasPrice)
		assert.Equal(t, expLimit, estLimit)

		maxFee, maxEstLimit, err := wrapped.GetMaxFee(t.Context(), calldata, gasLimit, maxFeePrice, nil, nil)
		require.NoError(t, err)
		expectedMaxFeeCap := assets.NewWei(baseFee).AddPercentage(gas.NoMempoolBaseFeeBufferPercentage).Add(assets.NewWei(priorityThreshold))
		assert.Equal(t, expectedMaxFeeCap, maxFee.GasFeeCap)
		assert.Equal(t, assets.NewWei(priorityThreshold), maxFee.GasTipCap)
		assert.Equal(t, expLimit, maxEstLimit)
	})

	t.Run("EIP1559 dynamic fees return error when L1-adjusted limit exceeds LimitMax", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		var arbPrices bytes.Buffer
		arbPrices.Write(common.BigToHash(big.NewInt(500_000)).Bytes())
		arbPrices.Write(common.BigToHash(big.NewInt(100_000)).Bytes())
		arbPrices.Write(common.BigToHash(big.NewInt(0)).Bytes())

		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), (*big.Int)(nil)).
			Return(common.BigToHash(big.NewInt(1)).Bytes(), nil).Maybe()
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), big.NewInt(-1)).
			Return(arbPrices.Bytes(), nil)

		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock: big.NewInt(1),
			Reward:      [][]*big.Int{{big.NewInt(10), big.NewInt(25)}},
			BaseFee:     []*big.Int{big.NewInt(100)},
		}
		feeEstimatorClient.On("FeeHistory", mock.Anything, uint64(5), (*big.Int)(nil), []float64{60, gas.ConnectivityPercentile}).
			Return(feeHistoryResult, nil)

		cfg := &arbConfig{
			v:            maxGasLimit,
			bumpPercent:  bumpPercent,
			bumpMin:      bumpMin,
			eip1559:      true,
			feeHistory:   arbFeeHistoryConfig{cacheTimeout: gas.MinimumCacheTimeout},
			blockHistory: gas.MockBlockHistoryConfig{BlockHistorySizeF: 5, TransactionPercentileF: 60},
		}

		arbEst := gas.NewArbitrumEstimator(logger.Test(t), cfg, feeEstimatorClient, chainID, l1Oracle)
		wrapped := gas.NewEvmFeeEstimator(logger.Test(t), arbEst, true, cfg, feeEstimatorClient)
		servicetest.RunHealthy(t, wrapped)

		maxFeePrice := assets.NewWeiI(1000)
		require.Eventually(t, func() bool {
			_, _, err = wrapped.GetFee(t.Context(), calldata, gasLimit, maxFeePrice, nil, nil)
			return err != nil && err.Error() != "dynamic fee not set"
		}, 5*time.Second, 20*time.Millisecond)
		assert.ErrorContains(t, err, "is greater than the maximum gas limit configured")
	})

	t.Run("EIP1559 dynamic fees support OptForceRefetch for L1 gas limit", func(t *testing.T) {
		feeEstimatorClient := mocks.NewFeeEstimatorClient(t)
		l1Oracle, err := rollups.NewArbitrumL1GasOracle(logger.Test(t), feeEstimatorClient)
		require.NoError(t, err)

		var initialPrices bytes.Buffer
		initialPrices.Write(common.BigToHash(big.NewInt(10_000)).Bytes())
		initialPrices.Write(common.BigToHash(big.NewInt(1_000)).Bytes())
		initialPrices.Write(common.BigToHash(big.NewInt(0)).Bytes())

		var refreshedPrices bytes.Buffer
		refreshedPrices.Write(common.BigToHash(big.NewInt(20_000)).Bytes())
		refreshedPrices.Write(common.BigToHash(big.NewInt(2_000)).Bytes())
		refreshedPrices.Write(common.BigToHash(big.NewInt(0)).Bytes())

		// First call is from arbitrumL1Oracle.Start() (getL1BaseFeeEstimate, blockNumber == nil),
		// subsequent calls with blockNumber == -1 are getPricesInArbGas().
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), (*big.Int)(nil)).
			Return(common.BigToHash(big.NewInt(1)).Bytes(), nil).Maybe()
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), big.NewInt(-1)).
			Return(initialPrices.Bytes(), nil).Once()
		feeEstimatorClient.On("CallContract", mock.Anything, mock.IsType(ethereum.CallMsg{}), big.NewInt(-1)).
			Return(refreshedPrices.Bytes(), nil)

		feeHistoryResult := &ethereum.FeeHistory{
			OldestBlock: big.NewInt(1),
			Reward:      [][]*big.Int{{big.NewInt(10), big.NewInt(25)}},
			BaseFee:     []*big.Int{big.NewInt(100)},
		}
		feeEstimatorClient.On("FeeHistory", mock.Anything, uint64(5), (*big.Int)(nil), []float64{60, gas.ConnectivityPercentile}).
			Return(feeHistoryResult, nil)

		cfg := &arbConfig{
			v:            maxGasLimit,
			bumpPercent:  bumpPercent,
			bumpMin:      bumpMin,
			eip1559:      true,
			feeHistory:   arbFeeHistoryConfig{cacheTimeout: gas.MinimumCacheTimeout},
			blockHistory: gas.MockBlockHistoryConfig{BlockHistorySizeF: 5, TransactionPercentileF: 60},
		}

		arbEst := gas.NewArbitrumEstimator(logger.Test(t), cfg, feeEstimatorClient, chainID, l1Oracle)
		wrapped := gas.NewEvmFeeEstimator(logger.Test(t), arbEst, true, cfg, feeEstimatorClient)
		servicetest.RunHealthy(t, wrapped)

		maxFeePrice := assets.NewWeiI(1000)
		var estLimit uint64
		require.Eventually(t, func() bool {
			_, estLimit, err = wrapped.GetFee(t.Context(), calldata, gasLimit, maxFeePrice, nil, nil)
			return err == nil
		}, 5*time.Second, 20*time.Millisecond)
		assert.Equal(t, gasLimit+10_000+1_000*uint64(len(calldata)), estLimit)

		_, estLimit, err = wrapped.GetFee(t.Context(), calldata, gasLimit, maxFeePrice, nil, nil, fees.OptForceRefetch)
		require.NoError(t, err)
		assert.Equal(t, gasLimit+20_000+2_000*uint64(len(calldata)), estLimit)
	})
}
