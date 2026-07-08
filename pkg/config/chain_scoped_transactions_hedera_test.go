package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/config/toml"
)

func TestTransactionsConfig_HederaBroadcastValidation(t *testing.T) {
	t.Parallel()

	t.Run("returns nil when unset", func(t *testing.T) {
		t.Parallel()

		cfg := &transactionsConfig{c: toml.Transactions{}}
		assert.Nil(t, cfg.HederaSequencePollTimeout())
		assert.Nil(t, cfg.HederaSequencePollInterval())
	})

	t.Run("returns configured values", func(t *testing.T) {
		t.Parallel()

		timeout := commonconfig.MustNewDuration(30 * time.Second)
		interval := commonconfig.MustNewDuration(2 * time.Second)
		cfg := &transactionsConfig{c: toml.Transactions{
			HederaBroadcastValidation: toml.HederaBroadcastValidationConfig{
				SequencePollTimeout:  timeout,
				SequencePollInterval: interval,
			},
		}}

		gotTimeout := cfg.HederaSequencePollTimeout()
		require.NotNil(t, gotTimeout)
		assert.Equal(t, 30*time.Second, *gotTimeout)

		gotInterval := cfg.HederaSequencePollInterval()
		require.NotNil(t, gotInterval)
		assert.Equal(t, 2*time.Second, *gotInterval)
	})
}

func TestHederaBroadcastValidationConfig_ValidateConfig(t *testing.T) {
	t.Parallel()

	t.Run("accepts valid config", func(t *testing.T) {
		t.Parallel()

		cfg := toml.HederaBroadcastValidationConfig{
			SequencePollTimeout: commonconfig.MustNewDuration(time.Second),
		}
		require.NoError(t, cfg.ValidateConfig())
	})

	t.Run("rejects interval without timeout", func(t *testing.T) {
		t.Parallel()

		cfg := toml.HederaBroadcastValidationConfig{
			SequencePollInterval: commonconfig.MustNewDuration(time.Second),
		}
		require.Error(t, cfg.ValidateConfig())
	})
}
