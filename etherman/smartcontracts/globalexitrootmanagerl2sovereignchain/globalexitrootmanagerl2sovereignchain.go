// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalexitrootmanagerl2sovereignchain

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

// Globalexitrootmanagerl2sovereignchainMetaData contains all meta data concerning the Globalexitrootmanagerl2sovereignchain contract.
var Globalexitrootmanagerl2sovereignchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughGlobalExitRootsInserted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLastInsertedGlobalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"InsertGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"removedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"RemoveLastGlobalExitRoot\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"insertGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insertedGERCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"gersToRemove\",\"type\":\"bytes32[]\"}],\"name\":\"removeLastGlobalExitRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801562000010575f80fd5b5060405162000f9838038062000f988339818101604052810190620000369190620001af565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506200007c6200008360201b60201c565b50620002bb565b603460019054906101000a900460ff1615620000d6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000cd9062000263565b60405180910390fd5b60ff801660345f9054906101000a900460ff1660ff1614620001485760ff60345f6101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff6040516200013f9190620002a0565b60405180910390a15b565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000179826200014e565b9050919050565b6200018b816200016d565b811462000196575f80fd5b50565b5f81519050620001a98162000180565b92915050565b5f60208284031215620001c757620001c66200014a565b5b5f620001d68482850162000199565b91505092915050565b5f82825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e6974695f8201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b5f6200024b602783620001df565b91506200025882620001ef565b604082019050919050565b5f6020820190508181035f8301526200027c816200023d565b9050919050565b5f60ff82169050919050565b6200029a8162000283565b82525050565b5f602082019050620002b55f8301846200028f565b92915050565b608051610cbd620002db5f395f81816103b601526106e40152610cbd5ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c806357dfb5721161006457806357dfb5721461011b5780637c314ce3146101375780638bd0eb1c14610155578063a3c573eb14610173578063c4d66de81461019157610091565b806301fd90441461009557806312da06b2146100b3578063257b3632146100cf57806333d6247d146100ff575b5f80fd5b61009d6101ad565b6040516100aa91906108ba565b60405180910390f35b6100cd60048036038101906100c89190610905565b6101b3565b005b6100e960048036038101906100e49190610905565b6103a0565b6040516100f69190610948565b60405180910390f35b61011960048036038101906101149190610905565b6103b4565b005b610135600480360381019061013091906109c2565b610443565b005b61013f6106b6565b60405161014c9190610a4c565b60405180910390f35b61015d6106dc565b60405161016a9190610948565b60405180910390f35b61017b6106e2565b6040516101889190610a4c565b60405180910390f35b6101ab60048036038101906101a69190610a8f565b610706565b005b60015481565b5f73ffffffffffffffffffffffffffffffffffffffff16603460029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610272573373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161461026d576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102fa565b3373ffffffffffffffffffffffffffffffffffffffff16603460029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146102f9576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5f805f8381526020019081526020015f20540361036b5760355f815461031f90610ae7565b9190508190555f808381526020019081526020015f2081905550807fb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d560405160405180910390a261039d565b6040517f1f97a58200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b5f602052805f5260405f205f915090505481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610439576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060018190555050565b5f73ffffffffffffffffffffffffffffffffffffffff16603460029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610502573373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff16146104fd576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61058a565b3373ffffffffffffffffffffffffffffffffffffffff16603460029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610589576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5f6035549050808383905011156105cd576040517f56feb4f500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b838390508110156106b0575f8484838181106105ee576105ed610b2e565b5b9050602002013590505f805f8381526020019081526020015f20549050838114610644576040517fae765ff600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f808381526020019081526020015f205f905560355f81548092919061066990610b5b565b9190505550817f605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a260405160405180910390a2505080806106a890610ae7565b9150506105cf565b50505050565b603460029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60355481565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f603460019054906101000a900460ff161590508080156107385750600160345f9054906101000a900460ff1660ff16105b80610766575061074730610880565b1580156107655750600160345f9054906101000a900460ff1660ff16145b5b6107a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079c90610c02565b60405180910390fd5b600160345f6101000a81548160ff021916908360ff16021790555080156107e2576001603460016101000a81548160ff0219169083151502179055505b81603460026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550801561087c575f603460016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516108739190610c6e565b60405180910390a15b5050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f819050919050565b6108b4816108a2565b82525050565b5f6020820190506108cd5f8301846108ab565b92915050565b5f80fd5b5f80fd5b6108e4816108a2565b81146108ee575f80fd5b50565b5f813590506108ff816108db565b92915050565b5f6020828403121561091a576109196108d3565b5b5f610927848285016108f1565b91505092915050565b5f819050919050565b61094281610930565b82525050565b5f60208201905061095b5f830184610939565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261098257610981610961565b5b8235905067ffffffffffffffff81111561099f5761099e610965565b5b6020830191508360208202830111156109bb576109ba610969565b5b9250929050565b5f80602083850312156109d8576109d76108d3565b5b5f83013567ffffffffffffffff8111156109f5576109f46108d7565b5b610a018582860161096d565b92509250509250929050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a3682610a0d565b9050919050565b610a4681610a2c565b82525050565b5f602082019050610a5f5f830184610a3d565b92915050565b610a6e81610a2c565b8114610a78575f80fd5b50565b5f81359050610a8981610a65565b92915050565b5f60208284031215610aa457610aa36108d3565b5b5f610ab184828501610a7b565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610af182610930565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610b2357610b22610aba565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f610b6582610930565b91505f8203610b7757610b76610aba565b5b600182039050919050565b5f82825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f610bec602e83610b82565b9150610bf782610b92565b604082019050919050565b5f6020820190508181035f830152610c1981610be0565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f610c58610c53610c4e84610c20565b610c35565b610c29565b9050919050565b610c6881610c3e565b82525050565b5f602082019050610c815f830184610c5f565b9291505056fea264697066735822122074d4b5710dd22e5484705695828b56da48cad4b79877dd7ffc8ab8c31a8f171e64736f6c63430008140033",
}

// Globalexitrootmanagerl2sovereignchainABI is the input ABI used to generate the binding from.
// Deprecated: Use Globalexitrootmanagerl2sovereignchainMetaData.ABI instead.
var Globalexitrootmanagerl2sovereignchainABI = Globalexitrootmanagerl2sovereignchainMetaData.ABI

// Globalexitrootmanagerl2sovereignchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Globalexitrootmanagerl2sovereignchainMetaData.Bin instead.
var Globalexitrootmanagerl2sovereignchainBin = Globalexitrootmanagerl2sovereignchainMetaData.Bin

// DeployGlobalexitrootmanagerl2sovereignchain deploys a new Ethereum contract, binding an instance of Globalexitrootmanagerl2sovereignchain to it.
func DeployGlobalexitrootmanagerl2sovereignchain(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Globalexitrootmanagerl2sovereignchain, error) {
	parsed, err := Globalexitrootmanagerl2sovereignchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Globalexitrootmanagerl2sovereignchainBin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Globalexitrootmanagerl2sovereignchain{Globalexitrootmanagerl2sovereignchainCaller: Globalexitrootmanagerl2sovereignchainCaller{contract: contract}, Globalexitrootmanagerl2sovereignchainTransactor: Globalexitrootmanagerl2sovereignchainTransactor{contract: contract}, Globalexitrootmanagerl2sovereignchainFilterer: Globalexitrootmanagerl2sovereignchainFilterer{contract: contract}}, nil
}

// Globalexitrootmanagerl2sovereignchain is an auto generated Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchain struct {
	Globalexitrootmanagerl2sovereignchainCaller     // Read-only binding to the contract
	Globalexitrootmanagerl2sovereignchainTransactor // Write-only binding to the contract
	Globalexitrootmanagerl2sovereignchainFilterer   // Log filterer for contract events
}

// Globalexitrootmanagerl2sovereignchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootmanagerl2sovereignchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootmanagerl2sovereignchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Globalexitrootmanagerl2sovereignchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Globalexitrootmanagerl2sovereignchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Globalexitrootmanagerl2sovereignchainSession struct {
	Contract     *Globalexitrootmanagerl2sovereignchain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                          // Call options to use throughout this session
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// Globalexitrootmanagerl2sovereignchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Globalexitrootmanagerl2sovereignchainCallerSession struct {
	Contract *Globalexitrootmanagerl2sovereignchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                // Call options to use throughout this session
}

// Globalexitrootmanagerl2sovereignchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Globalexitrootmanagerl2sovereignchainTransactorSession struct {
	Contract     *Globalexitrootmanagerl2sovereignchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                // Transaction auth options to use throughout this session
}

// Globalexitrootmanagerl2sovereignchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainRaw struct {
	Contract *Globalexitrootmanagerl2sovereignchain // Generic contract binding to access the raw methods on
}

// Globalexitrootmanagerl2sovereignchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainCallerRaw struct {
	Contract *Globalexitrootmanagerl2sovereignchainCaller // Generic read-only contract binding to access the raw methods on
}

// Globalexitrootmanagerl2sovereignchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Globalexitrootmanagerl2sovereignchainTransactorRaw struct {
	Contract *Globalexitrootmanagerl2sovereignchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalexitrootmanagerl2sovereignchain creates a new instance of Globalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchain(address common.Address, backend bind.ContractBackend) (*Globalexitrootmanagerl2sovereignchain, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchain{Globalexitrootmanagerl2sovereignchainCaller: Globalexitrootmanagerl2sovereignchainCaller{contract: contract}, Globalexitrootmanagerl2sovereignchainTransactor: Globalexitrootmanagerl2sovereignchainTransactor{contract: contract}, Globalexitrootmanagerl2sovereignchainFilterer: Globalexitrootmanagerl2sovereignchainFilterer{contract: contract}}, nil
}

// NewGlobalexitrootmanagerl2sovereignchainCaller creates a new read-only instance of Globalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainCaller(address common.Address, caller bind.ContractCaller) (*Globalexitrootmanagerl2sovereignchainCaller, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainCaller{contract: contract}, nil
}

// NewGlobalexitrootmanagerl2sovereignchainTransactor creates a new write-only instance of Globalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainTransactor(address common.Address, transactor bind.ContractTransactor) (*Globalexitrootmanagerl2sovereignchainTransactor, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainTransactor{contract: contract}, nil
}

// NewGlobalexitrootmanagerl2sovereignchainFilterer creates a new log filterer instance of Globalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewGlobalexitrootmanagerl2sovereignchainFilterer(address common.Address, filterer bind.ContractFilterer) (*Globalexitrootmanagerl2sovereignchainFilterer, error) {
	contract, err := bindGlobalexitrootmanagerl2sovereignchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainFilterer{contract: contract}, nil
}

// bindGlobalexitrootmanagerl2sovereignchain binds a generic wrapper to an already deployed contract.
func bindGlobalexitrootmanagerl2sovereignchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Globalexitrootmanagerl2sovereignchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Globalexitrootmanagerl2sovereignchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Globalexitrootmanagerl2sovereignchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Globalexitrootmanagerl2sovereignchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Globalexitrootmanagerl2sovereignchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) BridgeAddress() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.BridgeAddress(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) BridgeAddress() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.BridgeAddress(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootMap(&_Globalexitrootmanagerl2sovereignchain.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootMap(&_Globalexitrootmanagerl2sovereignchain.CallOpts, arg0)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) GlobalExitRootUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) GlobalExitRootUpdater() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootUpdater() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) InsertedGERCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "insertedGERCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) InsertedGERCount() (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertedGERCount(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) InsertedGERCount() (*big.Int, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertedGERCount(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) LastRollupExitRoot() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.LastRollupExitRoot(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.LastRollupExitRoot(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) Initialize(opts *bind.TransactOpts, _globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "initialize", _globalExitRootUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) Initialize(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) Initialize(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) InsertGlobalExitRoot(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "insertGlobalExitRoot", _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertGlobalExitRoot(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertGlobalExitRoot(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _newRoot)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) RemoveLastGlobalExitRoots(opts *bind.TransactOpts, gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "removeLastGlobalExitRoots", gersToRemove)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) RemoveLastGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemoveLastGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) RemoveLastGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemoveLastGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.UpdateExitRoot(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.UpdateExitRoot(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, newRoot)
}

// Globalexitrootmanagerl2sovereignchainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainInitializedIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainInitialized // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainInitialized)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainInitialized)
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
func (it *Globalexitrootmanagerl2sovereignchainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainInitialized represents a Initialized event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterInitialized(opts *bind.FilterOpts) (*Globalexitrootmanagerl2sovereignchainInitializedIterator, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainInitializedIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainInitialized) (event.Subscription, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainInitialized)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseInitialized(log types.Log) (*Globalexitrootmanagerl2sovereignchainInitialized, error) {
	event := new(Globalexitrootmanagerl2sovereignchainInitialized)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator is returned from FilterInsertGlobalExitRoot and is used to iterate over the raw logs and unpacked data for InsertGlobalExitRoot events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
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
func (it *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot represents a InsertGlobalExitRoot event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot struct {
	NewGlobalExitRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInsertGlobalExitRoot is a free log retrieval operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterInsertGlobalExitRoot(opts *bind.FilterOpts, newGlobalExitRoot [][32]byte) (*Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "InsertGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchInsertGlobalExitRoot is a free log subscription operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchInsertGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot, newGlobalExitRoot [][32]byte) (event.Subscription, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "InsertGlobalExitRoot", log); err != nil {
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

// ParseInsertGlobalExitRoot is a log parse operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseInsertGlobalExitRoot(log types.Log) (*Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot, error) {
	event := new(Globalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "InsertGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator is returned from FilterRemoveLastGlobalExitRoot and is used to iterate over the raw logs and unpacked data for RemoveLastGlobalExitRoot events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
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
func (it *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot represents a RemoveLastGlobalExitRoot event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot struct {
	RemovedGlobalExitRoot [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLastGlobalExitRoot is a free log retrieval operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterRemoveLastGlobalExitRoot(opts *bind.FilterOpts, removedGlobalExitRoot [][32]byte) (*Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "RemoveLastGlobalExitRoot", removedGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "RemoveLastGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchRemoveLastGlobalExitRoot is a free log subscription operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchRemoveLastGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot, removedGlobalExitRoot [][32]byte) (event.Subscription, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "RemoveLastGlobalExitRoot", removedGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "RemoveLastGlobalExitRoot", log); err != nil {
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

// ParseRemoveLastGlobalExitRoot is a log parse operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseRemoveLastGlobalExitRoot(log types.Log) (*Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot, error) {
	event := new(Globalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "RemoveLastGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
