// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oldglobalexitrootmanagerl2sovereignchain

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

// Oldglobalexitrootmanagerl2sovereignchainMetaData contains all meta data concerning the Oldglobalexitrootmanagerl2sovereignchain contract.
var Oldglobalexitrootmanagerl2sovereignchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughGlobalExitRootsInserted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLastInsertedGlobalExitRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"InsertGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"removedGlobalExitRoot\",\"type\":\"bytes32\"}],\"name\":\"RemoveLastGlobalExitRoot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGlobalExitRootRemover\",\"type\":\"address\"}],\"name\":\"SetGlobalExitRootRemover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"SetGlobalExitRootUpdater\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootRemover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalExitRootRemover\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"insertGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insertedGERCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"gersToRemove\",\"type\":\"bytes32[]\"}],\"name\":\"removeLastGlobalExitRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootRemover\",\"type\":\"address\"}],\"name\":\"setGlobalExitRootRemover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"setGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b5060405161096d38038061096d83398101604081905261002e91610108565b6001600160a01b038116608052610043610049565b50610135565b603454610100900460ff16156100b55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60345460ff90811614610106576034805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610118575f80fd5b81516001600160a01b038116811461012e575f80fd5b9392505050565b6080516108196101545f395f81816101a301526102cb01526108195ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c80636da0e4ab1161006e5780636da0e4ab1461013e5780637c314ce3146101515780638bd0eb1c1461018257806391eb796d1461018b578063a3c573eb1461019e578063d0267f39146101c5575f80fd5b806301fd9044146100b557806312da06b2146100d1578063257b3632146100e657806333d6247d14610105578063485cc9551461011857806357dfb5721461012b575b5f80fd5b6100be60015481565b6040519081526020015b60405180910390f35b6100e46100df36600461069c565b6101d8565b005b6100be6100f436600461069c565b5f6020819052908152604090205481565b6100e461011336600461069c565b6102c0565b6100e46101263660046106ce565b61030e565b6100e46101393660046106ff565b61045e565b6100e461014c36600461076e565b61056a565b60345461016a906201000090046001600160a01b031681565b6040516001600160a01b0390911681526020016100c8565b6100be60365481565b60355461016a906001600160a01b031681565b61016a7f000000000000000000000000000000000000000000000000000000000000000081565b6100e46101d336600461076e565b610629565b6034546201000090046001600160a01b03166102135741331461020e576040516363ac7e0d60e11b815260040160405180910390fd5b610244565b6034546201000090046001600160a01b03163314610244576040516363ac7e0d60e11b815260040160405180910390fd5b5f8181526020819052604081205490036102a75760365f8154610266906107a2565b91829055505f8281526020819052604080822092909255905182917fb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d591a250565b604051630fcbd2c160e11b815260040160405180910390fd5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103095760405163b49365dd60e01b815260040160405180910390fd5b600155565b603454610100900460ff161580801561032e5750603454600160ff909116105b806103485750303b158015610348575060345460ff166001145b6103af5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b6034805460ff1916600117905580156103d2576034805461ff0019166101001790555b6034805462010000600160b01b031916620100006001600160a01b038681169190910291909117909155603580546001600160a01b0319169184169190911790558015610459576034805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6035546001600160a01b031633146104895760405163a34ddeb160e01b815260040160405180910390fd5b603654808211156104ad576040516356feb4f560e01b815260040160405180910390fd5b5f5b82811015610562575f8484838181106104ca576104ca6107ba565b9050602002013590505f805f8381526020019081526020015f205490508381146105075760405163573b2ffb60e11b815260040160405180910390fd5b5f828152602081905260408120558361051f816107ce565b6040519095508391507f605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2905f90a25050808061055a906107a2565b9150506104af565b506036555050565b6034546201000090046001600160a01b03166105a5574133146105a0576040516363ac7e0d60e11b815260040160405180910390fd5b6105d6565b6034546201000090046001600160a01b031633146105d6576040516363ac7e0d60e11b815260040160405180910390fd5b6034805462010000600160b01b031916620100006001600160a01b038416908102919091179091556040517f992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c905f90a250565b6035546001600160a01b031633146106545760405163a34ddeb160e01b815260040160405180910390fd5b603580546001600160a01b0319166001600160a01b0383169081179091556040517eb4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758905f90a250565b5f602082840312156106ac575f80fd5b5035919050565b80356001600160a01b03811681146106c9575f80fd5b919050565b5f80604083850312156106df575f80fd5b6106e8836106b3565b91506106f6602084016106b3565b90509250929050565b5f8060208385031215610710575f80fd5b823567ffffffffffffffff80821115610727575f80fd5b818501915085601f83011261073a575f80fd5b813581811115610748575f80fd5b8660208260051b850101111561075c575f80fd5b60209290920196919550909350505050565b5f6020828403121561077e575f80fd5b610787826106b3565b9392505050565b634e487b7160e01b5f52601160045260245ffd5b5f600182016107b3576107b361078e565b5060010190565b634e487b7160e01b5f52603260045260245ffd5b5f816107dc576107dc61078e565b505f19019056fea264697066735822122028489e152b03d240177e44e3649d90f4b6305fd1183026889eb37073f90b4c3864736f6c63430008140033",
}

// Oldglobalexitrootmanagerl2sovereignchainABI is the input ABI used to generate the binding from.
// Deprecated: Use Oldglobalexitrootmanagerl2sovereignchainMetaData.ABI instead.
var Oldglobalexitrootmanagerl2sovereignchainABI = Oldglobalexitrootmanagerl2sovereignchainMetaData.ABI

// Oldglobalexitrootmanagerl2sovereignchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Oldglobalexitrootmanagerl2sovereignchainMetaData.Bin instead.
var Oldglobalexitrootmanagerl2sovereignchainBin = Oldglobalexitrootmanagerl2sovereignchainMetaData.Bin

// DeployOldglobalexitrootmanagerl2sovereignchain deploys a new Ethereum contract, binding an instance of Oldglobalexitrootmanagerl2sovereignchain to it.
func DeployOldglobalexitrootmanagerl2sovereignchain(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Oldglobalexitrootmanagerl2sovereignchain, error) {
	parsed, err := Oldglobalexitrootmanagerl2sovereignchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Oldglobalexitrootmanagerl2sovereignchainBin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oldglobalexitrootmanagerl2sovereignchain{Oldglobalexitrootmanagerl2sovereignchainCaller: Oldglobalexitrootmanagerl2sovereignchainCaller{contract: contract}, Oldglobalexitrootmanagerl2sovereignchainTransactor: Oldglobalexitrootmanagerl2sovereignchainTransactor{contract: contract}, Oldglobalexitrootmanagerl2sovereignchainFilterer: Oldglobalexitrootmanagerl2sovereignchainFilterer{contract: contract}}, nil
}

// Oldglobalexitrootmanagerl2sovereignchain is an auto generated Go binding around an Ethereum contract.
type Oldglobalexitrootmanagerl2sovereignchain struct {
	Oldglobalexitrootmanagerl2sovereignchainCaller     // Read-only binding to the contract
	Oldglobalexitrootmanagerl2sovereignchainTransactor // Write-only binding to the contract
	Oldglobalexitrootmanagerl2sovereignchainFilterer   // Log filterer for contract events
}

// Oldglobalexitrootmanagerl2sovereignchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type Oldglobalexitrootmanagerl2sovereignchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Oldglobalexitrootmanagerl2sovereignchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Oldglobalexitrootmanagerl2sovereignchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Oldglobalexitrootmanagerl2sovereignchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Oldglobalexitrootmanagerl2sovereignchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Oldglobalexitrootmanagerl2sovereignchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Oldglobalexitrootmanagerl2sovereignchainSession struct {
	Contract     *Oldglobalexitrootmanagerl2sovereignchain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                             // Call options to use throughout this session
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// Oldglobalexitrootmanagerl2sovereignchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Oldglobalexitrootmanagerl2sovereignchainCallerSession struct {
	Contract *Oldglobalexitrootmanagerl2sovereignchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                                   // Call options to use throughout this session
}

// Oldglobalexitrootmanagerl2sovereignchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Oldglobalexitrootmanagerl2sovereignchainTransactorSession struct {
	Contract     *Oldglobalexitrootmanagerl2sovereignchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                                   // Transaction auth options to use throughout this session
}

// Oldglobalexitrootmanagerl2sovereignchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type Oldglobalexitrootmanagerl2sovereignchainRaw struct {
	Contract *Oldglobalexitrootmanagerl2sovereignchain // Generic contract binding to access the raw methods on
}

// Oldglobalexitrootmanagerl2sovereignchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Oldglobalexitrootmanagerl2sovereignchainCallerRaw struct {
	Contract *Oldglobalexitrootmanagerl2sovereignchainCaller // Generic read-only contract binding to access the raw methods on
}

// Oldglobalexitrootmanagerl2sovereignchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Oldglobalexitrootmanagerl2sovereignchainTransactorRaw struct {
	Contract *Oldglobalexitrootmanagerl2sovereignchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOldglobalexitrootmanagerl2sovereignchain creates a new instance of Oldglobalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewOldglobalexitrootmanagerl2sovereignchain(address common.Address, backend bind.ContractBackend) (*Oldglobalexitrootmanagerl2sovereignchain, error) {
	contract, err := bindOldglobalexitrootmanagerl2sovereignchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oldglobalexitrootmanagerl2sovereignchain{Oldglobalexitrootmanagerl2sovereignchainCaller: Oldglobalexitrootmanagerl2sovereignchainCaller{contract: contract}, Oldglobalexitrootmanagerl2sovereignchainTransactor: Oldglobalexitrootmanagerl2sovereignchainTransactor{contract: contract}, Oldglobalexitrootmanagerl2sovereignchainFilterer: Oldglobalexitrootmanagerl2sovereignchainFilterer{contract: contract}}, nil
}

// NewOldglobalexitrootmanagerl2sovereignchainCaller creates a new read-only instance of Oldglobalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewOldglobalexitrootmanagerl2sovereignchainCaller(address common.Address, caller bind.ContractCaller) (*Oldglobalexitrootmanagerl2sovereignchainCaller, error) {
	contract, err := bindOldglobalexitrootmanagerl2sovereignchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Oldglobalexitrootmanagerl2sovereignchainCaller{contract: contract}, nil
}

// NewOldglobalexitrootmanagerl2sovereignchainTransactor creates a new write-only instance of Oldglobalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewOldglobalexitrootmanagerl2sovereignchainTransactor(address common.Address, transactor bind.ContractTransactor) (*Oldglobalexitrootmanagerl2sovereignchainTransactor, error) {
	contract, err := bindOldglobalexitrootmanagerl2sovereignchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Oldglobalexitrootmanagerl2sovereignchainTransactor{contract: contract}, nil
}

// NewOldglobalexitrootmanagerl2sovereignchainFilterer creates a new log filterer instance of Oldglobalexitrootmanagerl2sovereignchain, bound to a specific deployed contract.
func NewOldglobalexitrootmanagerl2sovereignchainFilterer(address common.Address, filterer bind.ContractFilterer) (*Oldglobalexitrootmanagerl2sovereignchainFilterer, error) {
	contract, err := bindOldglobalexitrootmanagerl2sovereignchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Oldglobalexitrootmanagerl2sovereignchainFilterer{contract: contract}, nil
}

// bindOldglobalexitrootmanagerl2sovereignchain binds a generic wrapper to an already deployed contract.
func bindOldglobalexitrootmanagerl2sovereignchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Oldglobalexitrootmanagerl2sovereignchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.Oldglobalexitrootmanagerl2sovereignchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.Oldglobalexitrootmanagerl2sovereignchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.Oldglobalexitrootmanagerl2sovereignchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oldglobalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) BridgeAddress() (common.Address, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.BridgeAddress(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCallerSession) BridgeAddress() (common.Address, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.BridgeAddress(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCaller) GlobalExitRootMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Oldglobalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootMap(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts, arg0)
}

// GlobalExitRootMap is a free data retrieval call binding the contract method 0x257b3632.
//
// Solidity: function globalExitRootMap(bytes32 ) view returns(uint256)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootMap(arg0 [32]byte) (*big.Int, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootMap(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts, arg0)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCaller) GlobalExitRootRemover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oldglobalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootRemover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) GlobalExitRootRemover() (common.Address, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootRemover(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootRemover() (common.Address, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootRemover(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCaller) GlobalExitRootUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oldglobalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) GlobalExitRootUpdater() (common.Address, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootUpdater(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootUpdater is a free data retrieval call binding the contract method 0x7c314ce3.
//
// Solidity: function globalExitRootUpdater() view returns(address)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootUpdater() (common.Address, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootUpdater(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCaller) InsertedGERCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Oldglobalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "insertedGERCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) InsertedGERCount() (*big.Int, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.InsertedGERCount(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// InsertedGERCount is a free data retrieval call binding the contract method 0x8bd0eb1c.
//
// Solidity: function insertedGERCount() view returns(uint256)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCallerSession) InsertedGERCount() (*big.Int, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.InsertedGERCount(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCaller) LastRollupExitRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Oldglobalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "lastRollupExitRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) LastRollupExitRoot() ([32]byte, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.LastRollupExitRoot(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// LastRollupExitRoot is a free data retrieval call binding the contract method 0x01fd9044.
//
// Solidity: function lastRollupExitRoot() view returns(bytes32)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainCallerSession) LastRollupExitRoot() ([32]byte, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.LastRollupExitRoot(&_Oldglobalexitrootmanagerl2sovereignchain.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactor) Initialize(opts *bind.TransactOpts, _globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.contract.Transact(opts, "initialize", _globalExitRootUpdater, _globalExitRootRemover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) Initialize(_globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater, _globalExitRootRemover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactorSession) Initialize(_globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater, _globalExitRootRemover)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactor) InsertGlobalExitRoot(opts *bind.TransactOpts, _newRoot [32]byte) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.contract.Transact(opts, "insertGlobalExitRoot", _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.InsertGlobalExitRoot(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, _newRoot)
}

// InsertGlobalExitRoot is a paid mutator transaction binding the contract method 0x12da06b2.
//
// Solidity: function insertGlobalExitRoot(bytes32 _newRoot) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactorSession) InsertGlobalExitRoot(_newRoot [32]byte) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.InsertGlobalExitRoot(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, _newRoot)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactor) RemoveLastGlobalExitRoots(opts *bind.TransactOpts, gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.contract.Transact(opts, "removeLastGlobalExitRoots", gersToRemove)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) RemoveLastGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.RemoveLastGlobalExitRoots(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
}

// RemoveLastGlobalExitRoots is a paid mutator transaction binding the contract method 0x57dfb572.
//
// Solidity: function removeLastGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactorSession) RemoveLastGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.RemoveLastGlobalExitRoots(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactor) SetGlobalExitRootRemover(opts *bind.TransactOpts, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.contract.Transact(opts, "setGlobalExitRootRemover", _globalExitRootRemover)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) SetGlobalExitRootRemover(_globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootRemover(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootRemover)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactorSession) SetGlobalExitRootRemover(_globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootRemover(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootRemover)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactor) SetGlobalExitRootUpdater(opts *bind.TransactOpts, _globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.contract.Transact(opts, "setGlobalExitRootUpdater", _globalExitRootUpdater)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) SetGlobalExitRootUpdater(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootUpdater(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactorSession) SetGlobalExitRootUpdater(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootUpdater(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactor) UpdateExitRoot(opts *bind.TransactOpts, newRoot [32]byte) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.contract.Transact(opts, "updateExitRoot", newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.UpdateExitRoot(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, newRoot)
}

// UpdateExitRoot is a paid mutator transaction binding the contract method 0x33d6247d.
//
// Solidity: function updateExitRoot(bytes32 newRoot) returns()
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainTransactorSession) UpdateExitRoot(newRoot [32]byte) (*types.Transaction, error) {
	return _Oldglobalexitrootmanagerl2sovereignchain.Contract.UpdateExitRoot(&_Oldglobalexitrootmanagerl2sovereignchain.TransactOpts, newRoot)
}

// Oldglobalexitrootmanagerl2sovereignchainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainInitializedIterator struct {
	Event *Oldglobalexitrootmanagerl2sovereignchainInitialized // Event containing the contract specifics and raw log

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
func (it *Oldglobalexitrootmanagerl2sovereignchainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oldglobalexitrootmanagerl2sovereignchainInitialized)
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
		it.Event = new(Oldglobalexitrootmanagerl2sovereignchainInitialized)
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
func (it *Oldglobalexitrootmanagerl2sovereignchainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oldglobalexitrootmanagerl2sovereignchainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oldglobalexitrootmanagerl2sovereignchainInitialized represents a Initialized event raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) FilterInitialized(opts *bind.FilterOpts) (*Oldglobalexitrootmanagerl2sovereignchainInitializedIterator, error) {

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Oldglobalexitrootmanagerl2sovereignchainInitializedIterator{contract: _Oldglobalexitrootmanagerl2sovereignchain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Oldglobalexitrootmanagerl2sovereignchainInitialized) (event.Subscription, error) {

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oldglobalexitrootmanagerl2sovereignchainInitialized)
				if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) ParseInitialized(log types.Log) (*Oldglobalexitrootmanagerl2sovereignchainInitialized, error) {
	event := new(Oldglobalexitrootmanagerl2sovereignchainInitialized)
	if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator is returned from FilterInsertGlobalExitRoot and is used to iterate over the raw logs and unpacked data for InsertGlobalExitRoot events raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator struct {
	Event *Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
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
		it.Event = new(Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
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
func (it *Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRoot represents a InsertGlobalExitRoot event raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRoot struct {
	NewGlobalExitRoot [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInsertGlobalExitRoot is a free log retrieval operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) FilterInsertGlobalExitRoot(opts *bind.FilterOpts, newGlobalExitRoot [][32]byte) (*Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRootIterator{contract: _Oldglobalexitrootmanagerl2sovereignchain.contract, event: "InsertGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchInsertGlobalExitRoot is a free log subscription operation binding the contract event 0xb1b866fe5fac68e8f1a4ab2520c7a6b493a954934bbd0f054bd91d6674a4c0d5.
//
// Solidity: event InsertGlobalExitRoot(bytes32 indexed newGlobalExitRoot)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) WatchInsertGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRoot, newGlobalExitRoot [][32]byte) (event.Subscription, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "InsertGlobalExitRoot", newGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
				if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "InsertGlobalExitRoot", log); err != nil {
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
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) ParseInsertGlobalExitRoot(log types.Log) (*Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRoot, error) {
	event := new(Oldglobalexitrootmanagerl2sovereignchainInsertGlobalExitRoot)
	if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "InsertGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator is returned from FilterRemoveLastGlobalExitRoot and is used to iterate over the raw logs and unpacked data for RemoveLastGlobalExitRoot events raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator struct {
	Event *Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot // Event containing the contract specifics and raw log

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
func (it *Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
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
		it.Event = new(Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
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
func (it *Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot represents a RemoveLastGlobalExitRoot event raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot struct {
	RemovedGlobalExitRoot [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLastGlobalExitRoot is a free log retrieval operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) FilterRemoveLastGlobalExitRoot(opts *bind.FilterOpts, removedGlobalExitRoot [][32]byte) (*Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "RemoveLastGlobalExitRoot", removedGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return &Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRootIterator{contract: _Oldglobalexitrootmanagerl2sovereignchain.contract, event: "RemoveLastGlobalExitRoot", logs: logs, sub: sub}, nil
}

// WatchRemoveLastGlobalExitRoot is a free log subscription operation binding the contract event 0x605764d0b65b62ecf05dc90f674a00a2e2531fabaf120fdde65790e407fcb7a2.
//
// Solidity: event RemoveLastGlobalExitRoot(bytes32 indexed removedGlobalExitRoot)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) WatchRemoveLastGlobalExitRoot(opts *bind.WatchOpts, sink chan<- *Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot, removedGlobalExitRoot [][32]byte) (event.Subscription, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "RemoveLastGlobalExitRoot", removedGlobalExitRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
				if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "RemoveLastGlobalExitRoot", log); err != nil {
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
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) ParseRemoveLastGlobalExitRoot(log types.Log) (*Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot, error) {
	event := new(Oldglobalexitrootmanagerl2sovereignchainRemoveLastGlobalExitRoot)
	if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "RemoveLastGlobalExitRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator is returned from FilterSetGlobalExitRootRemover and is used to iterate over the raw logs and unpacked data for SetGlobalExitRootRemover events raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator struct {
	Event *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover // Event containing the contract specifics and raw log

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
func (it *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
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
		it.Event = new(Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
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
func (it *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover represents a SetGlobalExitRootRemover event raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover struct {
	NewGlobalExitRootRemover common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalExitRootRemover is a free log retrieval operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) FilterSetGlobalExitRootRemover(opts *bind.FilterOpts, newGlobalExitRootRemover []common.Address) (*Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator, error) {

	var newGlobalExitRootRemoverRule []interface{}
	for _, newGlobalExitRootRemoverItem := range newGlobalExitRootRemover {
		newGlobalExitRootRemoverRule = append(newGlobalExitRootRemoverRule, newGlobalExitRootRemoverItem)
	}

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "SetGlobalExitRootRemover", newGlobalExitRootRemoverRule)
	if err != nil {
		return nil, err
	}
	return &Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator{contract: _Oldglobalexitrootmanagerl2sovereignchain.contract, event: "SetGlobalExitRootRemover", logs: logs, sub: sub}, nil
}

// WatchSetGlobalExitRootRemover is a free log subscription operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) WatchSetGlobalExitRootRemover(opts *bind.WatchOpts, sink chan<- *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover, newGlobalExitRootRemover []common.Address) (event.Subscription, error) {

	var newGlobalExitRootRemoverRule []interface{}
	for _, newGlobalExitRootRemoverItem := range newGlobalExitRootRemover {
		newGlobalExitRootRemoverRule = append(newGlobalExitRootRemoverRule, newGlobalExitRootRemoverItem)
	}

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "SetGlobalExitRootRemover", newGlobalExitRootRemoverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
				if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootRemover", log); err != nil {
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

// ParseSetGlobalExitRootRemover is a log parse operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) ParseSetGlobalExitRootRemover(log types.Log) (*Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover, error) {
	event := new(Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
	if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootRemover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator is returned from FilterSetGlobalExitRootUpdater and is used to iterate over the raw logs and unpacked data for SetGlobalExitRootUpdater events raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator struct {
	Event *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater // Event containing the contract specifics and raw log

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
func (it *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
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
		it.Event = new(Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
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
func (it *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater represents a SetGlobalExitRootUpdater event raised by the Oldglobalexitrootmanagerl2sovereignchain contract.
type Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater struct {
	NewGlobalExitRootUpdater common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalExitRootUpdater is a free log retrieval operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) FilterSetGlobalExitRootUpdater(opts *bind.FilterOpts, newGlobalExitRootUpdater []common.Address) (*Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator, error) {

	var newGlobalExitRootUpdaterRule []interface{}
	for _, newGlobalExitRootUpdaterItem := range newGlobalExitRootUpdater {
		newGlobalExitRootUpdaterRule = append(newGlobalExitRootUpdaterRule, newGlobalExitRootUpdaterItem)
	}

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "SetGlobalExitRootUpdater", newGlobalExitRootUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator{contract: _Oldglobalexitrootmanagerl2sovereignchain.contract, event: "SetGlobalExitRootUpdater", logs: logs, sub: sub}, nil
}

// WatchSetGlobalExitRootUpdater is a free log subscription operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) WatchSetGlobalExitRootUpdater(opts *bind.WatchOpts, sink chan<- *Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater, newGlobalExitRootUpdater []common.Address) (event.Subscription, error) {

	var newGlobalExitRootUpdaterRule []interface{}
	for _, newGlobalExitRootUpdaterItem := range newGlobalExitRootUpdater {
		newGlobalExitRootUpdaterRule = append(newGlobalExitRootUpdaterRule, newGlobalExitRootUpdaterItem)
	}

	logs, sub, err := _Oldglobalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "SetGlobalExitRootUpdater", newGlobalExitRootUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
				if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootUpdater", log); err != nil {
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

// ParseSetGlobalExitRootUpdater is a log parse operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Oldglobalexitrootmanagerl2sovereignchain *Oldglobalexitrootmanagerl2sovereignchainFilterer) ParseSetGlobalExitRootUpdater(log types.Log) (*Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater, error) {
	event := new(Oldglobalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
	if err := _Oldglobalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
