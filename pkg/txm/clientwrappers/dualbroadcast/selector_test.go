package dualbroadcast

import (
	"math/big"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
)

func mustParseURL(t *testing.T, raw string) *url.URL {
	t.Helper()
	u, err := url.Parse(raw)
	require.NoError(t, err)
	return u
}

func TestSelectClient_FlashbotsPrimaryOnly(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	c, eh, err := SelectClient(logger.Test(t), mockClient, nil,
		[]*url.URL{mustParseURL(t, "https://relay.flashbots.net")},
		big.NewInt(1), nil, false, nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Nil(t, eh)

	mux, ok := c.(*multiOfaClient)
	require.True(t, ok, "single flashbots URL should still use multiOfaClient as the outer shell")
	assert.Empty(t, mux.secondaries)
	fb, ok := mux.primary.(*ofaBackend)
	require.True(t, ok)
	assert.Equal(t, ofaFlashbots, fb.ofa)
}

func TestSelectClient_FlashbotsPrimaryWithNovaSecondary(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	urls := []*url.URL{
		mustParseURL(t, "https://relay.flashbots.net"),
		mustParseURL(t, "https://eth.novarpc.xyz?api_key=test"),
	}

	c, eh, err := SelectClient(logger.Test(t), mockClient, nil, urls, big.NewInt(1), nil, false, nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Nil(t, eh)

	mux, ok := c.(*multiOfaClient)
	require.True(t, ok)
	pri, ok := mux.primary.(*ofaBackend)
	require.True(t, ok)
	assert.Equal(t, ofaFlashbots, pri.ofa)
	require.Len(t, mux.secondaries, 1)
	sec, ok := mux.secondaries[0].(*ofaBackend)
	require.True(t, ok)
	assert.Equal(t, ofaNova, sec.ofa)
}

func TestSelectClient_NovaPrimaryOnly(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	c, _, err := SelectClient(logger.Test(t), mockClient, nil,
		[]*url.URL{mustParseURL(t, "https://eth.novarpc.xyz?api_key=test")},
		big.NewInt(1), nil, false, nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, c)

	mux, ok := c.(*multiOfaClient)
	require.True(t, ok)
	assert.Empty(t, mux.secondaries)
	nc, ok := mux.primary.(*ofaBackend)
	require.True(t, ok, "nova URL should use OFA primary inside multiplex shell")
	assert.Equal(t, ofaNova, nc.ofa)
}

func TestSelectClient_MetaPrimaryOnly(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	c, eh, err := SelectClient(logger.Test(t), mockClient, nil,
		[]*url.URL{mustParseURL(t, "https://custom-auction.example.com")},
		big.NewInt(1), nil, false, nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.NotNil(t, eh)

	_, ok := c.(*MetaClient)
	require.True(t, ok, "Meta auction URL should return MetaClient directly, not multiOfaClient")
}

func TestSelectClient_IgnoresNonPrimaryURLsWhenPrimaryIsMeta(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	urls := []*url.URL{
		mustParseURL(t, "https://custom-auction-a.example.com"),
		mustParseURL(t, "https://custom-auction-b.example.com"),
	}
	c, eh, err := SelectClient(logger.Test(t), mockClient, nil, urls, big.NewInt(1), nil, false, nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.NotNil(t, eh)

	_, ok := c.(*MetaClient)
	require.True(t, ok, "Meta auction URL should return MetaClient directly, not multiOfaClient")
}
