// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimism_sequencer_uptime_feed

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

var OptimismSequencerUptimeFeedMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"l1SenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l2CrossDomainMessengerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialStatus\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"description\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAnswer\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoundData\",\"inputs\":[{\"name\":\"_roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"outputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"answeredInRound\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimestamp\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1Sender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestAnswer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRoundData\",\"inputs\":[],\"outputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"answeredInRound\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferL1Sender\",\"inputs\":[{\"name\":\"newSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AnswerUpdated\",\"inputs\":[{\"name\":\"current\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"roundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1SenderTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"startedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoundUpdated\",\"inputs\":[{\"name\":\"status\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"updatedAt\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateIgnored\",\"inputs\":[{\"name\":\"latestStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"latestTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"incomingStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"incomingTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoDataPresent\",\"inputs\":[]}]",
	Bin: "0x6101206040525f60a081905260c081905260e081905261010052600480546001600160d81b031916905534801562000035575f80fd5b5060405162001ea138038062001ea18339810160408190526200005891620003eb565b828133805f81620000b05760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000e257620000e28162000127565b50506001805460ff60a01b1916600160a01b179055506200010382620001d1565b62000111600182426200023a565b5050506001600160a01b03166080525062000439565b336001600160a01b03821603620001815760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401620000a7565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6003546001600160a01b0390811690821681146200023657600380546001600160a01b0319166001600160a01b0384811691821790925560405190918316907f8e6da65f164d652f378f48652c0e1ca58d7c9cc52ceaa40c1dad055cd7681d18905f90a35b5050565b60408051606080820183526001600160401b0384811680845242821660208086018281528915158789018181526001600160501b038d165f818152600586528b902099518a54945192511515600160801b0260ff60801b19938a1668010000000000000000026001600160801b03199096169190991617939093171695909517909655865160808101885286815280820184905280880183905290940183905260048054600160d01b90940260ff60d01b19600160901b90930292909216600160901b600160d81b03196a010000000000000000000085026001600160901b0319909616881795909517949094169390931717909155925192835233927f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271910160405180910390a36001600160501b0383166200037783620003b8565b6040516001600160401b03841681527f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f9060200160405180910390a3505050565b5f81620003c6575f620003c9565b60015b92915050565b80516001600160a01b0381168114620003e6575f80fd5b919050565b5f805f60608486031215620003fe575f80fd5b6200040984620003cf565b92506200041960208501620003cf565b9150604084015180151581146200042e575f80fd5b809150509250925092565b608051611a48620004595f395f818161124601526112870152611a485ff3fe608060405234801561000f575f80fd5b5060043610610184575f3560e01c80638205bf6a116100dd578063b633620c11610088578063ed8378f511610063578063ed8378f514610390578063f2fde38b146103a3578063feaf968c146103b6575f80fd5b8063b633620c1461033a578063b7558b7a1461034d578063dc7f01241461036b575f80fd5b80639a6fc8f5116100b85780639a6fc8f5146102ca578063a118f24914610314578063b5ab58dc14610327575f80fd5b80638205bf6a146102715780638823da6c146102795780638da5cb5b1461028c575f80fd5b806354fd4d501161013d5780637284e416116101185780637284e4161461022557806379ba5097146102615780638038e4a114610269575f80fd5b806354fd4d50146101f2578063668a0f02146101fa5780636b14daf814610202575f80fd5b8063284afc081161016d578063284afc08146101b0578063313ce567146101c357806350d25bcd146101dc575f80fd5b80630a75698314610188578063181f5a7714610192575b5f80fd5b6101906103be565b005b61019a61043c565b6040516101a79190611776565b60405180910390f35b6101906101be366004611801565b610458565b6101ca5f81565b60405160ff90911681526020016101a7565b6101e461046c565b6040519081526020016101a7565b6101e4600181565b6101e4610549565b610215610210366004611849565b610602565b60405190151581526020016101a7565b61019a6040518060400160405280601f81526020017f4c322053657175656e63657220557074696d652053746174757320466565640081525081565b610190610634565b610190610730565b6101e46107c3565b610190610287366004611801565b610888565b5f5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a7565b6102dd6102d8366004611925565b61093f565b6040805169ffffffffffffffffffff968716815260208101959095528401929092526060830152909116608082015260a0016101a7565b610190610322366004611801565b610a9f565b6101e461033536600461194e565b610b51565b6101e461034836600461194e565b610c77565b60035473ffffffffffffffffffffffffffffffffffffffff166102a5565b6001546102159074010000000000000000000000000000000000000000900460ff1681565b61019061039e366004611965565b610d84565b6101906103b1366004611801565b610f32565b6102dd610f43565b6103c6611076565b60015474010000000000000000000000000000000000000000900460ff161561043a57600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638905f90a15b565b604051806060016040528060258152602001611a176025913981565b610460611076565b610469816110f6565b50565b5f6104ac335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061060292505050565b610517576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f20616363657373000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600454610544907a010000000000000000000000000000000000000000000000000000900460ff16611190565b905090565b5f610589335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061060292505050565b6105ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f206163636573730000000000000000000000000000000000000000000000604482015260640161050e565b5060045469ffffffffffffffffffff1690565b5f61060d83836111a4565b8061062d575073ffffffffffffffffffffffffffffffffffffffff831632145b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146106b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e657200000000000000000000604482015260640161050e565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610738611076565b60015474010000000000000000000000000000000000000000900460ff1661043a57600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517faebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480905f90a1565b5f610803335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061060292505050565b610869576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f206163636573730000000000000000000000000000000000000000000000604482015260640161050e565b506004546a0100000000000000000000900467ffffffffffffffff1690565b610890611076565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff16156104695773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527f3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d191015b60405180910390a150565b5f805f805f610983335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061060292505050565b6109e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f206163636573730000000000000000000000000000000000000000000000604482015260640161050e565b6109fe8669ffffffffffffffffffff166111f8565b610a34576040517fbb25870000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b69ffffffffffffffffffff86165f90815260056020526040902080548790610a7290700100000000000000000000000000000000900460ff16611190565b91549098919767ffffffffffffffff80831698506801000000000000000090920490911695509350915050565b610aa7611076565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff166104695773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527f87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db49101610934565b5f610b91335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061060292505050565b610bf7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f206163636573730000000000000000000000000000000000000000000000604482015260640161050e565b610c00826111f8565b610c36576040517fbb25870000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b69ffffffffffffffffffff82165f90815260056020526040902054610c7190700100000000000000000000000000000000900460ff16611190565b92915050565b5f610cb7335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061060292505050565b610d1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f206163636573730000000000000000000000000000000000000000000000604482015260640161050e565b610d26826111f8565b610d5c576040517fbb25870000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5069ffffffffffffffffffff165f9081526005602052604090205467ffffffffffffffff1690565b600354610da69073ffffffffffffffffffffffffffffffffffffffff1661122e565b5f610e53604080516080810182525f808252602082018190529181018290526060810191909152506040805160808101825260045469ffffffffffffffffffff811682526a0100000000000000000000810467ffffffffffffffff90811660208401527201000000000000000000000000000000000000820416928201929092527a01000000000000000000000000000000000000000000000000000090910460ff161515606082015290565b90508167ffffffffffffffff16816020015167ffffffffffffffff161115610ee4577fe4a6e16301740042c17431042adb8f60454c18fb5934dd4c456269c0dc388fdf816060015182602001518585604051610ed79493929190931515845267ffffffffffffffff9283166020850152901515604084015216606082015260800190565b60405180910390a1505050565b8215158160600151151503610f04578051610eff9084611362565b505050565b6001815f01818151610f1691906119ac565b69ffffffffffffffffffff169052508051610eff908484611450565b610f3a611076565b61046981611682565b5f805f805f610f87335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061060292505050565b610fed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f206163636573730000000000000000000000000000000000000000000000604482015260640161050e565b6004805469ffffffffffffffffffff811690611029907a010000000000000000000000000000000000000000000000000000900460ff16611190565b9154909791965067ffffffffffffffff6a01000000000000000000008204811696507201000000000000000000000000000000000000820416945069ffffffffffffffffffff1692509050565b5f5473ffffffffffffffffffffffffffffffffffffffff16331461043a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161050e565b60035473ffffffffffffffffffffffffffffffffffffffff908116908216811461118c57600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691821790925560405190918316907f8e6da65f164d652f378f48652c0e1ca58d7c9cc52ceaa40c1dad055cd7681d18905f90a35b5050565b5f8161119c575f610c71565b600192915050565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526002602052604081205460ff168061062d57505060015474010000000000000000000000000000000000000000900460ff161592915050565b5f8082118015611212575069ffffffffffffffffffff8211155b8015610c7157505060045469ffffffffffffffffffff16101590565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614158061132b57508073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112ee573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061131291906119fb565b73ffffffffffffffffffffffffffffffffffffffff1614155b15610469576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b69ffffffffffffffffffff82165f90815260056020526040902080547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff16680100000000000000004267ffffffffffffffff1690810291909117909155600480547fffffffffffff0000000000000000ffffffffffffffffffffffffffffffffffff1672010000000000000000000000000000000000009092029190911790557f297642343ed2faefb1a411b39fc449eae700e54223d5d0499a9421eb6f68f66a61142c82611190565b6040805191825267ffffffffffffffff421660208301520160405180910390a15050565b604080516060808201835267ffffffffffffffff848116808452428216602080860182815289151587890181815269ffffffffffffffffffff8d165f818152600586528b902099518a54945192511515700100000000000000000000000000000000027fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff938a1668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000000000000000000090961691909916179390931716959095179096558651608081018852868152808201849052808801839052909401839052600480547a0100000000000000000000000000000000000000000000000000009094027fffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffff7201000000000000000000000000000000000000909302929092167fffffffffff000000000000000000ffffffffffffffffffffffffffffffffffff6a010000000000000000000085027fffffffffffffffffffffffffffff000000000000000000000000000000000000909616881795909517949094169390931717909155925192835233927f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271910160405180910390a38269ffffffffffffffffffff1661164083611190565b60405167ffffffffffffffff841681527f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f9060200160405180910390a3505050565b3373ffffffffffffffffffffffffffffffffffffffff821603611701576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161050e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b5f602080835283518060208501525f5b818110156117a257858101830151858201604001528201611786565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b73ffffffffffffffffffffffffffffffffffffffff81168114610469575f80fd5b5f60208284031215611811575f80fd5b813561062d816117e0565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f806040838503121561185a575f80fd5b8235611865816117e0565b9150602083013567ffffffffffffffff80821115611881575f80fd5b818501915085601f830112611894575f80fd5b8135818111156118a6576118a661181c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156118ec576118ec61181c565b81604052828152886020848701011115611904575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215611935575f80fd5b813569ffffffffffffffffffff8116811461062d575f80fd5b5f6020828403121561195e575f80fd5b5035919050565b5f8060408385031215611976575f80fd5b82358015158114611985575f80fd5b9150602083013567ffffffffffffffff811681146119a1575f80fd5b809150509250929050565b69ffffffffffffffffffff8181168382160190808211156119f4577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5092915050565b5f60208284031215611a0b575f80fd5b815161062d816117e056fe4f7074696d69736d53657175656e636572557074696d654665656420312e312e302d646576a164736f6c6343000818000a",
}

var OptimismSequencerUptimeFeedABI = OptimismSequencerUptimeFeedMetaData.ABI

var OptimismSequencerUptimeFeedBin = OptimismSequencerUptimeFeedMetaData.Bin

func DeployOptimismSequencerUptimeFeed(auth *bind.TransactOpts, backend bind.ContractBackend, l1SenderAddress common.Address, l2CrossDomainMessengerAddr common.Address, initialStatus bool) (common.Address, *types.Transaction, *OptimismSequencerUptimeFeed, error) {
	parsed, err := OptimismSequencerUptimeFeedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptimismSequencerUptimeFeedBin), backend, l1SenderAddress, l2CrossDomainMessengerAddr, initialStatus)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptimismSequencerUptimeFeed{address: address, abi: *parsed, OptimismSequencerUptimeFeedCaller: OptimismSequencerUptimeFeedCaller{contract: contract}, OptimismSequencerUptimeFeedTransactor: OptimismSequencerUptimeFeedTransactor{contract: contract}, OptimismSequencerUptimeFeedFilterer: OptimismSequencerUptimeFeedFilterer{contract: contract}}, nil
}

type OptimismSequencerUptimeFeed struct {
	address common.Address
	abi     abi.ABI
	OptimismSequencerUptimeFeedCaller
	OptimismSequencerUptimeFeedTransactor
	OptimismSequencerUptimeFeedFilterer
}

type OptimismSequencerUptimeFeedCaller struct {
	contract *bind.BoundContract
}

type OptimismSequencerUptimeFeedTransactor struct {
	contract *bind.BoundContract
}

type OptimismSequencerUptimeFeedFilterer struct {
	contract *bind.BoundContract
}

type OptimismSequencerUptimeFeedSession struct {
	Contract     *OptimismSequencerUptimeFeed
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type OptimismSequencerUptimeFeedCallerSession struct {
	Contract *OptimismSequencerUptimeFeedCaller
	CallOpts bind.CallOpts
}

type OptimismSequencerUptimeFeedTransactorSession struct {
	Contract     *OptimismSequencerUptimeFeedTransactor
	TransactOpts bind.TransactOpts
}

type OptimismSequencerUptimeFeedRaw struct {
	Contract *OptimismSequencerUptimeFeed
}

type OptimismSequencerUptimeFeedCallerRaw struct {
	Contract *OptimismSequencerUptimeFeedCaller
}

type OptimismSequencerUptimeFeedTransactorRaw struct {
	Contract *OptimismSequencerUptimeFeedTransactor
}

func NewOptimismSequencerUptimeFeed(address common.Address, backend bind.ContractBackend) (*OptimismSequencerUptimeFeed, error) {
	abi, err := abi.JSON(strings.NewReader(OptimismSequencerUptimeFeedABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindOptimismSequencerUptimeFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeed{address: address, abi: abi, OptimismSequencerUptimeFeedCaller: OptimismSequencerUptimeFeedCaller{contract: contract}, OptimismSequencerUptimeFeedTransactor: OptimismSequencerUptimeFeedTransactor{contract: contract}, OptimismSequencerUptimeFeedFilterer: OptimismSequencerUptimeFeedFilterer{contract: contract}}, nil
}

func NewOptimismSequencerUptimeFeedCaller(address common.Address, caller bind.ContractCaller) (*OptimismSequencerUptimeFeedCaller, error) {
	contract, err := bindOptimismSequencerUptimeFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedCaller{contract: contract}, nil
}

func NewOptimismSequencerUptimeFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimismSequencerUptimeFeedTransactor, error) {
	contract, err := bindOptimismSequencerUptimeFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedTransactor{contract: contract}, nil
}

func NewOptimismSequencerUptimeFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimismSequencerUptimeFeedFilterer, error) {
	contract, err := bindOptimismSequencerUptimeFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedFilterer{contract: contract}, nil
}

func bindOptimismSequencerUptimeFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimismSequencerUptimeFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismSequencerUptimeFeed.Contract.OptimismSequencerUptimeFeedCaller.contract.Call(opts, result, method, params...)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.OptimismSequencerUptimeFeedTransactor.contract.Transfer(opts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.OptimismSequencerUptimeFeedTransactor.contract.Transact(opts, method, params...)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimismSequencerUptimeFeed.Contract.contract.Call(opts, result, method, params...)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.contract.Transfer(opts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.contract.Transact(opts, method, params...)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "checkEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) CheckEnabled() (bool, error) {
	return _OptimismSequencerUptimeFeed.Contract.CheckEnabled(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) CheckEnabled() (bool, error) {
	return _OptimismSequencerUptimeFeed.Contract.CheckEnabled(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) Decimals() (uint8, error) {
	return _OptimismSequencerUptimeFeed.Contract.Decimals(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) Decimals() (uint8, error) {
	return _OptimismSequencerUptimeFeed.Contract.Decimals(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) Description() (string, error) {
	return _OptimismSequencerUptimeFeed.Contract.Description(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) Description() (string, error) {
	return _OptimismSequencerUptimeFeed.Contract.Description(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "getAnswer", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.GetAnswer(&_OptimismSequencerUptimeFeed.CallOpts, roundId)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.GetAnswer(&_OptimismSequencerUptimeFeed.CallOpts, roundId)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (GetRoundData,

	error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(GetRoundData)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) GetRoundData(_roundId *big.Int) (GetRoundData,

	error) {
	return _OptimismSequencerUptimeFeed.Contract.GetRoundData(&_OptimismSequencerUptimeFeed.CallOpts, _roundId)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) GetRoundData(_roundId *big.Int) (GetRoundData,

	error) {
	return _OptimismSequencerUptimeFeed.Contract.GetRoundData(&_OptimismSequencerUptimeFeed.CallOpts, _roundId)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "getTimestamp", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.GetTimestamp(&_OptimismSequencerUptimeFeed.CallOpts, roundId)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.GetTimestamp(&_OptimismSequencerUptimeFeed.CallOpts, roundId)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "hasAccess", _user, _calldata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _OptimismSequencerUptimeFeed.Contract.HasAccess(&_OptimismSequencerUptimeFeed.CallOpts, _user, _calldata)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _OptimismSequencerUptimeFeed.Contract.HasAccess(&_OptimismSequencerUptimeFeed.CallOpts, _user, _calldata)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) L1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "l1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) L1Sender() (common.Address, error) {
	return _OptimismSequencerUptimeFeed.Contract.L1Sender(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) L1Sender() (common.Address, error) {
	return _OptimismSequencerUptimeFeed.Contract.L1Sender(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) LatestAnswer() (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.LatestAnswer(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) LatestAnswer() (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.LatestAnswer(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) LatestRound() (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.LatestRound(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) LatestRound() (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.LatestRound(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) LatestRoundData(opts *bind.CallOpts) (LatestRoundData,

	error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(LatestRoundData)
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) LatestRoundData() (LatestRoundData,

	error) {
	return _OptimismSequencerUptimeFeed.Contract.LatestRoundData(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) LatestRoundData() (LatestRoundData,

	error) {
	return _OptimismSequencerUptimeFeed.Contract.LatestRoundData(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) LatestTimestamp() (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.LatestTimestamp(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) LatestTimestamp() (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.LatestTimestamp(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) Owner() (common.Address, error) {
	return _OptimismSequencerUptimeFeed.Contract.Owner(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) Owner() (common.Address, error) {
	return _OptimismSequencerUptimeFeed.Contract.Owner(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) TypeAndVersion() (string, error) {
	return _OptimismSequencerUptimeFeed.Contract.TypeAndVersion(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) TypeAndVersion() (string, error) {
	return _OptimismSequencerUptimeFeed.Contract.TypeAndVersion(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptimismSequencerUptimeFeed.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) Version() (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.Version(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedCallerSession) Version() (*big.Int, error) {
	return _OptimismSequencerUptimeFeed.Contract.Version(&_OptimismSequencerUptimeFeed.CallOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.contract.Transact(opts, "acceptOwnership")
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) AcceptOwnership() (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.AcceptOwnership(&_OptimismSequencerUptimeFeed.TransactOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.AcceptOwnership(&_OptimismSequencerUptimeFeed.TransactOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.contract.Transact(opts, "addAccess", _user)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.AddAccess(&_OptimismSequencerUptimeFeed.TransactOpts, _user)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.AddAccess(&_OptimismSequencerUptimeFeed.TransactOpts, _user)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.contract.Transact(opts, "disableAccessCheck")
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) DisableAccessCheck() (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.DisableAccessCheck(&_OptimismSequencerUptimeFeed.TransactOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.DisableAccessCheck(&_OptimismSequencerUptimeFeed.TransactOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.contract.Transact(opts, "enableAccessCheck")
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) EnableAccessCheck() (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.EnableAccessCheck(&_OptimismSequencerUptimeFeed.TransactOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.EnableAccessCheck(&_OptimismSequencerUptimeFeed.TransactOpts)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.contract.Transact(opts, "removeAccess", _user)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.RemoveAccess(&_OptimismSequencerUptimeFeed.TransactOpts, _user)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.RemoveAccess(&_OptimismSequencerUptimeFeed.TransactOpts, _user)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactor) TransferL1Sender(opts *bind.TransactOpts, newSender common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.contract.Transact(opts, "transferL1Sender", newSender)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) TransferL1Sender(newSender common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.TransferL1Sender(&_OptimismSequencerUptimeFeed.TransactOpts, newSender)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorSession) TransferL1Sender(newSender common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.TransferL1Sender(&_OptimismSequencerUptimeFeed.TransactOpts, newSender)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.contract.Transact(opts, "transferOwnership", to)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.TransferOwnership(&_OptimismSequencerUptimeFeed.TransactOpts, to)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.TransferOwnership(&_OptimismSequencerUptimeFeed.TransactOpts, to)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactor) UpdateStatus(opts *bind.TransactOpts, status bool, timestamp uint64) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.contract.Transact(opts, "updateStatus", status, timestamp)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedSession) UpdateStatus(status bool, timestamp uint64) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.UpdateStatus(&_OptimismSequencerUptimeFeed.TransactOpts, status, timestamp)
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedTransactorSession) UpdateStatus(status bool, timestamp uint64) (*types.Transaction, error) {
	return _OptimismSequencerUptimeFeed.Contract.UpdateStatus(&_OptimismSequencerUptimeFeed.TransactOpts, status, timestamp)
}

type OptimismSequencerUptimeFeedAddedAccessIterator struct {
	Event *OptimismSequencerUptimeFeedAddedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedAddedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedAddedAccess)
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
		it.Event = new(OptimismSequencerUptimeFeedAddedAccess)
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

func (it *OptimismSequencerUptimeFeedAddedAccessIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedAddedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedAddedAccessIterator, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedAddedAccessIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedAddedAccess) (event.Subscription, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedAddedAccess)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "AddedAccess", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseAddedAccess(log types.Log) (*OptimismSequencerUptimeFeedAddedAccess, error) {
	event := new(OptimismSequencerUptimeFeedAddedAccess)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedAnswerUpdatedIterator struct {
	Event *OptimismSequencerUptimeFeedAnswerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedAnswerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedAnswerUpdated)
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
		it.Event = new(OptimismSequencerUptimeFeedAnswerUpdated)
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

func (it *OptimismSequencerUptimeFeedAnswerUpdatedIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*OptimismSequencerUptimeFeedAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedAnswerUpdatedIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedAnswerUpdated)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseAnswerUpdated(log types.Log) (*OptimismSequencerUptimeFeedAnswerUpdated, error) {
	event := new(OptimismSequencerUptimeFeedAnswerUpdated)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedCheckAccessDisabledIterator struct {
	Event *OptimismSequencerUptimeFeedCheckAccessDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedCheckAccessDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedCheckAccessDisabled)
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
		it.Event = new(OptimismSequencerUptimeFeedCheckAccessDisabled)
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

func (it *OptimismSequencerUptimeFeedCheckAccessDisabledIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedCheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedCheckAccessDisabled struct {
	Raw types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedCheckAccessDisabledIterator, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedCheckAccessDisabledIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedCheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedCheckAccessDisabled)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseCheckAccessDisabled(log types.Log) (*OptimismSequencerUptimeFeedCheckAccessDisabled, error) {
	event := new(OptimismSequencerUptimeFeedCheckAccessDisabled)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedCheckAccessEnabledIterator struct {
	Event *OptimismSequencerUptimeFeedCheckAccessEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedCheckAccessEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedCheckAccessEnabled)
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
		it.Event = new(OptimismSequencerUptimeFeedCheckAccessEnabled)
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

func (it *OptimismSequencerUptimeFeedCheckAccessEnabledIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedCheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedCheckAccessEnabled struct {
	Raw types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedCheckAccessEnabledIterator, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedCheckAccessEnabledIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedCheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedCheckAccessEnabled)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseCheckAccessEnabled(log types.Log) (*OptimismSequencerUptimeFeedCheckAccessEnabled, error) {
	event := new(OptimismSequencerUptimeFeedCheckAccessEnabled)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedL1SenderTransferredIterator struct {
	Event *OptimismSequencerUptimeFeedL1SenderTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedL1SenderTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedL1SenderTransferred)
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
		it.Event = new(OptimismSequencerUptimeFeedL1SenderTransferred)
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

func (it *OptimismSequencerUptimeFeedL1SenderTransferredIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedL1SenderTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedL1SenderTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterL1SenderTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismSequencerUptimeFeedL1SenderTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "L1SenderTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedL1SenderTransferredIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "L1SenderTransferred", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchL1SenderTransferred(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedL1SenderTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "L1SenderTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedL1SenderTransferred)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "L1SenderTransferred", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseL1SenderTransferred(log types.Log) (*OptimismSequencerUptimeFeedL1SenderTransferred, error) {
	event := new(OptimismSequencerUptimeFeedL1SenderTransferred)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "L1SenderTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedNewRoundIterator struct {
	Event *OptimismSequencerUptimeFeedNewRound

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedNewRoundIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedNewRound)
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
		it.Event = new(OptimismSequencerUptimeFeedNewRound)
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

func (it *OptimismSequencerUptimeFeedNewRoundIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*OptimismSequencerUptimeFeedNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedNewRoundIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedNewRound)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "NewRound", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseNewRound(log types.Log) (*OptimismSequencerUptimeFeedNewRound, error) {
	event := new(OptimismSequencerUptimeFeedNewRound)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedOwnershipTransferRequestedIterator struct {
	Event *OptimismSequencerUptimeFeedOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedOwnershipTransferRequested)
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
		it.Event = new(OptimismSequencerUptimeFeedOwnershipTransferRequested)
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

func (it *OptimismSequencerUptimeFeedOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismSequencerUptimeFeedOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedOwnershipTransferRequestedIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedOwnershipTransferRequested)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseOwnershipTransferRequested(log types.Log) (*OptimismSequencerUptimeFeedOwnershipTransferRequested, error) {
	event := new(OptimismSequencerUptimeFeedOwnershipTransferRequested)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedOwnershipTransferredIterator struct {
	Event *OptimismSequencerUptimeFeedOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedOwnershipTransferred)
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
		it.Event = new(OptimismSequencerUptimeFeedOwnershipTransferred)
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

func (it *OptimismSequencerUptimeFeedOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismSequencerUptimeFeedOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedOwnershipTransferredIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedOwnershipTransferred)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseOwnershipTransferred(log types.Log) (*OptimismSequencerUptimeFeedOwnershipTransferred, error) {
	event := new(OptimismSequencerUptimeFeedOwnershipTransferred)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedRemovedAccessIterator struct {
	Event *OptimismSequencerUptimeFeedRemovedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedRemovedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedRemovedAccess)
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
		it.Event = new(OptimismSequencerUptimeFeedRemovedAccess)
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

func (it *OptimismSequencerUptimeFeedRemovedAccessIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedRemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedRemovedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedRemovedAccessIterator, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedRemovedAccessIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedRemovedAccess) (event.Subscription, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedRemovedAccess)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseRemovedAccess(log types.Log) (*OptimismSequencerUptimeFeedRemovedAccess, error) {
	event := new(OptimismSequencerUptimeFeedRemovedAccess)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedRoundUpdatedIterator struct {
	Event *OptimismSequencerUptimeFeedRoundUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedRoundUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedRoundUpdated)
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
		it.Event = new(OptimismSequencerUptimeFeedRoundUpdated)
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

func (it *OptimismSequencerUptimeFeedRoundUpdatedIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedRoundUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedRoundUpdated struct {
	Status    *big.Int
	UpdatedAt uint64
	Raw       types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterRoundUpdated(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedRoundUpdatedIterator, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "RoundUpdated")
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedRoundUpdatedIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "RoundUpdated", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchRoundUpdated(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedRoundUpdated) (event.Subscription, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "RoundUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedRoundUpdated)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "RoundUpdated", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseRoundUpdated(log types.Log) (*OptimismSequencerUptimeFeedRoundUpdated, error) {
	event := new(OptimismSequencerUptimeFeedRoundUpdated)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "RoundUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type OptimismSequencerUptimeFeedUpdateIgnoredIterator struct {
	Event *OptimismSequencerUptimeFeedUpdateIgnored

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *OptimismSequencerUptimeFeedUpdateIgnoredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OptimismSequencerUptimeFeedUpdateIgnored)
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
		it.Event = new(OptimismSequencerUptimeFeedUpdateIgnored)
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

func (it *OptimismSequencerUptimeFeedUpdateIgnoredIterator) Error() error {
	return it.fail
}

func (it *OptimismSequencerUptimeFeedUpdateIgnoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type OptimismSequencerUptimeFeedUpdateIgnored struct {
	LatestStatus      bool
	LatestTimestamp   uint64
	IncomingStatus    bool
	IncomingTimestamp uint64
	Raw               types.Log
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) FilterUpdateIgnored(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedUpdateIgnoredIterator, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.FilterLogs(opts, "UpdateIgnored")
	if err != nil {
		return nil, err
	}
	return &OptimismSequencerUptimeFeedUpdateIgnoredIterator{contract: _OptimismSequencerUptimeFeed.contract, event: "UpdateIgnored", logs: logs, sub: sub}, nil
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) WatchUpdateIgnored(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedUpdateIgnored) (event.Subscription, error) {

	logs, sub, err := _OptimismSequencerUptimeFeed.contract.WatchLogs(opts, "UpdateIgnored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(OptimismSequencerUptimeFeedUpdateIgnored)
				if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "UpdateIgnored", log); err != nil {
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

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeedFilterer) ParseUpdateIgnored(log types.Log) (*OptimismSequencerUptimeFeedUpdateIgnored, error) {
	event := new(OptimismSequencerUptimeFeedUpdateIgnored)
	if err := _OptimismSequencerUptimeFeed.contract.UnpackLog(event, "UpdateIgnored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type GetRoundData struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}
type LatestRoundData struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}

func (OptimismSequencerUptimeFeedAddedAccess) Topic() common.Hash {
	return common.HexToHash("0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4")
}

func (OptimismSequencerUptimeFeedAnswerUpdated) Topic() common.Hash {
	return common.HexToHash("0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f")
}

func (OptimismSequencerUptimeFeedCheckAccessDisabled) Topic() common.Hash {
	return common.HexToHash("0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638")
}

func (OptimismSequencerUptimeFeedCheckAccessEnabled) Topic() common.Hash {
	return common.HexToHash("0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480")
}

func (OptimismSequencerUptimeFeedL1SenderTransferred) Topic() common.Hash {
	return common.HexToHash("0x8e6da65f164d652f378f48652c0e1ca58d7c9cc52ceaa40c1dad055cd7681d18")
}

func (OptimismSequencerUptimeFeedNewRound) Topic() common.Hash {
	return common.HexToHash("0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271")
}

func (OptimismSequencerUptimeFeedOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (OptimismSequencerUptimeFeedOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (OptimismSequencerUptimeFeedRemovedAccess) Topic() common.Hash {
	return common.HexToHash("0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1")
}

func (OptimismSequencerUptimeFeedRoundUpdated) Topic() common.Hash {
	return common.HexToHash("0x297642343ed2faefb1a411b39fc449eae700e54223d5d0499a9421eb6f68f66a")
}

func (OptimismSequencerUptimeFeedUpdateIgnored) Topic() common.Hash {
	return common.HexToHash("0xe4a6e16301740042c17431042adb8f60454c18fb5934dd4c456269c0dc388fdf")
}

func (_OptimismSequencerUptimeFeed *OptimismSequencerUptimeFeed) Address() common.Address {
	return _OptimismSequencerUptimeFeed.address
}

type OptimismSequencerUptimeFeedInterface interface {
	CheckEnabled(opts *bind.CallOpts) (bool, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	Description(opts *bind.CallOpts) (string, error)

	GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error)

	GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (GetRoundData,

		error)

	GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error)

	HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error)

	L1Sender(opts *bind.CallOpts) (common.Address, error)

	LatestAnswer(opts *bind.CallOpts) (*big.Int, error)

	LatestRound(opts *bind.CallOpts) (*big.Int, error)

	LatestRoundData(opts *bind.CallOpts) (LatestRoundData,

		error)

	LatestTimestamp(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Version(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error)

	DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error)

	EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error)

	TransferL1Sender(opts *bind.TransactOpts, newSender common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateStatus(opts *bind.TransactOpts, status bool, timestamp uint64) (*types.Transaction, error)

	FilterAddedAccess(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedAddedAccessIterator, error)

	WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedAddedAccess) (event.Subscription, error)

	ParseAddedAccess(log types.Log) (*OptimismSequencerUptimeFeedAddedAccess, error)

	FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*OptimismSequencerUptimeFeedAnswerUpdatedIterator, error)

	WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error)

	ParseAnswerUpdated(log types.Log) (*OptimismSequencerUptimeFeedAnswerUpdated, error)

	FilterCheckAccessDisabled(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedCheckAccessDisabledIterator, error)

	WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedCheckAccessDisabled) (event.Subscription, error)

	ParseCheckAccessDisabled(log types.Log) (*OptimismSequencerUptimeFeedCheckAccessDisabled, error)

	FilterCheckAccessEnabled(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedCheckAccessEnabledIterator, error)

	WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedCheckAccessEnabled) (event.Subscription, error)

	ParseCheckAccessEnabled(log types.Log) (*OptimismSequencerUptimeFeedCheckAccessEnabled, error)

	FilterL1SenderTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismSequencerUptimeFeedL1SenderTransferredIterator, error)

	WatchL1SenderTransferred(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedL1SenderTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1SenderTransferred(log types.Log) (*OptimismSequencerUptimeFeedL1SenderTransferred, error)

	FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*OptimismSequencerUptimeFeedNewRoundIterator, error)

	WatchNewRound(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error)

	ParseNewRound(log types.Log) (*OptimismSequencerUptimeFeedNewRound, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismSequencerUptimeFeedOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*OptimismSequencerUptimeFeedOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OptimismSequencerUptimeFeedOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*OptimismSequencerUptimeFeedOwnershipTransferred, error)

	FilterRemovedAccess(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedRemovedAccessIterator, error)

	WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedRemovedAccess) (event.Subscription, error)

	ParseRemovedAccess(log types.Log) (*OptimismSequencerUptimeFeedRemovedAccess, error)

	FilterRoundUpdated(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedRoundUpdatedIterator, error)

	WatchRoundUpdated(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedRoundUpdated) (event.Subscription, error)

	ParseRoundUpdated(log types.Log) (*OptimismSequencerUptimeFeedRoundUpdated, error)

	FilterUpdateIgnored(opts *bind.FilterOpts) (*OptimismSequencerUptimeFeedUpdateIgnoredIterator, error)

	WatchUpdateIgnored(opts *bind.WatchOpts, sink chan<- *OptimismSequencerUptimeFeedUpdateIgnored) (event.Subscription, error)

	ParseUpdateIgnored(log types.Log) (*OptimismSequencerUptimeFeedUpdateIgnored, error)

	Address() common.Address
}
