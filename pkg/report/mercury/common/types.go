package common

import (
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Report represents a simplified Mercury report which only contains the feedID
// Used to extract the report type, and to choose the correct decoder
type Report struct {
	FeedID [32]byte
}

var schema = GetSchema()

func GetSchema() abi.Arguments {
	mustNewType := func(t string) abi.Type {
		result, err := abi.NewType(t, "", []abi.ArgumentMarshaling{})
		if err != nil {
			panic(fmt.Sprintf("Unexpected error during abi.NewType: %s", err))
		}
		return result
	}
	return abi.Arguments([]abi.Argument{
		{Name: "feedID", Type: mustNewType("bytes32")},
	})
}

// Decode is made available to external users (i.e. mercury server)
func Decode(report []byte) (*Report, error) {
	values, err := schema.Unpack(report)
	if err != nil {
		return nil, fmt.Errorf("failed to decode report: %w", err)
	}
	decoded := new(Report)
	if err = schema.Copy(decoded, values); err != nil {
		return nil, fmt.Errorf("failed to copy report values to struct: %w", err)
	}
	return decoded, nil
}

// GetReportType returns the report type sourced from the feedID
// Notice: Data Stream (Asset DON) feed ID and the Data Stream feed ID are separate type of IDs, with different schemas
func GetReportType(feedID [32]byte) uint16 {
	// Get the first 2 bytes of the feedID
	return binary.BigEndian.Uint16(feedID[:2])
}
