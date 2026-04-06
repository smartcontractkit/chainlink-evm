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

func SelectClient(lggr logger.Logger, client client.Client, keyStore keys.ChainStore, primaryURL *url.URL, secondaryURL *url.URL, chainID *big.Int, txStore txm.TxStore, readRequestsToMultipleNodes bool, bundles *bool, auctionRequestTimeout *time.Duration) (txm.Client, txm.ErrorHandler, error) {
	chainClient, err := clientwrappers.NewChainClient(lggr, client, readRequestsToMultipleNodes)
	if err != nil {
		return nil, nil, err
	}

	primary, errHandler, err := selectSingleClient(lggr, chainClient, keyStore, primaryURL, chainID, txStore, bundles, auctionRequestTimeout)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create primary client for %s: %w", primaryURL.Redacted(), err)
	}

	if secondaryURL == nil {
		return primary, errHandler, nil
	}

	// TODO(gg): hardcode this to nova instead to prevent accidental use of flashbots/meta as secondary
	secondary, _, err := selectSingleClient(lggr, chainClient, keyStore, secondaryURL, chainID, txStore, bundles, auctionRequestTimeout)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create secondary client for %s: %w", secondaryURL.Redacted(), err)
	}

	return newMultiplexClient(lggr, primary, secondary), errHandler, nil
}

func selectSingleClient(lggr logger.Logger, chainClient *clientwrappers.ChainClient, keyStore keys.ChainStore, u *url.URL, chainID *big.Int, txStore txm.TxStore, bundles *bool, auctionRequestTimeout *time.Duration) (txm.Client, txm.ErrorHandler, error) {
	urlString := u.String()
	switch {
	case strings.Contains(urlString, "flashbots"):
		metrics, err := NewOFAMetrics(chainID.String(), "flashbots")
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create OFA metrics for flashbots: %w", err)
		}
		return NewFlashbotsClient(lggr, chainClient, keyStore, u, txStore, bundles, metrics), nil, nil
	case strings.Contains(urlString, "novarpc"):
		metrics, err := NewOFAMetrics(chainID.String(), "nova")
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
