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

// GovernanceMetaData contains all meta data concerning the Governance contract.
var GovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_governanceAdmins\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_minSoleAdminAgeSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotRemoveLastGovernanceAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyGovernanceAdmins\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"provided\",\"type\":\"uint64\"}],\"name\":\"MinAgeTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"addedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nowTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minAge\",\"type\":\"uint64\"}],\"name\":\"RemainingAdminTooNew\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ARBITER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_SOLE_ADMIN_AGE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arbiter\",\"type\":\"address\"}],\"name\":\"addArbiter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addGovernanceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"getAdminAddedAt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governanceAdminAddedAt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"arbiter\",\"type\":\"address\"}],\"name\":\"removeArbiter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeGovernanceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051611fa3380380611fa383398181016040528101906100319190610762565b5f82510361006b576040517f5ddbb04700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f8267ffffffffffffffff16146100835781610088565b62093a805b90506301e1338067ffffffffffffffff168167ffffffffffffffff1611156100e757806040517fafaf86290000000000000000000000000000000000000000000000000000000081526004016100de91906107cb565b60405180910390fd5b8067ffffffffffffffff1660808167ffffffffffffffff16815250506101155f5f1b3361026f60201b60201c565b506101467fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e9727806102b860201b60201c565b6101967fbb08418a67729a078f87bbc8d02a770929bb68f5bfdf134ae2ead6ed38e2f4ae7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e97276102b860201b60201c565b5f5f90505b8351811015610266575f8482815181106101b8576101b76107e4565b5b602002602001015190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610227576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102577fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e97278261026f60201b60201c565b5050808060010191505061019b565b50505050610811565b5f5f610281848461031660201b60201c565b905080156102ae576102ac8360015f8781526020019081526020015f2061040b60201b90919060201c565b505b8091505092915050565b5f6102c88361043e60201b60201c565b9050815f5f8581526020019081526020015f20600101819055508181847fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff60405160405180910390a4505050565b5f610327838361045a60201b60201c565b6104015760015f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555061039e61047360201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050610405565b5f90505b92915050565b5f610436835f018373ffffffffffffffffffffffffffffffffffffffff165f1b61047a60201b60201c565b905092915050565b5f5f5f8381526020019081526020015f20600101549050919050565b5f61046b83836104e760201b60201c565b905092915050565b5f33905090565b5f61048b838361054a60201b60201c565b6104dd57825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f2081905550600190506104e1565b5f90505b92915050565b5f5f5f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f5f836001015f8481526020019081526020015f20541415905092915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6105c58261057f565b810181811067ffffffffffffffff821117156105e4576105e361058f565b5b80604052505050565b5f6105f661056a565b905061060282826105bc565b919050565b5f67ffffffffffffffff8211156106215761062061058f565b5b602082029050602081019050919050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61065f82610636565b9050919050565b61066f81610655565b8114610679575f5ffd5b50565b5f8151905061068a81610666565b92915050565b5f6106a261069d84610607565b6105ed565b905080838252602082019050602084028301858111156106c5576106c4610632565b5b835b818110156106ee57806106da888261067c565b8452602084019350506020810190506106c7565b5050509392505050565b5f82601f83011261070c5761070b61057b565b5b815161071c848260208601610690565b91505092915050565b5f67ffffffffffffffff82169050919050565b61074181610725565b811461074b575f5ffd5b50565b5f8151905061075c81610738565b92915050565b5f5f6040838503121561077857610777610573565b5b5f83015167ffffffffffffffff81111561079557610794610577565b5b6107a1858286016106f8565b92505060206107b28582860161074e565b9150509250929050565b6107c581610725565b82525050565b5f6020820190506107de5f8301846107bc565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b60805161176c6108375f395f81816108460152818161088f0152610937015261176c5ff3fe608060405234801561000f575f5ffd5b506004361061011f575f3560e01c80639010d07c116100ab578063a3246ad31161006f578063a3246ad314610317578063b538d3bc14610347578063ca15c87314610363578063d547741f14610393578063ef60fa42146103af5761011f565b80639010d07c1461025f5780639053978c1461028f57806391d14854146102ab5780639222e175146102db578063a217fddf146102f95761011f565b80633487e08c116100f25780633487e08c146101cf57806336568abe146101eb5780634d104adf1461020757806377fb8d7a14610225578063840cab22146102435761011f565b806301ffc9a71461012357806308dd0fec14610153578063248a9ca3146101835780632f2ff15d146101b3575b5f5ffd5b61013d6004803603810190610138919061126b565b6103df565b60405161014a91906112b0565b60405180910390f35b61016d60048036038101906101689190611323565b6103f0565b60405161017a9190611370565b60405180910390f35b61019d600480360381019061019891906113bc565b610414565b6040516101aa91906113f6565b60405180910390f35b6101cd60048036038101906101c8919061140f565b610430565b005b6101e960048036038101906101e49190611323565b610452565b005b6102056004803603810190610200919061140f565b6104ab565b005b61020f610526565b60405161021c91906113f6565b60405180910390f35b61022d61054a565b60405161023a91906113f6565b60405180910390f35b61025d60048036038101906102589190611323565b61056e565b005b61027960048036038101906102749190611480565b61068f565b60405161028691906114cd565b60405180910390f35b6102a960048036038101906102a49190611323565b6106bb565b005b6102c560048036038101906102c0919061140f565b610922565b6040516102d291906112b0565b60405180910390f35b6102e3610935565b6040516102f09190611370565b60405180910390f35b610301610959565b60405161030e91906113f6565b60405180910390f35b610331600480360381019061032c91906113bc565b61095f565b60405161033e919061159d565b60405180910390f35b610361600480360381019061035c9190611323565b610981565b005b61037d600480360381019061037891906113bc565b610a3f565b60405161038a91906115cc565b60405180910390f35b6103ad60048036038101906103a8919061140f565b610a60565b005b6103c960048036038101906103c49190611323565b610a82565b6040516103d69190611370565b60405180910390f35b5f6103e982610adb565b9050919050565b6002602052805f5260405f205f915054906101000a900467ffffffffffffffff1681565b5f5f5f8381526020019081526020015f20600101549050919050565b61043982610414565b61044281610b54565b61044c8383610b68565b50505050565b7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e972761047c81610b54565b6104a67fbb08418a67729a078f87bbc8d02a770929bb68f5bfdf134ae2ead6ed38e2f4ae83610bab565b505050565b6104b3610bee565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610517576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105218282610bab565b505050565b7fbb08418a67729a078f87bbc8d02a770929bb68f5bfdf134ae2ead6ed38e2f4ae81565b7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e972781565b7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e972761059881610b54565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105fd576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106277fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e972783610b68565b504260025f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b5f6106b38260015f8681526020019081526020015f20610bf590919063ffffffff16565b905092915050565b7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e97276106e581610b54565b5f61070f7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e9727610a3f565b90506001811161074b576040517fd107f79000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600281036108f2575f61077e7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e97275f61068f565b90505f6107ac7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e9727600161068f565b90505f8273ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16146107e857826107ea565b815b90505f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff1690505f4290507f000000000000000000000000000000000000000000000000000000000000000067ffffffffffffffff16828261087b9190611612565b67ffffffffffffffff1610156108ec5781817f00000000000000000000000000000000000000000000000000000000000000006040517ff6209d840000000000000000000000000000000000000000000000000000000081526004016108e39392919061164d565b60405180910390fd5b50505050505b61091c7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e972784610bab565b50505050565b5f61092d8383610c0c565b905092915050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f1b81565b606061097a60015f8481526020019081526020015f20610c6f565b9050919050565b7fe7e0b301de1c1e268073a6b56ff0876cffca0becba0dca08f945c1394e3e97276109ab81610b54565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a10576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a3a7fbb08418a67729a078f87bbc8d02a770929bb68f5bfdf134ae2ead6ed38e2f4ae83610b68565b505050565b5f610a5960015f8481526020019081526020015f20610c8e565b9050919050565b610a6982610414565b610a7281610b54565b610a7c8383610bab565b50505050565b5f60025f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900467ffffffffffffffff169050919050565b5f7f5a05180f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610b4d5750610b4c82610ca1565b5b9050919050565b610b6581610b60610bee565b610d1a565b50565b5f5f610b748484610d6b565b90508015610ba157610b9f8360015f8781526020019081526020015f20610e5490919063ffffffff16565b505b8091505092915050565b5f5f610bb78484610e81565b90508015610be457610be28360015f8781526020019081526020015f20610f6a90919063ffffffff16565b505b8091505092915050565b5f33905090565b5f610c02835f0183610f97565b5f1c905092915050565b5f5f5f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b60605f610c7d835f01610fbe565b905060608190508092505050919050565b5f610c9a825f01611017565b9050919050565b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610d135750610d1282611026565b5b9050919050565b610d248282610922565b610d675780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401610d5e929190611682565b60405180910390fd5b5050565b5f610d768383610922565b610e4a5760015f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610de7610bee565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050610e4e565b5f90505b92915050565b5f610e79835f018373ffffffffffffffffffffffffffffffffffffffff165f1b61108f565b905092915050565b5f610e8c8383610922565b15610f60575f5f5f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610efd610bee565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a460019050610f64565b5f90505b92915050565b5f610f8f835f018373ffffffffffffffffffffffffffffffffffffffff165f1b6110f6565b905092915050565b5f825f018281548110610fad57610fac6116a9565b5b905f5260205f200154905092915050565b6060815f0180548060200260200160405190810160405280929190818152602001828054801561100b57602002820191905f5260205f20905b815481526020019060010190808311610ff7575b50505050509050919050565b5f815f01805490509050919050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f61109a83836111f2565b6110ec57825f0182908060018154018082558091505060019003905f5260205f20015f9091909190915055825f0180549050836001015f8481526020019081526020015f2081905550600190506110f0565b5f90505b92915050565b5f5f836001015f8481526020019081526020015f205490505f81146111e7575f60018261112391906116d6565b90505f6001865f018054905061113991906116d6565b905080821461119f575f865f018281548110611158576111576116a9565b5b905f5260205f200154905080875f018481548110611179576111786116a9565b5b905f5260205f20018190555083876001015f8381526020019081526020015f2081905550505b855f018054806111b2576111b1611709565b5b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506111ec565b5f9150505b92915050565b5f5f836001015f8481526020019081526020015f20541415905092915050565b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61124a81611216565b8114611254575f5ffd5b50565b5f8135905061126581611241565b92915050565b5f602082840312156112805761127f611212565b5b5f61128d84828501611257565b91505092915050565b5f8115159050919050565b6112aa81611296565b82525050565b5f6020820190506112c35f8301846112a1565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112f2826112c9565b9050919050565b611302816112e8565b811461130c575f5ffd5b50565b5f8135905061131d816112f9565b92915050565b5f6020828403121561133857611337611212565b5b5f6113458482850161130f565b91505092915050565b5f67ffffffffffffffff82169050919050565b61136a8161134e565b82525050565b5f6020820190506113835f830184611361565b92915050565b5f819050919050565b61139b81611389565b81146113a5575f5ffd5b50565b5f813590506113b681611392565b92915050565b5f602082840312156113d1576113d0611212565b5b5f6113de848285016113a8565b91505092915050565b6113f081611389565b82525050565b5f6020820190506114095f8301846113e7565b92915050565b5f5f6040838503121561142557611424611212565b5b5f611432858286016113a8565b92505060206114438582860161130f565b9150509250929050565b5f819050919050565b61145f8161144d565b8114611469575f5ffd5b50565b5f8135905061147a81611456565b92915050565b5f5f6040838503121561149657611495611212565b5b5f6114a3858286016113a8565b92505060206114b48582860161146c565b9150509250929050565b6114c7816112e8565b82525050565b5f6020820190506114e05f8301846114be565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611518816112e8565b82525050565b5f611529838361150f565b60208301905092915050565b5f602082019050919050565b5f61154b826114e6565b61155581856114f0565b935061156083611500565b805f5b83811015611590578151611577888261151e565b975061158283611535565b925050600181019050611563565b5085935050505092915050565b5f6020820190508181035f8301526115b58184611541565b905092915050565b6115c68161144d565b82525050565b5f6020820190506115df5f8301846115bd565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61161c8261134e565b91506116278361134e565b9250828203905067ffffffffffffffff811115611647576116466115e5565b5b92915050565b5f6060820190506116605f830186611361565b61166d6020830185611361565b61167a6040830184611361565b949350505050565b5f6040820190506116955f8301856114be565b6116a260208301846113e7565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6116e08261144d565b91506116eb8361144d565b9250828203905081811115611703576117026115e5565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffdfea26469706673582212204eb97a25f4d2ca6c55f1acf5ea3c94baceb1505ecf3a6392f40ae0e58482dbf364736f6c63430008210033",
}

// GovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceMetaData.ABI instead.
var GovernanceABI = GovernanceMetaData.ABI

// GovernanceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernanceMetaData.Bin instead.
var GovernanceBin = GovernanceMetaData.Bin

// DeployGovernance deploys a new Ethereum contract, binding an instance of Governance to it.
func DeployGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _governanceAdmins []common.Address, _minSoleAdminAgeSeconds uint64) (common.Address, *types.Transaction, *Governance, error) {
	parsed, err := GovernanceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceBin), backend, _governanceAdmins, _minSoleAdminAgeSeconds)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// ARBITERROLE is a free data retrieval call binding the contract method 0x4d104adf.
//
// Solidity: function ARBITER_ROLE() view returns(bytes32)
func (_Governance *GovernanceCaller) ARBITERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "ARBITER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ARBITERROLE is a free data retrieval call binding the contract method 0x4d104adf.
//
// Solidity: function ARBITER_ROLE() view returns(bytes32)
func (_Governance *GovernanceSession) ARBITERROLE() ([32]byte, error) {
	return _Governance.Contract.ARBITERROLE(&_Governance.CallOpts)
}

// ARBITERROLE is a free data retrieval call binding the contract method 0x4d104adf.
//
// Solidity: function ARBITER_ROLE() view returns(bytes32)
func (_Governance *GovernanceCallerSession) ARBITERROLE() ([32]byte, error) {
	return _Governance.Contract.ARBITERROLE(&_Governance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Governance *GovernanceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Governance *GovernanceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Governance.Contract.DEFAULTADMINROLE(&_Governance.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Governance *GovernanceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Governance.Contract.DEFAULTADMINROLE(&_Governance.CallOpts)
}

// GOVERNANCEADMINROLE is a free data retrieval call binding the contract method 0x77fb8d7a.
//
// Solidity: function GOVERNANCE_ADMIN_ROLE() view returns(bytes32)
func (_Governance *GovernanceCaller) GOVERNANCEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "GOVERNANCE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEADMINROLE is a free data retrieval call binding the contract method 0x77fb8d7a.
//
// Solidity: function GOVERNANCE_ADMIN_ROLE() view returns(bytes32)
func (_Governance *GovernanceSession) GOVERNANCEADMINROLE() ([32]byte, error) {
	return _Governance.Contract.GOVERNANCEADMINROLE(&_Governance.CallOpts)
}

// GOVERNANCEADMINROLE is a free data retrieval call binding the contract method 0x77fb8d7a.
//
// Solidity: function GOVERNANCE_ADMIN_ROLE() view returns(bytes32)
func (_Governance *GovernanceCallerSession) GOVERNANCEADMINROLE() ([32]byte, error) {
	return _Governance.Contract.GOVERNANCEADMINROLE(&_Governance.CallOpts)
}

// MINSOLEADMINAGE is a free data retrieval call binding the contract method 0x9222e175.
//
// Solidity: function MIN_SOLE_ADMIN_AGE() view returns(uint64)
func (_Governance *GovernanceCaller) MINSOLEADMINAGE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "MIN_SOLE_ADMIN_AGE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MINSOLEADMINAGE is a free data retrieval call binding the contract method 0x9222e175.
//
// Solidity: function MIN_SOLE_ADMIN_AGE() view returns(uint64)
func (_Governance *GovernanceSession) MINSOLEADMINAGE() (uint64, error) {
	return _Governance.Contract.MINSOLEADMINAGE(&_Governance.CallOpts)
}

// MINSOLEADMINAGE is a free data retrieval call binding the contract method 0x9222e175.
//
// Solidity: function MIN_SOLE_ADMIN_AGE() view returns(uint64)
func (_Governance *GovernanceCallerSession) MINSOLEADMINAGE() (uint64, error) {
	return _Governance.Contract.MINSOLEADMINAGE(&_Governance.CallOpts)
}

// GetAdminAddedAt is a free data retrieval call binding the contract method 0xef60fa42.
//
// Solidity: function getAdminAddedAt(address _admin) view returns(uint64)
func (_Governance *GovernanceCaller) GetAdminAddedAt(opts *bind.CallOpts, _admin common.Address) (uint64, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getAdminAddedAt", _admin)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAdminAddedAt is a free data retrieval call binding the contract method 0xef60fa42.
//
// Solidity: function getAdminAddedAt(address _admin) view returns(uint64)
func (_Governance *GovernanceSession) GetAdminAddedAt(_admin common.Address) (uint64, error) {
	return _Governance.Contract.GetAdminAddedAt(&_Governance.CallOpts, _admin)
}

// GetAdminAddedAt is a free data retrieval call binding the contract method 0xef60fa42.
//
// Solidity: function getAdminAddedAt(address _admin) view returns(uint64)
func (_Governance *GovernanceCallerSession) GetAdminAddedAt(_admin common.Address) (uint64, error) {
	return _Governance.Contract.GetAdminAddedAt(&_Governance.CallOpts, _admin)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Governance *GovernanceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Governance *GovernanceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Governance.Contract.GetRoleAdmin(&_Governance.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Governance *GovernanceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Governance.Contract.GetRoleAdmin(&_Governance.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Governance *GovernanceCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Governance *GovernanceSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Governance.Contract.GetRoleMember(&_Governance.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Governance *GovernanceCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Governance.Contract.GetRoleMember(&_Governance.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Governance *GovernanceCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Governance *GovernanceSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Governance.Contract.GetRoleMemberCount(&_Governance.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Governance.Contract.GetRoleMemberCount(&_Governance.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Governance *GovernanceCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Governance *GovernanceSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _Governance.Contract.GetRoleMembers(&_Governance.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Governance *GovernanceCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _Governance.Contract.GetRoleMembers(&_Governance.CallOpts, role)
}

// GovernanceAdminAddedAt is a free data retrieval call binding the contract method 0x08dd0fec.
//
// Solidity: function governanceAdminAddedAt(address ) view returns(uint64)
func (_Governance *GovernanceCaller) GovernanceAdminAddedAt(opts *bind.CallOpts, arg0 common.Address) (uint64, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "governanceAdminAddedAt", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GovernanceAdminAddedAt is a free data retrieval call binding the contract method 0x08dd0fec.
//
// Solidity: function governanceAdminAddedAt(address ) view returns(uint64)
func (_Governance *GovernanceSession) GovernanceAdminAddedAt(arg0 common.Address) (uint64, error) {
	return _Governance.Contract.GovernanceAdminAddedAt(&_Governance.CallOpts, arg0)
}

// GovernanceAdminAddedAt is a free data retrieval call binding the contract method 0x08dd0fec.
//
// Solidity: function governanceAdminAddedAt(address ) view returns(uint64)
func (_Governance *GovernanceCallerSession) GovernanceAdminAddedAt(arg0 common.Address) (uint64, error) {
	return _Governance.Contract.GovernanceAdminAddedAt(&_Governance.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 _role, address _account) view returns(bool)
func (_Governance *GovernanceCaller) HasRole(opts *bind.CallOpts, _role [32]byte, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "hasRole", _role, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 _role, address _account) view returns(bool)
func (_Governance *GovernanceSession) HasRole(_role [32]byte, _account common.Address) (bool, error) {
	return _Governance.Contract.HasRole(&_Governance.CallOpts, _role, _account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 _role, address _account) view returns(bool)
func (_Governance *GovernanceCallerSession) HasRole(_role [32]byte, _account common.Address) (bool, error) {
	return _Governance.Contract.HasRole(&_Governance.CallOpts, _role, _account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Governance *GovernanceCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Governance *GovernanceSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Governance.Contract.SupportsInterface(&_Governance.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_Governance *GovernanceCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Governance.Contract.SupportsInterface(&_Governance.CallOpts, _interfaceId)
}

// AddArbiter is a paid mutator transaction binding the contract method 0xb538d3bc.
//
// Solidity: function addArbiter(address arbiter) returns()
func (_Governance *GovernanceTransactor) AddArbiter(opts *bind.TransactOpts, arbiter common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addArbiter", arbiter)
}

// AddArbiter is a paid mutator transaction binding the contract method 0xb538d3bc.
//
// Solidity: function addArbiter(address arbiter) returns()
func (_Governance *GovernanceSession) AddArbiter(arbiter common.Address) (*types.Transaction, error) {
	return _Governance.Contract.AddArbiter(&_Governance.TransactOpts, arbiter)
}

// AddArbiter is a paid mutator transaction binding the contract method 0xb538d3bc.
//
// Solidity: function addArbiter(address arbiter) returns()
func (_Governance *GovernanceTransactorSession) AddArbiter(arbiter common.Address) (*types.Transaction, error) {
	return _Governance.Contract.AddArbiter(&_Governance.TransactOpts, arbiter)
}

// AddGovernanceAdmin is a paid mutator transaction binding the contract method 0x840cab22.
//
// Solidity: function addGovernanceAdmin(address admin) returns()
func (_Governance *GovernanceTransactor) AddGovernanceAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "addGovernanceAdmin", admin)
}

// AddGovernanceAdmin is a paid mutator transaction binding the contract method 0x840cab22.
//
// Solidity: function addGovernanceAdmin(address admin) returns()
func (_Governance *GovernanceSession) AddGovernanceAdmin(admin common.Address) (*types.Transaction, error) {
	return _Governance.Contract.AddGovernanceAdmin(&_Governance.TransactOpts, admin)
}

// AddGovernanceAdmin is a paid mutator transaction binding the contract method 0x840cab22.
//
// Solidity: function addGovernanceAdmin(address admin) returns()
func (_Governance *GovernanceTransactorSession) AddGovernanceAdmin(admin common.Address) (*types.Transaction, error) {
	return _Governance.Contract.AddGovernanceAdmin(&_Governance.TransactOpts, admin)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Governance *GovernanceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Governance *GovernanceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GrantRole(&_Governance.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Governance *GovernanceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GrantRole(&_Governance.TransactOpts, role, account)
}

// RemoveArbiter is a paid mutator transaction binding the contract method 0x3487e08c.
//
// Solidity: function removeArbiter(address arbiter) returns()
func (_Governance *GovernanceTransactor) RemoveArbiter(opts *bind.TransactOpts, arbiter common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "removeArbiter", arbiter)
}

// RemoveArbiter is a paid mutator transaction binding the contract method 0x3487e08c.
//
// Solidity: function removeArbiter(address arbiter) returns()
func (_Governance *GovernanceSession) RemoveArbiter(arbiter common.Address) (*types.Transaction, error) {
	return _Governance.Contract.RemoveArbiter(&_Governance.TransactOpts, arbiter)
}

// RemoveArbiter is a paid mutator transaction binding the contract method 0x3487e08c.
//
// Solidity: function removeArbiter(address arbiter) returns()
func (_Governance *GovernanceTransactorSession) RemoveArbiter(arbiter common.Address) (*types.Transaction, error) {
	return _Governance.Contract.RemoveArbiter(&_Governance.TransactOpts, arbiter)
}

// RemoveGovernanceAdmin is a paid mutator transaction binding the contract method 0x9053978c.
//
// Solidity: function removeGovernanceAdmin(address admin) returns()
func (_Governance *GovernanceTransactor) RemoveGovernanceAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "removeGovernanceAdmin", admin)
}

// RemoveGovernanceAdmin is a paid mutator transaction binding the contract method 0x9053978c.
//
// Solidity: function removeGovernanceAdmin(address admin) returns()
func (_Governance *GovernanceSession) RemoveGovernanceAdmin(admin common.Address) (*types.Transaction, error) {
	return _Governance.Contract.RemoveGovernanceAdmin(&_Governance.TransactOpts, admin)
}

// RemoveGovernanceAdmin is a paid mutator transaction binding the contract method 0x9053978c.
//
// Solidity: function removeGovernanceAdmin(address admin) returns()
func (_Governance *GovernanceTransactorSession) RemoveGovernanceAdmin(admin common.Address) (*types.Transaction, error) {
	return _Governance.Contract.RemoveGovernanceAdmin(&_Governance.TransactOpts, admin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Governance *GovernanceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Governance *GovernanceSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Governance.Contract.RenounceRole(&_Governance.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Governance *GovernanceTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Governance.Contract.RenounceRole(&_Governance.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Governance *GovernanceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Governance *GovernanceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Governance.Contract.RevokeRole(&_Governance.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Governance *GovernanceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Governance.Contract.RevokeRole(&_Governance.TransactOpts, role, account)
}

// GovernanceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Governance contract.
type GovernanceRoleAdminChangedIterator struct {
	Event *GovernanceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GovernanceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceRoleAdminChanged)
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
		it.Event = new(GovernanceRoleAdminChanged)
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
func (it *GovernanceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceRoleAdminChanged represents a RoleAdminChanged event raised by the Governance contract.
type GovernanceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Governance *GovernanceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GovernanceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Governance.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceRoleAdminChangedIterator{contract: _Governance.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Governance *GovernanceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GovernanceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Governance.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceRoleAdminChanged)
				if err := _Governance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Governance *GovernanceFilterer) ParseRoleAdminChanged(log types.Log) (*GovernanceRoleAdminChanged, error) {
	event := new(GovernanceRoleAdminChanged)
	if err := _Governance.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Governance contract.
type GovernanceRoleGrantedIterator struct {
	Event *GovernanceRoleGranted // Event containing the contract specifics and raw log

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
func (it *GovernanceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceRoleGranted)
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
		it.Event = new(GovernanceRoleGranted)
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
func (it *GovernanceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceRoleGranted represents a RoleGranted event raised by the Governance contract.
type GovernanceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Governance *GovernanceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GovernanceRoleGrantedIterator, error) {

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

	logs, sub, err := _Governance.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceRoleGrantedIterator{contract: _Governance.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Governance *GovernanceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GovernanceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Governance.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceRoleGranted)
				if err := _Governance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Governance *GovernanceFilterer) ParseRoleGranted(log types.Log) (*GovernanceRoleGranted, error) {
	event := new(GovernanceRoleGranted)
	if err := _Governance.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Governance contract.
type GovernanceRoleRevokedIterator struct {
	Event *GovernanceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GovernanceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceRoleRevoked)
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
		it.Event = new(GovernanceRoleRevoked)
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
func (it *GovernanceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceRoleRevoked represents a RoleRevoked event raised by the Governance contract.
type GovernanceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Governance *GovernanceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GovernanceRoleRevokedIterator, error) {

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

	logs, sub, err := _Governance.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceRoleRevokedIterator{contract: _Governance.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Governance *GovernanceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GovernanceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Governance.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceRoleRevoked)
				if err := _Governance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Governance *GovernanceFilterer) ParseRoleRevoked(log types.Log) (*GovernanceRoleRevoked, error) {
	event := new(GovernanceRoleRevoked)
	if err := _Governance.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
