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

	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
	"github.com/smartcontractkit/chainlink-evm/pkg/bindings"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/sdk"
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
)

var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"DataNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"AccessLogged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"logAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"onReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"readData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structDataStorage.UserData\",\"name\":\"userData\",\"type\":\"tuple\"}],\"name\":\"storeUserData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Structs
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

// Errors
type DataNotFound struct {
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

// Main Binding Type for Bindings
type Bindings struct {
	Address   []byte
	Options   *bindings.ContractInitOptions
	ABI       *abi.ABI
	evmClient bindings.EVMClient
	Codec     BindingsCodec
}

type BindingsCodec interface {
	EncodeLogAccessMethodCall(in LogAccessInput) ([]byte, error)
	EncodeOnReportMethodCall(in OnReportInput) ([]byte, error)
	EncodeReadDataMethodCall(in ReadDataInput) ([]byte, error)
	DecodeReadDataMethodOutput(data []byte) (string, error)
	EncodeStoreDataMethodCall(in StoreDataInput) ([]byte, error)
	EncodeStoreUserDataMethodCall(in StoreUserDataInput) ([]byte, error)
	EncodeDataStorageUserDataStruct(in DataStorageUserData) ([]byte, error)
	AccessLoggedLogHash() []byte
	DecodeAccessLogged(log *evm.Log) (*AccessLogged, error)
	DataStoredLogHash() []byte
	DecodeDataStored(log *evm.Log) (*DataStored, error)
}

func NewBindings(
	client bindings.EVMClient,
	address []byte,
	options *bindings.ContractInitOptions,
) (*Bindings, error) {
	parsed, err := abi.JSON(strings.NewReader(BindingsMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewBindingsCodec()
	if err != nil {
		return nil, err
	}
	return &Bindings{
		Address:   address,
		Options:   options,
		ABI:       &parsed,
		evmClient: client,
		Codec:     codec,
	}, nil
}

type bindingsCodecImpl struct {
	abi *abi.ABI
}

func NewBindingsCodec() (BindingsCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(BindingsMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &bindingsCodecImpl{abi: &parsed}, nil
}

func (c *bindingsCodecImpl) EncodeLogAccessMethodCall(in LogAccessInput) ([]byte, error) {
	return c.abi.Pack("logAccess", in.Message)
}

func (c *bindingsCodecImpl) EncodeOnReportMethodCall(in OnReportInput) ([]byte, error) {
	return c.abi.Pack("onReport", in.Metadata, in.Payload)
}

func (c *bindingsCodecImpl) EncodeReadDataMethodCall(in ReadDataInput) ([]byte, error) {
	return c.abi.Pack("readData", in.User, in.Key)
}

func (c *bindingsCodecImpl) DecodeReadDataMethodOutput(data []byte) (string, error) {
	vals, err := c.abi.Methods["readData"].Outputs.Unpack(data)
	if err != nil {
		return *new(string), err
	}
	return vals[0].(string), nil
}

func (c *bindingsCodecImpl) EncodeStoreDataMethodCall(in StoreDataInput) ([]byte, error) {
	return c.abi.Pack("storeData", in.Key, in.Value)
}

func (c *bindingsCodecImpl) EncodeStoreUserDataMethodCall(in StoreUserDataInput) ([]byte, error) {
	return c.abi.Pack("storeUserData", in.UserData)
}

func (c *bindingsCodecImpl) EncodeDataStorageUserDataStruct(in DataStorageUserData) ([]byte, error) {
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

func (c *bindingsCodecImpl) AccessLoggedLogHash() []byte {
	return c.abi.Events["AccessLogged"].ID.Bytes()
}

// DecodeAccessLogged decodes a log into a AccessLogged struct.
func (c *bindingsCodecImpl) DecodeAccessLogged(log *evm.Log) (*AccessLogged, error) {
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

func (c *bindingsCodecImpl) DataStoredLogHash() []byte {
	return c.abi.Events["DataStored"].ID.Bytes()
}

// DecodeDataStored decodes a log into a DataStored struct.
func (c *bindingsCodecImpl) DecodeDataStored(log *evm.Log) (*DataStored, error) {
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

func (c Bindings) ReadData(
	runtime sdk.Runtime,
	args ReadDataInput,
	options *bindings.ReadOptions,
) (sdk.Promise[*evm.CallContractReply], error) {
	calldata, err := c.Codec.EncodeReadDataMethodCall(args)
	if err != nil {
		return nil, err
	}
	if options == nil {
		options = &bindings.ReadOptions{BlockNumber: nil}
	}
	return c.evmClient.CallContract(runtime, &evm.CallContractRequest{
		Call:        &evm.CallMsg{To: c.Address, Data: calldata},
		BlockNumber: pb.NewBigIntFromInt(options.BlockNumber),
	}), nil
}

// DecodeDataNotFoundError decodes a DataNotFound error from revert data.
func (c *Bindings) DecodeDataNotFoundError(data []byte) (*DataNotFound, error) {
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

func (c *Bindings) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	case common.Bytes2Hex(c.ABI.Errors["DataNotFound"].ID.Bytes()[:4]):
		return c.DecodeDataNotFoundError(data)
	default:
		return nil, errors.New("unknown error selector")
	}
}

// Error implements the error interface for DataNotFound.
func (e *DataNotFound) Error() string {
	return fmt.Sprintf("DataNotFound error: requester=%v; key=%v; reason=%v;", e.Requester, e.Key, e.Reason)
}

func (c *Bindings) LogTriggerAccessLoggedLog(confidence evm.ConfidenceLevel, values []AccessLogged) (sdk.Trigger[*evm.Log, *evm.Log], error) {
	event := ds.ABI.Events["AccessLogged"]
	var topicValues []*evm.TopicValues
	for _, v := range values {
		encoded, err := bindings.EncodeTopics(event, v)
		if err != nil {
			return nil, fmt.Errorf("failed to encode AccessLogged topics: %w", err)
		}
		topicValues = append(topicValues, &evm.TopicValues{
			Values: encoded,
		})
	}
	return ds.evmClient.LogTrigger(&evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{ds.Address},
		Topics:     topicValues,
		Confidence: confidence,
	}), nil
}

func (c *Bindings) RegisterLogTrackingAccessLogged(runtime sdk.Runtime, options *bindings.LogTrackingOptions) {
	bindings.ValidateLogTrackingOptions(options)
	c.evmClient.RegisterLogTracking(runtime, &evm.RegisterLogTrackingRequest{
		Filter: &evm.LPFilter{
			Name:          "AccessLogged-" + common.Bytes2Hex(c.Address),
			Addresses:     [][]byte{c.Address},
			EventSigs:     [][]byte{c.Codec.AccessLoggedLogHash()},
			MaxLogsKept:   options.MaxLogsKept,
			RetentionTime: options.RetentionTime,
			LogsPerBlock:  options.LogsPerBlock,
			Topic2:        options.Topic2,
			Topic3:        options.Topic3,
			Topic4:        options.Topic4,
		},
	})
}

func (c *Bindings) UnregisterLogTrackingAccessLogged(runtime sdk.Runtime) {
	c.evmClient.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "AccessLogged-" + common.Bytes2Hex(c.Address),
	})
}

func (c *Bindings) FilterLogsAccessLogged(runtime sdk.Runtime, options *bindings.FilterOptions) sdk.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
	}
	return c.evmClient.FilterLogs(runtime, &evm.FilterLogsRequest{
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

func (c *Bindings) LogTriggerDataStoredLog(confidence evm.ConfidenceLevel, values []DataStored) (sdk.Trigger[*evm.Log, *evm.Log], error) {
	event := ds.ABI.Events["DataStored"]
	var topicValues []*evm.TopicValues
	for _, v := range values {
		encoded, err := bindings.EncodeTopics(event, v)
		if err != nil {
			return nil, fmt.Errorf("failed to encode DataStored topics: %w", err)
		}
		topicValues = append(topicValues, &evm.TopicValues{
			Values: encoded,
		})
	}
	return ds.evmClient.LogTrigger(&evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{ds.Address},
		Topics:     topicValues,
		Confidence: confidence,
	}), nil
}

func (c *Bindings) RegisterLogTrackingDataStored(runtime sdk.Runtime, options *bindings.LogTrackingOptions) {
	bindings.ValidateLogTrackingOptions(options)
	c.evmClient.RegisterLogTracking(runtime, &evm.RegisterLogTrackingRequest{
		Filter: &evm.LPFilter{
			Name:          "DataStored-" + common.Bytes2Hex(c.Address),
			Addresses:     [][]byte{c.Address},
			EventSigs:     [][]byte{c.Codec.DataStoredLogHash()},
			MaxLogsKept:   options.MaxLogsKept,
			RetentionTime: options.RetentionTime,
			LogsPerBlock:  options.LogsPerBlock,
			Topic2:        options.Topic2,
			Topic3:        options.Topic3,
			Topic4:        options.Topic4,
		},
	})
}

func (c *Bindings) UnregisterLogTrackingDataStored(runtime sdk.Runtime) {
	c.evmClient.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "DataStored-" + common.Bytes2Hex(c.Address),
	})
}

func (c *Bindings) FilterLogsDataStored(runtime sdk.Runtime, options *bindings.FilterOptions) sdk.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: options.ToBlock,
		}
	}
	return c.evmClient.FilterLogs(runtime, &evm.FilterLogsRequest{
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

// TODO: implement
func getChainID(e bindings.EVMClient) uint32 { return 123 }
