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
	"github.com/ethereum/go-ethereum/common/hexutil"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys/keystest"
	txmtypes "github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

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
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), keystest.TxSigner(nil), nil)

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

func TestNovaClient_SendTransaction_NonDual_doesNothing(t *testing.T) {
	rpc := &testNovaRPC{}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), nil, nil)

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
		Meta:        nil, // No meta, so not dual-broadcast
	}
	attempt := &txmtypes.Attempt{SignedTransaction: signedTx}

	err = client.SendTransaction(context.Background(), tx, attempt)
	require.NoError(t, err)

	assert.Empty(t, rpc.sendTxCalls)
}

func TestNovaClient_SendTransaction_Purgeable_doesNothing(t *testing.T) {
	rpc := &testNovaRPC{}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), nil, nil)

	tx, attempt := newDualBroadcastTx(t, 1)
	tx.IsPurgeable = true

	err = client.SendTransaction(context.Background(), tx, attempt)
	require.NoError(t, err)

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
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), keystest.TxSigner(nil), nil)

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
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), keystest.TxSigner(nil), nil)

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
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), nil, nil)

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
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), nil, nil)

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
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), nil, nil)

	_, err = client.PendingNonceAt(context.Background(), common.HexToAddress("0x123"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown address")
}

func TestNovaClient_NonceAt(t *testing.T) {
	rpc := &testNovaRPC{nonceAtNonce: 7}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), nil, nil)

	addr := common.HexToAddress("0x123")
	blockNum := big.NewInt(100)
	nonce, err := client.NonceAt(context.Background(), addr, blockNum)
	require.NoError(t, err)
	assert.Equal(t, uint64(7), nonce)

	require.Len(t, rpc.nonceAtCalls, 1)
	assert.Equal(t, addr, rpc.nonceAtCalls[0].Address)
	assert.Equal(t, blockNum, rpc.nonceAtCalls[0].BlockNumber)
}

func TestNovaClient_TierValue(t *testing.T) {
	t.Parallel()

	tier2Addr := common.HexToAddress("0xaaa")
	regularAddr := common.HexToAddress("0xbbb")

	rpc := &testNovaRPC{}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), nil, []common.Address{tier2Addr})

	assert.Equal(t, big.NewInt(2), client.tierValue(tier2Addr))
	assert.Equal(t, big.NewInt(1), client.tierValue(regularAddr))
}

func TestNovaClient_TierValue_EmptyFeeds(t *testing.T) {
	t.Parallel()

	rpc := &testNovaRPC{}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), nil, nil)

	assert.Equal(t, big.NewInt(1), client.tierValue(common.HexToAddress("0xaaa")))
}

func TestNovaClient_ResignWithTierValue_Legacy(t *testing.T) {
	t.Parallel()

	tier2Addr := common.HexToAddress("0x456")
	signer := keystest.TxSigner(nil) // returns tx unchanged

	rpc := &testNovaRPC{}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), signer, []common.Address{tier2Addr})

	tx, attempt := newDualBroadcastTx(t, 42)
	signedTx, err := client.resignWithTierValue(context.Background(), tx, attempt, tier2Addr)
	require.NoError(t, err)

	assert.Equal(t, big.NewInt(2), signedTx.Value())
	assert.Equal(t, attempt.SignedTransaction.Nonce(), signedTx.Nonce())
	assert.Equal(t, attempt.SignedTransaction.Gas(), signedTx.Gas())
	assert.Equal(t, attempt.SignedTransaction.GasPrice(), signedTx.GasPrice())
	assert.Equal(t, attempt.SignedTransaction.Data(), signedTx.Data())
	assert.Equal(t, evmtypes.LegacyTxType, int(signedTx.Type()))
}

func TestNovaClient_ResignWithTierValue_DynamicFee(t *testing.T) {
	t.Parallel()

	tier2Addr := common.HexToAddress("0x456")
	signer := keystest.TxSigner(nil)

	rpc := &testNovaRPC{}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), signer, []common.Address{tier2Addr})

	toAddress := common.HexToAddress("0x456")
	accessList := evmtypes.AccessList{
		{Address: common.HexToAddress("0xaaa"), StorageKeys: []common.Hash{common.HexToHash("0x01")}},
	}
	origTx := evmtypes.NewTx(&evmtypes.DynamicFeeTx{
		ChainID:    big.NewInt(1),
		Nonce:      42,
		To:         &toAddress,
		Value:      big.NewInt(0),
		Gas:        21000,
		GasFeeCap:  big.NewInt(100),
		GasTipCap:  big.NewInt(10),
		Data:       []byte{0xde, 0xad},
		AccessList: accessList,
	})

	dualBroadcast := true
	metaJSON, err := json.Marshal(txmtypes.TxMeta{DualBroadcast: &dualBroadcast})
	require.NoError(t, err)
	meta := sqlutil.JSON(metaJSON)
	nonce := uint64(42)

	tx := &txmtypes.Transaction{
		Nonce:       &nonce,
		FromAddress: common.HexToAddress("0x123"),
		ToAddress:   toAddress,
		Meta:        &meta,
	}
	attempt := &txmtypes.Attempt{
		ID:                1,
		SignedTransaction: origTx,
	}

	signedTx, err := client.resignWithTierValue(context.Background(), tx, attempt, tier2Addr)
	require.NoError(t, err)

	assert.Equal(t, big.NewInt(2), signedTx.Value())
	assert.Equal(t, uint64(42), signedTx.Nonce())
	assert.Equal(t, uint64(21000), signedTx.Gas())
	assert.Equal(t, big.NewInt(100), signedTx.GasFeeCap())
	assert.Equal(t, big.NewInt(10), signedTx.GasTipCap())
	assert.Equal(t, big.NewInt(1), signedTx.ChainId())
	assert.Equal(t, []byte{0xde, 0xad}, signedTx.Data())
	assert.Equal(t, accessList, signedTx.AccessList())
	assert.Equal(t, evmtypes.DynamicFeeTxType, int(signedTx.Type()))
}

func TestNovaClient_ResignWithTierValue_DefaultTier(t *testing.T) {
	t.Parallel()

	signer := keystest.TxSigner(nil)

	rpc := &testNovaRPC{}
	customURL, err := url.Parse("http://localhost?api_key=test")
	require.NoError(t, err)

	// No tier2 feeds — everything defaults to tier 1
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), signer, nil)

	tx, attempt := newDualBroadcastTx(t, 42)
	signedTx, err := client.resignWithTierValue(context.Background(), tx, attempt, tx.ToAddress)
	require.NoError(t, err)

	assert.Equal(t, big.NewInt(1), signedTx.Value())
}

func TestNovaClient_SendTransaction_ReSignsTx(t *testing.T) {
	var receivedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		receivedBody, err = io.ReadAll(r.Body)
		assert.NoError(t, err)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xabc"}`))
		assert.NoError(t, err)
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL + "?api_key=test_key")
	require.NoError(t, err)

	tier2Addr := common.HexToAddress("0x456") // matches newDualBroadcastTx toAddress
	signer := keystest.TxSigner(nil)

	rpc := &testNovaRPC{}
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), signer, []common.Address{tier2Addr})

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

	// Decode the sent raw tx and verify Value=2 (tier-2 feed)
	sentRLP, err := hexutil.Decode(req.Params[0])
	require.NoError(t, err)
	var sentTx evmtypes.Transaction
	require.NoError(t, sentTx.UnmarshalBinary(sentRLP))
	assert.Equal(t, big.NewInt(2), sentTx.Value(), "sent tx should have Value=2 for tier-2 feed")
	assert.Equal(t, attempt.SignedTransaction.Nonce(), sentTx.Nonce(), "nonce should be preserved")

	// Verify original attempt was NOT mutated
	assert.Equal(t, big.NewInt(0), attempt.SignedTransaction.Value(),
		"original attempt's signed transaction should keep Value=0")
}

func TestNovaClient_SendTransaction_ForwarderRouted(t *testing.T) {
	var receivedBody []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		receivedBody, err = io.ReadAll(r.Body)
		assert.NoError(t, err)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xabc"}`))
		assert.NoError(t, err)
	}))
	defer server.Close()

	customURL, err := url.Parse(server.URL + "?api_key=test_key")
	require.NoError(t, err)

	feedAddr := common.HexToAddress("0xFEED")
	forwarderAddr := common.HexToAddress("0xFWD0")
	signer := keystest.TxSigner(nil)

	rpc := &testNovaRPC{}
	client := newNovaClient(logger.Test(t), rpc, customURL, testOFAMetrics(t), signer, []common.Address{feedAddr})

	// Build a tx where ToAddress is the forwarder, but the real feed is in meta.FwdrDestAddress
	dualBroadcast := true
	metaJSON, err := json.Marshal(txmtypes.TxMeta{DualBroadcast: &dualBroadcast, FwdrDestAddress: &feedAddr})
	require.NoError(t, err)
	meta := sqlutil.JSON(metaJSON)
	nonce := uint64(42)

	signedTx := evmtypes.NewTx(&evmtypes.LegacyTx{
		Nonce:    nonce,
		To:       &forwarderAddr,
		Gas:      21000,
		GasPrice: big.NewInt(1),
	})

	tx := &txmtypes.Transaction{
		Nonce:       &nonce,
		FromAddress: common.HexToAddress("0x123"),
		ToAddress:   forwarderAddr,
		Meta:        &meta,
	}
	attempt := &txmtypes.Attempt{ID: 1, SignedTransaction: signedTx}

	err = client.SendTransaction(context.Background(), tx, attempt)
	require.NoError(t, err)

	var req struct {
		Method string   `json:"method"`
		Params []string `json:"params"`
	}
	require.NoError(t, json.Unmarshal(receivedBody, &req))

	sentRLP, err := hexutil.Decode(req.Params[0])
	require.NoError(t, err)
	var sentEvmTx evmtypes.Transaction
	require.NoError(t, sentEvmTx.UnmarshalBinary(sentRLP))
	assert.Equal(t, big.NewInt(2), sentEvmTx.Value(),
		"should use Value=2 based on meta.FwdrDestAddress, not the forwarder in tx.ToAddress")
}
