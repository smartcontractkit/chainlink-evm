package legacyevm

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	evmclient "github.com/smartcontractkit/chainlink-evm/pkg/client"
	evmconfig "github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas/rollups"
	evmheads "github.com/smartcontractkit/chainlink-evm/pkg/heads"
	"github.com/smartcontractkit/chainlink-evm/pkg/logpoller"
	"github.com/smartcontractkit/chainlink-evm/pkg/txmgr"
)

func newEvmTxm(
	ds sqlutil.DataSource,
	cfg evmconfig.EVM,
	databaseConfig txmgr.DatabaseConfig,
	listenerConfig txmgr.ListenerConfig,
	client evmclient.Client,
	lggr logger.Logger,
	logPoller logpoller.LogPoller,
	opts ChainRelayOpts,
	headTracker evmheads.Tracker,
	estimator gas.EvmFeeEstimator,
	clientsByChainID map[string]rollups.DAClient,
) (txm txmgr.TxManager,
	err error,
) {
	chainID := cfg.ChainID()

	lggr = logger.Named(lggr, "Txm")
	lggr.Infow("Initializing EVM transaction manager",
		"bumpTxDepth", cfg.GasEstimator().BumpTxDepth(),
		"maxInFlightTransactions", cfg.Transactions().MaxInFlight(),
		"maxQueuedTransactions", cfg.Transactions().MaxQueued(),
		"nonceAutoSync", cfg.NonceAutoSync(),
		"limitDefault", cfg.GasEstimator().LimitDefault(),
	)

	err = validateConfirmationTimeout(cfg)
	if err != nil {
		return nil, err
	}

	if opts.GenTxManager == nil {
		var txmv2 txmgr.TxManager
		txV2Cfg := cfg.Transactions().TransactionManagerV2()
		dualBroadcastEnabled := txV2Cfg.Enabled() && txV2Cfg.DualBroadcast() != nil &&
			*txV2Cfg.DualBroadcast() && txV2Cfg.CustomURL() != nil

		if txV2Cfg.Enabled() {
			// TxM v2 gets its own FeeHistory estimator with a potentially different TransactionPercentile,
			// allowing SVR transactions to target a different gas price aggressiveness than primary transactions.
			txmV2Estimator, estErr := newTxmV2FeeHistoryEstimator(cfg, client, lggr, clientsByChainID)
			if estErr != nil {
				return nil, fmt.Errorf("failed to initialize TxM v2 FeeHistory estimator: %w", estErr)
			}
			txmv2, err = txmgr.NewTxmV2(
				ds,
				cfg,
				txmgr.NewEvmTxmFeeConfig(cfg.GasEstimator()),
				cfg.Transactions(),
				txV2Cfg,
				client,
				lggr,
				logPoller,
				opts.KeyStore,
				txmV2Estimator,
				cfg.GasEstimator(),
			)
			if err != nil {
				return nil, err
			}

			if !dualBroadcastEnabled {
				return txmv2, nil
			}
		}
		txm, err = txmgr.NewTxm(
			ds,
			cfg,
			txmgr.NewEvmTxmFeeConfig(cfg.GasEstimator()),
			cfg.Transactions(),
			cfg.NodePool().Errors(),
			databaseConfig,
			listenerConfig,
			client,
			lggr,
			logPoller,
			opts.KeyStore,
			estimator,
			headTracker,
			txmv2,
			dualBroadcastEnabled)
	} else {
		txm = opts.GenTxManager(chainID)
	}
	return
}

const maximumConfirmationTimeout = time.Second * 600

func validateConfirmationTimeout(cfg evmconfig.EVM) error {
	if cfg.ConfirmationTimeout() > maximumConfirmationTimeout {
		return fmt.Errorf("ConfirmationTimeout cannot be greater than 10 minutes, got %s", cfg.ConfirmationTimeout())
	}
	return nil
}

func newGasEstimator(
	cfg evmconfig.EVM,
	client evmclient.Client,
	lggr logger.Logger,
	opts ChainRelayOpts,
	clientsByChainID map[string]rollups.DAClient,
) (estimator gas.EvmFeeEstimator, err error) {
	lggr = logger.Named(lggr, "GasEstimator")
	chainID := cfg.ChainID()
	// build estimator from factory
	if opts.GenGasEstimator == nil {
		if estimator, err = gas.NewEstimator(lggr, client, cfg.ChainType(), chainID, cfg.GasEstimator(), clientsByChainID); err != nil {
			return nil, fmt.Errorf("failed to initialize estimator: %w", err)
		}
	} else {
		estimator = opts.GenGasEstimator(chainID)
	}
	return
}

// newTxmV2FeeHistoryEstimator creates a FeeHistory estimator for TxM v2.
// It derives all config from the shared GasEstimator config, except RewardPercentile
// which is overridden by TransactionManagerV2.TransactionPercentile if set.
func newTxmV2FeeHistoryEstimator(
	cfg evmconfig.EVM,
	client evmclient.Client,
	lggr logger.Logger,
	clientsByChainID map[string]rollups.DAClient,
) (gas.EvmFeeEstimator, error) {
	lggr = logger.Named(lggr, "TxmV2FeeHistoryEstimator")
	chainID := cfg.ChainID()
	geCfg := cfg.GasEstimator()
	txmV2Cfg := cfg.Transactions().TransactionManagerV2()

	// Determine the RewardPercentile: use TxM v2 override if set, otherwise fall back to the shared config
	rewardPercentile := float64(geCfg.BlockHistory().TransactionPercentile())
	if txmV2Cfg.TransactionPercentile() != nil {
		rewardPercentile = float64(*txmV2Cfg.TransactionPercentile())
	}

	l1Oracle, err := rollups.NewL1GasOracle(lggr, client, cfg.ChainType(), geCfg.DAOracle(), clientsByChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize L1 oracle for TxM v2: %w", err)
	}

	fhCfg := gas.FeeHistoryEstimatorConfig{
		BumpPercent:      geCfg.BumpPercent(),
		CacheTimeout:     geCfg.FeeHistory().CacheTimeout(),
		EIP1559:          geCfg.EIP1559DynamicFees(),
		BlockHistorySize: uint64(geCfg.BlockHistory().BlockHistorySize()),
		RewardPercentile: rewardPercentile,
	}

	lggr.Infow("Initializing TxM v2 FeeHistory gas estimator",
		"rewardPercentile", rewardPercentile,
		"blockHistorySize", fhCfg.BlockHistorySize,
		"eip1559DynamicFees", fhCfg.EIP1559,
		"bumpPercent", fhCfg.BumpPercent,
	)

	newEstimator := func(l logger.Logger) gas.EvmEstimator {
		return gas.NewFeeHistoryEstimator(lggr, client, fhCfg, chainID, l1Oracle)
	}
	return gas.NewEvmFeeEstimator(lggr, newEstimator, geCfg.EIP1559DynamicFees(), geCfg, client), nil
}
