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

// SettlementEngineMetaData contains all meta data concerning the SettlementEngine contract.
var SettlementEngineMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_categories\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bondVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_profiles\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPacket\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"PreflightFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bondVault\",\"outputs\":[{\"internalType\":\"contractIBondVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"categories\",\"outputs\":[{\"internalType\":\"contractICategoryRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_merchantIdentity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_merchantPayout\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_categoryId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_createdAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_s\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profiles\",\"outputs\":[{\"internalType\":\"contractIProfileRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"}],\"name\":\"triggerSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_categories\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_profiles\",\"type\":\"address\"}],\"name\":\"updateRegistries\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIPaymentVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051611922380380611922833981810160405281019061003191906101bc565b60015f819055508473ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250505050505050610233565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61018b82610162565b9050919050565b61019b81610181565b81146101a5575f5ffd5b50565b5f815190506101b681610192565b92915050565b5f5f5f5f5f60a086880312156101d5576101d461015e565b5b5f6101e2888289016101a8565b95505060206101f3888289016101a8565b9450506040610204888289016101a8565b9350506060610215888289016101a8565b9250506080610226888289016101a8565b9150509295509295909350565b60805160a05160c0516116786102aa5f395f818161019b0152818161059d015281816105d90152818161077d01526107b901525f818161031f015281816104470152818161057901528181610a220152610b4a01525f81816104d0015281816108dd01528181610c400152610cd001526116785ff3fe608060405234801561000f575f5ffd5b5060043610610086575f3560e01c8063b712694911610059578063b712694914610100578063d3b7197b1461011c578063fbfa77cf14610138578063fdb330631461015657610086565b8063495057cf1461008a5780635aa6e675146100a857806367597af1146100c6578063990826b3146100e2575b5f5ffd5b610092610174565b60405161009f9190610dde565b60405180910390f35b6100b0610199565b6040516100bd9190610e17565b60405180910390f35b6100e060048036038101906100db9190610f12565b6101bd565b005b6100ea610577565b6040516100f79190610fe3565b60405180910390f35b61011a60048036038101906101159190610ffc565b61059b565b005b6101366004803603810190610131919061103a565b61077b565b005b610140610cce565b60405161014d9190611085565b60405180910390f35b61015e610cf2565b60405161016b91906110be565b60405180910390f35b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b6101c5610d17565b5f8282856040516020016101db9392919061112b565b60405160208183030381529060405280519060200120905084811461022c576040517fce09c71500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620151808461023b9190611194565b67ffffffffffffffff1642111561027e576040517f203d82d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d47875d08a6040518263ffffffff1660e01b81526004016102d991906111de565b602060405180830381865afa1580156102f4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610318919061120b565b036103a6577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f10bd31b898b6040518363ffffffff1660e01b8152600401610378929190611245565b5f604051808303815f87803b15801561038f575f5ffd5b505af11580156103a1573d5f5f3e3d5ffd5b505050505b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d47875d0336040518263ffffffff1660e01b815260040161040191906111de565b602060405180830381865afa15801561041c573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610440919061120b565b036104ce577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f10bd31b338b6040518363ffffffff1660e01b81526004016104a0929190611245565b5f604051808303815f87803b1580156104b7575f5ffd5b505af11580156104c9573d5f5f3e3d5ffd5b505050505b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16633f72ca41338b8b8b8b8b8b8b8b6040518a63ffffffff1660e01b81526004016105379998979695949392919061128a565b5f604051808303815f87803b15801561054e575f5ffd5b505af1158015610560573d5f5f3e3d5ffd5b505050505061056d610d5b565b5050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377fb8d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610640573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106649190611329565b336040518363ffffffff1660e01b8152600401610682929190611354565b602060405180830381865afa15801561069d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106c191906113b0565b6106f7576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634d104adf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610820573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108449190611329565b336040518363ffffffff1660e01b8152600401610862929190611354565b602060405180830381865afa15801561087d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108a191906113b0565b6108d7576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635fd13996866040518263ffffffff1660e01b815260040161093491906113db565b61012060405180830381865afa158015610950573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610974919061143f565b505050955095509550505093505f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d47875d0856040518263ffffffff1660e01b81526004016109dc91906111de565b602060405180830381865afa1580156109f7573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a1b919061120b565b03610aa9577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16638b4fb30d84836040518363ffffffff1660e01b8152600401610a7b929190611245565b5f604051808303815f87803b158015610a92575f5ffd5b505af1158015610aa4573d5f5f3e3d5ffd5b505050505b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d47875d0866040518263ffffffff1660e01b8152600401610b0491906111de565b602060405180830381865afa158015610b1f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b43919061120b565b03610bd1577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16638b4fb30d85836040518363ffffffff1660e01b8152600401610ba3929190611245565b5f604051808303815f87803b158015610bba575f5ffd5b505af1158015610bcc573d5f5f3e3d5ffd5b505050505b5f600167ffffffffffffffff811115610bed57610bec611503565b5b604051908082528060200260200182016040528015610c1b5781602001602082028036833780820191505090505b50905085815f81518110610c3257610c31611530565b5b6020026020010181815250507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477c5f0182856040518363ffffffff1660e01b8152600401610c99929190611614565b5f604051808303815f87803b158015610cb0575f5ffd5b505af1158015610cc2573d5f5f3e3d5ffd5b50505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f5403610d52576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f81905550565b60015f81905550565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610da6610da1610d9c84610d64565b610d83565b610d64565b9050919050565b5f610db782610d8c565b9050919050565b5f610dc882610dad565b9050919050565b610dd881610dbe565b82525050565b5f602082019050610df15f830184610dcf565b92915050565b5f610e0182610dad565b9050919050565b610e1181610df7565b82525050565b5f602082019050610e2a5f830184610e08565b92915050565b5f5ffd5b5f819050919050565b610e4681610e34565b8114610e50575f5ffd5b50565b5f81359050610e6181610e3d565b92915050565b5f610e7182610d64565b9050919050565b610e8181610e67565b8114610e8b575f5ffd5b50565b5f81359050610e9c81610e78565b92915050565b5f819050919050565b610eb481610ea2565b8114610ebe575f5ffd5b50565b5f81359050610ecf81610eab565b92915050565b5f67ffffffffffffffff82169050919050565b610ef181610ed5565b8114610efb575f5ffd5b50565b5f81359050610f0c81610ee8565b92915050565b5f5f5f5f5f5f5f5f610100898b031215610f2f57610f2e610e30565b5b5f610f3c8b828c01610e53565b9850506020610f4d8b828c01610e8e565b9750506040610f5e8b828c01610e8e565b9650506060610f6f8b828c01610ec1565b9550506080610f808b828c01610ec1565b94505060a0610f918b828c01610efe565b93505060c0610fa28b828c01610e53565b92505060e0610fb38b828c01610e53565b9150509295985092959890939650565b5f610fcd82610dad565b9050919050565b610fdd81610fc3565b82525050565b5f602082019050610ff65f830184610fd4565b92915050565b5f5f6040838503121561101257611011610e30565b5b5f61101f85828601610e8e565b925050602061103085828601610e8e565b9150509250929050565b5f6020828403121561104f5761104e610e30565b5b5f61105c84828501610ec1565b91505092915050565b5f61106f82610dad565b9050919050565b61107f81611065565b82525050565b5f6020820190506110985f830184611076565b92915050565b5f6110a882610dad565b9050919050565b6110b88161109e565b82525050565b5f6020820190506110d15f8301846110af565b92915050565b5f819050919050565b6110f16110ec82610e34565b6110d7565b82525050565b5f8160c01b9050919050565b5f61110d826110f7565b9050919050565b61112561112082610ed5565b611103565b82525050565b5f61113682866110e0565b60208201915061114682856110e0565b6020820191506111568284611114565b600882019150819050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61119e82610ed5565b91506111a983610ed5565b9250828201905067ffffffffffffffff8111156111c9576111c8611167565b5b92915050565b6111d881610e67565b82525050565b5f6020820190506111f15f8301846111cf565b92915050565b5f8151905061120581610e3d565b92915050565b5f602082840312156112205761121f610e30565b5b5f61122d848285016111f7565b91505092915050565b61123f81610e34565b82525050565b5f6040820190506112585f8301856111cf565b6112656020830184611236565b9392505050565b61127581610ea2565b82525050565b61128481610ed5565b82525050565b5f6101208201905061129e5f83018c6111cf565b6112ab602083018b611236565b6112b8604083018a6111cf565b6112c560608301896111cf565b6112d2608083018861126c565b6112df60a083018761126c565b6112ec60c083018661127b565b6112f960e0830185611236565b611307610100830184611236565b9a9950505050505050505050565b5f8151905061132381610eab565b92915050565b5f6020828403121561133e5761133d610e30565b5b5f61134b84828501611315565b91505092915050565b5f6040820190506113675f83018561126c565b61137460208301846111cf565b9392505050565b5f8115159050919050565b61138f8161137b565b8114611399575f5ffd5b50565b5f815190506113aa81611386565b92915050565b5f602082840312156113c5576113c4610e30565b5b5f6113d28482850161139c565b91505092915050565b5f6020820190506113ee5f83018461126c565b92915050565b5f8151905061140281610e78565b92915050565b5f8151905061141681610ee8565b92915050565b60048110611428575f5ffd5b50565b5f815190506114398161141c565b92915050565b5f5f5f5f5f5f5f5f5f6101208a8c03121561145d5761145c610e30565b5b5f61146a8c828d016113f4565b995050602061147b8c828d01611408565b985050604061148c8c828d0161142b565b975050606061149d8c828d016113f4565b96505060806114ae8c828d016113f4565b95505060a06114bf8c828d016111f7565b94505060c06114d08c828d01611315565b93505060e06114e18c828d016111f7565b9250506101006114f38c828d016111f7565b9150509295985092959850929598565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61158f81610ea2565b82525050565b5f6115a08383611586565b60208301905092915050565b5f602082019050919050565b5f6115c28261155d565b6115cc8185611567565b93506115d783611577565b805f5b838110156116075781516115ee8882611595565b97506115f9836115ac565b9250506001810190506115da565b5085935050505092915050565b5f6040820190508181035f83015261162c81856115b8565b905061163b60208301846111cf565b939250505056fea2646970667358221220e22760dff8f6d244409cbf6b5c6c06456354eab6f90ec0994f472c7173714c0d64736f6c63430008210033",
}

// SettlementEngineABI is the input ABI used to generate the binding from.
// Deprecated: Use SettlementEngineMetaData.ABI instead.
var SettlementEngineABI = SettlementEngineMetaData.ABI

// SettlementEngineBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SettlementEngineMetaData.Bin instead.
var SettlementEngineBin = SettlementEngineMetaData.Bin

// DeploySettlementEngine deploys a new Ethereum contract, binding an instance of SettlementEngine to it.
func DeploySettlementEngine(auth *bind.TransactOpts, backend bind.ContractBackend, _vault common.Address, _categories common.Address, _bondVault common.Address, _profiles common.Address, _gov common.Address) (common.Address, *types.Transaction, *SettlementEngine, error) {
	parsed, err := SettlementEngineMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SettlementEngineBin), backend, _vault, _categories, _bondVault, _profiles, _gov)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SettlementEngine{SettlementEngineCaller: SettlementEngineCaller{contract: contract}, SettlementEngineTransactor: SettlementEngineTransactor{contract: contract}, SettlementEngineFilterer: SettlementEngineFilterer{contract: contract}}, nil
}

// SettlementEngine is an auto generated Go binding around an Ethereum contract.
type SettlementEngine struct {
	SettlementEngineCaller     // Read-only binding to the contract
	SettlementEngineTransactor // Write-only binding to the contract
	SettlementEngineFilterer   // Log filterer for contract events
}

// SettlementEngineCaller is an auto generated read-only Go binding around an Ethereum contract.
type SettlementEngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementEngineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SettlementEngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementEngineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SettlementEngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementEngineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SettlementEngineSession struct {
	Contract     *SettlementEngine // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SettlementEngineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SettlementEngineCallerSession struct {
	Contract *SettlementEngineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SettlementEngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SettlementEngineTransactorSession struct {
	Contract     *SettlementEngineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SettlementEngineRaw is an auto generated low-level Go binding around an Ethereum contract.
type SettlementEngineRaw struct {
	Contract *SettlementEngine // Generic contract binding to access the raw methods on
}

// SettlementEngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SettlementEngineCallerRaw struct {
	Contract *SettlementEngineCaller // Generic read-only contract binding to access the raw methods on
}

// SettlementEngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SettlementEngineTransactorRaw struct {
	Contract *SettlementEngineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSettlementEngine creates a new instance of SettlementEngine, bound to a specific deployed contract.
func NewSettlementEngine(address common.Address, backend bind.ContractBackend) (*SettlementEngine, error) {
	contract, err := bindSettlementEngine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SettlementEngine{SettlementEngineCaller: SettlementEngineCaller{contract: contract}, SettlementEngineTransactor: SettlementEngineTransactor{contract: contract}, SettlementEngineFilterer: SettlementEngineFilterer{contract: contract}}, nil
}

// NewSettlementEngineCaller creates a new read-only instance of SettlementEngine, bound to a specific deployed contract.
func NewSettlementEngineCaller(address common.Address, caller bind.ContractCaller) (*SettlementEngineCaller, error) {
	contract, err := bindSettlementEngine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementEngineCaller{contract: contract}, nil
}

// NewSettlementEngineTransactor creates a new write-only instance of SettlementEngine, bound to a specific deployed contract.
func NewSettlementEngineTransactor(address common.Address, transactor bind.ContractTransactor) (*SettlementEngineTransactor, error) {
	contract, err := bindSettlementEngine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementEngineTransactor{contract: contract}, nil
}

// NewSettlementEngineFilterer creates a new log filterer instance of SettlementEngine, bound to a specific deployed contract.
func NewSettlementEngineFilterer(address common.Address, filterer bind.ContractFilterer) (*SettlementEngineFilterer, error) {
	contract, err := bindSettlementEngine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SettlementEngineFilterer{contract: contract}, nil
}

// bindSettlementEngine binds a generic wrapper to an already deployed contract.
func bindSettlementEngine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SettlementEngineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SettlementEngine *SettlementEngineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SettlementEngine.Contract.SettlementEngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SettlementEngine *SettlementEngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SettlementEngine.Contract.SettlementEngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SettlementEngine *SettlementEngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SettlementEngine.Contract.SettlementEngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SettlementEngine *SettlementEngineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SettlementEngine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SettlementEngine *SettlementEngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SettlementEngine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SettlementEngine *SettlementEngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SettlementEngine.Contract.contract.Transact(opts, method, params...)
}

// BondVault is a free data retrieval call binding the contract method 0x990826b3.
//
// Solidity: function bondVault() view returns(address)
func (_SettlementEngine *SettlementEngineCaller) BondVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SettlementEngine.contract.Call(opts, &out, "bondVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondVault is a free data retrieval call binding the contract method 0x990826b3.
//
// Solidity: function bondVault() view returns(address)
func (_SettlementEngine *SettlementEngineSession) BondVault() (common.Address, error) {
	return _SettlementEngine.Contract.BondVault(&_SettlementEngine.CallOpts)
}

// BondVault is a free data retrieval call binding the contract method 0x990826b3.
//
// Solidity: function bondVault() view returns(address)
func (_SettlementEngine *SettlementEngineCallerSession) BondVault() (common.Address, error) {
	return _SettlementEngine.Contract.BondVault(&_SettlementEngine.CallOpts)
}

// Categories is a free data retrieval call binding the contract method 0xfdb33063.
//
// Solidity: function categories() view returns(address)
func (_SettlementEngine *SettlementEngineCaller) Categories(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SettlementEngine.contract.Call(opts, &out, "categories")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Categories is a free data retrieval call binding the contract method 0xfdb33063.
//
// Solidity: function categories() view returns(address)
func (_SettlementEngine *SettlementEngineSession) Categories() (common.Address, error) {
	return _SettlementEngine.Contract.Categories(&_SettlementEngine.CallOpts)
}

// Categories is a free data retrieval call binding the contract method 0xfdb33063.
//
// Solidity: function categories() view returns(address)
func (_SettlementEngine *SettlementEngineCallerSession) Categories() (common.Address, error) {
	return _SettlementEngine.Contract.Categories(&_SettlementEngine.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_SettlementEngine *SettlementEngineCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SettlementEngine.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_SettlementEngine *SettlementEngineSession) Governance() (common.Address, error) {
	return _SettlementEngine.Contract.Governance(&_SettlementEngine.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_SettlementEngine *SettlementEngineCallerSession) Governance() (common.Address, error) {
	return _SettlementEngine.Contract.Governance(&_SettlementEngine.CallOpts)
}

// Profiles is a free data retrieval call binding the contract method 0x495057cf.
//
// Solidity: function profiles() view returns(address)
func (_SettlementEngine *SettlementEngineCaller) Profiles(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SettlementEngine.contract.Call(opts, &out, "profiles")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Profiles is a free data retrieval call binding the contract method 0x495057cf.
//
// Solidity: function profiles() view returns(address)
func (_SettlementEngine *SettlementEngineSession) Profiles() (common.Address, error) {
	return _SettlementEngine.Contract.Profiles(&_SettlementEngine.CallOpts)
}

// Profiles is a free data retrieval call binding the contract method 0x495057cf.
//
// Solidity: function profiles() view returns(address)
func (_SettlementEngine *SettlementEngineCallerSession) Profiles() (common.Address, error) {
	return _SettlementEngine.Contract.Profiles(&_SettlementEngine.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_SettlementEngine *SettlementEngineCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SettlementEngine.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_SettlementEngine *SettlementEngineSession) Vault() (common.Address, error) {
	return _SettlementEngine.Contract.Vault(&_SettlementEngine.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_SettlementEngine *SettlementEngineCallerSession) Vault() (common.Address, error) {
	return _SettlementEngine.Contract.Vault(&_SettlementEngine.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x67597af1.
//
// Solidity: function deposit(uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId, uint64 _createdAt, uint256 _r, uint256 _s) returns()
func (_SettlementEngine *SettlementEngineTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte, _createdAt uint64, _r *big.Int, _s *big.Int) (*types.Transaction, error) {
	return _SettlementEngine.contract.Transact(opts, "deposit", _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId, _createdAt, _r, _s)
}

// Deposit is a paid mutator transaction binding the contract method 0x67597af1.
//
// Solidity: function deposit(uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId, uint64 _createdAt, uint256 _r, uint256 _s) returns()
func (_SettlementEngine *SettlementEngineSession) Deposit(_amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte, _createdAt uint64, _r *big.Int, _s *big.Int) (*types.Transaction, error) {
	return _SettlementEngine.Contract.Deposit(&_SettlementEngine.TransactOpts, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId, _createdAt, _r, _s)
}

// Deposit is a paid mutator transaction binding the contract method 0x67597af1.
//
// Solidity: function deposit(uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId, uint64 _createdAt, uint256 _r, uint256 _s) returns()
func (_SettlementEngine *SettlementEngineTransactorSession) Deposit(_amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte, _createdAt uint64, _r *big.Int, _s *big.Int) (*types.Transaction, error) {
	return _SettlementEngine.Contract.Deposit(&_SettlementEngine.TransactOpts, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId, _createdAt, _r, _s)
}

// TriggerSettlement is a paid mutator transaction binding the contract method 0xd3b7197b.
//
// Solidity: function triggerSettlement(bytes32 _packetId) returns()
func (_SettlementEngine *SettlementEngineTransactor) TriggerSettlement(opts *bind.TransactOpts, _packetId [32]byte) (*types.Transaction, error) {
	return _SettlementEngine.contract.Transact(opts, "triggerSettlement", _packetId)
}

// TriggerSettlement is a paid mutator transaction binding the contract method 0xd3b7197b.
//
// Solidity: function triggerSettlement(bytes32 _packetId) returns()
func (_SettlementEngine *SettlementEngineSession) TriggerSettlement(_packetId [32]byte) (*types.Transaction, error) {
	return _SettlementEngine.Contract.TriggerSettlement(&_SettlementEngine.TransactOpts, _packetId)
}

// TriggerSettlement is a paid mutator transaction binding the contract method 0xd3b7197b.
//
// Solidity: function triggerSettlement(bytes32 _packetId) returns()
func (_SettlementEngine *SettlementEngineTransactorSession) TriggerSettlement(_packetId [32]byte) (*types.Transaction, error) {
	return _SettlementEngine.Contract.TriggerSettlement(&_SettlementEngine.TransactOpts, _packetId)
}

// UpdateRegistries is a paid mutator transaction binding the contract method 0xb7126949.
//
// Solidity: function updateRegistries(address _categories, address _profiles) returns()
func (_SettlementEngine *SettlementEngineTransactor) UpdateRegistries(opts *bind.TransactOpts, _categories common.Address, _profiles common.Address) (*types.Transaction, error) {
	return _SettlementEngine.contract.Transact(opts, "updateRegistries", _categories, _profiles)
}

// UpdateRegistries is a paid mutator transaction binding the contract method 0xb7126949.
//
// Solidity: function updateRegistries(address _categories, address _profiles) returns()
func (_SettlementEngine *SettlementEngineSession) UpdateRegistries(_categories common.Address, _profiles common.Address) (*types.Transaction, error) {
	return _SettlementEngine.Contract.UpdateRegistries(&_SettlementEngine.TransactOpts, _categories, _profiles)
}

// UpdateRegistries is a paid mutator transaction binding the contract method 0xb7126949.
//
// Solidity: function updateRegistries(address _categories, address _profiles) returns()
func (_SettlementEngine *SettlementEngineTransactorSession) UpdateRegistries(_categories common.Address, _profiles common.Address) (*types.Transaction, error) {
	return _SettlementEngine.Contract.UpdateRegistries(&_SettlementEngine.TransactOpts, _categories, _profiles)
}
