package datafeeds

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/platform/on-chain/forwarder"
)

func NewEVMTestReport(t *testing.T, ccip bool) []byte {
	feedReport := FeedReport{
		FeedID:    [32]byte{0x01},
		Price:     big.NewInt(1234567890123456789),
		Timestamp: 1620000000,
	}

	reports := &Reports{
		feedReport,
	}

	data, err := GetSchema(ccip).Pack(reports)
	require.NoError(t, err)

	encoded, err := forwarder.NewTestReport(t, data)
	require.NoError(t, err)

	return encoded
}
