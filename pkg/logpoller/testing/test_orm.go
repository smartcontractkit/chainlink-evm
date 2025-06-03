package testing

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
)

type TestDSORM struct {
	ds sqlutil.DataSource
}

// NewTestORM creates a test DSORM which contains method only used by tests
func NewTestORM(ds sqlutil.DataSource) *TestDSORM {
	return &TestDSORM{
		ds: ds,
	}
}

// HasFilterByEventSig checks if a filter exists for the provided event sig and contract address
func (o *TestDSORM) HasFilterByEventSig(ctx context.Context, chainID string, eventSig common.Hash, address []byte) (bool, error) {
	query := `SELECT COUNT(1) FROM evm.log_poller_filters
		WHERE evm_chain_id = $1 AND event = $2 AND address = $3 LIMIT 1`

	var exists int
	if err := o.ds.GetContext(ctx, &exists, query, chainID, eventSig.Bytes(), address); err != nil {
		return false, err
	}

	return exists != 0, nil
}
