package dualbroadcast

import (
	"errors"
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
func SelectClient(lggr logger.Logger, client client.Client, keyStore keys.ChainStore, ofaURLs []*url.URL, chainID *big.Int, txStore txm.TxStore, readRequestsToMultipleNodes bool, bundles *bool, auctionRequestTimeout *time.Duration, lifecycleMetrics txm.Metrics) (txm.Client, txm.ErrorHandler, error) {
	if len(ofaURLs) == 0 {
		return nil, nil, errors.New("ofaURLs must not be empty")
	}

	chainClient, err := clientwrappers.NewChainClient(lggr, client, readRequestsToMultipleNodes)
	if err != nil {
		return nil, nil, err
	}

	var mc txm.Client
	// for fastlane atlas we support single URL only
	if len(ofaURLs) == 1 && strings.Contains(ofaURLs[0].String(), "fastlane") {
		mc, err = NewMetaClient(lggr, chainClient, keyStore, ofaURLs[0], chainID, txStore, auctionRequestTimeout, lifecycleMetrics)
		if err != nil {
			return nil, nil, err
		}

		return mc, NewErrorHandler(), nil
	}

	// use multiOfaClient for all other cases
	mc, err = newMultiOfaClient(lggr, chainClient, keyStore, ofaURLs, chainID, txStore, bundles)
	if err != nil {
		return nil, nil, err
	}
	return mc, nil, err
}
