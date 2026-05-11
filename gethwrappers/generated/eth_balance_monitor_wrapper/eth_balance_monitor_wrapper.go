// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth_balance_monitor_wrapper

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

var EthBalanceMonitorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"keeperRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minWaitPeriodSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkUpkeep\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"upkeepNeeded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"performData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccountInfo\",\"inputs\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minBalanceWei\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"topUpAmountWei\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"lastTopUpTimestamp\",\"type\":\"uint56\",\"internalType\":\"uint56\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getKeeperRegistryAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"keeperRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinWaitPeriodSeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnderfundedAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWatchList\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"performUpkeep\",\"inputs\":[{\"name\":\"performData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeeperRegistryAddress\",\"inputs\":[{\"name\":\"keeperRegistryAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinWaitPeriodSeconds\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWatchList\",\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"minBalancesWei\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"topUpAmountsWei\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"topUp\",\"inputs\":[{\"name\":\"needsFunding\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"payee\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FundsAdded\",\"inputs\":[{\"name\":\"amountAdded\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FundsWithdrawn\",\"inputs\":[{\"name\":\"amountWithdrawn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"payee\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"KeeperRegistryAddressUpdated\",\"inputs\":[{\"name\":\"oldAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinWaitPeriodUpdated\",\"inputs\":[{\"name\":\"oldMinWaitPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMinWaitPeriod\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TopUpFailed\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TopUpSucceeded\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DuplicateAddress\",\"inputs\":[{\"name\":\"duplicate\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWatchList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyKeeperRegistry\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001ffc38038062001ffc8339810160408190526200003491620002c7565b33806000816200008b5760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615620000be57620000be81620000ec565b50506001805460ff60a01b1916905550620000d98262000197565b620000e4816200021e565b505062000303565b336001600160a01b03821603620001465760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640162000082565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b620001a162000269565b6001600160a01b038116620001b557600080fd5b600254604080516001600160a01b03928316815291831660208301527fb732223055abcde751d7a24272ffc8a3aa571cb72b443969a4199b7ecd59f8b9910160405180910390a1600280546001600160a01b0319166001600160a01b0392909216919091179055565b6200022862000269565b60035460408051918252602082018390527f04330086c73b1fe1e13cd47a61c692e7c4399b5de08ed94b7ab824684af09323910160405180910390a1600355565b6000546001600160a01b03163314620002c55760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640162000082565b565b60008060408385031215620002db57600080fd5b82516001600160a01b0381168114620002f357600080fd5b6020939093015192949293505050565b611ce980620003136000396000f3fe60806040526004361061012c5760003560e01c8063728584b7116100a55780638456cb591161007457806394555114116100595780639455511414610473578063b1d52fa014610493578063f2fde38b146104b357600080fd5b80638456cb59146104335780638da5cb5b1461044857600080fd5b8063728584b7146102bf57806379ba5097146102d45780637b510fe8146102e9578063810623e3146103e757600080fd5b80633f85861f116100fc5780634585e33b116100e15780634585e33b146102365780635c975abb146102565780636e04ff0d1461029157600080fd5b80633f85861f146101f857806341d2052e1461021857600080fd5b8062f714ce146101765780630b67ddce146101985780633e4ca677146101c35780633f4ba83a146101e357600080fd5b366101715760408051348152476020820152338183015290517fc6f3fb0fec49e4877342d4625d77a632541f55b7aae0f9d0b34c69b3478706dc9181900360600190a1005b600080fd5b34801561018257600080fd5b5061019661019136600461183c565b6104d3565b005b3480156101a457600080fd5b506101ad610591565b6040516101ba919061186c565b60405180910390f35b3480156101cf57600080fd5b506101966101de366004611905565b610868565b3480156101ef57600080fd5b50610196610ba9565b34801561020457600080fd5b506101966102133660046119e8565b610bbb565b34801561022457600080fd5b506003546040519081526020016101ba565b34801561024257600080fd5b50610196610251366004611a01565b610c04565b34801561026257600080fd5b5060015474010000000000000000000000000000000000000000900460ff1660405190151581526020016101ba565b34801561029d57600080fd5b506102b16102ac366004611a01565b610c76565b6040516101ba929190611a73565b3480156102cb57600080fd5b506101ad610cc0565b3480156102e057600080fd5b50610196610d2f565b3480156102f557600080fd5b506103a7610304366004611ae9565b73ffffffffffffffffffffffffffffffffffffffff166000908152600560209081526040918290208251608081018452905460ff8116151580835261010082046bffffffffffffffffffffffff9081169484018590526d010000000000000000000000000083041694830185905279010000000000000000000000000000000000000000000000000090910466ffffffffffffff16606090920182905293919291565b6040516101ba949392919093151584526bffffffffffffffffffffffff92831660208501529116604083015266ffffffffffffff16606082015260800190565b3480156103f357600080fd5b5060025473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ba565b34801561043f57600080fd5b50610196610e31565b34801561045457600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff1661040e565b34801561047f57600080fd5b5061019661048e366004611ae9565b610e41565b34801561049f57600080fd5b506101966104ae366004611b52565b610f04565b3480156104bf57600080fd5b506101966104ce366004611ae9565b611401565b6104db611412565b73ffffffffffffffffffffffffffffffffffffffff81166104fb57600080fd5b6040805183815273ffffffffffffffffffffffffffffffffffffffff831660208201527f6141b54b56b8a52a8c6f5cd2a857f6117b18ffbf4d46bd3106f300a839cbf5ea910160405180910390a160405173ffffffffffffffffffffffffffffffffffffffff82169083156108fc029084906000818181858888f1935050505015801561058c573d6000803e3d6000fd5b505050565b6060600060048054806020026020016040519081016040528092919081815260200182805480156105f857602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116105cd575b505050505090506000815167ffffffffffffffff81111561061b5761061b6118c6565b604051908082528060200260200182016040528015610644578160200160208202803683370190505b50600354604080516080810182526000808252602082018190529181018290526060810182905292935091479060005b8651811015610850576005600088838151811061069357610693611bec565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528181019290925260409081016000208151608081018352905460ff81161515825261010081046bffffffffffffffffffffffff908116948301949094526d010000000000000000000000000081049093169181019190915279010000000000000000000000000000000000000000000000000090910466ffffffffffffff1660608201819052909250429061074d908690611c4a565b1115801561076d575081604001516bffffffffffffffffffffffff168310155b80156107ba575081602001516bffffffffffffffffffffffff1687828151811061079957610799611bec565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1631105b1561083e578681815181106107d1576107d1611bec565b60200260200101518686815181106107eb576107eb611bec565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101528461081a81611c63565b95505081604001516bffffffffffffffffffffffff168361083b9190611c9b565b92505b8061084881611c63565b915050610674565b508551841461085d578385525b509295945050505050565b610870611493565b6003546040805160808101825260008082526020820181905291810182905260608101829052905b8351811015610ba257600560008583815181106108b7576108b7611bec565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528181019290925260409081016000208151608081018352905460ff81161580158084526bffffffffffffffffffffffff61010084048116968501969096526d010000000000000000000000000083049095169383019390935266ffffffffffffff790100000000000000000000000000000000000000000000000000909104166060820152935061098457504283836060015166ffffffffffffff166109819190611c4a565b11155b80156109d1575081602001516bffffffffffffffffffffffff168482815181106109b0576109b0611bec565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1631105b15610b805760008482815181106109ea576109ea611bec565b602002602001015173ffffffffffffffffffffffffffffffffffffffff166108fc84604001516bffffffffffffffffffffffff169081150290604051600060405180830381858888f1935050505090508015610b21574260056000878581518110610a5757610a57611bec565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160196101000a81548166ffffffffffffff021916908366ffffffffffffff160217905550848281518110610ad257610ad2611bec565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167f9eec55c371a49ce19e0a5792787c79b32dcf7d3490aa737436b49c0978ce9ce960405160405180910390a2610b7e565b848281518110610b3357610b33611bec565b602002602001015173ffffffffffffffffffffffffffffffffffffffff167fa9ff7a9b96721b0e16adb7de9db0764fbfd6a4516d4d165f9564e8c3755eb10560405160405180910390a25b505b61d6d85a1015610b905750505050565b80610b9a81611c63565b915050610898565b5050505b50565b610bb1611412565b610bb9611518565b565b610bc3611412565b60035460408051918252602082018390527f04330086c73b1fe1e13cd47a61c692e7c4399b5de08ed94b7ab824684af09323910160405180910390a1600355565b60025473ffffffffffffffffffffffffffffffffffffffff163314610c55576040517fd3a6803400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c5d611493565b6000610c6b82840184611905565b905061058c81610868565b60006060610c82611493565b6000610c8c610591565b90506000815111925080604051602001610ca6919061186c565b6040516020818303038152906040529150505b9250929050565b60606004805480602002602001604051908101604052809291908181526020018280548015610d2557602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610cfa575b5050505050905090565b60015473ffffffffffffffffffffffffffffffffffffffff163314610db5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b610e39611412565b610bb9611595565b610e49611412565b73ffffffffffffffffffffffffffffffffffffffff8116610e6957600080fd5b6002546040805173ffffffffffffffffffffffffffffffffffffffff928316815291831660208301527fb732223055abcde751d7a24272ffc8a3aa571cb72b443969a4199b7ecd59f8b9910160405180910390a1600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b610f0c611412565b8483141580610f1b5750848114155b15610f52576040517f3869bbe600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006004805480602002602001604051908101604052809291908181526020018280548015610fb757602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610f8c575b5050505050905060005b815181101561105157600060056000848481518110610fe257610fe2611bec565b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff16825281019190915260400160002080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169115159190911790558061104981611c63565b915050610fc1565b5060005b868110156113ea576005600089898481811061107357611073611bec565b90506020020160208101906110889190611ae9565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040016000205460ff1615611129578787828181106110c9576110c9611bec565b90506020020160208101906110de9190611ae9565b6040517f9f2277f300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602401610dac565b600088888381811061113d5761113d611bec565b90506020020160208101906111529190611ae9565b73ffffffffffffffffffffffffffffffffffffffff160361119f576040517f3869bbe600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8383828181106111b1576111b1611bec565b90506020020160208101906111c69190611cae565b6bffffffffffffffffffffffff1660000361120d576040517f3869bbe600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604051806080016040528060011515815260200187878481811061123357611233611bec565b90506020020160208101906112489190611cae565b6bffffffffffffffffffffffff16815260200185858481811061126d5761126d611bec565b90506020020160208101906112829190611cae565b6bffffffffffffffffffffffff168152602001600066ffffffffffffff16815250600560008a8a858181106112b9576112b9611bec565b90506020020160208101906112ce9190611ae9565b73ffffffffffffffffffffffffffffffffffffffff168152602080820192909252604090810160002083518154938501519285015160609095015166ffffffffffffff167901000000000000000000000000000000000000000000000000000278ffffffffffffffffffffffffffffffffffffffffffffffffff6bffffffffffffffffffffffff9687166d010000000000000000000000000002166cffffffffffffffffffffffffff96909416610100027fffffffffffffffffffffffffffffffffffffff000000000000000000000000ff921515929092167fffffffffffffffffffffffffffffffffffffff0000000000000000000000000090951694909417179390931617179055806113e281611c63565b915050611055565b506113f76004888861177d565b5050505050505050565b611409611412565b610ba681611604565b60005473ffffffffffffffffffffffffffffffffffffffff163314610bb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610dac565b60015474010000000000000000000000000000000000000000900460ff1615610bb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610dac565b6115206116f9565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b61159d611493565b600180547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861156b3390565b3373ffffffffffffffffffffffffffffffffffffffff821603611683576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610dac565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60015474010000000000000000000000000000000000000000900460ff16610bb9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610dac565b8280548282559060005260206000209081019282156117f5579160200282015b828111156117f55781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84351617825560209092019160019091019061179d565b50611801929150611805565b5090565b5b808211156118015760008155600101611806565b73ffffffffffffffffffffffffffffffffffffffff81168114610ba657600080fd5b6000806040838503121561184f57600080fd5b8235915060208301356118618161181a565b809150509250929050565b6020808252825182820181905260009190848201906040850190845b818110156118ba57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101611888565b50909695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b80356119008161181a565b919050565b6000602080838503121561191857600080fd5b823567ffffffffffffffff8082111561193057600080fd5b818501915085601f83011261194457600080fd5b813581811115611956576119566118c6565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f83011681018181108582111715611999576119996118c6565b6040529182528482019250838101850191888311156119b757600080fd5b938501935b828510156119dc576119cd856118f5565b845293850193928501926119bc565b98975050505050505050565b6000602082840312156119fa57600080fd5b5035919050565b60008060208385031215611a1457600080fd5b823567ffffffffffffffff80821115611a2c57600080fd5b818501915085601f830112611a4057600080fd5b813581811115611a4f57600080fd5b866020828501011115611a6157600080fd5b60209290920196919550909350505050565b821515815260006020604081840152835180604085015260005b81811015611aa957858101830151858201606001528201611a8d565b5060006060828601015260607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509392505050565b600060208284031215611afb57600080fd5b8135611b068161181a565b9392505050565b60008083601f840112611b1f57600080fd5b50813567ffffffffffffffff811115611b3757600080fd5b6020830191508360208260051b8501011115610cb957600080fd5b60008060008060008060608789031215611b6b57600080fd5b863567ffffffffffffffff80821115611b8357600080fd5b611b8f8a838b01611b0d565b90985096506020890135915080821115611ba857600080fd5b611bb48a838b01611b0d565b90965094506040890135915080821115611bcd57600080fd5b50611bda89828a01611b0d565b979a9699509497509295939492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115611c5d57611c5d611c1b565b92915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c9457611c94611c1b565b5060010190565b81810381811115611c5d57611c5d611c1b565b600060208284031215611cc057600080fd5b81356bffffffffffffffffffffffff81168114611b0657600080fdfea164736f6c6343000813000a",
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

func (_EthBalanceMonitor *EthBalanceMonitor) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _EthBalanceMonitor.abi.Events["FundsAdded"].ID:
		return _EthBalanceMonitor.ParseFundsAdded(log)
	case _EthBalanceMonitor.abi.Events["FundsWithdrawn"].ID:
		return _EthBalanceMonitor.ParseFundsWithdrawn(log)
	case _EthBalanceMonitor.abi.Events["KeeperRegistryAddressUpdated"].ID:
		return _EthBalanceMonitor.ParseKeeperRegistryAddressUpdated(log)
	case _EthBalanceMonitor.abi.Events["MinWaitPeriodUpdated"].ID:
		return _EthBalanceMonitor.ParseMinWaitPeriodUpdated(log)
	case _EthBalanceMonitor.abi.Events["OwnershipTransferRequested"].ID:
		return _EthBalanceMonitor.ParseOwnershipTransferRequested(log)
	case _EthBalanceMonitor.abi.Events["OwnershipTransferred"].ID:
		return _EthBalanceMonitor.ParseOwnershipTransferred(log)
	case _EthBalanceMonitor.abi.Events["Paused"].ID:
		return _EthBalanceMonitor.ParsePaused(log)
	case _EthBalanceMonitor.abi.Events["TopUpFailed"].ID:
		return _EthBalanceMonitor.ParseTopUpFailed(log)
	case _EthBalanceMonitor.abi.Events["TopUpSucceeded"].ID:
		return _EthBalanceMonitor.ParseTopUpSucceeded(log)
	case _EthBalanceMonitor.abi.Events["Unpaused"].ID:
		return _EthBalanceMonitor.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
