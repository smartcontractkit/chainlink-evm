// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transparent_upgradeable_proxy

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

var TransparentUpgradeableProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_logic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProxyDeniedAdminAccess\",\"inputs\":[]}]",
	Bin: "0x60a06040526040516111283803806111288339810160408190526100229161039d565b828161002e828261008f565b50508160405161003d9061033a565b6001600160a01b039091168152602001604051809103906000f080158015610069573d6000803e3d6000fd5b506001600160a01b031660805261008761008260805190565b6100ee565b50505061048f565b6100988261015c565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156100e2576100dd82826101db565b505050565b6100ea610252565b5050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61012e600080516020611108833981519152546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a161015981610273565b50565b806001600160a01b03163b60000361019757604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80546001600160a01b0319166001600160a01b039290921691909117905550565b6060600080846001600160a01b0316846040516101f89190610473565b600060405180830381855af49150503d8060008114610233576040519150601f19603f3d011682016040523d82523d6000602084013e610238565b606091505b5090925090506102498583836102b2565b95945050505050565b34156102715760405163b398979f60e01b815260040160405180910390fd5b565b6001600160a01b03811661029d57604051633173bdd160e11b81526000600482015260240161018e565b806000805160206111088339815191526101ba565b6060826102c7576102c282610311565b61030a565b81511580156102de57506001600160a01b0384163b155b1561030757604051639996b31560e01b81526001600160a01b038516600482015260240161018e565b50805b9392505050565b8051156103215780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b61065b80610aad83390190565b80516001600160a01b038116811461035e57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561039457818101518382015260200161037c565b50506000910152565b6000806000606084860312156103b257600080fd5b6103bb84610347565b92506103c960208501610347565b60408501519092506001600160401b038111156103e557600080fd5b8401601f810186136103f657600080fd5b80516001600160401b0381111561040f5761040f610363565b604051601f8201601f19908116603f011681016001600160401b038111828210171561043d5761043d610363565b60405281815282820160200188101561045557600080fd5b610466826020830160208601610379565b8093505050509250925092565b60008251610485818460208701610379565b9190910192915050565b6080516106046104a96000396000601001526106046000f3fe608060405261000c61000e565b005b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1633036100d2576000357fffffffff00000000000000000000000000000000000000000000000000000000167f4f1ef28600000000000000000000000000000000000000000000000000000000146100c8576040517fd2b576ec00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6100d06100da565b565b6100d0610109565b6000806100ea366004818461044d565b8101906100f791906104a6565b915091506101058282610119565b5050565b6100d0610114610181565b6101c6565b610122826101ea565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156101795761017482826102be565b505050565b610105610341565b60006101c17f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5473ffffffffffffffffffffffffffffffffffffffff1690565b905090565b3660008037600080366000845af43d6000803e8080156101e5573d6000f35b3d6000fd5b8073ffffffffffffffffffffffffffffffffffffffff163b600003610258576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60606000808473ffffffffffffffffffffffffffffffffffffffff16846040516102e891906105c8565b600060405180830381855af49150503d8060008114610323576040519150601f19603f3d011682016040523d82523d6000602084013e610328565b606091505b5091509150610338858383610379565b95945050505050565b34156100d0576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608261038e576103898261040b565b610404565b81511580156103b2575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610401576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161024f565b50805b9392505050565b80511561041b5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000808585111561045d57600080fd5b8386111561046a57600080fd5b5050820193919092039150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080604083850312156104b957600080fd5b823573ffffffffffffffffffffffffffffffffffffffff811681146104dd57600080fd5b9150602083013567ffffffffffffffff8111156104f957600080fd5b8301601f8101851361050a57600080fd5b803567ffffffffffffffff81111561052457610524610477565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561059057610590610477565b6040528181528282016020018710156105a857600080fd5b816020840160208301376000602083830101528093505050509250929050565b6000825160005b818110156105e957602081860181015185830152016105cf565b50600092019182525091905056fea164736f6c634300081a000a608060405234801561001057600080fd5b5060405161065b38038061065b83398101604081905261002f916100be565b806001600160a01b03811661005e57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b6100678161006e565b50506100ee565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100d057600080fd5b81516001600160a01b03811681146100e757600080fd5b9392505050565b61055e806100fd6000396000f3fe60806040526004361061005a5760003560e01c80639623609d116100435780639623609d146100b0578063ad3cb1cc146100c3578063f2fde38b1461011957600080fd5b8063715018a61461005f5780638da5cb5b14610076575b600080fd5b34801561006b57600080fd5b50610074610139565b005b34801561008257600080fd5b5060005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100746100be366004610364565b61014d565b3480156100cf57600080fd5b5061010c6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516100a791906104e3565b34801561012557600080fd5b506100746101343660046104fd565b6101e2565b61014161024b565b61014b600061029e565b565b61015561024b565b6040517f4f1ef28600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841690634f1ef2869034906101ab908690869060040161051a565b6000604051808303818588803b1580156101c457600080fd5b505af11580156101d8573d6000803e3d6000fd5b5050505050505050565b6101ea61024b565b73ffffffffffffffffffffffffffffffffffffffff811661023f576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024015b60405180910390fd5b6102488161029e565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461014b576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610236565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b73ffffffffffffffffffffffffffffffffffffffff8116811461024857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060006060848603121561037957600080fd5b833561038481610313565b9250602084013561039481610313565b9150604084013567ffffffffffffffff8111156103b057600080fd5b8401601f810186136103c157600080fd5b803567ffffffffffffffff8111156103db576103db610335565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561044757610447610335565b60405281815282820160200188101561045f57600080fd5b816020840160208301376000602083830101528093505050509250925092565b6000815180845260005b818110156104a557602081850181015186830182015201610489565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006104f6602083018461047f565b9392505050565b60006020828403121561050f57600080fd5b81356104f681610313565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000610549604083018461047f565b94935050505056fea164736f6c634300081a000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103",
}

var TransparentUpgradeableProxyABI = TransparentUpgradeableProxyMetaData.ABI

var TransparentUpgradeableProxyBin = TransparentUpgradeableProxyMetaData.Bin

func DeployTransparentUpgradeableProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _logic common.Address, initialOwner common.Address, _data []byte) (common.Address, *types.Transaction, *TransparentUpgradeableProxy, error) {
	parsed, err := TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransparentUpgradeableProxyBin), backend, _logic, initialOwner, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransparentUpgradeableProxy{address: address, abi: *parsed, TransparentUpgradeableProxyCaller: TransparentUpgradeableProxyCaller{contract: contract}, TransparentUpgradeableProxyTransactor: TransparentUpgradeableProxyTransactor{contract: contract}, TransparentUpgradeableProxyFilterer: TransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

type TransparentUpgradeableProxy struct {
	address common.Address
	abi     abi.ABI
	TransparentUpgradeableProxyCaller
	TransparentUpgradeableProxyTransactor
	TransparentUpgradeableProxyFilterer
}

type TransparentUpgradeableProxyCaller struct {
	contract *bind.BoundContract
}

type TransparentUpgradeableProxyTransactor struct {
	contract *bind.BoundContract
}

type TransparentUpgradeableProxyFilterer struct {
	contract *bind.BoundContract
}

type TransparentUpgradeableProxySession struct {
	Contract     *TransparentUpgradeableProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TransparentUpgradeableProxyCallerSession struct {
	Contract *TransparentUpgradeableProxyCaller
	CallOpts bind.CallOpts
}

type TransparentUpgradeableProxyTransactorSession struct {
	Contract     *TransparentUpgradeableProxyTransactor
	TransactOpts bind.TransactOpts
}

type TransparentUpgradeableProxyRaw struct {
	Contract *TransparentUpgradeableProxy
}

type TransparentUpgradeableProxyCallerRaw struct {
	Contract *TransparentUpgradeableProxyCaller
}

type TransparentUpgradeableProxyTransactorRaw struct {
	Contract *TransparentUpgradeableProxyTransactor
}

func NewTransparentUpgradeableProxy(address common.Address, backend bind.ContractBackend) (*TransparentUpgradeableProxy, error) {
	abi, err := abi.JSON(strings.NewReader(TransparentUpgradeableProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTransparentUpgradeableProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxy{address: address, abi: abi, TransparentUpgradeableProxyCaller: TransparentUpgradeableProxyCaller{contract: contract}, TransparentUpgradeableProxyTransactor: TransparentUpgradeableProxyTransactor{contract: contract}, TransparentUpgradeableProxyFilterer: TransparentUpgradeableProxyFilterer{contract: contract}}, nil
}

func NewTransparentUpgradeableProxyCaller(address common.Address, caller bind.ContractCaller) (*TransparentUpgradeableProxyCaller, error) {
	contract, err := bindTransparentUpgradeableProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyCaller{contract: contract}, nil
}

func NewTransparentUpgradeableProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*TransparentUpgradeableProxyTransactor, error) {
	contract, err := bindTransparentUpgradeableProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyTransactor{contract: contract}, nil
}

func NewTransparentUpgradeableProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*TransparentUpgradeableProxyFilterer, error) {
	contract, err := bindTransparentUpgradeableProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyFilterer{contract: contract}, nil
}

func bindTransparentUpgradeableProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransparentUpgradeableProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyCaller.contract.Call(opts, result, method, params...)
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyTransactor.contract.Transfer(opts)
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.TransparentUpgradeableProxyTransactor.contract.Transact(opts, method, params...)
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransparentUpgradeableProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.contract.Transfer(opts)
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.contract.Transact(opts, method, params...)
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.contract.RawTransact(opts, calldata)
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Fallback(&_TransparentUpgradeableProxy.TransactOpts, calldata)
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TransparentUpgradeableProxy.Contract.Fallback(&_TransparentUpgradeableProxy.TransactOpts, calldata)
}

type TransparentUpgradeableProxyAdminChangedIterator struct {
	Event *TransparentUpgradeableProxyAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TransparentUpgradeableProxyAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyAdminChanged)
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
		it.Event = new(TransparentUpgradeableProxyAdminChanged)
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

func (it *TransparentUpgradeableProxyAdminChangedIterator) Error() error {
	return it.fail
}

func (it *TransparentUpgradeableProxyAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TransparentUpgradeableProxyAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TransparentUpgradeableProxyAdminChangedIterator, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyAdminChangedIterator{contract: _TransparentUpgradeableProxy.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyAdminChanged) (event.Subscription, error) {

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TransparentUpgradeableProxyAdminChanged)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseAdminChanged(log types.Log) (*TransparentUpgradeableProxyAdminChanged, error) {
	event := new(TransparentUpgradeableProxyAdminChanged)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TransparentUpgradeableProxyUpgradedIterator struct {
	Event *TransparentUpgradeableProxyUpgraded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TransparentUpgradeableProxyUpgradedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransparentUpgradeableProxyUpgraded)
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
		it.Event = new(TransparentUpgradeableProxyUpgraded)
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

func (it *TransparentUpgradeableProxyUpgradedIterator) Error() error {
	return it.fail
}

func (it *TransparentUpgradeableProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TransparentUpgradeableProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TransparentUpgradeableProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TransparentUpgradeableProxyUpgradedIterator{contract: _TransparentUpgradeableProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TransparentUpgradeableProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TransparentUpgradeableProxyUpgraded)
				if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxyFilterer) ParseUpgraded(log types.Log) (*TransparentUpgradeableProxyUpgraded, error) {
	event := new(TransparentUpgradeableProxyUpgraded)
	if err := _TransparentUpgradeableProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (TransparentUpgradeableProxyAdminChanged) Topic() common.Hash {
	return common.HexToHash("0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f")
}

func (TransparentUpgradeableProxyUpgraded) Topic() common.Hash {
	return common.HexToHash("0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b")
}

func (_TransparentUpgradeableProxy *TransparentUpgradeableProxy) Address() common.Address {
	return _TransparentUpgradeableProxy.address
}

type TransparentUpgradeableProxyInterface interface {
	Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error)

	FilterAdminChanged(opts *bind.FilterOpts) (*TransparentUpgradeableProxyAdminChangedIterator, error)

	WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyAdminChanged) (event.Subscription, error)

	ParseAdminChanged(log types.Log) (*TransparentUpgradeableProxyAdminChanged, error)

	FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TransparentUpgradeableProxyUpgradedIterator, error)

	WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TransparentUpgradeableProxyUpgraded, implementation []common.Address) (event.Subscription, error)

	ParseUpgraded(log types.Log) (*TransparentUpgradeableProxyUpgraded, error)

	Address() common.Address
}
