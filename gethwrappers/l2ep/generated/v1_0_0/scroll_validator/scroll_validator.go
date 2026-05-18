// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package scroll_validator

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

var ScrollValidatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"l1CrossDomainMessengerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l2UptimeFeedAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1MessageQueueAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"L1_CROSS_DOMAIN_MESSENGER_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L1_MSG_QUEUE_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L2_UPTIME_FEED_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validate\",\"inputs\":[{\"name\":\"previousRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"previousAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"currentRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigUpdated\",\"inputs\":[{\"name\":\"l1CrossDomainMessengerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"l2UptimeFeedAddr\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasLimitUpdated\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatedStatus\",\"inputs\":[{\"name\":\"previousRoundId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"previousAnswer\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"currentRoundId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currentAnswer\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"L1CrossDomainMessengerAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2UptimeFeedAddrZero\",\"inputs\":[]}]",
	Bin: "0x60e060405234801562000010575f80fd5b5060405162001237380380620012378339810160408190526200003391620002cb565b83838233805f816200008c5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be8162000205565b50506001805460ff60a01b1916600160a01b179055506001600160a01b038316620000fc57604051639c9087c160e01b815260040160405180910390fd5b6001600160a01b0382166200012457604051632ef8090560e01b815260040160405180910390fd5b6001600160a01b03838116608081905290831660a08190526003805463ffffffff191663ffffffff8516908117909155604080519384526020840192909252908201527f50ac1a0ab48727faf611b51ad202bb80d8cc334ca8e64154c3d430fbb23040cb9060600160405180910390a15050506001600160a01b038216620001ef5760405162461bcd60e51b815260206004820181905260248201527f496e76616c6964204c31206d6573736167652071756575652061646472657373604482015260640162000083565b506001600160a01b031660c052506200032e9050565b336001600160a01b038216036200025f5760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000083565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b80516001600160a01b0381168114620002c6575f80fd5b919050565b5f805f8060808587031215620002df575f80fd5b620002ea85620002af565b9350620002fa60208601620002af565b92506200030a60408601620002af565b9150606085015163ffffffff8116811462000323575f80fd5b939692955090935050565b60805160a05160c051610ec96200036e5f395f81816101fa01526108f601525f81816101240152610a0d01525f818161035a01526108cc0152610ec95ff3fe6080604052600436106100f2575f3560e01c80638038e4a111610087578063beed9b5111610057578063beed9b51146102f9578063dc7f012414610318578063eda066f714610349578063f2fde38b1461037c575f80fd5b80638038e4a11461027e5780638823da6c146102925780638da5cb5b146102b1578063a118f249146102da575f80fd5b80634394a48a116100c25780634394a48a146101e957806352d84c621461021c5780636b14daf81461023b57806379ba50971461026a575f80fd5b80630a756983146100fd578063122555ff14610113578063181f5a77146101705780631a93d1c3146101c5575f80fd5b366100f957005b5f80fd5b348015610108575f80fd5b5061011161039b565b005b34801561011e575f80fd5b506101467f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561017b575f80fd5b506101b86040518060400160405280601581526020017f5363726f6c6c56616c696461746f7220312e312e30000000000000000000000081525081565b6040516101679190610cae565b3480156101d0575f80fd5b5060035460405163ffffffff9091168152602001610167565b3480156101f4575f80fd5b506101467f000000000000000000000000000000000000000000000000000000000000000081565b348015610227575f80fd5b50610111610236366004610cc0565b610419565b348015610246575f80fd5b5061025a610255366004610d38565b61048b565b6040519015158152602001610167565b348015610275575f80fd5b506101116104e0565b348015610289575f80fd5b506101116105e1565b34801561029d575f80fd5b506101116102ac366004610e12565b610674565b3480156102bc575f80fd5b505f5473ffffffffffffffffffffffffffffffffffffffff16610146565b3480156102e5575f80fd5b506101116102f4366004610e12565b610727565b348015610304575f80fd5b5061025a610313366004610e2b565b6107d9565b348015610323575f80fd5b5060015461025a9074010000000000000000000000000000000000000000900460ff1681565b348015610354575f80fd5b506101467f000000000000000000000000000000000000000000000000000000000000000081565b348015610387575f80fd5b50610111610396366004610e12565b610ac8565b6103a3610ad9565b60015474010000000000000000000000000000000000000000900460ff161561041757600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638905f90a15b565b610421610ad9565b600380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff83169081179091556040519081527f501c51f34bf3d8589156f7fd37055e69375b40034d9f99365cb3290b5d983c91906020015b60405180910390a150565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526002602052604081205460ff16806104d9575060015474010000000000000000000000000000000000000000900460ff16155b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610566576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b6105e9610ad9565b60015474010000000000000000000000000000000000000000900460ff1661041757600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517faebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480905f90a1565b61067c610ad9565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff16156107245773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527f3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d19101610480565b50565b61072f610ad9565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff166107245773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527f87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db49101610480565b5f610819335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061048b92505050565b61087f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f206163636573730000000000000000000000000000000000000000000000604482015260640161055d565b6003546040517fd7704bae00000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081169163b2267a7b917f0000000000000000000000000000000000000000000000000000000000000000169063d7704bae90602401602060405180830381865afa15801561093b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095f9190610e5a565b604080516001871460248201524267ffffffffffffffff1660448083019190915282518083039091018152606490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fed8378f500000000000000000000000000000000000000000000000000000000179052600354915160e085901b7fffffffff00000000000000000000000000000000000000000000000000000000168152610a40927f0000000000000000000000000000000000000000000000000000000000000000925f92909163ffffffff1690600401610e71565b5f604051808303818588803b158015610a57575f80fd5b505af1158015610a69573d5f803e3d5ffd5b50506040805189815260208101899052908101879052606081018690527fdf570e5bda6b7c0c0415ceeafd5f9321908b69023cae41e030a6df91cad90b0393506080019150610ab59050565b60405180910390a1506001949350505050565b610ad0610ad9565b61072481610b59565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610417576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161055d565b3373ffffffffffffffffffffffffffffffffffffffff821603610bd8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161055d565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b5f81518084525f5b81811015610c7157602081850181015186830182015201610c55565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6104d96020830184610c4d565b5f60208284031215610cd0575f80fd5b813563ffffffff811681146104d9575f80fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114610d06575f80fd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215610d49575f80fd5b610d5283610ce3565b9150602083013567ffffffffffffffff80821115610d6e575f80fd5b818501915085601f830112610d81575f80fd5b813581811115610d9357610d93610d0b565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610dd957610dd9610d0b565b81604052828152886020848701011115610df1575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215610e22575f80fd5b6104d982610ce3565b5f805f8060808587031215610e3e575f80fd5b5050823594602084013594506040840135936060013592509050565b5f60208284031215610e6a575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152608060408201525f610ea56080830185610c4d565b905063ffffffff831660608301529594505050505056fea164736f6c6343000818000a",
}

var ScrollValidatorABI = ScrollValidatorMetaData.ABI

var ScrollValidatorBin = ScrollValidatorMetaData.Bin

func DeployScrollValidator(auth *bind.TransactOpts, backend bind.ContractBackend, l1CrossDomainMessengerAddress common.Address, l2UptimeFeedAddr common.Address, l1MessageQueueAddr common.Address, gasLimit uint32) (common.Address, *types.Transaction, *ScrollValidator, error) {
	parsed, err := ScrollValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScrollValidatorBin), backend, l1CrossDomainMessengerAddress, l2UptimeFeedAddr, l1MessageQueueAddr, gasLimit)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ScrollValidator{address: address, abi: *parsed, ScrollValidatorCaller: ScrollValidatorCaller{contract: contract}, ScrollValidatorTransactor: ScrollValidatorTransactor{contract: contract}, ScrollValidatorFilterer: ScrollValidatorFilterer{contract: contract}}, nil
}

type ScrollValidator struct {
	address common.Address
	abi     abi.ABI
	ScrollValidatorCaller
	ScrollValidatorTransactor
	ScrollValidatorFilterer
}

type ScrollValidatorCaller struct {
	contract *bind.BoundContract
}

type ScrollValidatorTransactor struct {
	contract *bind.BoundContract
}

type ScrollValidatorFilterer struct {
	contract *bind.BoundContract
}

type ScrollValidatorSession struct {
	Contract     *ScrollValidator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ScrollValidatorCallerSession struct {
	Contract *ScrollValidatorCaller
	CallOpts bind.CallOpts
}

type ScrollValidatorTransactorSession struct {
	Contract     *ScrollValidatorTransactor
	TransactOpts bind.TransactOpts
}

type ScrollValidatorRaw struct {
	Contract *ScrollValidator
}

type ScrollValidatorCallerRaw struct {
	Contract *ScrollValidatorCaller
}

type ScrollValidatorTransactorRaw struct {
	Contract *ScrollValidatorTransactor
}

func NewScrollValidator(address common.Address, backend bind.ContractBackend) (*ScrollValidator, error) {
	abi, err := abi.JSON(strings.NewReader(ScrollValidatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindScrollValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScrollValidator{address: address, abi: abi, ScrollValidatorCaller: ScrollValidatorCaller{contract: contract}, ScrollValidatorTransactor: ScrollValidatorTransactor{contract: contract}, ScrollValidatorFilterer: ScrollValidatorFilterer{contract: contract}}, nil
}

func NewScrollValidatorCaller(address common.Address, caller bind.ContractCaller) (*ScrollValidatorCaller, error) {
	contract, err := bindScrollValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorCaller{contract: contract}, nil
}

func NewScrollValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ScrollValidatorTransactor, error) {
	contract, err := bindScrollValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorTransactor{contract: contract}, nil
}

func NewScrollValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ScrollValidatorFilterer, error) {
	contract, err := bindScrollValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorFilterer{contract: contract}, nil
}

func bindScrollValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScrollValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ScrollValidator *ScrollValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollValidator.Contract.ScrollValidatorCaller.contract.Call(opts, result, method, params...)
}

func (_ScrollValidator *ScrollValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollValidator.Contract.ScrollValidatorTransactor.contract.Transfer(opts)
}

func (_ScrollValidator *ScrollValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollValidator.Contract.ScrollValidatorTransactor.contract.Transact(opts, method, params...)
}

func (_ScrollValidator *ScrollValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollValidator.Contract.contract.Call(opts, result, method, params...)
}

func (_ScrollValidator *ScrollValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollValidator.Contract.contract.Transfer(opts)
}

func (_ScrollValidator *ScrollValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollValidator.Contract.contract.Transact(opts, method, params...)
}

func (_ScrollValidator *ScrollValidatorCaller) L1CROSSDOMAINMESSENGERADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollValidator.contract.Call(opts, &out, "L1_CROSS_DOMAIN_MESSENGER_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollValidator *ScrollValidatorSession) L1CROSSDOMAINMESSENGERADDRESS() (common.Address, error) {
	return _ScrollValidator.Contract.L1CROSSDOMAINMESSENGERADDRESS(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCallerSession) L1CROSSDOMAINMESSENGERADDRESS() (common.Address, error) {
	return _ScrollValidator.Contract.L1CROSSDOMAINMESSENGERADDRESS(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCaller) L1MSGQUEUEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollValidator.contract.Call(opts, &out, "L1_MSG_QUEUE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollValidator *ScrollValidatorSession) L1MSGQUEUEADDR() (common.Address, error) {
	return _ScrollValidator.Contract.L1MSGQUEUEADDR(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCallerSession) L1MSGQUEUEADDR() (common.Address, error) {
	return _ScrollValidator.Contract.L1MSGQUEUEADDR(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCaller) L2UPTIMEFEEDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollValidator.contract.Call(opts, &out, "L2_UPTIME_FEED_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollValidator *ScrollValidatorSession) L2UPTIMEFEEDADDR() (common.Address, error) {
	return _ScrollValidator.Contract.L2UPTIMEFEEDADDR(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCallerSession) L2UPTIMEFEEDADDR() (common.Address, error) {
	return _ScrollValidator.Contract.L2UPTIMEFEEDADDR(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ScrollValidator.contract.Call(opts, &out, "checkEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ScrollValidator *ScrollValidatorSession) CheckEnabled() (bool, error) {
	return _ScrollValidator.Contract.CheckEnabled(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCallerSession) CheckEnabled() (bool, error) {
	return _ScrollValidator.Contract.CheckEnabled(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCaller) GetGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ScrollValidator.contract.Call(opts, &out, "getGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_ScrollValidator *ScrollValidatorSession) GetGasLimit() (uint32, error) {
	return _ScrollValidator.Contract.GetGasLimit(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCallerSession) GetGasLimit() (uint32, error) {
	return _ScrollValidator.Contract.GetGasLimit(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCaller) HasAccess(opts *bind.CallOpts, _user common.Address, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _ScrollValidator.contract.Call(opts, &out, "hasAccess", _user, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ScrollValidator *ScrollValidatorSession) HasAccess(_user common.Address, arg1 []byte) (bool, error) {
	return _ScrollValidator.Contract.HasAccess(&_ScrollValidator.CallOpts, _user, arg1)
}

func (_ScrollValidator *ScrollValidatorCallerSession) HasAccess(_user common.Address, arg1 []byte) (bool, error) {
	return _ScrollValidator.Contract.HasAccess(&_ScrollValidator.CallOpts, _user, arg1)
}

func (_ScrollValidator *ScrollValidatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollValidator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ScrollValidator *ScrollValidatorSession) Owner() (common.Address, error) {
	return _ScrollValidator.Contract.Owner(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCallerSession) Owner() (common.Address, error) {
	return _ScrollValidator.Contract.Owner(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ScrollValidator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ScrollValidator *ScrollValidatorSession) TypeAndVersion() (string, error) {
	return _ScrollValidator.Contract.TypeAndVersion(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorCallerSession) TypeAndVersion() (string, error) {
	return _ScrollValidator.Contract.TypeAndVersion(&_ScrollValidator.CallOpts)
}

func (_ScrollValidator *ScrollValidatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollValidator.contract.Transact(opts, "acceptOwnership")
}

func (_ScrollValidator *ScrollValidatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ScrollValidator.Contract.AcceptOwnership(&_ScrollValidator.TransactOpts)
}

func (_ScrollValidator *ScrollValidatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ScrollValidator.Contract.AcceptOwnership(&_ScrollValidator.TransactOpts)
}

func (_ScrollValidator *ScrollValidatorTransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _ScrollValidator.contract.Transact(opts, "addAccess", _user)
}

func (_ScrollValidator *ScrollValidatorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _ScrollValidator.Contract.AddAccess(&_ScrollValidator.TransactOpts, _user)
}

func (_ScrollValidator *ScrollValidatorTransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _ScrollValidator.Contract.AddAccess(&_ScrollValidator.TransactOpts, _user)
}

func (_ScrollValidator *ScrollValidatorTransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollValidator.contract.Transact(opts, "disableAccessCheck")
}

func (_ScrollValidator *ScrollValidatorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _ScrollValidator.Contract.DisableAccessCheck(&_ScrollValidator.TransactOpts)
}

func (_ScrollValidator *ScrollValidatorTransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _ScrollValidator.Contract.DisableAccessCheck(&_ScrollValidator.TransactOpts)
}

func (_ScrollValidator *ScrollValidatorTransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollValidator.contract.Transact(opts, "enableAccessCheck")
}

func (_ScrollValidator *ScrollValidatorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _ScrollValidator.Contract.EnableAccessCheck(&_ScrollValidator.TransactOpts)
}

func (_ScrollValidator *ScrollValidatorTransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _ScrollValidator.Contract.EnableAccessCheck(&_ScrollValidator.TransactOpts)
}

func (_ScrollValidator *ScrollValidatorTransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _ScrollValidator.contract.Transact(opts, "removeAccess", _user)
}

func (_ScrollValidator *ScrollValidatorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _ScrollValidator.Contract.RemoveAccess(&_ScrollValidator.TransactOpts, _user)
}

func (_ScrollValidator *ScrollValidatorTransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _ScrollValidator.Contract.RemoveAccess(&_ScrollValidator.TransactOpts, _user)
}

func (_ScrollValidator *ScrollValidatorTransactor) SetGasLimit(opts *bind.TransactOpts, gasLimit uint32) (*types.Transaction, error) {
	return _ScrollValidator.contract.Transact(opts, "setGasLimit", gasLimit)
}

func (_ScrollValidator *ScrollValidatorSession) SetGasLimit(gasLimit uint32) (*types.Transaction, error) {
	return _ScrollValidator.Contract.SetGasLimit(&_ScrollValidator.TransactOpts, gasLimit)
}

func (_ScrollValidator *ScrollValidatorTransactorSession) SetGasLimit(gasLimit uint32) (*types.Transaction, error) {
	return _ScrollValidator.Contract.SetGasLimit(&_ScrollValidator.TransactOpts, gasLimit)
}

func (_ScrollValidator *ScrollValidatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ScrollValidator.contract.Transact(opts, "transferOwnership", to)
}

func (_ScrollValidator *ScrollValidatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ScrollValidator.Contract.TransferOwnership(&_ScrollValidator.TransactOpts, to)
}

func (_ScrollValidator *ScrollValidatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ScrollValidator.Contract.TransferOwnership(&_ScrollValidator.TransactOpts, to)
}

func (_ScrollValidator *ScrollValidatorTransactor) Validate(opts *bind.TransactOpts, previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _ScrollValidator.contract.Transact(opts, "validate", previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_ScrollValidator *ScrollValidatorSession) Validate(previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _ScrollValidator.Contract.Validate(&_ScrollValidator.TransactOpts, previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_ScrollValidator *ScrollValidatorTransactorSession) Validate(previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _ScrollValidator.Contract.Validate(&_ScrollValidator.TransactOpts, previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_ScrollValidator *ScrollValidatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollValidator.contract.RawTransact(opts, nil)
}

func (_ScrollValidator *ScrollValidatorSession) Receive() (*types.Transaction, error) {
	return _ScrollValidator.Contract.Receive(&_ScrollValidator.TransactOpts)
}

func (_ScrollValidator *ScrollValidatorTransactorSession) Receive() (*types.Transaction, error) {
	return _ScrollValidator.Contract.Receive(&_ScrollValidator.TransactOpts)
}

type ScrollValidatorAddedAccessIterator struct {
	Event *ScrollValidatorAddedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollValidatorAddedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollValidatorAddedAccess)
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
		it.Event = new(ScrollValidatorAddedAccess)
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

func (it *ScrollValidatorAddedAccessIterator) Error() error {
	return it.fail
}

func (it *ScrollValidatorAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollValidatorAddedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_ScrollValidator *ScrollValidatorFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*ScrollValidatorAddedAccessIterator, error) {

	logs, sub, err := _ScrollValidator.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorAddedAccessIterator{contract: _ScrollValidator.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

func (_ScrollValidator *ScrollValidatorFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *ScrollValidatorAddedAccess) (event.Subscription, error) {

	logs, sub, err := _ScrollValidator.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollValidatorAddedAccess)
				if err := _ScrollValidator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
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

func (_ScrollValidator *ScrollValidatorFilterer) ParseAddedAccess(log types.Log) (*ScrollValidatorAddedAccess, error) {
	event := new(ScrollValidatorAddedAccess)
	if err := _ScrollValidator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollValidatorCheckAccessDisabledIterator struct {
	Event *ScrollValidatorCheckAccessDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollValidatorCheckAccessDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollValidatorCheckAccessDisabled)
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
		it.Event = new(ScrollValidatorCheckAccessDisabled)
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

func (it *ScrollValidatorCheckAccessDisabledIterator) Error() error {
	return it.fail
}

func (it *ScrollValidatorCheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollValidatorCheckAccessDisabled struct {
	Raw types.Log
}

func (_ScrollValidator *ScrollValidatorFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*ScrollValidatorCheckAccessDisabledIterator, error) {

	logs, sub, err := _ScrollValidator.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorCheckAccessDisabledIterator{contract: _ScrollValidator.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

func (_ScrollValidator *ScrollValidatorFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *ScrollValidatorCheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _ScrollValidator.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollValidatorCheckAccessDisabled)
				if err := _ScrollValidator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
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

func (_ScrollValidator *ScrollValidatorFilterer) ParseCheckAccessDisabled(log types.Log) (*ScrollValidatorCheckAccessDisabled, error) {
	event := new(ScrollValidatorCheckAccessDisabled)
	if err := _ScrollValidator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollValidatorCheckAccessEnabledIterator struct {
	Event *ScrollValidatorCheckAccessEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollValidatorCheckAccessEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollValidatorCheckAccessEnabled)
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
		it.Event = new(ScrollValidatorCheckAccessEnabled)
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

func (it *ScrollValidatorCheckAccessEnabledIterator) Error() error {
	return it.fail
}

func (it *ScrollValidatorCheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollValidatorCheckAccessEnabled struct {
	Raw types.Log
}

func (_ScrollValidator *ScrollValidatorFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*ScrollValidatorCheckAccessEnabledIterator, error) {

	logs, sub, err := _ScrollValidator.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorCheckAccessEnabledIterator{contract: _ScrollValidator.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

func (_ScrollValidator *ScrollValidatorFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *ScrollValidatorCheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _ScrollValidator.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollValidatorCheckAccessEnabled)
				if err := _ScrollValidator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
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

func (_ScrollValidator *ScrollValidatorFilterer) ParseCheckAccessEnabled(log types.Log) (*ScrollValidatorCheckAccessEnabled, error) {
	event := new(ScrollValidatorCheckAccessEnabled)
	if err := _ScrollValidator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollValidatorConfigUpdatedIterator struct {
	Event *ScrollValidatorConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollValidatorConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollValidatorConfigUpdated)
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
		it.Event = new(ScrollValidatorConfigUpdated)
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

func (it *ScrollValidatorConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *ScrollValidatorConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollValidatorConfigUpdated struct {
	L1CrossDomainMessengerAddress common.Address
	L2UptimeFeedAddr              common.Address
	GasLimit                      uint32
	Raw                           types.Log
}

func (_ScrollValidator *ScrollValidatorFilterer) FilterConfigUpdated(opts *bind.FilterOpts) (*ScrollValidatorConfigUpdatedIterator, error) {

	logs, sub, err := _ScrollValidator.contract.FilterLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorConfigUpdatedIterator{contract: _ScrollValidator.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

func (_ScrollValidator *ScrollValidatorFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *ScrollValidatorConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _ScrollValidator.contract.WatchLogs(opts, "ConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollValidatorConfigUpdated)
				if err := _ScrollValidator.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

func (_ScrollValidator *ScrollValidatorFilterer) ParseConfigUpdated(log types.Log) (*ScrollValidatorConfigUpdated, error) {
	event := new(ScrollValidatorConfigUpdated)
	if err := _ScrollValidator.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollValidatorGasLimitUpdatedIterator struct {
	Event *ScrollValidatorGasLimitUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollValidatorGasLimitUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollValidatorGasLimitUpdated)
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
		it.Event = new(ScrollValidatorGasLimitUpdated)
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

func (it *ScrollValidatorGasLimitUpdatedIterator) Error() error {
	return it.fail
}

func (it *ScrollValidatorGasLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollValidatorGasLimitUpdated struct {
	GasLimit uint32
	Raw      types.Log
}

func (_ScrollValidator *ScrollValidatorFilterer) FilterGasLimitUpdated(opts *bind.FilterOpts) (*ScrollValidatorGasLimitUpdatedIterator, error) {

	logs, sub, err := _ScrollValidator.contract.FilterLogs(opts, "GasLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorGasLimitUpdatedIterator{contract: _ScrollValidator.contract, event: "GasLimitUpdated", logs: logs, sub: sub}, nil
}

func (_ScrollValidator *ScrollValidatorFilterer) WatchGasLimitUpdated(opts *bind.WatchOpts, sink chan<- *ScrollValidatorGasLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _ScrollValidator.contract.WatchLogs(opts, "GasLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollValidatorGasLimitUpdated)
				if err := _ScrollValidator.contract.UnpackLog(event, "GasLimitUpdated", log); err != nil {
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

func (_ScrollValidator *ScrollValidatorFilterer) ParseGasLimitUpdated(log types.Log) (*ScrollValidatorGasLimitUpdated, error) {
	event := new(ScrollValidatorGasLimitUpdated)
	if err := _ScrollValidator.contract.UnpackLog(event, "GasLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollValidatorOwnershipTransferRequestedIterator struct {
	Event *ScrollValidatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollValidatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollValidatorOwnershipTransferRequested)
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
		it.Event = new(ScrollValidatorOwnershipTransferRequested)
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

func (it *ScrollValidatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ScrollValidatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollValidatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollValidator *ScrollValidatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollValidatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollValidator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorOwnershipTransferRequestedIterator{contract: _ScrollValidator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ScrollValidator *ScrollValidatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollValidatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollValidator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollValidatorOwnershipTransferRequested)
				if err := _ScrollValidator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ScrollValidator *ScrollValidatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ScrollValidatorOwnershipTransferRequested, error) {
	event := new(ScrollValidatorOwnershipTransferRequested)
	if err := _ScrollValidator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollValidatorOwnershipTransferredIterator struct {
	Event *ScrollValidatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollValidatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollValidatorOwnershipTransferred)
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
		it.Event = new(ScrollValidatorOwnershipTransferred)
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

func (it *ScrollValidatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ScrollValidatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollValidatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ScrollValidator *ScrollValidatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollValidatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollValidator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorOwnershipTransferredIterator{contract: _ScrollValidator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ScrollValidator *ScrollValidatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollValidatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ScrollValidator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollValidatorOwnershipTransferred)
				if err := _ScrollValidator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ScrollValidator *ScrollValidatorFilterer) ParseOwnershipTransferred(log types.Log) (*ScrollValidatorOwnershipTransferred, error) {
	event := new(ScrollValidatorOwnershipTransferred)
	if err := _ScrollValidator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollValidatorRemovedAccessIterator struct {
	Event *ScrollValidatorRemovedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollValidatorRemovedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollValidatorRemovedAccess)
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
		it.Event = new(ScrollValidatorRemovedAccess)
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

func (it *ScrollValidatorRemovedAccessIterator) Error() error {
	return it.fail
}

func (it *ScrollValidatorRemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollValidatorRemovedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_ScrollValidator *ScrollValidatorFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*ScrollValidatorRemovedAccessIterator, error) {

	logs, sub, err := _ScrollValidator.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorRemovedAccessIterator{contract: _ScrollValidator.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

func (_ScrollValidator *ScrollValidatorFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *ScrollValidatorRemovedAccess) (event.Subscription, error) {

	logs, sub, err := _ScrollValidator.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollValidatorRemovedAccess)
				if err := _ScrollValidator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
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

func (_ScrollValidator *ScrollValidatorFilterer) ParseRemovedAccess(log types.Log) (*ScrollValidatorRemovedAccess, error) {
	event := new(ScrollValidatorRemovedAccess)
	if err := _ScrollValidator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ScrollValidatorValidatedStatusIterator struct {
	Event *ScrollValidatorValidatedStatus

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ScrollValidatorValidatedStatusIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollValidatorValidatedStatus)
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
		it.Event = new(ScrollValidatorValidatedStatus)
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

func (it *ScrollValidatorValidatedStatusIterator) Error() error {
	return it.fail
}

func (it *ScrollValidatorValidatedStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ScrollValidatorValidatedStatus struct {
	PreviousRoundId *big.Int
	PreviousAnswer  *big.Int
	CurrentRoundId  *big.Int
	CurrentAnswer   *big.Int
	Raw             types.Log
}

func (_ScrollValidator *ScrollValidatorFilterer) FilterValidatedStatus(opts *bind.FilterOpts) (*ScrollValidatorValidatedStatusIterator, error) {

	logs, sub, err := _ScrollValidator.contract.FilterLogs(opts, "ValidatedStatus")
	if err != nil {
		return nil, err
	}
	return &ScrollValidatorValidatedStatusIterator{contract: _ScrollValidator.contract, event: "ValidatedStatus", logs: logs, sub: sub}, nil
}

func (_ScrollValidator *ScrollValidatorFilterer) WatchValidatedStatus(opts *bind.WatchOpts, sink chan<- *ScrollValidatorValidatedStatus) (event.Subscription, error) {

	logs, sub, err := _ScrollValidator.contract.WatchLogs(opts, "ValidatedStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ScrollValidatorValidatedStatus)
				if err := _ScrollValidator.contract.UnpackLog(event, "ValidatedStatus", log); err != nil {
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

func (_ScrollValidator *ScrollValidatorFilterer) ParseValidatedStatus(log types.Log) (*ScrollValidatorValidatedStatus, error) {
	event := new(ScrollValidatorValidatedStatus)
	if err := _ScrollValidator.contract.UnpackLog(event, "ValidatedStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ScrollValidatorAddedAccess) Topic() common.Hash {
	return common.HexToHash("0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4")
}

func (ScrollValidatorCheckAccessDisabled) Topic() common.Hash {
	return common.HexToHash("0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638")
}

func (ScrollValidatorCheckAccessEnabled) Topic() common.Hash {
	return common.HexToHash("0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480")
}

func (ScrollValidatorConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0x50ac1a0ab48727faf611b51ad202bb80d8cc334ca8e64154c3d430fbb23040cb")
}

func (ScrollValidatorGasLimitUpdated) Topic() common.Hash {
	return common.HexToHash("0x501c51f34bf3d8589156f7fd37055e69375b40034d9f99365cb3290b5d983c91")
}

func (ScrollValidatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ScrollValidatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ScrollValidatorRemovedAccess) Topic() common.Hash {
	return common.HexToHash("0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1")
}

func (ScrollValidatorValidatedStatus) Topic() common.Hash {
	return common.HexToHash("0xdf570e5bda6b7c0c0415ceeafd5f9321908b69023cae41e030a6df91cad90b03")
}

func (_ScrollValidator *ScrollValidator) Address() common.Address {
	return _ScrollValidator.address
}

type ScrollValidatorInterface interface {
	L1CROSSDOMAINMESSENGERADDRESS(opts *bind.CallOpts) (common.Address, error)

	L1MSGQUEUEADDR(opts *bind.CallOpts) (common.Address, error)

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

	FilterAddedAccess(opts *bind.FilterOpts) (*ScrollValidatorAddedAccessIterator, error)

	WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *ScrollValidatorAddedAccess) (event.Subscription, error)

	ParseAddedAccess(log types.Log) (*ScrollValidatorAddedAccess, error)

	FilterCheckAccessDisabled(opts *bind.FilterOpts) (*ScrollValidatorCheckAccessDisabledIterator, error)

	WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *ScrollValidatorCheckAccessDisabled) (event.Subscription, error)

	ParseCheckAccessDisabled(log types.Log) (*ScrollValidatorCheckAccessDisabled, error)

	FilterCheckAccessEnabled(opts *bind.FilterOpts) (*ScrollValidatorCheckAccessEnabledIterator, error)

	WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *ScrollValidatorCheckAccessEnabled) (event.Subscription, error)

	ParseCheckAccessEnabled(log types.Log) (*ScrollValidatorCheckAccessEnabled, error)

	FilterConfigUpdated(opts *bind.FilterOpts) (*ScrollValidatorConfigUpdatedIterator, error)

	WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *ScrollValidatorConfigUpdated) (event.Subscription, error)

	ParseConfigUpdated(log types.Log) (*ScrollValidatorConfigUpdated, error)

	FilterGasLimitUpdated(opts *bind.FilterOpts) (*ScrollValidatorGasLimitUpdatedIterator, error)

	WatchGasLimitUpdated(opts *bind.WatchOpts, sink chan<- *ScrollValidatorGasLimitUpdated) (event.Subscription, error)

	ParseGasLimitUpdated(log types.Log) (*ScrollValidatorGasLimitUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollValidatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ScrollValidatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ScrollValidatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ScrollValidatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollValidatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ScrollValidatorOwnershipTransferred, error)

	FilterRemovedAccess(opts *bind.FilterOpts) (*ScrollValidatorRemovedAccessIterator, error)

	WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *ScrollValidatorRemovedAccess) (event.Subscription, error)

	ParseRemovedAccess(log types.Log) (*ScrollValidatorRemovedAccess, error)

	FilterValidatedStatus(opts *bind.FilterOpts) (*ScrollValidatorValidatedStatusIterator, error)

	WatchValidatedStatus(opts *bind.WatchOpts, sink chan<- *ScrollValidatorValidatedStatus) (event.Subscription, error)

	ParseValidatedStatus(log types.Log) (*ScrollValidatorValidatedStatus, error)

	Address() common.Address
}
