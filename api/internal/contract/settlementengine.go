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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_categories\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bondVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_profiles\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"PreflightFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"bondVault\",\"outputs\":[{\"internalType\":\"contractIBondVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"categories\",\"outputs\":[{\"internalType\":\"contractICategoryRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_merchantIdentity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_merchantPayout\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_categoryId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profiles\",\"outputs\":[{\"internalType\":\"contractIProfileRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"}],\"name\":\"triggerSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_categories\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_profiles\",\"type\":\"address\"}],\"name\":\"updateRegistries\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIPaymentVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f5ffd5b506040516116b73803806116b7833981810160405281019061003191906101bc565b60015f819055508473ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508273ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508160025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff1660c08173ffffffffffffffffffffffffffffffffffffffff16815250505050505050610233565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61018b82610162565b9050919050565b61019b81610181565b81146101a5575f5ffd5b50565b5f815190506101b681610192565b92915050565b5f5f5f5f5f60a086880312156101d5576101d461015e565b5b5f6101e2888289016101a8565b95505060206101f3888289016101a8565b9450506040610204888289016101a8565b9350506060610215888289016101a8565b9250506080610226888289016101a8565b9150509295509295909350565b60805160a05160c05161140d6102aa5f395f818161019b015281816101e30152818161021f015281816106ba01526106f601525f81816101bf0152818161046a015281816105920152818161095e0152610a8601525f818161061b0152818161081b01528181610b7c0152610c0d015261140d5ff3fe608060405234801561000f575f5ffd5b5060043610610086575f3560e01c8063bc1ca92d11610059578063bc1ca92d14610100578063d3b7197b1461011c578063fbfa77cf14610138578063fdb330631461015657610086565b8063495057cf1461008a5780635aa6e675146100a8578063990826b3146100c6578063b7126949146100e4575b5f5ffd5b610092610174565b60405161009f9190610d1b565b60405180910390f35b6100b0610199565b6040516100bd9190610d54565b60405180910390f35b6100ce6101bd565b6040516100db9190610d8d565b60405180910390f35b6100fe60048036038101906100f99190610de5565b6101e1565b005b61011a60048036038101906101159190610e89565b6103c1565b005b61013660048036038101906101319190610f00565b6106b8565b005b610140610c0b565b60405161014d9190610f4b565b60405180910390f35b61015e610c2f565b60405161016b9190610f84565b60405180910390f35b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377fb8d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610286573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102aa9190610fb1565b336040518363ffffffff1660e01b81526004016102c8929190610ffa565b602060405180830381865afa1580156102e3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103079190611056565b61033d576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6103c9610c54565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d47875d0866040518263ffffffff1660e01b81526004016104249190611081565b602060405180830381865afa15801561043f573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061046391906110ae565b036104f1577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f10bd31b85876040518363ffffffff1660e01b81526004016104c39291906110e8565b5f604051808303815f87803b1580156104da575f5ffd5b505af11580156104ec573d5f5f3e3d5ffd5b505050505b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d47875d0336040518263ffffffff1660e01b815260040161054c9190611081565b602060405180830381865afa158015610567573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061058b91906110ae565b03610619577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f10bd31b33876040518363ffffffff1660e01b81526004016105eb9291906110e8565b5f604051808303815f87803b158015610602575f5ffd5b505af1158015610614573d5f5f3e3d5ffd5b505050505b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f59942313387878787876040518763ffffffff1660e01b815260040161067c9695949392919061110f565b5f604051808303815f87803b158015610693575f5ffd5b505af11580156106a5573d5f5f3e3d5ffd5b505050506106b1610c98565b5050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634d104adf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561075d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107819190610fb1565b336040518363ffffffff1660e01b815260040161079f929190610ffa565b602060405180830381865afa1580156107ba573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107de9190611056565b610814576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f5f5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16635fd13996876040518263ffffffff1660e01b8152600401610872919061116e565b60e060405180830381865afa15801561088d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108b191906111fb565b9650965050955050945094505f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d47875d0856040518263ffffffff1660e01b81526004016109189190611081565b602060405180830381865afa158015610933573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061095791906110ae565b036109e5577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16638b4fb30d84866040518363ffffffff1660e01b81526004016109b79291906110e8565b5f604051808303815f87803b1580156109ce575f5ffd5b505af11580156109e0573d5f5f3e3d5ffd5b505050505b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d47875d0876040518263ffffffff1660e01b8152600401610a409190611081565b602060405180830381865afa158015610a5b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a7f91906110ae565b03610b0d577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16638b4fb30d86866040518363ffffffff1660e01b8152600401610adf9291906110e8565b5f604051808303815f87803b158015610af6575f5ffd5b505af1158015610b08573d5f5f3e3d5ffd5b505050505b5f600167ffffffffffffffff811115610b2957610b28611298565b5b604051908082528060200260200182016040528015610b575781602001602082028036833780820191505090505b50905086815f81518110610b6e57610b6d6112c5565b5b6020026020010181815250507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663477c5f0182856040518363ffffffff1660e01b8152600401610bd59291906113a9565b5f604051808303815f87803b158015610bec575f5ffd5b505af1158015610bfe573d5f5f3e3d5ffd5b5050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f5403610c8f576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f81905550565b60015f81905550565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610ce3610cde610cd984610ca1565b610cc0565b610ca1565b9050919050565b5f610cf482610cc9565b9050919050565b5f610d0582610cea565b9050919050565b610d1581610cfb565b82525050565b5f602082019050610d2e5f830184610d0c565b92915050565b5f610d3e82610cea565b9050919050565b610d4e81610d34565b82525050565b5f602082019050610d675f830184610d45565b92915050565b5f610d7782610cea565b9050919050565b610d8781610d6d565b82525050565b5f602082019050610da05f830184610d7e565b92915050565b5f5ffd5b5f610db482610ca1565b9050919050565b610dc481610daa565b8114610dce575f5ffd5b50565b5f81359050610ddf81610dbb565b92915050565b5f5f60408385031215610dfb57610dfa610da6565b5b5f610e0885828601610dd1565b9250506020610e1985828601610dd1565b9150509250929050565b5f819050919050565b610e3581610e23565b8114610e3f575f5ffd5b50565b5f81359050610e5081610e2c565b92915050565b5f819050919050565b610e6881610e56565b8114610e72575f5ffd5b50565b5f81359050610e8381610e5f565b92915050565b5f5f5f5f5f60a08688031215610ea257610ea1610da6565b5b5f610eaf88828901610e42565b9550506020610ec088828901610dd1565b9450506040610ed188828901610dd1565b9350506060610ee288828901610e75565b9250506080610ef388828901610e75565b9150509295509295909350565b5f60208284031215610f1557610f14610da6565b5b5f610f2284828501610e75565b91505092915050565b5f610f3582610cea565b9050919050565b610f4581610f2b565b82525050565b5f602082019050610f5e5f830184610f3c565b92915050565b5f610f6e82610cea565b9050919050565b610f7e81610f64565b82525050565b5f602082019050610f975f830184610f75565b92915050565b5f81519050610fab81610e5f565b92915050565b5f60208284031215610fc657610fc5610da6565b5b5f610fd384828501610f9d565b91505092915050565b610fe581610e56565b82525050565b610ff481610daa565b82525050565b5f60408201905061100d5f830185610fdc565b61101a6020830184610feb565b9392505050565b5f8115159050919050565b61103581611021565b811461103f575f5ffd5b50565b5f815190506110508161102c565b92915050565b5f6020828403121561106b5761106a610da6565b5b5f61107884828501611042565b91505092915050565b5f6020820190506110945f830184610feb565b92915050565b5f815190506110a881610e2c565b92915050565b5f602082840312156110c3576110c2610da6565b5b5f6110d08482850161109a565b91505092915050565b6110e281610e23565b82525050565b5f6040820190506110fb5f830185610feb565b61110860208301846110d9565b9392505050565b5f60c0820190506111225f830189610feb565b61112f60208301886110d9565b61113c6040830187610feb565b6111496060830186610feb565b6111566080830185610fdc565b61116360a0830184610fdc565b979650505050505050565b5f6020820190506111815f830184610fdc565b92915050565b5f8151905061119581610dbb565b92915050565b5f67ffffffffffffffff82169050919050565b6111b78161119b565b81146111c1575f5ffd5b50565b5f815190506111d2816111ae565b92915050565b600481106111e4575f5ffd5b50565b5f815190506111f5816111d8565b92915050565b5f5f5f5f5f5f5f60e0888a03121561121657611215610da6565b5b5f6112238a828b01611187565b97505060206112348a828b0161109a565b96505060406112458a828b016111c4565b95505060606112568a828b01611187565b94505060806112678a828b01610f9d565b93505060a06112788a828b01611187565b92505060c06112898a828b016111e7565b91505092959891949750929550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61132481610e56565b82525050565b5f611335838361131b565b60208301905092915050565b5f602082019050919050565b5f611357826112f2565b61136181856112fc565b935061136c8361130c565b805f5b8381101561139c578151611383888261132a565b975061138e83611341565b92505060018101905061136f565b5085935050505092915050565b5f6040820190508181035f8301526113c1818561134d565b90506113d06020830184610feb565b939250505056fea2646970667358221220f81ebb6b4e5302d0e11af23e3d7210c2fabc6609cb79eb6756dd68cc2b5bb65964736f6c63430008210033",
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

// Deposit is a paid mutator transaction binding the contract method 0xbc1ca92d.
//
// Solidity: function deposit(uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId) returns()
func (_SettlementEngine *SettlementEngineTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte) (*types.Transaction, error) {
	return _SettlementEngine.contract.Transact(opts, "deposit", _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc1ca92d.
//
// Solidity: function deposit(uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId) returns()
func (_SettlementEngine *SettlementEngineSession) Deposit(_amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte) (*types.Transaction, error) {
	return _SettlementEngine.Contract.Deposit(&_SettlementEngine.TransactOpts, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc1ca92d.
//
// Solidity: function deposit(uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId) returns()
func (_SettlementEngine *SettlementEngineTransactorSession) Deposit(_amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte) (*types.Transaction, error) {
	return _SettlementEngine.Contract.Deposit(&_SettlementEngine.TransactOpts, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId)
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
