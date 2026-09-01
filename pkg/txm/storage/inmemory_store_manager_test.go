package storage

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	fromAddress := testutils.NewAddress()
	m := NewInMemoryStoreManager(logger.Test(t), testutils.FixtureChainID)
	// Adds a new address
	err := m.Add(fromAddress)
	require.NoError(t, err)
	existingStore, exists := m.GetStoreSafe(fromAddress)
	require.True(t, exists)

	// Fails if address exists
	err = m.Add(fromAddress)
	require.Error(t, err)
	store, exists := m.GetStoreSafe(fromAddress)
	require.True(t, exists)
	assert.Same(t, existingStore, store)

	// Adds multiple addresses
	fromAddress1 := testutils.NewAddress()
	fromAddress2 := testutils.NewAddress()
	addresses := []common.Address{fromAddress1, fromAddress2}
	err = m.Add(addresses...)
	require.NoError(t, err)
	_, exists = m.GetStoreSafe(fromAddress1)
	require.True(t, exists)
	_, exists = m.GetStoreSafe(fromAddress2)
	require.True(t, exists)
}
