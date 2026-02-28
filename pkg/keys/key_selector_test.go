package keys

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

var (
	addr1 = common.HexToAddress("0x1111111111111111111111111111111111111111")
	addr2 = common.HexToAddress("0x2222222222222222222222222222222222222222")
	addr3 = common.HexToAddress("0x3333333333333333333333333333333333333333")
)

func TestNewSelector(t *testing.T) {
	lggr := logger.Test(t)

	t.Run("fails with empty addresses", func(t *testing.T) {
		_, err := NewSelector(lggr, []common.Address{}, StrategyHealthBasedFallback)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "at least one sending key address is required")
	})

	t.Run("fails with unknown strategy", func(t *testing.T) {
		_, err := NewSelector(lggr, []common.Address{addr1}, "UnknownStrategy")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unknown key selection strategy")
	})

	t.Run("creates health-based selector", func(t *testing.T) {
		sel, err := NewSelector(lggr, []common.Address{addr1, addr2}, StrategyHealthBasedFallback)
		require.NoError(t, err)
		require.NotNil(t, sel)
	})

	t.Run("creates round-robin selector", func(t *testing.T) {
		sel, err := NewSelector(lggr, []common.Address{addr1, addr2}, StrategyRoundRobin)
		require.NoError(t, err)
		require.NotNil(t, sel)
	})
}

func TestHealthBasedSelector(t *testing.T) {
	lggr := logger.Test(t)
	ctx := context.Background()

	t.Run("selects primary key when healthy", func(t *testing.T) {
		sel := newHealthBasedSelector(lggr, []common.Address{addr1, addr2})

		selected, err := sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr1, selected)

		// Should keep returning primary
		selected, err = sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr1, selected)
	})

	t.Run("falls back to secondary when primary is unhealthy", func(t *testing.T) {
		sel := newHealthBasedSelector(lggr, []common.Address{addr1, addr2})

		sel.MarkUnhealthy(addr1, "stuck nonce")

		selected, err := sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr2, selected)
	})

	t.Run("returns to primary after recovery", func(t *testing.T) {
		sel := newHealthBasedSelector(lggr, []common.Address{addr1, addr2})

		sel.MarkUnhealthy(addr1, "stuck nonce")
		selected, err := sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr2, selected)

		sel.MarkHealthy(addr1)
		selected, err = sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr1, selected)
	})

	t.Run("falls back to primary when all keys are unhealthy", func(t *testing.T) {
		sel := newHealthBasedSelector(lggr, []common.Address{addr1, addr2})

		sel.MarkUnhealthy(addr1, "stuck nonce")
		sel.MarkUnhealthy(addr2, "insufficient funds")

		selected, err := sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr1, selected, "should fall back to primary even when unhealthy")
	})

	t.Run("cascades through multiple fallbacks", func(t *testing.T) {
		sel := newHealthBasedSelector(lggr, []common.Address{addr1, addr2, addr3})

		sel.MarkUnhealthy(addr1, "stuck")
		sel.MarkUnhealthy(addr2, "stuck")

		selected, err := sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr3, selected)
	})

	t.Run("single key always returns that key", func(t *testing.T) {
		sel := newHealthBasedSelector(lggr, []common.Address{addr1})

		selected, err := sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr1, selected)

		sel.MarkUnhealthy(addr1, "stuck")
		selected, err = sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr1, selected)
	})

	t.Run("HealthStatus returns all keys", func(t *testing.T) {
		sel := newHealthBasedSelector(lggr, []common.Address{addr1, addr2})

		sel.MarkUnhealthy(addr2, "insufficient funds")

		statuses := sel.HealthStatus()
		require.Len(t, statuses, 2)
		assert.True(t, statuses[0].Healthy)
		assert.Equal(t, addr1, statuses[0].Address)
		assert.False(t, statuses[1].Healthy)
		assert.Equal(t, "insufficient funds", statuses[1].Reason)
	})

	t.Run("ActiveAddresses returns configured addresses", func(t *testing.T) {
		sel := newHealthBasedSelector(lggr, []common.Address{addr1, addr2})
		assert.Equal(t, []common.Address{addr1, addr2}, sel.ActiveAddresses())
	})
}

func TestRoundRobinSelector(t *testing.T) {
	lggr := logger.Test(t)
	ctx := context.Background()

	t.Run("cycles through keys", func(t *testing.T) {
		sel := newRoundRobinSelector(lggr, []common.Address{addr1, addr2, addr3})

		selected1, err := sel.SelectKey(ctx)
		require.NoError(t, err)

		selected2, err := sel.SelectKey(ctx)
		require.NoError(t, err)

		selected3, err := sel.SelectKey(ctx)
		require.NoError(t, err)

		// Should have cycled through all three
		keys := []common.Address{selected1, selected2, selected3}
		assert.Contains(t, keys, addr1)
		assert.Contains(t, keys, addr2)
		assert.Contains(t, keys, addr3)

		// Fourth call should wrap around
		selected4, err := sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, selected1, selected4)
	})

	t.Run("skips unhealthy keys", func(t *testing.T) {
		sel := newRoundRobinSelector(lggr, []common.Address{addr1, addr2, addr3})

		sel.MarkUnhealthy(addr2, "stuck")

		// Should get addr1, skip addr2, get addr3, then back to addr1
		results := make([]common.Address, 4)
		for i := 0; i < 4; i++ {
			selected, err := sel.SelectKey(ctx)
			require.NoError(t, err)
			results[i] = selected
		}

		for _, r := range results {
			assert.NotEqual(t, addr2, r, "should never select unhealthy addr2")
		}
	})

	t.Run("recovers unhealthy key back into rotation", func(t *testing.T) {
		sel := newRoundRobinSelector(lggr, []common.Address{addr1, addr2})

		sel.MarkUnhealthy(addr2, "stuck")

		selected, err := sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr1, selected)

		selected, err = sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Equal(t, addr1, selected, "should keep returning addr1 while addr2 is unhealthy")

		sel.MarkHealthy(addr2)

		// Now both should be in rotation
		seen := map[common.Address]bool{}
		for i := 0; i < 4; i++ {
			selected, err = sel.SelectKey(ctx)
			require.NoError(t, err)
			seen[selected] = true
		}
		assert.True(t, seen[addr1], "addr1 should be in rotation")
		assert.True(t, seen[addr2], "addr2 should be back in rotation")
	})

	t.Run("all unhealthy uses next in rotation", func(t *testing.T) {
		sel := newRoundRobinSelector(lggr, []common.Address{addr1, addr2})

		sel.MarkUnhealthy(addr1, "stuck")
		sel.MarkUnhealthy(addr2, "stuck")

		// Should still return keys (not error)
		selected, err := sel.SelectKey(ctx)
		require.NoError(t, err)
		assert.Contains(t, []common.Address{addr1, addr2}, selected)
	})

	t.Run("HealthStatus reflects state", func(t *testing.T) {
		sel := newRoundRobinSelector(lggr, []common.Address{addr1, addr2})

		statuses := sel.HealthStatus()
		require.Len(t, statuses, 2)
		assert.True(t, statuses[0].Healthy)
		assert.True(t, statuses[1].Healthy)

		sel.MarkUnhealthy(addr1, "test reason")
		statuses = sel.HealthStatus()
		assert.False(t, statuses[0].Healthy)
		assert.Equal(t, "test reason", statuses[0].Reason)
		assert.True(t, statuses[1].Healthy)
	})
}
