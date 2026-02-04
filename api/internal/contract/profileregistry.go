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

// ProfileRegistryMetaData contains all meta data concerning the ProfileRegistry contract.
var ProfileRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newScore\",\"type\":\"uint8\"}],\"name\":\"ProfileUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"getScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"profiles\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"trustScore\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"lastReversalAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"totalVolume\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalReversedVolume\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_score\",\"type\":\"uint8\"}],\"name\":\"setScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b506040516108b53803806108b5833981810160405281019061003191906100c9565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506100f4565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100988261006f565b9050919050565b6100a88161008e565b81146100b2575f5ffd5b50565b5f815190506100c38161009f565b92915050565b5f602082840312156100de576100dd61006b565b5b5f6100eb848285016100b5565b91505092915050565b60805161079c6101195f395f818160ed015281816101b101526101ed015261079c5ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c80635aa6e6751461004e578063bbe156271461006c578063d47875d01461009f578063db2ebdad146100cf575b5f5ffd5b6100566100eb565b6040516100639190610436565b60405180910390f35b6100866004803603810190610081919061048e565b61010f565b604051610096949392919061050e565b60405180910390f35b6100b960048036038101906100b4919061048e565b610159565b6040516100c69190610551565b60405180910390f35b6100e960048036038101906100e49190610594565b6101af565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b5f602052805f5260405f205f91509050805f015f9054906101000a900460ff1690805f0160019054906101000a900467ffffffffffffffff16908060010154908060020154905084565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff1660ff169050919050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16634d104adf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610254573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102789190610605565b336040518363ffffffff1660e01b815260040161029692919061064e565b602060405180830381865afa1580156102b1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102d591906106aa565b610314576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030b9061072f565b60405180910390fd5b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f6101000a81548160ff021916908360ff1602179055508173ffffffffffffffffffffffffffffffffffffffff167ffa27c1120fac3187a86799ce2fa4bab01117fc62929447129818338e8f764d53826040516103b0919061074d565b60405180910390a25050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f6103fe6103f96103f4846103bc565b6103db565b6103bc565b9050919050565b5f61040f826103e4565b9050919050565b5f61042082610405565b9050919050565b61043081610416565b82525050565b5f6020820190506104495f830184610427565b92915050565b5f5ffd5b5f61045d826103bc565b9050919050565b61046d81610453565b8114610477575f5ffd5b50565b5f8135905061048881610464565b92915050565b5f602082840312156104a3576104a261044f565b5b5f6104b08482850161047a565b91505092915050565b5f60ff82169050919050565b6104ce816104b9565b82525050565b5f67ffffffffffffffff82169050919050565b6104f0816104d4565b82525050565b5f819050919050565b610508816104f6565b82525050565b5f6080820190506105215f8301876104c5565b61052e60208301866104e7565b61053b60408301856104ff565b61054860608301846104ff565b95945050505050565b5f6020820190506105645f8301846104ff565b92915050565b610573816104b9565b811461057d575f5ffd5b50565b5f8135905061058e8161056a565b92915050565b5f5f604083850312156105aa576105a961044f565b5b5f6105b78582860161047a565b92505060206105c885828601610580565b9150509250929050565b5f819050919050565b6105e4816105d2565b81146105ee575f5ffd5b50565b5f815190506105ff816105db565b92915050565b5f6020828403121561061a5761061961044f565b5b5f610627848285016105f1565b91505092915050565b610639816105d2565b82525050565b61064881610453565b82525050565b5f6040820190506106615f830185610630565b61066e602083018461063f565b9392505050565b5f8115159050919050565b61068981610675565b8114610693575f5ffd5b50565b5f815190506106a481610680565b92915050565b5f602082840312156106bf576106be61044f565b5b5f6106cc84828501610696565b91505092915050565b5f82825260208201905092915050565b7f556e617574686f72697a656400000000000000000000000000000000000000005f82015250565b5f610719600c836106d5565b9150610724826106e5565b602082019050919050565b5f6020820190508181035f8301526107468161070d565b9050919050565b5f6020820190506107605f8301846104c5565b9291505056fea2646970667358221220dcb56613bca3aef71ac5ebb8add03145bfd906b24b2227fcbad80980be8e112c64736f6c63430008210033",
}

// ProfileRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfileRegistryMetaData.ABI instead.
var ProfileRegistryABI = ProfileRegistryMetaData.ABI

// ProfileRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProfileRegistryMetaData.Bin instead.
var ProfileRegistryBin = ProfileRegistryMetaData.Bin

// DeployProfileRegistry deploys a new Ethereum contract, binding an instance of ProfileRegistry to it.
func DeployProfileRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _gov common.Address) (common.Address, *types.Transaction, *ProfileRegistry, error) {
	parsed, err := ProfileRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProfileRegistryBin), backend, _gov)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProfileRegistry{ProfileRegistryCaller: ProfileRegistryCaller{contract: contract}, ProfileRegistryTransactor: ProfileRegistryTransactor{contract: contract}, ProfileRegistryFilterer: ProfileRegistryFilterer{contract: contract}}, nil
}

// ProfileRegistry is an auto generated Go binding around an Ethereum contract.
type ProfileRegistry struct {
	ProfileRegistryCaller     // Read-only binding to the contract
	ProfileRegistryTransactor // Write-only binding to the contract
	ProfileRegistryFilterer   // Log filterer for contract events
}

// ProfileRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfileRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfileRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfileRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfileRegistrySession struct {
	Contract     *ProfileRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfileRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfileRegistryCallerSession struct {
	Contract *ProfileRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProfileRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfileRegistryTransactorSession struct {
	Contract     *ProfileRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProfileRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfileRegistryRaw struct {
	Contract *ProfileRegistry // Generic contract binding to access the raw methods on
}

// ProfileRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfileRegistryCallerRaw struct {
	Contract *ProfileRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProfileRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfileRegistryTransactorRaw struct {
	Contract *ProfileRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfileRegistry creates a new instance of ProfileRegistry, bound to a specific deployed contract.
func NewProfileRegistry(address common.Address, backend bind.ContractBackend) (*ProfileRegistry, error) {
	contract, err := bindProfileRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistry{ProfileRegistryCaller: ProfileRegistryCaller{contract: contract}, ProfileRegistryTransactor: ProfileRegistryTransactor{contract: contract}, ProfileRegistryFilterer: ProfileRegistryFilterer{contract: contract}}, nil
}

// NewProfileRegistryCaller creates a new read-only instance of ProfileRegistry, bound to a specific deployed contract.
func NewProfileRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProfileRegistryCaller, error) {
	contract, err := bindProfileRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistryCaller{contract: contract}, nil
}

// NewProfileRegistryTransactor creates a new write-only instance of ProfileRegistry, bound to a specific deployed contract.
func NewProfileRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfileRegistryTransactor, error) {
	contract, err := bindProfileRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistryTransactor{contract: contract}, nil
}

// NewProfileRegistryFilterer creates a new log filterer instance of ProfileRegistry, bound to a specific deployed contract.
func NewProfileRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfileRegistryFilterer, error) {
	contract, err := bindProfileRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistryFilterer{contract: contract}, nil
}

// bindProfileRegistry binds a generic wrapper to an already deployed contract.
func bindProfileRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProfileRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProfileRegistry *ProfileRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProfileRegistry.Contract.ProfileRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProfileRegistry *ProfileRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileRegistry.Contract.ProfileRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProfileRegistry *ProfileRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProfileRegistry.Contract.ProfileRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProfileRegistry *ProfileRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProfileRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProfileRegistry *ProfileRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProfileRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProfileRegistry *ProfileRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProfileRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address _wallet) view returns(uint256)
func (_ProfileRegistry *ProfileRegistryCaller) GetScore(opts *bind.CallOpts, _wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProfileRegistry.contract.Call(opts, &out, "getScore", _wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address _wallet) view returns(uint256)
func (_ProfileRegistry *ProfileRegistrySession) GetScore(_wallet common.Address) (*big.Int, error) {
	return _ProfileRegistry.Contract.GetScore(&_ProfileRegistry.CallOpts, _wallet)
}

// GetScore is a free data retrieval call binding the contract method 0xd47875d0.
//
// Solidity: function getScore(address _wallet) view returns(uint256)
func (_ProfileRegistry *ProfileRegistryCallerSession) GetScore(_wallet common.Address) (*big.Int, error) {
	return _ProfileRegistry.Contract.GetScore(&_ProfileRegistry.CallOpts, _wallet)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_ProfileRegistry *ProfileRegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProfileRegistry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_ProfileRegistry *ProfileRegistrySession) Governance() (common.Address, error) {
	return _ProfileRegistry.Contract.Governance(&_ProfileRegistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_ProfileRegistry *ProfileRegistryCallerSession) Governance() (common.Address, error) {
	return _ProfileRegistry.Contract.Governance(&_ProfileRegistry.CallOpts)
}

// Profiles is a free data retrieval call binding the contract method 0xbbe15627.
//
// Solidity: function profiles(address ) view returns(uint8 trustScore, uint64 lastReversalAt, uint256 totalVolume, uint256 totalReversedVolume)
func (_ProfileRegistry *ProfileRegistryCaller) Profiles(opts *bind.CallOpts, arg0 common.Address) (struct {
	TrustScore          uint8
	LastReversalAt      uint64
	TotalVolume         *big.Int
	TotalReversedVolume *big.Int
}, error) {
	var out []interface{}
	err := _ProfileRegistry.contract.Call(opts, &out, "profiles", arg0)

	outstruct := new(struct {
		TrustScore          uint8
		LastReversalAt      uint64
		TotalVolume         *big.Int
		TotalReversedVolume *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TrustScore = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.LastReversalAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.TotalVolume = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalReversedVolume = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Profiles is a free data retrieval call binding the contract method 0xbbe15627.
//
// Solidity: function profiles(address ) view returns(uint8 trustScore, uint64 lastReversalAt, uint256 totalVolume, uint256 totalReversedVolume)
func (_ProfileRegistry *ProfileRegistrySession) Profiles(arg0 common.Address) (struct {
	TrustScore          uint8
	LastReversalAt      uint64
	TotalVolume         *big.Int
	TotalReversedVolume *big.Int
}, error) {
	return _ProfileRegistry.Contract.Profiles(&_ProfileRegistry.CallOpts, arg0)
}

// Profiles is a free data retrieval call binding the contract method 0xbbe15627.
//
// Solidity: function profiles(address ) view returns(uint8 trustScore, uint64 lastReversalAt, uint256 totalVolume, uint256 totalReversedVolume)
func (_ProfileRegistry *ProfileRegistryCallerSession) Profiles(arg0 common.Address) (struct {
	TrustScore          uint8
	LastReversalAt      uint64
	TotalVolume         *big.Int
	TotalReversedVolume *big.Int
}, error) {
	return _ProfileRegistry.Contract.Profiles(&_ProfileRegistry.CallOpts, arg0)
}

// SetScore is a paid mutator transaction binding the contract method 0xdb2ebdad.
//
// Solidity: function setScore(address _wallet, uint8 _score) returns()
func (_ProfileRegistry *ProfileRegistryTransactor) SetScore(opts *bind.TransactOpts, _wallet common.Address, _score uint8) (*types.Transaction, error) {
	return _ProfileRegistry.contract.Transact(opts, "setScore", _wallet, _score)
}

// SetScore is a paid mutator transaction binding the contract method 0xdb2ebdad.
//
// Solidity: function setScore(address _wallet, uint8 _score) returns()
func (_ProfileRegistry *ProfileRegistrySession) SetScore(_wallet common.Address, _score uint8) (*types.Transaction, error) {
	return _ProfileRegistry.Contract.SetScore(&_ProfileRegistry.TransactOpts, _wallet, _score)
}

// SetScore is a paid mutator transaction binding the contract method 0xdb2ebdad.
//
// Solidity: function setScore(address _wallet, uint8 _score) returns()
func (_ProfileRegistry *ProfileRegistryTransactorSession) SetScore(_wallet common.Address, _score uint8) (*types.Transaction, error) {
	return _ProfileRegistry.Contract.SetScore(&_ProfileRegistry.TransactOpts, _wallet, _score)
}

// ProfileRegistryProfileUpdatedIterator is returned from FilterProfileUpdated and is used to iterate over the raw logs and unpacked data for ProfileUpdated events raised by the ProfileRegistry contract.
type ProfileRegistryProfileUpdatedIterator struct {
	Event *ProfileRegistryProfileUpdated // Event containing the contract specifics and raw log

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
func (it *ProfileRegistryProfileUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileRegistryProfileUpdated)
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
		it.Event = new(ProfileRegistryProfileUpdated)
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
func (it *ProfileRegistryProfileUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileRegistryProfileUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileRegistryProfileUpdated represents a ProfileUpdated event raised by the ProfileRegistry contract.
type ProfileRegistryProfileUpdated struct {
	Wallet   common.Address
	NewScore uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProfileUpdated is a free log retrieval operation binding the contract event 0xfa27c1120fac3187a86799ce2fa4bab01117fc62929447129818338e8f764d53.
//
// Solidity: event ProfileUpdated(address indexed wallet, uint8 newScore)
func (_ProfileRegistry *ProfileRegistryFilterer) FilterProfileUpdated(opts *bind.FilterOpts, wallet []common.Address) (*ProfileRegistryProfileUpdatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ProfileRegistry.contract.FilterLogs(opts, "ProfileUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return &ProfileRegistryProfileUpdatedIterator{contract: _ProfileRegistry.contract, event: "ProfileUpdated", logs: logs, sub: sub}, nil
}

// WatchProfileUpdated is a free log subscription operation binding the contract event 0xfa27c1120fac3187a86799ce2fa4bab01117fc62929447129818338e8f764d53.
//
// Solidity: event ProfileUpdated(address indexed wallet, uint8 newScore)
func (_ProfileRegistry *ProfileRegistryFilterer) WatchProfileUpdated(opts *bind.WatchOpts, sink chan<- *ProfileRegistryProfileUpdated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _ProfileRegistry.contract.WatchLogs(opts, "ProfileUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileRegistryProfileUpdated)
				if err := _ProfileRegistry.contract.UnpackLog(event, "ProfileUpdated", log); err != nil {
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

// ParseProfileUpdated is a log parse operation binding the contract event 0xfa27c1120fac3187a86799ce2fa4bab01117fc62929447129818338e8f764d53.
//
// Solidity: event ProfileUpdated(address indexed wallet, uint8 newScore)
func (_ProfileRegistry *ProfileRegistryFilterer) ParseProfileUpdated(log types.Log) (*ProfileRegistryProfileUpdated, error) {
	event := new(ProfileRegistryProfileUpdated)
	if err := _ProfileRegistry.contract.UnpackLog(event, "ProfileUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
