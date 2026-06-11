package legacyevm

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-evm/pkg/assets"
	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/configtest"
	evmtoml "github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	gasmocks "github.com/smartcontractkit/chainlink-evm/pkg/gas/mocks"
	"github.com/smartcontractkit/chainlink-evm/pkg/txmgr"
	txmmocks "github.com/smartcontractkit/chainlink-evm/pkg/txmgr/mocks"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

func TestChain_TransactNoBalanceCheck(t *testing.T) {
	txManager := txmmocks.NewMockEvmTxManager(t)
	cfg := configtest.NewChainScopedConfig(t, func(c *evmtoml.EVMConfig) {
		v := uint64(42)
		c.GasEstimator.LimitTransfer = &v
	})

	c := &chain{
		id:  big.NewInt(0),
		cfg: cfg,
		txm: txManager,
	}

	ctx := t.Context()
	from := common.HexToAddress("0x123")
	to := common.HexToAddress("0x456")
	amount := big.NewInt(1)
	txManager.EXPECT().SendNativeToken(ctx, c.id, from, to, *amount, cfg.EVM().GasEstimator().LimitTransfer()).Return(txmgr.Tx{}, nil).Once()

	require.NoError(t, c.Transact(ctx, from.String(), to.String(), amount, false))
}

func TestChain_TransactWithBalanceCheck(t *testing.T) {
	from := common.HexToAddress("0x123")
	to := common.HexToAddress("0x456")
	amount := big.NewInt(1)
	fromPriceMax := assets.GWei(41)

	txManager := txmmocks.NewMockEvmTxManager(t)
	client := clienttest.NewClient(t)
	gasEstimator := gasmocks.NewEvmFeeEstimator(t)
	cfg := configtest.NewChainScopedConfig(t, func(c *evmtoml.EVMConfig) {
		v := uint64(42)
		c.GasEstimator.LimitTransfer = &v
		c.GasEstimator.PriceMax = fromPriceMax.Add(assets.GWei(1))
		c.KeySpecific = evmtoml.KeySpecificConfig{
			{Key: new(types.EIP55AddressFromAddress(from)),
				GasEstimator: evmtoml.KeySpecificGasEstimator{
					PriceMax: fromPriceMax,
				},
			},
		}
	})

	c := &chain{
		id:           big.NewInt(0),
		cfg:          cfg,
		txm:          txManager,
		client:       client,
		gasEstimator: gasEstimator,
	}

	ctx := t.Context()
	client.EXPECT().BalanceAt(ctx, from, (*big.Int)(nil)).Return(big.NewInt(1000), nil).Once()
	gasEstimator.EXPECT().GetMaxCost(ctx, (assets.Eth)(*amount), []byte(nil), cfg.EVM().GasEstimator().LimitTransfer(), fromPriceMax, &from, &to).Return(big.NewInt(1000), nil).Once()
	txManager.EXPECT().SendNativeToken(ctx, c.id, from, to, *amount, cfg.EVM().GasEstimator().LimitTransfer()).Return(txmgr.Tx{}, nil).Once()

	require.NoError(t, c.Transact(ctx, from.String(), to.String(), amount, true))
}
