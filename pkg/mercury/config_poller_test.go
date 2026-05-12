package mercury

import (
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/onsi/gomega"
	"github.com/pkg/errors"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2plus/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/umbracle/ethgo/abi"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/llo-feeds/generated/verifier"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/llo-feeds/generated/verifier_proxy"
	evmclient "github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/heads/headstest"
	"github.com/smartcontractkit/chainlink-evm/pkg/logpoller"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	evmutils "github.com/smartcontractkit/chainlink-evm/pkg/utils"
)

func TestMercuryConfigPoller(t *testing.T) {
	feedID := evmutils.NewHash()
	feedIDBytes := [32]byte(feedID)

	th := SetupTH(t, feedID)

	notify := th.configPoller.Notify()
	assert.Empty(t, notify)

	// Should have no config to begin with.
	_, config, err := th.configPoller.LatestConfigDetails(t.Context())
	require.NoError(t, err)
	require.Equal(t, ocrtypes2.ConfigDigest{}, config)

	// Create minimum number of nodes.
	n := 4
	var oracles []confighelper2.OracleIdentityExtra
	for range n {
		oracles = append(oracles, confighelper2.OracleIdentityExtra{
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  evmutils.RandomAddress().Bytes(),
				TransmitAccount:   ocrtypes2.Account(evmutils.RandomAddress().String()),
				OffchainPublicKey: evmutils.RandomBytes32(),
			},
			ConfigEncryptionPublicKey: evmutils.RandomBytes32(),
		})
	}
	f := uint8(1)
	// Setup config on contract
	configType := abi.MustNewType("tuple()")
	onchainConfigVal, err := abi.Encode(map[string]any{}, configType)
	require.NoError(t, err)
	signers, _, threshold, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgsForTests(
		2*time.Second,        // DeltaProgress
		20*time.Second,       // DeltaResend
		100*time.Millisecond, // DeltaRound
		0,                    // DeltaGrace
		1*time.Minute,        // DeltaStage
		100,                  // rMax
		[]int{len(oracles)},  // S
		oracles,
		[]byte{}, // reportingPluginConfig []byte,
		nil,
		0,                    // Max duration query
		250*time.Millisecond, // Max duration observation
		250*time.Millisecond, // MaxDurationReport
		250*time.Millisecond, // MaxDurationShouldAcceptFinalizedReport
		250*time.Millisecond, // MaxDurationShouldTransmitAcceptedReport
		int(f),               // f
		onchainConfigVal,
	)
	require.NoError(t, err)
	signerAddresses, err := onchainPublicKeyToAddress(signers)
	require.NoError(t, err)
	offchainTransmitters := make([][32]byte, n)
	encodedTransmitter := make([]ocrtypes2.Account, n)
	for i := range n {
		offchainTransmitters[i] = oracles[i].OffchainPublicKey
		encodedTransmitter[i] = ocrtypes2.Account(hex.EncodeToString(oracles[i].OffchainPublicKey[:]))
	}

	_, err = th.verifierContract.SetConfig(th.user, feedIDBytes, signerAddresses, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, nil)
	require.NoError(t, err, "failed to setConfig with feed ID")
	th.backend.Commit()

	latest, err := th.backend.Client().BlockByNumber(t.Context(), nil)
	require.NoError(t, err)
	// Ensure we capture this config set log.
	require.NoError(t, th.logPoller.Replay(t.Context(), latest.Number().Int64()-1))

	// Send blocks until we see the config updated.
	var configBlock uint64
	var digest [32]byte
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		th.backend.Commit()
		configBlock, digest, err = th.configPoller.LatestConfigDetails(t.Context())
		require.NoError(t, err)
		return ocrtypes2.ConfigDigest{} != digest
	}, testutils.WaitTimeout(t), 100*time.Millisecond).Should(gomega.BeTrue())

	// Assert the config returned is the one we configured.
	newConfig, err := th.configPoller.LatestConfig(t.Context(), configBlock)
	require.NoError(t, err)
	// Note we don't check onchainConfig, as that is populated in the contract itself.
	assert.Equal(t, digest, [32]byte(newConfig.ConfigDigest))
	assert.Equal(t, signers, newConfig.Signers)
	assert.Equal(t, threshold, newConfig.F)
	assert.Equal(t, encodedTransmitter, newConfig.Transmitters)
	assert.Equal(t, offchainConfigVersion, newConfig.OffchainConfigVersion)
	assert.Equal(t, offchainConfig, newConfig.OffchainConfig)
}

func onchainPublicKeyToAddress(publicKeys []types.OnchainPublicKey) (addresses []common.Address, err error) {
	for _, signer := range publicKeys {
		if len(signer) != 20 {
			return []common.Address{}, errors.Errorf("address is not 20 bytes %s", signer)
		}
		addresses = append(addresses, common.BytesToAddress(signer))
	}
	return addresses, nil
}

type TestHarness struct {
	configPoller     *ConfigPoller
	user             *bind.TransactOpts
	backend          *simulated.Backend
	verifierAddress  common.Address
	verifierContract *verifier.Verifier
	logPoller        logpoller.LogPoller
}

func SetupTH(t *testing.T, feedID common.Hash) TestHarness {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	user, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	require.NoError(t, err)
	b := simulated.NewBackend(ethtypes.GenesisAlloc{
		user.From: {Balance: big.NewInt(1000000000000000000)}},
		simulated.WithBlockGasLimit(5*ethconfig.Defaults.Miner.GasCeil))

	proxyAddress, _, verifierProxy, err := verifier_proxy.DeployVerifierProxy(user, b.Client(), common.Address{})
	require.NoError(t, err, "failed to deploy test mercury verifier proxy contract")
	b.Commit()
	verifierAddress, _, verifierContract, err := verifier.DeployVerifier(user, b.Client(), proxyAddress)
	require.NoError(t, err, "failed to deploy test mercury verifier contract")
	b.Commit()
	_, err = verifierProxy.InitializeVerifier(user, verifierAddress)
	require.NoError(t, err)
	b.Commit()

	db := testutils.NewSqlxDB(t)
	ethClient := evmclient.NewSimulatedBackendClient(t, b, big.NewInt(1337))
	lggr := logger.Test(t)
	lorm := logpoller.NewORM(big.NewInt(1337), db, lggr)

	lpOpts := logpoller.Opts{
		PollPeriod:               100 * time.Millisecond,
		FinalityDepth:            1,
		BackfillBatchSize:        2,
		RPCBatchSize:             2,
		KeepFinalizedBlocksDepth: 1000,
	}
	ht := headstest.NewSimulatedHeadTracker(ethClient, lpOpts.UseFinalityTag, lpOpts.FinalityDepth)
	lp := logpoller.NewLogPoller(lorm, ethClient, lggr, ht, lpOpts)
	servicetest.Run(t, lp)

	configPoller, err := NewConfigPoller(t.Context(), lggr, lp, verifierAddress, feedID)
	require.NoError(t, err)

	configPoller.Start()
	t.Cleanup(func() {
		assert.NoError(t, configPoller.Close())
	})

	return TestHarness{
		configPoller:     configPoller,
		user:             user,
		backend:          b,
		verifierAddress:  verifierAddress,
		verifierContract: verifierContract,
		logPoller:        lp,
	}
}
