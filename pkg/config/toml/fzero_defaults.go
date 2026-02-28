package toml

import (
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// FZeroDefaults holds the recommended configuration defaults for F=0 single-node DON deployments.
// With F=0, there is no redundancy — a stuck nonce blocks all writes with no backup transmitter.
type FZeroDefaults struct {
	AutoPurgeEnabled     bool
	AutoPurgeThreshold   uint32
	AutoPurgeMinAttempts uint32
}

// DefaultFZeroConfig returns the recommended defaults for F=0 deployments.
func DefaultFZeroConfig() FZeroDefaults {
	return FZeroDefaults{
		AutoPurgeEnabled:     true,
		AutoPurgeThreshold:   5,
		AutoPurgeMinAttempts: 3,
	}
}

// ApplyFZeroDefaults applies F=0-specific defaults to the EVM config.
// Only sets values that haven't been explicitly configured by the operator.
// Returns true if any defaults were applied.
func ApplyFZeroDefaults(lggr logger.Logger, txCfg *Transactions) bool {
	if txCfg == nil {
		return false
	}

	defaults := DefaultFZeroConfig()
	applied := false

	// Auto-Purge: default ON for F=0
	if txCfg.AutoPurge.Enabled == nil {
		txCfg.AutoPurge.Enabled = &defaults.AutoPurgeEnabled
		applied = true
		lggr.Infow("F=0 default applied: AutoPurge.Enabled = true",
			"reason", "With F=0, a stuck nonce blocks all writes with no backup transmitter")
	} else if !*txCfg.AutoPurge.Enabled {
		lggr.Warnw("Auto-Purge is disabled for F=0 DON. Stuck nonces will block all writes with no fallback. Consider enabling EVM.Transactions.AutoPurge.Enabled")
	}

	// Set AutoPurge threshold and min attempts if not already set and AutoPurge is enabled
	if txCfg.AutoPurge.Enabled != nil && *txCfg.AutoPurge.Enabled {
		if txCfg.AutoPurge.Threshold == nil {
			txCfg.AutoPurge.Threshold = &defaults.AutoPurgeThreshold
			applied = true
			lggr.Infow("F=0 default applied: AutoPurge.Threshold", "value", defaults.AutoPurgeThreshold)
		}
		if txCfg.AutoPurge.MinAttempts == nil {
			txCfg.AutoPurge.MinAttempts = &defaults.AutoPurgeMinAttempts
			applied = true
			lggr.Infow("F=0 default applied: AutoPurge.MinAttempts", "value", defaults.AutoPurgeMinAttempts)
		}
	}

	return applied
}
