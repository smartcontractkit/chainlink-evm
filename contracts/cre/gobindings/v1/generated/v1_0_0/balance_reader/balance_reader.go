// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balance_reader

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

var BalanceReaderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNativeBalances\",\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x60808060405234610016576103ae908161001c8239f35b600080fdfe6080604090808252600436101561001557600080fd5b600090813560e01c908163181f5a77146101ca5750634c04bf991461003957600080fd5b346101c7576020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101c3576004359267ffffffffffffffff84116101bf57366023850112156101bf5783600401359361009e61009986610346565b6102d3565b9460248487838152019160051b830101913683116101bb57602401905b82821061018b575050508351916100e06100d761009985610346565b93808552610346565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08185019201368337845b865181101561014b578073ffffffffffffffffffffffffffffffffffffffff6101376001938a61035e565b511631610144828861035e565b520161010c565b509190848483519485948186019282875251809352850193925b82811061017457505050500390f35b835185528695509381019392810192600101610165565b813573ffffffffffffffffffffffffffffffffffffffff811681036101b75781529084019084016100bb565b8680fd5b8580fd5b8280fd5b5080fd5b80fd5b839150346101bf57827ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101bf5781810181811067ffffffffffffffff8211176102a6578252601381526020907f42616c616e636552656164657220312e302e30000000000000000000000000006020820152825193849260208452825192836020860152825b84811061029057505050828201840152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168101030190f35b8181018301518882018801528795508201610254565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f604051930116820182811067ffffffffffffffff82111761031757604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b67ffffffffffffffff81116103175760051b60200190565b80518210156103725760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c6343000818000a",
}

var BalanceReaderABI = BalanceReaderMetaData.ABI

var BalanceReaderBin = BalanceReaderMetaData.Bin

func DeployBalanceReader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BalanceReader, error) {
	parsed, err := BalanceReaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BalanceReaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BalanceReader{address: address, abi: *parsed, BalanceReaderCaller: BalanceReaderCaller{contract: contract}, BalanceReaderTransactor: BalanceReaderTransactor{contract: contract}, BalanceReaderFilterer: BalanceReaderFilterer{contract: contract}}, nil
}

type BalanceReader struct {
	address common.Address
	abi     abi.ABI
	BalanceReaderCaller
	BalanceReaderTransactor
	BalanceReaderFilterer
}

type BalanceReaderCaller struct {
	contract *bind.BoundContract
}

type BalanceReaderTransactor struct {
	contract *bind.BoundContract
}

type BalanceReaderFilterer struct {
	contract *bind.BoundContract
}

type BalanceReaderSession struct {
	Contract     *BalanceReader
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BalanceReaderCallerSession struct {
	Contract *BalanceReaderCaller
	CallOpts bind.CallOpts
}

type BalanceReaderTransactorSession struct {
	Contract     *BalanceReaderTransactor
	TransactOpts bind.TransactOpts
}

type BalanceReaderRaw struct {
	Contract *BalanceReader
}

type BalanceReaderCallerRaw struct {
	Contract *BalanceReaderCaller
}

type BalanceReaderTransactorRaw struct {
	Contract *BalanceReaderTransactor
}

func NewBalanceReader(address common.Address, backend bind.ContractBackend) (*BalanceReader, error) {
	abi, err := abi.JSON(strings.NewReader(BalanceReaderABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBalanceReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceReader{address: address, abi: abi, BalanceReaderCaller: BalanceReaderCaller{contract: contract}, BalanceReaderTransactor: BalanceReaderTransactor{contract: contract}, BalanceReaderFilterer: BalanceReaderFilterer{contract: contract}}, nil
}

func NewBalanceReaderCaller(address common.Address, caller bind.ContractCaller) (*BalanceReaderCaller, error) {
	contract, err := bindBalanceReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceReaderCaller{contract: contract}, nil
}

func NewBalanceReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceReaderTransactor, error) {
	contract, err := bindBalanceReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceReaderTransactor{contract: contract}, nil
}

func NewBalanceReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceReaderFilterer, error) {
	contract, err := bindBalanceReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceReaderFilterer{contract: contract}, nil
}

func bindBalanceReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalanceReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BalanceReader *BalanceReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceReader.Contract.BalanceReaderCaller.contract.Call(opts, result, method, params...)
}

func (_BalanceReader *BalanceReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceReader.Contract.BalanceReaderTransactor.contract.Transfer(opts)
}

func (_BalanceReader *BalanceReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceReader.Contract.BalanceReaderTransactor.contract.Transact(opts, method, params...)
}

func (_BalanceReader *BalanceReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalanceReader.Contract.contract.Call(opts, result, method, params...)
}

func (_BalanceReader *BalanceReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceReader.Contract.contract.Transfer(opts)
}

func (_BalanceReader *BalanceReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceReader.Contract.contract.Transact(opts, method, params...)
}

func (_BalanceReader *BalanceReaderCaller) GetNativeBalances(opts *bind.CallOpts, addresses []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _BalanceReader.contract.Call(opts, &out, "getNativeBalances", addresses)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_BalanceReader *BalanceReaderSession) GetNativeBalances(addresses []common.Address) ([]*big.Int, error) {
	return _BalanceReader.Contract.GetNativeBalances(&_BalanceReader.CallOpts, addresses)
}

func (_BalanceReader *BalanceReaderCallerSession) GetNativeBalances(addresses []common.Address) ([]*big.Int, error) {
	return _BalanceReader.Contract.GetNativeBalances(&_BalanceReader.CallOpts, addresses)
}

func (_BalanceReader *BalanceReaderCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalanceReader.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BalanceReader *BalanceReaderSession) TypeAndVersion() (string, error) {
	return _BalanceReader.Contract.TypeAndVersion(&_BalanceReader.CallOpts)
}

func (_BalanceReader *BalanceReaderCallerSession) TypeAndVersion() (string, error) {
	return _BalanceReader.Contract.TypeAndVersion(&_BalanceReader.CallOpts)
}

func (_BalanceReader *BalanceReader) Address() common.Address {
	return _BalanceReader.address
}

type BalanceReaderInterface interface {
	GetNativeBalances(opts *bind.CallOpts, addresses []common.Address) ([]*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Address() common.Address
}
