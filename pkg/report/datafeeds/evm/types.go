package evm

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"
)

type Report struct {
	// TODO: replace with common type
	types.Metadata
	Data []byte
}

// FeedPrice represents a decoded data record with the new types.
type FeedReport struct {
	FeedID    [32]byte
	Price     *big.Int // *big.Int is used because go-ethereum converts large uints to *big.Int.
	Timestamp uint32
}

type Reports = []FeedReport

// Define the ABI schema for our tuple: (bytes32, uint224, uint32)[].
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
			Type: mustNewType("tuple(bytes32,uint224,uint32)[]", "", []abi.ArgumentMarshaling{
				{Name: "feedID", Type: "bytes32"},
				{Name: "price", Type: "uint224"},
				{Name: "timestamp", Type: "uint32"},
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

func (r Report) Encode() ([]byte, error) {
	// Encode the metadata
	metadataBytes, err := r.Metadata.Encode()
	if err != nil {
		return nil, err
	}

	return append(metadataBytes, r.Data...), nil
}
