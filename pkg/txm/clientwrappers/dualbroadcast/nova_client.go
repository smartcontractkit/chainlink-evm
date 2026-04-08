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

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

// TODO(gg): where is the refund set and any other potential params? (And is it tested?)

const novaRPCTimeout = 10 * time.Second

type novaClientRPC interface {
	NonceAt(context.Context, common.Address, *big.Int) (uint64, error)
	SendTransaction(context.Context, *types.Transaction, *types.Attempt) error
}

type novaClient struct {
	lggr      logger.SugaredLogger
	c         novaClientRPC
	customURL *url.URL
	metrics   ofaMetrics
}

var _ txm.Client = (*novaClient)(nil)

func newNovaClient(lggr logger.Logger, c novaClientRPC, customURL *url.URL, metrics ofaMetrics) *novaClient {
	return &novaClient{
		lggr:      logger.Sugared(logger.Named(lggr, "Txm.NovaClient")),
		c:         c,
		customURL: customURL,
		metrics:   metrics,
	}
}

func (n *novaClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return n.c.NonceAt(ctx, address, blockNumber)
}

func (n *novaClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	body := []byte(fmt.Sprintf(
		`{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["%s","pending"],"id":1}`,
		address.Hex(),
	))

	result, err := n.postToNovaWithResult(ctx, body)
	if err != nil {
		return 0, fmt.Errorf("nova eth_getTransactionCount failed: %w", err)
	}

	var nonce hexutil.Uint64
	if err = json.Unmarshal(result, (*hexutil.Uint64)(&nonce)); err != nil {
		return 0, fmt.Errorf("failed to parse nonce from nova response: %w: %s", err, string(result))
	}
	return uint64(nonce), nil
}

func (n *novaClient) SendTransaction(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	if meta != nil && meta.DualBroadcast != nil && *meta.DualBroadcast && !tx.IsPurgeable {
		return n.sendToNova(ctx, tx, attempt)
	}

	// fallback to chain client if not dual-broadcast
	return n.c.SendTransaction(ctx, nil, attempt)
}

func (n *novaClient) sendToNova(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	data, err := attempt.SignedTransaction.MarshalBinary()
	if err != nil {
		return err
	}

	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["%s"],"id":1}`, hexutil.Encode(data)))

	start := time.Now()
	err = n.postToNova(ctx, body)
	n.metrics.RecordSendTx(ctx, time.Since(start), err)
	if err != nil {
		return err
	}
	n.lggr.Debugw("Sent transaction to Nova", "txHash", attempt.Hash, "transactionLifecycleID", tx.GetTransactionLifecycleID(n.lggr))
	return nil
}

func (n *novaClient) postToNova(ctx context.Context, body []byte) error {
	_, err := n.postToNovaWithResult(ctx, body)
	return err
}

func (n *novaClient) postToNovaWithResult(ctx context.Context, body []byte) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, novaRPCTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, n.customURL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("nova request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nova request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	var response postResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal nova response: %w: %s", err, string(respBody))
	}
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response.Result, nil
}
