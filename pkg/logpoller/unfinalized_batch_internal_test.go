package logpoller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-evm/pkg/client/clienttest"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
)

type batchHeaderMockOpts struct {
	failAt map[int64]error
	nullAt map[int64]bool
	calls  *atomic.Int32
}

func copyHeadFields(dst, src *evmtypes.Head) {
	dst.Number = src.Number
	dst.Hash = src.Hash
	dst.ParentHash = src.ParentHash
	dst.Timestamp = src.Timestamp
}

func batchTestChain(start, end int64, logsPerBlock map[int64]int) (map[int64]*evmtypes.Head, map[common.Hash][]types.Log) {
	headers := make(map[int64]*evmtypes.Head, end-start+1)
	logsByHash := make(map[common.Hash][]types.Log)
	for n := start; n <= end; n++ {
		head := &evmtypes.Head{
			Number:     n,
			Hash:       hashOf(n),
			ParentHash: hashOf(n - 1),
			Timestamp:  time.Unix(n, 0).UTC(),
		}
		headers[n] = head
		for i := 0; i < logsPerBlock[n]; i++ {
			idx := int64(i) //nolint:gosec // G115: test log indexes are small
			logsByHash[head.Hash] = append(logsByHash[head.Hash], types.Log{
				BlockNumber: uint64(n), //nolint:gosec // G115: test values are small
				BlockHash:   head.Hash,
				Index:       uint(i),
				Topics:      []common.Hash{common.BigToHash(big.NewInt(n*10 + idx))},
				Address:     common.BigToAddress(big.NewInt(n)),
				TxHash:      common.BigToHash(big.NewInt(n*100 + idx)),
			})
		}
	}
	return headers, logsByHash
}

func mockBatchHeaders(t *testing.T, ec *clienttest.Client, headers map[int64]*evmtypes.Head, opts batchHeaderMockOpts) {
	t.Helper()
	ec.EXPECT().BatchCallContext(mock.Anything, mock.Anything).RunAndReturn(func(_ context.Context, elems []rpc.BatchElem) error {
		if opts.calls != nil {
			opts.calls.Add(1)
		}
		for i := range elems {
			raw, ok := elems[i].Args[0].(string)
			require.True(t, ok)
			n, err := hexutil.DecodeUint64(raw)
			require.NoError(t, err)
			num := int64(n) //nolint:gosec // G115: test block numbers are small
			if opts.failAt[num] != nil {
				elems[i].Error = opts.failAt[num]
				continue
			}
			if opts.nullAt[num] {
				continue
			}
			head, ok := headers[num]
			if !ok {
				elems[i].Error = fmt.Errorf("missing header %d", num)
				continue
			}
			copyHeadFields(elems[i].Result.(*evmtypes.Head), head)
		}
		return nil
	}).Maybe()
}

func mockHeadByNumber(t *testing.T, ec *clienttest.Client, headers map[int64]*evmtypes.Head) {
	t.Helper()
	ec.EXPECT().HeadByNumber(mock.Anything, mock.Anything).RunAndReturn(func(_ context.Context, n *big.Int) (*evmtypes.Head, error) {
		head, ok := headers[n.Int64()]
		require.True(t, ok, "missing head %d", n.Int64())
		return head, nil
	}).Maybe()
}

func mockFilterLogs(t *testing.T, ec *clienttest.Client, logsByHash map[common.Hash][]types.Log) {
	t.Helper()
	ec.EXPECT().FilterLogs(mock.Anything, mock.Anything).RunAndReturn(func(_ context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
		require.NotNil(t, q.BlockHash)
		return logsByHash[*q.BlockHash], nil
	}).Maybe()
}

func batchOpts(skipEmpty bool, chunkSize int64) Opts {
	return Opts{
		PollPeriod:                      time.Hour,
		FinalityDepth:                   2,
		BackfillBatchSize:               10,
		RPCBatchSize:                    10,
		SkipEmptyBlocks:                 skipEmpty,
		BatchUnfinalizedBlocks:          true,
		BatchUnfinalizedBlocksChunkSize: chunkSize,
	}
}

func TestGetUnfinalizedLogsBatched_MatchesLegacy(t *testing.T) {
	t.Parallel()

	for _, skipEmpty := range []bool{false, true} {
		for _, chunkSize := range []int64{1, 2, 7, 16} {
			t.Run(fmt.Sprintf("skipEmpty=%t/chunk=%d", skipEmpty, chunkSize), func(t *testing.T) {
				t.Parallel()

				const (
					start     int64 = 9
					latest    int64 = 30
					safe      int64 = 12
					finalized int64 = 10
				)
				headers, logsByHash := batchTestChain(start, latest, map[int64]int{
					start:     1,
					start + 1: 0, // finalized but empty
					start + 3: 2,
					start + 7: 1,
					latest:    1,
				})

				m := newMockedLP(t, batchOpts(skipEmpty, chunkSize))
				chainID := testutils.NewRandomEVMChainID()
				m.Client.EXPECT().ConfiguredChainID().Return(chainID)
				mockHeadByNumber(t, m.Client, headers)
				mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})
				mockFilterLogs(t, m.Client, logsByHash)

				legacyBlocks, legacyLogs, err := m.LP.getUnfinalizedLogs(t.Context(), headers[start], latest, safe, finalized, false)
				require.NoError(t, err)

				batchBlocks, batchLogs, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, safe, finalized, false)
				require.NoError(t, err)

				require.Equal(t, legacyBlocks, batchBlocks)
				require.Equal(t, legacyLogs, batchLogs)
			})
		}
	}
}

func TestGetUnfinalizedLogsBatched_AnchorIncluded(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 14
	)
	headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 2, start + 2: 1})
	m := newMockedLP(t, batchOpts(false, 2))
	chainID := testutils.NewRandomEVMChainID()
	m.Client.EXPECT().ConfiguredChainID().Return(chainID)
	mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})

	var calls atomic.Int32
	m.Client.EXPECT().FilterLogs(mock.Anything, mock.Anything).RunAndReturn(func(_ context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
		calls.Add(1)
		require.NotNil(t, q.BlockHash)
		return logsByHash[*q.BlockHash], nil
	})

	blocks, logs, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, false)
	require.NoError(t, err)
	require.NotEmpty(t, blocks)
	require.Equal(t, int32(latest-start+1), calls.Load(), "one log query per processed block, including the anchor")
	require.NotEmpty(t, logs)
}

func TestGetUnfinalizedLogsBatched_ParentChainReorg(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 30
	)
	for _, reorgAt := range []int64{start + 1, start + 8, start + 20} {
		t.Run(fmt.Sprintf("reorgAt=%d", reorgAt), func(t *testing.T) {
			t.Parallel()

			headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 1, start + 4: 1})
			headers[reorgAt].ParentHash = hashOf(reorgAt + 1000)

			m := newMockedLP(t, batchOpts(false, 4))
			chainID := testutils.NewRandomEVMChainID()
			m.Client.EXPECT().ConfiguredChainID().Return(chainID)
			mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})
			mockFilterLogs(t, m.Client, logsByHash)

			blocks, logs, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, false)
			var reorgErr *reorgError
			require.ErrorAs(t, err, &reorgErr)
			require.Equal(t, reorgAt, reorgErr.ReorgedAt.Number)

			for _, block := range blocks {
				require.Less(t, block.BlockNumber, reorgAt)
			}
			for _, log := range logs {
				require.Less(t, log.BlockNumber, reorgAt)
			}
			if len(blocks) > 0 {
				require.Less(t, blocks[len(blocks)-1].BlockNumber, reorgAt)
			}
		})
	}
}

func TestGetUnfinalizedLogsBatched_PartialHeaderBatch(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 30
	)
	for _, failAt := range []int64{start + 1, start + 5, start + 12} {
		t.Run(fmt.Sprintf("failAt=%d", failAt), func(t *testing.T) {
			t.Parallel()

			headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 1, start + 4: 1})
			m := newMockedLP(t, batchOpts(false, 4))
			chainID := testutils.NewRandomEVMChainID()
			m.Client.EXPECT().ConfiguredChainID().Return(chainID)
			mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{failAt: map[int64]error{failAt: errors.New("header unavailable")}})
			mockFilterLogs(t, m.Client, logsByHash)

			blocks, logs, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, false)
			require.Error(t, err)
			var reorgErr *reorgError
			require.NotErrorAs(t, err, &reorgErr)

			for _, block := range blocks {
				require.Less(t, block.BlockNumber, failAt)
			}
			for _, log := range logs {
				require.Less(t, log.BlockNumber, failAt)
			}
			if len(blocks) > 0 {
				require.Equal(t, failAt-1, blocks[len(blocks)-1].BlockNumber)
			}
		})
	}
}

func TestGetUnfinalizedLogsBatched_BatchUnavailable(t *testing.T) {
	t.Parallel()

	const start int64 = 9
	headers, logsByHash := batchTestChain(start, start+4, map[int64]int{start: 1})
	m := newMockedLP(t, batchOpts(false, 4))
	chainID := testutils.NewRandomEVMChainID()
	m.Client.EXPECT().ConfiguredChainID().Return(chainID)
	m.Client.EXPECT().BatchCallContext(mock.Anything, mock.Anything).Return(errors.New("batch unavailable")).Once()
	mockFilterLogs(t, m.Client, logsByHash)

	blocks, _, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], start+4, start+4, start+1, false)
	require.Error(t, err)
	require.Len(t, blocks, 1, "the anchor is the only safe checkpoint when the header batch fails")
	require.Equal(t, start, blocks[0].BlockNumber)
}

func TestGetUnfinalizedLogsBatched_NullHeaderAtTip(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 30
		nullAt int64 = 15
	)
	headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 1, start + 4: 1})
	m := newMockedLP(t, batchOpts(false, 16))
	chainID := testutils.NewRandomEVMChainID()
	m.Client.EXPECT().ConfiguredChainID().Return(chainID)
	mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{nullAt: map[int64]bool{nullAt: true}})
	mockFilterLogs(t, m.Client, logsByHash)

	blocks, _, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, false)
	require.NoError(t, err)
	require.NotEmpty(t, blocks)
	require.Equal(t, nullAt-1, blocks[len(blocks)-1].BlockNumber)
	for _, block := range blocks {
		require.Less(t, block.BlockNumber, nullAt)
	}
}

func TestGetUnfinalizedLogsBatched_LogFailureStopsAtFirstIncomplete(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 20
		failAt int64 = start + 5
	)
	headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 1, start + 4: 1})
	failHash := headers[failAt].Hash

	m := newMockedLP(t, batchOpts(false, 8))
	chainID := testutils.NewRandomEVMChainID()
	m.Client.EXPECT().ConfiguredChainID().Return(chainID)
	mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})
	m.Client.EXPECT().FilterLogs(mock.Anything, mock.Anything).RunAndReturn(func(_ context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
		if *q.BlockHash == failHash {
			return nil, errors.New("log query failed")
		}
		return logsByHash[*q.BlockHash], nil
	})

	blocks, logs, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, false)
	require.Error(t, err)
	for _, block := range blocks {
		require.Less(t, block.BlockNumber, failAt)
	}
	for _, log := range logs {
		require.Less(t, log.BlockNumber, failAt)
	}
	if len(blocks) > 0 {
		// In-flight sibling queries may be canceled by the first failure, so the
		// saved prefix can end earlier than failAt-1. It must never reach failAt.
		require.Less(t, blocks[len(blocks)-1].BlockNumber, failAt)
	}
}

func TestGetUnfinalizedLogsBatched_ReplayDB(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 12
	)
	headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 1})

	t.Run("matching block passes", func(t *testing.T) {
		m := newMockedLP(t, batchOpts(false, 2))
		chainID := testutils.NewRandomEVMChainID()
		m.Client.EXPECT().ConfiguredChainID().Return(chainID)
		mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})
		mockFilterLogs(t, m.Client, logsByHash)
		m.ORM.EXPECT().SelectBlockByNumber(mock.Anything, mock.Anything).RunAndReturn(func(_ context.Context, n int64) (*Block, error) {
			return &Block{BlockNumber: n, BlockHash: hashOf(n)}, nil
		}).Maybe()

		_, _, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, true)
		require.NoError(t, err)
	})

	t.Run("missing DB block passes", func(t *testing.T) {
		m := newMockedLP(t, batchOpts(false, 2))
		chainID := testutils.NewRandomEVMChainID()
		m.Client.EXPECT().ConfiguredChainID().Return(chainID)
		mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})
		mockFilterLogs(t, m.Client, logsByHash)
		m.ORM.EXPECT().SelectBlockByNumber(mock.Anything, mock.Anything).Return(nil, sql.ErrNoRows).Maybe()

		_, _, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, true)
		require.NoError(t, err)
	})

	t.Run("mismatched DB block is a reorg", func(t *testing.T) {
		m := newMockedLP(t, batchOpts(false, 2))
		chainID := testutils.NewRandomEVMChainID()
		m.Client.EXPECT().ConfiguredChainID().Return(chainID)
		mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})
		mockFilterLogs(t, m.Client, logsByHash)
		mismatchAt := start + 3
		m.ORM.EXPECT().SelectBlockByNumber(mock.Anything, mock.Anything).RunAndReturn(func(_ context.Context, n int64) (*Block, error) {
			if n == mismatchAt {
				return &Block{BlockNumber: n, BlockHash: common.HexToHash("0xdead")}, nil
			}
			return &Block{BlockNumber: n, BlockHash: hashOf(n)}, nil
		}).Maybe()

		_, _, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, true)
		var reorgErr *reorgError
		require.ErrorAs(t, err, &reorgErr)
		require.Equal(t, mismatchAt, reorgErr.ReorgedAt.Number)
	})

	t.Run("DB error aborts the chunk", func(t *testing.T) {
		m := newMockedLP(t, batchOpts(false, 2))
		chainID := testutils.NewRandomEVMChainID()
		m.Client.EXPECT().ConfiguredChainID().Return(chainID)
		mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})
		mockFilterLogs(t, m.Client, logsByHash)
		m.ORM.EXPECT().SelectBlockByNumber(mock.Anything, mock.Anything).Return(nil, errors.New("db unavailable")).Maybe()

		_, _, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, true)
		require.Error(t, err)
		var reorgErr *reorgError
		require.NotErrorAs(t, err, &reorgErr)
	})
}

func TestGetUnfinalizedLogsBatched_WalkedCap(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 1
		latest int64 = 5000
	)
	headers, logsByHash := batchTestChain(start, latest, nil)

	var maxRequested atomic.Int64
	m := newMockedLP(t, batchOpts(true, 16))
	chainID := testutils.NewRandomEVMChainID()
	m.Client.EXPECT().ConfiguredChainID().Return(chainID)
	m.Client.EXPECT().BatchCallContext(mock.Anything, mock.Anything).RunAndReturn(func(_ context.Context, elems []rpc.BatchElem) error {
		for i := range elems {
			raw := elems[i].Args[0].(string)
			n, err := hexutil.DecodeUint64(raw)
			require.NoError(t, err)
			num := int64(n) //nolint:gosec // G115: test block numbers are small
			for {
				old := maxRequested.Load()
				if num <= old || maxRequested.CompareAndSwap(old, num) {
					break
				}
			}
			copyHeadFields(elems[i].Result.(*evmtypes.Head), headers[num])
		}
		return nil
	}).Maybe()
	mockFilterLogs(t, m.Client, logsByHash)

	blocks, _, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start, false)
	require.NoError(t, err)
	require.NotEmpty(t, blocks)
	require.LessOrEqual(t, maxRequested.Load(), start+maxUnfinalizedBlocksWalked-1)
}

func TestGetUnfinalizedLogsBatched_ConcurrencyBound(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 20
	)
	headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 1, start + 2: 1})
	anchorHash := headers[start].Hash

	var inFlight, maxInFlight atomic.Int32
	started := make(chan struct{}, 32)
	release := make(chan struct{})

	m := newMockedLP(t, batchOpts(false, 8))
	chainID := testutils.NewRandomEVMChainID()
	m.Client.EXPECT().ConfiguredChainID().Return(chainID)
	mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})
	m.Client.EXPECT().FilterLogs(mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
		if q.BlockHash != nil && *q.BlockHash == anchorHash {
			return logsByHash[anchorHash], nil
		}
		current := inFlight.Add(1)
		for {
			old := maxInFlight.Load()
			if current <= old || maxInFlight.CompareAndSwap(old, current) {
				break
			}
		}
		started <- struct{}{}
		select {
		case <-release:
		case <-ctx.Done():
		}
		inFlight.Add(-1)
		return logsByHash[*q.BlockHash], nil
	})

	done := make(chan struct{})
	var (
		blocks []Block
		logs   []Log
		err    error
	)
	go func() {
		defer close(done)
		blocks, logs, err = m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, false)
	}()

	for i := 0; i < 4; i++ {
		select {
		case <-started:
		case <-time.After(5 * time.Second):
			close(release)
			<-done
			t.Fatalf("only %d concurrent log queries started", i)
		}
	}
	close(release)
	<-done

	require.NoError(t, err)
	require.NotEmpty(t, blocks)
	require.NotEmpty(t, logs)
	require.LessOrEqual(t, maxInFlight.Load(), int32(defaultUnfinalizedBatchLogConcurrency))
	require.Greater(t, maxInFlight.Load(), int32(1))
}

func TestGetUnfinalizedLogsBatched_ConvertLogsPerBlock(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 14
	)
	headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 2, start + 1: 1, start + 2: 3})
	m := newMockedLP(t, batchOpts(false, 3))
	chainID := testutils.NewRandomEVMChainID()
	m.Client.EXPECT().ConfiguredChainID().Return(chainID)
	mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{})
	mockFilterLogs(t, m.Client, logsByHash)

	_, logs, err := m.LP.getUnfinalizedLogsBatched(t.Context(), headers[start], latest, latest, start+1, false)
	require.NoError(t, err)
	require.NotEmpty(t, logs)

	timestamps := make(map[common.Hash]time.Time, len(headers))
	for _, head := range headers {
		timestamps[head.Hash] = head.Timestamp
	}
	for _, log := range logs {
		expected, ok := timestamps[log.BlockHash]
		require.True(t, ok)
		require.True(t, expected.Equal(log.BlockTimestamp), "block %d timestamp mismatch: expected %s got %s", log.BlockNumber, expected, log.BlockTimestamp)
	}
}

func TestGetUnfinalizedLogsForMode_FlagOffUsesLegacy(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 14
	)
	headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 1, start + 2: 1})
	opts := batchOpts(false, 4)
	opts.BatchUnfinalizedBlocks = false

	m := newMockedLP(t, opts)
	chainID := testutils.NewRandomEVMChainID()
	m.Client.EXPECT().ConfiguredChainID().Return(chainID)
	mockHeadByNumber(t, m.Client, headers)
	mockFilterLogs(t, m.Client, logsByHash)
	var batchCalled atomic.Bool
	m.Client.On("BatchCallContext", mock.Anything, mock.Anything).Run(func(mock.Arguments) {
		batchCalled.Store(true)
	}).Return(nil).Maybe()

	blocks, logs, err := m.LP.getUnfinalizedLogsForMode(t.Context(), headers[start], latest, latest, start+1, false)
	require.NoError(t, err)
	require.NotEmpty(t, blocks)
	require.NotEmpty(t, logs)
	require.False(t, batchCalled.Load(), "BatchCallContext must not be called when the flag is off")
}

func TestGetUnfinalizedLogsForMode_FlagOnUsesBatched(t *testing.T) {
	t.Parallel()

	const (
		start  int64 = 9
		latest int64 = 14
	)
	headers, logsByHash := batchTestChain(start, latest, map[int64]int{start: 1, start + 2: 1})

	m := newMockedLP(t, batchOpts(false, 3))
	chainID := testutils.NewRandomEVMChainID()
	m.Client.EXPECT().ConfiguredChainID().Return(chainID)
	var batchCalls atomic.Int32
	mockBatchHeaders(t, m.Client, headers, batchHeaderMockOpts{calls: &batchCalls})
	mockFilterLogs(t, m.Client, logsByHash)

	var headByNumberCalled atomic.Bool
	m.Client.On("HeadByNumber", mock.Anything, mock.Anything).Run(func(mock.Arguments) {
		headByNumberCalled.Store(true)
	}).Return(nil, errors.New("HeadByNumber must not be used in batched mode")).Maybe()

	blocks, logs, err := m.LP.getUnfinalizedLogsForMode(t.Context(), headers[start], latest, latest, start+1, false)
	require.NoError(t, err)
	require.NotEmpty(t, blocks)
	require.NotEmpty(t, logs)
	require.Positive(t, batchCalls.Load(), "batched mode must fetch headers in a batch")
	require.False(t, headByNumberCalled.Load(), "batched mode must not use the legacy per-block header path")
}

func TestNewLogPoller_BatchModeSettings(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		opts        Opts
		wantEnabled bool
		wantChunk   int64
	}{
		{
			name:        "disabled by default",
			opts:        Opts{},
			wantEnabled: false,
			wantChunk:   defaultUnfinalizedBatchChunkSize,
		},
		{
			name:        "enabled with explicit chunk size",
			opts:        Opts{BatchUnfinalizedBlocks: true, BatchUnfinalizedBlocksChunkSize: 8},
			wantEnabled: true,
			wantChunk:   8,
		},
		{
			name:        "enabled with default chunk size",
			opts:        Opts{BatchUnfinalizedBlocks: true},
			wantEnabled: true,
			wantChunk:   defaultUnfinalizedBatchChunkSize,
		},
		{
			name:        "chunk size clamps to max",
			opts:        Opts{BatchUnfinalizedBlocks: true, BatchUnfinalizedBlocksChunkSize: 500},
			wantEnabled: true,
			wantChunk:   maxUnfinalizedBatchChunkSize,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			m := newMockedLP(t, tc.opts)
			require.Equal(t, tc.wantEnabled, m.LP.batchUnfinalizedBlocks)
			require.Equal(t, tc.wantChunk, m.LP.unfinalizedBatchChunkSize)
			require.Equal(t, int64(defaultUnfinalizedBatchLogConcurrency), m.LP.unfinalizedBatchLogConcurrency)
		})
	}
}
