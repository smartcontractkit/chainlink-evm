package dualbroadcast

import (
	"context"
	"math/big"
	"net/url"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

type novaClient struct {
	lggr      logger.SugaredLogger
	ofaClient *ofaClient
}

var _ txm.Client = (*novaClient)(nil)

func newNovaClient(lggr logger.Logger, c publicMempoolRPC, customURL *url.URL, metrics ofaMetrics) *novaClient {
	log := logger.Sugared(logger.Named(lggr, "Txm.NovaClient"))
	ofaClient := newOFAClient(c, customURL, noAuth{}, metrics, "nova")

	return &novaClient{
		lggr:      log,
		ofaClient: ofaClient,
	}
}

func (n *novaClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return n.ofaClient.NonceAt(ctx, address, blockNumber)
}

func (n *novaClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return n.ofaClient.PendingNonceAt(ctx, address)
}

func (n *novaClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	// If not dual-broadcast, don't do anything
	if meta == nil || meta.DualBroadcast == nil || !*meta.DualBroadcast || tx.IsPurgeable {
		return nil
	}

	_, err2 := n.ofaClient.sendDualBroadcastTx(ctx, tx, attempt)
	return err2
}
