// Code generated — DO NOT EDIT.

package bindings

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
	pb2 "github.com/smartcontractkit/chainlink-common/pkg/workflows/sdk/v2/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
)

var EmptyContractMetaData = &bind.MetaData{
	ABI: "[]",
}

// Structs

// Contract Method Inputs

// Errors

// Events

// Main Binding Type for EmptyContract
type EmptyContract struct {
	Address []byte
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   EmptyContractCodec
}

type EmptyContractCodec interface {
}

func NewEmptyContract(
	client *evm.Client,
	address []byte,
	options *bindings.ContractInitOptions,
) (*EmptyContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EmptyContractMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewEmptyContractCodec()
	if err != nil {
		return nil, err
	}
	return &EmptyContract{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type emptyContractCodecImpl struct {
	abi *abi.ABI
}

func NewEmptyContractCodec() (EmptyContractCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(EmptyContractMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &emptyContractCodecImpl{abi: &parsed}, nil
}

func (c *EmptyContract) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	default:
		return nil, errors.New("unknown error selector")
	}
}
