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
	"strconv"
	"strings"
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
// and Nova. Flashbots-only pieces (signing headers, bundle RPC) live here rather than in the multiOfaClient.

// rpcTimeout bounds OFA HTTP JSON-RPC calls. The multiOfaClient reuses this for secondary fan-out deadlines.
const rpcTimeout = 10 * time.Second

// ofa selects URL shape, signing headers, logger name, and bundle behavior.
type ofa uint8

// ofa represents the type of OFA backend: Flashbots or Nova. Meta (Fastlane Atlas) is separate from this at the moment.
const (
	ofaFlashbots ofa = iota
	ofaNova
)

func (k ofa) name() string {
	switch k {
	case ofaFlashbots:
		return "flashbots"
	case ofaNova:
		return "nova"
	default:
		return "unknown"
	}
}

// ofaBackendRPCClient is used for public mempool fallback and reads shared by OFA implementations.
type ofaBackendRPCClient interface {
	BlockByNumber(ctx context.Context, number *big.Int) (*evmtypes.Block, error)
	NonceAt(context.Context, common.Address, *big.Int) (uint64, error)
	SendTransaction(context.Context, *types.Transaction, *types.Attempt) error
}

type postResponse struct {
	Result json.RawMessage `json:"result,omitempty"`
	Error  postError
}

type postError struct {
	Message string `json:"message,omitempty"`
}

// FlashbotsTxStore is required when bundle sending is enabled for Flashbots.
type FlashbotsTxStore interface {
	FetchUnconfirmedTransactions(context.Context, common.Address) ([]*types.Transaction, error)
}

// ofaBackend is an HTTP JSON-RPC OFA backend (Flashbots MEV-Share, Nova RPC)
type ofaBackend struct {
	lggr      logger.SugaredLogger
	c         ofaBackendRPCClient
	customURL *url.URL
	ofa       ofa
	keystore  keys.MessageSigner // Only if authentication is required
	metrics   ofaMetrics
	txStore   FlashbotsTxStore
	bundles   bool
}

var _ multiOfaBackend = (*ofaBackend)(nil)

func newFlashbotsClient(lggr logger.Logger, c ofaBackendRPCClient, keystore keys.MessageSigner, customURL *url.URL, txStore FlashbotsTxStore, bundlesEnabled bool, metrics ofaMetrics) *ofaBackend {
	return &ofaBackend{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.FlashbotsClient")),
		c:         c,
		customURL: customURL,
		ofa:       ofaFlashbots,
		keystore:  keystore,
		metrics:   metrics,
		txStore:   txStore,
		bundles:   bundlesEnabled,
	}
}

func newNovaClient(lggr logger.Logger, c ofaBackendRPCClient, customURL *url.URL, metrics ofaMetrics) *ofaBackend {
	return &ofaBackend{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.NovaClient")),
		c:         c,
		customURL: customURL,
		ofa:       ofaNova,
		metrics:   metrics,
	}
}

func (d *ofaBackend) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return d.c.NonceAt(ctx, address, blockNumber)
}

func (d *ofaBackend) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["%s","pending"], "id":1}`, address.Hex()))
	raw, err := d.postJSONRPC(ctx, address, body, nil)
	if err != nil {
		return 0, fmt.Errorf("%s eth_getTransactionCount failed: %w", d.ofa.name(), err)
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

func (d *ofaBackend) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	if meta == nil || meta.DualBroadcast == nil || !*meta.DualBroadcast || tx.IsPurgeable {
		return errors.New("ofaTXClient: SendTransaction called for a non-dual-broadcast transaction")
	}

	if err := d.sendDualBroadcastTx(ctx, tx, attempt, meta); err != nil {
		return err
	}

	if d.ofa == ofaFlashbots && d.bundles {
		if err := d.sendBundle(ctx, tx.FromAddress, meta); err != nil {
			d.lggr.Errorw("error sending bundle", "err", err, "transactionLifecycleID", tx.GetTransactionLifecycleID(d.lggr))
		}
	}

	return nil
}

func (d *ofaBackend) Label() string {
	return d.ofa.name()
}

func (d *ofaBackend) sendDualBroadcastTx(ctx context.Context, tx *types.Transaction, attempt *types.Attempt, meta *types.TxMeta) error {
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

func (d *ofaBackend) postJSONRPC(ctx context.Context, from common.Address, body []byte, meta *types.TxMeta) (json.RawMessage, error) {
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
		return nil, fmt.Errorf("%s request failed: %w", d.ofa.name(), err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s request failed with status %d: %s", d.ofa.name(), resp.StatusCode, string(respBody))
	}

	var response postResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal OFA response: %w: %s", err, string(respBody))
	}
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response.Result, nil
}

func (d *ofaBackend) signRequest(ctx context.Context, req *http.Request, body []byte, from common.Address) error {
	if d.ofa != ofaFlashbots || d.keystore == nil {
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

func (d *ofaBackend) postURL(meta *types.TxMeta) string {
	// Only Flashbots needs URL parameters
	if d.ofa != ofaFlashbots {
		return d.customURL.String()
	}

	var params string
	if meta != nil && meta.DualBroadcastParams != nil {
		params = *meta.DualBroadcastParams
	}
	return d.customURL.String() + "?" + params
}

// ----------------------------------------------
// Flashbots bundle submission (mev_sendBundle)
// ----------------------------------------------

const maxBlockDiff = 24

// sendBundle sends a bundle of all the in-flight transactions.
func (d *ofaBackend) sendBundle(ctx context.Context, fromAddress common.Address, meta *types.TxMeta) error {
	var urlParams string
	if meta != nil && meta.DualBroadcastParams != nil {
		urlParams = *meta.DualBroadcastParams
	}

	unconfirmedTxs, err := d.txStore.FetchUnconfirmedTransactions(ctx, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to fetch unconfirmed transactions: %w", err)
	}

	// We fetch all the unconfirmed transactions in an ascending nonce order.
	// For the bundle we need a signed transaction so we get the last attempt from each transaction.
	// TODO: Implement a more sophisticated attempt selection logic if necessary.
	attempts := make([]*types.Attempt, 0, len(unconfirmedTxs))
	attemptIDs := make([]uint64, 0, len(unconfirmedTxs))
	nonces := make([]uint64, 0, len(unconfirmedTxs))
	ids := make([]uint64, 0, len(unconfirmedTxs))
	for _, unconfirmedTx := range unconfirmedTxs {
		if len(unconfirmedTx.Attempts) > 0 && unconfirmedTx.Nonce != nil && unconfirmedTx.Attempts[len(unconfirmedTx.Attempts)-1].SignedTransaction != nil {
			latestAttempt := unconfirmedTx.Attempts[len(unconfirmedTx.Attempts)-1]
			attempts = append(attempts, latestAttempt)
			attemptIDs = append(attemptIDs, latestAttempt.ID)
			ids = append(ids, unconfirmedTx.ID)
		}
	}

	// Need at least 2 transactions to send a bundle
	if len(attempts) < 2 {
		return nil
	}

	prevNonce := attempts[0].SignedTransaction.Nonce()
	nonces = append(nonces, prevNonce)
	for _, attempt := range attempts[1:] {
		nonce := attempt.SignedTransaction.Nonce()
		nonces = append(nonces, nonce)
		expectedNonce := prevNonce + 1
		if nonce != expectedNonce {
			return fmt.Errorf("bundle attempts must be contiguous and strictly increasing: expected nonce %d, got nonce %d", expectedNonce, nonce)
		}
		prevNonce = nonce
	}

	// We make an RPC call to get the current block height. Some async caching may help with the overhead.
	currentBlock, err := d.c.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get current block height: %w", err)
	}
	maxBlock := currentBlock.NumberU64() + maxBlockDiff

	// For reference, Flashbots Bundle definition can be found here: https://docs.flashbots.net/flashbots-mev-share/searchers/understanding-bundles#bundle-definition
	// Keep in mind the docs might be outdated and latest features might not be documented.
	bodyItems := make([]map[string]any, 0, len(attempts))
	for _, attempt := range attempts {
		txData, err := attempt.SignedTransaction.MarshalBinary()
		if err != nil {
			return fmt.Errorf("failed to marshal transaction for attempt ID %d: %w", attempt.ID, err)
		}

		bodyItems = append(bodyItems, map[string]any{
			"tx":         hexutil.Encode(txData),
			"revertMode": "allow", // we always want to allow reverts so bundles are valid even if a single transaction within the bundle goes through
		})
	}
	privacy, refundConfig, err := parseURLParams(urlParams)
	if err != nil {
		return err
	}

	bundleParams := map[string]any{
		"body": bodyItems,
		"inclusion": map[string]any{
			"block":    hexutil.EncodeBig(new(big.Int).SetUint64(currentBlock.NumberU64())),
			"maxBlock": hexutil.EncodeBig(new(big.Int).SetUint64(maxBlock)),
		},
		"privacy": privacy,
		"version": "v0.1",
	}
	if refundConfig.Address != "" {
		bundleParams["validity"] = map[string]any{
			"refundConfig": []any{refundConfig},
		}
	}

	requestBody := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "mev_sendBundle",
		"params":  []any{bundleParams},
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal bundle request: %w", err)
	}

	raw, err := d.postJSONRPC(ctx, fromAddress, bodyBytes, nil)
	if err != nil {
		return err
	}

	var bundleResult struct {
		BundleHash string `json:"bundleHash"`
	}
	if err := json.Unmarshal(raw, &bundleResult); err != nil {
		return fmt.Errorf("failed to decode response %s into bundle result: %w", string(raw), err)
	}
	d.lggr.Infow("Broadcasted transaction bundle", "txIDs", ids, "attemptIDs", attemptIDs, "nonces", nonces, "bundleHash", bundleResult.BundleHash)
	return nil
}

func parseURLParams(params string) (privacy, refundConfig, error) {
	values, err := url.ParseQuery(params)
	if err != nil {
		return privacy{}, refundConfig{}, fmt.Errorf("unable to parse params: %w", err)
	}

	pvc := privacy{}
	if timeout, err := strconv.Atoi(values.Get("auctionTimeout")); err == nil {
		pvc.AuctionTimeout = timeout
	}

	pvc.Builders = append(pvc.Builders, values["builder"]...)

	pvc.Hints = append(pvc.Hints, values["hint"]...)

	refundCfg := refundConfig{}
	refundRaw := values.Get("refund")
	if refundRaw != "" {
		parts := strings.Split(refundRaw, ":")
		if len(parts) != 2 {
			return privacy{}, refundConfig{}, fmt.Errorf("unable to parse refund: %s. Expected format: address:percent", refundRaw)
		}
		address := parts[0]
		percentVal, err := strconv.Atoi(parts[1])
		if err != nil {
			return privacy{}, refundConfig{}, fmt.Errorf("unable to parse percentage: %w", err)
		}

		pvc.WantRefund = percentVal
		refundCfg = refundConfig{
			Address: address,
			Percent: 100, // wantRefund is an absolute percent of the refund, and refundConfig.percent=100 means entire refund goes to this address (no longer supported)
		}
	}
	return pvc, refundCfg, nil
}

type privacy struct {
	WantRefund     int      `json:"wantRefund"`
	AuctionTimeout int      `json:"auctionTimeout"`
	Builders       []string `json:"builders"`
	Hints          []string `json:"hints"`
}

type refundConfig struct {
	Address string `json:"address"`
	Percent int    `json:"percent"`
}
