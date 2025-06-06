package config

import (
	"time"

	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
	"github.com/smartcontractkit/chainlink-evm/pkg/types"
)

type workflowConfig struct {
	c toml.Workflow
}

func (b *workflowConfig) FromAddress() *types.EIP55Address {
	return b.c.FromAddress
}

func (b *workflowConfig) ForwarderAddress() *types.EIP55Address {
	return b.c.ForwarderAddress
}

func (b *workflowConfig) GasLimitDefault() *uint64 {
	return b.c.GasLimitDefault
}

func (b *workflowConfig) TxAcceptanceState() *commontypes.TransactionStatus {
	return b.c.TxAcceptanceState
}

func (b *workflowConfig) PollPeriod() time.Duration {
	return b.c.PollPeriod.Duration()
}
func (b *workflowConfig) AcceptanceTimeout() time.Duration {
	return b.c.AcceptanceTimeout.Duration()
}
