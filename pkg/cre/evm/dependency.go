package evm

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	evmclient "github.com/smartcontractkit/chainlink-evm/pkg/client"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/standalone"
)

// Defaults for the settings evmclient.Config exposes, mirroring chainlink's own EVM chain
// defaults for a generic EVM chain, so a standalone binary polling the same RPCs behaves like
// the node does. They live here rather than in the client: this is the instance the flags are
// bound to and decoded into, so an unset setting keeps the value it was given here.
const (
	defaultFinalityDepth = uint32(50)
	defaultPollInterval  = 10 * time.Second
)

var defaultConfig = evmclient.Config{
	FinalityTagEnabled: true,
	FinalityDepth:      defaultFinalityDepth,
	PollInterval:       *config.MustNewDuration(defaultPollInterval),
}

// Dependency returns a standalone.BootstrapDependency that resolves a dialed,
// multinode-backed EVM client. The settings and the dialing belong to the client
// itself (evmclient.Config); this binds them as bootstrap configuration under
// evm.*, defaults them, and resolves the client once.
func Dependency(lggr logger.Logger) standalone.BootstrapDependency[evmclient.Client] {
	// Wrap in OnceBootstrapper so Get (which dials every configured RPC) runs at
	// most once even if several services resolve this dependency.
	return standalone.OnceBootstrapper[evmclient.Client](&dependency{lggr: lggr, cfg: defaultConfig})
}

type dependency struct {
	lggr logger.Logger

	client evmclient.Client
	cfg    evmclient.Config
}

var _ standalone.BootstrapDependency[evmclient.Client] = (*dependency)(nil)

// Namespace groups the EVM settings under evm.* (--evm.http-url, CRE_EVM_HTTP_URL).
func (d *dependency) Namespace() string { return "evm" }

func (d *dependency) Config() any {
	return &d.cfg
}

func (d *dependency) Dependencies() []standalone.BootstrapCommand {
	return []standalone.BootstrapCommand{}
}

// ForEmbedding returns the receiver, so every embedded instance shares one client: they are
// reading the same chain over the same RPCs, and a client per instance would multiply the request
// rate without telling any of them anything new.
func (d *dependency) ForEmbedding(int) standalone.BootstrapDependency[evmclient.Client] { return d }

func (d *dependency) Get(ctx context.Context, _ standalone.CommonConfig) (evmclient.Client, error) {
	cl, err := evmclient.NewClientFromConfig(ctx, d.lggr, d.cfg)
	if err != nil {
		return nil, err
	}
	d.client = cl
	return cl, nil
}

func (d *dependency) Close() {
	if d.client != nil {
		d.client.Close()
	}
}
