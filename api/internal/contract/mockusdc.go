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

// MockUSDCMetaData contains all meta data concerning the MockUSDC contract.
var MockUSDCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280600981526020017f4d6f636b205553444300000000000000000000000000000000000000000000008152506040518060400160405280600581526020017f6d55534443000000000000000000000000000000000000000000000000000000815250816003908161008b91906102f1565b50806004908161009b91906102f1565b5050506103c0565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061011e57607f821691505b602082108103610131576101306100da565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610158565b61019d8683610158565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6101e16101dc6101d7846101b5565b6101be565b6101b5565b9050919050565b5f819050919050565b6101fa836101c7565b61020e610206826101e8565b848454610164565b825550505050565b5f5f905090565b610225610216565b6102308184846101f1565b505050565b5f5b828110156102565761024b5f82840161021d565b600181019050610237565b505050565b601f8211156102a957828211156102a85761027581610137565b61027e83610149565b61028785610149565b6020861015610294575f90505b8083016102a382840382610235565b505050505b5b505050565b5f82821c905092915050565b5f6102c95f19846008026102ae565b1980831691505092915050565b5f6102e183836102ba565b9150826002028217905092915050565b6102fa826100a3565b67ffffffffffffffff811115610313576103126100ad565b5b61031d8254610107565b61032882828561025b565b5f60209050601f831160018114610359575f8415610347578287015190505b61035185826102d6565b8655506103b8565b601f19841661036786610137565b5f5b8281101561038e57848901518255600182019150602085019450602081019050610369565b868310156103ab57848901516103a7601f8916826102ba565b8355505b6001600288020188555050505b505050505050565b61106e806103cd5f395ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c806340c10f191161006457806340c10f191461015a57806370a082311461017657806395d89b41146101a6578063a9059cbb146101c4578063dd62ed3e146101f45761009c565b806306fdde03146100a0578063095ea7b3146100be57806318160ddd146100ee57806323b872dd1461010c578063313ce5671461013c575b5f5ffd5b6100a8610224565b6040516100b59190610b2d565b60405180910390f35b6100d860048036038101906100d39190610bde565b6102b4565b6040516100e59190610c36565b60405180910390f35b6100f66102d6565b6040516101039190610c5e565b60405180910390f35b61012660048036038101906101219190610c77565b6102df565b6040516101339190610c36565b60405180910390f35b61014461030d565b6040516101519190610ce2565b60405180910390f35b610174600480360381019061016f9190610bde565b610315565b005b610190600480360381019061018b9190610cfb565b610341565b60405161019d9190610c5e565b60405180910390f35b6101ae610386565b6040516101bb9190610b2d565b60405180910390f35b6101de60048036038101906101d99190610bde565b610416565b6040516101eb9190610c36565b60405180910390f35b61020e60048036038101906102099190610d26565b610438565b60405161021b9190610c5e565b60405180910390f35b60606003805461023390610d91565b80601f016020809104026020016040519081016040528092919081815260200182805461025f90610d91565b80156102aa5780601f10610281576101008083540402835291602001916102aa565b820191905f5260205f20905b81548152906001019060200180831161028d57829003601f168201915b5050505050905090565b5f5f6102be6104ba565b90506102cb8185856104c1565b600191505092915050565b5f600254905090565b5f5f6102e96104ba565b90506102f68582856104d3565b610301858585610566565b60019150509392505050565b5f6006905090565b61033d8261032161030d565b600a61032d9190610f1d565b836103389190610f67565b610656565b5050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60606004805461039590610d91565b80601f01602080910402602001604051908101604052809291908181526020018280546103c190610d91565b801561040c5780601f106103e35761010080835404028352916020019161040c565b820191905f5260205f20905b8154815290600101906020018083116103ef57829003601f168201915b5050505050905090565b5f5f6104206104ba565b905061042d818585610566565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f33905090565b6104ce83838360016106d5565b505050565b5f6104de8484610438565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156105605781811015610551578281836040517ffb8f41b200000000000000000000000000000000000000000000000000000000815260040161054893929190610fb7565b60405180910390fd5b61055f84848484035f6106d5565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036105d6575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016105cd9190610fec565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610646575f6040517fec442f0500000000000000000000000000000000000000000000000000000000815260040161063d9190610fec565b60405180910390fd5b6106518383836108a4565b505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036106c6575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016106bd9190610fec565b60405180910390fd5b6106d15f83836108a4565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610745575f6040517fe602df0500000000000000000000000000000000000000000000000000000000815260040161073c9190610fec565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036107b5575f6040517f94280d620000000000000000000000000000000000000000000000000000000081526004016107ac9190610fec565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550801561089e578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516108959190610c5e565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108f4578060025f8282546108e89190611005565b925050819055506109c2565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490508181101561097d578381836040517fe450d38c00000000000000000000000000000000000000000000000000000000815260040161097493929190610fb7565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a09578060025f8282540392505081905550610a53565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610ab09190610c5e565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610aff82610abd565b610b098185610ac7565b9350610b19818560208601610ad7565b610b2281610ae5565b840191505092915050565b5f6020820190508181035f830152610b458184610af5565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b7a82610b51565b9050919050565b610b8a81610b70565b8114610b94575f5ffd5b50565b5f81359050610ba581610b81565b92915050565b5f819050919050565b610bbd81610bab565b8114610bc7575f5ffd5b50565b5f81359050610bd881610bb4565b92915050565b5f5f60408385031215610bf457610bf3610b4d565b5b5f610c0185828601610b97565b9250506020610c1285828601610bca565b9150509250929050565b5f8115159050919050565b610c3081610c1c565b82525050565b5f602082019050610c495f830184610c27565b92915050565b610c5881610bab565b82525050565b5f602082019050610c715f830184610c4f565b92915050565b5f5f5f60608486031215610c8e57610c8d610b4d565b5b5f610c9b86828701610b97565b9350506020610cac86828701610b97565b9250506040610cbd86828701610bca565b9150509250925092565b5f60ff82169050919050565b610cdc81610cc7565b82525050565b5f602082019050610cf55f830184610cd3565b92915050565b5f60208284031215610d1057610d0f610b4d565b5b5f610d1d84828501610b97565b91505092915050565b5f5f60408385031215610d3c57610d3b610b4d565b5b5f610d4985828601610b97565b9250506020610d5a85828601610b97565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610da857607f821691505b602082108103610dbb57610dba610d64565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b6001851115610e4357808604811115610e1f57610e1e610dc1565b5b6001851615610e2e5780820291505b8081029050610e3c85610dee565b9450610e03565b94509492505050565b5f82610e5b5760019050610f16565b81610e68575f9050610f16565b8160018114610e7e5760028114610e8857610eb7565b6001915050610f16565b60ff841115610e9a57610e99610dc1565b5b8360020a915084821115610eb157610eb0610dc1565b5b50610f16565b5060208310610133831016604e8410600b8410161715610eec5782820a905083811115610ee757610ee6610dc1565b5b610f16565b610ef98484846001610dfa565b92509050818404811115610f1057610f0f610dc1565b5b81810290505b9392505050565b5f610f2782610bab565b9150610f3283610cc7565b9250610f5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484610e4c565b905092915050565b5f610f7182610bab565b9150610f7c83610bab565b9250828202610f8a81610bab565b91508282048414831517610fa157610fa0610dc1565b5b5092915050565b610fb181610b70565b82525050565b5f606082019050610fca5f830186610fa8565b610fd76020830185610c4f565b610fe46040830184610c4f565b949350505050565b5f602082019050610fff5f830184610fa8565b92915050565b5f61100f82610bab565b915061101a83610bab565b925082820190508082111561103257611031610dc1565b5b9291505056fea2646970667358221220e664a3fd8e2007da36ce73027da707d9e7336cbd07e6b2876f128bf75b06157d64736f6c63430008210033",
}

// MockUSDCABI is the input ABI used to generate the binding from.
// Deprecated: Use MockUSDCMetaData.ABI instead.
var MockUSDCABI = MockUSDCMetaData.ABI

// MockUSDCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockUSDCMetaData.Bin instead.
var MockUSDCBin = MockUSDCMetaData.Bin

// DeployMockUSDC deploys a new Ethereum contract, binding an instance of MockUSDC to it.
func DeployMockUSDC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MockUSDC, error) {
	parsed, err := MockUSDCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockUSDCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockUSDC{MockUSDCCaller: MockUSDCCaller{contract: contract}, MockUSDCTransactor: MockUSDCTransactor{contract: contract}, MockUSDCFilterer: MockUSDCFilterer{contract: contract}}, nil
}

// MockUSDC is an auto generated Go binding around an Ethereum contract.
type MockUSDC struct {
	MockUSDCCaller     // Read-only binding to the contract
	MockUSDCTransactor // Write-only binding to the contract
	MockUSDCFilterer   // Log filterer for contract events
}

// MockUSDCCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockUSDCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUSDCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockUSDCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUSDCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockUSDCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockUSDCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockUSDCSession struct {
	Contract     *MockUSDC         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockUSDCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockUSDCCallerSession struct {
	Contract *MockUSDCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MockUSDCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockUSDCTransactorSession struct {
	Contract     *MockUSDCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MockUSDCRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockUSDCRaw struct {
	Contract *MockUSDC // Generic contract binding to access the raw methods on
}

// MockUSDCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockUSDCCallerRaw struct {
	Contract *MockUSDCCaller // Generic read-only contract binding to access the raw methods on
}

// MockUSDCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockUSDCTransactorRaw struct {
	Contract *MockUSDCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockUSDC creates a new instance of MockUSDC, bound to a specific deployed contract.
func NewMockUSDC(address common.Address, backend bind.ContractBackend) (*MockUSDC, error) {
	contract, err := bindMockUSDC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockUSDC{MockUSDCCaller: MockUSDCCaller{contract: contract}, MockUSDCTransactor: MockUSDCTransactor{contract: contract}, MockUSDCFilterer: MockUSDCFilterer{contract: contract}}, nil
}

// NewMockUSDCCaller creates a new read-only instance of MockUSDC, bound to a specific deployed contract.
func NewMockUSDCCaller(address common.Address, caller bind.ContractCaller) (*MockUSDCCaller, error) {
	contract, err := bindMockUSDC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockUSDCCaller{contract: contract}, nil
}

// NewMockUSDCTransactor creates a new write-only instance of MockUSDC, bound to a specific deployed contract.
func NewMockUSDCTransactor(address common.Address, transactor bind.ContractTransactor) (*MockUSDCTransactor, error) {
	contract, err := bindMockUSDC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTransactor{contract: contract}, nil
}

// NewMockUSDCFilterer creates a new log filterer instance of MockUSDC, bound to a specific deployed contract.
func NewMockUSDCFilterer(address common.Address, filterer bind.ContractFilterer) (*MockUSDCFilterer, error) {
	contract, err := bindMockUSDC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockUSDCFilterer{contract: contract}, nil
}

// bindMockUSDC binds a generic wrapper to an already deployed contract.
func bindMockUSDC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockUSDCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockUSDC *MockUSDCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUSDC.Contract.MockUSDCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockUSDC *MockUSDCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUSDC.Contract.MockUSDCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockUSDC *MockUSDCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUSDC.Contract.MockUSDCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockUSDC *MockUSDCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockUSDC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockUSDC *MockUSDCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockUSDC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockUSDC *MockUSDCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockUSDC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockUSDC *MockUSDCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockUSDC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockUSDC *MockUSDCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockUSDC.Contract.Allowance(&_MockUSDC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MockUSDC *MockUSDCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MockUSDC.Contract.Allowance(&_MockUSDC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUSDC *MockUSDCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MockUSDC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUSDC *MockUSDCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockUSDC.Contract.BalanceOf(&_MockUSDC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MockUSDC *MockUSDCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MockUSDC.Contract.BalanceOf(&_MockUSDC.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MockUSDC *MockUSDCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MockUSDC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MockUSDC *MockUSDCSession) Decimals() (uint8, error) {
	return _MockUSDC.Contract.Decimals(&_MockUSDC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_MockUSDC *MockUSDCCallerSession) Decimals() (uint8, error) {
	return _MockUSDC.Contract.Decimals(&_MockUSDC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockUSDC *MockUSDCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockUSDC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockUSDC *MockUSDCSession) Name() (string, error) {
	return _MockUSDC.Contract.Name(&_MockUSDC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MockUSDC *MockUSDCCallerSession) Name() (string, error) {
	return _MockUSDC.Contract.Name(&_MockUSDC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockUSDC *MockUSDCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MockUSDC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockUSDC *MockUSDCSession) Symbol() (string, error) {
	return _MockUSDC.Contract.Symbol(&_MockUSDC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MockUSDC *MockUSDCCallerSession) Symbol() (string, error) {
	return _MockUSDC.Contract.Symbol(&_MockUSDC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockUSDC *MockUSDCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockUSDC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockUSDC *MockUSDCSession) TotalSupply() (*big.Int, error) {
	return _MockUSDC.Contract.TotalSupply(&_MockUSDC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MockUSDC *MockUSDCCallerSession) TotalSupply() (*big.Int, error) {
	return _MockUSDC.Contract.TotalSupply(&_MockUSDC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MockUSDC *MockUSDCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockUSDC.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MockUSDC *MockUSDCSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockUSDC.Contract.Approve(&_MockUSDC.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MockUSDC *MockUSDCTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockUSDC.Contract.Approve(&_MockUSDC.TransactOpts, spender, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockUSDC *MockUSDCTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSDC.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockUSDC *MockUSDCSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSDC.Contract.Mint(&_MockUSDC.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_MockUSDC *MockUSDCTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MockUSDC.Contract.Mint(&_MockUSDC.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MockUSDC *MockUSDCTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockUSDC.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MockUSDC *MockUSDCSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockUSDC.Contract.Transfer(&_MockUSDC.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MockUSDC *MockUSDCTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockUSDC.Contract.Transfer(&_MockUSDC.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MockUSDC *MockUSDCTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockUSDC.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MockUSDC *MockUSDCSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockUSDC.Contract.TransferFrom(&_MockUSDC.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MockUSDC *MockUSDCTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MockUSDC.Contract.TransferFrom(&_MockUSDC.TransactOpts, from, to, value)
}

// MockUSDCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MockUSDC contract.
type MockUSDCApprovalIterator struct {
	Event *MockUSDCApproval // Event containing the contract specifics and raw log

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
func (it *MockUSDCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUSDCApproval)
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
		it.Event = new(MockUSDCApproval)
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
func (it *MockUSDCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUSDCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUSDCApproval represents a Approval event raised by the MockUSDC contract.
type MockUSDCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockUSDC *MockUSDCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MockUSDCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockUSDC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MockUSDCApprovalIterator{contract: _MockUSDC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockUSDC *MockUSDCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MockUSDCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MockUSDC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUSDCApproval)
				if err := _MockUSDC.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MockUSDC *MockUSDCFilterer) ParseApproval(log types.Log) (*MockUSDCApproval, error) {
	event := new(MockUSDCApproval)
	if err := _MockUSDC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockUSDCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MockUSDC contract.
type MockUSDCTransferIterator struct {
	Event *MockUSDCTransfer // Event containing the contract specifics and raw log

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
func (it *MockUSDCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockUSDCTransfer)
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
		it.Event = new(MockUSDCTransfer)
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
func (it *MockUSDCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockUSDCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockUSDCTransfer represents a Transfer event raised by the MockUSDC contract.
type MockUSDCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockUSDC *MockUSDCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MockUSDCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockUSDC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MockUSDCTransferIterator{contract: _MockUSDC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockUSDC *MockUSDCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MockUSDCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MockUSDC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockUSDCTransfer)
				if err := _MockUSDC.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MockUSDC *MockUSDCFilterer) ParseTransfer(log types.Log) (*MockUSDCTransfer, error) {
	event := new(MockUSDCTransfer)
	if err := _MockUSDC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
