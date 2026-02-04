package dualbroadcast

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/url"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
	pb "github.com/smartcontractkit/chainlink-protos/svr/v1"
)

func TestMetaMetrics(t *testing.T) {
	t.Parallel()
	chainID := "1"

	t.Run("NewMetaMetrics", func(t *testing.T) {
		t.Parallel()
		metrics, err := NewMetaMetrics(chainID, logger.Test(t))
		require.NoError(t, err)
		assert.NotNil(t, metrics)
		assert.Equal(t, chainID, metrics.chainID)
	})

	t.Run("RecordBasicMetrics", func(t *testing.T) {
		t.Parallel()
		metrics, err := NewMetaMetrics(chainID, logger.Test(t))
		require.NoError(t, err)

		ctx := t.Context()

		// Test that these don't panic - all metrics methods
		metrics.RecordStatusCode(ctx, 200)
		metrics.RecordLatency(ctx, time.Millisecond*100)
		metrics.RecordBidsReceived(ctx, 5)
		metrics.RecordSendRequestError(ctx)
		metrics.RecordSendOperationError(ctx)
	})
}

// mockBeholderEmitter is a mock for beholder.Emitter
type mockBeholderEmitter struct {
	mock.Mock
}

func (m *mockBeholderEmitter) Emit(ctx context.Context, body []byte, attrKVs ...any) error {
	args := m.Called(ctx, body, attrKVs)
	return args.Error(0)
}

func (m *mockBeholderEmitter) Close() error {
	args := m.Called()
	return args.Error(0)
}

func TestMetaClient_emitAtlasError(t *testing.T) {
	t.Parallel()
	testChainID := big.NewInt(1)
	testURL, _ := url.Parse("https://atlas.example.com")
	lggr := logger.Test(t)

	t.Run("emits error with all fields populated", func(t *testing.T) {
		t.Parallel()
		mockEmitter := new(mockBeholderEmitter)
		metrics, err := NewMetaMetrics(testChainID.String(), lggr)
		require.NoError(t, err)
		metrics.emitter = mockEmitter

		u, err := url.Parse("https://example.com")
		require.NoError(t, err)

		nonce := uint64(450)
		fwdrDestAddress := common.HexToAddress("0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")
		metaBytes, err := json.Marshal(types.TxMeta{FwdrDestAddress: &fwdrDestAddress})
		require.NoError(t, err)
		meta := sqlutil.JSON(metaBytes)

		tx := &types.Transaction{
			ID:          123,
			Nonce:       &nonce,
			FromAddress: common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			ToAddress:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
			Meta:        &meta,
		}

		var capturedBody []byte
		mockEmitter.On("Emit", mock.Anything, mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				capturedBody = args.Get(1).([]byte)
			}).
			Return(nil)

		metrics.emitAtlasError(t.Context(), "send_request", u, errors.New("test error message"), tx)

		mockEmitter.AssertExpectations(t)

		// Verify the emitted message
		var emittedMsg pb.FastLaneAtlasError
		err = proto.Unmarshal(capturedBody, &emittedMsg)
		require.NoError(t, err)

		assert.Equal(t, testChainID.String(), emittedMsg.ChainId)
		assert.Equal(t, tx.FromAddress.Hex(), emittedMsg.FromAddress)
		assert.Equal(t, tx.ToAddress.Hex(), emittedMsg.ToAddress)
		assert.Equal(t, fwdrDestAddress.String(), emittedMsg.FeedAddress)
		assert.Equal(t, "450", emittedMsg.Nonce)
		assert.Equal(t, "send_request", emittedMsg.ErrorType)
		assert.Equal(t, "test error message", emittedMsg.ErrorMessage)
		assert.Equal(t, uint64(123), emittedMsg.TransactionId)
		assert.Equal(t, u.String(), emittedMsg.AtlasUrl)
	})

	t.Run("emits error with nil nonce", func(t *testing.T) {
		t.Parallel()
		mockEmitter := new(mockBeholderEmitter)
		metrics, err := NewMetaMetrics(testChainID.String(), lggr)
		require.NoError(t, err)
		metrics.emitter = mockEmitter

		tx := &types.Transaction{
			ID:          456,
			Nonce:       nil, // nil nonce
			FromAddress: common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			ToAddress:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
			Meta:        nil,
		}

		var capturedBody []byte
		mockEmitter.On("Emit", mock.Anything, mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				capturedBody = args.Get(1).([]byte)
			}).
			Return(nil)

		metrics.emitAtlasError(t.Context(), "error_type", testURL, errors.New("some error"), tx)

		mockEmitter.AssertExpectations(t)

		var emittedMsg pb.FastLaneAtlasError
		err = proto.Unmarshal(capturedBody, &emittedMsg)
		require.NoError(t, err)

		assert.Equal(t, "", emittedMsg.Nonce)       // empty string when nonce is nil
		assert.Equal(t, "", emittedMsg.FeedAddress) // empty string when meta is nil
	})

	t.Run("handles emit error gracefully", func(t *testing.T) {
		t.Parallel()
		mockEmitter := new(mockBeholderEmitter)
		metrics, err := NewMetaMetrics(testChainID.String(), lggr)
		require.NoError(t, err)
		metrics.emitter = mockEmitter

		tx := &types.Transaction{
			ID:          999,
			FromAddress: common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			ToAddress:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
		}

		mockEmitter.On("Emit", mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("emit failed"))

		// Should not panic, just log the error
		metrics.emitAtlasError(t.Context(), "error_type", testURL, errors.New("some error"), tx)

		mockEmitter.AssertExpectations(t)
	})

	t.Run("handles invalid meta JSON gracefully", func(t *testing.T) {
		t.Parallel()
		mockEmitter := new(mockBeholderEmitter)
		metrics, err := NewMetaMetrics(testChainID.String(), lggr)
		require.NoError(t, err)
		metrics.emitter = mockEmitter

		invalidJSON := sqlutil.JSON([]byte("invalid json"))
		tx := &types.Transaction{
			ID:          111,
			FromAddress: common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			ToAddress:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
			Meta:        &invalidJSON,
		}

		// Should not call Emit because GetMeta will fail
		metrics.emitAtlasError(t.Context(), "error_type", testURL, errors.New("some error"), tx)

		// Emit should not be called when meta parsing fails
		mockEmitter.AssertNotCalled(t, "Emit", mock.Anything, mock.Anything, mock.Anything)
	})
}
