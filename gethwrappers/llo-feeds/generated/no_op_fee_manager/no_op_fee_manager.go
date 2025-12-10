// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package no_op_fee_manager

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

type CommonAddressAndWeight struct {
	Addr   common.Address
	Weight uint64
}

type CommonAsset struct {
	AssetAddress common.Address
	Amount       *big.Int
}

var NoOpFeeManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getFeeAndReward\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structCommon.Asset\",\"components\":[{\"name\":\"assetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"reward\",\"type\":\"tuple\",\"internalType\":\"structCommon.Asset\",\"components\":[{\"name\":\"assetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"appliedDiscount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"linkAvailableForPayment\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"payLinkDeficit\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processFee\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"processFeeBulk\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setFeeRecipients\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCommon.AddressAndWeight[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNativeSurcharge\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSubscriberDiscount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSubscriberGlobalDiscount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"RefundFailed\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5061086c806100206000396000f3fe6080604052600436106100bc5760003560e01c8063ce7817d111610074578063e03dab1a1161004e578063e03dab1a146101af578063e389d9a414610248578063f65df9621461026357600080fd5b8063ce7817d114610164578063d09dc33914610185578063dba45fe0146101a157600080fd5b806350538094116100a557806350538094146101185780636c2f1a171461013657806376cf31871461014957600080fd5b806301ffc9a7146100c15780631d4d84a2146100f6575b600080fd5b3480156100cd57600080fd5b506100e16100dc3660046103cb565b61027e565b60405190151581526020015b60405180910390f35b34801561010257600080fd5b5061011661011136600461043d565b505050565b005b34801561012457600080fd5b506101166101333660046104b9565b50565b61011661014436600461051d565b610317565b34801561015557600080fd5b506101166101113660046105cc565b34801561017057600080fd5b5061011661017f36600461060f565b50505050565b34801561019157600080fd5b50604051600081526020016100ed565b61011661014436600461065c565b3480156101bb57600080fd5b506101fd6101ca3660046106e0565b60408051808201909152600080825260208201526040805180820190915260008082526020820152600093509350939050565b60408051845173ffffffffffffffffffffffffffffffffffffffff9081168252602095860151868301528451169181019190915292909101516060830152608082015260a0016100ed565b34801561025457600080fd5b506101166101333660046107c7565b34801561026f57600080fd5b506101166101113660046107e0565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167fdba45fe000000000000000000000000000000000000000000000000000000000148061031157507fffffffff0000000000000000000000000000000000000000000000000000000082167f6c2f1a1700000000000000000000000000000000000000000000000000000000145b92915050565b61032081610327565b5050505050565b34156101335760008173ffffffffffffffffffffffffffffffffffffffff163460405160006040518083038185875af1925050503d8060008114610387576040519150601f19603f3d011682016040523d82523d6000602084013e61038c565b606091505b50509050806103c7576040517ff0c49d4400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6000602082840312156103dd57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461040d57600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461043857600080fd5b919050565b60008060006060848603121561045257600080fd5b61045b84610414565b925061046960208501610414565b9150604084013577ffffffffffffffffffffffffffffffffffffffffffffffff8116811461049657600080fd5b809150509250925092565b803567ffffffffffffffff8116811461043857600080fd5b6000602082840312156104cb57600080fd5b61040d826104a1565b60008083601f8401126104e657600080fd5b50813567ffffffffffffffff8111156104fe57600080fd5b60208301915083602082850101111561051657600080fd5b9250929050565b60008060008060006060868803121561053557600080fd5b853567ffffffffffffffff8082111561054d57600080fd5b818801915088601f83011261056157600080fd5b81358181111561057057600080fd5b8960208260051b850101111561058557600080fd5b6020928301975095509087013590808211156105a057600080fd5b506105ad888289016104d4565b90945092506105c0905060408701610414565b90509295509295909350565b6000806000606084860312156105e157600080fd5b6105ea84610414565b92506105f860208501610414565b9150610606604085016104a1565b90509250925092565b6000806000806080858703121561062557600080fd5b61062e85610414565b93506020850135925061064360408601610414565b9150610651606086016104a1565b905092959194509250565b60008060008060006060868803121561067457600080fd5b853567ffffffffffffffff8082111561068c57600080fd5b61069889838a016104d4565b909750955060208801359150808211156105a057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000606084860312156106f557600080fd5b6106fe84610414565b9250602084013567ffffffffffffffff8082111561071b57600080fd5b818601915086601f83011261072f57600080fd5b813581811115610741576107416106b1565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610787576107876106b1565b816040528281528960208487010111156107a057600080fd5b82602086016020830137600060208483010152809650505050505061060660408501610414565b6000602082840312156107d957600080fd5b5035919050565b6000806000604084860312156107f557600080fd5b83359250602084013567ffffffffffffffff8082111561081457600080fd5b818601915086601f83011261082857600080fd5b81358181111561083757600080fd5b8760208260061b850101111561084c57600080fd5b602083019450809350505050925092509256fea164736f6c6343000813000a",
}

var NoOpFeeManagerABI = NoOpFeeManagerMetaData.ABI

var NoOpFeeManagerBin = NoOpFeeManagerMetaData.Bin

func DeployNoOpFeeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoOpFeeManager, error) {
	parsed, err := NoOpFeeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NoOpFeeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoOpFeeManager{address: address, abi: *parsed, NoOpFeeManagerCaller: NoOpFeeManagerCaller{contract: contract}, NoOpFeeManagerTransactor: NoOpFeeManagerTransactor{contract: contract}, NoOpFeeManagerFilterer: NoOpFeeManagerFilterer{contract: contract}}, nil
}

type NoOpFeeManager struct {
	address common.Address
	abi     abi.ABI
	NoOpFeeManagerCaller
	NoOpFeeManagerTransactor
	NoOpFeeManagerFilterer
}

type NoOpFeeManagerCaller struct {
	contract *bind.BoundContract
}

type NoOpFeeManagerTransactor struct {
	contract *bind.BoundContract
}

type NoOpFeeManagerFilterer struct {
	contract *bind.BoundContract
}

type NoOpFeeManagerSession struct {
	Contract     *NoOpFeeManager
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type NoOpFeeManagerCallerSession struct {
	Contract *NoOpFeeManagerCaller
	CallOpts bind.CallOpts
}

type NoOpFeeManagerTransactorSession struct {
	Contract     *NoOpFeeManagerTransactor
	TransactOpts bind.TransactOpts
}

type NoOpFeeManagerRaw struct {
	Contract *NoOpFeeManager
}

type NoOpFeeManagerCallerRaw struct {
	Contract *NoOpFeeManagerCaller
}

type NoOpFeeManagerTransactorRaw struct {
	Contract *NoOpFeeManagerTransactor
}

func NewNoOpFeeManager(address common.Address, backend bind.ContractBackend) (*NoOpFeeManager, error) {
	abi, err := abi.JSON(strings.NewReader(NoOpFeeManagerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindNoOpFeeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoOpFeeManager{address: address, abi: abi, NoOpFeeManagerCaller: NoOpFeeManagerCaller{contract: contract}, NoOpFeeManagerTransactor: NoOpFeeManagerTransactor{contract: contract}, NoOpFeeManagerFilterer: NoOpFeeManagerFilterer{contract: contract}}, nil
}

func NewNoOpFeeManagerCaller(address common.Address, caller bind.ContractCaller) (*NoOpFeeManagerCaller, error) {
	contract, err := bindNoOpFeeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpFeeManagerCaller{contract: contract}, nil
}

func NewNoOpFeeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NoOpFeeManagerTransactor, error) {
	contract, err := bindNoOpFeeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpFeeManagerTransactor{contract: contract}, nil
}

func NewNoOpFeeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NoOpFeeManagerFilterer, error) {
	contract, err := bindNoOpFeeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoOpFeeManagerFilterer{contract: contract}, nil
}

func bindNoOpFeeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NoOpFeeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_NoOpFeeManager *NoOpFeeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpFeeManager.Contract.NoOpFeeManagerCaller.contract.Call(opts, result, method, params...)
}

func (_NoOpFeeManager *NoOpFeeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.NoOpFeeManagerTransactor.contract.Transfer(opts)
}

func (_NoOpFeeManager *NoOpFeeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.NoOpFeeManagerTransactor.contract.Transact(opts, method, params...)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpFeeManager.Contract.contract.Call(opts, result, method, params...)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.contract.Transfer(opts)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.contract.Transact(opts, method, params...)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) GetFeeAndReward(opts *bind.CallOpts, arg0 common.Address, arg1 []byte, arg2 common.Address) (GetFeeAndReward,

	error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "getFeeAndReward", arg0, arg1, arg2)

	outstruct := new(GetFeeAndReward)
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fee = *abi.ConvertType(out[0], new(CommonAsset)).(*CommonAsset)
	outstruct.Reward = *abi.ConvertType(out[1], new(CommonAsset)).(*CommonAsset)
	outstruct.AppliedDiscount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) GetFeeAndReward(arg0 common.Address, arg1 []byte, arg2 common.Address) (GetFeeAndReward,

	error) {
	return _NoOpFeeManager.Contract.GetFeeAndReward(&_NoOpFeeManager.CallOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) GetFeeAndReward(arg0 common.Address, arg1 []byte, arg2 common.Address) (GetFeeAndReward,

	error) {
	return _NoOpFeeManager.Contract.GetFeeAndReward(&_NoOpFeeManager.CallOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "linkAvailableForPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _NoOpFeeManager.Contract.LinkAvailableForPayment(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) LinkAvailableForPayment() (*big.Int, error) {
	return _NoOpFeeManager.Contract.LinkAvailableForPayment(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NoOpFeeManager.Contract.SupportsInterface(&_NoOpFeeManager.CallOpts, interfaceId)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NoOpFeeManager.Contract.SupportsInterface(&_NoOpFeeManager.CallOpts, interfaceId)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) PayLinkDeficit(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "payLinkDeficit", arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) PayLinkDeficit(arg0 [32]byte) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.PayLinkDeficit(&_NoOpFeeManager.TransactOpts, arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) PayLinkDeficit(arg0 [32]byte) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.PayLinkDeficit(&_NoOpFeeManager.TransactOpts, arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) ProcessFee(opts *bind.TransactOpts, arg0 []byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "processFee", arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) ProcessFee(arg0 []byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.ProcessFee(&_NoOpFeeManager.TransactOpts, arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) ProcessFee(arg0 []byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.ProcessFee(&_NoOpFeeManager.TransactOpts, arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) ProcessFeeBulk(opts *bind.TransactOpts, arg0 [][]byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "processFeeBulk", arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) ProcessFeeBulk(arg0 [][]byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.ProcessFeeBulk(&_NoOpFeeManager.TransactOpts, arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) ProcessFeeBulk(arg0 [][]byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.ProcessFeeBulk(&_NoOpFeeManager.TransactOpts, arg0, arg1, subscriber)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) SetFeeRecipients(opts *bind.TransactOpts, arg0 [32]byte, arg1 []CommonAddressAndWeight) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "setFeeRecipients", arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) SetFeeRecipients(arg0 [32]byte, arg1 []CommonAddressAndWeight) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.SetFeeRecipients(&_NoOpFeeManager.TransactOpts, arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) SetFeeRecipients(arg0 [32]byte, arg1 []CommonAddressAndWeight) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.SetFeeRecipients(&_NoOpFeeManager.TransactOpts, arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) SetNativeSurcharge(opts *bind.TransactOpts, arg0 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "setNativeSurcharge", arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) SetNativeSurcharge(arg0 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.SetNativeSurcharge(&_NoOpFeeManager.TransactOpts, arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) SetNativeSurcharge(arg0 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.SetNativeSurcharge(&_NoOpFeeManager.TransactOpts, arg0)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) UpdateSubscriberDiscount(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address, arg3 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "updateSubscriberDiscount", arg0, arg1, arg2, arg3)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) UpdateSubscriberDiscount(arg0 common.Address, arg1 [32]byte, arg2 common.Address, arg3 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.UpdateSubscriberDiscount(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2, arg3)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) UpdateSubscriberDiscount(arg0 common.Address, arg1 [32]byte, arg2 common.Address, arg3 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.UpdateSubscriberDiscount(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2, arg3)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) UpdateSubscriberGlobalDiscount(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "updateSubscriberGlobalDiscount", arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) UpdateSubscriberGlobalDiscount(arg0 common.Address, arg1 common.Address, arg2 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.UpdateSubscriberGlobalDiscount(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) UpdateSubscriberGlobalDiscount(arg0 common.Address, arg1 common.Address, arg2 uint64) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.UpdateSubscriberGlobalDiscount(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactor) Withdraw(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _NoOpFeeManager.contract.Transact(opts, "withdraw", arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerSession) Withdraw(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.Withdraw(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerTransactorSession) Withdraw(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _NoOpFeeManager.Contract.Withdraw(&_NoOpFeeManager.TransactOpts, arg0, arg1, arg2)
}

type GetFeeAndReward struct {
	Fee             CommonAsset
	Reward          CommonAsset
	AppliedDiscount *big.Int
}

func (_NoOpFeeManager *NoOpFeeManager) Address() common.Address {
	return _NoOpFeeManager.address
}

type NoOpFeeManagerInterface interface {
	GetFeeAndReward(opts *bind.CallOpts, arg0 common.Address, arg1 []byte, arg2 common.Address) (GetFeeAndReward,

		error)

	LinkAvailableForPayment(opts *bind.CallOpts) (*big.Int, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	PayLinkDeficit(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error)

	ProcessFee(opts *bind.TransactOpts, arg0 []byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error)

	ProcessFeeBulk(opts *bind.TransactOpts, arg0 [][]byte, arg1 []byte, subscriber common.Address) (*types.Transaction, error)

	SetFeeRecipients(opts *bind.TransactOpts, arg0 [32]byte, arg1 []CommonAddressAndWeight) (*types.Transaction, error)

	SetNativeSurcharge(opts *bind.TransactOpts, arg0 uint64) (*types.Transaction, error)

	UpdateSubscriberDiscount(opts *bind.TransactOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address, arg3 uint64) (*types.Transaction, error)

	UpdateSubscriberGlobalDiscount(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 uint64) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error)

	Address() common.Address
}
