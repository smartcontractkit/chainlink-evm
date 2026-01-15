// Package keys/v2 provides key management functionality for EVM transactions and OCR2 keyrings.
package keys

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-common/keystore"
)

const (
	// PrefixEVM is the prefix for EVM-related keys.
	PrefixEVM = "evm"
	// PrefixTxKeystore is the prefix for transaction keys.
	PrefixTxKeystore = "tx"
)

// TxKey represents an EVM transaction signing key.
type TxKey struct {
	ks interface {
		keystore.Reader
		keystore.Signer
	}
	keyPath keystore.KeyPath
	addr    common.Address
}

// SignTxRequest contains the request to sign a transaction.
type SignTxRequest struct {
	ChainID *big.Int
	Tx      *gethtypes.Transaction
}

// SignTxResponse contains the signed transaction.
type SignTxResponse struct {
	Tx *gethtypes.Transaction
}

// SignRawDataRequest contains the request to sign raw data.
type SignRawDataRequest struct {
	Data []byte
}

// SignRawDataResponse contains the signed raw data.
type SignRawDataResponse struct {
	Signature []byte
}

// KeyPath returns the key path for this transaction key.
func (k *TxKey) KeyPath() keystore.KeyPath {
	return k.keyPath
}

// Address returns the Ethereum address for this transaction key.
func (k *TxKey) Address() common.Address {
	return k.addr
}

// SignTx signs a transaction using this key.
func (k *TxKey) SignTx(ctx context.Context, req SignTxRequest) (SignTxResponse, error) {
	if req.ChainID == nil {
		return SignTxResponse{}, errors.New("chainID is nil")
	}
	signer := gethtypes.LatestSignerForChainID(req.ChainID)
	h := signer.Hash(req.Tx)
	signReq := keystore.SignRequest{
		KeyName: k.keyPath.String(),
		Data:    h[:],
	}
	signResp, err := k.ks.Sign(ctx, signReq)
	if err != nil {
		return SignTxResponse{}, err
	}
	req.Tx, err = req.Tx.WithSignature(signer, signResp.Signature)
	if err != nil {
		return SignTxResponse{}, err
	}
	return SignTxResponse{Tx: req.Tx}, nil
}

func (k *TxKey) SignRaw(ctx context.Context, req SignRawDataRequest) (SignRawDataResponse, error) {
	signReq := keystore.SignRequest{
		KeyName: k.keyPath.String(),
		Data:    req.Data,
	}
	signResp, err := k.ks.Sign(ctx, signReq)
	if err != nil {
		return SignRawDataResponse{}, err
	}
	return SignRawDataResponse{Signature: signResp.Signature}, nil
}

// GetTransactOpts returns transaction options for this key.
func (k *TxKey) GetTransactOpts(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	if chainID == nil {
		return nil, errors.New("chainID is nil")
	}
	return &bind.TransactOpts{
		From: k.addr,
		Signer: func(address common.Address, tx *gethtypes.Transaction) (*gethtypes.Transaction, error) {
			if k.Address() != address {
				return nil, bind.ErrNotAuthorized
			}
			resp, err := k.SignTx(ctx, SignTxRequest{
				ChainID: chainID,
				Tx:      tx,
			})
			if err != nil {
				return nil, err
			}
			return resp.Tx, nil
		},
	}, nil
}

// CreateTxKey creates a new transaction signing key.
// Note that key names are prefixed with PrefixEVM and PrefixTxKeystore.
// For example, a key named "test-key" will be stored at the path "evm/tx/test-key".
func CreateTxKey(ks keystore.Keystore, name string) (*TxKey, error) {
	path := keystore.NewKeyPath(PrefixEVM, PrefixTxKeystore, name)
	createReq := keystore.CreateKeysRequest{
		Keys: []keystore.CreateKeyRequest{
			{
				KeyName: path.String(),
				KeyType: keystore.ECDSA_S256,
			},
		},
	}
	resp, err := ks.CreateKeys(context.Background(), createReq)
	if err != nil {
		return nil, err
	}
	if len(resp.Keys) == 0 {
		return nil, errors.New("no keys created")
	}
	publicKey, err := gethcrypto.UnmarshalPubkey(resp.Keys[0].KeyInfo.PublicKey)
	if err != nil {
		return nil, err
	}
	addr := gethcrypto.PubkeyToAddress(*publicKey)
	return &TxKey{
		ks:      ks,
		keyPath: path,
		addr:    addr,
	}, nil
}

// GetTxKeysOption configures GetTxKeys behavior.
type GetTxKeysOption func(*getTxKeysOptions)

type getTxKeysOptions struct {
	noPrefix bool
}

// WithNoPrefix disables adding the evm/tx prefix to key names.
// When set, names are used as-is (useful for keystores with externally managed names
// like KMS-backed keystores).
func WithNoPrefix() GetTxKeysOption {
	return func(opts *getTxKeysOptions) {
		opts.noPrefix = true
	}
}

// GetTxKeys retrieves transaction keys by name.
// By default, prepends the evm/tx prefix to names.
// For example, a key named "test-key" will be retrieved at the path "evm/tx/test-key".
// Use WithNoPrefix() to use names as-is (for KMS-backed keystores).
func GetTxKeys(ctx context.Context, ks interface {
	keystore.Reader
	keystore.Signer
}, names []string, opts ...GetTxKeysOption) ([]*TxKey, error) {
	options := &getTxKeysOptions{}
	for _, opt := range opts {
		opt(options)
	}

	fullNames := make([]string, 0, len(names))
	if options.noPrefix {
		fullNames = names
	} else {
		for _, name := range names {
			fullNames = append(fullNames, keystore.NewKeyPath(PrefixEVM, PrefixTxKeystore, name).String())
		}
	}

	resp, err := ks.GetKeys(ctx, keystore.GetKeysRequest{KeyNames: fullNames})
	if err != nil {
		return nil, err
	}

	// Always require all requested keys to be found
	if len(names) > 0 && len(resp.Keys) != len(names) {
		return nil, errors.New("some keys not found")
	}

	// Note we rely on deterministic order of keys in the response
	keys := make([]*TxKey, 0, len(resp.Keys))
	prefixPath := keystore.NewKeyPath(PrefixEVM, PrefixTxKeystore)
	for _, key := range resp.Keys {
		path := keystore.NewKeyPathFromString(key.KeyInfo.Name)

		// If no prefix, sanity check key type (for KMS-backed keystores)
		if options.noPrefix {
			if key.KeyInfo.KeyType != keystore.ECDSA_S256 {
				return nil, errors.New("tried to load a non-ECDSA_S256 key: " + key.KeyInfo.Name)
			}
		}

		// If no names are provided and we're using prefix, filter only evm keys
		if !options.noPrefix && len(names) == 0 && !path.HasPrefix(prefixPath) {
			continue
		}

		publicKey, err := gethcrypto.UnmarshalPubkey(key.KeyInfo.PublicKey)
		if err != nil {
			return nil, err
		}
		addr := gethcrypto.PubkeyToAddress(*publicKey)
		keys = append(keys, &TxKey{
			ks:      ks,
			keyPath: path,
			addr:    addr,
		})
	}
	return keys, nil
}
