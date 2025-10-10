package keysv2_test

import (
	"context"
	"testing"
	"time"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	commonks "github.com/smartcontractkit/chainlink-common/keystore"
	ksstorage "github.com/smartcontractkit/chainlink-common/keystore/storage"
	logger "github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/keysv2"
	evmks "github.com/smartcontractkit/chainlink-evm/pkg/keysv2"
	ocr2agg "github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
	median "github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median"
	"github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median/evmreportcodec"
	libocr "github.com/smartcontractkit/libocr/offchainreporting2plus"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/chains/evmutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"
)

type ds struct{ lggr logger.Logger }

func (d *ds) Observe(ctx context.Context, repts ocrtypes.ReportTimestamp) (*big.Int, error) {
	d.lggr.Info("Observe", "repts", repts)
	return nil, nil
}

type transmitter struct {
	backend *simulated.Backend
	lggr    logger.Logger
	ocr2agg *ocr2agg.OCR2Aggregator
	opts    *bind.TransactOpts
}

func (t *transmitter) Transmit(
	ctx context.Context,
	reportContext ocrtypes.ReportContext,
	report ocrtypes.Report,
	signatures []ocrtypes.AttributedOnchainSignature,
) error {
	t.lggr.Info("Transmit", "report", report)
	var rs [][32]byte
	var ss [][32]byte
	var vs [32]byte
	for i, as := range signatures {
		r, s, v, err := evmutil.SplitSignature(as.Signature)
		if err != nil {
			panic("eventTransmit(ev): error in SplitSignature")
		}
		rs = append(rs, r)
		ss = append(ss, s)
		vs[i] = v
	}
	t.ocr2agg.Transmit(t.opts,
		evmutil.RawReportContext(reportContext),
		report, rs, ss, vs)
	t.backend.Commit()
	return nil
}

func (t *transmitter) LatestRoundRequested(ctx context.Context, _ time.Duration) (ocrtypes.ConfigDigest, uint32, uint8, error) {
	res, err := t.ocr2agg.LatestConfigDigestAndEpoch(&bind.CallOpts{Context: ctx})
	return ocrtypes.ConfigDigest(res.ConfigDigest), res.Epoch, 0, err
}

func (t *transmitter) LatestConfigDigestAndEpoch(ctx context.Context) (ocrtypes.ConfigDigest, uint32, error) {
	res, err := t.ocr2agg.LatestConfigDigestAndEpoch(&bind.CallOpts{Context: ctx})
	return ocrtypes.ConfigDigest(res.ConfigDigest), res.Epoch, err
}

func (t *transmitter) LatestTransmissionDetails(ctx context.Context) (ocrtypes.ConfigDigest, uint32, uint8, *big.Int, time.Time, error) {
	res, err := t.ocr2agg.LatestConfigDigestAndEpoch(&bind.CallOpts{Context: ctx})
	return ocrtypes.ConfigDigest(res.ConfigDigest), res.Epoch, 0, nil, time.Time{}, err
}

func (t *transmitter) FromAccount(ctx context.Context) (ocrtypes.Account, error) {
	return ocrtypes.Account(t.opts.From.String()), nil
}

// TestOCR2Keyring_Integration tests the OCR2 keyrings integration
// with libocr to ensure that the keyrings can actually be used
// to sign and verify reports.
func TestOCR2Keyring_Integration(t *testing.T) {
	lggr := logger.Test(t)
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
	aggAddress, _, agg, err := ocr2agg.DeployOCR2Aggregator(opts, backend.Client(), common.HexToAddress("0x0"), big.NewInt(0), big.NewInt(0), common.HexToAddress("0x0"), common.HexToAddress("0x0"), 18, "Test")
	require.NoError(t, err)

	transmitter := transmitter{backend: backend, lggr: lggr, ocr2agg: agg, opts: opts}

	_, err = libocr.NewOracle(libocr.OCR2OracleArgs{
		ReportingPluginFactory: &median.NumericalMedianFactory{
			ContractTransmitter:                  &transmitter,
			DataSource:                           &ds{},
			JuelsPerFeeCoinDataSource:            nil,
			GasPriceSubunitsDataSource:           nil,
			IncludeGasPriceSubunitsInObservation: false,
			Logger:                               logger.NewOCRWrapper(lggr, true, func(string) {}),
			OnchainConfigCodec:                   median.StandardOnchainConfigCodec{},
			ReportCodec:                          evmreportcodec.ReportCodec{},
			DeviationFunc:                        nil,
		},
		ContractConfigTracker: nil,
		ContractTransmitter:   &transmitter,
		Database:              nil,
		LocalConfig: ocrtypes.LocalConfig{
			BlockchainTimeout:                  10 * time.Second,
			ContractConfigConfirmations:        1,
			ContractConfigTrackerPollInterval:  10 * time.Second,
			ContractConfigLoadTimeout:          10 * time.Second,
			ContractTransmitterTransmitTimeout: 10 * time.Second,
			DatabaseTimeout:                    10 * time.Second,
			DevelopmentMode:                    ocrtypes.EnableDangerousDevelopmentMode,
			EnableTransmissionTelemetry:        false,
			MinOCR2MaxDurationQuery:            10 * time.Second,
			SkipContractConfigConfirmations:    false,
		},
		Logger:             logger.NewOCRWrapper(lggr, true, func(string) {}),
		MonitoringEndpoint: nil,
		MetricsRegisterer:  nil,
		OffchainConfigDigester: evmutil.EVMOffchainConfigDigester{
			ChainID:         uint64(1337),
			ContractAddress: aggAddress,
		},
		OffchainKeyring: offchainKeyring,
		OnchainKeyring:  onchainKeyring,
	})
	require.NoError(t, err)
}
