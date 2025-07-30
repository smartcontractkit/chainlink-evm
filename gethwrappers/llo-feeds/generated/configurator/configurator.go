// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package configurator

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

var ConfiguratorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"promoteStagingConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isGreenProduction\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProductionConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"onchainConfig\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainConfigVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offchainConfig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStagingConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"onchainConfig\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"offchainConfigVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offchainConfig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"isVerifier\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProductionConfigSet\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"configCount\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"signers\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"},{\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"onchainConfig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"offchainConfigVersion\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"offchainConfig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"isGreenProduction\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PromoteStagingConfig\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"retiredConfigDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"isGreenProduction\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StagingConfigSet\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"configCount\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"signers\",\"type\":\"bytes[]\",\"indexed\":false,\"internalType\":\"bytes[]\"},{\"name\":\"offchainTransmitters\",\"type\":\"bytes32[]\",\"indexed\":false,\"internalType\":\"bytes32[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"onchainConfig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"offchainConfigVersion\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"offchainConfig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"isGreenProduction\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ConfigUnset\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ConfigUnsetProduction\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isGreenProduction\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"ConfigUnsetStaging\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isGreenProduction\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"ExcessSigners\",\"inputs\":[{\"name\":\"numSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientSigners\",\"inputs\":[{\"name\":\"numSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidOnchainLength\",\"inputs\":[{\"name\":\"onchainConfigLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidPredecessorConfigDigest\",\"inputs\":[{\"name\":\"predecessorConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"IsGreenProductionMustMatchContractState\",\"inputs\":[{\"name\":\"configId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isGreenProductionContractState\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"type\":\"error\",\"name\":\"NonZeroPredecessorConfigDigest\",\"inputs\":[{\"name\":\"predecessorConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UnsupportedOnchainConfigVersion\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6113e5806101576000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b14610153578063dfb533d01461017b578063e6e7c5a41461018e578063f2fde38b146101a157600080fd5b806301ffc9a71461008d578063181f5a77146100f7578063790464e01461013657806379ba50971461014b575b600080fd5b6100e261009b366004610cfd565b7fffffffff00000000000000000000000000000000000000000000000000000000167f40569294000000000000000000000000000000000000000000000000000000001490565b60405190151581526020015b60405180910390f35b604080518082018252601281527f436f6e666967757261746f7220302e352e300000000000000000000000000000602082015290516100ee9190610daa565b610149610144366004611006565b6101b4565b005b610149610353565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ee565b610149610189366004611006565b610450565b61014961019c3660046110de565b610687565b6101496101af366004611113565b610893565b85518460ff16601f821115610204576040517f61750f4000000000000000000000000000000000000000000000000000000000815260048101839052601f60248201526044015b60405180910390fd5b61020f816003611178565b82116102675781610221826003611178565b61022c906001611195565b6040517f9dd9e6d8000000000000000000000000000000000000000000000000000000008152600481019290925260248201526044016101fb565b61026f6108a7565b6040855110156102b05784516040517f3e936ca80000000000000000000000000000000000000000000000000000000081526004016101fb91815260200190565b6020850151604086015160018210156102f8576040517f8f01e0d7000000000000000000000000000000000000000000000000000000008152600481018390526024016101fb565b8015610333576040517fb96bb760000000000000000000000000000000000000000000000000000000008152600481018290526024016101fb565b6103468b46308d8d8d8d8d8d600161092a565b5050505050505050505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146103d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016101fb565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b85518460ff16601f82111561049b576040517f61750f4000000000000000000000000000000000000000000000000000000000815260048101839052601f60248201526044016101fb565b6104a6816003611178565b82116104b85781610221826003611178565b6104c06108a7565b6040855110156105015784516040517f3e936ca80000000000000000000000000000000000000000000000000000000081526004016101fb91815260200190565b602085015160408601516001821015610549576040517f8f01e0d7000000000000000000000000000000000000000000000000000000008152600481018390526024016101fb565b60008b81526002602081815260408084208151608081018352815467ffffffffffffffff8116825268010000000000000000810463ffffffff16948201949094526c0100000000000000000000000090930460ff161515838301528151808301928390529293909260608501929091600185019182845b8154815260200190600101908083116105c05750505050508152505090506000801b82148061062c5750600260008d81526020019081526020016000206001018160400151610610576000610613565b60015b60ff1660028110610626576106266111a8565b01548214155b15610666576040517f7d78c2a1000000000000000000000000000000000000000000000000000000008152600481018390526024016101fb565b6106798c46308e8e8e8e8e8e600061092a565b505050505050505050505050565b61068f6108a7565b600082815260026020526040902080546c01000000000000000000000000900460ff161515821515146106f8576040517f85fa3a3700000000000000000000000000000000000000000000000000000000815260048101849052821560248201526044016101fb565b805467ffffffffffffffff16600003610740576040517f90e6f6dc000000000000000000000000000000000000000000000000000000008152600481018490526024016101fb565b60006001820183610752576001610755565b60005b60ff1660028110610768576107686111a8565b0154036107ac576040517f5b7f63570000000000000000000000000000000000000000000000000000000081526004810184905282151560248201526044016101fb565b600081600101836107be5760006107c1565b60015b60ff16600281106107d4576107d46111a8565b015490508061081a576040517fcaf1e7730000000000000000000000000000000000000000000000000000000081526004810185905283151560248201526044016101fb565b81547fffffffffffffffffffffffffffffffffffffff00ffffffffffffffffffffffff1683156c010000000000000000000000008102919091178355604051908152819085907f1062aa08ac6046a0e69e3eafdf12d1eba63a67b71a874623e86eb06348a1d84f9060200160405180910390a350505050565b61089b6108a7565b6108a481610b5a565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610928576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016101fb565b565b60008a81526002602052604081208054909190829082906109549067ffffffffffffffff166111d7565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790559050600061098f8d8d8d858e8e8e8e8e8e610c4f565b90508315610a57578c7f261b20c2ecd99d86d6e936279e4f78db34603a3de3a4a84d6f3d4e0dd55e24788460000160089054906101000a900463ffffffff1683858e8e8e8e8e8e8d600001600c9054906101000a900460ff166040516109fe9a9998979695949392919061128b565b60405180910390a260008d815260026020526040902083548291600101906c01000000000000000000000000900460ff16610a3a576000610a3d565b60015b60ff1660028110610a5057610a506111a8565b0155610b13565b8c7fef1b5f9d1b927b0fe871b12c7e7846457602d67b2bc36b0bc95feaf480e890568460000160089054906101000a900463ffffffff1683858e8e8e8e8e8e8d600001600c9054906101000a900460ff16604051610abe9a9998979695949392919061128b565b60405180910390a260008d815260026020526040902083548291600101906c01000000000000000000000000900460ff16610afa576001610afd565b60005b60ff1660028110610b1057610b106111a8565b01555b505080547fffffffffffffffffffffffffffffffffffffffff00000000ffffffffffffffff16680100000000000000004363ffffffff160217905550505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610bd9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016101fb565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000808b8b8b8b8b8b8b8b8b8b604051602001610c759a9998979695949392919061132b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101207dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e09000000000000000000000000000000000000000000000000000000000000179150509a9950505050505050505050565b600060208284031215610d0f57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610d3f57600080fd5b9392505050565b6000815180845260005b81811015610d6c57602081850181015186830182015201610d50565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610d3f6020830184610d46565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610e3357610e33610dbd565b604052919050565b600067ffffffffffffffff821115610e5557610e55610dbd565b5060051b60200190565b600082601f830112610e7057600080fd5b813567ffffffffffffffff811115610e8a57610e8a610dbd565b610ebb60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610dec565b818152846020838601011115610ed057600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112610efe57600080fd5b81356020610f13610f0e83610e3b565b610dec565b82815260059290921b84018101918181019086841115610f3257600080fd5b8286015b84811015610f7257803567ffffffffffffffff811115610f565760008081fd5b610f648986838b0101610e5f565b845250918301918301610f36565b509695505050505050565b600082601f830112610f8e57600080fd5b81356020610f9e610f0e83610e3b565b82815260059290921b84018101918181019086841115610fbd57600080fd5b8286015b84811015610f725780358352918301918301610fc1565b803560ff81168114610fe957600080fd5b919050565b803567ffffffffffffffff81168114610fe957600080fd5b600080600080600080600060e0888a03121561102157600080fd5b87359650602088013567ffffffffffffffff8082111561104057600080fd5b61104c8b838c01610eed565b975060408a013591508082111561106257600080fd5b61106e8b838c01610f7d565b965061107c60608b01610fd8565b955060808a013591508082111561109257600080fd5b61109e8b838c01610e5f565b94506110ac60a08b01610fee565b935060c08a01359150808211156110c257600080fd5b506110cf8a828b01610e5f565b91505092959891949750929550565b600080604083850312156110f157600080fd5b823591506020830135801515811461110857600080fd5b809150509250929050565b60006020828403121561112557600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610d3f57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761118f5761118f611149565b92915050565b8082018082111561118f5761118f611149565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600067ffffffffffffffff8083168181036111f4576111f4611149565b6001019392505050565b6000815180845260208085019450848260051b860182860160005b85811015611243578383038952611231838351610d46565b98850198925090840190600101611219565b5090979650505050505050565b600081518084526020808501945080840160005b8381101561128057815187529582019590820190600101611264565b509495945050505050565b600061014063ffffffff8d1683528b602084015267ffffffffffffffff808c1660408501528160608501526112c28285018c6111fe565b915083820360808501526112d6828b611250565b915060ff891660a085015283820360c08501526112f38289610d46565b90871660e085015283810361010085015290506113108186610d46565b9150508215156101208301529b9a5050505050505050505050565b60006101408c83528b602084015273ffffffffffffffffffffffffffffffffffffffff8b16604084015267ffffffffffffffff808b1660608501528160808501526113788285018b6111fe565b915083820360a085015261138c828a611250565b915060ff881660c085015283820360e08501526113a98288610d46565b90861661010085015283810361012085015290506113c78185610d46565b9d9c5050505050505050505050505056fea164736f6c6343000813000a",
}

var ConfiguratorABI = ConfiguratorMetaData.ABI

var ConfiguratorBin = ConfiguratorMetaData.Bin

func DeployConfigurator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Configurator, error) {
	parsed, err := ConfiguratorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConfiguratorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Configurator{address: address, abi: *parsed, ConfiguratorCaller: ConfiguratorCaller{contract: contract}, ConfiguratorTransactor: ConfiguratorTransactor{contract: contract}, ConfiguratorFilterer: ConfiguratorFilterer{contract: contract}}, nil
}

type Configurator struct {
	address common.Address
	abi     abi.ABI
	ConfiguratorCaller
	ConfiguratorTransactor
	ConfiguratorFilterer
}

type ConfiguratorCaller struct {
	contract *bind.BoundContract
}

type ConfiguratorTransactor struct {
	contract *bind.BoundContract
}

type ConfiguratorFilterer struct {
	contract *bind.BoundContract
}

type ConfiguratorSession struct {
	Contract     *Configurator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ConfiguratorCallerSession struct {
	Contract *ConfiguratorCaller
	CallOpts bind.CallOpts
}

type ConfiguratorTransactorSession struct {
	Contract     *ConfiguratorTransactor
	TransactOpts bind.TransactOpts
}

type ConfiguratorRaw struct {
	Contract *Configurator
}

type ConfiguratorCallerRaw struct {
	Contract *ConfiguratorCaller
}

type ConfiguratorTransactorRaw struct {
	Contract *ConfiguratorTransactor
}

func NewConfigurator(address common.Address, backend bind.ContractBackend) (*Configurator, error) {
	abi, err := abi.JSON(strings.NewReader(ConfiguratorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindConfigurator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Configurator{address: address, abi: abi, ConfiguratorCaller: ConfiguratorCaller{contract: contract}, ConfiguratorTransactor: ConfiguratorTransactor{contract: contract}, ConfiguratorFilterer: ConfiguratorFilterer{contract: contract}}, nil
}

func NewConfiguratorCaller(address common.Address, caller bind.ContractCaller) (*ConfiguratorCaller, error) {
	contract, err := bindConfigurator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorCaller{contract: contract}, nil
}

func NewConfiguratorTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfiguratorTransactor, error) {
	contract, err := bindConfigurator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorTransactor{contract: contract}, nil
}

func NewConfiguratorFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfiguratorFilterer, error) {
	contract, err := bindConfigurator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorFilterer{contract: contract}, nil
}

func bindConfigurator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConfiguratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Configurator *ConfiguratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Configurator.Contract.ConfiguratorCaller.contract.Call(opts, result, method, params...)
}

func (_Configurator *ConfiguratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Configurator.Contract.ConfiguratorTransactor.contract.Transfer(opts)
}

func (_Configurator *ConfiguratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Configurator.Contract.ConfiguratorTransactor.contract.Transact(opts, method, params...)
}

func (_Configurator *ConfiguratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Configurator.Contract.contract.Call(opts, result, method, params...)
}

func (_Configurator *ConfiguratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Configurator.Contract.contract.Transfer(opts)
}

func (_Configurator *ConfiguratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Configurator.Contract.contract.Transact(opts, method, params...)
}

func (_Configurator *ConfiguratorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Configurator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Configurator *ConfiguratorSession) Owner() (common.Address, error) {
	return _Configurator.Contract.Owner(&_Configurator.CallOpts)
}

func (_Configurator *ConfiguratorCallerSession) Owner() (common.Address, error) {
	return _Configurator.Contract.Owner(&_Configurator.CallOpts)
}

func (_Configurator *ConfiguratorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Configurator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Configurator *ConfiguratorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Configurator.Contract.SupportsInterface(&_Configurator.CallOpts, interfaceId)
}

func (_Configurator *ConfiguratorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Configurator.Contract.SupportsInterface(&_Configurator.CallOpts, interfaceId)
}

func (_Configurator *ConfiguratorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Configurator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Configurator *ConfiguratorSession) TypeAndVersion() (string, error) {
	return _Configurator.Contract.TypeAndVersion(&_Configurator.CallOpts)
}

func (_Configurator *ConfiguratorCallerSession) TypeAndVersion() (string, error) {
	return _Configurator.Contract.TypeAndVersion(&_Configurator.CallOpts)
}

func (_Configurator *ConfiguratorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "acceptOwnership")
}

func (_Configurator *ConfiguratorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Configurator.Contract.AcceptOwnership(&_Configurator.TransactOpts)
}

func (_Configurator *ConfiguratorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Configurator.Contract.AcceptOwnership(&_Configurator.TransactOpts)
}

func (_Configurator *ConfiguratorTransactor) PromoteStagingConfig(opts *bind.TransactOpts, configId [32]byte, isGreenProduction bool) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "promoteStagingConfig", configId, isGreenProduction)
}

func (_Configurator *ConfiguratorSession) PromoteStagingConfig(configId [32]byte, isGreenProduction bool) (*types.Transaction, error) {
	return _Configurator.Contract.PromoteStagingConfig(&_Configurator.TransactOpts, configId, isGreenProduction)
}

func (_Configurator *ConfiguratorTransactorSession) PromoteStagingConfig(configId [32]byte, isGreenProduction bool) (*types.Transaction, error) {
	return _Configurator.Contract.PromoteStagingConfig(&_Configurator.TransactOpts, configId, isGreenProduction)
}

func (_Configurator *ConfiguratorTransactor) SetProductionConfig(opts *bind.TransactOpts, configId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setProductionConfig", configId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_Configurator *ConfiguratorSession) SetProductionConfig(configId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _Configurator.Contract.SetProductionConfig(&_Configurator.TransactOpts, configId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_Configurator *ConfiguratorTransactorSession) SetProductionConfig(configId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _Configurator.Contract.SetProductionConfig(&_Configurator.TransactOpts, configId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_Configurator *ConfiguratorTransactor) SetStagingConfig(opts *bind.TransactOpts, configId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "setStagingConfig", configId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_Configurator *ConfiguratorSession) SetStagingConfig(configId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _Configurator.Contract.SetStagingConfig(&_Configurator.TransactOpts, configId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_Configurator *ConfiguratorTransactorSession) SetStagingConfig(configId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error) {
	return _Configurator.Contract.SetStagingConfig(&_Configurator.TransactOpts, configId, signers, offchainTransmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
}

func (_Configurator *ConfiguratorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Configurator.contract.Transact(opts, "transferOwnership", to)
}

func (_Configurator *ConfiguratorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.TransferOwnership(&_Configurator.TransactOpts, to)
}

func (_Configurator *ConfiguratorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Configurator.Contract.TransferOwnership(&_Configurator.TransactOpts, to)
}

type ConfiguratorOwnershipTransferRequestedIterator struct {
	Event *ConfiguratorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConfiguratorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorOwnershipTransferRequested)
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
		it.Event = new(ConfiguratorOwnershipTransferRequested)
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

func (it *ConfiguratorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ConfiguratorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConfiguratorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Configurator *ConfiguratorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConfiguratorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorOwnershipTransferRequestedIterator{contract: _Configurator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Configurator *ConfiguratorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ConfiguratorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConfiguratorOwnershipTransferRequested)
				if err := _Configurator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Configurator *ConfiguratorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ConfiguratorOwnershipTransferRequested, error) {
	event := new(ConfiguratorOwnershipTransferRequested)
	if err := _Configurator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConfiguratorOwnershipTransferredIterator struct {
	Event *ConfiguratorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConfiguratorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorOwnershipTransferred)
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
		it.Event = new(ConfiguratorOwnershipTransferred)
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

func (it *ConfiguratorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ConfiguratorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConfiguratorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Configurator *ConfiguratorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConfiguratorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorOwnershipTransferredIterator{contract: _Configurator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Configurator *ConfiguratorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConfiguratorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConfiguratorOwnershipTransferred)
				if err := _Configurator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Configurator *ConfiguratorFilterer) ParseOwnershipTransferred(log types.Log) (*ConfiguratorOwnershipTransferred, error) {
	event := new(ConfiguratorOwnershipTransferred)
	if err := _Configurator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConfiguratorProductionConfigSetIterator struct {
	Event *ConfiguratorProductionConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConfiguratorProductionConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorProductionConfigSet)
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
		it.Event = new(ConfiguratorProductionConfigSet)
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

func (it *ConfiguratorProductionConfigSetIterator) Error() error {
	return it.fail
}

func (it *ConfiguratorProductionConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConfiguratorProductionConfigSet struct {
	ConfigId                  [32]byte
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   [][]byte
	OffchainTransmitters      [][32]byte
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	IsGreenProduction         bool
	Raw                       types.Log
}

func (_Configurator *ConfiguratorFilterer) FilterProductionConfigSet(opts *bind.FilterOpts, configId [][32]byte) (*ConfiguratorProductionConfigSetIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "ProductionConfigSet", configIdRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorProductionConfigSetIterator{contract: _Configurator.contract, event: "ProductionConfigSet", logs: logs, sub: sub}, nil
}

func (_Configurator *ConfiguratorFilterer) WatchProductionConfigSet(opts *bind.WatchOpts, sink chan<- *ConfiguratorProductionConfigSet, configId [][32]byte) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "ProductionConfigSet", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConfiguratorProductionConfigSet)
				if err := _Configurator.contract.UnpackLog(event, "ProductionConfigSet", log); err != nil {
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

func (_Configurator *ConfiguratorFilterer) ParseProductionConfigSet(log types.Log) (*ConfiguratorProductionConfigSet, error) {
	event := new(ConfiguratorProductionConfigSet)
	if err := _Configurator.contract.UnpackLog(event, "ProductionConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConfiguratorPromoteStagingConfigIterator struct {
	Event *ConfiguratorPromoteStagingConfig

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConfiguratorPromoteStagingConfigIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorPromoteStagingConfig)
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
		it.Event = new(ConfiguratorPromoteStagingConfig)
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

func (it *ConfiguratorPromoteStagingConfigIterator) Error() error {
	return it.fail
}

func (it *ConfiguratorPromoteStagingConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConfiguratorPromoteStagingConfig struct {
	ConfigId            [32]byte
	RetiredConfigDigest [32]byte
	IsGreenProduction   bool
	Raw                 types.Log
}

func (_Configurator *ConfiguratorFilterer) FilterPromoteStagingConfig(opts *bind.FilterOpts, configId [][32]byte, retiredConfigDigest [][32]byte) (*ConfiguratorPromoteStagingConfigIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}
	var retiredConfigDigestRule []interface{}
	for _, retiredConfigDigestItem := range retiredConfigDigest {
		retiredConfigDigestRule = append(retiredConfigDigestRule, retiredConfigDigestItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "PromoteStagingConfig", configIdRule, retiredConfigDigestRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorPromoteStagingConfigIterator{contract: _Configurator.contract, event: "PromoteStagingConfig", logs: logs, sub: sub}, nil
}

func (_Configurator *ConfiguratorFilterer) WatchPromoteStagingConfig(opts *bind.WatchOpts, sink chan<- *ConfiguratorPromoteStagingConfig, configId [][32]byte, retiredConfigDigest [][32]byte) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}
	var retiredConfigDigestRule []interface{}
	for _, retiredConfigDigestItem := range retiredConfigDigest {
		retiredConfigDigestRule = append(retiredConfigDigestRule, retiredConfigDigestItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "PromoteStagingConfig", configIdRule, retiredConfigDigestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConfiguratorPromoteStagingConfig)
				if err := _Configurator.contract.UnpackLog(event, "PromoteStagingConfig", log); err != nil {
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

func (_Configurator *ConfiguratorFilterer) ParsePromoteStagingConfig(log types.Log) (*ConfiguratorPromoteStagingConfig, error) {
	event := new(ConfiguratorPromoteStagingConfig)
	if err := _Configurator.contract.UnpackLog(event, "PromoteStagingConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ConfiguratorStagingConfigSetIterator struct {
	Event *ConfiguratorStagingConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ConfiguratorStagingConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfiguratorStagingConfigSet)
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
		it.Event = new(ConfiguratorStagingConfigSet)
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

func (it *ConfiguratorStagingConfigSetIterator) Error() error {
	return it.fail
}

func (it *ConfiguratorStagingConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ConfiguratorStagingConfigSet struct {
	ConfigId                  [32]byte
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   [][]byte
	OffchainTransmitters      [][32]byte
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	IsGreenProduction         bool
	Raw                       types.Log
}

func (_Configurator *ConfiguratorFilterer) FilterStagingConfigSet(opts *bind.FilterOpts, configId [][32]byte) (*ConfiguratorStagingConfigSetIterator, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _Configurator.contract.FilterLogs(opts, "StagingConfigSet", configIdRule)
	if err != nil {
		return nil, err
	}
	return &ConfiguratorStagingConfigSetIterator{contract: _Configurator.contract, event: "StagingConfigSet", logs: logs, sub: sub}, nil
}

func (_Configurator *ConfiguratorFilterer) WatchStagingConfigSet(opts *bind.WatchOpts, sink chan<- *ConfiguratorStagingConfigSet, configId [][32]byte) (event.Subscription, error) {

	var configIdRule []interface{}
	for _, configIdItem := range configId {
		configIdRule = append(configIdRule, configIdItem)
	}

	logs, sub, err := _Configurator.contract.WatchLogs(opts, "StagingConfigSet", configIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ConfiguratorStagingConfigSet)
				if err := _Configurator.contract.UnpackLog(event, "StagingConfigSet", log); err != nil {
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

func (_Configurator *ConfiguratorFilterer) ParseStagingConfigSet(log types.Log) (*ConfiguratorStagingConfigSet, error) {
	event := new(ConfiguratorStagingConfigSet)
	if err := _Configurator.contract.UnpackLog(event, "StagingConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Configurator *Configurator) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Configurator.abi.Events["OwnershipTransferRequested"].ID:
		return _Configurator.ParseOwnershipTransferRequested(log)
	case _Configurator.abi.Events["OwnershipTransferred"].ID:
		return _Configurator.ParseOwnershipTransferred(log)
	case _Configurator.abi.Events["ProductionConfigSet"].ID:
		return _Configurator.ParseProductionConfigSet(log)
	case _Configurator.abi.Events["PromoteStagingConfig"].ID:
		return _Configurator.ParsePromoteStagingConfig(log)
	case _Configurator.abi.Events["StagingConfigSet"].ID:
		return _Configurator.ParseStagingConfigSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ConfiguratorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ConfiguratorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ConfiguratorProductionConfigSet) Topic() common.Hash {
	return common.HexToHash("0x261b20c2ecd99d86d6e936279e4f78db34603a3de3a4a84d6f3d4e0dd55e2478")
}

func (ConfiguratorPromoteStagingConfig) Topic() common.Hash {
	return common.HexToHash("0x1062aa08ac6046a0e69e3eafdf12d1eba63a67b71a874623e86eb06348a1d84f")
}

func (ConfiguratorStagingConfigSet) Topic() common.Hash {
	return common.HexToHash("0xef1b5f9d1b927b0fe871b12c7e7846457602d67b2bc36b0bc95feaf480e89056")
}

func (_Configurator *Configurator) Address() common.Address {
	return _Configurator.address
}

type ConfiguratorInterface interface {
	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	PromoteStagingConfig(opts *bind.TransactOpts, configId [32]byte, isGreenProduction bool) (*types.Transaction, error)

	SetProductionConfig(opts *bind.TransactOpts, configId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	SetStagingConfig(opts *bind.TransactOpts, configId [32]byte, signers [][]byte, offchainTransmitters [][32]byte, f uint8, onchainConfig []byte, offchainConfigVersion uint64, offchainConfig []byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConfiguratorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ConfiguratorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ConfiguratorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConfiguratorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConfiguratorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ConfiguratorOwnershipTransferred, error)

	FilterProductionConfigSet(opts *bind.FilterOpts, configId [][32]byte) (*ConfiguratorProductionConfigSetIterator, error)

	WatchProductionConfigSet(opts *bind.WatchOpts, sink chan<- *ConfiguratorProductionConfigSet, configId [][32]byte) (event.Subscription, error)

	ParseProductionConfigSet(log types.Log) (*ConfiguratorProductionConfigSet, error)

	FilterPromoteStagingConfig(opts *bind.FilterOpts, configId [][32]byte, retiredConfigDigest [][32]byte) (*ConfiguratorPromoteStagingConfigIterator, error)

	WatchPromoteStagingConfig(opts *bind.WatchOpts, sink chan<- *ConfiguratorPromoteStagingConfig, configId [][32]byte, retiredConfigDigest [][32]byte) (event.Subscription, error)

	ParsePromoteStagingConfig(log types.Log) (*ConfiguratorPromoteStagingConfig, error)

	FilterStagingConfigSet(opts *bind.FilterOpts, configId [][32]byte) (*ConfiguratorStagingConfigSetIterator, error)

	WatchStagingConfigSet(opts *bind.WatchOpts, sink chan<- *ConfiguratorStagingConfigSet, configId [][32]byte) (event.Subscription, error)

	ParseStagingConfigSet(log types.Log) (*ConfiguratorStagingConfigSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
