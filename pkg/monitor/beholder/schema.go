//nolint:gosimple // disable gosimple
package monitor

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// TODO: move to another directory / repo

const (
	AttrKeyBeholderDataSchema = "beholder_data_schema"
	AttrKeyBeholderEntity     = "beholder_entity"
	AttrKeyBeholderDomain     = "beholder_domain"
)

// patternSnake is a regular expression to match CamelCase words
// Notice: we use the Unicode property 'Lu' (uppercase letter) to match
// the first letter of the word, and 'P{Lu}' (not uppercase letter) to match
// the rest of the word.
var patternSnake = regexp.MustCompile("(\\p{Lu}+\\P{Lu}*)")

// toSnakeCase converts a CamelCase to snake_case (used for type -> file name mapping)
func toSnakeCase(s string) string {
	s = patternSnake.ReplaceAllString(s, "_${1}")
	s, _ = strings.CutPrefix(strings.ToLower(s), "_")
	return s
}

// toSchemaName returns a protobuf message name (short)
func toSchemaName(m proto.Message) string {
	return string(protoimpl.X.MessageTypeOf(m).Descriptor().Name())
}

// toSchemaName returns a protobuf message name (full)
func toSchemaFullName(m proto.Message) string {
	return string(protoimpl.X.MessageTypeOf(m).Descriptor().FullName())
}

// toSchemaPath maps a protobuf message to a Beholder schema path
func toSchemaPath(m proto.Message, basePath string) string {
	// Notice: a name like 'platform.on_chain.forwarder.ReportProcessed'
	protoName := toSchemaFullName(m)

	// We map to a Beholder schema path like '<basePath>/platform/on-chain/forwarder/report_processed.proto'
	protoPath := protoName
	protoPath = strings.ReplaceAll(protoPath, ".", "/")
	protoPath = strings.ReplaceAll(protoPath, "_", "-")

	// Split the path components (at least one component)
	pp := strings.Split(protoPath, "/")
	pp[len(pp)-1] = toSnakeCase(pp[len(pp)-1])

	// Join the path components again
	protoPath = strings.Join(pp, "/")
	protoPath = fmt.Sprintf("%s.proto", protoPath)

	// Return the full schema path
	return path.Join(basePath, protoPath)
}

// appendRequiredAttrDataSchema adds the message schema path as an attribute (required)
func appendRequiredAttrDataSchema(m proto.Message, attrKVs []any, basePath string) []any {
	key := AttrKeyBeholderDataSchema
	for i := 0; i < len(attrKVs); i += 2 {
		if attrKVs[i] == key {
			return attrKVs
		}
	}

	attrKVs = append(attrKVs, key)
	// Needs to be an URI (Beholder requirement)
	val := toSchemaPath(m, basePath)
	attrKVs = append(attrKVs, val)
	return attrKVs
}

// appendRequiredAttrEntity adds the message entity type as an attribute (required)
func appendRequiredAttrEntity(m proto.Message, attrKVs []any) []any {
	key := AttrKeyBeholderEntity
	for i := 0; i < len(attrKVs); i += 2 {
		if attrKVs[i] == key {
			return attrKVs
		}
	}

	attrKVs = append(attrKVs, key)
	attrKVs = append(attrKVs, toSchemaName(m))
	return attrKVs
}

// appendRequiredAttrDomain adds the message domain as an attribute (required)
func appendRequiredAttrDomain(m proto.Message, attrKVs []any) []any {
	key := AttrKeyBeholderDomain
	for i := 0; i < len(attrKVs); i += 2 {
		if attrKVs[i] == key {
			return attrKVs
		}
	}

	// Notice: a name like 'platform.on_chain.forwarder.ReportProcessed'
	protoName := toSchemaFullName(m)

	// Extract first path component (entrypoint package) as a domain
	domain := "unknown"
	if strings.Contains(protoName, ".") {
		domain = strings.Split(protoName, ".")[0]
	}

	attrKVs = append(attrKVs, key)
	attrKVs = append(attrKVs, domain)
	return attrKVs
}
