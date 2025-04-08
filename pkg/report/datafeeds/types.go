package datafeeds

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// This is ABI encoding - abi: "(bytes32 FeedID, bytes RawReport)[] Reports" (set in workflow)
// Encoded with: https://github.com/smartcontractkit/chainlink/blob/develop/core/services/relay/evm/cap_encoder.go
type FeedReport struct {
	FeedID [32]byte
	Data   []byte
}

type Reports = []FeedReport

// Define the ABI schema
var schema = GetSchema()

func GetSchema() abi.Arguments {
	mustNewType := func(t string, internalType string, components []abi.ArgumentMarshaling) abi.Type {
		result, err := abi.NewType(t, internalType, components)
		if err != nil {
			panic(fmt.Sprintf("Unexpected error during abi.NewType: %s", err))
		}
		return result
	}

	return abi.Arguments([]abi.Argument{
		// TODO: why is the workflow encoder_config "(bytes32 FeedID, bytes RawReport)[] Reports"?
		{
			Type: mustNewType("tuple(bytes32, bytes)[]", "", []abi.ArgumentMarshaling{
				{Name: "feedID", Type: "bytes32"},
				{Name: "data", Type: "bytes"},
			}),
		},
	})
}

// Decode is made available to external users (i.e. mercury server)
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
