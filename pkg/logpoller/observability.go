package logpoller

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
	"github.com/smartcontractkit/chainlink-framework/metrics"
)

// ObservedORM is a decorator layer for ORM used by LogPoller, responsible for pushing Prometheus metrics reporting duration and size of result set for the queries.
// It doesn't change internal logic, because all calls are delegated to the origin ORM
type ObservedORM struct {
	ORM
	metrics        metrics.GenericLogPollerMetrics
	queryDuration  *prometheus.HistogramVec
	datasetSize    *prometheus.GaugeVec
	logsInserted   *prometheus.CounterVec
	blocksInserted *prometheus.CounterVec
	chainID        string
}

// NewObservedORM creates an observed version of log poller's ORM created by NewORM
// Please see ObservedLogPoller for more details on how latencies are measured
func NewObservedORM(chainID *big.Int, ds sqlutil.DataSource, lggr logger.Logger) (*ObservedORM, error) {
	lpMetrics, err := metrics.NewGenericLogPollerMetrics(chainID.String(), metrics.EVM)
	if err != nil {
		return nil, err
	}
	return &ObservedORM{
		ORM:            NewORM(chainID, ds, lggr),
		metrics:        lpMetrics,
		queryDuration:  metrics.PromLpQueryDuration,
		datasetSize:    metrics.PromLpQueryDataSets,
		logsInserted:   metrics.PromLpLogsInserted,
		blocksInserted: metrics.PromLpBlocksInserted,
		chainID:        chainID.String(),
	}, nil
}

func (o *ObservedORM) InsertLogs(ctx context.Context, logs []Log) error {
	err := withObservedExec(o, "InsertLogs", metrics.Create, func() error {
		return o.ORM.InsertLogs(ctx, logs)
	})
	trackInsertedLogsAndBlock(o, logs, nil, err)
	return err
}

func (o *ObservedORM) InsertLogsWithBlock(ctx context.Context, logs []Log, block Block) error {
	err := withObservedExec(o, "InsertLogsWithBlock", metrics.Create, func() error {
		return o.ORM.InsertLogsWithBlock(ctx, logs, block)
	})
	trackInsertedLogsAndBlock(o, logs, &block, err)
	return err
}

func (o *ObservedORM) InsertFilter(ctx context.Context, filter Filter) error {
	return withObservedExec(o, "InsertFilter", metrics.Create, func() error {
		return o.ORM.InsertFilter(ctx, filter)
	})
}

func (o *ObservedORM) LoadFilters(ctx context.Context) (map[string]Filter, error) {
	return withObservedQuery(o, "LoadFilters", func() (map[string]Filter, error) {
		return o.ORM.LoadFilters(ctx)
	})
}

func (o *ObservedORM) DeleteFilter(ctx context.Context, name string) error {
	return withObservedExec(o, "DeleteFilter", metrics.Del, func() error {
		return o.ORM.DeleteFilter(ctx, name)
	})
}

func (o *ObservedORM) DeleteBlocksBefore(ctx context.Context, end int64, limit int64) (int64, error) {
	return withObservedExecAndRowsAffected(o, "DeleteBlocksBefore", metrics.Del, func() (int64, error) {
		return o.ORM.DeleteBlocksBefore(ctx, end, limit)
	})
}

func (o *ObservedORM) DeleteLogsAndBlocksAfter(ctx context.Context, start int64) error {
	return withObservedExec(o, "DeleteLogsAndBlocksAfter", metrics.Del, func() error {
		return o.ORM.DeleteLogsAndBlocksAfter(ctx, start)
	})
}

func (o *ObservedORM) DeleteExpiredLogs(ctx context.Context, limit int64) (int64, error) {
	return withObservedExecAndRowsAffected(o, "DeleteExpiredLogs", metrics.Del, func() (int64, error) {
		return o.ORM.DeleteExpiredLogs(ctx, limit)
	})
}

func (o *ObservedORM) SelectUnmatchedLogIDs(ctx context.Context, limit int64) (ids []uint64, err error) {
	return withObservedQueryAndResults[uint64](o, "SelectUnmatchedLogIDs", func() ([]uint64, error) {
		return o.ORM.SelectUnmatchedLogIDs(ctx, limit)
	})
}

func (o *ObservedORM) SelectExcessLogIDs(ctx context.Context, limit int64) ([]uint64, error) {
	return withObservedQueryAndResults[uint64](o, "SelectExcessLogIDs", func() ([]uint64, error) {
		return o.ORM.SelectExcessLogIDs(ctx, limit)
	})
}

func (o *ObservedORM) DeleteLogsByRowID(ctx context.Context, rowIDs []uint64) (int64, error) {
	return withObservedExecAndRowsAffected(o, "DeleteLogsByRowID", metrics.Del, func() (int64, error) {
		return o.ORM.DeleteLogsByRowID(ctx, rowIDs)
	})
}

func (o *ObservedORM) SelectBlockByNumber(ctx context.Context, n int64) (*Block, error) {
	return withObservedQuery(o, "SelectBlockByNumber", func() (*Block, error) {
		return o.ORM.SelectBlockByNumber(ctx, n)
	})
}

func (o *ObservedORM) SelectLatestBlock(ctx context.Context) (*Block, error) {
	return withObservedQuery(o, "SelectLatestBlock", func() (*Block, error) {
		return o.ORM.SelectLatestBlock(ctx)
	})
}

func (o *ObservedORM) SelectOldestBlock(ctx context.Context, minAllowedBlockNumber int64) (*Block, error) {
	return withObservedQuery(o, "SelectOldestBlock", func() (*Block, error) {
		return o.ORM.SelectOldestBlock(ctx, minAllowedBlockNumber)
	})
}

func (o *ObservedORM) SelectLatestLogByEventSigWithConfs(ctx context.Context, eventSig common.Hash, address common.Address, confs evmtypes.Confirmations) (*Log, error) {
	return withObservedQuery(o, "SelectLatestLogByEventSigWithConfs", func() (*Log, error) {
		return o.ORM.SelectLatestLogByEventSigWithConfs(ctx, eventSig, address, confs)
	})
}

func (o *ObservedORM) SelectLogsWithSigs(ctx context.Context, start, end int64, address common.Address, eventSigs []common.Hash) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectLogsWithSigs", func() ([]Log, error) {
		return o.ORM.SelectLogsWithSigs(ctx, start, end, address, eventSigs)
	})
}

func (o *ObservedORM) SelectLogsCreatedAfter(ctx context.Context, address common.Address, eventSig common.Hash, after time.Time, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectLogsCreatedAfter", func() ([]Log, error) {
		return o.ORM.SelectLogsCreatedAfter(ctx, address, eventSig, after, confs)
	})
}

func (o *ObservedORM) SelectIndexedLogs(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectIndexedLogs", func() ([]Log, error) {
		return o.ORM.SelectIndexedLogs(ctx, address, eventSig, topicIndex, topicValues, confs)
	})
}

func (o *ObservedORM) SelectIndexedLogsByBlockRange(ctx context.Context, start, end int64, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectIndexedLogsByBlockRange", func() ([]Log, error) {
		return o.ORM.SelectIndexedLogsByBlockRange(ctx, start, end, address, eventSig, topicIndex, topicValues)
	})
}

func (o *ObservedORM) SelectIndexedLogsCreatedAfter(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValues []common.Hash, after time.Time, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectIndexedLogsCreatedAfter", func() ([]Log, error) {
		return o.ORM.SelectIndexedLogsCreatedAfter(ctx, address, eventSig, topicIndex, topicValues, after, confs)
	})
}

func (o *ObservedORM) SelectIndexedLogsWithSigsExcluding(ctx context.Context, sigA, sigB common.Hash, topicIndex int, address common.Address, startBlock, endBlock int64, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectIndexedLogsWithSigsExcluding", func() ([]Log, error) {
		return o.ORM.SelectIndexedLogsWithSigsExcluding(ctx, sigA, sigB, topicIndex, address, startBlock, endBlock, confs)
	})
}

func (o *ObservedORM) SelectLogs(ctx context.Context, start, end int64, address common.Address, eventSig common.Hash) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectLogs", func() ([]Log, error) {
		return o.ORM.SelectLogs(ctx, start, end, address, eventSig)
	})
}

func (o *ObservedORM) SelectIndexedLogsByTxHash(ctx context.Context, address common.Address, eventSig common.Hash, txHash common.Hash) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectIndexedLogsByTxHash", func() ([]Log, error) {
		return o.ORM.SelectIndexedLogsByTxHash(ctx, address, eventSig, txHash)
	})
}

func (o *ObservedORM) GetBlocksRange(ctx context.Context, start int64, end int64) ([]Block, error) {
	return withObservedQueryAndResults(o, "GetBlocksRange", func() ([]Block, error) {
		return o.ORM.GetBlocksRange(ctx, start, end)
	})
}

func (o *ObservedORM) SelectLatestLogEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, addresses []common.Address, eventSigs []common.Hash, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectLatestLogEventSigsAddrsWithConfs", func() ([]Log, error) {
		return o.ORM.SelectLatestLogEventSigsAddrsWithConfs(ctx, fromBlock, addresses, eventSigs, confs)
	})
}

func (o *ObservedORM) SelectLatestBlockByEventSigsAddrsWithConfs(ctx context.Context, fromBlock int64, eventSigs []common.Hash, addresses []common.Address, confs evmtypes.Confirmations) (int64, error) {
	return withObservedQuery(o, "SelectLatestBlockByEventSigsAddrsWithConfs", func() (int64, error) {
		return o.ORM.SelectLatestBlockByEventSigsAddrsWithConfs(ctx, fromBlock, eventSigs, addresses, confs)
	})
}

func (o *ObservedORM) SelectLogsDataWordRange(ctx context.Context, address common.Address, eventSig common.Hash, wordIndex int, wordValueMin, wordValueMax common.Hash, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectLogsDataWordRange", func() ([]Log, error) {
		return o.ORM.SelectLogsDataWordRange(ctx, address, eventSig, wordIndex, wordValueMin, wordValueMax, confs)
	})
}

func (o *ObservedORM) SelectLogsDataWordGreaterThan(ctx context.Context, address common.Address, eventSig common.Hash, wordIndex int, wordValueMin common.Hash, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectLogsDataWordGreaterThan", func() ([]Log, error) {
		return o.ORM.SelectLogsDataWordGreaterThan(ctx, address, eventSig, wordIndex, wordValueMin, confs)
	})
}

func (o *ObservedORM) SelectLogsDataWordBetween(ctx context.Context, address common.Address, eventSig common.Hash, wordIndexMin int, wordIndexMax int, wordValue common.Hash, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectLogsDataWordBetween", func() ([]Log, error) {
		return o.ORM.SelectLogsDataWordBetween(ctx, address, eventSig, wordIndexMin, wordIndexMax, wordValue, confs)
	})
}

func (o *ObservedORM) SelectIndexedLogsTopicGreaterThan(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValueMin common.Hash, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectIndexedLogsTopicGreaterThan", func() ([]Log, error) {
		return o.ORM.SelectIndexedLogsTopicGreaterThan(ctx, address, eventSig, topicIndex, topicValueMin, confs)
	})
}

func (o *ObservedORM) SelectIndexedLogsTopicRange(ctx context.Context, address common.Address, eventSig common.Hash, topicIndex int, topicValueMin, topicValueMax common.Hash, confs evmtypes.Confirmations) ([]Log, error) {
	return withObservedQueryAndResults(o, "SelectIndexedLogsTopicRange", func() ([]Log, error) {
		return o.ORM.SelectIndexedLogsTopicRange(ctx, address, eventSig, topicIndex, topicValueMin, topicValueMax, confs)
	})
}

func (o *ObservedORM) FilteredLogs(ctx context.Context, filter []query.Expression, limitAndSort query.LimitAndSort, queryName string) ([]Log, error) {
	return withObservedQueryAndResults(o, queryName, func() ([]Log, error) {
		return o.ORM.FilteredLogs(ctx, filter, limitAndSort, queryName)
	})
}

func withObservedQueryAndResults[T any](o *ObservedORM, queryName string, query func() ([]T, error)) ([]T, error) {
	results, err := withObservedQuery(o, queryName, query)
	if err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), client.QueryTimeout)
		defer cancel()
		o.metrics.RecordQueryDatasetSize(ctx, queryName, metrics.Read, int64(len(results)))
	}
	return results, err
}

func withObservedExecAndRowsAffected(o *ObservedORM, queryName string, queryType metrics.QueryType, exec func() (int64, error)) (int64, error) {
	queryStarted := time.Now()
	rowsAffected, err := exec()
	ctx, cancel := context.WithTimeout(context.Background(), client.QueryTimeout)
	defer cancel()
	duration := float64(time.Since(queryStarted))
	o.metrics.RecordQueryDuration(ctx, queryName, queryType, duration)
	if err == nil {
		o.metrics.RecordQueryDatasetSize(ctx, queryName, queryType, rowsAffected)
	}

	return rowsAffected, err
}

func withObservedQuery[T any](o *ObservedORM, queryName string, query func() (T, error)) (T, error) {
	queryStarted := time.Now()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), client.QueryTimeout)
		defer cancel()
		o.metrics.RecordQueryDuration(ctx, queryName, metrics.Read, float64(time.Since(queryStarted)))
	}()
	return query()
}

func withObservedExec(o *ObservedORM, query string, queryType metrics.QueryType, exec func() error) error {
	queryStarted := time.Now()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), client.QueryTimeout)
		defer cancel()
		o.metrics.RecordQueryDuration(ctx, query, queryType, float64(time.Since(queryStarted)))
	}()
	return exec()
}

func trackInsertedLogsAndBlock(o *ObservedORM, logs []Log, block *Block, err error) {
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), client.QueryTimeout)
	defer cancel()
	o.metrics.IncrementLogsInserted(ctx, int64(len(logs)))
	if block != nil {
		o.metrics.IncrementBlocksInserted(ctx, 1)
	}
}
