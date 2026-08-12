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
func Dependency(lggr logger.Logger) standalone.BootstrapDependency[evmclient.Client] {
	return standalone.OnceBootstrapper[evmclient.Client](&dependency{lggr: lggr, cfg: defaultConfig})
}

type dependency struct {
	lggr logger.Logger

	client evmclient.Client
	cfg    evmclient.Config
}

var _ standalone.BootstrapDependency[evmclient.Client] = (*dependency)(nil)

func (d *dependency) Namespace() string { return "evm" }

func (d *dependency) Config() any {
	return &d.cfg
}

func (d *dependency) Dependencies() []standalone.BootstrapCommand {
	return []standalone.BootstrapCommand{}
}

func (d *dependency) ForEmbedding(_ int) standalone.BootstrapDependency[evmclient.Client] { return d }

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
