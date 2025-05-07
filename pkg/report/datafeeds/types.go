package datafeeds

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// This is ABI encoding - abi: "(bytes32 FeedID, bytes RawReport)[] Reports" (set in workflow)
// Encoded with: https://github.com/smartcontractkit/chainlink/blob/develop/core/services/relay/evm/cap_encoder.go
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

type Reports = []FeedReport
type CCIPReports = []CCIPFeedReport

func GetSchema(ccip bool) abi.Arguments {
	// Helper function to simplify error handling when creating new ABI types.
	mustNewType := func(typ string, internalType string, components []abi.ArgumentMarshaling) abi.Type {
		result, err := abi.NewType(typ, internalType, components)
		if err != nil {
			panic(fmt.Sprintf("Unexpected error during abi.NewType: %s", err))
		}
		return result
	}

	// need special handling for ccip DF since price and timestamp are swapped
	if ccip {
		return abi.Arguments([]abi.Argument{
			{
				// This defines the array of tuple records.
				Type: mustNewType("tuple(bytes32,uint224,uint32)[]", "", []abi.ArgumentMarshaling{
					{Name: "feedID", Type: "bytes32"},
					{Name: "price", Type: "uint224"},
					{Name: "timestamp", Type: "uint32"},
				}),
			},
		})
	}

	return abi.Arguments([]abi.Argument{
		{
			// This defines the array of tuple records.
			Type: mustNewType("tuple(bytes32,uint32,uint224)[]", "", []abi.ArgumentMarshaling{
				{Name: "feedID", Type: "bytes32"},
				{Name: "timestamp", Type: "uint32"},
				{Name: "price", Type: "uint224"},
			}),
		},
	})
}

func Decode(data []byte, ccip bool) (*Reports, error) {
	schema := GetSchema(ccip)

	// Unpack into a raw slice of tuples
	values, err := schema.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	// CCIP branch: copy into CCIPReports then remap into FeedReport
	if ccip {
		var ccipRaw CCIPReports
		if err := schema.Copy(&ccipRaw, values); err != nil {
			return nil, fmt.Errorf("failed to copy CCIP reports: %w", err)
		}

		out := make(Reports, len(ccipRaw))
		for i, r := range ccipRaw {
			out[i] = FeedReport{
				FeedID:    r.FeedID,
				Timestamp: r.Timestamp,
				Price:     r.Price,
			}
		}
		return &out, nil
	}

	// Normal DF branch: copy directly into FeedReport slice
	var out Reports
	if err := schema.Copy(&out, values); err != nil {
		return nil, fmt.Errorf("failed to copy normal reports: %w", err)
	}
	return &out, nil
}
