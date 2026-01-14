// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shard_config

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

var ShardConfigMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_desiredShardCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"desiredShardCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDesiredShardCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDesiredShardCount\",\"inputs\":[{\"name\":\"_newCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ShardCountUpdated\",\"inputs\":[{\"name\":\"newCount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161069338038061069383398101604081905261002f91610190565b3360008161005057604051639b15e16f60e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b03848116919091179091558116156100805761008081610117565b5050600081116100e15760405162461bcd60e51b815260206004820152602260248201527f536861726420636f756e74206d7573742062652067726561746572207468616e604482015261020360f41b606482015260840160405180910390fd5b600281905560405181907f14786ca9a16162bb91b8495eb0dfc22ade4352450ed6c8bcc2adb933162b877990600090a2506101a9565b336001600160a01b0382160361014057604051636d6c4ee560e11b815260040160405180910390fd5b600080546001600160a01b0319166001600160a01b03838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000602082840312156101a257600080fd5b5051919050565b6104db806101b86000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636bc66efb1161005b5780636bc66efb146100f757806379ba5097146101005780638da5cb5b14610108578063f2fde38b1461013057600080fd5b80630164a01014610082578063181f5a77146100995780632e3cd5aa146100e2575b600080fd5b6002545b6040519081526020015b60405180910390f35b6100d56040518060400160405280601581526020017f5368617264436f6e66696720312e302e302d646576000000000000000000000081525081565b604051610090919061040b565b6100f56100f0366004610478565b610143565b005b61008660025481565b6100f5610212565b60015460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610090565b6100f561013e366004610491565b6102e0565b61014b6102f4565b600081116101df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f536861726420636f756e74206d7573742062652067726561746572207468616e60448201527f2030000000000000000000000000000000000000000000000000000000000000606482015260840160405180910390fd5b600281905560405181907f14786ca9a16162bb91b8495eb0dfc22ade4352450ed6c8bcc2adb933162b877990600090a250565b60005473ffffffffffffffffffffffffffffffffffffffff163314610263576040517f02b543c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560008054909116815560405173ffffffffffffffffffffffffffffffffffffffff909216929183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6102e86102f4565b6102f181610347565b50565b60015473ffffffffffffffffffffffffffffffffffffffff163314610345576040517f2b5c74de00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b3373ffffffffffffffffffffffffffffffffffffffff821603610396576040517fdad89dca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff838116918217835560015460405192939116917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020808352835180602085015260005b818110156104395785810183015185820160400152820161041d565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b60006020828403121561048a57600080fd5b5035919050565b6000602082840312156104a357600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146104c757600080fd5b939250505056fea164736f6c6343000818000a",
}

var ShardConfigABI = ShardConfigMetaData.ABI

var ShardConfigBin = ShardConfigMetaData.Bin

func DeployShardConfig(auth *bind.TransactOpts, backend bind.ContractBackend, _desiredShardCount *big.Int) (common.Address, *types.Transaction, *ShardConfig, error) {
	parsed, err := ShardConfigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ShardConfigBin), backend, _desiredShardCount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShardConfig{address: address, abi: *parsed, ShardConfigCaller: ShardConfigCaller{contract: contract}, ShardConfigTransactor: ShardConfigTransactor{contract: contract}, ShardConfigFilterer: ShardConfigFilterer{contract: contract}}, nil
}

type ShardConfig struct {
	address common.Address
	abi     abi.ABI
	ShardConfigCaller
	ShardConfigTransactor
	ShardConfigFilterer
}

type ShardConfigCaller struct {
	contract *bind.BoundContract
}

type ShardConfigTransactor struct {
	contract *bind.BoundContract
}

type ShardConfigFilterer struct {
	contract *bind.BoundContract
}

type ShardConfigSession struct {
	Contract     *ShardConfig
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ShardConfigCallerSession struct {
	Contract *ShardConfigCaller
	CallOpts bind.CallOpts
}

type ShardConfigTransactorSession struct {
	Contract     *ShardConfigTransactor
	TransactOpts bind.TransactOpts
}

type ShardConfigRaw struct {
	Contract *ShardConfig
}

type ShardConfigCallerRaw struct {
	Contract *ShardConfigCaller
}

type ShardConfigTransactorRaw struct {
	Contract *ShardConfigTransactor
}

func NewShardConfig(address common.Address, backend bind.ContractBackend) (*ShardConfig, error) {
	abi, err := abi.JSON(strings.NewReader(ShardConfigABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindShardConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShardConfig{address: address, abi: abi, ShardConfigCaller: ShardConfigCaller{contract: contract}, ShardConfigTransactor: ShardConfigTransactor{contract: contract}, ShardConfigFilterer: ShardConfigFilterer{contract: contract}}, nil
}

func NewShardConfigCaller(address common.Address, caller bind.ContractCaller) (*ShardConfigCaller, error) {
	contract, err := bindShardConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShardConfigCaller{contract: contract}, nil
}

func NewShardConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ShardConfigTransactor, error) {
	contract, err := bindShardConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShardConfigTransactor{contract: contract}, nil
}

func NewShardConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ShardConfigFilterer, error) {
	contract, err := bindShardConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShardConfigFilterer{contract: contract}, nil
}

func bindShardConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShardConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ShardConfig *ShardConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShardConfig.Contract.ShardConfigCaller.contract.Call(opts, result, method, params...)
}

func (_ShardConfig *ShardConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardConfig.Contract.ShardConfigTransactor.contract.Transfer(opts)
}

func (_ShardConfig *ShardConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShardConfig.Contract.ShardConfigTransactor.contract.Transact(opts, method, params...)
}

func (_ShardConfig *ShardConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShardConfig.Contract.contract.Call(opts, result, method, params...)
}

func (_ShardConfig *ShardConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardConfig.Contract.contract.Transfer(opts)
}

func (_ShardConfig *ShardConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShardConfig.Contract.contract.Transact(opts, method, params...)
}

func (_ShardConfig *ShardConfigCaller) DesiredShardCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardConfig.contract.Call(opts, &out, "desiredShardCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ShardConfig *ShardConfigSession) DesiredShardCount() (*big.Int, error) {
	return _ShardConfig.Contract.DesiredShardCount(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCallerSession) DesiredShardCount() (*big.Int, error) {
	return _ShardConfig.Contract.DesiredShardCount(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCaller) GetDesiredShardCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardConfig.contract.Call(opts, &out, "getDesiredShardCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ShardConfig *ShardConfigSession) GetDesiredShardCount() (*big.Int, error) {
	return _ShardConfig.Contract.GetDesiredShardCount(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCallerSession) GetDesiredShardCount() (*big.Int, error) {
	return _ShardConfig.Contract.GetDesiredShardCount(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShardConfig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ShardConfig *ShardConfigSession) Owner() (common.Address, error) {
	return _ShardConfig.Contract.Owner(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCallerSession) Owner() (common.Address, error) {
	return _ShardConfig.Contract.Owner(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShardConfig.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ShardConfig *ShardConfigSession) TypeAndVersion() (string, error) {
	return _ShardConfig.Contract.TypeAndVersion(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigCallerSession) TypeAndVersion() (string, error) {
	return _ShardConfig.Contract.TypeAndVersion(&_ShardConfig.CallOpts)
}

func (_ShardConfig *ShardConfigTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardConfig.contract.Transact(opts, "acceptOwnership")
}

func (_ShardConfig *ShardConfigSession) AcceptOwnership() (*types.Transaction, error) {
	return _ShardConfig.Contract.AcceptOwnership(&_ShardConfig.TransactOpts)
}

func (_ShardConfig *ShardConfigTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ShardConfig.Contract.AcceptOwnership(&_ShardConfig.TransactOpts)
}

func (_ShardConfig *ShardConfigTransactor) SetDesiredShardCount(opts *bind.TransactOpts, _newCount *big.Int) (*types.Transaction, error) {
	return _ShardConfig.contract.Transact(opts, "setDesiredShardCount", _newCount)
}

func (_ShardConfig *ShardConfigSession) SetDesiredShardCount(_newCount *big.Int) (*types.Transaction, error) {
	return _ShardConfig.Contract.SetDesiredShardCount(&_ShardConfig.TransactOpts, _newCount)
}

func (_ShardConfig *ShardConfigTransactorSession) SetDesiredShardCount(_newCount *big.Int) (*types.Transaction, error) {
	return _ShardConfig.Contract.SetDesiredShardCount(&_ShardConfig.TransactOpts, _newCount)
}

func (_ShardConfig *ShardConfigTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ShardConfig.contract.Transact(opts, "transferOwnership", to)
}

func (_ShardConfig *ShardConfigSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ShardConfig.Contract.TransferOwnership(&_ShardConfig.TransactOpts, to)
}

func (_ShardConfig *ShardConfigTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ShardConfig.Contract.TransferOwnership(&_ShardConfig.TransactOpts, to)
}

type ShardConfigOwnershipTransferRequestedIterator struct {
	Event *ShardConfigOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ShardConfigOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardConfigOwnershipTransferRequested)
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
		it.Event = new(ShardConfigOwnershipTransferRequested)
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

func (it *ShardConfigOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ShardConfigOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ShardConfigOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ShardConfig *ShardConfigFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShardConfigOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShardConfig.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShardConfigOwnershipTransferRequestedIterator{contract: _ShardConfig.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ShardConfig *ShardConfigFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ShardConfigOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShardConfig.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ShardConfigOwnershipTransferRequested)
				if err := _ShardConfig.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ShardConfig *ShardConfigFilterer) ParseOwnershipTransferRequested(log types.Log) (*ShardConfigOwnershipTransferRequested, error) {
	event := new(ShardConfigOwnershipTransferRequested)
	if err := _ShardConfig.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ShardConfigOwnershipTransferredIterator struct {
	Event *ShardConfigOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ShardConfigOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardConfigOwnershipTransferred)
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
		it.Event = new(ShardConfigOwnershipTransferred)
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

func (it *ShardConfigOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ShardConfigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ShardConfigOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ShardConfig *ShardConfigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShardConfigOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShardConfig.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ShardConfigOwnershipTransferredIterator{contract: _ShardConfig.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ShardConfig *ShardConfigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShardConfigOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ShardConfig.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ShardConfigOwnershipTransferred)
				if err := _ShardConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ShardConfig *ShardConfigFilterer) ParseOwnershipTransferred(log types.Log) (*ShardConfigOwnershipTransferred, error) {
	event := new(ShardConfigOwnershipTransferred)
	if err := _ShardConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ShardConfigShardCountUpdatedIterator struct {
	Event *ShardConfigShardCountUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ShardConfigShardCountUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardConfigShardCountUpdated)
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
		it.Event = new(ShardConfigShardCountUpdated)
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

func (it *ShardConfigShardCountUpdatedIterator) Error() error {
	return it.fail
}

func (it *ShardConfigShardCountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ShardConfigShardCountUpdated struct {
	NewCount *big.Int
	Raw      types.Log
}

func (_ShardConfig *ShardConfigFilterer) FilterShardCountUpdated(opts *bind.FilterOpts, newCount []*big.Int) (*ShardConfigShardCountUpdatedIterator, error) {

	var newCountRule []interface{}
	for _, newCountItem := range newCount {
		newCountRule = append(newCountRule, newCountItem)
	}

	logs, sub, err := _ShardConfig.contract.FilterLogs(opts, "ShardCountUpdated", newCountRule)
	if err != nil {
		return nil, err
	}
	return &ShardConfigShardCountUpdatedIterator{contract: _ShardConfig.contract, event: "ShardCountUpdated", logs: logs, sub: sub}, nil
}

func (_ShardConfig *ShardConfigFilterer) WatchShardCountUpdated(opts *bind.WatchOpts, sink chan<- *ShardConfigShardCountUpdated, newCount []*big.Int) (event.Subscription, error) {

	var newCountRule []interface{}
	for _, newCountItem := range newCount {
		newCountRule = append(newCountRule, newCountItem)
	}

	logs, sub, err := _ShardConfig.contract.WatchLogs(opts, "ShardCountUpdated", newCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ShardConfigShardCountUpdated)
				if err := _ShardConfig.contract.UnpackLog(event, "ShardCountUpdated", log); err != nil {
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

func (_ShardConfig *ShardConfigFilterer) ParseShardCountUpdated(log types.Log) (*ShardConfigShardCountUpdated, error) {
	event := new(ShardConfigShardCountUpdated)
	if err := _ShardConfig.contract.UnpackLog(event, "ShardCountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ShardConfigOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ShardConfigOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ShardConfigShardCountUpdated) Topic() common.Hash {
	return common.HexToHash("0x14786ca9a16162bb91b8495eb0dfc22ade4352450ed6c8bcc2adb933162b8779")
}

func (_ShardConfig *ShardConfig) Address() common.Address {
	return _ShardConfig.address
}

type ShardConfigInterface interface {
	DesiredShardCount(opts *bind.CallOpts) (*big.Int, error)

	GetDesiredShardCount(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetDesiredShardCount(opts *bind.TransactOpts, _newCount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShardConfigOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ShardConfigOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ShardConfigOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ShardConfigOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShardConfigOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ShardConfigOwnershipTransferred, error)

	FilterShardCountUpdated(opts *bind.FilterOpts, newCount []*big.Int) (*ShardConfigShardCountUpdatedIterator, error)

	WatchShardCountUpdated(opts *bind.WatchOpts, sink chan<- *ShardConfigShardCountUpdated, newCount []*big.Int) (event.Subscription, error)

	ParseShardCountUpdated(log types.Log) (*ShardConfigShardCountUpdated, error)

	Address() common.Address
}
