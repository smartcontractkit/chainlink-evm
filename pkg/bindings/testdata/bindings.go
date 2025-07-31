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

var DataStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"DataNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"DataNotFound2\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"AccessLogged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"indexed\":true,\"internalType\":\"structDataStorage.UserData\",\"name\":\"userData\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sender\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes[]\",\"name\":\"metadataArray\",\"type\":\"bytes[]\"}],\"name\":\"DynamicEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NoFields\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReserve\",\"type\":\"uint256\"}],\"internalType\":\"structDataStorage.UpdateReserves\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"logAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"onReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"readData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structDataStorage.UserData\",\"name\":\"userData\",\"type\":\"tuple\"}],\"name\":\"storeUserData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"updateData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"oldValue\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Structs
type DataStorageUpdateReserves struct {
	TotalMinted  *big.Int
	TotalReserve *big.Int
}

type DataStorageUserData struct {
	Key   string
	Value string
}

// Contract Method Inputs
type LogAccessInput struct {
	Message string
}

type OnReportInput struct {
	Metadata []byte
	Payload  []byte
}

type ReadDataInput struct {
	User common.Address
	Key  string
}

type StoreDataInput struct {
	Key   string
	Value string
}

type StoreUserDataInput struct {
	UserData DataStorageUserData
}

type UpdateDataInput struct {
	Key      string
	NewValue string
}

// Errors
type DataNotFound struct {
	Requester common.Address
	Key       string
	Reason    string
}

type DataNotFound2 struct {
	Requester common.Address
	Key       string
	Reason    string
}

// Events
type AccessLogged struct {
	Caller  common.Address
	Message string
}

type DataStored struct {
	Sender common.Address
	Key    string
	Value  string
}

type DynamicEvent struct {
	Key           string
	UserData      DataStorageUserData
	Sender        string
	Metadata      common.Hash
	MetadataArray [][]byte
}

type NoFields struct {
}

// Main Binding Type for DataStorage
type DataStorage struct {
	Address []byte
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   DataStorageCodec
}

type DataStorageCodec interface {
	EncodeGetReservesMethodCall() ([]byte, error)
	DecodeGetReservesMethodOutput(data []byte) (DataStorageUpdateReserves, error)
	EncodeGetValueMethodCall() ([]byte, error)
	DecodeGetValueMethodOutput(data []byte) (string, error)
	EncodeLogAccessMethodCall(in LogAccessInput) ([]byte, error)
	EncodeOnReportMethodCall(in OnReportInput) ([]byte, error)
	EncodeReadDataMethodCall(in ReadDataInput) ([]byte, error)
	DecodeReadDataMethodOutput(data []byte) (string, error)
	EncodeStoreDataMethodCall(in StoreDataInput) ([]byte, error)
	EncodeStoreUserDataMethodCall(in StoreUserDataInput) ([]byte, error)
	EncodeUpdateDataMethodCall(in UpdateDataInput) ([]byte, error)
	DecodeUpdateDataMethodOutput(data []byte) (string, error)
	EncodeDataStorageUpdateReservesStruct(in DataStorageUpdateReserves) ([]byte, error)
	EncodeDataStorageUserDataStruct(in DataStorageUserData) ([]byte, error)
	AccessLoggedLogHash() []byte
	EncodeAccessLoggedTopics(evt abi.Event, values []AccessLogged) ([]*evm.TopicValues, error)
	DecodeAccessLogged(log *evm.Log) (*AccessLogged, error)
	DataStoredLogHash() []byte
	EncodeDataStoredTopics(evt abi.Event, values []DataStored) ([]*evm.TopicValues, error)
	DecodeDataStored(log *evm.Log) (*DataStored, error)
	DynamicEventLogHash() []byte
	EncodeDynamicEventTopics(evt abi.Event, values []DynamicEvent) ([]*evm.TopicValues, error)
	DecodeDynamicEvent(log *evm.Log) (*DynamicEvent, error)
	NoFieldsLogHash() []byte
	EncodeNoFieldsTopics(evt abi.Event, values []NoFields) ([]*evm.TopicValues, error)
	DecodeNoFields(log *evm.Log) (*NoFields, error)
}

func NewDataStorage(
	client *evm.Client,
	address []byte,
	options *bindings.ContractInitOptions,
) (*DataStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(DataStorageMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewDataStorageCodec()
	if err != nil {
		return nil, err
	}
	return &DataStorage{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type dataStorageCodecImpl struct {
	abi *abi.ABI
}

func NewDataStorageCodec() (DataStorageCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(DataStorageMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &dataStorageCodecImpl{abi: &parsed}, nil
}

func (c *dataStorageCodecImpl) EncodeGetReservesMethodCall() ([]byte, error) {
	return c.abi.Pack("getReserves")
}

func (c *dataStorageCodecImpl) DecodeGetReservesMethodOutput(data []byte) (DataStorageUpdateReserves, error) {
	vals, err := c.abi.Methods["getReserves"].Outputs.Unpack(data)
	if err != nil {
		return *new(DataStorageUpdateReserves), err
	}
	return vals[0].(DataStorageUpdateReserves), nil
}

func (c *dataStorageCodecImpl) EncodeGetValueMethodCall() ([]byte, error) {
	return c.abi.Pack("getValue")
}

func (c *dataStorageCodecImpl) DecodeGetValueMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["getValue"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	return vals[0].(string), nil
}

func (c *dataStorageCodecImpl) EncodeLogAccessMethodCall(in LogAccessInput) ([]byte, error) {
	return c.abi.Pack("logAccess", in.Message)
}

func (c *dataStorageCodecImpl) EncodeOnReportMethodCall(in OnReportInput) ([]byte, error) {
	return c.abi.Pack("onReport", in.Metadata, in.Payload)
}

func (c *dataStorageCodecImpl) EncodeReadDataMethodCall(in ReadDataInput) ([]byte, error) {
	return c.abi.Pack("readData", in.User, in.Key)
}

func (c *dataStorageCodecImpl) DecodeReadDataMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["readData"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	return vals[0].(string), nil
}

func (c *dataStorageCodecImpl) EncodeStoreDataMethodCall(in StoreDataInput) ([]byte, error) {
	return c.abi.Pack("storeData", in.Key, in.Value)
}

func (c *dataStorageCodecImpl) EncodeStoreUserDataMethodCall(in StoreUserDataInput) ([]byte, error) {
	return c.abi.Pack("storeUserData", in.UserData)
}

func (c *dataStorageCodecImpl) EncodeUpdateDataMethodCall(in UpdateDataInput) ([]byte, error) {
	return c.abi.Pack("updateData", in.Key, in.NewValue)
}

func (c *dataStorageCodecImpl) DecodeUpdateDataMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["updateData"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	return vals[0].(string), nil
}

func (c *dataStorageCodecImpl) EncodeDataStorageUpdateReservesStruct(in DataStorageUpdateReserves) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "totalMinted", Type: "uint256"},
			{Name: "totalReserve", Type: "uint256"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for DataStorageUpdateReserves: %w", err)
	}
	args := abi.Arguments{
		{Name: "dataStorageUpdateReserves", Type: tupleType},
	}

	return args.Pack(in)
}
func (c *dataStorageCodecImpl) EncodeDataStorageUserDataStruct(in DataStorageUserData) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "key", Type: "string"},
			{Name: "value", Type: "string"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for DataStorageUserData: %w", err)
	}
	args := abi.Arguments{
		{Name: "dataStorageUserData", Type: tupleType},
	}

	return args.Pack(in)
}

func (c *dataStorageCodecImpl) AccessLoggedLogHash() []byte {
	return c.abi.Events["AccessLogged"].ID.Bytes()
}

func (c *dataStorageCodecImpl) EncodeAccessLoggedTopics(
	evt abi.Event,
	values []AccessLogged,
) ([]*evm.TopicValues, error) {
	var callerRule []interface{}
	for _, v := range values {
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Caller)
		if err != nil {
			return nil, err
		}
		callerRule = append(callerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		callerRule,
	)
	if err != nil {
		return nil, err
	}

	topics := make([]*evm.TopicValues, len(rawTopics)+1)
	topics[0] = &evm.TopicValues{
		Values: [][]byte{evt.ID.Bytes()},
	}
	for i, hashList := range rawTopics {
		bs := make([][]byte, len(hashList))
		for j, h := range hashList {
			bs[j] = h.Bytes()
		}
		topics[i+1] = &evm.TopicValues{Values: bs}
	}
	return topics, nil
}

// DecodeAccessLogged decodes a log into a AccessLogged struct.
func (c *dataStorageCodecImpl) DecodeAccessLogged(log *evm.Log) (*AccessLogged, error) {
	event := new(AccessLogged)
	if err := c.abi.UnpackIntoInterface(event, "AccessLogged", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["AccessLogged"].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *dataStorageCodecImpl) DataStoredLogHash() []byte {
	return c.abi.Events["DataStored"].ID.Bytes()
}

func (c *dataStorageCodecImpl) EncodeDataStoredTopics(
	evt abi.Event,
	values []DataStored,
) ([]*evm.TopicValues, error) {
	var senderRule []interface{}
	for _, v := range values {
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Sender)
		if err != nil {
			return nil, err
		}
		senderRule = append(senderRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		senderRule,
	)
	if err != nil {
		return nil, err
	}

	topics := make([]*evm.TopicValues, len(rawTopics)+1)
	topics[0] = &evm.TopicValues{
		Values: [][]byte{evt.ID.Bytes()},
	}
	for i, hashList := range rawTopics {
		bs := make([][]byte, len(hashList))
		for j, h := range hashList {
			bs[j] = h.Bytes()
		}
		topics[i+1] = &evm.TopicValues{Values: bs}
	}
	return topics, nil
}

// DecodeDataStored decodes a log into a DataStored struct.
func (c *dataStorageCodecImpl) DecodeDataStored(log *evm.Log) (*DataStored, error) {
	event := new(DataStored)
	if err := c.abi.UnpackIntoInterface(event, "DataStored", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["DataStored"].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *dataStorageCodecImpl) DynamicEventLogHash() []byte {
	return c.abi.Events["DynamicEvent"].ID.Bytes()
}

func (c *dataStorageCodecImpl) EncodeDynamicEventTopics(
	evt abi.Event,
	values []DynamicEvent,
) ([]*evm.TopicValues, error) {
	var userDataRule []interface{}
	for _, v := range values {
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.UserData)
		if err != nil {
			return nil, err
		}
		userDataRule = append(userDataRule, fieldVal)
	}
	var metadataRule []interface{}
	for _, v := range values {
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[3], v.Metadata)
		if err != nil {
			return nil, err
		}
		metadataRule = append(metadataRule, fieldVal)
	}
	var metadataArrayRule []interface{}
	for _, v := range values {
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[4], v.MetadataArray)
		if err != nil {
			return nil, err
		}
		metadataArrayRule = append(metadataArrayRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		userDataRule,
		metadataRule,
		metadataArrayRule,
	)
	if err != nil {
		return nil, err
	}

	topics := make([]*evm.TopicValues, len(rawTopics)+1)
	topics[0] = &evm.TopicValues{
		Values: [][]byte{evt.ID.Bytes()},
	}
	for i, hashList := range rawTopics {
		bs := make([][]byte, len(hashList))
		for j, h := range hashList {
			bs[j] = h.Bytes()
		}
		topics[i+1] = &evm.TopicValues{Values: bs}
	}
	return topics, nil
}

// DecodeDynamicEvent decodes a log into a DynamicEvent struct.
func (c *dataStorageCodecImpl) DecodeDynamicEvent(log *evm.Log) (*DynamicEvent, error) {
	event := new(DynamicEvent)
	if err := c.abi.UnpackIntoInterface(event, "DynamicEvent", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["DynamicEvent"].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *dataStorageCodecImpl) NoFieldsLogHash() []byte {
	return c.abi.Events["NoFields"].ID.Bytes()
}

func (c *dataStorageCodecImpl) EncodeNoFieldsTopics(
	evt abi.Event,
	values []NoFields,
) ([]*evm.TopicValues, error) {

	rawTopics, err := abi.MakeTopics()
	if err != nil {
		return nil, err
	}

	topics := make([]*evm.TopicValues, len(rawTopics)+1)
	topics[0] = &evm.TopicValues{
		Values: [][]byte{evt.ID.Bytes()},
	}
	for i, hashList := range rawTopics {
		bs := make([][]byte, len(hashList))
		for j, h := range hashList {
			bs[j] = h.Bytes()
		}
		topics[i+1] = &evm.TopicValues{Values: bs}
	}
	return topics, nil
}

// DecodeNoFields decodes a log into a NoFields struct.
func (c *dataStorageCodecImpl) DecodeNoFields(log *evm.Log) (*NoFields, error) {
	event := new(NoFields)
	if err := c.abi.UnpackIntoInterface(event, "NoFields", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["NoFields"].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c DataStorage) GetReserves(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeGetReservesMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(-3)), // -3 means latest finalized block
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c DataStorage) GetValue(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeGetValueMethodCall()
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(-3)), // -3 means latest finalized block
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c DataStorage) ReadData(
	runtime cre.Runtime,
	args ReadDataInput,
	blockNumber *big.Int,
) cre.Promise[*evm.CallContractReply] {
	calldata, err := c.Codec.EncodeReadDataMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[*evm.CallContractReply](nil, err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: pb.NewBigIntFromInt(big.NewInt(-3)), // -3 means latest finalized block
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	return cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address, Data: calldata},
			BlockNumber: bn,
		})
	})

}

func (c DataStorage) WriteReportDataStorageUpdateReserves(
	runtime cre.Runtime,
	input DataStorageUpdateReserves,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeDataStorageUpdateReservesStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *pb2.ReportResponse) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteReportRequest{
			Receiver:  c.Address,
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c DataStorage) WriteReportDataStorageUserData(
	runtime cre.Runtime,
	input DataStorageUserData,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeDataStorageUserDataStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *pb2.ReportResponse) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteReportRequest{
			Receiver:  c.Address,
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

// DecodeDataNotFoundError decodes a DataNotFound error from revert data.
func (c *DataStorage) DecodeDataNotFoundError(data []byte) (*DataNotFound, error) {
	args := c.ABI.Errors["DataNotFound"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 3 {
		return nil, fmt.Errorf("expected 3 values, got %d", len(values))
	}

	requester, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for requester in DataNotFound error")
	}

	key, ok1 := values[1].(string)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for key in DataNotFound error")
	}

	reason, ok2 := values[2].(string)
	if !ok2 {
		return nil, fmt.Errorf("unexpected type for reason in DataNotFound error")
	}

	return &DataNotFound{
		Requester: requester,
		Key:       key,
		Reason:    reason,
	}, nil
}

// Error implements the error interface for DataNotFound.
func (e *DataNotFound) Error() string {
	return fmt.Sprintf("DataNotFound error: requester=%v; key=%v; reason=%v;", e.Requester, e.Key, e.Reason)
}

// DecodeDataNotFound2Error decodes a DataNotFound2 error from revert data.
func (c *DataStorage) DecodeDataNotFound2Error(data []byte) (*DataNotFound2, error) {
	args := c.ABI.Errors["DataNotFound2"].Inputs
	values, err := args.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack error: %w", err)
	}
	if len(values) != 3 {
		return nil, fmt.Errorf("expected 3 values, got %d", len(values))
	}

	requester, ok0 := values[0].(common.Address)
	if !ok0 {
		return nil, fmt.Errorf("unexpected type for requester in DataNotFound2 error")
	}

	key, ok1 := values[1].(string)
	if !ok1 {
		return nil, fmt.Errorf("unexpected type for key in DataNotFound2 error")
	}

	reason, ok2 := values[2].(string)
	if !ok2 {
		return nil, fmt.Errorf("unexpected type for reason in DataNotFound2 error")
	}

	return &DataNotFound2{
		Requester: requester,
		Key:       key,
		Reason:    reason,
	}, nil
}

// Error implements the error interface for DataNotFound2.
func (e *DataNotFound2) Error() string {
	return fmt.Sprintf("DataNotFound2 error: requester=%v; key=%v; reason=%v;", e.Requester, e.Key, e.Reason)
}

func (c *DataStorage) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["DataNotFound"].ID.Bytes()[:4]):
		return c.DecodeDataNotFoundError(data)
	case common.Bytes2Hex(c.ABI.Errors["DataNotFound2"].ID.Bytes()[:4]):
		return c.DecodeDataNotFound2Error(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

func (c *DataStorage) LogTriggerAccessLoggedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []AccessLogged) (cre.Trigger[*evm.Log, *evm.Log], error) {
	event := c.ABI.Events["AccessLogged"]
	topics, err := c.Codec.EncodeAccessLoggedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for AccessLogged: %w", err)
	}

	return evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address},
		Topics:     topics,
		Confidence: confidence,
	}), nil
}

func (c *DataStorage) RegisterLogTrackingAccessLogged(runtime cre.Runtime, options *bindings.LogTrackingOptions[AccessLogged]) cre.Promise[*emptypb.Empty] {
	bindings.ValidateLogTrackingOptions[AccessLogged](options)
	topics, err := c.Codec.EncodeAccessLoggedTopics(c.ABI.Events["AccessLogged"], options.Filters)
	if err != nil {
		return cre.PromiseFromResult[*emptypb.Empty](nil, fmt.Errorf("failed to encode topics for AccessLogged: %w", err))
	}
	padded := bindings.PadTopics(topics)
	return c.client.RegisterLogTracking(runtime, &evm.RegisterLogTrackingRequest{
		Filter: &evm.LPFilter{
			Name:          "AccessLogged-" + common.Bytes2Hex(c.Address),
			Addresses:     [][]byte{c.Address},
			EventSigs:     [][]byte{c.Codec.AccessLoggedLogHash()},
			MaxLogsKept:   options.MaxLogsKept,
			RetentionTime: options.RetentionTime,
			LogsPerBlock:  options.LogsPerBlock,
			Topic2:        padded[1].Values,
			Topic3:        padded[2].Values,
			Topic4:        padded[3].Values,
		},
	})
}

func (c *DataStorage) UnregisterLogTrackingAccessLogged(runtime cre.Runtime) cre.Promise[*emptypb.Empty] {
	return c.client.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "AccessLogged-" + common.Bytes2Hex(c.Address),
	})
}

func (c *DataStorage) FilterLogsAccessLogged(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.AccessLoggedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	})
}

func (c *DataStorage) LogTriggerDataStoredLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DataStored) (cre.Trigger[*evm.Log, *evm.Log], error) {
	event := c.ABI.Events["DataStored"]
	topics, err := c.Codec.EncodeDataStoredTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for DataStored: %w", err)
	}

	return evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address},
		Topics:     topics,
		Confidence: confidence,
	}), nil
}

func (c *DataStorage) RegisterLogTrackingDataStored(runtime cre.Runtime, options *bindings.LogTrackingOptions[DataStored]) cre.Promise[*emptypb.Empty] {
	bindings.ValidateLogTrackingOptions[DataStored](options)
	topics, err := c.Codec.EncodeDataStoredTopics(c.ABI.Events["DataStored"], options.Filters)
	if err != nil {
		return cre.PromiseFromResult[*emptypb.Empty](nil, fmt.Errorf("failed to encode topics for DataStored: %w", err))
	}
	padded := bindings.PadTopics(topics)
	return c.client.RegisterLogTracking(runtime, &evm.RegisterLogTrackingRequest{
		Filter: &evm.LPFilter{
			Name:          "DataStored-" + common.Bytes2Hex(c.Address),
			Addresses:     [][]byte{c.Address},
			EventSigs:     [][]byte{c.Codec.DataStoredLogHash()},
			MaxLogsKept:   options.MaxLogsKept,
			RetentionTime: options.RetentionTime,
			LogsPerBlock:  options.LogsPerBlock,
			Topic2:        padded[1].Values,
			Topic3:        padded[2].Values,
			Topic4:        padded[3].Values,
		},
	})
}

func (c *DataStorage) UnregisterLogTrackingDataStored(runtime cre.Runtime) cre.Promise[*emptypb.Empty] {
	return c.client.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "DataStored-" + common.Bytes2Hex(c.Address),
	})
}

func (c *DataStorage) FilterLogsDataStored(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DataStoredLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	})
}

func (c *DataStorage) LogTriggerDynamicEventLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DynamicEvent) (cre.Trigger[*evm.Log, *evm.Log], error) {
	event := c.ABI.Events["DynamicEvent"]
	topics, err := c.Codec.EncodeDynamicEventTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for DynamicEvent: %w", err)
	}

	return evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address},
		Topics:     topics,
		Confidence: confidence,
	}), nil
}

func (c *DataStorage) RegisterLogTrackingDynamicEvent(runtime cre.Runtime, options *bindings.LogTrackingOptions[DynamicEvent]) cre.Promise[*emptypb.Empty] {
	bindings.ValidateLogTrackingOptions[DynamicEvent](options)
	topics, err := c.Codec.EncodeDynamicEventTopics(c.ABI.Events["DynamicEvent"], options.Filters)
	if err != nil {
		return cre.PromiseFromResult[*emptypb.Empty](nil, fmt.Errorf("failed to encode topics for DynamicEvent: %w", err))
	}
	padded := bindings.PadTopics(topics)
	return c.client.RegisterLogTracking(runtime, &evm.RegisterLogTrackingRequest{
		Filter: &evm.LPFilter{
			Name:          "DynamicEvent-" + common.Bytes2Hex(c.Address),
			Addresses:     [][]byte{c.Address},
			EventSigs:     [][]byte{c.Codec.DynamicEventLogHash()},
			MaxLogsKept:   options.MaxLogsKept,
			RetentionTime: options.RetentionTime,
			LogsPerBlock:  options.LogsPerBlock,
			Topic2:        padded[1].Values,
			Topic3:        padded[2].Values,
			Topic4:        padded[3].Values,
		},
	})
}

func (c *DataStorage) UnregisterLogTrackingDynamicEvent(runtime cre.Runtime) cre.Promise[*emptypb.Empty] {
	return c.client.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "DynamicEvent-" + common.Bytes2Hex(c.Address),
	})
}

func (c *DataStorage) FilterLogsDynamicEvent(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DynamicEventLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	})
}

func (c *DataStorage) LogTriggerNoFieldsLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []NoFields) (cre.Trigger[*evm.Log, *evm.Log], error) {
	event := c.ABI.Events["NoFields"]
	topics, err := c.Codec.EncodeNoFieldsTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for NoFields: %w", err)
	}

	return evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address},
		Topics:     topics,
		Confidence: confidence,
	}), nil
}

func (c *DataStorage) RegisterLogTrackingNoFields(runtime cre.Runtime, options *bindings.LogTrackingOptions[NoFields]) cre.Promise[*emptypb.Empty] {
	bindings.ValidateLogTrackingOptions[NoFields](options)
	topics, err := c.Codec.EncodeNoFieldsTopics(c.ABI.Events["NoFields"], options.Filters)
	if err != nil {
		return cre.PromiseFromResult[*emptypb.Empty](nil, fmt.Errorf("failed to encode topics for NoFields: %w", err))
	}
	padded := bindings.PadTopics(topics)
	return c.client.RegisterLogTracking(runtime, &evm.RegisterLogTrackingRequest{
		Filter: &evm.LPFilter{
			Name:          "NoFields-" + common.Bytes2Hex(c.Address),
			Addresses:     [][]byte{c.Address},
			EventSigs:     [][]byte{c.Codec.NoFieldsLogHash()},
			MaxLogsKept:   options.MaxLogsKept,
			RetentionTime: options.RetentionTime,
			LogsPerBlock:  options.LogsPerBlock,
			Topic2:        padded[1].Values,
			Topic3:        padded[2].Values,
			Topic4:        padded[3].Values,
		},
	})
}

func (c *DataStorage) UnregisterLogTrackingNoFields(runtime cre.Runtime) cre.Promise[*emptypb.Empty] {
	return c.client.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "NoFields-" + common.Bytes2Hex(c.Address),
	})
}

func (c *DataStorage) FilterLogsNoFields(runtime cre.Runtime, options *bindings.FilterOptions) cre.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.NoFieldsLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	})
}
