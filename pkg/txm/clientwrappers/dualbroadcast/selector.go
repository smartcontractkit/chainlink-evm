package dualbroadcast

import (
	"fmt"
	"math/big"
	"net/url"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
)

// SelectClient builds the txm.Client for dual broadcast. ofaURLs must be non-empty; index 0 is the
// primary (determines broadcast outcome and nonce queries). The implementation is always a multiplexClient:
// it classifies each URL (Flashbots, Nova, Meta, …), builds one backend per entry, and with one URL has
// no secondaries; with more, additional URLs are best-effort secondaries (fire-and-forget with a separate timeout).
func SelectClient(lggr logger.Logger, client client.Client, keyStore keys.ChainStore, ofaURLs []*url.URL, chainID *big.Int, txStore txm.TxStore, readRequestsToMultipleNodes bool, bundles *bool, auctionRequestTimeout *time.Duration) (txm.Client, txm.ErrorHandler, error) {
	if len(ofaURLs) == 0 {
		return nil, nil, fmt.Errorf("ofaURLs must not be empty")
	}

	chainClient, err := clientwrappers.NewChainClient(lggr, client, readRequestsToMultipleNodes)
	if err != nil {
		return nil, nil, err
	}

	return newMultiplexClientFromOFAURLs(lggr, chainClient, keyStore, ofaURLs, chainID, txStore, bundles, auctionRequestTimeout)
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
