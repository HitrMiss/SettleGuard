// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Reference imports to suppress errors if they are not otherwise used.
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

// BondVaultMetaData contains all meta data concerning the BondVault contract.
var BondVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_withdrawFeeBps\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawFlatFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRequested\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"readyAt\",\"type\":\"uint256\"}],\"name\":\"ClaimTooFresh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBond\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"unlockAt\",\"type\":\"uint64\"}],\"name\":\"Locked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWithdrawRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEngine\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestExceedsBond\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"unlockAt\",\"type\":\"uint64\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"WithdrawCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"}],\"name\":\"WithdrawProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"requestedAt\",\"type\":\"uint64\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLAIM_MIN_AGE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_LOCK_SECS\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bonds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelWithdrawRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_party\",\"type\":\"address\"}],\"name\":\"getActiveBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"}],\"name\":\"getAvailableBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lockRisk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockedRisk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"processWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newEngine\",\"type\":\"address\"}],\"name\":\"setSettlementEngine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementEngine\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"slashBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockAt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlockRisk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawClaims\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"requestedAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeBps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFlatFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b5060405161308d38038061308d833981810160405281019061003291906102a4565b60015f819055505f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16148061009e57505f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16145b806100d457505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16145b1561010b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508373ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508273ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250508167ffffffffffffffff1660e08167ffffffffffffffff1681525050806101008181525050505050505061031b565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610203826101da565b9050919050565b610213816101f9565b811461021d575f5ffd5b50565b5f8151905061022e8161020a565b92915050565b5f67ffffffffffffffff82169050919050565b61025081610234565b811461025a575f5ffd5b50565b5f8151905061026b81610247565b92915050565b5f819050919050565b61028381610271565b811461028d575f5ffd5b50565b5f8151905061029e8161027a565b92915050565b5f5f5f5f5f60a086880312156102bd576102bc6101d6565b5b5f6102ca88828901610220565b95505060206102db88828901610220565b94505060406102ec88828901610220565b93505060606102fd8882890161025d565b925050608061030e88828901610290565b9150509295509295909350565b60805160a05160c05160e05161010051612ca96103e45f395f81816105af01528181610d1401528181610da901526124f401525f818161098401528181610d6e01526124b401525f81816105f70152818161135401526120ed01525f818161055c015281816105d301528181610b40015281816113000152818161137a01528181611a570152818161209d015261210f01525f818161061b01528181610e2e01528181610e6a015281816115e10152818161161d01528181611b2a0152611b660152612ca95ff3fe608060405234801561000f575f5ffd5b5060043610610171575f3560e01c8063ab6103cb116100dc578063e6d0919d11610095578063f22436dd1161006f578063f22436dd146103ff578063f67ed3f71461041b578063f89f4acf14610437578063fe10d7741461046757610171565b8063e6d0919d146103a9578063edb00098146103c7578063f10bd31b146103e357610171565b8063ab6103cb146102f9578063aeadff771461032b578063b6b55f251461035b578063cb9609cb14610377578063e19e716814610395578063e5cd8b6a1461039f57610171565b80635aa6e6751161012e5780635aa6e67514610237578063745400c9146102555780637a16f47b146102715780637aca2b591461028f57806389b7e586146102ad5780638b4fb30d146102dd57610171565b806301e36610146101755780632e1a7d4d146101a55780632f4f21e2146101c1578063334c03f6146101dd5780633e413bee146101fb5780634690484014610219575b5f5ffd5b61018f600480360381019061018a919061269c565b610497565b60405161019c91906126df565b60405180910390f35b6101bf60048036038101906101ba9190612722565b610525565b005b6101db60048036038101906101d6919061274d565b610542565b005b6101e56105ad565b6040516101f291906126df565b60405180910390f35b6102036105d1565b60405161021091906127e6565b60405180910390f35b6102216105f5565b60405161022e919061280e565b60405180910390f35b61023f610619565b60405161024c9190612847565b60405180910390f35b61026f600480360381019061026a9190612722565b61063d565b005b61027961097b565b6040516102869190612882565b60405180910390f35b610297610982565b6040516102a49190612882565b60405180910390f35b6102c760048036038101906102c2919061269c565b6109a6565b6040516102d49190612882565b60405180910390f35b6102f760048036038101906102f2919061274d565b6109ca565b005b610313600480360381019061030e919061269c565b610aa7565b6040516103229392919061289b565b60405180910390f35b6103456004803603810190610340919061269c565b610ae0565b60405161035291906126df565b60405180910390f35b61037560048036038101906103709190612722565b610b26565b005b61037f610b90565b60405161038c9190612882565b60405180910390f35b61039d610b95565b005b6103a7610cc9565b005b6103b1610dff565b6040516103be919061280e565b60405180910390f35b6103e160048036038101906103dc919061269c565b610e24565b005b6103fd60048036038101906103f8919061274d565b61143b565b005b610419600480360381019061041491906128d0565b6115d7565b005b6104356004803603810190610430919061269c565b611b28565b005b610451600480360381019061044c919061269c565b611d2c565b60405161045e91906126df565b60405180910390f35b610481600480360381019061047c919061269c565b611d41565b60405161048e91906126df565b60405180910390f35b5f60045f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205460025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205461051e919061294d565b9050919050565b61052d611d56565b6105373382611d9a565b61053f6121ab565b50565b61054a611d56565b61055482826121b4565b6105a13330837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661242c909392919063ffffffff16565b6105a96121ab565b5050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b610645611d56565b5f810361067e576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1642101561076d5760035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff166040517feada26580000000000000000000000000000000000000000000000000000000081526004016107649190612882565b60405180910390fd5b5f60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f0154146107e5576040517f5303ed8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6107ef826124ae565b90505f81836107fe9190612980565b90508060025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610877576040517fe92c469f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60405180606001604052808481526020018381526020014267ffffffffffffffff1681525060055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f820151815f0155602082015181600101556040820151816002015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055509050503373ffffffffffffffffffffffffffffffffffffffff167f5b5919a27b29ed9ea8fdb44f39ca081298229768813e2b5ffbd9d83eeb1de61e8484426040516109669392919061289b565b60405180910390a250506109786121ab565b50565b62093a8081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6003602052805f5260405f205f915054906101000a900467ffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a50576040517f8256cbfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610a9c919061294d565b925050819055505050565b6005602052805f5260405f205f91509050805f015490806001015490806002015f9054906101000a900467ffffffffffffffff16905083565b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b610b2e611d56565b610b3833826121b4565b610b853330837f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661242c909392919063ffffffff16565b610b8d6121ab565b50565b603c81565b610b9d611d56565b5f60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015403610c15576040517fa124d31700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f5f82015f9055600182015f9055600282015f6101000a81549067ffffffffffffffff021916905550503373ffffffffffffffffffffffffffffffffffffffff167fe4c65151920d34ba91eb49b98f4904c54654d4e5f450dc8387700c09c25a846a60405160405180910390a2610cc76121ab565b565b610cd1611d56565b5f60025f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490507f00000000000000000000000000000000000000000000000000000000000000008111610d6b576040517fe92c469f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f0000000000000000000000000000000000000000000000000000000000000000612710610d9a91906129b3565b67ffffffffffffffff166127107f000000000000000000000000000000000000000000000000000000000000000084610dd3919061294d565b610ddd91906129ee565b610de79190612a5c565b9050610df33382611d9a565b5050610dfd6121ab565b565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610e2c611d56565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634d104adf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ed1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ef59190612abf565b336040518363ffffffff1660e01b8152600401610f13929190612af9565b602060405180830381865afa158015610f2e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f529190612b55565b610f88576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060600160405290815f820154815260200160018201548152602001600282015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505090505f815f015103611053576040517fa124d31700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c816040015161106491906129b3565b67ffffffffffffffff164210156110c257603c816040015161108691906129b3565b6040517fbe0772cc0000000000000000000000000000000000000000000000000000000081526004016110b99190612bb0565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff164210156111b15760035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff166040517feada26580000000000000000000000000000000000000000000000000000000081526004016111a89190612882565b60405180910390fd5b5f8160200151825f01516111c59190612980565b90505f60025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015611242576040517fce622d5b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b818161124e919061294d565b60025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f5f82015f9055600182015f9055600282015f6101000a81549067ffffffffffffffff0219169055505061134484845f01517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166125249092919063ffffffff16565b5f8360200151146113bf576113be7f000000000000000000000000000000000000000000000000000000000000000084602001517f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166125249092919063ffffffff16565b5b3373ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fd4297feff9aaad2ad18f51ed5932f7d8396d1071010eb5126cf627285b3efb1d855f01518660200151604051611425929190612bc9565b60405180910390a35050506114386121ab565b50565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146114c1576040517f8256cbfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205460025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054611548919061294d565b1015611580576040517fe92c469f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060045f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546115cc9190612980565b925050819055505050565b6115df611d56565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634d104adf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611684573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116a89190612abf565b336040518363ffffffff1660e01b81526004016116c6929190612af9565b602060405180830381865afa1580156116e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117059190612b55565b61173b576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614806117a057505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b156117d7576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8203611810576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60025f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f810361188a576040517fe92c469f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f818411611898578361189a565b815b90505f81836118a9919061294d565b90508060025f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f60055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f206040518060600160405290815f820154815260200160018201548152602001600282015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152505090505f815f0151141580156119a057508060200151815f015161199d9190612980565b82105b15611a505760055f8873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f5f82015f9055600182015f9055600282015f6101000a81549067ffffffffffffffff021916905550508673ffffffffffffffffffffffffffffffffffffffff167fe4c65151920d34ba91eb49b98f4904c54654d4e5f450dc8387700c09c25a846a60405160405180910390a25b611a9b85847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166125249092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff167f1b9817da702c2ae84a03f2a48072c6a81bfc3b7d1d5615bffce3db4856831bb786604051611b0f91906126df565b60405180910390a450505050611b236121ab565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377fb8d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611bcd573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bf19190612abf565b336040518363ffffffff1660e01b8152600401611c0f929190612af9565b602060405180830381865afa158015611c2a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c4e9190612b55565b611c84576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611ce9576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6004602052805f5260405f205f915090505481565b6002602052805f5260405f205f915090505481565b60025f5403611d91576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f81905550565b5f8103611dd3576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16421015611ec25760035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff166040517feada2658000000000000000000000000000000000000000000000000000000008152600401611eb99190612882565b60405180910390fd5b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015414611fb35760055f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f5f82015f9055600182015f9055600282015f6101000a81549067ffffffffffffffff021916905550508173ffffffffffffffffffffffffffffffffffffffff167fe4c65151920d34ba91eb49b98f4904c54654d4e5f450dc8387700c09c25a846a60405160405180910390a25b5f611fbd826124ae565b90505f8183611fcc9190612980565b90505f60025f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015612049576040517fe92c469f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8181612055919061294d565b60025f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055506120e185857f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166125249092919063ffffffff16565b5f8314612154576121537f0000000000000000000000000000000000000000000000000000000000000000847f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166125249092919063ffffffff16565b5b8473ffffffffffffffffffffffffffffffffffffffff167f92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6858560405161219c929190612bc9565b60405180910390a25050505050565b60015f81905550565b5f81036121ed576040517f1f2a200500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603612252576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461229e9190612980565b925050819055505f62093a80426122b591906129b3565b905060035f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161115612387578060035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b8273ffffffffffffffffffffffffffffffffffffffff167fcd50943caf3a035e3821448b2743faef90aae73ff20c6c1b29ea8213ad1fee1d8360035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1660405161241f929190612bf0565b60405180910390a2505050565b6124a8848573ffffffffffffffffffffffffffffffffffffffff166323b872dd86868660405160240161246193929190612c17565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506125a3565b50505050565b5f6127107f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16836124e891906129ee565b6124f29190612a5c565b7f000000000000000000000000000000000000000000000000000000000000000061251d9190612980565b9050919050565b61259e838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8585604051602401612557929190612c4c565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506125a3565b505050565b5f5f60205f8451602086015f885af1806125c2576040513d5f823e3d81fd5b3d92505f519150505f82146125db5760018114156125f6565b5f8473ffffffffffffffffffffffffffffffffffffffff163b145b1561263857836040517f5274afe700000000000000000000000000000000000000000000000000000000815260040161262f919061280e565b60405180910390fd5b50505050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61266b82612642565b9050919050565b61267b81612661565b8114612685575f5ffd5b50565b5f8135905061269681612672565b92915050565b5f602082840312156126b1576126b061263e565b5b5f6126be84828501612688565b91505092915050565b5f819050919050565b6126d9816126c7565b82525050565b5f6020820190506126f25f8301846126d0565b92915050565b612701816126c7565b811461270b575f5ffd5b50565b5f8135905061271c816126f8565b92915050565b5f602082840312156127375761273661263e565b5b5f6127448482850161270e565b91505092915050565b5f5f604083850312156127635761276261263e565b5b5f61277085828601612688565b92505060206127818582860161270e565b9150509250929050565b5f819050919050565b5f6127ae6127a96127a484612642565b61278b565b612642565b9050919050565b5f6127bf82612794565b9050919050565b5f6127d0826127b5565b9050919050565b6127e0816127c6565b82525050565b5f6020820190506127f95f8301846127d7565b92915050565b61280881612661565b82525050565b5f6020820190506128215f8301846127ff565b92915050565b5f612831826127b5565b9050919050565b61284181612827565b82525050565b5f60208201905061285a5f830184612838565b92915050565b5f67ffffffffffffffff82169050919050565b61287c81612860565b82525050565b5f6020820190506128955f830184612873565b92915050565b5f6060820190506128ae5f8301866126d0565b6128bb60208301856126d0565b6128c86040830184612873565b949350505050565b5f5f5f606084860312156128e7576128e661263e565b5b5f6128f486828701612688565b93505060206129058682870161270e565b925050604061291686828701612688565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f612957826126c7565b9150612962836126c7565b925082820390508181111561297a57612979612920565b5b92915050565b5f61298a826126c7565b9150612995836126c7565b92508282019050808211156129ad576129ac612920565b5b92915050565b5f6129bd82612860565b91506129c883612860565b9250828201905067ffffffffffffffff8111156129e8576129e7612920565b5b92915050565b5f6129f8826126c7565b9150612a03836126c7565b9250828202612a11816126c7565b91508282048414831517612a2857612a27612920565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f612a66826126c7565b9150612a71836126c7565b925082612a8157612a80612a2f565b5b828204905092915050565b5f819050919050565b612a9e81612a8c565b8114612aa8575f5ffd5b50565b5f81519050612ab981612a95565b92915050565b5f60208284031215612ad457612ad361263e565b5b5f612ae184828501612aab565b91505092915050565b612af381612a8c565b82525050565b5f604082019050612b0c5f830185612aea565b612b1960208301846127ff565b9392505050565b5f8115159050919050565b612b3481612b20565b8114612b3e575f5ffd5b50565b5f81519050612b4f81612b2b565b92915050565b5f60208284031215612b6a57612b6961263e565b5b5f612b7784828501612b41565b91505092915050565b5f612b9a612b95612b9084612860565b61278b565b6126c7565b9050919050565b612baa81612b80565b82525050565b5f602082019050612bc35f830184612ba1565b92915050565b5f604082019050612bdc5f8301856126d0565b612be960208301846126d0565b9392505050565b5f604082019050612c035f8301856126d0565b612c106020830184612873565b9392505050565b5f606082019050612c2a5f8301866127ff565b612c3760208301856127ff565b612c4460408301846126d0565b949350505050565b5f604082019050612c5f5f8301856127ff565b612c6c60208301846126d0565b939250505056fea2646970667358221220c3ffb00e5dc552afb820948087347daf5397582e8e31b4ff8b1f1d9d48032d9664736f6c63430008210033",
}

// BondVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use BondVaultMetaData.ABI instead.
var BondVaultABI = BondVaultMetaData.ABI

// BondVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BondVaultMetaData.Bin instead.
var BondVaultBin = BondVaultMetaData.Bin

// DeployBondVault deploys a new Ethereum contract, binding an instance of BondVault to it.
func DeployBondVault(auth *bind.TransactOpts, backend bind.ContractBackend, _gov common.Address, _usdc common.Address, _feeRecipient common.Address, _withdrawFeeBps uint64, _withdrawFlatFee *big.Int) (common.Address, *types.Transaction, *BondVault, error) {
	parsed, err := BondVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BondVaultBin), backend, _gov, _usdc, _feeRecipient, _withdrawFeeBps, _withdrawFlatFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BondVault{BondVaultCaller: BondVaultCaller{contract: contract}, BondVaultTransactor: BondVaultTransactor{contract: contract}, BondVaultFilterer: BondVaultFilterer{contract: contract}}, nil
}

// BondVault is an auto generated Go binding around an Ethereum contract.
type BondVault struct {
	BondVaultCaller     // Read-only binding to the contract
	BondVaultTransactor // Write-only binding to the contract
	BondVaultFilterer   // Log filterer for contract events
}

// BondVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondVaultSession struct {
	Contract     *BondVault        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondVaultCallerSession struct {
	Contract *BondVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BondVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondVaultTransactorSession struct {
	Contract     *BondVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BondVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondVaultRaw struct {
	Contract *BondVault // Generic contract binding to access the raw methods on
}

// BondVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondVaultCallerRaw struct {
	Contract *BondVaultCaller // Generic read-only contract binding to access the raw methods on
}

// BondVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondVaultTransactorRaw struct {
	Contract *BondVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondVault creates a new instance of BondVault, bound to a specific deployed contract.
func NewBondVault(address common.Address, backend bind.ContractBackend) (*BondVault, error) {
	contract, err := bindBondVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BondVault{BondVaultCaller: BondVaultCaller{contract: contract}, BondVaultTransactor: BondVaultTransactor{contract: contract}, BondVaultFilterer: BondVaultFilterer{contract: contract}}, nil
}

// NewBondVaultCaller creates a new read-only instance of BondVault, bound to a specific deployed contract.
func NewBondVaultCaller(address common.Address, caller bind.ContractCaller) (*BondVaultCaller, error) {
	contract, err := bindBondVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondVaultCaller{contract: contract}, nil
}

// NewBondVaultTransactor creates a new write-only instance of BondVault, bound to a specific deployed contract.
func NewBondVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*BondVaultTransactor, error) {
	contract, err := bindBondVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondVaultTransactor{contract: contract}, nil
}

// NewBondVaultFilterer creates a new log filterer instance of BondVault, bound to a specific deployed contract.
func NewBondVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*BondVaultFilterer, error) {
	contract, err := bindBondVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondVaultFilterer{contract: contract}, nil
}

// bindBondVault binds a generic wrapper to an already deployed contract.
func bindBondVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BondVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondVault *BondVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondVault.Contract.BondVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondVault *BondVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondVault.Contract.BondVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondVault *BondVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondVault.Contract.BondVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondVault *BondVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondVault *BondVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondVault *BondVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondVault.Contract.contract.Transact(opts, method, params...)
}

// CLAIMMINAGE is a free data retrieval call binding the contract method 0xcb9609cb.
//
// Solidity: function CLAIM_MIN_AGE() view returns(uint64)
func (_BondVault *BondVaultCaller) CLAIMMINAGE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "CLAIM_MIN_AGE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// CLAIMMINAGE is a free data retrieval call binding the contract method 0xcb9609cb.
//
// Solidity: function CLAIM_MIN_AGE() view returns(uint64)
func (_BondVault *BondVaultSession) CLAIMMINAGE() (uint64, error) {
	return _BondVault.Contract.CLAIMMINAGE(&_BondVault.CallOpts)
}

// CLAIMMINAGE is a free data retrieval call binding the contract method 0xcb9609cb.
//
// Solidity: function CLAIM_MIN_AGE() view returns(uint64)
func (_BondVault *BondVaultCallerSession) CLAIMMINAGE() (uint64, error) {
	return _BondVault.Contract.CLAIMMINAGE(&_BondVault.CallOpts)
}

// DEPOSITLOCKSECS is a free data retrieval call binding the contract method 0x7a16f47b.
//
// Solidity: function DEPOSIT_LOCK_SECS() view returns(uint64)
func (_BondVault *BondVaultCaller) DEPOSITLOCKSECS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "DEPOSIT_LOCK_SECS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DEPOSITLOCKSECS is a free data retrieval call binding the contract method 0x7a16f47b.
//
// Solidity: function DEPOSIT_LOCK_SECS() view returns(uint64)
func (_BondVault *BondVaultSession) DEPOSITLOCKSECS() (uint64, error) {
	return _BondVault.Contract.DEPOSITLOCKSECS(&_BondVault.CallOpts)
}

// DEPOSITLOCKSECS is a free data retrieval call binding the contract method 0x7a16f47b.
//
// Solidity: function DEPOSIT_LOCK_SECS() view returns(uint64)
func (_BondVault *BondVaultCallerSession) DEPOSITLOCKSECS() (uint64, error) {
	return _BondVault.Contract.DEPOSITLOCKSECS(&_BondVault.CallOpts)
}

// Bonds is a free data retrieval call binding the contract method 0xfe10d774.
//
// Solidity: function bonds(address ) view returns(uint256)
func (_BondVault *BondVaultCaller) Bonds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "bonds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bonds is a free data retrieval call binding the contract method 0xfe10d774.
//
// Solidity: function bonds(address ) view returns(uint256)
func (_BondVault *BondVaultSession) Bonds(arg0 common.Address) (*big.Int, error) {
	return _BondVault.Contract.Bonds(&_BondVault.CallOpts, arg0)
}

// Bonds is a free data retrieval call binding the contract method 0xfe10d774.
//
// Solidity: function bonds(address ) view returns(uint256)
func (_BondVault *BondVaultCallerSession) Bonds(arg0 common.Address) (*big.Int, error) {
	return _BondVault.Contract.Bonds(&_BondVault.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BondVault *BondVaultCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BondVault *BondVaultSession) FeeRecipient() (common.Address, error) {
	return _BondVault.Contract.FeeRecipient(&_BondVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BondVault *BondVaultCallerSession) FeeRecipient() (common.Address, error) {
	return _BondVault.Contract.FeeRecipient(&_BondVault.CallOpts)
}

// GetActiveBond is a free data retrieval call binding the contract method 0xaeadff77.
//
// Solidity: function getActiveBond(address _party) view returns(uint256)
func (_BondVault *BondVaultCaller) GetActiveBond(opts *bind.CallOpts, _party common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "getActiveBond", _party)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveBond is a free data retrieval call binding the contract method 0xaeadff77.
//
// Solidity: function getActiveBond(address _party) view returns(uint256)
func (_BondVault *BondVaultSession) GetActiveBond(_party common.Address) (*big.Int, error) {
	return _BondVault.Contract.GetActiveBond(&_BondVault.CallOpts, _party)
}

// GetActiveBond is a free data retrieval call binding the contract method 0xaeadff77.
//
// Solidity: function getActiveBond(address _party) view returns(uint256)
func (_BondVault *BondVaultCallerSession) GetActiveBond(_party common.Address) (*big.Int, error) {
	return _BondVault.Contract.GetActiveBond(&_BondVault.CallOpts, _party)
}

// GetAvailableBond is a free data retrieval call binding the contract method 0x01e36610.
//
// Solidity: function getAvailableBond(address _merchant) view returns(uint256)
func (_BondVault *BondVaultCaller) GetAvailableBond(opts *bind.CallOpts, _merchant common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "getAvailableBond", _merchant)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableBond is a free data retrieval call binding the contract method 0x01e36610.
//
// Solidity: function getAvailableBond(address _merchant) view returns(uint256)
func (_BondVault *BondVaultSession) GetAvailableBond(_merchant common.Address) (*big.Int, error) {
	return _BondVault.Contract.GetAvailableBond(&_BondVault.CallOpts, _merchant)
}

// GetAvailableBond is a free data retrieval call binding the contract method 0x01e36610.
//
// Solidity: function getAvailableBond(address _merchant) view returns(uint256)
func (_BondVault *BondVaultCallerSession) GetAvailableBond(_merchant common.Address) (*big.Int, error) {
	return _BondVault.Contract.GetAvailableBond(&_BondVault.CallOpts, _merchant)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_BondVault *BondVaultCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_BondVault *BondVaultSession) Governance() (common.Address, error) {
	return _BondVault.Contract.Governance(&_BondVault.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_BondVault *BondVaultCallerSession) Governance() (common.Address, error) {
	return _BondVault.Contract.Governance(&_BondVault.CallOpts)
}

// LockedRisk is a free data retrieval call binding the contract method 0xf89f4acf.
//
// Solidity: function lockedRisk(address ) view returns(uint256)
func (_BondVault *BondVaultCaller) LockedRisk(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "lockedRisk", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedRisk is a free data retrieval call binding the contract method 0xf89f4acf.
//
// Solidity: function lockedRisk(address ) view returns(uint256)
func (_BondVault *BondVaultSession) LockedRisk(arg0 common.Address) (*big.Int, error) {
	return _BondVault.Contract.LockedRisk(&_BondVault.CallOpts, arg0)
}

// LockedRisk is a free data retrieval call binding the contract method 0xf89f4acf.
//
// Solidity: function lockedRisk(address ) view returns(uint256)
func (_BondVault *BondVaultCallerSession) LockedRisk(arg0 common.Address) (*big.Int, error) {
	return _BondVault.Contract.LockedRisk(&_BondVault.CallOpts, arg0)
}

// SettlementEngine is a free data retrieval call binding the contract method 0xe6d0919d.
//
// Solidity: function settlementEngine() view returns(address)
func (_BondVault *BondVaultCaller) SettlementEngine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "settlementEngine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SettlementEngine is a free data retrieval call binding the contract method 0xe6d0919d.
//
// Solidity: function settlementEngine() view returns(address)
func (_BondVault *BondVaultSession) SettlementEngine() (common.Address, error) {
	return _BondVault.Contract.SettlementEngine(&_BondVault.CallOpts)
}

// SettlementEngine is a free data retrieval call binding the contract method 0xe6d0919d.
//
// Solidity: function settlementEngine() view returns(address)
func (_BondVault *BondVaultCallerSession) SettlementEngine() (common.Address, error) {
	return _BondVault.Contract.SettlementEngine(&_BondVault.CallOpts)
}

// UnlockAt is a free data retrieval call binding the contract method 0x89b7e586.
//
// Solidity: function unlockAt(address ) view returns(uint64)
func (_BondVault *BondVaultCaller) UnlockAt(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "unlockAt", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// UnlockAt is a free data retrieval call binding the contract method 0x89b7e586.
//
// Solidity: function unlockAt(address ) view returns(uint64)
func (_BondVault *BondVaultSession) UnlockAt(arg0 common.Address) (uint64, error) {
	return _BondVault.Contract.UnlockAt(&_BondVault.CallOpts, arg0)
}

// UnlockAt is a free data retrieval call binding the contract method 0x89b7e586.
//
// Solidity: function unlockAt(address ) view returns(uint64)
func (_BondVault *BondVaultCallerSession) UnlockAt(arg0 common.Address) (uint64, error) {
	return _BondVault.Contract.UnlockAt(&_BondVault.CallOpts, arg0)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_BondVault *BondVaultCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_BondVault *BondVaultSession) Usdc() (common.Address, error) {
	return _BondVault.Contract.Usdc(&_BondVault.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_BondVault *BondVaultCallerSession) Usdc() (common.Address, error) {
	return _BondVault.Contract.Usdc(&_BondVault.CallOpts)
}

// WithdrawClaims is a free data retrieval call binding the contract method 0xab6103cb.
//
// Solidity: function withdrawClaims(address ) view returns(uint256 amount, uint256 fee, uint64 requestedAt)
func (_BondVault *BondVaultCaller) WithdrawClaims(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount      *big.Int
	Fee         *big.Int
	RequestedAt uint64
}, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "withdrawClaims", arg0)

	outstruct := new(struct {
		Amount      *big.Int
		Fee         *big.Int
		RequestedAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestedAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// WithdrawClaims is a free data retrieval call binding the contract method 0xab6103cb.
//
// Solidity: function withdrawClaims(address ) view returns(uint256 amount, uint256 fee, uint64 requestedAt)
func (_BondVault *BondVaultSession) WithdrawClaims(arg0 common.Address) (struct {
	Amount      *big.Int
	Fee         *big.Int
	RequestedAt uint64
}, error) {
	return _BondVault.Contract.WithdrawClaims(&_BondVault.CallOpts, arg0)
}

// WithdrawClaims is a free data retrieval call binding the contract method 0xab6103cb.
//
// Solidity: function withdrawClaims(address ) view returns(uint256 amount, uint256 fee, uint64 requestedAt)
func (_BondVault *BondVaultCallerSession) WithdrawClaims(arg0 common.Address) (struct {
	Amount      *big.Int
	Fee         *big.Int
	RequestedAt uint64
}, error) {
	return _BondVault.Contract.WithdrawClaims(&_BondVault.CallOpts, arg0)
}

// WithdrawFeeBps is a free data retrieval call binding the contract method 0x7aca2b59.
//
// Solidity: function withdrawFeeBps() view returns(uint64)
func (_BondVault *BondVaultCaller) WithdrawFeeBps(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "withdrawFeeBps")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// WithdrawFeeBps is a free data retrieval call binding the contract method 0x7aca2b59.
//
// Solidity: function withdrawFeeBps() view returns(uint64)
func (_BondVault *BondVaultSession) WithdrawFeeBps() (uint64, error) {
	return _BondVault.Contract.WithdrawFeeBps(&_BondVault.CallOpts)
}

// WithdrawFeeBps is a free data retrieval call binding the contract method 0x7aca2b59.
//
// Solidity: function withdrawFeeBps() view returns(uint64)
func (_BondVault *BondVaultCallerSession) WithdrawFeeBps() (uint64, error) {
	return _BondVault.Contract.WithdrawFeeBps(&_BondVault.CallOpts)
}

// WithdrawFlatFee is a free data retrieval call binding the contract method 0x334c03f6.
//
// Solidity: function withdrawFlatFee() view returns(uint256)
func (_BondVault *BondVaultCaller) WithdrawFlatFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BondVault.contract.Call(opts, &out, "withdrawFlatFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFlatFee is a free data retrieval call binding the contract method 0x334c03f6.
//
// Solidity: function withdrawFlatFee() view returns(uint256)
func (_BondVault *BondVaultSession) WithdrawFlatFee() (*big.Int, error) {
	return _BondVault.Contract.WithdrawFlatFee(&_BondVault.CallOpts)
}

// WithdrawFlatFee is a free data retrieval call binding the contract method 0x334c03f6.
//
// Solidity: function withdrawFlatFee() view returns(uint256)
func (_BondVault *BondVaultCallerSession) WithdrawFlatFee() (*big.Int, error) {
	return _BondVault.Contract.WithdrawFlatFee(&_BondVault.CallOpts)
}

// CancelWithdrawRequest is a paid mutator transaction binding the contract method 0xe19e7168.
//
// Solidity: function cancelWithdrawRequest() returns()
func (_BondVault *BondVaultTransactor) CancelWithdrawRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "cancelWithdrawRequest")
}

// CancelWithdrawRequest is a paid mutator transaction binding the contract method 0xe19e7168.
//
// Solidity: function cancelWithdrawRequest() returns()
func (_BondVault *BondVaultSession) CancelWithdrawRequest() (*types.Transaction, error) {
	return _BondVault.Contract.CancelWithdrawRequest(&_BondVault.TransactOpts)
}

// CancelWithdrawRequest is a paid mutator transaction binding the contract method 0xe19e7168.
//
// Solidity: function cancelWithdrawRequest() returns()
func (_BondVault *BondVaultTransactorSession) CancelWithdrawRequest() (*types.Transaction, error) {
	return _BondVault.Contract.CancelWithdrawRequest(&_BondVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_BondVault *BondVaultTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_BondVault *BondVaultSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.Deposit(&_BondVault.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_BondVault *BondVaultTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.Deposit(&_BondVault.TransactOpts, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address _user, uint256 _amount) returns()
func (_BondVault *BondVaultTransactor) DepositFor(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "depositFor", _user, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address _user, uint256 _amount) returns()
func (_BondVault *BondVaultSession) DepositFor(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.DepositFor(&_BondVault.TransactOpts, _user, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address _user, uint256 _amount) returns()
func (_BondVault *BondVaultTransactorSession) DepositFor(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.DepositFor(&_BondVault.TransactOpts, _user, _amount)
}

// LockRisk is a paid mutator transaction binding the contract method 0xf10bd31b.
//
// Solidity: function lockRisk(address _merchant, uint256 _amount) returns()
func (_BondVault *BondVaultTransactor) LockRisk(opts *bind.TransactOpts, _merchant common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "lockRisk", _merchant, _amount)
}

// LockRisk is a paid mutator transaction binding the contract method 0xf10bd31b.
//
// Solidity: function lockRisk(address _merchant, uint256 _amount) returns()
func (_BondVault *BondVaultSession) LockRisk(_merchant common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.LockRisk(&_BondVault.TransactOpts, _merchant, _amount)
}

// LockRisk is a paid mutator transaction binding the contract method 0xf10bd31b.
//
// Solidity: function lockRisk(address _merchant, uint256 _amount) returns()
func (_BondVault *BondVaultTransactorSession) LockRisk(_merchant common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.LockRisk(&_BondVault.TransactOpts, _merchant, _amount)
}

// ProcessWithdraw is a paid mutator transaction binding the contract method 0xedb00098.
//
// Solidity: function processWithdraw(address _user) returns()
func (_BondVault *BondVaultTransactor) ProcessWithdraw(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "processWithdraw", _user)
}

// ProcessWithdraw is a paid mutator transaction binding the contract method 0xedb00098.
//
// Solidity: function processWithdraw(address _user) returns()
func (_BondVault *BondVaultSession) ProcessWithdraw(_user common.Address) (*types.Transaction, error) {
	return _BondVault.Contract.ProcessWithdraw(&_BondVault.TransactOpts, _user)
}

// ProcessWithdraw is a paid mutator transaction binding the contract method 0xedb00098.
//
// Solidity: function processWithdraw(address _user) returns()
func (_BondVault *BondVaultTransactorSession) ProcessWithdraw(_user common.Address) (*types.Transaction, error) {
	return _BondVault.Contract.ProcessWithdraw(&_BondVault.TransactOpts, _user)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_BondVault *BondVaultTransactor) RequestWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "requestWithdraw", _amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_BondVault *BondVaultSession) RequestWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.RequestWithdraw(&_BondVault.TransactOpts, _amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_BondVault *BondVaultTransactorSession) RequestWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.RequestWithdraw(&_BondVault.TransactOpts, _amount)
}

// SetSettlementEngine is a paid mutator transaction binding the contract method 0xf67ed3f7.
//
// Solidity: function setSettlementEngine(address _newEngine) returns()
func (_BondVault *BondVaultTransactor) SetSettlementEngine(opts *bind.TransactOpts, _newEngine common.Address) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "setSettlementEngine", _newEngine)
}

// SetSettlementEngine is a paid mutator transaction binding the contract method 0xf67ed3f7.
//
// Solidity: function setSettlementEngine(address _newEngine) returns()
func (_BondVault *BondVaultSession) SetSettlementEngine(_newEngine common.Address) (*types.Transaction, error) {
	return _BondVault.Contract.SetSettlementEngine(&_BondVault.TransactOpts, _newEngine)
}

// SetSettlementEngine is a paid mutator transaction binding the contract method 0xf67ed3f7.
//
// Solidity: function setSettlementEngine(address _newEngine) returns()
func (_BondVault *BondVaultTransactorSession) SetSettlementEngine(_newEngine common.Address) (*types.Transaction, error) {
	return _BondVault.Contract.SetSettlementEngine(&_BondVault.TransactOpts, _newEngine)
}

// SlashBond is a paid mutator transaction binding the contract method 0xf22436dd.
//
// Solidity: function slashBond(address _user, uint256 _amount, address _recipient) returns()
func (_BondVault *BondVaultTransactor) SlashBond(opts *bind.TransactOpts, _user common.Address, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "slashBond", _user, _amount, _recipient)
}

// SlashBond is a paid mutator transaction binding the contract method 0xf22436dd.
//
// Solidity: function slashBond(address _user, uint256 _amount, address _recipient) returns()
func (_BondVault *BondVaultSession) SlashBond(_user common.Address, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _BondVault.Contract.SlashBond(&_BondVault.TransactOpts, _user, _amount, _recipient)
}

// SlashBond is a paid mutator transaction binding the contract method 0xf22436dd.
//
// Solidity: function slashBond(address _user, uint256 _amount, address _recipient) returns()
func (_BondVault *BondVaultTransactorSession) SlashBond(_user common.Address, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _BondVault.Contract.SlashBond(&_BondVault.TransactOpts, _user, _amount, _recipient)
}

// UnlockRisk is a paid mutator transaction binding the contract method 0x8b4fb30d.
//
// Solidity: function unlockRisk(address _merchant, uint256 _amount) returns()
func (_BondVault *BondVaultTransactor) UnlockRisk(opts *bind.TransactOpts, _merchant common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "unlockRisk", _merchant, _amount)
}

// UnlockRisk is a paid mutator transaction binding the contract method 0x8b4fb30d.
//
// Solidity: function unlockRisk(address _merchant, uint256 _amount) returns()
func (_BondVault *BondVaultSession) UnlockRisk(_merchant common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.UnlockRisk(&_BondVault.TransactOpts, _merchant, _amount)
}

// UnlockRisk is a paid mutator transaction binding the contract method 0x8b4fb30d.
//
// Solidity: function unlockRisk(address _merchant, uint256 _amount) returns()
func (_BondVault *BondVaultTransactorSession) UnlockRisk(_merchant common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.UnlockRisk(&_BondVault.TransactOpts, _merchant, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_BondVault *BondVaultTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_BondVault *BondVaultSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.Withdraw(&_BondVault.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_BondVault *BondVaultTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _BondVault.Contract.Withdraw(&_BondVault.TransactOpts, _amount)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xe5cd8b6a.
//
// Solidity: function withdrawMax() returns()
func (_BondVault *BondVaultTransactor) WithdrawMax(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondVault.contract.Transact(opts, "withdrawMax")
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xe5cd8b6a.
//
// Solidity: function withdrawMax() returns()
func (_BondVault *BondVaultSession) WithdrawMax() (*types.Transaction, error) {
	return _BondVault.Contract.WithdrawMax(&_BondVault.TransactOpts)
}

// WithdrawMax is a paid mutator transaction binding the contract method 0xe5cd8b6a.
//
// Solidity: function withdrawMax() returns()
func (_BondVault *BondVaultTransactorSession) WithdrawMax() (*types.Transaction, error) {
	return _BondVault.Contract.WithdrawMax(&_BondVault.TransactOpts)
}

// BondVaultDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the BondVault contract.
type BondVaultDepositedIterator struct {
	Event *BondVaultDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondVaultDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondVaultDeposited)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondVaultDeposited)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondVaultDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondVaultDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondVaultDeposited represents a Deposited event raised by the BondVault contract.
type BondVaultDeposited struct {
	User     common.Address
	Amount   *big.Int
	UnlockAt uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xcd50943caf3a035e3821448b2743faef90aae73ff20c6c1b29ea8213ad1fee1d.
//
// Solidity: event Deposited(address indexed user, uint256 amount, uint64 unlockAt)
func (_BondVault *BondVaultFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address) (*BondVaultDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondVault.contract.FilterLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return &BondVaultDepositedIterator{contract: _BondVault.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xcd50943caf3a035e3821448b2743faef90aae73ff20c6c1b29ea8213ad1fee1d.
//
// Solidity: event Deposited(address indexed user, uint256 amount, uint64 unlockAt)
func (_BondVault *BondVaultFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *BondVaultDeposited, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondVault.contract.WatchLogs(opts, "Deposited", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondVaultDeposited)
				if err := _BondVault.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xcd50943caf3a035e3821448b2743faef90aae73ff20c6c1b29ea8213ad1fee1d.
//
// Solidity: event Deposited(address indexed user, uint256 amount, uint64 unlockAt)
func (_BondVault *BondVaultFilterer) ParseDeposited(log types.Log) (*BondVaultDeposited, error) {
	event := new(BondVaultDeposited)
	if err := _BondVault.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondVaultSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the BondVault contract.
type BondVaultSlashedIterator struct {
	Event *BondVaultSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondVaultSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondVaultSlashed)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondVaultSlashed)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondVaultSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondVaultSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondVaultSlashed represents a Slashed event raised by the BondVault contract.
type BondVaultSlashed struct {
	User      common.Address
	Recipient common.Address
	Amount    *big.Int
	By        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x1b9817da702c2ae84a03f2a48072c6a81bfc3b7d1d5615bffce3db4856831bb7.
//
// Solidity: event Slashed(address indexed user, address indexed recipient, uint256 amount, address indexed by)
func (_BondVault *BondVaultFilterer) FilterSlashed(opts *bind.FilterOpts, user []common.Address, recipient []common.Address, by []common.Address) (*BondVaultSlashedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BondVault.contract.FilterLogs(opts, "Slashed", userRule, recipientRule, byRule)
	if err != nil {
		return nil, err
	}
	return &BondVaultSlashedIterator{contract: _BondVault.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x1b9817da702c2ae84a03f2a48072c6a81bfc3b7d1d5615bffce3db4856831bb7.
//
// Solidity: event Slashed(address indexed user, address indexed recipient, uint256 amount, address indexed by)
func (_BondVault *BondVaultFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *BondVaultSlashed, user []common.Address, recipient []common.Address, by []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BondVault.contract.WatchLogs(opts, "Slashed", userRule, recipientRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondVaultSlashed)
				if err := _BondVault.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x1b9817da702c2ae84a03f2a48072c6a81bfc3b7d1d5615bffce3db4856831bb7.
//
// Solidity: event Slashed(address indexed user, address indexed recipient, uint256 amount, address indexed by)
func (_BondVault *BondVaultFilterer) ParseSlashed(log types.Log) (*BondVaultSlashed, error) {
	event := new(BondVaultSlashed)
	if err := _BondVault.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondVaultWithdrawCanceledIterator is returned from FilterWithdrawCanceled and is used to iterate over the raw logs and unpacked data for WithdrawCanceled events raised by the BondVault contract.
type BondVaultWithdrawCanceledIterator struct {
	Event *BondVaultWithdrawCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondVaultWithdrawCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondVaultWithdrawCanceled)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondVaultWithdrawCanceled)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondVaultWithdrawCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondVaultWithdrawCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondVaultWithdrawCanceled represents a WithdrawCanceled event raised by the BondVault contract.
type BondVaultWithdrawCanceled struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCanceled is a free log retrieval operation binding the contract event 0xe4c65151920d34ba91eb49b98f4904c54654d4e5f450dc8387700c09c25a846a.
//
// Solidity: event WithdrawCanceled(address indexed user)
func (_BondVault *BondVaultFilterer) FilterWithdrawCanceled(opts *bind.FilterOpts, user []common.Address) (*BondVaultWithdrawCanceledIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondVault.contract.FilterLogs(opts, "WithdrawCanceled", userRule)
	if err != nil {
		return nil, err
	}
	return &BondVaultWithdrawCanceledIterator{contract: _BondVault.contract, event: "WithdrawCanceled", logs: logs, sub: sub}, nil
}

// WatchWithdrawCanceled is a free log subscription operation binding the contract event 0xe4c65151920d34ba91eb49b98f4904c54654d4e5f450dc8387700c09c25a846a.
//
// Solidity: event WithdrawCanceled(address indexed user)
func (_BondVault *BondVaultFilterer) WatchWithdrawCanceled(opts *bind.WatchOpts, sink chan<- *BondVaultWithdrawCanceled, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondVault.contract.WatchLogs(opts, "WithdrawCanceled", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondVaultWithdrawCanceled)
				if err := _BondVault.contract.UnpackLog(event, "WithdrawCanceled", log); err != nil {
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

// ParseWithdrawCanceled is a log parse operation binding the contract event 0xe4c65151920d34ba91eb49b98f4904c54654d4e5f450dc8387700c09c25a846a.
//
// Solidity: event WithdrawCanceled(address indexed user)
func (_BondVault *BondVaultFilterer) ParseWithdrawCanceled(log types.Log) (*BondVaultWithdrawCanceled, error) {
	event := new(BondVaultWithdrawCanceled)
	if err := _BondVault.contract.UnpackLog(event, "WithdrawCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondVaultWithdrawProcessedIterator is returned from FilterWithdrawProcessed and is used to iterate over the raw logs and unpacked data for WithdrawProcessed events raised by the BondVault contract.
type BondVaultWithdrawProcessedIterator struct {
	Event *BondVaultWithdrawProcessed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondVaultWithdrawProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondVaultWithdrawProcessed)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondVaultWithdrawProcessed)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondVaultWithdrawProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondVaultWithdrawProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondVaultWithdrawProcessed represents a WithdrawProcessed event raised by the BondVault contract.
type BondVaultWithdrawProcessed struct {
	User   common.Address
	Amount *big.Int
	Fee    *big.Int
	By     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawProcessed is a free log retrieval operation binding the contract event 0xd4297feff9aaad2ad18f51ed5932f7d8396d1071010eb5126cf627285b3efb1d.
//
// Solidity: event WithdrawProcessed(address indexed user, uint256 amount, uint256 fee, address indexed by)
func (_BondVault *BondVaultFilterer) FilterWithdrawProcessed(opts *bind.FilterOpts, user []common.Address, by []common.Address) (*BondVaultWithdrawProcessedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BondVault.contract.FilterLogs(opts, "WithdrawProcessed", userRule, byRule)
	if err != nil {
		return nil, err
	}
	return &BondVaultWithdrawProcessedIterator{contract: _BondVault.contract, event: "WithdrawProcessed", logs: logs, sub: sub}, nil
}

// WatchWithdrawProcessed is a free log subscription operation binding the contract event 0xd4297feff9aaad2ad18f51ed5932f7d8396d1071010eb5126cf627285b3efb1d.
//
// Solidity: event WithdrawProcessed(address indexed user, uint256 amount, uint256 fee, address indexed by)
func (_BondVault *BondVaultFilterer) WatchWithdrawProcessed(opts *bind.WatchOpts, sink chan<- *BondVaultWithdrawProcessed, user []common.Address, by []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _BondVault.contract.WatchLogs(opts, "WithdrawProcessed", userRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondVaultWithdrawProcessed)
				if err := _BondVault.contract.UnpackLog(event, "WithdrawProcessed", log); err != nil {
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

// ParseWithdrawProcessed is a log parse operation binding the contract event 0xd4297feff9aaad2ad18f51ed5932f7d8396d1071010eb5126cf627285b3efb1d.
//
// Solidity: event WithdrawProcessed(address indexed user, uint256 amount, uint256 fee, address indexed by)
func (_BondVault *BondVaultFilterer) ParseWithdrawProcessed(log types.Log) (*BondVaultWithdrawProcessed, error) {
	event := new(BondVaultWithdrawProcessed)
	if err := _BondVault.contract.UnpackLog(event, "WithdrawProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondVaultWithdrawRequestedIterator is returned from FilterWithdrawRequested and is used to iterate over the raw logs and unpacked data for WithdrawRequested events raised by the BondVault contract.
type BondVaultWithdrawRequestedIterator struct {
	Event *BondVaultWithdrawRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondVaultWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondVaultWithdrawRequested)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondVaultWithdrawRequested)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondVaultWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondVaultWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondVaultWithdrawRequested represents a WithdrawRequested event raised by the BondVault contract.
type BondVaultWithdrawRequested struct {
	User        common.Address
	Amount      *big.Int
	Fee         *big.Int
	RequestedAt uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequested is a free log retrieval operation binding the contract event 0x5b5919a27b29ed9ea8fdb44f39ca081298229768813e2b5ffbd9d83eeb1de61e.
//
// Solidity: event WithdrawRequested(address indexed user, uint256 amount, uint256 fee, uint64 requestedAt)
func (_BondVault *BondVaultFilterer) FilterWithdrawRequested(opts *bind.FilterOpts, user []common.Address) (*BondVaultWithdrawRequestedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondVault.contract.FilterLogs(opts, "WithdrawRequested", userRule)
	if err != nil {
		return nil, err
	}
	return &BondVaultWithdrawRequestedIterator{contract: _BondVault.contract, event: "WithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequested is a free log subscription operation binding the contract event 0x5b5919a27b29ed9ea8fdb44f39ca081298229768813e2b5ffbd9d83eeb1de61e.
//
// Solidity: event WithdrawRequested(address indexed user, uint256 amount, uint256 fee, uint64 requestedAt)
func (_BondVault *BondVaultFilterer) WatchWithdrawRequested(opts *bind.WatchOpts, sink chan<- *BondVaultWithdrawRequested, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondVault.contract.WatchLogs(opts, "WithdrawRequested", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondVaultWithdrawRequested)
				if err := _BondVault.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
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

// ParseWithdrawRequested is a log parse operation binding the contract event 0x5b5919a27b29ed9ea8fdb44f39ca081298229768813e2b5ffbd9d83eeb1de61e.
//
// Solidity: event WithdrawRequested(address indexed user, uint256 amount, uint256 fee, uint64 requestedAt)
func (_BondVault *BondVaultFilterer) ParseWithdrawRequested(log types.Log) (*BondVaultWithdrawRequested, error) {
	event := new(BondVaultWithdrawRequested)
	if err := _BondVault.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondVaultWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the BondVault contract.
type BondVaultWithdrawnIterator struct {
	Event *BondVaultWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondVaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondVaultWithdrawn)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondVaultWithdrawn)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondVaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondVaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondVaultWithdrawn represents a Withdrawn event raised by the BondVault contract.
type BondVaultWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 fee)
func (_BondVault *BondVaultFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*BondVaultWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondVault.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &BondVaultWithdrawnIterator{contract: _BondVault.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 fee)
func (_BondVault *BondVaultFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *BondVaultWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BondVault.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondVaultWithdrawn)
				if err := _BondVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 fee)
func (_BondVault *BondVaultFilterer) ParseWithdrawn(log types.Log) (*BondVaultWithdrawn, error) {
	event := new(BondVaultWithdrawn)
	if err := _BondVault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
