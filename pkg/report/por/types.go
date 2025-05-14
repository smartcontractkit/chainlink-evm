//nolint:revive // disable revive
package por

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Report struct {
	DataId    [32]byte
	Timestamp uint32
	Bundle    []byte
}

type Reports = []Report

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
		{
			Type: mustNewType("tuple(bytes32,uint32,bytes)[]", "", []abi.ArgumentMarshaling{
				{Name: "dataId", Type: "bytes32"},
				{Name: "timestamp", Type: "uint32"},
				{Name: "bundle", Type: "bytes"},
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

	var decoded []Report
	if err = schema.Copy(&decoded, values); err != nil {
		return nil, fmt.Errorf("failed to copy report values to struct: %w", err)
	}

	return &decoded, nil
}
