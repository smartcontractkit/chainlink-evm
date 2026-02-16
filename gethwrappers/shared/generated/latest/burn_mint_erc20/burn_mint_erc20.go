// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc20

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

var BurnMintERC20MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"maxSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"preMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60c060405234801561001057600080fd5b50604051611bb3380380611bb383398101604081905261002f916103d9565b858585858585858560036100438382610517565b5060046100508282610517565b50505060ff841660805260a083905260006001600160a01b038216156100765781610078565b335b600680546001600160a01b0319166001600160a01b038316179055905082156100a5576100a581846100c3565b6100b0600082610102565b50505050505050505050505050506105f6565b6001600160a01b0382166100f25760405163ec442f0560e01b8152600060048201526024015b60405180910390fd5b6100fe600083836101b2565b5050565b60008281526005602090815260408083206001600160a01b038516845290915281205460ff166101a85760008381526005602090815260408083206001600160a01b03861684529091529020805460ff191660011790556101603390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101ac565b5060005b92915050565b306001600160a01b038316036101e657604051630bc2c5df60e11b81526001600160a01b03831660048201526024016100e9565b6101f18383836101f6565b505050565b6001600160a01b03831661022157806002600082825461021691906105d5565b909155506102939050565b6001600160a01b038316600090815260208190526040902054818110156102745760405163391434e360e21b81526001600160a01b038516600482015260248101829052604481018390526064016100e9565b6001600160a01b03841660009081526020819052604090209082900390555b6001600160a01b0382166102af576002805482900390556102ce565b6001600160a01b03821660009081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161031391815260200190565b60405180910390a3505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261034757600080fd5b81516001600160401b0381111561036057610360610320565b604051601f8201601f19908116603f011681016001600160401b038111828210171561038e5761038e610320565b6040528181528382016020018510156103a657600080fd5b60005b828110156103c5576020818601810151838301820152016103a9565b506000918101602001919091529392505050565b60008060008060008060c087890312156103f257600080fd5b86516001600160401b0381111561040857600080fd5b61041489828a01610336565b602089015190975090506001600160401b0381111561043257600080fd5b61043e89828a01610336565b955050604087015160ff8116811461045557600080fd5b6060880151608089015160a08a015192965090945092506001600160a01b038116811461048157600080fd5b809150509295509295509295565b600181811c908216806104a357607f821691505b6020821081036104c357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101f157806000526020600020601f840160051c810160208510156104f05750805b601f840160051c820191505b8181101561051057600081556001016104fc565b5050505050565b81516001600160401b0381111561053057610530610320565b6105448161053e845461048f565b846104c9565b6020601f82116001811461057857600083156105605750848201515b600019600385901b1c1916600184901b178455610510565b600084815260208120601f198516915b828110156105a85787850151825560209485019460019092019101610588565b50848210156105c65786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156101ac57634e487b7160e01b600052601160045260246000fd5b60805160a05161158a610629600039600081816104840152818161073a0152610764015260006102d2015261158a6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806370a08231116100f9578063a8fa343c11610097578063d539139311610071578063d539139314610448578063d547741f1461046f578063d5abeb0114610482578063dd62ed3e146104a857600080fd5b8063a8fa343c1461040f578063a9059cbb14610422578063c630948d1461043557600080fd5b806391d14854116100d357806391d14854146103a657806395d89b41146103ec5780639dc29fac146103f4578063a217fddf1461040757600080fd5b806370a082311461033557806379cc67901461036b5780638fd6a6ac1461037e57600080fd5b8063248a9ca311610166578063313ce56711610140578063313ce567146102cb57806336568abe146102fc57806340c10f191461030f57806342966c681461032257600080fd5b8063248a9ca31461026c578063282c51f31461028f5780632f2ff15d146102b657600080fd5b806318160ddd1161019757806318160ddd1461020e578063181f5a771461022057806323b872dd1461025957600080fd5b806301ffc9a7146101be57806306fdde03146101e6578063095ea7b3146101fb575b600080fd5b6101d16101cc366004611321565b6104ee565b60405190151581526020015b60405180910390f35b6101ee61054b565b6040516101dd919061136a565b6101d16102093660046113ff565b6105dd565b6002545b6040519081526020016101dd565b60408051808201909152600d81527f43435420322e302e302d6465760000000000000000000000000000000000000060208201526101ee565b6101d1610267366004611429565b6105f5565b61021261027a366004611466565b60009081526005602052604090206001015490565b6102127f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b6102c96102c436600461147f565b610619565b005b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101dd565b6102c961030a36600461147f565b610644565b6102c961031d3660046113ff565b6106a2565b6102c9610330366004611466565b6107f1565b6102126103433660046114ab565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6102c96103793660046113ff565b610829565b60065460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101dd565b6101d16103b436600461147f565b600091825260056020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b6101ee610868565b6102c96104023660046113ff565b610877565b610212600081565b6102c961041d3660046114ab565b610881565b6101d16104303660046113ff565b610904565b6102c96104433660046114ab565b610912565b6102127f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6102c961047d36600461147f565b610969565b7f0000000000000000000000000000000000000000000000000000000000000000610212565b6102126104b63660046114c6565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60006104f98261098e565b8061054557507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b92915050565b60606003805461055a906114f0565b80601f0160208091040260200160405190810160405280929190818152602001828054610586906114f0565b80156105d35780601f106105a8576101008083540402835291602001916105d3565b820191906000526020600020905b8154815290600101906020018083116105b657829003601f168201915b5050505050905090565b6000336105eb818585610abe565b5060019392505050565b600033610603858285610acb565b61060e858585610b95565b506001949350505050565b60008281526005602052604090206001015461063481610c40565b61063e8383610c4a565b50505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610693576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61069d8282610d4a565b505050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66106cc81610c40565b3073ffffffffffffffffffffffffffffffffffffffff841603610738576040517f17858bbe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024015b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000001580159061079957507f00000000000000000000000000000000000000000000000000000000000000008261078d60025490565b6107979190611543565b115b156107e757816107a860025490565b6107b29190611543565b6040517fcbbf111300000000000000000000000000000000000000000000000000000000815260040161072f91815260200190565b61069d8383610e09565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84861081b81610c40565b6108253383610e65565b5050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84861085381610c40565b61085e833384610acb565b61069d8383610e65565b60606004805461055a906114f0565b6108258282610829565b600061088c81610c40565b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d724290600090a3505050565b6000336105eb818585610b95565b61093c7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a682610619565b6109667f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84882610619565b50565b60008281526005602052604090206001015461098481610c40565b61063e8383610d4a565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b07000000000000000000000000000000000000000000000000000000001480610a2157507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b80610a6d57507fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000145b8061054557507fffffffff0000000000000000000000000000000000000000000000000000000082167f8fd6a6ac000000000000000000000000000000000000000000000000000000001492915050565b61069d8383836001610ec1565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81101561063e5781811015610b86576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602481018290526044810183905260640161072f565b61063e84848484036000610ec1565b73ffffffffffffffffffffffffffffffffffffffff8316610be5576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526000600482015260240161072f565b73ffffffffffffffffffffffffffffffffffffffff8216610c35576040517fec442f050000000000000000000000000000000000000000000000000000000081526000600482015260240161072f565b61069d838383610f34565b6109668133610fa6565b600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff16610d4257600083815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff86168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610ce03390565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610545565b506000610545565b600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915281205460ff1615610d4257600083815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610545565b73ffffffffffffffffffffffffffffffffffffffff8216610e59576040517fec442f050000000000000000000000000000000000000000000000000000000081526000600482015260240161072f565b61082560008383610f34565b73ffffffffffffffffffffffffffffffffffffffff8216610eb5576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526000600482015260240161072f565b61082582600083610f34565b3073ffffffffffffffffffffffffffffffffffffffff841603610f28576040517f17858bbe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161072f565b61063e8484848461102e565b3073ffffffffffffffffffffffffffffffffffffffff831603610f9b576040517f17858bbe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161072f565b61069d838383611176565b600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff16610825576040517fe2517d3f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024810183905260440161072f565b73ffffffffffffffffffffffffffffffffffffffff841661107e576040517fe602df050000000000000000000000000000000000000000000000000000000081526000600482015260240161072f565b73ffffffffffffffffffffffffffffffffffffffff83166110ce576040517f94280d620000000000000000000000000000000000000000000000000000000081526000600482015260240161072f565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561063e578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161116891815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff83166111ae5780600260008282546111a39190611543565b909155506112609050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015611234576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602481018290526044810183905260640161072f565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216611289576002805482900390556112b5565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161131491815260200190565b60405180910390a3505050565b60006020828403121561133357600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461136357600080fd5b9392505050565b602081526000825180602084015260005b81811015611398576020818601810151604086840101520161137b565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146113fa57600080fd5b919050565b6000806040838503121561141257600080fd5b61141b836113d6565b946020939093013593505050565b60008060006060848603121561143e57600080fd5b611447846113d6565b9250611455602085016113d6565b929592945050506040919091013590565b60006020828403121561147857600080fd5b5035919050565b6000806040838503121561149257600080fd5b823591506114a2602084016113d6565b90509250929050565b6000602082840312156114bd57600080fd5b611363826113d6565b600080604083850312156114d957600080fd5b6114e2836113d6565b91506114a2602084016113d6565b600181811c9082168061150457607f821691505b60208210810361153d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b80820180821115610545577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea164736f6c634300081a000a",
}

var BurnMintERC20ABI = BurnMintERC20MetaData.ABI

var BurnMintERC20Bin = BurnMintERC20MetaData.Bin

func DeployBurnMintERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals_ uint8, maxSupply_ *big.Int, preMint *big.Int, newOwner common.Address) (common.Address, *types.Transaction, *BurnMintERC20, error) {
	parsed, err := BurnMintERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMintERC20Bin), backend, name, symbol, decimals_, maxSupply_, preMint, newOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMintERC20{address: address, abi: *parsed, BurnMintERC20Caller: BurnMintERC20Caller{contract: contract}, BurnMintERC20Transactor: BurnMintERC20Transactor{contract: contract}, BurnMintERC20Filterer: BurnMintERC20Filterer{contract: contract}}, nil
}

type BurnMintERC20 struct {
	address common.Address
	abi     abi.ABI
	BurnMintERC20Caller
	BurnMintERC20Transactor
	BurnMintERC20Filterer
}

type BurnMintERC20Caller struct {
	contract *bind.BoundContract
}

type BurnMintERC20Transactor struct {
	contract *bind.BoundContract
}

type BurnMintERC20Filterer struct {
	contract *bind.BoundContract
}

type BurnMintERC20Session struct {
	Contract     *BurnMintERC20
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type BurnMintERC20CallerSession struct {
	Contract *BurnMintERC20Caller
	CallOpts bind.CallOpts
}

type BurnMintERC20TransactorSession struct {
	Contract     *BurnMintERC20Transactor
	TransactOpts bind.TransactOpts
}

type BurnMintERC20Raw struct {
	Contract *BurnMintERC20
}

type BurnMintERC20CallerRaw struct {
	Contract *BurnMintERC20Caller
}

type BurnMintERC20TransactorRaw struct {
	Contract *BurnMintERC20Transactor
}

func NewBurnMintERC20(address common.Address, backend bind.ContractBackend) (*BurnMintERC20, error) {
	abi, err := abi.JSON(strings.NewReader(BurnMintERC20ABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindBurnMintERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20{address: address, abi: abi, BurnMintERC20Caller: BurnMintERC20Caller{contract: contract}, BurnMintERC20Transactor: BurnMintERC20Transactor{contract: contract}, BurnMintERC20Filterer: BurnMintERC20Filterer{contract: contract}}, nil
}

func NewBurnMintERC20Caller(address common.Address, caller bind.ContractCaller) (*BurnMintERC20Caller, error) {
	contract, err := bindBurnMintERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20Caller{contract: contract}, nil
}

func NewBurnMintERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*BurnMintERC20Transactor, error) {
	contract, err := bindBurnMintERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20Transactor{contract: contract}, nil
}

func NewBurnMintERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*BurnMintERC20Filterer, error) {
	contract, err := bindBurnMintERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20Filterer{contract: contract}, nil
}

func bindBurnMintERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMintERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_BurnMintERC20 *BurnMintERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20.Contract.BurnMintERC20Caller.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20 *BurnMintERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.BurnMintERC20Transactor.contract.Transfer(opts)
}

func (_BurnMintERC20 *BurnMintERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.BurnMintERC20Transactor.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20 *BurnMintERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMintERC20.Contract.contract.Call(opts, result, method, params...)
}

func (_BurnMintERC20 *BurnMintERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.contract.Transfer(opts)
}

func (_BurnMintERC20 *BurnMintERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.contract.Transact(opts, method, params...)
}

func (_BurnMintERC20 *BurnMintERC20Caller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20.Contract.BURNERROLE(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) BURNERROLE() ([32]byte, error) {
	return _BurnMintERC20.Contract.BURNERROLE(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20.Contract.DEFAULTADMINROLE(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BurnMintERC20.Contract.DEFAULTADMINROLE(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Caller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20.Contract.MINTERROLE(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) MINTERROLE() ([32]byte, error) {
	return _BurnMintERC20.Contract.MINTERROLE(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20.Contract.Allowance(&_BurnMintERC20.CallOpts, owner, spender)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BurnMintERC20.Contract.Allowance(&_BurnMintERC20.CallOpts, owner, spender)
}

func (_BurnMintERC20 *BurnMintERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20.Contract.BalanceOf(&_BurnMintERC20.CallOpts, account)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BurnMintERC20.Contract.BalanceOf(&_BurnMintERC20.CallOpts, account)
}

func (_BurnMintERC20 *BurnMintERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) Decimals() (uint8, error) {
	return _BurnMintERC20.Contract.Decimals(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) Decimals() (uint8, error) {
	return _BurnMintERC20.Contract.Decimals(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Caller) GetCCIPAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "getCCIPAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20.Contract.GetCCIPAdmin(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) GetCCIPAdmin() (common.Address, error) {
	return _BurnMintERC20.Contract.GetCCIPAdmin(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20.Contract.GetRoleAdmin(&_BurnMintERC20.CallOpts, role)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BurnMintERC20.Contract.GetRoleAdmin(&_BurnMintERC20.CallOpts, role)
}

func (_BurnMintERC20 *BurnMintERC20Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20.Contract.HasRole(&_BurnMintERC20.CallOpts, role, account)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BurnMintERC20.Contract.HasRole(&_BurnMintERC20.CallOpts, role, account)
}

func (_BurnMintERC20 *BurnMintERC20Caller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20.Contract.MaxSupply(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) MaxSupply() (*big.Int, error) {
	return _BurnMintERC20.Contract.MaxSupply(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) Name() (string, error) {
	return _BurnMintERC20.Contract.Name(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) Name() (string, error) {
	return _BurnMintERC20.Contract.Name(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20.Contract.SupportsInterface(&_BurnMintERC20.CallOpts, interfaceId)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BurnMintERC20.Contract.SupportsInterface(&_BurnMintERC20.CallOpts, interfaceId)
}

func (_BurnMintERC20 *BurnMintERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) Symbol() (string, error) {
	return _BurnMintERC20.Contract.Symbol(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) Symbol() (string, error) {
	return _BurnMintERC20.Contract.Symbol(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20.Contract.TotalSupply(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _BurnMintERC20.Contract.TotalSupply(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BurnMintERC20.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_BurnMintERC20 *BurnMintERC20Session) TypeAndVersion() (string, error) {
	return _BurnMintERC20.Contract.TypeAndVersion(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20CallerSession) TypeAndVersion() (string, error) {
	return _BurnMintERC20.Contract.TypeAndVersion(&_BurnMintERC20.CallOpts)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "approve", spender, value)
}

func (_BurnMintERC20 *BurnMintERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Approve(&_BurnMintERC20.TransactOpts, spender, value)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Approve(&_BurnMintERC20.TransactOpts, spender, value)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "burn", amount)
}

func (_BurnMintERC20 *BurnMintERC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Burn(&_BurnMintERC20.TransactOpts, amount)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Burn(&_BurnMintERC20.TransactOpts, amount)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "burn0", account, amount)
}

func (_BurnMintERC20 *BurnMintERC20Session) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Burn0(&_BurnMintERC20.TransactOpts, account, amount)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) Burn0(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Burn0(&_BurnMintERC20.TransactOpts, account, amount)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "burnFrom", account, amount)
}

func (_BurnMintERC20 *BurnMintERC20Session) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.BurnFrom(&_BurnMintERC20.TransactOpts, account, amount)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.BurnFrom(&_BurnMintERC20.TransactOpts, account, amount)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (_BurnMintERC20 *BurnMintERC20Session) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.GrantMintAndBurnRoles(&_BurnMintERC20.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) GrantMintAndBurnRoles(burnAndMinter common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.GrantMintAndBurnRoles(&_BurnMintERC20.TransactOpts, burnAndMinter)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "grantRole", role, account)
}

func (_BurnMintERC20 *BurnMintERC20Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.GrantRole(&_BurnMintERC20.TransactOpts, role, account)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.GrantRole(&_BurnMintERC20.TransactOpts, role, account)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "mint", account, amount)
}

func (_BurnMintERC20 *BurnMintERC20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Mint(&_BurnMintERC20.TransactOpts, account, amount)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Mint(&_BurnMintERC20.TransactOpts, account, amount)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

func (_BurnMintERC20 *BurnMintERC20Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.RenounceRole(&_BurnMintERC20.TransactOpts, role, callerConfirmation)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.RenounceRole(&_BurnMintERC20.TransactOpts, role, callerConfirmation)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "revokeRole", role, account)
}

func (_BurnMintERC20 *BurnMintERC20Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.RevokeRole(&_BurnMintERC20.TransactOpts, role, account)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.RevokeRole(&_BurnMintERC20.TransactOpts, role, account)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "setCCIPAdmin", newAdmin)
}

func (_BurnMintERC20 *BurnMintERC20Session) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.SetCCIPAdmin(&_BurnMintERC20.TransactOpts, newAdmin)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) SetCCIPAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.SetCCIPAdmin(&_BurnMintERC20.TransactOpts, newAdmin)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "transfer", to, value)
}

func (_BurnMintERC20 *BurnMintERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Transfer(&_BurnMintERC20.TransactOpts, to, value)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.Transfer(&_BurnMintERC20.TransactOpts, to, value)
}

func (_BurnMintERC20 *BurnMintERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

func (_BurnMintERC20 *BurnMintERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.TransferFrom(&_BurnMintERC20.TransactOpts, from, to, value)
}

func (_BurnMintERC20 *BurnMintERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20.Contract.TransferFrom(&_BurnMintERC20.TransactOpts, from, to, value)
}

type BurnMintERC20ApprovalIterator struct {
	Event *BurnMintERC20Approval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20ApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20Approval)
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
		it.Event = new(BurnMintERC20Approval)
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

func (it *BurnMintERC20ApprovalIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_BurnMintERC20 *BurnMintERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20ApprovalIterator{contract: _BurnMintERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20 *BurnMintERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnMintERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20Approval)
				if err := _BurnMintERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_BurnMintERC20 *BurnMintERC20Filterer) ParseApproval(log types.Log) (*BurnMintERC20Approval, error) {
	event := new(BurnMintERC20Approval)
	if err := _BurnMintERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20CCIPAdminTransferredIterator struct {
	Event *BurnMintERC20CCIPAdminTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20CCIPAdminTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20CCIPAdminTransferred)
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
		it.Event = new(BurnMintERC20CCIPAdminTransferred)
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

func (it *BurnMintERC20CCIPAdminTransferredIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20CCIPAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20CCIPAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log
}

func (_BurnMintERC20 *BurnMintERC20Filterer) FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20CCIPAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20.contract.FilterLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20CCIPAdminTransferredIterator{contract: _BurnMintERC20.contract, event: "CCIPAdminTransferred", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20 *BurnMintERC20Filterer) WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20CCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _BurnMintERC20.contract.WatchLogs(opts, "CCIPAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20CCIPAdminTransferred)
				if err := _BurnMintERC20.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
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

func (_BurnMintERC20 *BurnMintERC20Filterer) ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20CCIPAdminTransferred, error) {
	event := new(BurnMintERC20CCIPAdminTransferred)
	if err := _BurnMintERC20.contract.UnpackLog(event, "CCIPAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20RoleAdminChangedIterator struct {
	Event *BurnMintERC20RoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20RoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20RoleAdminChanged)
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
		it.Event = new(BurnMintERC20RoleAdminChanged)
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

func (it *BurnMintERC20RoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_BurnMintERC20 *BurnMintERC20Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20RoleAdminChangedIterator, error) {

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

	logs, sub, err := _BurnMintERC20.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20RoleAdminChangedIterator{contract: _BurnMintERC20.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20 *BurnMintERC20Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20RoleAdminChanged)
				if err := _BurnMintERC20.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

func (_BurnMintERC20 *BurnMintERC20Filterer) ParseRoleAdminChanged(log types.Log) (*BurnMintERC20RoleAdminChanged, error) {
	event := new(BurnMintERC20RoleAdminChanged)
	if err := _BurnMintERC20.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20RoleGrantedIterator struct {
	Event *BurnMintERC20RoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20RoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20RoleGranted)
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
		it.Event = new(BurnMintERC20RoleGranted)
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

func (it *BurnMintERC20RoleGrantedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20 *BurnMintERC20Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20RoleGrantedIterator, error) {

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

	logs, sub, err := _BurnMintERC20.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20RoleGrantedIterator{contract: _BurnMintERC20.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20 *BurnMintERC20Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20RoleGranted)
				if err := _BurnMintERC20.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

func (_BurnMintERC20 *BurnMintERC20Filterer) ParseRoleGranted(log types.Log) (*BurnMintERC20RoleGranted, error) {
	event := new(BurnMintERC20RoleGranted)
	if err := _BurnMintERC20.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20RoleRevokedIterator struct {
	Event *BurnMintERC20RoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20RoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20RoleRevoked)
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
		it.Event = new(BurnMintERC20RoleRevoked)
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

func (it *BurnMintERC20RoleRevokedIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_BurnMintERC20 *BurnMintERC20Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20RoleRevokedIterator, error) {

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

	logs, sub, err := _BurnMintERC20.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20RoleRevokedIterator{contract: _BurnMintERC20.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20 *BurnMintERC20Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BurnMintERC20.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20RoleRevoked)
				if err := _BurnMintERC20.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

func (_BurnMintERC20 *BurnMintERC20Filterer) ParseRoleRevoked(log types.Log) (*BurnMintERC20RoleRevoked, error) {
	event := new(BurnMintERC20RoleRevoked)
	if err := _BurnMintERC20.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type BurnMintERC20TransferIterator struct {
	Event *BurnMintERC20Transfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *BurnMintERC20TransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnMintERC20Transfer)
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
		it.Event = new(BurnMintERC20Transfer)
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

func (it *BurnMintERC20TransferIterator) Error() error {
	return it.fail
}

func (it *BurnMintERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type BurnMintERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_BurnMintERC20 *BurnMintERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnMintERC20TransferIterator{contract: _BurnMintERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_BurnMintERC20 *BurnMintERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnMintERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(BurnMintERC20Transfer)
				if err := _BurnMintERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_BurnMintERC20 *BurnMintERC20Filterer) ParseTransfer(log types.Log) (*BurnMintERC20Transfer, error) {
	event := new(BurnMintERC20Transfer)
	if err := _BurnMintERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (BurnMintERC20Approval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (BurnMintERC20CCIPAdminTransferred) Topic() common.Hash {
	return common.HexToHash("0x9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d7242")
}

func (BurnMintERC20RoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (BurnMintERC20RoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (BurnMintERC20RoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (BurnMintERC20Transfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (_BurnMintERC20 *BurnMintERC20) Address() common.Address {
	return _BurnMintERC20.address
}

type BurnMintERC20Interface interface {
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

	GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnMintERC20ApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnMintERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*BurnMintERC20Approval, error)

	FilterCCIPAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*BurnMintERC20CCIPAdminTransferredIterator, error)

	WatchCCIPAdminTransferred(opts *bind.WatchOpts, sink chan<- *BurnMintERC20CCIPAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error)

	ParseCCIPAdminTransferred(log types.Log) (*BurnMintERC20CCIPAdminTransferred, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BurnMintERC20RoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BurnMintERC20RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*BurnMintERC20RoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20RoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BurnMintERC20RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*BurnMintERC20RoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BurnMintERC20RoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BurnMintERC20RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*BurnMintERC20RoleRevoked, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnMintERC20TransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnMintERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*BurnMintERC20Transfer, error)

	Address() common.Address
}
