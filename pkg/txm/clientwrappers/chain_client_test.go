package clientwrappers

import (
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
)

func TestPendingNonceAtWithFallback_ReturnsNonce(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClientWithDefaultChainID(t)
	address := common.HexToAddress("0x1111111111111111111111111111111111111111")
	c, err := NewChainClient(logger.Test(t), m, true)
	require.NoError(t, err)

	m.On("PendingNonceAtWithFallback", mock.Anything, address).Return(uint64(10), nil).Once()

	nonce, err := c.PendingNonceAt(t.Context(), address)
	require.NoError(t, err)
	require.Equal(t, uint64(10), nonce)
}

func TestPendingNonceAtWithFallback_ErrorsWhenNoSuccessfulResults(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClientWithDefaultChainID(t)
	address := common.HexToAddress("0x2222222222222222222222222222222222222222")
	c, err := NewChainClient(logger.Test(t), m, true)
	require.NoError(t, err)

	m.On("PendingNonceAtWithFallback", mock.Anything, address).Return(uint64(0), errors.New("all nodes failed for pending nonce")).Once()

	_, err = c.PendingNonceAt(t.Context(), address)
	require.ErrorContains(t, err, "all nodes failed")
}

func TestNonceAtWithFallback_ReturnsNonce(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClientWithDefaultChainID(t)
	address := common.HexToAddress("0x3333333333333333333333333333333333333333")
	blockNumber := big.NewInt(7)
	c, err := NewChainClient(logger.Test(t), m, true)
	require.NoError(t, err)

	m.On("NonceAtWithFallback", mock.Anything, address, blockNumber).Return(uint64(11), nil).Once()

	nonce, err := c.NonceAt(t.Context(), address, blockNumber)
	require.NoError(t, err)
	require.Equal(t, uint64(11), nonce)
}

func TestNonceAtWithFallback_ErrorsWhenNoSuccessfulResults(t *testing.T) {
	t.Parallel()

	m := clienttest.NewClientWithDefaultChainID(t)
	address := common.HexToAddress("0x4444444444444444444444444444444444444444")
	c, err := NewChainClient(logger.Test(t), m, true)
	require.NoError(t, err)

	m.On("NonceAtWithFallback", mock.Anything, address, (*big.Int)(nil)).Return(uint64(0), errors.New("all nodes failed for nonce")).Once()

	_, err = c.NonceAt(t.Context(), address, nil)
	require.ErrorContains(t, err, "all nodes failed")
}
