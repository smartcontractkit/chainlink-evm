package config

import (
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
)

type balanceMonitorConfig struct {
	c toml.BalanceMonitor
}

func (b *balanceMonitorConfig) Enabled() bool {
	return *b.c.Enabled
}
