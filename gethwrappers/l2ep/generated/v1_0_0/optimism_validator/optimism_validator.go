// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimism_validator

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

var OptimismValidatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"l1CrossDomainMessengerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l2UptimeFeedAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"L1_CROSS_DOMAIN_MESSENGER_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L2_UPTIME_FEED_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validate\",\"inputs\":[{\"name\":\"previousRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"previousAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"currentRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigUpdated\",\"inputs\":[{\"name\":\"l1CrossDomainMessengerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"l2UptimeFeedAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasLimitUpdated\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatedStatus\",\"inputs\":[{\"name\":\"previousRoundId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"previousAnswer\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"currentRoundId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currentAnswer\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"L1CrossDomainMessengerAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2UptimeFeedAddrZero\",\"inputs\":[]}]",
	Bin: "0x60c060405234801562000010575f80fd5b50604051620010f6380380620010f6833981016040819052620000339162000266565b82828233805f816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620001a0565b50506001805460ff60a01b1916600160a01b179055506001600160a01b038316620000fc57604051639c9087c160e01b815260040160405180910390fd5b6001600160a01b0382166200012457604051632ef8090560e01b815260040160405180910390fd5b6001600160a01b03838116608081905290831660a08190526003805463ffffffff191663ffffffff8516908117909155604080519384526020840192909252908201527f50ac1a0ab48727faf611b51ad202bb80d8cc334ca8e64154c3d430fbb23040cb9060600160405180910390a1505050505050620002b8565b336001600160a01b03821603620001fa5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b038116811462000261575f80fd5b919050565b5f805f6060848603121562000279575f80fd5b62000284846200024a565b925062000294602085016200024a565b9150604084015163ffffffff81168114620002ad575f80fd5b809150509250925092565b60805160a051610e0e620002e85f395f8181610119015261097501525f818161031c01526109480152610e0e5ff3fe6080604052600436106100e7575f3560e01c80638038e4a111610087578063beed9b5111610057578063beed9b51146102bb578063dc7f0124146102da578063eda066f71461030b578063f2fde38b1461033e575f80fd5b80638038e4a1146102405780638823da6c146102545780638da5cb5b14610273578063a118f2491461029c575f80fd5b80631a93d1c3116100c25780631a93d1c3146101ba57806352d84c62146101de5780636b14daf8146101fd57806379ba50971461022c575f80fd5b80630a756983146100f2578063122555ff14610108578063181f5a7714610165575f80fd5b366100ee57005b5f80fd5b3480156100fd575f80fd5b5061010661035d565b005b348015610113575f80fd5b5061013b7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b348015610170575f80fd5b506101ad6040518060400160405280601b81526020017f4f7074696d69736d56616c696461746f7220312e312e302d646576000000000081525081565b60405161015c9190610c11565b3480156101c5575f80fd5b5060035460405163ffffffff909116815260200161015c565b3480156101e9575f80fd5b506101066101f8366004610c23565b6103db565b348015610208575f80fd5b5061021c610217366004610c9b565b61044d565b604051901515815260200161015c565b348015610237575f80fd5b506101066104a2565b34801561024b575f80fd5b506101066105a3565b34801561025f575f80fd5b5061010661026e366004610d75565b610636565b34801561027e575f80fd5b505f5473ffffffffffffffffffffffffffffffffffffffff1661013b565b3480156102a7575f80fd5b506101066102b6366004610d75565b6106e9565b3480156102c6575f80fd5b5061021c6102d5366004610d8e565b61079b565b3480156102e5575f80fd5b5060015461021c9074010000000000000000000000000000000000000000900460ff1681565b348015610316575f80fd5b5061013b7f000000000000000000000000000000000000000000000000000000000000000081565b348015610349575f80fd5b50610106610358366004610d75565b610a2b565b610365610a3c565b60015474010000000000000000000000000000000000000000900460ff16156103d957600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638905f90a15b565b6103e3610a3c565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff83169081179091556040519081527f501c51f34bf3d8589156f7fd37055e69375b40034d9f99365cb3290b5d983c91906020015b60405180910390a150565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526002602052604081205460ff168061049b575060015474010000000000000000000000000000000000000000900460ff16155b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610528576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6105ab610a3c565b60015474010000000000000000000000000000000000000000900460ff166103d957600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517faebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480905f90a1565b61063e610a3c565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff16156106e65773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527f3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d19101610442565b50565b6106f1610a3c565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff166106e65773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527f87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db49101610442565b5f6107db335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061044d92505050565b610841576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f206163636573730000000000000000000000000000000000000000000000604482015260640161051f565b60405160018314602482018190524267ffffffffffffffff811660448401527fed8378f500000000000000000000000000000000000000000000000000000000925f908490606401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925260035491517f3dbb202b00000000000000000000000000000000000000000000000000000000815290925073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001691633dbb202b916109a6917f000000000000000000000000000000000000000000000000000000000000000091869163ffffffff1690600401610dbd565b5f604051808303815f87803b1580156109bd575f80fd5b505af11580156109cf573d5f803e3d5ffd5b5050604080518c8152602081018c90529081018a9052606081018990527fdf570e5bda6b7c0c0415ceeafd5f9321908b69023cae41e030a6df91cad90b039250608001905060405180910390a150600198975050505050505050565b610a33610a3c565b6106e681610abc565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146103d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161051f565b3373ffffffffffffffffffffffffffffffffffffffff821603610b3b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161051f565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b5f81518084525f5b81811015610bd457602081850181015186830182015201610bb8565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61049b6020830184610bb0565b5f60208284031215610c33575f80fd5b813563ffffffff8116811461049b575f80fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114610c69575f80fd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215610cac575f80fd5b610cb583610c46565b9150602083013567ffffffffffffffff80821115610cd1575f80fd5b818501915085601f830112610ce4575f80fd5b813581811115610cf657610cf6610c6e565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610d3c57610d3c610c6e565b81604052828152886020848701011115610d54575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215610d85575f80fd5b61049b82610c46565b5f805f8060808587031215610da1575f80fd5b5050823594602084013594506040840135936060013592509050565b73ffffffffffffffffffffffffffffffffffffffff84168152606060208201525f610deb6060830185610bb0565b905063ffffffff8316604083015294935050505056fea164736f6c6343000818000a",
}

var OptimismValidatorABI = OptimismValidatorMetaData.ABI

var OptimismValidatorBin = OptimismValidatorMetaData.Bin

func DeployOptimismValidator(auth *bind.TransactOpts, backend bind.ContractBackend, l1CrossDomainMessengerAddress common.Address, l2UptimeFeedAddr common.Address, gasLimit uint32) (common.Address, *types.Transaction, *OptimismValidator, error) {
	parsed, err := OptimismValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismValidatorBin), backend, l1CrossDomainMessengerAddress, l2UptimeFeedAddr, gasLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismValidator{address: address, abi: *parsed, OptimismValidatorCaller: OptimismValidatorCaller{contract: contract}, OptimismValidatorTransactor: OptimismValidatorTransactor{contract: contract}, OptimismValidatorFilterer: OptimismValidatorFilterer{contract: contract}}, nil
}

type OptimismValidator struct {
	address common.Address
	abi     abi.ABI
	OptimismValidatorCaller
	OptimismValidatorTransactor
	OptimismValidatorFilterer
}

type OptimismValidatorCaller struct {
	contract *bind.BoundContract
}

type OptimismValidatorTransactor struct {
	contract *bind.BoundContract
}

type OptimismValidatorFilterer struct {
	contract *bind.BoundContract
}

type OptimismValidatorSession struct {
	Contract     *OptimismValidator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OptimismValidatorCallerSession struct {
	Contract *OptimismValidatorCaller
	CallOpts bind.CallOpts
}

type OptimismValidatorTransactorSession struct {
	Contract     *OptimismValidatorTransactor
	TransactOpts bind.TransactOpts
}

type OptimismValidatorRaw struct {
	Contract *OptimismValidator
}

type OptimismValidatorCallerRaw struct {
	Contract *OptimismValidatorCaller
}

type OptimismValidatorTransactorRaw struct {
	Contract *OptimismValidatorTransactor
}

func NewOptimismValidator(address common.Address, backend bind.ContractBackend) (*OptimismValidator, error) {
	abi, err := abi.JSON(strings.NewReader(OptimismValidatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOptimismValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismValidator{address: address, abi: abi, OptimismValidatorCaller: OptimismValidatorCaller{contract: contract}, OptimismValidatorTransactor: OptimismValidatorTransactor{contract: contract}, OptimismValidatorFilterer: OptimismValidatorFilterer{contract: contract}}, nil
}

func NewOptimismValidatorCaller(address common.Address, caller bind.ContractCaller) (*OptimismValidatorCaller, error) {
	contract, err := bindOptimismValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorCaller{contract: contract}, nil
}

func NewOptimismValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismValidatorTransactor, error) {
	contract, err := bindOptimismValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorTransactor{contract: contract}, nil
}

func NewOptimismValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismValidatorFilterer, error) {
	contract, err := bindOptimismValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorFilterer{contract: contract}, nil
}

func bindOptimismValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OptimismValidator *OptimismValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismValidator.Contract.OptimismValidatorCaller.contract.Call(opts, result, method, params...)
}

func (_OptimismValidator *OptimismValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismValidator.Contract.OptimismValidatorTransactor.contract.Transfer(opts)
}

func (_OptimismValidator *OptimismValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismValidator.Contract.OptimismValidatorTransactor.contract.Transact(opts, method, params...)
}

func (_OptimismValidator *OptimismValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismValidator.Contract.contract.Call(opts, result, method, params...)
}

func (_OptimismValidator *OptimismValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismValidator.Contract.contract.Transfer(opts)
}

func (_OptimismValidator *OptimismValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismValidator.Contract.contract.Transact(opts, method, params...)
}

func (_OptimismValidator *OptimismValidatorCaller) L1CROSSDOMAINMESSENGERADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismValidator.contract.Call(opts, &out, "L1_CROSS_DOMAIN_MESSENGER_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismValidator *OptimismValidatorSession) L1CROSSDOMAINMESSENGERADDRESS() (common.Address, error) {
	return _OptimismValidator.Contract.L1CROSSDOMAINMESSENGERADDRESS(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCallerSession) L1CROSSDOMAINMESSENGERADDRESS() (common.Address, error) {
	return _OptimismValidator.Contract.L1CROSSDOMAINMESSENGERADDRESS(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCaller) L2UPTIMEFEEDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismValidator.contract.Call(opts, &out, "L2_UPTIME_FEED_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismValidator *OptimismValidatorSession) L2UPTIMEFEEDADDR() (common.Address, error) {
	return _OptimismValidator.Contract.L2UPTIMEFEEDADDR(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCallerSession) L2UPTIMEFEEDADDR() (common.Address, error) {
	return _OptimismValidator.Contract.L2UPTIMEFEEDADDR(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OptimismValidator.contract.Call(opts, &out, "checkEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OptimismValidator *OptimismValidatorSession) CheckEnabled() (bool, error) {
	return _OptimismValidator.Contract.CheckEnabled(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCallerSession) CheckEnabled() (bool, error) {
	return _OptimismValidator.Contract.CheckEnabled(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCaller) GetGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OptimismValidator.contract.Call(opts, &out, "getGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_OptimismValidator *OptimismValidatorSession) GetGasLimit() (uint32, error) {
	return _OptimismValidator.Contract.GetGasLimit(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCallerSession) GetGasLimit() (uint32, error) {
	return _OptimismValidator.Contract.GetGasLimit(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCaller) HasAccess(opts *bind.CallOpts, _user common.Address, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _OptimismValidator.contract.Call(opts, &out, "hasAccess", _user, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OptimismValidator *OptimismValidatorSession) HasAccess(_user common.Address, arg1 []byte) (bool, error) {
	return _OptimismValidator.Contract.HasAccess(&_OptimismValidator.CallOpts, _user, arg1)
}

func (_OptimismValidator *OptimismValidatorCallerSession) HasAccess(_user common.Address, arg1 []byte) (bool, error) {
	return _OptimismValidator.Contract.HasAccess(&_OptimismValidator.CallOpts, _user, arg1)
}

func (_OptimismValidator *OptimismValidatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismValidator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismValidator *OptimismValidatorSession) Owner() (common.Address, error) {
	return _OptimismValidator.Contract.Owner(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCallerSession) Owner() (common.Address, error) {
	return _OptimismValidator.Contract.Owner(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OptimismValidator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OptimismValidator *OptimismValidatorSession) TypeAndVersion() (string, error) {
	return _OptimismValidator.Contract.TypeAndVersion(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorCallerSession) TypeAndVersion() (string, error) {
	return _OptimismValidator.Contract.TypeAndVersion(&_OptimismValidator.CallOpts)
}

func (_OptimismValidator *OptimismValidatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismValidator.contract.Transact(opts, "acceptOwnership")
}

func (_OptimismValidator *OptimismValidatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OptimismValidator.Contract.AcceptOwnership(&_OptimismValidator.TransactOpts)
}

func (_OptimismValidator *OptimismValidatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OptimismValidator.Contract.AcceptOwnership(&_OptimismValidator.TransactOpts)
}

func (_OptimismValidator *OptimismValidatorTransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _OptimismValidator.contract.Transact(opts, "addAccess", _user)
}

func (_OptimismValidator *OptimismValidatorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _OptimismValidator.Contract.AddAccess(&_OptimismValidator.TransactOpts, _user)
}

func (_OptimismValidator *OptimismValidatorTransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _OptimismValidator.Contract.AddAccess(&_OptimismValidator.TransactOpts, _user)
}

func (_OptimismValidator *OptimismValidatorTransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismValidator.contract.Transact(opts, "disableAccessCheck")
}

func (_OptimismValidator *OptimismValidatorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _OptimismValidator.Contract.DisableAccessCheck(&_OptimismValidator.TransactOpts)
}

func (_OptimismValidator *OptimismValidatorTransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _OptimismValidator.Contract.DisableAccessCheck(&_OptimismValidator.TransactOpts)
}

func (_OptimismValidator *OptimismValidatorTransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismValidator.contract.Transact(opts, "enableAccessCheck")
}

func (_OptimismValidator *OptimismValidatorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _OptimismValidator.Contract.EnableAccessCheck(&_OptimismValidator.TransactOpts)
}

func (_OptimismValidator *OptimismValidatorTransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _OptimismValidator.Contract.EnableAccessCheck(&_OptimismValidator.TransactOpts)
}

func (_OptimismValidator *OptimismValidatorTransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _OptimismValidator.contract.Transact(opts, "removeAccess", _user)
}

func (_OptimismValidator *OptimismValidatorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _OptimismValidator.Contract.RemoveAccess(&_OptimismValidator.TransactOpts, _user)
}

func (_OptimismValidator *OptimismValidatorTransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _OptimismValidator.Contract.RemoveAccess(&_OptimismValidator.TransactOpts, _user)
}

func (_OptimismValidator *OptimismValidatorTransactor) SetGasLimit(opts *bind.TransactOpts, gasLimit uint32) (*types.Transaction, error) {
	return _OptimismValidator.contract.Transact(opts, "setGasLimit", gasLimit)
}

func (_OptimismValidator *OptimismValidatorSession) SetGasLimit(gasLimit uint32) (*types.Transaction, error) {
	return _OptimismValidator.Contract.SetGasLimit(&_OptimismValidator.TransactOpts, gasLimit)
}

func (_OptimismValidator *OptimismValidatorTransactorSession) SetGasLimit(gasLimit uint32) (*types.Transaction, error) {
	return _OptimismValidator.Contract.SetGasLimit(&_OptimismValidator.TransactOpts, gasLimit)
}

func (_OptimismValidator *OptimismValidatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OptimismValidator.contract.Transact(opts, "transferOwnership", to)
}

func (_OptimismValidator *OptimismValidatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OptimismValidator.Contract.TransferOwnership(&_OptimismValidator.TransactOpts, to)
}

func (_OptimismValidator *OptimismValidatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OptimismValidator.Contract.TransferOwnership(&_OptimismValidator.TransactOpts, to)
}

func (_OptimismValidator *OptimismValidatorTransactor) Validate(opts *bind.TransactOpts, previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _OptimismValidator.contract.Transact(opts, "validate", previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_OptimismValidator *OptimismValidatorSession) Validate(previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _OptimismValidator.Contract.Validate(&_OptimismValidator.TransactOpts, previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_OptimismValidator *OptimismValidatorTransactorSession) Validate(previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _OptimismValidator.Contract.Validate(&_OptimismValidator.TransactOpts, previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_OptimismValidator *OptimismValidatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismValidator.contract.RawTransact(opts, nil)
}

func (_OptimismValidator *OptimismValidatorSession) Receive() (*types.Transaction, error) {
	return _OptimismValidator.Contract.Receive(&_OptimismValidator.TransactOpts)
}

func (_OptimismValidator *OptimismValidatorTransactorSession) Receive() (*types.Transaction, error) {
	return _OptimismValidator.Contract.Receive(&_OptimismValidator.TransactOpts)
}

type OptimismValidatorAddedAccessIterator struct {
	Event *OptimismValidatorAddedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismValidatorAddedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismValidatorAddedAccess)
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
		it.Event = new(OptimismValidatorAddedAccess)
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

func (it *OptimismValidatorAddedAccessIterator) Error() error {
	return it.fail
}

func (it *OptimismValidatorAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismValidatorAddedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_OptimismValidator *OptimismValidatorFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*OptimismValidatorAddedAccessIterator, error) {

	logs, sub, err := _OptimismValidator.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorAddedAccessIterator{contract: _OptimismValidator.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

func (_OptimismValidator *OptimismValidatorFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *OptimismValidatorAddedAccess) (event.Subscription, error) {

	logs, sub, err := _OptimismValidator.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismValidatorAddedAccess)
				if err := _OptimismValidator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
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

func (_OptimismValidator *OptimismValidatorFilterer) ParseAddedAccess(log types.Log) (*OptimismValidatorAddedAccess, error) {
	event := new(OptimismValidatorAddedAccess)
	if err := _OptimismValidator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismValidatorCheckAccessDisabledIterator struct {
	Event *OptimismValidatorCheckAccessDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismValidatorCheckAccessDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismValidatorCheckAccessDisabled)
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
		it.Event = new(OptimismValidatorCheckAccessDisabled)
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

func (it *OptimismValidatorCheckAccessDisabledIterator) Error() error {
	return it.fail
}

func (it *OptimismValidatorCheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismValidatorCheckAccessDisabled struct {
	Raw types.Log
}

func (_OptimismValidator *OptimismValidatorFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*OptimismValidatorCheckAccessDisabledIterator, error) {

	logs, sub, err := _OptimismValidator.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorCheckAccessDisabledIterator{contract: _OptimismValidator.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

func (_OptimismValidator *OptimismValidatorFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *OptimismValidatorCheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _OptimismValidator.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismValidatorCheckAccessDisabled)
				if err := _OptimismValidator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
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

func (_OptimismValidator *OptimismValidatorFilterer) ParseCheckAccessDisabled(log types.Log) (*OptimismValidatorCheckAccessDisabled, error) {
	event := new(OptimismValidatorCheckAccessDisabled)
	if err := _OptimismValidator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismValidatorCheckAccessEnabledIterator struct {
	Event *OptimismValidatorCheckAccessEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismValidatorCheckAccessEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismValidatorCheckAccessEnabled)
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
		it.Event = new(OptimismValidatorCheckAccessEnabled)
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

func (it *OptimismValidatorCheckAccessEnabledIterator) Error() error {
	return it.fail
}

func (it *OptimismValidatorCheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismValidatorCheckAccessEnabled struct {
	Raw types.Log
}

func (_OptimismValidator *OptimismValidatorFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*OptimismValidatorCheckAccessEnabledIterator, error) {

	logs, sub, err := _OptimismValidator.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorCheckAccessEnabledIterator{contract: _OptimismValidator.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

func (_OptimismValidator *OptimismValidatorFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *OptimismValidatorCheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _OptimismValidator.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismValidatorCheckAccessEnabled)
				if err := _OptimismValidator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
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

func (_OptimismValidator *OptimismValidatorFilterer) ParseCheckAccessEnabled(log types.Log) (*OptimismValidatorCheckAccessEnabled, error) {
	event := new(OptimismValidatorCheckAccessEnabled)
	if err := _OptimismValidator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismValidatorConfigUpdatedIterator struct {
	Event *OptimismValidatorConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismValidatorConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismValidatorConfigUpdated)
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
		it.Event = new(OptimismValidatorConfigUpdated)
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

func (it *OptimismValidatorConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *OptimismValidatorConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismValidatorConfigUpdated struct {
	L1CrossDomainMessengerAddress common.Address
	L2UptimeFeedAddr              common.Address
	GasLimit                      uint32
	Raw                           types.Log
}

func (_OptimismValidator *OptimismValidatorFilterer) FilterConfigUpdated(opts *bind.FilterOpts) (*OptimismValidatorConfigUpdatedIterator, error) {

	logs, sub, err := _OptimismValidator.contract.FilterLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorConfigUpdatedIterator{contract: _OptimismValidator.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

func (_OptimismValidator *OptimismValidatorFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *OptimismValidatorConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _OptimismValidator.contract.WatchLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismValidatorConfigUpdated)
				if err := _OptimismValidator.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

func (_OptimismValidator *OptimismValidatorFilterer) ParseConfigUpdated(log types.Log) (*OptimismValidatorConfigUpdated, error) {
	event := new(OptimismValidatorConfigUpdated)
	if err := _OptimismValidator.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismValidatorGasLimitUpdatedIterator struct {
	Event *OptimismValidatorGasLimitUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismValidatorGasLimitUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismValidatorGasLimitUpdated)
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
		it.Event = new(OptimismValidatorGasLimitUpdated)
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

func (it *OptimismValidatorGasLimitUpdatedIterator) Error() error {
	return it.fail
}

func (it *OptimismValidatorGasLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismValidatorGasLimitUpdated struct {
	GasLimit uint32
	Raw      types.Log
}

func (_OptimismValidator *OptimismValidatorFilterer) FilterGasLimitUpdated(opts *bind.FilterOpts) (*OptimismValidatorGasLimitUpdatedIterator, error) {

	logs, sub, err := _OptimismValidator.contract.FilterLogs(opts, "GasLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorGasLimitUpdatedIterator{contract: _OptimismValidator.contract, event: "GasLimitUpdated", logs: logs, sub: sub}, nil
}

func (_OptimismValidator *OptimismValidatorFilterer) WatchGasLimitUpdated(opts *bind.WatchOpts, sink chan<- *OptimismValidatorGasLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _OptimismValidator.contract.WatchLogs(opts, "GasLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismValidatorGasLimitUpdated)
				if err := _OptimismValidator.contract.UnpackLog(event, "GasLimitUpdated", log); err != nil {
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

func (_OptimismValidator *OptimismValidatorFilterer) ParseGasLimitUpdated(log types.Log) (*OptimismValidatorGasLimitUpdated, error) {
	event := new(OptimismValidatorGasLimitUpdated)
	if err := _OptimismValidator.contract.UnpackLog(event, "GasLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismValidatorOwnershipTransferRequestedIterator struct {
	Event *OptimismValidatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismValidatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismValidatorOwnershipTransferRequested)
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
		it.Event = new(OptimismValidatorOwnershipTransferRequested)
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

func (it *OptimismValidatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OptimismValidatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismValidatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismValidator *OptimismValidatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismValidatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismValidator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorOwnershipTransferRequestedIterator{contract: _OptimismValidator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OptimismValidator *OptimismValidatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismValidatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismValidator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismValidatorOwnershipTransferRequested)
				if err := _OptimismValidator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OptimismValidator *OptimismValidatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*OptimismValidatorOwnershipTransferRequested, error) {
	event := new(OptimismValidatorOwnershipTransferRequested)
	if err := _OptimismValidator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismValidatorOwnershipTransferredIterator struct {
	Event *OptimismValidatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismValidatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismValidatorOwnershipTransferred)
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
		it.Event = new(OptimismValidatorOwnershipTransferred)
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

func (it *OptimismValidatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OptimismValidatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismValidatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismValidator *OptimismValidatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismValidatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismValidator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorOwnershipTransferredIterator{contract: _OptimismValidator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OptimismValidator *OptimismValidatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismValidatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismValidator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismValidatorOwnershipTransferred)
				if err := _OptimismValidator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OptimismValidator *OptimismValidatorFilterer) ParseOwnershipTransferred(log types.Log) (*OptimismValidatorOwnershipTransferred, error) {
	event := new(OptimismValidatorOwnershipTransferred)
	if err := _OptimismValidator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismValidatorRemovedAccessIterator struct {
	Event *OptimismValidatorRemovedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismValidatorRemovedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismValidatorRemovedAccess)
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
		it.Event = new(OptimismValidatorRemovedAccess)
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

func (it *OptimismValidatorRemovedAccessIterator) Error() error {
	return it.fail
}

func (it *OptimismValidatorRemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismValidatorRemovedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_OptimismValidator *OptimismValidatorFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*OptimismValidatorRemovedAccessIterator, error) {

	logs, sub, err := _OptimismValidator.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorRemovedAccessIterator{contract: _OptimismValidator.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

func (_OptimismValidator *OptimismValidatorFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *OptimismValidatorRemovedAccess) (event.Subscription, error) {

	logs, sub, err := _OptimismValidator.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismValidatorRemovedAccess)
				if err := _OptimismValidator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
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

func (_OptimismValidator *OptimismValidatorFilterer) ParseRemovedAccess(log types.Log) (*OptimismValidatorRemovedAccess, error) {
	event := new(OptimismValidatorRemovedAccess)
	if err := _OptimismValidator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismValidatorValidatedStatusIterator struct {
	Event *OptimismValidatorValidatedStatus

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismValidatorValidatedStatusIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismValidatorValidatedStatus)
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
		it.Event = new(OptimismValidatorValidatedStatus)
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

func (it *OptimismValidatorValidatedStatusIterator) Error() error {
	return it.fail
}

func (it *OptimismValidatorValidatedStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismValidatorValidatedStatus struct {
	PreviousRoundId *big.Int
	PreviousAnswer  *big.Int
	CurrentRoundId  *big.Int
	CurrentAnswer   *big.Int
	Raw             types.Log
}

func (_OptimismValidator *OptimismValidatorFilterer) FilterValidatedStatus(opts *bind.FilterOpts) (*OptimismValidatorValidatedStatusIterator, error) {

	logs, sub, err := _OptimismValidator.contract.FilterLogs(opts, "ValidatedStatus")
	if err != nil {
		return nil, err
	}
	return &OptimismValidatorValidatedStatusIterator{contract: _OptimismValidator.contract, event: "ValidatedStatus", logs: logs, sub: sub}, nil
}

func (_OptimismValidator *OptimismValidatorFilterer) WatchValidatedStatus(opts *bind.WatchOpts, sink chan<- *OptimismValidatorValidatedStatus) (event.Subscription, error) {

	logs, sub, err := _OptimismValidator.contract.WatchLogs(opts, "ValidatedStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismValidatorValidatedStatus)
				if err := _OptimismValidator.contract.UnpackLog(event, "ValidatedStatus", log); err != nil {
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

func (_OptimismValidator *OptimismValidatorFilterer) ParseValidatedStatus(log types.Log) (*OptimismValidatorValidatedStatus, error) {
	event := new(OptimismValidatorValidatedStatus)
	if err := _OptimismValidator.contract.UnpackLog(event, "ValidatedStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (OptimismValidatorAddedAccess) Topic() common.Hash {
	return common.HexToHash("0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4")
}

func (OptimismValidatorCheckAccessDisabled) Topic() common.Hash {
	return common.HexToHash("0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638")
}

func (OptimismValidatorCheckAccessEnabled) Topic() common.Hash {
	return common.HexToHash("0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480")
}

func (OptimismValidatorConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x50ac1a0ab48727faf611b51ad202bb80d8cc334ca8e64154c3d430fbb23040cb")
}

func (OptimismValidatorGasLimitUpdated) Topic() common.Hash {
	return common.HexToHash("0x501c51f34bf3d8589156f7fd37055e69375b40034d9f99365cb3290b5d983c91")
}

func (OptimismValidatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OptimismValidatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OptimismValidatorRemovedAccess) Topic() common.Hash {
	return common.HexToHash("0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1")
}

func (OptimismValidatorValidatedStatus) Topic() common.Hash {
	return common.HexToHash("0xdf570e5bda6b7c0c0415ceeafd5f9321908b69023cae41e030a6df91cad90b03")
}

func (_OptimismValidator *OptimismValidator) Address() common.Address {
	return _OptimismValidator.address
}

type OptimismValidatorInterface interface {
	L1CROSSDOMAINMESSENGERADDRESS(opts *bind.CallOpts) (common.Address, error)

	L2UPTIMEFEEDADDR(opts *bind.CallOpts) (common.Address, error)

	CheckEnabled(opts *bind.CallOpts) (bool, error)

	GetGasLimit(opts *bind.CallOpts) (uint32, error)

	HasAccess(opts *bind.CallOpts, _user common.Address, arg1 []byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error)

	DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error)

	EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error)

	SetGasLimit(opts *bind.TransactOpts, gasLimit uint32) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Validate(opts *bind.TransactOpts, previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAddedAccess(opts *bind.FilterOpts) (*OptimismValidatorAddedAccessIterator, error)

	WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *OptimismValidatorAddedAccess) (event.Subscription, error)

	ParseAddedAccess(log types.Log) (*OptimismValidatorAddedAccess, error)

	FilterCheckAccessDisabled(opts *bind.FilterOpts) (*OptimismValidatorCheckAccessDisabledIterator, error)

	WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *OptimismValidatorCheckAccessDisabled) (event.Subscription, error)

	ParseCheckAccessDisabled(log types.Log) (*OptimismValidatorCheckAccessDisabled, error)

	FilterCheckAccessEnabled(opts *bind.FilterOpts) (*OptimismValidatorCheckAccessEnabledIterator, error)

	WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *OptimismValidatorCheckAccessEnabled) (event.Subscription, error)

	ParseCheckAccessEnabled(log types.Log) (*OptimismValidatorCheckAccessEnabled, error)

	FilterConfigUpdated(opts *bind.FilterOpts) (*OptimismValidatorConfigUpdatedIterator, error)

	WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *OptimismValidatorConfigUpdated) (event.Subscription, error)

	ParseConfigUpdated(log types.Log) (*OptimismValidatorConfigUpdated, error)

	FilterGasLimitUpdated(opts *bind.FilterOpts) (*OptimismValidatorGasLimitUpdatedIterator, error)

	WatchGasLimitUpdated(opts *bind.WatchOpts, sink chan<- *OptimismValidatorGasLimitUpdated) (event.Subscription, error)

	ParseGasLimitUpdated(log types.Log) (*OptimismValidatorGasLimitUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismValidatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismValidatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OptimismValidatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismValidatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismValidatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OptimismValidatorOwnershipTransferred, error)

	FilterRemovedAccess(opts *bind.FilterOpts) (*OptimismValidatorRemovedAccessIterator, error)

	WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *OptimismValidatorRemovedAccess) (event.Subscription, error)

	ParseRemovedAccess(log types.Log) (*OptimismValidatorRemovedAccess, error)

	FilterValidatedStatus(opts *bind.FilterOpts) (*OptimismValidatorValidatedStatusIterator, error)

	WatchValidatedStatus(opts *bind.WatchOpts, sink chan<- *OptimismValidatorValidatedStatus) (event.Subscription, error)

	ParseValidatedStatus(log types.Log) (*OptimismValidatorValidatedStatus, error)

	Address() common.Address
}
