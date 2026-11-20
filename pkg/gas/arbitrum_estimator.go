package gas

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"sync"
	"time"

	pkgerrors "github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"

	"github.com/smartcontractkit/chainlink-framework/chains/fees"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	evmconfig "github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/chaintype"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/rollups"
)

type ArbConfig interface {
	LimitMax() uint64
	BumpPercent() uint16
	BumpMin() *assets.Wei
	EIP1559DynamicFees() bool
	FeeHistory() evmconfig.FeeHistory
	BlockHistory() evmconfig.BlockHistory
}

// arbitrumEstimator wraps a price estimator and uses getPricesInArbGas() to add the L1 calldata cost to the gas limit.
// Prices are delegated to the wrapped estimator:
//   - SuggestedPriceEstimator for legacy transactions, on chains that process transactions in FIFO order.
//   - FeeHistoryEstimator for dynamic fee transactions, on chains that order transactions by priority fee.
//
// For legacy transactions the gas limit is adjusted in GetLegacyGas. For dynamic fee transactions GetDynamicFee doesn't
// return a gas limit, so evmFeeEstimator applies gasLimitWithL1Cost directly.
type arbitrumEstimator struct {
	services.StateMachine
	cfg ArbConfig

	EvmEstimator // *SuggestedPriceEstimator or *FeeHistoryEstimator

	pollPeriod time.Duration
	logger     logger.Logger

	getPricesInArbGasMu sync.RWMutex
	perL2Tx             uint32
	perL1CalldataUnit   uint32

	chForceRefetch chan (chan struct{})
	chInitialised  chan struct{}
	chStop         services.StopChan
	chDone         chan struct{}

	l1Oracle rollups.ArbL1GasOracle
}

const arbitrumPollPeriod = 2 * time.Second

// NewArbitrumEstimator returns an estimator that adds the L1 calldata cost to the gas limit. Pricing is delegated to
// FeeHistoryEstimator if EIP1559DynamicFees is enabled, or to SuggestedPriceEstimator otherwise. Both are created with
// the Arbitrum chain type, as the Arbitrum estimator implies Arbitrum's pricing semantics.
func NewArbitrumEstimator(lggr logger.Logger, cfg ArbConfig, ethClient FeeEstimatorClient, chainID *big.Int, l1Oracle rollups.ArbL1GasOracle) EvmEstimator {
	lggr = logger.Named(lggr, "ArbitrumEstimator")

	var priceEstimator EvmEstimator
	if cfg.EIP1559DynamicFees() {
		fhCfg := FeeHistoryEstimatorConfig{
			BumpPercent:      cfg.BumpPercent(),
			CacheTimeout:     cfg.FeeHistory().CacheTimeout(),
			EIP1559:          cfg.EIP1559DynamicFees(),
			BlockHistorySize: uint64(cfg.BlockHistory().BlockHistorySize()),
			RewardPercentile: float64(cfg.BlockHistory().TransactionPercentile()),
		}
		priceEstimator = NewFeeHistoryEstimator(lggr, ethClient, fhCfg, chainID, chaintype.ChainArbitrum, l1Oracle)
	} else {
		priceEstimator = NewSuggestedPriceEstimator(lggr, ethClient, cfg, chaintype.ChainArbitrum, l1Oracle)
	}

	return &arbitrumEstimator{
		cfg:            cfg,
		EvmEstimator:   priceEstimator,
		pollPeriod:     arbitrumPollPeriod,
		logger:         lggr,
		chForceRefetch: make(chan (chan struct{})),
		chInitialised:  make(chan struct{}),
		chStop:         make(chan struct{}),
		chDone:         make(chan struct{}),
		l1Oracle:       l1Oracle,
	}
}

func (a *arbitrumEstimator) Name() string {
	return a.logger.Name()
}

func (a *arbitrumEstimator) Start(ctx context.Context) error {
	return a.StartOnce("ArbitrumEstimator", func() error {
		if err := a.EvmEstimator.Start(ctx); err != nil {
			return pkgerrors.Wrap(err, "failed to start gas price estimator")
		}
		go a.run()
		<-a.chInitialised
		return nil
	})
}
func (a *arbitrumEstimator) Close() error {
	return a.StopOnce("ArbitrumEstimator", func() (err error) {
		close(a.chStop)
		err = pkgerrors.Wrap(a.EvmEstimator.Close(), "failed to stop gas price estimator")
		<-a.chDone
		return
	})
}

func (a *arbitrumEstimator) Ready() error { return a.StateMachine.Ready() }

func (a *arbitrumEstimator) HealthReport() map[string]error {
	hp := map[string]error{a.Name(): a.Healthy()}
	services.CopyHealth(hp, a.EvmEstimator.HealthReport())
	return hp
}

// GetLegacyGas estimates both the gas price and the gas limit.
//   - Price is delegated to the wrapped estimator, with a buffer added on top (see gasPriceWithBuffer).
//   - Limit is l2GasLimit plus the L1 calldata cost (see gasLimitWithL1Cost).
func (a *arbitrumEstimator) GetLegacyGas(ctx context.Context, calldata []byte, l2GasLimit uint64, maxGasPriceWei *assets.Wei, opts ...fees.Opt) (gasPrice *assets.Wei, chainSpecificGasLimit uint64, err error) {
	gasPrice, _, err = a.EvmEstimator.GetLegacyGas(ctx, calldata, l2GasLimit, maxGasPriceWei, opts...)
	if err != nil {
		return
	}
	gasPrice = a.gasPriceWithBuffer(gasPrice, maxGasPriceWei)
	chainSpecificGasLimit, err = a.gasLimitWithL1Cost(ctx, calldata, l2GasLimit, opts...)
	return
}

// gasLimitWithL1Cost adds the dynamic values perL2Tx and perL1CalldataUnit, provided by the getPricesInArbGas() method
// of the precompile contract at ArbGasInfoAddress, to l2GasLimit. perL2Tx is a constant amount of gas, and
// perL1CalldataUnit is multiplied by the length of the tx calldata. It returns an error if the estimator is not started
// or if the result exceeds LimitMax.
func (a *arbitrumEstimator) gasLimitWithL1Cost(ctx context.Context, calldata []byte, l2GasLimit uint64, opts ...fees.Opt) (chainSpecificGasLimit uint64, err error) {
	ok := a.IfStarted(func() {
		if slices.Contains(opts, fees.OptForceRefetch) {
			ch := make(chan struct{})
			select {
			case a.chForceRefetch <- ch:
			case <-a.chStop:
				err = pkgerrors.New("estimator stopped")
				return
			case <-ctx.Done():
				err = ctx.Err()
				return
			}
			select {
			case <-ch:
			case <-a.chStop:
				err = pkgerrors.New("estimator stopped")
				return
			case <-ctx.Done():
				err = ctx.Err()
				return
			}
		}
		perL2Tx, perL1CalldataUnit := a.getPricesInArbGas()
		chainSpecificGasLimit = l2GasLimit + uint64(perL2Tx) + uint64(len(calldata))*uint64(perL1CalldataUnit)
		a.logger.Debugw("gasLimitWithL1Cost", "l2GasLimit", l2GasLimit, "calldataLen", len(calldata), "perL2Tx", perL2Tx,
			"perL1CalldataUnit", perL1CalldataUnit, "chainSpecificGasLimit", chainSpecificGasLimit)
	})
	if !ok {
		return 0, pkgerrors.New("estimator is not started")
	}
	if err != nil {
		return 0, err
	}
	if limitMax := a.cfg.LimitMax(); chainSpecificGasLimit > limitMax {
		return 0, fmt.Errorf("estimated gas limit: %d is greater than the maximum gas limit configured: %d", chainSpecificGasLimit, limitMax)
	}
	return chainSpecificGasLimit, nil
}

// GetMaxLegacyGas returns the result of GetLegacyGas. Legacy pricing is used on chains that process transactions in
// FIFO order, so a higher price doesn't improve inclusion and a standard gas estimation has the same effect.
func (a *arbitrumEstimator) GetMaxLegacyGas(ctx context.Context, calldata []byte, gasLimit uint64, maxGasPriceWei *assets.Wei, opts ...fees.Opt) (gasPrice *assets.Wei, chainSpecificGasLimit uint64, err error) {
	return a.GetLegacyGas(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
}

// During network congestion Arbitrum's suggested gas price can be extremely volatile, making gas estimations less accurate. For any transaction, Arbitrum will only charge
// the block's base fee. If the base fee increases rapidly there is a chance the suggested gas price will fall under that value, resulting in a fee too low error.
// We use gasPriceWithBuffer to increase the estimated gas price by some percentage to avoid fee too low errors. Eventually, only the base fee will be paid, regardless of the price.
func (a *arbitrumEstimator) gasPriceWithBuffer(gasPrice *assets.Wei, maxGasPriceWei *assets.Wei) *assets.Wei {
	const gasPriceBufferPercentage = 90

	gasPrice = gasPrice.AddPercentage(gasPriceBufferPercentage)
	if gasPrice.Cmp(maxGasPriceWei) > 0 {
		a.logger.Warnw("Updated gasPrice with buffer is higher than the max gas price limit. Falling back to max gas price", "gasPriceWithBuffer", gasPrice, "maxGasPriceWei", maxGasPriceWei)
		gasPrice = maxGasPriceWei
	}
	a.logger.Debugw("gasPriceWithBuffer", "updatedGasPrice", gasPrice)
	return gasPrice
}

func (a *arbitrumEstimator) getPricesInArbGas() (perL2Tx uint32, perL1CalldataUnit uint32) {
	a.getPricesInArbGasMu.RLock()
	perL2Tx, perL1CalldataUnit = a.perL2Tx, a.perL1CalldataUnit
	a.getPricesInArbGasMu.RUnlock()
	return
}

func (a *arbitrumEstimator) run() {
	defer close(a.chDone)

	a.refreshPricesInArbGas()
	close(a.chInitialised)

	t := services.TickerConfig{
		Initial:   a.pollPeriod,
		JitterPct: services.DefaultJitter,
	}.NewTicker(a.pollPeriod)
	defer t.Stop()

	for {
		select {
		case <-a.chStop:
			return
		case ch := <-a.chForceRefetch:
			a.refreshPricesInArbGas()
			t.Reset()
			close(ch)
		case <-t.C:
			a.refreshPricesInArbGas()
		}
	}
}

// refreshPricesInArbGas calls getPricesInArbGas() and caches the refreshed prices.
func (a *arbitrumEstimator) refreshPricesInArbGas() {
	perL2Tx, perL1CalldataUnit, err := a.l1Oracle.GetPricesInArbGas()
	if err != nil {
		a.logger.Warnw("Failed to refresh prices", "err", err)
		return
	}

	a.logger.Debugw("refreshPricesInArbGas", "perL2Tx", perL2Tx, "perL2CalldataUnit", perL1CalldataUnit)

	a.getPricesInArbGasMu.Lock()
	a.perL2Tx = perL2Tx
	a.perL1CalldataUnit = perL1CalldataUnit
	a.getPricesInArbGasMu.Unlock()
}
