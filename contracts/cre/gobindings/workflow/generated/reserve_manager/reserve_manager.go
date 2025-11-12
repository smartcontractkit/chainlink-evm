// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reserve_manager

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

type ReserveManagerUpdateReserves struct {
	TotalMinted  *big.Int
	TotalReserve *big.Int
}

var ReserveManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"lastTotalMinted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastTotalReserve\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReport\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"report\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"RequestReserveUpdate\",\"inputs\":[{\"name\":\"u\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"struct ReserveManager.UpdateReserves\",\"components\":[{\"name\":\"totalMinted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalReserve\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false}]",
	Bin: "0x6080806040523460155761026d908161001b8239f35b600080fdfe608080604052600436101561001357600080fd5b60003560e01c90816301ffc9a71461019457508063624bb9e414610176578063805f21321461006e576384a760091461004b57600080fd5b34610069576000366003190112610069576020600054604051908152f35b600080fd5b346100695760403660031901126100695760043567ffffffffffffffff81116100695761009f903690600401610232565b505060243567ffffffffffffffff8111610069576100c36040913690600401610232565b908092918101031261006957600090604051906040820182811067ffffffffffffffff821117610162576040526020813591828452013560208301918183528455600155600254600019811461014e576001016002556040805192518352905160208301527f5e7ff2d8ad6b6eac88310759fab38a6228ed5bff1f5258edf5302b1094503b3891a180f35b602484634e487b7160e01b81526011600452fd5b602484634e487b7160e01b81526041600452fd5b34610069576000366003190112610069576020600154604051908152f35b3461006957602036600319011261006957600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361006957817f805f21320000000000000000000000000000000000000000000000000000000060209314908115610208575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483610201565b9181601f840112156100695782359167ffffffffffffffff831161006957602083818601950101116100695756fea164736f6c634300081a000a",
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

	Address() common.Address
}
