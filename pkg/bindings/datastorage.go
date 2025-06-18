// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// Reference imports to suppress errors if they are not otherwise used.
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

// DataStorageUserData is an auto generated low-level Go binding around an user-defined struct.
type DataStorageUserData struct {
	Key   string
	Value string
}

// DatastorageMetaData contains all meta data concerning the Datastorage contract.
var DatastorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"DataNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"AccessLogged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"logAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"readData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structDataStorage.UserData\",\"name\":\"userData\",\"type\":\"tuple\"}],\"name\":\"storeUserData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610aa4806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634ece5b4c1461005157806398458c5d1461006d578063ccf1582714610089578063f5bfa815146100a5575b600080fd5b61006b60048036038101906100669190610605565b6100d5565b005b61008760048036038101906100829190610686565b61019b565b005b6100a3600480360381019061009e91906105b8565b610296565b005b6100bf60048036038101906100ba9190610558565b6102ea565b6040516100cc9190610849565b60405180910390f35b81816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020868660405161012592919061078c565b9081526020016040518091039020919061014092919061042b565b503373ffffffffffffffffffffffffffffffffffffffff167fc95c7d5d3ac582f659cd004afbea77723e1315567b6557f3c059e8eb9586518f8585858560405161018d949392919061080e565b60405180910390a250505050565b8080602001906101ab919061086b565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208380600001906101f9919061086b565b60405161020792919061078c565b9081526020016040518091039020919061022292919061042b565b503373ffffffffffffffffffffffffffffffffffffffff167fc95c7d5d3ac582f659cd004afbea77723e1315567b6557f3c059e8eb9586518f82806000019061026b919061086b565b84806020019061027b919061086b565b60405161028b949392919061080e565b60405180910390a250565b3373ffffffffffffffffffffffffffffffffffffffff167fe2ab1536af9681ad9e5927bca61830526c4cd932e970162eef77328af1fdcfb583836040516102de9291906107ea565b60405180910390a25050565b606060008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020848460405161033c92919061078c565b9081526020016040518091039020805461035590610969565b80601f016020809104026020016040519081016040528092919081815260200182805461038190610969565b80156103ce5780601f106103a3576101008083540402835291602001916103ce565b820191906000526020600020905b8154815290600101906020018083116103b157829003601f168201915b50505050509050600081511415610420578484846040517ff1e50209000000000000000000000000000000000000000000000000000000008152600401610417939291906107a5565b60405180910390fd5b809150509392505050565b82805461043790610969565b90600052602060002090601f01602090048101928261045957600085556104a0565b82601f1061047257803560ff19168380011785556104a0565b828001600101855582156104a0579182015b8281111561049f578235825591602001919060010190610484565b5b5090506104ad91906104b1565b5090565b5b808211156104ca5760008160009055506001016104b2565b5090565b6000813590506104dd81610a57565b92915050565b60008083601f8401126104f9576104f86109cf565b5b8235905067ffffffffffffffff811115610516576105156109ca565b5b602083019150836001820283011115610532576105316109e3565b5b9250929050565b60006040828403121561054f5761054e6109d9565b5b81905092915050565b600080600060408486031215610571576105706109f2565b5b600061057f868287016104ce565b935050602084013567ffffffffffffffff8111156105a05761059f6109ed565b5b6105ac868287016104e3565b92509250509250925092565b600080602083850312156105cf576105ce6109f2565b5b600083013567ffffffffffffffff8111156105ed576105ec6109ed565b5b6105f9858286016104e3565b92509250509250929050565b6000806000806040858703121561061f5761061e6109f2565b5b600085013567ffffffffffffffff81111561063d5761063c6109ed565b5b610649878288016104e3565b9450945050602085013567ffffffffffffffff81111561066c5761066b6109ed565b5b610678878288016104e3565b925092505092959194509250565b60006020828403121561069c5761069b6109f2565b5b600082013567ffffffffffffffff8111156106ba576106b96109ed565b5b6106c684828501610539565b91505092915050565b6106d8816108f5565b82525050565b60006106ea83856108d9565b93506106f7838584610927565b610700836109f7565b840190509392505050565b600061071783856108ea565b9350610724838584610927565b82840190509392505050565b600061073b826108ce565b61074581856108d9565b9350610755818560208601610936565b61075e816109f7565b840191505092915050565b60006107766021836108d9565b915061078182610a08565b604082019050919050565b600061079982848661070b565b91508190509392505050565b60006060820190506107ba60008301866106cf565b81810360208301526107cd8184866106de565b905081810360408301526107e081610769565b9050949350505050565b600060208201905081810360008301526108058184866106de565b90509392505050565b600060408201905081810360008301526108298186886106de565b9050818103602083015261083e8184866106de565b905095945050505050565b600060208201905081810360008301526108638184610730565b905092915050565b60008083356001602003843603038112610888576108876109de565b5b80840192508235915067ffffffffffffffff8211156108aa576108a96109d4565b5b6020830192506001820236038313156108c6576108c56109e8565b5b509250929050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600061090082610907565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b82818337600083830152505050565b60005b83811015610954578082015181840152602081019050610939565b83811115610963576000848401525b50505050565b6000600282049050600182168061098157607f821691505b602082108114156109955761099461099b565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e6f2064617461206173736f63696174656420776974682074686973206b657960008201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b610a60816108f5565b8114610a6b57600080fd5b5056fea2646970667358221220a124db93ed569560f075e2bcc61bd20a4b0585d1c44842b0523350244a1a0f3664736f6c63430008060033",
}

// DatastorageABI is the input ABI used to generate the binding from.
// Deprecated: Use DatastorageMetaData.ABI instead.
var DatastorageABI = DatastorageMetaData.ABI

// DatastorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DatastorageMetaData.Bin instead.
var DatastorageBin = DatastorageMetaData.Bin

// DeployDatastorage deploys a new Ethereum contract, binding an instance of Datastorage to it.
func DeployDatastorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Datastorage, error) {
	parsed, err := DatastorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DatastorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Datastorage{DatastorageCaller: DatastorageCaller{contract: contract}, DatastorageTransactor: DatastorageTransactor{contract: contract}, DatastorageFilterer: DatastorageFilterer{contract: contract}}, nil
}

// Datastorage is an auto generated Go binding around an Ethereum contract.
type Datastorage struct {
	DatastorageCaller     // Read-only binding to the contract
	DatastorageTransactor // Write-only binding to the contract
	DatastorageFilterer   // Log filterer for contract events
}

// DatastorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type DatastorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DatastorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DatastorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatastorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DatastorageSession struct {
	Contract     *Datastorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatastorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DatastorageCallerSession struct {
	Contract *DatastorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DatastorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DatastorageTransactorSession struct {
	Contract     *DatastorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DatastorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type DatastorageRaw struct {
	Contract *Datastorage // Generic contract binding to access the raw methods on
}

// DatastorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DatastorageCallerRaw struct {
	Contract *DatastorageCaller // Generic read-only contract binding to access the raw methods on
}

// DatastorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DatastorageTransactorRaw struct {
	Contract *DatastorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatastorage creates a new instance of Datastorage, bound to a specific deployed contract.
func NewDatastorage(address common.Address, backend bind.ContractBackend) (*Datastorage, error) {
	contract, err := bindDatastorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Datastorage{DatastorageCaller: DatastorageCaller{contract: contract}, DatastorageTransactor: DatastorageTransactor{contract: contract}, DatastorageFilterer: DatastorageFilterer{contract: contract}}, nil
}

// NewDatastorageCaller creates a new read-only instance of Datastorage, bound to a specific deployed contract.
func NewDatastorageCaller(address common.Address, caller bind.ContractCaller) (*DatastorageCaller, error) {
	contract, err := bindDatastorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatastorageCaller{contract: contract}, nil
}

// NewDatastorageTransactor creates a new write-only instance of Datastorage, bound to a specific deployed contract.
func NewDatastorageTransactor(address common.Address, transactor bind.ContractTransactor) (*DatastorageTransactor, error) {
	contract, err := bindDatastorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatastorageTransactor{contract: contract}, nil
}

// NewDatastorageFilterer creates a new log filterer instance of Datastorage, bound to a specific deployed contract.
func NewDatastorageFilterer(address common.Address, filterer bind.ContractFilterer) (*DatastorageFilterer, error) {
	contract, err := bindDatastorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatastorageFilterer{contract: contract}, nil
}

// bindDatastorage binds a generic wrapper to an already deployed contract.
func bindDatastorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DatastorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datastorage *DatastorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Datastorage.Contract.DatastorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datastorage *DatastorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datastorage.Contract.DatastorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datastorage *DatastorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datastorage.Contract.DatastorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datastorage *DatastorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Datastorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datastorage *DatastorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datastorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datastorage *DatastorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datastorage.Contract.contract.Transact(opts, method, params...)
}

// ReadData is a free data retrieval call binding the contract method 0xf5bfa815.
//
// Solidity: function readData(address user, string key) view returns(string)
func (_Datastorage *DatastorageCaller) ReadData(opts *bind.CallOpts, user common.Address, key string) (string, error) {
	var out []interface{}
	err := _Datastorage.contract.Call(opts, &out, "readData", user, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadData is a free data retrieval call binding the contract method 0xf5bfa815.
//
// Solidity: function readData(address user, string key) view returns(string)
func (_Datastorage *DatastorageSession) ReadData(user common.Address, key string) (string, error) {
	return _Datastorage.Contract.ReadData(&_Datastorage.CallOpts, user, key)
}

// ReadData is a free data retrieval call binding the contract method 0xf5bfa815.
//
// Solidity: function readData(address user, string key) view returns(string)
func (_Datastorage *DatastorageCallerSession) ReadData(user common.Address, key string) (string, error) {
	return _Datastorage.Contract.ReadData(&_Datastorage.CallOpts, user, key)
}

// LogAccess is a paid mutator transaction binding the contract method 0xccf15827.
//
// Solidity: function logAccess(string message) returns()
func (_Datastorage *DatastorageTransactor) LogAccess(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _Datastorage.contract.Transact(opts, "logAccess", message)
}

// LogAccess is a paid mutator transaction binding the contract method 0xccf15827.
//
// Solidity: function logAccess(string message) returns()
func (_Datastorage *DatastorageSession) LogAccess(message string) (*types.Transaction, error) {
	return _Datastorage.Contract.LogAccess(&_Datastorage.TransactOpts, message)
}

// LogAccess is a paid mutator transaction binding the contract method 0xccf15827.
//
// Solidity: function logAccess(string message) returns()
func (_Datastorage *DatastorageTransactorSession) LogAccess(message string) (*types.Transaction, error) {
	return _Datastorage.Contract.LogAccess(&_Datastorage.TransactOpts, message)
}

// StoreData is a paid mutator transaction binding the contract method 0x4ece5b4c.
//
// Solidity: function storeData(string key, string value) returns()
func (_Datastorage *DatastorageTransactor) StoreData(opts *bind.TransactOpts, key string, value string) (*types.Transaction, error) {
	return _Datastorage.contract.Transact(opts, "storeData", key, value)
}

// StoreData is a paid mutator transaction binding the contract method 0x4ece5b4c.
//
// Solidity: function storeData(string key, string value) returns()
func (_Datastorage *DatastorageSession) StoreData(key string, value string) (*types.Transaction, error) {
	return _Datastorage.Contract.StoreData(&_Datastorage.TransactOpts, key, value)
}

// StoreData is a paid mutator transaction binding the contract method 0x4ece5b4c.
//
// Solidity: function storeData(string key, string value) returns()
func (_Datastorage *DatastorageTransactorSession) StoreData(key string, value string) (*types.Transaction, error) {
	return _Datastorage.Contract.StoreData(&_Datastorage.TransactOpts, key, value)
}

// StoreUserData is a paid mutator transaction binding the contract method 0x98458c5d.
//
// Solidity: function storeUserData((string,string) userData) returns()
func (_Datastorage *DatastorageTransactor) StoreUserData(opts *bind.TransactOpts, userData DataStorageUserData) (*types.Transaction, error) {
	return _Datastorage.contract.Transact(opts, "storeUserData", userData)
}

// StoreUserData is a paid mutator transaction binding the contract method 0x98458c5d.
//
// Solidity: function storeUserData((string,string) userData) returns()
func (_Datastorage *DatastorageSession) StoreUserData(userData DataStorageUserData) (*types.Transaction, error) {
	return _Datastorage.Contract.StoreUserData(&_Datastorage.TransactOpts, userData)
}

// StoreUserData is a paid mutator transaction binding the contract method 0x98458c5d.
//
// Solidity: function storeUserData((string,string) userData) returns()
func (_Datastorage *DatastorageTransactorSession) StoreUserData(userData DataStorageUserData) (*types.Transaction, error) {
	return _Datastorage.Contract.StoreUserData(&_Datastorage.TransactOpts, userData)
}

// DatastorageAccessLoggedIterator is returned from FilterAccessLogged and is used to iterate over the raw logs and unpacked data for AccessLogged events raised by the Datastorage contract.
type DatastorageAccessLoggedIterator struct {
	Event *DatastorageAccessLogged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DatastorageAccessLoggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatastorageAccessLogged)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DatastorageAccessLogged)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DatastorageAccessLoggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatastorageAccessLoggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatastorageAccessLogged represents a AccessLogged event raised by the Datastorage contract.
type DatastorageAccessLogged struct {
	Caller  common.Address
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccessLogged is a free log retrieval operation binding the contract event 0xe2ab1536af9681ad9e5927bca61830526c4cd932e970162eef77328af1fdcfb5.
//
// Solidity: event AccessLogged(address indexed caller, string message)
func (_Datastorage *DatastorageFilterer) FilterAccessLogged(opts *bind.FilterOpts, caller []common.Address) (*DatastorageAccessLoggedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Datastorage.contract.FilterLogs(opts, "AccessLogged", callerRule)
	if err != nil {
		return nil, err
	}
	return &DatastorageAccessLoggedIterator{contract: _Datastorage.contract, event: "AccessLogged", logs: logs, sub: sub}, nil
}

// WatchAccessLogged is a free log subscription operation binding the contract event 0xe2ab1536af9681ad9e5927bca61830526c4cd932e970162eef77328af1fdcfb5.
//
// Solidity: event AccessLogged(address indexed caller, string message)
func (_Datastorage *DatastorageFilterer) WatchAccessLogged(opts *bind.WatchOpts, sink chan<- *DatastorageAccessLogged, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _Datastorage.contract.WatchLogs(opts, "AccessLogged", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatastorageAccessLogged)
				if err := _Datastorage.contract.UnpackLog(event, "AccessLogged", log); err != nil {
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

// ParseAccessLogged is a log parse operation binding the contract event 0xe2ab1536af9681ad9e5927bca61830526c4cd932e970162eef77328af1fdcfb5.
//
// Solidity: event AccessLogged(address indexed caller, string message)
func (_Datastorage *DatastorageFilterer) ParseAccessLogged(log types.Log) (*DatastorageAccessLogged, error) {
	event := new(DatastorageAccessLogged)
	if err := _Datastorage.contract.UnpackLog(event, "AccessLogged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DatastorageDataStoredIterator is returned from FilterDataStored and is used to iterate over the raw logs and unpacked data for DataStored events raised by the Datastorage contract.
type DatastorageDataStoredIterator struct {
	Event *DatastorageDataStored // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DatastorageDataStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DatastorageDataStored)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DatastorageDataStored)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DatastorageDataStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DatastorageDataStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DatastorageDataStored represents a DataStored event raised by the Datastorage contract.
type DatastorageDataStored struct {
	Sender common.Address
	Key    string
	Value  string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDataStored is a free log retrieval operation binding the contract event 0xc95c7d5d3ac582f659cd004afbea77723e1315567b6557f3c059e8eb9586518f.
//
// Solidity: event DataStored(address indexed sender, string key, string value)
func (_Datastorage *DatastorageFilterer) FilterDataStored(opts *bind.FilterOpts, sender []common.Address) (*DatastorageDataStoredIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Datastorage.contract.FilterLogs(opts, "DataStored", senderRule)
	if err != nil {
		return nil, err
	}
	return &DatastorageDataStoredIterator{contract: _Datastorage.contract, event: "DataStored", logs: logs, sub: sub}, nil
}

// WatchDataStored is a free log subscription operation binding the contract event 0xc95c7d5d3ac582f659cd004afbea77723e1315567b6557f3c059e8eb9586518f.
//
// Solidity: event DataStored(address indexed sender, string key, string value)
func (_Datastorage *DatastorageFilterer) WatchDataStored(opts *bind.WatchOpts, sink chan<- *DatastorageDataStored, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Datastorage.contract.WatchLogs(opts, "DataStored", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DatastorageDataStored)
				if err := _Datastorage.contract.UnpackLog(event, "DataStored", log); err != nil {
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

// ParseDataStored is a log parse operation binding the contract event 0xc95c7d5d3ac582f659cd004afbea77723e1315567b6557f3c059e8eb9586518f.
//
// Solidity: event DataStored(address indexed sender, string key, string value)
func (_Datastorage *DatastorageFilterer) ParseDataStored(log types.Log) (*DatastorageDataStored, error) {
	event := new(DatastorageDataStored)
	if err := _Datastorage.contract.UnpackLog(event, "DataStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
