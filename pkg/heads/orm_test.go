package heads_test

import (
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
	orm := heads.NewORM(*testutils.FixtureChainID, db, 0)

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

func TestORM_IdempotentInsertHead_Batch(t *testing.T) {
	t.Parallel()
	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db, 2)

	// Returns nil when inserting first head
	head := testutils.Head(0)
	require.NoError(t, orm.IdempotentInsertHead(t.Context(), head))

	head2 := testutils.Head(1)
	require.NoError(t, orm.IdempotentInsertHead(t.Context(), head2))

	// Head is inserted
	foundHead, err := orm.LatestHead(t.Context())
	require.NoError(t, err)
	assert.Equal(t, head2.Hash, foundHead.Hash)

	// Returns nil when inserting same head again
	require.NoError(t, orm.IdempotentInsertHead(t.Context(), head))

	// Head is still inserted
	foundHead, err = orm.LatestHead(t.Context())
	require.NoError(t, err)
	assert.Equal(t, head2.Hash, foundHead.Hash)
}

func TestORM_TrimOldHeads(t *testing.T) {
	t.Parallel()
	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db, 0)

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
	orm := heads.NewORM(*testutils.FixtureChainID, db, 2)

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

func TestORM_HeadByHash(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db, 0)

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
	orm := heads.NewORM(*testutils.FixtureChainID, db, 0)

	hash := testutils.Head(123).Hash
	head, err := orm.HeadByHash(tests.Context(t), hash)

	require.Nil(t, head)
	require.NoError(t, err)
}

func TestORM_LatestHeads_NoRows(t *testing.T) {
	t.Parallel()

	db := testutils.NewSqlxDB(t)
	orm := heads.NewORM(*testutils.FixtureChainID, db, 0)

	heads, err := orm.LatestHeads(tests.Context(t), 100)

	require.Empty(t, heads)
	require.NoError(t, err)
}
