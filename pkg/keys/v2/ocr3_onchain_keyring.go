package keys

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/keystore"
	"github.com/smartcontractkit/chainlink-common/keystore/ocr3util"
)

// OCR3OnchainKeyringWithPublicKey pairs an OCR3 onchain keyring with the public
// key (an EVM address) that it signs from.
type OCR3OnchainKeyringWithPublicKey[RI any] struct {
	Keyring   ocr3types.OnchainKeyring2[RI]
	PublicKey common.Address
}

// CreateOCR3OnchainKeyringIgnoringRI creates a new OCR3 onchain keyring.
// Note that key names are prefixed with PrefixEVM and PrefixOCR2Onchain (to preserve compatibility with OCR2).
// For example, a key named "test-key" will be stored at the path "evm/ocr2_onchain/test-key".
// Be careful! The generic RI parameter is not used anywhere in the implementation. It acts as a defensive interface boundary to avoid signing a wrong type of reports.
// If you have security-critical information in RI, you should make your own ocr3types.OnchainKeyring2 that does use the information in RI for signing.
func CreateOCR3OnchainKeyringIgnoringRI[RI any](ctx context.Context, ks keystore.Keystore, keyringName string) (OCR3OnchainKeyringWithPublicKey[RI], error) {
	if keyringName == "" {
		return OCR3OnchainKeyringWithPublicKey[RI]{}, errors.New("keyring name cannot be empty")
	}
	onchainKeyPath := keystore.NewKeyPath(PrefixEVM, PrefixOCR2Onchain, keyringName)
	createReq := keystore.CreateKeysRequest{
		Keys: []keystore.CreateKeyRequest{
			{
				KeyName: onchainKeyPath.String(),
				KeyType: keystore.ECDSA_S256,
			},
		},
	}
	resp, err := ks.CreateKeys(ctx, createReq)
	if err != nil {
		return OCR3OnchainKeyringWithPublicKey[RI]{}, err
	}
	if len(resp.Keys) != 1 {
		return OCR3OnchainKeyringWithPublicKey[RI]{}, fmt.Errorf("expected 1 key, got %d", len(resp.Keys))
	}
	publicKey, err := crypto.UnmarshalPubkey(resp.Keys[0].KeyInfo.PublicKey)
	if err != nil {
		return OCR3OnchainKeyringWithPublicKey[RI]{}, err
	}
	addr := crypto.PubkeyToAddress(*publicKey)
	return OCR3OnchainKeyringWithPublicKey[RI]{
		Keyring:   ocr3util.AsOCR3OnchainKeyring2IgnoringRI[RI](&evmOnchainKeyring2{ks: ks, addr: addr, keyPath: onchainKeyPath}),
		PublicKey: addr,
	}, nil
}

// Be careful! The generic RI parameter is not used anywhere in the implementation. It acts as a defensive interface boundary to avoid signing a wrong type of reports.
// If you have security-critical information in RI, you should make your own ocr3types.OnchainKeyring2 that does use the information in RI for signing.
func ListOCR3OnchainKeyrings[RI any](ctx context.Context, ks keystore.Keystore, keyringNames ...string) ([]OCR3OnchainKeyringWithPublicKey[RI], error) {
	var names []string
	if len(keyringNames) > 0 {
		for _, krn := range keyringNames {
			names = append(names, keystore.NewKeyPath(PrefixEVM, PrefixOCR2Onchain, krn).String())
		}
	}

	getReq := keystore.GetKeysRequest{KeyNames: names}
	resp, err := ks.GetKeys(ctx, getReq)
	if err != nil {
		return nil, err
	}

	keyrings := make([]OCR3OnchainKeyringWithPublicKey[RI], 0, len(resp.Keys))
	for _, key := range resp.Keys {
		if !strings.HasPrefix(key.KeyInfo.Name, keystore.NewKeyPath(PrefixEVM, PrefixOCR2Onchain).String()) {
			continue
		}
		keyPath := keystore.NewKeyPathFromString(key.KeyInfo.Name)
		publicKey, err := crypto.UnmarshalPubkey(key.KeyInfo.PublicKey)
		if err != nil {
			return nil, err
		}
		addr := crypto.PubkeyToAddress(*publicKey)
		keyrings = append(keyrings, OCR3OnchainKeyringWithPublicKey[RI]{
			Keyring:   ocr3util.AsOCR3OnchainKeyring2IgnoringRI[RI](&evmOnchainKeyring2{ks: ks, addr: addr, keyPath: keyPath}),
			PublicKey: addr,
		})
	}
	return keyrings, nil
}

var _ ocr3util.OnchainKeyring2 = &evmOnchainKeyring2{}

type evmOnchainKeyring2 struct {
	ks      keystore.Keystore
	addr    common.Address
	keyPath keystore.KeyPath
}

func (e *evmOnchainKeyring2) Sign(configDigest ocrtypes.ConfigDigest, seqNr uint64, report types.Report) (signature []byte, err error) {
	hash := hashReport(configDigest, seqNr, report)
	signResp, err := e.ks.Sign(context.Background(), keystore.SignRequest{
		KeyName: e.keyPath.String(),
		Data:    hash,
	})
	if err != nil {
		return nil, err
	}
	return compressSignature(signResp.Signature), nil
}

func (e *evmOnchainKeyring2) Verify(publicKey ocrtypes.OnchainPublicKey, configDigest ocrtypes.ConfigDigest, seqNr uint64, report types.Report, signature []byte) bool {
	hash := hashReport(configDigest, seqNr, report)
	recoveredPublicKey, err := crypto.SigToPub(hash, uncompressSignature(signature))
	if err != nil {
		return false
	}
	signerAddress := crypto.PubkeyToAddress(*recoveredPublicKey)
	return bytes.Equal(publicKey, signerAddress[:])
}

func (e *evmOnchainKeyring2) Has(publicKey ocrtypes.OnchainPublicKey) bool {
	return bytes.Equal(publicKey, e.addr.Bytes())
}

func (e *evmOnchainKeyring2) MaxSignatureLength() int {
	return 64
}

func (e *evmOnchainKeyring2) DebugIdentifier() string {
	return fmt.Sprintf("addr = %s, path = %s", e.addr.Hex(), e.keyPath.String())
}

// Compresses the given ECDSA signature in its internal format (r, s, v) into the
// corresponding compressed format (r, s'), where s' in {s, n-s} depending on v.
func compressSignature(sig []byte) []byte {
	// v is either 0 or 1; stores the recovery id of the underlying public key.
	v := sig[64]
	if v != 0 && v != 1 {
		panic("v not in {0, 1}, cannot happen with cryptographic certainty")
	}

	result := make([]byte, 64)
	copy(result, sig[:64])

	// In the smart contracts, a value of v=0 is assumed. (See https://github.com/smartcontractkit/offchain-reporting/blob/03dfcfe3e36f89a0d50678860b4b35ddf45a1c7e/contract3/src/OCR3ECDSAAttestationVerifierLib.sol#L62 .)
	// So, if the recovery id is 1, s needs to flipped to "simulate" v=0.
	if v == 1 {
		// Extract the value of s from the signature bytes.
		s := new(big.Int).SetBytes(result[32:64])

		// Get the order of the elliptic curve n, and compute s' = n - s.
		N := crypto.S256().Params().N
		s.Sub(N, s)

		// Store s back into the result bytes.
		s.FillBytes(result[32:64])
	}

	return result
}

// Uncompresses the given ECDSA signature in its external format (r, s') into the
// format expected internally (r, s', v=0).
// Note that uncompression does not guarantee that the resulting signature verifies
// using the "github.com/ethereum/go-ethereum/crypto".VerifySignature function. Instead, it guarantees that
// the correct public key is recovered using "github.com/ethereum/go-ethereum/crypto".SigToPub.
func uncompressSignature(sig []byte) []byte {
	out := make([]byte, 65)
	copy(out, sig)
	return out
}

// hashReport computes an EVM-friendly cryptographic hash function over configDigest, seqNr, and reportWithInfo.Report.
// The main difference from OCR2 hashing (https://github.com/smartcontractkit/chainlink-evm/blob/b4068bf735e6e1af28602eece9e29a0bd31eed37/pkg/keys/v2/ocr2_onchain_keyring.go#L98)
// is that we hash the sequence number `seqNr` instead of `epoch||round||extraHash`.
// Domain separation is guaranteed because the result of this function is a Keccak hash computed over 96 bytes
// while the aforementioned OCR2 function is a Keccak hash computed over 128 bytes.
func hashReport(configDigest ocrtypes.ConfigDigest, seqNr uint64, report types.Report) []byte {
	data := make([]byte, 96)
	copy(data, configDigest[:])
	binary.BigEndian.PutUint64(data[56:64], seqNr)
	copy(data[64:96], crypto.Keccak256(report))
	return crypto.Keccak256(data)
}
