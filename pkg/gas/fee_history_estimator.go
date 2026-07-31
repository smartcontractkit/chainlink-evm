package gas

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	bigmath "github.com/smartcontractkit/chainlink-common/pkg/utils/big_math"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/rollups"
	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
	"github.com/smartcontractkit/chainlink-framework/chains/fees"
)

const (
	// MinimumBumpPercentage is the minimum percentage that a bumped fee must be above the original fee.
	//This is based on geth's spec, which requires a minimum of 10% bump for bumped transactions to be accepted by the RPC.
	MinimumBumpPercentage = 10
	// ConnectivityPercentile is the highest percentile of the maxPriorityFeePerGas that we're willing to pay.
	// If a bumped attempt exceeds this value then there is a good chance there is a connectivity issue and we shouldn't bump.
	ConnectivityPercentile = 85
	// BaseFeeBufferPercentage covers the base fee growth until the transaction gets included. It also helps to cover potential
	// delays from the time the prices are fetched until the transaction is transmitted. On Ethereum the base fee moves by
	// 12.5% * (gasUsed - target) / target per block. This value covers ~8 blocks at u = 0.67, which is a demand spike rather than regular traffic.
	BaseFeeBufferPercentage = 40
	// NoMempoolBaseFeeBufferPercentage covers potential cache staleness and delays from the time the prices are fetched until transmission.
	// Chains without a mempool sequence transactions almost immediately, so the exposure is the CacheTimeout window
	// (especially on super fast chains like Arbitrum) rather than the mempool wait.
	// We avoid using an even larger buffer because even though the excess fee is refunded, it still increases the pre-transactional cost
	// and can cause the transaction to be rejected with an insufficient funds error.
	NoMempoolBaseFeeBufferPercentage = 90
	MinimumCacheTimeout              = 500 * time.Millisecond
)

type FeeHistoryEstimatorConfig struct {
	BumpPercent  uint16
	CacheTimeout time.Duration

	EIP1559          bool
	BlockHistorySize uint64
	RewardPercentile float64
}

// dynamicFeeCache holds the cached results of a single RefreshDynamicPrice round.
type dynamicFeeCache struct {
	dynamicFee           DynamicFee
	priorityFeeThreshold *assets.Wei
	nextBaseFee          *assets.Wei
}

// getDynamicFee returns the cached dynamic price, or an error if no refresh has populated it yet.
func (c dynamicFeeCache) getDynamicFee() (fee DynamicFee, err error) {
	if c.dynamicFee.GasFeeCap == nil || c.dynamicFee.GasTipCap == nil {
		return fee, errors.New("dynamic fee not set")
	}
	return c.dynamicFee, nil
}

// getPriorityFeeThreshold returns the cached connectivity threshold, or an error if no refresh has populated it yet.
func (c dynamicFeeCache) getPriorityFeeThreshold() (*assets.Wei, error) {
	if c.priorityFeeThreshold == nil {
		return nil, errors.New("priorityFeeThreshold not set")
	}
	return c.priorityFeeThreshold, nil
}

type feeHistoryEstimatorClient interface {
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	FeeHistory(ctx context.Context, blockCount uint64, lastBlock *big.Int, rewardPercentiles []float64) (feeHistory *ethereum.FeeHistory, err error)
}

type FeeHistoryEstimator struct {
	services.StateMachine

	client  feeHistoryEstimatorClient
	logger  logger.Logger
	config  FeeHistoryEstimatorConfig
	chainID *big.Int

	gasPriceMu sync.RWMutex
	gasPrice   *assets.Wei

	dynamicMu       sync.RWMutex
	dynamicFeeCache dynamicFeeCache

	l1Oracle rollups.L1Oracle

	metrics *feeHistoryEstimatorMetrics

	wg        *sync.WaitGroup
	stopCh    services.StopChan
	refreshCh chan struct{}
}

func NewFeeHistoryEstimator(lggr logger.Logger, client feeHistoryEstimatorClient, cfg FeeHistoryEstimatorConfig, chainID *big.Int, l1Oracle rollups.L1Oracle) *FeeHistoryEstimator {
	l := logger.Sugared(logger.Named(lggr, "FeeHistoryEstimator"))
	return &FeeHistoryEstimator{
		client:    client,
		logger:    l,
		config:    cfg,
		chainID:   chainID,
		l1Oracle:  l1Oracle,
		metrics:   newFeeHistoryEstimatorMetrics(l, chainID),
		wg:        new(sync.WaitGroup),
		stopCh:    make(chan struct{}),
		refreshCh: make(chan struct{}, 1),
	}
}

func (f *FeeHistoryEstimator) Start(context.Context) error {
	return f.StartOnce("FeeHistoryEstimator", func() error {
		if f.config.BumpPercent < MinimumBumpPercentage {
			return fmt.Errorf("BumpPercent: %s is less than minimum allowed percentage: %s",
				strconv.FormatUint(uint64(f.config.BumpPercent), 10), strconv.Itoa(MinimumBumpPercentage))
		}
		if f.config.EIP1559 && f.config.RewardPercentile > ConnectivityPercentile {
			return fmt.Errorf("RewardPercentile: %s is greater than maximum allowed percentile: %s",
				strconv.FormatUint(uint64(f.config.RewardPercentile), 10), strconv.Itoa(ConnectivityPercentile))
		}
		if f.config.CacheTimeout < MinimumCacheTimeout {
			return fmt.Errorf("CacheTimeout: %s must be at least %s", f.config.CacheTimeout, MinimumCacheTimeout)
		}
		f.wg.Add(1)
		go f.run()

		f.logger.Infof("Started FeeHistoryEstimator")
		return nil
	})
}

func (f *FeeHistoryEstimator) Close() error {
	return f.StopOnce("FeeHistoryEstimator", func() error {
		close(f.stopCh)
		f.wg.Wait()
		return nil
	})
}

func (f *FeeHistoryEstimator) run() {
	defer f.wg.Done()

	t := services.TickerConfig{
		JitterPct: services.DefaultJitter,
	}.NewTicker(f.config.CacheTimeout)

	for {
		select {
		case <-f.stopCh:
			return
		case <-f.refreshCh:
			t.Reset()
		case <-t.C:
			if f.config.EIP1559 {
				if err := f.RefreshDynamicFee(); err != nil {
					f.logger.Error(err)
				}
			} else {
				if _, err := f.RefreshGasPrice(); err != nil {
					f.logger.Error(err)
				}
			}
		}
	}
}

// GetLegacyGas will fetch the cached gas price value.
func (f *FeeHistoryEstimator) GetLegacyGas(ctx context.Context, _ []byte, gasLimit uint64, maxPrice *assets.Wei, opts ...fees.Opt) (gasPrice *assets.Wei, chainSpecificGasLimit uint64, err error) {
	chainSpecificGasLimit = gasLimit
	if gasPrice, err = f.getGasPrice(); err != nil {
		return
	}

	if gasPrice.Cmp(maxPrice) > 0 {
		f.logger.Warnf("estimated gas price: %s is greater than the maximum gas price configured: %s, returning the maximum price instead.", gasPrice, maxPrice)
		return maxPrice, chainSpecificGasLimit, nil
	}
	return
}

// GetMaxLegacyGas is not supported for FeeHistoryEstimator because eth_gasPrice fetches a single value gas price, so there is no way to increase the priority via this method.
// If there is a need to support this, we can always return the result of GetLegacyGas but that wouldn't increase the priority.
func (f *FeeHistoryEstimator) GetMaxLegacyGas(_ context.Context, _ []byte, gasLimit uint64, maxGasPriceWei *assets.Wei, _ ...fees.Opt) (gasPrice *assets.Wei, chainSpecificGasLimit uint64, err error) {
	return nil, 0, errors.New("max legacy gas is not supported for FeeHistoryEstimator")
}

// RefreshGasPrice will use eth_gasPrice to fetch and cache the latest gas price from the RPC.
func (f *FeeHistoryEstimator) RefreshGasPrice() (*assets.Wei, error) {
	ctx, cancel := f.stopCh.CtxWithTimeout(client.QueryTimeout)
	defer cancel()

	gasPrice, err := f.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	f.metrics.RecordGasPrice(ctx, float64(gasPrice.Int64()))

	gasPriceWei := assets.NewWei(gasPrice)

	f.logger.Debugf("Fetched new gas price: %v", gasPriceWei)

	f.gasPriceMu.Lock()
	defer f.gasPriceMu.Unlock()
	f.gasPrice = gasPriceWei
	return f.gasPrice, nil
}

func (f *FeeHistoryEstimator) getGasPrice() (*assets.Wei, error) {
	f.gasPriceMu.RLock()
	defer f.gasPriceMu.RUnlock()
	if f.gasPrice == nil {
		return f.gasPrice, errors.New("gas price not set")
	}
	return f.gasPrice, nil
}

// GetDynamicFee will fetch the cached dynamic prices.
func (f *FeeHistoryEstimator) GetDynamicFee(ctx context.Context, maxPrice *assets.Wei) (fee DynamicFee, err error) {
	if fee, err = f.getDynamicFeeCache().getDynamicFee(); err != nil {
		return
	}

	if fee.GasFeeCap.Cmp(maxPrice) > 0 {
		f.logger.Warnf("estimated maxFeePerGas: %v is greater than the maximum price configured: %v, returning the maximum price instead.",
			fee.GasFeeCap, maxPrice)
		fee.GasFeeCap = maxPrice
		if fee.GasTipCap.Cmp(maxPrice) > 0 {
			f.logger.Warnf("estimated maxPriorityFeePerGas: %v is greater than the maximum price configured: %v, returning the maximum price instead.",
				fee.GasTipCap, maxPrice)
			fee.GasTipCap = maxPrice
		}
	}
	return
}

func (f *FeeHistoryEstimator) GetMaxDynamicFee(maxPrice *assets.Wei) (fee DynamicFee, err error) {
	cache := f.getDynamicFeeCache()
	priorityFeeThreshold, err := cache.getPriorityFeeThreshold()
	if err != nil {
		return fee, err
	}
	if cache.nextBaseFee == nil {
		return fee, errors.New("nextBaseFee not set")
	}
	maxFeeCap := cache.nextBaseFee.AddPercentage(f.baseFeeBufferPercentage()).Add(priorityFeeThreshold)
	return DynamicFee{GasFeeCap: assets.WeiMin(maxFeeCap, maxPrice), GasTipCap: assets.WeiMin(priorityFeeThreshold, maxPrice)}, nil
}

// RefreshDynamicFee uses eth_feeHistory to fetch the baseFee of the next block and the Nth maxPriorityFeePerGas percentiles
// of the past X blocks. It also fetches the highest 85th maxPriorityFeePerGas percentile of the past X blocks, which represents
// the highest percentile we're willing to pay. A buffer is added on top of the latest baseFee to catch fluctuations in the next
// blocks. On Ethereum the increase is baseFee * 1.125 per block, however in some chains that may vary. See
// baseFeeBufferPercentage for how the buffer is picked.
func (f *FeeHistoryEstimator) RefreshDynamicFee() error {
	ctx, cancel := f.stopCh.CtxWithTimeout(client.QueryTimeout)
	defer cancel()

	// RewardPercentile will be used for maxPriorityFeePerGas estimations and connectivityPercentile to set the highest threshold for bumping.
	feeHistory, err := f.client.FeeHistory(ctx, max(f.config.BlockHistorySize, 1), nil, []float64{f.config.RewardPercentile, ConnectivityPercentile})
	if err != nil {
		return err
	}
	if feeHistory == nil {
		return errors.New("feeHistory result is nil")
	}

	// Start from the currently cached values so that partial responses only overwrite what they actually carry.
	cache := f.getDynamicFeeCache()

	var nextBlock *big.Int
	// If the BaseFee list is empty, maintain the cached base fee to continue updating the priority fee threshold.
	if len(feeHistory.BaseFee) != 0 {
		// eth_feeHistory doesn't return the latest baseFee of the range but rather the latest + 1, because it can be derived from the existing
		// values. Source: https://github.com/ethereum/go-ethereum/blob/b0f66e34ca2a4ea7ae23475224451c8c9a569826/eth/gasprice/feehistory.go#L235
		cache.nextBaseFee = assets.NewWei(feeHistory.BaseFee[len(feeHistory.BaseFee)-1])
		nextBlock = new(big.Int).Add(feeHistory.OldestBlock, big.NewInt(int64(len(feeHistory.BaseFee)-1)))
	}
	// If the cached base fee is nil and the BaseFee list is empty, return an error since a proper next base fee isn't set
	if cache.nextBaseFee == nil {
		return errors.New("nextBaseFee not set")
	}

	// If BlockHistorySize is 0 it means priority fees will be ignored from the calculations, so we set them to 0.
	var maxPriorityFeePerGas, priorityFeeThreshold *assets.Wei
	if f.config.BlockHistorySize == 0 {
		maxPriorityFeePerGas, priorityFeeThreshold = assets.NewWeiI(0), assets.NewWeiI(0)
	} else {
		maxPriorityFeePerGas, priorityFeeThreshold = calcPriorityFees(feeHistory.Reward)
	}

	// The two priority fees are always cached as a pair. priorityFeeThreshold is a higher percentile of the very same
	// sample as maxPriorityFeePerGas, so within a response maxPriorityFeePerGas <= priorityFeeThreshold always holds.
	// Mixing values from different rounds could invert that and make BumpDynamicFee report a bogus connectivity issue.
	if maxPriorityFeePerGas != nil && priorityFeeThreshold != nil {
		if maxPriorityFeePerGas.Cmp(priorityFeeThreshold) > 0 {
			// Only reachable if the RPC returns non-monotonic percentiles. Raise the threshold to keep bumping possible.
			f.logger.Warnw("eth_feeHistory returned a maxPriorityFeePerGas above its ConnectivityPercentile, raising the threshold to match",
				"maxPriorityFeePerGas", maxPriorityFeePerGas, "priorityFeeThreshold", priorityFeeThreshold)
			priorityFeeThreshold = maxPriorityFeePerGas
		}
		cache.dynamicFee.GasTipCap = maxPriorityFeePerGas
		cache.priorityFeeThreshold = priorityFeeThreshold
		f.metrics.RecordMaxPriorityFeePerGas(ctx, float64(maxPriorityFeePerGas.Int64()))
	} else {
		f.logger.Warnw("eth_feeHistory returned no usable priority fees, keeping the previously cached ones",
			"blocks", len(feeHistory.Reward), "maxPriorityFeePerGas", maxPriorityFeePerGas, "priorityFeeThreshold", priorityFeeThreshold)
	}

	// The fee cap is recomputed even when the priority fees above were retained, because the base fee moves much
	// faster than they do. It can only be set once a priority fee has been cached at least once, so that the
	// dynamic fee is never published half-populated.
	if cache.dynamicFee.GasTipCap != nil {
		cache.dynamicFee.GasFeeCap = cache.nextBaseFee.AddPercentage(f.baseFeeBufferPercentage()).Add(cache.dynamicFee.GasTipCap)
		f.metrics.RecordMaxFeePerGas(ctx, float64(cache.dynamicFee.GasFeeCap.Int64()))
	}

	f.metrics.RecordBaseFee(ctx, float64(cache.nextBaseFee.Int64()))

	f.logger.Debugf("Fetched new dynamic prices, nextBlock#: %v - oldestBlock#: %v - nextBaseFee: %v - maxFeePerGas: %v - maxPriorityFeePerGas: %v - maxPriorityFeeThreshold: %v",
		nextBlock, feeHistory.OldestBlock, cache.nextBaseFee, cache.dynamicFee.GasFeeCap, cache.dynamicFee.GasTipCap, cache.priorityFeeThreshold)

	f.dynamicMu.Lock()
	defer f.dynamicMu.Unlock()
	f.dynamicFeeCache = cache
	return nil
}

// calcPriorityFees returns the average of the non-zero RewardPercentile priority fees and the maximum
// ConnectivityPercentile priority fee of the given eth_feeHistory rewards. Either value is nil when the
// response carries no usable data for it.
//
// Zero priced priority fees are excluded, even though some networks allow them, because eth_feeHistory returns
// 0 values for empty blocks and discarding them yields a more representative sample.
func calcPriorityFees(rewards [][]*big.Int) (avg *assets.Wei, threshold *assets.Wei) {
	var nonZeroRewardsLen int64
	priorityFeeSum := big.NewInt(0)
	priorityFeeThreshold := big.NewInt(0)
	for _, reward := range rewards {
		// reward needs to have values for two percentiles. Some chains may return an empty slice instead of 0x0 values, so we
		// skip it instead of throwing an error.
		if len(reward) < 2 || reward[0] == nil || reward[1] == nil {
			continue
		}
		// We'll calculate the average of non-zero priority fees
		if reward[0].Sign() > 0 {
			priorityFeeSum.Add(priorityFeeSum, reward[0])
			nonZeroRewardsLen++
		}
		// We take the max value for the bumping threshold
		if reward[1].Sign() > 0 {
			priorityFeeThreshold = bigmath.Max(priorityFeeThreshold, reward[1])
		}
	}

	if nonZeroRewardsLen > 0 {
		avg = assets.NewWei(priorityFeeSum.Div(priorityFeeSum, big.NewInt(nonZeroRewardsLen)))
	}
	if priorityFeeThreshold.Sign() > 0 {
		threshold = assets.NewWei(priorityFeeThreshold)
	}
	return avg, threshold
}

// baseFeeBufferPercentage returns the safety buffer applied on top of the latest base fee. BlockHistorySize of 0 is used
// as the signal for a chain without a mempool, which is the only signal the estimator has for it.
func (f *FeeHistoryEstimator) baseFeeBufferPercentage() uint16 {
	if f.config.BlockHistorySize == 0 {
		return NoMempoolBaseFeeBufferPercentage
	}
	return BaseFeeBufferPercentage
}

func (f *FeeHistoryEstimator) getDynamicFeeCache() dynamicFeeCache {
	f.dynamicMu.RLock()
	defer f.dynamicMu.RUnlock()
	return f.dynamicFeeCache
}

// BumpLegacyGas provides a bumped gas price value by bumping the previous one by BumpPercent.
// If the original value is higher than the max price it returns an error as there is no room for bumping.
// It aggregates the market, bumped, and max gas price to provide a correct value.
func (f *FeeHistoryEstimator) BumpLegacyGas(ctx context.Context, originalGasPrice *assets.Wei, gasLimit uint64, maxPrice *assets.Wei, _ []EvmPriorAttempt) (*assets.Wei, uint64, error) {
	// Sanitize original fee input
	if originalGasPrice == nil || originalGasPrice.Cmp(maxPrice) >= 0 {
		return nil, 0, fmt.Errorf("%w: error while retrieving original gas price: originalGasPrice: %s. Maximum price configured: %s",
			fees.ErrBump, originalGasPrice, maxPrice)
	}

	currentGasPrice, err := f.RefreshGasPrice()
	if err != nil {
		return nil, 0, err
	}
	f.IfStarted(func() { f.signalNonBlockingRefresh() })

	bumpedGasPrice := originalGasPrice.AddPercentage(f.config.BumpPercent)
	bumpedGasPrice, err = LimitBumpedFee(originalGasPrice, currentGasPrice, bumpedGasPrice, maxPrice)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to limit gas price: %w", err)
	}

	f.logger.Debugw("bumped gas price", "originalGasPrice", originalGasPrice, "marketGasPrice", currentGasPrice, "bumpedGasPrice", bumpedGasPrice)

	return bumpedGasPrice, gasLimit, nil
}

// BumpDynamicFee provides a bumped dynamic fee by bumping the previous one by BumpPercent.
// If the original values are higher than the max price it returns an error as there is no room for bumping. If maxPriorityFeePerGas is bumped
// above the priority fee threshold then there is a good chance there is a connectivity issue and we shouldn't bump.
// Both maxFeePerGas as well as maxPriorityFeePerGas need to be bumped otherwise the RPC won't accept the transaction and throw an error.
// See: https://github.com/ethereum/go-ethereum/issues/24284
// It aggregates the market, bumped, and max price to provide a correct value, for both maxFeePerGas as well as maxPriorityFerPergas.
//
// If BlockHistorySize is 0 the chain has no mempool, so there is no concept of gas bumping and no priority fee to bump: the
// originalFee is ignored and the freshly refreshed market fee is returned instead.
func (f *FeeHistoryEstimator) BumpDynamicFee(ctx context.Context, originalFee DynamicFee, maxPrice *assets.Wei, _ []EvmPriorAttempt) (bumped DynamicFee, err error) {
	if f.config.BlockHistorySize == 0 {
		if err = f.refreshDynamicFeeNow(); err != nil {
			return bumped, err
		}
		return f.GetDynamicFee(ctx, maxPrice)
	}

	// Sanitize original fee input
	// According to geth's spec we need to bump both maxFeePerGas and maxPriorityFeePerGas for the new attempt to be accepted by the RPC
	if originalFee.GasFeeCap == nil ||
		originalFee.GasTipCap == nil ||
		((originalFee.GasTipCap.Cmp(originalFee.GasFeeCap)) > 0) ||
		(originalFee.GasFeeCap.Cmp(maxPrice) >= 0) {
		return bumped, fmt.Errorf("%w: error while retrieving original dynamic fees: (originalFeePerGas: %s - originalPriorityFeePerGas: %s). Maximum price configured: %s",
			fees.ErrBump, originalFee.GasFeeCap, originalFee.GasTipCap, maxPrice)
	}

	cache := f.getDynamicFeeCache()
	currentDynamicFee, err := cache.getDynamicFee()
	if err != nil {
		return
	}

	bumpedMaxPriorityFeePerGas := originalFee.GasTipCap.AddPercentage(f.config.BumpPercent)
	bumpedMaxFeePerGas := originalFee.GasFeeCap.AddPercentage(f.config.BumpPercent)

	bumpedMaxPriorityFeePerGas, err = LimitBumpedFee(originalFee.GasTipCap, currentDynamicFee.GasTipCap, bumpedMaxPriorityFeePerGas, maxPrice)
	if err != nil {
		return bumped, fmt.Errorf("failed to limit maxPriorityFeePerGas: %w", err)
	}

	priorityFeeThreshold, err := cache.getPriorityFeeThreshold()
	if err != nil {
		return bumped, err
	}

	if bumpedMaxPriorityFeePerGas.Cmp(priorityFeeThreshold) > 0 {
		return bumped, fmt.Errorf("%w: bumpedMaxPriorityFeePerGas: %s is above market's %sth percentile: %s, bumping is halted",
			fees.ErrConnectivity, bumpedMaxPriorityFeePerGas, strconv.Itoa(ConnectivityPercentile), priorityFeeThreshold)
	}

	bumpedMaxFeePerGas, err = LimitBumpedFee(originalFee.GasFeeCap, currentDynamicFee.GasFeeCap, bumpedMaxFeePerGas, maxPrice)
	if err != nil {
		return bumped, fmt.Errorf("failed to limit maxFeePerGas: %w", err)
	}

	bumpedFee := DynamicFee{GasFeeCap: bumpedMaxFeePerGas, GasTipCap: bumpedMaxPriorityFeePerGas}
	f.logger.Debugw("bumped dynamic fee", "originalFee", originalFee, "marketFee", currentDynamicFee, "bumpedFee", bumpedFee)

	return bumpedFee, nil
}

// LimitBumpedFee selects the maximum value between the bumped attempt and the current fee, if there is one. If the result is higher than the max price it gets capped.
// Geth's implementation has a hard 10% minimum limit for the bumped values, otherwise it rejects the transaction with an error.
// See: https://github.com/ethereum/go-ethereum/blob/bff330335b94af3643ac2fb809793f77de3069d4/core/tx_list.go#L298
//
// Note: for chains that support EIP-1559 but we still choose to send Legacy transactions to them, the limit is still enforcable due to the fact that Legacy transactions
// are treated the same way as Dynamic transactions under the hood. For chains that don't support EIP-1559 at all, the limit isn't enforcable but a 10% minimum bump percentage
// makes sense anyway.
func LimitBumpedFee(originalFee *assets.Wei, currentFee *assets.Wei, bumpedFee *assets.Wei, maxPrice *assets.Wei) (*assets.Wei, error) {
	if currentFee != nil {
		bumpedFee = assets.WeiMax(currentFee, bumpedFee)
	}
	bumpedFee = assets.WeiMin(bumpedFee, maxPrice)

	// The first check is added for the following edge case:
	// If originalFee is below 10 wei, then adding the minimum bump percentage won't have any effect on the final value because of rounding down.
	// Similarly for bumpedFee, it can have the exact same value as the originalFee, even if we bumped, given an originalFee of less than 10 wei
	// and a small enough BumpPercent.
	if bumpedFee.Cmp(originalFee) == 0 ||
		bumpedFee.Cmp(originalFee.AddPercentage(MinimumBumpPercentage)) < 0 {
		return nil, fmt.Errorf("%w: %s is bumped less than minimum allowed percentage(%s) from originalFee: %s - maxPrice: %s",
			fees.ErrBump, bumpedFee, strconv.Itoa(MinimumBumpPercentage), originalFee, maxPrice)
	}
	return bumpedFee, nil
}

// refreshDynamicFeeNow refreshes the cached dynamic fees synchronously and pushes the next scheduled refresh back,
// so the run loop doesn't immediately repeat the call we just made.
func (f *FeeHistoryEstimator) refreshDynamicFeeNow() error {
	var err error
	if !f.IfStarted(func() {
		if err = f.RefreshDynamicFee(); err == nil {
			f.signalNonBlockingRefresh()
		}
	}) {
		return errors.New("estimator not started")
	}
	return err
}

// signalNonBlockingRefresh resets the run loop's ticker without blocking, deferring the next periodic refresh.
func (f *FeeHistoryEstimator) signalNonBlockingRefresh() {
	select {
	case f.refreshCh <- struct{}{}:
	default:
	}
}

func (f *FeeHistoryEstimator) Name() string                                      { return f.logger.Name() }
func (f *FeeHistoryEstimator) L1Oracle() rollups.L1Oracle                        { return f.l1Oracle }
func (f *FeeHistoryEstimator) HealthReport() map[string]error                    { return map[string]error{f.Name(): nil} }
func (f *FeeHistoryEstimator) OnNewLongestChain(context.Context, *evmtypes.Head) {}
