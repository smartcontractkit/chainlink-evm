package txm

import (
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder/beholdertest"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
	svrv1 "github.com/smartcontractkit/chainlink-protos/svr/v1"

	"google.golang.org/protobuf/proto"
)

func TestEmitTxMessage(t *testing.T) {
	t.Run("overrides 0x0 as ToAddress if tx is purgeable", func(t *testing.T) {
		// GIVEN
		ctx := t.Context()
		beholderTester := beholdertest.NewObserver(t)

		toAddress := testutils.NewAddress()
		fromAddress := testutils.NewAddress()

		expectedFromAddress := fromAddress
		expectedToAddress := common.Address{}
		expectedHash := common.Hash{}
		expectedChain := testutils.FixtureChainID
		expectedNonce := uint64(256)
		var actualMessage svrv1.TxMessage

		txmMetrics, err := NewTxmMetrics(expectedChain)
		require.NoError(t, err)

		tx := &types.Transaction{
			IsPurgeable: true,
			FromAddress: fromAddress,
			ToAddress:   toAddress,
			Nonce:       &expectedNonce,
		}

		// WHEN
		err = txmMetrics.EmitTxMessage(
			ctx,
			expectedHash,
			fromAddress,
			tx,
		)
		require.NoError(t, err)

		// THEN
		messages := beholderTester.Messages(t)

		assert.Len(t, messages, 1)
		actualMessageBody := messages[0]
		err = proto.Unmarshal(actualMessageBody.Body, &actualMessage)
		require.NoError(t, err)

		assert.Equal(t, expectedFromAddress.String(), actualMessage.FromAddress)
		assert.Equal(t, expectedToAddress.String(), actualMessage.ToAddress)
		assert.Equal(t, strconv.FormatUint(expectedNonce, 10), actualMessage.Nonce)
		assert.Equal(t, expectedChain.String(), actualMessage.ChainId)
		assert.Equal(t, "", actualMessage.FeedAddress)
	})

	t.Run("sends original ToAddress if tx is not purgeable", func(t *testing.T) {
		// GIVEN
		ctx := t.Context()
		beholderTester := beholdertest.NewObserver(t)

		toAddress := testutils.NewAddress()
		fromAddress := testutils.NewAddress()

		expectedFromAddress := fromAddress
		expectedToAddress := toAddress
		expectedHash := common.Hash{}
		expectedChain := testutils.FixtureChainID
		expectedNonce := uint64(256)
		var actualMessage svrv1.TxMessage

		txmMetrics, err := NewTxmMetrics(expectedChain)
		require.NoError(t, err)

		tx := &types.Transaction{
			IsPurgeable: false,
			FromAddress: fromAddress,
			ToAddress:   toAddress,
			Nonce:       &expectedNonce,
		}

		// WHEN
		err = txmMetrics.EmitTxMessage(
			ctx,
			expectedHash,
			fromAddress,
			tx,
		)
		require.NoError(t, err)

		// THEN
		messages := beholderTester.Messages(t)

		assert.Len(t, messages, 1)
		actualMessageBody := messages[0]
		err = proto.Unmarshal(actualMessageBody.Body, &actualMessage)
		require.NoError(t, err)

		assert.Equal(t, expectedFromAddress.String(), actualMessage.FromAddress)
		assert.Equal(t, expectedToAddress.String(), actualMessage.ToAddress)
		assert.Equal(t, strconv.FormatUint(expectedNonce, 10), actualMessage.Nonce)
		assert.Equal(t, expectedChain.String(), actualMessage.ChainId)
		assert.Equal(t, "", actualMessage.FeedAddress)
	})
}

func TestReachedMaxAttempts(t *testing.T) {
	ctx := t.Context()

	expectedChain := testutils.FixtureChainID
	txmMetrics, err := NewTxmMetrics(expectedChain)
	require.NoError(t, err)

	txmMetrics.ReachedMaxAttempts(ctx, true)
	value := testutil.ToFloat64(promReachedMaxAttempts.WithLabelValues(testutils.FixtureChainID.String()))
	require.InEpsilon(t, float64(1), value, 0.00001)

	txmMetrics.ReachedMaxAttempts(ctx, false)
	value = testutil.ToFloat64(promReachedMaxAttempts.WithLabelValues(testutils.FixtureChainID.String()))
	require.InDelta(t, float64(0), value, 0.00001)
}
