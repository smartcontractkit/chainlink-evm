package config

import (
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

type writeCapabilityConfig struct {
	c toml.WriteCapability
}

func (b *writeCapabilityConfig) FromAddress() *types.EIP55Address {
	return b.c.FromAddress
}

func (b *writeCapabilityConfig) ForwarderAddress() *types.EIP55Address {
	return b.c.ForwarderAddress
}

func (b *writeCapabilityConfig) GasLimitDefault() *uint64 {
	return b.c.GasLimitDefault
}
