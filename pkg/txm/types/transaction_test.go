package types

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
)

func TestTransaction_GetMeta(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		meta     *sqlutil.JSON
		expected *TxMeta
		wantErr  bool
	}{
		{
			name: "successful parse with all fields",
			meta: func() *sqlutil.JSON {
				meta := TxMeta{
					JobID:               new(int32(123)),
					FailOnRevert:        null.BoolFrom(true),
					RequestID:           &common.Hash{1},
					RequestTxHash:       &common.Hash{2},
					RequestIDs:          []common.Hash{{3}, {4}},
					RequestTxHashes:     []common.Hash{{5}, {6}},
					MaxLink:             new("1000000000000000000"),
					SubID:               new(uint64(5)),
					GlobalSubID:         new("abc123"),
					MaxEth:              new("2000000000000000000"),
					ForceFulfilled:      new(true),
					UpkeepID:            new("7890"),
					WorkflowExecutionID: new("workflow1"),
					FwdrDestAddress:     &common.Address{7},
					MessageIDs:          []string{"msg1", "msg2"},
					SeqNumbers:          []uint64{1, 2},
					DualBroadcast:       new(true),
					DualBroadcastParams: new("params123"),
				}
				b, err := json.Marshal(meta)
				require.NoError(t, err)
				j := sqlutil.JSON(b)
				return &j
			}(),
			expected: &TxMeta{
				JobID:               new(int32(123)),
				FailOnRevert:        null.BoolFrom(true),
				RequestID:           &common.Hash{1},
				RequestTxHash:       &common.Hash{2},
				RequestIDs:          []common.Hash{{3}, {4}},
				RequestTxHashes:     []common.Hash{{5}, {6}},
				MaxLink:             new("1000000000000000000"),
				SubID:               new(uint64(5)),
				GlobalSubID:         new("abc123"),
				MaxEth:              new("2000000000000000000"),
				ForceFulfilled:      new(true),
				UpkeepID:            new("7890"),
				WorkflowExecutionID: new("workflow1"),
				FwdrDestAddress:     &common.Address{7},
				MessageIDs:          []string{"msg1", "msg2"},
				SeqNumbers:          []uint64{1, 2},
				DualBroadcast:       new(true),
				DualBroadcastParams: new("params123"),
			},
			wantErr: false,
		},
		{
			name:     "nil meta returns nil",
			meta:     nil,
			expected: nil,
			wantErr:  false,
		},
		{
			name: "invalid json returns error",
			meta: func() *sqlutil.JSON {
				j := sqlutil.JSON([]byte(`{invalid json`))
				return &j
			}(),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tx := Transaction{
				Meta: tt.meta,
			}

			got, err := tx.GetMeta()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}
