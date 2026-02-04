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

// ICategoryRegistryCategoryRule is an auto generated low-level Go binding around an user-defined struct.
type ICategoryRegistryCategoryRule struct {
	Existence     bool
	Enabled       bool
	MultiplierBps uint16
	MinHoldSecs   uint32
	MaxHoldSecs   uint32
	BufferSecs    uint32
	FormulaParam  uint32
}

// CategoryRegistryMetaData contains all meta data concerning the CategoryRegistry contract.
var CategoryRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadRule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CategoryDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedArrays\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"multiplier\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"buffer\",\"type\":\"uint32\"}],\"name\":\"RuleUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_ids\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"multiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minHoldSecs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxHoldSecs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"bufferSecs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"formulaParam\",\"type\":\"uint32\"}],\"internalType\":\"structICategoryRegistry.CategoryRule[]\",\"name\":\"_newRules\",\"type\":\"tuple[]\"}],\"name\":\"batchUpsertRules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_expectedSecs\",\"type\":\"uint32\"}],\"name\":\"calculateExecutionWindow\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"categoryList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getRule\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"multiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minHoldSecs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxHoldSecs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"bufferSecs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"formulaParam\",\"type\":\"uint32\"}],\"internalType\":\"structICategoryRegistry.CategoryRule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"multiplierBps\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minHoldSecs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxHoldSecs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"bufferSecs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"formulaParam\",\"type\":\"uint32\"}],\"internalType\":\"structICategoryRegistry.CategoryRule\",\"name\":\"_rule\",\"type\":\"tuple\"}],\"name\":\"upsertRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051611777380380611777833981810160405281019061003191906100c9565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506100f4565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100988261006f565b9050919050565b6100a88161008e565b81146100b2575f5ffd5b50565b5f815190506100c38161009f565b92915050565b5f602082840312156100de576100dd61006b565b5b5f6100eb848285016100b5565b91505092915050565b60805161164f6101285f395f818161028d015281816102c901528181610593015281816105b701526105f3015261164f5ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80630c74109214610064578063549c3571146100945780635aa6e675146100b05780636bdeac8e146100ce578063c21260a5146100ea578063f7cab8471461011a575b5f5ffd5b61007e60048036038101906100799190610b28565b61014a565b60405161008b9190610b75565b60405180910390f35b6100ae60048036038101906100a99190610c44565b61028b565b005b6100b8610591565b6040516100c59190610d3c565b60405180910390f35b6100e860048036038101906100e39190610d77565b6105b5565b005b61010460048036038101906100ff9190610de9565b610821565b6040516101119190610e23565b60405180910390f35b610134600480360381019061012f9190610e3c565b610841565b6040516101419190610f38565b60405180910390f35b5f5f5f5f8581526020019081526020015f209050805f0160019054906101000a900460ff166101a5576040517f5c46569700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f612710825f0160029054906101000a900461ffff1661ffff168563ffffffff166101d09190610f7e565b6101da9190610fec565b9050815f01600c9054906101000a900463ffffffff1663ffffffff1681610201919061101c565b9050815f0160049054906101000a900463ffffffff1663ffffffff1681101561024157815f0160049054906101000a900463ffffffff1692505050610285565b815f0160089054906101000a900463ffffffff1663ffffffff1681111561027f57815f0160089054906101000a900463ffffffff1692505050610285565b80925050505b92915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377fb8d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610330573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103549190611063565b336040518363ffffffff1660e01b81526004016103729291906110ae565b602060405180830381865afa15801561038d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103b191906110ff565b6103e7576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f84849050905082829050811461042a576040517fa121188700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f90505b81811015610589575f86868381811061044b5761044a61112a565b5b905060200201359050368585848181106104685761046761112a565b5b905060e0020190505f5f1b82036104ab576040517fac8fb3c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104b481610951565b5f5f8381526020019081526020015f205f015f9054906101000a900460ff166104fe57600182908060018154018082558091505060019003905f5260205f20015f90919091909150555b805f5f8481526020019081526020015f20818161051b919061156b565b905050817f6dcb02bda6c9267939741fcbb5f280a19d124b22033b609cd60d982979a6f9d0826040016020810190610553919061158d565b8360a001602081019061056691906115b8565b6040516105749291906115f2565b60405180910390a2826001019250505061042f565b505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377fb8d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561065a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061067e9190611063565b336040518363ffffffff1660e01b815260040161069c9291906110ae565b602060405180830381865afa1580156106b7573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106db91906110ff565b610711576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f1b820361074c576040517fac8fb3c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61075581610951565b5f5f8381526020019081526020015f205f015f9054906101000a900460ff1661079f57600182908060018154018082558091505060019003905f5260205f20015f90919091909150555b805f5f8481526020019081526020015f2081816107bc919061156b565b905050817f6dcb02bda6c9267939741fcbb5f280a19d124b22033b609cd60d982979a6f9d08260400160208101906107f4919061158d565b8360a001602081019061080791906115b8565b6040516108159291906115f2565b60405180910390a25050565b60018181548110610830575f80fd5b905f5260205f20015f915090505481565b610849610a5e565b5f5f8381526020019081526020015f206040518060e00160405290815f82015f9054906101000a900460ff161515151581526020015f820160019054906101000a900460ff161515151581526020015f820160029054906101000a900461ffff1661ffff1661ffff1681526020015f820160049054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020015f820160089054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020015f8201600c9054906101000a900463ffffffff1663ffffffff1663ffffffff1681526020015f820160109054906101000a900463ffffffff1663ffffffff1663ffffffff16815250509050919050565b5f816040016020810190610965919061158d565b61ffff16036109a0576040517fca7c57d400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8160600160208101906109b491906115b8565b63ffffffff16036109f1576040517fca7c57d400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806060016020810190610a0491906115b8565b63ffffffff16816080016020810190610a1d91906115b8565b63ffffffff161015610a5b576040517fca7c57d400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b6040518060e001604052805f151581526020015f151581526020015f61ffff1681526020015f63ffffffff1681526020015f63ffffffff1681526020015f63ffffffff1681526020015f63ffffffff1681525090565b5f5ffd5b5f5ffd5b5f819050919050565b610ace81610abc565b8114610ad8575f5ffd5b50565b5f81359050610ae981610ac5565b92915050565b5f63ffffffff82169050919050565b610b0781610aef565b8114610b11575f5ffd5b50565b5f81359050610b2281610afe565b92915050565b5f5f60408385031215610b3e57610b3d610ab4565b5b5f610b4b85828601610adb565b9250506020610b5c85828601610b14565b9150509250929050565b610b6f81610aef565b82525050565b5f602082019050610b885f830184610b66565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610baf57610bae610b8e565b5b8235905067ffffffffffffffff811115610bcc57610bcb610b92565b5b602083019150836020820283011115610be857610be7610b96565b5b9250929050565b5f5f83601f840112610c0457610c03610b8e565b5b8235905067ffffffffffffffff811115610c2157610c20610b92565b5b6020830191508360e0820283011115610c3d57610c3c610b96565b5b9250929050565b5f5f5f5f60408587031215610c5c57610c5b610ab4565b5b5f85013567ffffffffffffffff811115610c7957610c78610ab8565b5b610c8587828801610b9a565b9450945050602085013567ffffffffffffffff811115610ca857610ca7610ab8565b5b610cb487828801610bef565b925092505092959194509250565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610d04610cff610cfa84610cc2565b610ce1565b610cc2565b9050919050565b5f610d1582610cea565b9050919050565b5f610d2682610d0b565b9050919050565b610d3681610d1c565b82525050565b5f602082019050610d4f5f830184610d2d565b92915050565b5f5ffd5b5f60e08284031215610d6e57610d6d610d55565b5b81905092915050565b5f5f6101008385031215610d8e57610d8d610ab4565b5b5f610d9b85828601610adb565b9250506020610dac85828601610d59565b9150509250929050565b5f819050919050565b610dc881610db6565b8114610dd2575f5ffd5b50565b5f81359050610de381610dbf565b92915050565b5f60208284031215610dfe57610dfd610ab4565b5b5f610e0b84828501610dd5565b91505092915050565b610e1d81610abc565b82525050565b5f602082019050610e365f830184610e14565b92915050565b5f60208284031215610e5157610e50610ab4565b5b5f610e5e84828501610adb565b91505092915050565b5f8115159050919050565b610e7b81610e67565b82525050565b5f61ffff82169050919050565b610e9781610e81565b82525050565b610ea681610aef565b82525050565b60e082015f820151610ec05f850182610e72565b506020820151610ed36020850182610e72565b506040820151610ee66040850182610e8e565b506060820151610ef96060850182610e9d565b506080820151610f0c6080850182610e9d565b5060a0820151610f1f60a0850182610e9d565b5060c0820151610f3260c0850182610e9d565b50505050565b5f60e082019050610f4b5f830184610eac565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f8882610db6565b9150610f9383610db6565b9250828202610fa181610db6565b91508282048414831517610fb857610fb7610f51565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610ff682610db6565b915061100183610db6565b92508261101157611010610fbf565b5b828204905092915050565b5f61102682610db6565b915061103183610db6565b925082820190508082111561104957611048610f51565b5b92915050565b5f8151905061105d81610ac5565b92915050565b5f6020828403121561107857611077610ab4565b5b5f6110858482850161104f565b91505092915050565b5f61109882610cc2565b9050919050565b6110a88161108e565b82525050565b5f6040820190506110c15f830185610e14565b6110ce602083018461109f565b9392505050565b6110de81610e67565b81146110e8575f5ffd5b50565b5f815190506110f9816110d5565b92915050565b5f6020828403121561111457611113610ab4565b5b5f611121848285016110eb565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f8135611163816110d5565b80915050919050565b5f815f1b9050919050565b5f60ff6111838461116c565b9350801983169250808416831791505092915050565b5f6111a382610e67565b9050919050565b5f819050919050565b6111bc82611199565b6111cf6111c8826111aa565b8354611177565b8255505050565b5f8160081b9050919050565b5f61ff006111ef846111d6565b9350801983169250808416831791505092915050565b61120e82611199565b61122161121a826111aa565b83546111e2565b8255505050565b61123181610e81565b811461123b575f5ffd5b50565b5f813561124a81611228565b80915050919050565b5f8160101b9050919050565b5f63ffff000061126e84611253565b9350801983169250808416831791505092915050565b5f61129e61129961129484610e81565b610ce1565b610e81565b9050919050565b5f819050919050565b6112b782611284565b6112ca6112c3826112a5565b835461125f565b8255505050565b5f81356112dd81610afe565b80915050919050565b5f8160201b9050919050565b5f67ffffffff00000000611305846112e6565b9350801983169250808416831791505092915050565b5f61133561133061132b84610aef565b610ce1565b610aef565b9050919050565b5f819050919050565b61134e8261131b565b61136161135a8261133c565b83546112f2565b8255505050565b5f8160401b9050919050565b5f6bffffffff000000000000000061138b84611368565b9350801983169250808416831791505092915050565b6113aa8261131b565b6113bd6113b68261133c565b8354611374565b8255505050565b5f8160601b9050919050565b5f6fffffffff0000000000000000000000006113eb846113c4565b9350801983169250808416831791505092915050565b61140a8261131b565b61141d6114168261133c565b83546113d0565b8255505050565b5f8160801b9050919050565b5f73ffffffff0000000000000000000000000000000061144f84611424565b9350801983169250808416831791505092915050565b61146e8261131b565b61148161147a8261133c565b8354611430565b8255505050565b5f81015f83018061149881611157565b90506114a481846111b3565b5050505f810160208301806114b881611157565b90506114c48184611205565b5050505f810160408301806114d88161123e565b90506114e481846112ae565b5050505f810160608301806114f8816112d1565b90506115048184611345565b5050505f81016080830180611518816112d1565b905061152481846113a1565b5050505f810160a0830180611538816112d1565b90506115448184611401565b5050505f810160c0830180611558816112d1565b90506115648184611465565b5050505050565b6115758282611488565b5050565b5f8135905061158781611228565b92915050565b5f602082840312156115a2576115a1610ab4565b5b5f6115af84828501611579565b91505092915050565b5f602082840312156115cd576115cc610ab4565b5b5f6115da84828501610b14565b91505092915050565b6115ec81610e81565b82525050565b5f6040820190506116055f8301856115e3565b6116126020830184610b66565b939250505056fea26469706673582212204b6e2f8dc3cda055592751164efbb47da77dc4fa901761d6a85d2c4c7c8b8d5764736f6c63430008210033",
}

// CategoryRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CategoryRegistryMetaData.ABI instead.
var CategoryRegistryABI = CategoryRegistryMetaData.ABI

// CategoryRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CategoryRegistryMetaData.Bin instead.
var CategoryRegistryBin = CategoryRegistryMetaData.Bin

// DeployCategoryRegistry deploys a new Ethereum contract, binding an instance of CategoryRegistry to it.
func DeployCategoryRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _gov common.Address) (common.Address, *types.Transaction, *CategoryRegistry, error) {
	parsed, err := CategoryRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CategoryRegistryBin), backend, _gov)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CategoryRegistry{CategoryRegistryCaller: CategoryRegistryCaller{contract: contract}, CategoryRegistryTransactor: CategoryRegistryTransactor{contract: contract}, CategoryRegistryFilterer: CategoryRegistryFilterer{contract: contract}}, nil
}

// CategoryRegistry is an auto generated Go binding around an Ethereum contract.
type CategoryRegistry struct {
	CategoryRegistryCaller     // Read-only binding to the contract
	CategoryRegistryTransactor // Write-only binding to the contract
	CategoryRegistryFilterer   // Log filterer for contract events
}

// CategoryRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CategoryRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CategoryRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CategoryRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CategoryRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CategoryRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CategoryRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CategoryRegistrySession struct {
	Contract     *CategoryRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CategoryRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CategoryRegistryCallerSession struct {
	Contract *CategoryRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CategoryRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CategoryRegistryTransactorSession struct {
	Contract     *CategoryRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CategoryRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CategoryRegistryRaw struct {
	Contract *CategoryRegistry // Generic contract binding to access the raw methods on
}

// CategoryRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CategoryRegistryCallerRaw struct {
	Contract *CategoryRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CategoryRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CategoryRegistryTransactorRaw struct {
	Contract *CategoryRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCategoryRegistry creates a new instance of CategoryRegistry, bound to a specific deployed contract.
func NewCategoryRegistry(address common.Address, backend bind.ContractBackend) (*CategoryRegistry, error) {
	contract, err := bindCategoryRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CategoryRegistry{CategoryRegistryCaller: CategoryRegistryCaller{contract: contract}, CategoryRegistryTransactor: CategoryRegistryTransactor{contract: contract}, CategoryRegistryFilterer: CategoryRegistryFilterer{contract: contract}}, nil
}

// NewCategoryRegistryCaller creates a new read-only instance of CategoryRegistry, bound to a specific deployed contract.
func NewCategoryRegistryCaller(address common.Address, caller bind.ContractCaller) (*CategoryRegistryCaller, error) {
	contract, err := bindCategoryRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CategoryRegistryCaller{contract: contract}, nil
}

// NewCategoryRegistryTransactor creates a new write-only instance of CategoryRegistry, bound to a specific deployed contract.
func NewCategoryRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CategoryRegistryTransactor, error) {
	contract, err := bindCategoryRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CategoryRegistryTransactor{contract: contract}, nil
}

// NewCategoryRegistryFilterer creates a new log filterer instance of CategoryRegistry, bound to a specific deployed contract.
func NewCategoryRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CategoryRegistryFilterer, error) {
	contract, err := bindCategoryRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CategoryRegistryFilterer{contract: contract}, nil
}

// bindCategoryRegistry binds a generic wrapper to an already deployed contract.
func bindCategoryRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CategoryRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CategoryRegistry *CategoryRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CategoryRegistry.Contract.CategoryRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CategoryRegistry *CategoryRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CategoryRegistry.Contract.CategoryRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CategoryRegistry *CategoryRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CategoryRegistry.Contract.CategoryRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CategoryRegistry *CategoryRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CategoryRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CategoryRegistry *CategoryRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CategoryRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CategoryRegistry *CategoryRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CategoryRegistry.Contract.contract.Transact(opts, method, params...)
}

// CalculateExecutionWindow is a free data retrieval call binding the contract method 0x0c741092.
//
// Solidity: function calculateExecutionWindow(bytes32 _id, uint32 _expectedSecs) view returns(uint32)
func (_CategoryRegistry *CategoryRegistryCaller) CalculateExecutionWindow(opts *bind.CallOpts, _id [32]byte, _expectedSecs uint32) (uint32, error) {
	var out []interface{}
	err := _CategoryRegistry.contract.Call(opts, &out, "calculateExecutionWindow", _id, _expectedSecs)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CalculateExecutionWindow is a free data retrieval call binding the contract method 0x0c741092.
//
// Solidity: function calculateExecutionWindow(bytes32 _id, uint32 _expectedSecs) view returns(uint32)
func (_CategoryRegistry *CategoryRegistrySession) CalculateExecutionWindow(_id [32]byte, _expectedSecs uint32) (uint32, error) {
	return _CategoryRegistry.Contract.CalculateExecutionWindow(&_CategoryRegistry.CallOpts, _id, _expectedSecs)
}

// CalculateExecutionWindow is a free data retrieval call binding the contract method 0x0c741092.
//
// Solidity: function calculateExecutionWindow(bytes32 _id, uint32 _expectedSecs) view returns(uint32)
func (_CategoryRegistry *CategoryRegistryCallerSession) CalculateExecutionWindow(_id [32]byte, _expectedSecs uint32) (uint32, error) {
	return _CategoryRegistry.Contract.CalculateExecutionWindow(&_CategoryRegistry.CallOpts, _id, _expectedSecs)
}

// CategoryList is a free data retrieval call binding the contract method 0xc21260a5.
//
// Solidity: function categoryList(uint256 ) view returns(bytes32)
func (_CategoryRegistry *CategoryRegistryCaller) CategoryList(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CategoryRegistry.contract.Call(opts, &out, "categoryList", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CategoryList is a free data retrieval call binding the contract method 0xc21260a5.
//
// Solidity: function categoryList(uint256 ) view returns(bytes32)
func (_CategoryRegistry *CategoryRegistrySession) CategoryList(arg0 *big.Int) ([32]byte, error) {
	return _CategoryRegistry.Contract.CategoryList(&_CategoryRegistry.CallOpts, arg0)
}

// CategoryList is a free data retrieval call binding the contract method 0xc21260a5.
//
// Solidity: function categoryList(uint256 ) view returns(bytes32)
func (_CategoryRegistry *CategoryRegistryCallerSession) CategoryList(arg0 *big.Int) ([32]byte, error) {
	return _CategoryRegistry.Contract.CategoryList(&_CategoryRegistry.CallOpts, arg0)
}

// GetRule is a free data retrieval call binding the contract method 0xf7cab847.
//
// Solidity: function getRule(bytes32 _id) view returns((bool,bool,uint16,uint32,uint32,uint32,uint32))
func (_CategoryRegistry *CategoryRegistryCaller) GetRule(opts *bind.CallOpts, _id [32]byte) (ICategoryRegistryCategoryRule, error) {
	var out []interface{}
	err := _CategoryRegistry.contract.Call(opts, &out, "getRule", _id)

	if err != nil {
		return *new(ICategoryRegistryCategoryRule), err
	}

	out0 := *abi.ConvertType(out[0], new(ICategoryRegistryCategoryRule)).(*ICategoryRegistryCategoryRule)

	return out0, err

}

// GetRule is a free data retrieval call binding the contract method 0xf7cab847.
//
// Solidity: function getRule(bytes32 _id) view returns((bool,bool,uint16,uint32,uint32,uint32,uint32))
func (_CategoryRegistry *CategoryRegistrySession) GetRule(_id [32]byte) (ICategoryRegistryCategoryRule, error) {
	return _CategoryRegistry.Contract.GetRule(&_CategoryRegistry.CallOpts, _id)
}

// GetRule is a free data retrieval call binding the contract method 0xf7cab847.
//
// Solidity: function getRule(bytes32 _id) view returns((bool,bool,uint16,uint32,uint32,uint32,uint32))
func (_CategoryRegistry *CategoryRegistryCallerSession) GetRule(_id [32]byte) (ICategoryRegistryCategoryRule, error) {
	return _CategoryRegistry.Contract.GetRule(&_CategoryRegistry.CallOpts, _id)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_CategoryRegistry *CategoryRegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CategoryRegistry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_CategoryRegistry *CategoryRegistrySession) Governance() (common.Address, error) {
	return _CategoryRegistry.Contract.Governance(&_CategoryRegistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_CategoryRegistry *CategoryRegistryCallerSession) Governance() (common.Address, error) {
	return _CategoryRegistry.Contract.Governance(&_CategoryRegistry.CallOpts)
}

// BatchUpsertRules is a paid mutator transaction binding the contract method 0x549c3571.
//
// Solidity: function batchUpsertRules(bytes32[] _ids, (bool,bool,uint16,uint32,uint32,uint32,uint32)[] _newRules) returns()
func (_CategoryRegistry *CategoryRegistryTransactor) BatchUpsertRules(opts *bind.TransactOpts, _ids [][32]byte, _newRules []ICategoryRegistryCategoryRule) (*types.Transaction, error) {
	return _CategoryRegistry.contract.Transact(opts, "batchUpsertRules", _ids, _newRules)
}

// BatchUpsertRules is a paid mutator transaction binding the contract method 0x549c3571.
//
// Solidity: function batchUpsertRules(bytes32[] _ids, (bool,bool,uint16,uint32,uint32,uint32,uint32)[] _newRules) returns()
func (_CategoryRegistry *CategoryRegistrySession) BatchUpsertRules(_ids [][32]byte, _newRules []ICategoryRegistryCategoryRule) (*types.Transaction, error) {
	return _CategoryRegistry.Contract.BatchUpsertRules(&_CategoryRegistry.TransactOpts, _ids, _newRules)
}

// BatchUpsertRules is a paid mutator transaction binding the contract method 0x549c3571.
//
// Solidity: function batchUpsertRules(bytes32[] _ids, (bool,bool,uint16,uint32,uint32,uint32,uint32)[] _newRules) returns()
func (_CategoryRegistry *CategoryRegistryTransactorSession) BatchUpsertRules(_ids [][32]byte, _newRules []ICategoryRegistryCategoryRule) (*types.Transaction, error) {
	return _CategoryRegistry.Contract.BatchUpsertRules(&_CategoryRegistry.TransactOpts, _ids, _newRules)
}

// UpsertRule is a paid mutator transaction binding the contract method 0x6bdeac8e.
//
// Solidity: function upsertRule(bytes32 _id, (bool,bool,uint16,uint32,uint32,uint32,uint32) _rule) returns()
func (_CategoryRegistry *CategoryRegistryTransactor) UpsertRule(opts *bind.TransactOpts, _id [32]byte, _rule ICategoryRegistryCategoryRule) (*types.Transaction, error) {
	return _CategoryRegistry.contract.Transact(opts, "upsertRule", _id, _rule)
}

// UpsertRule is a paid mutator transaction binding the contract method 0x6bdeac8e.
//
// Solidity: function upsertRule(bytes32 _id, (bool,bool,uint16,uint32,uint32,uint32,uint32) _rule) returns()
func (_CategoryRegistry *CategoryRegistrySession) UpsertRule(_id [32]byte, _rule ICategoryRegistryCategoryRule) (*types.Transaction, error) {
	return _CategoryRegistry.Contract.UpsertRule(&_CategoryRegistry.TransactOpts, _id, _rule)
}

// UpsertRule is a paid mutator transaction binding the contract method 0x6bdeac8e.
//
// Solidity: function upsertRule(bytes32 _id, (bool,bool,uint16,uint32,uint32,uint32,uint32) _rule) returns()
func (_CategoryRegistry *CategoryRegistryTransactorSession) UpsertRule(_id [32]byte, _rule ICategoryRegistryCategoryRule) (*types.Transaction, error) {
	return _CategoryRegistry.Contract.UpsertRule(&_CategoryRegistry.TransactOpts, _id, _rule)
}

// CategoryRegistryRuleUpdatedIterator is returned from FilterRuleUpdated and is used to iterate over the raw logs and unpacked data for RuleUpdated events raised by the CategoryRegistry contract.
type CategoryRegistryRuleUpdatedIterator struct {
	Event *CategoryRegistryRuleUpdated // Event containing the contract specifics and raw log

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
func (it *CategoryRegistryRuleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CategoryRegistryRuleUpdated)
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
		it.Event = new(CategoryRegistryRuleUpdated)
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
func (it *CategoryRegistryRuleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CategoryRegistryRuleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CategoryRegistryRuleUpdated represents a RuleUpdated event raised by the CategoryRegistry contract.
type CategoryRegistryRuleUpdated struct {
	Id         [32]byte
	Multiplier uint16
	Buffer     uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRuleUpdated is a free log retrieval operation binding the contract event 0x6dcb02bda6c9267939741fcbb5f280a19d124b22033b609cd60d982979a6f9d0.
//
// Solidity: event RuleUpdated(bytes32 indexed id, uint16 multiplier, uint32 buffer)
func (_CategoryRegistry *CategoryRegistryFilterer) FilterRuleUpdated(opts *bind.FilterOpts, id [][32]byte) (*CategoryRegistryRuleUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CategoryRegistry.contract.FilterLogs(opts, "RuleUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &CategoryRegistryRuleUpdatedIterator{contract: _CategoryRegistry.contract, event: "RuleUpdated", logs: logs, sub: sub}, nil
}

// WatchRuleUpdated is a free log subscription operation binding the contract event 0x6dcb02bda6c9267939741fcbb5f280a19d124b22033b609cd60d982979a6f9d0.
//
// Solidity: event RuleUpdated(bytes32 indexed id, uint16 multiplier, uint32 buffer)
func (_CategoryRegistry *CategoryRegistryFilterer) WatchRuleUpdated(opts *bind.WatchOpts, sink chan<- *CategoryRegistryRuleUpdated, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CategoryRegistry.contract.WatchLogs(opts, "RuleUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CategoryRegistryRuleUpdated)
				if err := _CategoryRegistry.contract.UnpackLog(event, "RuleUpdated", log); err != nil {
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

// ParseRuleUpdated is a log parse operation binding the contract event 0x6dcb02bda6c9267939741fcbb5f280a19d124b22033b609cd60d982979a6f9d0.
//
// Solidity: event RuleUpdated(bytes32 indexed id, uint16 multiplier, uint32 buffer)
func (_CategoryRegistry *CategoryRegistryFilterer) ParseRuleUpdated(log types.Log) (*CategoryRegistryRuleUpdated, error) {
	event := new(CategoryRegistryRuleUpdated)
	if err := _CategoryRegistry.contract.UnpackLog(event, "RuleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
