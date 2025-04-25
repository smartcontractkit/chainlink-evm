package forwarder

import (
	"math/big"
	"testing"

	ocr3types "github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"

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
		Metadata: ocr3types.Metadata{
			Version:          1,
			ExecutionID:      "0102030405060708090a0b0c0d0e0f1000000000000000000000000000000000",
			Timestamp:        1620000000,
			DONID:            1,
			DONConfigVersion: 1,
			WorkflowID:       "1234567890123456789012345678901234567890123456789012345678901234",
			WorkflowName:     "12",
			WorkflowOwner:    "1234567890123456789012345678901234567890",
			ReportID:         "1234",
		},
		Data: data,
	}

	encoded, err := report.Encode()
	require.NoError(t, err)

	return encoded
}
