package keys

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-common/keystore"
	"github.com/stretchr/testify/require"
)

func TestCoreKeystore(t *testing.T) {
	ctx := context.Background()
	tmpfile, err := os.CreateTemp("", "keystore.json")
	require.NoError(t, err)
	defer os.Remove(tmpfile.Name())
	ks, err := keystore.LoadKeystore(ctx, keystore.NewFileStorage(tmpfile.Name()), "password")
	require.NoError(t, err)

	coreKs := NewTxKeyCoreKeystore(ks)
	accounts, err := coreKs.Accounts(ctx)
	require.NoError(t, err)
	require.Len(t, accounts, 0)

	txKey, err := CreateTxKey(ks, "key1")
	require.NoError(t, err)

	keys, err := ks.GetKeys(ctx, keystore.GetKeysRequest{
		KeyNames: []string{txKey.KeyPath().String()},
	})
	require.NoError(t, err)
	require.Len(t, keys.Keys, 1)

	data := crypto.Keccak256([]byte("data"))
	signature, err := coreKs.Sign(ctx, txKey.Address().String(), data)
	require.NoError(t, err)
	require.NotNil(t, signature)

	resp, err := ks.Verify(ctx, keystore.VerifyRequest{
		KeyType:   keystore.ECDSA_S256,
		PublicKey: keys.Keys[0].KeyInfo.PublicKey,
		Data:      data,
		Signature: signature,
	})
	require.NoError(t, err)
	require.True(t, resp.Valid)

	// Make sure the cache populated.
	require.Equal(t, txKey.KeyPath().String(), coreKs.cache[txKey.Address().String()])

	// Sign again with cached.
	signature, err = coreKs.Sign(ctx, txKey.Address().String(), data)
	require.NoError(t, err)
	require.NotNil(t, signature)
	resp, err = ks.Verify(ctx, keystore.VerifyRequest{
		KeyType:   keystore.ECDSA_S256,
		PublicKey: keys.Keys[0].KeyInfo.PublicKey,
		Data:      data,
		Signature: signature,
	})
	require.NoError(t, err)
	require.True(t, resp.Valid)

	// Test WithAllowedKeyNames
	_, err = CreateTxKey(ks, "key2")
	require.NoError(t, err)
	filteredCoreKs := NewTxKeyCoreKeystore(ks, WithAllowedKeyNames([]string{txKey.KeyPath().String()}))
	accounts, err = filteredCoreKs.Accounts(ctx)
	require.NoError(t, err)
	require.Len(t, accounts, 1)
	require.Equal(t, txKey.Address().String(), accounts[0])
}
