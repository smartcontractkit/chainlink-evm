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
	Bin: "0x608060405234801561000f575f80fd5b5033805f816100655760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b038481169190911790915581161561009457610094816100b5565b5050305f908152600260205260409020805460ff191660011790555061015d565b336001600160a01b0382160361010d5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005c565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6111b18061016a5f395ff3fe608060405234801561000f575f80fd5b50600436106100cf575f3560e01c80635c41d2fe1161007d5780638da5cb5b116100585780638da5cb5b14610321578063abcef5541461033e578063f2fde38b14610376575f80fd5b80635c41d2fe1461022e57806379ba5097146102415780638864b86414610249575f80fd5b8063272cbd93116100ad578063272cbd931461015d578063354bdd661461017d5780634d93172d1461021b575f80fd5b806311289565146100d3578063181f5a77146100e8578063233fd52d1461013a575b5f80fd5b6100e66100e1366004610d4b565b610389565b005b6101246040518060400160405280601f81526020017f4d6f636b4b657973746f6e65466f7277617264657220312e302e302d6465760081525081565b6040516101319190610e2c565b60405180910390f35b61014d610148366004610e7f565b61059e565b6040519015158152602001610131565b61017061016b366004610f18565b61073c565b6040516101319190610fa6565b61020d61018b366004610f18565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606085901b166020820152603481018390527fffff000000000000000000000000000000000000000000000000000000000000821660548201525f906056016040516020818303038152906040528051906020012090509392505050565b604051908152602001610131565b6100e6610229366004611054565b61093f565b6100e661023c366004611054565b6109ba565b6100e6610a38565b6102fc610257366004610f18565b6040805160609490941b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208086019190915260348501939093527fffff00000000000000000000000000000000000000000000000000000000000091909116605484015280516036818503018152605690930181528251928201929092205f9081526003909152205473ffffffffffffffffffffffffffffffffffffffff1690565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610131565b5f5473ffffffffffffffffffffffffffffffffffffffff166102fc565b61014d61034c366004611054565b73ffffffffffffffffffffffffffffffffffffffff165f9081526002602052604090205460ff1690565b6100e6610384366004611054565b610b39565b606d8510156103c4576040517fb55ac75400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f805f61040589898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610b4d92505050565b6040805160608f901b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602080830191909152603482018690527fffff00000000000000000000000000000000000000000000000000000000000084166054830152825160368184030181526056909201909252805191012092955093505f9250309163233fd52d9150338d8d8d602d90606d926104a693929190611074565b8f8f606d9080926104b993929190611074565b6040518863ffffffff1660e01b81526004016104db97969594939291906110e2565b6020604051808303815f875af11580156104f7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061051b9190611154565b9050817dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916838b73ffffffffffffffffffffffffffffffffffffffff167f3617b009e9785c42daebadb6d3fb553243a4bf586d07ea72d65d80013ce116b58460405161058a911515815260200190565b60405180910390a450505050505050505050565b5f87815260036020526040812080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff89161790555a5f89815260036020526040808220805469ffffffffffffffffffff949094167601000000000000000000000000000000000000000000000275ffffffffffffffffffffffffffffffffffffffffffff909416939093179092559051819061065b908890889088908890602401611173565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f805f21320000000000000000000000000000000000000000000000000000000017815281519192505f918291828c5af15f9a8b5260036020526040909a2080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000008c151502179055509798975050505050505050565b6040805160c0810182525f808252602080830182905282840182905260608084018390526080840183905260a0840183905284519088901b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001681830152603481018790527fffff000000000000000000000000000000000000000000000000000000000000861660548201528451603681830301815260568201808752815191840191909120808552600390935285842060d68301909652945473ffffffffffffffffffffffffffffffffffffffff811680875274010000000000000000000000000000000000000000820460ff9081161515607685015275010000000000000000000000000000000000000000008304161515609684015276010000000000000000000000000000000000000000000090910469ffffffffffffffffffff1660b6909201919091529293909290919061089857505f6108c0565b8160200151156108aa575060026108c0565b81604001516108ba5760036108bd565b60015b90505b6040518060c001604052808481526020018260038111156108e3576108e3610f79565b8152602001835f015173ffffffffffffffffffffffffffffffffffffffff168152602001836020015115158152602001836040015115158152602001836060015169ffffffffffffffffffff1681525093505050509392505050565b610947610b68565b73ffffffffffffffffffffffffffffffffffffffff81165f8181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055517fb96d15bf9258c7b8df062753a6a262864611fc7b060a5ee2e57e79b85f898d389190a250565b6109c2610b68565b73ffffffffffffffffffffffffffffffffffffffff81165f8181526002602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f0ea0ce2c048ff45a4a95f2947879de3fb94abec2f152190400cab2d1272a68e79190a250565b60015473ffffffffffffffffffffffffffffffffffffffff163314610abe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610b41610b68565b610b4a81610bea565b50565b60218101516045820151608b90920151909260c09290921c91565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610be8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610ab5565b565b3373ffffffffffffffffffffffffffffffffffffffff821603610c69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610ab5565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d01575f80fd5b919050565b5f8083601f840112610d16575f80fd5b50813567ffffffffffffffff811115610d2d575f80fd5b602083019150836020828501011115610d44575f80fd5b9250929050565b5f805f805f805f6080888a031215610d61575f80fd5b610d6a88610cde565b9650602088013567ffffffffffffffff811115610d85575f80fd5b610d918a828b01610d06565b909750955050604088013567ffffffffffffffff811115610db0575f80fd5b610dbc8a828b01610d06565b909550935050606088013567ffffffffffffffff811115610ddb575f80fd5b8801601f81018a13610deb575f80fd5b803567ffffffffffffffff811115610e01575f80fd5b8a60208260051b8401011115610e15575f80fd5b602082019350809250505092959891949750929550565b602081525f82518060208401528060208501604085015e5f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f805f805f805f60a0888a031215610e95575f80fd5b87359650610ea560208901610cde565b9550610eb360408901610cde565b9450606088013567ffffffffffffffff811115610ece575f80fd5b610eda8a828b01610d06565b909550935050608088013567ffffffffffffffff811115610ef9575f80fd5b610f058a828b01610d06565b989b979a50959850939692959293505050565b5f805f60608486031215610f2a575f80fd5b610f3384610cde565b92506020840135915060408401357fffff00000000000000000000000000000000000000000000000000000000000081168114610f6e575f80fd5b809150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b81518152602082015160c082019060048110610fe9577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b8060208401525073ffffffffffffffffffffffffffffffffffffffff6040840151166040830152606083015115156060830152608083015161102f608084018215159052565b5060a083015161104d60a084018269ffffffffffffffffffff169052565b5092915050565b5f60208284031215611064575f80fd5b61106d82610cde565b9392505050565b5f8085851115611082575f80fd5b8386111561108e575f80fd5b5050820193919092039150565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b87815273ffffffffffffffffffffffffffffffffffffffff8716602082015273ffffffffffffffffffffffffffffffffffffffff8616604082015260a060608201525f61113360a08301868861109b565b828103608084015261114681858761109b565b9a9950505050505050505050565b5f60208284031215611164575f80fd5b8151801515811461106d575f80fd5b604081525f61118660408301868861109b565b828103602084015261119981858761109b565b97965050505050505056fea164736f6c634300081a000a",
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
