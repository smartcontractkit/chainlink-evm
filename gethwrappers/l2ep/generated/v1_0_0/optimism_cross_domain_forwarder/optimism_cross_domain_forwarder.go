// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimism_cross_domain_forwarder

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

var OptimismCrossDomainForwarderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"crossDomainMessengerAddr\",\"type\":\"address\",\"internalType\":\"contractiOVM_CrossDomainMessenger\"},{\"name\":\"l1OwnerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptL1Ownership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forward\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferL1Ownership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"L1OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801562000010575f80fd5b50604051620012c2380380620012c2833981016040819052620000339162000264565b8033805f816200008a5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000bc57620000bc8162000146565b505050620000d081620001f060201b60201c565b506001600160a01b038216620001335760405162461bcd60e51b815260206004820152602160248201527f496e76616c69642078446f6d61696e204d657373656e676572206164647265736044820152607360f81b606482015260840162000081565b506001600160a01b0316608052620002a1565b336001600160a01b03821603620001a05760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000081565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600280546001600160a01b038381166001600160a01b03198084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b6001600160a01b038116811462000261575f80fd5b50565b5f806040838503121562000276575f80fd5b825162000283816200024c565b602084015190925062000296816200024c565b809150509250929050565b608051610fe5620002dd5f395f818161016a015281816101c6015281816102830152818161050a015281816105c201526106fe0152610fe55ff3fe608060405234801561000f575f80fd5b506004361061009f575f3560e01c806396b8d7c011610072578063d2db637211610058578063d2db637214610137578063f2fde38b14610155578063f43b361314610168575f80fd5b806396b8d7c01461011c578063b2ec07871461012f575f80fd5b8063181f5a77146100a35780636fadcf72146100c157806379ba5097146100d65780638da5cb5b146100de575b5f80fd5b6100ab61018e565b6040516100b89190610deb565b60405180910390f35b6100d46100cf366004610e89565b6101ae565b005b6100d46103f6565b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b8565b6100d461012a366004610f65565b6104f2565b6100d46106fc565b60025473ffffffffffffffffffffffffffffffffffffffff166100f7565b6100d4610163366004610f65565b6108cf565b7f00000000000000000000000000000000000000000000000000000000000000006100f7565b6060604051806060016040528060228152602001610fb760229139905090565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610252576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572000060448201526064015b60405180910390fd5b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102ea573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061030e9190610f80565b73ffffffffffffffffffffffffffffffffffffffff16146103b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f65720000000000000000000000000000000000000000000000000000000000006064820152608401610249565b6103f182826040518060400160405280601781526020017f466f727761726465722063616c6c2072657665727465640000000000000000008152506108e0565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610477576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610249565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610591576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e67657200006044820152606401610249565b60025473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610629573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061064d9190610f80565b73ffffffffffffffffffffffffffffffffffffffff16146106f0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f78446f6d61696e2073656e646572206973206e6f7420746865204c31206f776e60448201527f65720000000000000000000000000000000000000000000000000000000000006064820152608401610249565b6106f9816108f8565b50565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff8116331461079c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e67657200006044820152606401610249565b600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff92831692841691636e296e459160048083019260209291908290030181865afa15801561080c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108309190610f80565b73ffffffffffffffffffffffffffffffffffffffff16146108ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4d7573742062652070726f706f736564204c31206f776e6572000000000000006044820152606401610249565b6003546106f99073ffffffffffffffffffffffffffffffffffffffff166109ed565b6108d7610a6e565b6106f981610af0565b60606108ee84845f85610be4565b90505b9392505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610977576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610249565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600254604051919216907f55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f905f90a350565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff00000000000000000000000000000000000000008084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610aee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610249565b565b3373ffffffffffffffffffffffffffffffffffffffff821603610b6f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610249565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606082471015610c76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610249565b73ffffffffffffffffffffffffffffffffffffffff85163b610cf4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610249565b5f808673ffffffffffffffffffffffffffffffffffffffff168587604051610d1c9190610f9b565b5f6040518083038185875af1925050503d805f8114610d56576040519150601f19603f3d011682016040523d82523d5f602084013e610d5b565b606091505b5091509150610d6b828286610d76565b979650505050505050565b60608315610d855750816108f1565b825115610d955782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102499190610deb565b5f5b83811015610de3578181015183820152602001610dcb565b50505f910152565b602081525f8251806020840152610e09816040850160208701610dc9565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b73ffffffffffffffffffffffffffffffffffffffff811681146106f9575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215610e9a575f80fd5b8235610ea581610e3b565b9150602083013567ffffffffffffffff80821115610ec1575f80fd5b818501915085601f830112610ed4575f80fd5b813581811115610ee657610ee6610e5c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610f2c57610f2c610e5c565b81604052828152886020848701011115610f44575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215610f75575f80fd5b81356108f181610e3b565b5f60208284031215610f90575f80fd5b81516108f181610e3b565b5f8251610fac818460208701610dc9565b919091019291505056fe4f7074696d69736d43726f7373446f6d61696e466f7277617264657220312e302e30a164736f6c6343000818000a",
}

var OptimismCrossDomainForwarderABI = OptimismCrossDomainForwarderMetaData.ABI

var OptimismCrossDomainForwarderBin = OptimismCrossDomainForwarderMetaData.Bin

func DeployOptimismCrossDomainForwarder(auth *bind.TransactOpts, backend bind.ContractBackend, crossDomainMessengerAddr common.Address, l1OwnerAddr common.Address) (common.Address, *types.Transaction, *OptimismCrossDomainForwarder, error) {
	parsed, err := OptimismCrossDomainForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismCrossDomainForwarderBin), backend, crossDomainMessengerAddr, l1OwnerAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismCrossDomainForwarder{address: address, abi: *parsed, OptimismCrossDomainForwarderCaller: OptimismCrossDomainForwarderCaller{contract: contract}, OptimismCrossDomainForwarderTransactor: OptimismCrossDomainForwarderTransactor{contract: contract}, OptimismCrossDomainForwarderFilterer: OptimismCrossDomainForwarderFilterer{contract: contract}}, nil
}

type OptimismCrossDomainForwarder struct {
	address common.Address
	abi     abi.ABI
	OptimismCrossDomainForwarderCaller
	OptimismCrossDomainForwarderTransactor
	OptimismCrossDomainForwarderFilterer
}

type OptimismCrossDomainForwarderCaller struct {
	contract *bind.BoundContract
}

type OptimismCrossDomainForwarderTransactor struct {
	contract *bind.BoundContract
}

type OptimismCrossDomainForwarderFilterer struct {
	contract *bind.BoundContract
}

type OptimismCrossDomainForwarderSession struct {
	Contract     *OptimismCrossDomainForwarder
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OptimismCrossDomainForwarderCallerSession struct {
	Contract *OptimismCrossDomainForwarderCaller
	CallOpts bind.CallOpts
}

type OptimismCrossDomainForwarderTransactorSession struct {
	Contract     *OptimismCrossDomainForwarderTransactor
	TransactOpts bind.TransactOpts
}

type OptimismCrossDomainForwarderRaw struct {
	Contract *OptimismCrossDomainForwarder
}

type OptimismCrossDomainForwarderCallerRaw struct {
	Contract *OptimismCrossDomainForwarderCaller
}

type OptimismCrossDomainForwarderTransactorRaw struct {
	Contract *OptimismCrossDomainForwarderTransactor
}

func NewOptimismCrossDomainForwarder(address common.Address, backend bind.ContractBackend) (*OptimismCrossDomainForwarder, error) {
	abi, err := abi.JSON(strings.NewReader(OptimismCrossDomainForwarderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOptimismCrossDomainForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainForwarder{address: address, abi: abi, OptimismCrossDomainForwarderCaller: OptimismCrossDomainForwarderCaller{contract: contract}, OptimismCrossDomainForwarderTransactor: OptimismCrossDomainForwarderTransactor{contract: contract}, OptimismCrossDomainForwarderFilterer: OptimismCrossDomainForwarderFilterer{contract: contract}}, nil
}

func NewOptimismCrossDomainForwarderCaller(address common.Address, caller bind.ContractCaller) (*OptimismCrossDomainForwarderCaller, error) {
	contract, err := bindOptimismCrossDomainForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainForwarderCaller{contract: contract}, nil
}

func NewOptimismCrossDomainForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismCrossDomainForwarderTransactor, error) {
	contract, err := bindOptimismCrossDomainForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainForwarderTransactor{contract: contract}, nil
}

func NewOptimismCrossDomainForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismCrossDomainForwarderFilterer, error) {
	contract, err := bindOptimismCrossDomainForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainForwarderFilterer{contract: contract}, nil
}

func bindOptimismCrossDomainForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismCrossDomainForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismCrossDomainForwarder.Contract.OptimismCrossDomainForwarderCaller.contract.Call(opts, result, method, params...)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.OptimismCrossDomainForwarderTransactor.contract.Transfer(opts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.OptimismCrossDomainForwarderTransactor.contract.Transact(opts, method, params...)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismCrossDomainForwarder.Contract.contract.Call(opts, result, method, params...)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.contract.Transfer(opts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.contract.Transact(opts, method, params...)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderCaller) CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismCrossDomainForwarder.contract.Call(opts, &out, "crossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderSession) CrossDomainMessenger() (common.Address, error) {
	return _OptimismCrossDomainForwarder.Contract.CrossDomainMessenger(&_OptimismCrossDomainForwarder.CallOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderCallerSession) CrossDomainMessenger() (common.Address, error) {
	return _OptimismCrossDomainForwarder.Contract.CrossDomainMessenger(&_OptimismCrossDomainForwarder.CallOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderCaller) L1Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismCrossDomainForwarder.contract.Call(opts, &out, "l1Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderSession) L1Owner() (common.Address, error) {
	return _OptimismCrossDomainForwarder.Contract.L1Owner(&_OptimismCrossDomainForwarder.CallOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderCallerSession) L1Owner() (common.Address, error) {
	return _OptimismCrossDomainForwarder.Contract.L1Owner(&_OptimismCrossDomainForwarder.CallOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismCrossDomainForwarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderSession) Owner() (common.Address, error) {
	return _OptimismCrossDomainForwarder.Contract.Owner(&_OptimismCrossDomainForwarder.CallOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderCallerSession) Owner() (common.Address, error) {
	return _OptimismCrossDomainForwarder.Contract.Owner(&_OptimismCrossDomainForwarder.CallOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OptimismCrossDomainForwarder.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderSession) TypeAndVersion() (string, error) {
	return _OptimismCrossDomainForwarder.Contract.TypeAndVersion(&_OptimismCrossDomainForwarder.CallOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderCallerSession) TypeAndVersion() (string, error) {
	return _OptimismCrossDomainForwarder.Contract.TypeAndVersion(&_OptimismCrossDomainForwarder.CallOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactor) AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.contract.Transact(opts, "acceptL1Ownership")
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.AcceptL1Ownership(&_OptimismCrossDomainForwarder.TransactOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.AcceptL1Ownership(&_OptimismCrossDomainForwarder.TransactOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.contract.Transact(opts, "acceptOwnership")
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderSession) AcceptOwnership() (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.AcceptOwnership(&_OptimismCrossDomainForwarder.TransactOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.AcceptOwnership(&_OptimismCrossDomainForwarder.TransactOpts)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactor) Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.contract.Transact(opts, "forward", target, data)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.Forward(&_OptimismCrossDomainForwarder.TransactOpts, target, data)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.Forward(&_OptimismCrossDomainForwarder.TransactOpts, target, data)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactor) TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.contract.Transact(opts, "transferL1Ownership", to)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.TransferL1Ownership(&_OptimismCrossDomainForwarder.TransactOpts, to)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.TransferL1Ownership(&_OptimismCrossDomainForwarder.TransactOpts, to)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.contract.Transact(opts, "transferOwnership", to)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.TransferOwnership(&_OptimismCrossDomainForwarder.TransactOpts, to)
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OptimismCrossDomainForwarder.Contract.TransferOwnership(&_OptimismCrossDomainForwarder.TransactOpts, to)
}

type OptimismCrossDomainForwarderL1OwnershipTransferRequestedIterator struct {
	Event *OptimismCrossDomainForwarderL1OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismCrossDomainForwarderL1OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismCrossDomainForwarderL1OwnershipTransferRequested)
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
		it.Event = new(OptimismCrossDomainForwarderL1OwnershipTransferRequested)
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

func (it *OptimismCrossDomainForwarderL1OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OptimismCrossDomainForwarderL1OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismCrossDomainForwarderL1OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainForwarderL1OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainForwarder.contract.FilterLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainForwarderL1OwnershipTransferRequestedIterator{contract: _OptimismCrossDomainForwarder.contract, event: "L1OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainForwarderL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainForwarder.contract.WatchLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismCrossDomainForwarderL1OwnershipTransferRequested)
				if err := _OptimismCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
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

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) ParseL1OwnershipTransferRequested(log types.Log) (*OptimismCrossDomainForwarderL1OwnershipTransferRequested, error) {
	event := new(OptimismCrossDomainForwarderL1OwnershipTransferRequested)
	if err := _OptimismCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismCrossDomainForwarderL1OwnershipTransferredIterator struct {
	Event *OptimismCrossDomainForwarderL1OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismCrossDomainForwarderL1OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismCrossDomainForwarderL1OwnershipTransferred)
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
		it.Event = new(OptimismCrossDomainForwarderL1OwnershipTransferred)
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

func (it *OptimismCrossDomainForwarderL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OptimismCrossDomainForwarderL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismCrossDomainForwarderL1OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainForwarderL1OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainForwarder.contract.FilterLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainForwarderL1OwnershipTransferredIterator{contract: _OptimismCrossDomainForwarder.contract, event: "L1OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainForwarderL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainForwarder.contract.WatchLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismCrossDomainForwarderL1OwnershipTransferred)
				if err := _OptimismCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
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

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) ParseL1OwnershipTransferred(log types.Log) (*OptimismCrossDomainForwarderL1OwnershipTransferred, error) {
	event := new(OptimismCrossDomainForwarderL1OwnershipTransferred)
	if err := _OptimismCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismCrossDomainForwarderOwnershipTransferRequestedIterator struct {
	Event *OptimismCrossDomainForwarderOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismCrossDomainForwarderOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismCrossDomainForwarderOwnershipTransferRequested)
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
		it.Event = new(OptimismCrossDomainForwarderOwnershipTransferRequested)
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

func (it *OptimismCrossDomainForwarderOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OptimismCrossDomainForwarderOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismCrossDomainForwarderOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainForwarderOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainForwarder.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainForwarderOwnershipTransferRequestedIterator{contract: _OptimismCrossDomainForwarder.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainForwarder.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismCrossDomainForwarderOwnershipTransferRequested)
				if err := _OptimismCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) ParseOwnershipTransferRequested(log types.Log) (*OptimismCrossDomainForwarderOwnershipTransferRequested, error) {
	event := new(OptimismCrossDomainForwarderOwnershipTransferRequested)
	if err := _OptimismCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismCrossDomainForwarderOwnershipTransferredIterator struct {
	Event *OptimismCrossDomainForwarderOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismCrossDomainForwarderOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismCrossDomainForwarderOwnershipTransferred)
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
		it.Event = new(OptimismCrossDomainForwarderOwnershipTransferred)
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

func (it *OptimismCrossDomainForwarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OptimismCrossDomainForwarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismCrossDomainForwarderOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainForwarderOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainForwarder.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismCrossDomainForwarderOwnershipTransferredIterator{contract: _OptimismCrossDomainForwarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismCrossDomainForwarder.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismCrossDomainForwarderOwnershipTransferred)
				if err := _OptimismCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarderFilterer) ParseOwnershipTransferred(log types.Log) (*OptimismCrossDomainForwarderOwnershipTransferred, error) {
	event := new(OptimismCrossDomainForwarderOwnershipTransferred)
	if err := _OptimismCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (OptimismCrossDomainForwarderL1OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f")
}

func (OptimismCrossDomainForwarderL1OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0xb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272")
}

func (OptimismCrossDomainForwarderOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OptimismCrossDomainForwarderOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_OptimismCrossDomainForwarder *OptimismCrossDomainForwarder) Address() common.Address {
	return _OptimismCrossDomainForwarder.address
}

type OptimismCrossDomainForwarderInterface interface {
	CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error)

	L1Owner(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error)

	TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainForwarderL1OwnershipTransferRequestedIterator, error)

	WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainForwarderL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferRequested(log types.Log) (*OptimismCrossDomainForwarderL1OwnershipTransferRequested, error)

	FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainForwarderL1OwnershipTransferredIterator, error)

	WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainForwarderL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferred(log types.Log) (*OptimismCrossDomainForwarderL1OwnershipTransferred, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainForwarderOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OptimismCrossDomainForwarderOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismCrossDomainForwarderOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismCrossDomainForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OptimismCrossDomainForwarderOwnershipTransferred, error)

	Address() common.Address
}
