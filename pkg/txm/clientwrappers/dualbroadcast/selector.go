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
)

func SelectClient(lggr logger.Logger, client client.Client, keyStore keys.ChainStore, url *url.URL, chainID *big.Int, txStore txm.TxStore, bundles *bool, auctionRequestTimeout *time.Duration) (txm.Client, txm.ErrorHandler, error) {
	urlString := url.String()
	switch {
	case strings.Contains(urlString, "flashbots"):
		return NewFlashbotsClient(lggr, client, keyStore, url, txStore, bundles), nil, nil
	default:
		mc, err := NewMetaClient(lggr, client, keyStore, url, chainID, txStore, auctionRequestTimeout)
		if err != nil {
			return nil, nil, err
		}
		return mc, NewErrorHandler(), nil
	}
}
