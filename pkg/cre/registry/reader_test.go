package registry

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ragetypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	capregv2 "github.com/smartcontractkit/chainlink-evm/gethwrappers/workflow/generated/capabilities_registry_wrapper_v2"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/standalone/registry"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// peer builds a recognisable peer ID for tests.
func peer(b byte) ragetypes.PeerID {
	var p ragetypes.PeerID
	p[0] = b
	return p
}

// fakeCaller stands in for the generated v2 contract caller.
type fakeCaller struct {
	mu sync.Mutex

	caps  []capregv2.CapabilitiesRegistryCapabilityInfo
	dons  []capregv2.CapabilitiesRegistryDONInfo
	nodes []capregv2.INodeInfoProviderNodeInfo

	capsErr  error
	donsErr  error
	nodesErr error

	calls int
	// lastStart/lastLimit record the pagination arguments actually sent.
	lastStart, lastLimit *big.Int
}

func (f *fakeCaller) GetCapabilities(_ *bind.CallOpts, start, limit *big.Int) ([]capregv2.CapabilitiesRegistryCapabilityInfo, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls++
	f.lastStart, f.lastLimit = start, limit
	return f.caps, f.capsErr
}

func (f *fakeCaller) GetDONs(*bind.CallOpts, *big.Int, *big.Int) ([]capregv2.CapabilitiesRegistryDONInfo, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.dons, f.donsErr
}

func (f *fakeCaller) GetNodes(*bind.CallOpts, *big.Int, *big.Int) ([]capregv2.INodeInfoProviderNodeInfo, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.nodes, f.nodesErr
}

func (f *fakeCaller) callCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.calls
}

// newTestReader builds a Reader around a fake caller, bypassing NewReader's contract binding
// (which needs a chain backend).
func newTestReader(t *testing.T, caller registryCaller) *Reader {
	t.Helper()
	return &Reader{lggr: logger.Test(t), caller: caller}
}

// snapshot reads through a fake caller and fails the test if the read does.
func snapshot(t *testing.T, caller registryCaller) *registry.Snapshot {
	t.Helper()
	snap, err := newTestReader(t, caller).Read(context.Background())
	require.NoError(t, err)
	return snap
}

func fullCaller() *fakeCaller {
	p1, p2 := peer(1), peer(2)
	return &fakeCaller{
		caps: []capregv2.CapabilitiesRegistryCapabilityInfo{
			{CapabilityId: "act@1.0.0", Metadata: []byte(`{"capabilityType":1}`)},
			{CapabilityId: "cron@1.0.0", Metadata: []byte(`{"capabilityType":0}`)},
		},
		dons: []capregv2.CapabilitiesRegistryDONInfo{{
			Id:               2,
			ConfigCount:      5,
			F:                1,
			IsPublic:         true,
			AcceptsWorkflows: true,
			NodeP2PIds:       [][32]byte{p1, p2},
			DonFamilies:      []string{"zone-a"},
			Name:             "cap-don",
			Config:           []byte("don-config"),
			CapabilityConfigurations: []capregv2.CapabilitiesRegistryCapabilityConfiguration{
				{CapabilityId: "act@1.0.0", Config: []byte("act-cfg")},
			},
		}},
		nodes: []capregv2.INodeInfoProviderNodeInfo{
			{NodeOperatorId: 11, ConfigCount: 1, WorkflowDONId: 2, P2pId: p1, Signer: [32]byte{0xaa}, CapabilityIds: []string{"act@1.0.0"}},
			{NodeOperatorId: 12, ConfigCount: 1, WorkflowDONId: 2, P2pId: p2},
		},
	}
}

func TestReader_ReadBuildsSnapshot(t *testing.T) {
	snap := snapshot(t, fullCaller())

	require.Len(t, snap.Capabilities, 2)
	assert.Equal(t, capabilities.CapabilityTypeAction, snap.Capabilities["act@1.0.0"].CapabilityType)
	assert.Equal(t, capabilities.CapabilityTypeTrigger, snap.Capabilities["cron@1.0.0"].CapabilityType)

	require.Len(t, snap.DONs, 1)
	don := snap.DONs[2]
	assert.Equal(t, "cap-don", don.Name)
	assert.Equal(t, uint8(1), don.F)
	// ConfigVersion comes from the contract's ConfigCount.
	assert.Equal(t, uint32(5), don.ConfigVersion)
	assert.Equal(t, []string{"zone-a"}, don.Families)
	assert.True(t, don.IsPublic)
	assert.True(t, don.AcceptsWorkflows)
	assert.Equal(t, []byte("don-config"), don.Config)
	require.Len(t, don.Members, 2)

	// Capability config is carried through undecoded.
	assert.Equal(t, []byte("act-cfg"), don.CapabilityConfigurations["act@1.0.0"])

	require.Len(t, snap.Nodes, 2)
	assert.Equal(t, uint32(11), snap.Nodes[peer(1)].NodeOperatorID)
	assert.Equal(t, [32]byte{0xaa}, snap.Nodes[peer(1)].Signer)

	// Which of these nodes is "me" is not answered here: a snapshot carries the registry and
	// nothing about who is reading it. Whoever keeps it in sync supplies that.
	assert.Equal(t, uint32(2), snap.Nodes[peer(1)].WorkflowDONID)
}

func TestReader_ReadSendsPagination(t *testing.T) {
	caller := fullCaller()
	snapshot(t, caller)

	assert.Equal(t, int64(0), caller.lastStart.Int64())
	assert.Equal(t, int64(pageLimit), caller.lastLimit.Int64())
}

func TestReader_SkipsDeprecatedAndUnparseableCapabilities(t *testing.T) {
	caller := fullCaller()
	caller.caps = append(caller.caps,
		// Deprecated capabilities are still returned by the contract; treating them as live would
		// let workflows resolve a retired capability.
		capregv2.CapabilitiesRegistryCapabilityInfo{
			CapabilityId: "old@1.0.0", IsDeprecated: true, Metadata: []byte(`{"capabilityType":1}`),
		},
		// One bad metadata blob must not fail the whole read.
		capregv2.CapabilitiesRegistryCapabilityInfo{CapabilityId: "garbage@1.0.0", Metadata: []byte("{")},
		capregv2.CapabilitiesRegistryCapabilityInfo{CapabilityId: "empty@1.0.0", Metadata: nil},
		capregv2.CapabilitiesRegistryCapabilityInfo{CapabilityId: "weird@1.0.0", Metadata: []byte(`{"capabilityType":99}`)},
	)

	snap := snapshot(t, caller)

	assert.Len(t, snap.Capabilities, 2)
	for _, skipped := range []string{"old@1.0.0", "garbage@1.0.0", "empty@1.0.0", "weird@1.0.0"} {
		_, ok := snap.Capabilities[skipped]
		assert.False(t, ok, "%s should have been skipped", skipped)
	}
}

// A failed read reports which call failed, so an operator knows whether the contract, the address
// or the RPC is at fault.
func TestReader_ReadPropagatesEachCallError(t *testing.T) {
	for _, tc := range []struct {
		name    string
		fail    func(*fakeCaller)
		wantErr string
	}{
		{"capabilities", func(c *fakeCaller) { c.capsErr = errors.New("boom") }, "getCapabilities"},
		{"dons", func(c *fakeCaller) { c.donsErr = errors.New("boom") }, "getDONs"},
		{"nodes", func(c *fakeCaller) { c.nodesErr = errors.New("boom") }, "getNodes"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			caller := fullCaller()
			tc.fail(caller)

			_, err := newTestReader(t, caller).Read(context.Background())
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.wantErr)
		})
	}
}

func TestParseCapabilityType(t *testing.T) {
	for _, tc := range []struct {
		name    string
		in      []byte
		want    capabilities.CapabilityType
		wantErr bool
	}{
		{"trigger", []byte(`{"capabilityType":0}`), capabilities.CapabilityTypeTrigger, false},
		{"action", []byte(`{"capabilityType":1}`), capabilities.CapabilityTypeAction, false},
		{"consensus", []byte(`{"capabilityType":2}`), capabilities.CapabilityTypeConsensus, false},
		{"target", []byte(`{"capabilityType":3}`), capabilities.CapabilityTypeTarget, false},
		{"out of range", []byte(`{"capabilityType":9}`), capabilities.CapabilityTypeUnknown, true},
		{"empty", nil, capabilities.CapabilityTypeUnknown, true},
		{"malformed", []byte("nope"), capabilities.CapabilityTypeUnknown, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseCapabilityType(tc.in)
			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
