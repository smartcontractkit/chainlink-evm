package logpoller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"

	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
)

func (lp *logPoller) recordUnfinalizedBatchChunk(outcome string, headers int, start time.Time) {
	chainID := lp.ec.ConfiguredChainID().String()
	promLpUnfinalizedBatchChunks.WithLabelValues(chainID, outcome).Inc()
	promLpUnfinalizedBatchChunkDuration.WithLabelValues(chainID).Observe(time.Since(start).Seconds())
	promLpUnfinalizedBatchHeaders.WithLabelValues(chainID).Add(float64(headers))
}

func (lp *logPoller) unfinalizedBlock(head *evmtypes.Head, safe, finalized int64) Block {
	return Block{
		BlockHash:            head.Hash,
		BlockNumber:          head.Number,
		BlockTimestamp:       head.Timestamp,
		FinalizedBlockNumber: finalized,
		SafeBlockNumber:      safe,
	}
}

func (lp *logPoller) appendUnfinalizedBlock(head *evmtypes.Head, rpcLogs []types.Log, blocks []Block, logs []Log, safe, finalized int64) ([]Block, []Log) {
	block := lp.unfinalizedBlock(head, safe, finalized)
	logs = append(logs, convertLogs(rpcLogs, []Block{block}, lp.lggr, lp.ec.ConfiguredChainID())...)
	// Preserve the legacy skipEmptyBlocks rule: save the block if it has logs,
	// if skipping is disabled, or if it is the finalized checkpoint.
	if len(rpcLogs) > 0 || !lp.skipEmptyBlocks || head.Number == finalized {
		blocks = append(blocks, block)
	}
	return blocks, logs
}

// batchFetchHeaders fetches [from, to] in one JSON-RPC batch and returns the
// contiguous prefix of headers that are available and internally consistent.
// A nil error with a short prefix means the endpoint is behind and the caller
// should save the prefix and retry the missing height on the next poll.
func (lp *logPoller) batchFetchHeaders(ctx context.Context, from, to int64) ([]*evmtypes.Head, error) {
	if from > to {
		return nil, nil
	}

	reqs := make([]rpc.BatchElem, 0, int(to-from+1)) //nolint:gosec // G115: bounded by unfinalizedBatchChunkSize (<=50)
	for n := from; n <= to; n++ {
		reqs = append(reqs, newBlockReq(hexutil.EncodeBig(big.NewInt(n))))
	}

	if err := lp.ec.BatchCallContext(ctx, reqs); err != nil {
		return nil, fmt.Errorf("failed to batch fetch unfinalized headers %d-%d: %w", from, to, err)
	}

	headers := make([]*evmtypes.Head, 0, len(reqs))
	for i, req := range reqs {
		expected := from + int64(i) //nolint:gosec // G115: batch index is bounded by the chunk size
		if req.Error != nil {
			return headers, fmt.Errorf("failed to fetch unfinalized header %d: %w", expected, req.Error)
		}

		head, ok := req.Result.(*evmtypes.Head)
		if !ok {
			return headers, fmt.Errorf("failed to fetch unfinalized header %d: unexpected result type %T", expected, req.Result)
		}
		if head == nil || (head.Hash == common.Hash{} && head.Number == 0) {
			// eth_getBlockByNumber returns null for a height the endpoint has not
			// seen yet. Do not treat that as a fatal error; stop at the last
			// contiguous header and retry next poll.
			return headers, nil
		}
		if head.Number != expected {
			return headers, fmt.Errorf("failed to fetch unfinalized header %d: got block number %d", expected, head.Number)
		}
		if err := validateHead(head); err != nil {
			return headers, fmt.Errorf("failed to fetch unfinalized header %d: %w", expected, err)
		}
		headers = append(headers, head)
	}

	return headers, nil
}

// batchFilterLogsByHash queries logs for each header with bounded concurrency.
// It returns results in header order, the number of contiguous completed
// headers, and the first real error. A short prefix is always safe to persist:
// the next poll resumes at the first incomplete block, so no block is skipped.
func (lp *logPoller) batchFilterLogsByHash(ctx context.Context, headers []*evmtypes.Head) ([][]types.Log, int, error) {
	results := make([][]types.Log, len(headers))
	completed := make([]bool, len(headers))
	if len(headers) == 0 {
		return results, 0, nil
	}

	parentCtx := ctx
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	concurrency := int(lp.unfinalizedBatchLogConcurrency) //nolint:gosec // G115: bounded by the number of headers in the chunk
	if concurrency < 1 {
		concurrency = 1
	}
	if concurrency > len(headers) {
		concurrency = len(headers)
	}

	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstErr error
	firstErrIndex := len(headers)

	for i, header := range headers {
		wg.Add(1)
		go func(i int, header *evmtypes.Head) {
			defer wg.Done()

			select {
			case sem <- struct{}{}:
				defer func() { <-sem }()
			case <-ctx.Done():
				return
			}

			blockHash := header.Hash
			rpcLogs, err := lp.latencyMonitor.FilterLogs(ctx, lp.Filter(nil, nil, &blockHash))
			if err != nil {
				if errors.Is(err, context.Canceled) && parentCtx.Err() == nil {
					// A sibling failed and canceled this call; the sibling records
					// the real error and the caller will see the shorter prefix.
					return
				}
				mu.Lock()
				if firstErr == nil || i < firstErrIndex {
					firstErr = fmt.Errorf("unable to query logs for block %d (%s): %w", header.Number, header.Hash, err)
					firstErrIndex = i
				}
				mu.Unlock()
				cancel()
				return
			}

			results[i] = rpcLogs
			completed[i] = true
		}(i, header)
	}
	wg.Wait()

	firstIncomplete := 0
	for firstIncomplete < len(headers) && completed[firstIncomplete] {
		firstIncomplete++
	}
	if firstIncomplete == len(headers) {
		return results, len(headers), nil
	}
	if firstErr == nil {
		if err := parentCtx.Err(); err != nil {
			firstErr = err
		} else {
			firstErr = fmt.Errorf("unfinalized log query stopped before block %d", headers[firstIncomplete].Number)
		}
	}
	return results, firstIncomplete, firstErr
}

// getUnfinalizedLogsBatched is the opt-in replacement for getUnfinalizedLogs.
// It preserves the legacy anchor query, reorg rule, replay DB check,
// skipEmptyBlocks behavior, and deferred last-block checkpoint.
func (lp *logPoller) getUnfinalizedLogsBatched(ctx context.Context, currentBlock *evmtypes.Head, latest, safe, finalized int64, isReplay bool) (blocks []Block, logs []Log, err error) {
	var lastProcessed *evmtypes.Head
	defer func() {
		// Ensure the last processed block is always a checkpoint for the next poll.
		if lastProcessed != nil && (len(blocks) == 0 || blocks[len(blocks)-1].BlockNumber != lastProcessed.Number) {
			blocks = append(blocks, lp.unfinalizedBlock(lastProcessed, safe, finalized))
		}
	}()

	// The anchor block's logs are queried on every poll, exactly as the legacy path.
	anchorHash := currentBlock.Hash
	rpcLogs, err := lp.latencyMonitor.FilterLogs(ctx, lp.Filter(nil, nil, &anchorHash))
	if err != nil {
		lp.lggr.Warnw("Unable to query for logs, retrying on next poll", "err", err, "block", currentBlock.Number)
		return nil, nil, fmt.Errorf("unable to query for logs: %w", err)
	}
	lp.lggr.Debugw("Unfinalized log query", "logs", len(logs), "currentBlockNumber", currentBlock.Number, "blockHash", currentBlock.Hash, "timestamp", currentBlock.Timestamp)
	blocks, logs = lp.appendUnfinalizedBlock(currentBlock, rpcLogs, blocks, logs, safe, finalized)
	lastProcessed = currentBlock

	if currentBlock.Number >= latest {
		return blocks, logs, nil
	}

	walked := int64(1) // the anchor counts toward the per-poll cap
	for currentBlock.Number < latest {
		remaining := int64(maxUnfinalizedBlocksWalked) - walked
		if remaining <= 0 {
			lp.lggr.Warnw("Too many unfinalized blocks, stopping log retrieval to avoid OOM", "currentBlockNumber", currentBlock.Number, "latestBlockNumber", latest)
			return blocks, logs, nil
		}

		chunkSize := lp.unfinalizedBatchChunkSize
		if chunkSize > remaining {
			chunkSize = remaining
		}
		to := min(latest, currentBlock.Number+chunkSize)

		chunkStart := time.Now()
		headers, headerErr := lp.batchFetchHeaders(ctx, currentBlock.Number+1, to)
		if len(headers) == 0 {
			outcome := "partial"
			if headerErr != nil {
				outcome = "error"
			}
			lp.recordUnfinalizedBatchChunk(outcome, 0, chunkStart)
			return blocks, logs, headerErr
		}

		// Verify the whole chunk before issuing any log query. If the parent chain
		// or the replay DB state disagrees, discard the chunk and let the existing
		// reorg handler run; no orphaned logs can be saved.
		prev := currentBlock
		for _, header := range headers {
			if header.ParentHash != prev.Hash {
				lp.recordUnfinalizedBatchChunk("reorg", 0, chunkStart)
				return blocks, logs, newReorgError(header)
			}
			if isReplay {
				dbBlock, dbErr := lp.orm.SelectBlockByNumber(ctx, header.Number)
				if dbErr != nil && !errors.Is(dbErr, sql.ErrNoRows) {
					lp.lggr.Warnw("Unable to get next block from DB during replay, retrying on next poll", "err", dbErr, "block", header.Number)
					lp.recordUnfinalizedBatchChunk("error", 0, chunkStart)
					return blocks, logs, fmt.Errorf("failed to get next block from DB during replay: %w", dbErr)
				}
				if dbBlock != nil && header.Hash != dbBlock.BlockHash {
					lp.recordUnfinalizedBatchChunk("reorg", 0, chunkStart)
					return blocks, logs, newReorgError(header)
				}
			}
			prev = header
		}

		logsByBlock, completed, logErr := lp.batchFilterLogsByHash(ctx, headers)
		for i := 0; i < completed; i++ {
			header := headers[i]
			lp.lggr.Debugw("Unfinalized log query", "logs", len(logs), "currentBlockNumber", header.Number, "blockHash", header.Hash, "timestamp", header.Timestamp)
			blocks, logs = lp.appendUnfinalizedBlock(header, logsByBlock[i], blocks, logs, safe, finalized)
			lastProcessed = header
			walked++
		}
		if logErr != nil {
			lp.recordUnfinalizedBatchChunk("error", completed, chunkStart)
			return blocks, logs, logErr
		}

		currentBlock = headers[len(headers)-1]
		if currentBlock.Number < to {
			// The endpoint was behind or a header failed. Persist the prefix and
			// retry the missing height on the next poll.
			outcome := "partial"
			if headerErr != nil {
				outcome = "error"
			}
			lp.recordUnfinalizedBatchChunk(outcome, completed, chunkStart)
			return blocks, logs, headerErr
		}
		lp.recordUnfinalizedBatchChunk("success", len(headers), chunkStart)
	}

	return blocks, logs, nil
}
