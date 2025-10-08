package keystore

import (
	"context"
	"fmt"
	"strings"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-common/keystore"
)

const (
	EVMPrefix        = "evm"
	TxKeystorePrefix = "tx"
)

// JoinKeySegments joins path-like key name segments using "/" and avoids double slashes.
// Empty segments are skipped so JoinKeySegments("EVM", "TX", "my-key") => "EVM/TX/my-key".
func JoinKeySegments(segments ...string) string {
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

func GetTxKeystoreName(localName string) string {
	return JoinKeySegments(EVMPrefix, TxKeystorePrefix, localName)
}

type TxKey struct {
	ks keystore.Keystore
	// Fully qualified name in keystore. Use for administration.
	fullName string
	name     string
	addr     common.Address
}

type SignTxRequest struct {
	ChainID *big.Int
	Tx      *gethtypes.Transaction
}

type SignTxResponse struct {
	Tx *gethtypes.Transaction
}

func (k *TxKey) Name() string {
	return k.name
}

func (k *TxKey) FullName() string {
	return k.fullName
}

func (k *TxKey) Address() common.Address {
	return k.addr
}

func (k *TxKey) SignTx(ctx context.Context, req SignTxRequest) (SignTxResponse, error) {
	signer := gethtypes.LatestSignerForChainID(req.ChainID)
	h := signer.Hash(req.Tx)
	signReq := keystore.SignRequest{
		KeyName: k.FullName(),
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

func CreateTxKey(ks keystore.Keystore, localName string) (*TxKey, error) {
	createReq := keystore.CreateKeysRequest{
		Keys: []keystore.CreateKeyRequest{
			{
				KeyName: GetTxKeystoreName(localName),
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
		ks:       ks,
		name:     localName,
		fullName: GetTxKeystoreName(localName),
		addr:     addr,
	}, nil
}

func GetTxKeys(ctx context.Context, ks keystore.Keystore, names []string) ([]*TxKey, error) {
	var fullNames []string
	for _, name := range names {
		fullNames = append(fullNames, GetTxKeystoreName(name))
	}
	resp, err := ks.GetKeys(ctx, keystore.GetKeysRequest{KeyNames: fullNames})
	if err != nil {
		return nil, err
	}

	var keys []*TxKey
	for _, key := range resp.Keys {
		publicKey, err := gethcrypto.UnmarshalPubkey(key.KeyInfo.PublicKey)
		if err != nil {
			return nil, err
		}
		addr := gethcrypto.PubkeyToAddress(*publicKey)
		keys = append(keys, &TxKey{
			ks:       ks,
			fullName: key.KeyInfo.Name,
			name:     key.KeyInfo.Name,
			addr:     addr,
		})
	}
	return keys, nil
}
