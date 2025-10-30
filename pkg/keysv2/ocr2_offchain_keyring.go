package keysv2

import (
	"context"
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-common/keystore"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"golang.org/x/crypto/curve25519"
)

const (
	OCR2OffchainSigning    = "ocr2_offchain_signing"
	OCR2OffchainEncryption = "ocr2_offchain_encryption"
	OCR2OffchainPrefix     = "ocr2_offchain"
)

func CreateOCR2OffchainKeyring(ctx context.Context, ks keystore.Keystore, keyringName string) (ocrtypes.OffchainKeyring, error) {
	signingKeyPath := NewKeyPath(EVMPrefix, OCR2OffchainPrefix, keyringName, OCR2OffchainSigning)
	encryptionKeyPath := NewKeyPath(EVMPrefix, OCR2OffchainPrefix, keyringName, OCR2OffchainEncryption)
	createReq := keystore.CreateKeysRequest{
		Keys: []keystore.CreateKeyRequest{
			{
				KeyName: signingKeyPath.String(),
				KeyType: keystore.Ed25519,
			},
			{
				KeyName: encryptionKeyPath.String(),
				KeyType: keystore.X25519,
			},
		},
	}
	resp, err := ks.CreateKeys(ctx, createReq)
	if err != nil {
		return nil, err
	}
	if len(resp.Keys) != 2 {
		return nil, fmt.Errorf("expected 2 keys, got %d", len(resp.Keys))
	}
	return &evmOffchainKeyring{
		ks:                    ks,
		signingKeyPath:        signingKeyPath,
		encryptionKeyPath:     encryptionKeyPath,
		offchainKey:           resp.Keys[0].KeyInfo,
		offchainEncryptionKey: resp.Keys[1].KeyInfo,
	}, nil
}

// ListOCR2OffchainKeyrings lists OCR2 offchain keyrings. If no local names provided, returns all OCR2 offchain keyrings.
func GetOCR2OffchainKeyrings(ctx context.Context, ks keystore.Keystore, keyRingNames []string) ([]ocrtypes.OffchainKeyring, error) {
	var names []string
	if len(keyRingNames) > 0 {
		for _, keyRingName := range keyRingNames {
			names = append(names, NewKeyPath(EVMPrefix, OCR2OffchainPrefix, keyRingName, OCR2OffchainSigning).String())
			names = append(names, NewKeyPath(EVMPrefix, OCR2OffchainPrefix, keyRingName, OCR2OffchainEncryption).String())
		}
	}

	getReq := keystore.GetKeysRequest{KeyNames: names}
	resp, err := ks.GetKeys(ctx, getReq)
	if err != nil {
		return nil, err
	}

	// Group by keyrings.
	keyRingMap := make(map[string][]keystore.KeyInfo)
	for _, key := range resp.Keys {
		if !strings.HasPrefix(key.KeyInfo.Name, NewKeyPath(EVMPrefix, OCR2OffchainPrefix).String()) {
			continue
		}
		keyPath := NewKeyPathFromString(key.KeyInfo.Name)
		// Example:
		// /evm/ocr2_offchain/keyring_name/ocr2_offchain_signing
		// /evm/ocr2_offchain/keyring_name/ocr2_offchain_encryption
		keyRingMap[keyPath[:2].String()] = append(keyRingMap[keyPath[:2].String()], key.KeyInfo)
	}

	var keyrings []ocrtypes.OffchainKeyring
	for _, keyInfos := range keyRingMap {
		keyrings = append(keyrings, &evmOffchainKeyring{
			ks:                ks,
			signingKeyPath:    NewKeyPathFromString(keyInfos[0].Name),
			encryptionKeyPath: NewKeyPathFromString(keyInfos[1].Name),
		})
	}
	return keyrings, nil
}

var _ ocrtypes.OffchainKeyring = &evmOffchainKeyring{}

type evmOffchainKeyring struct {
	ks                    keystore.Keystore
	signingKeyPath        KeyPath
	encryptionKeyPath     KeyPath
	offchainKey           keystore.KeyInfo
	offchainEncryptionKey keystore.KeyInfo
}

func (k *evmOffchainKeyring) ConfigEncryptionKeyPath() KeyPath {
	return k.encryptionKeyPath
}

func (k *evmOffchainKeyring) ConfigSigningKeyPath() KeyPath {
	return k.signingKeyPath
}

func (k *evmOffchainKeyring) OffchainPublicKey() ocrtypes.OffchainPublicKey {
	var pubKey ocrtypes.OffchainPublicKey
	copy(pubKey[:], k.offchainKey.PublicKey)
	return pubKey
}

func (k *evmOffchainKeyring) ConfigEncryptionPublicKey() ocrtypes.ConfigEncryptionPublicKey {
	var pubKey ocrtypes.ConfigEncryptionPublicKey
	copy(pubKey[:], k.offchainEncryptionKey.PublicKey)
	return pubKey
}

func (k *evmOffchainKeyring) OffchainSign(msg []byte) ([]byte, error) {
	signResp, err := k.ks.Sign(context.Background(), keystore.SignRequest{
		KeyName: k.offchainKey.Name,
		Data:    msg,
	})
	return signResp.Signature, err
}

func (k *evmOffchainKeyring) ConfigDiffieHellman(point [curve25519.PointSize]byte) ([curve25519.PointSize]byte, error) {
	resp, err := k.ks.DeriveSharedSecret(context.Background(), keystore.DeriveSharedSecretRequest{
		KeyName:      k.offchainEncryptionKey.Name,
		RemotePubKey: point[:],
	})
	if err != nil {
		return [curve25519.PointSize]byte{}, err
	}

	var sharedPoint [curve25519.PointSize]byte
	copy(sharedPoint[:], resp.SharedSecret)
	return sharedPoint, nil
}
