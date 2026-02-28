package toml

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestDefaultFZeroConfig(t *testing.T) {
	defaults := DefaultFZeroConfig()
	assert.True(t, defaults.AutoPurgeEnabled)
	assert.Equal(t, uint32(5), defaults.AutoPurgeThreshold)
	assert.Equal(t, uint32(3), defaults.AutoPurgeMinAttempts)
}

func TestApplyFZeroDefaults(t *testing.T) {
	lggr := logger.Test(t)

	t.Run("nil config returns false", func(t *testing.T) {
		applied := ApplyFZeroDefaults(lggr, nil)
		assert.False(t, applied)
	})

	t.Run("applies all defaults when nothing is set", func(t *testing.T) {
		txCfg := &Transactions{}

		applied := ApplyFZeroDefaults(lggr, txCfg)

		assert.True(t, applied)
		require.NotNil(t, txCfg.AutoPurge.Enabled)
		assert.True(t, *txCfg.AutoPurge.Enabled)
		require.NotNil(t, txCfg.AutoPurge.Threshold)
		assert.Equal(t, uint32(5), *txCfg.AutoPurge.Threshold)
		require.NotNil(t, txCfg.AutoPurge.MinAttempts)
		assert.Equal(t, uint32(3), *txCfg.AutoPurge.MinAttempts)
	})

	t.Run("respects operator override: AutoPurge explicitly enabled", func(t *testing.T) {
		txCfg := &Transactions{}
		txCfg.AutoPurge.Enabled = ptr(true)
		txCfg.AutoPurge.Threshold = ptr(uint32(10))
		txCfg.AutoPurge.MinAttempts = ptr(uint32(5))

		applied := ApplyFZeroDefaults(lggr, txCfg)

		assert.False(t, applied, "should not override operator-set values")
		assert.Equal(t, uint32(10), *txCfg.AutoPurge.Threshold)
		assert.Equal(t, uint32(5), *txCfg.AutoPurge.MinAttempts)
	})

	t.Run("respects operator override: AutoPurge explicitly disabled", func(t *testing.T) {
		txCfg := &Transactions{}
		txCfg.AutoPurge.Enabled = ptr(false)

		applied := ApplyFZeroDefaults(lggr, txCfg)

		assert.False(t, applied, "should not override explicit disable")
		assert.False(t, *txCfg.AutoPurge.Enabled, "should remain disabled")
		// Threshold and MinAttempts should NOT be set since AutoPurge is disabled
		assert.Nil(t, txCfg.AutoPurge.Threshold)
		assert.Nil(t, txCfg.AutoPurge.MinAttempts)
	})

	t.Run("sets Threshold and MinAttempts when only Enabled is set", func(t *testing.T) {
		txCfg := &Transactions{}
		txCfg.AutoPurge.Enabled = ptr(true)

		applied := ApplyFZeroDefaults(lggr, txCfg)

		assert.True(t, applied)
		assert.True(t, *txCfg.AutoPurge.Enabled)
		require.NotNil(t, txCfg.AutoPurge.Threshold)
		assert.Equal(t, uint32(5), *txCfg.AutoPurge.Threshold)
		require.NotNil(t, txCfg.AutoPurge.MinAttempts)
		assert.Equal(t, uint32(3), *txCfg.AutoPurge.MinAttempts)
	})

	t.Run("partial override: operator sets Threshold but not MinAttempts", func(t *testing.T) {
		txCfg := &Transactions{}
		txCfg.AutoPurge.Enabled = ptr(true)
		txCfg.AutoPurge.Threshold = ptr(uint32(8))

		applied := ApplyFZeroDefaults(lggr, txCfg)

		assert.True(t, applied, "should apply MinAttempts default")
		assert.Equal(t, uint32(8), *txCfg.AutoPurge.Threshold, "operator value preserved")
		assert.Equal(t, uint32(3), *txCfg.AutoPurge.MinAttempts, "default applied")
	})
}
