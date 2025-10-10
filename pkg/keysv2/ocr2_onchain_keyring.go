package keysv2

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-common/keystore"
	evmutil "github.com/smartcontractkit/libocr/offchainreporting2plus/chains/evmutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

const (
	OCR2OnchainPrefix = "ocr2_onchain"
)

func GetOCR2OnchainKeystoreName(localName string) string {
	return JoinKeySegments(EVMPrefix, OCR2OnchainPrefix, localName)
}

func IsOCR2OnchainKey(name string) bool {
	return strings.HasPrefix(name, JoinKeySegments(EVMPrefix, OCR2OnchainPrefix, ""))
}

type OCR2OnchainKeyringCreateRequest struct {
	LocalName string
}

type OCR2OnchainKeyringCreateResponse struct {
	Keyring ocrtypes.OnchainKeyring
}

type OCR2OnchainKeyringGetKeyringsRequest struct {
	Names []string // Empty slice means get all OCR2 onchain keyrings
}

type OCR2OnchainKeyringGetKeyringsResponse struct {
	Keyrings []ocrtypes.OnchainKeyring
}

// CreateOCR2OnchainKeyring creates an OCR2 onchain keyring using the base keystore and returns the handle.
func CreateOCR2OnchainKeyring(ctx context.Context, ks keystore.Keystore, localName string) (ocrtypes.OnchainKeyring, error) {
	createReq := keystore.CreateKeysRequest{
		Keys: []keystore.CreateKeyRequest{
			{
				KeyName: GetOCR2OnchainKeystoreName(localName),
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
	publicKey, err := gethcrypto.UnmarshalPubkey(resp.Keys[0].KeyInfo.PublicKey)
	if err != nil {
		return nil, err
	}
	addr := gethcrypto.PubkeyToAddress(*publicKey)
	return &evmOnchainKeyring{ks: ks, onchainKey: resp.Keys[0].KeyInfo, addr: addr}, nil
}

// ListOCR2OnchainKeyrings lists OCR2 onchain keyrings. If no local names provided, returns all OCR2 onchain keyrings.
func ListOCR2OnchainKeyrings(ctx context.Context, ks keystore.Keystore, localNames ...string) ([]ocrtypes.OnchainKeyring, error) {
	// Build names if explicitly provided
	var names []string
	if len(localNames) > 0 {
		for _, ln := range localNames {
			names = append(names, GetOCR2OnchainKeystoreName(ln))
		}
	}

	getReq := keystore.GetKeysRequest{KeyNames: names}
	resp, err := ks.GetKeys(ctx, getReq)
	if err != nil {
		return nil, err
	}

	var keyrings []ocrtypes.OnchainKeyring
	for _, key := range resp.Keys {
		if IsOCR2OnchainKey(key.KeyInfo.Name) {
			keyrings = append(keyrings, &evmOnchainKeyring{ks: ks, onchainKey: key.KeyInfo})
		}
	}
	return keyrings, nil
}

var _ ocrtypes.OnchainKeyring = &evmOnchainKeyring{}

type evmOnchainKeyring struct {
	ks         keystore.Keystore
	onchainKey keystore.KeyInfo
	addr       common.Address
}

func (k *evmOnchainKeyring) PublicKey() ocrtypes.OnchainPublicKey {
	return k.addr.Bytes()
}

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
		KeyName: k.onchainKey.Name,
		Data:    ReportToSigData(reportCtx, report),
	})
	return signResp.Signature, err
}

func (k *evmOnchainKeyring) Verify(publicKey ocrtypes.OnchainPublicKey, reportCtx ocrtypes.ReportContext, report ocrtypes.Report, signature []byte) bool {
	verifyResp, err := k.ks.Verify(context.Background(), keystore.VerifyRequest{
		KeyType:   keystore.ECDSA_S256,
		PublicKey: k.onchainKey.PublicKey,
		Data:      ReportToSigData(reportCtx, report),
		Signature: signature,
	})
	if err != nil {
		// Log?
		return false
	}
	return verifyResp.Valid
}

func (k *evmOnchainKeyring) MaxSignatureLength() int {
	return 65
}
