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

const (
	flashbotsRPCTimeout = 10 * time.Second
	maxBlockDiff        = 24
)

type FlashbotsTxStore interface {
	FetchUnconfirmedTransactions(context.Context, common.Address) ([]*types.Transaction, error)
}

type FlashbotsClientRPC interface {
	BlockByNumber(ctx context.Context, number *big.Int) (*evmtypes.Block, error)
	NonceAt(context.Context, common.Address, *big.Int) (uint64, error)
	SendTransaction(context.Context, *evmtypes.Transaction) error
}

type FlashbotsClient struct {
	lggr      logger.SugaredLogger
	c         FlashbotsClientRPC
	keystore  keys.MessageSigner
	customURL *url.URL
	txStore   FlashbotsTxStore
	bundles   bool
}

func NewFlashbotsClient(lggr logger.Logger, c FlashbotsClientRPC, keystore keys.MessageSigner, customURL *url.URL, txStore FlashbotsTxStore, bundles *bool) *FlashbotsClient {
	b := bundles != nil && *bundles
	return &FlashbotsClient{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.FlashbotsClient")),
		c:         c,
		keystore:  keystore,
		customURL: customURL,
		txStore:   txStore,
		bundles:   b,
	}
}

func (d *FlashbotsClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return d.c.NonceAt(ctx, address, blockNumber)
}

func (d *FlashbotsClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, flashbotsRPCTimeout)
	defer cancel()
	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["%s","pending"], "id":1}`, address.String()))
	raw, err := d.signAndPostMessage(ctx, address, body, "")
	if err != nil {
		return 0, err
	}

	var resultStr string
	if err := json.Unmarshal(raw, &resultStr); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response %s into string: %w", string(raw), err)
	}
	nonce, err := hexutil.DecodeUint64(resultStr)
	if err != nil {
		return 0, fmt.Errorf("failed to decode response %v into uint64: %w", resultStr, err)
	}
	return nonce, nil
}

func (d *FlashbotsClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	if meta != nil && meta.DualBroadcast != nil && *meta.DualBroadcast && !tx.IsPurgeable {
		data, err := attempt.SignedTransaction.MarshalBinary()
		if err != nil {
			return err
		}
		params := ""
		if meta.DualBroadcastParams != nil {
			params = *meta.DualBroadcastParams
		}
		body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["%s"], "id":1}`, hexutil.Encode(data)))
		_, err = d.signAndPostMessage(ctx, tx.FromAddress, body, params)
		if err != nil {
			return err
		}

		// After successfully sending the transaction, send a bundle with all unconfirmed transactions
		// Don't act on a bundle error - this is a fire and forget operation but we do want to log the error.
		if d.bundles {
			if err := d.SendBundle(ctx, tx.FromAddress, params); err != nil {
				d.lggr.Errorw("error sending bundle", "err", err)
			}
		}
		return nil
	}

	return d.c.SendTransaction(ctx, attempt.SignedTransaction)
}

func (d *FlashbotsClient) signAndPostMessage(ctx context.Context, address common.Address, body []byte, urlParams string) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, flashbotsRPCTimeout)
	defer cancel()
	bodyReader := bytes.NewReader(body)
	postReq, err := http.NewRequestWithContext(ctx, http.MethodPost, d.customURL.String()+"?"+urlParams, bodyReader)
	if err != nil {
		return nil, err
	}

	hashedBody := crypto.Keccak256Hash(body).Hex()
	signedMessage, err := d.keystore.SignMessage(ctx, address, []byte(hashedBody))
	if err != nil {
		return nil, err
	}

	postReq.Header.Add("X-Flashbots-signature", address.String()+":"+hexutil.Encode(signedMessage))
	postReq.Header.Add("X-Flashbots-Origin", "chainlink")
	postReq.Header.Add("Content-Type", "application/json")

	reqDesc := fmt.Sprintf("%s %s body: %s", postReq.Method, postReq.URL.String(), string(body))
	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		return nil, fmt.Errorf("request %s failed: %w", reqDesc, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request %s failed with status: %d", reqDesc, resp.StatusCode)
	}

	keyJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response postResponse
	err = json.Unmarshal(keyJSON, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response into struct: %w: %s", err, string(keyJSON))
	}
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response.Result, nil
}

type postResponse struct {
	Result json.RawMessage `json:"result,omitempty"`
	Error  postError
}

type postError struct {
	Message string `json:"message,omitempty"`
}

// SendBundle sends a bundle of all the in-flight transactions.
func (d *FlashbotsClient) SendBundle(ctx context.Context, fromAddress common.Address, urlParams string) error {
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

	raw, err := d.signAndPostMessage(ctx, fromAddress, bodyBytes, "")
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

func parseURLParams(params string) (Privacy, RefundConfig, error) {
	values, err := url.ParseQuery(params)
	if err != nil {
		return Privacy{}, RefundConfig{}, fmt.Errorf("unable to parse params: %w", err)
	}

	privacy := Privacy{}
	if timeout, err := strconv.Atoi(values.Get("auctionTimeout")); err == nil {
		privacy.AuctionTimeout = timeout
	}

	privacy.Builders = append(privacy.Builders, values["builder"]...)

	privacy.Hints = append(privacy.Hints, values["hint"]...)

	refundConfig := RefundConfig{}
	refundRaw := values.Get("refund")
	if refundRaw != "" {
		parts := strings.Split(refundRaw, ":")
		if len(parts) != 2 {
			return Privacy{}, RefundConfig{}, fmt.Errorf("unable to parse refund: %s. Expected format: address:percent", refundRaw)
		}
		address := parts[0]
		percentVal, err := strconv.Atoi(parts[1])
		if err != nil {
			return Privacy{}, RefundConfig{}, fmt.Errorf("unable to parse percentage: %w", err)
		}

		privacy.WantRefund = percentVal
		refundConfig = RefundConfig{
			Address: address,
			Percent: 100, // wantRefund is an absolute percent of the refund, and refundConfig.percent=100 means entire refund goes to this address (no longer supported)
		}
	}
	return privacy, refundConfig, nil
}

type Privacy struct {
	WantRefund     int      `json:"wantRefund"`
	AuctionTimeout int      `json:"auctionTimeout"`
	Builders       []string `json:"builders"`
	Hints          []string `json:"hints"`
}

type RefundConfig struct {
	Address string `json:"address"`
	Percent int    `json:"percent"`
}
