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

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

var _ txm.Client = &FlashbotsClient{}

type FlashbotsClient struct {
	c         client.Client
	keystore  keys.MessageSigner
	customURL *url.URL
}

func NewFlashbotsClient(c client.Client, keystore keys.MessageSigner, customURL *url.URL) *FlashbotsClient {
	return &FlashbotsClient{
		c:         c,
		keystore:  keystore,
		customURL: customURL,
	}
}

func (d *FlashbotsClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return d.c.NonceAt(ctx, address, blockNumber)
}

func (d *FlashbotsClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["%s","pending"], "id":1}`, address.String()))
	response, err := d.signAndPostMessage(ctx, address, body, "")
	if err != nil {
		return 0, err
	}

	nonce, err := hexutil.DecodeUint64(response)
	if err != nil {
		return 0, fmt.Errorf("failed to decode response %v into uint64: %w", response, err)
	}
	return nonce, nil
}

func (d *FlashbotsClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt, txStore txm.TxStore) error {
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
		_ = d.SendBundle(ctx, txStore, tx.FromAddress, params)
		// Don't return bundle error - the single transaction was already sent successfully
		return nil
	}

	return d.c.SendTransaction(ctx, attempt.SignedTransaction)
}

func (d *FlashbotsClient) signAndPostMessage(ctx context.Context, address common.Address, body []byte, urlParams string) (result string, err error) {
	bodyReader := bytes.NewReader(body)
	postReq, err := http.NewRequestWithContext(ctx, http.MethodPost, d.customURL.String()+"?"+urlParams, bodyReader)
	if err != nil {
		return
	}

	hashedBody := crypto.Keccak256Hash(body).Hex()
	signedMessage, err := d.keystore.SignMessage(ctx, address, []byte(hashedBody))
	if err != nil {
		return
	}

	postReq.Header.Add("X-Flashbots-signature", address.String()+":"+hexutil.Encode(signedMessage))
	postReq.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		return result, fmt.Errorf("request %v failed: %w", postReq, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("request %v failed with status: %d", postReq, resp.StatusCode)
	}

	keyJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var response postResponse
	err = json.Unmarshal(keyJSON, &response)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal response into struct: %w: %s", err, string(keyJSON))
	}
	if response.Error.Message != "" {
		return result, errors.New(response.Error.Message)
	}
	return response.Result, nil
}

type postResponse struct {
	Result string `json:"result,omitempty"`
	Error  postError
}

type postError struct {
	Message string `json:"message,omitempty"`
}

func (d *FlashbotsClient) SendBundle(ctx context.Context, txStore txm.TxStore, fromAddress common.Address, urlParams string) error {
	unconfirmedTxs, err := txStore.FetchUnconfirmedTransactions(ctx, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to fetch unconfirmed transactions: %w", err)
	}

	// TODO: Get the first attempt from each transaction for now.
	attempts := make([]*types.Attempt, 0, len(unconfirmedTxs))
	for _, unconfirmedTx := range unconfirmedTxs {
		if len(unconfirmedTx.Attempts) > 0 {
			attempts = append(attempts, unconfirmedTx.Attempts[0])
		}
	}

	// Need at least 2 transactions to send a bundle
	if len(attempts) < 2 {
		return nil
	}

	// TODO: we don't have a good way to get this other than making an RPC call. Some async caching may help with the overhead.
	currentBlock, err := d.c.LatestBlockHeight(ctx)
	if err != nil {
		return fmt.Errorf("failed to get current block height: %w", err)
	}
	targetBlock := new(big.Int).Add(currentBlock, big.NewInt(1))

	bodyItems := make([]map[string]any, 0, len(attempts))
	for _, attempt := range attempts {
		if attempt.SignedTransaction == nil {
			return fmt.Errorf("attempt with ID %d has nil SignedTransaction", attempt.ID)
		}

		txData, err := attempt.SignedTransaction.MarshalBinary()
		if err != nil {
			return fmt.Errorf("failed to marshal transaction for attempt ID %d: %w", attempt.ID, err)
		}

		bodyItems = append(bodyItems, map[string]interface{}{
			"tx":        hexutil.Encode(txData),
			"canRevert": false,
		})
	}

	bundleParams := map[string]interface{}{
		"version": "v0.1",
		"inclusion": map[string]interface{}{
			"block": hexutil.EncodeBig(targetBlock),
		},
		"body": bodyItems,
	}

	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "mev_sendBundle",
		"params":  []any{bundleParams},
	}

	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal bundle request: %w", err)
	}

	_, err = d.signAndPostMessage(ctx, fromAddress, bodyBytes, urlParams)
	return err
}
