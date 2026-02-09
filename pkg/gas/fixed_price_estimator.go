package gas

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-framework/chains/fees"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/rollups"
	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
)

var _ EvmEstimator = (*fixedPriceEstimator)(nil)

type fixedPriceEstimator struct {
	config                    fixedPriceEstimatorConfig
	EIP1559FeeCapBufferBlocks uint16
	lggr                      logger.SugaredLogger
	l1Oracle                  rollups.L1Oracle
}
type fixedPriceEstimatorConfig interface {
	BumpThreshold() uint64
	FeeCapDefault() *assets.Wei
	PriceDefault() *assets.Wei
	Mode() string
	bumpConfig
}

// NewFixedPriceEstimator returns a new "FixedPrice" estimator which will
// always use the config default values for gas prices and limits
func NewFixedPriceEstimator(cfg fixedPriceEstimatorConfig, ethClient FeeEstimatorClient, EIP1559FeeCapBufferBlocks uint16, lggr logger.Logger, l1Oracle rollups.L1Oracle) EvmEstimator {
	return &fixedPriceEstimator{cfg, EIP1559FeeCapBufferBlocks, logger.Sugared(logger.Named(lggr, "FixedPriceEstimator")), l1Oracle}
}

func (f *fixedPriceEstimator) Start(context.Context) error {
	if f.config.BumpThreshold() == 0 && f.config.Mode() == "FixedPrice" {
		// EvmGasFeeCapDefault is ignored if fixed estimator mode is on and gas bumping is disabled
		if f.config.FeeCapDefault().Cmp(f.config.PriceMax()) != 0 {
			f.lggr.Infof("You are using FixedPrice estimator with gas bumping disabled. EVM.GasEstimator.PriceMax (value: %s) will be used as the FeeCap for transactions", f.config.PriceMax())
		}
	}
	return nil
}

func (f *fixedPriceEstimator) GetLegacyGas(_ context.Context, _ []byte, gasLimit uint64, maxGasPriceWei *assets.Wei, _ ...fees.Opt) (*assets.Wei, uint64, error) {
	gasPrice := fees.CalculateFee(f.config.PriceDefault().ToInt(), maxGasPriceWei.ToInt(), f.config.PriceMax().ToInt())
	chainSpecificGasLimit := gasLimit
	return assets.NewWei(gasPrice), chainSpecificGasLimit, nil
}

func (f *fixedPriceEstimator) BumpLegacyGas(
	_ context.Context,
	originalGasPrice *assets.Wei,
	originalGasLimit uint64,
	maxGasPriceWei *assets.Wei,
	_ []EvmPriorAttempt,
) (*assets.Wei, uint64, error) {
	gasPrice, err := fees.CalculateBumpedFee(
		f.lggr,
		f.config.PriceDefault().ToInt(),
		originalGasPrice.ToInt(),
		maxGasPriceWei.ToInt(),
		f.config.PriceMax().ToInt(),
		f.config.BumpMin().ToInt(),
		f.config.BumpPercent(),
		assets.FormatWei,
	)
	if err != nil {
		return nil, 0, err
	}

	chainSpecificGasLimit := originalGasLimit
	return assets.NewWei(gasPrice), chainSpecificGasLimit, err
}

// GetMaxLegacyGas returns the result of GetLegacyGas. FixedPriceEstimator provides fixed gas prices, which generally indicates there is no priority queue or the network
// expects fixed prices. Either way, fetching a standard gas estimation will have the same effect for transaction inclusion.
func (f *fixedPriceEstimator) GetMaxLegacyGas(ctx context.Context, calldata []byte, gasLimit uint64, maxGasPriceWei *assets.Wei, opts ...fees.Opt) (gasPrice *assets.Wei, chainSpecificGasLimit uint64, err error) {
	return f.GetLegacyGas(ctx, calldata, gasLimit, maxGasPriceWei, opts...)
}

func (f *fixedPriceEstimator) GetDynamicFee(_ context.Context, maxGasPriceWei *assets.Wei) (d DynamicFee, err error) {
	gasTipCap := f.config.TipCapDefault()

	if gasTipCap == nil {
		return d, pkgerrors.New("cannot calculate dynamic fee: EthGasTipCapDefault was not set")
	}

	var feeCap *assets.Wei
	if f.config.BumpThreshold() == 0 {
		// Gas bumping is disabled, just use the max fee cap
		feeCap = getMaxGasPrice(maxGasPriceWei, f.config.PriceMax())
	} else {
		// Need to leave headroom for bumping so we fallback to the default value here
		feeCap = f.config.FeeCapDefault()
	}

	return DynamicFee{
		GasFeeCap: feeCap,
		GasTipCap: gasTipCap,
	}, nil
}

func (f *fixedPriceEstimator) BumpDynamicFee(
	_ context.Context,
	originalFee DynamicFee,
	maxGasPriceWei *assets.Wei,
	_ []EvmPriorAttempt,
) (bumped DynamicFee, err error) {
	return BumpDynamicFeeOnly(
		f.config,
		f.EIP1559FeeCapBufferBlocks,
		f.lggr,
		f.config.TipCapDefault(),
		nil,
		originalFee,
		maxGasPriceWei,
	)
}

func (f *fixedPriceEstimator) L1Oracle() rollups.L1Oracle {
	return f.l1Oracle
}

func (f *fixedPriceEstimator) GetMaxDynamicFee(maxGasPriceWei *assets.Wei) (fee DynamicFee, err error) {
	return f.GetDynamicFee(context.Background(), maxGasPriceWei) // context is not used by this method
}

func (f *fixedPriceEstimator) Name() string                                          { return f.lggr.Name() }
func (f *fixedPriceEstimator) Ready() error                                          { return nil }
func (f *fixedPriceEstimator) HealthReport() map[string]error                        { return map[string]error{} }
func (f *fixedPriceEstimator) Close() error                                          { return nil }
func (f *fixedPriceEstimator) OnNewLongestChain(_ context.Context, _ *evmtypes.Head) {}
