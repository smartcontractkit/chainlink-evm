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

// multiplexPrimary is the authoritative backend: broadcast outcome and all nonce reads.
type multiplexPrimary interface {
	PendingNonceAt(ctx context.Context, address common.Address) (uint64, error)
	NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error
}

// multiplexSecondary is only used for best-effort duplicate sends; multiplex never queries nonces from it.
type multiplexSecondary interface {
	SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error
}

// multiplexClient implements txm.Client: it owns the OFA URL list, constructs one backend per URL,
// fans out sends to secondaries (best-effort), and delegates nonce queries to the primary only.
type multiplexClient struct {
	lggr                 logger.SugaredLogger
	primaryBackend       string
	primary              multiplexPrimary
	secondaries          []multiplexSecondary
	secondarySendTimeout time.Duration
}

var _ txm.Client = (*multiplexClient)(nil)

func backendLabel(c any) string {
	switch x := c.(type) {
	case *ofaTXClient:
		return x.kind.name()
	case *MetaClient:
		return "meta"
	default:
		return fmt.Sprintf("%T", c)
	}
}

// newMultiplexClient wires an already-built primary and optional secondaries. Tests use this;
// production uses newMultiplexClientFromOFAURLs.
func newMultiplexClient(lggr logger.Logger, primaryBackend string, primary multiplexPrimary, secondaries ...multiplexSecondary) *multiplexClient {
	return &multiplexClient{
		lggr:                 logger.Sugared(logger.Named(lggr, "Txm.MultiplexClient")),
		primaryBackend:       primaryBackend,
		primary:              primary,
		secondaries:          secondaries,
		secondarySendTimeout: rpcTimeout,
	}
}

func (m *multiplexClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	for _, secondary := range m.secondaries {
		sec := secondary
		secLabel := backendLabel(sec)
		go func() {
			// Inherit cancellation from the caller so shutdown (ctx done) stops secondary work; timeout caps wall time.
			secondaryCtx, cancel := context.WithTimeout(ctx, m.secondarySendTimeout)
			defer cancel()

			if err := sec.SendTransaction(secondaryCtx, tx, attempt); err != nil {
				m.lggr.Errorw("Secondary backend send failed",
					"err", err,
					"primaryBackend", m.primaryBackend,
					"secondaryBackend", secLabel,
					"transactionLifecycleID", tx.GetTransactionLifecycleID(m.lggr))
			}
		}()
	}

	primaryCtx, cancel := context.WithTimeout(ctx, rpcTimeout)
	defer cancel()
	return m.primary.SendTransaction(primaryCtx, tx, attempt)
}

func (m *multiplexClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return m.primary.PendingNonceAt(ctx, address)
}

func (m *multiplexClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return m.primary.NonceAt(ctx, address, blockNumber)
}

// newMultiplexClientFromOFAURLs builds backends from URLs: index 0 is primary (outcome and nonces); the rest are secondaries.
func newMultiplexClientFromOFAURLs(
	lggr logger.Logger,
	chainClient *clientwrappers.ChainClient,
	keyStore keys.ChainStore,
	ofaURLs []*url.URL,
	chainID *big.Int,
	txStore txm.TxStore,
	bundles *bool,
	auctionRequestTimeout *time.Duration,
) (*multiplexClient, txm.ErrorHandler, error) {
	if len(ofaURLs) == 0 {
		return nil, nil, fmt.Errorf("ofaURLs must not be empty")
	}

	primary, errHandler, err := newClientForOFAURL(lggr, chainClient, keyStore, ofaURLs[0], chainID, txStore, bundles, auctionRequestTimeout)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create primary client for %s: %w", redactURL(ofaURLs[0]), err)
	}

	secondaries := make([]multiplexSecondary, 0, len(ofaURLs)-1)
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

	primaryLabel := backendLabel(primary)
	lggr.Infow("TransactionManagerV2 OFA client created",
		"primaryURL", urlStrs[0],
		"secondaryURLs", urlStrs[1:],
		"primaryBackend", primaryLabel)

	return newMultiplexClient(lggr, primaryLabel, primary, secondaries...), errHandler, nil
}

func newClientForOFAURL(lggr logger.Logger, chainClient *clientwrappers.ChainClient, keyStore keys.ChainStore, u *url.URL, chainID *big.Int, txStore txm.TxStore, bundles *bool, auctionRequestTimeout *time.Duration) (multiplexPrimary, txm.ErrorHandler, error) {
	urlString := u.String()
	switch {
	case strings.Contains(urlString, "flashbots"):
		metrics, err := newOFAMetrics(chainID.String(), ofaKindFlashbots.name())
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create OFA metrics for flashbots: %w", err)
		}
		bundlesEnabled := bundles != nil && *bundles
		return newFlashbotsClient(lggr, chainClient, keyStore, u, txStore, bundlesEnabled, metrics), nil, nil
	case strings.Contains(urlString, "novarpc"):
		metrics, err := newOFAMetrics(chainID.String(), ofaKindNova.name())
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

var (
	_ multiplexPrimary   = (*ofaTXClient)(nil)
	_ multiplexSecondary = (*ofaTXClient)(nil)
	_ multiplexPrimary   = (*MetaClient)(nil)
	_ multiplexSecondary = (*MetaClient)(nil)
)
