// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reserve_manager

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

type ReserveManagerUpdateReserves struct {
	TotalMinted  *big.Int
	TotalReserve *big.Int
}

var ReserveManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"lastTotalMinted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastTotalReserve\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"RequestReserveUpdate\",\"inputs\":[{\"name\":\"u\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structReserveManager.UpdateReserves\",\"components\":[{\"name\":\"totalMinted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalReserve\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f80fd5b5061038b8061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806301ffc9a71461004e578063624bb9e414610076578063805f21321461008d57806384a76009146100a2575b5f80fd5b61006161005c3660046101b7565b6100aa565b60405190151581526020015b60405180910390f35b61007f60015481565b60405190815260200161006d565b6100a061009b366004610242565b610142565b005b61007f5f5481565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f805f213200000000000000000000000000000000000000000000000000000000148061013c57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b5f61014f828401846102ae565b80515f9081556020820151600155600280549293509061016e83610322565b90915550506040805182518152602080840151908201527f5e7ff2d8ad6b6eac88310759fab38a6228ed5bff1f5258edf5302b1094503b38910160405180910390a15050505050565b5f602082840312156101c7575f80fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146101f6575f80fd5b9392505050565b5f8083601f84011261020d575f80fd5b50813567ffffffffffffffff811115610224575f80fd5b60208301915083602082850101111561023b575f80fd5b9250929050565b5f805f8060408587031215610255575f80fd5b843567ffffffffffffffff81111561026b575f80fd5b610277878288016101fd565b909550935050602085013567ffffffffffffffff811115610296575f80fd5b6102a2878288016101fd565b95989497509550505050565b5f60408284031280156102bf575f80fd5b506040805190810167ffffffffffffffff81118282101715610308577f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604052823581526020928301359281019290925250919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610377577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b506001019056fea164736f6c634300081a000a",
}

var ReserveManagerABI = ReserveManagerMetaData.ABI

var ReserveManagerBin = ReserveManagerMetaData.Bin

func DeployReserveManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReserveManager, error) {
	parsed, err := ReserveManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReserveManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReserveManager{address: address, abi: *parsed, ReserveManagerCaller: ReserveManagerCaller{contract: contract}, ReserveManagerTransactor: ReserveManagerTransactor{contract: contract}, ReserveManagerFilterer: ReserveManagerFilterer{contract: contract}}, nil
}

type ReserveManager struct {
	address common.Address
	abi     abi.ABI
	ReserveManagerCaller
	ReserveManagerTransactor
	ReserveManagerFilterer
}

type ReserveManagerCaller struct {
	contract *bind.BoundContract
}

type ReserveManagerTransactor struct {
	contract *bind.BoundContract
}

type ReserveManagerFilterer struct {
	contract *bind.BoundContract
}

type ReserveManagerSession struct {
	Contract     *ReserveManager
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ReserveManagerCallerSession struct {
	Contract *ReserveManagerCaller
	CallOpts bind.CallOpts
}

type ReserveManagerTransactorSession struct {
	Contract     *ReserveManagerTransactor
	TransactOpts bind.TransactOpts
}

type ReserveManagerRaw struct {
	Contract *ReserveManager
}

type ReserveManagerCallerRaw struct {
	Contract *ReserveManagerCaller
}

type ReserveManagerTransactorRaw struct {
	Contract *ReserveManagerTransactor
}

func NewReserveManager(address common.Address, backend bind.ContractBackend) (*ReserveManager, error) {
	abi, err := abi.JSON(strings.NewReader(ReserveManagerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindReserveManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReserveManager{address: address, abi: abi, ReserveManagerCaller: ReserveManagerCaller{contract: contract}, ReserveManagerTransactor: ReserveManagerTransactor{contract: contract}, ReserveManagerFilterer: ReserveManagerFilterer{contract: contract}}, nil
}

func NewReserveManagerCaller(address common.Address, caller bind.ContractCaller) (*ReserveManagerCaller, error) {
	contract, err := bindReserveManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveManagerCaller{contract: contract}, nil
}

func NewReserveManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveManagerTransactor, error) {
	contract, err := bindReserveManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveManagerTransactor{contract: contract}, nil
}

func NewReserveManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ReserveManagerFilterer, error) {
	contract, err := bindReserveManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReserveManagerFilterer{contract: contract}, nil
}

func bindReserveManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReserveManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ReserveManager *ReserveManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReserveManager.Contract.ReserveManagerCaller.contract.Call(opts, result, method, params...)
}

func (_ReserveManager *ReserveManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveManager.Contract.ReserveManagerTransactor.contract.Transfer(opts)
}

func (_ReserveManager *ReserveManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveManager.Contract.ReserveManagerTransactor.contract.Transact(opts, method, params...)
}

func (_ReserveManager *ReserveManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReserveManager.Contract.contract.Call(opts, result, method, params...)
}

func (_ReserveManager *ReserveManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveManager.Contract.contract.Transfer(opts)
}

func (_ReserveManager *ReserveManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveManager.Contract.contract.Transact(opts, method, params...)
}

func (_ReserveManager *ReserveManagerCaller) LastTotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReserveManager.contract.Call(opts, &out, "lastTotalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ReserveManager *ReserveManagerSession) LastTotalMinted() (*big.Int, error) {
	return _ReserveManager.Contract.LastTotalMinted(&_ReserveManager.CallOpts)
}

func (_ReserveManager *ReserveManagerCallerSession) LastTotalMinted() (*big.Int, error) {
	return _ReserveManager.Contract.LastTotalMinted(&_ReserveManager.CallOpts)
}

func (_ReserveManager *ReserveManagerCaller) LastTotalReserve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReserveManager.contract.Call(opts, &out, "lastTotalReserve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ReserveManager *ReserveManagerSession) LastTotalReserve() (*big.Int, error) {
	return _ReserveManager.Contract.LastTotalReserve(&_ReserveManager.CallOpts)
}

func (_ReserveManager *ReserveManagerCallerSession) LastTotalReserve() (*big.Int, error) {
	return _ReserveManager.Contract.LastTotalReserve(&_ReserveManager.CallOpts)
}

func (_ReserveManager *ReserveManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ReserveManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ReserveManager *ReserveManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ReserveManager.Contract.SupportsInterface(&_ReserveManager.CallOpts, interfaceId)
}

func (_ReserveManager *ReserveManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ReserveManager.Contract.SupportsInterface(&_ReserveManager.CallOpts, interfaceId)
}

func (_ReserveManager *ReserveManagerTransactor) OnReport(opts *bind.TransactOpts, arg0 []byte, report []byte) (*types.Transaction, error) {
	return _ReserveManager.contract.Transact(opts, "onReport", arg0, report)
}

func (_ReserveManager *ReserveManagerSession) OnReport(arg0 []byte, report []byte) (*types.Transaction, error) {
	return _ReserveManager.Contract.OnReport(&_ReserveManager.TransactOpts, arg0, report)
}

func (_ReserveManager *ReserveManagerTransactorSession) OnReport(arg0 []byte, report []byte) (*types.Transaction, error) {
	return _ReserveManager.Contract.OnReport(&_ReserveManager.TransactOpts, arg0, report)
}

type ReserveManagerRequestReserveUpdateIterator struct {
	Event *ReserveManagerRequestReserveUpdate

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ReserveManagerRequestReserveUpdateIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveManagerRequestReserveUpdate)
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
		it.Event = new(ReserveManagerRequestReserveUpdate)
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

func (it *ReserveManagerRequestReserveUpdateIterator) Error() error {
	return it.fail
}

func (it *ReserveManagerRequestReserveUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ReserveManagerRequestReserveUpdate struct {
	U   ReserveManagerUpdateReserves
	Raw types.Log
}

func (_ReserveManager *ReserveManagerFilterer) FilterRequestReserveUpdate(opts *bind.FilterOpts) (*ReserveManagerRequestReserveUpdateIterator, error) {

	logs, sub, err := _ReserveManager.contract.FilterLogs(opts, "RequestReserveUpdate")
	if err != nil {
		return nil, err
	}
	return &ReserveManagerRequestReserveUpdateIterator{contract: _ReserveManager.contract, event: "RequestReserveUpdate", logs: logs, sub: sub}, nil
}

func (_ReserveManager *ReserveManagerFilterer) WatchRequestReserveUpdate(opts *bind.WatchOpts, sink chan<- *ReserveManagerRequestReserveUpdate) (event.Subscription, error) {

	logs, sub, err := _ReserveManager.contract.WatchLogs(opts, "RequestReserveUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ReserveManagerRequestReserveUpdate)
				if err := _ReserveManager.contract.UnpackLog(event, "RequestReserveUpdate", log); err != nil {
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

func (_ReserveManager *ReserveManagerFilterer) ParseRequestReserveUpdate(log types.Log) (*ReserveManagerRequestReserveUpdate, error) {
	event := new(ReserveManagerRequestReserveUpdate)
	if err := _ReserveManager.contract.UnpackLog(event, "RequestReserveUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ReserveManager *ReserveManager) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ReserveManager.abi.Events["RequestReserveUpdate"].ID:
		return _ReserveManager.ParseRequestReserveUpdate(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ReserveManagerRequestReserveUpdate) Topic() common.Hash {
	return common.HexToHash("0x5e7ff2d8ad6b6eac88310759fab38a6228ed5bff1f5258edf5302b1094503b38")
}

func (_ReserveManager *ReserveManager) Address() common.Address {
	return _ReserveManager.address
}

type ReserveManagerInterface interface {
	LastTotalMinted(opts *bind.CallOpts) (*big.Int, error)

	LastTotalReserve(opts *bind.CallOpts) (*big.Int, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	OnReport(opts *bind.TransactOpts, arg0 []byte, report []byte) (*types.Transaction, error)

	FilterRequestReserveUpdate(opts *bind.FilterOpts) (*ReserveManagerRequestReserveUpdateIterator, error)

	WatchRequestReserveUpdate(opts *bind.WatchOpts, sink chan<- *ReserveManagerRequestReserveUpdate) (event.Subscription, error)

	ParseRequestReserveUpdate(log types.Log) (*ReserveManagerRequestReserveUpdate, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
