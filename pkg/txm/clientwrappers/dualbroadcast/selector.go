package dualbroadcast

import (
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
)

// SelectClient builds the txm.Client for dual broadcast. ofaURLs must be non-empty; index 0 is the
// primary (determines broadcast outcome and nonce queries). The client is always a multiplexClient:
// with one URL it has no secondaries; with more, additional URLs are best-effort secondaries
// (fire-and-forget with a separate timeout).
func SelectClient(lggr logger.Logger, client client.Client, keyStore keys.ChainStore, ofaURLs []*url.URL, chainID *big.Int, txStore txm.TxStore, readRequestsToMultipleNodes bool, bundles *bool, auctionRequestTimeout *time.Duration) (txm.Client, txm.ErrorHandler, error) {
	if len(ofaURLs) == 0 {
		return nil, nil, fmt.Errorf("ofaURLs must not be empty")
	}

	chainClient, err := clientwrappers.NewChainClient(lggr, client, readRequestsToMultipleNodes)
	if err != nil {
		return nil, nil, err
	}

	primary, errHandler, err := selectSingleClient(lggr, chainClient, keyStore, ofaURLs[0], chainID, txStore, bundles, auctionRequestTimeout)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create primary client for %s: %w", redactURL(ofaURLs[0]), err)
	}

	secondaries := make([]txm.Client, 0, len(ofaURLs)-1)
	for _, u := range ofaURLs[1:] {
		sec, _, err := selectSingleClient(lggr, chainClient, keyStore, u, chainID, txStore, bundles, auctionRequestTimeout)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create secondary client for %s: %w", redactURL(u), err)
		}
		secondaries = append(secondaries, sec)
	}

	urlStrs := make([]string, len(ofaURLs))
	for i, u := range ofaURLs {
		urlStrs[i] = redactURL(u)
	}

	lggr.Infow("TransactionManagerV2 OFA client created",
		"primaryURL", urlStrs[0],
		"secondaryURLs", urlStrs[1:])

	return newMultiplexClient(lggr, primary, secondaries...), errHandler, nil
}

func selectSingleClient(lggr logger.Logger, chainClient *clientwrappers.ChainClient, keyStore keys.ChainStore, u *url.URL, chainID *big.Int, txStore txm.TxStore, bundles *bool, auctionRequestTimeout *time.Duration) (txm.Client, txm.ErrorHandler, error) {
	urlString := u.String()
	switch {
	case strings.Contains(urlString, "flashbots"):
		metrics, err := newOFAMetrics(chainID.String(), ofaKindFlashbots.name())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create OFA metrics for flashbots: %w", err)
		}
		bundlesEnabled := bundles != nil && *bundles
		return newFlashbotsClient(lggr, chainClient, keyStore, u, txStore, bundlesEnabled, metrics), nil, nil
	case strings.Contains(urlString, "novarpc"):
		metrics, err := newOFAMetrics(chainID.String(), ofaKindNova.name())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create OFA metrics for nova: %w", err)
		}
		return newNovaClient(lggr, chainClient, u, metrics), nil, nil
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
