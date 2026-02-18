package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRPCClientMetrics(t *testing.T) {
	m, err := newRPCClientMetrics()
	require.NoError(t, err)
	require.NotNil(t, m)
	assert.NotNil(t, m.callsTotal)
	assert.NotNil(t, m.callsSuccess)
	assert.NotNil(t, m.callsFailed)
}

func TestRPCClientMetrics_Increment(t *testing.T) {
	m, err := newRPCClientMetrics()
	require.NoError(t, err)

	ctx := context.Background()

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
