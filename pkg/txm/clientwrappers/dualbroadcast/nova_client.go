package dualbroadcast

import (
	"context"
	"fmt"
	"math/big"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

type novaClient struct {
	lggr       logger.SugaredLogger
	ofaClient  *ofaClient
	keystore   keys.TxSigner
	tier2Feeds map[common.Address]struct{}
}

var _ txm.Client = (*novaClient)(nil)

func newNovaClient(lggr logger.Logger, c ofaRPCClient, customURL *url.URL, metrics ofaMetrics, keystore keys.TxSigner, tier2Feeds []common.Address) *novaClient {
	log := logger.Sugared(logger.Named(lggr, "Txm.NovaClient"))
	ofaClient := newOFAClient(c, customURL, noAuth{}, metrics, "nova")

	feedSet := make(map[common.Address]struct{}, len(tier2Feeds))
	for _, addr := range tier2Feeds {
		feedSet[addr] = struct{}{}
	}

	return &novaClient{
		lggr:       log,
		ofaClient:  ofaClient,
		keystore:   keystore,
		tier2Feeds: feedSet,
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

	if meta == nil || meta.DualBroadcast == nil || !*meta.DualBroadcast || tx.IsPurgeable {
		return nil
	}

	feedAddr := tx.ToAddress
	if meta.FwdrDestAddress != nil { // Aggregator address
		feedAddr = *meta.FwdrDestAddress
	}

	signedTx, err := n.resignWithTierValue(ctx, tx, attempt, feedAddr)
	if err != nil {
		return fmt.Errorf("failed to re-sign transaction for Nova: %w", err)
	}

	novaAttempt := *attempt
	novaAttempt.SignedTransaction = signedTx
	novaAttempt.Hash = signedTx.Hash()
	return n.ofaClient.sendDualBroadcastTx(ctx, tx, &novaAttempt)
}

func (n *novaClient) tierValue(toAddress common.Address) *big.Int {
	if _, ok := n.tier2Feeds[toAddress]; ok {
		return big.NewInt(2)
	}
	return big.NewInt(1)
}

func (n *novaClient) resignWithTierValue(ctx context.Context, tx *types.Transaction, attempt *types.Attempt, feedAddr common.Address) (*evmtypes.Transaction, error) {
	origTx := attempt.SignedTransaction
	value := n.tierValue(feedAddr)

	var inner evmtypes.TxData
	switch origTx.Type() {
	case evmtypes.LegacyTxType:
		inner = &evmtypes.LegacyTx{
			Nonce:    origTx.Nonce(),
			To:       origTx.To(),
			Value:    value,
			Gas:      origTx.Gas(),
			GasPrice: origTx.GasPrice(),
			Data:     origTx.Data(),
		}
	case evmtypes.DynamicFeeTxType:
		inner = &evmtypes.DynamicFeeTx{
			ChainID:    origTx.ChainId(),
			Nonce:      origTx.Nonce(),
			To:         origTx.To(),
			Value:      value,
			Gas:        origTx.Gas(),
			GasFeeCap:  origTx.GasFeeCap(),
			GasTipCap:  origTx.GasTipCap(),
			Data:       origTx.Data(),
			AccessList: origTx.AccessList(),
		}
	default:
		return nil, fmt.Errorf("unsupported transaction type for re-signing: %d", origTx.Type())
	}

	return n.keystore.SignTx(ctx, tx.FromAddress, evmtypes.NewTx(inner))
}
