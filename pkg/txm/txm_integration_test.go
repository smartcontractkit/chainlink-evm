package txm

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/google/uuid"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/gas"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys/keystest"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/storage"
	txmtypes "github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/stretchr/testify/require"
)

func Test_StuckTx(t *testing.T) {
	lggr := logger.Test(t)
	address := common.HexToAddress(blockchain.DefaultGethPublicKey)
	priveKeyStr := blockchain.DefaultGethPrivateKey
	// Initialize KeyStore and add pre-funded key
	keyStore := keystest.NewMemoryChainStore()
	privKey, err := crypto.HexToECDSA(priveKeyStr)
	require.NoError(t, err)
	keyStore.AddKey(address, privKey)
	chainStore := keys.NewChainStore(keyStore, testutils.FixtureChainID)
	
	input := &blockchain.Input{
		Type:    "geth",
		Port:    "8211",
		ChainID: testutils.FixtureChainID.String(),
	}
	output, err := blockchain.NewBlockchainNetwork(input)
	require.NoError(t, err)
	require.NotEmpty(t, output.Nodes)
	httpUrl := output.Nodes[0].ExternalHTTPUrl
	// Wait for endpoint to respond
	require.Eventually(t, func() bool {
		client, err := rpc.DialContext(t.Context(), httpUrl)
			if err != nil {
				return false
			}

			var blockNumber string
			if err := client.CallContext(t.Context(), &blockNumber, "eth_blockNumber"); err != nil {
				return false
			}

			client.Close()
			// If we get here, the endpoint is responding
			return true
	}, 2*time.Minute, time.Second, "rpc endpoint never responded")

	ethClient, err := ethclient.DialContext(t.Context(), httpUrl)
	gethClient := clientwrappers.NewGethClient(ethClient)

	txStore := storage.NewInMemoryStoreManager(lggr, testutils.FixtureChainID)
	require.NoError(t, txStore.Add(address))

	priceMax := func(common.Address) *assets.Wei{return assets.NewWeiI(1_000_000)}
	estimatorCfg := TestGasEstimatorConfig{bumpThreshold: 5}
	fixedPriceEstimator := gas.NewFixedPriceEstimator(estimatorCfg, gethClient, nil, lggr, nil)
	newEstimator := func(logger.Logger) gas.EvmEstimator{return fixedPriceEstimator}
	estimator := gas.NewEvmFeeEstimator(lggr, newEstimator, false, estimatorCfg, gethClient)
	ab := NewAttemptBuilder(priceMax, estimator, chainStore)
	config := Config{BlockTime: 1 * time.Minute, RetryBlockThreshold: 10}

	stuckTxDetectorConfig := StuckTxDetectorConfig{
		BlockTime:             10 * time.Second,
		StuckTxBlockThreshold: 2,
	}
	stuckTxDetector := NewStuckTxDetector(lggr, "", stuckTxDetectorConfig)

	txm := NewTxm(lggr, testutils.FixtureChainID, gethClient, ab, txStore, stuckTxDetector, config, chainStore)
	servicetest.Run(t, txm)
	toAddress := testutils.NewAddress()

	// Create a bad tx that will hold up the TXM
	IDK := uuid.NewString()	
	txRequest := &txmtypes.TxRequest{
		IdempotencyKey:    &IDK,
		ChainID:           testutils.FixtureChainID,
		FromAddress:       address,
		ToAddress:         toAddress,
		SpecifiedGasLimit: 1_000_000_000, // set insane gas limit to fail initial broadcast
		Value:             big.NewInt(1),
	}

	tx, err := txm.CreateTransaction(t.Context(), txRequest)
	require.NoError(t, err)
	require.NotNil(t, tx)

	// Create good tx that gets stuck because of the previous
	IDK2 := uuid.NewString()
	txRequest2 := &txmtypes.TxRequest{
		IdempotencyKey:    &IDK2,
		ChainID:           testutils.FixtureChainID,
		FromAddress:       address,
		ToAddress:         toAddress,
		SpecifiedGasLimit: 21_000, // normal gas limit
		Value:             big.NewInt(1),
	}

	tx2, err := txm.CreateTransaction(t.Context(), txRequest2)
	require.NoError(t, err)
	require.NotNil(t, tx2)

	require.Eventually(t, func() bool {
		curTXMNonce := txm.getNonce(address)
		// Wait till both transactions have been picked up for broadcast
		if curTXMNonce > 1 {
			return true
		}
		return false
	}, time.Minute, time.Second, "transactions never picked up for broadcast")

	require.Eventually(t, func() bool {
		txCount, err := gethClient.NonceAt(t.Context(), address, nil)
		require.NoError(t, err)
		// Only expect 1 transaction to succeed broadcast
		if txCount > 0 {
			return true
		}
		return false
	}, time.Minute, time.Second, "transactions are stuck")
}

type TestGasEstimatorConfig struct {
	bumpThreshold uint64
}

func (g TestGasEstimatorConfig) EIP1559DynamicFees() bool           { return false }
func (g TestGasEstimatorConfig) LimitDefault() uint64               { return 42 }
func (g TestGasEstimatorConfig) BumpPercent() uint16                { return 42 }
func (g TestGasEstimatorConfig) BumpThreshold() uint64              { return g.bumpThreshold }
func (g TestGasEstimatorConfig) BumpMin() *assets.Wei               { return assets.NewWeiI(42) }
func (g TestGasEstimatorConfig) FeeCapDefault() *assets.Wei         { return assets.NewWeiI(42) }
func (g TestGasEstimatorConfig) PriceDefault() *assets.Wei          { return assets.NewWeiI(42) }
func (g TestGasEstimatorConfig) TipCapDefault() *assets.Wei         { return assets.NewWeiI(42) }
func (g TestGasEstimatorConfig) TipCapMin() *assets.Wei             { return assets.NewWeiI(42) }
func (g TestGasEstimatorConfig) LimitMax() uint64                   { return 1_000_000_000 }
func (g TestGasEstimatorConfig) LimitMultiplier() float32           { return 1 }
func (g TestGasEstimatorConfig) BumpTxDepth() uint32                { return 42 }
func (g TestGasEstimatorConfig) LimitTransfer() uint64              { return 42 }
func (g TestGasEstimatorConfig) PriceMax() *assets.Wei              { return assets.NewWeiI(42) }
func (g TestGasEstimatorConfig) PriceMin() *assets.Wei              { return assets.NewWeiI(42) }
func (g TestGasEstimatorConfig) Mode() string                       { return "FixedPrice" }
func (g TestGasEstimatorConfig) EstimateLimit() bool                { return false }
func (g TestGasEstimatorConfig) SenderAddress() *types.EIP55Address { return nil }

func (g TestGasEstimatorConfig) PriceMaxKey(addr common.Address) *assets.Wei {
	return assets.NewWeiI(42)
}
