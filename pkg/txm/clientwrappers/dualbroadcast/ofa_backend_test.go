package dualbroadcast

import (
	"context"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys/keystest"
	txmtypes "github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

// --- Flashbots test doubles

type testFlashbotsRPC struct {
	block *evmtypes.Block
}

func (m *testFlashbotsRPC) BlockByNumber(context.Context, *big.Int) (*evmtypes.Block, error) {
	return m.block, nil
}

func (m *testFlashbotsRPC) NonceAt(context.Context, common.Address, *big.Int) (uint64, error) {
	return 0, nil
}

func (m *testFlashbotsRPC) SendTransaction(context.Context, *txmtypes.Transaction, *txmtypes.Attempt) error {
	return nil
}

type testFlashbotsTxStore struct {
	txs []*txmtypes.Transaction
}

func (s *testFlashbotsTxStore) FetchUnconfirmedTransactions(context.Context, common.Address) ([]*txmtypes.Transaction, error) {
	return s.txs, nil
}

func TestParseURLParams(t *testing.T) {
	tests := []struct {
		name           string
		params         string
		wantPrivacy    privacy
		wantRefund     refundConfig
		wantErr        bool
		wantErrContain string
	}{
		{
			name:        "empty params",
			params:      "",
			wantPrivacy: privacy{},
			wantRefund:  refundConfig{},
		},
		{
			name:        "auctionTimeout",
			params:      "auctionTimeout=60",
			wantPrivacy: privacy{AuctionTimeout: 60},
			wantRefund:  refundConfig{},
		},
		{
			name:        "auctionTimeout invalid ignored",
			params:      "auctionTimeout=notanint",
			wantPrivacy: privacy{},
			wantRefund:  refundConfig{},
		},
		{
			name:        "single builder",
			params:      "builder=test_builder",
			wantPrivacy: privacy{Builders: []string{"test_builder"}},
			wantRefund:  refundConfig{},
		},
		{
			name:        "multiple builders",
			params:      "builder=test_builder_1&builder=test_builder_2",
			wantPrivacy: privacy{Builders: []string{"test_builder_1", "test_builder_2"}},
			wantRefund:  refundConfig{},
		},
		{
			name:        "single hint",
			params:      "hint=calldata",
			wantPrivacy: privacy{Hints: []string{"calldata"}},
			wantRefund:  refundConfig{},
		},
		{
			name:        "multiple hints",
			params:      "hint=calldata&hint=hash",
			wantPrivacy: privacy{Hints: []string{"calldata", "hash"}},
			wantRefund:  refundConfig{},
		},
		{
			name:        "refund valid",
			params:      "refund=0xRefundAddr:50",
			wantPrivacy: privacy{WantRefund: 50},
			wantRefund:  refundConfig{Address: "0xRefundAddr", Percent: 100},
		},
		{
			name:           "refund invalid percent",
			params:         "refund=0xRefundAddr:bad",
			wantErr:        true,
			wantErrContain: "unable to parse percentage",
		},
		{
			name:           "refund single part returns error",
			params:         "refund=0xRefundAddr",
			wantErr:        true,
			wantErrContain: "unable to parse refund",
		},
		{
			name:           "refund three parts returns error",
			params:         "refund=0xRefundAddr:50:extra",
			wantErr:        true,
			wantErrContain: "unable to parse refund",
		},
		{
			name:           "invalid query",
			params:         "%",
			wantErr:        true,
			wantErrContain: "unable to parse params",
		},
		{
			name:   "combined params",
			params: "auctionTimeout=120&builder=test_builder_1&builder=test_builder_2&hint=h1&refund=0xR:75",
			wantPrivacy: privacy{
				AuctionTimeout: 120,
				Builders:       []string{"test_builder_1", "test_builder_2"},
				Hints:          []string{"h1"},
				WantRefund:     75,
			},
			wantRefund: refundConfig{Address: "0xR", Percent: 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			privacy, refund, err := parseURLParams(tt.params)
			if tt.wantErr {
				require.Error(t, err)
				if tt.wantErrContain != "" {
					assert.Contains(t, err.Error(), tt.wantErrContain)
				}
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantPrivacy, privacy)
			assert.Equal(t, tt.wantRefund, refund)
		})
	}
}

func TestSendBundle_UsesLatestAttemptPerTransaction(t *testing.T) {
	fromAddress := common.HexToAddress("0x123")
	toAddress := common.HexToAddress("0x456")

	makeTx := func(nonce uint64, marker byte) *evmtypes.Transaction {
		return evmtypes.NewTx(&evmtypes.LegacyTx{
			Nonce:    nonce,
			To:       &toAddress,
			Gas:      21000,
			GasPrice: big.NewInt(1),
			Value:    big.NewInt(0),
			Data:     []byte{marker},
		})
	}

	oldAttemptTx := makeTx(1, 0x01)
	latestAttemptTx := makeTx(1, 0x02)
	secondTx := makeTx(2, 0x03)

	nonce1 := uint64(1)
	nonce2 := uint64(2)
	txStore := &testFlashbotsTxStore{txs: []*txmtypes.Transaction{
		{
			Nonce: &nonce1,
			Attempts: []*txmtypes.Attempt{
				{ID: 10, SignedTransaction: oldAttemptTx},
				{ID: 11, SignedTransaction: latestAttemptTx},
			},
		},
		{
			Nonce: &nonce2,
			Attempts: []*txmtypes.Attempt{
				{ID: 20, SignedTransaction: secondTx},
			},
		},
	}}

	var requestBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		requestBody, err = io.ReadAll(r.Body)
		require.NoError(t, err)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"bundleHash":"0xabc"}}`))
		require.NoError(t, err)
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	rpc := &testFlashbotsRPC{block: evmtypes.NewBlockWithHeader(&evmtypes.Header{Number: big.NewInt(100)})}
	metrics, mErr := newOFAMetrics("1", "flashbots")
	require.NoError(t, mErr)
	client := newFlashbotsClient(logger.Test(t), rpc, keystest.MessageSigner(nil), customURL, txStore, false, metrics)

	err = client.sendBundle(context.Background(), fromAddress, nil)
	require.NoError(t, err)

	var req struct {
		Method string `json:"method"`
		Params []struct {
			Body []struct {
				Tx string `json:"tx"`
			} `json:"body"`
		} `json:"params"`
	}
	require.NoError(t, json.Unmarshal(requestBody, &req))
	require.Equal(t, "mev_sendBundle", req.Method)
	require.Len(t, req.Params, 1)
	require.Len(t, req.Params[0].Body, 2)

	expectedLatestTx, err := latestAttemptTx.MarshalBinary()
	require.NoError(t, err)
	assert.Equal(t, "0x"+common.Bytes2Hex(expectedLatestTx), req.Params[0].Body[0].Tx)
}

func TestSendBundle_SucceedsOnIncreasingNonces(t *testing.T) {
	fromAddress := common.HexToAddress("0x123")
	toAddress := common.HexToAddress("0x456")

	makeTx := func(nonce uint64) *evmtypes.Transaction {
		return evmtypes.NewTx(&evmtypes.LegacyTx{
			Nonce:    nonce,
			To:       &toAddress,
			Gas:      21000,
			GasPrice: big.NewInt(1),
			Value:    big.NewInt(0),
		})
	}

	nonce7 := uint64(7)
	nonce8 := uint64(8)
	nonce9 := uint64(9)
	txStore := &testFlashbotsTxStore{txs: []*txmtypes.Transaction{
		{Nonce: &nonce7, Attempts: []*txmtypes.Attempt{{ID: 1, SignedTransaction: makeTx(7)}}},
		{Nonce: &nonce8, Attempts: []*txmtypes.Attempt{{ID: 2, SignedTransaction: makeTx(8)}}},
		{Nonce: &nonce9, Attempts: []*txmtypes.Attempt{{ID: 3, SignedTransaction: makeTx(9)}}},
	}}

	var requestBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		requestBody, err = io.ReadAll(r.Body)
		require.NoError(t, err)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"bundleHash":"0xabc"}}`))
		require.NoError(t, err)
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL)
	require.NoError(t, err)

	rpc := &testFlashbotsRPC{block: evmtypes.NewBlockWithHeader(&evmtypes.Header{Number: big.NewInt(100)})}
	metrics, mErr := newOFAMetrics("1", "flashbots")
	require.NoError(t, mErr)
	client := newFlashbotsClient(logger.Test(t), rpc, keystest.MessageSigner(nil), customURL, txStore, false, metrics)
	err = client.sendBundle(context.Background(), fromAddress, nil)
	require.NoError(t, err)

	var req struct {
		Params []struct {
			Body []struct {
				Tx string `json:"tx"`
			} `json:"body"`
		} `json:"params"`
	}
	require.NoError(t, json.Unmarshal(requestBody, &req))
	require.Len(t, req.Params, 1)
	require.Len(t, req.Params[0].Body, 3)
}

func TestSendBundle_ReturnsErrorOnNonceGap(t *testing.T) {
	fromAddress := common.HexToAddress("0x123")
	toAddress := common.HexToAddress("0x456")

	makeTx := func(nonce uint64) *evmtypes.Transaction {
		return evmtypes.NewTx(&evmtypes.LegacyTx{
			Nonce:    nonce,
			To:       &toAddress,
			Gas:      21000,
			GasPrice: big.NewInt(1),
			Value:    big.NewInt(0),
		})
	}

	nonce8 := uint64(8)
	nonce6 := uint64(6)
	txStore := &testFlashbotsTxStore{txs: []*txmtypes.Transaction{
		{Nonce: &nonce6, Attempts: []*txmtypes.Attempt{{ID: 1, SignedTransaction: makeTx(nonce6)}}},
		{Nonce: &nonce8, Attempts: []*txmtypes.Attempt{{ID: 2, SignedTransaction: makeTx(nonce8)}}},
	}}

	customURL, err := url.Parse("http://localhost")
	require.NoError(t, err)

	metrics, mErr := newOFAMetrics("1", "flashbots")
	require.NoError(t, mErr)
	client := newFlashbotsClient(logger.Test(t), &testFlashbotsRPC{}, keystest.MessageSigner(nil), customURL, txStore, false, metrics)
	err = client.sendBundle(context.Background(), fromAddress, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "must be contiguous and strictly increasing")
}

// --- Nova test doubles

type testNovaRPC struct {
	nonceAtCalls []nonceAtCall
	nonceAtNonce uint64
	nonceAtErr   error
	sendTxCalls  []sendTxCall
	sendTxErr    error
}

type nonceAtCall struct {
	Address     common.Address
	BlockNumber *big.Int
}

type sendTxCall struct {
	Tx      *txmtypes.Transaction
	Attempt *txmtypes.Attempt
}

func (m *testNovaRPC) NonceAt(_ context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	m.nonceAtCalls = append(m.nonceAtCalls, nonceAtCall{Address: address, BlockNumber: blockNumber})
	return m.nonceAtNonce, m.nonceAtErr
}

func (m *testNovaRPC) SendTransaction(_ context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error {
	m.sendTxCalls = append(m.sendTxCalls, sendTxCall{Tx: tx, Attempt: attempt})
	return m.sendTxErr
}

func (m *testNovaRPC) BlockByNumber(_ context.Context, number *big.Int) (*evmtypes.Block, error) {
	return nil, nil
}

func testOFAMetrics(t *testing.T) ofaMetrics {
	t.Helper()
	m, err := newOFAMetrics("1", "nova")
	require.NoError(t, err)
	return m
}

func newDualBroadcastTx(t *testing.T, nonce uint64) (*txmtypes.Transaction, *txmtypes.Attempt) {
	t.Helper()
	toAddress := common.HexToAddress("0x456")
	signedTx := evmtypes.NewTx(&evmtypes.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Gas:      21000,
		GasPrice: big.NewInt(1),
	})
	dualBroadcast := true
	metaJSON, err := json.Marshal(txmtypes.TxMeta{DualBroadcast: &dualBroadcast})
	require.NoError(t, err)
	meta := sqlutil.JSON(metaJSON)

	tx := &txmtypes.Transaction{
		Nonce:       &nonce,
		FromAddress: common.HexToAddress("0x123"),
		ToAddress:   toAddress,
		Meta:        &meta,
	}

	attempt := &txmtypes.Attempt{
		ID:                1,
		SignedTransaction: signedTx,
	}
	return tx, attempt
}

func TestNovaClient_SendTransaction_DualBroadcast(t *testing.T) {
	var receivedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		receivedBody, err = io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		// Verify no Flashbots signature header
		assert.Empty(t, r.Header.Get("X-Flashbots-signature"))
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xabc"}`))
		assert.NoError(t, err)
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL + "?api_key=test_key")
	require.NoError(t, err)

	rpc := &testNovaRPC{}
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	tx, attempt := newDualBroadcastTx(t, 42)
	err = client.SendTransaction(context.Background(), tx, attempt)
	require.NoError(t, err)

	var req struct {
		Method string   `json:"method"`
		Params []string `json:"params"`
	}
	require.NoError(t, json.Unmarshal(receivedBody, &req))
	assert.Equal(t, "eth_sendRawTransaction", req.Method)
	assert.Len(t, req.Params, 1)
	assert.Contains(t, req.Params[0], "0x") // hex-encoded RLP

	assert.Empty(t, rpc.sendTxCalls, "dual-broadcast tx should go to Nova, not chain RPC fallback")
}

func TestNovaClient_SendTransaction_NonDual_errors(t *testing.T) {
	rpc := &testNovaRPC{}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	nonce := uint64(1)
	toAddress := common.HexToAddress("0x456")
	signedTx := evmtypes.NewTx(&evmtypes.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Gas:      21000,
		GasPrice: big.NewInt(1),
	})
	tx := &txmtypes.Transaction{
		Nonce:       &nonce,
		FromAddress: common.HexToAddress("0x123"),
		ToAddress:   toAddress,
		Meta:        nil,
	}
	attempt := &txmtypes.Attempt{SignedTransaction: signedTx}

	err = client.SendTransaction(context.Background(), tx, attempt)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "non-dual-broadcast")
	assert.Empty(t, rpc.sendTxCalls)
}

func TestNovaClient_SendTransaction_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`internal server error`))
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL + "?api_key=test")
	require.NoError(t, err)

	rpc := &testNovaRPC{}
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	tx, attempt := newDualBroadcastTx(t, 1)
	err = client.SendTransaction(context.Background(), tx, attempt)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "status 500")
}

func TestNovaClient_SendTransaction_RPCError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"error":{"message":"nonce too low"}}`))
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL + "?api_key=test")
	require.NoError(t, err)

	rpc := &testNovaRPC{}
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	tx, attempt := newDualBroadcastTx(t, 1)
	err = client.SendTransaction(context.Background(), tx, attempt)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nonce too low")
}

func TestNovaClient_PendingNonceAt(t *testing.T) {
	var receivedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		receivedBody, err = io.ReadAll(r.Body)
		assert.NoError(t, err)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x2a"}`))
		assert.NoError(t, err)
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL + "?api_key=test")
	require.NoError(t, err)

	rpc := &testNovaRPC{}
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	addr := common.HexToAddress("0x123")
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	require.NoError(t, err)
	assert.Equal(t, uint64(42), nonce)

	var req struct {
		Method string   `json:"method"`
		Params []string `json:"params"`
	}
	require.NoError(t, json.Unmarshal(receivedBody, &req))
	assert.Equal(t, "eth_getTransactionCount", req.Method)
	assert.Equal(t, []string{addr.Hex(), "pending"}, req.Params)

	assert.Empty(t, rpc.nonceAtCalls, "PendingNonceAt should call Nova directly, not chain RPC")
}

func TestNovaClient_PendingNonceAt_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`internal server error`))
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL + "?api_key=test")
	require.NoError(t, err)

	rpc := &testNovaRPC{}
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	_, err = client.PendingNonceAt(context.Background(), common.HexToAddress("0x123"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nova eth_getTransactionCount failed")
}

func TestNovaClient_PendingNonceAt_RPCError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"error":{"message":"unknown address"}}`))
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL + "?api_key=test")
	require.NoError(t, err)

	rpc := &testNovaRPC{}
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	_, err = client.PendingNonceAt(context.Background(), common.HexToAddress("0x123"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown address")
}

func TestNovaClient_NonceAt(t *testing.T) {
	rpc := &testNovaRPC{nonceAtNonce: 7}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	addr := common.HexToAddress("0x123")
	blockNum := big.NewInt(100)
	nonce, err := client.NonceAt(context.Background(), addr, blockNum)
	require.NoError(t, err)
	assert.Equal(t, uint64(7), nonce)

	require.Len(t, rpc.nonceAtCalls, 1)
	assert.Equal(t, addr, rpc.nonceAtCalls[0].Address)
	assert.Equal(t, blockNum, rpc.nonceAtCalls[0].BlockNumber)
}
