package keystore

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	commonks "github.com/smartcontractkit/chainlink-common/keystore"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

var _ keys.Store = &Keystore{}

type Keystore struct {
	ks commonks.Keystore
}

// Backwards compatibility with the usage of the keys package.
func NewKeystore(ks commonks.Keystore) *Keystore {
	return &Keystore{ks: ks}
}

func (k *Keystore) CheckEnabled(ctx context.Context, address common.Address) error {
	return nil
}

func (k *Keystore) EnabledAddresses(ctx context.Context) ([]common.Address, error) {
	return nil, nil
}

func (k *Keystore) SignMessage(ctx context.Context, address common.Address, message []byte) ([]byte, error) {
	return nil, nil
}

func (k *Keystore) Sign(ctx context.Context, address common.Address, bytes []byte) ([]byte, error) {
	return nil, nil
}

func (k *Keystore) GetNextAddress(ctx context.Context, addresses ...common.Address) (common.Address, error) {
	return common.Address{}, nil
}

func (k *Keystore) SignTx(ctx context.Context, fromAddress common.Address, tx *types.Transaction) (*types.Transaction, error) {
	return nil, nil
}

func (k *Keystore) GetMutex(address common.Address) *keys.Mutex {
	return &keys.Mutex{}
}
