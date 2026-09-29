package config

import (
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

type balanceMonitorConfig struct {
	c toml.BalanceMonitor
}

func (b *balanceMonitorConfig) Enabled() bool {
	return *b.c.Enabled
}

func (b *balanceMonitorConfig) ERC20TokenAddress() *types.EIP55Address {
	return b.c.ERC20TokenAddress
}
