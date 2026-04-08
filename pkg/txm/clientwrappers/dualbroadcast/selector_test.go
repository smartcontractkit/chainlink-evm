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

	primaryURL := mustParseURL(t, "https://relay.flashbots.net")

	c, eh, err := SelectClient(logger.Test(t), mockClient, nil, primaryURL, nil, big.NewInt(1), nil, false, nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Nil(t, eh)

	_, isMultiplex := c.(*multiplexClient)
	assert.False(t, isMultiplex, "should return FlashbotsClient directly, not wrapped in multiplexClient")
}

func TestSelectClient_FlashbotsPrimaryWithNovaSecondary(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	primaryURL := mustParseURL(t, "https://relay.flashbots.net")
	secondaryURL := mustParseURL(t, "https://eth.novarpc.xyz?api_key=test")

	c, eh, err := SelectClient(logger.Test(t), mockClient, nil, primaryURL, secondaryURL, big.NewInt(1), nil, false, nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Nil(t, eh)

	_, isMultiplex := c.(*multiplexClient)
	assert.True(t, isMultiplex, "should return a multiplexClient when secondary URL is provided")
}

func TestSelectClient_RejectsNonNovaSecondary(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	primaryURL := mustParseURL(t, "https://relay.flashbots.net")
	secondaryURL := mustParseURL(t, "https://relay.flashbots.net")

	_, _, err := SelectClient(logger.Test(t), mockClient, nil, primaryURL, secondaryURL, big.NewInt(1), nil, false, nil, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "secondary URL must be a Nova RPC endpoint")
}

func TestSelectClient_NovaPrimaryOnly(t *testing.T) {
	mockClient := clienttest.NewClient(t)
	mockClient.EXPECT().ConfiguredChainID().Return(big.NewInt(1))

	primaryURL := mustParseURL(t, "https://eth.novarpc.xyz?api_key=test")

	c, _, err := SelectClient(logger.Test(t), mockClient, nil, primaryURL, nil, big.NewInt(1), nil, false, nil, nil)
	require.NoError(t, err)
	assert.NotNil(t, c)

	_, isNova := c.(*novaClient)
	assert.True(t, isNova, "nova URL should create a novaClient")
}

func TestRedactURL(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "", redactURL(nil))

	u := mustParseURL(t, "https://eth.novarpc.xyz?api_key=secret&foo=bar")
	assert.Equal(t, "https://eth.novarpc.xyz?api_key=xxxxx&foo=bar", redactURL(u))
	assert.Equal(t, "secret", u.Query().Get("api_key"), "must not mutate original URL")

	uPass := mustParseURL(t, "https://user:pass@eth.novarpc.xyz?api_key=secret")
	assert.Equal(t, "https://user:xxxxx@eth.novarpc.xyz?api_key=xxxxx", redactURL(uPass))
}
