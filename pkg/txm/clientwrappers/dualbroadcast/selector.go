package dualbroadcast

import (
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
)

func SelectClient(lggr logger.Logger, client client.Client, keyStore keys.ChainStore, primaryURL *url.URL, secondaryURL *url.URL, chainID *big.Int, txStore txm.TxStore, readRequestsToMultipleNodes bool, bundles *bool, auctionRequestTimeout *time.Duration, tier2Feeds []common.Address) (txm.Client, txm.ErrorHandler, error) {
	chainClient, err := clientwrappers.NewChainClient(lggr, client, readRequestsToMultipleNodes)
	if err != nil {
		return nil, nil, err
	}

	primary, errHandler, err := selectSingleClient(lggr, chainClient, keyStore, primaryURL, chainID, txStore, bundles, auctionRequestTimeout, tier2Feeds)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create primary client for %s: %w", redactURL(primaryURL), err)
	}

	if secondaryURL == nil {
		return primary, errHandler, nil
	}

	// secondary must be a Nova RPC endpoint
	if !strings.Contains(secondaryURL.String(), "novarpc") {
		return nil, nil, fmt.Errorf("secondary URL must be a Nova RPC endpoint, got: %s", redactURL(secondaryURL))
	}

	secondary, _, err := selectSingleClient(lggr, chainClient, keyStore, secondaryURL, chainID, txStore, bundles, auctionRequestTimeout, tier2Feeds)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create secondary client for %s: %w", redactURL(secondaryURL), err)
	}

	return newMultiplexClient(lggr, primary, secondary), errHandler, nil
}

func selectSingleClient(lggr logger.Logger, chainClient *clientwrappers.ChainClient, keyStore keys.ChainStore, u *url.URL, chainID *big.Int, txStore txm.TxStore, bundles *bool, auctionRequestTimeout *time.Duration, tier2Feeds []common.Address) (txm.Client, txm.ErrorHandler, error) {
	urlString := u.String()
	switch {
	case strings.Contains(urlString, "flashbots"):
		metrics, err := newOFAMetrics(chainID.String(), "flashbots")
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create OFA metrics for flashbots: %w", err)
		}
		return NewFlashbotsClient(lggr, chainClient, keyStore, u, txStore, bundles, metrics), nil, nil
	case strings.Contains(urlString, "novarpc"):
		metrics, err := newOFAMetrics(chainID.String(), "nova")
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create OFA metrics for nova: %w", err)
		}
		return newNovaClient(lggr, chainClient, u, metrics, keyStore, tier2Feeds), nil, nil
	default:
		mc, err := NewMetaClient(lggr, chainClient, keyStore, u, chainID, txStore, auctionRequestTimeout)
		if err != nil {
			return nil, nil, err
		}
		return mc, NewErrorHandler(), nil
	}
}

// redactURL returns u as a string safe for logs: same redaction as url.URL.Redacted for userinfo, and api_key query values are replaced with "xxxxx". It does not mutate the original URL.
func redactURL(u *url.URL) string {
	if u == nil {
		return ""
	}
	cp := *u
	q := cp.Query()
	if _, has := q["api_key"]; has {
		q.Set("api_key", "xxxxx")
		cp.RawQuery = q.Encode()
	}
	return cp.Redacted()
}
