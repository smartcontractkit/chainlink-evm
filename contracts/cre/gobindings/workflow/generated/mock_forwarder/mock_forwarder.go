// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock_forwarder

import (
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
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

type IRouterTransmissionInfo struct {
	TransmissionId  [32]byte
	State           uint8
	Transmitter     common.Address
	InvalidReceiver bool
	Success         bool
	GasLimit        *big.Int
}

var MockKeystoneForwarderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTransmissionId\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowExecutionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reportId\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getTransmissionInfo\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowExecutionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reportId\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIRouter.TransmissionInfo\",\"components\":[{\"name\":\"transmissionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumIRouter.TransmissionState\"},{\"name\":\"transmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"invalidReceiver\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"gasLimit\",\"type\":\"uint80\",\"internalType\":\"uint80\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransmitter\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"workflowExecutionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reportId\",\"type\":\"bytes2\",\"internalType\":\"bytes2\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeForwarder\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"report\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rawReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"route\",\"inputs\":[{\"name\":\"transmissionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transmitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadata\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"validatedReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ForwarderAdded\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForwarderRemoved\",\"inputs\":[{\"name\":\"forwarder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportProcessed\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"workflowExecutionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"reportId\",\"type\":\"bytes2\",\"indexed\":true,\"internalType\":\"bytes2\"},{\"name\":\"result\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyAttempted\",\"inputs\":[{\"name\":\"transmissionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasForRouting\",\"inputs\":[{\"name\":\"transmissionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidReport\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedForwarder\",\"inputs\":[]}]",
	Bin: "0x608080604052346088573315604657600080546001600160a01b031916331781553081526002602052604090819020805460ff1916600117905551611237908161008e8239f35b62461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f00000000000000006044820152606490fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c806311289565146100d7578063181f5a77146100d2578063233fd52d146100cd578063272cbd93146100c8578063354bdd66146100c35780634d93172d146100be5780635c41d2fe146100b957806379ba5097146100b45780638864b864146100af5780638da5cb5b146100aa578063abcef554146100a55763f2fde38b146100a057600080fd5b610c04565b610b9a565b610b48565b610b09565b6109ae565b6108fb565b61084b565b61082a565b6106d4565b610394565b6102c4565b610178565b6004359073ffffffffffffffffffffffffffffffffffffffff821682036100ff57565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff821682036100ff57565b6044359073ffffffffffffffffffffffffffffffffffffffff821682036100ff57565b9181601f840112156100ff5782359167ffffffffffffffff83116100ff57602083818601950101116100ff57565b346100ff5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ff576101af6100dc565b60243567ffffffffffffffff81116100ff576101cf90369060040161014a565b9060443567ffffffffffffffff81116100ff576101f090369060040161014a565b505060643567ffffffffffffffff81116100ff57366023820112156100ff5780600401359067ffffffffffffffff82116100ff57602490369260051b0101116100ff5761023c92610eac565b005b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176102ae57604052565b61023e565b604051906102c260c08361026d565b565b346100ff5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ff576040805190610302818361026d565b601f82527f4d6f636b4b657973746f6e65466f7277617264657220312e302e302d646576006020830152805180926020825280519081602084015260005b82811061037d5750506000828201840152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168101030190f35b602082820181015187830187015286945001610340565b346100ff5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ff576004356103ce610104565b6103d6610127565b9160643567ffffffffffffffff81116100ff576103f790369060040161014a565b91909260843567ffffffffffffffff81116100ff576105aa9561054f95600080946105386105989861050c6104836104348699369060040161014a565b9190958c88526003602052604088209073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055565b5a8b87526003602052604087209075ffffffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffff0000000000000000000000000000000000000000000083549260b01b16911617905560405194859360208501987f805f2132000000000000000000000000000000000000000000000000000000008a5260248601611063565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0810183528261026d565b51925af19283916000526003602052604060002090565b907fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff75ff000000000000000000000000000000000000000000835492151560a81b169116179055565b60405190151581529081906020820190565b0390f35b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60609101126100ff5760043573ffffffffffffffffffffffffffffffffffffffff811681036100ff5790602435906044357fffff000000000000000000000000000000000000000000000000000000000000811681036100ff5790565b919060c08301928151815260208201519160048310156106a55760a080916102c294602085015273ffffffffffffffffffffffffffffffffffffffff60408201511660408501526060810151151560608501526106926080820151608086019015159052565b015169ffffffffffffffffffff16910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b346100ff576105aa6106f76106e8366105ae565b916106f161108a565b50611138565b6107e9610716610711836000526003602052604060002090565b6110cc565b6107d6610737825173ffffffffffffffffffffffffffffffffffffffff1690565b9173ffffffffffffffffffffffffffffffffffffffff83166107f5576107cd6000915b6107c4610774606083015169ffffffffffffffffffff1690565b956107a7610792604061078a6020870151151590565b950151151590565b9561079b6102b3565b9a8b5260208b0161112c565b73ffffffffffffffffffffffffffffffffffffffff166040890152565b15156060870152565b15156080850152565b69ffffffffffffffffffff1660a0830152565b6040519182918261062c565b60208101511561080a576107cd60029161075a565b604081015115610820576107cd60015b9161075a565b6107cd600361081a565b346100ff57602061084361083d366105ae565b91611138565b604051908152f35b346100ff5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ff5773ffffffffffffffffffffffffffffffffffffffff6108976100dc565b61089f6111ab565b1680600052600260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690557fb96d15bf9258c7b8df062753a6a262864611fc7b060a5ee2e57e79b85f898d38600080a2005b346100ff5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ff5773ffffffffffffffffffffffffffffffffffffffff6109476100dc565b61094f6111ab565b16806000526002602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790557f0ea0ce2c048ff45a4a95f2947879de3fb94abec2f152190400cab2d1272a68e7600080a2005b346100ff5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ff5773ffffffffffffffffffffffffffffffffffffffff600154163303610aab5760005473ffffffffffffffffffffffffffffffffffffffff16600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055610a6c7fffffffffffffffffffffffff000000000000000000000000000000000000000060015416600155565b73ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152fd5b346100ff57610b1a61083d366105ae565b6000526003602052602073ffffffffffffffffffffffffffffffffffffffff60406000205416604051908152f35b346100ff5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ff57602073ffffffffffffffffffffffffffffffffffffffff60005416604051908152f35b346100ff5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ff5773ffffffffffffffffffffffffffffffffffffffff610be66100dc565b166000526002602052602060ff604060002054166040519015158152f35b346100ff5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126100ff5773ffffffffffffffffffffffffffffffffffffffff610c506100dc565b610c586111ab565b16338114610ce657807fffffffffffffffffffffffff0000000000000000000000000000000000000000600154161760015573ffffffffffffffffffffffffffffffffffffffff610cbe60005473ffffffffffffffffffffffffffffffffffffffff1690565b167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152fd5b92919267ffffffffffffffff82116102ae5760405191610d8c601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0166020018461026d565b8294818452818301116100ff578281602093846000960137010152565b90606d116100ff57602d0190604090565b9092919283606d116100ff5783116100ff57606d01917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff930190565b908160209103126100ff575180151581036100ff5790565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b959391610e8f9373ffffffffffffffffffffffffffffffffffffffff8092610e9d9a98948a5216602089015216604087015260a0606087015260a0860191610e0d565b926080818503910152610e0d565b90565b6040513d6000823e3d90fd5b9190606d821061103957826020610edd610ec7368686610d44565b602181015191608b604583015160c01c92015190565b9391959050610f3e610ef0858885611138565b91610f0781610eff818a610da9565b939099610dba565b9160405198899687967f233fd52d000000000000000000000000000000000000000000000000000000008852339060048901610e4c565b03816000305af191821561103457600092610fbf575b5060405191151582527fffff000000000000000000000000000000000000000000000000000000000000169273ffffffffffffffffffffffffffffffffffffffff16907f3617b009e9785c42daebadb6d3fb553243a4bf586d07ea72d65d80013ce116b590602090a4565b7f3617b009e9785c42daebadb6d3fb553243a4bf586d07ea72d65d80013ce116b59192506110257fffff0000000000000000000000000000000000000000000000000000000000009160203d60201161102d575b61101d818361026d565b810190610df5565b929150610f54565b503d611013565b610ea0565b7fb55ac7540000000000000000000000000000000000000000000000000000000060005260046000fd5b929061107c90610e9d9593604086526040860191610e0d565b926020818503910152610e0d565b6040519060c0820182811067ffffffffffffffff8211176102ae57604052600060a0838281528260208201528260408201528260608201528260808201520152565b906040516080810181811067ffffffffffffffff8211176102ae57604052606081935473ffffffffffffffffffffffffffffffffffffffff8116835260ff8160a01c161515602084015260ff8160a81c161515604084015260b01c910152565b60048210156106a55752565b917fffff00000000000000000000000000000000000000000000000000000000000090604051927fffffffffffffffffffffffffffffffffffffffff000000000000000000000000602085019560601b1685526034840152166054820152603681526111a560568261026d565b51902090565b73ffffffffffffffffffffffffffffffffffffffff6000541633036111cc57565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152fdfea164736f6c634300081e000a",
}

var MockKeystoneForwarderABI = MockKeystoneForwarderMetaData.ABI

var MockKeystoneForwarderBin = MockKeystoneForwarderMetaData.Bin

func DeployMockKeystoneForwarder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockKeystoneForwarder, error) {
	parsed, err := MockKeystoneForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockKeystoneForwarderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockKeystoneForwarder{address: address, abi: *parsed, MockKeystoneForwarderCaller: MockKeystoneForwarderCaller{contract: contract}, MockKeystoneForwarderTransactor: MockKeystoneForwarderTransactor{contract: contract}, MockKeystoneForwarderFilterer: MockKeystoneForwarderFilterer{contract: contract}}, nil
}

type MockKeystoneForwarder struct {
	address common.Address
	abi     abi.ABI
	MockKeystoneForwarderCaller
	MockKeystoneForwarderTransactor
	MockKeystoneForwarderFilterer
}

type MockKeystoneForwarderCaller struct {
	contract *bind.BoundContract
}

type MockKeystoneForwarderTransactor struct {
	contract *bind.BoundContract
}

type MockKeystoneForwarderFilterer struct {
	contract *bind.BoundContract
}

type MockKeystoneForwarderSession struct {
	Contract     *MockKeystoneForwarder
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MockKeystoneForwarderCallerSession struct {
	Contract *MockKeystoneForwarderCaller
	CallOpts bind.CallOpts
}

type MockKeystoneForwarderTransactorSession struct {
	Contract     *MockKeystoneForwarderTransactor
	TransactOpts bind.TransactOpts
}

type MockKeystoneForwarderRaw struct {
	Contract *MockKeystoneForwarder
}

type MockKeystoneForwarderCallerRaw struct {
	Contract *MockKeystoneForwarderCaller
}

type MockKeystoneForwarderTransactorRaw struct {
	Contract *MockKeystoneForwarderTransactor
}

func NewMockKeystoneForwarder(address common.Address, backend bind.ContractBackend) (*MockKeystoneForwarder, error) {
	abi, err := abi.JSON(strings.NewReader(MockKeystoneForwarderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMockKeystoneForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockKeystoneForwarder{address: address, abi: abi, MockKeystoneForwarderCaller: MockKeystoneForwarderCaller{contract: contract}, MockKeystoneForwarderTransactor: MockKeystoneForwarderTransactor{contract: contract}, MockKeystoneForwarderFilterer: MockKeystoneForwarderFilterer{contract: contract}}, nil
}

func NewMockKeystoneForwarderCaller(address common.Address, caller bind.ContractCaller) (*MockKeystoneForwarderCaller, error) {
	contract, err := bindMockKeystoneForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockKeystoneForwarderCaller{contract: contract}, nil
}

func NewMockKeystoneForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*MockKeystoneForwarderTransactor, error) {
	contract, err := bindMockKeystoneForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockKeystoneForwarderTransactor{contract: contract}, nil
}

func NewMockKeystoneForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*MockKeystoneForwarderFilterer, error) {
	contract, err := bindMockKeystoneForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockKeystoneForwarderFilterer{contract: contract}, nil
}

func bindMockKeystoneForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockKeystoneForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockKeystoneForwarder.Contract.MockKeystoneForwarderCaller.contract.Call(opts, result, method, params...)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.MockKeystoneForwarderTransactor.contract.Transfer(opts)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.MockKeystoneForwarderTransactor.contract.Transact(opts, method, params...)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockKeystoneForwarder.Contract.contract.Call(opts, result, method, params...)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.contract.Transfer(opts)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.contract.Transact(opts, method, params...)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCaller) GetTransmissionId(opts *bind.CallOpts, receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) ([32]byte, error) {
	var out []interface{}
	err := _MockKeystoneForwarder.contract.Call(opts, &out, "getTransmissionId", receiver, workflowExecutionId, reportId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) GetTransmissionId(receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) ([32]byte, error) {
	return _MockKeystoneForwarder.Contract.GetTransmissionId(&_MockKeystoneForwarder.CallOpts, receiver, workflowExecutionId, reportId)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCallerSession) GetTransmissionId(receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) ([32]byte, error) {
	return _MockKeystoneForwarder.Contract.GetTransmissionId(&_MockKeystoneForwarder.CallOpts, receiver, workflowExecutionId, reportId)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCaller) GetTransmissionInfo(opts *bind.CallOpts, receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) (IRouterTransmissionInfo, error) {
	var out []interface{}
	err := _MockKeystoneForwarder.contract.Call(opts, &out, "getTransmissionInfo", receiver, workflowExecutionId, reportId)

	if err != nil {
		return *new(IRouterTransmissionInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRouterTransmissionInfo)).(*IRouterTransmissionInfo)

	return out0, err

}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) GetTransmissionInfo(receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) (IRouterTransmissionInfo, error) {
	return _MockKeystoneForwarder.Contract.GetTransmissionInfo(&_MockKeystoneForwarder.CallOpts, receiver, workflowExecutionId, reportId)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCallerSession) GetTransmissionInfo(receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) (IRouterTransmissionInfo, error) {
	return _MockKeystoneForwarder.Contract.GetTransmissionInfo(&_MockKeystoneForwarder.CallOpts, receiver, workflowExecutionId, reportId)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCaller) GetTransmitter(opts *bind.CallOpts, receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) (common.Address, error) {
	var out []interface{}
	err := _MockKeystoneForwarder.contract.Call(opts, &out, "getTransmitter", receiver, workflowExecutionId, reportId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) GetTransmitter(receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) (common.Address, error) {
	return _MockKeystoneForwarder.Contract.GetTransmitter(&_MockKeystoneForwarder.CallOpts, receiver, workflowExecutionId, reportId)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCallerSession) GetTransmitter(receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) (common.Address, error) {
	return _MockKeystoneForwarder.Contract.GetTransmitter(&_MockKeystoneForwarder.CallOpts, receiver, workflowExecutionId, reportId)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCaller) IsForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _MockKeystoneForwarder.contract.Call(opts, &out, "isForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) IsForwarder(forwarder common.Address) (bool, error) {
	return _MockKeystoneForwarder.Contract.IsForwarder(&_MockKeystoneForwarder.CallOpts, forwarder)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCallerSession) IsForwarder(forwarder common.Address) (bool, error) {
	return _MockKeystoneForwarder.Contract.IsForwarder(&_MockKeystoneForwarder.CallOpts, forwarder)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockKeystoneForwarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) Owner() (common.Address, error) {
	return _MockKeystoneForwarder.Contract.Owner(&_MockKeystoneForwarder.CallOpts)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCallerSession) Owner() (common.Address, error) {
	return _MockKeystoneForwarder.Contract.Owner(&_MockKeystoneForwarder.CallOpts)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockKeystoneForwarder.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) TypeAndVersion() (string, error) {
	return _MockKeystoneForwarder.Contract.TypeAndVersion(&_MockKeystoneForwarder.CallOpts)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderCallerSession) TypeAndVersion() (string, error) {
	return _MockKeystoneForwarder.Contract.TypeAndVersion(&_MockKeystoneForwarder.CallOpts)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockKeystoneForwarder.contract.Transact(opts, "acceptOwnership")
}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.AcceptOwnership(&_MockKeystoneForwarder.TransactOpts)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.AcceptOwnership(&_MockKeystoneForwarder.TransactOpts)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactor) AddForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _MockKeystoneForwarder.contract.Transact(opts, "addForwarder", forwarder)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) AddForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.AddForwarder(&_MockKeystoneForwarder.TransactOpts, forwarder)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactorSession) AddForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.AddForwarder(&_MockKeystoneForwarder.TransactOpts, forwarder)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactor) RemoveForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error) {
	return _MockKeystoneForwarder.contract.Transact(opts, "removeForwarder", forwarder)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) RemoveForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.RemoveForwarder(&_MockKeystoneForwarder.TransactOpts, forwarder)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactorSession) RemoveForwarder(forwarder common.Address) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.RemoveForwarder(&_MockKeystoneForwarder.TransactOpts, forwarder)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactor) Report(opts *bind.TransactOpts, receiver common.Address, rawReport []byte, arg2 []byte, arg3 [][]byte) (*types.Transaction, error) {
	return _MockKeystoneForwarder.contract.Transact(opts, "report", receiver, rawReport, arg2, arg3)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) Report(receiver common.Address, rawReport []byte, arg2 []byte, arg3 [][]byte) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.Report(&_MockKeystoneForwarder.TransactOpts, receiver, rawReport, arg2, arg3)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactorSession) Report(receiver common.Address, rawReport []byte, arg2 []byte, arg3 [][]byte) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.Report(&_MockKeystoneForwarder.TransactOpts, receiver, rawReport, arg2, arg3)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactor) Route(opts *bind.TransactOpts, transmissionId [32]byte, transmitter common.Address, receiver common.Address, metadata []byte, validatedReport []byte) (*types.Transaction, error) {
	return _MockKeystoneForwarder.contract.Transact(opts, "route", transmissionId, transmitter, receiver, metadata, validatedReport)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) Route(transmissionId [32]byte, transmitter common.Address, receiver common.Address, metadata []byte, validatedReport []byte) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.Route(&_MockKeystoneForwarder.TransactOpts, transmissionId, transmitter, receiver, metadata, validatedReport)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactorSession) Route(transmissionId [32]byte, transmitter common.Address, receiver common.Address, metadata []byte, validatedReport []byte) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.Route(&_MockKeystoneForwarder.TransactOpts, transmissionId, transmitter, receiver, metadata, validatedReport)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _MockKeystoneForwarder.contract.Transact(opts, "transferOwnership", to)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.TransferOwnership(&_MockKeystoneForwarder.TransactOpts, to)
}

func (_MockKeystoneForwarder *MockKeystoneForwarderTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _MockKeystoneForwarder.Contract.TransferOwnership(&_MockKeystoneForwarder.TransactOpts, to)
}

type MockKeystoneForwarderForwarderAddedIterator struct {
	Event *MockKeystoneForwarderForwarderAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockKeystoneForwarderForwarderAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockKeystoneForwarderForwarderAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockKeystoneForwarderForwarderAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockKeystoneForwarderForwarderAddedIterator) Error() error {
	return it.fail
}

func (it *MockKeystoneForwarderForwarderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockKeystoneForwarderForwarderAdded struct {
	Forwarder common.Address
	Raw       types.Log
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) FilterForwarderAdded(opts *bind.FilterOpts, forwarder []common.Address) (*MockKeystoneForwarderForwarderAddedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.FilterLogs(opts, "ForwarderAdded", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &MockKeystoneForwarderForwarderAddedIterator{contract: _MockKeystoneForwarder.contract, event: "ForwarderAdded", logs: logs, sub: sub}, nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) WatchForwarderAdded(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderForwarderAdded, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.WatchLogs(opts, "ForwarderAdded", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockKeystoneForwarderForwarderAdded)
				if err := _MockKeystoneForwarder.contract.UnpackLog(event, "ForwarderAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) ParseForwarderAdded(log types.Log) (*MockKeystoneForwarderForwarderAdded, error) {
	event := new(MockKeystoneForwarderForwarderAdded)
	if err := _MockKeystoneForwarder.contract.UnpackLog(event, "ForwarderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockKeystoneForwarderForwarderRemovedIterator struct {
	Event *MockKeystoneForwarderForwarderRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockKeystoneForwarderForwarderRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockKeystoneForwarderForwarderRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockKeystoneForwarderForwarderRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockKeystoneForwarderForwarderRemovedIterator) Error() error {
	return it.fail
}

func (it *MockKeystoneForwarderForwarderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockKeystoneForwarderForwarderRemoved struct {
	Forwarder common.Address
	Raw       types.Log
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) FilterForwarderRemoved(opts *bind.FilterOpts, forwarder []common.Address) (*MockKeystoneForwarderForwarderRemovedIterator, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.FilterLogs(opts, "ForwarderRemoved", forwarderRule)
	if err != nil {
		return nil, err
	}
	return &MockKeystoneForwarderForwarderRemovedIterator{contract: _MockKeystoneForwarder.contract, event: "ForwarderRemoved", logs: logs, sub: sub}, nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) WatchForwarderRemoved(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderForwarderRemoved, forwarder []common.Address) (event.Subscription, error) {

	var forwarderRule []interface{}
	for _, forwarderItem := range forwarder {
		forwarderRule = append(forwarderRule, forwarderItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.WatchLogs(opts, "ForwarderRemoved", forwarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockKeystoneForwarderForwarderRemoved)
				if err := _MockKeystoneForwarder.contract.UnpackLog(event, "ForwarderRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) ParseForwarderRemoved(log types.Log) (*MockKeystoneForwarderForwarderRemoved, error) {
	event := new(MockKeystoneForwarderForwarderRemoved)
	if err := _MockKeystoneForwarder.contract.UnpackLog(event, "ForwarderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockKeystoneForwarderOwnershipTransferRequestedIterator struct {
	Event *MockKeystoneForwarderOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockKeystoneForwarderOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockKeystoneForwarderOwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockKeystoneForwarderOwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockKeystoneForwarderOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *MockKeystoneForwarderOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockKeystoneForwarderOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockKeystoneForwarderOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockKeystoneForwarderOwnershipTransferRequestedIterator{contract: _MockKeystoneForwarder.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockKeystoneForwarderOwnershipTransferRequested)
				if err := _MockKeystoneForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) ParseOwnershipTransferRequested(log types.Log) (*MockKeystoneForwarderOwnershipTransferRequested, error) {
	event := new(MockKeystoneForwarderOwnershipTransferRequested)
	if err := _MockKeystoneForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockKeystoneForwarderOwnershipTransferredIterator struct {
	Event *MockKeystoneForwarderOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockKeystoneForwarderOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockKeystoneForwarderOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockKeystoneForwarderOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockKeystoneForwarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *MockKeystoneForwarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockKeystoneForwarderOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockKeystoneForwarderOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockKeystoneForwarderOwnershipTransferredIterator{contract: _MockKeystoneForwarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockKeystoneForwarderOwnershipTransferred)
				if err := _MockKeystoneForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) ParseOwnershipTransferred(log types.Log) (*MockKeystoneForwarderOwnershipTransferred, error) {
	event := new(MockKeystoneForwarderOwnershipTransferred)
	if err := _MockKeystoneForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type MockKeystoneForwarderReportProcessedIterator struct {
	Event *MockKeystoneForwarderReportProcessed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MockKeystoneForwarderReportProcessedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockKeystoneForwarderReportProcessed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(MockKeystoneForwarderReportProcessed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *MockKeystoneForwarderReportProcessedIterator) Error() error {
	return it.fail
}

func (it *MockKeystoneForwarderReportProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MockKeystoneForwarderReportProcessed struct {
	Receiver            common.Address
	WorkflowExecutionId [32]byte
	ReportId            [2]byte
	Result              bool
	Raw                 types.Log
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) FilterReportProcessed(opts *bind.FilterOpts, receiver []common.Address, workflowExecutionId [][32]byte, reportId [][2]byte) (*MockKeystoneForwarderReportProcessedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var workflowExecutionIdRule []interface{}
	for _, workflowExecutionIdItem := range workflowExecutionId {
		workflowExecutionIdRule = append(workflowExecutionIdRule, workflowExecutionIdItem)
	}
	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.FilterLogs(opts, "ReportProcessed", receiverRule, workflowExecutionIdRule, reportIdRule)
	if err != nil {
		return nil, err
	}
	return &MockKeystoneForwarderReportProcessedIterator{contract: _MockKeystoneForwarder.contract, event: "ReportProcessed", logs: logs, sub: sub}, nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) WatchReportProcessed(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderReportProcessed, receiver []common.Address, workflowExecutionId [][32]byte, reportId [][2]byte) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var workflowExecutionIdRule []interface{}
	for _, workflowExecutionIdItem := range workflowExecutionId {
		workflowExecutionIdRule = append(workflowExecutionIdRule, workflowExecutionIdItem)
	}
	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _MockKeystoneForwarder.contract.WatchLogs(opts, "ReportProcessed", receiverRule, workflowExecutionIdRule, reportIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MockKeystoneForwarderReportProcessed)
				if err := _MockKeystoneForwarder.contract.UnpackLog(event, "ReportProcessed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarderFilterer) ParseReportProcessed(log types.Log) (*MockKeystoneForwarderReportProcessed, error) {
	event := new(MockKeystoneForwarderReportProcessed)
	if err := _MockKeystoneForwarder.contract.UnpackLog(event, "ReportProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MockKeystoneForwarder *MockKeystoneForwarder) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MockKeystoneForwarder.abi.Events["ForwarderAdded"].ID:
		return _MockKeystoneForwarder.ParseForwarderAdded(log)
	case _MockKeystoneForwarder.abi.Events["ForwarderRemoved"].ID:
		return _MockKeystoneForwarder.ParseForwarderRemoved(log)
	case _MockKeystoneForwarder.abi.Events["OwnershipTransferRequested"].ID:
		return _MockKeystoneForwarder.ParseOwnershipTransferRequested(log)
	case _MockKeystoneForwarder.abi.Events["OwnershipTransferred"].ID:
		return _MockKeystoneForwarder.ParseOwnershipTransferred(log)
	case _MockKeystoneForwarder.abi.Events["ReportProcessed"].ID:
		return _MockKeystoneForwarder.ParseReportProcessed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MockKeystoneForwarderForwarderAdded) Topic() common.Hash {
	return common.HexToHash("0x0ea0ce2c048ff45a4a95f2947879de3fb94abec2f152190400cab2d1272a68e7")
}

func (MockKeystoneForwarderForwarderRemoved) Topic() common.Hash {
	return common.HexToHash("0xb96d15bf9258c7b8df062753a6a262864611fc7b060a5ee2e57e79b85f898d38")
}

func (MockKeystoneForwarderOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (MockKeystoneForwarderOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (MockKeystoneForwarderReportProcessed) Topic() common.Hash {
	return common.HexToHash("0x3617b009e9785c42daebadb6d3fb553243a4bf586d07ea72d65d80013ce116b5")
}

func (_MockKeystoneForwarder *MockKeystoneForwarder) Address() common.Address {
	return _MockKeystoneForwarder.address
}

type MockKeystoneForwarderInterface interface {
	GetTransmissionId(opts *bind.CallOpts, receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) ([32]byte, error)

	GetTransmissionInfo(opts *bind.CallOpts, receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) (IRouterTransmissionInfo, error)

	GetTransmitter(opts *bind.CallOpts, receiver common.Address, workflowExecutionId [32]byte, reportId [2]byte) (common.Address, error)

	IsForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error)

	RemoveForwarder(opts *bind.TransactOpts, forwarder common.Address) (*types.Transaction, error)

	Report(opts *bind.TransactOpts, receiver common.Address, rawReport []byte, arg2 []byte, arg3 [][]byte) (*types.Transaction, error)

	Route(opts *bind.TransactOpts, transmissionId [32]byte, transmitter common.Address, receiver common.Address, metadata []byte, validatedReport []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterForwarderAdded(opts *bind.FilterOpts, forwarder []common.Address) (*MockKeystoneForwarderForwarderAddedIterator, error)

	WatchForwarderAdded(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderForwarderAdded, forwarder []common.Address) (event.Subscription, error)

	ParseForwarderAdded(log types.Log) (*MockKeystoneForwarderForwarderAdded, error)

	FilterForwarderRemoved(opts *bind.FilterOpts, forwarder []common.Address) (*MockKeystoneForwarderForwarderRemovedIterator, error)

	WatchForwarderRemoved(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderForwarderRemoved, forwarder []common.Address) (event.Subscription, error)

	ParseForwarderRemoved(log types.Log) (*MockKeystoneForwarderForwarderRemoved, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockKeystoneForwarderOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*MockKeystoneForwarderOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockKeystoneForwarderOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*MockKeystoneForwarderOwnershipTransferred, error)

	FilterReportProcessed(opts *bind.FilterOpts, receiver []common.Address, workflowExecutionId [][32]byte, reportId [][2]byte) (*MockKeystoneForwarderReportProcessedIterator, error)

	WatchReportProcessed(opts *bind.WatchOpts, sink chan<- *MockKeystoneForwarderReportProcessed, receiver []common.Address, workflowExecutionId [][32]byte, reportId [][2]byte) (event.Subscription, error)

	ParseReportProcessed(log types.Log) (*MockKeystoneForwarderReportProcessed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
