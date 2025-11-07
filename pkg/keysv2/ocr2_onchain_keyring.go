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

func CreateOCR2OnchainKeyring(ctx context.Context, ks keystore.Keystore, keyringName string) (ocrtypes.OnchainKeyring, error) {
	onchainKeyPath := keystore.NewKeyPath(EVMPrefix, OCR2OnchainPrefix, keyringName)
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
	publicKey, err := gethcrypto.UnmarshalPubkey(resp.Keys[0].KeyInfo.PublicKey)
	if err != nil {
		return nil, err
	}
	addr := gethcrypto.PubkeyToAddress(*publicKey)
	return &evmOnchainKeyring{ks: ks, onchainKey: resp.Keys[0].KeyInfo, addr: addr}, nil
}

func ListOCR2OnchainKeyrings(ctx context.Context, ks keystore.Keystore, localNames ...string) ([]ocrtypes.OnchainKeyring, error) {
	var names []string
	if len(localNames) > 0 {
		for _, ln := range localNames {
			names = append(names, keystore.NewKeyPath(EVMPrefix, OCR2OnchainPrefix, ln).String())
		}
	}

	getReq := keystore.GetKeysRequest{KeyNames: names}
	resp, err := ks.GetKeys(ctx, getReq)
	if err != nil {
		return nil, err
	}

	keyrings := make([]ocrtypes.OnchainKeyring, 0, len(resp.Keys))
	for _, key := range resp.Keys {
		if !strings.HasPrefix(key.KeyInfo.Name, keystore.NewKeyPath(EVMPrefix, OCR2OnchainPrefix).String()) {
			continue
		}
		keyPath := keystore.NewKeyPathFromString(key.KeyInfo.Name)
		publicKey, err := gethcrypto.UnmarshalPubkey(key.KeyInfo.PublicKey)
		if err != nil {
			return nil, err
		}
		addr := gethcrypto.PubkeyToAddress(*publicKey)
		keyrings = append(keyrings, &evmOnchainKeyring{ks: ks, onchainKey: key.KeyInfo, addr: addr, keyPath: keyPath})
	}
	return keyrings, nil
}

var _ ocrtypes.OnchainKeyring = &evmOnchainKeyring{}

type evmOnchainKeyring struct {
	ks         keystore.Keystore
	onchainKey keystore.KeyInfo
	addr       common.Address
	keyPath    keystore.KeyPath
}

func (k *evmOnchainKeyring) KeyPath() keystore.KeyPath {
	return k.keyPath
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
