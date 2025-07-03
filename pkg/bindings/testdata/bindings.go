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

	evmcappb "github.com/smartcontractkit/chainlink-common/pkg/capabilities/v2/chain-capabilities/evm"
	"github.com/smartcontractkit/chainlink-common/pkg/chains/evm"
	"github.com/smartcontractkit/chainlink-common/pkg/values/pb"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows/sdk/v2"
	"github.com/smartcontractkit/chainlink-evm/pkg/bindings"
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

var DataStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"DataNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"AccessLogged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"logAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"onReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"readData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structDataStorage.UserData\",\"name\":\"userData\",\"type\":\"tuple\"}],\"name\":\"storeUserData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newValue\",\"type\":\"string\"}],\"name\":\"updateData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"oldValue\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061139b8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80634ece5b4c14610064578063805f21321461008057806398458c5d1461009c578063bddbb023146100b8578063ccf15827146100e8578063f5bfa81514610104575b5f5ffd5b61007e6004803603810190610079919061076c565b610134565b005b61009a6004803603810190610095919061083f565b6101f8565b005b6100b660048036038101906100b191906108df565b6102d5565b005b6100d260048036038101906100cd919061076c565b6103cc565b6040516100df9190610996565b60405180910390f35b61010260048036038101906100fd91906109b6565b61056c565b005b61011e60048036038101906101199190610a5b565b6105c0565b60405161012b9190610996565b60405180910390f35b81815f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208686604051610182929190610af4565b9081526020016040518091039020918261019d929190610d49565b503373ffffffffffffffffffffffffffffffffffffffff167fc95c7d5d3ac582f659cd004afbea77723e1315567b6557f3c059e8eb9586518f858585856040516101ea9493929190610e42565b60405180910390a250505050565b5f82828101906102089190610ff5565b905080602001515f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20825f015160405161025c919061106c565b908152602001604051809103902090816102769190611082565b503373ffffffffffffffffffffffffffffffffffffffff167fc95c7d5d3ac582f659cd004afbea77723e1315567b6557f3c059e8eb9586518f825f015183602001516040516102c6929190611151565b60405180910390a25050505050565b8080602001906102e59190611192565b5f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2083805f01906103309190611192565b60405161033e929190610af4565b90815260200160405180910390209182610359929190610d49565b503373ffffffffffffffffffffffffffffffffffffffff167fc95c7d5d3ac582f659cd004afbea77723e1315567b6557f3c059e8eb9586518f82805f01906103a19190611192565b8480602001906103b19190611192565b6040516103c19493929190610e42565b60405180910390a250565b60605f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20858560405161041a929190610af4565b9081526020016040518091039020805461043390610b70565b80601f016020809104026020016040519081016040528092919081815260200182805461045f90610b70565b80156104aa5780601f10610481576101008083540402835291602001916104aa565b820191905f5260205f20905b81548152906001019060200180831161048d57829003601f168201915b505050505090505f8151036104fa573385856040517ff1e502090000000000000000000000000000000000000000000000000000000081526004016104f19392919061124d565b60405180910390fd5b82825f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208787604051610548929190610af4565b90815260200160405180910390209182610563929190610d49565b50949350505050565b3373ffffffffffffffffffffffffffffffffffffffff167fe2ab1536af9681ad9e5927bca61830526c4cd932e970162eef77328af1fdcfb583836040516105b4929190611290565b60405180910390a25050565b60605f5f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20848460405161060f929190610af4565b9081526020016040518091039020805461062890610b70565b80601f016020809104026020016040519081016040528092919081815260200182805461065490610b70565b801561069f5780601f106106765761010080835404028352916020019161069f565b820191905f5260205f20905b81548152906001019060200180831161068257829003601f168201915b505050505090505f8151036106ef578484846040517ff1e502090000000000000000000000000000000000000000000000000000000081526004016106e693929190611322565b60405180910390fd5b809150509392505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261072c5761072b61070b565b5b8235905067ffffffffffffffff8111156107495761074861070f565b5b60208301915083600182028301111561076557610764610713565b5b9250929050565b5f5f5f5f6040858703121561078457610783610703565b5b5f85013567ffffffffffffffff8111156107a1576107a0610707565b5b6107ad87828801610717565b9450945050602085013567ffffffffffffffff8111156107d0576107cf610707565b5b6107dc87828801610717565b925092505092959194509250565b5f5f83601f8401126107ff576107fe61070b565b5b8235905067ffffffffffffffff81111561081c5761081b61070f565b5b60208301915083600182028301111561083857610837610713565b5b9250929050565b5f5f5f5f6040858703121561085757610856610703565b5b5f85013567ffffffffffffffff81111561087457610873610707565b5b610880878288016107ea565b9450945050602085013567ffffffffffffffff8111156108a3576108a2610707565b5b6108af878288016107ea565b925092505092959194509250565b5f5ffd5b5f604082840312156108d6576108d56108bd565b5b81905092915050565b5f602082840312156108f4576108f3610703565b5b5f82013567ffffffffffffffff81111561091157610910610707565b5b61091d848285016108c1565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61096882610926565b6109728185610930565b9350610982818560208601610940565b61098b8161094e565b840191505092915050565b5f6020820190508181035f8301526109ae818461095e565b905092915050565b5f5f602083850312156109cc576109cb610703565b5b5f83013567ffffffffffffffff8111156109e9576109e8610707565b5b6109f585828601610717565b92509250509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a2a82610a01565b9050919050565b610a3a81610a20565b8114610a44575f5ffd5b50565b5f81359050610a5581610a31565b92915050565b5f5f5f60408486031215610a7257610a71610703565b5b5f610a7f86828701610a47565b935050602084013567ffffffffffffffff811115610aa057610a9f610707565b5b610aac86828701610717565b92509250509250925092565b5f81905092915050565b828183375f83830152505050565b5f610adb8385610ab8565b9350610ae8838584610ac2565b82840190509392505050565b5f610b00828486610ad0565b91508190509392505050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610b8757607f821691505b602082108103610b9a57610b99610b43565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610bfc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610bc1565b610c068683610bc1565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610c4a610c45610c4084610c1e565b610c27565b610c1e565b9050919050565b5f819050919050565b610c6383610c30565b610c77610c6f82610c51565b848454610bcd565b825550505050565b5f5f905090565b610c8e610c7f565b610c99818484610c5a565b505050565b5b81811015610cbc57610cb15f82610c86565b600181019050610c9f565b5050565b601f821115610d0157610cd281610ba0565b610cdb84610bb2565b81016020851015610cea578190505b610cfe610cf685610bb2565b830182610c9e565b50505b505050565b5f82821c905092915050565b5f610d215f1984600802610d06565b1980831691505092915050565b5f610d398383610d12565b9150826002028217905092915050565b610d538383610b0c565b67ffffffffffffffff811115610d6c57610d6b610b16565b5b610d768254610b70565b610d81828285610cc0565b5f601f831160018114610dae575f8415610d9c578287013590505b610da68582610d2e565b865550610e0d565b601f198416610dbc86610ba0565b5f5b82811015610de357848901358255600182019150602085019450602081019050610dbe565b86831015610e005784890135610dfc601f891682610d12565b8355505b6001600288020188555050505b50505050505050565b5f610e218385610930565b9350610e2e838584610ac2565b610e378361094e565b840190509392505050565b5f6040820190508181035f830152610e5b818688610e16565b90508181036020830152610e70818486610e16565b905095945050505050565b5f5ffd5b610e888261094e565b810181811067ffffffffffffffff82111715610ea757610ea6610b16565b5b80604052505050565b5f610eb96106fa565b9050610ec58282610e7f565b919050565b5f5ffd5b5f5ffd5b5f67ffffffffffffffff821115610eec57610eeb610b16565b5b610ef58261094e565b9050602081019050919050565b5f610f14610f0f84610ed2565b610eb0565b905082815260208101848484011115610f3057610f2f610ece565b5b610f3b848285610ac2565b509392505050565b5f82601f830112610f5757610f5661070b565b5b8135610f67848260208601610f02565b91505092915050565b5f60408284031215610f8557610f84610e7b565b5b610f8f6040610eb0565b90505f82013567ffffffffffffffff811115610fae57610fad610eca565b5b610fba84828501610f43565b5f83015250602082013567ffffffffffffffff811115610fdd57610fdc610eca565b5b610fe984828501610f43565b60208301525092915050565b5f6020828403121561100a57611009610703565b5b5f82013567ffffffffffffffff81111561102757611026610707565b5b61103384828501610f70565b91505092915050565b5f61104682610926565b6110508185610ab8565b9350611060818560208601610940565b80840191505092915050565b5f611077828461103c565b915081905092915050565b61108b82610926565b67ffffffffffffffff8111156110a4576110a3610b16565b5b6110ae8254610b70565b6110b9828285610cc0565b5f60209050601f8311600181146110ea575f84156110d8578287015190505b6110e28582610d2e565b865550611149565b601f1984166110f886610ba0565b5f5b8281101561111f578489015182556001820191506020850194506020810190506110fa565b8683101561113c5784890151611138601f891682610d12565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f830152611169818561095e565b9050818103602083015261117d818461095e565b90509392505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f833560016020038436030381126111ae576111ad611186565b5b80840192508235915067ffffffffffffffff8211156111d0576111cf61118a565b5b6020830192506001820236038313156111ec576111eb61118e565b5b509250929050565b6111fd81610a20565b82525050565b7f4e6f206578697374696e67206461746120746f207570646174650000000000005f82015250565b5f611237601a83610930565b915061124282611203565b602082019050919050565b5f6060820190506112605f8301866111f4565b8181036020830152611273818486610e16565b905081810360408301526112868161122b565b9050949350505050565b5f6020820190508181035f8301526112a9818486610e16565b90509392505050565b7f4e6f2064617461206173736f63696174656420776974682074686973206b65795f8201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f61130c602183610930565b9150611317826112b2565b604082019050919050565b5f6060820190506113355f8301866111f4565b8181036020830152611348818486610e16565b9050818103604083015261135b81611300565b905094935050505056fea2646970667358221220edbc6493bb8a529c59e4d71f5c772673cb566e139e9f126bde94d7b143a7080c64736f6c634300081e0033",
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

// Main Binding Type for DataStorage
type DataStorage struct {
	Address   []byte
	Options   *bindings.ContractInitOptions
	ABI       *abi.ABI
	evmClient bindings.EVMClient
	Codec     DataStorageCodec
}

type DataStorageCodec interface {
	EncodeLogAccessMethodCall(in LogAccessInput) ([]byte, error)
	EncodeOnReportMethodCall(in OnReportInput) ([]byte, error)
	EncodeReadDataMethodCall(in ReadDataInput) ([]byte, error)
	DecodeReadDataMethodOutput(data []byte) (string, error)
	EncodeStoreDataMethodCall(in StoreDataInput) ([]byte, error)
	EncodeStoreUserDataMethodCall(in StoreUserDataInput) ([]byte, error)
	EncodeUpdateDataMethodCall(in UpdateDataInput) ([]byte, error)
	DecodeUpdateDataMethodOutput(data []byte) (string, error)
	EncodeDataStorageUserDataStruct(in DataStorageUserData) ([]byte, error)
	AccessLoggedLogHash() []byte
	DecodeAccessLogged(log *evm.Log) (*AccessLogged, error)
	DataStoredLogHash() []byte
	DecodeDataStored(log *evm.Log) (*DataStored, error)
}

func NewDataStorage(
	client bindings.EVMClient,
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
		Address:   address,
		Options:   options,
		ABI:       &parsed,
		evmClient: client,
		Codec:     codec,
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

func (c DataStorage) ReadData(
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

func (c DataStorage) WriteReportDataStorageUserData(
	runtime sdk.Runtime,
	input DataStorageUserData,
	gasConfig *evmcappb.GasConfig,
) (sdk.Promise[*evmcappb.WriteReportReply], error) {
	encoded, err := c.Codec.EncodeDataStorageUserDataStruct(input)
	if err != nil {
		return nil, err
	}
	report := bindings.GenerateReport(getChainID(c.evmClient), encoded)
	return c.evmClient.WriteReport(runtime, &evmcappb.WriteReportRequest{
		Receiver: c.Address,
		Report: &evmcappb.SignedReport{
			RawReport:     report.RawReport,
			ReportContext: report.ReportContext,
			Signatures:    report.Signatures,
			Id:            report.ID,
		},
		GasConfig: gasConfig,
	}), nil
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

func (c *DataStorage) UnpackError(data []byte) (any, error) {
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

func (c *DataStorage) RegisterLogTrackingAccessLogged(runtime sdk.Runtime, options *bindings.LogTrackingOptions) {
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

func (c *DataStorage) UnregisterLogTrackingAccessLogged(runtime sdk.Runtime) {
	c.evmClient.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "AccessLogged-" + common.Bytes2Hex(c.Address),
	})
}

func (c *DataStorage) FilterLogsAccessLogged(runtime sdk.Runtime, options *bindings.FilterOptions) sdk.Promise[*evm.FilterLogsReply] {
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

func (c *DataStorage) RegisterLogTrackingDataStored(runtime sdk.Runtime, options *bindings.LogTrackingOptions) {
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

func (c *DataStorage) UnregisterLogTrackingDataStored(runtime sdk.Runtime) {
	c.evmClient.UnregisterLogTracking(runtime, &evm.UnregisterLogTrackingRequest{
		FilterName: "DataStored-" + common.Bytes2Hex(c.Address),
	})
}

func (c *DataStorage) FilterLogsDataStored(runtime sdk.Runtime, options *bindings.FilterOptions) sdk.Promise[*evm.FilterLogsReply] {
	if options == nil {
		options = &bindings.FilterOptions{
			ToBlock: big.NewInt(options.ToBlock),
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
