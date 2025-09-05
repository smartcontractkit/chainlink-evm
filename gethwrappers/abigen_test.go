package gethwrappers

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/log_emitter"
)

// Test that the generated Deploy method fill all the required fields and returns the correct address.
// We perform this test using the generated LogEmitter wrapper.
func TestGeneratedDeployMethodAddressField(t *testing.T) {
	owner := MustNewSimTransactor(t)
	ec := simulated.NewBackend(types.GenesisAlloc{
		owner.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)),
		},
	}, simulated.WithBlockGasLimit(10e6)).Client()

	emitterAddr, _, emitter, err := log_emitter.DeployLogEmitter(owner, ec)
	require.NoError(t, err)
	require.Equal(t, emitterAddr, emitter.Address())
}

func MustNewSimTransactor(t testing.TB) *bind.TransactOpts {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	require.NoError(t, err)
	return transactor
}
