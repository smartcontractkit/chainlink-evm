package keysv2

import (
	"testing"

	commonks "github.com/smartcontractkit/chainlink-common/keystore"
	"github.com/stretchr/testify/require"
)

func TestPeerKeyring(t *testing.T) {
	storage := commonks.NewMemoryStorage()
	ctx := t.Context()
	ks, err := commonks.LoadKeystore(ctx, storage, commonks.EncryptionParams{
		Password:     "test-password",
		ScryptParams: commonks.FastScryptParams,
	})
	require.NoError(t, err)
	peerKeyring, err := CreatePeerKeyring(ctx, ks, "test-peer-keyring")
	require.NoError(t, err)
	msg := []byte("test-message")
	signature, err := peerKeyring.Sign(msg)
	require.NoError(t, err)
	require.NotNil(t, signature)
}
