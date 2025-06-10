// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package workflow_registry_wrapper_v2

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

var WorkflowRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canLinkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canUnlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDefaults\",\"inputs\":[],\"outputs\":[{\"name\":\"maxPerDON\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerUserDON\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLinkedOwners\",\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"batchSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerDON\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxWorkflowsPerUserDON\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowedSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOwnerLinked\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"linkOwner\",\"inputs\":[{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDONOverride\",\"inputs\":[{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaults\",\"inputs\":[{\"name\":\"maxPerDON\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerUserDON\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserDONOverride\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"donLabel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"limit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalLinkedOwners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlinkOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumWorkflowRegistry.PreUnlinkAction\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAllowedSigners\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowedSignersUpdatedV1\",\"inputs\":[{\"name\":\"signers\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipLinkUpdatedV1\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"added\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotTransferToSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOwnershipLink\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"recoverErrorId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"recoverErrorArg\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MustBeProposedOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCallableByOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnershipLinkAlreadyExists\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkDoesNotExist\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnershipLinkProofDoesNotMatch\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"storedProof\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"UnlinkOwnerRequestExpired\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"currentTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiryTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x608080604052346054573315604357600180546001600160a01b03191633179055600680546001600160401b03191664c8000001f4179055611adc908161005a8239f35b639b15e16f60e01b60005260046000fd5b600080fdfe6080604052600436101561001257600080fd5b60003560e01c80630987294c14610e9c57806317e0edfc14610e04578063181f5a7714610d3457806323f381bb14610cef57806344d4154e14610b2657806379ba509714610a3d5780637ab06d02146109ae578063809b23cc146109645780638da5cb5b14610912578063a0b8a4fe146108d6578063af71bab21461074e578063b172ed4d14610702578063c27fb6bb14610606578063cabb9e7a1461059c578063d8e4a724146103ad578063dc1019691461034f578063dfcb0b3114610333578063f23e37ff146101e25763f2fde38b146100ed57600080fd5b346101dd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd5773ffffffffffffffffffffffffffffffffffffffff610139610f0b565b61014161165a565b163381146101b357807fffffffffffffffffffffffff0000000000000000000000000000000000000000600054161760005573ffffffffffffffffffffffffffffffffffffffff600154167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278600080a3005b7fdad89dca0000000000000000000000000000000000000000000000000000000060005260046000fd5b600080fd5b346101dd5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd57610219610f0b565b602435906044359163ffffffff83168093036101dd5760643580151581036101dd5761024361165a565b156102fb576040519261025584610f2e565b835273ffffffffffffffffffffffffffffffffffffffff60208401926001845216600052600860205260406000209060005260205263ffffffff806040600020935116167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000083549260201b169116179055600080f35b915073ffffffffffffffffffffffffffffffffffffffff16600052600860205260406000209060005260205260006040812055600080f35b346101dd5761034d610344366110b9565b9291909161142a565b005b346101dd57600161036b610362366110b9565b9180949361142a565b33600052600560205280604060002055610384336119cd565b50337f2195f84ef0d57d74c40638caac1ace3d9c91f7d4bd107849140f5437425ef65b600080a4005b346101dd5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd5760043567ffffffffffffffff81116101dd57366023820112156101dd57806004013567ffffffffffffffff81116101dd576024820191602436918360051b0101116101dd57602435918215158093036101dd5761043761165a565b60ff831660005b8381106104e257505060405191806040840160408552526060830191906000905b808210610495577f7b317b86e2b8023639eee242d6f7a311def02336487b07a136bfe8f67b51b39c8580868960208301520390a1005b90919283359073ffffffffffffffffffffffffffffffffffffffff821682036101dd576020809173ffffffffffffffffffffffffffffffffffffffff60019416815201940192019061045f565b73ffffffffffffffffffffffffffffffffffffffff61050a6105058387876113ca565b611409565b1615610572578073ffffffffffffffffffffffffffffffffffffffff61053661050560019488886113ca565b1660005260026020526040600020837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790550161043e565b7f8579befe0000000000000000000000000000000000000000000000000000000060005260046000fd5b346101dd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd5773ffffffffffffffffffffffffffffffffffffffff6105e8610f0b565b166000526002602052602060ff604060002054166040519015158152f35b346101dd5760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd576004356106406110a6565b9060443580151581036101dd5761065561165a565b156106f15763ffffffff6040519261066c84610f2e565b168252602082019060018252600052600760205263ffffffff806040600020935116167fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000008354161782555115157fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff64ff0000000083549260201b169116179055600080f35b600090815260076020526040812055005b346101dd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd57604060065463ffffffff825191818116835260201c166020820152f35b346101dd5761075c36611022565b5090839493421161088a5773ffffffffffffffffffffffffffffffffffffffff8416610795816000526004602052604060002054151590565b1561085d578060005260056020526040600020549084820361082857505073ffffffffffffffffffffffffffffffffffffffff6107d583838689896116a5565b16600052600260205260ff60406000205416156107ee57005b610824926040519586957f335d4ce100000000000000000000000000000000000000000000000000000000875260048701611391565b0390fd5b84907fc24eb43a0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b7fc2dda3f90000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8473ffffffffffffffffffffffffffffffffffffffff857f3d8a511600000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b346101dd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd576020600354604051908152f35b346101dd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd57602073ffffffffffffffffffffffffffffffffffffffff60015416604051908152f35b346101dd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd5760206109a060043561130c565b63ffffffff60405191168152f35b346101dd5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd5760043563ffffffff81168091036101dd576109f66110a6565b906109ff61165a565b7fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000067ffffffff000000006006549360201b1692161717600655600080f35b346101dd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd5760005473ffffffffffffffffffffffffffffffffffffffff81163303610afc577fffffffffffffffffffffffff00000000000000000000000000000000000000006001549133828416176001551660005573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a3005b7f02b543c60000000000000000000000000000000000000000000000000000000060005260046000fd5b346101dd57610b3436611022565b5083949192934211610ca35773ffffffffffffffffffffffffffffffffffffffff821692610b6f846000526004602052604060002054151590565b15610c7557836000526005602052604060002054858103610c40575073ffffffffffffffffffffffffffffffffffffffff610bad8383888a886116a5565b16600052600260205260ff6040600020541615610c0657600085858083526005602052826040812055610bdf816117fb565b507f2195f84ef0d57d74c40638caac1ace3d9c91f7d4bd107849140f5437425ef65b8380a4005b90919361082493506040519586957f335d4ce100000000000000000000000000000000000000000000000000000000875260048701611391565b85857fc24eb43a0000000000000000000000000000000000000000000000000000000060005260045260245260445260646000fd5b837fc2dda3f90000000000000000000000000000000000000000000000000000000060005260045260246000fd5b8473ffffffffffffffffffffffffffffffffffffffff837f3d8a511600000000000000000000000000000000000000000000000000000000600052166004524260245260445260646000fd5b346101dd5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd5760206109a0610d2b610f0b565b6024359061129e565b346101dd5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd576040805190610d728183610f79565b601a82527f576f726b666c6f77526567697374727920322e302e302d6465760000000000006020830152805180926020825280519081602084015260005b828110610ded5750506000828201840152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168101030190f35b602082820181015187830187015286945001610db0565b346101dd5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd57610e41602435600435611162565b60405180916020820160208352815180915260206040840192019060005b818110610e6d575050500390f35b825173ffffffffffffffffffffffffffffffffffffffff16845285945060209384019390920191600101610e5f565b346101dd5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126101dd576020610f0173ffffffffffffffffffffffffffffffffffffffff610eed610f0b565b166000526004602052604060002054151590565b6040519015158152f35b6004359073ffffffffffffffffffffffffffffffffffffffff821682036101dd57565b6040810190811067ffffffffffffffff821117610f4a57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610f4a57604052565b67ffffffffffffffff8111610f4a57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b9181601f840112156101dd5782359167ffffffffffffffff83116101dd57602083818601950101116101dd57565b9060a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8301126101dd5760043573ffffffffffffffffffffffffffffffffffffffff811681036101dd579160243591604435916064359067ffffffffffffffff82116101dd5761109591600401610ff4565b909160843560038110156101dd5790565b6024359063ffffffff821682036101dd57565b60607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc8201126101dd5760043591602435916044359067ffffffffffffffff82116101dd5761110a91600401610ff4565b9091565b67ffffffffffffffff8111610f4a5760051b60200190565b9190820180921161113357565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906003549081831015611281578161117a8285611126565b11156112705750905b80820391821161113357906111978161110e565b916111a56040519384610f79565b8183527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06111d28361110e565b013660208501376003549060005b8381106111ee575050505090565b6111f88183611126565b6000848210156112435780600360209252200173ffffffffffffffffffffffffffffffffffffffff60009154169086518310156112435750600582901b8601602001526001016111e0565b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526032600452fd5b61127b915082611126565b90611183565b505050604051611292602082610f79565b60008152600036813790565b73ffffffffffffffffffffffffffffffffffffffff16600052600860205260406000209060005260205260406000206020604051916112dc83610f2e565b549160ff63ffffffff841693848352831c1615159182910152611309575063ffffffff60065460201c1690565b90565b6000526007602052604060002060206040519161132883610f2e565b549160ff63ffffffff841693848352831c1615159182910152611309575063ffffffff6006541690565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b909260809273ffffffffffffffffffffffffffffffffffffffff6113099795168352602083015260408201528160608201520191611352565b91908110156113da5760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b3573ffffffffffffffffffffffffffffffffffffffff811681036101dd5790565b919290928242116116245761144c336000526004602052604060002054151590565b6115f657600060405160208101908282523360408201524660608201523060808201528560a08201528660c082015260c0815261148a60e082610f79565b5190207f19457468657265756d205369676e6564204d6573736167653a0a3332000000008252601c52603c81206114c084610fba565b906114ce6040519283610f79565b84825236858501116115f257906114f391858560208401378360208784010152611991565b90919260048310156115c557826115725750505073ffffffffffffffffffffffffffffffffffffffff1660009081526002602052604090205460ff161561153a5750505050565b90610824916040519485947f335d4ce10000000000000000000000000000000000000000000000000000000086523360048701611391565b8593505060ff6115b56040519586957fd36ab6b9000000000000000000000000000000000000000000000000000000008752606060048801526064870191611352565b9216602484015260448301520390fd5b807f4e487b7100000000000000000000000000000000000000000000000000000000602492526021600452fd5b8280fd5b7fd9a5f5ca000000000000000000000000000000000000000000000000000000006000523360045260246000fd5b827f502d038700000000000000000000000000000000000000000000000000000000600052336004524260245260445260646000fd5b73ffffffffffffffffffffffffffffffffffffffff60015416330361167b57565b7f2b5c74de0000000000000000000000000000000000000000000000000000000060005260046000fd5b916040519173ffffffffffffffffffffffffffffffffffffffff60208401946001865216604084015246606084015230608084015260a083015260c082015260c081526116f360e082610f79565b5190207f19457468657265756d205369676e6564204d6573736167653a0a333200000000600052601c52603c60002061172b83610fba565b6117386040519182610f79565b83815236848401116101dd5761175d9184846020840137600060208684010152611991565b60048295939510156117b45781611775575050505090565b60ff6115b56040519586957fd36ab6b9000000000000000000000000000000000000000000000000000000008752606060048801526064870191611352565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b80548210156113da5760005260206000200190600090565b600081815260046020526040902054801561198a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff810181811161113357600354907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82019182116111335781810361191b575b50505060035480156118ec577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016118a98160036117e3565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600355600052600460205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61197261192c61193d9360036117e3565b90549060031b1c92839260036117e3565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b90556000526004602052604060002055388080611870565b5050600090565b81519190604183036119c2576119bb92506020820151906060604084015193015160001a90611a2d565b9192909190565b505060009160029190565b80600052600460205260406000205415600014611a275760035468010000000000000000811015610f4a57611a0e61193d82600185940160035560036117e3565b9055600354906000526004602052604060002055600190565b50600090565b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411611ac3579160209360809260ff60009560405194855216868401526040830152606082015282805260015afa15611ab75760005173ffffffffffffffffffffffffffffffffffffffff811615611aab5790600090600090565b50600090600190600090565b6040513d6000823e3d90fd5b5050506000916003919056fea164736f6c634300081a000a",
}

var WorkflowRegistryABI = WorkflowRegistryMetaData.ABI

var WorkflowRegistryBin = WorkflowRegistryMetaData.Bin

func DeployWorkflowRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WorkflowRegistry, error) {
	parsed, err := WorkflowRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WorkflowRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WorkflowRegistry{address: address, abi: *parsed, WorkflowRegistryCaller: WorkflowRegistryCaller{contract: contract}, WorkflowRegistryTransactor: WorkflowRegistryTransactor{contract: contract}, WorkflowRegistryFilterer: WorkflowRegistryFilterer{contract: contract}}, nil
}

type WorkflowRegistry struct {
	address common.Address
	abi     abi.ABI
	WorkflowRegistryCaller
	WorkflowRegistryTransactor
	WorkflowRegistryFilterer
}

type WorkflowRegistryCaller struct {
	contract *bind.BoundContract
}

type WorkflowRegistryTransactor struct {
	contract *bind.BoundContract
}

type WorkflowRegistryFilterer struct {
	contract *bind.BoundContract
}

type WorkflowRegistrySession struct {
	Contract     *WorkflowRegistry
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type WorkflowRegistryCallerSession struct {
	Contract *WorkflowRegistryCaller
	CallOpts bind.CallOpts
}

type WorkflowRegistryTransactorSession struct {
	Contract     *WorkflowRegistryTransactor
	TransactOpts bind.TransactOpts
}

type WorkflowRegistryRaw struct {
	Contract *WorkflowRegistry
}

type WorkflowRegistryCallerRaw struct {
	Contract *WorkflowRegistryCaller
}

type WorkflowRegistryTransactorRaw struct {
	Contract *WorkflowRegistryTransactor
}

func NewWorkflowRegistry(address common.Address, backend bind.ContractBackend) (*WorkflowRegistry, error) {
	abi, err := abi.JSON(strings.NewReader(WorkflowRegistryABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindWorkflowRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistry{address: address, abi: abi, WorkflowRegistryCaller: WorkflowRegistryCaller{contract: contract}, WorkflowRegistryTransactor: WorkflowRegistryTransactor{contract: contract}, WorkflowRegistryFilterer: WorkflowRegistryFilterer{contract: contract}}, nil
}

func NewWorkflowRegistryCaller(address common.Address, caller bind.ContractCaller) (*WorkflowRegistryCaller, error) {
	contract, err := bindWorkflowRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryCaller{contract: contract}, nil
}

func NewWorkflowRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkflowRegistryTransactor, error) {
	contract, err := bindWorkflowRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryTransactor{contract: contract}, nil
}

func NewWorkflowRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkflowRegistryFilterer, error) {
	contract, err := bindWorkflowRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryFilterer{contract: contract}, nil
}

func bindWorkflowRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WorkflowRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_WorkflowRegistry *WorkflowRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkflowRegistry.Contract.WorkflowRegistryCaller.contract.Call(opts, result, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.WorkflowRegistryTransactor.contract.Transfer(opts)
}

func (_WorkflowRegistry *WorkflowRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.WorkflowRegistryTransactor.contract.Transact(opts, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkflowRegistry.Contract.contract.Call(opts, result, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.contract.Transfer(opts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.contract.Transact(opts, method, params...)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) CanLinkOwner(opts *bind.CallOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canLinkOwner", validityTimestamp, proof, signature)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanLinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanLinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) error {
	return _WorkflowRegistry.Contract.CanLinkOwner(&_WorkflowRegistry.CallOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte, arg4 uint8) error {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "canUnlinkOwner", owner, validityTimestamp, proof, signature, arg4)

	if err != nil {
		return err
	}

	return err

}

func (_WorkflowRegistry *WorkflowRegistrySession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte, arg4 uint8) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, proof, signature, arg4)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) CanUnlinkOwner(owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte, arg4 uint8) error {
	return _WorkflowRegistry.Contract.CanUnlinkOwner(&_WorkflowRegistry.CallOpts, owner, validityTimestamp, proof, signature, arg4)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetDefaults(opts *bind.CallOpts) (GetDefaults,

	error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getDefaults")

	outstruct := new(GetDefaults)
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxPerDON = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.MaxPerUserDON = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetDefaults() (GetDefaults,

	error) {
	return _WorkflowRegistry.Contract.GetDefaults(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetDefaults() (GetDefaults,

	error) {
	return _WorkflowRegistry.Contract.GetDefaults(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetLinkedOwners(opts *bind.CallOpts, start *big.Int, batchSize *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getLinkedOwners", start, batchSize)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetLinkedOwners(start *big.Int, batchSize *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetLinkedOwners(&_WorkflowRegistry.CallOpts, start, batchSize)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetLinkedOwners(start *big.Int, batchSize *big.Int) ([]common.Address, error) {
	return _WorkflowRegistry.Contract.GetLinkedOwners(&_WorkflowRegistry.CallOpts, start, batchSize)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerDON(opts *bind.CallOpts, donLabel [32]byte) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerDON", donLabel)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerDON(donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerDON(donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerDON(&_WorkflowRegistry.CallOpts, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donLabel [32]byte) (uint32, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "getMaxWorkflowsPerUserDON", user, donLabel)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) GetMaxWorkflowsPerUserDON(user common.Address, donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) GetMaxWorkflowsPerUserDON(user common.Address, donLabel [32]byte) (uint32, error) {
	return _WorkflowRegistry.Contract.GetMaxWorkflowsPerUserDON(&_WorkflowRegistry.CallOpts, user, donLabel)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "isAllowedSigner", signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) IsAllowedSigner(signer common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsAllowedSigner(&_WorkflowRegistry.CallOpts, signer)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) IsAllowedSigner(signer common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsAllowedSigner(&_WorkflowRegistry.CallOpts, signer)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "isOwnerLinked", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) IsOwnerLinked(owner common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsOwnerLinked(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) IsOwnerLinked(owner common.Address) (bool, error) {
	return _WorkflowRegistry.Contract.IsOwnerLinked(&_WorkflowRegistry.CallOpts, owner)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) Owner() (common.Address, error) {
	return _WorkflowRegistry.Contract.Owner(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) Owner() (common.Address, error) {
	return _WorkflowRegistry.Contract.Owner(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "totalLinkedOwners")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TotalLinkedOwners() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalLinkedOwners(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TotalLinkedOwners() (*big.Int, error) {
	return _WorkflowRegistry.Contract.TotalLinkedOwners(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WorkflowRegistry.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_WorkflowRegistry *WorkflowRegistrySession) TypeAndVersion() (string, error) {
	return _WorkflowRegistry.Contract.TypeAndVersion(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryCallerSession) TypeAndVersion() (string, error) {
	return _WorkflowRegistry.Contract.TypeAndVersion(&_WorkflowRegistry.CallOpts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "acceptOwnership")
}

func (_WorkflowRegistry *WorkflowRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AcceptOwnership(&_WorkflowRegistry.TransactOpts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.AcceptOwnership(&_WorkflowRegistry.TransactOpts)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "linkOwner", validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistrySession) LinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.LinkOwner(&_WorkflowRegistry.TransactOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) LinkOwner(validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.LinkOwner(&_WorkflowRegistry.TransactOpts, validityTimestamp, proof, signature)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDONOverride(opts *bind.TransactOpts, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDONOverride", donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDONOverride(donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONOverride(&_WorkflowRegistry.TransactOpts, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDONOverride(donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDONOverride(&_WorkflowRegistry.TransactOpts, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetDefaults(opts *bind.TransactOpts, maxPerDON uint32, maxPerUserDON uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setDefaults", maxPerDON, maxPerUserDON)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetDefaults(maxPerDON uint32, maxPerUserDON uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDefaults(&_WorkflowRegistry.TransactOpts, maxPerDON, maxPerUserDON)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetDefaults(maxPerDON uint32, maxPerUserDON uint32) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetDefaults(&_WorkflowRegistry.TransactOpts, maxPerDON, maxPerUserDON)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "setUserDONOverride", user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistrySession) SetUserDONOverride(user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) SetUserDONOverride(user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.SetUserDONOverride(&_WorkflowRegistry.TransactOpts, user, donLabel, limit, enabled)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "transferOwnership", to)
}

func (_WorkflowRegistry *WorkflowRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.TransferOwnership(&_WorkflowRegistry.TransactOpts, to)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.TransferOwnership(&_WorkflowRegistry.TransactOpts, to)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte, arg4 uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "unlinkOwner", owner, validityTimestamp, proof, signature, arg4)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte, arg4 uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, proof, signature, arg4)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UnlinkOwner(owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte, arg4 uint8) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UnlinkOwner(&_WorkflowRegistry.TransactOpts, owner, validityTimestamp, proof, signature, arg4)
}

func (_WorkflowRegistry *WorkflowRegistryTransactor) UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error) {
	return _WorkflowRegistry.contract.Transact(opts, "updateAllowedSigners", signers, allowed)
}

func (_WorkflowRegistry *WorkflowRegistrySession) UpdateAllowedSigners(signers []common.Address, allowed bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateAllowedSigners(&_WorkflowRegistry.TransactOpts, signers, allowed)
}

func (_WorkflowRegistry *WorkflowRegistryTransactorSession) UpdateAllowedSigners(signers []common.Address, allowed bool) (*types.Transaction, error) {
	return _WorkflowRegistry.Contract.UpdateAllowedSigners(&_WorkflowRegistry.TransactOpts, signers, allowed)
}

type WorkflowRegistryAllowedSignersUpdatedV1Iterator struct {
	Event *WorkflowRegistryAllowedSignersUpdatedV1

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryAllowedSignersUpdatedV1Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryAllowedSignersUpdatedV1)
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
		it.Event = new(WorkflowRegistryAllowedSignersUpdatedV1)
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

func (it *WorkflowRegistryAllowedSignersUpdatedV1Iterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryAllowedSignersUpdatedV1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryAllowedSignersUpdatedV1 struct {
	Signers []common.Address
	Allowed bool
	Raw     types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterAllowedSignersUpdatedV1(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedV1Iterator, error) {

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "AllowedSignersUpdatedV1")
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryAllowedSignersUpdatedV1Iterator{contract: _WorkflowRegistry.contract, event: "AllowedSignersUpdatedV1", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchAllowedSignersUpdatedV1(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdatedV1) (event.Subscription, error) {

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "AllowedSignersUpdatedV1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryAllowedSignersUpdatedV1)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "AllowedSignersUpdatedV1", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseAllowedSignersUpdatedV1(log types.Log) (*WorkflowRegistryAllowedSignersUpdatedV1, error) {
	event := new(WorkflowRegistryAllowedSignersUpdatedV1)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "AllowedSignersUpdatedV1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryOwnershipLinkUpdatedV1Iterator struct {
	Event *WorkflowRegistryOwnershipLinkUpdatedV1

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryOwnershipLinkUpdatedV1Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryOwnershipLinkUpdatedV1)
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
		it.Event = new(WorkflowRegistryOwnershipLinkUpdatedV1)
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

func (it *WorkflowRegistryOwnershipLinkUpdatedV1Iterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryOwnershipLinkUpdatedV1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryOwnershipLinkUpdatedV1 struct {
	Owner common.Address
	Proof [32]byte
	Added bool
	Raw   types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterOwnershipLinkUpdatedV1(opts *bind.FilterOpts, owner []common.Address, proof [][32]byte, added []bool) (*WorkflowRegistryOwnershipLinkUpdatedV1Iterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var proofRule []interface{}
	for _, proofItem := range proof {
		proofRule = append(proofRule, proofItem)
	}
	var addedRule []interface{}
	for _, addedItem := range added {
		addedRule = append(addedRule, addedItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "OwnershipLinkUpdatedV1", ownerRule, proofRule, addedRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryOwnershipLinkUpdatedV1Iterator{contract: _WorkflowRegistry.contract, event: "OwnershipLinkUpdatedV1", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchOwnershipLinkUpdatedV1(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipLinkUpdatedV1, owner []common.Address, proof [][32]byte, added []bool) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var proofRule []interface{}
	for _, proofItem := range proof {
		proofRule = append(proofRule, proofItem)
	}
	var addedRule []interface{}
	for _, addedItem := range added {
		addedRule = append(addedRule, addedItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "OwnershipLinkUpdatedV1", ownerRule, proofRule, addedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryOwnershipLinkUpdatedV1)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipLinkUpdatedV1", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseOwnershipLinkUpdatedV1(log types.Log) (*WorkflowRegistryOwnershipLinkUpdatedV1, error) {
	event := new(WorkflowRegistryOwnershipLinkUpdatedV1)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipLinkUpdatedV1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryOwnershipTransferRequestedIterator struct {
	Event *WorkflowRegistryOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryOwnershipTransferRequested)
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
		it.Event = new(WorkflowRegistryOwnershipTransferRequested)
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

func (it *WorkflowRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryOwnershipTransferRequestedIterator{contract: _WorkflowRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryOwnershipTransferRequested)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*WorkflowRegistryOwnershipTransferRequested, error) {
	event := new(WorkflowRegistryOwnershipTransferRequested)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type WorkflowRegistryOwnershipTransferredIterator struct {
	Event *WorkflowRegistryOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *WorkflowRegistryOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowRegistryOwnershipTransferred)
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
		it.Event = new(WorkflowRegistryOwnershipTransferred)
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

func (it *WorkflowRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *WorkflowRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type WorkflowRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowRegistryOwnershipTransferredIterator{contract: _WorkflowRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_WorkflowRegistry *WorkflowRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WorkflowRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(WorkflowRegistryOwnershipTransferred)
				if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_WorkflowRegistry *WorkflowRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*WorkflowRegistryOwnershipTransferred, error) {
	event := new(WorkflowRegistryOwnershipTransferred)
	if err := _WorkflowRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetDefaults struct {
	MaxPerDON     uint32
	MaxPerUserDON uint32
}

func (_WorkflowRegistry *WorkflowRegistry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _WorkflowRegistry.abi.Events["AllowedSignersUpdatedV1"].ID:
		return _WorkflowRegistry.ParseAllowedSignersUpdatedV1(log)
	case _WorkflowRegistry.abi.Events["OwnershipLinkUpdatedV1"].ID:
		return _WorkflowRegistry.ParseOwnershipLinkUpdatedV1(log)
	case _WorkflowRegistry.abi.Events["OwnershipTransferRequested"].ID:
		return _WorkflowRegistry.ParseOwnershipTransferRequested(log)
	case _WorkflowRegistry.abi.Events["OwnershipTransferred"].ID:
		return _WorkflowRegistry.ParseOwnershipTransferred(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (WorkflowRegistryAllowedSignersUpdatedV1) Topic() common.Hash {
	return common.HexToHash("0x7b317b86e2b8023639eee242d6f7a311def02336487b07a136bfe8f67b51b39c")
}

func (WorkflowRegistryOwnershipLinkUpdatedV1) Topic() common.Hash {
	return common.HexToHash("0x2195f84ef0d57d74c40638caac1ace3d9c91f7d4bd107849140f5437425ef65b")
}

func (WorkflowRegistryOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (WorkflowRegistryOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (_WorkflowRegistry *WorkflowRegistry) Address() common.Address {
	return _WorkflowRegistry.address
}

type WorkflowRegistryInterface interface {
	CanLinkOwner(opts *bind.CallOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) error

	CanUnlinkOwner(opts *bind.CallOpts, owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte, arg4 uint8) error

	GetDefaults(opts *bind.CallOpts) (GetDefaults,

		error)

	GetLinkedOwners(opts *bind.CallOpts, start *big.Int, batchSize *big.Int) ([]common.Address, error)

	GetMaxWorkflowsPerDON(opts *bind.CallOpts, donLabel [32]byte) (uint32, error)

	GetMaxWorkflowsPerUserDON(opts *bind.CallOpts, user common.Address, donLabel [32]byte) (uint32, error)

	IsAllowedSigner(opts *bind.CallOpts, signer common.Address) (bool, error)

	IsOwnerLinked(opts *bind.CallOpts, owner common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TotalLinkedOwners(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	LinkOwner(opts *bind.TransactOpts, validityTimestamp *big.Int, proof [32]byte, signature []byte) (*types.Transaction, error)

	SetDONOverride(opts *bind.TransactOpts, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error)

	SetDefaults(opts *bind.TransactOpts, maxPerDON uint32, maxPerUserDON uint32) (*types.Transaction, error)

	SetUserDONOverride(opts *bind.TransactOpts, user common.Address, donLabel [32]byte, limit uint32, enabled bool) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UnlinkOwner(opts *bind.TransactOpts, owner common.Address, validityTimestamp *big.Int, proof [32]byte, signature []byte, arg4 uint8) (*types.Transaction, error)

	UpdateAllowedSigners(opts *bind.TransactOpts, signers []common.Address, allowed bool) (*types.Transaction, error)

	FilterAllowedSignersUpdatedV1(opts *bind.FilterOpts) (*WorkflowRegistryAllowedSignersUpdatedV1Iterator, error)

	WatchAllowedSignersUpdatedV1(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryAllowedSignersUpdatedV1) (event.Subscription, error)

	ParseAllowedSignersUpdatedV1(log types.Log) (*WorkflowRegistryAllowedSignersUpdatedV1, error)

	FilterOwnershipLinkUpdatedV1(opts *bind.FilterOpts, owner []common.Address, proof [][32]byte, added []bool) (*WorkflowRegistryOwnershipLinkUpdatedV1Iterator, error)

	WatchOwnershipLinkUpdatedV1(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipLinkUpdatedV1, owner []common.Address, proof [][32]byte, added []bool) (event.Subscription, error)

	ParseOwnershipLinkUpdatedV1(log types.Log) (*WorkflowRegistryOwnershipLinkUpdatedV1, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*WorkflowRegistryOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WorkflowRegistryOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkflowRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*WorkflowRegistryOwnershipTransferred, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
