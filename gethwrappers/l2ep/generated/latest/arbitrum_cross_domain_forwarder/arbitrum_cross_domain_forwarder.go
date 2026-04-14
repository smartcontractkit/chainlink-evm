// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_cross_domain_forwarder

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

var ArbitrumCrossDomainForwarderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"l1OwnerAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptL1Ownership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forward\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferL1Ownership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"L1OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610dcb380380610dcb83398101604081905261002e916101d0565b8033805f816100845760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b03848116919091179091558116156100b3576100b3816100cc565b5050506100c58161017460201b60201c565b50506101fd565b336001600160a01b038216036101245760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161007b565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600280546001600160a01b038381166001600160a01b03198084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f602082840312156101e0575f80fd5b81516001600160a01b03811681146101f6575f80fd5b9392505050565b610bc18061020a5f395ff3fe608060405234801561000f575f80fd5b506004361061009f575f3560e01c806396b8d7c011610072578063d2db637211610058578063d2db637214610137578063f2fde38b14610155578063f43b361314610168575f80fd5b806396b8d7c01461011c578063b2ec07871461012f575f80fd5b8063181f5a77146100a35780636fadcf72146100c157806379ba5097146100d65780638da5cb5b146100de575b5f80fd5b6100ab610170565b6040516100b891906109db565b60405180910390f35b6100d46100cf366004610a80565b610190565b005b6100d4610240565b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100b8565b6100d461012a366004610b5a565b61033c565b6100d46103e4565b60025473ffffffffffffffffffffffffffffffffffffffff166100f7565b6100d4610163366004610b5a565b6104d2565b6100f76104e3565b6060604051806060016040528060268152602001610b8f60269139905090565b6101986104e3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610231576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e676572000060448201526064015b60405180910390fd5b61023b8282610524565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146102c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610228565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6103446104e3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f53656e646572206973206e6f7420746865204c32206d657373656e67657200006044820152606401610228565b6103e181610538565b50565b60035461041a9073ffffffffffffffffffffffffffffffffffffffff167311110000000000000000000000000000000011110190565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4d7573742062652070726f706f736564204c31206f776e6572000000000000006044820152606401610228565b6003546104d09073ffffffffffffffffffffffffffffffffffffffff1661062d565b565b6104da6106ae565b6103e18161072e565b5f61051f61050660025473ffffffffffffffffffffffffffffffffffffffff1690565b7311110000000000000000000000000000000011110190565b905090565b606061053183835f610822565b9392505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036105b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610228565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217909255600254604051919216907f55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f905f90a350565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff00000000000000000000000000000000000000008084168217909455600380549094169093556040519116919082907fb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146104d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610228565b3373ffffffffffffffffffffffffffffffffffffffff8216036107ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610228565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606081471015610867576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101839052604401610228565b5f808573ffffffffffffffffffffffffffffffffffffffff16848660405161088f9190610b73565b5f6040518083038185875af1925050503d805f81146108c9576040519150601f19603f3d011682016040523d82523d5f602084013e6108ce565b606091505b50915091506108de8683836108e8565b9695505050505050565b6060826108fd576108f882610977565b610531565b8151158015610921575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610970576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610228565b5080610531565b8051156109875780518082602001fd5b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b838110156109d35781810151838201526020016109bb565b50505f910152565b602081525f82518060208401526109f98160408501602087016109b9565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610a4e575f80fd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215610a91575f80fd5b610a9a83610a2b565b9150602083013567ffffffffffffffff80821115610ab6575f80fd5b818501915085601f830112610ac9575f80fd5b813581811115610adb57610adb610a53565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610b2157610b21610a53565b81604052828152886020848701011115610b39575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215610b6a575f80fd5b61053182610a2b565b5f8251610b848184602087016109b9565b919091019291505056fe417262697472756d43726f7373446f6d61696e466f7277617264657220312e312e302d646576a164736f6c6343000818000a",
}

var ArbitrumCrossDomainForwarderABI = ArbitrumCrossDomainForwarderMetaData.ABI

var ArbitrumCrossDomainForwarderBin = ArbitrumCrossDomainForwarderMetaData.Bin

func DeployArbitrumCrossDomainForwarder(auth *bind.TransactOpts, backend bind.ContractBackend, l1OwnerAddr common.Address) (common.Address, *types.Transaction, *ArbitrumCrossDomainForwarder, error) {
	parsed, err := ArbitrumCrossDomainForwarderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumCrossDomainForwarderBin), backend, l1OwnerAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumCrossDomainForwarder{address: address, abi: *parsed, ArbitrumCrossDomainForwarderCaller: ArbitrumCrossDomainForwarderCaller{contract: contract}, ArbitrumCrossDomainForwarderTransactor: ArbitrumCrossDomainForwarderTransactor{contract: contract}, ArbitrumCrossDomainForwarderFilterer: ArbitrumCrossDomainForwarderFilterer{contract: contract}}, nil
}

type ArbitrumCrossDomainForwarder struct {
	address common.Address
	abi     abi.ABI
	ArbitrumCrossDomainForwarderCaller
	ArbitrumCrossDomainForwarderTransactor
	ArbitrumCrossDomainForwarderFilterer
}

type ArbitrumCrossDomainForwarderCaller struct {
	contract *bind.BoundContract
}

type ArbitrumCrossDomainForwarderTransactor struct {
	contract *bind.BoundContract
}

type ArbitrumCrossDomainForwarderFilterer struct {
	contract *bind.BoundContract
}

type ArbitrumCrossDomainForwarderSession struct {
	Contract     *ArbitrumCrossDomainForwarder
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbitrumCrossDomainForwarderCallerSession struct {
	Contract *ArbitrumCrossDomainForwarderCaller
	CallOpts bind.CallOpts
}

type ArbitrumCrossDomainForwarderTransactorSession struct {
	Contract     *ArbitrumCrossDomainForwarderTransactor
	TransactOpts bind.TransactOpts
}

type ArbitrumCrossDomainForwarderRaw struct {
	Contract *ArbitrumCrossDomainForwarder
}

type ArbitrumCrossDomainForwarderCallerRaw struct {
	Contract *ArbitrumCrossDomainForwarderCaller
}

type ArbitrumCrossDomainForwarderTransactorRaw struct {
	Contract *ArbitrumCrossDomainForwarderTransactor
}

func NewArbitrumCrossDomainForwarder(address common.Address, backend bind.ContractBackend) (*ArbitrumCrossDomainForwarder, error) {
	abi, err := abi.JSON(strings.NewReader(ArbitrumCrossDomainForwarderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindArbitrumCrossDomainForwarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainForwarder{address: address, abi: abi, ArbitrumCrossDomainForwarderCaller: ArbitrumCrossDomainForwarderCaller{contract: contract}, ArbitrumCrossDomainForwarderTransactor: ArbitrumCrossDomainForwarderTransactor{contract: contract}, ArbitrumCrossDomainForwarderFilterer: ArbitrumCrossDomainForwarderFilterer{contract: contract}}, nil
}

func NewArbitrumCrossDomainForwarderCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumCrossDomainForwarderCaller, error) {
	contract, err := bindArbitrumCrossDomainForwarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainForwarderCaller{contract: contract}, nil
}

func NewArbitrumCrossDomainForwarderTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumCrossDomainForwarderTransactor, error) {
	contract, err := bindArbitrumCrossDomainForwarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainForwarderTransactor{contract: contract}, nil
}

func NewArbitrumCrossDomainForwarderFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumCrossDomainForwarderFilterer, error) {
	contract, err := bindArbitrumCrossDomainForwarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainForwarderFilterer{contract: contract}, nil
}

func bindArbitrumCrossDomainForwarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrumCrossDomainForwarderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumCrossDomainForwarder.Contract.ArbitrumCrossDomainForwarderCaller.contract.Call(opts, result, method, params...)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.ArbitrumCrossDomainForwarderTransactor.contract.Transfer(opts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.ArbitrumCrossDomainForwarderTransactor.contract.Transact(opts, method, params...)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumCrossDomainForwarder.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.contract.Transfer(opts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.contract.Transact(opts, method, params...)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderCaller) CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainForwarder.contract.Call(opts, &out, "crossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderSession) CrossDomainMessenger() (common.Address, error) {
	return _ArbitrumCrossDomainForwarder.Contract.CrossDomainMessenger(&_ArbitrumCrossDomainForwarder.CallOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderCallerSession) CrossDomainMessenger() (common.Address, error) {
	return _ArbitrumCrossDomainForwarder.Contract.CrossDomainMessenger(&_ArbitrumCrossDomainForwarder.CallOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderCaller) L1Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainForwarder.contract.Call(opts, &out, "l1Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderSession) L1Owner() (common.Address, error) {
	return _ArbitrumCrossDomainForwarder.Contract.L1Owner(&_ArbitrumCrossDomainForwarder.CallOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderCallerSession) L1Owner() (common.Address, error) {
	return _ArbitrumCrossDomainForwarder.Contract.L1Owner(&_ArbitrumCrossDomainForwarder.CallOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainForwarder.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderSession) Owner() (common.Address, error) {
	return _ArbitrumCrossDomainForwarder.Contract.Owner(&_ArbitrumCrossDomainForwarder.CallOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderCallerSession) Owner() (common.Address, error) {
	return _ArbitrumCrossDomainForwarder.Contract.Owner(&_ArbitrumCrossDomainForwarder.CallOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArbitrumCrossDomainForwarder.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderSession) TypeAndVersion() (string, error) {
	return _ArbitrumCrossDomainForwarder.Contract.TypeAndVersion(&_ArbitrumCrossDomainForwarder.CallOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderCallerSession) TypeAndVersion() (string, error) {
	return _ArbitrumCrossDomainForwarder.Contract.TypeAndVersion(&_ArbitrumCrossDomainForwarder.CallOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactor) AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.contract.Transact(opts, "acceptL1Ownership")
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.AcceptL1Ownership(&_ArbitrumCrossDomainForwarder.TransactOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactorSession) AcceptL1Ownership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.AcceptL1Ownership(&_ArbitrumCrossDomainForwarder.TransactOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.contract.Transact(opts, "acceptOwnership")
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.AcceptOwnership(&_ArbitrumCrossDomainForwarder.TransactOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.AcceptOwnership(&_ArbitrumCrossDomainForwarder.TransactOpts)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactor) Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.contract.Transact(opts, "forward", target, data)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.Forward(&_ArbitrumCrossDomainForwarder.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactorSession) Forward(target common.Address, data []byte) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.Forward(&_ArbitrumCrossDomainForwarder.TransactOpts, target, data)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactor) TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.contract.Transact(opts, "transferL1Ownership", to)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.TransferL1Ownership(&_ArbitrumCrossDomainForwarder.TransactOpts, to)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactorSession) TransferL1Ownership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.TransferL1Ownership(&_ArbitrumCrossDomainForwarder.TransactOpts, to)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.contract.Transact(opts, "transferOwnership", to)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.TransferOwnership(&_ArbitrumCrossDomainForwarder.TransactOpts, to)
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumCrossDomainForwarder.Contract.TransferOwnership(&_ArbitrumCrossDomainForwarder.TransactOpts, to)
}

type ArbitrumCrossDomainForwarderL1OwnershipTransferRequestedIterator struct {
	Event *ArbitrumCrossDomainForwarderL1OwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainForwarderL1OwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainForwarderL1OwnershipTransferRequested)
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
		it.Event = new(ArbitrumCrossDomainForwarderL1OwnershipTransferRequested)
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

func (it *ArbitrumCrossDomainForwarderL1OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainForwarderL1OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainForwarderL1OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainForwarderL1OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainForwarder.contract.FilterLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainForwarderL1OwnershipTransferRequestedIterator{contract: _ArbitrumCrossDomainForwarder.contract, event: "L1OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainForwarderL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainForwarder.contract.WatchLogs(opts, "L1OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainForwarderL1OwnershipTransferRequested)
				if err := _ArbitrumCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
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

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) ParseL1OwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainForwarderL1OwnershipTransferRequested, error) {
	event := new(ArbitrumCrossDomainForwarderL1OwnershipTransferRequested)
	if err := _ArbitrumCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumCrossDomainForwarderL1OwnershipTransferredIterator struct {
	Event *ArbitrumCrossDomainForwarderL1OwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainForwarderL1OwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainForwarderL1OwnershipTransferred)
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
		it.Event = new(ArbitrumCrossDomainForwarderL1OwnershipTransferred)
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

func (it *ArbitrumCrossDomainForwarderL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainForwarderL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainForwarderL1OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainForwarderL1OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainForwarder.contract.FilterLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainForwarderL1OwnershipTransferredIterator{contract: _ArbitrumCrossDomainForwarder.contract, event: "L1OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainForwarderL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainForwarder.contract.WatchLogs(opts, "L1OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainForwarderL1OwnershipTransferred)
				if err := _ArbitrumCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
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

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) ParseL1OwnershipTransferred(log types.Log) (*ArbitrumCrossDomainForwarderL1OwnershipTransferred, error) {
	event := new(ArbitrumCrossDomainForwarderL1OwnershipTransferred)
	if err := _ArbitrumCrossDomainForwarder.contract.UnpackLog(event, "L1OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumCrossDomainForwarderOwnershipTransferRequestedIterator struct {
	Event *ArbitrumCrossDomainForwarderOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainForwarderOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainForwarderOwnershipTransferRequested)
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
		it.Event = new(ArbitrumCrossDomainForwarderOwnershipTransferRequested)
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

func (it *ArbitrumCrossDomainForwarderOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainForwarderOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainForwarderOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainForwarderOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainForwarder.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainForwarderOwnershipTransferRequestedIterator{contract: _ArbitrumCrossDomainForwarder.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainForwarder.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainForwarderOwnershipTransferRequested)
				if err := _ArbitrumCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) ParseOwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainForwarderOwnershipTransferRequested, error) {
	event := new(ArbitrumCrossDomainForwarderOwnershipTransferRequested)
	if err := _ArbitrumCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumCrossDomainForwarderOwnershipTransferredIterator struct {
	Event *ArbitrumCrossDomainForwarderOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumCrossDomainForwarderOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumCrossDomainForwarderOwnershipTransferred)
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
		it.Event = new(ArbitrumCrossDomainForwarderOwnershipTransferred)
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

func (it *ArbitrumCrossDomainForwarderOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumCrossDomainForwarderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumCrossDomainForwarderOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainForwarderOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainForwarder.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumCrossDomainForwarderOwnershipTransferredIterator{contract: _ArbitrumCrossDomainForwarder.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumCrossDomainForwarder.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumCrossDomainForwarderOwnershipTransferred)
				if err := _ArbitrumCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarderFilterer) ParseOwnershipTransferred(log types.Log) (*ArbitrumCrossDomainForwarderOwnershipTransferred, error) {
	event := new(ArbitrumCrossDomainForwarderOwnershipTransferred)
	if err := _ArbitrumCrossDomainForwarder.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ArbitrumCrossDomainForwarderL1OwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0x55eb94b097b1cd6441a2403b6e04742bd36cfd6e6905ea13cbc5879d26e5b20f")
}

func (ArbitrumCrossDomainForwarderL1OwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0xb0121dbcab67289d5cb13a1f35ca086715f35ef418e0eda134c4145a086b6272")
}

func (ArbitrumCrossDomainForwarderOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ArbitrumCrossDomainForwarderOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_ArbitrumCrossDomainForwarder *ArbitrumCrossDomainForwarder) Address() common.Address {
	return _ArbitrumCrossDomainForwarder.address
}

type ArbitrumCrossDomainForwarderInterface interface {
	CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error)

	L1Owner(opts *bind.CallOpts) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptL1Ownership(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Forward(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error)

	TransferL1Ownership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterL1OwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainForwarderL1OwnershipTransferRequestedIterator, error)

	WatchL1OwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainForwarderL1OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainForwarderL1OwnershipTransferRequested, error)

	FilterL1OwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainForwarderL1OwnershipTransferredIterator, error)

	WatchL1OwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainForwarderL1OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1OwnershipTransferred(log types.Log) (*ArbitrumCrossDomainForwarderL1OwnershipTransferred, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainForwarderOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainForwarderOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ArbitrumCrossDomainForwarderOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumCrossDomainForwarderOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumCrossDomainForwarderOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ArbitrumCrossDomainForwarderOwnershipTransferred, error)

	Address() common.Address
}
