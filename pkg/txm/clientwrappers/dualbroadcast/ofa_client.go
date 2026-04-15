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

	"github.com/smartcontractkit/chainlink-evm/pkg/keys"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

const rpcTimeout = 10 * time.Second

// publicMempoolRPC is the chain RPC surface used when a transaction is not sent to the relay.
type publicMempoolRPC interface {
	BlockByNumber(ctx context.Context, number *big.Int) (*evmtypes.Block, error)
	NonceAt(context.Context, common.Address, *big.Int) (uint64, error)
	SendTransaction(context.Context, *types.Transaction, *types.Attempt) error
}

// ofaHTTPAuth configures how JSON-RPC POSTs are addressed and authenticated per OFA provider.
type ofaHTTPAuth interface { // TODO(gg): simplify
	requestURL(base *url.URL, extraQuery string) string
	apply(ctx context.Context, req *http.Request, body []byte, from common.Address) error
}

type flashbotsHTTPAuth struct {
	keystore keys.MessageSigner
}

func (a *flashbotsHTTPAuth) requestURL(base *url.URL, extraQuery string) string {
	return base.String() + "?" + extraQuery
}

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

// novaHTTPAuth posts to the configured base URL (API key lives in the URL query). No signing.
type novaHTTPAuth struct{}

func (novaHTTPAuth) requestURL(base *url.URL, extraQuery string) string {
	_ = extraQuery // dual-broadcast URL params are Flashbots-specific; Nova uses the configured URL only
	return base.String()
}

func (novaHTTPAuth) apply(context.Context, *http.Request, []byte, common.Address) error {
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
	raw, err := r.postJSONRPC(ctx, address, body, "")
	if err != nil {
		return 0, fmt.Errorf("%s eth_getTransactionCount failed: %w", r.errHTTPPrefix, err)
	}

	var nonce hexutil.Uint64
	if err = json.Unmarshal(raw, (*hexutil.Uint64)(&nonce)); err != nil {
		return 0, fmt.Errorf("failed to parse nonce from relay response: %w: %s", err, string(raw))
	}
	return uint64(nonce), nil
}

// sendDualBroadcastTx sends a dual-broadcast transaction to the OFA. It returns URL-encoded params from meta (for Flashbots bundle follow-up).
func (r *ofaClient) sendDualBroadcastTx(ctx context.Context, tx *types.Transaction, attempt *types.Attempt) (params string, err error) {
	meta, err := tx.GetMeta()
	if err != nil {
		return "", err
	}

	data, err := attempt.SignedTransaction.MarshalBinary()
	if err != nil {
		return "", err
	}

	params = ""
	if meta.DualBroadcastParams != nil {
		params = *meta.DualBroadcastParams
	}

	body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["%s"], "id":1}`, hexutil.Encode(data)))
	start := time.Now()
	_, err = r.postJSONRPC(ctx, tx.FromAddress, body, params)
	r.metrics.RecordSendTx(ctx, time.Since(start), err)

	return params, err
}

func (r *ofaClient) postJSONRPC(ctx context.Context, from common.Address, body []byte, urlParams string) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, rpcTimeout)
	defer cancel()

	postURL := r.auth.requestURL(r.customURL, urlParams)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if err := r.auth.apply(ctx, req, body, from); err != nil {
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
