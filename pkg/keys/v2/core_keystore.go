package keys

import (
	"context"
	"errors"

	"github.com/smartcontractkit/chainlink-common/keystore"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
)

var _ core.Keystore = &TxKeyCoreKeystore{}

type TxKeyCoreKeystore struct {
	ks interface {
		keystore.Reader
		keystore.Signer
	}
	cache map[string]string
}

// NewTxKeyCoreKeystore creates a new CoreKeystore for transaction keys.
// This wrapper is required for using TxKeys with the txm
// which requires address based lookups.
func NewTxKeyCoreKeystore(ks interface {
	keystore.Reader
	keystore.Signer
}) *TxKeyCoreKeystore {
	return &TxKeyCoreKeystore{
		ks:    ks,
		cache: make(map[string]string),
	}
}

func (s *TxKeyCoreKeystore) Accounts(ctx context.Context) ([]string, error) {
	keys, err := GetTxKeys(ctx, s.ks, []string{})
	if err != nil {
		return nil, err
	}
	accounts := make([]string, 0, len(keys))
	for _, key := range keys {
		accounts = append(accounts, key.Address().String())
	}
	return accounts, nil
}

func (s *TxKeyCoreKeystore) Sign(ctx context.Context, account string, data []byte) ([]byte, error) {
	if addr, ok := s.cache[account]; ok {
		resp, err := s.ks.Sign(ctx, keystore.SignRequest{
			KeyName: addr,
			Data:    data,
		})
		if err != nil {
			return nil, err
		}
		return resp.Signature, nil
	}
	// Otherwise do the first time lookup to find the key by address.
	keys, err := GetTxKeys(ctx, s.ks, []string{})
	if err != nil {
		return nil, err
	}
	if len(keys) == 0 {
		return nil, errors.New("no keys found")
	}
	for _, key := range keys {
		if key.Address().String() == account {
			s.cache[account] = key.KeyPath().String()
			resp, err := key.SignRaw(ctx, SignRawDataRequest{Data: data})
			if err != nil {
				return nil, err
			}
			return resp.Signature, nil
		}
	}
	return nil, errors.New("key not found")
}

func (s *TxKeyCoreKeystore) Decrypt(ctx context.Context, account string, data []byte) ([]byte, error) {
	return nil, errors.New("not implemented")
}
