// Package keys provides key management functionality for EVM transactions and OCR2 keyrings.
package keys

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/chains/evmutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/keystore"
)

const (
	// PrefixOCR2Onchain is the prefix for OCR2 onchain keyrings.
	PrefixOCR2Onchain = "ocr2_onchain"
)

// CreateOCR2OnchainKeyring creates a new OCR2 onchain keyring.
// Note that key names are prefixed with PrefixEVM and PrefixOCR2Onchain.
// For example, a key named "test-key" will be stored at the path "evm/ocr2_onchain/test-key".
func CreateOCR2OnchainKeyring(ctx context.Context, ks keystore.Keystore, keyringName string) (ocrtypes.OnchainKeyring, error) {
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
		return nil, err
	}
	if len(resp.Keys) != 1 {
		return nil, fmt.Errorf("expected 1 key, got %d", len(resp.Keys))
	}
	publicKey, err := crypto.UnmarshalPubkey(resp.Keys[0].KeyInfo.PublicKey)
	if err != nil {
		return nil, err
	}
	addr := crypto.PubkeyToAddress(*publicKey)
	return &evmOnchainKeyring{ks: ks, addr: addr, keyPath: onchainKeyPath}, nil
}

// ListOCR2OnchainKeyrings lists OCR2 onchain keyrings, optionally filtered by keyringnames.
func ListOCR2OnchainKeyrings(ctx context.Context, ks keystore.Keystore, keyringNames ...string) ([]ocrtypes.OnchainKeyring, error) {
	fullNames, err := ListCompatibleKeys(ctx, ks, []string{PrefixOCR2Onchain}, keyringNames...)
	if err != nil {
		return nil, err
	}

	getReq := keystore.GetKeysRequest{KeyNames: fullNames}
	resp, err := ks.GetKeys(ctx, getReq)
	if err != nil {
		return nil, err
	}

	keyrings := make([]ocrtypes.OnchainKeyring, 0, len(resp.Keys))
	for _, key := range resp.Keys {
		keyPath := keystore.NewKeyPathFromString(key.KeyInfo.Name)
		publicKey, err := crypto.UnmarshalPubkey(key.KeyInfo.PublicKey)
		if err != nil {
			return nil, err
		}
		addr := crypto.PubkeyToAddress(*publicKey)
		keyrings = append(keyrings, &evmOnchainKeyring{ks: ks, addr: addr, keyPath: keyPath})
	}
	return keyrings, nil
}

func ListOCR3PlusOnchainKeyrings(ctx context.Context, ks keystore.Keystore, keyringNames ...string) ([]ocrtypes.OnchainKeyringPlus, error) {
	fullNames, err := ListCompatibleKeys(ctx, ks, []string{PrefixOCR2Onchain, PrefixOCR3Onchain}, keyringNames...)
	if err != nil {
		return nil, err
	}

	getReq := keystore.GetKeysRequest{KeyNames: fullNames}
	resp, err := ks.GetKeys(ctx, getReq)
	if err != nil {
		return nil, err
	}

	keyrings := make([]ocrtypes.OnchainKeyringPlus, 0, len(resp.Keys))
	for _, key := range resp.Keys {
		keyrings = append(keyrings, &evmOnchainKeyring{something with the key})
	}
	return keyrings, nil
}

// Gets the list of fully qualified keyring names for given prioritizedPrefixes.
// If there are multiple keys with the same keyring name (e.g, "evm/ocr2/mykey" and "evm/ocr3/mykey") and different
// prefixes, then the one with the higher order prefix in prioritizedPrefixes is returned.
func ListCompatibleKeys(ctx context.Context, ks keystore.Keystore, prioritizedPrefixes []string, keyringNames ...string) ([]string, error) {
	resp, err := ks.GetKeys(ctx, keystore.GetKeysRequest{})
	if err != nil {
		return nil, err
	}

	var m map[string]struct{}
	for _, key := range resp.Keys {
		m[key.KeyInfo.Name] = struct{}{}
	}

	result := make([]string, 0, len(keyringNames))
	for i, keyringName := range keyringNames {
		for _, prefix := range prioritizedPrefixes {
			fullKeyringName := keystore.NewKeyPath(PrefixEVM, prefix, keyringName).String()
			if _, ok := m[fullKeyringName]; ok {
				result[i] = fullKeyringName
			}
		}
		if result[i] == "" {
			return nil, fmt.Errorf("no compatible keyrings found for %s", keyringName)
		}
	}
	return result, nil
}

var _ ocrtypes.OnchainKeyring = &evmOnchainKeyring{}

type evmOnchainKeyring struct {
	ks      keystore.Keystore
	addr    common.Address
	keyPath keystore.KeyPath
}

func (k *evmOnchainKeyring) KeyPath() keystore.KeyPath {
	return k.keyPath
}

func (k *evmOnchainKeyring) PublicKey() ocrtypes.OnchainPublicKey {
	return k.addr.Bytes()
}

// ReportToSigData converts a report and report context into signature data.
func ReportToSigData(reportCtx ocrtypes.ReportContext, report ocrtypes.Report) []byte {
	rawReportContext := evmutil.RawReportContext(reportCtx)
	sigData := crypto.Keccak256(report)
	sigData = append(sigData, rawReportContext[0][:]...)
	sigData = append(sigData, rawReportContext[1][:]...)
	sigData = append(sigData, rawReportContext[2][:]...)
	return crypto.Keccak256(sigData)
}

func (k *evmOnchainKeyring) Sign(reportCtx ocrtypes.ReportContext, report ocrtypes.Report) ([]byte, error) {
	signResp, err := k.ks.Sign(context.Background(), keystore.SignRequest{
		KeyName: k.keyPath.String(),
		Data:    ReportToSigData(reportCtx, report),
	})
	return signResp.Signature, err
}

func (k *evmOnchainKeyring) Verify(publicKey ocrtypes.OnchainPublicKey, reportCtx ocrtypes.ReportContext, report ocrtypes.Report, signature []byte) bool {
	sigData := ReportToSigData(reportCtx, report)
	authorPubkey, err := crypto.SigToPub(sigData, signature)
	if err != nil {
		return false
	}
	pubKey := crypto.S256().Marshal(authorPubkey.X, authorPubkey.Y)
	verifyResp, err := k.ks.Verify(context.Background(), keystore.VerifyRequest{
		KeyType:   keystore.ECDSA_S256,
		PublicKey: pubKey,
		Data:      sigData,
		Signature: signature,
	})
	if err != nil {
		return false
	}
	return verifyResp.Valid
}

func (k *evmOnchainKeyring) MaxSignatureLength() int {
	return 65
}
