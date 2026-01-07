// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verifier_proxy_v0_5_1

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

var VerifierProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"accessController\",\"type\":\"address\",\"internalType\":\"contractAccessControllerInterface\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getVerifier\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeVerifier\",\"inputs\":[{\"name\":\"verifierAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_accessController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAccessControllerInterface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"s_feeManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVerifierFeeManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAccessController\",\"inputs\":[{\"name\":\"accessController\",\"type\":\"address\",\"internalType\":\"contractAccessControllerInterface\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeManager\",\"inputs\":[{\"name\":\"feeManager\",\"type\":\"address\",\"internalType\":\"contractIVerifierFeeManager\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVerifier\",\"inputs\":[{\"name\":\"currentConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newConfigDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"addressesAndWeights\",\"type\":\"tuple[]\",\"internalType\":\"structCommon.AddressAndWeight[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unsetVerifier\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"parameterPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyBulk\",\"inputs\":[{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"parameterPayload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"verifiedReports\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyBulkView\",\"inputs\":[{\"name\":\"payloads\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"verifiedReports\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyView\",\"inputs\":[{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AccessControllerSet\",\"inputs\":[{\"name\":\"oldAccessController\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAccessController\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeManagerSet\",\"inputs\":[{\"name\":\"oldFeeManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newFeeManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierInitialized\",\"inputs\":[{\"name\":\"verifierAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierSet\",\"inputs\":[{\"name\":\"oldConfigDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"newConfigDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"verifierAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifierUnset\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"verifierAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessForbidden\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BadVerification\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ConfigDigestAlreadySet\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"FeeManagerInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerifierAlreadyInitialized\",\"inputs\":[{\"name\":\"verifier\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"VerifierInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerifierNotFound\",\"inputs\":[{\"name\":\"configDigest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620021b3380380620021b3833981016040819052620000349162000193565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000e8565b5050600480546001600160a01b0319166001600160a01b03939093169290921790915550620001c5565b336001600160a01b03821603620001425760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600060208284031215620001a657600080fd5b81516001600160a01b0381168114620001be57600080fd5b9392505050565b611fde80620001d56000396000f3fe6080604052600436106100f35760003560e01c80638da5cb5b1161008a578063f08391d811610059578063f08391d814610321578063f2fde38b14610341578063f7e83aee14610361578063f873a61c1461037457600080fd5b80638da5cb5b1461026657806394ba284614610291578063b011b247146102be578063eeb7b248146102de57600080fd5b8063665bc7a3116100c6578063665bc7a3146101e45780636e9140941461021157806379ba5097146102315780638c2a4d531461024657600080fd5b8063181f5a77146100f857806338416b5b14610150578063472d35b9146101a257806351b34c90146101c4575b600080fd5b34801561010457600080fd5b5060408051808201909152601381527f566572696669657250726f787920322e302e310000000000000000000000000060208201525b60405161014791906117b6565b60405180910390f35b34801561015c57600080fd5b5060055461017d9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610147565b3480156101ae57600080fd5b506101c26101bd3660046117f2565b610387565b005b3480156101d057600080fd5b5061013a6101df366004611858565b6105ff565b3480156101f057600080fd5b506102046101ff3660046118df565b610708565b6040516101479190611915565b34801561021d57600080fd5b506101c261022c366004611995565b6108bb565b34801561023d57600080fd5b506101c26109ac565b34801561025257600080fd5b506101c26102613660046117f2565b610aa9565b34801561027257600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff1661017d565b34801561029d57600080fd5b5060045461017d9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102ca57600080fd5b506101c26102d93660046119ae565b610cda565b3480156102ea57600080fd5b5061017d6102f9366004611995565b60009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b34801561032d57600080fd5b506101c261033c3660046117f2565b610f00565b34801561034d57600080fd5b506101c261035c3660046117f2565b610f87565b61013a61036f366004611a31565b610f9b565b610204610382366004611a9d565b611155565b61038f6113b9565b73ffffffffffffffffffffffffffffffffffffffff81166103dc576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527fdba45fe000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa158015610466573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048a9190611ad7565b158061054157506040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f6c2f1a1700000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa15801561051b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053f9190611ad7565b155b15610578576040517f8238941900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f04628abcaa6b1674651352125cb94b65b289145bc2bc4d67720bb7d966372f0391015b60405180910390a15050565b60045460609073ffffffffffffffffffffffffffffffffffffffff1680158015906106bf57506040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690636b14daf89061067c9033906000903690600401611b42565b602060405180830381865afa158015610699573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106bd9190611ad7565b155b156106f6576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610700848461143c565b949350505050565b60045460609073ffffffffffffffffffffffffffffffffffffffff1680158015906107c857506040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690636b14daf8906107859033906000903690600401611b42565b602060405180830381865afa1580156107a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c69190611ad7565b155b156107ff576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8267ffffffffffffffff81111561081857610818611b7b565b60405190808252806020026020018201604052801561084b57816020015b60608152602001906001900390816108365790505b50915060005b838110156108b35761088585858381811061086e5761086e611baa565b90506020028101906108809190611bd9565b61143c565b83828151811061089757610897611baa565b6020026020010181905250806108ac90611c3e565b9050610851565b505092915050565b6108c36113b9565b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610927576040517fb151802b000000000000000000000000000000000000000000000000000000008152600481018390526024015b60405180910390fd5b6000828152600360205260409081902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055517f11dc15c4b8ac2b183166cc8427e5385a5ece8308217a4217338c6a7614845c4c906105f3908490849091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a2d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161091e565b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610ab16113b9565b8073ffffffffffffffffffffffffffffffffffffffff8116610aff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f01ffc9a70000000000000000000000000000000000000000000000000000000081527f3d3ac1b500000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff8216906301ffc9a790602401602060405180830381865afa158015610b89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bad9190611ad7565b610be3576040517f75b0527a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660009081526002602052604090205460ff1615610c5b576040517f4e01ccfd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161091e565b73ffffffffffffffffffffffffffffffffffffffff821660008181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527f1f2cd7c97f4d801b5efe26cc409617c1fd6c5ef786e79aacb90af40923e4e8e991016105f3565b600083815260036020526040902054839073ffffffffffffffffffffffffffffffffffffffff168015610d58576040517f375d1fe60000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8216602482015260440161091e565b3360009081526002602052604090205460ff16610da1576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600085815260036020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016331790558215610eb95760055473ffffffffffffffffffffffffffffffffffffffff16610e2c576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6005546040517ff65df96200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063f65df96290610e8690889088908890600401611c9d565b600060405180830381600087803b158015610ea057600080fd5b505af1158015610eb4573d6000803e3d6000fd5b505050505b6040805187815260208101879052338183015290517fbeb513e532542a562ac35699e7cd9ae7d198dcd3eee15bada6c857d28ceaddcf9181900360600190a1505050505050565b610f086113b9565b6004805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f953e92b1a6442e9c3242531154a3f6f6eb00b4e9c719ba8118fa6235e4ce89b691016105f3565b610f8f6113b9565b610f988161156e565b50565b60045460609073ffffffffffffffffffffffffffffffffffffffff16801580159061105b57506040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690636b14daf8906110189033906000903690600401611b42565b602060405180830381865afa158015611035573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110599190611ad7565b155b15611092576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055473ffffffffffffffffffffffffffffffffffffffff168015611140576040517fdba45fe000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82169063dba45fe090349061110d908b908b908b908b903390600401611d26565b6000604051808303818588803b15801561112657600080fd5b505af115801561113a573d6000803e3d6000fd5b50505050505b61114a8787611663565b979650505050505050565b60045460609073ffffffffffffffffffffffffffffffffffffffff16801580159061121557506040517f6b14daf800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690636b14daf8906111d29033906000903690600401611b42565b602060405180830381865afa1580156111ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112139190611ad7565b155b1561124c576040517fef67f5d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055473ffffffffffffffffffffffffffffffffffffffff1680156112fa576040517f6c2f1a1700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690636c2f1a179034906112c7908b908b908b908b903390600401611d76565b6000604051808303818588803b1580156112e057600080fd5b505af11580156112f4573d6000803e3d6000fd5b50505050505b8567ffffffffffffffff81111561131357611313611b7b565b60405190808252806020026020018201604052801561134657816020015b60608152602001906001900390816113315790505b50925060005b868110156113ae5761138088888381811061136957611369611baa565b905060200281019061137b9190611bd9565b611663565b84828151811061139257611392611baa565b6020026020010181905250806113a790611c3e565b905061134c565b505050949350505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461143a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161091e565b565b6060600061144a8385611e87565b60008181526003602052604090205490915073ffffffffffffffffffffffffffffffffffffffff16806114ac576040517fb151802b0000000000000000000000000000000000000000000000000000000081526004810183905260240161091e565b6040517f51b34c9000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216906351b34c90906115009088908890600401611ec3565b600060405180830381865afa15801561151d573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526115639190810190611ed7565b925050505b92915050565b3373ffffffffffffffffffffffffffffffffffffffff8216036115ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161091e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b606060006116718385611e87565b60008181526003602052604090205490915073ffffffffffffffffffffffffffffffffffffffff16806116d3576040517fb151802b0000000000000000000000000000000000000000000000000000000081526004810183905260240161091e565b6040517f3d3ac1b500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821690633d3ac1b59061172990889088903390600401611f97565b6000604051808303816000875af115801561151d573d6000803e3d6000fd5b60005b8381101561176357818101518382015260200161174b565b50506000910152565b60008151808452611784816020860160208601611748565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006117c9602083018461176c565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610f9857600080fd5b60006020828403121561180457600080fd5b81356117c9816117d0565b60008083601f84011261182157600080fd5b50813567ffffffffffffffff81111561183957600080fd5b60208301915083602082850101111561185157600080fd5b9250929050565b6000806020838503121561186b57600080fd5b823567ffffffffffffffff81111561188257600080fd5b61188e8582860161180f565b90969095509350505050565b60008083601f8401126118ac57600080fd5b50813567ffffffffffffffff8111156118c457600080fd5b6020830191508360208260051b850101111561185157600080fd5b600080602083850312156118f257600080fd5b823567ffffffffffffffff81111561190957600080fd5b61188e8582860161189a565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611988577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc088860301845261197685835161176c565b9450928501929085019060010161193c565b5092979650505050505050565b6000602082840312156119a757600080fd5b5035919050565b600080600080606085870312156119c457600080fd5b8435935060208501359250604085013567ffffffffffffffff808211156119ea57600080fd5b818701915087601f8301126119fe57600080fd5b813581811115611a0d57600080fd5b8860208260061b8501011115611a2257600080fd5b95989497505060200194505050565b60008060008060408587031215611a4757600080fd5b843567ffffffffffffffff80821115611a5f57600080fd5b611a6b8883890161180f565b90965094506020870135915080821115611a8457600080fd5b50611a918782880161180f565b95989497509550505050565b60008060008060408587031215611ab357600080fd5b843567ffffffffffffffff80821115611acb57600080fd5b611a6b8883890161189a565b600060208284031215611ae957600080fd5b815180151581146117c957600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff84168152604060208201526000611b72604083018486611af9565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112611c0e57600080fd5b83018035915067ffffffffffffffff821115611c2957600080fd5b60200191503681900382131561185157600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c96577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b838152604060208083018290528282018490526000919085906060850184805b88811015611d17578435611cd0816117d0565b73ffffffffffffffffffffffffffffffffffffffff1683528484013567ffffffffffffffff8116808214611d02578384fd5b84860152509385019391850191600101611cbd565b50909998505050505050505050565b606081526000611d3a606083018789611af9565b8281036020840152611d4d818688611af9565b91505073ffffffffffffffffffffffffffffffffffffffff831660408301529695505050505050565b6060808252810185905260006080600587901b8301810190830188835b89811015611e42577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8086850301835281357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18c3603018112611df457600080fd5b8b01602081810191359067ffffffffffffffff821115611e1357600080fd5b813603831315611e2257600080fd5b611e2d878385611af9565b96509485019493909301925050600101611d93565b5050508281036020840152611e58818688611af9565b915050611e7d604083018473ffffffffffffffffffffffffffffffffffffffff169052565b9695505050505050565b80356020831015611568577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602084900360031b1b1692915050565b602081526000610700602083018486611af9565b600060208284031215611ee957600080fd5b815167ffffffffffffffff80821115611f0157600080fd5b818401915084601f830112611f1557600080fd5b815181811115611f2757611f27611b7b565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611f6d57611f6d611b7b565b81604052828152876020848701011115611f8657600080fd5b61114a836020830160208801611748565b604081526000611fab604083018587611af9565b905073ffffffffffffffffffffffffffffffffffffffff8316602083015294935050505056fea164736f6c6343000813000a",
}

var VerifierProxyABI = VerifierProxyMetaData.ABI

var VerifierProxyBin = VerifierProxyMetaData.Bin

func DeployVerifierProxy(auth *bind.TransactOpts, backend bind.ContractBackend, accessController common.Address) (common.Address, *types.Transaction, *VerifierProxy, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierProxyBin), backend, accessController)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifierProxy{address: address, abi: *parsed, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

type VerifierProxy struct {
	address common.Address
	abi     abi.ABI
	VerifierProxyCaller
	VerifierProxyTransactor
	VerifierProxyFilterer
}

type VerifierProxyCaller struct {
	contract *bind.BoundContract
}

type VerifierProxyTransactor struct {
	contract *bind.BoundContract
}

type VerifierProxyFilterer struct {
	contract *bind.BoundContract
}

type VerifierProxySession struct {
	Contract     *VerifierProxy
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type VerifierProxyCallerSession struct {
	Contract *VerifierProxyCaller
	CallOpts bind.CallOpts
}

type VerifierProxyTransactorSession struct {
	Contract     *VerifierProxyTransactor
	TransactOpts bind.TransactOpts
}

type VerifierProxyRaw struct {
	Contract *VerifierProxy
}

type VerifierProxyCallerRaw struct {
	Contract *VerifierProxyCaller
}

type VerifierProxyTransactorRaw struct {
	Contract *VerifierProxyTransactor
}

func NewVerifierProxy(address common.Address, backend bind.ContractBackend) (*VerifierProxy, error) {
	abi, err := abi.JSON(strings.NewReader(VerifierProxyABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindVerifierProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifierProxy{address: address, abi: abi, VerifierProxyCaller: VerifierProxyCaller{contract: contract}, VerifierProxyTransactor: VerifierProxyTransactor{contract: contract}, VerifierProxyFilterer: VerifierProxyFilterer{contract: contract}}, nil
}

func NewVerifierProxyCaller(address common.Address, caller bind.ContractCaller) (*VerifierProxyCaller, error) {
	contract, err := bindVerifierProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyCaller{contract: contract}, nil
}

func NewVerifierProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierProxyTransactor, error) {
	contract, err := bindVerifierProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyTransactor{contract: contract}, nil
}

func NewVerifierProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierProxyFilterer, error) {
	contract, err := bindVerifierProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFilterer{contract: contract}, nil
}

func bindVerifierProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_VerifierProxy *VerifierProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.VerifierProxyCaller.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifierProxyTransactor.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifierProxy.Contract.contract.Call(opts, result, method, params...)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transfer(opts)
}

func (_VerifierProxy *VerifierProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifierProxy.Contract.contract.Transact(opts, method, params...)
}

func (_VerifierProxy *VerifierProxyCaller) GetVerifier(opts *bind.CallOpts, configDigest [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "getVerifier", configDigest)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) GetVerifier(configDigest [32]byte) (common.Address, error) {
	return _VerifierProxy.Contract.GetVerifier(&_VerifierProxy.CallOpts, configDigest)
}

func (_VerifierProxy *VerifierProxyCallerSession) GetVerifier(configDigest [32]byte) (common.Address, error) {
	return _VerifierProxy.Contract.GetVerifier(&_VerifierProxy.CallOpts, configDigest)
}

func (_VerifierProxy *VerifierProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) Owner() (common.Address, error) {
	return _VerifierProxy.Contract.Owner(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) SAccessController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "s_accessController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) SAccessController() (common.Address, error) {
	return _VerifierProxy.Contract.SAccessController(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) SAccessController() (common.Address, error) {
	return _VerifierProxy.Contract.SAccessController(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) SFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "s_feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) SFeeManager() (common.Address, error) {
	return _VerifierProxy.Contract.SFeeManager(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) SFeeManager() (common.Address, error) {
	return _VerifierProxy.Contract.SFeeManager(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCallerSession) TypeAndVersion() (string, error) {
	return _VerifierProxy.Contract.TypeAndVersion(&_VerifierProxy.CallOpts)
}

func (_VerifierProxy *VerifierProxyCaller) VerifyBulkView(opts *bind.CallOpts, payloads [][]byte) ([][]byte, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "verifyBulkView", payloads)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) VerifyBulkView(payloads [][]byte) ([][]byte, error) {
	return _VerifierProxy.Contract.VerifyBulkView(&_VerifierProxy.CallOpts, payloads)
}

func (_VerifierProxy *VerifierProxyCallerSession) VerifyBulkView(payloads [][]byte) ([][]byte, error) {
	return _VerifierProxy.Contract.VerifyBulkView(&_VerifierProxy.CallOpts, payloads)
}

func (_VerifierProxy *VerifierProxyCaller) VerifyView(opts *bind.CallOpts, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _VerifierProxy.contract.Call(opts, &out, "verifyView", payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_VerifierProxy *VerifierProxySession) VerifyView(payload []byte) ([]byte, error) {
	return _VerifierProxy.Contract.VerifyView(&_VerifierProxy.CallOpts, payload)
}

func (_VerifierProxy *VerifierProxyCallerSession) VerifyView(payload []byte) ([]byte, error) {
	return _VerifierProxy.Contract.VerifyView(&_VerifierProxy.CallOpts, payload)
}

func (_VerifierProxy *VerifierProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "acceptOwnership")
}

func (_VerifierProxy *VerifierProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VerifierProxy.Contract.AcceptOwnership(&_VerifierProxy.TransactOpts)
}

func (_VerifierProxy *VerifierProxyTransactor) InitializeVerifier(opts *bind.TransactOpts, verifierAddress common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "initializeVerifier", verifierAddress)
}

func (_VerifierProxy *VerifierProxySession) InitializeVerifier(verifierAddress common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.InitializeVerifier(&_VerifierProxy.TransactOpts, verifierAddress)
}

func (_VerifierProxy *VerifierProxyTransactorSession) InitializeVerifier(verifierAddress common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.InitializeVerifier(&_VerifierProxy.TransactOpts, verifierAddress)
}

func (_VerifierProxy *VerifierProxyTransactor) SetAccessController(opts *bind.TransactOpts, accessController common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setAccessController", accessController)
}

func (_VerifierProxy *VerifierProxySession) SetAccessController(accessController common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetAccessController(&_VerifierProxy.TransactOpts, accessController)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetAccessController(accessController common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetAccessController(&_VerifierProxy.TransactOpts, accessController)
}

func (_VerifierProxy *VerifierProxyTransactor) SetFeeManager(opts *bind.TransactOpts, feeManager common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setFeeManager", feeManager)
}

func (_VerifierProxy *VerifierProxySession) SetFeeManager(feeManager common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetFeeManager(&_VerifierProxy.TransactOpts, feeManager)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetFeeManager(feeManager common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetFeeManager(&_VerifierProxy.TransactOpts, feeManager)
}

func (_VerifierProxy *VerifierProxyTransactor) SetVerifier(opts *bind.TransactOpts, currentConfigDigest [32]byte, newConfigDigest [32]byte, addressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "setVerifier", currentConfigDigest, newConfigDigest, addressesAndWeights)
}

func (_VerifierProxy *VerifierProxySession) SetVerifier(currentConfigDigest [32]byte, newConfigDigest [32]byte, addressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetVerifier(&_VerifierProxy.TransactOpts, currentConfigDigest, newConfigDigest, addressesAndWeights)
}

func (_VerifierProxy *VerifierProxyTransactorSession) SetVerifier(currentConfigDigest [32]byte, newConfigDigest [32]byte, addressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error) {
	return _VerifierProxy.Contract.SetVerifier(&_VerifierProxy.TransactOpts, currentConfigDigest, newConfigDigest, addressesAndWeights)
}

func (_VerifierProxy *VerifierProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "transferOwnership", to)
}

func (_VerifierProxy *VerifierProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VerifierProxy.Contract.TransferOwnership(&_VerifierProxy.TransactOpts, to)
}

func (_VerifierProxy *VerifierProxyTransactor) UnsetVerifier(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "unsetVerifier", configDigest)
}

func (_VerifierProxy *VerifierProxySession) UnsetVerifier(configDigest [32]byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.UnsetVerifier(&_VerifierProxy.TransactOpts, configDigest)
}

func (_VerifierProxy *VerifierProxyTransactorSession) UnsetVerifier(configDigest [32]byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.UnsetVerifier(&_VerifierProxy.TransactOpts, configDigest)
}

func (_VerifierProxy *VerifierProxyTransactor) Verify(opts *bind.TransactOpts, payload []byte, parameterPayload []byte) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "verify", payload, parameterPayload)
}

func (_VerifierProxy *VerifierProxySession) Verify(payload []byte, parameterPayload []byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.Verify(&_VerifierProxy.TransactOpts, payload, parameterPayload)
}

func (_VerifierProxy *VerifierProxyTransactorSession) Verify(payload []byte, parameterPayload []byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.Verify(&_VerifierProxy.TransactOpts, payload, parameterPayload)
}

func (_VerifierProxy *VerifierProxyTransactor) VerifyBulk(opts *bind.TransactOpts, payloads [][]byte, parameterPayload []byte) (*types.Transaction, error) {
	return _VerifierProxy.contract.Transact(opts, "verifyBulk", payloads, parameterPayload)
}

func (_VerifierProxy *VerifierProxySession) VerifyBulk(payloads [][]byte, parameterPayload []byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifyBulk(&_VerifierProxy.TransactOpts, payloads, parameterPayload)
}

func (_VerifierProxy *VerifierProxyTransactorSession) VerifyBulk(payloads [][]byte, parameterPayload []byte) (*types.Transaction, error) {
	return _VerifierProxy.Contract.VerifyBulk(&_VerifierProxy.TransactOpts, payloads, parameterPayload)
}

type VerifierProxyAccessControllerSetIterator struct {
	Event *VerifierProxyAccessControllerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyAccessControllerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyAccessControllerSet)
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
		it.Event = new(VerifierProxyAccessControllerSet)
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

func (it *VerifierProxyAccessControllerSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyAccessControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyAccessControllerSet struct {
	OldAccessController common.Address
	NewAccessController common.Address
	Raw                 types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterAccessControllerSet(opts *bind.FilterOpts) (*VerifierProxyAccessControllerSetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "AccessControllerSet")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyAccessControllerSetIterator{contract: _VerifierProxy.contract, event: "AccessControllerSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchAccessControllerSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyAccessControllerSet) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "AccessControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyAccessControllerSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "AccessControllerSet", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseAccessControllerSet(log types.Log) (*VerifierProxyAccessControllerSet, error) {
	event := new(VerifierProxyAccessControllerSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "AccessControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyFeeManagerSetIterator struct {
	Event *VerifierProxyFeeManagerSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyFeeManagerSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyFeeManagerSet)
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
		it.Event = new(VerifierProxyFeeManagerSet)
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

func (it *VerifierProxyFeeManagerSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyFeeManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyFeeManagerSet struct {
	OldFeeManager common.Address
	NewFeeManager common.Address
	Raw           types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterFeeManagerSet(opts *bind.FilterOpts) (*VerifierProxyFeeManagerSetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "FeeManagerSet")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyFeeManagerSetIterator{contract: _VerifierProxy.contract, event: "FeeManagerSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchFeeManagerSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeManagerSet) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "FeeManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyFeeManagerSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "FeeManagerSet", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseFeeManagerSet(log types.Log) (*VerifierProxyFeeManagerSet, error) {
	event := new(VerifierProxyFeeManagerSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "FeeManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferRequestedIterator struct {
	Event *VerifierProxyOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferRequested)
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
		it.Event = new(VerifierProxyOwnershipTransferRequested)
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

func (it *VerifierProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferRequestedIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferRequested)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error) {
	event := new(VerifierProxyOwnershipTransferRequested)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyOwnershipTransferredIterator struct {
	Event *VerifierProxyOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyOwnershipTransferred)
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
		it.Event = new(VerifierProxyOwnershipTransferred)
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

func (it *VerifierProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VerifierProxyOwnershipTransferredIterator{contract: _VerifierProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyOwnershipTransferred)
				if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error) {
	event := new(VerifierProxyOwnershipTransferred)
	if err := _VerifierProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyVerifierInitializedIterator struct {
	Event *VerifierProxyVerifierInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyVerifierInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyVerifierInitialized)
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
		it.Event = new(VerifierProxyVerifierInitialized)
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

func (it *VerifierProxyVerifierInitializedIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyVerifierInitialized struct {
	VerifierAddress common.Address
	Raw             types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterVerifierInitialized(opts *bind.FilterOpts) (*VerifierProxyVerifierInitializedIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "VerifierInitialized")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyVerifierInitializedIterator{contract: _VerifierProxy.contract, event: "VerifierInitialized", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchVerifierInitialized(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "VerifierInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyVerifierInitialized)
				if err := _VerifierProxy.contract.UnpackLog(event, "VerifierInitialized", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseVerifierInitialized(log types.Log) (*VerifierProxyVerifierInitialized, error) {
	event := new(VerifierProxyVerifierInitialized)
	if err := _VerifierProxy.contract.UnpackLog(event, "VerifierInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyVerifierSetIterator struct {
	Event *VerifierProxyVerifierSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyVerifierSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyVerifierSet)
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
		it.Event = new(VerifierProxyVerifierSet)
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

func (it *VerifierProxyVerifierSetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyVerifierSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyVerifierSet struct {
	OldConfigDigest [32]byte
	NewConfigDigest [32]byte
	VerifierAddress common.Address
	Raw             types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterVerifierSet(opts *bind.FilterOpts) (*VerifierProxyVerifierSetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyVerifierSetIterator{contract: _VerifierProxy.contract, event: "VerifierSet", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchVerifierSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierSet) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "VerifierSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyVerifierSet)
				if err := _VerifierProxy.contract.UnpackLog(event, "VerifierSet", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseVerifierSet(log types.Log) (*VerifierProxyVerifierSet, error) {
	event := new(VerifierProxyVerifierSet)
	if err := _VerifierProxy.contract.UnpackLog(event, "VerifierSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type VerifierProxyVerifierUnsetIterator struct {
	Event *VerifierProxyVerifierUnset

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *VerifierProxyVerifierUnsetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifierProxyVerifierUnset)
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
		it.Event = new(VerifierProxyVerifierUnset)
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

func (it *VerifierProxyVerifierUnsetIterator) Error() error {
	return it.fail
}

func (it *VerifierProxyVerifierUnsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type VerifierProxyVerifierUnset struct {
	ConfigDigest    [32]byte
	VerifierAddress common.Address
	Raw             types.Log
}

func (_VerifierProxy *VerifierProxyFilterer) FilterVerifierUnset(opts *bind.FilterOpts) (*VerifierProxyVerifierUnsetIterator, error) {

	logs, sub, err := _VerifierProxy.contract.FilterLogs(opts, "VerifierUnset")
	if err != nil {
		return nil, err
	}
	return &VerifierProxyVerifierUnsetIterator{contract: _VerifierProxy.contract, event: "VerifierUnset", logs: logs, sub: sub}, nil
}

func (_VerifierProxy *VerifierProxyFilterer) WatchVerifierUnset(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierUnset) (event.Subscription, error) {

	logs, sub, err := _VerifierProxy.contract.WatchLogs(opts, "VerifierUnset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(VerifierProxyVerifierUnset)
				if err := _VerifierProxy.contract.UnpackLog(event, "VerifierUnset", log); err != nil {
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

func (_VerifierProxy *VerifierProxyFilterer) ParseVerifierUnset(log types.Log) (*VerifierProxyVerifierUnset, error) {
	event := new(VerifierProxyVerifierUnset)
	if err := _VerifierProxy.contract.UnpackLog(event, "VerifierUnset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_VerifierProxy *VerifierProxy) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _VerifierProxy.abi.Events["AccessControllerSet"].ID:
		return _VerifierProxy.ParseAccessControllerSet(log)
	case _VerifierProxy.abi.Events["FeeManagerSet"].ID:
		return _VerifierProxy.ParseFeeManagerSet(log)
	case _VerifierProxy.abi.Events["OwnershipTransferRequested"].ID:
		return _VerifierProxy.ParseOwnershipTransferRequested(log)
	case _VerifierProxy.abi.Events["OwnershipTransferred"].ID:
		return _VerifierProxy.ParseOwnershipTransferred(log)
	case _VerifierProxy.abi.Events["VerifierInitialized"].ID:
		return _VerifierProxy.ParseVerifierInitialized(log)
	case _VerifierProxy.abi.Events["VerifierSet"].ID:
		return _VerifierProxy.ParseVerifierSet(log)
	case _VerifierProxy.abi.Events["VerifierUnset"].ID:
		return _VerifierProxy.ParseVerifierUnset(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (VerifierProxyAccessControllerSet) Topic() common.Hash {
	return common.HexToHash("0x953e92b1a6442e9c3242531154a3f6f6eb00b4e9c719ba8118fa6235e4ce89b6")
}

func (VerifierProxyFeeManagerSet) Topic() common.Hash {
	return common.HexToHash("0x04628abcaa6b1674651352125cb94b65b289145bc2bc4d67720bb7d966372f03")
}

func (VerifierProxyOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (VerifierProxyOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (VerifierProxyVerifierInitialized) Topic() common.Hash {
	return common.HexToHash("0x1f2cd7c97f4d801b5efe26cc409617c1fd6c5ef786e79aacb90af40923e4e8e9")
}

func (VerifierProxyVerifierSet) Topic() common.Hash {
	return common.HexToHash("0xbeb513e532542a562ac35699e7cd9ae7d198dcd3eee15bada6c857d28ceaddcf")
}

func (VerifierProxyVerifierUnset) Topic() common.Hash {
	return common.HexToHash("0x11dc15c4b8ac2b183166cc8427e5385a5ece8308217a4217338c6a7614845c4c")
}

func (_VerifierProxy *VerifierProxy) Address() common.Address {
	return _VerifierProxy.address
}

type VerifierProxyInterface interface {
	GetVerifier(opts *bind.CallOpts, configDigest [32]byte) (common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SAccessController(opts *bind.CallOpts) (common.Address, error)

	SFeeManager(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	VerifyBulkView(opts *bind.CallOpts, payloads [][]byte) ([][]byte, error)

	VerifyView(opts *bind.CallOpts, payload []byte) ([]byte, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	InitializeVerifier(opts *bind.TransactOpts, verifierAddress common.Address) (*types.Transaction, error)

	SetAccessController(opts *bind.TransactOpts, accessController common.Address) (*types.Transaction, error)

	SetFeeManager(opts *bind.TransactOpts, feeManager common.Address) (*types.Transaction, error)

	SetVerifier(opts *bind.TransactOpts, currentConfigDigest [32]byte, newConfigDigest [32]byte, addressesAndWeights []CommonAddressAndWeight) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnsetVerifier(opts *bind.TransactOpts, configDigest [32]byte) (*types.Transaction, error)

	Verify(opts *bind.TransactOpts, payload []byte, parameterPayload []byte) (*types.Transaction, error)

	VerifyBulk(opts *bind.TransactOpts, payloads [][]byte, parameterPayload []byte) (*types.Transaction, error)

	FilterAccessControllerSet(opts *bind.FilterOpts) (*VerifierProxyAccessControllerSetIterator, error)

	WatchAccessControllerSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyAccessControllerSet) (event.Subscription, error)

	ParseAccessControllerSet(log types.Log) (*VerifierProxyAccessControllerSet, error)

	FilterFeeManagerSet(opts *bind.FilterOpts) (*VerifierProxyFeeManagerSetIterator, error)

	WatchFeeManagerSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyFeeManagerSet) (event.Subscription, error)

	ParseFeeManagerSet(log types.Log) (*VerifierProxyFeeManagerSet, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*VerifierProxyOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VerifierProxyOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifierProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*VerifierProxyOwnershipTransferred, error)

	FilterVerifierInitialized(opts *bind.FilterOpts) (*VerifierProxyVerifierInitializedIterator, error)

	WatchVerifierInitialized(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierInitialized) (event.Subscription, error)

	ParseVerifierInitialized(log types.Log) (*VerifierProxyVerifierInitialized, error)

	FilterVerifierSet(opts *bind.FilterOpts) (*VerifierProxyVerifierSetIterator, error)

	WatchVerifierSet(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierSet) (event.Subscription, error)

	ParseVerifierSet(log types.Log) (*VerifierProxyVerifierSet, error)

	FilterVerifierUnset(opts *bind.FilterOpts) (*VerifierProxyVerifierUnsetIterator, error)

	WatchVerifierUnset(opts *bind.WatchOpts, sink chan<- *VerifierProxyVerifierUnset) (event.Subscription, error)

	ParseVerifierUnset(log types.Log) (*VerifierProxyVerifierUnset, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
