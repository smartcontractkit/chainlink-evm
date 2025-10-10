package keysv2_test

import (
	"context"
	"testing"
	"time"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	commonks "github.com/smartcontractkit/chainlink-common/keystore"
	ksstorage "github.com/smartcontractkit/chainlink-common/keystore/storage"
	"github.com/smartcontractkit/chainlink-evm/pkg/keysv2"
	evmks "github.com/smartcontractkit/chainlink-evm/pkg/keysv2"
	ocr2agg "github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
	median "github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median"
	libocr "github.com/smartcontractkit/libocr/offchainreporting2plus"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"
)

// TestOCR2Keyring_Integration tests the OCR2 keyrings integration
// with libocr to ensure that the keyrings can actually be used
// to sign and verify reports.
func TestOCR2Keyring_Integration(t *testing.T) {
	storage := ksstorage.NewMemoryStorage()
	ctx := t.Context()
	ks, err := commonks.LoadKeystore(ctx, storage, commonks.EncryptionParams{
		Password:     "test-password",
		ScryptParams: commonks.FastScryptParams,
	})
	require.NoError(t, err)

	onchainKeyring, err := keysv2.CreateOCR2OnchainKeyring(context.Background(), ks, "test-onchain-keyring")
	require.NoError(t, err)
	offchainKeyring, err := keysv2.CreateOCR2OffchainKeyring(context.Background(), ks, "test-offchain-keyring")
	require.NoError(t, err)

	testKey, err := evmks.CreateTxKey(ks, "test-tx-key")
	require.NoError(t, err)
	backend := simulated.NewBackend(types.GenesisAlloc{
		testKey.Address(): {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)), // 10 ETH
		},
	}, simulated.WithBlockGasLimit(10e6))
	defer backend.Close()
	opts, err := testKey.GetTransactOpts(ctx, big.NewInt(1337))
	require.NoError(t, err)
	_, _, _, err = ocr2agg.DeployOCR2Aggregator(opts, backend.Client(), common.HexToAddress("0x0"), big.NewInt(0), big.NewInt(0), common.HexToAddress("0x0"), common.HexToAddress("0x0"), 18, "Test")
	require.NoError(t, err)

	_, err = libocr.NewOracle(libocr.OCR2OracleArgs{
		ReportingPluginFactory: &median.NumericalMedianFactory{
			ContractTransmitter:                  nil,
			DataSource:                           nil,
			JuelsPerFeeCoinDataSource:            nil,
			GasPriceSubunitsDataSource:           nil,
			IncludeGasPriceSubunitsInObservation: false,
			Logger:                               nil,
			OnchainConfigCodec:                   nil,
			ReportCodec:                          nil,
			DeviationFunc:                        nil,
		},
		ContractConfigTracker: nil,
		ContractTransmitter:   nil,
		Database:              nil,
		LocalConfig: ocrtypes.LocalConfig{
			BlockchainTimeout:                  10 * time.Second,
			ContractConfigConfirmations:        1,
			ContractConfigTrackerPollInterval:  10 * time.Second,
			ContractConfigLoadTimeout:          10 * time.Second,
			ContractTransmitterTransmitTimeout: 10 * time.Second,
			DatabaseTimeout:                    10 * time.Second,
			DevelopmentMode:                    "enable dangerous development mode",
			EnableTransmissionTelemetry:        false,
			MinOCR2MaxDurationQuery:            10 * time.Second,
			SkipContractConfigConfirmations:    false,
		},
		Logger:                 nil,
		MonitoringEndpoint:     nil,
		MetricsRegisterer:      nil,
		OffchainConfigDigester: nil,
		OffchainKeyring:        offchainKeyring,
		OnchainKeyring:         onchainKeyring,
	})
	require.NoError(t, err)
}
