package keysv2_test

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	commonks "github.com/smartcontractkit/chainlink-common/keystore"
	logger "github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/keysv2"
	evmks "github.com/smartcontractkit/chainlink-evm/pkg/keysv2"
	"github.com/smartcontractkit/libocr/commontypes"
	ocr2agg "github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
	"github.com/smartcontractkit/libocr/networking"
	nettypes "github.com/smartcontractkit/libocr/networking/types"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	median "github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median"
	"github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median/evmreportcodec"
	libocr "github.com/smartcontractkit/libocr/offchainreporting2plus"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/chains/evmutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	ragetypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"github.com/stretchr/testify/require"
)

var _ ocrtypes.ContractConfigTracker = (*helper)(nil)
var _ ocrtypes.ContractTransmitter = (*helper)(nil)
var _ median.DataSource = (*helper)(nil)
var _ ocrtypes.Database = (*helper)(nil)
var _ ragetypes.PeerKeyring = (*peerKeyring)(nil)
var _ nettypes.DiscovererDatabase = (*memoryDiscovererDatabase)(nil)

// peerKeyring wraps an Ed25519 private key for P2P networking
type peerKeyring struct {
	privKey ed25519.PrivateKey
}

func (k *peerKeyring) Sign(msg []byte) ([]byte, error) {
	return ed25519.Sign(k.privKey, msg), nil
}

func (k *peerKeyring) PublicKey() ragetypes.PeerPublicKey {
	pubKey := k.privKey.Public().(ed25519.PublicKey)
	var peerPubKey ragetypes.PeerPublicKey
	copy(peerPubKey[:], pubKey)
	return peerPubKey
}

// memoryDiscovererDatabase is a simple in-memory implementation of DiscovererDatabase
type memoryDiscovererDatabase struct {
	announcements map[string][]byte
}

func newMemoryDiscovererDatabase() *memoryDiscovererDatabase {
	return &memoryDiscovererDatabase{
		announcements: make(map[string][]byte),
	}
}

func (m *memoryDiscovererDatabase) StoreAnnouncement(ctx context.Context, peerID string, ann []byte) error {
	m.announcements[peerID] = ann
	return nil
}

func (m *memoryDiscovererDatabase) ReadAnnouncements(ctx context.Context, peerIDs []string) (map[string][]byte, error) {
	result := make(map[string][]byte)
	for _, peerID := range peerIDs {
		if ann, ok := m.announcements[peerID]; ok {
			result[peerID] = ann
		}
	}
	return result, nil
}

type helper struct {
	backend      *simulated.Backend
	lggr         logger.Logger
	ocr2agg      *ocr2agg.OCR2Aggregator
	opts         *bind.TransactOpts
	observeValue *big.Int    // Per-oracle value to observe
	backendMutex *sync.Mutex // Shared mutex for backend access
}

func (t *helper) Observe(ctx context.Context, repts ocrtypes.ReportTimestamp) (*big.Int, error) {
	t.lggr.Infow("Observe", "repts", repts, "value", t.observeValue)
	return t.observeValue, nil
}

func (t *helper) Transmit(
	ctx context.Context,
	reportContext ocrtypes.ReportContext,
	report ocrtypes.Report,
	signatures []ocrtypes.AttributedOnchainSignature,
) error {
	t.backendMutex.Lock()
	defer t.backendMutex.Unlock()

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
	_, err := t.ocr2agg.Transmit(t.opts,
		evmutil.RawReportContext(reportContext),
		report, rs, ss, vs)
	if err != nil {
		return err
	}
	t.backend.Commit()
	return nil
}

func (t *helper) LatestRoundRequested(ctx context.Context, _ time.Duration) (ocrtypes.ConfigDigest, uint32, uint8, error) {
	t.backendMutex.Lock()
	defer t.backendMutex.Unlock()
	t.lggr.Info("LatestRoundRequested")
	res, err := t.ocr2agg.LatestConfigDigestAndEpoch(&bind.CallOpts{Context: context.Background()})
	return ocrtypes.ConfigDigest(res.ConfigDigest), res.Epoch, 0, err
}

func (t *helper) LatestConfigDigestAndEpoch(ctx context.Context) (ocrtypes.ConfigDigest, uint32, error) {
	t.backendMutex.Lock()
	defer t.backendMutex.Unlock()
	t.lggr.Info("LatestConfigDigestAndEpoch")
	res, err := t.ocr2agg.LatestConfigDigestAndEpoch(&bind.CallOpts{Context: context.Background()})
	return ocrtypes.ConfigDigest(res.ConfigDigest), res.Epoch, err
}

func (t *helper) LatestTransmissionDetails(ctx context.Context) (ocrtypes.ConfigDigest, uint32, uint8, *big.Int, time.Time, error) {
	t.backendMutex.Lock()
	defer t.backendMutex.Unlock()
	t.lggr.Info("LatestTransmissionDetails")
	res, err := t.ocr2agg.LatestTransmissionDetails(&bind.CallOpts{Context: context.Background()})
	return ocrtypes.ConfigDigest(res.ConfigDigest), res.Epoch, res.Round, res.LatestAnswer, time.Unix(int64(res.LatestTimestamp), 0), err
}

func (t *helper) FromAccount(ctx context.Context) (ocrtypes.Account, error) {
	t.lggr.Info("FromAccount")
	return ocrtypes.Account(t.opts.From.String()), nil
}

func (t *helper) LatestBlockHeight(ctx context.Context) (uint64, error) {
	// Simulated backend doesn't handle concurrent calls well
	// Use mutex but keep original context
	t.backendMutex.Lock()
	defer t.backendMutex.Unlock()

	header, err := t.backend.Client().HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return header.Number.Uint64(), nil
}

func (t *helper) LatestConfig(ctx context.Context, changedInBlock uint64) (ocrtypes.ContractConfig, error) {
	t.backendMutex.Lock()
	defer t.backendMutex.Unlock()
	t.lggr.Info("LatestConfig", "changedInBlock", changedInBlock, "ocr2agg", t.ocr2agg)
	c, err := t.ocr2agg.FilterConfigSet(&bind.FilterOpts{Context: context.Background(), Start: uint64(changedInBlock)})
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
	t.backendMutex.Lock()
	defer t.backendMutex.Unlock()
	c, err := t.ocr2agg.LatestConfigDetails(&bind.CallOpts{Context: context.Background()})
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

// TestOCR2Keyring_Integration tests the OCR2 keyrings integration
// with libocr to ensure that the keyrings can actually be used
// to sign and verify reports.
func TestOCR2Keyring_Integration(t *testing.T) {
	lggr := logger.Test(t)
	storage := commonks.NewMemoryStorage()
	ctx := context.Background()
	ks, err := commonks.LoadKeystore(ctx, storage, commonks.EncryptionParams{
		Password:     "test-password",
		ScryptParams: commonks.FastScryptParams,
	})
	require.NoError(t, err)
	ownerKey, err := evmks.CreateTxKey(ks, "test-tx-key")
	require.NoError(t, err)

	// Create keyrings and identities for 4 oracles
	var oracles []confighelper.OracleIdentityExtra
	var offchainKeyrings []ocrtypes.OffchainKeyring
	var onchainKeyrings []ocrtypes.OnchainKeyring
	var peerIDs []string
	var p2pKeys []ed25519.PrivateKey
	var oracleTxOpts []*bind.TransactOpts

	for i := 0; i < 4; i++ {
		onchainKeyring, err := keysv2.CreateOCR2OnchainKeyring(ctx, ks, fmt.Sprintf("test-onchain-keyring-%d", i))
		require.NoError(t, err)
		offchainKeyring, err := keysv2.CreateOCR2OffchainKeyring(ctx, ks, fmt.Sprintf("test-offchain-keyring-%d", i))
		require.NoError(t, err)

		// Create P2P key for networking (not stored in keystore)
		_, p2pPrivKey, err := ed25519.GenerateKey(nil)
		require.NoError(t, err)
		p2pPubKey := p2pPrivKey.Public().(ed25519.PublicKey)
		p2pKeys = append(p2pKeys, p2pPrivKey)

		// Get peer ID from P2P public key
		peerID, err := ragetypes.PeerIDFromPublicKey(p2pPubKey)
		require.NoError(t, err)
		peerIDs = append(peerIDs, peerID.String())

		// Create transmit key in keystore
		txKeyName := fmt.Sprintf("test-transmit-key-%d", i)
		keys, err := ks.CreateKeys(ctx, commonks.CreateKeysRequest{
			Keys: []commonks.CreateKeyRequest{
				{KeyName: txKeyName, KeyType: commonks.ECDSA_S256},
			},
		})
		require.NoError(t, err)
		require.Equal(t, 1, len(keys.Keys))

		// Get transmit account
		transmitKey, err := gethcrypto.UnmarshalPubkey(keys.Keys[0].KeyInfo.PublicKey)
		require.NoError(t, err)
		transmitAccount := gethcrypto.PubkeyToAddress(*transmitKey)

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

		// Build per-oracle transact opts that sign via keystore using the transmit key
		chainID := big.NewInt(1337)
		oracleTxOpts = append(oracleTxOpts, &bind.TransactOpts{
			From: transmitAccount,
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				if address != transmitAccount {
					return nil, bind.ErrNotAuthorized
				}
				signer := types.LatestSignerForChainID(chainID)
				h := signer.Hash(tx)
				resp, err := ks.Sign(context.Background(), commonks.SignRequest{
					KeyName: txKeyName,
					Data:    h[:],
				})
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(signer, resp.Signature)
			},
		})
	}

	// Create simulated backend with funded accounts
	alloc := types.GenesisAlloc{
		ownerKey.Address(): {
			Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18)), // 100 ETH
		},
	}
	for _, oracle := range oracles {
		alloc[common.HexToAddress(string(oracle.OracleIdentity.TransmitAccount))] = types.Account{
			Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18)), // 100 ETH
		}
	}
	backend := simulated.NewBackend(alloc, simulated.WithBlockGasLimit(10e6))
	defer backend.Close()

	// Deploy OCR2Aggregator contract
	opts, err := ownerKey.GetTransactOpts(ctx, big.NewInt(1337))
	require.NoError(t, err)
	aggAddress, tx, agg, err := ocr2agg.DeployOCR2Aggregator(
		opts, backend.Client(),
		common.HexToAddress("0x0"), // link token (not used in test)
		big.NewInt(1),              // min answer
		big.NewInt(10),             // max answer
		common.HexToAddress("0x0"), // billing access controller
		common.HexToAddress("0x0"), // requester access controller
		18,                         // decimals
		"Test Aggregator",
	)
	require.NoError(t, err)
	backend.Commit()
	receipt, err := backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	lggr.Infow("Deployed OCR2Aggregator", "address", aggAddress.Hex())

	// Generate OCR config
	signers, transmitters, f, _, offchainConfigVersion, offchainConfig, err := confighelper.ContractSetConfigArgsForEthereumIntegrationTest(
		oracles, 1, 1000000)
	require.NoError(t, err)

	onchainConfig, err := median.StandardOnchainConfigCodec{}.Encode(ctx, median.OnchainConfig{
		Min: big.NewInt(1),
		Max: big.NewInt(10),
	})
	require.NoError(t, err)

	// Set config on contract
	tx, err = agg.SetConfig(opts, signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	require.NoError(t, err)
	backend.Commit()
	receipt, err = backend.Client().TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)
	lggr.Info("Set config on OCR2Aggregator")

	// Verify config was set correctly
	configDetails, err := agg.LatestConfigDetails(&bind.CallOpts{Context: ctx})
	require.NoError(t, err)
	lggr.Infow("Config details after SetConfig",
		"blockNumber", configDetails.BlockNumber,
		"configDigest", fmt.Sprintf("%x", configDetails.ConfigDigest),
		"configCount", configDetails.ConfigCount)
	require.NotEqual(t, uint32(0), configDetails.ConfigCount, "Config count should be > 0")

	// Create P2P peers for each oracle using networking.NewPeer
	basePort := 30000
	var peers []ocrtypes.BinaryNetworkEndpointFactory
	var bootstrapLocators []commontypes.BootstrapperLocator

	// Oracle 0 will be the bootstrap node
	bootstrapLocators = append(bootstrapLocators, commontypes.BootstrapperLocator{
		PeerID: peerIDs[0],
		Addrs:  []string{fmt.Sprintf("127.0.0.1:%d", basePort)},
	})

	// Create all P2P peers
	for i := 0; i < 4; i++ {
		port := basePort + i
		listenAddr := fmt.Sprintf("127.0.0.1:%d", port)

		// Create peer with networking package
		peer, err := networking.NewPeer(networking.PeerConfig{
			PeerKeyring:          &peerKeyring{privKey: p2pKeys[i]},
			Logger:               logger.NewOCRWrapper(lggr, true, func(string) {}),
			V2ListenAddresses:    []string{listenAddr},
			V2DeltaReconcile:     1 * time.Second,
			V2DeltaDial:          100 * time.Millisecond,
			V2DiscovererDatabase: newMemoryDiscovererDatabase(),
			V2EndpointConfig: networking.EndpointConfigV2{
				IncomingMessageBufferSize: 100,
				OutgoingMessageBufferSize: 100,
			},
		})
		require.NoError(t, err)
		defer func() { _ = peer.Close() }()
		peers = append(peers, peer.OCR2BinaryNetworkEndpointFactory())
		lggr.Infow("Started P2P peer", "oracle", i, "peerID", peerIDs[i], "listenAddr", listenAddr)
	}

	// Give peers time to discover each other
	time.Sleep(2 * time.Second)

	// Create shared mutex for backend access
	backendMutex := &sync.Mutex{}

	// Start auto-mining blocks in the simulated backend
	stopMining := make(chan struct{})
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				backendMutex.Lock()
				backend.Commit()
				backendMutex.Unlock()
			case <-stopMining:
				return
			}
		}
	}()
	defer close(stopMining)

	// Create helpers and oracles
	for i := 0; i < 4; i++ {
		h := &helper{
			backend:      backend,
			lggr:         lggr,
			ocr2agg:      agg,
			opts:         oracleTxOpts[i],
			observeValue: big.NewInt(int64(i + 1)), // Oracle i observes value i+1 (1, 2, 3, 4)
			backendMutex: backendMutex,
		}

		oracle, err := libocr.NewOracle(libocr.OCR2OracleArgs{
			BinaryNetworkEndpointFactory: peers[i],
			V2Bootstrappers:              bootstrapLocators,
			ReportingPluginFactory: &median.NumericalMedianFactory{
				ContractTransmitter:                  h,
				DataSource:                           h,
				JuelsPerFeeCoinDataSource:            h,
				GasPriceSubunitsDataSource:           h,
				IncludeGasPriceSubunitsInObservation: false,
				Logger:                               logger.NewOCRWrapper(h.lggr, true, func(string) {}),
				OnchainConfigCodec:                   median.StandardOnchainConfigCodec{},
				ReportCodec:                          evmreportcodec.ReportCodec{},
			},
			ContractConfigTracker: h,
			ContractTransmitter:   h,
			Database:              h,
			LocalConfig: ocrtypes.LocalConfig{
				BlockchainTimeout:                  30 * time.Second,
				ContractConfigConfirmations:        1,
				ContractConfigTrackerPollInterval:  2 * time.Second,
				ContractTransmitterTransmitTimeout: 30 * time.Second,
				DatabaseTimeout:                    30 * time.Second,
				DevelopmentMode:                    ocrtypes.EnableDangerousDevelopmentMode,
			},
			Logger:             logger.NewOCRWrapper(h.lggr, true, func(string) {}),
			MonitoringEndpoint: nil,
			OffchainConfigDigester: evmutil.EVMOffchainConfigDigester{
				ChainID:         uint64(1337),
				ContractAddress: aggAddress,
			},
			OffchainKeyring: offchainKeyrings[i],
			OnchainKeyring:  onchainKeyrings[i],
		})
		require.NoError(t, err)

		err = oracle.Start()
		require.NoError(t, err)
		defer oracle.Close()
		lggr.Infow("Started oracle", "index", i, "peerID", peerIDs[i])
	}

	// Poll for transmission with timeout
	lggr.Info("Waiting for OCR report transmission...")
	timeout := time.After(60 * time.Second)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	var latestAnswer *big.Int
	var foundTransmission bool

	for !foundTransmission {
		select {
		case <-timeout:
			t.Fatal("Timed out waiting for OCR report transmission after 60 seconds")
		case <-ticker.C:
			// Check latest answer (blocks are auto-mined in background)
			backendMutex.Lock()
			result, err := agg.LatestAnswer(&bind.CallOpts{Context: ctx})
			backendMutex.Unlock()

			if err != nil {
				lggr.Warnw("Error getting latest answer", "error", err)
				continue
			}

			// Check if we have a non-zero answer
			if result.Cmp(big.NewInt(0)) > 0 {
				latestAnswer = result
				foundTransmission = true
				lggr.Infow("Found transmission!", "latestAnswer", latestAnswer)
			} else {
				lggr.Info("No transmission yet, waiting...")
			}
		}
	}

	// Verify results
	// With values 1, 2, 3, 4, the median should be 2 or 3
	// (In OCR2 median plugin, for even number of values, it typically picks the lower middle value)
	expectedMedian := big.NewInt(2)
	expectedMedianAlt := big.NewInt(3) // Some median implementations use average or upper value

	require.NotNil(t, latestAnswer, "Latest answer should not be nil")

	// Check that answer is one of the expected medians
	isExpectedMedian := latestAnswer.Cmp(expectedMedian) == 0 || latestAnswer.Cmp(expectedMedianAlt) == 0
	require.True(t, isExpectedMedian,
		"Expected median to be 2 or 3, got %s", latestAnswer.String())

	lggr.Infow("Test completed successfully!",
		"latestAnswer", latestAnswer,
		"expectedMedian", expectedMedian)

	// Verify that the transaction succeeded (no revert)
	backendMutex.Lock()
	details, err := agg.LatestTransmissionDetails(&bind.CallOpts{Context: ctx})
	backendMutex.Unlock()
	require.NoError(t, err)
	require.NotEqual(t, uint32(0), details.Epoch, "Should have transmitted at least once")

	lggr.Infow("Transmission details",
		"configDigest", fmt.Sprintf("%x", details.ConfigDigest),
		"epoch", details.Epoch,
		"latestAnswer", details.LatestAnswer)
}
