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
	txmtypes "github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

type testNovaRPC struct {
	nonceAtCalls    []nonceAtCall
	nonceAtNonce    uint64
	nonceAtErr      error
	sendTxCalls     []sendTxCall
	sendTxErr       error
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
		require.NoError(t, err)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		// Verify no Flashbots signature header
		assert.Empty(t, r.Header.Get("X-Flashbots-signature"))
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xabc"}`))
		require.NoError(t, err)
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

func TestNovaClient_SendTransaction_NonDual(t *testing.T) {
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
	}
	attempt := &txmtypes.Attempt{SignedTransaction: signedTx}

	err = client.SendTransaction(context.Background(), tx, attempt)
	require.NoError(t, err)

	require.Len(t, rpc.sendTxCalls, 1)
	assert.Nil(t, rpc.sendTxCalls[0].Tx, "fallback should pass nil tx to chain RPC")
	assert.Equal(t, attempt, rpc.sendTxCalls[0].Attempt)
}

func TestNovaClient_SendTransaction_Purgeable(t *testing.T) {
	rpc := &testNovaRPC{}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	tx, attempt := newDualBroadcastTx(t, 1)
	tx.IsPurgeable = true

	err = client.SendTransaction(context.Background(), tx, attempt)
	require.NoError(t, err)

	require.Len(t, rpc.sendTxCalls, 1)
	assert.Nil(t, rpc.sendTxCalls[0].Tx, "purgeable tx should fallback to chain RPC with nil tx")
	assert.Equal(t, attempt, rpc.sendTxCalls[0].Attempt)
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
	rpc := &testNovaRPC{nonceAtNonce: 42}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t))

	addr := common.HexToAddress("0x123")
	nonce, err := client.PendingNonceAt(context.Background(), addr)
	require.NoError(t, err)
	assert.Equal(t, uint64(42), nonce)

	require.Len(t, rpc.nonceAtCalls, 1)
	assert.Equal(t, addr, rpc.nonceAtCalls[0].Address)
	assert.Nil(t, rpc.nonceAtCalls[0].BlockNumber, "PendingNonceAt should delegate to NonceAt with nil block (latest)")
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
