// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc20_with_drip

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

var BurnMintERC20WithDripMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"drip\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051611ba0380380611ba083398101604081905261002f916103e3565b81816012600080808585858380808585600361004b83826104d4565b50600461005882826104d4565b50505060ff841660805260a083905260006001600160a01b0382161561007e5781610080565b335b600680546001600160a01b0319166001600160a01b038316179055905082156100ad576100ad81846100cd565b6100b860008261010c565b505050505050505050505050505050506105b3565b6001600160a01b0382166100fc5760405163ec442f0560e01b8152600060048201526024015b60405180910390fd5b610108600083836101bc565b5050565b60008281526005602090815260408083206001600160a01b038516845290915281205460ff166101b25760008381526005602090815260408083206001600160a01b03861684529091529020805460ff1916600117905561016a3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101b6565b5060005b92915050565b306001600160a01b038316036101f057604051630bc2c5df60e11b81526001600160a01b03831660048201526024016100f3565b6101fb838383610200565b505050565b6001600160a01b03831661022b5780600260008282546102209190610592565b9091555061029d9050565b6001600160a01b0383166000908152602081905260409020548181101561027e5760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016100f3565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b0382166102b9576002805482900390556102d8565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161031d91815260200190565b60405180910390a3505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261035157600080fd5b81516001600160401b0381111561036a5761036a61032a565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103985761039861032a565b6040528181528382016020018510156103b057600080fd5b60005b828110156103cf576020818601810151838301820152016103b3565b506000918101602001919091529392505050565b600080604083850312156103f657600080fd5b82516001600160401b0381111561040c57600080fd5b61041885828601610340565b602085015190935090506001600160401b0381111561043657600080fd5b61044285828601610340565b9150509250929050565b600181811c9082168061046057607f821691505b60208210810361048057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101fb57806000526020600020601f840160051c810160208510156104ad5750805b601f840160051c820191505b818110156104cd57600081556001016104b9565b5050505050565b81516001600160401b038111156104ed576104ed61032a565b610501816104fb845461044c565b84610486565b6020601f821160018114610535576000831561051d5750848201515b600019600385901b1c1916600184901b1784556104cd565b600084815260208120601f198516915b828110156105655787850151825560209485019460019092019101610545565b50848210156105835786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156101b657634e487b7160e01b600052601160045260246000fd5b60805160a0516115ba6105e6600039600081816104a2015281816107580152610782015260006102dd01526115ba6000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c806370a08231116100f9578063a8fa343c11610097578063d539139311610071578063d539139314610466578063d547741f1461048d578063d5abeb01146104a0578063dd62ed3e146104c657600080fd5b8063a8fa343c1461042d578063a9059cbb14610440578063c630948d1461045357600080fd5b806391d14854116100d357806391d14854146103c457806395d89b411461040a5780639dc29fac14610412578063a217fddf1461042557600080fd5b806370a082311461035357806379cc6790146103895780638fd6a6ac1461039c57600080fd5b8063282c51f31161016657806336568abe1161014057806336568abe1461030757806340c10f191461031a57806342966c681461032d57806367a5cd061461034057600080fd5b8063282c51f31461029a5780632f2ff15d146102c1578063313ce567146102d657600080fd5b806318160ddd116101a257806318160ddd14610219578063181f5a771461022b57806323b872dd14610264578063248a9ca31461027757600080fd5b806301ffc9a7146101c957806306fdde03146101f1578063095ea7b314610206575b600080fd5b6101dc6101d7366004611351565b61050c565b60405190151581526020015b60405180910390f35b6101f9610569565b6040516101e8919061139a565b6101dc61021436600461142f565b6105fb565b6002545b6040519081526020016101e8565b60408051808201909152600d81527f43435420322e302e302d6465760000000000000000000000000000000000000060208201526101f9565b6101dc610272366004611459565b610613565b61021d610285366004611496565b60009081526005602052604090206001015490565b61021d7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b6102d46102cf3660046114af565b610637565b005b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101e8565b6102d46103153660046114af565b610662565b6102d461032836600461142f565b6106c0565b6102d461033b366004611496565b61080f565b6102d461034e3660046114db565b610847565b61021d6103613660046114db565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6102d461039736600461142f565b61085c565b60065460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e8565b6101dc6103d23660046114af565b600091825260056020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101f961089b565b6102d461042036600461142f565b6108aa565b61021d600081565b6102d461043b3660046114db565b6108b4565b6101dc61044e36600461142f565b610937565b6102d46104613660046114db565b610945565b61021d7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6102d461049b3660046114af565b610999565b7f000000000000000000000000000000000000000000000000000000000000000061021d565b61021d6104d43660046114f6565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6000610517826109be565b8061056357507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b92915050565b60606003805461057890611520565b80601f01602080910402602001604051908101604052809291908181526020018280546105a490611520565b80156105f15780601f106105c6576101008083540402835291602001916105f1565b820191906000526020600020905b8154815290600101906020018083116105d457829003601f168201915b5050505050905090565b600033610609818585610aee565b5060019392505050565b600033610621858285610afb565b61062c858585610bc5565b506001949350505050565b60008281526005602052604090206001015461065281610c70565b61065c8383610c7a565b50505050565b73ffffffffffffffffffffffffffffffffffffffff811633146106b1576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106bb8282610d7a565b505050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66106ea81610c70565b3073ffffffffffffffffffffffffffffffffffffffff841603610756576040517f17858bbe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b7f0000000000000000000000000000000000000000000000000000000000000000158015906107b757507f0000000000000000000000000000000000000000000000000000000000000000826107ab60025490565b6107b59190611573565b115b1561080557816107c660025490565b6107d09190611573565b6040517fcbbf111300000000000000000000000000000000000000000000000000000000815260040161074d91815260200190565b6106bb8383610e39565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84861083981610c70565b6108433383610e95565b5050565b61085981670de0b6b3a7640000610e39565b50565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84861088681610c70565b610891833384610afb565b6106bb8383610e95565b60606004805461057890611520565b610843828261085c565b60006108bf81610c70565b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d724290600090a3505050565b600033610609818585610bc5565b61096f7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a682610637565b6108597f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84882610637565b6000828152600560205260409020600101546109b481610c70565b61065c8383610d7a565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b07000000000000000000000000000000000000000000000000000000001480610a5157507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b80610a9d57507fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000145b8061056357507fffffffff0000000000000000000000000000000000000000000000000000000082167f8fd6a6ac000000000000000000000000000000000000000000000000000000001492915050565b6106bb8383836001610ef1565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81101561065c5781811015610bb6576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602481018290526044810183905260640161074d565b61065c84848484036000610ef1565b73ffffffffffffffffffffffffffffffffffffffff8316610c15576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526000600482015260240161074d565b73ffffffffffffffffffffffffffffffffffffffff8216610c65576040517fec442f050000000000000000000000000000000000000000000000000000000081526000600482015260240161074d565b6106bb838383610f64565b6108598133610fd6565b600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610d7257600083815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610d103390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610563565b506000610563565b600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610d7257600083815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610563565b73ffffffffffffffffffffffffffffffffffffffff8216610e89576040517fec442f050000000000000000000000000000000000000000000000000000000081526000600482015260240161074d565b61084360008383610f64565b73ffffffffffffffffffffffffffffffffffffffff8216610ee5576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526000600482015260240161074d565b61084382600083610f64565b3073ffffffffffffffffffffffffffffffffffffffff841603610f58576040517f17858bbe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161074d565b61065c8484848461105e565b3073ffffffffffffffffffffffffffffffffffffffff831603610fcb576040517f17858bbe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161074d565b6106bb8383836111a6565b600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610843576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440161074d565b73ffffffffffffffffffffffffffffffffffffffff84166110ae576040517fe602df050000000000000000000000000000000000000000000000000000000081526000600482015260240161074d565b73ffffffffffffffffffffffffffffffffffffffff83166110fe576040517f94280d620000000000000000000000000000000000000000000000000000000081526000600482015260240161074d565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561065c578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161119891815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff83166111de5780600260008282546111d39190611573565b909155506112909050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015611264576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018290526044810183905260640161074d565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff82166112b9576002805482900390556112e5565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161134491815260200190565b60405180910390a3505050565b60006020828403121561136357600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461139357600080fd5b9392505050565b602081526000825180602084015260005b818110156113c857602081860181015160408684010152016113ab565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461142a57600080fd5b919050565b6000806040838503121561144257600080fd5b61144b83611406565b946020939093013593505050565b60008060006060848603121561146e57600080fd5b61147784611406565b925061148560208501611406565b929592945050506040919091013590565b6000602082840312156114a857600080fd5b5035919050565b600080604083850312156114c257600080fd5b823591506114d260208401611406565b90509250929050565b6000602082840312156114ed57600080fd5b61139382611406565b6000806040838503121561150957600080fd5b61151283611406565b91506114d260208401611406565b600181811c9082168061153457607f821691505b60208210810361156d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b80820180821115610563577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea164736f6c634300081a000a",
}

var BurnMintERC20WithDripABI = BurnMintERC20WithDripMetaData.ABI

var BurnMintERC20WithDripBin = BurnMintERC20WithDripMetaData.Bin

func DeployBurnMintERC20WithDrip(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *BurnMintERC20WithDrip, error) {
	parsed, err := BurnMintERC20WithDripMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC20WithDripBin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC20WithDrip{address: address, abi: *parsed, BurnMintERC20WithDripCaller: BurnMintERC20WithDripCaller{contract: contract}, BurnMintERC20WithDripTransactor: BurnMintERC20WithDripTransactor{contract: contract}, BurnMintERC20WithDripFilterer: BurnMintERC20WithDripFilterer{contract: contract}}, nil
}

type BurnMintERC20WithDrip struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC20WithDripCaller
	BurnMintERC20WithDripTransactor
	BurnMintERC20WithDripFilterer
}

type BurnMintERC20WithDripCaller struct {
	contract *bind.BoundContract
}

type BurnMintERC20WithDripTransactor struct {
	contract *bind.BoundContract
}

type BurnMintERC20WithDripFilterer struct {
	contract *bind.BoundContract
}

type BurnMintERC20WithDripSession struct {
	Contract     *BurnMintERC20WithDrip
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC20WithDripCallerSession struct {
	Contract *BurnMintERC20WithDripCaller
	CallOpts bind.CallOpts
}

type BurnMintERC20WithDripTransactorSession struct {
	Contract     *BurnMintERC20WithDripTransactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC20WithDripRaw struct {
	Contract *BurnMintERC20WithDrip
}

type BurnMintERC20WithDripCallerRaw struct {
	Contract *BurnMintERC20WithDripCaller
}

type BurnMintERC20WithDripTransactorRaw struct {
	Contract *BurnMintERC20WithDripTransactor
}

func NewBurnMintERC20WithDrip(address common.Address, backend bind.ContractBackend) (*BurnMintERC20WithDrip, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC20WithDripABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC20WithDrip(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDrip{address: address, abi: abi, BurnMintERC20WithDripCaller: BurnMintERC20WithDripCaller{contract: contract}, BurnMintERC20WithDripTransactor: BurnMintERC20WithDripTransactor{contract: contract}, BurnMintERC20WithDripFilterer: BurnMintERC20WithDripFilterer{contract: contract}}, nil
}

func NewBurnMintERC20WithDripCaller(address common.Address, caller bind.ContractCaller) (*BurnMintERC20WithDripCaller, error) {
	contract, err := bindBurnMintERC20WithDrip(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDripCaller{contract: contract}, nil
}

func NewBurnMintERC20WithDripTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC20WithDripTransactor, error) {
	contract, err := bindBurnMintERC20WithDrip(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDripTransactor{contract: contract}, nil
}

func NewBurnMintERC20WithDripFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC20WithDripFilterer, error) {
	contract, err := bindBurnMintERC20WithDrip(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDripFilterer{contract: contract}, nil
}

func bindBurnMintERC20WithDrip(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC20WithDripMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20WithDrip.Contract.BurnMintERC20WithDripCaller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.BurnMintERC20WithDripTransactor.contract.Transfer(opts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.BurnMintERC20WithDripTransactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20WithDrip.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.contract.Transfer(opts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20WithDrip.Contract.BURNERROLE(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20WithDrip.Contract.BURNERROLE(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20WithDrip.Contract.DEFAULTADMINROLE(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20WithDrip.Contract.DEFAULTADMINROLE(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20WithDrip.Contract.MINTERROLE(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20WithDrip.Contract.MINTERROLE(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20WithDrip.Contract.Allowance(&_BurnMintERC20WithDrip.CallOpts, owner, spender)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20WithDrip.Contract.Allowance(&_BurnMintERC20WithDrip.CallOpts, owner, spender)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20WithDrip.Contract.BalanceOf(&_BurnMintERC20WithDrip.CallOpts, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20WithDrip.Contract.BalanceOf(&_BurnMintERC20WithDrip.CallOpts, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Decimals() (uint8, error) {
	return _BurnMintERC20WithDrip.Contract.Decimals(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) Decimals() (uint8, error) {
	return _BurnMintERC20WithDrip.Contract.Decimals(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20WithDrip.Contract.GetCCIPAdmin(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20WithDrip.Contract.GetCCIPAdmin(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20WithDrip.Contract.GetRoleAdmin(&_BurnMintERC20WithDrip.CallOpts, role)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20WithDrip.Contract.GetRoleAdmin(&_BurnMintERC20WithDrip.CallOpts, role)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20WithDrip.Contract.HasRole(&_BurnMintERC20WithDrip.CallOpts, role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20WithDrip.Contract.HasRole(&_BurnMintERC20WithDrip.CallOpts, role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20WithDrip.Contract.MaxSupply(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20WithDrip.Contract.MaxSupply(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Name() (string, error) {
	return _BurnMintERC20WithDrip.Contract.Name(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) Name() (string, error) {
	return _BurnMintERC20WithDrip.Contract.Name(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20WithDrip.Contract.SupportsInterface(&_BurnMintERC20WithDrip.CallOpts, interfaceId)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20WithDrip.Contract.SupportsInterface(&_BurnMintERC20WithDrip.CallOpts, interfaceId)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Symbol() (string, error) {
	return _BurnMintERC20WithDrip.Contract.Symbol(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) Symbol() (string, error) {
	return _BurnMintERC20WithDrip.Contract.Symbol(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20WithDrip.Contract.TotalSupply(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20WithDrip.Contract.TotalSupply(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20WithDrip.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) TypeAndVersion() (string, error) {
	return _BurnMintERC20WithDrip.Contract.TypeAndVersion(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripCallerSession) TypeAndVersion() (string, error) {
	return _BurnMintERC20WithDrip.Contract.TypeAndVersion(&_BurnMintERC20WithDrip.CallOpts)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "approve", spender, value)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Approve(&_BurnMintERC20WithDrip.TransactOpts, spender, value)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Approve(&_BurnMintERC20WithDrip.TransactOpts, spender, value)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Burn(&_BurnMintERC20WithDrip.TransactOpts, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Burn(&_BurnMintERC20WithDrip.TransactOpts, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Burn0(&_BurnMintERC20WithDrip.TransactOpts, account, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Burn0(&_BurnMintERC20WithDrip.TransactOpts, account, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.BurnFrom(&_BurnMintERC20WithDrip.TransactOpts, account, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.BurnFrom(&_BurnMintERC20WithDrip.TransactOpts, account, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) Drip(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "drip", to)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Drip(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Drip(&_BurnMintERC20WithDrip.TransactOpts, to)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) Drip(to common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Drip(&_BurnMintERC20WithDrip.TransactOpts, to)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.GrantMintAndBurnRoles(&_BurnMintERC20WithDrip.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.GrantMintAndBurnRoles(&_BurnMintERC20WithDrip.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.GrantRole(&_BurnMintERC20WithDrip.TransactOpts, role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.GrantRole(&_BurnMintERC20WithDrip.TransactOpts, role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Mint(&_BurnMintERC20WithDrip.TransactOpts, account, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Mint(&_BurnMintERC20WithDrip.TransactOpts, account, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.RenounceRole(&_BurnMintERC20WithDrip.TransactOpts, role, callerConfirmation)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.RenounceRole(&_BurnMintERC20WithDrip.TransactOpts, role, callerConfirmation)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.RevokeRole(&_BurnMintERC20WithDrip.TransactOpts, role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.RevokeRole(&_BurnMintERC20WithDrip.TransactOpts, role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.SetCCIPAdmin(&_BurnMintERC20WithDrip.TransactOpts, newAdmin)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.SetCCIPAdmin(&_BurnMintERC20WithDrip.TransactOpts, newAdmin)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "transfer", to, value)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Transfer(&_BurnMintERC20WithDrip.TransactOpts, to, value)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Transfer(&_BurnMintERC20WithDrip.TransactOpts, to, value)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.TransferFrom(&_BurnMintERC20WithDrip.TransactOpts, from, to, value)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.TransferFrom(&_BurnMintERC20WithDrip.TransactOpts, from, to, value)
}

type BurnMintERC20WithDripApprovalIterator struct {
	Event *BurnMintERC20WithDripApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20WithDripApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20WithDripApproval)
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
		it.Event = new(BurnMintERC20WithDripApproval)
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

func (it *BurnMintERC20WithDripApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20WithDripApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20WithDripApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20WithDripApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDripApprovalIterator{contract: _BurnMintERC20WithDrip.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20WithDripApproval)
				if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) ParseApproval(log types.Log) (*BurnMintERC20WithDripApproval, error) {
	event := new(BurnMintERC20WithDripApproval)
	if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20WithDripCCIPAdminTransferredIterator struct {
	Event *BurnMintERC20WithDripCCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20WithDripCCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20WithDripCCIPAdminTransferred)
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
		it.Event = new(BurnMintERC20WithDripCCIPAdminTransferred)
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

func (it *BurnMintERC20WithDripCCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20WithDripCCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20WithDripCCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20WithDripCCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDripCCIPAdminTransferredIterator{contract: _BurnMintERC20WithDrip.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20WithDripCCIPAdminTransferred)
				if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20WithDripCCIPAdminTransferred, error) {
	event := new(BurnMintERC20WithDripCCIPAdminTransferred)
	if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20WithDripRoleAdminChangedIterator struct {
	Event *BurnMintERC20WithDripRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20WithDripRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20WithDripRoleAdminChanged)
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
		it.Event = new(BurnMintERC20WithDripRoleAdminChanged)
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

func (it *BurnMintERC20WithDripRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20WithDripRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20WithDripRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20WithDripRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDripRoleAdminChangedIterator{contract: _BurnMintERC20WithDrip.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20WithDripRoleAdminChanged)
				if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) ParseRoleAdminChanged(log types.Log) (*BurnMintERC20WithDripRoleAdminChanged, error) {
	event := new(BurnMintERC20WithDripRoleAdminChanged)
	if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20WithDripRoleGrantedIterator struct {
	Event *BurnMintERC20WithDripRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20WithDripRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20WithDripRoleGranted)
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
		it.Event = new(BurnMintERC20WithDripRoleGranted)
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

func (it *BurnMintERC20WithDripRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20WithDripRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20WithDripRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20WithDripRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDripRoleGrantedIterator{contract: _BurnMintERC20WithDrip.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20WithDripRoleGranted)
				if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) ParseRoleGranted(log types.Log) (*BurnMintERC20WithDripRoleGranted, error) {
	event := new(BurnMintERC20WithDripRoleGranted)
	if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20WithDripRoleRevokedIterator struct {
	Event *BurnMintERC20WithDripRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20WithDripRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20WithDripRoleRevoked)
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
		it.Event = new(BurnMintERC20WithDripRoleRevoked)
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

func (it *BurnMintERC20WithDripRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20WithDripRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20WithDripRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20WithDripRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDripRoleRevokedIterator{contract: _BurnMintERC20WithDrip.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20WithDripRoleRevoked)
				if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) ParseRoleRevoked(log types.Log) (*BurnMintERC20WithDripRoleRevoked, error) {
	event := new(BurnMintERC20WithDripRoleRevoked)
	if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20WithDripTransferIterator struct {
	Event *BurnMintERC20WithDripTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20WithDripTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20WithDripTransfer)
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
		it.Event = new(BurnMintERC20WithDripTransfer)
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

func (it *BurnMintERC20WithDripTransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20WithDripTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20WithDripTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20WithDripTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20WithDripTransferIterator{contract: _BurnMintERC20WithDrip.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20WithDrip.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20WithDripTransfer)
				if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripFilterer) ParseTransfer(log types.Log) (*BurnMintERC20WithDripTransfer, error) {
	event := new(BurnMintERC20WithDripTransfer)
	if err := _BurnMintERC20WithDrip.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (BurnMintERC20WithDripApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC20WithDripCCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (BurnMintERC20WithDripRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnMintERC20WithDripRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnMintERC20WithDripRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (BurnMintERC20WithDripTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDrip) Address() common.Address {
	return _BurnMintERC20WithDrip.address
}

type BurnMintERC20WithDripInterface interface {
	BURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	MINTERROLE(opts *bind.CallOpts) ([32]byte, error)

	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	MaxSupply(opts *bind.CallOpts) (*big.Int, error)

	Name(opts *bind.CallOpts) (string, error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	TypeAndVersion(opts *bind.CallOpts) (string, error)

	Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	Drip(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20WithDripApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC20WithDripApproval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20WithDripCCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripCCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20WithDripCCIPAdminTransferred, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20WithDripRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnMintERC20WithDripRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20WithDripRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnMintERC20WithDripRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20WithDripRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnMintERC20WithDripRoleRevoked, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20WithDripTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20WithDripTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC20WithDripTransfer, error)

	Address() common.Address
}
