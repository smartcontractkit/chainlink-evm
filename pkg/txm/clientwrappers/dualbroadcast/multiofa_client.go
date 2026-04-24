package dualbroadcast

import (
	"context"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

// ofaBackend is used to broadcast to an OFA and read nonces.
type ofaBackend interface {
	PendingNonceAt(ctx context.Context, address common.Address) (uint64, error)
	NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error
	Label() string
}

// multiOfaClient implements txm.Client: it owns the OFA URL list, constructs one backend per URL,
// fans out sends to secondaries (best-effort), and delegates nonce queries to the primary only.
type multiOfaClient struct {
	lggr                 logger.SugaredLogger
	chainClient          chainRPCClient
	primary              ofaBackend
	secondaries          []ofaBackend
	secondarySendTimeout time.Duration
}

var (
	_ txm.Client = (*multiOfaClient)(nil)
	_ ofaBackend = (*ofaTXClient)(nil)
	_ ofaBackend = (*MetaClient)(nil)
)

// newMultiOfaClient builds backends from URLs: index 0 is primary (outcome and nonces); the rest are secondaries.
func newMultiOfaClient(
	lggr logger.Logger,
	chainClient *clientwrappers.ChainClient,
	keyStore keys.ChainStore,
	ofaURLs []*url.URL,
	chainID *big.Int,
	txStore txm.TxStore,
	bundles *bool,
	auctionRequestTimeout *time.Duration,
) (*multiOfaClient, txm.ErrorHandler, error) {
	if len(ofaURLs) == 0 {
		return nil, nil, fmt.Errorf("ofaURLs must not be empty")
	}

	primary, errHandler, err := newClientForOFAURL(lggr, chainClient, keyStore, ofaURLs[0], chainID, txStore, bundles, auctionRequestTimeout)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create primary client for %s: %w", redactURL(ofaURLs[0]), err)
	}

	secondaries := make([]ofaBackend, 0, len(ofaURLs)-1)
	for _, u := range ofaURLs[1:] {
		sec, _, err := newClientForOFAURL(lggr, chainClient, keyStore, u, chainID, txStore, bundles, auctionRequestTimeout)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create secondary client for %s: %w", redactURL(u), err)
		}
		secondaries = append(secondaries, sec)
	}

	urlStrs := make([]string, len(ofaURLs))
	for i, u := range ofaURLs {
		urlStrs[i] = redactURL(u)
	}

	secondaryLabels := make([]string, len(secondaries))
	for i, secondary := range secondaries {
		secondaryLabels[i] = secondary.Label()
	}

	lggr.Infow("TransactionManagerV2 OFA client created",
		"primaryBackend", primary.Label(),
		"primaryURL", urlStrs[0],
		"secondaryBackends", secondaryLabels,
		"secondaryURLs", urlStrs[1:])

	return &multiOfaClient{
		lggr:                 logger.Sugared(logger.Named(lggr, "Txm.MultiOfaClient")),
		chainClient:          chainClient,
		primary:              primary,
		secondaries:          secondaries,
		secondarySendTimeout: rpcTimeout,
	}, errHandler, nil
}

// SendTransaction sends the transaction to the primary and secondaries, unless it is a non-dual-broadcast transaction.
// In that case, it falls back to sending the transaction to the chain RPC directly.
func (m *multiOfaClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	// TODO(gg): check that this is the correct behavior for meta_client as well
	// If not dual-broadcast, fall back to sending the transaction to the chain RPC directly
	if meta == nil || meta.DualBroadcast == nil || !*meta.DualBroadcast || tx.IsPurgeable {
		return m.chainClient.SendTransaction(ctx, tx, attempt)
	}

	// send to secondaries in parallel, fire-and-forget: in case of error, log and continue
	for _, secondary := range m.secondaries {
		sec := secondary
		go func() {

			// TODO(gg): add waitgroup, see this comment: https://github.com/smartcontractkit/chainlink-evm/pull/410#discussion_r3110773136

			secondaryCtx, cancel := context.WithTimeout(ctx, m.secondarySendTimeout)
			defer cancel()

			if err := sec.SendTransaction(secondaryCtx, tx, attempt); err != nil {
				m.lggr.Errorw("Secondary backend send failed",
					"err", err,
					"backend", sec.Label(),
					"txID", tx.ID,
					"attemptHash", attempt.Hash,
					"transactionLifecycleID", tx.GetTransactionLifecycleID(m.lggr))
			}
		}()
	}

	primaryCtx, cancel := context.WithTimeout(ctx, rpcTimeout)
	defer cancel()
	return m.primary.SendTransaction(primaryCtx, tx, attempt)
}

func (m *multiOfaClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return m.primary.PendingNonceAt(ctx, address)
}

func (m *multiOfaClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return m.primary.NonceAt(ctx, address, blockNumber)
}

func newClientForOFAURL(
	lggr logger.Logger,
	chainClient *clientwrappers.ChainClient,
	keyStore keys.ChainStore,
	u *url.URL,
	chainID *big.Int,
	txStore txm.TxStore,
	bundles *bool,
	auctionRequestTimeout *time.Duration) (ofaBackend, txm.ErrorHandler, error) {

	urlString := u.String()
	switch {
	case strings.Contains(urlString, "flashbots"):
		metrics, err := newOFAMetrics(chainID.String(), ofaFlashbots.name())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create OFA metrics for flashbots: %w", err)
		}
		bundlesEnabled := bundles != nil && *bundles
		return newFlashbotsClient(lggr, chainClient, keyStore, u, txStore, bundlesEnabled, metrics), nil, nil
	case strings.Contains(urlString, "novarpc"):
		metrics, err := newOFAMetrics(chainID.String(), ofaNova.name())
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
