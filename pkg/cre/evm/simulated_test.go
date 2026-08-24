package evm

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	evmclient "github.com/smartcontractkit/chainlink-evm/pkg/client"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/standalone"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// TestStartSimulated covers what a run pointed at nothing gets: a chain that
// answers over both of its RPCs, produces blocks on its own, and funds an account
// that was not known when it started.
func TestStartSimulated(t *testing.T) {
	chain, err := StartSimulated(t.Context(), logger.Test(t), SimulatedConfig{BlockTime: 50 * time.Millisecond})
	require.NoError(t, err)
	t.Cleanup(func() { assert.NoError(t, chain.Close()) })

	account := common.HexToAddress("0x00000000000000000000000000000000000000ff")
	require.NoError(t, chain.Fund(t.Context(), account, big.NewInt(params.Ether)))

	for name, url := range map[string]string{"http": chain.HTTPURL, "websocket": chain.WSURL} {
		t.Run("reachable over "+name, func(t *testing.T) {
			client, err := ethclient.DialContext(t.Context(), url)
			require.NoError(t, err)
			t.Cleanup(client.Close)

			id, err := client.ChainID(t.Context())
			require.NoError(t, err)
			assert.Equal(t, chain.ChainID, id)

			balance, err := client.BalanceAt(t.Context(), account, nil)
			require.NoError(t, err)
			assert.Equal(t, big.NewInt(params.Ether), balance, "the funded account must hold what it was sent")
		})
	}

	t.Run("has a finalized block that is not genesis", func(t *testing.T) {
		// go-ethereum's development beacon finalizes on epochs of 32 blocks, so a chain
		// handed over at genesis has nothing for a reader to read up to: a log poller
		// asked to back-fill to block zero fails against its own tables.
		client, err := ethclient.DialContext(t.Context(), chain.HTTPURL)
		require.NoError(t, err)
		t.Cleanup(client.Close)

		finalized, err := client.HeaderByNumber(t.Context(), big.NewInt(int64(rpc.FinalizedBlockNumber)))
		require.NoError(t, err)
		assert.Positive(t, finalized.Number.Sign(), "the chain must be wound past genesis finality before it is handed over")
	})

	t.Run("mines what is sent to it, without anyone committing a block", func(t *testing.T) {
		// The question this answers: a transaction from the layer above - a transaction
		// manager, say - arrives over the RPC and sits in the pool. Nothing above calls
		// Commit, so if the chain only produced blocks when asked, every write would sit
		// there forever.
		client, err := ethclient.DialContext(t.Context(), chain.HTTPURL)
		require.NoError(t, err)
		t.Cleanup(client.Close)

		sent := common.HexToAddress("0x00000000000000000000000000000000000000ee")
		require.NoError(t, send(t, client, chain, sent, big.NewInt(1)))

		assert.Eventually(t, func() bool {
			balance, err := client.BalanceAt(t.Context(), sent, nil)
			return err == nil && balance.Sign() > 0
		}, 10*time.Second, 100*time.Millisecond, "the mining loop must include what was sent to the pool")
	})

	t.Run("produces blocks on its own", func(t *testing.T) {
		client, err := ethclient.DialContext(t.Context(), chain.HTTPURL)
		require.NoError(t, err)
		t.Cleanup(client.Close)

		first, err := client.BlockNumber(t.Context())
		require.NoError(t, err)

		assert.Eventually(t, func() bool {
			next, err := client.BlockNumber(t.Context())
			return err == nil && next > first
		}, 5*time.Second, 50*time.Millisecond, "nothing was sent, and the chain must move anyway")
	})
}

// TestEmbeddedStartsAChainWhenToldOfNone is the behaviour every consumer of this
// dependency gets: embed and configure nothing, and there is a chain.
func TestEmbeddedStartsAChainWhenToldOfNone(t *testing.T) {
	dep := Dependency(logger.Test(t))
	embedded := dep.ForEmbedding(0, 1)
	chains := dep.Simulated().ForEmbedding(0, 1)

	client, err := embedded.Get(t.Context(), standalone.CommonConfig{})
	require.NoError(t, err)
	t.Cleanup(client.Close)

	chain, err := chains.Get(t.Context(), standalone.CommonConfig{})
	require.NoError(t, err)
	require.NotNil(t, chain, "a run told of no chain must have been given one")
	t.Cleanup(func() { assert.NoError(t, chain.Close()) })

	assert.Equal(t, chain.ChainID, client.ConfiguredChainID(), "the client must be dialled at the chain that was started")

	head, err := client.HeadByNumber(t.Context(), nil)
	require.NoError(t, err)
	assert.Positive(t, head.Number, "and must be able to read it")

	// Every instance shares it: they are a DON on one chain, not one chain each.
	again, err := dep.Simulated().ForEmbedding(1, 2).Get(t.Context(), standalone.CommonConfig{})
	require.NoError(t, err)
	assert.Same(t, chain, again)
}

// TestEmbeddedUsesTheChainItWasGiven is the other half: anything said about a
// chain is a chain the run meant, and starting one instead would ignore it.
func TestEmbeddedUsesTheChainItWasGiven(t *testing.T) {
	dep := Dependency(logger.Test(t))
	embedded, ok := dep.ForEmbedding(0, 1).(*ClientDependency)
	require.True(t, ok)

	// As a flag would have: decoded into the settings this form registered.
	embedded.cfg.HTTPURLs = []string{"http://127.0.0.1:1"}
	embedded.cfg.ChainID = "11155111"

	// The pool dials lazily, so nothing here reaches that address: what is asserted is
	// which chain it is for, and that none was started to be that chain instead.
	client, err := embedded.Get(t.Context(), standalone.CommonConfig{})
	require.NoError(t, err)
	t.Cleanup(client.Close)

	assert.Equal(t, big.NewInt(11155111), client.ConfiguredChainID())
	assert.Nil(t, embedded.simulated.chain, "a run that named a chain must not have been given another")

	chain, err := dep.Simulated().ForEmbedding(0, 1).Get(t.Context(), standalone.CommonConfig{})
	require.NoError(t, err)
	assert.Nil(t, chain, "and must be told there is no simulated chain, rather than handed one")
}

// TestNamed covers the question the embedded form asks before starting a chain:
// did this run say where its chain is.
//
// It is asked of the URLs alone. Asking the settings as a whole - "is any field
// set" - reads a defaulted field as an answer, and a run that tuned its pool
// without naming a chain would be left with nothing to dial.
func TestNamed(t *testing.T) {
	assert.False(t, named(evmclient.Config{}), "nothing named is what starting a chain is for")
	assert.False(t, named(defaultConfig), "and the defaults name nothing")
	assert.False(t, named(evmclient.Config{SyncThreshold: 9}), "a tuned pool is not a chain")

	assert.True(t, named(evmclient.Config{HTTPURLs: []string{"http://example.com"}}))
	assert.True(t, named(evmclient.Config{WSURLs: []string{"ws://example.com"}}))
}

// TestClientIsDialledOnce is what every service resolving this dependency relies
// on: they share one pool, with one set of connections and one head
// subscription. Two would be two subscriptions to drop and twice the polling -
// which is what a chain that keeps resubscribing looks like from the outside.
func TestClientIsDialledOnce(t *testing.T) {
	dep := Dependency(logger.Test(t))
	dep.cfg.HTTPURLs = []string{"http://127.0.0.1:1"}
	dep.cfg.ChainID = "1337"

	first, err := dep.Get(t.Context(), standalone.CommonConfig{})
	require.NoError(t, err)
	t.Cleanup(first.Close)

	second, err := dep.Get(t.Context(), standalone.CommonConfig{})
	require.NoError(t, err)
	assert.Same(t, first, second)
}

// send submits a transfer over the RPC and returns once the node has it, without
// mining anything: what mines it is what is under test.
func send(t *testing.T, client *ethclient.Client, chain *SimulatedChain, to common.Address, amount *big.Int) error {
	t.Helper()

	nonce, err := client.PendingNonceAt(t.Context(), chain.Deployer())
	require.NoError(t, err)
	head, err := client.HeaderByNumber(t.Context(), nil)
	require.NoError(t, err)

	key, err := deployerKey()
	require.NoError(t, err)

	tip := big.NewInt(params.GWei)
	tx, err := gethtypes.SignNewTx(key, gethtypes.LatestSignerForChainID(chain.ChainID), &gethtypes.DynamicFeeTx{
		ChainID:   chain.ChainID,
		Nonce:     nonce,
		To:        &to,
		Value:     amount,
		Gas:       params.TxGas,
		GasTipCap: tip,
		GasFeeCap: new(big.Int).Add(tip, new(big.Int).Mul(head.BaseFee, big.NewInt(2))),
	})
	require.NoError(t, err)

	return client.SendTransaction(t.Context(), tx)
}
