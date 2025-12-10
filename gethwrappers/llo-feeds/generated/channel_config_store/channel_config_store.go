// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package channel_config_store

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

var ChannelConfigStoreMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addChannelDefinitions\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"channelAdderId\",\"type\":\"uint32\",\"internalType\":\"IChannelConfigStore.ChannelAdderId\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sha\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllowedChannelAdders\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"allowedChannelAdderIds\",\"type\":\"uint32[]\",\"internalType\":\"IChannelConfigStore.ChannelAdderId[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChannelAdderAddress\",\"inputs\":[{\"name\":\"channelAdderId\",\"type\":\"uint32\",\"internalType\":\"IChannelConfigStore.ChannelAdderId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isChannelAdderAllowed\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"channelAdderId\",\"type\":\"uint32\",\"internalType\":\"IChannelConfigStore.ChannelAdderId\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChannelAdder\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"channelAdderId\",\"type\":\"uint32\",\"internalType\":\"IChannelConfigStore.ChannelAdderId\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChannelAdderAddress\",\"inputs\":[{\"name\":\"channelAdderId\",\"type\":\"uint32\",\"internalType\":\"IChannelConfigStore.ChannelAdderId\"},{\"name\":\"adderAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChannelDefinitions\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"sha\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"ChannelAdderAddressSet\",\"inputs\":[{\"name\":\"channelAdderId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"IChannelConfigStore.ChannelAdderId\"},{\"name\":\"adderAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelAdderSet\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"channelAdderId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"IChannelConfigStore.ChannelAdderId\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelDefinitionAdded\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"channelAdderId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"IChannelConfigStore.ChannelAdderId\"},{\"name\":\"url\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sha\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewChannelDefinition\",\"inputs\":[{\"name\":\"donId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"version\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"url\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"sha\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReservedChannelAdderId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedChannelAdder\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b611047806101576000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c80638ea4179d11610081578063a6a2d7021161005b578063a6a2d70214610253578063bfe37b8d14610266578063f2fde38b1461028657600080fd5b80638ea4179d146101f15780639050a7171461022d57806399bb1be31461024057600080fd5b80635ec39ed6116100b25780635ec39ed61461019757806379ba5097146101aa5780638da5cb5b146101b257600080fd5b806301ffc9a7146100d9578063181f5a77146101435780635ba5bac214610182575b600080fd5b61012e6100e7366004610b5a565b7fffffffff00000000000000000000000000000000000000000000000000000000167f9b6823f2000000000000000000000000000000000000000000000000000000001490565b60405190151581526020015b60405180910390f35b604080518082018252601881527f4368616e6e656c436f6e66696753746f726520312e302e3000000000000000006020820152905161013a9190610b9c565b610195610190366004610c63565b610299565b005b6101956101a5366004610ce8565b610316565b6101956103ee565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161013a565b6101cc6101ff366004610d1d565b63ffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b61019561023b366004610d3a565b6104f0565b61019561024e366004610d8a565b6105de565b61012e610261366004610df7565b6106ec565b610279610274366004610d1d565b61071c565b60405161013a9190610e30565b610195610294366004610e7a565b6107ec565b6102a1610800565b63ffffffff84166000908152600260205260408120805482906102c390610ec4565b91905081905590508463ffffffff167fe5b641a7879fb491e4e5a35a1ce950f0237b2537ee9b1b1e4fb65e29aff1f5e8828686866040516103079493929190610f45565b60405180910390a25050505050565b61031e610800565b6103e863ffffffff83161015610360576040517f405cfed700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff821660008181526003602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff861690811790915591519182527fd3c5db69981921b17e39aedfb7c2663d7605477421099c01ad72c980b469f8fe910160405180910390a25050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610474576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6104f8610800565b6103e863ffffffff8316101561053a576040517f405cfed700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b801561056a5763ffffffff8381166000908152600460205260409020610564918085169061088316565b50610590565b63ffffffff838116600090815260046020526040902061058e918085169061088f16565b505b8163ffffffff168363ffffffff167f53af2fd88370b4b325f7cbd06f89b37ae7a5ebd9b6ee7cb19f6ce9bc1c7ef02b836040516105d1911515815260200190565b60405180910390a3505050565b63ffffffff841660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610642576040517f3dc818df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff8581166000908152600460205260409020610666918087169061089b16565b61069c576040517f3dc818df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8363ffffffff168563ffffffff167f5e9e3e3cf7859ce74906f92017b202f25c2d278c32ef43735b43d6b35a49764e8585856040516106dd93929190610f76565b60405180910390a35050505050565b63ffffffff8281166000908152600460205260408120909161071391908085169061089b16565b90505b92915050565b63ffffffff811660009081526004602052604081206060919061073e906108b3565b9050805167ffffffffffffffff81111561075a5761075a610f9a565b604051908082528060200260200182016040528015610783578160200160208202803683370190505b50915060005b81518110156107e5578181815181106107a4576107a4610fc9565b60200260200101518382815181106107be576107be610fc9565b63ffffffff90921660209283029190910190910152806107dd81610ec4565b915050610789565b5050919050565b6107f4610800565b6107fd816108c7565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610881576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161046b565b565b600061071383836109bc565b60006107138383610a0b565b60008181526001830160205260408120541515610713565b606060006108c083610afe565b9392505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610946576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161046b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000818152600183016020526040812054610a0357508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610716565b506000610716565b60008181526001830160205260408120548015610af4576000610a2f600183610ff8565b8554909150600090610a4390600190610ff8565b9050818114610aa8576000866000018281548110610a6357610a63610fc9565b9060005260206000200154905080876000018481548110610a8657610a86610fc9565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610ab957610ab961100b565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610716565b6000915050610716565b606081600001805480602002602001604051908101604052809291908181526020018280548015610b4e57602002820191906000526020600020905b815481526020019060010190808311610b3a575b50505050509050919050565b600060208284031215610b6c57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146108c057600080fd5b600060208083528351808285015260005b81811015610bc957858101830151858201604001528201610bad565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b63ffffffff811681146107fd57600080fd5b60008083601f840112610c2c57600080fd5b50813567ffffffffffffffff811115610c4457600080fd5b602083019150836020828501011115610c5c57600080fd5b9250929050565b60008060008060608587031215610c7957600080fd5b8435610c8481610c08565b9350602085013567ffffffffffffffff811115610ca057600080fd5b610cac87828801610c1a565b9598909750949560400135949350505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610ce357600080fd5b919050565b60008060408385031215610cfb57600080fd5b8235610d0681610c08565b9150610d1460208401610cbf565b90509250929050565b600060208284031215610d2f57600080fd5b81356108c081610c08565b600080600060608486031215610d4f57600080fd5b8335610d5a81610c08565b92506020840135610d6a81610c08565b915060408401358015158114610d7f57600080fd5b809150509250925092565b600080600080600060808688031215610da257600080fd5b8535610dad81610c08565b94506020860135610dbd81610c08565b9350604086013567ffffffffffffffff811115610dd957600080fd5b610de588828901610c1a565b96999598509660600135949350505050565b60008060408385031215610e0a57600080fd5b8235610e1581610c08565b91506020830135610e2581610c08565b809150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015610e6e57835163ffffffff1683529284019291840191600101610e4c565b50909695505050505050565b600060208284031215610e8c57600080fd5b61071382610cbf565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610ef557610ef5610e95565b5060010190565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b63ffffffff85168152606060208201526000610f65606083018587610efc565b905082604083015295945050505050565b604081526000610f8a604083018587610efc565b9050826020830152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8181038181111561071657610716610e95565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c6343000813000a",
}

var ChannelConfigStoreABI = ChannelConfigStoreMetaData.ABI

var ChannelConfigStoreBin = ChannelConfigStoreMetaData.Bin

func DeployChannelConfigStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChannelConfigStore, error) {
	parsed, err := ChannelConfigStoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChannelConfigStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChannelConfigStore{address: address, abi: *parsed, ChannelConfigStoreCaller: ChannelConfigStoreCaller{contract: contract}, ChannelConfigStoreTransactor: ChannelConfigStoreTransactor{contract: contract}, ChannelConfigStoreFilterer: ChannelConfigStoreFilterer{contract: contract}}, nil
}

type ChannelConfigStore struct {
	address common.Address
	abi     abi.ABI
	ChannelConfigStoreCaller
	ChannelConfigStoreTransactor
	ChannelConfigStoreFilterer
}

type ChannelConfigStoreCaller struct {
	contract *bind.BoundContract
}

type ChannelConfigStoreTransactor struct {
	contract *bind.BoundContract
}

type ChannelConfigStoreFilterer struct {
	contract *bind.BoundContract
}

type ChannelConfigStoreSession struct {
	Contract     *ChannelConfigStore
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ChannelConfigStoreCallerSession struct {
	Contract *ChannelConfigStoreCaller
	CallOpts bind.CallOpts
}

type ChannelConfigStoreTransactorSession struct {
	Contract     *ChannelConfigStoreTransactor
	TransactOpts bind.TransactOpts
}

type ChannelConfigStoreRaw struct {
	Contract *ChannelConfigStore
}

type ChannelConfigStoreCallerRaw struct {
	Contract *ChannelConfigStoreCaller
}

type ChannelConfigStoreTransactorRaw struct {
	Contract *ChannelConfigStoreTransactor
}

func NewChannelConfigStore(address common.Address, backend bind.ContractBackend) (*ChannelConfigStore, error) {
	abi, err := abi.JSON(strings.NewReader(ChannelConfigStoreABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindChannelConfigStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStore{address: address, abi: abi, ChannelConfigStoreCaller: ChannelConfigStoreCaller{contract: contract}, ChannelConfigStoreTransactor: ChannelConfigStoreTransactor{contract: contract}, ChannelConfigStoreFilterer: ChannelConfigStoreFilterer{contract: contract}}, nil
}

func NewChannelConfigStoreCaller(address common.Address, caller bind.ContractCaller) (*ChannelConfigStoreCaller, error) {
	contract, err := bindChannelConfigStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStoreCaller{contract: contract}, nil
}

func NewChannelConfigStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelConfigStoreTransactor, error) {
	contract, err := bindChannelConfigStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStoreTransactor{contract: contract}, nil
}

func NewChannelConfigStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelConfigStoreFilterer, error) {
	contract, err := bindChannelConfigStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStoreFilterer{contract: contract}, nil
}

func bindChannelConfigStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChannelConfigStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ChannelConfigStore *ChannelConfigStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChannelConfigStore.Contract.ChannelConfigStoreCaller.contract.Call(opts, result, method, params...)
}

func (_ChannelConfigStore *ChannelConfigStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.ChannelConfigStoreTransactor.contract.Transfer(opts)
}

func (_ChannelConfigStore *ChannelConfigStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.ChannelConfigStoreTransactor.contract.Transact(opts, method, params...)
}

func (_ChannelConfigStore *ChannelConfigStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChannelConfigStore.Contract.contract.Call(opts, result, method, params...)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.contract.Transfer(opts)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.contract.Transact(opts, method, params...)
}

func (_ChannelConfigStore *ChannelConfigStoreCaller) GetAllowedChannelAdders(opts *bind.CallOpts, donId uint32) ([]uint32, error) {
	var out []interface{}
	err := _ChannelConfigStore.contract.Call(opts, &out, "getAllowedChannelAdders", donId)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

func (_ChannelConfigStore *ChannelConfigStoreSession) GetAllowedChannelAdders(donId uint32) ([]uint32, error) {
	return _ChannelConfigStore.Contract.GetAllowedChannelAdders(&_ChannelConfigStore.CallOpts, donId)
}

func (_ChannelConfigStore *ChannelConfigStoreCallerSession) GetAllowedChannelAdders(donId uint32) ([]uint32, error) {
	return _ChannelConfigStore.Contract.GetAllowedChannelAdders(&_ChannelConfigStore.CallOpts, donId)
}

func (_ChannelConfigStore *ChannelConfigStoreCaller) GetChannelAdderAddress(opts *bind.CallOpts, channelAdderId uint32) (common.Address, error) {
	var out []interface{}
	err := _ChannelConfigStore.contract.Call(opts, &out, "getChannelAdderAddress", channelAdderId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ChannelConfigStore *ChannelConfigStoreSession) GetChannelAdderAddress(channelAdderId uint32) (common.Address, error) {
	return _ChannelConfigStore.Contract.GetChannelAdderAddress(&_ChannelConfigStore.CallOpts, channelAdderId)
}

func (_ChannelConfigStore *ChannelConfigStoreCallerSession) GetChannelAdderAddress(channelAdderId uint32) (common.Address, error) {
	return _ChannelConfigStore.Contract.GetChannelAdderAddress(&_ChannelConfigStore.CallOpts, channelAdderId)
}

func (_ChannelConfigStore *ChannelConfigStoreCaller) IsChannelAdderAllowed(opts *bind.CallOpts, donId uint32, channelAdderId uint32) (bool, error) {
	var out []interface{}
	err := _ChannelConfigStore.contract.Call(opts, &out, "isChannelAdderAllowed", donId, channelAdderId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ChannelConfigStore *ChannelConfigStoreSession) IsChannelAdderAllowed(donId uint32, channelAdderId uint32) (bool, error) {
	return _ChannelConfigStore.Contract.IsChannelAdderAllowed(&_ChannelConfigStore.CallOpts, donId, channelAdderId)
}

func (_ChannelConfigStore *ChannelConfigStoreCallerSession) IsChannelAdderAllowed(donId uint32, channelAdderId uint32) (bool, error) {
	return _ChannelConfigStore.Contract.IsChannelAdderAllowed(&_ChannelConfigStore.CallOpts, donId, channelAdderId)
}

func (_ChannelConfigStore *ChannelConfigStoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChannelConfigStore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ChannelConfigStore *ChannelConfigStoreSession) Owner() (common.Address, error) {
	return _ChannelConfigStore.Contract.Owner(&_ChannelConfigStore.CallOpts)
}

func (_ChannelConfigStore *ChannelConfigStoreCallerSession) Owner() (common.Address, error) {
	return _ChannelConfigStore.Contract.Owner(&_ChannelConfigStore.CallOpts)
}

func (_ChannelConfigStore *ChannelConfigStoreCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ChannelConfigStore.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ChannelConfigStore *ChannelConfigStoreSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ChannelConfigStore.Contract.SupportsInterface(&_ChannelConfigStore.CallOpts, interfaceId)
}

func (_ChannelConfigStore *ChannelConfigStoreCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ChannelConfigStore.Contract.SupportsInterface(&_ChannelConfigStore.CallOpts, interfaceId)
}

func (_ChannelConfigStore *ChannelConfigStoreCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ChannelConfigStore.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ChannelConfigStore *ChannelConfigStoreSession) TypeAndVersion() (string, error) {
	return _ChannelConfigStore.Contract.TypeAndVersion(&_ChannelConfigStore.CallOpts)
}

func (_ChannelConfigStore *ChannelConfigStoreCallerSession) TypeAndVersion() (string, error) {
	return _ChannelConfigStore.Contract.TypeAndVersion(&_ChannelConfigStore.CallOpts)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChannelConfigStore.contract.Transact(opts, "acceptOwnership")
}

func (_ChannelConfigStore *ChannelConfigStoreSession) AcceptOwnership() (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.AcceptOwnership(&_ChannelConfigStore.TransactOpts)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.AcceptOwnership(&_ChannelConfigStore.TransactOpts)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactor) AddChannelDefinitions(opts *bind.TransactOpts, donId uint32, channelAdderId uint32, url string, sha [32]byte) (*types.Transaction, error) {
	return _ChannelConfigStore.contract.Transact(opts, "addChannelDefinitions", donId, channelAdderId, url, sha)
}

func (_ChannelConfigStore *ChannelConfigStoreSession) AddChannelDefinitions(donId uint32, channelAdderId uint32, url string, sha [32]byte) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.AddChannelDefinitions(&_ChannelConfigStore.TransactOpts, donId, channelAdderId, url, sha)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactorSession) AddChannelDefinitions(donId uint32, channelAdderId uint32, url string, sha [32]byte) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.AddChannelDefinitions(&_ChannelConfigStore.TransactOpts, donId, channelAdderId, url, sha)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactor) SetChannelAdder(opts *bind.TransactOpts, donId uint32, channelAdderId uint32, allowed bool) (*types.Transaction, error) {
	return _ChannelConfigStore.contract.Transact(opts, "setChannelAdder", donId, channelAdderId, allowed)
}

func (_ChannelConfigStore *ChannelConfigStoreSession) SetChannelAdder(donId uint32, channelAdderId uint32, allowed bool) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.SetChannelAdder(&_ChannelConfigStore.TransactOpts, donId, channelAdderId, allowed)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactorSession) SetChannelAdder(donId uint32, channelAdderId uint32, allowed bool) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.SetChannelAdder(&_ChannelConfigStore.TransactOpts, donId, channelAdderId, allowed)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactor) SetChannelAdderAddress(opts *bind.TransactOpts, channelAdderId uint32, adderAddress common.Address) (*types.Transaction, error) {
	return _ChannelConfigStore.contract.Transact(opts, "setChannelAdderAddress", channelAdderId, adderAddress)
}

func (_ChannelConfigStore *ChannelConfigStoreSession) SetChannelAdderAddress(channelAdderId uint32, adderAddress common.Address) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.SetChannelAdderAddress(&_ChannelConfigStore.TransactOpts, channelAdderId, adderAddress)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactorSession) SetChannelAdderAddress(channelAdderId uint32, adderAddress common.Address) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.SetChannelAdderAddress(&_ChannelConfigStore.TransactOpts, channelAdderId, adderAddress)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactor) SetChannelDefinitions(opts *bind.TransactOpts, donId uint32, url string, sha [32]byte) (*types.Transaction, error) {
	return _ChannelConfigStore.contract.Transact(opts, "setChannelDefinitions", donId, url, sha)
}

func (_ChannelConfigStore *ChannelConfigStoreSession) SetChannelDefinitions(donId uint32, url string, sha [32]byte) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.SetChannelDefinitions(&_ChannelConfigStore.TransactOpts, donId, url, sha)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactorSession) SetChannelDefinitions(donId uint32, url string, sha [32]byte) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.SetChannelDefinitions(&_ChannelConfigStore.TransactOpts, donId, url, sha)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ChannelConfigStore.contract.Transact(opts, "transferOwnership", to)
}

func (_ChannelConfigStore *ChannelConfigStoreSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.TransferOwnership(&_ChannelConfigStore.TransactOpts, to)
}

func (_ChannelConfigStore *ChannelConfigStoreTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ChannelConfigStore.Contract.TransferOwnership(&_ChannelConfigStore.TransactOpts, to)
}

type ChannelConfigStoreChannelAdderAddressSetIterator struct {
	Event *ChannelConfigStoreChannelAdderAddressSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChannelConfigStoreChannelAdderAddressSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelConfigStoreChannelAdderAddressSet)
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
		it.Event = new(ChannelConfigStoreChannelAdderAddressSet)
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

func (it *ChannelConfigStoreChannelAdderAddressSetIterator) Error() error {
	return it.fail
}

func (it *ChannelConfigStoreChannelAdderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChannelConfigStoreChannelAdderAddressSet struct {
	ChannelAdderId uint32
	AdderAddress   common.Address
	Raw            types.Log
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) FilterChannelAdderAddressSet(opts *bind.FilterOpts, channelAdderId []uint32) (*ChannelConfigStoreChannelAdderAddressSetIterator, error) {

	var channelAdderIdRule []interface{}
	for _, channelAdderIdItem := range channelAdderId {
		channelAdderIdRule = append(channelAdderIdRule, channelAdderIdItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.FilterLogs(opts, "ChannelAdderAddressSet", channelAdderIdRule)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStoreChannelAdderAddressSetIterator{contract: _ChannelConfigStore.contract, event: "ChannelAdderAddressSet", logs: logs, sub: sub}, nil
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) WatchChannelAdderAddressSet(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreChannelAdderAddressSet, channelAdderId []uint32) (event.Subscription, error) {

	var channelAdderIdRule []interface{}
	for _, channelAdderIdItem := range channelAdderId {
		channelAdderIdRule = append(channelAdderIdRule, channelAdderIdItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.WatchLogs(opts, "ChannelAdderAddressSet", channelAdderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChannelConfigStoreChannelAdderAddressSet)
				if err := _ChannelConfigStore.contract.UnpackLog(event, "ChannelAdderAddressSet", log); err != nil {
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

func (_ChannelConfigStore *ChannelConfigStoreFilterer) ParseChannelAdderAddressSet(log types.Log) (*ChannelConfigStoreChannelAdderAddressSet, error) {
	event := new(ChannelConfigStoreChannelAdderAddressSet)
	if err := _ChannelConfigStore.contract.UnpackLog(event, "ChannelAdderAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ChannelConfigStoreChannelAdderSetIterator struct {
	Event *ChannelConfigStoreChannelAdderSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChannelConfigStoreChannelAdderSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelConfigStoreChannelAdderSet)
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
		it.Event = new(ChannelConfigStoreChannelAdderSet)
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

func (it *ChannelConfigStoreChannelAdderSetIterator) Error() error {
	return it.fail
}

func (it *ChannelConfigStoreChannelAdderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChannelConfigStoreChannelAdderSet struct {
	DonId          *big.Int
	ChannelAdderId uint32
	Allowed        bool
	Raw            types.Log
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) FilterChannelAdderSet(opts *bind.FilterOpts, donId []*big.Int, channelAdderId []uint32) (*ChannelConfigStoreChannelAdderSetIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var channelAdderIdRule []interface{}
	for _, channelAdderIdItem := range channelAdderId {
		channelAdderIdRule = append(channelAdderIdRule, channelAdderIdItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.FilterLogs(opts, "ChannelAdderSet", donIdRule, channelAdderIdRule)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStoreChannelAdderSetIterator{contract: _ChannelConfigStore.contract, event: "ChannelAdderSet", logs: logs, sub: sub}, nil
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) WatchChannelAdderSet(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreChannelAdderSet, donId []*big.Int, channelAdderId []uint32) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var channelAdderIdRule []interface{}
	for _, channelAdderIdItem := range channelAdderId {
		channelAdderIdRule = append(channelAdderIdRule, channelAdderIdItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.WatchLogs(opts, "ChannelAdderSet", donIdRule, channelAdderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChannelConfigStoreChannelAdderSet)
				if err := _ChannelConfigStore.contract.UnpackLog(event, "ChannelAdderSet", log); err != nil {
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

func (_ChannelConfigStore *ChannelConfigStoreFilterer) ParseChannelAdderSet(log types.Log) (*ChannelConfigStoreChannelAdderSet, error) {
	event := new(ChannelConfigStoreChannelAdderSet)
	if err := _ChannelConfigStore.contract.UnpackLog(event, "ChannelAdderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ChannelConfigStoreChannelDefinitionAddedIterator struct {
	Event *ChannelConfigStoreChannelDefinitionAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChannelConfigStoreChannelDefinitionAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelConfigStoreChannelDefinitionAdded)
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
		it.Event = new(ChannelConfigStoreChannelDefinitionAdded)
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

func (it *ChannelConfigStoreChannelDefinitionAddedIterator) Error() error {
	return it.fail
}

func (it *ChannelConfigStoreChannelDefinitionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChannelConfigStoreChannelDefinitionAdded struct {
	DonId          *big.Int
	ChannelAdderId uint32
	Url            string
	Sha            [32]byte
	Raw            types.Log
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) FilterChannelDefinitionAdded(opts *bind.FilterOpts, donId []*big.Int, channelAdderId []uint32) (*ChannelConfigStoreChannelDefinitionAddedIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var channelAdderIdRule []interface{}
	for _, channelAdderIdItem := range channelAdderId {
		channelAdderIdRule = append(channelAdderIdRule, channelAdderIdItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.FilterLogs(opts, "ChannelDefinitionAdded", donIdRule, channelAdderIdRule)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStoreChannelDefinitionAddedIterator{contract: _ChannelConfigStore.contract, event: "ChannelDefinitionAdded", logs: logs, sub: sub}, nil
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) WatchChannelDefinitionAdded(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreChannelDefinitionAdded, donId []*big.Int, channelAdderId []uint32) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}
	var channelAdderIdRule []interface{}
	for _, channelAdderIdItem := range channelAdderId {
		channelAdderIdRule = append(channelAdderIdRule, channelAdderIdItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.WatchLogs(opts, "ChannelDefinitionAdded", donIdRule, channelAdderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChannelConfigStoreChannelDefinitionAdded)
				if err := _ChannelConfigStore.contract.UnpackLog(event, "ChannelDefinitionAdded", log); err != nil {
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

func (_ChannelConfigStore *ChannelConfigStoreFilterer) ParseChannelDefinitionAdded(log types.Log) (*ChannelConfigStoreChannelDefinitionAdded, error) {
	event := new(ChannelConfigStoreChannelDefinitionAdded)
	if err := _ChannelConfigStore.contract.UnpackLog(event, "ChannelDefinitionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ChannelConfigStoreNewChannelDefinitionIterator struct {
	Event *ChannelConfigStoreNewChannelDefinition

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChannelConfigStoreNewChannelDefinitionIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelConfigStoreNewChannelDefinition)
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
		it.Event = new(ChannelConfigStoreNewChannelDefinition)
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

func (it *ChannelConfigStoreNewChannelDefinitionIterator) Error() error {
	return it.fail
}

func (it *ChannelConfigStoreNewChannelDefinitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChannelConfigStoreNewChannelDefinition struct {
	DonId   *big.Int
	Version uint32
	Url     string
	Sha     [32]byte
	Raw     types.Log
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) FilterNewChannelDefinition(opts *bind.FilterOpts, donId []*big.Int) (*ChannelConfigStoreNewChannelDefinitionIterator, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.FilterLogs(opts, "NewChannelDefinition", donIdRule)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStoreNewChannelDefinitionIterator{contract: _ChannelConfigStore.contract, event: "NewChannelDefinition", logs: logs, sub: sub}, nil
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) WatchNewChannelDefinition(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreNewChannelDefinition, donId []*big.Int) (event.Subscription, error) {

	var donIdRule []interface{}
	for _, donIdItem := range donId {
		donIdRule = append(donIdRule, donIdItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.WatchLogs(opts, "NewChannelDefinition", donIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChannelConfigStoreNewChannelDefinition)
				if err := _ChannelConfigStore.contract.UnpackLog(event, "NewChannelDefinition", log); err != nil {
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

func (_ChannelConfigStore *ChannelConfigStoreFilterer) ParseNewChannelDefinition(log types.Log) (*ChannelConfigStoreNewChannelDefinition, error) {
	event := new(ChannelConfigStoreNewChannelDefinition)
	if err := _ChannelConfigStore.contract.UnpackLog(event, "NewChannelDefinition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ChannelConfigStoreOwnershipTransferRequestedIterator struct {
	Event *ChannelConfigStoreOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChannelConfigStoreOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelConfigStoreOwnershipTransferRequested)
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
		it.Event = new(ChannelConfigStoreOwnershipTransferRequested)
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

func (it *ChannelConfigStoreOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ChannelConfigStoreOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChannelConfigStoreOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChannelConfigStoreOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStoreOwnershipTransferRequestedIterator{contract: _ChannelConfigStore.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChannelConfigStoreOwnershipTransferRequested)
				if err := _ChannelConfigStore.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ChannelConfigStore *ChannelConfigStoreFilterer) ParseOwnershipTransferRequested(log types.Log) (*ChannelConfigStoreOwnershipTransferRequested, error) {
	event := new(ChannelConfigStoreOwnershipTransferRequested)
	if err := _ChannelConfigStore.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ChannelConfigStoreOwnershipTransferredIterator struct {
	Event *ChannelConfigStoreOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChannelConfigStoreOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelConfigStoreOwnershipTransferred)
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
		it.Event = new(ChannelConfigStoreOwnershipTransferred)
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

func (it *ChannelConfigStoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ChannelConfigStoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChannelConfigStoreOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChannelConfigStoreOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChannelConfigStoreOwnershipTransferredIterator{contract: _ChannelConfigStore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ChannelConfigStore *ChannelConfigStoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ChannelConfigStore.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChannelConfigStoreOwnershipTransferred)
				if err := _ChannelConfigStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ChannelConfigStore *ChannelConfigStoreFilterer) ParseOwnershipTransferred(log types.Log) (*ChannelConfigStoreOwnershipTransferred, error) {
	event := new(ChannelConfigStoreOwnershipTransferred)
	if err := _ChannelConfigStore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ChannelConfigStore *ChannelConfigStore) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ChannelConfigStore.abi.Events["ChannelAdderAddressSet"].ID:
		return _ChannelConfigStore.ParseChannelAdderAddressSet(log)
	case _ChannelConfigStore.abi.Events["ChannelAdderSet"].ID:
		return _ChannelConfigStore.ParseChannelAdderSet(log)
	case _ChannelConfigStore.abi.Events["ChannelDefinitionAdded"].ID:
		return _ChannelConfigStore.ParseChannelDefinitionAdded(log)
	case _ChannelConfigStore.abi.Events["NewChannelDefinition"].ID:
		return _ChannelConfigStore.ParseNewChannelDefinition(log)
	case _ChannelConfigStore.abi.Events["OwnershipTransferRequested"].ID:
		return _ChannelConfigStore.ParseOwnershipTransferRequested(log)
	case _ChannelConfigStore.abi.Events["OwnershipTransferred"].ID:
		return _ChannelConfigStore.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ChannelConfigStoreChannelAdderAddressSet) Topic() common.Hash {
	return common.HexToHash("0xd3c5db69981921b17e39aedfb7c2663d7605477421099c01ad72c980b469f8fe")
}

func (ChannelConfigStoreChannelAdderSet) Topic() common.Hash {
	return common.HexToHash("0x53af2fd88370b4b325f7cbd06f89b37ae7a5ebd9b6ee7cb19f6ce9bc1c7ef02b")
}

func (ChannelConfigStoreChannelDefinitionAdded) Topic() common.Hash {
	return common.HexToHash("0x5e9e3e3cf7859ce74906f92017b202f25c2d278c32ef43735b43d6b35a49764e")
}

func (ChannelConfigStoreNewChannelDefinition) Topic() common.Hash {
	return common.HexToHash("0xe5b641a7879fb491e4e5a35a1ce950f0237b2537ee9b1b1e4fb65e29aff1f5e8")
}

func (ChannelConfigStoreOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ChannelConfigStoreOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_ChannelConfigStore *ChannelConfigStore) Address() common.Address {
	return _ChannelConfigStore.address
}

type ChannelConfigStoreInterface interface {
	GetAllowedChannelAdders(opts *bind.CallOpts, donId uint32) ([]uint32, error)

	GetChannelAdderAddress(opts *bind.CallOpts, channelAdderId uint32) (common.Address, error)

	IsChannelAdderAllowed(opts *bind.CallOpts, donId uint32, channelAdderId uint32) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddChannelDefinitions(opts *bind.TransactOpts, donId uint32, channelAdderId uint32, url string, sha [32]byte) (*types.Transaction, error)

	SetChannelAdder(opts *bind.TransactOpts, donId uint32, channelAdderId uint32, allowed bool) (*types.Transaction, error)

	SetChannelAdderAddress(opts *bind.TransactOpts, channelAdderId uint32, adderAddress common.Address) (*types.Transaction, error)

	SetChannelDefinitions(opts *bind.TransactOpts, donId uint32, url string, sha [32]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterChannelAdderAddressSet(opts *bind.FilterOpts, channelAdderId []uint32) (*ChannelConfigStoreChannelAdderAddressSetIterator, error)

	WatchChannelAdderAddressSet(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreChannelAdderAddressSet, channelAdderId []uint32) (event.Subscription, error)

	ParseChannelAdderAddressSet(log types.Log) (*ChannelConfigStoreChannelAdderAddressSet, error)

	FilterChannelAdderSet(opts *bind.FilterOpts, donId []*big.Int, channelAdderId []uint32) (*ChannelConfigStoreChannelAdderSetIterator, error)

	WatchChannelAdderSet(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreChannelAdderSet, donId []*big.Int, channelAdderId []uint32) (event.Subscription, error)

	ParseChannelAdderSet(log types.Log) (*ChannelConfigStoreChannelAdderSet, error)

	FilterChannelDefinitionAdded(opts *bind.FilterOpts, donId []*big.Int, channelAdderId []uint32) (*ChannelConfigStoreChannelDefinitionAddedIterator, error)

	WatchChannelDefinitionAdded(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreChannelDefinitionAdded, donId []*big.Int, channelAdderId []uint32) (event.Subscription, error)

	ParseChannelDefinitionAdded(log types.Log) (*ChannelConfigStoreChannelDefinitionAdded, error)

	FilterNewChannelDefinition(opts *bind.FilterOpts, donId []*big.Int) (*ChannelConfigStoreNewChannelDefinitionIterator, error)

	WatchNewChannelDefinition(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreNewChannelDefinition, donId []*big.Int) (event.Subscription, error)

	ParseNewChannelDefinition(log types.Log) (*ChannelConfigStoreNewChannelDefinition, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChannelConfigStoreOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ChannelConfigStoreOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChannelConfigStoreOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChannelConfigStoreOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ChannelConfigStoreOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
