package logpoller

import (
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
	ubig "github.com/smartcontractkit/chainlink-evm/pkg/utils/big"
	"github.com/smartcontractkit/chainlink-framework/metrics"
)

const network = "EVM Test"

func TestMultipleMetricsArePublished(t *testing.T) {
	ctx := testutils.Context(t)
	orm := createObservedORM(t, 100)
	t.Cleanup(func() { resetMetrics(*orm) })
	require.Equal(t, 0, testutil.CollectAndCount(orm.queryDuration))

	_, _ = orm.SelectIndexedLogs(ctx, common.Address{}, common.Hash{}, 1, []common.Hash{}, 1)
	_, _ = orm.SelectIndexedLogsByBlockRange(ctx, 0, 1, common.Address{}, common.Hash{}, 1, []common.Hash{})
	_, _ = orm.SelectIndexedLogsTopicGreaterThan(ctx, common.Address{}, common.Hash{}, 1, common.Hash{}, 1)
	_, _ = orm.SelectIndexedLogsTopicRange(ctx, common.Address{}, common.Hash{}, 1, common.Hash{}, common.Hash{}, 1)
	_, _ = orm.SelectIndexedLogsWithSigsExcluding(ctx, common.Hash{}, common.Hash{}, 1, common.Address{}, 0, 1, 1)
	_, _ = orm.SelectLogsDataWordRange(ctx, common.Address{}, common.Hash{}, 0, common.Hash{}, common.Hash{}, 1)
	_, _ = orm.SelectLogsDataWordGreaterThan(ctx, common.Address{}, common.Hash{}, 0, common.Hash{}, 1)
	_, _ = orm.SelectLogsCreatedAfter(ctx, common.Address{}, common.Hash{}, time.Now(), 0)
	_, _ = orm.SelectLatestLogByEventSigWithConfs(ctx, common.Hash{}, common.Address{}, 0)
	_, _ = orm.SelectLatestLogEventSigsAddrsWithConfs(ctx, 0, []common.Address{{}}, []common.Hash{{}}, 1)
	_, _ = orm.SelectIndexedLogsCreatedAfter(ctx, common.Address{}, common.Hash{}, 1, []common.Hash{}, time.Now(), 0)
	_ = orm.InsertLogs(ctx, []Log{})
	_ = orm.InsertLogsWithBlock(ctx, []Log{}, Block{
		BlockNumber:    1,
		BlockTimestamp: time.Now(),
	})

	require.Equal(t, 13, testutil.CollectAndCount(orm.queryDuration))
	require.Equal(t, 10, testutil.CollectAndCount(orm.datasetSize))
}

func TestShouldPublishDurationInCaseOfError(t *testing.T) {
	ctx := testutils.Context(t)
	orm := createObservedORM(t, 200)
	t.Cleanup(func() { resetMetrics(*orm) })
	require.Equal(t, 0, testutil.CollectAndCount(orm.queryDuration))

	_, err := orm.SelectLatestLogByEventSigWithConfs(ctx, common.Hash{}, common.Address{}, 0)
	require.Error(t, err)

	require.Equal(t, 1, testutil.CollectAndCount(orm.queryDuration))
	require.Equal(t, 1, counterFromHistogramByLabels(t, orm.queryDuration, network, "200", "SelectLatestLogByEventSigWithConfs", "read"))
}

func TestMetricsAreProperlyPopulatedWithLabels(t *testing.T) {
	orm := createObservedORM(t, 420)
	t.Cleanup(func() { resetMetrics(*orm) })
	expectedCount := 9
	expectedSize := 2

	for i := 0; i < expectedCount; i++ {
		_, err := withObservedQueryAndResults(orm, "query", func() ([]string, error) { return []string{"value1", "value2"}, nil })
		require.NoError(t, err)
	}

	require.Equal(t, expectedCount, counterFromHistogramByLabels(t, orm.queryDuration, network, "420", "query", "read"))
	require.Equal(t, expectedSize, counterFromGaugeByLabels(orm.datasetSize, network, "420", "query", "read"))

	require.Equal(t, 0, counterFromHistogramByLabels(t, orm.queryDuration, network, "420", "other_query", "read"))
	require.Equal(t, 0, counterFromHistogramByLabels(t, orm.queryDuration, network, "5", "query", "read"))

	require.Equal(t, 0, counterFromGaugeByLabels(orm.datasetSize, network, "420", "other_query", "read"))
	require.Equal(t, 0, counterFromGaugeByLabels(orm.datasetSize, network, "5", "query", "read"))
}

func TestNotPublishingDatasetSizeInCaseOfError(t *testing.T) {
	orm := createObservedORM(t, 420)

	_, err := withObservedQueryAndResults(orm, "errorQuery", func() ([]string, error) { return nil, errors.New("error") })
	require.Error(t, err)

	require.Equal(t, 1, counterFromHistogramByLabels(t, orm.queryDuration, network, "420", "errorQuery", "read"))
	require.Equal(t, 0, counterFromGaugeByLabels(orm.datasetSize, network, "420", "errorQuery", "read"))
}

func TestMetricsAreProperlyPopulatedForWrites(t *testing.T) {
	orm := createObservedORM(t, 420)
	require.NoError(t, withObservedExec(orm, "execQuery", metrics.Create, func() error { return nil }))
	require.Error(t, withObservedExec(orm, "execQuery", metrics.Create, func() error { return errors.New("error") }))

	require.Equal(t, 2, counterFromHistogramByLabels(t, orm.queryDuration, network, "420", "execQuery", "create"))
}

func TestCountersAreProperlyPopulatedForWrites(t *testing.T) {
	ctx := testutils.Context(t)
	orm := createObservedORM(t, 420)
	logs := generateRandomLogs(420, 20)

	// First insert 10 logs
	require.NoError(t, orm.InsertLogs(ctx, logs[:10]))
	assert.Equal(t, 10, int(testutil.ToFloat64(orm.logsInserted.WithLabelValues(network, "420"))))

	// Insert 5 more logs with block
	require.NoError(t, orm.InsertLogsWithBlock(ctx, logs[10:15], Block{
		BlockHash:            utils.RandomBytes32(),
		BlockNumber:          10,
		BlockTimestamp:       time.Now(),
		FinalizedBlockNumber: 5,
	}))
	assert.Equal(t, 15, int(testutil.ToFloat64(orm.logsInserted.WithLabelValues(network, "420"))))
	assert.Equal(t, 1, int(testutil.ToFloat64(orm.blocksInserted.WithLabelValues(network, "420"))))

	// Insert 5 more logs with block
	require.NoError(t, orm.InsertLogsWithBlock(ctx, logs[15:], Block{
		BlockHash:            utils.RandomBytes32(),
		BlockNumber:          15,
		BlockTimestamp:       time.Now(),
		FinalizedBlockNumber: 5,
	}))
	assert.Equal(t, 20, int(testutil.ToFloat64(orm.logsInserted.WithLabelValues(network, "420"))))
	assert.Equal(t, 2, int(testutil.ToFloat64(orm.blocksInserted.WithLabelValues(network, "420"))))

	rowsAffected, err := orm.DeleteExpiredLogs(ctx, 3)
	require.NoError(t, err)
	require.Equal(t, int64(0), rowsAffected)
	assert.Equal(t, 0, counterFromGaugeByLabels(orm.datasetSize, network, "420", "DeleteExpiredLogs", "delete"))

	rowsAffected, err = orm.DeleteBlocksBefore(ctx, 30, 0)
	require.NoError(t, err)
	require.Equal(t, int64(2), rowsAffected)
	assert.Equal(t, 2, counterFromGaugeByLabels(orm.datasetSize, network, "420", "DeleteBlocksBefore", "delete"))

	// Don't update counters in case of an error
	require.Error(t, orm.InsertLogsWithBlock(ctx, logs, Block{
		BlockHash:      utils.RandomBytes32(),
		BlockTimestamp: time.Now(),
	}))
	assert.Equal(t, 20, int(testutil.ToFloat64(orm.logsInserted.WithLabelValues(network, "420"))))
	assert.Equal(t, 2, int(testutil.ToFloat64(orm.blocksInserted.WithLabelValues(network, "420"))))
}

func generateRandomLogs(chainID, count int) []Log {
	logs := make([]Log, count)
	for i := range logs {
		logs[i] = Log{
			EVMChainID:     ubig.NewI(int64(chainID)),
			LogIndex:       int64(i + 1),
			BlockHash:      utils.RandomBytes32(),
			BlockNumber:    int64(i + 1),
			BlockTimestamp: time.Now(),
			Topics:         [][]byte{},
			EventSig:       utils.RandomBytes32(),
			Address:        utils.RandomAddress(),
			TxHash:         utils.RandomBytes32(),
			Data:           []byte{},
			CreatedAt:      time.Now(),
		}
	}
	return logs
}

func NewTestObservedORM(chainID *big.Int, ds sqlutil.DataSource, lggr logger.Logger) (*ObservedORM, error) {
	lpMetrics, err := metrics.NewGenericLogPollerMetrics(chainID.String(), network)
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

func createObservedORM(t *testing.T, chainId int64) *ObservedORM {
	lggr := logger.Test(t)
	db := testutils.NewSqlxDB(t)
	observed, err := NewTestObservedORM(big.NewInt(chainId), db, lggr)
	require.NoError(t, err)
	return observed
}

func resetMetrics(lp ObservedORM) {
	lp.queryDuration.Reset()
	lp.datasetSize.Reset()
	lp.logsInserted.Reset()
	lp.blocksInserted.Reset()
}

func counterFromGaugeByLabels(gaugeVec *prometheus.GaugeVec, labels ...string) int {
	value := testutil.ToFloat64(gaugeVec.WithLabelValues(labels...))
	return int(value)
}

func counterFromHistogramByLabels(t *testing.T, histogramVec *prometheus.HistogramVec, labels ...string) int {
	observer, err := histogramVec.GetMetricWithLabelValues(labels...)
	require.NoError(t, err)

	metricCh := make(chan prometheus.Metric, 1)
	observer.(prometheus.Histogram).Collect(metricCh)
	close(metricCh)

	metric := <-metricCh
	pb := &io_prometheus_client.Metric{}
	err = metric.Write(pb)
	require.NoError(t, err)

	return int(pb.GetHistogram().GetSampleCount())
}
