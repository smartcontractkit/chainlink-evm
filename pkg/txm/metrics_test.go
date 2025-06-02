package txm

import (
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/txm/types"
	svrv1 "github.com/smartcontractkit/chainlink-protos/svr/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestEmitTxMessage(t *testing.T) {
	t.Parallel()

	toAddress := testutils.NewAddress()
	fromAddress := testutils.NewAddress()

	beholderEmitter := newMockEmitter(t)
	beholderClient := beholder.GetClient()
	beholderClient.Emitter = beholderEmitter

	t.Run("overrides 0x0 as ToAddress if tx is purgeable", func(t *testing.T) {
		// GIVEN
		ctx := t.Context()

		expectedToAddress := common.Address{}
		expectedHash := common.Hash{}
		expectedChain := testutils.FixtureChainID
		expectedNonce := uint64(256)

		txmMetrics, err := NewTxmMetrics(expectedChain)
		require.NoError(t, err)

		tx := &types.Transaction{
			IsPurgeable: true,
			FromAddress: fromAddress,
			ToAddress:   toAddress,
			Nonce:       &expectedNonce,
		}

		beholderEmitter.On(
			"Emit",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil)

		// WHEN
		err = txmMetrics.EmitTxMessage(
			ctx,
			expectedHash,
			fromAddress,
			tx,
		)
		require.NoError(t, err)

		// THEN
		beholderEmitter.AssertCalled(t,
			"Emit",
			ctx,
			mock.MatchedBy(func(b []byte) bool {
				var actual svrv1.TxMessage
				if err := proto.Unmarshal(b, &actual); err != nil {
					return false
				}

				return actual.FromAddress == fromAddress.String() &&
					actual.ToAddress == expectedToAddress.String() &&
					actual.Nonce == strconv.FormatUint(expectedNonce, 10) &&
					actual.ChainId == expectedChain.String() &&
					actual.FeedAddress == ""
			}),
			"beholder_domain", "svr",
			"beholder_entity", "svr.v1.TxMessage",
			"beholder_data_schema", "/beholder-tx-message/versions/2",
		)
	})

	t.Run("sends original ToAddress if tx is not purgeable", func(t *testing.T) {
		// GIVEN
		ctx := t.Context()

		expectedToAddress := toAddress
		expectedHash := common.Hash{}
		expectedChain := testutils.FixtureChainID
		expectedNonce := uint64(256)

		txmMetrics, err := NewTxmMetrics(expectedChain)
		require.NoError(t, err)

		tx := &types.Transaction{
			IsPurgeable: false,
			FromAddress: fromAddress,
			ToAddress:   toAddress,
			Nonce:       &expectedNonce,
		}

		beholderEmitter.On(
			"Emit",
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil)

		// WHEN
		err = txmMetrics.EmitTxMessage(
			ctx,
			expectedHash,
			fromAddress,
			tx,
		)
		require.NoError(t, err)

		// THEN
		beholderEmitter.AssertCalled(t,
			"Emit",
			ctx,
			mock.MatchedBy(func(b []byte) bool {
				var actual svrv1.TxMessage
				if err := proto.Unmarshal(b, &actual); err != nil {
					return false
				}

				return actual.FromAddress == fromAddress.String() &&
					actual.ToAddress == expectedToAddress.String() &&
					actual.Nonce == strconv.FormatUint(expectedNonce, 10) &&
					actual.ChainId == expectedChain.String() &&
					actual.FeedAddress == ""
			}),
			"beholder_domain", "svr",
			"beholder_entity", "svr.v1.TxMessage",
			"beholder_data_schema", "/beholder-tx-message/versions/2",
		)
	})
}
