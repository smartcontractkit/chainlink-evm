package dualbroadcast

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetaMetrics(t *testing.T) {
	chainID := "1"

	t.Run("NewMetaMetrics", func(t *testing.T) {
		metrics, err := NewMetaMetrics(chainID)
		require.NoError(t, err)
		assert.NotNil(t, metrics)
		assert.Equal(t, chainID, metrics.chainID)
	})

	t.Run("RecordBasicMetrics", func(t *testing.T) {
		metrics, err := NewMetaMetrics(chainID)
		require.NoError(t, err)

		ctx := t.Context()

		// Test that these don't panic - all metrics methods
		metrics.RecordStatusCode(ctx, 200)
		metrics.RecordLatency(ctx, time.Millisecond*100)
		metrics.RecordBidsReceived(ctx, 5)
		metrics.RecordSendRequestError(ctx)
		metrics.RecordSendOperationError(ctx)
	})
}
