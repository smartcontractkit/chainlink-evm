package platform

import (
	"fmt"

	ocr3types "github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/types"
)

// Report represents an OCR3 report with metadata and data
// version | workflow_execution_id | timestamp | don_id | config_version | ... | data
type Report struct {
	ocr3types.Metadata
	Data []byte
}

func Decode(raw []byte) (*Report, error) {
	md, tail, err := ocr3types.Decode(raw)
	if err != nil {
		return nil, err
	}
	if md.Version != 1 {
		return nil, fmt.Errorf("unsupported version %d", md.Version)
	}
	return &Report{Metadata: md, Data: tail}, nil
}

func (r Report) Encode() ([]byte, error) {
	// Encode the metadata
	metadataBytes, err := r.Metadata.Encode()
	if err != nil {
		return nil, err
	}

	return append(metadataBytes, r.Data...), nil
}
