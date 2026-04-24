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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/clientwrappers"
	txmtypes "github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
)

func createMultiOfaClient(t *testing.T, primary ofaBackend, secondaries ...ofaBackend) *multiOfaClient {
	return &multiOfaClient{
		lggr:                 logger.Sugared(logger.Test(t)),
		primary:              primary,
		secondaries:          secondaries,
		secondarySendTimeout: rpcTimeout,
	}
}

type mockClient struct {
	label          string
	sendTxFn       func(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error
	pendingNonceFn func(ctx context.Context, address common.Address) (uint64, error)
	nonceAtFn      func(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error)
	sendCalled     chan struct{}
}

func (m *mockClient) SendTransaction(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error {
	if m.sendCalled != nil {
		defer func() { m.sendCalled <- struct{}{} }()
	}
	if m.sendTxFn != nil {
		return m.sendTxFn(ctx, tx, attempt)
	}
	return nil
}

func (m *mockClient) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	if m.pendingNonceFn != nil {
		return m.pendingNonceFn(ctx, address)
	}
	return 0, nil
}

func (m *mockClient) NonceAt(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
	if m.nonceAtFn != nil {
		return m.nonceAtFn(ctx, address, blockNumber)
	}
	return 0, nil
}

func (m *mockClient) Label() string {
	return m.label
}

func TestMultiOfaClient_SendTransaction_TwoSecondaries(t *testing.T) {
	sec1 := make(chan struct{}, 1)
	sec2 := make(chan struct{}, 1)
	primary := &mockClient{label: "primary"}
	mc := createMultiOfaClient(t, primary, &mockClient{label: "secondary1", sendCalled: sec1}, &mockClient{label: "secondary2", sendCalled: sec2})
	err := mc.SendTransaction(context.Background(), &txmtypes.Transaction{}, &txmtypes.Attempt{})
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
	secondaryCalled := make(chan struct{}, 1)
	primary := &mockClient{label: "primary"}
	secondary := &mockClient{label: "secondary", sendCalled: secondaryCalled}

	mc := createMultiOfaClient(t, primary, secondary)
	err := mc.SendTransaction(context.Background(), &txmtypes.Transaction{}, &txmtypes.Attempt{})
	require.NoError(t, err)

	select {
	case <-secondaryCalled:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary was not called")
	}
}

func TestMultiOfaClient_SendTransaction_PrimaryFails(t *testing.T) {
	primaryErr := errors.New("flashbots rejected")
	primary := &mockClient{
		sendTxFn: func(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error {
			return primaryErr
		},
		label: "primary",
	}
	secondary := &mockClient{label: "secondary", sendCalled: make(chan struct{}, 1)}

	mc := createMultiOfaClient(t, primary, secondary)
	err := mc.SendTransaction(context.Background(), &txmtypes.Transaction{}, &txmtypes.Attempt{})
	require.ErrorIs(t, err, primaryErr)
}

func TestMultiOfaClient_SecondarySendRespectsTimeout(t *testing.T) {
	secondaryDone := make(chan struct{}, 1)
	primary := &mockClient{label: "primary"}
	secondary := &mockClient{
		sendTxFn: func(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error {
			<-ctx.Done()
			secondaryDone <- struct{}{}
			return ctx.Err()
		},
		sendCalled: make(chan struct{}, 1),
	}

	mc := createMultiOfaClient(t, primary, secondary)
	mc.secondarySendTimeout = 150 * time.Millisecond
	err := mc.SendTransaction(context.Background(), &txmtypes.Transaction{}, &txmtypes.Attempt{})
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
	secondaryCalled := make(chan struct{}, 1)
	primary := &mockClient{label: "primary"}
	secondary := &mockClient{
		sendTxFn: func(ctx context.Context, tx *txmtypes.Transaction, attempt *txmtypes.Attempt) error {
			return errors.New("nova failed")
		},
		sendCalled: secondaryCalled,
	}

	mc := createMultiOfaClient(t, primary, secondary)
	err := mc.SendTransaction(context.Background(), &txmtypes.Transaction{}, &txmtypes.Attempt{})
	require.NoError(t, err)

	select {
	case <-secondaryCalled:
	case <-time.After(2 * time.Second):
		t.Fatal("secondary was not called")
	}
}

func TestMultiOfaClient_PendingNonceAt_RoutesToPrimary(t *testing.T) {
	primary := &mockClient{label: "primary",
		pendingNonceFn: func(ctx context.Context, address common.Address) (uint64, error) {
			return 42, nil
		},
	}
	secondary := &mockClient{label: "secondary",
		pendingNonceFn: func(ctx context.Context, address common.Address) (uint64, error) {
			t.Fatal("secondary PendingNonceAt should not be called")
			return 0, nil
		},
	}

	mc := createMultiOfaClient(t, primary, secondary)
	nonce, err := mc.PendingNonceAt(context.Background(), common.HexToAddress("0x123"))
	require.NoError(t, err)
	assert.Equal(t, uint64(42), nonce)
}

func TestMultiOfaClient_NonceAt_RoutesToPrimary(t *testing.T) {
	primary := &mockClient{label: "primary",
		nonceAtFn: func(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
			return 99, nil
		},
	}
	secondary := &mockClient{label: "secondary",
		nonceAtFn: func(ctx context.Context, address common.Address, blockNumber *big.Int) (uint64, error) {
			t.Fatal("secondary NonceAt should not be called")
			return 0, nil
		},
	}

	mc := createMultiOfaClient(t, primary, secondary)
	nonce, err := mc.NonceAt(context.Background(), common.HexToAddress("0x123"), big.NewInt(100))
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

	mux, eh, err := newMultiOfaClient(
		logger.Test(t),
		cc,
		nil,
		[]*url.URL{uPrimary, uSecondary},
		big.NewInt(1),
		nil,
		nil,
		nil,
	)
	require.NoError(t, err)
	require.Nil(t, eh)

	tx, attempt := newDualBroadcastTx(t, 7)
	err = mux.SendTransaction(context.Background(), tx, attempt)
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		return primaryHits.Load() == 1 && secondaryHits.Load() == 1
	}, 2*time.Second, 5*time.Millisecond, "both OFA backends should receive eth_sendRawTransaction")
}
