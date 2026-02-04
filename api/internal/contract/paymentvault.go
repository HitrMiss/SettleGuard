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

// PaymentVaultMetaData contains all meta data concerning the PaymentVault contract.
var PaymentVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"packetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentRefunded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"}],\"name\":\"getPaymentDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"createdAt\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"merchantIdentity\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"categoryId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"merchantPayout\",\"type\":\"address\"},{\"internalType\":\"enumIPaymentVault.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_merchantIdentity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_merchantPayout\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_categoryId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"}],\"name\":\"lockFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"createdAt\",\"type\":\"uint64\"},{\"internalType\":\"enumIPaymentVault.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"merchantIdentity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"merchantPayout\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"categoryId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_packetIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"_settlementWallet\",\"type\":\"address\"}],\"name\":\"prepareSettlementBulk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_engine\",\"type\":\"address\"}],\"name\":\"setEngine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementEngine\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newEngine\",\"type\":\"address\"}],\"name\":\"updateSettlementEngine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161189a38038061189a83398181016040528101906100319190610105565b60015f819055508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505050610143565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100d4826100ab565b9050919050565b6100e4816100ca565b81146100ee575f5ffd5b50565b5f815190506100ff816100db565b92915050565b5f5f6040838503121561011b5761011a6100a7565b5b5f610128858286016100f1565b9250506020610139858286016100f1565b9150509250929050565b60805160a0516117056101955f395f818161044d015281816105cd0152818161087f0152610e0201525f81816102ae015281816102ea015281816106210152818161096b01526109a701526117055ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c80635fd13996116100645780635fd139961461014a5780637249fbb614610180578063e6d0919d1461019c578063e9e2dbab146101ba578063f5994231146101d65761009c565b80630716326d146100a05780630e830e49146100d65780633e413bee146100f2578063477c5f01146101105780635aa6e6751461012c575b5f5ffd5b6100ba60048036038101906100b5919061107b565b6101f2565b6040516100cd97969594939291906111a1565b60405180910390f35b6100f060048036038101906100eb9190611238565b6102ac565b005b6100fa61044b565b60405161010791906112be565b60405180910390f35b61012a60048036038101906101259190611338565b61046f565b005b61013461061f565b60405161014191906113b5565b60405180910390f35b610164600480360381019061015f919061107b565b610643565b60405161017797969594939291906113ce565b60405180910390f35b61019a6004803603810190610195919061107b565b610716565b005b6101a4610944565b6040516101b1919061143b565b60405180910390f35b6101d460048036038101906101cf9190611238565b610969565b005b6101f060048036038101906101eb919061147e565b610b08565b005b6002602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690805f0160149054906101000a900467ffffffffffffffff1690805f01601c9054906101000a900460ff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154905087565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377fb8d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610351573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610375919061151b565b336040518363ffffffff1660e01b8152600401610393929190611546565b602060405180830381865afa1580156103ae573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103d291906115a2565b610408576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104f5576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6104fd610e57565b5f5f90505f5f90505b848490508110156105c5575f60025f878785818110610528576105276115cd565b5b9050602002013581526020019081526020015f2090506001600381111561055257610551611107565b5b815f01601c9054906101000a900460ff16600381111561057557610574611107565b5b036105b9576002815f01601c6101000a81548160ff021916908360038111156105a1576105a0611107565b5b02179055508060030154836105b69190611627565b92505b81600101915050610506565b5061061182827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610e9b9092919063ffffffff16565b5061061a610f1a565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f5f5f5f5f5f60025f8a81526020019081526020015f209050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160030154825f0160149054906101000a900467ffffffffffffffff16836001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168460040154856002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16865f01601c9054906101000a900460ff16975097509750975097509750975050919395979092949650565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461079c576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107a4610e57565b5f60025f8381526020019081526020015f209050600160038111156107cc576107cb611107565b5b815f01601c9054906101000a900460ff1660038111156107ef576107ee611107565b5b14610826576040517f80cb55e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003815f01601c6101000a81548160ff0219169083600381111561084d5761084c611107565b5b02179055506108c3815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682600301547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610e9b9092919063ffffffff16565b805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16827f3b28817a65441bed1ac3a5496c6b71346aa302e80c7418c12e1c49eaa17211768360030154604051610930919061165a565b60405180910390a350610941610f1a565b50565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377fb8d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a0e573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a32919061151b565b336040518363ffffffff1660e01b8152600401610a50929190611546565b602060405180830381865afa158015610a6b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a8f91906115a2565b610ac5576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b8e576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b96610e57565b5f6003811115610ba957610ba8611107565b5b60025f8381526020019081526020015f205f01601c9054906101000a900460ff166003811115610bdc57610bdb611107565b5b14610c13576040517f23369fa600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060e001604052808773ffffffffffffffffffffffffffffffffffffffff1681526020014267ffffffffffffffff16815260200160016003811115610c5e57610c5d611107565b5b81526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018381525060025f8381526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151815f0160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040820151815f01601c6101000a81548160ff02191690836003811115610d5257610d51611107565b5b02179055506060820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816003015560c08201518160040155905050610e478630877f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610f23909392919063ffffffff16565b610e4f610f1a565b505050505050565b60025f5403610e92576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f81905550565b610f15838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8585604051602401610ece929190611673565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610fa5565b505050565b60015f81905550565b610f9f848573ffffffffffffffffffffffffffffffffffffffff166323b872dd868686604051602401610f589392919061169a565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610fa5565b50505050565b5f5f60205f8451602086015f885af180610fc4576040513d5f823e3d81fd5b3d92505f519150505f8214610fdd576001811415610ff8565b5f8473ffffffffffffffffffffffffffffffffffffffff163b145b1561103a57836040517f5274afe7000000000000000000000000000000000000000000000000000000008152600401611031919061143b565b60405180910390fd5b50505050565b5f5ffd5b5f5ffd5b5f819050919050565b61105a81611048565b8114611064575f5ffd5b50565b5f8135905061107581611051565b92915050565b5f602082840312156110905761108f611040565b5b5f61109d84828501611067565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6110cf826110a6565b9050919050565b6110df816110c5565b82525050565b5f67ffffffffffffffff82169050919050565b611101816110e5565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6004811061114557611144611107565b5b50565b5f81905061115582611134565b919050565b5f61116482611148565b9050919050565b6111748161115a565b82525050565b5f819050919050565b61118c8161117a565b82525050565b61119b81611048565b82525050565b5f60e0820190506111b45f83018a6110d6565b6111c160208301896110f8565b6111ce604083018861116b565b6111db60608301876110d6565b6111e860808301866110d6565b6111f560a0830185611183565b61120260c0830184611192565b98975050505050505050565b611217816110c5565b8114611221575f5ffd5b50565b5f813590506112328161120e565b92915050565b5f6020828403121561124d5761124c611040565b5b5f61125a84828501611224565b91505092915050565b5f819050919050565b5f61128661128161127c846110a6565b611263565b6110a6565b9050919050565b5f6112978261126c565b9050919050565b5f6112a88261128d565b9050919050565b6112b88161129e565b82525050565b5f6020820190506112d15f8301846112af565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126112f8576112f76112d7565b5b8235905067ffffffffffffffff811115611315576113146112db565b5b602083019150836020820283011115611331576113306112df565b5b9250929050565b5f5f5f6040848603121561134f5761134e611040565b5b5f84013567ffffffffffffffff81111561136c5761136b611044565b5b611378868287016112e3565b9350935050602061138b86828701611224565b9150509250925092565b5f61139f8261128d565b9050919050565b6113af81611395565b82525050565b5f6020820190506113c85f8301846113a6565b92915050565b5f60e0820190506113e15f83018a6110d6565b6113ee6020830189611183565b6113fb60408301886110f8565b61140860608301876110d6565b6114156080830186611192565b61142260a08301856110d6565b61142f60c083018461116b565b98975050505050505050565b5f60208201905061144e5f8301846110d6565b92915050565b61145d8161117a565b8114611467575f5ffd5b50565b5f8135905061147881611454565b92915050565b5f5f5f5f5f5f60c0878903121561149857611497611040565b5b5f6114a589828a01611224565b96505060206114b689828a0161146a565b95505060406114c789828a01611224565b94505060606114d889828a01611224565b93505060806114e989828a01611067565b92505060a06114fa89828a01611067565b9150509295509295509295565b5f8151905061151581611051565b92915050565b5f602082840312156115305761152f611040565b5b5f61153d84828501611507565b91505092915050565b5f6040820190506115595f830185611192565b61156660208301846110d6565b9392505050565b5f8115159050919050565b6115818161156d565b811461158b575f5ffd5b50565b5f8151905061159c81611578565b92915050565b5f602082840312156115b7576115b6611040565b5b5f6115c48482850161158e565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6116318261117a565b915061163c8361117a565b9250828201905080821115611654576116536115fa565b5b92915050565b5f60208201905061166d5f830184611183565b92915050565b5f6040820190506116865f8301856110d6565b6116936020830184611183565b9392505050565b5f6060820190506116ad5f8301866110d6565b6116ba60208301856110d6565b6116c76040830184611183565b94935050505056fea2646970667358221220b50efd4442c3ed34eaccd93a9fbd3da6b8a3a30d524c5ba37d961c94448ebf8764736f6c63430008210033",
}

// PaymentVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentVaultMetaData.ABI instead.
var PaymentVaultABI = PaymentVaultMetaData.ABI

// PaymentVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PaymentVaultMetaData.Bin instead.
var PaymentVaultBin = PaymentVaultMetaData.Bin

// DeployPaymentVault deploys a new Ethereum contract, binding an instance of PaymentVault to it.
func DeployPaymentVault(auth *bind.TransactOpts, backend bind.ContractBackend, _usdc common.Address, _gov common.Address) (common.Address, *types.Transaction, *PaymentVault, error) {
	parsed, err := PaymentVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymentVaultBin), backend, _usdc, _gov)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PaymentVault{PaymentVaultCaller: PaymentVaultCaller{contract: contract}, PaymentVaultTransactor: PaymentVaultTransactor{contract: contract}, PaymentVaultFilterer: PaymentVaultFilterer{contract: contract}}, nil
}

// PaymentVault is an auto generated Go binding around an Ethereum contract.
type PaymentVault struct {
	PaymentVaultCaller     // Read-only binding to the contract
	PaymentVaultTransactor // Write-only binding to the contract
	PaymentVaultFilterer   // Log filterer for contract events
}

// PaymentVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentVaultSession struct {
	Contract     *PaymentVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentVaultCallerSession struct {
	Contract *PaymentVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PaymentVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentVaultTransactorSession struct {
	Contract     *PaymentVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PaymentVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentVaultRaw struct {
	Contract *PaymentVault // Generic contract binding to access the raw methods on
}

// PaymentVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentVaultCallerRaw struct {
	Contract *PaymentVaultCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentVaultTransactorRaw struct {
	Contract *PaymentVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentVault creates a new instance of PaymentVault, bound to a specific deployed contract.
func NewPaymentVault(address common.Address, backend bind.ContractBackend) (*PaymentVault, error) {
	contract, err := bindPaymentVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentVault{PaymentVaultCaller: PaymentVaultCaller{contract: contract}, PaymentVaultTransactor: PaymentVaultTransactor{contract: contract}, PaymentVaultFilterer: PaymentVaultFilterer{contract: contract}}, nil
}

// NewPaymentVaultCaller creates a new read-only instance of PaymentVault, bound to a specific deployed contract.
func NewPaymentVaultCaller(address common.Address, caller bind.ContractCaller) (*PaymentVaultCaller, error) {
	contract, err := bindPaymentVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultCaller{contract: contract}, nil
}

// NewPaymentVaultTransactor creates a new write-only instance of PaymentVault, bound to a specific deployed contract.
func NewPaymentVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentVaultTransactor, error) {
	contract, err := bindPaymentVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultTransactor{contract: contract}, nil
}

// NewPaymentVaultFilterer creates a new log filterer instance of PaymentVault, bound to a specific deployed contract.
func NewPaymentVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentVaultFilterer, error) {
	contract, err := bindPaymentVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultFilterer{contract: contract}, nil
}

// bindPaymentVault binds a generic wrapper to an already deployed contract.
func bindPaymentVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentVault *PaymentVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentVault.Contract.PaymentVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentVault *PaymentVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentVault.Contract.PaymentVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentVault *PaymentVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentVault.Contract.PaymentVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentVault *PaymentVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentVault *PaymentVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentVault *PaymentVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentVault.Contract.contract.Transact(opts, method, params...)
}

// GetPaymentDetails is a free data retrieval call binding the contract method 0x5fd13996.
//
// Solidity: function getPaymentDetails(bytes32 _packetId) view returns(address payer, uint256 amount, uint64 createdAt, address merchantIdentity, bytes32 categoryId, address merchantPayout, uint8 status)
func (_PaymentVault *PaymentVaultCaller) GetPaymentDetails(opts *bind.CallOpts, _packetId [32]byte) (struct {
	Payer            common.Address
	Amount           *big.Int
	CreatedAt        uint64
	MerchantIdentity common.Address
	CategoryId       [32]byte
	MerchantPayout   common.Address
	Status           uint8
}, error) {
	var out []interface{}
	err := _PaymentVault.contract.Call(opts, &out, "getPaymentDetails", _packetId)

	outstruct := new(struct {
		Payer            common.Address
		Amount           *big.Int
		CreatedAt        uint64
		MerchantIdentity common.Address
		CategoryId       [32]byte
		MerchantPayout   common.Address
		Status           uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Payer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.MerchantIdentity = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.CategoryId = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.MerchantPayout = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetPaymentDetails is a free data retrieval call binding the contract method 0x5fd13996.
//
// Solidity: function getPaymentDetails(bytes32 _packetId) view returns(address payer, uint256 amount, uint64 createdAt, address merchantIdentity, bytes32 categoryId, address merchantPayout, uint8 status)
func (_PaymentVault *PaymentVaultSession) GetPaymentDetails(_packetId [32]byte) (struct {
	Payer            common.Address
	Amount           *big.Int
	CreatedAt        uint64
	MerchantIdentity common.Address
	CategoryId       [32]byte
	MerchantPayout   common.Address
	Status           uint8
}, error) {
	return _PaymentVault.Contract.GetPaymentDetails(&_PaymentVault.CallOpts, _packetId)
}

// GetPaymentDetails is a free data retrieval call binding the contract method 0x5fd13996.
//
// Solidity: function getPaymentDetails(bytes32 _packetId) view returns(address payer, uint256 amount, uint64 createdAt, address merchantIdentity, bytes32 categoryId, address merchantPayout, uint8 status)
func (_PaymentVault *PaymentVaultCallerSession) GetPaymentDetails(_packetId [32]byte) (struct {
	Payer            common.Address
	Amount           *big.Int
	CreatedAt        uint64
	MerchantIdentity common.Address
	CategoryId       [32]byte
	MerchantPayout   common.Address
	Status           uint8
}, error) {
	return _PaymentVault.Contract.GetPaymentDetails(&_PaymentVault.CallOpts, _packetId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PaymentVault *PaymentVaultCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentVault.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PaymentVault *PaymentVaultSession) Governance() (common.Address, error) {
	return _PaymentVault.Contract.Governance(&_PaymentVault.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PaymentVault *PaymentVaultCallerSession) Governance() (common.Address, error) {
	return _PaymentVault.Contract.Governance(&_PaymentVault.CallOpts)
}

// Payments is a free data retrieval call binding the contract method 0x0716326d.
//
// Solidity: function payments(bytes32 ) view returns(address payer, uint64 createdAt, uint8 status, address merchantIdentity, address merchantPayout, uint256 amount, bytes32 categoryId)
func (_PaymentVault *PaymentVaultCaller) Payments(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Payer            common.Address
	CreatedAt        uint64
	Status           uint8
	MerchantIdentity common.Address
	MerchantPayout   common.Address
	Amount           *big.Int
	CategoryId       [32]byte
}, error) {
	var out []interface{}
	err := _PaymentVault.contract.Call(opts, &out, "payments", arg0)

	outstruct := new(struct {
		Payer            common.Address
		CreatedAt        uint64
		Status           uint8
		MerchantIdentity common.Address
		MerchantPayout   common.Address
		Amount           *big.Int
		CategoryId       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Payer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.MerchantIdentity = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.MerchantPayout = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.CategoryId = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Payments is a free data retrieval call binding the contract method 0x0716326d.
//
// Solidity: function payments(bytes32 ) view returns(address payer, uint64 createdAt, uint8 status, address merchantIdentity, address merchantPayout, uint256 amount, bytes32 categoryId)
func (_PaymentVault *PaymentVaultSession) Payments(arg0 [32]byte) (struct {
	Payer            common.Address
	CreatedAt        uint64
	Status           uint8
	MerchantIdentity common.Address
	MerchantPayout   common.Address
	Amount           *big.Int
	CategoryId       [32]byte
}, error) {
	return _PaymentVault.Contract.Payments(&_PaymentVault.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x0716326d.
//
// Solidity: function payments(bytes32 ) view returns(address payer, uint64 createdAt, uint8 status, address merchantIdentity, address merchantPayout, uint256 amount, bytes32 categoryId)
func (_PaymentVault *PaymentVaultCallerSession) Payments(arg0 [32]byte) (struct {
	Payer            common.Address
	CreatedAt        uint64
	Status           uint8
	MerchantIdentity common.Address
	MerchantPayout   common.Address
	Amount           *big.Int
	CategoryId       [32]byte
}, error) {
	return _PaymentVault.Contract.Payments(&_PaymentVault.CallOpts, arg0)
}

// SettlementEngine is a free data retrieval call binding the contract method 0xe6d0919d.
//
// Solidity: function settlementEngine() view returns(address)
func (_PaymentVault *PaymentVaultCaller) SettlementEngine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentVault.contract.Call(opts, &out, "settlementEngine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SettlementEngine is a free data retrieval call binding the contract method 0xe6d0919d.
//
// Solidity: function settlementEngine() view returns(address)
func (_PaymentVault *PaymentVaultSession) SettlementEngine() (common.Address, error) {
	return _PaymentVault.Contract.SettlementEngine(&_PaymentVault.CallOpts)
}

// SettlementEngine is a free data retrieval call binding the contract method 0xe6d0919d.
//
// Solidity: function settlementEngine() view returns(address)
func (_PaymentVault *PaymentVaultCallerSession) SettlementEngine() (common.Address, error) {
	return _PaymentVault.Contract.SettlementEngine(&_PaymentVault.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_PaymentVault *PaymentVaultCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentVault.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_PaymentVault *PaymentVaultSession) Usdc() (common.Address, error) {
	return _PaymentVault.Contract.Usdc(&_PaymentVault.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_PaymentVault *PaymentVaultCallerSession) Usdc() (common.Address, error) {
	return _PaymentVault.Contract.Usdc(&_PaymentVault.CallOpts)
}

// LockFunds is a paid mutator transaction binding the contract method 0xf5994231.
//
// Solidity: function lockFunds(address _payer, uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId) returns()
func (_PaymentVault *PaymentVaultTransactor) LockFunds(opts *bind.TransactOpts, _payer common.Address, _amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "lockFunds", _payer, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId)
}

// LockFunds is a paid mutator transaction binding the contract method 0xf5994231.
//
// Solidity: function lockFunds(address _payer, uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId) returns()
func (_PaymentVault *PaymentVaultSession) LockFunds(_payer common.Address, _amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte) (*types.Transaction, error) {
	return _PaymentVault.Contract.LockFunds(&_PaymentVault.TransactOpts, _payer, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId)
}

// LockFunds is a paid mutator transaction binding the contract method 0xf5994231.
//
// Solidity: function lockFunds(address _payer, uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId) returns()
func (_PaymentVault *PaymentVaultTransactorSession) LockFunds(_payer common.Address, _amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte) (*types.Transaction, error) {
	return _PaymentVault.Contract.LockFunds(&_PaymentVault.TransactOpts, _payer, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId)
}

// PrepareSettlementBulk is a paid mutator transaction binding the contract method 0x477c5f01.
//
// Solidity: function prepareSettlementBulk(bytes32[] _packetIds, address _settlementWallet) returns()
func (_PaymentVault *PaymentVaultTransactor) PrepareSettlementBulk(opts *bind.TransactOpts, _packetIds [][32]byte, _settlementWallet common.Address) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "prepareSettlementBulk", _packetIds, _settlementWallet)
}

// PrepareSettlementBulk is a paid mutator transaction binding the contract method 0x477c5f01.
//
// Solidity: function prepareSettlementBulk(bytes32[] _packetIds, address _settlementWallet) returns()
func (_PaymentVault *PaymentVaultSession) PrepareSettlementBulk(_packetIds [][32]byte, _settlementWallet common.Address) (*types.Transaction, error) {
	return _PaymentVault.Contract.PrepareSettlementBulk(&_PaymentVault.TransactOpts, _packetIds, _settlementWallet)
}

// PrepareSettlementBulk is a paid mutator transaction binding the contract method 0x477c5f01.
//
// Solidity: function prepareSettlementBulk(bytes32[] _packetIds, address _settlementWallet) returns()
func (_PaymentVault *PaymentVaultTransactorSession) PrepareSettlementBulk(_packetIds [][32]byte, _settlementWallet common.Address) (*types.Transaction, error) {
	return _PaymentVault.Contract.PrepareSettlementBulk(&_PaymentVault.TransactOpts, _packetIds, _settlementWallet)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _packetId) returns()
func (_PaymentVault *PaymentVaultTransactor) Refund(opts *bind.TransactOpts, _packetId [32]byte) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "refund", _packetId)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _packetId) returns()
func (_PaymentVault *PaymentVaultSession) Refund(_packetId [32]byte) (*types.Transaction, error) {
	return _PaymentVault.Contract.Refund(&_PaymentVault.TransactOpts, _packetId)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 _packetId) returns()
func (_PaymentVault *PaymentVaultTransactorSession) Refund(_packetId [32]byte) (*types.Transaction, error) {
	return _PaymentVault.Contract.Refund(&_PaymentVault.TransactOpts, _packetId)
}

// SetEngine is a paid mutator transaction binding the contract method 0x0e830e49.
//
// Solidity: function setEngine(address _engine) returns()
func (_PaymentVault *PaymentVaultTransactor) SetEngine(opts *bind.TransactOpts, _engine common.Address) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "setEngine", _engine)
}

// SetEngine is a paid mutator transaction binding the contract method 0x0e830e49.
//
// Solidity: function setEngine(address _engine) returns()
func (_PaymentVault *PaymentVaultSession) SetEngine(_engine common.Address) (*types.Transaction, error) {
	return _PaymentVault.Contract.SetEngine(&_PaymentVault.TransactOpts, _engine)
}

// SetEngine is a paid mutator transaction binding the contract method 0x0e830e49.
//
// Solidity: function setEngine(address _engine) returns()
func (_PaymentVault *PaymentVaultTransactorSession) SetEngine(_engine common.Address) (*types.Transaction, error) {
	return _PaymentVault.Contract.SetEngine(&_PaymentVault.TransactOpts, _engine)
}

// UpdateSettlementEngine is a paid mutator transaction binding the contract method 0xe9e2dbab.
//
// Solidity: function updateSettlementEngine(address _newEngine) returns()
func (_PaymentVault *PaymentVaultTransactor) UpdateSettlementEngine(opts *bind.TransactOpts, _newEngine common.Address) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "updateSettlementEngine", _newEngine)
}

// UpdateSettlementEngine is a paid mutator transaction binding the contract method 0xe9e2dbab.
//
// Solidity: function updateSettlementEngine(address _newEngine) returns()
func (_PaymentVault *PaymentVaultSession) UpdateSettlementEngine(_newEngine common.Address) (*types.Transaction, error) {
	return _PaymentVault.Contract.UpdateSettlementEngine(&_PaymentVault.TransactOpts, _newEngine)
}

// UpdateSettlementEngine is a paid mutator transaction binding the contract method 0xe9e2dbab.
//
// Solidity: function updateSettlementEngine(address _newEngine) returns()
func (_PaymentVault *PaymentVaultTransactorSession) UpdateSettlementEngine(_newEngine common.Address) (*types.Transaction, error) {
	return _PaymentVault.Contract.UpdateSettlementEngine(&_PaymentVault.TransactOpts, _newEngine)
}

// PaymentVaultPaymentRefundedIterator is returned from FilterPaymentRefunded and is used to iterate over the raw logs and unpacked data for PaymentRefunded events raised by the PaymentVault contract.
type PaymentVaultPaymentRefundedIterator struct {
	Event *PaymentVaultPaymentRefunded // Event containing the contract specifics and raw log

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
func (it *PaymentVaultPaymentRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentVaultPaymentRefunded)
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
		it.Event = new(PaymentVaultPaymentRefunded)
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
func (it *PaymentVaultPaymentRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentVaultPaymentRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentVaultPaymentRefunded represents a PaymentRefunded event raised by the PaymentVault contract.
type PaymentVaultPaymentRefunded struct {
	PacketId [32]byte
	Payer    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaymentRefunded is a free log retrieval operation binding the contract event 0x3b28817a65441bed1ac3a5496c6b71346aa302e80c7418c12e1c49eaa1721176.
//
// Solidity: event PaymentRefunded(bytes32 indexed packetId, address indexed payer, uint256 amount)
func (_PaymentVault *PaymentVaultFilterer) FilterPaymentRefunded(opts *bind.FilterOpts, packetId [][32]byte, payer []common.Address) (*PaymentVaultPaymentRefundedIterator, error) {

	var packetIdRule []interface{}
	for _, packetIdItem := range packetId {
		packetIdRule = append(packetIdRule, packetIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentVault.contract.FilterLogs(opts, "PaymentRefunded", packetIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentVaultPaymentRefundedIterator{contract: _PaymentVault.contract, event: "PaymentRefunded", logs: logs, sub: sub}, nil
}

// WatchPaymentRefunded is a free log subscription operation binding the contract event 0x3b28817a65441bed1ac3a5496c6b71346aa302e80c7418c12e1c49eaa1721176.
//
// Solidity: event PaymentRefunded(bytes32 indexed packetId, address indexed payer, uint256 amount)
func (_PaymentVault *PaymentVaultFilterer) WatchPaymentRefunded(opts *bind.WatchOpts, sink chan<- *PaymentVaultPaymentRefunded, packetId [][32]byte, payer []common.Address) (event.Subscription, error) {

	var packetIdRule []interface{}
	for _, packetIdItem := range packetId {
		packetIdRule = append(packetIdRule, packetIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentVault.contract.WatchLogs(opts, "PaymentRefunded", packetIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentVaultPaymentRefunded)
				if err := _PaymentVault.contract.UnpackLog(event, "PaymentRefunded", log); err != nil {
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

// ParsePaymentRefunded is a log parse operation binding the contract event 0x3b28817a65441bed1ac3a5496c6b71346aa302e80c7418c12e1c49eaa1721176.
//
// Solidity: event PaymentRefunded(bytes32 indexed packetId, address indexed payer, uint256 amount)
func (_PaymentVault *PaymentVaultFilterer) ParsePaymentRefunded(log types.Log) (*PaymentVaultPaymentRefunded, error) {
	event := new(PaymentVaultPaymentRefunded)
	if err := _PaymentVault.contract.UnpackLog(event, "PaymentRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
