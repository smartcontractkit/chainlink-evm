package dualbroadcast

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetaMetrics(t *testing.T) {
	chainID := big.NewInt(1)

	t.Run("NewMetaMetrics", func(t *testing.T) {
		metrics, err := NewMetaMetrics(chainID)
		require.NoError(t, err)
		assert.NotNil(t, metrics)
		assert.Equal(t, chainID, metrics.chainID)
	})

	t.Run("RecordBasicMetrics", func(t *testing.T) {
		metrics, err := NewMetaMetrics(chainID)
		require.NoError(t, err)

		ctx := context.Background()

		// Test that these don't panic - only status code and latency remain
		metrics.RecordStatusCode(ctx, 200)
		metrics.RecordLatency(ctx, time.Millisecond*100)
	})

}
