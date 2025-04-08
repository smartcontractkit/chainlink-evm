package keystest

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"maps"
	"slices"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
)

var _ core.Keystore = &MemoryChainStore{}

func NewMemoryChainStore() *MemoryChainStore {
	return &MemoryChainStore{privKeys: make(map[string]*ecdsa.PrivateKey)}
}

type MemoryChainStore struct {
	mu       sync.RWMutex
	privKeys map[string]*ecdsa.PrivateKey
}

func (m *MemoryChainStore) MustCreate(t require.TestingT) common.Address {
	addr, err := m.Create()
	require.NoError(t, err)
	return addr
}
func (m *MemoryChainStore) Create() (common.Address, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	privKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return common.Address{}, err
	}
	addr := crypto.PubkeyToAddress(privKey.PublicKey)
	if m.privKeys == nil {
		m.privKeys = make(map[string]*ecdsa.PrivateKey)
	}
	m.privKeys[addr.String()] = privKey
	return addr, nil
}

func (m *MemoryChainStore) Delete(addr common.Address) {
	m.mu.Lock()
	if m.privKeys != nil {
		delete(m.privKeys, addr.String())
	}
	m.mu.Unlock()
}

func (m *MemoryChainStore) Accounts(ctx context.Context) (accounts []string, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return slices.Collect(maps.Keys(m.privKeys)), nil
}

func (m *MemoryChainStore) Sign(ctx context.Context, account string, data []byte) (signed []byte, err error) {
	m.mu.Lock()
	pk, ok := m.privKeys[account]
	m.mu.Unlock()
	if !ok {
		return nil, fmt.Errorf("account %s not found", account)
	}
	return crypto.Sign(data, pk)
}
