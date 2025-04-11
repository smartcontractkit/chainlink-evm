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

type Reports = []FeedReport

// Define the ABI schema
var schema = GetSchema()

func GetSchema() abi.Arguments {
	// Helper function to simplify error handling when creating new ABI types.
	mustNewType := func(typ string, internalType string, components []abi.ArgumentMarshaling) abi.Type {
		result, err := abi.NewType(typ, internalType, components)
		if err != nil {
			panic(fmt.Sprintf("Unexpected error during abi.NewType: %s", err))
		}
		return result
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

// Decode decodes the provided ABI-encoded data into a Prices slice.
func Decode(data []byte) (*Reports, error) {
	values, err := schema.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}

	var decoded []FeedReport
	if err = schema.Copy(&decoded, values); err != nil {
		return nil, fmt.Errorf("failed to copy report values to struct: %w", err)
	}

	return &decoded, nil
}
