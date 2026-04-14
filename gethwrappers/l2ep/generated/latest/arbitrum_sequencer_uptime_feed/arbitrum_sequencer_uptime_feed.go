// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum_sequencer_uptime_feed

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

var ArbitrumSequencerUptimeFeedMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"flagsAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1SenderAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"FLAGS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract IFlags\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FLAG_L2_SEQ_OFFLINE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aliasedL1MessageSender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"description\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableAccessCheck\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAnswer\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoundData\",\"inputs\":[{\"name\":\"_roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"outputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"answeredInRound\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimestamp\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1Sender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestAnswer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRoundData\",\"inputs\":[],\"outputs\":[{\"name\":\"roundId\",\"type\":\"uint80\",\"internalType\":\"uint80\"},{\"name\":\"answer\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"answeredInRound\",\"type\":\"uint80\",\"internalType\":\"uint80\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAccess\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferL1Sender\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AnswerUpdated\",\"inputs\":[{\"name\":\"current\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"roundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"updatedAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckAccessEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L1SenderTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewRound\",\"inputs\":[{\"name\":\"roundId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"startedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"startedAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedAccess\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateIgnored\",\"inputs\":[{\"name\":\"latestStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"latestTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"incomingStatus\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"incomingTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoDataPresent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Uninitialized\",\"inputs\":[]}]",
	Bin: "0x6101006040525f60a081905260c081905260e052600480546001600160981b03191690553480156200002f575f80fd5b5060405162001f5638038062001f5683398101604081905262000052916200023d565b33805f81620000a85760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000da57620000da816200010e565b50506001805460ff60a01b1916600160a01b17905550620000fb81620001b8565b506001600160a01b031660805262000273565b336001600160a01b03821603620001685760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016200009f565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6003546001600160a01b0390811690821681146200021d57600380546001600160a01b0319166001600160a01b0384811691821790925560405190918316907f8e6da65f164d652f378f48652c0e1ca58d7c9cc52ceaa40c1dad055cd7681d18905f90a35b5050565b80516001600160a01b038116811462000238575f80fd5b919050565b5f80604083850312156200024f575f80fd5b6200025a8362000221565b91506200026a6020840162000221565b90509250929050565b608051611cb5620002a15f395f8181610317015281816109b8015281816117a9015261187a0152611cb55ff3fe608060405234801561000f575f80fd5b50600436106101b0575f3560e01c80638205bf6a116100f3578063b5ab58dc11610093578063dc7f01241161006e578063dc7f0124146103da578063ed8378f5146103ff578063f2fde38b14610412578063feaf968c14610425575f80fd5b8063b5ab58dc14610396578063b633620c146103a9578063b7558b7a146103bc575f80fd5b8063926aa9f6116100ce578063926aa9f61461030a5780639728538f146103125780639a6fc8f514610339578063a118f24914610383575f80fd5b80638205bf6a146102d25780638823da6c146102da5780638da5cb5b146102ed575f80fd5b806354fd4d501161015e5780637284e416116101395780637284e4161461027e57806379ba5097146102ba5780638038e4a1146102c25780638129fc1c146102ca575f80fd5b806354fd4d501461024b578063668a0f02146102535780636b14daf81461025b575f80fd5b8063284afc081161018e578063284afc0814610209578063313ce5671461021c57806350d25bcd14610235575f80fd5b80630a756983146101b4578063181f5a77146101be5780631e8d27dd146101dc575b5f80fd5b6101bc61042d565b005b6101c66104ab565b6040516101d391906119be565b60405180910390f35b6101e46104cb565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6101bc610217366004611a4b565b6104fc565b6102235f81565b60405160ff90911681526020016101d3565b61023d610510565b6040519081526020016101d3565b61023d600181565b61023d61062e565b61026e610269366004611a91565b610746565b60405190151581526020016101d3565b6101c66040518060400160405280601f81526020017f4c322053657175656e63657220557074696d652053746174757320466565640081525081565b6101bc61077a565b6101bc610876565b6101bc610909565b61023d610abc565b6101bc6102e8366004611a4b565b610bd5565b5f5473ffffffffffffffffffffffffffffffffffffffff166101e4565b6101e4610c8c565b6101e47f000000000000000000000000000000000000000000000000000000000000000081565b61034c610347366004611b6b565b610ccd565b6040805169ffffffffffffffffffff968716815260208101959095528401929092526060830152909116608082015260a0016101d3565b6101bc610391366004611a4b565b610e27565b61023d6103a4366004611b9b565b610ed9565b61023d6103b7366004611b9b565b610fd4565b60035473ffffffffffffffffffffffffffffffffffffffff166101e4565b60015461026e9074010000000000000000000000000000000000000000900460ff1681565b6101bc61040d366004611bbf565b6110cd565b6101bc610420366004611a4b565b611276565b61034c611287565b6104356113c4565b60015474010000000000000000000000000000000000000000900460ff16156104a957600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638905f90a15b565b6060604051806060016040528060218152602001611c8860219139905090565b6104f660017fa438451d6458044c3c8cd2f6f31c91ac882a6d917fa1d50c2bc3074c4524952d611c2f565b60601c81565b6105046113c4565b61050d81611444565b50565b5f610550335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061074692505050565b6105bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f20616363657373000000000000000000000000000000000000000000000060448201526064015b60405180910390fd5b6040805160608101825260045469ffffffffffffffffffff81168083526a0100000000000000000000820460ff16151560208401526b01000000000000000000000090910467ffffffffffffffff16928201929092529061061b906114de565b6106288160200151611523565b91505090565b5f61066e335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061074692505050565b6106d4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f20616363657373000000000000000000000000000000000000000000000060448201526064016105b2565b6040805160608101825260045469ffffffffffffffffffff81168083526a0100000000000000000000820460ff16151560208401526b01000000000000000000000090910467ffffffffffffffff169282019290925290610734906114de565b5169ffffffffffffffffffff16905090565b5f6107518383611537565b80610771575073ffffffffffffffffffffffffffffffffffffffff831632145b90505b92915050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146107fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064016105b2565b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b61087e6113c4565b60015474010000000000000000000000000000000000000000900460ff166104a957600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517faebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480905f90a1565b6109116113c4565b6040805160608101825260045469ffffffffffffffffffff81168083526a0100000000000000000000820460ff16151560208401526b01000000000000000000000090910467ffffffffffffffff1692820192909252901561099f576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b425f73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663357e47fe610a0860017fa438451d6458044c3c8cd2f6f31c91ac882a6d917fa1d50c2bc3074c4524952d611c2f565b60405160e083901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260609190911c6004820152602401602060405180830381865afa158015610a5d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a819190611c42565b9050610a8f6001828461158b565b6040517f5daa87a0e9463431830481fd4b6e3403442dfb9a12b9c07597e9f61d50b633c8905f90a1505050565b5f610afc335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061074692505050565b610b62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f20616363657373000000000000000000000000000000000000000000000060448201526064016105b2565b6040805160608101825260045469ffffffffffffffffffff81168083526a0100000000000000000000820460ff16151560208401526b01000000000000000000000090910467ffffffffffffffff169282019290925290610bc2906114de565b6040015167ffffffffffffffff16905090565b610bdd6113c4565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff161561050d5773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905590519182527f3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d191015b60405180910390a150565b5f610cc8610caf60035473ffffffffffffffffffffffffffffffffffffffff1690565b7311110000000000000000000000000000000011110190565b905090565b5f805f805f610d11335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061074692505050565b610d77576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f20616363657373000000000000000000000000000000000000000000000060448201526064016105b2565b600454610d8f9069ffffffffffffffffffff166114de565b610da48669ffffffffffffffffffff16611756565b15610e125769ffffffffffffffffffff86165f9081526005602090815260409182902082518084019093525460ff8116151580845261010090910467ffffffffffffffff1691830191909152610df990611523565b9450806020015167ffffffffffffffff16935050610e19565b5f93505f92505b509394919350915081908490565b610e2f6113c4565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526002602052604090205460ff1661050d5773ffffffffffffffffffffffffffffffffffffffff81165f8181526002602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590519182527f87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db49101610c81565b5f610f19335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061074692505050565b610f7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f20616363657373000000000000000000000000000000000000000000000060448201526064016105b2565b600454610f979069ffffffffffffffffffff166114de565b610fa082611756565b15610fcc5769ffffffffffffffffffff82165f908152600560205260409020546107749060ff16611523565b505f5b919050565b5f611014335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061074692505050565b61107a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f20616363657373000000000000000000000000000000000000000000000060448201526064016105b2565b6004546110929069ffffffffffffffffffff166114de565b61109b82611756565b15610fcc575069ffffffffffffffffffff165f90815260056020526040902054610100900467ffffffffffffffff1690565b6040805160608101825260045469ffffffffffffffffffff81168083526a0100000000000000000000820460ff16151560208401526b01000000000000000000000090910467ffffffffffffffff16928201929092529061112d906114de565b611135610c8c565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611199576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8215158160200151151514806111c657508167ffffffffffffffff16816040015167ffffffffffffffff16115b1561123a577fe4a6e16301740042c17431042adb8f60454c18fb5934dd4c456269c0dc388fdf81602001518260400151858560405161122d9493929190931515845267ffffffffffffffff9283166020850152901515604084015216606082015260800190565b60405180910390a1505050565b6001815f0181815161124c9190611c5d565b69ffffffffffffffffffff16905250805161126890848461158b565b6112718361178c565b505050565b61127e6113c4565b61050d816118ca565b5f805f805f6112cb335f368080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061074692505050565b611331576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f4e6f20616363657373000000000000000000000000000000000000000000000060448201526064016105b2565b6040805160608101825260045469ffffffffffffffffffff81168083526a0100000000000000000000820460ff16151560208401526b01000000000000000000000090910467ffffffffffffffff169282019290925290611391906114de565b805f015195506113a48160200151611523565b6040909101519596909567ffffffffffffffff1694508493508692509050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146104a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016105b2565b60035473ffffffffffffffffffffffffffffffffffffffff90811690821681146114da57600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691821790925560405190918316907f8e6da65f164d652f378f48652c0e1ca58d7c9cc52ceaa40c1dad055cd7681d18905f90a35b5050565b8069ffffffffffffffffffff165f0361050d576040517f1c72fad400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8161152f575f610774565b600192915050565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526002602052604081205460ff168061077157505060015474010000000000000000000000000000000000000000900460ff161592915050565b60408051808201825283151580825267ffffffffffffffff8481166020808501828152865160608101885269ffffffffffffffffffff8b8116808352828501978852828a018681525f828152600587528b90208a51815496518a16610100027fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff911515919091167fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000009097169690961795909517909455825160048054995195519098166b010000000000000000000000027fffffffffffffffffffffffffff0000000000000000ffffffffffffffffffffff9515156a0100000000000000000000027fffffffffffffffffffffffffffffffffffffffffff0000000000000000000000909a1691909316179790971792909216919091179093559451908152929390923392917f0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271910160405180910390a38469ffffffffffffffffffff1661171285611523565b60405167ffffffffffffffff861681527f0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f9060200160405180910390a35050505050565b5f8082118015611770575069ffffffffffffffffffff8211155b801561077457505060045469ffffffffffffffffffff16101590565b80156118635773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663d74af2636117f960017fa438451d6458044c3c8cd2f6f31c91ac882a6d917fa1d50c2bc3074c4524952d611c2f565b60405160e083901b7fffffffff0000000000000000000000000000000000000000000000000000000016815260609190911c60048201526024015f604051808303815f87803b15801561184a575f80fd5b505af115801561185c573d5f803e3d5ffd5b5050505050565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663e37a83666117f960017fa438451d6458044c3c8cd2f6f31c91ac882a6d917fa1d50c2bc3074c4524952d611c2f565b3373ffffffffffffffffffffffffffffffffffffffff821603611949576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016105b2565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b5f602080835283518060208501525f5b818110156119ea578581018301518582016040015282016119ce565b505f6040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610fcf575f80fd5b5f60208284031215611a5b575f80fd5b61077182611a28565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f8060408385031215611aa2575f80fd5b611aab83611a28565b9150602083013567ffffffffffffffff80821115611ac7575f80fd5b818501915085601f830112611ada575f80fd5b813581811115611aec57611aec611a64565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611b3257611b32611a64565b81604052828152886020848701011115611b4a575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f60208284031215611b7b575f80fd5b813569ffffffffffffffffffff81168114611b94575f80fd5b9392505050565b5f60208284031215611bab575f80fd5b5035919050565b801515811461050d575f80fd5b5f8060408385031215611bd0575f80fd5b8235611bdb81611bb2565b9150602083013567ffffffffffffffff81168114611bf7575f80fd5b809150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8181038181111561077457610774611c02565b5f60208284031215611c52575f80fd5b8151611b9481611bb2565b69ffffffffffffffffffff818116838216019080821115611c8057611c80611c02565b509291505056fe417262697472756d53657175656e636572557074696d654665656420312e302e30a164736f6c6343000818000a",
}

var ArbitrumSequencerUptimeFeedABI = ArbitrumSequencerUptimeFeedMetaData.ABI

var ArbitrumSequencerUptimeFeedBin = ArbitrumSequencerUptimeFeedMetaData.Bin

func DeployArbitrumSequencerUptimeFeed(auth *bind.TransactOpts, backend bind.ContractBackend, flagsAddress common.Address, l1SenderAddress common.Address) (common.Address, *types.Transaction, *ArbitrumSequencerUptimeFeed, error) {
	parsed, err := ArbitrumSequencerUptimeFeedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ArbitrumSequencerUptimeFeedBin), backend, flagsAddress, l1SenderAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbitrumSequencerUptimeFeed{address: address, abi: *parsed, ArbitrumSequencerUptimeFeedCaller: ArbitrumSequencerUptimeFeedCaller{contract: contract}, ArbitrumSequencerUptimeFeedTransactor: ArbitrumSequencerUptimeFeedTransactor{contract: contract}, ArbitrumSequencerUptimeFeedFilterer: ArbitrumSequencerUptimeFeedFilterer{contract: contract}}, nil
}

type ArbitrumSequencerUptimeFeed struct {
	address common.Address
	abi     abi.ABI
	ArbitrumSequencerUptimeFeedCaller
	ArbitrumSequencerUptimeFeedTransactor
	ArbitrumSequencerUptimeFeedFilterer
}

type ArbitrumSequencerUptimeFeedCaller struct {
	contract *bind.BoundContract
}

type ArbitrumSequencerUptimeFeedTransactor struct {
	contract *bind.BoundContract
}

type ArbitrumSequencerUptimeFeedFilterer struct {
	contract *bind.BoundContract
}

type ArbitrumSequencerUptimeFeedSession struct {
	Contract     *ArbitrumSequencerUptimeFeed
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ArbitrumSequencerUptimeFeedCallerSession struct {
	Contract *ArbitrumSequencerUptimeFeedCaller
	CallOpts bind.CallOpts
}

type ArbitrumSequencerUptimeFeedTransactorSession struct {
	Contract     *ArbitrumSequencerUptimeFeedTransactor
	TransactOpts bind.TransactOpts
}

type ArbitrumSequencerUptimeFeedRaw struct {
	Contract *ArbitrumSequencerUptimeFeed
}

type ArbitrumSequencerUptimeFeedCallerRaw struct {
	Contract *ArbitrumSequencerUptimeFeedCaller
}

type ArbitrumSequencerUptimeFeedTransactorRaw struct {
	Contract *ArbitrumSequencerUptimeFeedTransactor
}

func NewArbitrumSequencerUptimeFeed(address common.Address, backend bind.ContractBackend) (*ArbitrumSequencerUptimeFeed, error) {
	abi, err := abi.JSON(strings.NewReader(ArbitrumSequencerUptimeFeedABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindArbitrumSequencerUptimeFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeed{address: address, abi: abi, ArbitrumSequencerUptimeFeedCaller: ArbitrumSequencerUptimeFeedCaller{contract: contract}, ArbitrumSequencerUptimeFeedTransactor: ArbitrumSequencerUptimeFeedTransactor{contract: contract}, ArbitrumSequencerUptimeFeedFilterer: ArbitrumSequencerUptimeFeedFilterer{contract: contract}}, nil
}

func NewArbitrumSequencerUptimeFeedCaller(address common.Address, caller bind.ContractCaller) (*ArbitrumSequencerUptimeFeedCaller, error) {
	contract, err := bindArbitrumSequencerUptimeFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedCaller{contract: contract}, nil
}

func NewArbitrumSequencerUptimeFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbitrumSequencerUptimeFeedTransactor, error) {
	contract, err := bindArbitrumSequencerUptimeFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedTransactor{contract: contract}, nil
}

func NewArbitrumSequencerUptimeFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbitrumSequencerUptimeFeedFilterer, error) {
	contract, err := bindArbitrumSequencerUptimeFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedFilterer{contract: contract}, nil
}

func bindArbitrumSequencerUptimeFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbitrumSequencerUptimeFeedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumSequencerUptimeFeed.Contract.ArbitrumSequencerUptimeFeedCaller.contract.Call(opts, result, method, params...)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.ArbitrumSequencerUptimeFeedTransactor.contract.Transfer(opts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.ArbitrumSequencerUptimeFeedTransactor.contract.Transact(opts, method, params...)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbitrumSequencerUptimeFeed.Contract.contract.Call(opts, result, method, params...)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.contract.Transfer(opts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.contract.Transact(opts, method, params...)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) FLAGS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "FLAGS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) FLAGS() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.FLAGS(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) FLAGS() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.FLAGS(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) FLAGL2SEQOFFLINE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "FLAG_L2_SEQ_OFFLINE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) FLAGL2SEQOFFLINE() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.FLAGL2SEQOFFLINE(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) FLAGL2SEQOFFLINE() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.FLAGL2SEQOFFLINE(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) AliasedL1MessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "aliasedL1MessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) AliasedL1MessageSender() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.AliasedL1MessageSender(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) AliasedL1MessageSender() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.AliasedL1MessageSender(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) CheckEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "checkEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) CheckEnabled() (bool, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.CheckEnabled(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) CheckEnabled() (bool, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.CheckEnabled(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) Decimals() (uint8, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Decimals(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) Decimals() (uint8, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Decimals(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) Description() (string, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Description(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) Description() (string, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Description(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "getAnswer", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.GetAnswer(&_ArbitrumSequencerUptimeFeed.CallOpts, roundId)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.GetAnswer(&_ArbitrumSequencerUptimeFeed.CallOpts, roundId)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (GetRoundData,

	error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "getRoundData", _roundId)

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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) GetRoundData(_roundId *big.Int) (GetRoundData,

	error) {
	return _ArbitrumSequencerUptimeFeed.Contract.GetRoundData(&_ArbitrumSequencerUptimeFeed.CallOpts, _roundId)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) GetRoundData(_roundId *big.Int) (GetRoundData,

	error) {
	return _ArbitrumSequencerUptimeFeed.Contract.GetRoundData(&_ArbitrumSequencerUptimeFeed.CallOpts, _roundId)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "getTimestamp", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.GetTimestamp(&_ArbitrumSequencerUptimeFeed.CallOpts, roundId)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.GetTimestamp(&_ArbitrumSequencerUptimeFeed.CallOpts, roundId)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) HasAccess(opts *bind.CallOpts, _user common.Address, _calldata []byte) (bool, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "hasAccess", _user, _calldata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.HasAccess(&_ArbitrumSequencerUptimeFeed.CallOpts, _user, _calldata)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) HasAccess(_user common.Address, _calldata []byte) (bool, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.HasAccess(&_ArbitrumSequencerUptimeFeed.CallOpts, _user, _calldata)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) L1Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "l1Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) L1Sender() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.L1Sender(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) L1Sender() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.L1Sender(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) LatestAnswer() (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.LatestAnswer(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) LatestAnswer() (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.LatestAnswer(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) LatestRound() (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.LatestRound(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) LatestRound() (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.LatestRound(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) LatestRoundData(opts *bind.CallOpts) (LatestRoundData,

	error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "latestRoundData")

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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) LatestRoundData() (LatestRoundData,

	error) {
	return _ArbitrumSequencerUptimeFeed.Contract.LatestRoundData(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) LatestRoundData() (LatestRoundData,

	error) {
	return _ArbitrumSequencerUptimeFeed.Contract.LatestRoundData(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) LatestTimestamp() (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.LatestTimestamp(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) LatestTimestamp() (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.LatestTimestamp(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) Owner() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Owner(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) Owner() (common.Address, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Owner(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) TypeAndVersion() (string, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.TypeAndVersion(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) TypeAndVersion() (string, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.TypeAndVersion(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbitrumSequencerUptimeFeed.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) Version() (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Version(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedCallerSession) Version() (*big.Int, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Version(&_ArbitrumSequencerUptimeFeed.CallOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.contract.Transact(opts, "acceptOwnership")
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.AcceptOwnership(&_ArbitrumSequencerUptimeFeed.TransactOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.AcceptOwnership(&_ArbitrumSequencerUptimeFeed.TransactOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactor) AddAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.contract.Transact(opts, "addAccess", _user)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.AddAccess(&_ArbitrumSequencerUptimeFeed.TransactOpts, _user)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorSession) AddAccess(_user common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.AddAccess(&_ArbitrumSequencerUptimeFeed.TransactOpts, _user)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactor) DisableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.contract.Transact(opts, "disableAccessCheck")
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) DisableAccessCheck() (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.DisableAccessCheck(&_ArbitrumSequencerUptimeFeed.TransactOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorSession) DisableAccessCheck() (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.DisableAccessCheck(&_ArbitrumSequencerUptimeFeed.TransactOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactor) EnableAccessCheck(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.contract.Transact(opts, "enableAccessCheck")
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) EnableAccessCheck() (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.EnableAccessCheck(&_ArbitrumSequencerUptimeFeed.TransactOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorSession) EnableAccessCheck() (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.EnableAccessCheck(&_ArbitrumSequencerUptimeFeed.TransactOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.contract.Transact(opts, "initialize")
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) Initialize() (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Initialize(&_ArbitrumSequencerUptimeFeed.TransactOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorSession) Initialize() (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.Initialize(&_ArbitrumSequencerUptimeFeed.TransactOpts)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactor) RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.contract.Transact(opts, "removeAccess", _user)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.RemoveAccess(&_ArbitrumSequencerUptimeFeed.TransactOpts, _user)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorSession) RemoveAccess(_user common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.RemoveAccess(&_ArbitrumSequencerUptimeFeed.TransactOpts, _user)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactor) TransferL1Sender(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.contract.Transact(opts, "transferL1Sender", to)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) TransferL1Sender(to common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.TransferL1Sender(&_ArbitrumSequencerUptimeFeed.TransactOpts, to)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorSession) TransferL1Sender(to common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.TransferL1Sender(&_ArbitrumSequencerUptimeFeed.TransactOpts, to)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.contract.Transact(opts, "transferOwnership", to)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.TransferOwnership(&_ArbitrumSequencerUptimeFeed.TransactOpts, to)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.TransferOwnership(&_ArbitrumSequencerUptimeFeed.TransactOpts, to)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactor) UpdateStatus(opts *bind.TransactOpts, status bool, timestamp uint64) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.contract.Transact(opts, "updateStatus", status, timestamp)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedSession) UpdateStatus(status bool, timestamp uint64) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.UpdateStatus(&_ArbitrumSequencerUptimeFeed.TransactOpts, status, timestamp)
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedTransactorSession) UpdateStatus(status bool, timestamp uint64) (*types.Transaction, error) {
	return _ArbitrumSequencerUptimeFeed.Contract.UpdateStatus(&_ArbitrumSequencerUptimeFeed.TransactOpts, status, timestamp)
}

type ArbitrumSequencerUptimeFeedAddedAccessIterator struct {
	Event *ArbitrumSequencerUptimeFeedAddedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedAddedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedAddedAccess)
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
		it.Event = new(ArbitrumSequencerUptimeFeedAddedAccess)
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

func (it *ArbitrumSequencerUptimeFeedAddedAccessIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedAddedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedAddedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterAddedAccess(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedAddedAccessIterator, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedAddedAccessIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "AddedAccess", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedAddedAccess) (event.Subscription, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "AddedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedAddedAccess)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "AddedAccess", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseAddedAccess(log types.Log) (*ArbitrumSequencerUptimeFeedAddedAccess, error) {
	event := new(ArbitrumSequencerUptimeFeedAddedAccess)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "AddedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedAnswerUpdatedIterator struct {
	Event *ArbitrumSequencerUptimeFeedAnswerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedAnswerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedAnswerUpdated)
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
		it.Event = new(ArbitrumSequencerUptimeFeedAnswerUpdated)
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

func (it *ArbitrumSequencerUptimeFeedAnswerUpdatedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ArbitrumSequencerUptimeFeedAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedAnswerUpdatedIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedAnswerUpdated)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseAnswerUpdated(log types.Log) (*ArbitrumSequencerUptimeFeedAnswerUpdated, error) {
	event := new(ArbitrumSequencerUptimeFeedAnswerUpdated)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedCheckAccessDisabledIterator struct {
	Event *ArbitrumSequencerUptimeFeedCheckAccessDisabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedCheckAccessDisabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedCheckAccessDisabled)
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
		it.Event = new(ArbitrumSequencerUptimeFeedCheckAccessDisabled)
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

func (it *ArbitrumSequencerUptimeFeedCheckAccessDisabledIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedCheckAccessDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedCheckAccessDisabled struct {
	Raw types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterCheckAccessDisabled(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedCheckAccessDisabledIterator, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedCheckAccessDisabledIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "CheckAccessDisabled", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedCheckAccessDisabled) (event.Subscription, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "CheckAccessDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedCheckAccessDisabled)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseCheckAccessDisabled(log types.Log) (*ArbitrumSequencerUptimeFeedCheckAccessDisabled, error) {
	event := new(ArbitrumSequencerUptimeFeedCheckAccessDisabled)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "CheckAccessDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedCheckAccessEnabledIterator struct {
	Event *ArbitrumSequencerUptimeFeedCheckAccessEnabled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedCheckAccessEnabledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedCheckAccessEnabled)
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
		it.Event = new(ArbitrumSequencerUptimeFeedCheckAccessEnabled)
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

func (it *ArbitrumSequencerUptimeFeedCheckAccessEnabledIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedCheckAccessEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedCheckAccessEnabled struct {
	Raw types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterCheckAccessEnabled(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedCheckAccessEnabledIterator, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedCheckAccessEnabledIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "CheckAccessEnabled", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedCheckAccessEnabled) (event.Subscription, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "CheckAccessEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedCheckAccessEnabled)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseCheckAccessEnabled(log types.Log) (*ArbitrumSequencerUptimeFeedCheckAccessEnabled, error) {
	event := new(ArbitrumSequencerUptimeFeedCheckAccessEnabled)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "CheckAccessEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedInitializedIterator struct {
	Event *ArbitrumSequencerUptimeFeedInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedInitialized)
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
		it.Event = new(ArbitrumSequencerUptimeFeedInitialized)
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

func (it *ArbitrumSequencerUptimeFeedInitializedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedInitialized struct {
	Raw types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterInitialized(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedInitializedIterator, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedInitializedIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedInitialized) (event.Subscription, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedInitialized)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "Initialized", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseInitialized(log types.Log) (*ArbitrumSequencerUptimeFeedInitialized, error) {
	event := new(ArbitrumSequencerUptimeFeedInitialized)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedL1SenderTransferredIterator struct {
	Event *ArbitrumSequencerUptimeFeedL1SenderTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedL1SenderTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedL1SenderTransferred)
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
		it.Event = new(ArbitrumSequencerUptimeFeedL1SenderTransferred)
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

func (it *ArbitrumSequencerUptimeFeedL1SenderTransferredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedL1SenderTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedL1SenderTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterL1SenderTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumSequencerUptimeFeedL1SenderTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "L1SenderTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedL1SenderTransferredIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "L1SenderTransferred", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchL1SenderTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedL1SenderTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "L1SenderTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedL1SenderTransferred)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "L1SenderTransferred", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseL1SenderTransferred(log types.Log) (*ArbitrumSequencerUptimeFeedL1SenderTransferred, error) {
	event := new(ArbitrumSequencerUptimeFeedL1SenderTransferred)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "L1SenderTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedNewRoundIterator struct {
	Event *ArbitrumSequencerUptimeFeedNewRound

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedNewRoundIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedNewRound)
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
		it.Event = new(ArbitrumSequencerUptimeFeedNewRound)
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

func (it *ArbitrumSequencerUptimeFeedNewRoundIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*ArbitrumSequencerUptimeFeedNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedNewRoundIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedNewRound)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "NewRound", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseNewRound(log types.Log) (*ArbitrumSequencerUptimeFeedNewRound, error) {
	event := new(ArbitrumSequencerUptimeFeedNewRound)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedOwnershipTransferRequestedIterator struct {
	Event *ArbitrumSequencerUptimeFeedOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedOwnershipTransferRequested)
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
		it.Event = new(ArbitrumSequencerUptimeFeedOwnershipTransferRequested)
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

func (it *ArbitrumSequencerUptimeFeedOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumSequencerUptimeFeedOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedOwnershipTransferRequestedIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedOwnershipTransferRequested)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseOwnershipTransferRequested(log types.Log) (*ArbitrumSequencerUptimeFeedOwnershipTransferRequested, error) {
	event := new(ArbitrumSequencerUptimeFeedOwnershipTransferRequested)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedOwnershipTransferredIterator struct {
	Event *ArbitrumSequencerUptimeFeedOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedOwnershipTransferred)
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
		it.Event = new(ArbitrumSequencerUptimeFeedOwnershipTransferred)
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

func (it *ArbitrumSequencerUptimeFeedOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumSequencerUptimeFeedOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedOwnershipTransferredIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedOwnershipTransferred)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseOwnershipTransferred(log types.Log) (*ArbitrumSequencerUptimeFeedOwnershipTransferred, error) {
	event := new(ArbitrumSequencerUptimeFeedOwnershipTransferred)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedRemovedAccessIterator struct {
	Event *ArbitrumSequencerUptimeFeedRemovedAccess

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedRemovedAccessIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedRemovedAccess)
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
		it.Event = new(ArbitrumSequencerUptimeFeedRemovedAccess)
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

func (it *ArbitrumSequencerUptimeFeedRemovedAccessIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedRemovedAccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedRemovedAccess struct {
	User common.Address
	Raw  types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterRemovedAccess(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedRemovedAccessIterator, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedRemovedAccessIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "RemovedAccess", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedRemovedAccess) (event.Subscription, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "RemovedAccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedRemovedAccess)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseRemovedAccess(log types.Log) (*ArbitrumSequencerUptimeFeedRemovedAccess, error) {
	event := new(ArbitrumSequencerUptimeFeedRemovedAccess)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "RemovedAccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ArbitrumSequencerUptimeFeedUpdateIgnoredIterator struct {
	Event *ArbitrumSequencerUptimeFeedUpdateIgnored

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ArbitrumSequencerUptimeFeedUpdateIgnoredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbitrumSequencerUptimeFeedUpdateIgnored)
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
		it.Event = new(ArbitrumSequencerUptimeFeedUpdateIgnored)
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

func (it *ArbitrumSequencerUptimeFeedUpdateIgnoredIterator) Error() error {
	return it.fail
}

func (it *ArbitrumSequencerUptimeFeedUpdateIgnoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ArbitrumSequencerUptimeFeedUpdateIgnored struct {
	LatestStatus      bool
	LatestTimestamp   uint64
	IncomingStatus    bool
	IncomingTimestamp uint64
	Raw               types.Log
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) FilterUpdateIgnored(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedUpdateIgnoredIterator, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.FilterLogs(opts, "UpdateIgnored")
	if err != nil {
		return nil, err
	}
	return &ArbitrumSequencerUptimeFeedUpdateIgnoredIterator{contract: _ArbitrumSequencerUptimeFeed.contract, event: "UpdateIgnored", logs: logs, sub: sub}, nil
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) WatchUpdateIgnored(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedUpdateIgnored) (event.Subscription, error) {

	logs, sub, err := _ArbitrumSequencerUptimeFeed.contract.WatchLogs(opts, "UpdateIgnored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ArbitrumSequencerUptimeFeedUpdateIgnored)
				if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "UpdateIgnored", log); err != nil {
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

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeedFilterer) ParseUpdateIgnored(log types.Log) (*ArbitrumSequencerUptimeFeedUpdateIgnored, error) {
	event := new(ArbitrumSequencerUptimeFeedUpdateIgnored)
	if err := _ArbitrumSequencerUptimeFeed.contract.UnpackLog(event, "UpdateIgnored", log); err != nil {
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

func (ArbitrumSequencerUptimeFeedAddedAccess) Topic() common.Hash {
	return common.HexToHash("0x87286ad1f399c8e82bf0c4ef4fcdc570ea2e1e92176e5c848b6413545b885db4")
}

func (ArbitrumSequencerUptimeFeedAnswerUpdated) Topic() common.Hash {
	return common.HexToHash("0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f")
}

func (ArbitrumSequencerUptimeFeedCheckAccessDisabled) Topic() common.Hash {
	return common.HexToHash("0x3be8a977a014527b50ae38adda80b56911c267328965c98ddc385d248f539638")
}

func (ArbitrumSequencerUptimeFeedCheckAccessEnabled) Topic() common.Hash {
	return common.HexToHash("0xaebf329500988c6488a0074e5a0a9ff304561fc5c6fc877aeb1d59c8282c3480")
}

func (ArbitrumSequencerUptimeFeedInitialized) Topic() common.Hash {
	return common.HexToHash("0x5daa87a0e9463431830481fd4b6e3403442dfb9a12b9c07597e9f61d50b633c8")
}

func (ArbitrumSequencerUptimeFeedL1SenderTransferred) Topic() common.Hash {
	return common.HexToHash("0x8e6da65f164d652f378f48652c0e1ca58d7c9cc52ceaa40c1dad055cd7681d18")
}

func (ArbitrumSequencerUptimeFeedNewRound) Topic() common.Hash {
	return common.HexToHash("0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271")
}

func (ArbitrumSequencerUptimeFeedOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (ArbitrumSequencerUptimeFeedOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ArbitrumSequencerUptimeFeedRemovedAccess) Topic() common.Hash {
	return common.HexToHash("0x3d68a6fce901d20453d1a7aa06bf3950302a735948037deb182a8db66df2a0d1")
}

func (ArbitrumSequencerUptimeFeedUpdateIgnored) Topic() common.Hash {
	return common.HexToHash("0xe4a6e16301740042c17431042adb8f60454c18fb5934dd4c456269c0dc388fdf")
}

func (_ArbitrumSequencerUptimeFeed *ArbitrumSequencerUptimeFeed) Address() common.Address {
	return _ArbitrumSequencerUptimeFeed.address
}

type ArbitrumSequencerUptimeFeedInterface interface {
	FLAGS(opts *bind.CallOpts) (common.Address, error)

	FLAGL2SEQOFFLINE(opts *bind.CallOpts) (common.Address, error)

	AliasedL1MessageSender(opts *bind.CallOpts) (common.Address, error)

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

	Initialize(opts *bind.TransactOpts) (*types.Transaction, error)

	RemoveAccess(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error)

	TransferL1Sender(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	UpdateStatus(opts *bind.TransactOpts, status bool, timestamp uint64) (*types.Transaction, error)

	FilterAddedAccess(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedAddedAccessIterator, error)

	WatchAddedAccess(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedAddedAccess) (event.Subscription, error)

	ParseAddedAccess(log types.Log) (*ArbitrumSequencerUptimeFeedAddedAccess, error)

	FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*ArbitrumSequencerUptimeFeedAnswerUpdatedIterator, error)

	WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error)

	ParseAnswerUpdated(log types.Log) (*ArbitrumSequencerUptimeFeedAnswerUpdated, error)

	FilterCheckAccessDisabled(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedCheckAccessDisabledIterator, error)

	WatchCheckAccessDisabled(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedCheckAccessDisabled) (event.Subscription, error)

	ParseCheckAccessDisabled(log types.Log) (*ArbitrumSequencerUptimeFeedCheckAccessDisabled, error)

	FilterCheckAccessEnabled(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedCheckAccessEnabledIterator, error)

	WatchCheckAccessEnabled(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedCheckAccessEnabled) (event.Subscription, error)

	ParseCheckAccessEnabled(log types.Log) (*ArbitrumSequencerUptimeFeedCheckAccessEnabled, error)

	FilterInitialized(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*ArbitrumSequencerUptimeFeedInitialized, error)

	FilterL1SenderTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumSequencerUptimeFeedL1SenderTransferredIterator, error)

	WatchL1SenderTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedL1SenderTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseL1SenderTransferred(log types.Log) (*ArbitrumSequencerUptimeFeedL1SenderTransferred, error)

	FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*ArbitrumSequencerUptimeFeedNewRoundIterator, error)

	WatchNewRound(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error)

	ParseNewRound(log types.Log) (*ArbitrumSequencerUptimeFeedNewRound, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumSequencerUptimeFeedOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*ArbitrumSequencerUptimeFeedOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ArbitrumSequencerUptimeFeedOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ArbitrumSequencerUptimeFeedOwnershipTransferred, error)

	FilterRemovedAccess(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedRemovedAccessIterator, error)

	WatchRemovedAccess(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedRemovedAccess) (event.Subscription, error)

	ParseRemovedAccess(log types.Log) (*ArbitrumSequencerUptimeFeedRemovedAccess, error)

	FilterUpdateIgnored(opts *bind.FilterOpts) (*ArbitrumSequencerUptimeFeedUpdateIgnoredIterator, error)

	WatchUpdateIgnored(opts *bind.WatchOpts, sink chan<- *ArbitrumSequencerUptimeFeedUpdateIgnored) (event.Subscription, error)

	ParseUpdateIgnored(log types.Log) (*ArbitrumSequencerUptimeFeedUpdateIgnored, error)

	Address() common.Address
}
