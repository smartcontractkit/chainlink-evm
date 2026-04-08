package client

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRPCClientMetrics(t *testing.T) {
	m, err := newRPCClientMetrics(big.NewInt(1))
	require.NoError(t, err)
	require.NotNil(t, m)
	assert.NotNil(t, m.rpcClientMetrics)
	assert.NotNil(t, m.callsTotal)
	assert.NotNil(t, m.callsSuccess)
	assert.NotNil(t, m.callsFailed)
}

func TestRPCClientMetrics_Increment(t *testing.T) {
	m, err := newRPCClientMetrics(big.NewInt(1))
	require.NoError(t, err)

	ctx := context.Background()

	assert.NotPanics(t, func() {
		m.RecordLatency(ctx, "rpc.example.com","eth_call", false, time.Second, nil)
	})
	assert.NotPanics(t, func() {
		m.IncrementTotal(ctx, "1", "node-1", "rpc.example.com", "eth_call")
	})
	assert.NotPanics(t, func() {
		m.IncrementSuccess(ctx, "1", "node-1", "rpc.example.com", "eth_call")
	})
	assert.NotPanics(t, func() {
		m.IncrementFailed(ctx, "1", "node-1", "rpc.example.com", "eth_call")
	})
}
