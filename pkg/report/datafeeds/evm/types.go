package evm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Report struct {
	// TODO: replace with common type
	ReportV1Metadata
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

func decodeReportMetadata(data []byte) (metadata ReportV1Metadata, err error) {
	if len(data) < metadata.Length() {
		return metadata, fmt.Errorf("data too short: %d bytes", len(data))
	}
	return metadata, binary.Read(bytes.NewReader(data[:metadata.Length()]), binary.BigEndian, &metadata)
}

func (rm ReportV1Metadata) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, rm)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (rm ReportV1Metadata) Length() int {
	bytes, err := rm.Encode()
	if err != nil {
		return 0
	}
	return len(bytes)
}

func (r Report) Encode() ([]byte, error) {
	// Encode the metadata
	metadataBytes, err := r.ReportV1Metadata.Encode()
	if err != nil {
		return nil, err
	}

	return append(metadataBytes, r.Data...), nil
}

// TODO: replace with https://github.com/smartcontractkit/chainlink-common/blob/39bc061d09ded8c6b87ff95ffaea53110a742f87/pkg/capabilities/consensus/ocr3/types/aggregator.go#L14-L24
type ReportV1Metadata struct {
	Version             uint8
	WorkflowExecutionID [32]byte
	Timestamp           uint32
	DonID               uint32
	DonConfigVersion    uint32
	WorkflowCID         [32]byte
	WorkflowName        [10]byte
	WorkflowOwner       [20]byte
	ReportID            [2]byte
}
