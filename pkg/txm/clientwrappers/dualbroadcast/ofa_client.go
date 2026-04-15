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

	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const rpcTimeout = 10 * time.Second

// publicMempoolRPC is the chain RPC surface used when a transaction is not sent to the relay.
type publicMempoolRPC interface {
	BlockByNumber(ctx context.Context, number *big.Int) (*evmtypes.Block, error)
	NonceAt(context.Context, common.Address, *big.Int) (uint64, error)
	SendTransaction(context.Context, *types.Transaction, *types.Attempt) error
}

// ofaHTTPAuth configures how JSON-RPC POSTs are addressed and authenticated.
type ofaHTTPAuth interface {
	requestURL(base *url.URL, meta *types.TxMeta) string
	apply(ctx context.Context, req *http.Request, body []byte, from common.Address) error
}

// noAuth does not add any additional URL parameters or signing.
type noAuth struct{}

func (noAuth) requestURL(base *url.URL, _ *types.TxMeta) string {
	return base.String()
}

func (noAuth) apply(context.Context, *http.Request, []byte, common.Address) error {
	return nil
}

// ofaClient performs JSON-RPC over HTTP to OFA relays (Flashbots, Nova, …). Signing and
// URL rules are injected via ofaHTTPAuth.
type ofaClient struct {
	c             publicMempoolRPC
	customURL     *url.URL
	auth          ofaHTTPAuth
	metrics       ofaMetrics
	errHTTPPrefix string
}

func newOFAClient(
	c publicMempoolRPC,
	customURL *url.URL,
	auth ofaHTTPAuth,
	metrics ofaMetrics,
	errHTTPPrefix string,
) *ofaClient {
	return &ofaClient{
		c:             c,
		customURL:     customURL,
		auth:          auth,
		metrics:       metrics,
		errHTTPPrefix: errHTTPPrefix,
	}
}

func (r *ofaClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	return r.c.NonceAt(ctx, address, blockNumber)
}

func (r *ofaClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, rpcTimeout)
	defer cancel()
	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getTransactionCount","params":["%s","pending"], "id":1}`, address.Hex()))
	raw, err := r.postJSONRPC(ctx, address, body, nil)
	if err != nil {
		return 0, fmt.Errorf("%s eth_getTransactionCount failed: %w", r.errHTTPPrefix, err)
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

// sendDualBroadcastTx sends a dual-broadcast transaction to the OFA relay.
func (r *ofaClient) sendDualBroadcastTx(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) error {
	meta, err := tx.GetMeta()
	if err != nil {
		return err
	}

	data, err := attempt.SignedTransaction.MarshalBinary()
	if err != nil {
		return err
	}

	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["%s"], "id":1}`, hexutil.Encode(data)))
	start := time.Now()
	_, err = r.postJSONRPC(ctx, tx.FromAddress, body, meta)
	r.metrics.RecordSendTx(ctx, time.Since(start), err)

	return err
}

func (r *ofaClient) postJSONRPC(ctx context.Context, from common.Address, body []byte, meta *types.TxMeta) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, rpcTimeout)
	defer cancel()

	postURL := r.auth.requestURL(r.customURL, meta)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if err = r.auth.apply(ctx, req, body, from); err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s request failed: %w", r.errHTTPPrefix, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s request failed with status %d: %s", r.errHTTPPrefix, resp.StatusCode, string(respBody))
	}

	var response postResponse
	if err = json.Unmarshal(respBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal relay response: %w: %s", err, string(respBody))
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
