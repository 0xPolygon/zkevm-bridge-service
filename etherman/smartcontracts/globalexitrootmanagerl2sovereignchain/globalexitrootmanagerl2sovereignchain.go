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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingGlobalExitRootUpdater\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGlobalExitRootRemover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGlobalExitRootRemover\",\"type\":\"address\"}],\"name\":\"AcceptGlobalExitRootRemover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldGlobalExitRootUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"AcceptGlobalExitRootUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentGlobalExitRootRemover\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pendingGlobalExitRootRemover\",\"type\":\"address\"}],\"name\":\"TransferGlobalExitRootRemover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentGlobalExitRootUpdater\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pendingGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"TransferGlobalExitRootUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newHashChainValue\",\"type\":\"bytes32\"}],\"name\":\"UpdateHashChainValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"removedGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRemovalHashChainValue\",\"type\":\"bytes32\"}],\"name\":\"UpdateRemovalHashChainValue\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GER_SOVEREIGN_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGlobalExitRootRemover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootRemover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalExitRootRemover\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"insertGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insertedGERHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGlobalExitRootRemover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGlobalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"gersToRemove\",\"type\":\"bytes32[]\"}],\"name\":\"removeGlobalExitRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removedGERHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGlobalExitRootRemover\",\"type\":\"address\"}],\"name\":\"transferGlobalExitRootRemover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"transferGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161109738038061109783398101604081905261002e91610109565b6001600160a01b038116608052610043610049565b50610136565b603454610100900460ff16156100b55760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60345460ff9081161015610107576034805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610119575f5ffd5b81516001600160a01b038116811461012f575f5ffd5b9392505050565b608051610f426101555f395f818161033301526106650152610f425ff3fe608060405234801561000f575f5ffd5b5060043610610149575f3560e01c806354fd4d50116100c757806391eb796d1161007d578063c053902a11610063578063c053902a14610355578063f5c2f0921461035d578063f5d2f04b14610365575f5ffd5b806391eb796d1461030e578063a3c573eb1461032e575f5ffd5b806368328bc1116100ad57806368328bc1146102995780636ee160d0146102ac5780637c314ce3146102e8575f5ffd5b806354fd4d501461024457806365f0e34714610286575f5ffd5b8063163bbb461161011c5780632d5ddf2b116101025780632d5ddf2b146101fe57806333d6247d1461021e578063485cc95514610231575f5ffd5b8063163bbb46146101d6578063257b3632146101df575f5ffd5b806301fd90441461014d5780630e1bbf9f1461016957806312da06b21461017e57806314770a9314610191575b5f5ffd5b61015660015481565b6040519081526020015b60405180910390f35b61017c610177366004610d9d565b61036e565b005b61017c61018c366004610dbd565b610501565b603a546101b19073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610160565b61015660375481565b6101566101ed366004610dbd565b5f6020819052908152604090205481565b6039546101b19073ffffffffffffffffffffffffffffffffffffffff1681565b61017c61022c366004610dbd565b61064d565b61017c61023f366004610dd4565b6106c1565b60408051808201909152600681527f76312e302e30000000000000000000000000000000000000000000000000000060208201525b6040516101609190610e05565b61017c610294366004610e6e565b6109a5565b61017c6102a7366004610d9d565b610abe565b6102796040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b6034546101b19062010000900473ffffffffffffffffffffffffffffffffffffffff1681565b6035546101b19073ffffffffffffffffffffffffffffffffffffffff1681565b6101b17f000000000000000000000000000000000000000000000000000000000000000081565b61017c610b90565b61017c610c99565b61015660385481565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff166103cf574133146103ca576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610426565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff163314610426576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610473576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092556034546040805162010000909204909316815260208101919091527f1b87468e424189ebdac99fd548646bdb9a48aa9708cfae9f96e6b2e76aee842591015b60405180910390a150565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff166105625741331461055d576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105b9565b60345462010000900473ffffffffffffffffffffffffffffffffffffffff1633146105b9576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81815260208190526040812054900361061b575f818152602081815260408083204290556037548352908390529020603781905560405182907f65d3bf36615f1f02a134d12dfa9ea6b1d4a52386e825973cd27ddb70895c2319905f90a350565b6040517f1f97a58200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106bc576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600155565b603454610100900460ff16158080156106e15750603454600160ff909116105b806106fb5750303b1580156106fb575060345460ff166001145b61078b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840160405180910390fd5b603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156107e957603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff8316610836576040517ff6b2911f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603480547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166201000073ffffffffffffffffffffffffffffffffffffffff86811682029290921792839055604080515f81529190930490911660208201527f8002020f64e628e4e2ff674f8bb88c2709216fc10e3ccdca75ac257faf494236910160405180910390a1603580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117909155604080515f815260208101929092527fabb4abb224bcd13da954d8616357fc9fcf0ccb7057f6dd0fbb7b10624c924ec5910160405180910390a180156109a057603480547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60355473ffffffffffffffffffffffffffffffffffffffff1633146109f6576040517fa34ddeb100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6038545f5b82811015610ab6575f848483818110610a1657610a16610edf565b9050602002013590505f5f8281526020019081526020015f20545f03610a68576040517ff4a66f9d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f9283526020818152604080852083865291859052808520859055519093849183917faafec9380147d2b2b14fe23b1343cbaa1b07f86c5adb060bd28cdf1af4c6f0d491a3506001016109fb565b506038555050565b60355473ffffffffffffffffffffffffffffffffffffffff163314610b0f576040517fa34ddeb100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8381169182179092556035546040805191909316815260208101919091527fea76dc66cb4397b7afec1a503c40b50f296f45c8a3e32d9f3eee2f7e07e7fba791016104f6565b60395473ffffffffffffffffffffffffffffffffffffffff163314610be1576040517f5f063f0100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603480546039805473ffffffffffffffffffffffffffffffffffffffff808216620100009081027fffffffffffffffffffff0000000000000000000000000000000000000000ffff861617958690557fffffffffffffffffffffffff000000000000000000000000000000000000000090921690925560408051938290048316808552919094049091166020830152917f8002020f64e628e4e2ff674f8bb88c2709216fc10e3ccdca75ac257faf49423691016104f6565b603a5473ffffffffffffffffffffffffffffffffffffffff163314610cea576040517f7ca4d27a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60358054603a805473ffffffffffffffffffffffffffffffffffffffff8082167fffffffffffffffffffffffff0000000000000000000000000000000000000000808616821790965594909116909155604080519190921680825260208201939093527fabb4abb224bcd13da954d8616357fc9fcf0ccb7057f6dd0fbb7b10624c924ec591016104f6565b803573ffffffffffffffffffffffffffffffffffffffff81168114610d98575f5ffd5b919050565b5f60208284031215610dad575f5ffd5b610db682610d75565b9392505050565b5f60208284031215610dcd575f5ffd5b5035919050565b5f5f60408385031215610de5575f5ffd5b610dee83610d75565b9150610dfc60208401610d75565b90509250929050565b602081525f82518060208401525f5b81811015610e315760208186018101516040868401015201610e14565b505f6040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b5f5f60208385031215610e7f575f5ffd5b823567ffffffffffffffff811115610e95575f5ffd5b8301601f81018513610ea5575f5ffd5b803567ffffffffffffffff811115610ebb575f5ffd5b8560208260051b8401011115610ecf575f5ffd5b6020919091019590945092505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220e5ce7ae2026903f6fc8638837b20b85f374846901e705e74566169edd19f530b64736f6c634300081c0033",
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

// GERSOVEREIGNVERSION is a free data retrieval call binding the contract method 0x6ee160d0.
//
// Solidity: function GER_SOVEREIGN_VERSION() view returns(string)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) GERSOVEREIGNVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "GER_SOVEREIGN_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GERSOVEREIGNVERSION is a free data retrieval call binding the contract method 0x6ee160d0.
//
// Solidity: function GER_SOVEREIGN_VERSION() view returns(string)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) GERSOVEREIGNVERSION() (string, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GERSOVEREIGNVERSION(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// GERSOVEREIGNVERSION is a free data retrieval call binding the contract method 0x6ee160d0.
//
// Solidity: function GER_SOVEREIGN_VERSION() view returns(string)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) GERSOVEREIGNVERSION() (string, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GERSOVEREIGNVERSION(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
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

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) GlobalExitRootRemover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "globalExitRootRemover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) GlobalExitRootRemover() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// GlobalExitRootRemover is a free data retrieval call binding the contract method 0x91eb796d.
//
// Solidity: function globalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) GlobalExitRootRemover() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.GlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
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

// InsertedGERHashChain is a free data retrieval call binding the contract method 0x163bbb46.
//
// Solidity: function insertedGERHashChain() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) InsertedGERHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "insertedGERHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InsertedGERHashChain is a free data retrieval call binding the contract method 0x163bbb46.
//
// Solidity: function insertedGERHashChain() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) InsertedGERHashChain() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertedGERHashChain(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// InsertedGERHashChain is a free data retrieval call binding the contract method 0x163bbb46.
//
// Solidity: function insertedGERHashChain() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) InsertedGERHashChain() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.InsertedGERHashChain(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
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

// PendingGlobalExitRootRemover is a free data retrieval call binding the contract method 0x14770a93.
//
// Solidity: function pendingGlobalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) PendingGlobalExitRootRemover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "pendingGlobalExitRootRemover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGlobalExitRootRemover is a free data retrieval call binding the contract method 0x14770a93.
//
// Solidity: function pendingGlobalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) PendingGlobalExitRootRemover() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.PendingGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// PendingGlobalExitRootRemover is a free data retrieval call binding the contract method 0x14770a93.
//
// Solidity: function pendingGlobalExitRootRemover() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) PendingGlobalExitRootRemover() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.PendingGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// PendingGlobalExitRootUpdater is a free data retrieval call binding the contract method 0x2d5ddf2b.
//
// Solidity: function pendingGlobalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) PendingGlobalExitRootUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "pendingGlobalExitRootUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGlobalExitRootUpdater is a free data retrieval call binding the contract method 0x2d5ddf2b.
//
// Solidity: function pendingGlobalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) PendingGlobalExitRootUpdater() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.PendingGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// PendingGlobalExitRootUpdater is a free data retrieval call binding the contract method 0x2d5ddf2b.
//
// Solidity: function pendingGlobalExitRootUpdater() view returns(address)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) PendingGlobalExitRootUpdater() (common.Address, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.PendingGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// RemovedGERHashChain is a free data retrieval call binding the contract method 0xf5d2f04b.
//
// Solidity: function removedGERHashChain() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) RemovedGERHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "removedGERHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RemovedGERHashChain is a free data retrieval call binding the contract method 0xf5d2f04b.
//
// Solidity: function removedGERHashChain() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) RemovedGERHashChain() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemovedGERHashChain(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// RemovedGERHashChain is a free data retrieval call binding the contract method 0xf5d2f04b.
//
// Solidity: function removedGERHashChain() view returns(bytes32)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) RemovedGERHashChain() ([32]byte, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemovedGERHashChain(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Globalexitrootmanagerl2sovereignchain.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) Version() (string, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Version(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainCallerSession) Version() (string, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Version(&_Globalexitrootmanagerl2sovereignchain.CallOpts)
}

// AcceptGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xf5c2f092.
//
// Solidity: function acceptGlobalExitRootRemover() returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) AcceptGlobalExitRootRemover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "acceptGlobalExitRootRemover")
}

// AcceptGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xf5c2f092.
//
// Solidity: function acceptGlobalExitRootRemover() returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) AcceptGlobalExitRootRemover() (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.AcceptGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.TransactOpts)
}

// AcceptGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xf5c2f092.
//
// Solidity: function acceptGlobalExitRootRemover() returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) AcceptGlobalExitRootRemover() (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.AcceptGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.TransactOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) AcceptGlobalExitRootUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "acceptGlobalExitRootUpdater")
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.AcceptGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.TransactOpts)
}

// AcceptGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0xc053902a.
//
// Solidity: function acceptGlobalExitRootUpdater() returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) AcceptGlobalExitRootUpdater() (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.AcceptGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) Initialize(opts *bind.TransactOpts, _globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "initialize", _globalExitRootUpdater, _globalExitRootRemover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) Initialize(_globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater, _globalExitRootRemover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _globalExitRootUpdater, address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) Initialize(_globalExitRootUpdater common.Address, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.Initialize(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater, _globalExitRootRemover)
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

// RemoveGlobalExitRoots is a paid mutator transaction binding the contract method 0x65f0e347.
//
// Solidity: function removeGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) RemoveGlobalExitRoots(opts *bind.TransactOpts, gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "removeGlobalExitRoots", gersToRemove)
}

// RemoveGlobalExitRoots is a paid mutator transaction binding the contract method 0x65f0e347.
//
// Solidity: function removeGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) RemoveGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemoveGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
}

// RemoveGlobalExitRoots is a paid mutator transaction binding the contract method 0x65f0e347.
//
// Solidity: function removeGlobalExitRoots(bytes32[] gersToRemove) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) RemoveGlobalExitRoots(gersToRemove [][32]byte) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.RemoveGlobalExitRoots(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, gersToRemove)
}

// TransferGlobalExitRootRemover is a paid mutator transaction binding the contract method 0x68328bc1.
//
// Solidity: function transferGlobalExitRootRemover(address _newGlobalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) TransferGlobalExitRootRemover(opts *bind.TransactOpts, _newGlobalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "transferGlobalExitRootRemover", _newGlobalExitRootRemover)
}

// TransferGlobalExitRootRemover is a paid mutator transaction binding the contract method 0x68328bc1.
//
// Solidity: function transferGlobalExitRootRemover(address _newGlobalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) TransferGlobalExitRootRemover(_newGlobalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.TransferGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _newGlobalExitRootRemover)
}

// TransferGlobalExitRootRemover is a paid mutator transaction binding the contract method 0x68328bc1.
//
// Solidity: function transferGlobalExitRootRemover(address _newGlobalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) TransferGlobalExitRootRemover(_newGlobalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.TransferGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _newGlobalExitRootRemover)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) TransferGlobalExitRootUpdater(opts *bind.TransactOpts, _newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "transferGlobalExitRootUpdater", _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.TransferGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _newGlobalExitRootUpdater)
}

// TransferGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x0e1bbf9f.
//
// Solidity: function transferGlobalExitRootUpdater(address _newGlobalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) TransferGlobalExitRootUpdater(_newGlobalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.TransferGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _newGlobalExitRootUpdater)
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

// Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemoverIterator is returned from FilterAcceptGlobalExitRootRemover and is used to iterate over the raw logs and unpacked data for AcceptGlobalExitRootRemover events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemoverIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemover // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemover)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemover)
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
func (it *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemover represents a AcceptGlobalExitRootRemover event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemover struct {
	OldGlobalExitRootRemover common.Address
	NewGlobalExitRootRemover common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptGlobalExitRootRemover is a free log retrieval operation binding the contract event 0xabb4abb224bcd13da954d8616357fc9fcf0ccb7057f6dd0fbb7b10624c924ec5.
//
// Solidity: event AcceptGlobalExitRootRemover(address oldGlobalExitRootRemover, address newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterAcceptGlobalExitRootRemover(opts *bind.FilterOpts) (*Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemoverIterator, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "AcceptGlobalExitRootRemover")
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemoverIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "AcceptGlobalExitRootRemover", logs: logs, sub: sub}, nil
}

// WatchAcceptGlobalExitRootRemover is a free log subscription operation binding the contract event 0xabb4abb224bcd13da954d8616357fc9fcf0ccb7057f6dd0fbb7b10624c924ec5.
//
// Solidity: event AcceptGlobalExitRootRemover(address oldGlobalExitRootRemover, address newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchAcceptGlobalExitRootRemover(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemover) (event.Subscription, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "AcceptGlobalExitRootRemover")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemover)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "AcceptGlobalExitRootRemover", log); err != nil {
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

// ParseAcceptGlobalExitRootRemover is a log parse operation binding the contract event 0xabb4abb224bcd13da954d8616357fc9fcf0ccb7057f6dd0fbb7b10624c924ec5.
//
// Solidity: event AcceptGlobalExitRootRemover(address oldGlobalExitRootRemover, address newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseAcceptGlobalExitRootRemover(log types.Log) (*Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemover, error) {
	event := new(Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootRemover)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "AcceptGlobalExitRootRemover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdaterIterator is returned from FilterAcceptGlobalExitRootUpdater and is used to iterate over the raw logs and unpacked data for AcceptGlobalExitRootUpdater events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdaterIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdater // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdater)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdater)
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
func (it *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdater represents a AcceptGlobalExitRootUpdater event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdater struct {
	OldGlobalExitRootUpdater common.Address
	NewGlobalExitRootUpdater common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptGlobalExitRootUpdater is a free log retrieval operation binding the contract event 0x8002020f64e628e4e2ff674f8bb88c2709216fc10e3ccdca75ac257faf494236.
//
// Solidity: event AcceptGlobalExitRootUpdater(address oldGlobalExitRootUpdater, address newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterAcceptGlobalExitRootUpdater(opts *bind.FilterOpts) (*Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdaterIterator, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "AcceptGlobalExitRootUpdater")
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdaterIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "AcceptGlobalExitRootUpdater", logs: logs, sub: sub}, nil
}

// WatchAcceptGlobalExitRootUpdater is a free log subscription operation binding the contract event 0x8002020f64e628e4e2ff674f8bb88c2709216fc10e3ccdca75ac257faf494236.
//
// Solidity: event AcceptGlobalExitRootUpdater(address oldGlobalExitRootUpdater, address newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchAcceptGlobalExitRootUpdater(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdater) (event.Subscription, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "AcceptGlobalExitRootUpdater")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdater)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "AcceptGlobalExitRootUpdater", log); err != nil {
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

// ParseAcceptGlobalExitRootUpdater is a log parse operation binding the contract event 0x8002020f64e628e4e2ff674f8bb88c2709216fc10e3ccdca75ac257faf494236.
//
// Solidity: event AcceptGlobalExitRootUpdater(address oldGlobalExitRootUpdater, address newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseAcceptGlobalExitRootUpdater(log types.Log) (*Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdater, error) {
	event := new(Globalexitrootmanagerl2sovereignchainAcceptGlobalExitRootUpdater)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "AcceptGlobalExitRootUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemoverIterator is returned from FilterTransferGlobalExitRootRemover and is used to iterate over the raw logs and unpacked data for TransferGlobalExitRootRemover events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemoverIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemover // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemover)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemover)
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
func (it *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemover represents a TransferGlobalExitRootRemover event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemover struct {
	CurrentGlobalExitRootRemover common.Address
	PendingGlobalExitRootRemover common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterTransferGlobalExitRootRemover is a free log retrieval operation binding the contract event 0xea76dc66cb4397b7afec1a503c40b50f296f45c8a3e32d9f3eee2f7e07e7fba7.
//
// Solidity: event TransferGlobalExitRootRemover(address currentGlobalExitRootRemover, address pendingGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterTransferGlobalExitRootRemover(opts *bind.FilterOpts) (*Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemoverIterator, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "TransferGlobalExitRootRemover")
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemoverIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "TransferGlobalExitRootRemover", logs: logs, sub: sub}, nil
}

// WatchTransferGlobalExitRootRemover is a free log subscription operation binding the contract event 0xea76dc66cb4397b7afec1a503c40b50f296f45c8a3e32d9f3eee2f7e07e7fba7.
//
// Solidity: event TransferGlobalExitRootRemover(address currentGlobalExitRootRemover, address pendingGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchTransferGlobalExitRootRemover(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemover) (event.Subscription, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "TransferGlobalExitRootRemover")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemover)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "TransferGlobalExitRootRemover", log); err != nil {
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

// ParseTransferGlobalExitRootRemover is a log parse operation binding the contract event 0xea76dc66cb4397b7afec1a503c40b50f296f45c8a3e32d9f3eee2f7e07e7fba7.
//
// Solidity: event TransferGlobalExitRootRemover(address currentGlobalExitRootRemover, address pendingGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseTransferGlobalExitRootRemover(log types.Log) (*Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemover, error) {
	event := new(Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootRemover)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "TransferGlobalExitRootRemover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdaterIterator is returned from FilterTransferGlobalExitRootUpdater and is used to iterate over the raw logs and unpacked data for TransferGlobalExitRootUpdater events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdaterIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdater // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdater)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdater)
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
func (it *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdater represents a TransferGlobalExitRootUpdater event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdater struct {
	CurrentGlobalExitRootUpdater common.Address
	PendingGlobalExitRootUpdater common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterTransferGlobalExitRootUpdater is a free log retrieval operation binding the contract event 0x1b87468e424189ebdac99fd548646bdb9a48aa9708cfae9f96e6b2e76aee8425.
//
// Solidity: event TransferGlobalExitRootUpdater(address currentGlobalExitRootUpdater, address pendingGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterTransferGlobalExitRootUpdater(opts *bind.FilterOpts) (*Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdaterIterator, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "TransferGlobalExitRootUpdater")
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdaterIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "TransferGlobalExitRootUpdater", logs: logs, sub: sub}, nil
}

// WatchTransferGlobalExitRootUpdater is a free log subscription operation binding the contract event 0x1b87468e424189ebdac99fd548646bdb9a48aa9708cfae9f96e6b2e76aee8425.
//
// Solidity: event TransferGlobalExitRootUpdater(address currentGlobalExitRootUpdater, address pendingGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchTransferGlobalExitRootUpdater(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdater) (event.Subscription, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "TransferGlobalExitRootUpdater")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdater)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "TransferGlobalExitRootUpdater", log); err != nil {
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

// ParseTransferGlobalExitRootUpdater is a log parse operation binding the contract event 0x1b87468e424189ebdac99fd548646bdb9a48aa9708cfae9f96e6b2e76aee8425.
//
// Solidity: event TransferGlobalExitRootUpdater(address currentGlobalExitRootUpdater, address pendingGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseTransferGlobalExitRootUpdater(log types.Log) (*Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdater, error) {
	event := new(Globalexitrootmanagerl2sovereignchainTransferGlobalExitRootUpdater)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "TransferGlobalExitRootUpdater", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainUpdateHashChainValueIterator is returned from FilterUpdateHashChainValue and is used to iterate over the raw logs and unpacked data for UpdateHashChainValue events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainUpdateHashChainValueIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainUpdateHashChainValue // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainUpdateHashChainValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainUpdateHashChainValue)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainUpdateHashChainValue)
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
func (it *Globalexitrootmanagerl2sovereignchainUpdateHashChainValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainUpdateHashChainValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainUpdateHashChainValue represents a UpdateHashChainValue event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainUpdateHashChainValue struct {
	NewGlobalExitRoot [32]byte
	NewHashChainValue [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateHashChainValue is a free log retrieval operation binding the contract event 0x65d3bf36615f1f02a134d12dfa9ea6b1d4a52386e825973cd27ddb70895c2319.
//
// Solidity: event UpdateHashChainValue(bytes32 indexed newGlobalExitRoot, bytes32 indexed newHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterUpdateHashChainValue(opts *bind.FilterOpts, newGlobalExitRoot [][32]byte, newHashChainValue [][32]byte) (*Globalexitrootmanagerl2sovereignchainUpdateHashChainValueIterator, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}
	var newHashChainValueRule []interface{}
	for _, newHashChainValueItem := range newHashChainValue {
		newHashChainValueRule = append(newHashChainValueRule, newHashChainValueItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "UpdateHashChainValue", newGlobalExitRootRule, newHashChainValueRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainUpdateHashChainValueIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "UpdateHashChainValue", logs: logs, sub: sub}, nil
}

// WatchUpdateHashChainValue is a free log subscription operation binding the contract event 0x65d3bf36615f1f02a134d12dfa9ea6b1d4a52386e825973cd27ddb70895c2319.
//
// Solidity: event UpdateHashChainValue(bytes32 indexed newGlobalExitRoot, bytes32 indexed newHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchUpdateHashChainValue(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainUpdateHashChainValue, newGlobalExitRoot [][32]byte, newHashChainValue [][32]byte) (event.Subscription, error) {

	var newGlobalExitRootRule []interface{}
	for _, newGlobalExitRootItem := range newGlobalExitRoot {
		newGlobalExitRootRule = append(newGlobalExitRootRule, newGlobalExitRootItem)
	}
	var newHashChainValueRule []interface{}
	for _, newHashChainValueItem := range newHashChainValue {
		newHashChainValueRule = append(newHashChainValueRule, newHashChainValueItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "UpdateHashChainValue", newGlobalExitRootRule, newHashChainValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainUpdateHashChainValue)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "UpdateHashChainValue", log); err != nil {
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

// ParseUpdateHashChainValue is a log parse operation binding the contract event 0x65d3bf36615f1f02a134d12dfa9ea6b1d4a52386e825973cd27ddb70895c2319.
//
// Solidity: event UpdateHashChainValue(bytes32 indexed newGlobalExitRoot, bytes32 indexed newHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseUpdateHashChainValue(log types.Log) (*Globalexitrootmanagerl2sovereignchainUpdateHashChainValue, error) {
	event := new(Globalexitrootmanagerl2sovereignchainUpdateHashChainValue)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "UpdateHashChainValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValueIterator is returned from FilterUpdateRemovalHashChainValue and is used to iterate over the raw logs and unpacked data for UpdateRemovalHashChainValue events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValueIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValue // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValue)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValue)
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
func (it *Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValue represents a UpdateRemovalHashChainValue event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValue struct {
	RemovedGlobalExitRoot    [32]byte
	NewRemovalHashChainValue [32]byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterUpdateRemovalHashChainValue is a free log retrieval operation binding the contract event 0xaafec9380147d2b2b14fe23b1343cbaa1b07f86c5adb060bd28cdf1af4c6f0d4.
//
// Solidity: event UpdateRemovalHashChainValue(bytes32 indexed removedGlobalExitRoot, bytes32 indexed newRemovalHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterUpdateRemovalHashChainValue(opts *bind.FilterOpts, removedGlobalExitRoot [][32]byte, newRemovalHashChainValue [][32]byte) (*Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValueIterator, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}
	var newRemovalHashChainValueRule []interface{}
	for _, newRemovalHashChainValueItem := range newRemovalHashChainValue {
		newRemovalHashChainValueRule = append(newRemovalHashChainValueRule, newRemovalHashChainValueItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "UpdateRemovalHashChainValue", removedGlobalExitRootRule, newRemovalHashChainValueRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValueIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "UpdateRemovalHashChainValue", logs: logs, sub: sub}, nil
}

// WatchUpdateRemovalHashChainValue is a free log subscription operation binding the contract event 0xaafec9380147d2b2b14fe23b1343cbaa1b07f86c5adb060bd28cdf1af4c6f0d4.
//
// Solidity: event UpdateRemovalHashChainValue(bytes32 indexed removedGlobalExitRoot, bytes32 indexed newRemovalHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchUpdateRemovalHashChainValue(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValue, removedGlobalExitRoot [][32]byte, newRemovalHashChainValue [][32]byte) (event.Subscription, error) {

	var removedGlobalExitRootRule []interface{}
	for _, removedGlobalExitRootItem := range removedGlobalExitRoot {
		removedGlobalExitRootRule = append(removedGlobalExitRootRule, removedGlobalExitRootItem)
	}
	var newRemovalHashChainValueRule []interface{}
	for _, newRemovalHashChainValueItem := range newRemovalHashChainValue {
		newRemovalHashChainValueRule = append(newRemovalHashChainValueRule, newRemovalHashChainValueItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "UpdateRemovalHashChainValue", removedGlobalExitRootRule, newRemovalHashChainValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValue)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "UpdateRemovalHashChainValue", log); err != nil {
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

// ParseUpdateRemovalHashChainValue is a log parse operation binding the contract event 0xaafec9380147d2b2b14fe23b1343cbaa1b07f86c5adb060bd28cdf1af4c6f0d4.
//
// Solidity: event UpdateRemovalHashChainValue(bytes32 indexed removedGlobalExitRoot, bytes32 indexed newRemovalHashChainValue)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseUpdateRemovalHashChainValue(log types.Log) (*Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValue, error) {
	event := new(Globalexitrootmanagerl2sovereignchainUpdateRemovalHashChainValue)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "UpdateRemovalHashChainValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
