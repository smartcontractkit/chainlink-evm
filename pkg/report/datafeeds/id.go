package datafeeds

import (
	"encoding/hex"
	"fmt"
)

// FeedID represents a 32-byte feed ID
type FeedID [32]byte

// NewFeedID creates a FeedID from a 32-byte array
func NewFeedID(feedID [32]byte) FeedID {
	return feedID
}

// NewFeedIDFromHex creates a FeedID from a hex string
func NewFeedIDFromHex(feedID string) (FeedID, error) {
	b, err := hexDecodeStringAsByte32(feedID)
	if err != nil {
		return FeedID{}, err
	}
	return NewFeedID(b), nil
}

// hexDecodeStringAsByte32 decodes a hex string into a 32-byte array
func hexDecodeStringAsByte32(s string) ([32]byte, error) {
	var b [32]byte
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return b, err
	}
	if len(decoded) != 32 {
		return b, fmt.Errorf("decoded string is not 32 bytes long")
	}
	copy(b[:], decoded)
	return b, nil
}

// String returns the FeedID as a 0x prefixed hex string
func (id FeedID) String() string {
	return "0x" + hex.EncodeToString(id[:])
}

// GetReportType returns the report type sourced from the feedID
//
// [DF2.0 | Data ID Final Specification](https://docs.google.com/document/d/13ciwTx8lSUfyz1IdETwpxlIVSn1lwYzGtzOBBTpl5Vg/edit?usp=sharing)
// Byte 0: ID Format - 256 options (base case)
//   - Incrementing from 0
//   - 0x00 = current Data Streams format
//   - 0x01 = this format
//   - 0x02 = PoR self-serve SA team allocated IDs (15 bytes)
//   - 0x03 = PoR from feeds team
//   - 0xFF can extend ID format to subsequent bytes, so 0xFF00 is first, then 0xFF01, etc.
func (id FeedID) GetReportType() uint8 {
	// Get the first byte of the feedID
	return id[0]
}

// GetDecimals returns the number of decimals for the feed, derived from the feedID
//
// [DF2.0 | Data ID Final Specification](https://docs.google.com/document/d/13ciwTx8lSUfyz1IdETwpxlIVSn1lwYzGtzOBBTpl5Vg/edit?usp=sharing)
// Byte 7: Data Type - 256 options
//   - Given the variety of buckets, a data type for the buckets will be useful for correct parsing
//   - 0x00 = Boolean
//   - 0x01= String
//   - 0x02 = Address
//   - 0x03 = Bytes
//   - 0x04 = Bundle (Encoded Struct)
//   - 0x05-0x1F reserved
//   - 0x20 = Decimal0 (Integer)
//   - 0x21 = Decimal1 (Float w/ 1 decimal place)
//   - …
//   - 0x28 = Decimal8
//   - …
//   - 0x32 = Decimal18
//   - …
//   - 0x60 = Decimal64
//   - 0x61-0xFF reserved
func (id FeedID) GetDataType() uint8 {
	// Get the 8th byte (index 7) of the feedID
	return id[7]
}

// GetDecimals returns the number of decimals for the fe7], derived from the data type
// Returns false if the data type is not a number
func GetDecimals(dataType uint8) (uint8, bool) {
	if dataType >= 0x20 && dataType <= 0x60 {
		return dataType - 0x20, true
	}
	// Else if the data type is not a number
	return 0, false
}
