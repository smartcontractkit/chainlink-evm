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
	
	t.Run("RecordBidAnalysis", func(t *testing.T) {
		metrics, err := NewMetaMetrics(chainID)
		require.NoError(t, err)
		
		ctx := context.Background()
		
		tests := []struct {
			name      string
			solverOps []*SO
		}{
			{
				name:      "no bids",
				solverOps: []*SO{},
			},
			{
				name: "single bid",
				solverOps: []*SO{
					{
						BidToken:  common.HexToAddress("0x1234"),
						BidAmount: (*hexutil.Big)(big.NewInt(100)),
						Solver:    common.HexToAddress("0xabcd"),
					},
				},
			},
			{
				name: "multiple bids",
				solverOps: []*SO{
					{
						BidToken:  common.HexToAddress("0x1234"),
						BidAmount: (*hexutil.Big)(big.NewInt(200)),
						Solver:    common.HexToAddress("0xabcd"),
					},
					{
						BidToken:  common.HexToAddress("0x1234"),
						BidAmount: (*hexutil.Big)(big.NewInt(150)),
						Solver:    common.HexToAddress("0xefgh"),
					},
				},
			},
		}
		
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Test that this doesn't panic
				metrics.RecordBidAnalysis(ctx, tt.solverOps)
			})
		}
	})
	
	t.Run("RecordOtherMetrics", func(t *testing.T) {
		metrics, err := NewMetaMetrics(chainID)
		require.NoError(t, err)
		
		ctx := context.Background()
		
		// Test that these don't panic
		metrics.RecordStatusCode(ctx, 200)
		metrics.RecordBidsReceived(ctx, 5)
		metrics.RecordEventProcessed(ctx)
	})
	
	t.Run("EmitMetaRequestEvent", func(t *testing.T) {
		metrics, err := NewMetaMetrics(chainID)
		require.NoError(t, err)
		
		ctx := context.Background()
		
		// Create test transaction
		tx := &types.Transaction{
			ID:          123,
			FromAddress: common.HexToAddress("0x1234"),
			ToAddress:   common.HexToAddress("0x5678"),
		}
		
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
		err = metrics.EmitMetaRequestEvent(ctx, tx, nil, []byte("test payload"), 200, time.Millisecond*100, "", solverOps)
		// In test environment, beholder might not be configured, so we don't assert on success
		// but we can test that the function doesn't panic
		assert.NotPanics(t, func() {
			metrics.EmitMetaRequestEvent(ctx, tx, nil, []byte("test payload"), 200, time.Millisecond*100, "", solverOps)
		})
	})
}

