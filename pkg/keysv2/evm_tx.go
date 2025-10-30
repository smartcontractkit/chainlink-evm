package keysv2

import (
	"context"
	"fmt"
	"strings"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-common/keystore"
)

const (
	EVMPrefix        = "evm"
	TxKeystorePrefix = "tx"
)

type KeyPath []string

func (k KeyPath) String() string {
	return joinKeySegments(k...)
}

func (k KeyPath) Leaf() string {
	return k[len(k)-1]
}

func NewKeyPath(segments ...string) KeyPath {
	return segments
}

func NewKeyPathFromString(fullName string) KeyPath {
	return strings.Split(fullName, "/")
}

// JoinKeySegments joins path-like key name segments using "/" and avoids double slashes.
// Empty segments are skipped so JoinKeySegments("EVM", "TX", "my-key") => "EVM/TX/my-key".
func joinKeySegments(segments ...string) string {
	cleaned := make([]string, 0, len(segments))
	for _, s := range segments {
		s = strings.Trim(s, "/")
		if s == "" {
			continue
		}
		cleaned = append(cleaned, s)
	}
	return strings.Join(cleaned, "/")
}

type TxKey struct {
	ks      keystore.Keystore
	keyPath KeyPath
	addr    common.Address
}

type SignTxRequest struct {
	ChainID *big.Int
	Tx      *gethtypes.Transaction
}

type SignTxResponse struct {
	Tx *gethtypes.Transaction
}

func (k *TxKey) KeyPath() KeyPath {
	return k.keyPath
}

func (k *TxKey) Address() common.Address {
	return k.addr
}

func (k *TxKey) SignTx(ctx context.Context, req SignTxRequest) (SignTxResponse, error) {
	if req.ChainID == nil {
		return SignTxResponse{}, fmt.Errorf("chainID is nil")
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

func (k *TxKey) GetTransactOpts(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	if chainID == nil {
		return nil, fmt.Errorf("chainID is nil")
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

func CreateTxKey(ks keystore.Keystore, name string) (*TxKey, error) {
	path := NewKeyPath(EVMPrefix, TxKeystorePrefix, name)
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
		return nil, fmt.Errorf("no keys created")
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

func GetTxKeys(ctx context.Context, ks keystore.Keystore, names []string) ([]*TxKey, error) {
	var fullNames []string
	for _, name := range names {
		fullNames = append(fullNames, NewKeyPath(EVMPrefix, TxKeystorePrefix, name).String())
	}
	resp, err := ks.GetKeys(ctx, keystore.GetKeysRequest{KeyNames: fullNames})
	if err != nil {
		return nil, err
	}

	// Note we rely on deterministic order of keys in the response
	var keys []*TxKey
	for _, key := range resp.Keys {
		publicKey, err := gethcrypto.UnmarshalPubkey(key.KeyInfo.PublicKey)
		if err != nil {
			return nil, err
		}
		addr := gethcrypto.PubkeyToAddress(*publicKey)
		keys = append(keys, &TxKey{
			ks:      ks,
			keyPath: NewKeyPathFromString(key.KeyInfo.Name),
			addr:    addr,
		})
	}
	return keys, nil
}
