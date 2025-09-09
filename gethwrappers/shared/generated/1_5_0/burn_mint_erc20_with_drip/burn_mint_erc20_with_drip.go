// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package burn_mint_erc20_with_drip

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

var BurnMintERC20WithDripMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"drip\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCCIPAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantMintAndBurnRoles\",\"inputs\":[{\"name\":\"burnAndMinter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCCIPAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CCIPAdminTransferred\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidRecipient\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"MaxSupplyExceeded\",\"inputs\":[{\"name\":\"supplyAfterMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200226d3803806200226d8339810160408190526200003491620002f0565b81816012600080848460036200004b8382620003e8565b5060046200005a8282620003e8565b50505060ff831660805260a0829052600680546001600160a01b0319163317905580156200008e576200008e3382620000a8565b6200009b6000336200016f565b50505050505050620004d6565b6001600160a01b038216620001035760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b8060026000828254620001179190620004b4565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b5050565b6200017b8282620001fe565b6200016b5760008281526005602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620001b53390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b505050565b60008281526005602090815260408083206001600160a01b038516845290915290205460ff165b92915050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200025357600080fd5b81516001600160401b03808211156200027057620002706200022b565b604051601f8301601f19908116603f011681019082821181831017156200029b576200029b6200022b565b81604052838152602092508683858801011115620002b857600080fd5b600091505b83821015620002dc5785820183015181830184015290820190620002bd565b600093810190920192909252949350505050565b600080604083850312156200030457600080fd5b82516001600160401b03808211156200031c57600080fd5b6200032a8683870162000241565b935060208501519150808211156200034157600080fd5b50620003508582860162000241565b9150509250929050565b600181811c908216806200036f57607f821691505b6020821081036200039057634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001f957600081815260208120601f850160051c81016020861015620003bf5750805b601f850160051c820191505b81811015620003e057828155600101620003cb565b505050505050565b81516001600160401b038111156200040457620004046200022b565b6200041c816200041584546200035a565b8462000396565b602080601f8311600181146200045457600084156200043b5750858301515b600019600386901b1c1916600185901b178555620003e0565b600085815260208120601f198616915b82811015620004855788860151825594840194600190910190840162000464565b5085821015620004a45787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200022557634e487b7160e01b600052601160045260246000fd5b60805160a051611d636200050a6000396000818161049a01528181610910015261093a015260006102af0152611d636000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c806370a0823111610104578063a457c2d7116100a2578063d539139311610071578063d53913931461045e578063d547741f14610485578063d5abeb0114610498578063dd62ed3e146104be57600080fd5b8063a457c2d714610412578063a8fa343c14610425578063a9059cbb14610438578063c630948d1461044b57600080fd5b806391d14854116100de57806391d14854146103a957806395d89b41146103ef5780639dc29fac146103f7578063a217fddf1461040a57600080fd5b806370a082311461033857806379cc67901461036e5780638fd6a6ac1461038157600080fd5b80632f2ff15d11610171578063395093511161014b57806339509351146102ec57806340c10f19146102ff57806342966c681461031257806367a5cd061461032557600080fd5b80632f2ff15d14610293578063313ce567146102a857806336568abe146102d957600080fd5b806318160ddd116101ad57806318160ddd1461022457806323b872dd14610236578063248a9ca314610249578063282c51f31461026c57600080fd5b806301ffc9a7146101d457806306fdde03146101fc578063095ea7b314610211575b600080fd5b6101e76101e23660046119c6565b610504565b60405190151581526020015b60405180910390f35b610204610681565b6040516101f39190611a2c565b6101e761021f366004611aa6565b610713565b6002545b6040519081526020016101f3565b6101e7610244366004611ad0565b61072b565b610228610257366004611b0c565b60009081526005602052604090206001015490565b6102287f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b6102a66102a1366004611b25565b61074f565b005b60405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101f3565b6102a66102e7366004611b25565b610779565b6101e76102fa366004611aa6565b610831565b6102a661030d366004611aa6565b61087d565b6102a6610320366004611b0c565b6109c7565b6102a6610333366004611b51565b6109fa565b610228610346366004611b51565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6102a661037c366004611aa6565b610a0f565b60065460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101f3565b6101e76103b7366004611b25565b600091825260056020908152604080842073ffffffffffffffffffffffffffffffffffffffff93909316845291905290205460ff1690565b610204610a43565b6102a6610405366004611aa6565b610a52565b610228600081565b6101e7610420366004611aa6565b610a5c565b6102a6610433366004611b51565b610b2d565b6101e7610446366004611aa6565b610bb0565b6102a6610459366004611b51565b610bbe565b6102287f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6102a6610493366004611b25565b610c12565b7f0000000000000000000000000000000000000000000000000000000000000000610228565b6102286104cc366004611b6c565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f36372b0700000000000000000000000000000000000000000000000000000000148061059757507fffffffff0000000000000000000000000000000000000000000000000000000082167fe6599b4d00000000000000000000000000000000000000000000000000000000145b806105e357507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b8061062f57507fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000145b8061067b57507fffffffff0000000000000000000000000000000000000000000000000000000082167f8fd6a6ac00000000000000000000000000000000000000000000000000000000145b92915050565b60606003805461069090611b96565b80601f01602080910402602001604051908101604052809291908181526020018280546106bc90611b96565b80156107095780601f106106de57610100808354040283529160200191610709565b820191906000526020600020905b8154815290600101906020018083116106ec57829003601f168201915b5050505050905090565b600033610721818585610c37565b5060019392505050565b600033610739858285610ca9565b610744858585610d80565b506001949350505050565b60008281526005602052604090206001015461076a81610df2565b6107748383610dfc565b505050565b73ffffffffffffffffffffffffffffffffffffffff81163314610823576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201527f20726f6c657320666f722073656c66000000000000000000000000000000000060648201526084015b60405180910390fd5b61082d8282610ef0565b5050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906107219082908690610878908790611c18565b610c37565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a66108a781610df2565b3073ffffffffffffffffffffffffffffffffffffffff84160361090e576040517f17858bbe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161081a565b7f00000000000000000000000000000000000000000000000000000000000000001580159061096f57507f00000000000000000000000000000000000000000000000000000000000000008261096360025490565b61096d9190611c18565b115b156109bd578161097e60025490565b6109889190611c18565b6040517fcbbf111300000000000000000000000000000000000000000000000000000000815260040161081a91815260200190565b6107748383610fab565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8486109f181610df2565b61082d8261109e565b610a0c81670de0b6b3a7640000610fab565b50565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a848610a3981610df2565b61077483836110a8565b60606004805461069090611b96565b61082d8282610a0f565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610b20576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f000000000000000000000000000000000000000000000000000000606482015260840161081a565b6107448286868403610c37565b6000610b3881610df2565b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f9524c9e4b0b61eb018dd58a1cd856e3e74009528328ab4a613b434fa631d724290600090a3505050565b600033610721818585610d80565b610be87f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a68261074f565b610a0c7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8488261074f565b600082815260056020526040902060010154610c2d81610df2565b6107748383610ef0565b3073ffffffffffffffffffffffffffffffffffffffff831603610c9e576040517f17858bbe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161081a565b6107748383836110bd565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610d7a5781811015610d6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161081a565b610d7a8484848403610c37565b50505050565b3073ffffffffffffffffffffffffffffffffffffffff831603610de7576040517f17858bbe00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8316600482015260240161081a565b610774838383611270565b610a0c81336114df565b600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661082d57600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff85168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610e923390565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff161561082d57600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b73ffffffffffffffffffffffffffffffffffffffff8216611028576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161081a565b806002600082825461103a9190611c18565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b610a0c3382611599565b6110b3823383610ca9565b61082d8282611599565b73ffffffffffffffffffffffffffffffffffffffff831661115f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161081a565b73ffffffffffffffffffffffffffffffffffffffff8216611202576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161081a565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611313576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015260840161081a565b73ffffffffffffffffffffffffffffffffffffffff82166113b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015260840161081a565b73ffffffffffffffffffffffffffffffffffffffff83166000908152602081905260409020548181101561146c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e63650000000000000000000000000000000000000000000000000000606482015260840161081a565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610d7a565b600082815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205460ff1661082d5761151f8161175d565b61152a83602061177c565b60405160200161153b929190611c2b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290527f08c379a000000000000000000000000000000000000000000000000000000000825261081a91600401611a2c565b73ffffffffffffffffffffffffffffffffffffffff821661163c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161081a565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260208190526040902054818110156116f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161081a565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3505050565b606061067b73ffffffffffffffffffffffffffffffffffffffff831660145b6060600061178b836002611cac565b611796906002611c18565b67ffffffffffffffff8111156117ae576117ae611cc3565b6040519080825280601f01601f1916602001820160405280156117d8576020820181803683370190505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811061180f5761180f611cf2565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061187257611872611cf2565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006118ae846002611cac565b6118b9906001611c18565b90505b6001811115611956577f303132333435363738396162636465660000000000000000000000000000000085600f16601081106118fa576118fa611cf2565b1a60f81b82828151811061191057611910611cf2565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c9361194f81611d21565b90506118bc565b5083156119bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604482015260640161081a565b9392505050565b6000602082840312156119d857600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146119bf57600080fd5b60005b83811015611a23578181015183820152602001611a0b565b50506000910152565b6020815260008251806020840152611a4b816040850160208701611a08565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611aa157600080fd5b919050565b60008060408385031215611ab957600080fd5b611ac283611a7d565b946020939093013593505050565b600080600060608486031215611ae557600080fd5b611aee84611a7d565b9250611afc60208501611a7d565b9150604084013590509250925092565b600060208284031215611b1e57600080fd5b5035919050565b60008060408385031215611b3857600080fd5b82359150611b4860208401611a7d565b90509250929050565b600060208284031215611b6357600080fd5b6119bf82611a7d565b60008060408385031215611b7f57600080fd5b611b8883611a7d565b9150611b4860208401611a7d565b600181811c90821680611baa57607f821691505b602082108103611be3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561067b5761067b611be9565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611c63816017850160208801611a08565b7f206973206d697373696e6720726f6c65200000000000000000000000000000006017918401918201528351611ca0816028840160208801611a08565b01602801949350505050565b808202811582820484141761067b5761067b611be9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600081611d3057611d30611be9565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019056fea164736f6c6343000813000a",
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "approve", spender, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Approve(&_BurnMintERC20WithDrip.TransactOpts, spender, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Approve(&_BurnMintERC20WithDrip.TransactOpts, spender, amount)
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.DecreaseAllowance(&_BurnMintERC20WithDrip.TransactOpts, spender, subtractedValue)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.DecreaseAllowance(&_BurnMintERC20WithDrip.TransactOpts, spender, subtractedValue)
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.IncreaseAllowance(&_BurnMintERC20WithDrip.TransactOpts, spender, addedValue)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.IncreaseAllowance(&_BurnMintERC20WithDrip.TransactOpts, spender, addedValue)
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "renounceRole", role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.RenounceRole(&_BurnMintERC20WithDrip.TransactOpts, role, account)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.RenounceRole(&_BurnMintERC20WithDrip.TransactOpts, role, account)
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "transfer", to, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Transfer(&_BurnMintERC20WithDrip.TransactOpts, to, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.Transfer(&_BurnMintERC20WithDrip.TransactOpts, to, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.TransferFrom(&_BurnMintERC20WithDrip.TransactOpts, from, to, amount)
}

func (_BurnMintERC20WithDrip *BurnMintERC20WithDripTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BurnMintERC20WithDrip.Contract.TransferFrom(&_BurnMintERC20WithDrip.TransactOpts, from, to, amount)
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

func (_BurnMintERC20WithDrip *BurnMintERC20WithDrip) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _BurnMintERC20WithDrip.abi.Events["Approval"].ID:
		return _BurnMintERC20WithDrip.ParseApproval(log)
	case _BurnMintERC20WithDrip.abi.Events["CCIPAdminTransferred"].ID:
		return _BurnMintERC20WithDrip.ParseCCIPAdminTransferred(log)
	case _BurnMintERC20WithDrip.abi.Events["RoleAdminChanged"].ID:
		return _BurnMintERC20WithDrip.ParseRoleAdminChanged(log)
	case _BurnMintERC20WithDrip.abi.Events["RoleGranted"].ID:
		return _BurnMintERC20WithDrip.ParseRoleGranted(log)
	case _BurnMintERC20WithDrip.abi.Events["RoleRevoked"].ID:
		return _BurnMintERC20WithDrip.ParseRoleRevoked(log)
	case _BurnMintERC20WithDrip.abi.Events["Transfer"].ID:
		return _BurnMintERC20WithDrip.ParseTransfer(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
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

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Burn0(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	Drip(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	SetCCIPAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

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

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
