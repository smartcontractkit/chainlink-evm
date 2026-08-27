// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_balance_monitor

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

var EthBalanceMonitorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"keeperRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minWaitPeriodSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkUpkeep\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"upkeepNeeded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"performData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccountInfo\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minBalanceWei\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"topUpAmountWei\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"lastTopUpTimestamp\",\"type\":\"uint56\",\"internalType\":\"uint56\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeeperRegistryAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"keeperRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinWaitPeriodSeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnderfundedAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWatchList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"performUpkeep\",\"inputs\":[{\"name\":\"performData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeeperRegistryAddress\",\"inputs\":[{\"name\":\"keeperRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWaitPeriodSeconds\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWatchList\",\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"minBalancesWei\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"topUpAmountsWei\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"topUp\",\"inputs\":[{\"name\":\"needsFunding\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payee\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FundsAdded\",\"inputs\":[{\"name\":\"amountAdded\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsWithdrawn\",\"inputs\":[{\"name\":\"amountWithdrawn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"payee\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperRegistryAddressUpdated\",\"inputs\":[{\"name\":\"oldAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinWaitPeriodUpdated\",\"inputs\":[{\"name\":\"oldMinWaitPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinWaitPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TopUpFailed\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TopUpSucceeded\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DuplicateAddress\",\"inputs\":[{\"name\":\"duplicate\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWatchList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyKeeperRegistry\",\"inputs\":[]}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001f5338038062001f538339810160408190526200003391620002c1565b33805f81620000895760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b5f80546001600160a01b0319166001600160a01b0384811691909117909155811615620000bb57620000bb81620000e9565b50506001805460ff60a01b1916905550620000d68262000193565b620000e18162000219565b5050620002fa565b336001600160a01b03821603620001435760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000080565b600180546001600160a01b0319166001600160a01b038381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6200019d62000264565b6001600160a01b038116620001b0575f80fd5b600254604080516001600160a01b03928316815291831660208301527fb732223055abcde751d7a24272ffc8a3aa571cb72b443969a4199b7ecd59f8b9910160405180910390a1600280546001600160a01b0319166001600160a01b0392909216919091179055565b6200022362000264565b60035460408051918252602082018390527f04330086c73b1fe1e13cd47a61c692e7c4399b5de08ed94b7ab824684af09323910160405180910390a1600355565b5f546001600160a01b03163314620002bf5760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640162000080565b565b5f8060408385031215620002d3575f80fd5b82516001600160a01b0381168114620002ea575f80fd5b6020939093015192949293505050565b611c4b80620003085f395ff3fe608060405260043610610125575f3560e01c8063728584b7116100a15780638456cb59116100715780639455511411610057578063945551141461045a578063b1d52fa014610479578063f2fde38b14610498575f80fd5b80638456cb591461041d5780638da5cb5b14610431575f80fd5b8063728584b7146102ae57806379ba5097146102c25780637b510fe8146102d6578063810623e3146103d2575f80fd5b80633f85861f116100f65780634585e33b116100dc5780634585e33b146102285780635c975abb146102475780636e04ff0d14610281575f80fd5b80633f85861f146101ec57806341d2052e1461020b575f80fd5b8062f714ce1461016e5780630b67ddce1461018f5780633e4ca677146101b95780633f4ba83a146101d8575f80fd5b3661016a5760408051348152476020820152338183015290517fc6f3fb0fec49e4877342d4625d77a632541f55b7aae0f9d0b34c69b3478706dc9181900360600190a1005b5f80fd5b348015610179575f80fd5b5061018d6101883660046117c7565b6104b7565b005b34801561019a575f80fd5b506101a3610571565b6040516101b091906117f5565b60405180910390f35b3480156101c4575f80fd5b5061018d6101d336600461188b565b610836565b3480156101e3575f80fd5b5061018d610b65565b3480156101f7575f80fd5b5061018d610206366004611969565b610b77565b348015610216575f80fd5b506003546040519081526020016101b0565b348015610233575f80fd5b5061018d610242366004611980565b610bc0565b348015610252575f80fd5b5060015474010000000000000000000000000000000000000000900460ff1660405190151581526020016101b0565b34801561028c575f80fd5b506102a061029b366004611980565b610c31565b6040516101b09291906119ec565b3480156102b9575f80fd5b506101a3610c78565b3480156102cd575f80fd5b5061018d610ce5565b3480156102e1575f80fd5b506103926102f0366004611a60565b73ffffffffffffffffffffffffffffffffffffffff165f908152600560209081526040918290208251608081018452905460ff8116151580835261010082046bffffffffffffffffffffffff9081169484018590526d010000000000000000000000000083041694830185905279010000000000000000000000000000000000000000000000000090910466ffffffffffffff16606090920182905293919291565b6040516101b0949392919093151584526bffffffffffffffffffffffff92831660208501529116604083015266ffffffffffffff16606082015260800190565b3480156103dd575f80fd5b5060025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b0565b348015610428575f80fd5b5061018d610de6565b34801561043c575f80fd5b505f5473ffffffffffffffffffffffffffffffffffffffff166103f8565b348015610465575f80fd5b5061018d610474366004611a60565b610df6565b348015610484575f80fd5b5061018d610493366004611ac3565b610eb8565b3480156104a3575f80fd5b5061018d6104b2366004611a60565b611392565b6104bf6113a3565b73ffffffffffffffffffffffffffffffffffffffff81166104de575f80fd5b6040805183815273ffffffffffffffffffffffffffffffffffffffff831660208201527f6141b54b56b8a52a8c6f5cd2a857f6117b18ffbf4d46bd3106f300a839cbf5ea910160405180910390a160405173ffffffffffffffffffffffffffffffffffffffff82169083156108fc029084905f818181858888f1935050505015801561056c573d5f803e3d5ffd5b505050565b60605f60048054806020026020016040519081016040528092919081815260200182805480156105d557602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116105aa575b505050505090505f815167ffffffffffffffff8111156105f7576105f761184e565b604051908082528060200260200182016040528015610620578160200160208202803683370190505b50600354604080516080810182525f80825260208201819052918101829052606081018290529293509147905f5b865181101561081e5760055f88838151811061066c5761066c611b56565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528181019290925260409081015f208151608081018352905460ff81161515825261010081046bffffffffffffffffffffffff908116948301949094526d010000000000000000000000000081049093169181019190915279010000000000000000000000000000000000000000000000000090910466ffffffffffffff16606082018190529092504290610725908690611bb0565b11158015610745575081604001516bffffffffffffffffffffffff168310155b8015610792575081602001516bffffffffffffffffffffffff1687828151811061077157610771611b56565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1631105b15610816578681815181106107a9576107a9611b56565b60200260200101518686815181106107c3576107c3611b56565b73ffffffffffffffffffffffffffffffffffffffff90921660209283029190910190910152846107f281611bc9565b95505081604001516bffffffffffffffffffffffff16836108139190611c00565b92505b60010161064e565b508551841461082b578385525b509295945050505050565b61083e611423565b600354604080516080810182525f8082526020820181905291810182905260608101829052905b8351811015610b5e5760055f85838151811061088357610883611b56565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528181019290925260409081015f208151608081018352905460ff81161580158084526bffffffffffffffffffffffff61010084048116968501969096526d010000000000000000000000000083049095169383019390935266ffffffffffffff790100000000000000000000000000000000000000000000000000909104166060820152935061094f57504283836060015166ffffffffffffff1661094c9190611bb0565b11155b801561099c575081602001516bffffffffffffffffffffffff1684828151811061097b5761097b611b56565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1631105b15610b46575f8482815181106109b4576109b4611b56565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166108fc84604001516bffffffffffffffffffffffff1690811502906040515f60405180830381858888f1935050505090508015610ae7574260055f878581518110610a1f57610a1f611b56565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0160196101000a81548166ffffffffffffff021916908366ffffffffffffff160217905550848281518110610a9857610a98611b56565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f9eec55c371a49ce19e0a5792787c79b32dcf7d3490aa737436b49c0978ce9ce960405160405180910390a2610b44565b848281518110610af957610af9611b56565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fa9ff7a9b96721b0e16adb7de9db0764fbfd6a4516d4d165f9564e8c3755eb10560405160405180910390a25b505b61d6d85a1015610b565750505050565b600101610865565b5050505b50565b610b6d6113a3565b610b756114a8565b565b610b7f6113a3565b60035460408051918252602082018390527f04330086c73b1fe1e13cd47a61c692e7c4399b5de08ed94b7ab824684af09323910160405180910390a1600355565b60025473ffffffffffffffffffffffffffffffffffffffff163314610c11576040517fd3a6803400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c19611423565b5f610c268284018461188b565b905061056c81610836565b5f6060610c3c611423565b5f610c45610571565b90505f815111925080604051602001610c5e91906117f5565b6040516020818303038152906040529150505b9250929050565b60606004805480602002602001604051908101604052809291908181526020018280548015610cdb57602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610cb0575b5050505050905090565b60015473ffffffffffffffffffffffffffffffffffffffff163314610d6b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b5f8054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610dee6113a3565b610b75611525565b610dfe6113a3565b73ffffffffffffffffffffffffffffffffffffffff8116610e1d575f80fd5b6002546040805173ffffffffffffffffffffffffffffffffffffffff928316815291831660208301527fb732223055abcde751d7a24272ffc8a3aa571cb72b443969a4199b7ecd59f8b9910160405180910390a1600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610ec06113a3565b8483141580610ecf5750848114155b15610f06576040517f3869bbe600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6004805480602002602001604051908101604052809291908181526020018280548015610f6857602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610f3d575b505050505090505f5b8151811015610ff4575f60055f848481518110610f9057610f90611b56565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528101919091526040015f2080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055600101610f71565b505f5b8681101561137b5760055f89898481811061101457611014611b56565b90506020020160208101906110299190611a60565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040015f205460ff16156110c95787878281811061106957611069611b56565b905060200201602081019061107e9190611a60565b6040517f9f2277f300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610d62565b5f8888838181106110dc576110dc611b56565b90506020020160208101906110f19190611a60565b73ffffffffffffffffffffffffffffffffffffffff160361113e576040517f3869bbe600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83838281811061115057611150611b56565b90506020020160208101906111659190611c13565b6bffffffffffffffffffffffff165f036111ab576040517f3869bbe600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180608001604052806001151581526020018787848181106111d1576111d1611b56565b90506020020160208101906111e69190611c13565b6bffffffffffffffffffffffff16815260200185858481811061120b5761120b611b56565b90506020020160208101906112209190611c13565b6bffffffffffffffffffffffff1681526020015f66ffffffffffffff1681525060055f8a8a8581811061125557611255611b56565b905060200201602081019061126a9190611a60565b73ffffffffffffffffffffffffffffffffffffffff16815260208082019290925260409081015f2083518154938501519285015160609095015166ffffffffffffff167901000000000000000000000000000000000000000000000000000278ffffffffffffffffffffffffffffffffffffffffffffffffff6bffffffffffffffffffffffff9687166d010000000000000000000000000002166cffffffffffffffffffffffffff96909416610100027fffffffffffffffffffffffffffffffffffffff000000000000000000000000ff921515929092167fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090951694909417179390931617179055600101610ff7565b506113886004888861170c565b5050505050505050565b61139a6113a3565b610b6281611594565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610b75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610d62565b60015474010000000000000000000000000000000000000000900460ff1615610b75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610d62565b6114b0611688565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b61152d611423565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586114fb3390565b3373ffffffffffffffffffffffffffffffffffffffff821603611613576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610d62565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092555f8054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60015474010000000000000000000000000000000000000000900460ff16610b75576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610d62565b828054828255905f5260205f20908101928215611782579160200282015b828111156117825781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84351617825560209092019160019091019061172a565b5061178e929150611792565b5090565b5b8082111561178e575f8155600101611793565b73ffffffffffffffffffffffffffffffffffffffff81168114610b62575f80fd5b5f80604083850312156117d8575f80fd5b8235915060208301356117ea816117a6565b809150509250929050565b602080825282518282018190525f9190848201906040850190845b8181101561184257835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611810565b50909695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b8035611886816117a6565b919050565b5f602080838503121561189c575f80fd5b823567ffffffffffffffff808211156118b3575f80fd5b818501915085601f8301126118c6575f80fd5b8135818111156118d8576118d861184e565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f8301168101818110858211171561191b5761191b61184e565b604052918252848201925083810185019188831115611938575f80fd5b938501935b8285101561195d5761194e8561187b565b8452938501939285019261193d565b98975050505050505050565b5f60208284031215611979575f80fd5b5035919050565b5f8060208385031215611991575f80fd5b823567ffffffffffffffff808211156119a8575f80fd5b818501915085601f8301126119bb575f80fd5b8135818111156119c9575f80fd5b8660208285010111156119da575f80fd5b60209290920196919550909350505050565b82151581525f60206040602084015283518060408501525f5b81811015611a2157858101830151858201606001528201611a05565b505f6060828601015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509392505050565b5f60208284031215611a70575f80fd5b8135611a7b816117a6565b9392505050565b5f8083601f840112611a92575f80fd5b50813567ffffffffffffffff811115611aa9575f80fd5b6020830191508360208260051b8501011115610c71575f80fd5b5f805f805f8060608789031215611ad8575f80fd5b863567ffffffffffffffff80821115611aef575f80fd5b611afb8a838b01611a82565b90985096506020890135915080821115611b13575f80fd5b611b1f8a838b01611a82565b90965094506040890135915080821115611b37575f80fd5b50611b4489828a01611a82565b979a9699509497509295939492505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820180821115611bc357611bc3611b83565b92915050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611bf957611bf9611b83565b5060010190565b81810381811115611bc357611bc3611b83565b5f60208284031215611c23575f80fd5b81356bffffffffffffffffffffffff81168114611a7b575f80fdfea164736f6c6343000818000a",
}

var EthBalanceMonitorABI = EthBalanceMonitorMetaData.ABI

var EthBalanceMonitorBin = EthBalanceMonitorMetaData.Bin

func DeployEthBalanceMonitor(auth *bind.TransactOpts, backend bind.ContractBackend, keeperRegistryAddress common.Address, minWaitPeriodSeconds *big.Int) (common.Address, *types.Transaction, *EthBalanceMonitor, error) {
	parsed, err := EthBalanceMonitorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthBalanceMonitorBin), backend, keeperRegistryAddress, minWaitPeriodSeconds)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthBalanceMonitor{address: address, abi: *parsed, EthBalanceMonitorCaller: EthBalanceMonitorCaller{contract: contract}, EthBalanceMonitorTransactor: EthBalanceMonitorTransactor{contract: contract}, EthBalanceMonitorFilterer: EthBalanceMonitorFilterer{contract: contract}}, nil
}

type EthBalanceMonitor struct {
	address common.Address
	abi     abi.ABI
	EthBalanceMonitorCaller
	EthBalanceMonitorTransactor
	EthBalanceMonitorFilterer
}

type EthBalanceMonitorCaller struct {
	contract *bind.BoundContract
}

type EthBalanceMonitorTransactor struct {
	contract *bind.BoundContract
}

type EthBalanceMonitorFilterer struct {
	contract *bind.BoundContract
}

type EthBalanceMonitorSession struct {
	Contract     *EthBalanceMonitor
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type EthBalanceMonitorCallerSession struct {
	Contract *EthBalanceMonitorCaller
	CallOpts bind.CallOpts
}

type EthBalanceMonitorTransactorSession struct {
	Contract     *EthBalanceMonitorTransactor
	TransactOpts bind.TransactOpts
}

type EthBalanceMonitorRaw struct {
	Contract *EthBalanceMonitor
}

type EthBalanceMonitorCallerRaw struct {
	Contract *EthBalanceMonitorCaller
}

type EthBalanceMonitorTransactorRaw struct {
	Contract *EthBalanceMonitorTransactor
}

func NewEthBalanceMonitor(address common.Address, backend bind.ContractBackend) (*EthBalanceMonitor, error) {
	abi, err := abi.JSON(strings.NewReader(EthBalanceMonitorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindEthBalanceMonitor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitor{address: address, abi: abi, EthBalanceMonitorCaller: EthBalanceMonitorCaller{contract: contract}, EthBalanceMonitorTransactor: EthBalanceMonitorTransactor{contract: contract}, EthBalanceMonitorFilterer: EthBalanceMonitorFilterer{contract: contract}}, nil
}

func NewEthBalanceMonitorCaller(address common.Address, caller bind.ContractCaller) (*EthBalanceMonitorCaller, error) {
	contract, err := bindEthBalanceMonitor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorCaller{contract: contract}, nil
}

func NewEthBalanceMonitorTransactor(address common.Address, transactor bind.ContractTransactor) (*EthBalanceMonitorTransactor, error) {
	contract, err := bindEthBalanceMonitor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorTransactor{contract: contract}, nil
}

func NewEthBalanceMonitorFilterer(address common.Address, filterer bind.ContractFilterer) (*EthBalanceMonitorFilterer, error) {
	contract, err := bindEthBalanceMonitor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorFilterer{contract: contract}, nil
}

func bindEthBalanceMonitor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthBalanceMonitorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_EthBalanceMonitor *EthBalanceMonitorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthBalanceMonitor.Contract.EthBalanceMonitorCaller.contract.Call(opts, result, method, params...)
}

func (_EthBalanceMonitor *EthBalanceMonitorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.EthBalanceMonitorTransactor.contract.Transfer(opts)
}

func (_EthBalanceMonitor *EthBalanceMonitorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.EthBalanceMonitorTransactor.contract.Transact(opts, method, params...)
}

func (_EthBalanceMonitor *EthBalanceMonitorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthBalanceMonitor.Contract.contract.Call(opts, result, method, params...)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.contract.Transfer(opts)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.contract.Transact(opts, method, params...)
}

func (_EthBalanceMonitor *EthBalanceMonitorCaller) CheckUpkeep(opts *bind.CallOpts, arg0 []byte) (CheckUpkeep,

	error) {
	var out []interface{}
	err := _EthBalanceMonitor.contract.Call(opts, &out, "checkUpkeep", arg0)

	outstruct := new(CheckUpkeep)
	if err != nil {
		return *outstruct, err
	}

	outstruct.UpkeepNeeded = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PerformData = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

func (_EthBalanceMonitor *EthBalanceMonitorSession) CheckUpkeep(arg0 []byte) (CheckUpkeep,

	error) {
	return _EthBalanceMonitor.Contract.CheckUpkeep(&_EthBalanceMonitor.CallOpts, arg0)
}

func (_EthBalanceMonitor *EthBalanceMonitorCallerSession) CheckUpkeep(arg0 []byte) (CheckUpkeep,

	error) {
	return _EthBalanceMonitor.Contract.CheckUpkeep(&_EthBalanceMonitor.CallOpts, arg0)
}

func (_EthBalanceMonitor *EthBalanceMonitorCaller) GetAccountInfo(opts *bind.CallOpts, targetAddress common.Address) (GetAccountInfo,

	error) {
	var out []interface{}
	err := _EthBalanceMonitor.contract.Call(opts, &out, "getAccountInfo", targetAddress)

	outstruct := new(GetAccountInfo)
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.MinBalanceWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TopUpAmountWei = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastTopUpTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_EthBalanceMonitor *EthBalanceMonitorSession) GetAccountInfo(targetAddress common.Address) (GetAccountInfo,

	error) {
	return _EthBalanceMonitor.Contract.GetAccountInfo(&_EthBalanceMonitor.CallOpts, targetAddress)
}

func (_EthBalanceMonitor *EthBalanceMonitorCallerSession) GetAccountInfo(targetAddress common.Address) (GetAccountInfo,

	error) {
	return _EthBalanceMonitor.Contract.GetAccountInfo(&_EthBalanceMonitor.CallOpts, targetAddress)
}

func (_EthBalanceMonitor *EthBalanceMonitorCaller) GetKeeperRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthBalanceMonitor.contract.Call(opts, &out, "getKeeperRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EthBalanceMonitor *EthBalanceMonitorSession) GetKeeperRegistryAddress() (common.Address, error) {
	return _EthBalanceMonitor.Contract.GetKeeperRegistryAddress(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCallerSession) GetKeeperRegistryAddress() (common.Address, error) {
	return _EthBalanceMonitor.Contract.GetKeeperRegistryAddress(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCaller) GetMinWaitPeriodSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthBalanceMonitor.contract.Call(opts, &out, "getMinWaitPeriodSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_EthBalanceMonitor *EthBalanceMonitorSession) GetMinWaitPeriodSeconds() (*big.Int, error) {
	return _EthBalanceMonitor.Contract.GetMinWaitPeriodSeconds(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCallerSession) GetMinWaitPeriodSeconds() (*big.Int, error) {
	return _EthBalanceMonitor.Contract.GetMinWaitPeriodSeconds(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCaller) GetUnderfundedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EthBalanceMonitor.contract.Call(opts, &out, "getUnderfundedAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EthBalanceMonitor *EthBalanceMonitorSession) GetUnderfundedAddresses() ([]common.Address, error) {
	return _EthBalanceMonitor.Contract.GetUnderfundedAddresses(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCallerSession) GetUnderfundedAddresses() ([]common.Address, error) {
	return _EthBalanceMonitor.Contract.GetUnderfundedAddresses(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCaller) GetWatchList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EthBalanceMonitor.contract.Call(opts, &out, "getWatchList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_EthBalanceMonitor *EthBalanceMonitorSession) GetWatchList() ([]common.Address, error) {
	return _EthBalanceMonitor.Contract.GetWatchList(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCallerSession) GetWatchList() ([]common.Address, error) {
	return _EthBalanceMonitor.Contract.GetWatchList(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthBalanceMonitor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_EthBalanceMonitor *EthBalanceMonitorSession) Owner() (common.Address, error) {
	return _EthBalanceMonitor.Contract.Owner(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCallerSession) Owner() (common.Address, error) {
	return _EthBalanceMonitor.Contract.Owner(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthBalanceMonitor.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_EthBalanceMonitor *EthBalanceMonitorSession) Paused() (bool, error) {
	return _EthBalanceMonitor.Contract.Paused(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorCallerSession) Paused() (bool, error) {
	return _EthBalanceMonitor.Contract.Paused(&_EthBalanceMonitor.CallOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "acceptOwnership")
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.AcceptOwnership(&_EthBalanceMonitor.TransactOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.AcceptOwnership(&_EthBalanceMonitor.TransactOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "pause")
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) Pause() (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.Pause(&_EthBalanceMonitor.TransactOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) Pause() (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.Pause(&_EthBalanceMonitor.TransactOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) PerformUpkeep(opts *bind.TransactOpts, performData []byte) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "performUpkeep", performData)
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.PerformUpkeep(&_EthBalanceMonitor.TransactOpts, performData)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.PerformUpkeep(&_EthBalanceMonitor.TransactOpts, performData)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) SetKeeperRegistryAddress(opts *bind.TransactOpts, keeperRegistryAddress common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "setKeeperRegistryAddress", keeperRegistryAddress)
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) SetKeeperRegistryAddress(keeperRegistryAddress common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.SetKeeperRegistryAddress(&_EthBalanceMonitor.TransactOpts, keeperRegistryAddress)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) SetKeeperRegistryAddress(keeperRegistryAddress common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.SetKeeperRegistryAddress(&_EthBalanceMonitor.TransactOpts, keeperRegistryAddress)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) SetMinWaitPeriodSeconds(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "setMinWaitPeriodSeconds", period)
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) SetMinWaitPeriodSeconds(period *big.Int) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.SetMinWaitPeriodSeconds(&_EthBalanceMonitor.TransactOpts, period)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) SetMinWaitPeriodSeconds(period *big.Int) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.SetMinWaitPeriodSeconds(&_EthBalanceMonitor.TransactOpts, period)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) SetWatchList(opts *bind.TransactOpts, addresses []common.Address, minBalancesWei []*big.Int, topUpAmountsWei []*big.Int) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "setWatchList", addresses, minBalancesWei, topUpAmountsWei)
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) SetWatchList(addresses []common.Address, minBalancesWei []*big.Int, topUpAmountsWei []*big.Int) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.SetWatchList(&_EthBalanceMonitor.TransactOpts, addresses, minBalancesWei, topUpAmountsWei)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) SetWatchList(addresses []common.Address, minBalancesWei []*big.Int, topUpAmountsWei []*big.Int) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.SetWatchList(&_EthBalanceMonitor.TransactOpts, addresses, minBalancesWei, topUpAmountsWei)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) TopUp(opts *bind.TransactOpts, needsFunding []common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "topUp", needsFunding)
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) TopUp(needsFunding []common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.TopUp(&_EthBalanceMonitor.TransactOpts, needsFunding)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) TopUp(needsFunding []common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.TopUp(&_EthBalanceMonitor.TransactOpts, needsFunding)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "transferOwnership", to)
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.TransferOwnership(&_EthBalanceMonitor.TransactOpts, to)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.TransferOwnership(&_EthBalanceMonitor.TransactOpts, to)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "unpause")
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) Unpause() (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.Unpause(&_EthBalanceMonitor.TransactOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) Unpause() (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.Unpause(&_EthBalanceMonitor.TransactOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, payee common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.Transact(opts, "withdraw", amount, payee)
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) Withdraw(amount *big.Int, payee common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.Withdraw(&_EthBalanceMonitor.TransactOpts, amount, payee)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) Withdraw(amount *big.Int, payee common.Address) (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.Withdraw(&_EthBalanceMonitor.TransactOpts, amount, payee)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthBalanceMonitor.contract.RawTransact(opts, nil)
}

func (_EthBalanceMonitor *EthBalanceMonitorSession) Receive() (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.Receive(&_EthBalanceMonitor.TransactOpts)
}

func (_EthBalanceMonitor *EthBalanceMonitorTransactorSession) Receive() (*types.Transaction, error) {
	return _EthBalanceMonitor.Contract.Receive(&_EthBalanceMonitor.TransactOpts)
}

type EthBalanceMonitorFundsAddedIterator struct {
	Event *EthBalanceMonitorFundsAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorFundsAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorFundsAdded)
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
		it.Event = new(EthBalanceMonitorFundsAdded)
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

func (it *EthBalanceMonitorFundsAddedIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorFundsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorFundsAdded struct {
	AmountAdded *big.Int
	NewBalance  *big.Int
	Sender      common.Address
	Raw         types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterFundsAdded(opts *bind.FilterOpts) (*EthBalanceMonitorFundsAddedIterator, error) {

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "FundsAdded")
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorFundsAddedIterator{contract: _EthBalanceMonitor.contract, event: "FundsAdded", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchFundsAdded(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorFundsAdded) (event.Subscription, error) {

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "FundsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorFundsAdded)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "FundsAdded", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParseFundsAdded(log types.Log) (*EthBalanceMonitorFundsAdded, error) {
	event := new(EthBalanceMonitorFundsAdded)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "FundsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EthBalanceMonitorFundsWithdrawnIterator struct {
	Event *EthBalanceMonitorFundsWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorFundsWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorFundsWithdrawn)
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
		it.Event = new(EthBalanceMonitorFundsWithdrawn)
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

func (it *EthBalanceMonitorFundsWithdrawnIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorFundsWithdrawn struct {
	AmountWithdrawn *big.Int
	Payee           common.Address
	Raw             types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts) (*EthBalanceMonitorFundsWithdrawnIterator, error) {

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "FundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorFundsWithdrawnIterator{contract: _EthBalanceMonitor.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorFundsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "FundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorFundsWithdrawn)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParseFundsWithdrawn(log types.Log) (*EthBalanceMonitorFundsWithdrawn, error) {
	event := new(EthBalanceMonitorFundsWithdrawn)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EthBalanceMonitorKeeperRegistryAddressUpdatedIterator struct {
	Event *EthBalanceMonitorKeeperRegistryAddressUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorKeeperRegistryAddressUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorKeeperRegistryAddressUpdated)
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
		it.Event = new(EthBalanceMonitorKeeperRegistryAddressUpdated)
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

func (it *EthBalanceMonitorKeeperRegistryAddressUpdatedIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorKeeperRegistryAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorKeeperRegistryAddressUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterKeeperRegistryAddressUpdated(opts *bind.FilterOpts) (*EthBalanceMonitorKeeperRegistryAddressUpdatedIterator, error) {

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "KeeperRegistryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorKeeperRegistryAddressUpdatedIterator{contract: _EthBalanceMonitor.contract, event: "KeeperRegistryAddressUpdated", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchKeeperRegistryAddressUpdated(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorKeeperRegistryAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "KeeperRegistryAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorKeeperRegistryAddressUpdated)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "KeeperRegistryAddressUpdated", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParseKeeperRegistryAddressUpdated(log types.Log) (*EthBalanceMonitorKeeperRegistryAddressUpdated, error) {
	event := new(EthBalanceMonitorKeeperRegistryAddressUpdated)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "KeeperRegistryAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EthBalanceMonitorMinWaitPeriodUpdatedIterator struct {
	Event *EthBalanceMonitorMinWaitPeriodUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorMinWaitPeriodUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorMinWaitPeriodUpdated)
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
		it.Event = new(EthBalanceMonitorMinWaitPeriodUpdated)
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

func (it *EthBalanceMonitorMinWaitPeriodUpdatedIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorMinWaitPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorMinWaitPeriodUpdated struct {
	OldMinWaitPeriod *big.Int
	NewMinWaitPeriod *big.Int
	Raw              types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterMinWaitPeriodUpdated(opts *bind.FilterOpts) (*EthBalanceMonitorMinWaitPeriodUpdatedIterator, error) {

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "MinWaitPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorMinWaitPeriodUpdatedIterator{contract: _EthBalanceMonitor.contract, event: "MinWaitPeriodUpdated", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchMinWaitPeriodUpdated(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorMinWaitPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "MinWaitPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorMinWaitPeriodUpdated)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "MinWaitPeriodUpdated", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParseMinWaitPeriodUpdated(log types.Log) (*EthBalanceMonitorMinWaitPeriodUpdated, error) {
	event := new(EthBalanceMonitorMinWaitPeriodUpdated)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "MinWaitPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EthBalanceMonitorOwnershipTransferRequestedIterator struct {
	Event *EthBalanceMonitorOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorOwnershipTransferRequested)
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
		it.Event = new(EthBalanceMonitorOwnershipTransferRequested)
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

func (it *EthBalanceMonitorOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EthBalanceMonitorOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorOwnershipTransferRequestedIterator{contract: _EthBalanceMonitor.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorOwnershipTransferRequested)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParseOwnershipTransferRequested(log types.Log) (*EthBalanceMonitorOwnershipTransferRequested, error) {
	event := new(EthBalanceMonitorOwnershipTransferRequested)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EthBalanceMonitorOwnershipTransferredIterator struct {
	Event *EthBalanceMonitorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorOwnershipTransferred)
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
		it.Event = new(EthBalanceMonitorOwnershipTransferred)
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

func (it *EthBalanceMonitorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EthBalanceMonitorOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorOwnershipTransferredIterator{contract: _EthBalanceMonitor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorOwnershipTransferred)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParseOwnershipTransferred(log types.Log) (*EthBalanceMonitorOwnershipTransferred, error) {
	event := new(EthBalanceMonitorOwnershipTransferred)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EthBalanceMonitorPausedIterator struct {
	Event *EthBalanceMonitorPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorPaused)
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
		it.Event = new(EthBalanceMonitorPaused)
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

func (it *EthBalanceMonitorPausedIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterPaused(opts *bind.FilterOpts) (*EthBalanceMonitorPausedIterator, error) {

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorPausedIterator{contract: _EthBalanceMonitor.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorPaused) (event.Subscription, error) {

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorPaused)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParsePaused(log types.Log) (*EthBalanceMonitorPaused, error) {
	event := new(EthBalanceMonitorPaused)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EthBalanceMonitorTopUpFailedIterator struct {
	Event *EthBalanceMonitorTopUpFailed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorTopUpFailedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorTopUpFailed)
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
		it.Event = new(EthBalanceMonitorTopUpFailed)
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

func (it *EthBalanceMonitorTopUpFailedIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorTopUpFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorTopUpFailed struct {
	Recipient common.Address
	Raw       types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterTopUpFailed(opts *bind.FilterOpts, recipient []common.Address) (*EthBalanceMonitorTopUpFailedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "TopUpFailed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorTopUpFailedIterator{contract: _EthBalanceMonitor.contract, event: "TopUpFailed", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchTopUpFailed(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorTopUpFailed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "TopUpFailed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorTopUpFailed)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "TopUpFailed", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParseTopUpFailed(log types.Log) (*EthBalanceMonitorTopUpFailed, error) {
	event := new(EthBalanceMonitorTopUpFailed)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "TopUpFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EthBalanceMonitorTopUpSucceededIterator struct {
	Event *EthBalanceMonitorTopUpSucceeded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorTopUpSucceededIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorTopUpSucceeded)
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
		it.Event = new(EthBalanceMonitorTopUpSucceeded)
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

func (it *EthBalanceMonitorTopUpSucceededIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorTopUpSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorTopUpSucceeded struct {
	Recipient common.Address
	Raw       types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterTopUpSucceeded(opts *bind.FilterOpts, recipient []common.Address) (*EthBalanceMonitorTopUpSucceededIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "TopUpSucceeded", recipientRule)
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorTopUpSucceededIterator{contract: _EthBalanceMonitor.contract, event: "TopUpSucceeded", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchTopUpSucceeded(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorTopUpSucceeded, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "TopUpSucceeded", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorTopUpSucceeded)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "TopUpSucceeded", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParseTopUpSucceeded(log types.Log) (*EthBalanceMonitorTopUpSucceeded, error) {
	event := new(EthBalanceMonitorTopUpSucceeded)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "TopUpSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type EthBalanceMonitorUnpausedIterator struct {
	Event *EthBalanceMonitorUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *EthBalanceMonitorUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceMonitorUnpaused)
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
		it.Event = new(EthBalanceMonitorUnpaused)
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

func (it *EthBalanceMonitorUnpausedIterator) Error() error {
	return it.fail
}

func (it *EthBalanceMonitorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type EthBalanceMonitorUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EthBalanceMonitorUnpausedIterator, error) {

	logs, sub, err := _EthBalanceMonitor.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EthBalanceMonitorUnpausedIterator{contract: _EthBalanceMonitor.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorUnpaused) (event.Subscription, error) {

	logs, sub, err := _EthBalanceMonitor.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(EthBalanceMonitorUnpaused)
				if err := _EthBalanceMonitor.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_EthBalanceMonitor *EthBalanceMonitorFilterer) ParseUnpaused(log types.Log) (*EthBalanceMonitorUnpaused, error) {
	event := new(EthBalanceMonitorUnpaused)
	if err := _EthBalanceMonitor.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type CheckUpkeep struct {
	UpkeepNeeded bool
	PerformData  []byte
}
type GetAccountInfo struct {
	IsActive           bool
	MinBalanceWei      *big.Int
	TopUpAmountWei     *big.Int
	LastTopUpTimestamp *big.Int
}

func (EthBalanceMonitorFundsAdded) Topic() common.Hash {
	return common.HexToHash("0xc6f3fb0fec49e4877342d4625d77a632541f55b7aae0f9d0b34c69b3478706dc")
}

func (EthBalanceMonitorFundsWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x6141b54b56b8a52a8c6f5cd2a857f6117b18ffbf4d46bd3106f300a839cbf5ea")
}

func (EthBalanceMonitorKeeperRegistryAddressUpdated) Topic() common.Hash {
	return common.HexToHash("0xb732223055abcde751d7a24272ffc8a3aa571cb72b443969a4199b7ecd59f8b9")
}

func (EthBalanceMonitorMinWaitPeriodUpdated) Topic() common.Hash {
	return common.HexToHash("0x04330086c73b1fe1e13cd47a61c692e7c4399b5de08ed94b7ab824684af09323")
}

func (EthBalanceMonitorOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (EthBalanceMonitorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (EthBalanceMonitorPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (EthBalanceMonitorTopUpFailed) Topic() common.Hash {
	return common.HexToHash("0xa9ff7a9b96721b0e16adb7de9db0764fbfd6a4516d4d165f9564e8c3755eb105")
}

func (EthBalanceMonitorTopUpSucceeded) Topic() common.Hash {
	return common.HexToHash("0x9eec55c371a49ce19e0a5792787c79b32dcf7d3490aa737436b49c0978ce9ce9")
}

func (EthBalanceMonitorUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_EthBalanceMonitor *EthBalanceMonitor) Address() common.Address {
	return _EthBalanceMonitor.address
}

type EthBalanceMonitorInterface interface {
	CheckUpkeep(opts *bind.CallOpts, arg0 []byte) (CheckUpkeep,

		error)

	GetAccountInfo(opts *bind.CallOpts, targetAddress common.Address) (GetAccountInfo,

		error)

	GetKeeperRegistryAddress(opts *bind.CallOpts) (common.Address, error)

	GetMinWaitPeriodSeconds(opts *bind.CallOpts) (*big.Int, error)

	GetUnderfundedAddresses(opts *bind.CallOpts) ([]common.Address, error)

	GetWatchList(opts *bind.CallOpts) ([]common.Address, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	PerformUpkeep(opts *bind.TransactOpts, performData []byte) (*types.Transaction, error)

	SetKeeperRegistryAddress(opts *bind.TransactOpts, keeperRegistryAddress common.Address) (*types.Transaction, error)

	SetMinWaitPeriodSeconds(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error)

	SetWatchList(opts *bind.TransactOpts, addresses []common.Address, minBalancesWei []*big.Int, topUpAmountsWei []*big.Int) (*types.Transaction, error)

	TopUp(opts *bind.TransactOpts, needsFunding []common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, amount *big.Int, payee common.Address) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterFundsAdded(opts *bind.FilterOpts) (*EthBalanceMonitorFundsAddedIterator, error)

	WatchFundsAdded(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorFundsAdded) (event.Subscription, error)

	ParseFundsAdded(log types.Log) (*EthBalanceMonitorFundsAdded, error)

	FilterFundsWithdrawn(opts *bind.FilterOpts) (*EthBalanceMonitorFundsWithdrawnIterator, error)

	WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorFundsWithdrawn) (event.Subscription, error)

	ParseFundsWithdrawn(log types.Log) (*EthBalanceMonitorFundsWithdrawn, error)

	FilterKeeperRegistryAddressUpdated(opts *bind.FilterOpts) (*EthBalanceMonitorKeeperRegistryAddressUpdatedIterator, error)

	WatchKeeperRegistryAddressUpdated(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorKeeperRegistryAddressUpdated) (event.Subscription, error)

	ParseKeeperRegistryAddressUpdated(log types.Log) (*EthBalanceMonitorKeeperRegistryAddressUpdated, error)

	FilterMinWaitPeriodUpdated(opts *bind.FilterOpts) (*EthBalanceMonitorMinWaitPeriodUpdatedIterator, error)

	WatchMinWaitPeriodUpdated(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorMinWaitPeriodUpdated) (event.Subscription, error)

	ParseMinWaitPeriodUpdated(log types.Log) (*EthBalanceMonitorMinWaitPeriodUpdated, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EthBalanceMonitorOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*EthBalanceMonitorOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EthBalanceMonitorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*EthBalanceMonitorOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*EthBalanceMonitorPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*EthBalanceMonitorPaused, error)

	FilterTopUpFailed(opts *bind.FilterOpts, recipient []common.Address) (*EthBalanceMonitorTopUpFailedIterator, error)

	WatchTopUpFailed(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorTopUpFailed, recipient []common.Address) (event.Subscription, error)

	ParseTopUpFailed(log types.Log) (*EthBalanceMonitorTopUpFailed, error)

	FilterTopUpSucceeded(opts *bind.FilterOpts, recipient []common.Address) (*EthBalanceMonitorTopUpSucceededIterator, error)

	WatchTopUpSucceeded(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorTopUpSucceeded, recipient []common.Address) (event.Subscription, error)

	ParseTopUpSucceeded(log types.Log) (*EthBalanceMonitorTopUpSucceeded, error)

	FilterUnpaused(opts *bind.FilterOpts) (*EthBalanceMonitorUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EthBalanceMonitorUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*EthBalanceMonitorUnpaused, error)

	Address() common.Address
}
