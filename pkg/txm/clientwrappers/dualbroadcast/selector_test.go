package dualbroadcast

import (
	"math/big"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
)

func testTxmMetrics(t *testing.T) txm.Metrics {
	t.Helper()
	return txm.NewTxmMetrics(logger.Test(t), big.NewInt(1))
}

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
		big.NewInt(1), nil, false, nil, nil, testTxmMetrics(t))
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

	c, eh, err := SelectClient(logger.Test(t), mockClient, nil, urls, big.NewInt(1), nil, false, nil, nil, testTxmMetrics(t))
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
		big.NewInt(1), nil, false, nil, nil, testTxmMetrics(t))
	require.NoError(t, err)
	assert.NotNil(t, c)

	mux, ok := c.(*multiOfaClient)
	require.True(t, ok)
	assert.Empty(t, mux.secondaries)
	nc, ok := mux.primary.(*ofaBackend)
	require.True(t, ok, "nova URL should use OFA primary inside multiplex shell")
	assert.Equal(t, ofaNova, nc.ofa)
}

func TestSelectClient_FastlaneSingleAuctioneerURL_ReturnsMetaClient(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	c, eh, err := SelectClient(logger.Test(t), mockClient, nil,
		[]*url.URL{mustParseURL(t, "https://auctioneer.fastlane.example.com/v1/submit")},
		big.NewInt(1), nil, false, nil, nil, testTxmMetrics(t))
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.NotNil(t, eh)

	_, ok := c.(*MetaClient)
	require.True(t, ok)
}

func TestSelectClient_NonRelaySingleURL_Errors(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	c, eh, err := SelectClient(logger.Test(t), mockClient, nil,
		[]*url.URL{mustParseURL(t, "https://custom-auction.example.com")},
		big.NewInt(1), nil, false, nil, nil, testTxmMetrics(t))
	require.Error(t, err)
	require.Nil(t, c)
	require.Nil(t, eh)
	assert.Contains(t, err.Error(), "does not support OFA URL")
}

func TestSelectClient_MultipleURLsWithAuctioneerPrimary_Errors(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	urls := []*url.URL{
		mustParseURL(t, "https://auctioneer.fastlane.example.com"),
		mustParseURL(t, "https://relay.flashbots.net"),
	}
	c, eh, err := SelectClient(logger.Test(t), mockClient, nil, urls, big.NewInt(1), nil, false, nil, nil, testTxmMetrics(t))
	require.Error(t, err)
	require.Nil(t, c)
	require.Nil(t, eh)
	assert.Contains(t, err.Error(), "does not support OFA URL")
}
