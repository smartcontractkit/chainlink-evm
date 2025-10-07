package dualbroadcast

import (
	"math/big"
	"net/url"
	"strings"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
)

func SelectClient(lggr logger.Logger, client client.Client, keyStore keys.ChainStore, url *url.URL, chainID *big.Int) (txm.Client, error) {
	urlString := url.String()
	switch {
	case strings.Contains(urlString, "flashbots"):
		return NewFlashbotsClient(client, keyStore, url), nil
	default:
		return NewMetaClient(lggr, client, keyStore, url, chainID)
	}
}
