package evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// SimulatedChain is a development chain running in this process.
//
// It exists because an embedded run has nothing to point at: it derives its keys
// from instance indices and runs its oracles as goroutines, and until now the one
// thing it could not derive was a chain. So every local run began by starting an
// anvil, funding an account on it and deploying to it by hand.
//
// The chain itself is in memory - state, blocks and the EVM. What it serves is an
// RPC, over HTTP and websockets, because what runs above it is the ordinary
// client: the same multinode pool, dialling the same way, subscribing to heads
// the same way, failing the same way. A client handed an in-process backend
// instead would make a local run prove less than it looks like it proves.
type SimulatedChain struct {
	// HTTPURL and WSURL are the chain's RPC, on ports the operating system chose
	// unless the run named them. Both are served: heads are a subscription when there
	// is a websocket, which is what a chain with one does.
	HTTPURL string
	WSURL   string

	// ChainID is what it answers eth_chainId with.
	ChainID *big.Int

	backend *simulated.Backend
	auth    *bind.TransactOpts
	stop    context.CancelFunc
	done    chan struct{}
	lggr    logger.Logger
}

// SimulatedConfig is what a simulated chain can be told. Everything here has a
// default that works, since the point of it is a run that is told nothing at all.
type SimulatedConfig struct {
	// Host and Port are where the RPC listens. A port of zero - the default - takes
	// whatever is free: the chain is in this process, and its port is nobody else's
	// business unless a run says otherwise, which is what naming one is for.
	Host string `usage:"host the simulated chain's RPC listens on"`
	Port uint16 `usage:"port the simulated chain's HTTP RPC listens on; 0 takes a free one"`
	// WSPort is the same for the websocket the client subscribes to heads on.
	WSPort uint16 `usage:"port the simulated chain's websocket RPC listens on; 0 takes a free one"`

	// ChainID is what it answers as. It defaults to the one local development
	// environments use, so what runs against those runs here unchanged.
	ChainID uint64 `usage:"chain ID the simulated chain answers as"`

	// BlockTime is how often it produces a block, whether or not anything was sent: a
	// head tracker wants heads, and a log poller wants a finalized block to read up to.
	BlockTime time.Duration `usage:"how often the simulated chain produces a block"`

	// FundingETH is what an account funded on this chain is given. It is deliberately
	// absurd: nothing here is worth anything, and a local run that stops because an
	// account ran out of gas is an afternoon wasted.
	FundingETH uint64 `usage:"ether an account is given when it is funded on the simulated chain"`
}

// DefaultSimulatedConfig is what a run that says nothing gets.
var DefaultSimulatedConfig = SimulatedConfig{
	Host:       "127.0.0.1",
	ChainID:    1337,
	BlockTime:  time.Second,
	FundingETH: 1_000_000,
}

// deployerSeed derives the account that funds everything else and deploys
// whatever a run puts on this chain. It is derived rather than generated so that
// what it deploys lands at the same address on every run - a contract's address
// is this account and a nonce - which is what lets a caller name one before it
// exists.
const deployerSeed = "cre/simulated/deployer"

// blocksBeforeReadingCap bounds how far the chain is taken before it is handed
// over. See advancePastGenesisFinality, which is what does the taking; this is
// only here so that a go-ethereum whose finality never arrives is a chain that
// starts anyway rather than a process that hangs.
const blocksBeforeReadingCap = 256

// StartSimulated brings up the chain and returns it running: mining on its own,
// with the deployer funded, and serving RPC.
func StartSimulated(ctx context.Context, lggr logger.Logger, cfg SimulatedConfig) (*SimulatedChain, error) {
	cfg = cfg.withDefaults()

	deployer, err := deployerKey()
	if err != nil {
		return nil, err
	}
	chainID := new(big.Int).SetUint64(cfg.ChainID)

	auth, err := bind.NewKeyedTransactorWithChainID(deployer, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to build the simulated chain's deployer: %w", err)
	}

	// Everything this chain will ever hand out, held by one account: an address is
	// funded by a transfer from here rather than at genesis, since who needs funding is
	// something a run discovers as it starts rather than knows beforehand.
	treasury := new(big.Int).Mul(new(big.Int).SetUint64(cfg.FundingETH), big.NewInt(params.Ether))
	treasury.Mul(treasury, big.NewInt(1_000))

	// Resolved before starting, because go-ethereum panics rather than returning when
	// its node cannot listen: a port already taken - a second run of this, or the first
	// still shutting down - would otherwise take the process down with a stack trace.
	httpPort, err := listenPort(cfg.Host, cfg.Port)
	if err != nil {
		return nil, err
	}
	wsPort, err := listenPort(cfg.Host, cfg.WSPort)
	if err != nil {
		return nil, err
	}

	alloc := gethtypes.GenesisAlloc{crypto.PubkeyToAddress(deployer.PublicKey): {Balance: treasury}}
	backend := simulated.NewBackend(alloc, func(nodeConf *node.Config, ethConf *ethconfig.Config) {
		nodeConf.HTTPHost, nodeConf.HTTPPort = cfg.Host, int(httpPort)
		nodeConf.HTTPModules = []string{"eth", "net", "web3", "debug", "txpool"}
		nodeConf.WSHost, nodeConf.WSPort = cfg.Host, int(wsPort)
		nodeConf.WSModules = nodeConf.HTTPModules

		// The chain config is copied before its ID is set: it is a package-level value in
		// go-ethereum, and renaming the chain for everything else in the process would be a
		// surprising thing for this to do.
		chainConfig := *params.AllDevChainProtocolChanges
		chainConfig.ChainID = chainID
		ethConf.Genesis = &core.Genesis{
			Config:   &chainConfig,
			GasLimit: ethConf.Miner.GasCeil,
			Alloc:    alloc,
		}
	})

	chain := &SimulatedChain{
		HTTPURL: fmt.Sprintf("http://%s:%d", cfg.Host, httpPort),
		WSURL:   fmt.Sprintf("ws://%s:%d", cfg.Host, wsPort),
		ChainID: chainID,
		backend: backend,
		auth:    auth,
		lggr:    lggr,
	}

	if err := advancePastGenesisFinality(ctx, chain); err != nil {
		_ = chain.Close()
		return nil, err
	}

	// Mined from here on, so that what runs above behaves as it does against a chain
	// that produces blocks whether or not anyone is sending to it.
	mining, stop := context.WithCancel(context.WithoutCancel(ctx))
	chain.stop, chain.done = stop, make(chan struct{})
	go func() {
		defer close(chain.done)

		ticker := time.NewTicker(cfg.BlockTime)
		defer ticker.Stop()
		for {
			select {
			case <-mining.Done():
				return
			case <-ticker.C:
				backend.Commit()
			}
		}
	}()

	lggr.Infow("Started a simulated chain",
		"httpURL", chain.HTTPURL, "wsURL", chain.WSURL, "chainID", chainID,
		"blockTime", cfg.BlockTime, "deployer", chain.Deployer())

	return chain, nil
}

func (c SimulatedConfig) withDefaults() SimulatedConfig {
	if c.Host == "" {
		c.Host = DefaultSimulatedConfig.Host
	}
	if c.ChainID == 0 {
		c.ChainID = DefaultSimulatedConfig.ChainID
	}
	if c.BlockTime <= 0 {
		c.BlockTime = DefaultSimulatedConfig.BlockTime
	}
	if c.FundingETH == 0 {
		c.FundingETH = DefaultSimulatedConfig.FundingETH
	}
	return c
}

// advancePastGenesisFinality takes the chain far enough that it has a finalized
// block that is not genesis.
//
// It matters because go-ethereum's development beacon finalizes on epochs of 32
// blocks: until the 33rd, the finalized block is block zero. Everything that
// reads a chain reads up to its finalized block, and a log poller asked to
// back-fill to block zero fails - its tables do not have a row for a block that
// is only a genesis hash. So a chain handed over at genesis is one every reader
// starts by failing against, retrying, and eventually recovering from a minute
// later, for no reason other than that nobody wound it forward.
//
// Blocks are committed rather than waited for, so this costs milliseconds; the
// mining loop that follows is paced for what runs above it, not for this.
func advancePastGenesisFinality(ctx context.Context, chain *SimulatedChain) error {
	client := chain.backend.Client()

	for range blocksBeforeReadingCap {
		finalized, err := client.HeaderByNumber(ctx, big.NewInt(int64(rpc.FinalizedBlockNumber)))
		if err == nil && finalized != nil && finalized.Number.Sign() > 0 {
			chain.lggr.Debugw("The simulated chain has a finalized block to read up to", "finalized", finalized.Number)
			return nil
		}
		// Not an error worth reporting: before the first epoch there is no finalized
		// block to answer with, which is exactly the state being wound past.
		chain.backend.Commit()
	}

	return fmt.Errorf("the simulated chain has no finalized block after %d blocks", blocksBeforeReadingCap)
}

// Deployer is the account that funds and deploys on this chain.
func (c *SimulatedChain) Deployer() common.Address { return c.auth.From }

// Transactor is what a caller deploying a contract on this chain signs with. It
// is the deployer's, so contracts land at addresses derived from one account.
func (c *SimulatedChain) Transactor() *bind.TransactOpts { return c.auth }

// Backend is this chain as go-ethereum's contract bindings take it, for deploying
// and calling directly rather than through everything the client is.
func (c *SimulatedChain) Backend() simulated.Client { return c.backend.Client() }

// Fund gives address the run's funding amount, and waits for it to land.
//
// It is a transfer rather than a genesis allocation because who needs funding is
// something a run works out as it starts: an embedded instance's account comes
// from its keystore, which is resolved after this chain is already up.
func (c *SimulatedChain) Fund(ctx context.Context, address common.Address, amount *big.Int) error {
	client := c.backend.Client()

	nonce, err := client.PendingNonceAt(ctx, c.Deployer())
	if err != nil {
		return fmt.Errorf("failed to read the deployer's nonce: %w", err)
	}
	head, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to read the chain's head: %w", err)
	}

	// A development chain's base fee is whatever the last block set; the tip is the
	// miner's minimum. Doubling the base fee leaves room for the next block to raise it.
	tip := big.NewInt(params.GWei)
	feeCap := new(big.Int).Add(tip, new(big.Int).Mul(head.BaseFee, big.NewInt(2)))

	tx, err := gethtypes.SignNewTx(deployerPrivateKey(c.auth), gethtypes.LatestSignerForChainID(c.ChainID), &gethtypes.DynamicFeeTx{
		ChainID:   c.ChainID,
		Nonce:     nonce,
		To:        &address,
		Value:     amount,
		Gas:       params.TxGas,
		GasTipCap: tip,
		GasFeeCap: feeCap,
	})
	if err != nil {
		return fmt.Errorf("failed to sign the funding transaction for %s: %w", address, err)
	}

	if err := client.SendTransaction(ctx, tx); err != nil {
		return fmt.Errorf("failed to send %s its funding: %w", address, err)
	}
	if err := c.Mined(ctx, tx); err != nil {
		return fmt.Errorf("failed to fund %s: %w", address, err)
	}

	c.lggr.Infow("Funded an account on the simulated chain", "account", address, "wei", amount)
	return nil
}

// Mined commits a block and waits for tx to land in it, which is what setting a
// chain up looks like: the mining loop is paced for what runs afterwards, and a
// setup that waited for it would just start slower.
func (c *SimulatedChain) Mined(ctx context.Context, tx *gethtypes.Transaction) error {
	c.backend.Commit()

	receipt, err := bind.WaitMined(ctx, c.backend.Client(), tx)
	if err != nil {
		return err
	}
	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		return errors.New("the transaction was mined but reverted")
	}
	return nil
}

// FundingAmount is what Fund gives an account when a caller has no figure of its
// own: this chain's whole point is that nothing runs out of gas.
func (c SimulatedConfig) FundingAmount() *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(c.withDefaults().FundingETH), big.NewInt(params.Ether))
}

// Close stops mining and shuts the chain down. Nothing outlives it: the state is
// in memory, and what was deployed on it was deployed for this run.
func (c *SimulatedChain) Close() error {
	if c.stop != nil {
		c.stop()
		<-c.done
	}
	return c.backend.Close()
}

func deployerKey() (*ecdsa.PrivateKey, error) {
	key, err := crypto.ToECDSA(crypto.Keccak256([]byte(deployerSeed)))
	if err != nil {
		return nil, fmt.Errorf("failed to derive the simulated chain's deployer key: %w", err)
	}
	return key, nil
}

// deployerPrivateKey is the key behind auth. It is re-derived rather than carried
// so that the key itself lives in one place, and it is a constant of this package
// anyway - a simulated chain holds nothing worth protecting.
func deployerPrivateKey(*bind.TransactOpts) *ecdsa.PrivateKey {
	key, _ := deployerKey()
	return key
}

// listenPort is the port to serve on: the one asked for, checked, or - for zero -
// one the operating system says is free.
//
// The listener is closed again rather than handed over, since go-ethereum opens
// its own. There is a moment in between, which is a race with anything else on
// this machine looking for a port; the alternative is reaching into a node the
// simulated backend does not expose.
func listenPort(host string, port uint16) (uint16, error) {
	listener, err := net.Listen("tcp", net.JoinHostPort(host, strconv.Itoa(int(port))))
	if err != nil {
		return 0, fmt.Errorf("the simulated chain cannot listen on %s:%d: %w (another run of this? name a free port)", host, port, err)
	}
	defer listener.Close()

	return uint16(listener.Addr().(*net.TCPAddr).Port), nil //#nosec G115 - a port is 16 bits
}
