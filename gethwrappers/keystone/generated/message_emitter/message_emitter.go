// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package message_emitter

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

var MessageEmitterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"emitMessage\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLastMessage\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMessage\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MessageEmitted\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50610944806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063181f5a77146100515780632ac0df26146100a3578063570c537e146100b8578063e0c01bfe146100cb575b600080fd5b61008d6040518060400160405280601581526020017f436f6e7472616374456d697474657220312e302e30000000000000000000000081525081565b60405161009a91906105a8565b60405180910390f35b6100b66100b1366004610615565b6100de565b005b61008d6100c63660046106b0565b6102a7565b61008d6100d93660046106da565b610430565b8061014a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f4d6573736167652063616e6e6f7420626520656d70747900000000000000000060448201526064015b60405180910390fd5b6040805133602080830191909152428284015282518083038401815260609092018352815191810191909120600081815291829052919020805461018d906106fc565b15905061021c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4d65737361676520616c72656164792065786973747320666f7220746865207360448201527f616d652073656e64657220616e6420626c6f636b2074696d657374616d7000006064820152608401610141565b60008181526020819052604090206102358385836107cf565b503360009081526001602052604090206102508385836107cf565b50423373ffffffffffffffffffffffffffffffffffffffff167fc799f359194674b273986b8c03283265390f642b631c04e6526b99d0d8f4c38d858560405161029a9291906108ea565b60405180910390a3505050565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602080830191909152818301849052825180830384018152606092830184528051908201206000818152918290529281208054929392610301906106fc565b905011610390576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f4d65737361676520646f6573206e6f7420657869737420666f7220746865206760448201527f6976656e2073656e64657220616e642074696d657374616d70000000000000006064820152608401610141565b600081815260208190526040902080546103a9906106fc565b80601f01602080910402602001604051908101604052809291908181526020018280546103d5906106fc565b80156104225780601f106103f757610100808354040283529160200191610422565b820191906000526020600020905b81548152906001019060200180831161040557829003601f168201915b505050505091505092915050565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260016020526040812080546060929190610465906106fc565b9050116104f3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4e6f206c617374206d65737361676520666f722074686520676976656e20736560448201527f6e646572000000000000000000000000000000000000000000000000000000006064820152608401610141565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604090208054610523906106fc565b80601f016020809104026020016040519081016040528092919081815260200182805461054f906106fc565b801561059c5780601f106105715761010080835404028352916020019161059c565b820191906000526020600020905b81548152906001019060200180831161057f57829003601f168201915b50505050509050919050565b60006020808352835180602085015260005b818110156105d6578581018301518582016040015282016105ba565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b6000806020838503121561062857600080fd5b823567ffffffffffffffff8082111561064057600080fd5b818501915085601f83011261065457600080fd5b81358181111561066357600080fd5b86602082850101111561067557600080fd5b60209290920196919550909350505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146106ab57600080fd5b919050565b600080604083850312156106c357600080fd5b6106cc83610687565b946020939093013593505050565b6000602082840312156106ec57600080fd5b6106f582610687565b9392505050565b600181811c9082168061071057607f821691505b602082108103610749577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f8211156107ca576000816000526020600020601f850160051c810160208610156107a75750805b601f850160051c820191505b818110156107c6578281556001016107b3565b5050505b505050565b67ffffffffffffffff8311156107e7576107e761074f565b6107fb836107f583546106fc565b8361077e565b6000601f84116001811461084d57600085156108175750838201355b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b1c1916600186901b1783556108e3565b6000838152602090207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0861690835b8281101561089c578685013582556020948501946001909201910161087c565b50868210156108d7577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010191905056fea164736f6c6343000818000a",
}

var MessageEmitterABI = MessageEmitterMetaData.ABI

var MessageEmitterBin = MessageEmitterMetaData.Bin

func DeployMessageEmitter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageEmitter, error) {
	parsed, err := MessageEmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageEmitterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageEmitter{address: address, abi: *parsed, MessageEmitterCaller: MessageEmitterCaller{contract: contract}, MessageEmitterTransactor: MessageEmitterTransactor{contract: contract}, MessageEmitterFilterer: MessageEmitterFilterer{contract: contract}}, nil
}

type MessageEmitter struct {
	address common.Address
	abi     abi.ABI
	MessageEmitterCaller
	MessageEmitterTransactor
	MessageEmitterFilterer
}

type MessageEmitterCaller struct {
	contract *bind.BoundContract
}

type MessageEmitterTransactor struct {
	contract *bind.BoundContract
}

type MessageEmitterFilterer struct {
	contract *bind.BoundContract
}

type MessageEmitterSession struct {
	Contract     *MessageEmitter
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type MessageEmitterCallerSession struct {
	Contract *MessageEmitterCaller
	CallOpts bind.CallOpts
}

type MessageEmitterTransactorSession struct {
	Contract     *MessageEmitterTransactor
	TransactOpts bind.TransactOpts
}

type MessageEmitterRaw struct {
	Contract *MessageEmitter
}

type MessageEmitterCallerRaw struct {
	Contract *MessageEmitterCaller
}

type MessageEmitterTransactorRaw struct {
	Contract *MessageEmitterTransactor
}

func NewMessageEmitter(address common.Address, backend bind.ContractBackend) (*MessageEmitter, error) {
	abi, err := abi.JSON(strings.NewReader(MessageEmitterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindMessageEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageEmitter{address: address, abi: abi, MessageEmitterCaller: MessageEmitterCaller{contract: contract}, MessageEmitterTransactor: MessageEmitterTransactor{contract: contract}, MessageEmitterFilterer: MessageEmitterFilterer{contract: contract}}, nil
}

func NewMessageEmitterCaller(address common.Address, caller bind.ContractCaller) (*MessageEmitterCaller, error) {
	contract, err := bindMessageEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageEmitterCaller{contract: contract}, nil
}

func NewMessageEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageEmitterTransactor, error) {
	contract, err := bindMessageEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageEmitterTransactor{contract: contract}, nil
}

func NewMessageEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageEmitterFilterer, error) {
	contract, err := bindMessageEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageEmitterFilterer{contract: contract}, nil
}

func bindMessageEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageEmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_MessageEmitter *MessageEmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageEmitter.Contract.MessageEmitterCaller.contract.Call(opts, result, method, params...)
}

func (_MessageEmitter *MessageEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageEmitter.Contract.MessageEmitterTransactor.contract.Transfer(opts)
}

func (_MessageEmitter *MessageEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageEmitter.Contract.MessageEmitterTransactor.contract.Transact(opts, method, params...)
}

func (_MessageEmitter *MessageEmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageEmitter.Contract.contract.Call(opts, result, method, params...)
}

func (_MessageEmitter *MessageEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageEmitter.Contract.contract.Transfer(opts)
}

func (_MessageEmitter *MessageEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageEmitter.Contract.contract.Transact(opts, method, params...)
}

func (_MessageEmitter *MessageEmitterCaller) GetLastMessage(opts *bind.CallOpts, emitter common.Address) (string, error) {
	var out []interface{}
	err := _MessageEmitter.contract.Call(opts, &out, "getLastMessage", emitter)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MessageEmitter *MessageEmitterSession) GetLastMessage(emitter common.Address) (string, error) {
	return _MessageEmitter.Contract.GetLastMessage(&_MessageEmitter.CallOpts, emitter)
}

func (_MessageEmitter *MessageEmitterCallerSession) GetLastMessage(emitter common.Address) (string, error) {
	return _MessageEmitter.Contract.GetLastMessage(&_MessageEmitter.CallOpts, emitter)
}

func (_MessageEmitter *MessageEmitterCaller) GetMessage(opts *bind.CallOpts, emitter common.Address, timestamp *big.Int) (string, error) {
	var out []interface{}
	err := _MessageEmitter.contract.Call(opts, &out, "getMessage", emitter, timestamp)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MessageEmitter *MessageEmitterSession) GetMessage(emitter common.Address, timestamp *big.Int) (string, error) {
	return _MessageEmitter.Contract.GetMessage(&_MessageEmitter.CallOpts, emitter, timestamp)
}

func (_MessageEmitter *MessageEmitterCallerSession) GetMessage(emitter common.Address, timestamp *big.Int) (string, error) {
	return _MessageEmitter.Contract.GetMessage(&_MessageEmitter.CallOpts, emitter, timestamp)
}

func (_MessageEmitter *MessageEmitterCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MessageEmitter.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_MessageEmitter *MessageEmitterSession) TypeAndVersion() (string, error) {
	return _MessageEmitter.Contract.TypeAndVersion(&_MessageEmitter.CallOpts)
}

func (_MessageEmitter *MessageEmitterCallerSession) TypeAndVersion() (string, error) {
	return _MessageEmitter.Contract.TypeAndVersion(&_MessageEmitter.CallOpts)
}

func (_MessageEmitter *MessageEmitterTransactor) EmitMessage(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _MessageEmitter.contract.Transact(opts, "emitMessage", message)
}

func (_MessageEmitter *MessageEmitterSession) EmitMessage(message string) (*types.Transaction, error) {
	return _MessageEmitter.Contract.EmitMessage(&_MessageEmitter.TransactOpts, message)
}

func (_MessageEmitter *MessageEmitterTransactorSession) EmitMessage(message string) (*types.Transaction, error) {
	return _MessageEmitter.Contract.EmitMessage(&_MessageEmitter.TransactOpts, message)
}

type MessageEmitterMessageEmittedIterator struct {
	Event *MessageEmitterMessageEmitted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *MessageEmitterMessageEmittedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageEmitterMessageEmitted)
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
		it.Event = new(MessageEmitterMessageEmitted)
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

func (it *MessageEmitterMessageEmittedIterator) Error() error {
	return it.fail
}

func (it *MessageEmitterMessageEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type MessageEmitterMessageEmitted struct {
	Emitter   common.Address
	Timestamp *big.Int
	Message   string
	Raw       types.Log
}

func (_MessageEmitter *MessageEmitterFilterer) FilterMessageEmitted(opts *bind.FilterOpts, emitter []common.Address, timestamp []*big.Int) (*MessageEmitterMessageEmittedIterator, error) {

	var emitterRule []interface{}
	for _, emitterItem := range emitter {
		emitterRule = append(emitterRule, emitterItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _MessageEmitter.contract.FilterLogs(opts, "MessageEmitted", emitterRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return &MessageEmitterMessageEmittedIterator{contract: _MessageEmitter.contract, event: "MessageEmitted", logs: logs, sub: sub}, nil
}

func (_MessageEmitter *MessageEmitterFilterer) WatchMessageEmitted(opts *bind.WatchOpts, sink chan<- *MessageEmitterMessageEmitted, emitter []common.Address, timestamp []*big.Int) (event.Subscription, error) {

	var emitterRule []interface{}
	for _, emitterItem := range emitter {
		emitterRule = append(emitterRule, emitterItem)
	}
	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _MessageEmitter.contract.WatchLogs(opts, "MessageEmitted", emitterRule, timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(MessageEmitterMessageEmitted)
				if err := _MessageEmitter.contract.UnpackLog(event, "MessageEmitted", log); err != nil {
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

func (_MessageEmitter *MessageEmitterFilterer) ParseMessageEmitted(log types.Log) (*MessageEmitterMessageEmitted, error) {
	event := new(MessageEmitterMessageEmitted)
	if err := _MessageEmitter.contract.UnpackLog(event, "MessageEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_MessageEmitter *MessageEmitter) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _MessageEmitter.abi.Events["MessageEmitted"].ID:
		return _MessageEmitter.ParseMessageEmitted(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (MessageEmitterMessageEmitted) Topic() common.Hash {
	return common.HexToHash("0xc799f359194674b273986b8c03283265390f642b631c04e6526b99d0d8f4c38d")
}

func (_MessageEmitter *MessageEmitter) Address() common.Address {
	return _MessageEmitter.address
}

type MessageEmitterInterface interface {
	GetLastMessage(opts *bind.CallOpts, emitter common.Address) (string, error)

	GetMessage(opts *bind.CallOpts, emitter common.Address, timestamp *big.Int) (string, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	EmitMessage(opts *bind.TransactOpts, message string) (*types.Transaction, error)

	FilterMessageEmitted(opts *bind.FilterOpts, emitter []common.Address, timestamp []*big.Int) (*MessageEmitterMessageEmittedIterator, error)

	WatchMessageEmitted(opts *bind.WatchOpts, sink chan<- *MessageEmitterMessageEmitted, emitter []common.Address, timestamp []*big.Int) (event.Subscription, error)

	ParseMessageEmitted(log types.Log) (*MessageEmitterMessageEmitted, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
