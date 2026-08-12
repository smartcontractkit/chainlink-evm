package registry

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/standalone"
	commonregistry "github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/standalone/registry"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	evmclient "github.com/smartcontractkit/chainlink-evm/pkg/client"
)

// Config is what reading the registry off a chain needs beyond the chain itself: where the contract
// is. Everything about reaching the chain belongs to the EVM client this wraps.
type Config struct {
	Address string `usage:"address of the on-chain CapabilitiesRegistry (v2) contract" validate:"required" example:"'0xYourRegistryAddress'"`
}

// Dependency returns a standalone.BootstrapDependency that resolves a registry Reader backed by the
// on-chain CapabilitiesRegistry, read through evm.
//
// It takes the EVM client as a dependency rather than settings of its own, so a binary that already
// talks to a chain does not configure it twice, and this stays the only place that knows the
// registry is on a chain at all: what it resolves to is chainlink-common's registry.Reader, which
// says nothing about EVM.
func Dependency(lggr logger.Logger, evm standalone.BootstrapDependency[evmclient.Client]) standalone.BootstrapDependency[commonregistry.Reader] {
	// Wrapped in OnceBootstrapper so the contract is bound once however many services resolve it.
	return standalone.OnceBootstrapper[commonregistry.Reader](&dependency{lggr: lggr, evm: evm})
}

type dependency struct {
	lggr logger.Logger
	evm  standalone.BootstrapDependency[evmclient.Client]
	cfg  Config
}

var _ standalone.BootstrapDependency[commonregistry.Reader] = (*dependency)(nil)

// Namespace groups the setting under capabilities-registry.*
// (--capabilities-registry.address, CRE_CAPABILITIES_REGISTRY_ADDRESS).
func (d *dependency) Namespace() string { return "capabilities-registry" }

func (d *dependency) Config() any { return &d.cfg }

func (d *dependency) Dependencies() []standalone.BootstrapCommand {
	return []standalone.BootstrapCommand{d.evm}
}

// ForEmbedding returns the receiver: every embedded instance reads the same registry, since they
// are members of the same DONs rather than separate deployments.
func (d *dependency) ForEmbedding(int) standalone.BootstrapDependency[commonregistry.Reader] {
	return d
}

func (d *dependency) Get(ctx context.Context, cc standalone.CommonConfig) (commonregistry.Reader, error) {
	if !common.IsHexAddress(d.cfg.Address) {
		return nil, fmt.Errorf("--capabilities-registry.address must be a hex address, got %q", d.cfg.Address)
	}

	client, err := d.evm.Get(ctx, cc)
	if err != nil {
		return nil, fmt.Errorf("failed to get evm client: %w", err)
	}

	return NewReader(d.lggr, client, common.HexToAddress(d.cfg.Address))
}
