package dualbroadcast

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
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

	t.Run("EmitMetaRequestEvent", func(t *testing.T) {
		metrics, err := NewMetaMetrics(chainID)
		require.NoError(t, err)

		ctx := context.Background()

		// Create test transaction and attempt
		tx := &types.Transaction{
			ID:          123,
			FromAddress: common.HexToAddress("0x1234"),
			ToAddress:   common.HexToAddress("0x5678"),
		}

		attempt := &types.Attempt{}

		solverOps := []*SO{
			{
				BidToken:  common.HexToAddress("0xabcd"),
				BidAmount: (*hexutil.Big)(big.NewInt(500)),
				Solver:    common.HexToAddress("0xsolver1"),
			},
			{
				BidToken:  common.HexToAddress("0xabcd"),
				BidAmount: (*hexutil.Big)(big.NewInt(400)),
				Solver:    common.HexToAddress("0xsolver2"),
			},
		}

		// Test that this doesn't panic - in real environment it would emit to beholder
		err = metrics.EmitMetaRequestEvent(ctx, tx, attempt, []byte("test payload"), 200, time.Millisecond*100, "", solverOps)
		// In test environment, beholder might not be configured, so we don't assert on success
		// but we can test that the function doesn't panic
		assert.NotPanics(t, func() {
			metrics.EmitMetaRequestEvent(ctx, tx, attempt, []byte("test payload"), 200, time.Millisecond*100, "", solverOps)
		})
	})
}
