package keysv2_test

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"testing"
	"time"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	commonks "github.com/smartcontractkit/chainlink-common/keystore"
	ksstorage "github.com/smartcontractkit/chainlink-common/keystore/storage"
	logger "github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/keysv2"
	evmks "github.com/smartcontractkit/chainlink-evm/pkg/keysv2"
	"github.com/smartcontractkit/libocr/commontypes"
	ocr2agg "github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	median "github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median"
	"github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median/evmreportcodec"
	libocr "github.com/smartcontractkit/libocr/offchainreporting2plus"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/chains/evmutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"github.com/stretchr/testify/require"
)

var _ ocrtypes.ContractConfigTracker = (*helper)(nil)
var _ ocrtypes.ContractTransmitter = (*helper)(nil)
var _ median.DataSource = (*helper)(nil)
var _ ocrtypes.Database = (*helper)(nil)

type helper struct {
	backend *simulated.Backend
	lggr    logger.Logger
	ocr2agg *ocr2agg.OCR2Aggregator
	opts    *bind.TransactOpts
}

func (t *helper) Observe(ctx context.Context, repts ocrtypes.ReportTimestamp) (*big.Int, error) {
	t.lggr.Info("Observe", "repts", repts)
	return big.NewInt(1), nil
}

func (t *helper) Transmit(
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

func (t *helper) LatestRoundRequested(ctx context.Context, _ time.Duration) (ocrtypes.ConfigDigest, uint32, uint8, error) {
	t.lggr.Info("LatestRoundRequested")
	res, err := t.ocr2agg.LatestConfigDigestAndEpoch(&bind.CallOpts{Context: ctx})
	return ocrtypes.ConfigDigest(res.ConfigDigest), res.Epoch, 0, err
}

func (t *helper) LatestConfigDigestAndEpoch(ctx context.Context) (ocrtypes.ConfigDigest, uint32, error) {
	t.lggr.Info("LatestConfigDigestAndEpoch")
	res, err := t.ocr2agg.LatestConfigDigestAndEpoch(&bind.CallOpts{Context: ctx})
	return ocrtypes.ConfigDigest(res.ConfigDigest), res.Epoch, err
}

func (t *helper) LatestTransmissionDetails(ctx context.Context) (ocrtypes.ConfigDigest, uint32, uint8, *big.Int, time.Time, error) {
	t.lggr.Info("LatestTransmissionDetails")
	res, err := t.ocr2agg.LatestTransmissionDetails(&bind.CallOpts{Context: ctx})
	return ocrtypes.ConfigDigest(res.ConfigDigest), res.Epoch, 0, nil, time.Time{}, err
}

func (t *helper) FromAccount(ctx context.Context) (ocrtypes.Account, error) {
	t.lggr.Info("FromAccount")
	return ocrtypes.Account(t.opts.From.String()), nil
}

func (t *helper) LatestBlockHeight(ctx context.Context) (uint64, error) {
	header, err := t.backend.Client().HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), err
}

func (t *helper) LatestConfig(ctx context.Context, changedInBlock uint64) (ocrtypes.ContractConfig, error) {
	t.lggr.Info("LatestConfig", "changedInBlock", changedInBlock, "ocr2agg", t.ocr2agg)
	c, err := t.ocr2agg.FilterConfigSet(&bind.FilterOpts{Context: ctx, Start: uint64(changedInBlock)})
	if err != nil {
		return ocrtypes.ContractConfig{}, err
	}
	ok := c.Next()
	if !ok {
		return ocrtypes.ContractConfig{}, fmt.Errorf("no config set event found")
	}
	t.lggr.Infof("ConfigSet %x\n", c.Event.ConfigDigest[:])
	return evmutil.ContractConfigFromConfigSetEvent(*c.Event), nil
}

func (t *helper) LatestConfigDetails(ctx context.Context) (uint64, ocrtypes.ConfigDigest, error) {
	c, err := t.ocr2agg.LatestConfigDetails(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, ocrtypes.ConfigDigest{}, err
	}
	t.lggr.Info("LatestConfigDetails", "c", c)
	return uint64(c.BlockNumber), ocrtypes.ConfigDigest(c.ConfigDigest), nil
}

func (t *helper) Notify() <-chan struct{} {
	return nil
}

func (t *helper) ReadState(ctx context.Context, configDigest ocrtypes.ConfigDigest) (*ocrtypes.PersistentState, error) {
	return nil, nil
}

func (t *helper) WriteState(ctx context.Context, configDigest ocrtypes.ConfigDigest, state ocrtypes.PersistentState) error {
	return nil
}

func (t *helper) StorePendingTransmission(ctx context.Context, reportTimestamp ocrtypes.ReportTimestamp, pendingTransmission ocrtypes.PendingTransmission) error {
	return nil
}

func (t *helper) DeletePendingTransmission(ctx context.Context, reportTimestamp ocrtypes.ReportTimestamp) error {
	return nil
}

func (t *helper) DeletePendingTransmissionsOlderThan(ctx context.Context, time time.Time) error {
	return nil
}

func (t *helper) PendingTransmissionsWithConfigDigest(ctx context.Context, configDigest ocrtypes.ConfigDigest) (map[ocrtypes.ReportTimestamp]ocrtypes.PendingTransmission, error) {
	return nil, nil
}

func (t *helper) ReadConfig(ctx context.Context) (*ocrtypes.ContractConfig, error) {
	return nil, nil
}

func (t *helper) WriteConfig(ctx context.Context, config ocrtypes.ContractConfig) error {
	return nil
}

func (t *helper) NewEndpoint(configDigest ocrtypes.ConfigDigest, peerIDs []string,
	v2bootstrappers []commontypes.BootstrapperLocator, f int, limits ocrtypes.BinaryNetworkEndpointLimits) (commontypes.BinaryNetworkEndpoint, error) {
	t.lggr.Info("NewEndpoint", "configDigest", configDigest, "peerIDs", peerIDs, "v2bootstrappers", v2bootstrappers, "f", f, "limits", limits)
	return nil, nil
}

func (t *helper) PeerID() string {
	t.lggr.Info("PeerID")
	return ""
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
	ownerKey, err := evmks.CreateTxKey(ks, "test-tx-key")
	require.NoError(t, err)

	var oracles []confighelper.OracleIdentityExtra
	var offchainKeyrings []ocrtypes.OffchainKeyring
	var onchainKeyrings []ocrtypes.OnchainKeyring
	for i := 0; i < 4; i++ {
		onchainKeyring, err := keysv2.CreateOCR2OnchainKeyring(context.Background(), ks, fmt.Sprintf("test-onchain-keyring-%d", i))
		require.NoError(t, err)
		offchainKeyring, err := keysv2.CreateOCR2OffchainKeyring(context.Background(), ks, fmt.Sprintf("test-offchain-keyring-%d", i))
		require.NoError(t, err)
		keys, err := ks.CreateKeys(context.Background(), commonks.CreateKeysRequest{
			Keys: []commonks.CreateKeyRequest{
				{KeyName: fmt.Sprintf("test-peer-id-%d", i), KeyType: commonks.Ed25519},
				{KeyName: fmt.Sprintf("test-transmit-key-%d", i), KeyType: commonks.ECDSA_S256},
			},
		})
		require.NoError(t, err)
		require.Equal(t, 2, len(keys.Keys))
		peerID, err := ragep2ptypes.PeerIDFromPublicKey(ed25519.PublicKey(keys.Keys[0].KeyInfo.PublicKey))
		require.NoError(t, err)
		transmitKey, err := gethcrypto.UnmarshalPubkey(keys.Keys[1].KeyInfo.PublicKey)
		require.NoError(t, err)
		transmitAccount := gethcrypto.PubkeyToAddress(*transmitKey)
		require.NoError(t, err)
		oracles = append(oracles, confighelper.OracleIdentityExtra{
			OracleIdentity: confighelper.OracleIdentity{
				OnchainPublicKey:  onchainKeyring.PublicKey(),
				OffchainPublicKey: offchainKeyring.OffchainPublicKey(),
				PeerID:            peerID.String(),
				TransmitAccount:   ocrtypes.Account(transmitAccount.String()),
			},
			ConfigEncryptionPublicKey: offchainKeyring.ConfigEncryptionPublicKey(),
		})
		offchainKeyrings = append(offchainKeyrings, offchainKeyring)
		onchainKeyrings = append(onchainKeyrings, onchainKeyring)
	}
	backend := simulated.NewBackend(types.GenesisAlloc{
		ownerKey.Address(): {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)), // 10 ETH
		},
		common.HexToAddress(string(oracles[0].OracleIdentity.TransmitAccount)): {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)), // 10 ETH
		},
		common.HexToAddress(string(oracles[1].OracleIdentity.TransmitAccount)): {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)), // 10 ETH
		},
		common.HexToAddress(string(oracles[2].OracleIdentity.TransmitAccount)): {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)), // 10 ETH
		},
		common.HexToAddress(string(oracles[3].OracleIdentity.TransmitAccount)): {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)), // 10 ETH
		},
	}, simulated.WithBlockGasLimit(10e6))
	defer backend.Close()

	opts, err := ownerKey.GetTransactOpts(ctx, big.NewInt(1337))
	require.NoError(t, err)
	aggAddress, tx, agg, err := ocr2agg.DeployOCR2Aggregator(opts, backend.Client(), common.HexToAddress("0x0"), big.NewInt(1), big.NewInt(10), common.HexToAddress("0x0"), common.HexToAddress("0x0"), 18, "Test")
	require.NoError(t, err)
	backend.Commit()
	receipt, err := backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper.ContractSetConfigArgsForEthereumIntegrationTest(
		oracles, 1, 1000000)
	require.NoError(t, err)
	onchainConfig, err = median.StandardOnchainConfigCodec{}.Encode(ctx, median.OnchainConfig{
		Min: big.NewInt(1),
		Max: big.NewInt(10),
	})
	require.NoError(t, err)

	tx, err = agg.SetConfig(opts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	require.NoError(t, err)
	backend.Commit()
	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	helper := helper{backend: backend, lggr: lggr, ocr2agg: agg, opts: opts}
	for i := range oracles {
		oracle, err := libocr.NewOracle(libocr.OCR2OracleArgs{
			BinaryNetworkEndpointFactory: &helper,
			ReportingPluginFactory: &median.NumericalMedianFactory{
				ContractTransmitter:                  &helper,
				DataSource:                           &helper,
				JuelsPerFeeCoinDataSource:            &helper,
				GasPriceSubunitsDataSource:           &helper,
				IncludeGasPriceSubunitsInObservation: false,
				Logger:                               logger.NewOCRWrapper(lggr, true, func(string) {}),
				OnchainConfigCodec:                   median.StandardOnchainConfigCodec{},
				ReportCodec:                          evmreportcodec.ReportCodec{},
				DeviationFunc:                        median.DefaultDeviationFunc,
			},
			ContractConfigTracker: &helper,
			ContractTransmitter:   &helper,
			Database:              &helper,
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
			OffchainKeyring: offchainKeyrings[i],
			OnchainKeyring:  onchainKeyrings[i],
		})
		require.NoError(t, err)
		err = oracle.Start()
		defer oracle.Close()
	}
	time.Sleep(10 * time.Second)
}
