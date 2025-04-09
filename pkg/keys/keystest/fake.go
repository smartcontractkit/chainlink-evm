package keystest

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys/internal"
)

var _ keys.ChainStore = &FakeChainStore{}

// FakeChainStore is an implementation of keys.ChainStore for testing.
// The zero value is usable.
type FakeChainStore struct {
	Addresses

	TxSigner
	MessageSigner

	internal.Locker[keys.Mutex]
}

var _ keys.Addresses = Addresses([]common.Address{})

type Addresses []common.Address

func (a Addresses) CheckEnabled(ctx context.Context, address common.Address) error {
	if slices.Contains(a, address) {
		return nil
	}
	return errors.New("address not found")
}

func (a Addresses) EnabledAddresses(ctx context.Context) (addresses []common.Address, err error) {
	return a, nil
}

func (a Addresses) GetNextAddress(ctx context.Context, addresses ...common.Address) (address common.Address, err error) {
	addresses = slices.DeleteFunc(addresses, func(address common.Address) bool {
		return !slices.Contains(a, address)
	})
	if len(addresses) == 0 {
		return common.Address{}, errors.New("no addresses")
	}
	return addresses[rand.Int()%len(addresses)], nil
}

var _ keys.TxSigner = TxSigner(nil)

type TxSigner func(ctx context.Context, from common.Address, tx *types.Transaction) (*types.Transaction, error)

func (f TxSigner) SignTx(ctx context.Context, from common.Address, tx *types.Transaction) (*types.Transaction, error) {
	if f == nil {
		return tx, nil
	}
	return f(ctx, from, tx)
}

var _ keys.MessageSigner = MessageSigner(nil)

type MessageSigner func(ctx context.Context, address common.Address, message []byte) ([]byte, error)

func (f MessageSigner) SignMessage(ctx context.Context, address common.Address, message []byte) ([]byte, error) {
	if f == nil {
		return message, nil
	}
	return f(ctx, address, message)
}

type ECDSAMessageSigner ecdsa.PrivateKey

func (k *ECDSAMessageSigner) SignMessage(ctx context.Context, address common.Address, message []byte) ([]byte, error) {
	if pub := crypto.PubkeyToAddress(k.PublicKey); address != pub {
		return nil, fmt.Errorf("unable to sign for %s with key: %s", address, pub)
	}
	return crypto.Sign(accounts.TextHash(message), (*ecdsa.PrivateKey)(k))
}

func (f MessageSigner) SignRawUnhashedBytes(ctx context.Context, address common.Address, message []byte) ([]byte, error) {
	if f == nil {
		return message, nil
	}

	return f(ctx, address, message)
}
