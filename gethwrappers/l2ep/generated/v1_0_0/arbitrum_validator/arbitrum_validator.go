// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_validator

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

type ArbitrumValidatorGasConfig struct {
	MaxGas             *big.Int
	GasPriceBid        *big.Int
	BaseFee            *big.Int
	GasPriceL1FeedAddr common.Address
}

var ArbitrumValidatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"crossDomainMessengerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l2ArbitrumSequencerUptimeFeedAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configACAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPriceBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPriceL1FeedAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_paymentStrategy\",\"type\":\"uint8\",\"internalType\":\"enum ArbitrumValidator.PaymentStrategy\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"CROSS_DOMAIN_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L2_ALIAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L2_SEQ_STATUS_RECORDER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configAC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gasConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"struct ArbitrumValidator.GasConfig\",\"components\":[{\"name\":\"maxGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPriceBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPriceL1FeedAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paymentStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enum ArbitrumValidator.PaymentStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setConfigAC\",\"inputs\":[{\"name\":\"accessController\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasConfig\",\"inputs\":[{\"name\":\"maxGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPriceBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"baseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasPriceL1FeedAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentStrategy\",\"inputs\":[{\"name\":\"_paymentStrategy\",\"type\":\"uint8\",\"internalType\":\"enum ArbitrumValidator.PaymentStrategy\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"validate\",\"inputs\":[{\"name\":\"previousRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"previousAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"currentRoundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentAnswer\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFunds\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFundsFromL2\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"refundAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFundsTo\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AddedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigACSet\",\"inputs\":[{\"name\":\"previous\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasConfigSet\",\"inputs\":[{\"name\":\"maxGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasPriceBid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasPriceL1FeedAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L2WithdrawalRequested\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"refundAddr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PaymentStrategySet\",\"inputs\":[{\"name\":\"paymentStrategy\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enum ArbitrumValidator.PaymentStrategy\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatedStatus\",\"inputs\":[{\"name\":\"previousRoundId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"previousAnswer\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"currentRoundId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"currentAnswer\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false}]",
	Bin: "0x60e060405273111100000000000000000000000000000000111130016001600160a01b031660c05234801562000033575f80fd5b506040516200224a3803806200224a833981016040819052620000569162000532565b33805f81620000ac5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000de57620000de8162000219565b50506001805460ff60a01b1916600160a01b179055506001600160a01b038816620001565760405162461bcd60e51b815260206004820152602160248201527f496e76616c69642078446f6d61696e204d657373656e676572206164647265736044820152607360f81b6064820152608401620000a3565b6001600160a01b038716620001d45760405162461bcd60e51b815260206004820152603460248201527f496e76616c696420417262697472756d53657175656e636572557074696d654660448201527f65656420636f6e747261637420616464726573730000000000000000000000006064820152608401620000a3565b6001600160a01b03808916608052871660a052620001f286620002c3565b62000200858585856200032c565b6200020b81620004b0565b5050505050505050620005d4565b336001600160a01b03821603620002735760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a3565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6008546001600160a01b0390811690821681146200032857600880546001600160a01b0319166001600160a01b0384811691821790925560405190918316907f6b0ce63879b19fa36f93da30a36f44140d8c3c727b6c75a16165c2017037d2cc905f90a35b5050565b5f84116200036f5760405162461bcd60e51b815260206004820152600f60248201526e4d617820676173206973207a65726f60881b6044820152606401620000a3565b5f8311620003c05760405162461bcd60e51b815260206004820152601560248201527f47617320707269636520626964206973207a65726f00000000000000000000006044820152606401620000a3565b6001600160a01b038116620004245760405162461bcd60e51b8152602060048201526024808201527f4761732070726963652041676772656761746f72206973207a65726f206164646044820152637265737360e01b6064820152608401620000a3565b6040805160808101825285815260208082018690528183018590526001600160a01b0384166060909201829052600487905560058690556006859055600780546001600160a01b03191683179055825187815290810186905290917f35674f8e28e701bef8f072d1034d588998f6966a59806b4299f6749c86b269e3910160405180910390a250505050565b6003805482919060ff191660018381811115620004d157620004d1620005c0565b0217905550806001811115620004eb57620004eb620005c0565b6040517fcc19f6868f2a8f851b1b59973487a53f0ac827f88698345b1529ef7c77b22bb5905f90a250565b80516001600160a01b03811681146200052d575f80fd5b919050565b5f805f805f805f80610100898b0312156200054b575f80fd5b620005568962000516565b97506200056660208a0162000516565b96506200057660408a0162000516565b9550606089015194506080890151935060a089015192506200059b60c08a0162000516565b915060e089015160028110620005af575f80fd5b809150509295985092959890939650565b634e487b7160e01b5f52602160045260245ffd5b60805160a05160c051611c28620006225f395f81816102b20152610f6401525f81816104c8015261100c01525f81816101cc015281816106a001528181610fcf01526113b60152611c285ff3fe60806040526004361061017b575f3560e01c80636b14daf8116100d1578063a118f2491161007c578063bfb9d1ab11610057578063bfb9d1ab146104b7578063dc7f0124146104ea578063f2fde38b1461051b575f80fd5b8063a118f2491461045a578063a2e6e0eb14610479578063beed9b5114610498575f80fd5b80638823da6c116100ac5780638823da6c146103f35780638a67ad0f146104125780638da5cb5b14610431575f80fd5b80636b14daf81461039c57806379ba5097146103cb5780638038e4a1146103df575f80fd5b8063276bf08f11610131578063490a44321161010c578063490a4432146103015780635c178f371461035e5780635e521eca1461037d575f80fd5b8063276bf08f146102775780633e48f015146102a157806344c2fa0c146102d4575f80fd5b806316be836c1161016157806316be836c146101bb578063181f5a771461021857806324600fc314610263575f80fd5b806302163e6c146101865780630a756983146101a7575f80fd5b3661018257005b5f80fd5b348015610191575f80fd5b506101a56101a0366004611801565b61053a565b005b3480156101b2575f80fd5b506101a561054e565b3480156101c6575f80fd5b506101ee7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b348015610223575f80fd5b50604080518082018252601781527f417262697472756d56616c696461746f7220322e302e300000000000000000006020820152905161020f919061187d565b34801561026e575f80fd5b506101a56105cc565b348015610282575f80fd5b5060085473ffffffffffffffffffffffffffffffffffffffff166101ee565b3480156102ac575f80fd5b506101ee7f000000000000000000000000000000000000000000000000000000000000000081565b3480156102df575f80fd5b506102f36102ee36600461188f565b6105e4565b60405190815260200161020f565b34801561030c575f80fd5b506103156107a7565b60405161020f919081518152602080830151908201526040808301519082015260609182015173ffffffffffffffffffffffffffffffffffffffff169181019190915260800190565b348015610369575f80fd5b506101a56103783660046118bd565b610827565b348015610388575f80fd5b5060035460ff1660405161020f9190611928565b3480156103a7575f80fd5b506103bb6103b6366004611994565b61097f565b604051901515815260200161020f565b3480156103d6575f80fd5b506101a56109d4565b3480156103ea575f80fd5b506101a5610ad0565b3480156103fe575f80fd5b506101a561040d366004611801565b610b63565b34801561041d575f80fd5b506101a561042c366004611a70565b610c1a565b34801561043c575f80fd5b505f5473ffffffffffffffffffffffffffffffffffffffff166101ee565b348015610465575f80fd5b506101a5610474366004611801565b610d64565b348015610484575f80fd5b506101a5610493366004611801565b610e16565b3480156104a3575f80fd5b506103bb6104b2366004611a8e565b610e29565b3480156104c2575f80fd5b506101ee7f000000000000000000000000000000000000000000000000000000000000000081565b3480156104f5575f80fd5b506001546103bb9074010000000000000000000000000000000000000000900460ff1681565b348015610526575f80fd5b506101a5610535366004611801565b6110f5565b610542611106565b61054b81611186565b50565b610556611106565b60015474010000000000000000000000000000000000000000900460ff16156105ca57600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638905f90a15b565b6105d4611106565b33476105e0828261121f565b5050565b5f6105ed611106565b60408051306024808301919091528251808303909101815260449091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f25e160630000000000000000000000000000000000000000000000000000000017905280515f906106629061137a565b6005549091506201d4c0905f8060035460ff166001811115610686576106866118fb565b14610691575f61069c565b61069c848484611434565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631b871c8d8260648b888c8d8a8a8e6040518a63ffffffff1660e01b8152600401610707989796959493929190611abd565b60206040518083038185885af1158015610723573d5f803e3d5ffd5b50505050506040513d601f19601f820116820180604052508101906107489190611b25565b95508673ffffffffffffffffffffffffffffffffffffffff16867f06f76b16d832d9e442e96306c36f3f2a819b64bd28441aa14fef67308a95c7168a60405161079391815260200190565b60405180910390a350505050505b92915050565b6107e460405180608001604052805f81526020015f81526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5060408051608081018252600454815260055460208201526006549181019190915260075473ffffffffffffffffffffffffffffffffffffffff16606082015290565b5f5473ffffffffffffffffffffffffffffffffffffffff16331480610902575060085473ffffffffffffffffffffffffffffffffffffffff161580159061090257506008546040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690636b14daf8906108c39033905f903690600401611b3c565b602060405180830381865afa1580156108de573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109029190611ba5565b61096d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f20616363657373000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b61097984848484611449565b50505050565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526002602052604081205460ff16806109cd575060015474010000000000000000000000000000000000000000900460ff16155b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a55576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610964565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610ad8611106565b60015474010000000000000000000000000000000000000000900460ff166105ca57600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517faebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480905f90a1565b610b6b611106565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff161561054b5773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527f3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d191015b60405180910390a150565b5f5473ffffffffffffffffffffffffffffffffffffffff16331480610cf5575060085473ffffffffffffffffffffffffffffffffffffffff1615801590610cf557506008546040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690636b14daf890610cb69033905f903690600401611b3c565b602060405180830381865afa158015610cd1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cf59190611ba5565b610d5b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610964565b61054b8161166e565b610d6c611106565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff1661054b5773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527f87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db49101610c0f565b610e1e611106565b476105e0828261121f565b5f610e69335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061097f92505050565b610ecf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f2061636365737300000000000000000000000000000000000000000000006044820152606401610964565b818403610ede575060016110ed565b60408051600184146024820181905267ffffffffffffffff429081166044808501919091528451808503909101815260649093019093526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fed8378f50000000000000000000000000000000000000000000000000000000090811790915282517f0000000000000000000000000000000000000000000000000000000000000000949193905f90610f929061137a565b600454600554919250905f8060035460ff166001811115610fb557610fb56118fb565b14610fc0575f610fcb565b610fcb848484611434565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631b871c8d827f00000000000000000000000000000000000000000000000000000000000000005f888e8f8a8a8e6040518a63ffffffff1660e01b8152600401611055989796959493929190611abd565b60206040518083038185885af1158015611071573d5f803e3d5ffd5b50505050506040513d601f19601f820116820180604052508101906110969190611b25565b50604080518f8152602081018f90529081018d9052606081018c90527fdf570e5bda6b7c0c0415ceeafd5f9321908b69023cae41e030a6df91cad90b039060800160405180910390a1600199505050505050505050505b949350505050565b6110fd611106565b61054b816116ec565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146105ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610964565b60085473ffffffffffffffffffffffffffffffffffffffff90811690821681146105e057600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691821790925560405190918316907f6b0ce63879b19fa36f93da30a36f44140d8c3c727b6c75a16165c2017037d2cc905f90a35050565b80471015611289576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610964565b5f8273ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f81146112df576040519150601f19603f3d011682016040523d82523d5f602084013e6112e4565b606091505b5050905080611375576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610964565b505050565b6006546040517fa66b327d0000000000000000000000000000000000000000000000000000000081526004810183905260248101919091525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a66b327d90604401602060405180830381865afa158015611410573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107a19190611b25565b5f61143f8284611bf1565b6110ed9085611c08565b5f84116114b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f4d617820676173206973207a65726f00000000000000000000000000000000006044820152606401610964565b5f831161151b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f47617320707269636520626964206973207a65726f00000000000000000000006044820152606401610964565b73ffffffffffffffffffffffffffffffffffffffff81166115bd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4761732070726963652041676772656761746f72206973207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610964565b60408051608081018252858152602080820186905281830185905273ffffffffffffffffffffffffffffffffffffffff84166060909201829052600487905560058690556006859055600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001683179055825187815290810186905290917f35674f8e28e701bef8f072d1034d588998f6966a59806b4299f6749c86b269e3910160405180910390a250505050565b600380548291907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183818111156116aa576116aa6118fb565b02179055508060018111156116c1576116c16118fb565b6040517fcc19f6868f2a8f851b1b59973487a53f0ac827f88698345b1529ef7c77b22bb5905f90a250565b3373ffffffffffffffffffffffffffffffffffffffff82160361176b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610964565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b73ffffffffffffffffffffffffffffffffffffffff8116811461054b575f80fd5b5f60208284031215611811575f80fd5b81356109cd816117e0565b5f81518084525f5b8181101561184057602081850181015186830182015201611824565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f6109cd602083018461181c565b5f80604083850312156118a0575f80fd5b8235915060208301356118b2816117e0565b809150509250929050565b5f805f80608085870312156118d0575f80fd5b84359350602085013592506040850135915060608501356118f0816117e0565b939692955090935050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6020810160028310611961577f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b91905290565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f80604083850312156119a5575f80fd5b82356119b0816117e0565b9150602083013567ffffffffffffffff808211156119cc575f80fd5b818501915085601f8301126119df575f80fd5b8135818111156119f1576119f1611967565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611a3757611a37611967565b81604052828152886020848701011115611a4f575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215611a80575f80fd5b8135600281106109cd575f80fd5b5f805f8060808587031215611aa1575f80fd5b5050823594602084013594506040840135936060013592509050565b5f61010073ffffffffffffffffffffffffffffffffffffffff808c1684528a602085015289604085015280891660608501528088166080850152508560a08401528460c08401528060e0840152611b168184018561181c565b9b9a5050505050505050505050565b5f60208284031215611b35575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff8416815260406020820152816040820152818360608301375f818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b5f60208284031215611bb5575f80fd5b815180151581146109cd575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820281158282048414176107a1576107a1611bc4565b808201808211156107a1576107a1611bc456fea164736f6c6343000818000a",
}

var ArbitrumValidatorABI = ArbitrumValidatorMetaData.ABI

var ArbitrumValidatorBin = ArbitrumValidatorMetaData.Bin

func DeployArbitrumValidator(auth *bind.TransactOpts, backend bind.ContractBackend, crossDomainMessengerAddr common.Address, l2ArbitrumSequencerUptimeFeedAddr common.Address, configACAddr common.Address, maxGas *big.Int, gasPriceBid *big.Int, baseFee *big.Int, gasPriceL1FeedAddr common.Address, _paymentStrategy uint8) (common.Address, *types.Transaction, *ArbitrumValidator, error) {
	parsed, err := ArbitrumValidatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumValidatorBin), backend, crossDomainMessengerAddr, l2ArbitrumSequencerUptimeFeedAddr, configACAddr, maxGas, gasPriceBid, baseFee, gasPriceL1FeedAddr, _paymentStrategy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumValidator{address: address, abi: *parsed, ArbitrumValidatorCaller: ArbitrumValidatorCaller{contract: contract}, ArbitrumValidatorTransactor: ArbitrumValidatorTransactor{contract: contract}, ArbitrumValidatorFilterer: ArbitrumValidatorFilterer{contract: contract}}, nil
}

type ArbitrumValidator struct {
	address common.Address
	abi     abi.ABI
	ArbitrumValidatorCaller
	ArbitrumValidatorTransactor
	ArbitrumValidatorFilterer
}

type ArbitrumValidatorCaller struct {
	contract *bind.BoundContract
}

type ArbitrumValidatorTransactor struct {
	contract *bind.BoundContract
}

type ArbitrumValidatorFilterer struct {
	contract *bind.BoundContract
}

type ArbitrumValidatorSession struct {
	Contract     *ArbitrumValidator
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbitrumValidatorCallerSession struct {
	Contract *ArbitrumValidatorCaller
	CallOpts bind.CallOpts
}

type ArbitrumValidatorTransactorSession struct {
	Contract     *ArbitrumValidatorTransactor
	TransactOpts bind.TransactOpts
}

type ArbitrumValidatorRaw struct {
	Contract *ArbitrumValidator
}

type ArbitrumValidatorCallerRaw struct {
	Contract *ArbitrumValidatorCaller
}

type ArbitrumValidatorTransactorRaw struct {
	Contract *ArbitrumValidatorTransactor
}

func NewArbitrumValidator(address common.Address, backend bind.ContractBackend) (*ArbitrumValidator, error) {
	abi, err := abi.JSON(strings.NewReader(ArbitrumValidatorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindArbitrumValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidator{address: address, abi: abi, ArbitrumValidatorCaller: ArbitrumValidatorCaller{contract: contract}, ArbitrumValidatorTransactor: ArbitrumValidatorTransactor{contract: contract}, ArbitrumValidatorFilterer: ArbitrumValidatorFilterer{contract: contract}}, nil
}

func NewArbitrumValidatorCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumValidatorCaller, error) {
	contract, err := bindArbitrumValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorCaller{contract: contract}, nil
}

func NewArbitrumValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumValidatorTransactor, error) {
	contract, err := bindArbitrumValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorTransactor{contract: contract}, nil
}

func NewArbitrumValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumValidatorFilterer, error) {
	contract, err := bindArbitrumValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorFilterer{contract: contract}, nil
}

func bindArbitrumValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrumValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ArbitrumValidator *ArbitrumValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumValidator.Contract.ArbitrumValidatorCaller.contract.Call(opts, result, method, params...)
}

func (_ArbitrumValidator *ArbitrumValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.ArbitrumValidatorTransactor.contract.Transfer(opts)
}

func (_ArbitrumValidator *ArbitrumValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.ArbitrumValidatorTransactor.contract.Transact(opts, method, params...)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumValidator.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.contract.Transfer(opts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.contract.Transact(opts, method, params...)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) CROSSDOMAINMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "CROSS_DOMAIN_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) CROSSDOMAINMESSENGER() (common.Address, error) {
	return _ArbitrumValidator.Contract.CROSSDOMAINMESSENGER(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) CROSSDOMAINMESSENGER() (common.Address, error) {
	return _ArbitrumValidator.Contract.CROSSDOMAINMESSENGER(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) L2ALIAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "L2_ALIAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) L2ALIAS() (common.Address, error) {
	return _ArbitrumValidator.Contract.L2ALIAS(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) L2ALIAS() (common.Address, error) {
	return _ArbitrumValidator.Contract.L2ALIAS(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) L2SEQSTATUSRECORDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "L2_SEQ_STATUS_RECORDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) L2SEQSTATUSRECORDER() (common.Address, error) {
	return _ArbitrumValidator.Contract.L2SEQSTATUSRECORDER(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) L2SEQSTATUSRECORDER() (common.Address, error) {
	return _ArbitrumValidator.Contract.L2SEQSTATUSRECORDER(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "checkEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) CheckEnabled() (bool, error) {
	return _ArbitrumValidator.Contract.CheckEnabled(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) CheckEnabled() (bool, error) {
	return _ArbitrumValidator.Contract.CheckEnabled(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) ConfigAC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "configAC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) ConfigAC() (common.Address, error) {
	return _ArbitrumValidator.Contract.ConfigAC(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) ConfigAC() (common.Address, error) {
	return _ArbitrumValidator.Contract.ConfigAC(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) GasConfig(opts *bind.CallOpts) (ArbitrumValidatorGasConfig, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "gasConfig")

	if err != nil {
		return *new(ArbitrumValidatorGasConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ArbitrumValidatorGasConfig)).(*ArbitrumValidatorGasConfig)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) GasConfig() (ArbitrumValidatorGasConfig, error) {
	return _ArbitrumValidator.Contract.GasConfig(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) GasConfig() (ArbitrumValidatorGasConfig, error) {
	return _ArbitrumValidator.Contract.GasConfig(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) HasAccess(opts *bind.CallOpts, _user common.Address, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "hasAccess", _user, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) HasAccess(_user common.Address, arg1 []byte) (bool, error) {
	return _ArbitrumValidator.Contract.HasAccess(&_ArbitrumValidator.CallOpts, _user, arg1)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) HasAccess(_user common.Address, arg1 []byte) (bool, error) {
	return _ArbitrumValidator.Contract.HasAccess(&_ArbitrumValidator.CallOpts, _user, arg1)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) Owner() (common.Address, error) {
	return _ArbitrumValidator.Contract.Owner(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) Owner() (common.Address, error) {
	return _ArbitrumValidator.Contract.Owner(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) PaymentStrategy(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "paymentStrategy")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) PaymentStrategy() (uint8, error) {
	return _ArbitrumValidator.Contract.PaymentStrategy(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) PaymentStrategy() (uint8, error) {
	return _ArbitrumValidator.Contract.PaymentStrategy(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArbitrumValidator.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ArbitrumValidator *ArbitrumValidatorSession) TypeAndVersion() (string, error) {
	return _ArbitrumValidator.Contract.TypeAndVersion(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorCallerSession) TypeAndVersion() (string, error) {
	return _ArbitrumValidator.Contract.TypeAndVersion(&_ArbitrumValidator.CallOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "acceptOwnership")
}

func (_ArbitrumValidator *ArbitrumValidatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.AcceptOwnership(&_ArbitrumValidator.TransactOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.AcceptOwnership(&_ArbitrumValidator.TransactOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "addAccess", _user)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.AddAccess(&_ArbitrumValidator.TransactOpts, _user)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.AddAccess(&_ArbitrumValidator.TransactOpts, _user)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "disableAccessCheck")
}

func (_ArbitrumValidator *ArbitrumValidatorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.DisableAccessCheck(&_ArbitrumValidator.TransactOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.DisableAccessCheck(&_ArbitrumValidator.TransactOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "enableAccessCheck")
}

func (_ArbitrumValidator *ArbitrumValidatorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.EnableAccessCheck(&_ArbitrumValidator.TransactOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.EnableAccessCheck(&_ArbitrumValidator.TransactOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "removeAccess", _user)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.RemoveAccess(&_ArbitrumValidator.TransactOpts, _user)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.RemoveAccess(&_ArbitrumValidator.TransactOpts, _user)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) SetConfigAC(opts *bind.TransactOpts, accessController common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "setConfigAC", accessController)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) SetConfigAC(accessController common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.SetConfigAC(&_ArbitrumValidator.TransactOpts, accessController)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) SetConfigAC(accessController common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.SetConfigAC(&_ArbitrumValidator.TransactOpts, accessController)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) SetGasConfig(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, baseFee *big.Int, gasPriceL1FeedAddr common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "setGasConfig", maxGas, gasPriceBid, baseFee, gasPriceL1FeedAddr)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) SetGasConfig(maxGas *big.Int, gasPriceBid *big.Int, baseFee *big.Int, gasPriceL1FeedAddr common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.SetGasConfig(&_ArbitrumValidator.TransactOpts, maxGas, gasPriceBid, baseFee, gasPriceL1FeedAddr)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) SetGasConfig(maxGas *big.Int, gasPriceBid *big.Int, baseFee *big.Int, gasPriceL1FeedAddr common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.SetGasConfig(&_ArbitrumValidator.TransactOpts, maxGas, gasPriceBid, baseFee, gasPriceL1FeedAddr)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) SetPaymentStrategy(opts *bind.TransactOpts, _paymentStrategy uint8) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "setPaymentStrategy", _paymentStrategy)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) SetPaymentStrategy(_paymentStrategy uint8) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.SetPaymentStrategy(&_ArbitrumValidator.TransactOpts, _paymentStrategy)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) SetPaymentStrategy(_paymentStrategy uint8) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.SetPaymentStrategy(&_ArbitrumValidator.TransactOpts, _paymentStrategy)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "transferOwnership", to)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.TransferOwnership(&_ArbitrumValidator.TransactOpts, to)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.TransferOwnership(&_ArbitrumValidator.TransactOpts, to)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) Validate(opts *bind.TransactOpts, previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "validate", previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) Validate(previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.Validate(&_ArbitrumValidator.TransactOpts, previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) Validate(previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.Validate(&_ArbitrumValidator.TransactOpts, previousRoundId, previousAnswer, currentRoundId, currentAnswer)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "withdrawFunds")
}

func (_ArbitrumValidator *ArbitrumValidatorSession) WithdrawFunds() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.WithdrawFunds(&_ArbitrumValidator.TransactOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.WithdrawFunds(&_ArbitrumValidator.TransactOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) WithdrawFundsFromL2(opts *bind.TransactOpts, amount *big.Int, refundAddr common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "withdrawFundsFromL2", amount, refundAddr)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) WithdrawFundsFromL2(amount *big.Int, refundAddr common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.WithdrawFundsFromL2(&_ArbitrumValidator.TransactOpts, amount, refundAddr)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) WithdrawFundsFromL2(amount *big.Int, refundAddr common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.WithdrawFundsFromL2(&_ArbitrumValidator.TransactOpts, amount, refundAddr)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) WithdrawFundsTo(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.Transact(opts, "withdrawFundsTo", recipient)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) WithdrawFundsTo(recipient common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.WithdrawFundsTo(&_ArbitrumValidator.TransactOpts, recipient)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) WithdrawFundsTo(recipient common.Address) (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.WithdrawFundsTo(&_ArbitrumValidator.TransactOpts, recipient)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumValidator.contract.RawTransact(opts, nil)
}

func (_ArbitrumValidator *ArbitrumValidatorSession) Receive() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.Receive(&_ArbitrumValidator.TransactOpts)
}

func (_ArbitrumValidator *ArbitrumValidatorTransactorSession) Receive() (*types.Transaction, error) {
	return _ArbitrumValidator.Contract.Receive(&_ArbitrumValidator.TransactOpts)
}

type ArbitrumValidatorAddedAccessIterator struct {
	Event *ArbitrumValidatorAddedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorAddedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorAddedAccess)
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
		it.Event = new(ArbitrumValidatorAddedAccess)
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

func (it *ArbitrumValidatorAddedAccessIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorAddedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*ArbitrumValidatorAddedAccessIterator, error) {

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorAddedAccessIterator{contract: _ArbitrumValidator.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorAddedAccess) (event.Subscription, error) {

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorAddedAccess)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseAddedAccess(log types.Log) (*ArbitrumValidatorAddedAccess, error) {
	event := new(ArbitrumValidatorAddedAccess)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorCheckAccessDisabledIterator struct {
	Event *ArbitrumValidatorCheckAccessDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorCheckAccessDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorCheckAccessDisabled)
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
		it.Event = new(ArbitrumValidatorCheckAccessDisabled)
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

func (it *ArbitrumValidatorCheckAccessDisabledIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorCheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorCheckAccessDisabled struct {
	Raw types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*ArbitrumValidatorCheckAccessDisabledIterator, error) {

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorCheckAccessDisabledIterator{contract: _ArbitrumValidator.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorCheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorCheckAccessDisabled)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseCheckAccessDisabled(log types.Log) (*ArbitrumValidatorCheckAccessDisabled, error) {
	event := new(ArbitrumValidatorCheckAccessDisabled)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorCheckAccessEnabledIterator struct {
	Event *ArbitrumValidatorCheckAccessEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorCheckAccessEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorCheckAccessEnabled)
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
		it.Event = new(ArbitrumValidatorCheckAccessEnabled)
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

func (it *ArbitrumValidatorCheckAccessEnabledIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorCheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorCheckAccessEnabled struct {
	Raw types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*ArbitrumValidatorCheckAccessEnabledIterator, error) {

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorCheckAccessEnabledIterator{contract: _ArbitrumValidator.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorCheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorCheckAccessEnabled)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseCheckAccessEnabled(log types.Log) (*ArbitrumValidatorCheckAccessEnabled, error) {
	event := new(ArbitrumValidatorCheckAccessEnabled)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorConfigACSetIterator struct {
	Event *ArbitrumValidatorConfigACSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorConfigACSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorConfigACSet)
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
		it.Event = new(ArbitrumValidatorConfigACSet)
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

func (it *ArbitrumValidatorConfigACSetIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorConfigACSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorConfigACSet struct {
	Previous common.Address
	Current  common.Address
	Raw      types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterConfigACSet(opts *bind.FilterOpts, previous []common.Address, current []common.Address) (*ArbitrumValidatorConfigACSetIterator, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "ConfigACSet", previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorConfigACSetIterator{contract: _ArbitrumValidator.contract, event: "ConfigACSet", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchConfigACSet(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorConfigACSet, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "ConfigACSet", previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorConfigACSet)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "ConfigACSet", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseConfigACSet(log types.Log) (*ArbitrumValidatorConfigACSet, error) {
	event := new(ArbitrumValidatorConfigACSet)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "ConfigACSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorGasConfigSetIterator struct {
	Event *ArbitrumValidatorGasConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorGasConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorGasConfigSet)
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
		it.Event = new(ArbitrumValidatorGasConfigSet)
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

func (it *ArbitrumValidatorGasConfigSetIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorGasConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorGasConfigSet struct {
	MaxGas             *big.Int
	GasPriceBid        *big.Int
	GasPriceL1FeedAddr common.Address
	Raw                types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterGasConfigSet(opts *bind.FilterOpts, gasPriceL1FeedAddr []common.Address) (*ArbitrumValidatorGasConfigSetIterator, error) {

	var gasPriceL1FeedAddrRule []interface{}
	for _, gasPriceL1FeedAddrItem := range gasPriceL1FeedAddr {
		gasPriceL1FeedAddrRule = append(gasPriceL1FeedAddrRule, gasPriceL1FeedAddrItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "GasConfigSet", gasPriceL1FeedAddrRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorGasConfigSetIterator{contract: _ArbitrumValidator.contract, event: "GasConfigSet", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchGasConfigSet(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorGasConfigSet, gasPriceL1FeedAddr []common.Address) (event.Subscription, error) {

	var gasPriceL1FeedAddrRule []interface{}
	for _, gasPriceL1FeedAddrItem := range gasPriceL1FeedAddr {
		gasPriceL1FeedAddrRule = append(gasPriceL1FeedAddrRule, gasPriceL1FeedAddrItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "GasConfigSet", gasPriceL1FeedAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorGasConfigSet)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "GasConfigSet", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseGasConfigSet(log types.Log) (*ArbitrumValidatorGasConfigSet, error) {
	event := new(ArbitrumValidatorGasConfigSet)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "GasConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorL2WithdrawalRequestedIterator struct {
	Event *ArbitrumValidatorL2WithdrawalRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorL2WithdrawalRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorL2WithdrawalRequested)
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
		it.Event = new(ArbitrumValidatorL2WithdrawalRequested)
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

func (it *ArbitrumValidatorL2WithdrawalRequestedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorL2WithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorL2WithdrawalRequested struct {
	Id         *big.Int
	Amount     *big.Int
	RefundAddr common.Address
	Raw        types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterL2WithdrawalRequested(opts *bind.FilterOpts, id []*big.Int, refundAddr []common.Address) (*ArbitrumValidatorL2WithdrawalRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var refundAddrRule []interface{}
	for _, refundAddrItem := range refundAddr {
		refundAddrRule = append(refundAddrRule, refundAddrItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "L2WithdrawalRequested", idRule, refundAddrRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorL2WithdrawalRequestedIterator{contract: _ArbitrumValidator.contract, event: "L2WithdrawalRequested", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchL2WithdrawalRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorL2WithdrawalRequested, id []*big.Int, refundAddr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var refundAddrRule []interface{}
	for _, refundAddrItem := range refundAddr {
		refundAddrRule = append(refundAddrRule, refundAddrItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "L2WithdrawalRequested", idRule, refundAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorL2WithdrawalRequested)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "L2WithdrawalRequested", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseL2WithdrawalRequested(log types.Log) (*ArbitrumValidatorL2WithdrawalRequested, error) {
	event := new(ArbitrumValidatorL2WithdrawalRequested)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "L2WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorOwnershipTransferRequestedIterator struct {
	Event *ArbitrumValidatorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorOwnershipTransferRequested)
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
		it.Event = new(ArbitrumValidatorOwnershipTransferRequested)
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

func (it *ArbitrumValidatorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumValidatorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorOwnershipTransferRequestedIterator{contract: _ArbitrumValidator.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorOwnershipTransferRequested)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseOwnershipTransferRequested(log types.Log) (*ArbitrumValidatorOwnershipTransferRequested, error) {
	event := new(ArbitrumValidatorOwnershipTransferRequested)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorOwnershipTransferredIterator struct {
	Event *ArbitrumValidatorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorOwnershipTransferred)
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
		it.Event = new(ArbitrumValidatorOwnershipTransferred)
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

func (it *ArbitrumValidatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumValidatorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorOwnershipTransferredIterator{contract: _ArbitrumValidator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorOwnershipTransferred)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseOwnershipTransferred(log types.Log) (*ArbitrumValidatorOwnershipTransferred, error) {
	event := new(ArbitrumValidatorOwnershipTransferred)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorPaymentStrategySetIterator struct {
	Event *ArbitrumValidatorPaymentStrategySet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorPaymentStrategySetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorPaymentStrategySet)
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
		it.Event = new(ArbitrumValidatorPaymentStrategySet)
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

func (it *ArbitrumValidatorPaymentStrategySetIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorPaymentStrategySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorPaymentStrategySet struct {
	PaymentStrategy uint8
	Raw             types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterPaymentStrategySet(opts *bind.FilterOpts, paymentStrategy []uint8) (*ArbitrumValidatorPaymentStrategySetIterator, error) {

	var paymentStrategyRule []interface{}
	for _, paymentStrategyItem := range paymentStrategy {
		paymentStrategyRule = append(paymentStrategyRule, paymentStrategyItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "PaymentStrategySet", paymentStrategyRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorPaymentStrategySetIterator{contract: _ArbitrumValidator.contract, event: "PaymentStrategySet", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchPaymentStrategySet(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorPaymentStrategySet, paymentStrategy []uint8) (event.Subscription, error) {

	var paymentStrategyRule []interface{}
	for _, paymentStrategyItem := range paymentStrategy {
		paymentStrategyRule = append(paymentStrategyRule, paymentStrategyItem)
	}

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "PaymentStrategySet", paymentStrategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorPaymentStrategySet)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "PaymentStrategySet", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParsePaymentStrategySet(log types.Log) (*ArbitrumValidatorPaymentStrategySet, error) {
	event := new(ArbitrumValidatorPaymentStrategySet)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "PaymentStrategySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorRemovedAccessIterator struct {
	Event *ArbitrumValidatorRemovedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorRemovedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorRemovedAccess)
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
		it.Event = new(ArbitrumValidatorRemovedAccess)
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

func (it *ArbitrumValidatorRemovedAccessIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorRemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorRemovedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*ArbitrumValidatorRemovedAccessIterator, error) {

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorRemovedAccessIterator{contract: _ArbitrumValidator.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorRemovedAccess) (event.Subscription, error) {

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorRemovedAccess)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseRemovedAccess(log types.Log) (*ArbitrumValidatorRemovedAccess, error) {
	event := new(ArbitrumValidatorRemovedAccess)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumValidatorValidatedStatusIterator struct {
	Event *ArbitrumValidatorValidatedStatus

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumValidatorValidatedStatusIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumValidatorValidatedStatus)
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
		it.Event = new(ArbitrumValidatorValidatedStatus)
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

func (it *ArbitrumValidatorValidatedStatusIterator) Error() error {
	return it.fail
}

func (it *ArbitrumValidatorValidatedStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumValidatorValidatedStatus struct {
	PreviousRoundId *big.Int
	PreviousAnswer  *big.Int
	CurrentRoundId  *big.Int
	CurrentAnswer   *big.Int
	Raw             types.Log
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) FilterValidatedStatus(opts *bind.FilterOpts) (*ArbitrumValidatorValidatedStatusIterator, error) {

	logs, sub, err := _ArbitrumValidator.contract.FilterLogs(opts, "ValidatedStatus")
	if err != nil {
		return nil, err
	}
	return &ArbitrumValidatorValidatedStatusIterator{contract: _ArbitrumValidator.contract, event: "ValidatedStatus", logs: logs, sub: sub}, nil
}

func (_ArbitrumValidator *ArbitrumValidatorFilterer) WatchValidatedStatus(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorValidatedStatus) (event.Subscription, error) {

	logs, sub, err := _ArbitrumValidator.contract.WatchLogs(opts, "ValidatedStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumValidatorValidatedStatus)
				if err := _ArbitrumValidator.contract.UnpackLog(event, "ValidatedStatus", log); err != nil {
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

func (_ArbitrumValidator *ArbitrumValidatorFilterer) ParseValidatedStatus(log types.Log) (*ArbitrumValidatorValidatedStatus, error) {
	event := new(ArbitrumValidatorValidatedStatus)
	if err := _ArbitrumValidator.contract.UnpackLog(event, "ValidatedStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (ArbitrumValidatorAddedAccess) Topic() common.Hash {
	return common.HexToHash("0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4")
}

func (ArbitrumValidatorCheckAccessDisabled) Topic() common.Hash {
	return common.HexToHash("0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638")
}

func (ArbitrumValidatorCheckAccessEnabled) Topic() common.Hash {
	return common.HexToHash("0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480")
}

func (ArbitrumValidatorConfigACSet) Topic() common.Hash {
	return common.HexToHash("0x6b0ce63879b19fa36f93da30a36f44140d8c3c727b6c75a16165c2017037d2cc")
}

func (ArbitrumValidatorGasConfigSet) Topic() common.Hash {
	return common.HexToHash("0x35674f8e28e701bef8f072d1034d588998f6966a59806b4299f6749c86b269e3")
}

func (ArbitrumValidatorL2WithdrawalRequested) Topic() common.Hash {
	return common.HexToHash("0x06f76b16d832d9e442e96306c36f3f2a819b64bd28441aa14fef67308a95c716")
}

func (ArbitrumValidatorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ArbitrumValidatorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ArbitrumValidatorPaymentStrategySet) Topic() common.Hash {
	return common.HexToHash("0xcc19f6868f2a8f851b1b59973487a53f0ac827f88698345b1529ef7c77b22bb5")
}

func (ArbitrumValidatorRemovedAccess) Topic() common.Hash {
	return common.HexToHash("0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1")
}

func (ArbitrumValidatorValidatedStatus) Topic() common.Hash {
	return common.HexToHash("0xdf570e5bda6b7c0c0415ceeafd5f9321908b69023cae41e030a6df91cad90b03")
}

func (_ArbitrumValidator *ArbitrumValidator) Address() common.Address {
	return _ArbitrumValidator.address
}

type ArbitrumValidatorInterface interface {
	CROSSDOMAINMESSENGER(opts *bind.CallOpts) (common.Address, error)

	L2ALIAS(opts *bind.CallOpts) (common.Address, error)

	L2SEQSTATUSRECORDER(opts *bind.CallOpts) (common.Address, error)

	CheckEnabled(opts *bind.CallOpts) (bool, error)

	ConfigAC(opts *bind.CallOpts) (common.Address, error)

	GasConfig(opts *bind.CallOpts) (ArbitrumValidatorGasConfig, error)

	HasAccess(opts *bind.CallOpts, _user common.Address, arg1 []byte) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PaymentStrategy(opts *bind.CallOpts) (uint8, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error)

	DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error)

	EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error)

	SetConfigAC(opts *bind.TransactOpts, accessController common.Address) (*types.Transaction, error)

	SetGasConfig(opts *bind.TransactOpts, maxGas *big.Int, gasPriceBid *big.Int, baseFee *big.Int, gasPriceL1FeedAddr common.Address) (*types.Transaction, error)

	SetPaymentStrategy(opts *bind.TransactOpts, _paymentStrategy uint8) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Validate(opts *bind.TransactOpts, previousRoundId *big.Int, previousAnswer *big.Int, currentRoundId *big.Int, currentAnswer *big.Int) (*types.Transaction, error)

	WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawFundsFromL2(opts *bind.TransactOpts, amount *big.Int, refundAddr common.Address) (*types.Transaction, error)

	WithdrawFundsTo(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAddedAccess(opts *bind.FilterOpts) (*ArbitrumValidatorAddedAccessIterator, error)

	WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorAddedAccess) (event.Subscription, error)

	ParseAddedAccess(log types.Log) (*ArbitrumValidatorAddedAccess, error)

	FilterCheckAccessDisabled(opts *bind.FilterOpts) (*ArbitrumValidatorCheckAccessDisabledIterator, error)

	WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorCheckAccessDisabled) (event.Subscription, error)

	ParseCheckAccessDisabled(log types.Log) (*ArbitrumValidatorCheckAccessDisabled, error)

	FilterCheckAccessEnabled(opts *bind.FilterOpts) (*ArbitrumValidatorCheckAccessEnabledIterator, error)

	WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorCheckAccessEnabled) (event.Subscription, error)

	ParseCheckAccessEnabled(log types.Log) (*ArbitrumValidatorCheckAccessEnabled, error)

	FilterConfigACSet(opts *bind.FilterOpts, previous []common.Address, current []common.Address) (*ArbitrumValidatorConfigACSetIterator, error)

	WatchConfigACSet(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorConfigACSet, previous []common.Address, current []common.Address) (event.Subscription, error)

	ParseConfigACSet(log types.Log) (*ArbitrumValidatorConfigACSet, error)

	FilterGasConfigSet(opts *bind.FilterOpts, gasPriceL1FeedAddr []common.Address) (*ArbitrumValidatorGasConfigSetIterator, error)

	WatchGasConfigSet(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorGasConfigSet, gasPriceL1FeedAddr []common.Address) (event.Subscription, error)

	ParseGasConfigSet(log types.Log) (*ArbitrumValidatorGasConfigSet, error)

	FilterL2WithdrawalRequested(opts *bind.FilterOpts, id []*big.Int, refundAddr []common.Address) (*ArbitrumValidatorL2WithdrawalRequestedIterator, error)

	WatchL2WithdrawalRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorL2WithdrawalRequested, id []*big.Int, refundAddr []common.Address) (event.Subscription, error)

	ParseL2WithdrawalRequested(log types.Log) (*ArbitrumValidatorL2WithdrawalRequested, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumValidatorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ArbitrumValidatorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumValidatorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ArbitrumValidatorOwnershipTransferred, error)

	FilterPaymentStrategySet(opts *bind.FilterOpts, paymentStrategy []uint8) (*ArbitrumValidatorPaymentStrategySetIterator, error)

	WatchPaymentStrategySet(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorPaymentStrategySet, paymentStrategy []uint8) (event.Subscription, error)

	ParsePaymentStrategySet(log types.Log) (*ArbitrumValidatorPaymentStrategySet, error)

	FilterRemovedAccess(opts *bind.FilterOpts) (*ArbitrumValidatorRemovedAccessIterator, error)

	WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorRemovedAccess) (event.Subscription, error)

	ParseRemovedAccess(log types.Log) (*ArbitrumValidatorRemovedAccess, error)

	FilterValidatedStatus(opts *bind.FilterOpts) (*ArbitrumValidatorValidatedStatusIterator, error)

	WatchValidatedStatus(opts *bind.WatchOpts, sink chan<- *ArbitrumValidatorValidatedStatus) (event.Subscription, error)

	ParseValidatedStatus(log types.Log) (*ArbitrumValidatorValidatedStatus, error)

	Address() common.Address
}
