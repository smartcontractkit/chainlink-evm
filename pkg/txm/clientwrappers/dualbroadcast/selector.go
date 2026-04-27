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

// SelectClient builds the txm.Client for OFA broadcast. ofaURLs must be non-empty; index 0 is the
// primary OFA.
//
// Currently supports:
// * a single fastlane OFA url
// * any Flashbots/Nova urls (either standalone or combined)
func SelectClient(lggr logger.Logger, client client.Client, keyStore keys.ChainStore, ofaURLs []*url.URL, chainID *big.Int, txStore txm.TxStore, readRequestsToMultipleNodes bool, bundles *bool, auctionRequestTimeout *time.Duration) (txm.Client, txm.ErrorHandler, error) {
	if len(ofaURLs) == 0 {
		return nil, nil, fmt.Errorf("ofaURLs must not be empty")
	}

	chainClient, err := clientwrappers.NewChainClient(lggr, client, readRequestsToMultipleNodes)
	if err != nil {
		return nil, nil, err
	}

	primaryUrl := ofaURLs[0].String()
	switch {
	case strings.Contains(primaryUrl, "flashbots") || strings.Contains(primaryUrl, "novarpc"):
		mc, err := newMultiOfaClient(lggr, chainClient, keyStore, ofaURLs, chainID, txStore, bundles)
		return mc, nil, err
	default:
		mc, err := NewMetaClient(lggr, chainClient, keyStore, ofaURLs[0], chainID, txStore, auctionRequestTimeout)
		if err != nil {
			return nil, nil, err
		}

		if len(ofaURLs) > 1 {
			lggr.Warnw("Created MetaClient for primary OFA URL, ignoring secondary OFA URLs",
				"primaryURL", redactURL(ofaURLs[0]),
				"secondaryURLs", redactURLs(ofaURLs[1:]),
			)
		}

		return mc, NewErrorHandler(), nil
	}
}
