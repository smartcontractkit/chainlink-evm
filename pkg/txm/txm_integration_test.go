package txm

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	input := &blockchain.Input{
		Type:    "geth",
		Port:    "8211",
		ChainID: "1337",
	}
	output, err := blockchain.NewBlockchainNetwork(input)
	httpUrl := output.Nodes[0].ExternalHTTPUrl
	require.NoError(t, err)

	ethClient, err := ethclient.DialContext(t.Context(), httpUrl)
	gethClient := clientwrappers.NewGethClient(ethClient)

	// Initialize KeyStore and create new key
	keyStore := keystest.NewMemoryChainStore()
	chainStore := keys.NewChainStore(keyStore, testutils.FixtureChainID)
	address, err := keyStore.Create()
	require.NoError(t, err)

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

	IDK := "IDK"
	txRequest := &txmtypes.TxRequest{
		Data:              []byte{100, 200},
		IdempotencyKey:    &IDK,
		ChainID:           testutils.FixtureChainID,
		FromAddress:       address,
		ToAddress:         testutils.NewAddress(),
		SpecifiedGasLimit: 1_000_000_000, // set insane gas limit to fail initial broadcast
	}

	tx, err := txm.CreateTransaction(t.Context(), txRequest)
	require.NoError(t, err)
	require.NotNil(t, tx)

	require.Eventually(t, func() bool {
		curTXMNonce := txm.getNonce(address)
		if curTXMNonce > 0 {
			return true
		}
		return false
	}, time.Minute, time.Second, "transaction never picked up for broadcast")

	require.Eventually(t, func() bool {
		txCount, err := gethClient.NonceAt(t.Context(), address, nil)
		require.NoError(t, err)
		if txCount > 0 {
			return true
		}
		return false
	}, 2*time.Minute, time.Second, "nonce never filled")
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
func (g TestGasEstimatorConfig) LimitMax() uint64                   { return 0 }
func (g TestGasEstimatorConfig) LimitMultiplier() float32           { return 0 }
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
