// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_cross_domain_governor

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

var ArbitrumCrossDomainGovernorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"l1OwnerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptL1Ownership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forward\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forwardDelegate\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferL1Ownership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"L1OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610fac380380610fac83398101604081905261002e916101d2565b808033805f816100855760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b03848116919091179091558116156100b4576100b4816100ce565b5050506100c68161017660201b60201c565b5050506101ff565b336001600160a01b038216036101265760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007c565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600280546001600160a01b038381166001600160a01b03198084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f602082840312156101e2575f80fd5b81516001600160a01b03811681146101f8575f80fd5b9392505050565b610da08061020c5f395ff3fe608060405234801561000f575f80fd5b50600436106100b9575f3560e01c806396b8d7c011610072578063d2db637211610058578063d2db637214610164578063f2fde38b14610182578063f43b361314610195575f80fd5b806396b8d7c014610149578063b2ec07871461015c575f80fd5b80636fadcf72116100a25780636fadcf72146100f057806379ba5097146101035780638da5cb5b1461010b575f80fd5b8063181f5a77146100bd57806326929eb6146100db575b5f80fd5b6100c561019d565b6040516100d29190610bbb565b60405180910390f35b6100ee6100e9366004610c60565b6101bd565b005b6100ee6100fe366004610c60565b6102b4565b6100ee6103a1565b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100d2565b6100ee610157366004610d3a565b61049d565b6100ee610545565b60025473ffffffffffffffffffffffffffffffffffffffff16610124565b6100ee610190366004610d3a565b610633565b610124610644565b6060604051806060016040528060258152602001610d6f60259139905090565b6101c5610644565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061021457505f5473ffffffffffffffffffffffffffffffffffffffff1633145b6102a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572206f60448201527f72206f776e65720000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102af8282610685565b505050565b6102bc610644565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061030b57505f5473ffffffffffffffffffffffffffffffffffffffff1633145b610397576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572206f60448201527f72206f776e657200000000000000000000000000000000000000000000000000606482015260840161029c565b6102af8282610704565b60015473ffffffffffffffffffffffffffffffffffffffff163314610422576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161029c565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6104a5610644565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610539576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e6765720000604482015260640161029c565b61054281610718565b50565b60035461057b9073ffffffffffffffffffffffffffffffffffffffff167311110000000000000000000000000000000011110190565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461060f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4d7573742062652070726f706f736564204c31206f776e657200000000000000604482015260640161029c565b6003546106319073ffffffffffffffffffffffffffffffffffffffff1661080d565b565b61063b61088e565b6105428161090e565b5f61068061066760025473ffffffffffffffffffffffffffffffffffffffff1690565b7311110000000000000000000000000000000011110190565b905090565b60605f808473ffffffffffffffffffffffffffffffffffffffff16846040516106ae9190610d53565b5f60405180830381855af49150503d805f81146106e6576040519150601f19603f3d011682016040523d82523d5f602084013e6106eb565b606091505b50915091506106fb858383610a02565b95945050505050565b606061071183835f610a91565b9392505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610797576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161029c565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600254604051919216907f55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f905f90a350565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff00000000000000000000000000000000000000008084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610631576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161029c565b3373ffffffffffffffffffffffffffffffffffffffff82160361098d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161029c565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606082610a1757610a1282610b57565b610711565b8151158015610a3b575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610a8a576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161029c565b5092915050565b606081471015610ad6576040517fcf4791810000000000000000000000000000000000000000000000000000000081524760048201526024810183905260440161029c565b5f808573ffffffffffffffffffffffffffffffffffffffff168486604051610afe9190610d53565b5f6040518083038185875af1925050503d805f8114610b38576040519150601f19603f3d011682016040523d82523d5f602084013e610b3d565b606091505b5091509150610b4d868383610a02565b9695505050505050565b805115610b675780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b83811015610bb3578181015183820152602001610b9b565b50505f910152565b602081525f8251806020840152610bd9816040850160208701610b99565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c2e575f80fd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215610c71575f80fd5b610c7a83610c0b565b9150602083013567ffffffffffffffff80821115610c96575f80fd5b818501915085601f830112610ca9575f80fd5b813581811115610cbb57610cbb610c33565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610d0157610d01610c33565b81604052828152886020848701011115610d19575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215610d4a575f80fd5b61071182610c0b565b5f8251610d64818460208701610b99565b919091019291505056fe417262697472756d43726f7373446f6d61696e476f7665726e6f7220312e312e302d646576a164736f6c6343000818000a",
}

var ArbitrumCrossDomainGovernorABI = ArbitrumCrossDomainGovernorMetaData.ABI

var ArbitrumCrossDomainGovernorBin = ArbitrumCrossDomainGovernorMetaData.Bin

func DeployArbitrumCrossDomainGovernor(auth *bind.TransactOpts, backend bind.ContractBackend, l1OwnerAddr common.Address) (common.Address, *types.Transaction, *ArbitrumCrossDomainGovernor, error) {
	parsed, err := ArbitrumCrossDomainGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumCrossDomainGovernorBin), backend, l1OwnerAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumCrossDomainGovernor{address: address, abi: *parsed, ArbitrumCrossDomainGovernorCaller: ArbitrumCrossDomainGovernorCaller{contract: contract}, ArbitrumCrossDomainGovernorTransactor: ArbitrumCrossDomainGovernorTransactor{contract: contract}, ArbitrumCrossDomainGovernorFilterer: ArbitrumCrossDomainGovernorFilterer{contract: contract}}, nil
}

type ArbitrumCrossDomainGovernor struct {
	address common.Address
	abi     abi.ABI
	ArbitrumCrossDomainGovernorCaller
	ArbitrumCrossDomainGovernorTransactor
	ArbitrumCrossDomainGovernorFilterer
}

type ArbitrumCrossDomainGovernorCaller struct {
	contract *bind.BoundContract
}

type ArbitrumCrossDomainGovernorTransactor struct {
	contract *bind.BoundContract
}

type ArbitrumCrossDomainGovernorFilterer struct {
	contract *bind.BoundContract
}

type ArbitrumCrossDomainGovernorSession struct {
	Contract     *ArbitrumCrossDomainGovernor
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbitrumCrossDomainGovernorCallerSession struct {
	Contract *ArbitrumCrossDomainGovernorCaller
	CallOpts bind.CallOpts
}

type ArbitrumCrossDomainGovernorTransactorSession struct {
	Contract     *ArbitrumCrossDomainGovernorTransactor
	TransactOpts bind.TransactOpts
}

type ArbitrumCrossDomainGovernorRaw struct {
	Contract *ArbitrumCrossDomainGovernor
}

type ArbitrumCrossDomainGovernorCallerRaw struct {
	Contract *ArbitrumCrossDomainGovernorCaller
}

type ArbitrumCrossDomainGovernorTransactorRaw struct {
	Contract *ArbitrumCrossDomainGovernorTransactor
}

func NewArbitrumCrossDomainGovernor(address common.Address, backend bind.ContractBackend) (*ArbitrumCrossDomainGovernor, error) {
	abi, err := abi.JSON(strings.NewReader(ArbitrumCrossDomainGovernorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindArbitrumCrossDomainGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernor{address: address, abi: abi, ArbitrumCrossDomainGovernorCaller: ArbitrumCrossDomainGovernorCaller{contract: contract}, ArbitrumCrossDomainGovernorTransactor: ArbitrumCrossDomainGovernorTransactor{contract: contract}, ArbitrumCrossDomainGovernorFilterer: ArbitrumCrossDomainGovernorFilterer{contract: contract}}, nil
}

func NewArbitrumCrossDomainGovernorCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumCrossDomainGovernorCaller, error) {
	contract, err := bindArbitrumCrossDomainGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorCaller{contract: contract}, nil
}

func NewArbitrumCrossDomainGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumCrossDomainGovernorTransactor, error) {
	contract, err := bindArbitrumCrossDomainGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorTransactor{contract: contract}, nil
}

func NewArbitrumCrossDomainGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumCrossDomainGovernorFilterer, error) {
	contract, err := bindArbitrumCrossDomainGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorFilterer{contract: contract}, nil
}

func bindArbitrumCrossDomainGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrumCrossDomainGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumCrossDomainGovernor.Contract.ArbitrumCrossDomainGovernorCaller.contract.Call(opts, result, method, params...)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.ArbitrumCrossDomainGovernorTransactor.contract.Transfer(opts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.ArbitrumCrossDomainGovernorTransactor.contract.Transact(opts, method, params...)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumCrossDomainGovernor.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.contract.Transfer(opts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.contract.Transact(opts, method, params...)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCaller) CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainGovernor.contract.Call(opts, &out, "crossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) CrossDomainMessenger() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.CrossDomainMessenger(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerSession) CrossDomainMessenger() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.CrossDomainMessenger(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCaller) L1Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainGovernor.contract.Call(opts, &out, "l1Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) L1Owner() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.L1Owner(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerSession) L1Owner() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.L1Owner(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainGovernor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) Owner() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.Owner(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerSession) Owner() (common.Address, error) {
	return _ArbitrumCrossDomainGovernor.Contract.Owner(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainGovernor.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) TypeAndVersion() (string, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TypeAndVersion(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorCallerSession) TypeAndVersion() (string, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TypeAndVersion(&_ArbitrumCrossDomainGovernor.CallOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "acceptL1Ownership")
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.AcceptL1Ownership(&_ArbitrumCrossDomainGovernor.TransactOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.AcceptL1Ownership(&_ArbitrumCrossDomainGovernor.TransactOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "acceptOwnership")
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.AcceptOwnership(&_ArbitrumCrossDomainGovernor.TransactOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.AcceptOwnership(&_ArbitrumCrossDomainGovernor.TransactOpts)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "forward", target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.Forward(&_ArbitrumCrossDomainGovernor.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.Forward(&_ArbitrumCrossDomainGovernor.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) ForwardDelegate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "forwardDelegate", target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) ForwardDelegate(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.ForwardDelegate(&_ArbitrumCrossDomainGovernor.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) ForwardDelegate(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.ForwardDelegate(&_ArbitrumCrossDomainGovernor.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "transferL1Ownership", to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TransferL1Ownership(&_ArbitrumCrossDomainGovernor.TransactOpts, to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TransferL1Ownership(&_ArbitrumCrossDomainGovernor.TransactOpts, to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.contract.Transact(opts, "transferOwnership", to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TransferOwnership(&_ArbitrumCrossDomainGovernor.TransactOpts, to)
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainGovernor.Contract.TransferOwnership(&_ArbitrumCrossDomainGovernor.TransactOpts, to)
}

type ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator struct {
	Event *ArbitrumCrossDomainGovernorL1OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainGovernorL1OwnershipTransferRequested)
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
		it.Event = new(ArbitrumCrossDomainGovernorL1OwnershipTransferRequested)
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

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainGovernorL1OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.FilterLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator{contract: _ArbitrumCrossDomainGovernor.contract, event: "L1OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.WatchLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainGovernorL1OwnershipTransferRequested)
				if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
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

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) ParseL1OwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainGovernorL1OwnershipTransferRequested, error) {
	event := new(ArbitrumCrossDomainGovernorL1OwnershipTransferRequested)
	if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator struct {
	Event *ArbitrumCrossDomainGovernorL1OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainGovernorL1OwnershipTransferred)
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
		it.Event = new(ArbitrumCrossDomainGovernorL1OwnershipTransferred)
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

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainGovernorL1OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.FilterLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator{contract: _ArbitrumCrossDomainGovernor.contract, event: "L1OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.WatchLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainGovernorL1OwnershipTransferred)
				if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
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

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) ParseL1OwnershipTransferred(log types.Log) (*ArbitrumCrossDomainGovernorL1OwnershipTransferred, error) {
	event := new(ArbitrumCrossDomainGovernorL1OwnershipTransferred)
	if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator struct {
	Event *ArbitrumCrossDomainGovernorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainGovernorOwnershipTransferRequested)
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
		it.Event = new(ArbitrumCrossDomainGovernorOwnershipTransferRequested)
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

func (it *ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainGovernorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator{contract: _ArbitrumCrossDomainGovernor.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainGovernorOwnershipTransferRequested)
				if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainGovernorOwnershipTransferRequested, error) {
	event := new(ArbitrumCrossDomainGovernorOwnershipTransferRequested)
	if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumCrossDomainGovernorOwnershipTransferredIterator struct {
	Event *ArbitrumCrossDomainGovernorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainGovernorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainGovernorOwnershipTransferred)
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
		it.Event = new(ArbitrumCrossDomainGovernorOwnershipTransferred)
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

func (it *ArbitrumCrossDomainGovernorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainGovernorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainGovernorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainGovernorOwnershipTransferredIterator{contract: _ArbitrumCrossDomainGovernor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainGovernor.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainGovernorOwnershipTransferred)
				if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernorFilterer) ParseOwnershipTransferred(log types.Log) (*ArbitrumCrossDomainGovernorOwnershipTransferred, error) {
	event := new(ArbitrumCrossDomainGovernorOwnershipTransferred)
	if err := _ArbitrumCrossDomainGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ArbitrumCrossDomainGovernorL1OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f")
}

func (ArbitrumCrossDomainGovernorL1OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0xb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272")
}

func (ArbitrumCrossDomainGovernorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ArbitrumCrossDomainGovernorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_ArbitrumCrossDomainGovernor *ArbitrumCrossDomainGovernor) Address() common.Address {
	return _ArbitrumCrossDomainGovernor.address
}

type ArbitrumCrossDomainGovernorInterface interface {
	CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error)

	L1Owner(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error)

	ForwardDelegate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error)

	TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorL1OwnershipTransferRequestedIterator, error)

	WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainGovernorL1OwnershipTransferRequested, error)

	FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorL1OwnershipTransferredIterator, error)

	WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferred(log types.Log) (*ArbitrumCrossDomainGovernorL1OwnershipTransferred, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainGovernorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainGovernorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainGovernorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ArbitrumCrossDomainGovernorOwnershipTransferred, error)

	Address() common.Address
}
