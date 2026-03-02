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

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/keys/keystest"
	txmtypes "github.com/smartcontractkit/chainlink-evm/pkg/txm/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testFlashbotsRPC struct {
	block *evmtypes.Block
}

func (m *testFlashbotsRPC) BlockByNumber(context.Context, *big.Int) (*evmtypes.Block, error) {
	return m.block, nil
}

func (m *testFlashbotsRPC) NonceAt(context.Context, common.Address, *big.Int) (uint64, error) {
	return 0, nil
}

func (m *testFlashbotsRPC) SendTransaction(context.Context, *evmtypes.Transaction) error {
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
		wantPrivacy    Privacy
		wantRefund     RefundConfig
		wantErr        bool
		wantErrContain string
	}{
		{
			name:        "empty params",
			params:      "",
			wantPrivacy: Privacy{},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "auctionTimeout",
			params:      "auctionTimeout=60",
			wantPrivacy: Privacy{AuctionTimeout: 60},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "auctionTimeout invalid ignored",
			params:      "auctionTimeout=notanint",
			wantPrivacy: Privacy{},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "single builder",
			params:      "builder=test_builder",
			wantPrivacy: Privacy{Builders: []string{"test_builder"}},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "multiple builders",
			params:      "builder=test_builder_1&builder=test_builder_2",
			wantPrivacy: Privacy{Builders: []string{"test_builder_1", "test_builder_2"}},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "single hint",
			params:      "hint=calldata",
			wantPrivacy: Privacy{Hints: []string{"calldata"}},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "multiple hints",
			params:      "hint=calldata&hint=hash",
			wantPrivacy: Privacy{Hints: []string{"calldata", "hash"}},
			wantRefund:  RefundConfig{},
		},
		{
			name:        "refund valid",
			params:      "refund=0xRefundAddr:50",
			wantPrivacy: Privacy{WantRefund: 50},
			wantRefund:  RefundConfig{Address: "0xRefundAddr", Percent: 100},
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
			wantPrivacy: Privacy{
				AuctionTimeout: 120,
				Builders:       []string{"test_builder_1", "test_builder_2"},
				Hints:          []string{"h1"},
				WantRefund:     75,
			},
			wantRefund: RefundConfig{Address: "0xR", Percent: 100},
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
	client := NewFlashbotsClient(logger.Test(t), rpc, keystest.MessageSigner(nil), customURL, txStore, nil)

	err = client.SendBundle(context.Background(), fromAddress, "")
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
	client := NewFlashbotsClient(logger.Test(t), rpc, keystest.MessageSigner(nil), customURL, txStore, nil)
	err = client.SendBundle(context.Background(), fromAddress, "")
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

	client := NewFlashbotsClient(logger.Test(t), &testFlashbotsRPC{}, keystest.MessageSigner(nil), customURL, txStore, nil)
	err = client.SendBundle(context.Background(), fromAddress, "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "must be contiguous and strictly increasing")
}
