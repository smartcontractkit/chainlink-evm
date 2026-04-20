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

	_, isMultiplex := c.(*multiplexClient)
	assert.False(t, isMultiplex, "should return FlashbotsClient directly, not wrapped in multiplexClient")
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

	_, isMultiplex := c.(*multiplexClient)
	assert.True(t, isMultiplex, "should return a multiplexClient when more than one URL is provided")
}

func TestSelectClient_NovaPrimaryOnly(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	c, _, err := SelectClient(logger.Test(t), mockClient, nil,
		[]*url.URL{mustParseURL(t, "https://eth.novarpc.xyz?api_key=test")},
		big.NewInt(1), nil, false, nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, c)

	nc, isOFA := c.(*ofaTXClient)
	assert.True(t, isOFA, "nova URL should create an OFA client")
	assert.Equal(t, ofaKindNova, nc.ofa.kind)
}

func TestRedactURL(t *testing.T) {
	t.Parallel()

	assert.Empty(t, redactURL(nil))

	u := mustParseURL(t, "https://eth.novarpc.xyz?api_key=secret&foo=bar")
	assert.Equal(t, "https://eth.novarpc.xyz?api_key=xxxxx&foo=bar", redactURL(u))
	assert.Equal(t, "secret", u.Query().Get("api_key"), "must not mutate original URL")

	uPass := mustParseURL(t, "https://user:pass@eth.novarpc.xyz?api_key=secret")
	assert.Equal(t, "https://user:xxxxx@eth.novarpc.xyz?api_key=xxxxx", redactURL(uPass))
}
