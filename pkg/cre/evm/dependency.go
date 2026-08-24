package evm

import (
	"context"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	evmclient "github.com/smartcontractkit/chainlink-evm/pkg/client"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/standalone"
)

// Defaults for the settings evmclient.Config exposes, mirroring chainlink's own EVM chain
// defaults for a generic EVM chain at the time of creation.
var defaultConfig = evmclient.Config{
	ConfirmationTimeout:        *config.MustNewDuration(60 * time.Second),
	DeathDeclarationDelay:      *config.MustNewDuration(10 * time.Second),
	EnforceRepeatableRead:      true,
	FinalityTagEnabled:         true,
	FinalityDepth:              uint32(50),
	FinalizedBlockPollInterval: *config.MustNewDuration(5 * time.Second),
	NoNewHeadsThreshold:        *config.MustNewDuration(3 * time.Minute),
	PollInterval:               *config.MustNewDuration(10 * time.Second),
	PollFailureThreshold:       5,
	SelectionMode:              "HighestHead",
	SyncThreshold:              5,
}

// Dependency returns a standalone.BootstrapDependency that resolves a dialed,
// multinode-backed EVM client. The settings and the dialing belong to the client
// itself (evmclient.Config); this binds them as bootstrap configuration under
// evm.*, defaults them, and resolves the client once.
//
// An embedded run told about no chain at all starts one: see SimulatedChain, and
// Simulated for reaching it. That is what makes "embed" a command that needs
// nothing, which is what a local run, a test or a first look wants - and what
// every consumer of this dependency used to assemble by hand.
func Dependency(lggr logger.Logger) *ClientDependency {
	cfg := defaultConfig
	return &ClientDependency{lggr: lggr, cfg: &cfg}
}

// ClientDependency is the EVM client as a bootstrap dependency.
//
// It is a named type rather than the interface so that what it resolves besides
// the client - the chain it may have started - is reachable. See Simulated.
type ClientDependency struct {
	lggr logger.Logger

	// mu guards the resolved client, which every service asking for one shares.
	mu     sync.Mutex
	client evmclient.Client
	err    error

	cfg *evmclient.Config

	// simulated is the chain this dependency starts when it is given none, and nil for
	// a configured run, which was told where its chain is.
	simulated *simulatedDependency

	// embedded is the one embedded form. The bootstrapper asks for one per instance,
	// and once more - as (0, 1) - to collect the settings to register; every answer has
	// to be the same object, or what is read at startup is not what was bound to flags.
	embedded *ClientDependency
}

var _ standalone.BootstrapDependency[evmclient.Client] = (*ClientDependency)(nil)

func (d *ClientDependency) Namespace() string { return "evm" }

func (d *ClientDependency) Config() any { return d.cfg }

func (d *ClientDependency) Dependencies() []standalone.BootstrapCommand {
	if d.simulated == nil {
		return nil
	}
	return []standalone.BootstrapCommand{d.simulated}
}

// ForEmbedding returns the form that starts a chain when it is given none.
//
// Its settings are the same defaults a configured run gets, so an embedded run
// tunes its pool with the same flags and the same values. What differs is the
// URLs: they have none until either the run names one or this starts a chain to
// take their place.
func (d *ClientDependency) ForEmbedding(_, _ int) standalone.BootstrapDependency[evmclient.Client] {
	if d.embedded == nil {
		cfg := defaultConfig
		d.embedded = &ClientDependency{
			lggr:      d.lggr,
			cfg:       &cfg,
			simulated: &simulatedDependency{lggr: d.lggr, cfg: &SimulatedConfig{}},
		}
	}
	return d.embedded
}

// Get dials the chain, once, however many services resolve this: they are meant
// to share a pool - one set of connections, one head subscription - not open one
// each. Two pools against one chain is two subscriptions to drop and two nodes'
// worth of polling, which is what a caller sees as a chain that keeps
// resubscribing.
func (d *ClientDependency) Get(ctx context.Context, cc standalone.CommonConfig) (evmclient.Client, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.client != nil || d.err != nil {
		return d.client, d.err
	}

	d.client, d.err = d.dial(ctx, cc)
	return d.client, d.err
}

func (d *ClientDependency) dial(ctx context.Context, cc standalone.CommonConfig) (evmclient.Client, error) {
	cfg := *d.cfg

	if d.simulated != nil && !named(cfg) {
		chain, err := d.simulated.Get(ctx, cc)
		if err != nil {
			return nil, err
		}

		// Pointed at the chain that was just started, over both of the RPCs it serves:
		// given a websocket the pool subscribes to heads rather than polling for them,
		// which is what it does against any chain that has one.
		cfg.HTTPURLs = []string{chain.HTTPURL}
		cfg.WSURLs = []string{chain.WSURL}
		cfg.ChainID = chain.ChainID.String()
	}

	return evmclient.NewClientFromConfig(ctx, d.lggr, cfg)
}

func (d *ClientDependency) Close() {
	if d.client != nil {
		d.client.Close()
	}
}

// Simulated is the chain this dependency starts when it is given none, as a
// dependency of its own: nil for a configured run, and nil for an embedded run
// that was told where its chain is.
//
// It is how a consumer does to that chain what a deployment does to a real one -
// fund the accounts it will send from, deploy the contracts it will write through
// - without having to know which kind of run it is in.
func (d *ClientDependency) Simulated() standalone.BootstrapDependency[*SimulatedChain] {
	return &simulatedView{client: d}
}

// named reports whether the run said where its chain is.
//
// A URL is the question, and the only one: everything else here tunes how a chain
// is followed rather than which one it is, and a run that tuned its pool without
// naming a chain still wants one started for it. Asking the settings as a whole -
// "is any field set" - would answer it with a defaulted field, and would flip an
// explicit --evm.finality-tag-enabled=false back to true on the way past.
func named(cfg evmclient.Config) bool {
	return len(cfg.HTTPURLs) > 0 || len(cfg.WSURLs) > 0
}

// simulatedDependency starts the chain, once, however many ask for it: the
// instances of an embedded run are a DON on one chain, not one chain each.
type simulatedDependency struct {
	lggr logger.Logger
	cfg  *SimulatedConfig

	started bool
	chain   *SimulatedChain
	err     error
}

var _ standalone.BootstrapDependency[*SimulatedChain] = (*simulatedDependency)(nil)

func (d *simulatedDependency) Namespace() string { return "simulated" }

func (d *simulatedDependency) Config() any { return d.cfg }

func (d *simulatedDependency) Dependencies() []standalone.BootstrapCommand { return nil }

func (d *simulatedDependency) ForEmbedding(_, _ int) standalone.BootstrapDependency[*SimulatedChain] {
	return d
}

func (d *simulatedDependency) Get(ctx context.Context, _ standalone.CommonConfig) (*SimulatedChain, error) {
	if !d.started {
		d.started = true
		d.chain, d.err = StartSimulated(ctx, d.lggr, *d.cfg)
	}
	return d.chain, d.err
}

func (d *simulatedDependency) Close() {
	if d.chain != nil {
		_ = d.chain.Close()
	}
}

// simulatedView hands back the chain the client started, and nothing if it
// started none.
//
// It resolves through the client rather than the chain because whether there is
// one is the client's answer: it is what compares the configuration against
// nothing at all.
type simulatedView struct {
	client *ClientDependency
}

var _ standalone.BootstrapDependency[*SimulatedChain] = (*simulatedView)(nil)

// Namespace is empty and Config is nil: the settings belong to the client and to
// the chain, and both register their own.
func (d *simulatedView) Namespace() string { return "" }

func (d *simulatedView) Config() any { return nil }

func (d *simulatedView) Dependencies() []standalone.BootstrapCommand {
	return []standalone.BootstrapCommand{d.client}
}

func (d *simulatedView) ForEmbedding(i, instances int) standalone.BootstrapDependency[*SimulatedChain] {
	embedded, ok := d.client.ForEmbedding(i, instances).(*ClientDependency)
	if !ok {
		return d
	}
	return &simulatedView{client: embedded}
}

func (d *simulatedView) Get(ctx context.Context, cc standalone.CommonConfig) (*SimulatedChain, error) {
	if _, err := d.client.Get(ctx, cc); err != nil {
		return nil, err
	}
	if d.client.simulated == nil {
		return nil, nil
	}
	return d.client.simulated.chain, d.client.simulated.err
}
