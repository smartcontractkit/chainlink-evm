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

const flashbotsRPCTimeout = 10 * time.Second

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
}

func NewFlashbotsClient(lggr logger.Logger, c FlashbotsClientRPC, keystore keys.MessageSigner, customURL *url.URL, txStore FlashbotsTxStore) *FlashbotsClient {
	return &FlashbotsClient{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.FlashbotsClient")),
		c:         c,
		keystore:  keystore,
		customURL: customURL,
		txStore:   txStore,
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
		return 0, fmt.Errorf("failed to unmarshal response into string: %w", err)
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
		if err := d.SendBundle(ctx, tx.FromAddress, params); err != nil {
			d.lggr.Error("error sending bundle: ", err)
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
	postReq.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		return nil, fmt.Errorf("request %v failed: %w", postReq, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request %v failed with status: %d", postReq, resp.StatusCode)
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

func (d *FlashbotsClient) SendBundle(ctx context.Context, fromAddress common.Address, urlParams string) error {
	unconfirmedTxs, err := d.txStore.FetchUnconfirmedTransactions(ctx, fromAddress)
	if err != nil {
		return fmt.Errorf("failed to fetch unconfirmed transactions: %w", err)
	}

	// TODO: Get the first attempt from each transaction for now.
	attempts := make([]*types.Attempt, 0, len(unconfirmedTxs))
	nonces := make([]uint64, 0, len(unconfirmedTxs))
	for _, unconfirmedTx := range unconfirmedTxs {
		if len(unconfirmedTx.Attempts) > 0 {
			attempts = append(attempts, unconfirmedTx.Attempts[0])
			if unconfirmedTx.Nonce != nil {
				nonces = append(nonces, *unconfirmedTx.Nonce)
			}
		}
	}

	// Need at least 2 transactions to send a bundle
	if len(attempts) < 2 {
		return nil
	}

	// TODO: we don't have a good way to get this other than making an RPC call. Some async caching may help with the overhead.
	currentBlock, err := d.c.BlockByNumber(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to get current block height: %w", err)
	}
	targetBlock := currentBlock.NumberU64() + 10

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
			"canRevert": true,
		})
	}

	bundleParams := map[string]interface{}{
		"version": "v0.1",
		"inclusion": map[string]interface{}{
			"block": hexutil.EncodeBig(new(big.Int).SetUint64(targetBlock)),
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

	raw, err := d.signAndPostMessage(ctx, fromAddress, bodyBytes, urlParams)
	if err != nil {
		return err
	}

	var bundleResult struct {
		BundleHash string `json:"bundleHash"`
	}
	if err := json.Unmarshal(raw, &bundleResult); err != nil {
		return fmt.Errorf("failed to decode response %v into bundle result: %w", raw, err)
	}
	d.lggr.Infow("Broadcasted transaction bundle", "nonces", nonces, "bundleHash", bundleResult.BundleHash)
	return nil
}
