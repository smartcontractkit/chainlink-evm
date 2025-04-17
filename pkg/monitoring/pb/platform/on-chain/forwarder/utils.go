package forwarder

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-evm/pkg/report/datafeeds"
	"github.com/smartcontractkit/chainlink-evm/pkg/report/platform"
	"github.com/test-go/testify/require"
)

func NewTestReport(t *testing.T) []byte {
	feedReport := datafeeds.FeedReport{
		FeedID:    [32]byte{0x01},
		Price:     big.NewInt(1234567890123456789),
		Timestamp: 1620000000,
	}

	reports := &datafeeds.Reports{
		feedReport,
	}

	data, err := datafeeds.GetSchema().Pack(reports)
	require.NoError(t, err)

	report := platform.Report{
		Metadata: datafeeds.Metadata{
			Version:             1,
			WorkflowExecutionID: [32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			Timestamp:           1620000000,
			DonID:               1,
			DonConfigVersion:    1,
			WorkflowCID:         [32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			WorkflowName:        [10]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a},
			WorkflowOwner:       [20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10},
			ReportID:            [2]byte{0x01},
		},
		Data: data,
	}

	encoded, err := report.Encode()
	require.NoError(t, err)

	return encoded
}
