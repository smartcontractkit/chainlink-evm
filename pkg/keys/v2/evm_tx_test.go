package keys_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	commonks "github.com/smartcontractkit/chainlink-common/keystore"
	evmks "github.com/smartcontractkit/chainlink-evm/pkg/keys/v2"
	"github.com/stretchr/testify/require"
)

func setupBackend(t *testing.T, testKey common.Address) (*simulated.Backend, func() error) {
	backend := simulated.NewBackend(types.GenesisAlloc{
		testKey: {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)), // 10 ETH
		},
	}, simulated.WithBlockGasLimit(10e6))
	return backend, func() error {
		return backend.Close()
	}
}

func TestTxKey(t *testing.T) {
	storage := commonks.NewMemoryStorage()
	ctx := t.Context()
	ks, err := commonks.LoadKeystore(ctx, storage, "test-password", commonks.WithScryptParams(commonks.FastScryptParams))
	require.NoError(t, err)
	testKey, err := evmks.CreateTxKey(ks, "test-tx-key")
	require.NoError(t, err)
	testKey2, err := evmks.CreateTxKey(ks, "test-tx-key-2")
	require.NoError(t, err)

	backend, cleanup := setupBackend(t, testKey.Address())
	defer cleanup()

	testTransaction := types.NewTransaction(
		0,                       // Nonce
		testKey2.Address(),      // To other key
		big.NewInt(1),           // Value
		21000,                   // Gas Limit
		big.NewInt(20000000000), // Gas Price
		nil)
	resp, err := testKey.SignTx(ctx, evmks.SignTxRequest{
		ChainID: big.NewInt(1337), // Use a test chain ID
		Tx:      testTransaction,
	})
	require.NoError(t, err)
	require.NotNil(t, resp.Tx)
	require.NoError(t, backend.Client().SendTransaction(ctx, resp.Tx))
	backend.Commit()
	receipt, err := backend.Client().TransactionReceipt(ctx, resp.Tx.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

	endBalance, err := backend.Client().BalanceAt(ctx, testKey2.Address(), nil)
	require.NoError(t, err)
	require.Equal(t, endBalance, big.NewInt(1))

	// Admin operation will invalidate the keys.
	_, err = ks.DeleteKeys(ctx, commonks.DeleteKeysRequest{
		KeyNames: []string{testKey.KeyPath().String(), testKey2.KeyPath().String()},
	})
	require.NoError(t, err)

	// Empty names will return all keys.
	keys, err := evmks.GetTxKeys(ctx, ks, []string{})
	require.NoError(t, err)
	require.Empty(t, keys)

	// Create a non-EVM key and verify filtering
	_, err = ks.CreateKeys(ctx, commonks.CreateKeysRequest{
		Keys: []commonks.CreateKeyRequest{
			{KeyName: "solana/tx/non-evm-key", KeyType: commonks.Ed25519},
		},
	})
	require.NoError(t, err)
	keys, err = evmks.GetTxKeys(ctx, ks, []string{})
	require.NoError(t, err)
	require.Empty(t, keys) // Should filter out non-EVM keys

	// Signing will now error.
	_, err = testKey.SignTx(ctx, evmks.SignTxRequest{
		ChainID: big.NewInt(1337), // Use a test chain ID
		Tx:      testTransaction,
	})
	require.Error(t, err)
	require.ErrorIs(t, err, commonks.ErrKeyNotFound)
}
