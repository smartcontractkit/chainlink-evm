package mercury

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services/servicetest"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/llo-feeds/generated/verifier"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/llo-feeds/generated/verifier_proxy"
	evmclient "github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/heads/headstest"
	"github.com/smartcontractkit/chainlink-evm/pkg/logpoller"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
)

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
	b := simulated.NewBackend(types.GenesisAlloc{
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
