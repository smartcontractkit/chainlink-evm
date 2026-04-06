package integrationtests

import (
	"context"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers/dualbroadcast"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/storage"
	txmtypes "github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const (
	SimulatedChainID     = 1337
	NumberOfTransactions = 20
)

func setupBackend(t *testing.T, simulationMode SimulationMode) (*txm.Txm, *storage.InMemoryStoreManager, *big.Int, logger.Logger, common.Address, SimulatedClient, AppConfig) {
	env := os.Getenv("ENV")

	// Logger
	lggr, observedLogs := logger.TestObserved(t, zap.DebugLevel)

	switch env {
	case "TESTNET":
		// Env Vars
		envs := LoadEnvVariables(t)
		// Default client
		client, err := ethclient.Dial(envs.RPC)
		require.NoError(t, err)
		chainID, err := client.ChainID(t.Context())
		require.NoError(t, err)

		// Configs
		configs := LoadConfigVariablesWithDefaults(t)
		// TXM
		txm, store, simulatedClient := setupTestnetTXM(t, client, lggr, observedLogs, simulationMode, envs, configs)
		fromAddress := common.HexToAddress(envs.FromAddress)
		return txm, store, chainID, lggr, fromAddress, simulatedClient, configs
	case "DEVNET":
	default:
		txm, store, chainID, fromAddress, simulatedClient, configs := setupDevnetTXM(t, lggr, observedLogs, simulationMode)
		return txm, store, chainID, lggr, fromAddress, simulatedClient, configs
	}
	return nil, nil, nil, nil, common.Address{}, nil, AppConfig{}
}

func setupGasEstimator(t *testing.T, lggr logger.Logger, observedLogs *observer.ObservedLogs, client gas.FeeEstimatorClient, chainID *big.Int, configs AppConfig) gas.EvmFeeEstimator {
	estimator, err := gas.NewEstimator(lggr, client, "", chainID, configs, nil)
	require.NoError(t, err)
	servicetest.Run(t, estimator)
	tests.AssertLogEventually(t, observedLogs, "Fetched") // Ensure there is at least one successful gas estimation stored
	return estimator
}

func setupTestnetTXM(
	t *testing.T,
	c *ethclient.Client,
	lggr logger.Logger,
	observedLogs *observer.ObservedLogs,
	simulationMode SimulationMode,
	envs *EnvVariables,
	configs AppConfig,
) (*txm.Txm, *storage.InMemoryStoreManager, SimulatedClient) {
	// Client
	client := NewGethSimulatedClient(clientwrappers.NewGethClient(c), simulationMode)
	require.NotNil(t, client)
	chainID, err := client.ChainID(t.Context())
	require.NoError(t, err)

	// Gas Estimator
	estimator := setupGasEstimator(t, lggr, observedLogs, client, chainID, configs)

	// Keystore
	keystore := txm.NewKeystore(chainID)
	err = keystore.Add(envs.PrivateKey)
	require.NoError(t, err, "failed to add private key to keystore")

	// AttemptBuilder
	ab := txm.NewAttemptBuilder(configs.PriceMaxKey, estimator, keystore, configs.LimitTransfer())

	// InMemory storage
	store := storage.NewInMemoryStoreManager(lggr, chainID)
	fromAddress := common.HexToAddress(envs.FromAddress)
	err = store.Add(fromAddress)
	require.NoError(t, err, "failed to add address to InMemory store")

	txmConfig := txm.Config{
		EIP1559:             configs.EIP1559DynamicFees(),
		BlockTime:           configs.BlockTime(),
		RetryBlockThreshold: uint16(configs.BumpThreshold()),
		EmptyTxLimitDefault: configs.LimitTransfer(),
	}

	var stuckTxDetector txm.StuckTxDetector
	if simulationMode == StuckTxDetection {
		stuckTxDetector = txm.NewStuckTxDetector(lggr, "", txm.StuckTxDetectorConfig{
			BlockTime:             configs.BlockTime(),
			StuckTxBlockThreshold: uint32(configs.BumpThreshold() - 1),
		})
	}

	var errorHandler txm.ErrorHandler
	if simulationMode == ErrorHandling {
		errorHandler = dualbroadcast.NewErrorHandler()
	}

	// TXM
	txm := txm.NewTxm(lggr, chainID, client, ab, store, stuckTxDetector, txmConfig, keystore, errorHandler)
	require.NotNil(t, txm)
	servicetest.Run(t, txm)
	return txm, store, client
}

func setupDevnetTXM(
	t *testing.T,
	lggr logger.Logger,
	observedLogs *observer.ObservedLogs,
	simulationMode SimulationMode,
) (*txm.Txm, *storage.InMemoryStoreManager, *big.Int, common.Address, SimulatedClient, AppConfig) {
	// Configs
	configs := defaultConfigs()

	// ChainID
	chainID := big.NewInt(SimulatedChainID) // Default chainID for SimulatedBackend

	client, fromAddress, privateKeyHex := setupSimulatedBackendClient(t, chainID, simulationMode, configs.EIP1559DynamicFees())

	// Gas Estimator
	estimator := setupGasEstimator(t, lggr, observedLogs, client, chainID, configs)

	// Keystore
	keystore := txm.NewKeystore(chainID)
	require.NoError(t, keystore.Add(privateKeyHex), "failed to add private key to keystore")

	// AttemptBuilder
	ab := txm.NewAttemptBuilder(configs.PriceMaxKey, estimator, keystore, configs.LimitDefault())

	// InMemory storage
	store := storage.NewInMemoryStoreManager(lggr, chainID)
	require.NoError(t, store.Add(fromAddress), "failed to add address to InMemory store")

	txmConfig := txm.Config{
		EIP1559:             configs.EIP1559DynamicFees(),
		BlockTime:           configs.BlockTime(),
		RetryBlockThreshold: uint16(configs.BumpThreshold()),
		EmptyTxLimitDefault: configs.LimitTransfer(),
	}

	var stuckTxDetector txm.StuckTxDetector
	if simulationMode == StuckTxDetection {
		stuckTxDetector = txm.NewStuckTxDetector(lggr, "", txm.StuckTxDetectorConfig{
			BlockTime:             configs.BlockTime(),
			StuckTxBlockThreshold: uint32(configs.BumpThreshold() - 1),
		})
	}

	var errorHandler txm.ErrorHandler
	if simulationMode == ErrorHandling {
		errorHandler = dualbroadcast.NewErrorHandler()
	}

	// TXM
	txm := txm.NewTxm(lggr, chainID, client, ab, store, stuckTxDetector, txmConfig, keystore, errorHandler)
	require.NotNil(t, txm)
	servicetest.Run(t, txm)

	return txm, store, chainID, fromAddress, client, configs
}

func setupSimulatedBackendClient(t *testing.T, chainID *big.Int, simulationMode SimulationMode, EIP1559DynamicFees bool) (SimulatedClient, common.Address, string) {
	// Keys
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	require.NoError(t, err)
	privateKeyBytes := crypto.FromECDSA(privateKey)

	prefixedHex := hexutil.Encode(privateKeyBytes)
	privateKeyHex := strings.TrimPrefix(prefixedHex, "0x")

	address := auth.From
	genesisAlloc := map[common.Address]types.Account{
		address: {
			Balance: big.NewInt(1000000000000000000), // 1 ETH
		},
	}
	blockGasLimit := uint64(5000000)
	backend := simulated.NewBackend(genesisAlloc, simulated.WithBlockGasLimit(blockGasLimit))

	// Estimators don't play nice with no transactions at all on the network, so we add a history manually.
	client := backend.Client()
	nonce, err := client.PendingNonceAt(context.Background(), address)
	require.NoError(t, err)
	gasPrice, err := client.SuggestGasPrice(t.Context())
	require.NoError(t, err)
	if EIP1559DynamicFees {
		gasTipCap, err := client.SuggestGasTipCap(t.Context())
		require.NoError(t, err)

		tx1 := types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce,
			To:        &address,
			Value:     big.NewInt(1000),
			Gas:       21000,
			GasTipCap: gasTipCap,
			GasFeeCap: gasPrice,
		})
		signedTx1, err := types.SignTx(tx1, types.LatestSignerForChainID(chainID), privateKey)
		require.NoError(t, err)
		require.NoError(t, client.SendTransaction(t.Context(), signedTx1))

		// tx2 needs to have a higher tip and fee cap than tx1 to account for the connectivity threshold.
		tx2 := types.NewTx(&types.DynamicFeeTx{
			ChainID:   chainID,
			Nonce:     nonce + 1,
			To:        &address,
			Value:     big.NewInt(2000),
			Gas:       21000,
			GasTipCap: gasTipCap.Mul(gasTipCap, big.NewInt(2)),
			GasFeeCap: gasPrice.Mul(gasPrice, big.NewInt(2)),
		})
		signedTx2, err := types.SignTx(tx2, types.LatestSignerForChainID(chainID), privateKey)
		require.NoError(t, err)
		require.NoError(t, client.SendTransaction(t.Context(), signedTx2))
	} else {

		// Create and send first legacy transaction
		tx1 := types.NewTransaction(nonce, address, big.NewInt(1000), 21000, gasPrice, nil)
		signedTx1, err := types.SignTx(tx1, types.LatestSignerForChainID(chainID), privateKey)
		require.NoError(t, err)
		err = client.SendTransaction(t.Context(), signedTx1)
		require.NoError(t, err)

		// Create and send second legacy transaction
		tx2 := types.NewTransaction(nonce+1, address, big.NewInt(2000), 21000, gasPrice, nil)
		signedTx2, err := types.SignTx(tx2, types.LatestSignerForChainID(chainID), privateKey)
		require.NoError(t, err)
		err = client.SendTransaction(t.Context(), signedTx2)
		require.NoError(t, err)
	}

	backend.Commit()

	c := NewBackendSimulatedClient(backend, simulationMode)
	return c, address, privateKeyHex
}

func waitUntilQueuesAreEmpty(
	t *testing.T,
	store *storage.InMemoryStoreManager,
	fromAddress common.Address,
	lggr logger.Logger,
	client SimulatedClient,
	blockTime time.Duration,
) {
	require.Eventually(t, func() bool {
		client.Commit()
		unstartedCount, err := store.CountUnstartedTransactions(fromAddress)
		require.NoError(t, err)
		_, unconfirmedCount, err := store.FetchUnconfirmedTransactionAtNonceWithCount(context.TODO(), 0, fromAddress)
		require.NoError(t, err)
		lggr.Debugw("Queue status", "unstarted", unstartedCount, "unconfirmed", unconfirmedCount)
		return unstartedCount == 0 && unconfirmedCount == 0
	}, testutils.WaitTimeout(t), blockTime)
}

// TestIntegration_StandardFlow creates 20 transaction requests and queues them in the TXM.
// It then triggers the TXM to process the transactions It then waits until the queues are empty
// This test is used to verify TXM's standard flow. The number 20 was chosen because it will
// trigger a thottling scenario and will enable the TXM to gracefully handle it.
func TestIntegration_StandardFlow(t *testing.T) {
	txm, store, chainID, lggr, fromAddress, simulatedClient, configs := setupBackend(t, Standard)

	createTransactions(t, txm, chainID, fromAddress, NumberOfTransactions)
	txm.Trigger(fromAddress) // Trigger instantly triggers the TXM instead of waiting for the next cycle.
	waitUntilQueuesAreEmpty(t, store, fromAddress, lggr, simulatedClient, configs.BlockTime())
}

// TestIntegration_Retransmission utilizes the Retransmission simulation mode to test the TXM's
// retransmission logic. That means every other attempt will be assumed to be successful while it wasn't,
// so the TXM will have to handle that.
func TestIntegration_Retransmission(t *testing.T) {
	txm, store, chainID, lggr, fromAddress, simulatedClient, configs := setupBackend(t, Retransmission)

	createTransactions(t, txm, chainID, fromAddress, NumberOfTransactions)
	txm.Trigger(fromAddress) // Trigger instantly triggers the TXM instead of waiting for the next cycle.
	waitUntilQueuesAreEmpty(t, store, fromAddress, lggr, simulatedClient, configs.BlockTime())
}

// TestIntegration_StuckTxDetection tests the TXM's stuck tx detection logic. It injects a mix of stuck and
// non-stuck transactions and attempts.
func TestIntegration_StuckTxDetection(t *testing.T) {
	txm, store, chainID, lggr, fromAddress, simulatedClient, configs := setupBackend(t, StuckTxDetection)

	createTransactions(t, txm, chainID, fromAddress, NumberOfTransactions)
	txm.Trigger(fromAddress) // Trigger instantly triggers the TXM instead of waiting for the next cycle.
	waitUntilQueuesAreEmpty(t, store, fromAddress, lggr, simulatedClient, configs.BlockTime())
}

// TestIntegration_ErrorHandling tests the TXM's error handling logic by injecting transactions with certain error messages
// and checking if the TXM will handle the nonce reassignment correctly.
func TestIntegration_ErrorHandling(t *testing.T) {
	txm, store, chainID, lggr, fromAddress, simulatedClient, configs := setupBackend(t, ErrorHandling)

	createTransactions(t, txm, chainID, fromAddress, NumberOfTransactions)
	txm.Trigger(fromAddress) // Trigger instantly triggers the TXM instead of waiting for the next cycle.
	waitUntilQueuesAreEmpty(t, store, fromAddress, lggr, simulatedClient, configs.BlockTime())
}

func createTransactions(t *testing.T, txm *txm.Txm, chainID *big.Int, fromAddress common.Address, count int) {
	for range count {
		//Create request
		txRequest := txmtypes.TxRequest{
			ChainID:           chainID,
			FromAddress:       fromAddress,
			ToAddress:         fromAddress,
			Value:             big.NewInt(50),
			Data:              []byte{128, 100, 11},
			SpecifiedGasLimit: 40000,
		}
		_, err := txm.CreateTransaction(t.Context(), &txRequest)
		require.NoError(t, err, "failed to create transaction")
	}
}
