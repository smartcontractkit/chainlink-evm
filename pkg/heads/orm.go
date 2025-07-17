package heads

import (
	"context"
	"database/sql"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	pkgerrors "github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"

	evmtypes "github.com/smartcontractkit/chainlink-evm/pkg/types"
	ubig "github.com/smartcontractkit/chainlink-evm/pkg/utils/big"
)

type ORM interface {
	// IdempotentInsertHead inserts a head only if the hash is new. Will do nothing if hash exists already.
	// No advisory lock required because this is thread safe.
	IdempotentInsertHead(ctx context.Context, head *evmtypes.Head) error
	// TrimOldHeads deletes heads such that only blocks >= minBlockNumber remain
	TrimOldHeads(ctx context.Context, minBlockNumber int64) (err error)
	// LatestHead returns the highest seen head
	LatestHead(ctx context.Context) (head *evmtypes.Head, err error)
	// LatestHeads returns the latest heads with blockNumbers >= minBlockNumber
	LatestHeads(ctx context.Context, minBlockNumber int64) (heads []*evmtypes.Head, err error)
	// HeadByHash fetches the head with the given hash from the db, returns nil if none exists
	HeadByHash(ctx context.Context, hash common.Hash) (head *evmtypes.Head, err error)
}

var _ ORM = &DbORM{}

type DbORM struct {
	chainID                ubig.Big
	ds                     sqlutil.DataSource
	lastTrimmedBlockNumber int64            // the last block number that was trimmed
	headsBatch             []*evmtypes.Head // used to batch insert heads
	mu                     sync.RWMutex
	batchSize              int64 // the number of heads to batch insert/delete
}

// NewORM creates an ORM scoped to chainID.
func NewORM(chainID big.Int, ds sqlutil.DataSource, batchSize int64) *DbORM {
	return &DbORM{
		chainID:                ubig.Big(chainID),
		ds:                     ds,
		lastTrimmedBlockNumber: -1,
		headsBatch:             make([]*evmtypes.Head, 0),
		batchSize:              batchSize,
	}
}

func (orm *DbORM) setHeadsBatch(appendHead *evmtypes.Head) []*evmtypes.Head {
	orm.mu.Lock()
	defer orm.mu.Unlock()

	// if we are appending a head, we need to check if the batch is big enough to insert
	// if it is, copy it and return it
	if appendHead != nil {
		orm.headsBatch = append(orm.headsBatch, appendHead)
		if len(orm.headsBatch) < int(orm.batchSize) {
			// Not enough to batch insert yet
			return nil
		}
		copied := make([]*evmtypes.Head, len(orm.headsBatch))
		copy(copied, orm.headsBatch)

		return copied
	}
	// this will be only called to reset headsBatch
	orm.headsBatch = orm.headsBatch[:0]
	return nil
}

func (orm *DbORM) IdempotentInsertHead(ctx context.Context, head *evmtypes.Head) error {
	batch := orm.setHeadsBatch(head)

	if batch == nil {
		// Not enough to batch insert yet
		return nil
	}

	// listener guarantees head.EVMChainID to be equal to DbORM.chainID
	query := `INSERT INTO evm.heads
				(hash, number, parent_hash, created_at, timestamp, l1_block_number, evm_chain_id, base_fee_per_gas)
			VALUES (:hash, :number, :parent_hash, NOW(), :timestamp, :l1_block_number, :evm_chain_id, :base_fee_per_gas)
				ON CONFLICT DO NOTHING`

	_, err := orm.ds.NamedExecContext(ctx, query, batch)
	if err != nil {
		return pkgerrors.Wrap(err, "IdempotentInsertHead failed to insert heads")
	}
	// reset the heads batch
	orm.setHeadsBatch(nil)

	return nil
}

// the return value tells the caller if the batch is big enough to trim
func (orm *DbORM) setLastTrimmedBlockNumber(minBlockNumber int64) bool {
	orm.mu.Lock()
	defer orm.mu.Unlock()
	if orm.lastTrimmedBlockNumber == -1 {
		// we delete everything before the minBlockNumber, so we need to set the lastTrimmedBlockNumber to the block before the minBlockNumber
		orm.lastTrimmedBlockNumber = minBlockNumber - 1
	}
	if minBlockNumber-orm.lastTrimmedBlockNumber <= orm.batchSize {
		// Batch not big enough to trim yet
		return false
	}
	// we delete everything before the minBlockNumber, so we need to set the lastTrimmedBlockNumber to the block before the minBlockNumber
	orm.lastTrimmedBlockNumber = minBlockNumber - 1
	return true
}

func (orm *DbORM) TrimOldHeads(ctx context.Context, minBlockNumber int64) (err error) {
	shouldTrim := orm.setLastTrimmedBlockNumber(minBlockNumber)
	if !shouldTrim {
		return nil
	}
	query := `DELETE FROM evm.heads WHERE evm_chain_id = $1 AND number < $2`
	_, err = orm.ds.ExecContext(ctx, query, orm.chainID, minBlockNumber)

	return err
}

func (orm *DbORM) LatestHead(ctx context.Context) (head *evmtypes.Head, err error) {
	head = new(evmtypes.Head)
	err = orm.ds.GetContext(ctx, head, `SELECT * FROM evm.heads WHERE evm_chain_id = $1 ORDER BY number DESC, created_at DESC, id DESC LIMIT 1`, orm.chainID)
	if pkgerrors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	err = pkgerrors.Wrap(err, "LatestHead failed")
	return
}

func (orm *DbORM) LatestHeads(ctx context.Context, minBlockNumer int64) (heads []*evmtypes.Head, err error) {
	err = orm.ds.SelectContext(ctx, &heads, `SELECT * FROM evm.heads WHERE evm_chain_id = $1 AND number >= $2 ORDER BY number DESC, created_at DESC, id DESC`, orm.chainID, minBlockNumer)
	err = pkgerrors.Wrap(err, "LatestHeads failed")
	return
}

func (orm *DbORM) HeadByHash(ctx context.Context, hash common.Hash) (head *evmtypes.Head, err error) {
	head = new(evmtypes.Head)
	err = orm.ds.GetContext(ctx, head, `SELECT * FROM evm.heads WHERE evm_chain_id = $1 AND hash = $2`, orm.chainID, hash)
	if pkgerrors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return head, err
}

type nullORM struct{}

func NewNullORM() ORM {
	return &nullORM{}
}

func (orm *nullORM) IdempotentInsertHead(ctx context.Context, head *evmtypes.Head) error {
	return nil
}

func (orm *nullORM) TrimOldHeads(ctx context.Context, minBlockNumber int64) (err error) {
	return nil
}

func (orm *nullORM) LatestHead(ctx context.Context) (head *evmtypes.Head, err error) {
	return nil, nil
}

func (orm *nullORM) LatestHeads(ctx context.Context, minBlockNumer int64) (heads []*evmtypes.Head, err error) {
	return nil, nil
}

func (orm *nullORM) HeadByHash(ctx context.Context, hash common.Hash) (head *evmtypes.Head, err error) {
	return nil, nil
}
