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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"packetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentRefunded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"}],\"name\":\"getPaymentDetails\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"createdAt\",\"type\":\"uint64\"},{\"internalType\":\"enumIPaymentVault.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"merchantIdentity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"merchantPayout\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"categoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"contractIGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_merchantIdentity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_merchantPayout\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_categoryId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_createdAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_s\",\"type\":\"uint256\"}],\"name\":\"lockFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"payments\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"createdAt\",\"type\":\"uint64\"},{\"internalType\":\"enumIPaymentVault.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"merchantIdentity\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"merchantPayout\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"categoryId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_packetIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"_settlementWallet\",\"type\":\"address\"}],\"name\":\"prepareSettlementBulk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_packetId\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_engine\",\"type\":\"address\"}],\"name\":\"setEngine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementEngine\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newEngine\",\"type\":\"address\"}],\"name\":\"updateSettlementEngine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516118f73803806118f783398181016040528101906100319190610105565b60015f819055508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505050610143565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100d4826100ab565b9050919050565b6100e4816100ca565b81146100ee575f5ffd5b50565b5f815190506100ff816100db565b92915050565b5f5f6040838503121561011b5761011a6100a7565b5b5f610128858286016100f1565b9250506020610139858286016100f1565b9150509250929050565b60805160a0516117626101955f395f818161045d0152818161079b015281816109510152610c1501525f81816102be015281816102fa015281816109a501528181610d010152610d3d01526117625ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c80635aa6e675116100645780635aa6e6751461014a5780635fd13996146101685780637249fbb6146101a0578063e6d0919d146101bc578063e9e2dbab146101da5761009c565b80630716326d146100a05780630e830e49146100d85780633e413bee146100f45780633f72ca4114610112578063477c5f011461012e575b5f5ffd5b6100ba60048036038101906100b591906110c2565b6101f6565b6040516100cf999897969594939291906111e8565b60405180910390f35b6100f260048036038101906100ed919061129d565b6102bc565b005b6100fc61045b565b6040516101099190611323565b60405180910390f35b61012c60048036038101906101279190611390565b61047f565b005b610148600480360381019061014391906114b5565b6107f3565b005b6101526109a3565b60405161015f9190611532565b60405180910390f35b610182600480360381019061017d91906110c2565b6109c7565b604051610197999897969594939291906111e8565b60405180910390f35b6101ba60048036038101906101b591906110c2565b610aac565b005b6101c4610cda565b6040516101d1919061154b565b60405180910390f35b6101f460048036038101906101ef919061129d565b610cff565b005b6002602052805f5260405f205f91509050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690805f0160149054906101000a900467ffffffffffffffff1690805f01601c9054906101000a900460ff1690806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154908060050154908060060154905089565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377fb8d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610361573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103859190611578565b336040518363ffffffff1660e01b81526004016103a39291906115a3565b602060405180830381865afa1580156103be573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103e291906115ff565b610418576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610505576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61050d610e9e565b5f60038111156105205761051f61114e565b5b60025f8681526020019081526020015f205f01601c9054906101000a900460ff1660038111156105535761055261114e565b5b1461058a576040517f23369fa600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518061012001604052808a73ffffffffffffffffffffffffffffffffffffffff1681526020018467ffffffffffffffff168152602001600160038111156105d6576105d561114e565b5b81526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018981526020018681526020018381526020018281525060025f8681526020019081526020015f205f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151815f0160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506040820151815f01601c6101000a81548160ff021916908360038111156106d6576106d561114e565b5b02179055506060820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816003015560c0820151816004015560e0820151816005015561010082015181600601559050506107e089308a7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610ee2909392919063ffffffff16565b6107e8610f64565b505050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610879576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610881610e9e565b5f5f90505f5f90505b84849050811015610949575f60025f8787858181106108ac576108ab61162a565b5b9050602002013581526020019081526020015f209050600160038111156108d6576108d561114e565b5b815f01601c9054906101000a900460ff1660038111156108f9576108f861114e565b5b0361093d576002815f01601c6101000a81548160ff021916908360038111156109255761092461114e565b5b021790555080600301548361093a9190611684565b92505b8160010191505061088a565b5061099582827f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610f6d9092919063ffffffff16565b5061099e610f64565b505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f5f5f5f5f5f5f5f60025f8c81526020019081526020015f209050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815f0160149054906101000a900467ffffffffffffffff16825f01601c9054906101000a900460ff16836001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846002015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168560030154866004015487600501548860060154995099509950995099509950995099509950509193959799909294969850565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b32576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b3a610e9e565b5f60025f8381526020019081526020015f20905060016003811115610b6257610b6161114e565b5b815f01601c9054906101000a900460ff166003811115610b8557610b8461114e565b5b14610bbc576040517f80cb55e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003815f01601c6101000a81548160ff02191690836003811115610be357610be261114e565b5b0217905550610c59815f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682600301547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610f6d9092919063ffffffff16565b805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16827f3b28817a65441bed1ac3a5496c6b71346aa302e80c7418c12e1c49eaa17211768360030154604051610cc691906116b7565b60405180910390a350610cd7610f64565b50565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166391d148547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166377fb8d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610da4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dc89190611578565b336040518363ffffffff1660e01b8152600401610de69291906115a3565b602060405180830381865afa158015610e01573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e2591906115ff565b610e5b576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60025f5403610ed9576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025f81905550565b610f5e848573ffffffffffffffffffffffffffffffffffffffff166323b872dd868686604051602401610f17939291906116d0565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610fec565b50505050565b60015f81905550565b610fe7838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8585604051602401610fa0929190611705565b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610fec565b505050565b5f5f60205f8451602086015f885af18061100b576040513d5f823e3d81fd5b3d92505f519150505f821461102457600181141561103f565b5f8473ffffffffffffffffffffffffffffffffffffffff163b145b1561108157836040517f5274afe7000000000000000000000000000000000000000000000000000000008152600401611078919061154b565b60405180910390fd5b50505050565b5f5ffd5b5f5ffd5b5f819050919050565b6110a18161108f565b81146110ab575f5ffd5b50565b5f813590506110bc81611098565b92915050565b5f602082840312156110d7576110d6611087565b5b5f6110e4848285016110ae565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611116826110ed565b9050919050565b6111268161110c565b82525050565b5f67ffffffffffffffff82169050919050565b6111488161112c565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6004811061118c5761118b61114e565b5b50565b5f81905061119c8261117b565b919050565b5f6111ab8261118f565b9050919050565b6111bb816111a1565b82525050565b5f819050919050565b6111d3816111c1565b82525050565b6111e28161108f565b82525050565b5f610120820190506111fc5f83018c61111d565b611209602083018b61113f565b611216604083018a6111b2565b611223606083018961111d565b611230608083018861111d565b61123d60a08301876111ca565b61124a60c08301866111d9565b61125760e08301856111ca565b6112656101008301846111ca565b9a9950505050505050505050565b61127c8161110c565b8114611286575f5ffd5b50565b5f8135905061129781611273565b92915050565b5f602082840312156112b2576112b1611087565b5b5f6112bf84828501611289565b91505092915050565b5f819050919050565b5f6112eb6112e66112e1846110ed565b6112c8565b6110ed565b9050919050565b5f6112fc826112d1565b9050919050565b5f61130d826112f2565b9050919050565b61131d81611303565b82525050565b5f6020820190506113365f830184611314565b92915050565b611345816111c1565b811461134f575f5ffd5b50565b5f813590506113608161133c565b92915050565b61136f8161112c565b8114611379575f5ffd5b50565b5f8135905061138a81611366565b92915050565b5f5f5f5f5f5f5f5f5f6101208a8c0312156113ae576113ad611087565b5b5f6113bb8c828d01611289565b99505060206113cc8c828d01611352565b98505060406113dd8c828d01611289565b97505060606113ee8c828d01611289565b96505060806113ff8c828d016110ae565b95505060a06114108c828d016110ae565b94505060c06114218c828d0161137c565b93505060e06114328c828d01611352565b9250506101006114448c828d01611352565b9150509295985092959850929598565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261147557611474611454565b5b8235905067ffffffffffffffff81111561149257611491611458565b5b6020830191508360208202830111156114ae576114ad61145c565b5b9250929050565b5f5f5f604084860312156114cc576114cb611087565b5b5f84013567ffffffffffffffff8111156114e9576114e861108b565b5b6114f586828701611460565b9350935050602061150886828701611289565b9150509250925092565b5f61151c826112f2565b9050919050565b61152c81611512565b82525050565b5f6020820190506115455f830184611523565b92915050565b5f60208201905061155e5f83018461111d565b92915050565b5f8151905061157281611098565b92915050565b5f6020828403121561158d5761158c611087565b5b5f61159a84828501611564565b91505092915050565b5f6040820190506115b65f8301856111d9565b6115c3602083018461111d565b9392505050565b5f8115159050919050565b6115de816115ca565b81146115e8575f5ffd5b50565b5f815190506115f9816115d5565b92915050565b5f6020828403121561161457611613611087565b5b5f611621848285016115eb565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61168e826111c1565b9150611699836111c1565b92508282019050808211156116b1576116b0611657565b5b92915050565b5f6020820190506116ca5f8301846111ca565b92915050565b5f6060820190506116e35f83018661111d565b6116f0602083018561111d565b6116fd60408301846111ca565b949350505050565b5f6040820190506117185f83018561111d565b61172560208301846111ca565b939250505056fea2646970667358221220b6ac5419d6468fa8362c5d49055dc11d1405d8ed7f3f25ac60ec16ea0278e5af64736f6c63430008210033",
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
// Solidity: function getPaymentDetails(bytes32 _packetId) view returns(address payer, uint64 createdAt, uint8 status, address merchantIdentity, address merchantPayout, uint256 amount, bytes32 categoryId, uint256 r, uint256 s)
func (_PaymentVault *PaymentVaultCaller) GetPaymentDetails(opts *bind.CallOpts, _packetId [32]byte) (struct {
	Payer            common.Address
	CreatedAt        uint64
	Status           uint8
	MerchantIdentity common.Address
	MerchantPayout   common.Address
	Amount           *big.Int
	CategoryId       [32]byte
	R                *big.Int
	S                *big.Int
}, error) {
	var out []interface{}
	err := _PaymentVault.contract.Call(opts, &out, "getPaymentDetails", _packetId)

	outstruct := new(struct {
		Payer            common.Address
		CreatedAt        uint64
		Status           uint8
		MerchantIdentity common.Address
		MerchantPayout   common.Address
		Amount           *big.Int
		CategoryId       [32]byte
		R                *big.Int
		S                *big.Int
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
	outstruct.R = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.S = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPaymentDetails is a free data retrieval call binding the contract method 0x5fd13996.
//
// Solidity: function getPaymentDetails(bytes32 _packetId) view returns(address payer, uint64 createdAt, uint8 status, address merchantIdentity, address merchantPayout, uint256 amount, bytes32 categoryId, uint256 r, uint256 s)
func (_PaymentVault *PaymentVaultSession) GetPaymentDetails(_packetId [32]byte) (struct {
	Payer            common.Address
	CreatedAt        uint64
	Status           uint8
	MerchantIdentity common.Address
	MerchantPayout   common.Address
	Amount           *big.Int
	CategoryId       [32]byte
	R                *big.Int
	S                *big.Int
}, error) {
	return _PaymentVault.Contract.GetPaymentDetails(&_PaymentVault.CallOpts, _packetId)
}

// GetPaymentDetails is a free data retrieval call binding the contract method 0x5fd13996.
//
// Solidity: function getPaymentDetails(bytes32 _packetId) view returns(address payer, uint64 createdAt, uint8 status, address merchantIdentity, address merchantPayout, uint256 amount, bytes32 categoryId, uint256 r, uint256 s)
func (_PaymentVault *PaymentVaultCallerSession) GetPaymentDetails(_packetId [32]byte) (struct {
	Payer            common.Address
	CreatedAt        uint64
	Status           uint8
	MerchantIdentity common.Address
	MerchantPayout   common.Address
	Amount           *big.Int
	CategoryId       [32]byte
	R                *big.Int
	S                *big.Int
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
// Solidity: function payments(bytes32 ) view returns(address payer, uint64 createdAt, uint8 status, address merchantIdentity, address merchantPayout, uint256 amount, bytes32 categoryId, uint256 r, uint256 s)
func (_PaymentVault *PaymentVaultCaller) Payments(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Payer            common.Address
	CreatedAt        uint64
	Status           uint8
	MerchantIdentity common.Address
	MerchantPayout   common.Address
	Amount           *big.Int
	CategoryId       [32]byte
	R                *big.Int
	S                *big.Int
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
		R                *big.Int
		S                *big.Int
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
	outstruct.R = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.S = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Payments is a free data retrieval call binding the contract method 0x0716326d.
//
// Solidity: function payments(bytes32 ) view returns(address payer, uint64 createdAt, uint8 status, address merchantIdentity, address merchantPayout, uint256 amount, bytes32 categoryId, uint256 r, uint256 s)
func (_PaymentVault *PaymentVaultSession) Payments(arg0 [32]byte) (struct {
	Payer            common.Address
	CreatedAt        uint64
	Status           uint8
	MerchantIdentity common.Address
	MerchantPayout   common.Address
	Amount           *big.Int
	CategoryId       [32]byte
	R                *big.Int
	S                *big.Int
}, error) {
	return _PaymentVault.Contract.Payments(&_PaymentVault.CallOpts, arg0)
}

// Payments is a free data retrieval call binding the contract method 0x0716326d.
//
// Solidity: function payments(bytes32 ) view returns(address payer, uint64 createdAt, uint8 status, address merchantIdentity, address merchantPayout, uint256 amount, bytes32 categoryId, uint256 r, uint256 s)
func (_PaymentVault *PaymentVaultCallerSession) Payments(arg0 [32]byte) (struct {
	Payer            common.Address
	CreatedAt        uint64
	Status           uint8
	MerchantIdentity common.Address
	MerchantPayout   common.Address
	Amount           *big.Int
	CategoryId       [32]byte
	R                *big.Int
	S                *big.Int
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

// LockFunds is a paid mutator transaction binding the contract method 0x3f72ca41.
//
// Solidity: function lockFunds(address _payer, uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId, uint64 _createdAt, uint256 _r, uint256 _s) returns()
func (_PaymentVault *PaymentVaultTransactor) LockFunds(opts *bind.TransactOpts, _payer common.Address, _amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte, _createdAt uint64, _r *big.Int, _s *big.Int) (*types.Transaction, error) {
	return _PaymentVault.contract.Transact(opts, "lockFunds", _payer, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId, _createdAt, _r, _s)
}

// LockFunds is a paid mutator transaction binding the contract method 0x3f72ca41.
//
// Solidity: function lockFunds(address _payer, uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId, uint64 _createdAt, uint256 _r, uint256 _s) returns()
func (_PaymentVault *PaymentVaultSession) LockFunds(_payer common.Address, _amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte, _createdAt uint64, _r *big.Int, _s *big.Int) (*types.Transaction, error) {
	return _PaymentVault.Contract.LockFunds(&_PaymentVault.TransactOpts, _payer, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId, _createdAt, _r, _s)
}

// LockFunds is a paid mutator transaction binding the contract method 0x3f72ca41.
//
// Solidity: function lockFunds(address _payer, uint256 _amount, address _merchantIdentity, address _merchantPayout, bytes32 _categoryId, bytes32 _packetId, uint64 _createdAt, uint256 _r, uint256 _s) returns()
func (_PaymentVault *PaymentVaultTransactorSession) LockFunds(_payer common.Address, _amount *big.Int, _merchantIdentity common.Address, _merchantPayout common.Address, _categoryId [32]byte, _packetId [32]byte, _createdAt uint64, _r *big.Int, _s *big.Int) (*types.Transaction, error) {
	return _PaymentVault.Contract.LockFunds(&_PaymentVault.TransactOpts, _payer, _amount, _merchantIdentity, _merchantPayout, _categoryId, _packetId, _createdAt, _r, _s)
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
