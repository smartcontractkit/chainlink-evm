package datafeeds

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-framework/capabilities/writetarget/monitoring/pb/platform/on-chain/forwarder"
)

type FeedReport struct {
	FeedID    [32]byte
	Timestamp uint32
	Price     *big.Int // *big.Int is used because go-ethereum converts large uints to *big.Int.
}

type CCIPFeedReport struct {
	FeedID    [32]byte
	Price     *big.Int
	Timestamp uint32
}

type PORReport struct {
	DataID    [32]byte
	Timestamp uint32
	Bundle    []byte
}

type Reports = []FeedReport
type CCIPReports = []CCIPFeedReport
type PORReports = []PORReport

func NewEVMTestReport(t *testing.T, ccip bool) []byte {
	feedReport := FeedReport{
		FeedID:    [32]byte{0x01},
		Price:     big.NewInt(1234567890123456789),
		Timestamp: 1620000000,
	}

	reports := &Reports{
		feedReport,
	}

	var schema abi.Arguments
	if ccip {
		schema = GetSchema("tuple(bytes32,uint32,uint224)[]", "", []abi.ArgumentMarshaling{
			{Name: "feedID", Type: "bytes32"},
			{Name: "price", Type: "uint224"},
			{Name: "timestamp", Type: "uint32"},
		})
	} else {
		schema = GetSchema("tuple(bytes32,uint32,uint224)[]", "", []abi.ArgumentMarshaling{
			{Name: "feedID", Type: "bytes32"},
			{Name: "timestamp", Type: "uint32"},
			{Name: "price", Type: "uint224"},
		})
	}
	data, err := schema.Pack(reports)
	require.NoError(t, err)

	encoded, err := forwarder.NewTestReport(t, data)
	require.NoError(t, err)

	return encoded
}
