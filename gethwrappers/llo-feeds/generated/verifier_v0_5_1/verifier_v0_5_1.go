// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_v0_5_1

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

type CommonAddressAndWeight struct {
	Addr   common.Address
	Weight uint64
}

var VerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"verifierProxyAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activateConfig\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deactivateConfig\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"latestConfigDetails\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recipientAddressesAndWeights\",\"type\":\"tuple[]\",\"internalType\":\"structCommon.AddressAndWeight[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"isVerifier\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateConfig\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prevSigners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"newSigners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"signedReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"verifierResponse\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyView\",\"inputs\":[{\"name\":\"signedReport\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifierResponse\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ConfigActivated\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigDeactivated\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"f\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConfigUpdated\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"prevSigners\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"newSigners\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportVerified\",\"inputs\":[{\"name\":\"feedId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"requester\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessForbidden\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadVerification\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConfigDigestAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DigestEmpty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DigestInactive\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"DigestNotSet\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ExcessSigners\",\"inputs\":[{\"name\":\"numSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"FaultToleranceMustBePositive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectSignatureCount\",\"inputs\":[{\"name\":\"numSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedNumSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientSigners\",\"inputs\":[{\"name\":\"numSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minSigners\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MismatchedSignatures\",\"inputs\":[{\"name\":\"rsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ssLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NonUniqueSignatures\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001f0b38038062001f0b8339810160408190526200003491620001a6565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000fb565b5050506001600160a01b038116620000e95760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0316608052620001d8565b336001600160a01b03821603620001555760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001b957600080fd5b81516001600160a01b0381168114620001d157600080fd5b9392505050565b608051611d10620001fb60003960008181610d8a0152610e550152611d106000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c806341e3df58116100815780638da5cb5b1161005b5780638da5cb5b14610201578063e84f128e14610229578063f2fde38b1461026457600080fd5b806341e3df58146101d357806351b34c90146101e657806379ba5097146101f957600080fd5b80630f672ef4116100b25780630f672ef41461016b578063181f5a771461017e5780633d3ac1b5146101c057600080fd5b806301ffc9a7146100d95780630d1d79af146101435780630e112e5414610158575b600080fd5b61012e6100e736600461137d565b7fffffffff00000000000000000000000000000000000000000000000000000000167f3d3ac1b5000000000000000000000000000000000000000000000000000000001490565b60405190151581526020015b60405180910390f35b6101566101513660046113c6565b610277565b005b610156610166366004611441565b61036f565b6101566101793660046113c6565b6106da565b60408051808201909152600e81527f566572696669657220322e302e3100000000000000000000000000000000000060208201525b60405161013a919061152e565b6101b36101ce3660046115a7565b6107c6565b6101566101e13660046116e9565b610830565b6101b36101f4366004611802565b6108f2565b610156610907565b60005460405173ffffffffffffffffffffffffffffffffffffffff909116815260200161013a565b61024f6102373660046113c6565b60009081526002602052604090205463ffffffff1690565b60405163ffffffff909116815260200161013a565b610156610272366004611844565b610a04565b61027f610a18565b6000818152600260205260409020816102c4576040517fe332262700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805465010000000000900460ff16600003610313576040517f74eb4b93000000000000000000000000000000000000000000000000000000008152600481018390526024015b60405180910390fd5b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff1664010000000017815560405182907fa543797a0501218bba8a3daf75a71c8df8d1a7f791f4e44d40e43b6450183cea90600090a25050565b8160ff821660008190036103af576040517f0743bae600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601f8211156103f4576040517f61750f4000000000000000000000000000000000000000000000000000000000815260048101839052601f602482015260440161030a565b6103ff81600361188e565b8211610457578161041182600361188e565b61041c9060016118a5565b6040517f9dd9e6d80000000000000000000000000000000000000000000000000000000081526004810192909252602482015260440161030a565b61045f610a18565b6000888152600260205260408120805490916501000000000090910460ff1690036104b9576040517f74eb4b93000000000000000000000000000000000000000000000000000000008152600481018a905260240161030a565b80546601000000000000900460ff168714610500576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8781101561063a5760008260010160008b8b85818110610525576105256118e7565b905060200201602081019061053a9190611844565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002054610100900460ff16600181111561057a5761057a6118b8565b036105b1576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160010160008a8a848181106105c9576105c96118e7565b90506020020160208101906105de9190611844565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905561063381611916565b9050610503565b5061069189878787600060405190808252806020026020018201604052801561068957816020015b60408051808201909152600080825260208201528152602001906001900390816106625790505b506001610a9b565b887fb0b75a854fab801413da6202fc07e875c54eaf371a1e3909fb2645364ba58616898989896040516106c794939291906119a2565b60405180910390a2505050505050505050565b6106e2610a18565b600081815260026020526040902081610727576040517fe332262700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805465010000000000900460ff16600003610771576040517f74eb4b930000000000000000000000000000000000000000000000000000000081526004810183905260240161030a565b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff16815560405182907f5bfaab86edc1b932e3c334327a591c9ded067cb521abae19b95ca927d607657990600090a25050565b606060006107d48585610e3b565b90506107df816119d4565b60405173ffffffffffffffffffffffffffffffffffffffff851681527f58ca9502e98a536e06e72d680fcc251e5d10b72291a281665a2c2dc0ac30fcc59060200160405180910390a2949350505050565b8260ff83166000819003610870576040517f0743bae600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601f8211156108b5576040517f61750f4000000000000000000000000000000000000000000000000000000000815260048101839052601f602482015260440161030a565b6108c081600361188e565b82116108d2578161041182600361188e565b6108da610a18565b6108e987878787876000610a9b565b50505050505050565b60606108fe8383610e3b565b90505b92915050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610988576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161030a565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610a0c610a18565b610a1581610f0c565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610a99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161030a565b565b6000868152600260205260409020805465010000000000900460ff1615801590610ac3575081155b15610afa576040517f961dba8800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805460ff8681166601000000000000027fffffffffffffffffffffffffffffffffffffffffffffffffff00ff00ffffffff91871665010000000000027fffffffffffffffffffffffffffffffffffffffffffffffffffff00ff0000000090931663ffffffff43161792909217161764010000000017815560005b60ff8116861115610d4757600087878360ff16818110610b9657610b966118e7565b9050602002016020810190610bab9190611844565b905073ffffffffffffffffffffffffffffffffffffffff8116610bfa576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008073ffffffffffffffffffffffffffffffffffffffff831660009081526001868101602052604090912054610100900460ff1690811115610c3f57610c3f6118b8565b1480159150610c7a576040517ff67bc7c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051808201825260ff85811682526001602080840182815273ffffffffffffffffffffffffffffffffffffffff881660009081528a84019092529490208351815493167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008416811782559451939490939284927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000090911690911790610100908490811115610d2c57610d2c6118b8565b0217905550905050505080610d4090611a19565b9050610b74565b50816108e9576040517fb011b24700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063b011b24790610dc4906000908b908890600401611a38565b600060405180830381600087803b158015610dde57600080fd5b505af1158015610df2573d6000803e3d6000fd5b50505050867f5b1f376eb2bda670fa39339616d0a73f45b61bec8faeba8ca834f2ebb49676e0878787604051610e2a93929190611ab8565b60405180910390a250505050505050565b60603373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610eac576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080808080610ebe87890189611bd3565b84516000818152600260205260409020959a509398509196509450925090610ee882868684611001565b85516020870120610efd818988888887611101565b50949998505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff821603610f8b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161030a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b805460009061101d9065010000000000900460ff166001611cae565b8254909150640100000000900460ff16611066576040517fd990d6210000000000000000000000000000000000000000000000000000000081526004810186905260240161030a565b8060ff168451146110b25783516040517f5348a282000000000000000000000000000000000000000000000000000000008152600481019190915260ff8216602482015260440161030a565b82518451146110fa57835183516040517ff0d314080000000000000000000000000000000000000000000000000000000081526004810192909252602482015260440161030a565b5050505050565b60008686604051602001611116929190611cc7565b604051602081830303815290604052805190602001209050600061114a604080518082019091526000808252602082015290565b8651600090815b818110156113155760018689836020811061116e5761116e6118e7565b61117b91901a601b611cae565b8c848151811061118d5761118d6118e7565b60200260200101518c85815181106111a7576111a76118e7565b6020026020010151604051600081526020016040526040516111e5949392919093845260ff9290921660208401526040830152606082015260800190565b6020604051602081039080840390855afa158015611207573d6000803e3d6000fd5b5050604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081015173ffffffffffffffffffffffffffffffffffffffff811660009081526001808d01602090815291859020848601909552845460ff80821686529399509395509085019261010090049091169081111561128c5761128c6118b8565b600181111561129d5761129d6118b8565b90525093506001846020015160018111156112ba576112ba6118b8565b146112f1576040517f4df18f0700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836000015160080260ff166001901b850194508061130e90611916565b9050611151565b50837e01010101010101010101010101010101010101010101010101010101010101851614611370576040517f4df18f0700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050505050505050565b60006020828403121561138f57600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113bf57600080fd5b9392505050565b6000602082840312156113d857600080fd5b5035919050565b60008083601f8401126113f157600080fd5b50813567ffffffffffffffff81111561140957600080fd5b6020830191508360208260051b850101111561142457600080fd5b9250929050565b803560ff8116811461143c57600080fd5b919050565b6000806000806000806080878903121561145a57600080fd5b86359550602087013567ffffffffffffffff8082111561147957600080fd5b6114858a838b016113df565b9097509550604089013591508082111561149e57600080fd5b506114ab89828a016113df565b90945092506114be90506060880161142b565b90509295509295509295565b6000815180845260005b818110156114f0576020818501810151868301820152016114d4565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006108fe60208301846114ca565b60008083601f84011261155357600080fd5b50813567ffffffffffffffff81111561156b57600080fd5b60208301915083602082850101111561142457600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461143c57600080fd5b6000806000604084860312156115bc57600080fd5b833567ffffffffffffffff8111156115d357600080fd5b6115df86828701611541565b90945092506115f2905060208501611583565b90509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561164d5761164d6115fb565b60405290565b6040516060810167ffffffffffffffff8111828210171561164d5761164d6115fb565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156116bd576116bd6115fb565b604052919050565b600067ffffffffffffffff8211156116df576116df6115fb565b5060051b60200190565b60008060008060006080868803121561170157600080fd5b8535945060208087013567ffffffffffffffff8082111561172157600080fd5b61172d8a838b016113df565b90975095506040915061174189830161142b565b945060608901358181111561175557600080fd5b8901601f81018b1361176657600080fd5b8035611779611774826116c5565b611676565b81815260069190911b8201850190858101908d83111561179857600080fd5b928601925b828410156117ee5785848f0312156117b55760008081fd5b6117bd61162a565b6117c685611583565b81528785013586811681146117db5760008081fd5b818901528252928501929086019061179d565b809750505050505050509295509295909350565b6000806020838503121561181557600080fd5b823567ffffffffffffffff81111561182c57600080fd5b61183885828601611541565b90969095509350505050565b60006020828403121561185657600080fd5b6108fe82611583565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820281158282048414176109015761090161185f565b808201808211156109015761090161185f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036119475761194761185f565b5060010190565b8183526000602080850194508260005b858110156119975773ffffffffffffffffffffffffffffffffffffffff61198483611583565b168752958201959082019060010161195e565b509495945050505050565b6040815260006119b660408301868861194e565b82810360208401526119c981858761194e565b979650505050505050565b80516020808301519190811015611a13577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8160200360031b1b821691505b50919050565b600060ff821660ff8103611a2f57611a2f61185f565b60010192915050565b600060608201858352602085818501526040606081860152828651808552608087019150838801945060005b81811015611aa9578551805173ffffffffffffffffffffffffffffffffffffffff16845285015167ffffffffffffffff16858401529484019491830191600101611a64565b50909998505050505050505050565b604081526000611acc60408301858761194e565b905060ff83166020830152949350505050565b600082601f830112611af057600080fd5b813567ffffffffffffffff811115611b0a57611b0a6115fb565b611b3b60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611676565b818152846020838601011115611b5057600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f830112611b7e57600080fd5b81356020611b8e611774836116c5565b82815260059290921b84018101918181019086841115611bad57600080fd5b8286015b84811015611bc85780358352918301918301611bb1565b509695505050505050565b600080600080600060e08688031215611beb57600080fd5b86601f870112611bfa57600080fd5b611c02611653565b806060880189811115611c1457600080fd5b885b81811015611c2e578035845260209384019301611c16565b5090965035905067ffffffffffffffff80821115611c4b57600080fd5b611c5789838a01611adf565b95506080880135915080821115611c6d57600080fd5b611c7989838a01611b6d565b945060a0880135915080821115611c8f57600080fd5b50611c9c88828901611b6d565b9598949750929560c001359392505050565b60ff81811683821601908111156109015761090161185f565b828152600060208083018460005b6003811015611cf257815183529183019190830190600101611cd5565b50505050608082019050939250505056fea164736f6c6343000813000a",
}

var VerifierABI = VerifierMetaData.ABI

var VerifierBin = VerifierMetaData.Bin

func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, verifierProxyAddr common.Address) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierBin), backend, verifierProxyAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{address: address, abi: *parsed, VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

type Verifier struct {
	address common.Address
	abi     abi.ABI
	VerifierCaller
	VerifierTransactor
	VerifierFilterer
}

type VerifierCaller struct {
	contract *bind.BoundContract
}

type VerifierTransactor struct {
	contract *bind.BoundContract
}

type VerifierFilterer struct {
	contract *bind.BoundContract
}

type VerifierSession struct {
	Contract     *Verifier
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierCallerSession struct {
	Contract *VerifierCaller
	CallOpts bind.CallOpts
}

type VerifierTransactorSession struct {
	Contract     *VerifierTransactor
	TransactOpts bind.TransactOpts
}

type VerifierRaw struct {
	Contract *Verifier
}

type VerifierCallerRaw struct {
	Contract *VerifierCaller
}

type VerifierTransactorRaw struct {
	Contract *VerifierTransactor
}

func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{address: address, abi: abi, VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

func (_Verifier *VerifierCaller) LatestConfigDetails(opts *bind.CallOpts, configDigest [32]byte) (uint32, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "latestConfigDetails", configDigest)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_Verifier *VerifierSession) LatestConfigDetails(configDigest [32]byte) (uint32, error) {
	return _Verifier.Contract.LatestConfigDetails(&_Verifier.CallOpts, configDigest)
}

func (_Verifier *VerifierCallerSession) LatestConfigDetails(configDigest [32]byte) (uint32, error) {
	return _Verifier.Contract.LatestConfigDetails(&_Verifier.CallOpts, configDigest)
}

func (_Verifier *VerifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Verifier *VerifierSession) Owner() (common.Address, error) {
	return _Verifier.Contract.Owner(&_Verifier.CallOpts)
}

func (_Verifier *VerifierCallerSession) Owner() (common.Address, error) {
	return _Verifier.Contract.Owner(&_Verifier.CallOpts)
}

func (_Verifier *VerifierCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Verifier *VerifierSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Verifier.Contract.SupportsInterface(&_Verifier.CallOpts, interfaceId)
}

func (_Verifier *VerifierCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Verifier.Contract.SupportsInterface(&_Verifier.CallOpts, interfaceId)
}

func (_Verifier *VerifierCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Verifier *VerifierSession) TypeAndVersion() (string, error) {
	return _Verifier.Contract.TypeAndVersion(&_Verifier.CallOpts)
}

func (_Verifier *VerifierCallerSession) TypeAndVersion() (string, error) {
	return _Verifier.Contract.TypeAndVersion(&_Verifier.CallOpts)
}

func (_Verifier *VerifierCaller) VerifyView(opts *bind.CallOpts, signedReport []byte) ([]byte, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyView", signedReport)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_Verifier *VerifierSession) VerifyView(signedReport []byte) ([]byte, error) {
	return _Verifier.Contract.VerifyView(&_Verifier.CallOpts, signedReport)
}

func (_Verifier *VerifierCallerSession) VerifyView(signedReport []byte) ([]byte, error) {
	return _Verifier.Contract.VerifyView(&_Verifier.CallOpts, signedReport)
}

func (_Verifier *VerifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "acceptOwnership")
}

func (_Verifier *VerifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _Verifier.Contract.AcceptOwnership(&_Verifier.TransactOpts)
}

func (_Verifier *VerifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Verifier.Contract.AcceptOwnership(&_Verifier.TransactOpts)
}

func (_Verifier *VerifierTransactor) ActivateConfig(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "activateConfig", configDigest)
}

func (_Verifier *VerifierSession) ActivateConfig(configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.ActivateConfig(&_Verifier.TransactOpts, configDigest)
}

func (_Verifier *VerifierTransactorSession) ActivateConfig(configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.ActivateConfig(&_Verifier.TransactOpts, configDigest)
}

func (_Verifier *VerifierTransactor) DeactivateConfig(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "deactivateConfig", configDigest)
}

func (_Verifier *VerifierSession) DeactivateConfig(configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.DeactivateConfig(&_Verifier.TransactOpts, configDigest)
}

func (_Verifier *VerifierTransactorSession) DeactivateConfig(configDigest [32]byte) (*types.Transaction, error) {
	return _Verifier.Contract.DeactivateConfig(&_Verifier.TransactOpts, configDigest)
}

func (_Verifier *VerifierTransactor) SetConfig(opts *bind.TransactOpts, configDigest [32]byte, signers []common.Address, f uint8, recipientAddressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "setConfig", configDigest, signers, f, recipientAddressesAndWeights)
}

func (_Verifier *VerifierSession) SetConfig(configDigest [32]byte, signers []common.Address, f uint8, recipientAddressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _Verifier.Contract.SetConfig(&_Verifier.TransactOpts, configDigest, signers, f, recipientAddressesAndWeights)
}

func (_Verifier *VerifierTransactorSession) SetConfig(configDigest [32]byte, signers []common.Address, f uint8, recipientAddressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _Verifier.Contract.SetConfig(&_Verifier.TransactOpts, configDigest, signers, f, recipientAddressesAndWeights)
}

func (_Verifier *VerifierTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "transferOwnership", to)
}

func (_Verifier *VerifierSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Verifier.Contract.TransferOwnership(&_Verifier.TransactOpts, to)
}

func (_Verifier *VerifierTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Verifier.Contract.TransferOwnership(&_Verifier.TransactOpts, to)
}

func (_Verifier *VerifierTransactor) UpdateConfig(opts *bind.TransactOpts, configDigest [32]byte, prevSigners []common.Address, newSigners []common.Address, f uint8) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "updateConfig", configDigest, prevSigners, newSigners, f)
}

func (_Verifier *VerifierSession) UpdateConfig(configDigest [32]byte, prevSigners []common.Address, newSigners []common.Address, f uint8) (*types.Transaction, error) {
	return _Verifier.Contract.UpdateConfig(&_Verifier.TransactOpts, configDigest, prevSigners, newSigners, f)
}

func (_Verifier *VerifierTransactorSession) UpdateConfig(configDigest [32]byte, prevSigners []common.Address, newSigners []common.Address, f uint8) (*types.Transaction, error) {
	return _Verifier.Contract.UpdateConfig(&_Verifier.TransactOpts, configDigest, prevSigners, newSigners, f)
}

func (_Verifier *VerifierTransactor) Verify(opts *bind.TransactOpts, signedReport []byte, sender common.Address) (*types.Transaction, error) {
	return _Verifier.contract.Transact(opts, "verify", signedReport, sender)
}

func (_Verifier *VerifierSession) Verify(signedReport []byte, sender common.Address) (*types.Transaction, error) {
	return _Verifier.Contract.Verify(&_Verifier.TransactOpts, signedReport, sender)
}

func (_Verifier *VerifierTransactorSession) Verify(signedReport []byte, sender common.Address) (*types.Transaction, error) {
	return _Verifier.Contract.Verify(&_Verifier.TransactOpts, signedReport, sender)
}

type VerifierConfigActivatedIterator struct {
	Event *VerifierConfigActivated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierConfigActivatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierConfigActivated)
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
		it.Event = new(VerifierConfigActivated)
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

func (it *VerifierConfigActivatedIterator) Error() error {
	return it.fail
}

func (it *VerifierConfigActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierConfigActivated struct {
	ConfigDigest [32]byte
	Raw          types.Log
}

func (_Verifier *VerifierFilterer) FilterConfigActivated(opts *bind.FilterOpts, configDigest [][32]byte) (*VerifierConfigActivatedIterator, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "ConfigActivated", configDigestRule)
	if err != nil {
		return nil, err
	}
	return &VerifierConfigActivatedIterator{contract: _Verifier.contract, event: "ConfigActivated", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchConfigActivated(opts *bind.WatchOpts, sink chan<- *VerifierConfigActivated, configDigest [][32]byte) (event.Subscription, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "ConfigActivated", configDigestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierConfigActivated)
				if err := _Verifier.contract.UnpackLog(event, "ConfigActivated", log); err != nil {
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

func (_Verifier *VerifierFilterer) ParseConfigActivated(log types.Log) (*VerifierConfigActivated, error) {
	event := new(VerifierConfigActivated)
	if err := _Verifier.contract.UnpackLog(event, "ConfigActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierConfigDeactivatedIterator struct {
	Event *VerifierConfigDeactivated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierConfigDeactivatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierConfigDeactivated)
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
		it.Event = new(VerifierConfigDeactivated)
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

func (it *VerifierConfigDeactivatedIterator) Error() error {
	return it.fail
}

func (it *VerifierConfigDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierConfigDeactivated struct {
	ConfigDigest [32]byte
	Raw          types.Log
}

func (_Verifier *VerifierFilterer) FilterConfigDeactivated(opts *bind.FilterOpts, configDigest [][32]byte) (*VerifierConfigDeactivatedIterator, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "ConfigDeactivated", configDigestRule)
	if err != nil {
		return nil, err
	}
	return &VerifierConfigDeactivatedIterator{contract: _Verifier.contract, event: "ConfigDeactivated", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchConfigDeactivated(opts *bind.WatchOpts, sink chan<- *VerifierConfigDeactivated, configDigest [][32]byte) (event.Subscription, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "ConfigDeactivated", configDigestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierConfigDeactivated)
				if err := _Verifier.contract.UnpackLog(event, "ConfigDeactivated", log); err != nil {
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

func (_Verifier *VerifierFilterer) ParseConfigDeactivated(log types.Log) (*VerifierConfigDeactivated, error) {
	event := new(VerifierConfigDeactivated)
	if err := _Verifier.contract.UnpackLog(event, "ConfigDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierConfigSetIterator struct {
	Event *VerifierConfigSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierConfigSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierConfigSet)
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
		it.Event = new(VerifierConfigSet)
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

func (it *VerifierConfigSetIterator) Error() error {
	return it.fail
}

func (it *VerifierConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierConfigSet struct {
	ConfigDigest [32]byte
	Signers      []common.Address
	F            uint8
	Raw          types.Log
}

func (_Verifier *VerifierFilterer) FilterConfigSet(opts *bind.FilterOpts, configDigest [][32]byte) (*VerifierConfigSetIterator, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "ConfigSet", configDigestRule)
	if err != nil {
		return nil, err
	}
	return &VerifierConfigSetIterator{contract: _Verifier.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierConfigSet, configDigest [][32]byte) (event.Subscription, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "ConfigSet", configDigestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierConfigSet)
				if err := _Verifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

func (_Verifier *VerifierFilterer) ParseConfigSet(log types.Log) (*VerifierConfigSet, error) {
	event := new(VerifierConfigSet)
	if err := _Verifier.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierConfigUpdatedIterator struct {
	Event *VerifierConfigUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierConfigUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierConfigUpdated)
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
		it.Event = new(VerifierConfigUpdated)
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

func (it *VerifierConfigUpdatedIterator) Error() error {
	return it.fail
}

func (it *VerifierConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierConfigUpdated struct {
	ConfigDigest [32]byte
	PrevSigners  []common.Address
	NewSigners   []common.Address
	Raw          types.Log
}

func (_Verifier *VerifierFilterer) FilterConfigUpdated(opts *bind.FilterOpts, configDigest [][32]byte) (*VerifierConfigUpdatedIterator, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "ConfigUpdated", configDigestRule)
	if err != nil {
		return nil, err
	}
	return &VerifierConfigUpdatedIterator{contract: _Verifier.contract, event: "ConfigUpdated", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *VerifierConfigUpdated, configDigest [][32]byte) (event.Subscription, error) {

	var configDigestRule []interface{}
	for _, configDigestItem := range configDigest {
		configDigestRule = append(configDigestRule, configDigestItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "ConfigUpdated", configDigestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierConfigUpdated)
				if err := _Verifier.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
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

func (_Verifier *VerifierFilterer) ParseConfigUpdated(log types.Log) (*VerifierConfigUpdated, error) {
	event := new(VerifierConfigUpdated)
	if err := _Verifier.contract.UnpackLog(event, "ConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierOwnershipTransferRequestedIterator struct {
	Event *VerifierOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierOwnershipTransferRequested)
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
		it.Event = new(VerifierOwnershipTransferRequested)
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

func (it *VerifierOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Verifier *VerifierFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierOwnershipTransferRequestedIterator{contract: _Verifier.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierOwnershipTransferRequested)
				if err := _Verifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_Verifier *VerifierFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierOwnershipTransferRequested, error) {
	event := new(VerifierOwnershipTransferRequested)
	if err := _Verifier.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierOwnershipTransferredIterator struct {
	Event *VerifierOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierOwnershipTransferred)
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
		it.Event = new(VerifierOwnershipTransferred)
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

func (it *VerifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_Verifier *VerifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierOwnershipTransferredIterator{contract: _Verifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierOwnershipTransferred)
				if err := _Verifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Verifier *VerifierFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierOwnershipTransferred, error) {
	event := new(VerifierOwnershipTransferred)
	if err := _Verifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierReportVerifiedIterator struct {
	Event *VerifierReportVerified

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierReportVerifiedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierReportVerified)
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
		it.Event = new(VerifierReportVerified)
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

func (it *VerifierReportVerifiedIterator) Error() error {
	return it.fail
}

func (it *VerifierReportVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierReportVerified struct {
	FeedId    [32]byte
	Requester common.Address
	Raw       types.Log
}

func (_Verifier *VerifierFilterer) FilterReportVerified(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierReportVerifiedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.FilterLogs(opts, "ReportVerified", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &VerifierReportVerifiedIterator{contract: _Verifier.contract, event: "ReportVerified", logs: logs, sub: sub}, nil
}

func (_Verifier *VerifierFilterer) WatchReportVerified(opts *bind.WatchOpts, sink chan<- *VerifierReportVerified, feedId [][32]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _Verifier.contract.WatchLogs(opts, "ReportVerified", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierReportVerified)
				if err := _Verifier.contract.UnpackLog(event, "ReportVerified", log); err != nil {
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

func (_Verifier *VerifierFilterer) ParseReportVerified(log types.Log) (*VerifierReportVerified, error) {
	event := new(VerifierReportVerified)
	if err := _Verifier.contract.UnpackLog(event, "ReportVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Verifier *Verifier) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Verifier.abi.Events["ConfigActivated"].ID:
		return _Verifier.ParseConfigActivated(log)
	case _Verifier.abi.Events["ConfigDeactivated"].ID:
		return _Verifier.ParseConfigDeactivated(log)
	case _Verifier.abi.Events["ConfigSet"].ID:
		return _Verifier.ParseConfigSet(log)
	case _Verifier.abi.Events["ConfigUpdated"].ID:
		return _Verifier.ParseConfigUpdated(log)
	case _Verifier.abi.Events["OwnershipTransferRequested"].ID:
		return _Verifier.ParseOwnershipTransferRequested(log)
	case _Verifier.abi.Events["OwnershipTransferred"].ID:
		return _Verifier.ParseOwnershipTransferred(log)
	case _Verifier.abi.Events["ReportVerified"].ID:
		return _Verifier.ParseReportVerified(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierConfigActivated) Topic() common.Hash {
	return common.HexToHash("0xa543797a0501218bba8a3daf75a71c8df8d1a7f791f4e44d40e43b6450183cea")
}

func (VerifierConfigDeactivated) Topic() common.Hash {
	return common.HexToHash("0x5bfaab86edc1b932e3c334327a591c9ded067cb521abae19b95ca927d6076579")
}

func (VerifierConfigSet) Topic() common.Hash {
	return common.HexToHash("0x5b1f376eb2bda670fa39339616d0a73f45b61bec8faeba8ca834f2ebb49676e0")
}

func (VerifierConfigUpdated) Topic() common.Hash {
	return common.HexToHash("0xb0b75a854fab801413da6202fc07e875c54eaf371a1e3909fb2645364ba58616")
}

func (VerifierOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VerifierReportVerified) Topic() common.Hash {
	return common.HexToHash("0x58ca9502e98a536e06e72d680fcc251e5d10b72291a281665a2c2dc0ac30fcc5")
}

func (_Verifier *Verifier) Address() common.Address {
	return _Verifier.address
}

type VerifierInterface interface {
	LatestConfigDetails(opts *bind.CallOpts, configDigest [32]byte) (uint32, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VerifyView(opts *bind.CallOpts, signedReport []byte) ([]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	ActivateConfig(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error)

	DeactivateConfig(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error)

	SetConfig(opts *bind.TransactOpts, configDigest [32]byte, signers []common.Address, f uint8, recipientAddressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateConfig(opts *bind.TransactOpts, configDigest [32]byte, prevSigners []common.Address, newSigners []common.Address, f uint8) (*types.Transaction, error)

	Verify(opts *bind.TransactOpts, signedReport []byte, sender common.Address) (*types.Transaction, error)

	FilterConfigActivated(opts *bind.FilterOpts, configDigest [][32]byte) (*VerifierConfigActivatedIterator, error)

	WatchConfigActivated(opts *bind.WatchOpts, sink chan<- *VerifierConfigActivated, configDigest [][32]byte) (event.Subscription, error)

	ParseConfigActivated(log types.Log) (*VerifierConfigActivated, error)

	FilterConfigDeactivated(opts *bind.FilterOpts, configDigest [][32]byte) (*VerifierConfigDeactivatedIterator, error)

	WatchConfigDeactivated(opts *bind.WatchOpts, sink chan<- *VerifierConfigDeactivated, configDigest [][32]byte) (event.Subscription, error)

	ParseConfigDeactivated(log types.Log) (*VerifierConfigDeactivated, error)

	FilterConfigSet(opts *bind.FilterOpts, configDigest [][32]byte) (*VerifierConfigSetIterator, error)

	WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VerifierConfigSet, configDigest [][32]byte) (event.Subscription, error)

	ParseConfigSet(log types.Log) (*VerifierConfigSet, error)

	FilterConfigUpdated(opts *bind.FilterOpts, configDigest [][32]byte) (*VerifierConfigUpdatedIterator, error)

	WatchConfigUpdated(opts *bind.WatchOpts, sink chan<- *VerifierConfigUpdated, configDigest [][32]byte) (event.Subscription, error)

	ParseConfigUpdated(log types.Log) (*VerifierConfigUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierOwnershipTransferred, error)

	FilterReportVerified(opts *bind.FilterOpts, feedId [][32]byte) (*VerifierReportVerifiedIterator, error)

	WatchReportVerified(opts *bind.WatchOpts, sink chan<- *VerifierReportVerified, feedId [][32]byte) (event.Subscription, error)

	ParseReportVerified(log types.Log) (*VerifierReportVerified, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
