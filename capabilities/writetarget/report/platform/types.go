package platform

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"
)

// Report represents an OCR3 report with metadata and data
// version | workflow_execution_id | timestamp | don_id | config_version | ... | data
type Report struct {
	types.Metadata
	Data []byte
}

// Decode decodes a report from raw bytes
func Decode(rawReport []byte) (*Report, error) {
	if len(rawReport) < 109 {
		return nil, fmt.Errorf("invalid report length")
	}

	report := &Report{}
	buf := bytes.NewReader(rawReport)

	// Decode version
	var versionByte byte
	if err := binary.Read(buf, binary.BigEndian, &versionByte); err != nil {
		return nil, err
	}
	report.Version = uint32(versionByte)

	// Notice: we only support version 1 currently
	if report.Version != 1 {
		return nil, fmt.Errorf("invalid report version: %d", report.Version)
	}

	// Decode workflow_execution_id
	var workflowExecutionIDBytes [32]byte
	if _, err := buf.Read(workflowExecutionIDBytes[:]); err != nil {
		return nil, err
	}
	// TODO: should we prefix with 0x?
	report.ExecutionID = hex.EncodeToString(workflowExecutionIDBytes[:])

	// Decode timestamp
	var timestampBytes [4]byte
	if _, err := buf.Read(timestampBytes[:]); err != nil {
		return nil, err
	}
	report.Timestamp = binary.BigEndian.Uint32(timestampBytes[:])

	// Decode don_id
	var donIDBytes [4]byte
	if _, err := buf.Read(donIDBytes[:]); err != nil {
		return nil, err
	}
	report.DONID = binary.BigEndian.Uint32(donIDBytes[:])

	// Decode config_version
	var configVersionBytes [4]byte
	if _, err := buf.Read(configVersionBytes[:]); err != nil {
		return nil, err
	}
	report.DONConfigVersion = binary.BigEndian.Uint32(configVersionBytes[:])

	// Decode workflow_id
	var workflowIDBytes [32]byte
	if _, err := buf.Read(workflowIDBytes[:]); err != nil {
		return nil, err
	}
	// TODO: should we prefix with 0x?
	report.WorkflowID = hex.EncodeToString(workflowIDBytes[:])

	// Decode workflow_name (UTF-8)
	var workflowNameBytes [10]byte
	if _, err := buf.Read(workflowNameBytes[:]); err != nil {
		return nil, err
	}
	report.WorkflowName = hex.EncodeToString(workflowNameBytes[:])

	// Decode workflow_owner
	var workflowOwnerBytes [20]byte
	if _, err := buf.Read(workflowOwnerBytes[:]); err != nil {
		return nil, err
	}
	// TODO: should we prefix with 0x?
	report.WorkflowOwner = hex.EncodeToString(workflowOwnerBytes[:])

	// Decode report_id
	var reportIDBytes [2]byte
	if _, err := buf.Read(reportIDBytes[:]); err != nil {
		return nil, err
	}
	// TODO: should we prefix with 0x?
	report.ReportID = hex.EncodeToString(reportIDBytes[:])

	// Decode data
	report.Data = make([]byte, buf.Len())
	if len(report.Data) == 0 {
		// No data to read
		return report, nil
	}
	if _, err := buf.Read(report.Data); err != nil {
		return nil, err
	}

	return report, nil
}
