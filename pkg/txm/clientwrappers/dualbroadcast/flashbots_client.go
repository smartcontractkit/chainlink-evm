package dualbroadcast

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const maxBlockDiff = 24

type FlashbotsTxStore interface {
	FetchUnconfirmedTransactions(context.Context, common.Address) ([]*types.Transaction, error)
}

type FlashbotsClient struct {
	lggr      logger.SugaredLogger
	ofaClient *ofaClient
	txStore   FlashbotsTxStore
	bundles   bool
}

var _ txm.Client = (*FlashbotsClient)(nil)

func NewFlashbotsClient(lggr logger.Logger, c publicMempoolRPC, keystore keys.MessageSigner, customURL *url.URL, txStore FlashbotsTxStore, bundles *bool, metrics ofaMetrics) *FlashbotsClient {
	log := logger.Sugared(logger.Named(lggr, "Txm.FlashbotsClient"))
	ofaClient := newOFAClient(c, customURL, &flashbotsHTTPAuth{keystore: keystore}, metrics, "flashbots")
	b := bundles != nil && *bundles

	return &FlashbotsClient{
		lggr:      log,
		ofaClient: ofaClient,
		txStore:   txStore,
		bundles:   b,
	}
}

func (d *FlashbotsClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return d.ofaClient.NonceAt(ctx, address, blockNumber)
}

func (d *FlashbotsClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return d.ofaClient.PendingNonceAt(ctx, address)
}

// SendTransaction sends a dual-broadcast transaction to Flashbots when meta says so; otherwise it falls back to the public chain RPC.
func (d *FlashbotsClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	// If not dual-broadcast, fall back to sending the transaction to the chain RPC directly
	if meta == nil || meta.DualBroadcast == nil || !*meta.DualBroadcast || tx.IsPurgeable {
		return d.ofaClient.c.SendTransaction(ctx, nil, attempt)
	}

	if err := d.ofaClient.sendDualBroadcastTx(ctx, tx, attempt); err != nil {
		return err
	}

	if d.bundles {
		if err := d.sendBundle(ctx, tx.FromAddress, meta); err != nil {
			d.lggr.Errorw("error sending bundle", "err", err, "transactionLifecycleID", tx.GetTransactionLifecycleID(d.lggr))
		}
	}

	return nil
}

type flashbotsHTTPAuth struct {
	keystore keys.MessageSigner
}

var _ ofaHTTPAuth = (*flashbotsHTTPAuth)(nil)

// requestURL returns the base URL with the extra Flashbots query parameters.
func (a *flashbotsHTTPAuth) requestURL(base *url.URL, meta *types.TxMeta) string {
	var params string
	if meta != nil && meta.DualBroadcastParams != nil {
		params = *meta.DualBroadcastParams
	}

	return base.String() + "?" + params
}

// apply signs the body and adds the Flashbots headers.
func (a *flashbotsHTTPAuth) apply(ctx context.Context, req *http.Request, body []byte, from common.Address) error {
	hashedBody := crypto.Keccak256Hash(body).Hex()
	signedMessage, err := a.keystore.SignMessage(ctx, from, []byte(hashedBody))
	if err != nil {
		return err
	}
	req.Header.Add("X-Flashbots-signature", from.String()+":"+hexutil.Encode(signedMessage))
	req.Header.Add("X-Flashbots-Origin", "chainlink")
	return nil
}

// sendBundle sends a bundle of all the in-flight transactions.
func (d *FlashbotsClient) sendBundle(ctx context.Context, fromAddress common.Address, meta *types.TxMeta) error {
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

	// TODO: we don't have a good way to get this other than making an RPC call. Some async caching may help with the overhead.
	currentBlock, err := d.ofaClient.c.BlockByNumber(ctx, nil)
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

	raw, err := d.ofaClient.postJSONRPC(ctx, fromAddress, bodyBytes, nil)
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
