// Package registry reads the on-chain CapabilitiesRegistry for a standalone binary.
//
// It is the chain half of chainlink-common's standalone/registry: that package decides when to
// read, what to do with a snapshot and who may see it, none of which needs a chain; this one turns
// the contract into a snapshot, which is the only part that does. So the only thing here is three
// paged view calls and the decoding of what they return.
//
// The reads go through the generated gethwrappers and an EVM client directly. There is no relayer,
// no ContractReader, no codec and no ReadIdentifier indirection: a binary that only wants to know
// who is in which DON should not have to stand up a relayer to find out.
package registry

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	ragetypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/standalone/registry"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	capregv2 "github.com/smartcontractkit/chainlink-evm/gethwrappers/workflow/generated/capabilities_registry_wrapper_v2"
)

// pageLimit bounds each view call. The v2 contract's getters are paginated; this reads one large
// page rather than implementing pagination, matching chainlink's registrysyncer/v2. A registry that
// ever exceeds it produces a warning rather than a silent truncation.
const pageLimit = 1024

// contractCapabilityType is the on-chain capability type enum, in the CapabilitiesRegistry
// contract's ordering.
type contractCapabilityType uint8

const (
	contractCapabilityTypeTrigger contractCapabilityType = iota
	contractCapabilityTypeAction
	contractCapabilityTypeConsensus
	contractCapabilityTypeTarget
)

// capabilityMetadata is the JSON blob the v2 contract stores per capability.
type capabilityMetadata struct {
	CapabilityType uint8 `json:"capabilityType"`
	ResponseType   uint8 `json:"responseType"`
}

// registryCaller is the subset of the generated v2 wrapper this reader uses. Narrowing it keeps the
// reader testable without a chain.
type registryCaller interface {
	GetCapabilities(opts *bind.CallOpts, start, limit *big.Int) ([]capregv2.CapabilitiesRegistryCapabilityInfo, error)
	GetDONs(opts *bind.CallOpts, start, limit *big.Int) ([]capregv2.CapabilitiesRegistryDONInfo, error)
	GetNodes(opts *bind.CallOpts, start, limit *big.Int) ([]capregv2.INodeInfoProviderNodeInfo, error)
}

// Reader reads the on-chain CapabilitiesRegistry.
type Reader struct {
	lggr   logger.Logger
	caller registryCaller
}

var _ registry.Reader = (*Reader)(nil)

// NewReader binds the CapabilitiesRegistry at registryAddress on an already-dialed backend.
//
// backend is typically chainlink-evm's multinode-backed client.Client, which satisfies
// bind.ContractCaller; node health, dead node declaration and primary selection come from there
// rather than being reimplemented here.
func NewReader(
	lggr logger.Logger,
	backend bind.ContractCaller,
	registryAddress common.Address,
) (*Reader, error) {
	caller, err := capregv2.NewCapabilitiesRegistryCaller(registryAddress, backend)
	if err != nil {
		return nil, fmt.Errorf("failed to bind CapabilitiesRegistry at %s: %w", registryAddress, err)
	}
	return &Reader{
		lggr:   logger.Named(lggr, "CapabilitiesRegistryReader"),
		caller: caller,
	}, nil
}

// Read returns the whole registry as one snapshot: capabilities, DONs and nodes, in three view
// calls decoded by the wrapper.
func (r *Reader) Read(ctx context.Context) (*registry.Snapshot, error) {
	opts := &bind.CallOpts{Context: ctx}
	start, limit := big.NewInt(0), big.NewInt(pageLimit)

	caps, err := r.caller.GetCapabilities(opts, start, limit)
	if err != nil {
		return nil, fmt.Errorf("getCapabilities: %w", err)
	}
	r.warnIfTruncated("capabilities", len(caps))

	idsToCapabilities := make(map[string]registry.Capability, len(caps))
	for _, c := range caps {
		if c.IsDeprecated {
			continue
		}
		capType, err := parseCapabilityType(c.Metadata)
		if err != nil {
			r.lggr.Warnw("failed to parse capability metadata, skipping",
				"capabilityID", c.CapabilityId, "err", err)
			continue
		}
		idsToCapabilities[c.CapabilityId] = registry.Capability{ID: c.CapabilityId, CapabilityType: capType}
	}

	dons, err := r.caller.GetDONs(opts, start, limit)
	if err != nil {
		return nil, fmt.Errorf("getDONs: %w", err)
	}
	r.warnIfTruncated("dons", len(dons))

	idsToDONs := make(map[registry.DonID]registry.DON, len(dons))
	for _, d := range dons {
		cfgs := make(map[string]registry.CapabilityConfiguration, len(d.CapabilityConfigurations))
		for _, dc := range d.CapabilityConfigurations {
			cfgs[dc.CapabilityId] = registry.CapabilityConfiguration{Config: dc.Config}
		}
		idsToDONs[registry.DonID(d.Id)] = registry.DON{
			DON:                      toDON(d),
			CapabilityConfigurations: cfgs,
		}
	}

	nodes, err := r.caller.GetNodes(opts, start, limit)
	if err != nil {
		return nil, fmt.Errorf("getNodes: %w", err)
	}
	r.warnIfTruncated("nodes", len(nodes))

	idsToNodes := make(map[ragetypes.PeerID]registry.NodeInfo, len(nodes))
	for _, n := range nodes {
		idsToNodes[n.P2pId] = registry.NodeInfo{
			NodeOperatorID:      n.NodeOperatorId,
			ConfigCount:         n.ConfigCount,
			WorkflowDONID:       n.WorkflowDONId,
			Signer:              n.Signer,
			P2pID:               n.P2pId,
			EncryptionPublicKey: n.EncryptionPublicKey,
			CsaKey:              n.CsaKey,
			CapabilityIDs:       n.CapabilityIds,
		}
	}

	return &registry.Snapshot{
		DONs:         idsToDONs,
		Nodes:        idsToNodes,
		Capabilities: idsToCapabilities,
	}, nil
}

// warnIfTruncated makes a hit against pageLimit visible instead of letting a full page read as
// "that is all there is".
func (r *Reader) warnIfTruncated(what string, n int) {
	if n >= pageLimit {
		r.lggr.Warnw("registry read hit the page limit; results may be truncated",
			"what", what, "count", n, "limit", pageLimit)
	}
}

func toDON(d capregv2.CapabilitiesRegistryDONInfo) capabilities.DON {
	members := make([]ragetypes.PeerID, 0, len(d.NodeP2PIds))
	for _, p := range d.NodeP2PIds {
		members = append(members, p)
	}
	return capabilities.DON{
		Name:             d.Name,
		ID:               d.Id,
		Families:         d.DonFamilies,
		ConfigVersion:    d.ConfigCount,
		Members:          members,
		F:                d.F,
		IsPublic:         d.IsPublic,
		AcceptsWorkflows: d.AcceptsWorkflows,
		Config:           d.Config,
	}
}

func parseCapabilityType(metadata []byte) (capabilities.CapabilityType, error) {
	if len(metadata) == 0 {
		return capabilities.CapabilityTypeUnknown, errors.New("metadata is empty")
	}
	var meta capabilityMetadata
	if err := json.Unmarshal(metadata, &meta); err != nil {
		return capabilities.CapabilityTypeUnknown, fmt.Errorf("invalid metadata: %w", err)
	}
	switch contractCapabilityType(meta.CapabilityType) {
	case contractCapabilityTypeTrigger:
		return capabilities.CapabilityTypeTrigger, nil
	case contractCapabilityTypeAction:
		return capabilities.CapabilityTypeAction, nil
	case contractCapabilityTypeConsensus:
		return capabilities.CapabilityTypeConsensus, nil
	case contractCapabilityTypeTarget:
		return capabilities.CapabilityTypeTarget, nil
	default:
		return capabilities.CapabilityTypeUnknown, fmt.Errorf("unknown on-chain capability type %d", meta.CapabilityType)
	}
}
