package dualbroadcast

import (
	"context"
	"errors"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	evmtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
	txmtypes "github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

type chainClientMock struct {
	sendTxCalls atomic.Int32
}

func (c *chainClientMock) BlockByNumber(context.Context, *big.Int) (*evmtypes.Block, error) {
	return nil, errors.New("unexpected BlockByNumber")
}
func (c *chainClientMock) NonceAt(context.Context, common.Address, *big.Int) (uint64, error) {
	return 0, errors.New("unexpected NonceAt")
}
func (c *chainClientMock) SendTransaction(context.Context, *txmtypes.Transaction, *txmtypes.Attempt) error {
	c.sendTxCalls.Add(1)
	return nil
}

var _ chainRPCClient = (*chainClientMock)(nil)

func createMultiOfaClient(t *testing.T, c chainRPCClient, primary multiOfaBackend, secondaries ...multiOfaBackend) *multiOfaClient {
	t.Helper()
	return &multiOfaClient{
		lggr:                 logger.Sugared(logger.Test(t)),
		chainClient:          c,
		primary:              primary,
		secondaries:          secondaries,
		secondarySendTimeout: rpcTimeout,
	}
}

type ofaBackendMock struct {
	sendTxFn       func(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error
	pendingNonceFn func(ctx context.Context, address common.Address) (uint64, error)
	nonceAtFn      func(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error)
	sendCalled     chan struct{}
}

func (m *ofaBackendMock) SendTransaction(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error {
	if m.sendCalled != nil {
		defer func() { m.sendCalled <- struct{}{} }()
	}
	if m.sendTxFn != nil {
		return m.sendTxFn(ctx, tx, attempt)
	}
	return nil
}

func (m *ofaBackendMock) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	if m.pendingNonceFn != nil {
		return m.pendingNonceFn(ctx, address)
	}
	return 0, nil
}

func (m *ofaBackendMock) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	if m.nonceAtFn != nil {
		return m.nonceAtFn(ctx, address, blockNumber)
	}
	return 0, nil
}

var _ multiOfaBackend = (*ofaBackendMock)(nil)

func TestMultiOfaClient_NonDual_RoutesOnlyToChainClient(t *testing.T) {
	chainClient := &chainClientMock{}

	primary := &ofaBackendMock{
		sendTxFn: func(context.Context, *txmtypes.Transaction, *txmtypes.Attempt) error {
			require.FailNow(t, "primary must not receive SendTransaction when not dual-broadcasting")
			return nil
		},
	}
	secCalled := make(chan struct{}, 1)
	secondary := &ofaBackendMock{
		sendTxFn: func(context.Context, *txmtypes.Transaction, *txmtypes.Attempt) error {
			secCalled <- struct{}{}
			return nil
		},
	}

	mc := createMultiOfaClient(t, chainClient, primary, secondary)

	nonce := uint64(1)
	to := common.HexToAddress("0x456")
	signedTx := evmtypes.NewTx(&evmtypes.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Gas:      21000,
		GasPrice: big.NewInt(1),
	})
	tx := &txmtypes.Transaction{
		Nonce:       &nonce,
		FromAddress: common.HexToAddress("0x123"),
		ToAddress:   to,
		Meta:        nil,
	}
	attempt := &txmtypes.Attempt{SignedTransaction: signedTx}

	err := mc.SendTransaction(testutils.Context(t), tx, attempt)
	require.NoError(t, err)
	require.Equal(t, int32(1), chainClient.sendTxCalls.Load(), "non-dual path must use chainClient only")

	select {
	case <-secCalled:
		require.FailNow(t, "secondaries must not run for non-dual-broadcast")
	default:
	}
}

func TestMultiOfaClient_SendTransaction_TwoSecondaries(t *testing.T) {
	chainClient := &chainClientMock{}
	sec1 := make(chan struct{}, 1)
	sec2 := make(chan struct{}, 1)
	primary := &ofaBackendMock{}
	mc := createMultiOfaClient(t, chainClient, primary, &ofaBackendMock{sendCalled: sec1}, &ofaBackendMock{sendCalled: sec2})
	tx, attempt := newDualBroadcastTx(t, 1)
	err := mc.SendTransaction(testutils.Context(t), tx, attempt)
	require.NoError(t, err)

	select {
	case <-sec1:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary 1 was not called")
	}
	select {
	case <-sec2:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary 2 was not called")
	}
}

func TestMultiOfaClient_SendTransaction_BothSucceed(t *testing.T) {
	chainClient := &chainClientMock{}
	secondaryCalled := make(chan struct{}, 1)
	primary := &ofaBackendMock{}
	secondary := &ofaBackendMock{sendCalled: secondaryCalled}

	mc := createMultiOfaClient(t, chainClient, primary, secondary)
	tx, attempt := newDualBroadcastTx(t, 1)
	err := mc.SendTransaction(testutils.Context(t), tx, attempt)
	require.NoError(t, err)

	select {
	case <-secondaryCalled:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary was not called")
	}
}

// TestMultiOfaClient_SecondarySend_DoesNotBlockReturn ensures the primary outcome is not delayed
// by slow secondaries (fan-out does not synchronize on secondary completion).
func TestMultiOfaClient_SecondarySend_DoesNotBlockReturn(t *testing.T) {
	releaseSecondary := make(chan struct{})
	chainClient := &chainClientMock{}
	primary := &ofaBackendMock{
		sendTxFn: func(context.Context, *txmtypes.Transaction, *txmtypes.Attempt) error {
			return nil
		},
	}
	secondary := &ofaBackendMock{
		sendTxFn: func(context.Context, *txmtypes.Transaction, *txmtypes.Attempt) error {
			<-releaseSecondary
			return nil
		},
	}
	mc := createMultiOfaClient(t, chainClient, primary, secondary)
	tx, attempt := newDualBroadcastTx(t, 1)

	done := make(chan error, 1)
	go func() {
		done <- mc.SendTransaction(testutils.Context(t), tx, attempt)
	}()

	select {
	case err := <-done:
		require.NoError(t, err)
	case <-time.After(2 * time.Second):
		t.Fatal("SendTransaction should return when primary completes, without waiting for secondary")
	}

	close(releaseSecondary)
}

func TestMultiOfaClient_SendTransaction_PrimaryFails(t *testing.T) {
	chainClient := &chainClientMock{}
	primaryErr := errors.New("flashbots rejected")
	primary := &ofaBackendMock{
		sendTxFn: func(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error {
			return primaryErr
		},
	}
	secondary := &ofaBackendMock{sendCalled: make(chan struct{}, 1)}

	mc := createMultiOfaClient(t, chainClient, primary, secondary)
	tx, attempt := newDualBroadcastTx(t, 1)
	err := mc.SendTransaction(testutils.Context(t), tx, attempt)
	require.ErrorIs(t, err, primaryErr)
}

func TestMultiOfaClient_SecondarySendRespectsTimeout(t *testing.T) {
	chainClient := &chainClientMock{}
	secondaryDone := make(chan struct{}, 1)
	primary := &ofaBackendMock{}
	secondary := &ofaBackendMock{
		sendTxFn: func(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error {
			<-ctx.Done()
			secondaryDone <- struct{}{}
			return ctx.Err()
		},
		sendCalled: make(chan struct{}, 1),
	}

	mc := createMultiOfaClient(t, chainClient, primary, secondary)
	mc.secondarySendTimeout = 150 * time.Millisecond
	tx, attempt := newDualBroadcastTx(t, 1)
	err := mc.SendTransaction(testutils.Context(t), tx, attempt)
	require.NoError(t, err)

	select {
	case <-secondaryDone:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary goroutine did not finish after context timeout")
	}
	select {
	case <-secondary.sendCalled:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary SendTransaction did not return")
	}
}

func TestMultiOfaClient_SendTransaction_SecondaryFails(t *testing.T) {
	chainClient := &chainClientMock{}
	secondaryCalled := make(chan struct{}, 1)
	primary := &ofaBackendMock{}
	secondary := &ofaBackendMock{
		sendTxFn: func(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error {
			return errors.New("nova failed")
		},
		sendCalled: secondaryCalled,
	}

	mc := createMultiOfaClient(t, chainClient, primary, secondary)
	tx, attempt := newDualBroadcastTx(t, 1)
	err := mc.SendTransaction(testutils.Context(t), tx, attempt)
	require.NoError(t, err)

	select {
	case <-secondaryCalled:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary was not called")
	}
}

func TestMultiOfaClient_PendingNonceAt_RoutesToPrimary(t *testing.T) {
	chainClient := &chainClientMock{}
	primary := &ofaBackendMock{
		pendingNonceFn: func(ctx context.Context, address common.Address) (uint64, error) {
			return 42, nil
		},
	}
	secondary := &ofaBackendMock{
		pendingNonceFn: func(ctx context.Context, address common.Address) (uint64, error) {
			t.Fatal("secondary PendingNonceAt should not be called")
			return 0, nil
		},
	}

	mc := createMultiOfaClient(t, chainClient, primary, secondary)
	nonce, err := mc.PendingNonceAt(testutils.Context(t), common.HexToAddress("0x123"))
	require.NoError(t, err)
	assert.Equal(t, uint64(42), nonce)
}

func TestMultiOfaClient_NonceAt_RoutesToPrimary(t *testing.T) {
	chainClient := &chainClientMock{}
	primary := &ofaBackendMock{
		nonceAtFn: func(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
			return 99, nil
		},
	}
	secondary := &ofaBackendMock{
		nonceAtFn: func(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
			t.Fatal("secondary NonceAt should not be called")
			return 0, nil
		},
	}

	mc := createMultiOfaClient(t, chainClient, primary, secondary)
	nonce, err := mc.NonceAt(testutils.Context(t), common.HexToAddress("0x123"), big.NewInt(100))
	require.NoError(t, err)
	assert.Equal(t, uint64(99), nonce)
}

func TestMultiOfaClient_FromOFAURLs_HTTPServers_DualBroadcast(t *testing.T) {
	var primaryHits, secondaryHits atomic.Int32

	primarySrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		primaryHits.Add(1)
		_, _ = io.Copy(io.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xabc"}`))
	}))
	defer primarySrv.Close()

	secondarySrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secondaryHits.Add(1)
		_, _ = io.Copy(io.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xdef"}`))
	}))
	defer secondarySrv.Close()

	uPrimary, err := url.Parse(primarySrv.URL + "/relay.flashbots.net")
	require.NoError(t, err)
	uSecondary, err := url.Parse(secondarySrv.URL + "/novarpc")
	require.NoError(t, err)

	mockEth := clienttest.NewClient(t)
	mockEth.EXPECT().ConfiguredChainID().Return(big.NewInt(1)).Maybe()

	cc, err := clientwrappers.NewChainClient(logger.Test(t), mockEth, false)
	require.NoError(t, err)

	mux, err := newMultiOfaClient(
		logger.Test(t),
		cc,
		nil,
		[]*url.URL{uPrimary, uSecondary},
		big.NewInt(1),
		nil,
		nil,
	)
	require.NoError(t, err)

	tx, attempt := newDualBroadcastTx(t, 7)
	err = mux.SendTransaction(testutils.Context(t), tx, attempt)
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		return primaryHits.Load() == 1 && secondaryHits.Load() == 1
	}, 2*time.Second, 5*time.Millisecond, "both OFA backends should receive eth_sendRawTransaction")

	// Dual-broadcast must not use the public mempool path
	mockEth.AssertNotCalled(t, "SendTransaction", mock.Anything, mock.Anything)
}

func TestMultiOfaClient_NonDual_NovaPrimary_RoutesToMempool(t *testing.T) {
	var relayHits atomic.Int32
	relaySrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		relayHits.Add(1)
		_, _ = io.Copy(io.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x"}`))
	}))
	defer relaySrv.Close()

	u, err := url.Parse(relaySrv.URL + "/novarpc")
	require.NoError(t, err)

	mockEth := clienttest.NewClient(t)
	mockEth.EXPECT().ConfiguredChainID().Return(big.NewInt(1)).Maybe()
	mockEth.EXPECT().SendTransaction(mock.Anything, mock.Anything).Return(nil).Times(1)

	cc, err := clientwrappers.NewChainClient(logger.Test(t), mockEth, false)
	require.NoError(t, err)

	mux, err := newMultiOfaClient(logger.Test(t), cc, nil, []*url.URL{u}, big.NewInt(1), nil, nil)
	require.NoError(t, err)

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

	err = mux.SendTransaction(testutils.Context(t), tx, attempt)
	require.NoError(t, err)
	require.Equal(t, int32(0), relayHits.Load(), "Nova relay must not receive sends when not dual-broadcasting")
}

func TestMultiOfaClient_NonDual_FlashbotsPrimaryNovaSecondary_NoOFAHTTPHits(t *testing.T) {
	var primaryHits, secondaryHits atomic.Int32
	primarySrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		primaryHits.Add(1)
		_, _ = io.Copy(io.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xabc"}`))
	}))
	defer primarySrv.Close()
	secondarySrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secondaryHits.Add(1)
		_, _ = io.Copy(io.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0xdef"}`))
	}))
	defer secondarySrv.Close()

	uPrimary, err := url.Parse(primarySrv.URL + "/relay.flashbots.net")
	require.NoError(t, err)
	uSecondary, err := url.Parse(secondarySrv.URL + "/novarpc")
	require.NoError(t, err)

	mockEth := clienttest.NewClient(t)
	mockEth.EXPECT().ConfiguredChainID().Return(big.NewInt(1)).Maybe()
	mockEth.EXPECT().SendTransaction(mock.Anything, mock.Anything).Return(nil).Times(1)

	cc, err := clientwrappers.NewChainClient(logger.Test(t), mockEth, false)
	require.NoError(t, err)

	mux, err := newMultiOfaClient(
		logger.Test(t),
		cc,
		nil,
		[]*url.URL{uPrimary, uSecondary},
		big.NewInt(1),
		nil,
		nil,
	)
	require.NoError(t, err)

	nonce := uint64(3)
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

	err = mux.SendTransaction(testutils.Context(t), tx, attempt)
	require.NoError(t, err)
	require.Equal(t, int32(0), primaryHits.Load())
	require.Equal(t, int32(0), secondaryHits.Load())
}

func TestMultiOfaClient_Purgeable_NovaPrimary_RoutesToMempool(t *testing.T) {
	var relayHits atomic.Int32
	relaySrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		relayHits.Add(1)
		_, _ = io.Copy(io.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x"}`))
	}))
	defer relaySrv.Close()

	u, err := url.Parse(relaySrv.URL + "/novarpc")
	require.NoError(t, err)

	mockEth := clienttest.NewClient(t)
	mockEth.EXPECT().ConfiguredChainID().Return(big.NewInt(1)).Maybe()
	mockEth.EXPECT().SendTransaction(mock.Anything, mock.Anything).Return(nil).Times(1)

	cc, err := clientwrappers.NewChainClient(logger.Test(t), mockEth, false)
	require.NoError(t, err)

	mux, err := newMultiOfaClient(logger.Test(t), cc, nil, []*url.URL{u}, big.NewInt(1), nil, nil)
	require.NoError(t, err)

	tx, attempt := newDualBroadcastTx(t, 1)
	tx.IsPurgeable = true

	err = mux.SendTransaction(testutils.Context(t), tx, attempt)
	require.NoError(t, err)
	require.Equal(t, int32(0), relayHits.Load(), "purgeable txs must not hit the Nova relay")
}

func TestRedactURL(t *testing.T) {
	t.Parallel()

	assert.Empty(t, redactURL(nil))

	u := mustParseURL(t, "https://eth.novarpc.xyz?api_key=secret&foo=bar")
	assert.Equal(t, "https://eth.novarpc.xyz?api_key=xxxxx&foo=bar", redactURL(u))
	assert.Equal(t, "secret", u.Query().Get("api_key"), "must not mutate original URL")

	uPass := mustParseURL(t, "https://user:pass@eth.novarpc.xyz?api_key=secret")
	assert.Equal(t, "https://user:xxxxx@eth.novarpc.xyz?api_key=xxxxx", redactURL(uPass))
}
