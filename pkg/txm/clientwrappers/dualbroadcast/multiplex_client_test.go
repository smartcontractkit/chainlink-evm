package dualbroadcast

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

type mockClient struct {
	sendTxFn       func(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error
	pendingNonceFn func(ctx context.Context, address common.Address) (uint64, error)
	nonceAtFn      func(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error)
	sendCalled     chan struct{}
}

func (m *mockClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	if m.sendCalled != nil {
		defer func() { m.sendCalled <- struct{}{} }()
	}
	if m.sendTxFn != nil {
		return m.sendTxFn(ctx, tx, attempt)
	}
	return nil
}

func (m *mockClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	if m.pendingNonceFn != nil {
		return m.pendingNonceFn(ctx, address)
	}
	return 0, nil
}

func (m *mockClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	if m.nonceAtFn != nil {
		return m.nonceAtFn(ctx, address, blockNumber)
	}
	return 0, nil
}

func TestMultiplexClient_SendTransaction_BothSucceed(t *testing.T) {
	secondaryCalled := make(chan struct{}, 1)
	primary := &mockClient{}
	secondary := &mockClient{sendCalled: secondaryCalled}

	mc := NewMultiplexClient(logger.Test(t), primary, secondary)
	err := mc.SendTransaction(context.Background(), &types.Transaction{}, &types.Attempt{})
	require.NoError(t, err)

	select {
	case <-secondaryCalled:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary was not called")
	}
}

func TestMultiplexClient_SendTransaction_PrimaryFails(t *testing.T) {
	primaryErr := errors.New("flashbots rejected")
	primary := &mockClient{
		sendTxFn: func(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
			return primaryErr
		},
	}
	secondary := &mockClient{sendCalled: make(chan struct{}, 1)}

	mc := NewMultiplexClient(logger.Test(t), primary, secondary)
	err := mc.SendTransaction(context.Background(), &types.Transaction{}, &types.Attempt{})
	require.ErrorIs(t, err, primaryErr)
}

func TestMultiplexClient_SendTransaction_SecondaryFails(t *testing.T) {
	secondaryCalled := make(chan struct{}, 1)
	primary := &mockClient{}
	secondary := &mockClient{
		sendTxFn: func(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
			return errors.New("nova failed")
		},
		sendCalled: secondaryCalled,
	}

	mc := NewMultiplexClient(logger.Test(t), primary, secondary)
	err := mc.SendTransaction(context.Background(), &types.Transaction{}, &types.Attempt{})
	require.NoError(t, err)

	select {
	case <-secondaryCalled:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary was not called")
	}
}

func TestMultiplexClient_PendingNonceAt_RoutesToPrimary(t *testing.T) {
	primary := &mockClient{
		pendingNonceFn: func(ctx context.Context, address common.Address) (uint64, error) {
			return 42, nil
		},
	}
	secondary := &mockClient{
		pendingNonceFn: func(ctx context.Context, address common.Address) (uint64, error) {
			t.Fatal("secondary PendingNonceAt should not be called")
			return 0, nil
		},
	}

	mc := NewMultiplexClient(logger.Test(t), primary, secondary)
	nonce, err := mc.PendingNonceAt(context.Background(), common.HexToAddress("0x123"))
	require.NoError(t, err)
	assert.Equal(t, uint64(42), nonce)
}

func TestMultiplexClient_NonceAt_RoutesToPrimary(t *testing.T) {
	primary := &mockClient{
		nonceAtFn: func(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
			return 99, nil
		},
	}
	secondary := &mockClient{
		nonceAtFn: func(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
			t.Fatal("secondary NonceAt should not be called")
			return 0, nil
		},
	}

	mc := NewMultiplexClient(logger.Test(t), primary, secondary)
	nonce, err := mc.NonceAt(context.Background(), common.HexToAddress("0x123"), big.NewInt(100))
	require.NoError(t, err)
	assert.Equal(t, uint64(99), nonce)
}
