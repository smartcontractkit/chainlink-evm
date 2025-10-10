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
	OCR2OffchainSigningPrefix    = "ocr2_offchain_signing"
	OCR2OffchainEncryptionPrefix = "ocr2_offchain_encryption"
)

func GetOCR2OffchainSigningKeystoreName(localName string) string {
	return JoinKeySegments(OCR2OffchainSigningPrefix, localName)
}

func GetOCR2OffchainEncryptionKeystoreName(localName string) string {
	return JoinKeySegments(OCR2OffchainEncryptionPrefix, localName)
}

func IsOCR2OffchainSigningKey(name string) bool {
	return strings.HasPrefix(name, JoinKeySegments(EVMPrefix, OCR2OffchainSigningPrefix, ""))
}

func IsOCR2OffchainEncryptionKey(name string) bool {
	return strings.HasPrefix(name, JoinKeySegments(EVMPrefix, OCR2OffchainEncryptionPrefix, ""))
}

type OCR2OffchainKeyringCreateRequest struct {
	LocalName string
}

type OCR2OffchainKeyringCreateResponse struct {
	Keyring ocrtypes.OffchainKeyring
}

type OCR2OffchainKeyringGetKeyringsRequest struct {
	Names []string // Empty slice means get all OCR2 offchain keyrings
}

type OCR2OffchainKeyringGetKeyringsResponse struct {
	Keyrings []ocrtypes.OffchainKeyring
}

// TODO: Maybe want to embed this interface and add a few other methods.
func CreateOCR2OffchainKeyring(ctx context.Context, ks keystore.Keystore, localName string) (ocrtypes.OffchainKeyring, error) {
	createReq := keystore.CreateKeysRequest{
		Keys: []keystore.CreateKeyRequest{
			{
				KeyName: GetOCR2OffchainSigningKeystoreName(localName),
				KeyType: keystore.Ed25519,
			},
			{
				KeyName: GetOCR2OffchainEncryptionKeystoreName(localName),
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
		OffchainKey:           resp.Keys[0].KeyInfo,
		OffchainEncryptionKey: resp.Keys[1].KeyInfo,
	}, nil
}

// ListOCR2OffchainKeyrings lists OCR2 offchain keyrings. If no local names provided, returns all OCR2 offchain keyrings.
func ListOCR2OffchainKeyrings(ctx context.Context, ks keystore.Keystore, localNames ...string) ([]ocrtypes.OffchainKeyring, error) {
	// Build names if explicitly provided
	var names []string
	if len(localNames) > 0 {
		for _, ln := range localNames {
			names = append(names, GetOCR2OffchainSigningKeystoreName(ln))
		}
	}

	getReq := keystore.GetKeysRequest{KeyNames: names}
	resp, err := ks.GetKeys(ctx, getReq)
	if err != nil {
		return nil, err
	}

	var keyrings []ocrtypes.OffchainKeyring
	for _, key := range resp.Keys {
		if IsOCR2OffchainSigningKey(key.KeyInfo.Name) {
			// Fetch the matching encryption key
			encryptionKeyName := GetOCR2OffchainEncryptionKeystoreName(strings.TrimPrefix(key.KeyInfo.Name, JoinKeySegments(EVMPrefix, OCR2OffchainSigningPrefix, "")))
			getReq := keystore.GetKeysRequest{KeyNames: []string{encryptionKeyName}}
			getResp, err := ks.GetKeys(context.Background(), getReq)
			if err != nil {
				return nil, err
			}
			if len(getResp.Keys) == 0 {
				return nil, fmt.Errorf("encryption key not found for keyring: %s", key.KeyInfo.Name)
			}
			keyrings = append(keyrings, &evmOffchainKeyring{
				ks:                    ks,
				OffchainKey:           key.KeyInfo,
				OffchainEncryptionKey: getResp.Keys[0].KeyInfo,
			})
		}
	}
	return keyrings, nil
}

var _ ocrtypes.OffchainKeyring = &evmOffchainKeyring{}

type evmOffchainKeyring struct {
	ks                    keystore.Keystore
	OffchainKey           keystore.KeyInfo
	OffchainEncryptionKey keystore.KeyInfo
}

func (k *evmOffchainKeyring) OffchainPublicKey() ocrtypes.OffchainPublicKey {
	var pubKey ocrtypes.OffchainPublicKey
	copy(pubKey[:], k.OffchainKey.PublicKey)
	return pubKey
}

func (k *evmOffchainKeyring) ConfigEncryptionPublicKey() ocrtypes.ConfigEncryptionPublicKey {
	var pubKey ocrtypes.ConfigEncryptionPublicKey
	copy(pubKey[:], k.OffchainEncryptionKey.PublicKey)
	return pubKey
}

func (k *evmOffchainKeyring) OffchainSign(msg []byte) ([]byte, error) {
	signResp, err := k.ks.Sign(context.Background(), keystore.SignRequest{
		KeyName: k.OffchainKey.Name,
		Data:    msg,
	})
	return signResp.Signature, err
}

func (k *evmOffchainKeyring) ConfigDiffieHellman(point [curve25519.PointSize]byte) ([curve25519.PointSize]byte, error) {
	resp, err := k.ks.DeriveSharedSecret(context.Background(), keystore.DeriveSharedSecretRequest{
		KeyName:      k.OffchainEncryptionKey.Name,
		RemotePubKey: point[:],
	})
	if err != nil {
		return [curve25519.PointSize]byte{}, err
	}

	var sharedPoint [curve25519.PointSize]byte
	copy(sharedPoint[:], resp.SharedSecret)
	return sharedPoint, nil
}
