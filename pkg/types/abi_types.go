package types

import (
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate go run ./gen/main.go

var typeMap = map[string]*ABIEncodingType{
	"bool": {
		native:  reflect.TypeFor[bool](),
		checked: reflect.TypeFor[bool](),
	},
	"int8": {
		native:  reflect.TypeFor[int8](),
		checked: reflect.TypeFor[int8](),
	},
	"int16": {
		native:  reflect.TypeFor[int16](),
		checked: reflect.TypeFor[int16](),
	},
	"int32": {
		native:  reflect.TypeFor[int32](),
		checked: reflect.TypeFor[int32](),
	},
	"int64": {
		native:  reflect.TypeFor[int64](),
		checked: reflect.TypeFor[int64](),
	},
	"uint8": {
		native:  reflect.TypeFor[uint8](),
		checked: reflect.TypeFor[uint8](),
	},
	"uint16": {
		native:  reflect.TypeFor[uint16](),
		checked: reflect.TypeFor[uint16](),
	},
	"uint32": {
		native:  reflect.TypeFor[uint32](),
		checked: reflect.TypeFor[uint32](),
	},
	"uint64": {
		native:  reflect.TypeFor[uint64](),
		checked: reflect.TypeFor[uint64](),
	},
	"string": {
		native:  reflect.TypeFor[string](),
		checked: reflect.TypeFor[string](),
	},
	"address": {
		native:  reflect.TypeFor[common.Address](),
		checked: reflect.TypeFor[common.Address](),
	},
	"bytes": {
		native:  reflect.TypeFor[[]byte](),
		checked: reflect.TypeFor[[]byte](),
	},
}

type ABIEncodingType struct {
	native  reflect.Type
	checked reflect.Type
}

func GetAbiEncodingType(name string) (*ABIEncodingType, bool) {
	abiType, ok := typeMap[name]
	return abiType, ok
}
