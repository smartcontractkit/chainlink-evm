package configtest

import (
	"testing"

	"github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils/big"
)

func NewChainScopedConfig(t testing.TB, overrideFn func(c *toml.EVMConfig)) *config.ChainScoped {
	chainID := big.NewI(0)
	evmCfg := &toml.EVMConfig{
		ChainID: chainID,
		Chain:   toml.Defaults(chainID),
	}

	if overrideFn != nil {
		// We need to get the chainID from the override function first to load the correct chain defaults.
		// Then we apply the override values on top
		overrideFn(evmCfg)
		evmCfg.Chain = toml.Defaults(evmCfg.ChainID)
		overrideFn(evmCfg)
	}

	return config.NewTOMLChainScopedConfig(evmCfg)
}
