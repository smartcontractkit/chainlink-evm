package keys

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"math/big"
	"slices"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys/internal"
)

type Store interface {
	AddressChecker
	AddressLister
	RoundRobin
	MessageSigner
	Locker
	RawUnhashedSigner
}

// ChainStore extends Store with methods that require a chain ID.
type ChainStore interface {
	Store
	TxSigner
}

type Addresses interface {
	AddressChecker
	AddressLister
}

type AddressChecker interface {
	CheckEnabled(ctx context.Context, address common.Address) error
}

type AddressLister interface {
	EnabledAddresses(ctx context.Context) (addresses []common.Address, err error)
}

// RawUnhashedSigner is an interface for signing raw bytes without hashing them first.
// This is useful for signing transactions or other data that needs to be signed as-is (required for TRON).
// The interface effectively exposes the Sign method from the keystore
type RawUnhashedSigner interface {
	SignRawUnhashedBytes(ctx context.Context, address common.Address, bytes []byte) ([]byte, error)
}

type MessageSigner interface {
	SignMessage(ctx context.Context, address common.Address, message []byte) ([]byte, error)
}

type TxSigner interface {
	SignTx(ctx context.Context, fromAddress common.Address, tx *types.Transaction) (*types.Transaction, error)
}

type Locker interface {
	GetMutex(address common.Address) *Mutex
}

type RoundRobin interface {
	GetNextAddress(ctx context.Context, addresses ...common.Address) (address common.Address, err error)
}

var _ Store = &store{}

type store struct {
	ks core.Keystore

	internal.Locker[Mutex]

	lastUsedMu sync.Mutex
	lastUsed   map[common.Address]time.Time
}

// NewStore returns a new Store backed by ks.
func NewStore(ks core.Keystore) Store {
	return &store{
		ks:       ks,
		lastUsed: make(map[common.Address]time.Time),
	}
}

func (s *store) CheckEnabled(ctx context.Context, address common.Address) error {
	as, err := s.ks.Accounts(ctx)
	if err != nil {
		return fmt.Errorf("failed to get accounts: %w", err)
	}
	if !slices.Contains(as, address.String()) {
		return errors.New("not enabled")
	}
	return nil
}

func (s *store) EnabledAddresses(ctx context.Context) ([]common.Address, error) {
	as, err := s.ks.Accounts(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts: %w", err)
	}
	addresses := make([]common.Address, 0, len(as))
	for _, a := range as {
		if !common.IsHexAddress(a) {
			return nil, fmt.Errorf("invalid address: %s", a)
		}
		addresses = append(addresses, common.HexToAddress(a))
	}
	return addresses, nil
}

func (s *store) SignMessage(ctx context.Context, address common.Address, message []byte) ([]byte, error) {
	return s.ks.Sign(ctx, address.String(), accounts.TextHash(message))
}

func (s *store) SignRawUnhashedBytes(ctx context.Context, address common.Address, bytes []byte) ([]byte, error) {
	return s.ks.Sign(ctx, address.String(), bytes)
}

func (s *store) GetNextAddress(ctx context.Context, whitelist ...common.Address) (next common.Address, err error) {
	s.lastUsedMu.Lock()
	defer s.lastUsedMu.Unlock()

	if len(whitelist) == 0 {
		whitelist, err = s.EnabledAddresses(ctx)
		if err != nil {
			return
		}
	} else {
		var enabled []common.Address
		enabled, err = s.EnabledAddresses(ctx)
		if err != nil {
			return
		}
		enabledSet := maps.Collect(func(yield func(common.Address, struct{}) bool) {
			for _, addr := range enabled {
				if !yield(addr, struct{}{}) {
					return
				}
			}
		})
		whitelist = slices.DeleteFunc(whitelist, func(a common.Address) bool {
			_, ok := enabledSet[a]
			return !ok
		})
	}

	var lru time.Time

	for _, addr := range whitelist {
		lastUsed, ok := s.lastUsed[addr]
		if !ok {
			// never
			next = addr
			break
		}
		if lru.IsZero() || lastUsed.Before(lru) {
			lru = lastUsed
			next = addr
		}
	}

	s.lastUsed[next] = time.Now()

	return
}

type chainStore struct {
	*store
	chainID *big.Int
}

// NewChainStore returns a new ChainStore for chainID backed by ks.
func NewChainStore(ks core.Keystore, chainID *big.Int) ChainStore {
	return &chainStore{
		store: &store{
			ks:       ks,
			lastUsed: make(map[common.Address]time.Time),
		},
		chainID: chainID,
	}
}

func (s *chainStore) SignTx(ctx context.Context, fromAddress common.Address, tx *types.Transaction) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(s.chainID)
	h := signer.Hash(tx)
	sig, err := s.ks.Sign(ctx, fromAddress.String(), h[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}
	return tx.WithSignature(signer, sig)
}
