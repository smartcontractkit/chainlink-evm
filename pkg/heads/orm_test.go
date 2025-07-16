package heads_test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-evm/pkg/heads"
	"github.com/smartcontractkit/chainlink-evm/pkg/testutils"
)

func TestORM_IdempotentInsertHead(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db)

	// Returns nil when inserting first head
	head := testutils.Head(0)
	require.NoError(t, orm.IdempotentInsertHead(t.Context(), head))

	// Head is inserted
	foundHead, err := orm.LatestHead(t.Context())
	require.NoError(t, err)
	assert.Equal(t, head.Hash, foundHead.Hash)

	// Returns nil when inserting same head again
	require.NoError(t, orm.IdempotentInsertHead(t.Context(), head))

	// Head is still inserted
	foundHead, err = orm.LatestHead(t.Context())
	require.NoError(t, err)
	assert.Equal(t, head.Hash, foundHead.Hash)
}

func TestORM_TrimOldHeads(t *testing.T) {
	t.Parallel()
	heads.BatchSize = 0
	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db)

	for i := 0; i < 10; i++ {
		head := testutils.Head(i)
		require.NoError(t, orm.IdempotentInsertHead(t.Context(), head))
	}

	uncleHead := testutils.Head(5)
	require.NoError(t, orm.IdempotentInsertHead(t.Context(), uncleHead))

	err := orm.TrimOldHeads(t.Context(), 5)
	require.NoError(t, err)

	heads, err := orm.LatestHeads(t.Context(), 0)
	require.NoError(t, err)

	// uncle block was loaded too
	require.Len(t, heads, 6)
	for i := 0; i < 5; i++ {
		require.LessOrEqual(t, int64(5), heads[i].Number)
	}
}

func TestORM_TrimOldHeads_Batch(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db)

	for i := 0; i < 10; i++ {
		head := testutils.Head(i)
		require.NoError(t, orm.IdempotentInsertHead(t.Context(), head))
	}

	uncleHead := testutils.Head(5)
	require.NoError(t, orm.IdempotentInsertHead(t.Context(), uncleHead))

	err := orm.TrimOldHeads(t.Context(), 5)
	require.NoError(t, err)

	err = orm.TrimOldHeads(t.Context(), 6)
	require.NoError(t, err)

	err = orm.TrimOldHeads(t.Context(), 7)
	require.NoError(t, err)

	heads, err := orm.LatestHeads(t.Context(), 0)
	require.NoError(t, err)

	// uncle block was loaded too
	require.Len(t, heads, 3)
	for i := 0; i < 3; i++ {
		require.LessOrEqual(t, int64(7), heads[i].Number)
	}
}

func TrimBlocks(blocks []int, minBlockNumber int) []int {
	for i, block := range blocks {
		if block >= minBlockNumber {
			return blocks[i:]
		}
	}
	// All blocks are less than minBlockNumber
	return []int{}
}

var lastTrimmedBlockNumber int64

func trim(minBlockNumber int64, batchSize int64, db []int) []int {
	if lastTrimmedBlockNumber == -1 {
		lastTrimmedBlockNumber = minBlockNumber - 1
	}
	if minBlockNumber-lastTrimmedBlockNumber <= batchSize {
		return db
	}
	lastTrimmedBlockNumber = minBlockNumber - 1
	db = TrimBlocks(db, int(minBlockNumber))
	return db
}

func TestTrim(t *testing.T) {
	lastTrimmedBlockNumber = -1
	db := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	batchSize := int64(2)
	db = trim(1, batchSize, db)
	fmt.Println(1, db)
	db = trim(2, batchSize, db)
	fmt.Println(2, db)
	db = trim(3, batchSize, db)
	fmt.Println(3, db)
	db = trim(4, batchSize, db)
	fmt.Println(4, db)
	db = trim(5, batchSize, db)
	fmt.Println(5, db)
	db = trim(6, batchSize, db)
	fmt.Println(6, db)
	db = trim(7, batchSize, db)
	fmt.Println(7, db)
	db = trim(8, batchSize, db)
	fmt.Println(8, db)
}

func TestORM_HeadByHash(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db)

	var hash common.Hash
	for i := 0; i < 10; i++ {
		head := testutils.Head(i)
		if i == 5 {
			hash = head.Hash
		}
		require.NoError(t, orm.IdempotentInsertHead(tests.Context(t), head))
	}

	head, err := orm.HeadByHash(tests.Context(t), hash)
	require.NoError(t, err)
	require.Equal(t, hash, head.Hash)
	require.Equal(t, int64(5), head.Number)
}

func TestORM_HeadByHash_NotFound(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db)

	hash := testutils.Head(123).Hash
	head, err := orm.HeadByHash(tests.Context(t), hash)

	require.Nil(t, head)
	require.NoError(t, err)
}

func TestORM_LatestHeads_NoRows(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db)

	heads, err := orm.LatestHeads(tests.Context(t), 100)

	require.Empty(t, heads)
	require.NoError(t, err)
}
