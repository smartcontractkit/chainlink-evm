// Package keysv2 provides key management functionality for EVM transactions and OCR2 keyrings.
package keysv2

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
	// EVMPrefix is the prefix for EVM-related keys.
	EVMPrefix        = "evm"
	// TxKeystorePrefix is the prefix for transaction keys.
	TxKeystorePrefix = "tx"
)

// TxKey represents an EVM transaction signing key.
type TxKey struct {
	ks      keystore.Keystore
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
func CreateTxKey(ks keystore.Keystore, name string) (*TxKey, error) {
	path := keystore.NewKeyPath(EVMPrefix, TxKeystorePrefix, name)
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

// GetTxKeys retrieves transaction keys by name.
func GetTxKeys(ctx context.Context, ks keystore.Keystore, names []string) ([]*TxKey, error) {
	fullNames := make([]string, 0, len(names))
	for _, name := range names {
		fullNames = append(fullNames, keystore.NewKeyPath(EVMPrefix, TxKeystorePrefix, name).String())
	}
	resp, err := ks.GetKeys(ctx, keystore.GetKeysRequest{KeyNames: fullNames})
	if err != nil {
		return nil, err
	}

	// Note we rely on deterministic order of keys in the response
	keys := make([]*TxKey, 0, len(resp.Keys))
	for _, key := range resp.Keys {
		publicKey, err := gethcrypto.UnmarshalPubkey(key.KeyInfo.PublicKey)
		if err != nil {
			return nil, err
		}
		addr := gethcrypto.PubkeyToAddress(*publicKey)
		keys = append(keys, &TxKey{
			ks:      ks,
			keyPath: keystore.NewKeyPathFromString(key.KeyInfo.Name),
			addr:    addr,
		})
	}
	return keys, nil
}
