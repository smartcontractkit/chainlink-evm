package dualbroadcast

import (
	"context"
	"errors"
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

// multiOfaBackend is used to broadcast to an OFA backend and read nonces.
type multiOfaBackend interface {
	PendingNonceAt(ctx context.Context, address common.Address) (uint64, error)
	NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error
}

// multiOfaClient implements txm.Client: it owns the OFA URL list, constructs one backend per URL,
// fans out sends to secondaries (best-effort), and delegates nonce queries to the primary only.
type multiOfaClient struct {
	lggr                 logger.SugaredLogger
	chainClient          chainRPCClient
	primary              multiOfaBackend
	secondaries          []multiOfaBackend
	secondarySendTimeout time.Duration // used in unit tests to configure the timeout
}

var _ txm.Client = (*multiOfaClient)(nil)

// newMultiOfaClient builds backends from URLs: index 0 is primary (outcome and nonces); the rest are secondaries.
func newMultiOfaClient(
	lggr logger.Logger,
	chainClient *clientwrappers.ChainClient,
	keyStore keys.ChainStore,
	ofaURLs []*url.URL,
	chainID *big.Int,
	txStore txm.TxStore,
	bundles *bool,
) (*multiOfaClient, error) {
	if len(ofaURLs) == 0 {
		return nil, errors.New("ofaURLs must not be empty")
	}

	primary, err := newClientForRelayOFAURL(lggr, chainClient, keyStore, ofaURLs[0], chainID, txStore, bundles)
	if err != nil {
		return nil, fmt.Errorf("failed to create primary client for %s: %w", redactURL(ofaURLs[0]), err)
	}

	secondaries := make([]multiOfaBackend, 0, len(ofaURLs)-1)
	for _, u := range ofaURLs[1:] {
		sec, err := newClientForRelayOFAURL(lggr, chainClient, keyStore, u, chainID, txStore, bundles)
		if err != nil {
			return nil, fmt.Errorf("failed to create secondary client for %s: %w", redactURL(u), err)
		}
		secondaries = append(secondaries, sec)
	}

	//nolint:gosec // disable G602 // using ofaURLs is safe here as we validated that it is not empty above
	lggr.Infow("MultiOfaClient created",
		"primaryURL", redactURL(ofaURLs[0]), "secondaryURLs", redactURLs(ofaURLs[1:]),
	)

	return &multiOfaClient{
		lggr:                 logger.Sugared(logger.Named(lggr, "Txm.MultiOfaClient")),
		chainClient:          chainClient,
		primary:              primary,
		secondaries:          secondaries,
		secondarySendTimeout: rpcTimeout,
	}, nil
}

// SendTransaction sends the transaction to the primary and secondaries, unless it is a non-dual-broadcast transaction.
// In that case, it falls back to sending the transaction to the chain RPC directly.
func (m *multiOfaClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	// If not dual-broadcast, fall back to sending the transaction to the chain RPC directly.
	if meta == nil || meta.DualBroadcast == nil || !*meta.DualBroadcast || tx.IsPurgeable {
		return m.chainClient.SendTransaction(ctx, tx, attempt)
	}

	// Secondaries are best-effort and do not block the primary. Each call is bounded by
	// secondarySendTimeout
	for _, secondary := range m.secondaries {
		sec := secondary
		go func() {
			secondaryCtx, cancel := context.WithTimeout(ctx, m.secondarySendTimeout)
			defer cancel()

			if secErr := sec.SendTransaction(secondaryCtx, tx, attempt); secErr != nil {
				m.lggr.Errorw("Secondary backend send failed",
					"err", secErr,
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

// newClientForRelayOFAURL builds only Flashbots or Nova backends for use inside multiOfaClient.
func newClientForRelayOFAURL(
	lggr logger.Logger,
	chainClient *clientwrappers.ChainClient,
	keyStore keys.ChainStore,
	u *url.URL,
	chainID *big.Int,
	txStore txm.TxStore,
	bundles *bool) (multiOfaBackend, error) {
	urlString := u.String()
	switch {
	case strings.Contains(urlString, "flashbots"):
		metrics, err := newOFAMetrics(chainID.String(), ofaFlashbots.name())
		if err != nil {
			return nil, fmt.Errorf("failed to create OFA metrics for flashbots: %w", err)
		}
		bundlesEnabled := bundles != nil && *bundles
		return newFlashbotsClient(lggr, chainClient, keyStore, u, txStore, bundlesEnabled, metrics), nil
	case strings.Contains(urlString, "novarpc"):
		metrics, err := newOFAMetrics(chainID.String(), ofaNova.name())
		if err != nil {
			return nil, fmt.Errorf("failed to create OFA metrics for nova: %w", err)
		}
		return newNovaClient(lggr, chainClient, u, metrics), nil
	default:
		return nil, fmt.Errorf("multiOfaClient does not support OFA URL %s", redactURL(u))
	}
}

// redactURL returns u as a string safe for logs: same redaction as url.URL.Redacted for userinfo, and api_key query values are replaced with "xxxxx". It does not mutate the original URL.
func redactURL(u *url.URL) string {
	if u == nil {
		return ""
	}
	cp := *u
	q := cp.Query()
	if _, has := q["api_key"]; has {
		q.Set("api_key", "xxxxx")
		cp.RawQuery = q.Encode()
	}
	return cp.Redacted()
}

func redactURLs(urls []*url.URL) []string {
	redacted := make([]string, len(urls))
	for i, u := range urls {
		redacted[i] = redactURL(u)
	}
	return redacted
}
