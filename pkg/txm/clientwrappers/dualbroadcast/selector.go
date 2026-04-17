package dualbroadcast

import (
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

func SelectClient(
	lggr logger.Logger,
	client client.Client,
	keyStore keys.ChainStore,
	url *url.URL,
	chainID *big.Int,
	txStore txm.TxStore,
	readRequestsToMultipleNodes bool,
	bundles *bool,
	auctionRequestTimeout *time.Duration,
	lifecycleMetrics txm.Metrics,
) (txm.Client, txm.ErrorHandler, error) {
	chainClient, err := clientwrappers.NewChainClient(lggr, client, readRequestsToMultipleNodes)
	if err != nil {
		return nil, nil, err
	}

	urlString := url.String()
	switch {
	case strings.Contains(urlString, "flashbots"):
		return NewFlashbotsClient(lggr, chainClient, keyStore, url, txStore, bundles), nil, nil
	default:
		mc, err := NewMetaClient(lggr, chainClient, keyStore, url, chainID, txStore, auctionRequestTimeout, lifecycleMetrics)
		if err != nil {
			return nil, nil, err
		}
		return mc, NewErrorHandler(), nil
	}
}
