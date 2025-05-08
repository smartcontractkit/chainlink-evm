// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc677_helper

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

var BurnMintERC677HelperMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreaseApproval\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"drip\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBurners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinters\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantBurnRole\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantMintRole\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseApproval\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBurner\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMinter\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revokeBurnRole\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeMintRole\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferAndCall\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnAccessGranted\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BurnAccessRevoked\",\"inputs\":[{\"name\":\"burner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintAccessGranted\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintAccessRevoked\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferRequested\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SenderNotBurner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SenderNotMinter\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60c0604052346103a4576128cd80380380610019816103a9565b9283398101906040818303126103a45780516001600160401b0381116103a457826100459183016103ce565b60208201519092906001600160401b0381116103a45761006592016103ce565b81516001600160401b0381116102af57600354600181811c9116801561039a575b602082101461028f57601f8111610335575b50602092601f82116001146102d057928192936000926102c5575b50508160011b916000199060031b1c1916176003555b80516001600160401b0381116102af57600454600181811c911680156102a5575b602082101461028f57601f811161022a575b50602091601f82116001146101c6579181926000926101bb575b50508160011b916000199060031b1c1916176004555b331561017657600580546001600160a01b031916331790556012608052600060a052604051612493908161043a82396080518161127e015260a0518181816104450152610f190152f35b60405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f00000000000000006044820152606490fd5b015190503880610116565b601f198216926004600052806000209160005b858110610212575083600195106101f9575b505050811b0160045561012c565b015160001960f88460031b161c191690553880806101eb565b919260206001819286850151815501940192016101d9565b60046000527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f830160051c81019160208410610285575b601f0160051c01905b81811061027957506100fc565b6000815560010161026c565b9091508190610263565b634e487b7160e01b600052602260045260246000fd5b90607f16906100ea565b634e487b7160e01b600052604160045260246000fd5b0151905038806100b3565b601f198216936003600052806000209160005b86811061031d5750836001959610610304575b505050811b016003556100c9565b015160001960f88460031b161c191690553880806102f6565b919260206001819286850151815501940192016102e3565b60036000527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f830160051c81019160208410610390575b601f0160051c01905b8181106103845750610098565b60008155600101610377565b909150819061036e565b90607f1690610086565b600080fd5b6040519190601f01601f191682016001600160401b038111838210176102af57604052565b81601f820112156103a4578051906001600160401b0382116102af576103fd601f8301601f19166020016103a9565b92828452602083830101116103a45760005b82811061042457505060206000918301015290565b8060208092840101518282870101520161040f56fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a7146117f75750806306fdde031461171c578063095ea7b31461156157806318160ddd1461152557806323b872dd146112a2578063313ce5671461124657806339509351146112095780634000aea01461101757806340c10f1914610ea157806342966c6814610e1d5780634334614a14610db75780634f5632f814610d2957806366188463146106d957806367a5cd0614610be35780636b32810b14610b5057806370a0823114610aee57806379ba5097146109d157806379cc6790146106de57806386fe8b43146109325780638da5cb5b146108e05780638fd6a6ac146108e057806395d89b411461077a5780639dc29fac146106de578063a457c2d7146106d9578063a9059cbb14610693578063aa271e1a14610623578063c2e3273d14610595578063c630948d146104f6578063c64d0ebc14610468578063d5abeb011461040f578063d73dd623146103c7578063dd62ed3e1461034c578063f2fde38b146102275763f81094f31461019757600080fd5b346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff6101e361197a565b6101eb611ec5565b166101f5816122a1565b6101fd575080f35b7fed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e98280a280f35b80fd5b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff61027461197a565b61027c611ec5565b163381146102ee57807fffffffffffffffffffffffff0000000000000000000000000000000000000000600654161760065573ffffffffffffffffffffffffffffffffffffffff600554167fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12788380a380f35b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152fd5b50346102245760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff604061039b61197a565b92826103a56119a2565b9416815260016020522091166000526020526020604060002054604051908152f35b50346102245760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245761040b61040261197a565b60243590611c34565b5080f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245760206040517f00000000000000000000000000000000000000000000000000000000000000008152f35b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff6104b561197a565b6104bd611ec5565b166104c78161242c565b6104cf575080f35b7f92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad8280a280f35b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff61054361197a565b61054b611ec5565b16610555816123cc565b61056b575b610562611ec5565b6104c78161242c565b807fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea8380a261055a565b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff6105e261197a565b6105ea611ec5565b166105f4816123cc565b6105fc575080f35b7fe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea8280a280f35b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602061068973ffffffffffffffffffffffffffffffffffffffff61067561197a565b166000526008602052604060002054151590565b6040519015158152f35b50346102245760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245760206106896106d061197a565b60243590611cd3565b611a35565b50346102245760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245761071661197a565b6024359061073133600052600a602052604060002054151590565b1561074e579061074b91610746823383611d6a565b611f44565b80f35b6024837fc820b10b00000000000000000000000000000000000000000000000000000000815233600452fd5b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610224576040519080600454908160011c916001811680156108d6575b6020841081146108a9578386529081156108645750600114610807575b610803846107ef818603826119c5565b60405191829160208352602083019061191b565b0390f35b600481527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b939250905b80821061084a575090915081016020016107ef826107df565b919260018160209254838588010152019101909291610831565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208087019190915292151560051b850190920192506107ef91508390506107df565b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526022600452fd5b92607f16926107c2565b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602073ffffffffffffffffffffffffffffffffffffffff60055416604051908152f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245760405180916020600954928381520191600982527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af915b8181106109bb57610803856109af818703826119c5565b60405191829182611ba8565b8254845260209093019260019283019201610998565b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245760065473ffffffffffffffffffffffffffffffffffffffff81163303610a90577fffffffffffffffffffffffff00000000000000000000000000000000000000006005549133828416176005551660065573ffffffffffffffffffffffffffffffffffffffff3391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152fd5b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457604060209173ffffffffffffffffffffffffffffffffffffffff610b4061197a565b1681528083522054604051908152f35b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245760405180916020600754928381520191600782527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688915b818110610bcd57610803856109af818703826119c5565b8254845260209093019260019283019201610bb6565b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff610c3061197a565b168015610ccb57600254670de0b6b3a76400008101809111610c9e576002558082528160205260408220670de0b6b3a76400008154019055817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020604051670de0b6b3a76400008152a380f35b6024837f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245773ffffffffffffffffffffffffffffffffffffffff610d7661197a565b610d7e611ec5565b16610d888161210b565b610d90575080f35b7f0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c8280a280f35b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602061068973ffffffffffffffffffffffffffffffffffffffff610e0961197a565b16600052600a602052604060002054151590565b50346102245760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457610e6433600052600a602052604060002054151590565b15610e755761074b60043533611f44565b807fc820b10b000000000000000000000000000000000000000000000000000000006024925233600452fd5b50346102245760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457610ed961197a565b60243590610ef4336000526008602052604060002054151590565b15610feb5773ffffffffffffffffffffffffffffffffffffffff1690308214610fe7577f00000000000000000000000000000000000000000000000000000000000000008015159081610fd2575b50610f9b578115610ccb577fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef602082610f7e8694600254611bf8565b60025584845283825260408420818154019055604051908152a380f35b82610faa602492600254611bf8565b7fcbbf1113000000000000000000000000000000000000000000000000000000008252600452fd5b9050610fe082600254611bf8565b1138610f42565b8280fd5b6024837fe2c8c9d500000000000000000000000000000000000000000000000000000000815233600452fd5b50346102245760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245761104f61197a565b6024359060443567ffffffffffffffff811161120557366023820112156112055780600401359067ffffffffffffffff82116111d8578490604051926110bd60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601856119c5565b8084523660248284010111610fe7578060246020930183860137830101526110e58383611cd3565b5073ffffffffffffffffffffffffffffffffffffffff82169182604051858152604060208201527fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1633918061113d604082018861191b565b0390a33b611151575b602060405160018152f35b8184923b15610fe7576111a793836040518096819582947fa4c0ed36000000000000000000000000000000000000000000000000000000008452336004850152602484015260606044840152606483019061191b565b03925af180156111cd576111bd575b8080611146565b816111c7916119c5565b386111b6565b6040513d84823e3d90fd5b6024857f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b8380fd5b50346102245760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602061068961040261197a565b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261022457602060405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b50346102245760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610224576112da61197a565b906112e36119a2565b73ffffffffffffffffffffffffffffffffffffffff60443591611307833387611d6a565b1692308414610fe75773ffffffffffffffffffffffffffffffffffffffff169182156114a157831561141d578281528060205260408120549180831061139957604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815280845220818154019055604051908152a3602060405160018152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152fd5b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610224576020600254604051908152f35b50346102245760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102245761159961197a565b73ffffffffffffffffffffffffffffffffffffffff6024359116913083146102245733156116995782156116155760408291338152600160205281812085825260205220556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152fd5b503461022457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610224576040519080600354908160011c916001811680156117ed575b6020841081146108a957838652908115610864575060011461179057610803846107ef818603826119c5565b600381527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b939250905b8082106117d3575090915081016020016107ef826107df565b9192600181602092548385880101520191019092916117ba565b92607f1692611764565b9050346119175760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112611917576004357fffffffff000000000000000000000000000000000000000000000000000000008116809103610fe757602092507f36372b070000000000000000000000000000000000000000000000000000000081149081156118ed575b81156118c3575b8115611899575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438611892565b7fe6599b4d000000000000000000000000000000000000000000000000000000008114915061188b565b7f4000aea00000000000000000000000000000000000000000000000000000000081149150611884565b5080fd5b919082519283825260005b8481106119655750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b80602080928401015182828601015201611926565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361199d57565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361199d57565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117611a0657604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b3461199d5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261199d57611a6c61197a565b6024359060009133835260016020526040832073ffffffffffffffffffffffffffffffffffffffff83168452602052604083205490808210611b245773ffffffffffffffffffffffffffffffffffffffff91039116913083146102245733156116995782156116155760408291338152600160205281812085825260205220556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a360206001604051908152f35b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152fd5b602060408183019282815284518094520192019060005b818110611bcc5750505090565b825173ffffffffffffffffffffffffffffffffffffffff16845260209384019390920191600101611bbf565b91908201809211611c0557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b90611c7073ffffffffffffffffffffffffffffffffffffffff913360005260016020526040600020838516600052602052604060002054611bf8565b91169030821461199d57331561169957811561161557336000526001602052604060002082600052602052806040600020556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3600190565b73ffffffffffffffffffffffffffffffffffffffff169030821461199d5733156114a157811561141d573360005260006020526040600020548181106113995781903360005260006020520360406000205581600052600060205260406000208181540190556040519081527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203392a3600190565b73ffffffffffffffffffffffffffffffffffffffff909291921690816000526001602052604060002073ffffffffffffffffffffffffffffffffffffffff8416600052602052604060002054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611de5575b50505050565b808210611e675773ffffffffffffffffffffffffffffffffffffffff910392169130831461199d5781156116995782156116155760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a338808080611ddf565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152fd5b73ffffffffffffffffffffffffffffffffffffffff600554163303611ee657565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152fd5b73ffffffffffffffffffffffffffffffffffffffff1680156120405780600052600060205260406000205491808310611fbc576020817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926000958587528684520360408620558060025403600255604051908152a3565b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152fd5b60846040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152fd5b80548210156120dc5760005260206000200190600090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000818152600a6020526040902054801561229a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611c0557600954907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611c055780820361222b575b50505060095480156121fc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016121b98160096120c4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600955600052600a60205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b61228261223c61224d9360096120c4565b90549060031b1c92839260096120c4565b81939154907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060031b92831b921b19161790565b9055600052600a602052604060002055388080612180565b5050600090565b600081815260086020526040902054801561229a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101818111611c0557600754907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201918211611c0557808203612392575b50505060075480156121fc577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161234f8160076120c4565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055600755600052600860205260006040812055600190565b6123b46123a361224d9360076120c4565b90549060031b1c92839260076120c4565b90556000526008602052604060002055388080612316565b806000526008602052604060002054156000146124265760075468010000000000000000811015611a065761240d61224d82600185940160075560076120c4565b9055600754906000526008602052604060002055600190565b50600090565b80600052600a602052604060002054156000146124265760095468010000000000000000811015611a065761246d61224d82600185940160095560096120c4565b905560095490600052600a60205260406000205560019056fea164736f6c634300081a000a",
}

var BurnMintERC677HelperABI = BurnMintERC677HelperMetaData.ABI

var BurnMintERC677HelperBin = BurnMintERC677HelperMetaData.Bin

func DeployBurnMintERC677Helper(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *BurnMintERC677Helper, error) {
	parsed, err := BurnMintERC677HelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC677HelperBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC677Helper{address: address, abi: *parsed, BurnMintERC677HelperCaller: BurnMintERC677HelperCaller{contract: contract}, BurnMintERC677HelperTransactor: BurnMintERC677HelperTransactor{contract: contract}, BurnMintERC677HelperFilterer: BurnMintERC677HelperFilterer{contract: contract}}, nil
}

type BurnMintERC677Helper struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC677HelperCaller
	BurnMintERC677HelperTransactor
	BurnMintERC677HelperFilterer
}

type BurnMintERC677HelperCaller struct {
	contract *bind.BoundContract
}

type BurnMintERC677HelperTransactor struct {
	contract *bind.BoundContract
}

type BurnMintERC677HelperFilterer struct {
	contract *bind.BoundContract
}

type BurnMintERC677HelperSession struct {
	Contract     *BurnMintERC677Helper
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC677HelperCallerSession struct {
	Contract *BurnMintERC677HelperCaller
	CallOpts bind.CallOpts
}

type BurnMintERC677HelperTransactorSession struct {
	Contract     *BurnMintERC677HelperTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC677HelperRaw struct {
	Contract *BurnMintERC677Helper
}

type BurnMintERC677HelperCallerRaw struct {
	Contract *BurnMintERC677HelperCaller
}

type BurnMintERC677HelperTransactorRaw struct {
	Contract *BurnMintERC677HelperTransactor
}

func NewBurnMintERC677Helper(address common.Address, backend bind.ContractBackend) (*BurnMintERC677Helper, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC677HelperABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC677Helper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677Helper{address: address, abi: abi, BurnMintERC677HelperCaller: BurnMintERC677HelperCaller{contract: contract}, BurnMintERC677HelperTransactor: BurnMintERC677HelperTransactor{contract: contract}, BurnMintERC677HelperFilterer: BurnMintERC677HelperFilterer{contract: contract}}, nil
}

func NewBurnMintERC677HelperCaller(address common.Address, caller bind.ContractCaller) (*BurnMintERC677HelperCaller, error) {
	contract, err := bindBurnMintERC677Helper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperCaller{contract: contract}, nil
}

func NewBurnMintERC677HelperTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC677HelperTransactor, error) {
	contract, err := bindBurnMintERC677Helper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperTransactor{contract: contract}, nil
}

func NewBurnMintERC677HelperFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC677HelperFilterer, error) {
	contract, err := bindBurnMintERC677Helper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperFilterer{contract: contract}, nil
}

func bindBurnMintERC677Helper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC677HelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC677Helper.Contract.BurnMintERC677HelperCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.BurnMintERC677HelperTransactor.contract.Transfer(opts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.BurnMintERC677HelperTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC677Helper.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.contract.Transfer(opts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.Allowance(&_BurnMintERC677Helper.CallOpts, owner, spender)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.Allowance(&_BurnMintERC677Helper.CallOpts, owner, spender)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.BalanceOf(&_BurnMintERC677Helper.CallOpts, account)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.BalanceOf(&_BurnMintERC677Helper.CallOpts, account)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Decimals() (uint8, error) {
	return _BurnMintERC677Helper.Contract.Decimals(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Decimals() (uint8, error) {
	return _BurnMintERC677Helper.Contract.Decimals(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) GetBurners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "getBurners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GetBurners() ([]common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetBurners(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) GetBurners() ([]common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetBurners(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetCCIPAdmin(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetCCIPAdmin(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) GetMinters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "getMinters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GetMinters() ([]common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetMinters(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) GetMinters() ([]common.Address, error) {
	return _BurnMintERC677Helper.Contract.GetMinters(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) IsBurner(opts *bind.CallOpts, burner common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "isBurner", burner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) IsBurner(burner common.Address) (bool, error) {
	return _BurnMintERC677Helper.Contract.IsBurner(&_BurnMintERC677Helper.CallOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) IsBurner(burner common.Address) (bool, error) {
	return _BurnMintERC677Helper.Contract.IsBurner(&_BurnMintERC677Helper.CallOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) IsMinter(opts *bind.CallOpts, minter common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "isMinter", minter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) IsMinter(minter common.Address) (bool, error) {
	return _BurnMintERC677Helper.Contract.IsMinter(&_BurnMintERC677Helper.CallOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) IsMinter(minter common.Address) (bool, error) {
	return _BurnMintERC677Helper.Contract.IsMinter(&_BurnMintERC677Helper.CallOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.MaxSupply(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.MaxSupply(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Name() (string, error) {
	return _BurnMintERC677Helper.Contract.Name(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Name() (string, error) {
	return _BurnMintERC677Helper.Contract.Name(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Owner() (common.Address, error) {
	return _BurnMintERC677Helper.Contract.Owner(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Owner() (common.Address, error) {
	return _BurnMintERC677Helper.Contract.Owner(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC677Helper.Contract.SupportsInterface(&_BurnMintERC677Helper.CallOpts, interfaceId)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC677Helper.Contract.SupportsInterface(&_BurnMintERC677Helper.CallOpts, interfaceId)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Symbol() (string, error) {
	return _BurnMintERC677Helper.Contract.Symbol(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) Symbol() (string, error) {
	return _BurnMintERC677Helper.Contract.Symbol(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC677Helper.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.TotalSupply(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC677Helper.Contract.TotalSupply(&_BurnMintERC677Helper.CallOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "acceptOwnership")
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.AcceptOwnership(&_BurnMintERC677Helper.TransactOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.AcceptOwnership(&_BurnMintERC677Helper.TransactOpts)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "approve", spender, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Approve(&_BurnMintERC677Helper.TransactOpts, spender, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Approve(&_BurnMintERC677Helper.TransactOpts, spender, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Burn(&_BurnMintERC677Helper.TransactOpts, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Burn(&_BurnMintERC677Helper.TransactOpts, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Burn0(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Burn0(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.BurnFrom(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.BurnFrom(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.DecreaseAllowance(&_BurnMintERC677Helper.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.DecreaseAllowance(&_BurnMintERC677Helper.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.DecreaseApproval(&_BurnMintERC677Helper.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.DecreaseApproval(&_BurnMintERC677Helper.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Drip(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "drip", to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Drip(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Drip(&_BurnMintERC677Helper.TransactOpts, to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Drip(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Drip(&_BurnMintERC677Helper.TransactOpts, to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) GrantBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "grantBurnRole", burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantBurnRole(&_BurnMintERC677Helper.TransactOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) GrantBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantBurnRole(&_BurnMintERC677Helper.TransactOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantMintAndBurnRoles(&_BurnMintERC677Helper.TransactOpts, burnAndMinter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantMintAndBurnRoles(&_BurnMintERC677Helper.TransactOpts, burnAndMinter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) GrantMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "grantMintRole", minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantMintRole(&_BurnMintERC677Helper.TransactOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) GrantMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.GrantMintRole(&_BurnMintERC677Helper.TransactOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.IncreaseAllowance(&_BurnMintERC677Helper.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.IncreaseAllowance(&_BurnMintERC677Helper.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.IncreaseApproval(&_BurnMintERC677Helper.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.IncreaseApproval(&_BurnMintERC677Helper.TransactOpts, spender, addedValue)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Mint(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Mint(&_BurnMintERC677Helper.TransactOpts, account, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) RevokeBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "revokeBurnRole", burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.RevokeBurnRole(&_BurnMintERC677Helper.TransactOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) RevokeBurnRole(burner common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.RevokeBurnRole(&_BurnMintERC677Helper.TransactOpts, burner)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) RevokeMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "revokeMintRole", minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.RevokeMintRole(&_BurnMintERC677Helper.TransactOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) RevokeMintRole(minter common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.RevokeMintRole(&_BurnMintERC677Helper.TransactOpts, minter)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "transfer", to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Transfer(&_BurnMintERC677Helper.TransactOpts, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.Transfer(&_BurnMintERC677Helper.TransactOpts, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "transferAndCall", to, amount, data)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) TransferAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferAndCall(&_BurnMintERC677Helper.TransactOpts, to, amount, data)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) TransferAndCall(to common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferAndCall(&_BurnMintERC677Helper.TransactOpts, to, amount, data)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferFrom(&_BurnMintERC677Helper.TransactOpts, from, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferFrom(&_BurnMintERC677Helper.TransactOpts, from, to, amount)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.contract.Transact(opts, "transferOwnership", to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferOwnership(&_BurnMintERC677Helper.TransactOpts, to)
}

func (_BurnMintERC677Helper *BurnMintERC677HelperTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC677Helper.Contract.TransferOwnership(&_BurnMintERC677Helper.TransactOpts, to)
}

type BurnMintERC677HelperApprovalIterator struct {
	Event *BurnMintERC677HelperApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperApproval)
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
		it.Event = new(BurnMintERC677HelperApproval)
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

func (it *BurnMintERC677HelperApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC677HelperApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperApprovalIterator{contract: _BurnMintERC677Helper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperApproval)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseApproval(log types.Log) (*BurnMintERC677HelperApproval, error) {
	event := new(BurnMintERC677HelperApproval)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperBurnAccessGrantedIterator struct {
	Event *BurnMintERC677HelperBurnAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperBurnAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperBurnAccessGranted)
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
		it.Event = new(BurnMintERC677HelperBurnAccessGranted)
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

func (it *BurnMintERC677HelperBurnAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperBurnAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperBurnAccessGranted struct {
	Burner common.Address
	Raw    types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterBurnAccessGranted(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677HelperBurnAccessGrantedIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "BurnAccessGranted", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperBurnAccessGrantedIterator{contract: _BurnMintERC677Helper.contract, event: "BurnAccessGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperBurnAccessGranted, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "BurnAccessGranted", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperBurnAccessGranted)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseBurnAccessGranted(log types.Log) (*BurnMintERC677HelperBurnAccessGranted, error) {
	event := new(BurnMintERC677HelperBurnAccessGranted)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "BurnAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperBurnAccessRevokedIterator struct {
	Event *BurnMintERC677HelperBurnAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperBurnAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperBurnAccessRevoked)
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
		it.Event = new(BurnMintERC677HelperBurnAccessRevoked)
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

func (it *BurnMintERC677HelperBurnAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperBurnAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperBurnAccessRevoked struct {
	Burner common.Address
	Raw    types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterBurnAccessRevoked(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677HelperBurnAccessRevokedIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "BurnAccessRevoked", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperBurnAccessRevokedIterator{contract: _BurnMintERC677Helper.contract, event: "BurnAccessRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperBurnAccessRevoked, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "BurnAccessRevoked", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperBurnAccessRevoked)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseBurnAccessRevoked(log types.Log) (*BurnMintERC677HelperBurnAccessRevoked, error) {
	event := new(BurnMintERC677HelperBurnAccessRevoked)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "BurnAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperMintAccessGrantedIterator struct {
	Event *BurnMintERC677HelperMintAccessGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperMintAccessGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperMintAccessGranted)
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
		it.Event = new(BurnMintERC677HelperMintAccessGranted)
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

func (it *BurnMintERC677HelperMintAccessGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperMintAccessGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperMintAccessGranted struct {
	Minter common.Address
	Raw    types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterMintAccessGranted(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677HelperMintAccessGrantedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "MintAccessGranted", minterRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperMintAccessGrantedIterator{contract: _BurnMintERC677Helper.contract, event: "MintAccessGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperMintAccessGranted, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "MintAccessGranted", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperMintAccessGranted)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseMintAccessGranted(log types.Log) (*BurnMintERC677HelperMintAccessGranted, error) {
	event := new(BurnMintERC677HelperMintAccessGranted)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "MintAccessGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperMintAccessRevokedIterator struct {
	Event *BurnMintERC677HelperMintAccessRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperMintAccessRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperMintAccessRevoked)
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
		it.Event = new(BurnMintERC677HelperMintAccessRevoked)
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

func (it *BurnMintERC677HelperMintAccessRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperMintAccessRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperMintAccessRevoked struct {
	Minter common.Address
	Raw    types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterMintAccessRevoked(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677HelperMintAccessRevokedIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "MintAccessRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperMintAccessRevokedIterator{contract: _BurnMintERC677Helper.contract, event: "MintAccessRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperMintAccessRevoked, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "MintAccessRevoked", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperMintAccessRevoked)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseMintAccessRevoked(log types.Log) (*BurnMintERC677HelperMintAccessRevoked, error) {
	event := new(BurnMintERC677HelperMintAccessRevoked)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "MintAccessRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperOwnershipTransferRequestedIterator struct {
	Event *BurnMintERC677HelperOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperOwnershipTransferRequested)
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
		it.Event = new(BurnMintERC677HelperOwnershipTransferRequested)
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

func (it *BurnMintERC677HelperOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperOwnershipTransferRequestedIterator{contract: _BurnMintERC677Helper.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperOwnershipTransferRequested)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseOwnershipTransferRequested(log types.Log) (*BurnMintERC677HelperOwnershipTransferRequested, error) {
	event := new(BurnMintERC677HelperOwnershipTransferRequested)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperOwnershipTransferredIterator struct {
	Event *BurnMintERC677HelperOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperOwnershipTransferred)
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
		it.Event = new(BurnMintERC677HelperOwnershipTransferred)
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

func (it *BurnMintERC677HelperOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperOwnershipTransferredIterator{contract: _BurnMintERC677Helper.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperOwnershipTransferred)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseOwnershipTransferred(log types.Log) (*BurnMintERC677HelperOwnershipTransferred, error) {
	event := new(BurnMintERC677HelperOwnershipTransferred)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperTransferIterator struct {
	Event *BurnMintERC677HelperTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperTransfer)
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
		it.Event = new(BurnMintERC677HelperTransfer)
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

func (it *BurnMintERC677HelperTransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperTransferIterator{contract: _BurnMintERC677Helper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperTransfer)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseTransfer(log types.Log) (*BurnMintERC677HelperTransfer, error) {
	event := new(BurnMintERC677HelperTransfer)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC677HelperTransfer0Iterator struct {
	Event *BurnMintERC677HelperTransfer0

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC677HelperTransfer0Iterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC677HelperTransfer0)
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
		it.Event = new(BurnMintERC677HelperTransfer0)
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

func (it *BurnMintERC677HelperTransfer0Iterator) Error() error {
	return it.fail
}

func (it *BurnMintERC677HelperTransfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC677HelperTransfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperTransfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC677HelperTransfer0Iterator{contract: _BurnMintERC677Helper.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperTransfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC677Helper.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC677HelperTransfer0)
				if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Transfer0", log); err != nil {
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

func (_BurnMintERC677Helper *BurnMintERC677HelperFilterer) ParseTransfer0(log types.Log) (*BurnMintERC677HelperTransfer0, error) {
	event := new(BurnMintERC677HelperTransfer0)
	if err := _BurnMintERC677Helper.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_BurnMintERC677Helper *BurnMintERC677Helper) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC677Helper.abi.Events["Approval"].ID:
		return _BurnMintERC677Helper.ParseApproval(log)
	case _BurnMintERC677Helper.abi.Events["BurnAccessGranted"].ID:
		return _BurnMintERC677Helper.ParseBurnAccessGranted(log)
	case _BurnMintERC677Helper.abi.Events["BurnAccessRevoked"].ID:
		return _BurnMintERC677Helper.ParseBurnAccessRevoked(log)
	case _BurnMintERC677Helper.abi.Events["MintAccessGranted"].ID:
		return _BurnMintERC677Helper.ParseMintAccessGranted(log)
	case _BurnMintERC677Helper.abi.Events["MintAccessRevoked"].ID:
		return _BurnMintERC677Helper.ParseMintAccessRevoked(log)
	case _BurnMintERC677Helper.abi.Events["OwnershipTransferRequested"].ID:
		return _BurnMintERC677Helper.ParseOwnershipTransferRequested(log)
	case _BurnMintERC677Helper.abi.Events["OwnershipTransferred"].ID:
		return _BurnMintERC677Helper.ParseOwnershipTransferred(log)
	case _BurnMintERC677Helper.abi.Events["Transfer"].ID:
		return _BurnMintERC677Helper.ParseTransfer(log)
	case _BurnMintERC677Helper.abi.Events["Transfer0"].ID:
		return _BurnMintERC677Helper.ParseTransfer0(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (BurnMintERC677HelperApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC677HelperBurnAccessGranted) Topic() common.Hash {
	return common.HexToHash("0x92308bb7573b2a3d17ddb868b39d8ebec433f3194421abc22d084f89658c9bad")
}

func (BurnMintERC677HelperBurnAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0x0a675452746933cefe3d74182e78db7afe57ba60eaa4234b5d85e9aa41b0610c")
}

func (BurnMintERC677HelperMintAccessGranted) Topic() common.Hash {
	return common.HexToHash("0xe46fef8bbff1389d9010703cf8ebb363fb3daf5bf56edc27080b67bc8d9251ea")
}

func (BurnMintERC677HelperMintAccessRevoked) Topic() common.Hash {
	return common.HexToHash("0xed998b960f6340d045f620c119730f7aa7995e7425c2401d3a5b64ff998a59e9")
}

func (BurnMintERC677HelperOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (BurnMintERC677HelperOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (BurnMintERC677HelperTransfer) Topic() common.Hash {
	return common.HexToHash("0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16")
}

func (BurnMintERC677HelperTransfer0) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_BurnMintERC677Helper *BurnMintERC677Helper) Address() common.Address {
	return _BurnMintERC677Helper.address
}

type BurnMintERC677HelperInterface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetBurners(opts *bind.CallOpts) ([]common.Address, error)

	GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error)

	GetMinters(opts *bind.CallOpts) ([]common.Address, error)

	IsBurner(opts *bind.CallOpts, burner common.Address) (bool, error)

	IsMinter(opts *bind.CallOpts, minter common.Address) (bool, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	Drip(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	GrantBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error)

	GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error)

	GrantMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error)

	IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	RevokeBurnRole(opts *bind.TransactOpts, burner common.Address) (*types.Transaction, error)

	RevokeMintRole(opts *bind.TransactOpts, minter common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferAndCall(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC677HelperApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC677HelperApproval, error)

	FilterBurnAccessGranted(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677HelperBurnAccessGrantedIterator, error)

	WatchBurnAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperBurnAccessGranted, burner []common.Address) (event.Subscription, error)

	ParseBurnAccessGranted(log types.Log) (*BurnMintERC677HelperBurnAccessGranted, error)

	FilterBurnAccessRevoked(opts *bind.FilterOpts, burner []common.Address) (*BurnMintERC677HelperBurnAccessRevokedIterator, error)

	WatchBurnAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperBurnAccessRevoked, burner []common.Address) (event.Subscription, error)

	ParseBurnAccessRevoked(log types.Log) (*BurnMintERC677HelperBurnAccessRevoked, error)

	FilterMintAccessGranted(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677HelperMintAccessGrantedIterator, error)

	WatchMintAccessGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperMintAccessGranted, minter []common.Address) (event.Subscription, error)

	ParseMintAccessGranted(log types.Log) (*BurnMintERC677HelperMintAccessGranted, error)

	FilterMintAccessRevoked(opts *bind.FilterOpts, minter []common.Address) (*BurnMintERC677HelperMintAccessRevokedIterator, error)

	WatchMintAccessRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperMintAccessRevoked, minter []common.Address) (event.Subscription, error)

	ParseMintAccessRevoked(log types.Log) (*BurnMintERC677HelperMintAccessRevoked, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*BurnMintERC677HelperOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*BurnMintERC677HelperOwnershipTransferred, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC677HelperTransfer, error)

	FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC677HelperTransfer0Iterator, error)

	WatchTransfer0(opts *bind.WatchOpts, sink chan<- *BurnMintERC677HelperTransfer0, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer0(log types.Log) (*BurnMintERC677HelperTransfer0, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
