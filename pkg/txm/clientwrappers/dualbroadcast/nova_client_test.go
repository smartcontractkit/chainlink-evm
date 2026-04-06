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

// TODO(gg): add/update tests to validate OFAMetrics
// TODO(gg): add test for TestParseURLParams? Similar to flashbots_client_test.go

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

	rpc := &testFlashbotsRPC{}
	client := NewNovaClient(logger.Test(t), rpc, customURL, nil)

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
}

func TestNovaClient_SendTransaction_NonDual(t *testing.T) {
	rpc := &testFlashbotsRPC{}
	customURL, err := url.Parse("https://eth.novarpc.xyz?api_key=test") // TODO(gg): use localhost instead
	require.NoError(t, err)

	client := NewNovaClient(logger.Test(t), rpc, customURL, nil)

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
}

func TestNovaClient_SendTransaction_Purgeable(t *testing.T) {
	rpc := &testFlashbotsRPC{}
	customURL, err := url.Parse("https://eth.novarpc.xyz?api_key=test")
	require.NoError(t, err)

	client := NewNovaClient(logger.Test(t), rpc, customURL, nil)

	tx, attempt := newDualBroadcastTx(t, 1)
	tx.IsPurgeable = true

	err = client.SendTransaction(context.Background(), tx, attempt)
	require.NoError(t, err)
}

func TestNovaClient_SendTransaction_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`internal server error`))
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL + "?api_key=test")
	require.NoError(t, err)

	rpc := &testFlashbotsRPC{}
	client := NewNovaClient(logger.Test(t), rpc, customURL, nil)

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

	rpc := &testFlashbotsRPC{}
	client := NewNovaClient(logger.Test(t), rpc, customURL, nil)

	tx, attempt := newDualBroadcastTx(t, 1)
	err = client.SendTransaction(context.Background(), tx, attempt)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nonce too low")
}

func TestNovaClient_PendingNonceAt(t *testing.T) {
	rpc := &testFlashbotsRPC{}
	customURL, err := url.Parse("https://eth.novarpc.xyz?api_key=test")
	require.NoError(t, err)

	client := NewNovaClient(logger.Test(t), rpc, customURL, nil)

	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress("0x123"))
	require.NoError(t, err)
	assert.Equal(t, uint64(0), nonce)
}

func TestNovaClient_NonceAt(t *testing.T) {
	rpc := &testFlashbotsRPC{}
	customURL, err := url.Parse("https://eth.novarpc.xyz?api_key=test")
	require.NoError(t, err)

	client := NewNovaClient(logger.Test(t), rpc, customURL, nil)

	nonce, err := client.NonceAt(context.Background(), common.HexToAddress("0x123"), big.NewInt(100))
	require.NoError(t, err)
	assert.Equal(t, uint64(0), nonce)
}
