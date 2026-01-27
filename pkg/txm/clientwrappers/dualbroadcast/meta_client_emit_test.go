package dualbroadcast

import (
	"context"
	"errors"
	"math/big"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
	svrv1 "github.com/smartcontractkit/chainlink-protos/svr/v1"
)

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

func TestMetaClient_emitAtlasErrorWithHttpStatusCode(t *testing.T) {
	testChainID := big.NewInt(1)
	testURL, _ := url.Parse("https://atlas.example.com")
	lggr := logger.Test(t)

	t.Run("emits error with all fields populated", func(t *testing.T) {
		mockEmitter := new(mockBeholderEmitter)
		metrics, err := NewMetaMetrics(testChainID.String())
		require.NoError(t, err)

		u, err := url.Parse("https://example.com")
		require.NoError(t, err)

		client := &MetaClient{
			lggr:            logger.Sugared(lggr),
			chainID:         big.NewInt(421614),
			customURL:       u,
			metrics:         metrics,
			beholderEmitter: mockEmitter,
		}

		nonce := uint64(450)

		tx := &types.Transaction{
			ID:          123,
			Nonce:       &nonce,
			FromAddress: common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			ToAddress:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
			// Meta:        nil,
		}

		var capturedBody []byte
		mockEmitter.On("Emit", mock.Anything, mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				capturedBody = args.Get(1).([]byte)
			}).
			Return(nil)

		client.emitAtlasErrorWithHttpStatusCode(t.Context(), "send_request", errors.New("test error message"), 500, tx)

		mockEmitter.AssertExpectations(t)

		// Verify the emitted message
		var emittedMsg svrv1.FastLaneAtlasError
		err = proto.Unmarshal(capturedBody, &emittedMsg)
		require.NoError(t, err)

		assert.Equal(t, testChainID.String(), emittedMsg.ChainId)
		assert.Equal(t, tx.FromAddress.Hex(), emittedMsg.FromAddress)
		assert.Equal(t, tx.ToAddress.Hex(), emittedMsg.ToAddress)
		// assert.Equal(t, fwdrDestAddress.String(), emittedMsg.FeedAddress)
		assert.Equal(t, "42", emittedMsg.Nonce)
		assert.Equal(t, "test_error_type", emittedMsg.ErrorType)
		assert.Equal(t, "test error message", emittedMsg.ErrorMessage)
		assert.Equal(t, int32(500), emittedMsg.HttpStatusCode)
		assert.Equal(t, int64(123), emittedMsg.TransactionId)
		assert.Equal(t, testURL.String(), emittedMsg.AtlasUrl)
	})

	t.Run("emits error with nil nonce", func(t *testing.T) {
		mockEmitter := new(mockBeholderEmitter)
		metrics, err := NewMetaMetrics(testChainID.String())
		require.NoError(t, err)

		client := &MetaClient{
			lggr:            logger.Sugared(lggr),
			chainID:         testChainID,
			customURL:       testURL,
			metrics:         metrics,
			beholderEmitter: mockEmitter,
		}

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

		client.emitAtlasErrorWithHttpStatusCode(t.Context(), "error_type", errors.New("some error"), 400, tx)

		mockEmitter.AssertExpectations(t)

		var emittedMsg svrv1.FastLaneAtlasError
		err = proto.Unmarshal(capturedBody, &emittedMsg)
		require.NoError(t, err)

		assert.Equal(t, "", emittedMsg.Nonce)       // empty string when nonce is nil
		assert.Equal(t, "", emittedMsg.FeedAddress) // empty string when meta is nil
	})

	t.Run("emits error with negative http status code for non-http errors", func(t *testing.T) {
		mockEmitter := new(mockBeholderEmitter)
		metrics, err := NewMetaMetrics(testChainID.String())
		require.NoError(t, err)

		client := &MetaClient{
			lggr:            logger.Sugared(lggr),
			chainID:         testChainID,
			customURL:       testURL,
			metrics:         metrics,
			beholderEmitter: mockEmitter,
		}

		tx := &types.Transaction{
			ID:          789,
			FromAddress: common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			ToAddress:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
		}

		var capturedBody []byte
		mockEmitter.On("Emit", mock.Anything, mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				capturedBody = args.Get(1).([]byte)
			}).
			Return(nil)

		client.emitAtlasErrorWithHttpStatusCode(t.Context(), "non_http_error", errors.New("network error"), -1, tx)

		mockEmitter.AssertExpectations(t)

		var emittedMsg svrv1.FastLaneAtlasError
		err = proto.Unmarshal(capturedBody, &emittedMsg)
		require.NoError(t, err)

		assert.Equal(t, int32(-1), emittedMsg.HttpStatusCode)
	})

	t.Run("handles emit error gracefully", func(t *testing.T) {
		mockEmitter := new(mockBeholderEmitter)
		metrics, err := NewMetaMetrics(testChainID.String())
		require.NoError(t, err)

		client := &MetaClient{
			lggr:            logger.Sugared(lggr),
			chainID:         testChainID,
			customURL:       testURL,
			metrics:         metrics,
			beholderEmitter: mockEmitter,
		}

		tx := &types.Transaction{
			ID:          999,
			FromAddress: common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			ToAddress:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
		}

		mockEmitter.On("Emit", mock.Anything, mock.Anything, mock.Anything).
			Return(errors.New("emit failed"))

		// Should not panic, just log the error
		client.emitAtlasErrorWithHttpStatusCode(t.Context(), "error_type", errors.New("some error"), 500, tx)

		mockEmitter.AssertExpectations(t)
	})

	t.Run("handles invalid meta JSON gracefully", func(t *testing.T) {
		mockEmitter := new(mockBeholderEmitter)
		metrics, err := NewMetaMetrics(testChainID.String())
		require.NoError(t, err)

		client := &MetaClient{
			lggr:            logger.Sugared(lggr),
			chainID:         testChainID,
			customURL:       testURL,
			metrics:         metrics,
			beholderEmitter: mockEmitter,
		}

		invalidJSON := sqlutil.JSON([]byte("invalid json"))
		tx := &types.Transaction{
			ID:          111,
			FromAddress: common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
			ToAddress:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
			Meta:        &invalidJSON,
		}

		// Should not call Emit because GetMeta will fail
		client.emitAtlasErrorWithHttpStatusCode(t.Context(), "error_type", errors.New("some error"), 500, tx)

		// Emit should not be called when meta parsing fails
		mockEmitter.AssertNotCalled(t, "Emit", mock.Anything, mock.Anything, mock.Anything)
	})
}
