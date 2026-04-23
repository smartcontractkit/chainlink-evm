package dualbroadcast

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

// This file holds HTTP JSON-RPC OFA backends: Flashbots-style signing plus shared send/nonce plumbing,
// and Nova. Flashbots-only pieces (signing headers, bundle RPC) live here rather than in the multiplex client.

// rpcTimeout bounds OFA HTTP JSON-RPC calls. The multiplex client reuses this for secondary fan-out deadlines.
const rpcTimeout = 10 * time.Second

// ofaKind selects URL shape, signing headers, logger name, non-dual fallback, and bundle behavior.
type ofaKind uint8

const (
	ofaKindFlashbots ofaKind = iota
	ofaKindNova
)

func (k ofaKind) name() string {
	switch k {
	case ofaKindFlashbots:
		return "flashbots"
	case ofaKindNova:
		return "nova"
	default:
		return "unknown"
	}
}

// chainRPCClient is used for public mempool fallback and reads shared by OFA implementations.
type chainRPCClient interface {
	BlockByNumber(ctx context.Context, number *big.Int) (*evmtypes.Block, error)
	NonceAt(context.Context, common.Address, *big.Int) (uint64, error)
	SendTransaction(context.Context, *types.Transaction, *types.Attempt) error
}

type ofaPostResponse struct {
	Result json.RawMessage `json:"result,omitempty"`
	Error  ofaPostError
}

type ofaPostError struct {
	Message string `json:"message,omitempty"`
}

// FlashbotsTxStore is required when bundle sending is enabled for Flashbots.
type FlashbotsTxStore interface {
	FetchUnconfirmedTransactions(context.Context, common.Address) ([]*types.Transaction, error)
}

// ofaTXClient is an HTTP JSON-RPC OFA backend (Flashbots MEV-Share, Nova RPC). Multiplex uses it as a
// multiplexPrimary (when authoritative) or multiplexSecondary (best-effort duplicate send only).
type ofaTXClient struct {
	lggr      logger.SugaredLogger
	c         chainRPCClient
	customURL *url.URL
	kind      ofaKind
	keystore  keys.MessageSigner // Only if authentication is required
	metrics   ofaMetrics
	txStore   FlashbotsTxStore
	bundles   bool
}

func newFlashbotsClient(lggr logger.Logger, c chainRPCClient, keystore keys.MessageSigner, customURL *url.URL, txStore FlashbotsTxStore, bundlesEnabled bool, metrics ofaMetrics) *ofaTXClient {
	return &ofaTXClient{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.FlashbotsClient")),
		c:         c,
		customURL: customURL,
		kind:      ofaKindFlashbots,
		keystore:  keystore,
		metrics:   metrics,
		txStore:   txStore,
		bundles:   bundlesEnabled,
	}
}

func newNovaClient(lggr logger.Logger, c chainRPCClient, customURL *url.URL, metrics ofaMetrics) *ofaTXClient {
	return &ofaTXClient{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.NovaClient")),
		c:         c,
		customURL: customURL,
		kind:      ofaKindNova,
		metrics:   metrics,
	}
}

func (d *ofaTXClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return d.c.NonceAt(ctx, address, blockNumber)
}

func (d *ofaTXClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["%s","pending"], "id":1}`, address.Hex()))
	raw, err := d.postJSONRPC(ctx, address, body, nil)
	if err != nil {
		return 0, fmt.Errorf("%s eth_getTransactionCount failed: %w", d.kind.name(), err)
	}

	var resultStr string
	if err = json.Unmarshal(raw, &resultStr); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response %s into string: %w", string(raw), err)
	}
	nonce, err := hexutil.DecodeUint64(resultStr)
	if err != nil {
		return 0, fmt.Errorf("failed to decode response %v into uint64: %w", resultStr, err)
	}
	return nonce, nil
}

func (d *ofaTXClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	if meta == nil || meta.DualBroadcast == nil || !*meta.DualBroadcast || tx.IsPurgeable {
		switch d.kind {
		case ofaKindFlashbots:
			// If not dual-broadcast, fall back to sending the transaction to the chain RPC directly
			return d.c.SendTransaction(ctx, nil, attempt)
		case ofaKindNova:
			return nil // assume we only use Nova for secondary broadcast, don't fall back to chain RPC
		default:
			return fmt.Errorf("ofaTXClient: unsupported OFA backend %q for dual-broadcast routing", d.kind.name())
		}
	}

	if err := d.sendDualBroadcastTx(ctx, tx, attempt, meta); err != nil {
		return err
	}

	if d.kind == ofaKindFlashbots && d.bundles {
		if err := d.sendBundle(ctx, tx.FromAddress, meta); err != nil {
			d.lggr.Errorw("error sending bundle", "err", err, "transactionLifecycleID", tx.GetTransactionLifecycleID(d.lggr))
		}
	}

	return nil
}

func (d *ofaTXClient) sendDualBroadcastTx(ctx context.Context, tx *types.Transaction, attempt *types.Attempt, meta *types.TxMeta) error {
	data, err := attempt.SignedTransaction.MarshalBinary()
	if err != nil {
		return err
	}

	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["%s"], "id":1}`, hexutil.Encode(data)))
	start := time.Now()
	_, err = d.postJSONRPC(ctx, tx.FromAddress, body, meta)
	d.metrics.RecordSendTx(ctx, time.Since(start), err)

	return err
}

func (d *ofaTXClient) postJSONRPC(ctx context.Context, from common.Address, body []byte, meta *types.TxMeta) (json.RawMessage, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, rpcTimeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.postURL(meta), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if err = d.signRequest(ctx, req, body, from); err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s request failed: %w", d.kind.name(), err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s request failed with status %d: %s", d.kind.name(), resp.StatusCode, string(respBody))
	}

	var response ofaPostResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal OFA response: %w: %s", err, string(respBody))
	}
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response.Result, nil
}

func (d *ofaTXClient) signRequest(ctx context.Context, req *http.Request, body []byte, from common.Address) error {
	if d.kind != ofaKindFlashbots || d.keystore == nil {
		// signing is only required for Flashbots
		return nil
	}

	hashedBody := crypto.Keccak256Hash(body).Hex()
	signedMessage, err := d.keystore.SignMessage(ctx, from, []byte(hashedBody))
	if err != nil {
		return err
	}
	req.Header.Add("X-Flashbots-signature", from.String()+":"+hexutil.Encode(signedMessage))
	req.Header.Add("X-Flashbots-Origin", "chainlink")
	return nil
}

func (d *ofaTXClient) postURL(meta *types.TxMeta) string {
	// Only Flashbots needs URL parameters
	if d.kind != ofaKindFlashbots {
		return d.customURL.String()
	}

	var params string
	if meta != nil && meta.DualBroadcastParams != nil {
		params = *meta.DualBroadcastParams
	}
	return d.customURL.String() + "?" + params
}
