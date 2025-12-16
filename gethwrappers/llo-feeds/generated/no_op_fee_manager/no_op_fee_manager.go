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
	ABI: "[{\"type\":\"function\",\"name\":\"getFeeAndReward\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"tuple\",\"internalType\":\"structCommon.Asset\",\"components\":[{\"name\":\"assetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"reward\",\"type\":\"tuple\",\"internalType\":\"structCommon.Asset\",\"components\":[{\"name\":\"assetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"appliedDiscount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"linkAvailableForPayment\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"payLinkDeficit\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processFee\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"processFeeBulk\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"subscriber\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"s_globalDiscounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"s_subscriberDiscounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"setFeeRecipients\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structCommon.AddressAndWeight[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNativeSurcharge\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateSubscriberDiscount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSubscriberGlobalDiscount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint192\",\"internalType\":\"uint192\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"RefundFailed\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b50610a09806100206000396000f3fe6080604052600436106100dd5760003560e01c806387d6d8431161007f578063dba45fe011610059578063dba45fe014610269578063e03dab1a14610277578063e389d9a414610313578063f65df9621461032e57600080fd5b806387d6d84314610209578063ce7817d114610234578063d09dc3391461025557600080fd5b80631d4d84a2116100bb5780631d4d84a21461019b57806350538094146101bd5780636c2f1a17146101db57806376cf3187146101ee57600080fd5b806301ffc9a7146100e2578063181f5a77146101175780631cc7f2d814610163575b600080fd5b3480156100ee57600080fd5b506101026100fd366004610496565b610349565b60405190151581526020015b60405180910390f35b34801561012357600080fd5b50604080518082018252601481527f4e6f4f704665654d616e6167657220302e352e300000000000000000000000006020820152905161010e91906104df565b34801561016f57600080fd5b5061018d61017e366004610574565b670de0b6b3a764000092915050565b60405190815260200161010e565b3480156101a757600080fd5b506101bb6101b63660046105a7565b505050565b005b3480156101c957600080fd5b506101bb6101d8366004610623565b50565b6101bb6101e9366004610687565b6103e2565b3480156101fa57600080fd5b506101bb6101b6366004610736565b34801561021557600080fd5b5061018d610224366004610779565b670de0b6b3a76400009392505050565b34801561024057600080fd5b506101bb61024f3660046107ac565b50505050565b34801561026157600080fd5b50600061018d565b6101bb6101e93660046107f9565b34801561028357600080fd5b506102c861029236600461087d565b505060408051808201825260008082526020808301829052835180850190945281845283015292909150670de0b6b3a764000090565b60408051845173ffffffffffffffffffffffffffffffffffffffff9081168252602095860151868301528451169181019190915292909101516060830152608082015260a00161010e565b34801561031f57600080fd5b506101bb6101d8366004610964565b34801561033a57600080fd5b506101bb6101b636600461097d565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167fdba45fe00000000000000000000000000000000000000000000000000000000014806103dc57507fffffffff0000000000000000000000000000000000000000000000000000000082167f6c2f1a1700000000000000000000000000000000000000000000000000000000145b92915050565b6103eb816103f2565b5050505050565b34156101d85760008173ffffffffffffffffffffffffffffffffffffffff163460405160006040518083038185875af1925050503d8060008114610452576040519150601f19603f3d011682016040523d82523d6000602084013e610457565b606091505b5050905080610492576040517ff0c49d4400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6000602082840312156104a857600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146104d857600080fd5b9392505050565b600060208083528351808285015260005b8181101561050c578581018301518582016040015282016104f0565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461056f57600080fd5b919050565b6000806040838503121561058757600080fd5b6105908361054b565b915061059e6020840161054b565b90509250929050565b6000806000606084860312156105bc57600080fd5b6105c58461054b565b92506105d36020850161054b565b9150604084013577ffffffffffffffffffffffffffffffffffffffffffffffff8116811461060057600080fd5b809150509250925092565b803567ffffffffffffffff8116811461056f57600080fd5b60006020828403121561063557600080fd5b6104d88261060b565b60008083601f84011261065057600080fd5b50813567ffffffffffffffff81111561066857600080fd5b60208301915083602082850101111561068057600080fd5b9250929050565b60008060008060006060868803121561069f57600080fd5b853567ffffffffffffffff808211156106b757600080fd5b818801915088601f8301126106cb57600080fd5b8135818111156106da57600080fd5b8960208260051b85010111156106ef57600080fd5b60209283019750955090870135908082111561070a57600080fd5b506107178882890161063e565b909450925061072a90506040870161054b565b90509295509295909350565b60008060006060848603121561074b57600080fd5b6107548461054b565b92506107626020850161054b565b91506107706040850161060b565b90509250925092565b60008060006060848603121561078e57600080fd5b6107978461054b565b9250602084013591506107706040850161054b565b600080600080608085870312156107c257600080fd5b6107cb8561054b565b9350602085013592506107e06040860161054b565b91506107ee6060860161060b565b905092959194509250565b60008060008060006060868803121561081157600080fd5b853567ffffffffffffffff8082111561082957600080fd5b61083589838a0161063e565b9097509550602088013591508082111561070a57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060006060848603121561089257600080fd5b61089b8461054b565b9250602084013567ffffffffffffffff808211156108b857600080fd5b818601915086601f8301126108cc57600080fd5b8135818111156108de576108de61084e565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156109245761092461084e565b8160405282815289602084870101111561093d57600080fd5b8260208601602083013760006020848301015280965050505050506107706040850161054b565b60006020828403121561097657600080fd5b5035919050565b60008060006040848603121561099257600080fd5b83359250602084013567ffffffffffffffff808211156109b157600080fd5b818601915086601f8301126109c557600080fd5b8135818111156109d457600080fd5b8760208260061b85010111156109e957600080fd5b602083019450809350505050925092509256fea164736f6c6343000813000a",
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

func (_NoOpFeeManager *NoOpFeeManagerCaller) SGlobalDiscounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "s_globalDiscounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) SGlobalDiscounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NoOpFeeManager.Contract.SGlobalDiscounts(&_NoOpFeeManager.CallOpts, arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) SGlobalDiscounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _NoOpFeeManager.Contract.SGlobalDiscounts(&_NoOpFeeManager.CallOpts, arg0, arg1)
}

func (_NoOpFeeManager *NoOpFeeManagerCaller) SSubscriberDiscounts(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "s_subscriberDiscounts", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) SSubscriberDiscounts(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _NoOpFeeManager.Contract.SSubscriberDiscounts(&_NoOpFeeManager.CallOpts, arg0, arg1, arg2)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) SSubscriberDiscounts(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _NoOpFeeManager.Contract.SSubscriberDiscounts(&_NoOpFeeManager.CallOpts, arg0, arg1, arg2)
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

func (_NoOpFeeManager *NoOpFeeManagerCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NoOpFeeManager.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_NoOpFeeManager *NoOpFeeManagerSession) TypeAndVersion() (string, error) {
	return _NoOpFeeManager.Contract.TypeAndVersion(&_NoOpFeeManager.CallOpts)
}

func (_NoOpFeeManager *NoOpFeeManagerCallerSession) TypeAndVersion() (string, error) {
	return _NoOpFeeManager.Contract.TypeAndVersion(&_NoOpFeeManager.CallOpts)
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

	SGlobalDiscounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error)

	SSubscriberDiscounts(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

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
