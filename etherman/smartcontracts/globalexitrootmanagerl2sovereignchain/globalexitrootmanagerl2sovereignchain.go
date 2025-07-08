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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GlobalExitRootAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAllowedContracts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootUpdater\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGlobalExitRootRemover\",\"type\":\"address\"}],\"name\":\"SetGlobalExitRootRemover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGlobalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"SetGlobalExitRootUpdater\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newHashChainValue\",\"type\":\"bytes32\"}],\"name\":\"UpdateHashChainValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"removedGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newRemovalHashChainValue\",\"type\":\"bytes32\"}],\"name\":\"UpdateRemovalHashChainValue\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"globalExitRootMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootRemover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalExitRootRemover\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newRoot\",\"type\":\"bytes32\"}],\"name\":\"insertGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insertedGERHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRollupExitRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"gersToRemove\",\"type\":\"bytes32[]\"}],\"name\":\"removeGlobalExitRoots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removedGERHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootRemover\",\"type\":\"address\"}],\"name\":\"setGlobalExitRootRemover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_globalExitRootUpdater\",\"type\":\"address\"}],\"name\":\"setGlobalExitRootUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"updateExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161118b38038061118b833981810160405281019061003191906101fe565b808073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505061007561007b60201b60201c565b50610264565b5f61008a61017960201b60201c565b9050805f0160089054906101000a900460ff16156100d4576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b67ffffffffffffffff8016815f015f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16146101765767ffffffffffffffff815f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d267ffffffffffffffff60405161016d919061024b565b60405180910390a15b50565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101cd826101a4565b9050919050565b6101dd816101c3565b81146101e7575f5ffd5b50565b5f815190506101f8816101d4565b92915050565b5f60208284031215610213576102126101a0565b5b5f610220848285016101ea565b91505092915050565b5f67ffffffffffffffff82169050919050565b61024581610229565b82525050565b5f60208201905061025e5f83018461023c565b92915050565b608051610f086102835f395f818161046c0152610a730152610f085ff3fe608060405234801561000f575f5ffd5b50600436106100cd575f3560e01c806365f0e3471161008a57806391eb796d1161006457806391eb796d146101e7578063a3c573eb14610205578063d0267f3914610223578063f5d2f04b1461023f576100cd565b806365f0e347146101915780636da0e4ab146101ad5780637c314ce3146101c9576100cd565b806301fd9044146100d157806312da06b2146100ef578063163bbb461461010b578063257b36321461012957806333d6247d14610159578063485cc95514610175575b5f5ffd5b6100d961025d565b6040516100e69190610bf9565b60405180910390f35b61010960048036038101906101049190610c44565b610263565b005b610113610450565b6040516101209190610bf9565b60405180910390f35b610143600480360381019061013e9190610c44565b610456565b6040516101509190610c87565b60405180910390f35b610173600480360381019061016e9190610c44565b61046a565b005b61018f600480360381019061018a9190610cfa565b6104f9565b005b6101ab60048036038101906101a69190610d99565b6106f2565b005b6101c760048036038101906101c29190610de4565b61085c565b005b6101d1610a27565b6040516101de9190610e1e565b60405180910390f35b6101ef610a4c565b6040516101fc9190610e1e565b60405180910390f35b61020d610a71565b60405161021a9190610e1e565b60405180910390f35b61023d60048036038101906102389190610de4565b610a95565b005b610247610ba0565b6040516102549190610bf9565b60405180910390f35b60015481565b5f73ffffffffffffffffffffffffffffffffffffffff1660345f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610321573373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff161461031c576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103a8565b3373ffffffffffffffffffffffffffffffffffffffff1660345f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146103a7576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5f5f5f8381526020019081526020015f20540361041b57425f5f8381526020019081526020015f20819055506103e060375482610ba6565b603781905550603754817f65d3bf36615f1f02a134d12dfa9ea6b1d4a52386e825973cd27ddb70895c231960405160405180910390a361044d565b6040517f1f97a58200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b60375481565b5f602052805f5260405f205f915090505481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ef576040517fb49365dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060018190555050565b5f610502610bba565b90505f815f0160089054906101000a900460ff161590505f825f015f9054906101000a900467ffffffffffffffff1690505f5f8267ffffffffffffffff1614801561054a5750825b90505f60018367ffffffffffffffff1614801561057d57505f3073ffffffffffffffffffffffffffffffffffffffff163b145b90508115801561058b575080155b156105c2576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001855f015f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550831561060f576001855f0160086101000a81548160ff0219169083151502179055505b8660345f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560355f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083156106e9575f855f0160086101000a81548160ff0219169083151502179055507fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d260016040516106e09190610e8c565b60405180910390a15b50505050505050565b3373ffffffffffffffffffffffffffffffffffffffff1660355f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610778576040517fa34ddeb100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60385490505f5f90505b8383905081101561084f575f8484838181106107a2576107a1610ea5565b5b9050602002013590505f5f5f8381526020019081526020015f2054036107f4576040517ff4a66f9d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107fe8382610ba6565b92505f5f8281526020019081526020015f205f905582817faafec9380147d2b2b14fe23b1343cbaa1b07f86c5adb060bd28cdf1af4c6f0d460405160405180910390a3508080600101915050610783565b5080603881905550505050565b5f73ffffffffffffffffffffffffffffffffffffffff1660345f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361091a573373ffffffffffffffffffffffffffffffffffffffff164173ffffffffffffffffffffffffffffffffffffffff1614610915576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109a1565b3373ffffffffffffffffffffffffffffffffffffffff1660345f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146109a0576040517fc758fc1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b8060345f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167f992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c60405160405180910390a250565b60345f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60355f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b3373ffffffffffffffffffffffffffffffffffffffff1660355f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610b1b576040517fa34ddeb100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060355f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff167eb4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b75860405160405180910390a250565b60385481565b5f825f528160205260405f20905092915050565b5f7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00905090565b5f819050919050565b610bf381610be1565b82525050565b5f602082019050610c0c5f830184610bea565b92915050565b5f5ffd5b5f5ffd5b610c2381610be1565b8114610c2d575f5ffd5b50565b5f81359050610c3e81610c1a565b92915050565b5f60208284031215610c5957610c58610c12565b5b5f610c6684828501610c30565b91505092915050565b5f819050919050565b610c8181610c6f565b82525050565b5f602082019050610c9a5f830184610c78565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610cc982610ca0565b9050919050565b610cd981610cbf565b8114610ce3575f5ffd5b50565b5f81359050610cf481610cd0565b92915050565b5f5f60408385031215610d1057610d0f610c12565b5b5f610d1d85828601610ce6565b9250506020610d2e85828601610ce6565b9150509250929050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610d5957610d58610d38565b5b8235905067ffffffffffffffff811115610d7657610d75610d3c565b5b602083019150836020820283011115610d9257610d91610d40565b5b9250929050565b5f5f60208385031215610daf57610dae610c12565b5b5f83013567ffffffffffffffff811115610dcc57610dcb610c16565b5b610dd885828601610d44565b92509250509250929050565b5f60208284031215610df957610df8610c12565b5b5f610e0684828501610ce6565b91505092915050565b610e1881610cbf565b82525050565b5f602082019050610e315f830184610e0f565b92915050565b5f819050919050565b5f67ffffffffffffffff82169050919050565b5f819050919050565b5f610e76610e71610e6c84610e37565b610e53565b610e40565b9050919050565b610e8681610e5c565b82525050565b5f602082019050610e9f5f830184610e7d565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea26469706673582212200c9c0d8affb171d68f68f8f69c6c9ad5816f657084db21706ca7d4b5529f4e5664736f6c634300081c0033",
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

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) SetGlobalExitRootRemover(opts *bind.TransactOpts, _globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "setGlobalExitRootRemover", _globalExitRootRemover)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) SetGlobalExitRootRemover(_globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootRemover)
}

// SetGlobalExitRootRemover is a paid mutator transaction binding the contract method 0xd0267f39.
//
// Solidity: function setGlobalExitRootRemover(address _globalExitRootRemover) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) SetGlobalExitRootRemover(_globalExitRootRemover common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootRemover(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootRemover)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactor) SetGlobalExitRootUpdater(opts *bind.TransactOpts, _globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.contract.Transact(opts, "setGlobalExitRootUpdater", _globalExitRootUpdater)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainSession) SetGlobalExitRootUpdater(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
}

// SetGlobalExitRootUpdater is a paid mutator transaction binding the contract method 0x6da0e4ab.
//
// Solidity: function setGlobalExitRootUpdater(address _globalExitRootUpdater) returns()
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainTransactorSession) SetGlobalExitRootUpdater(_globalExitRootUpdater common.Address) (*types.Transaction, error) {
	return _Globalexitrootmanagerl2sovereignchain.Contract.SetGlobalExitRootUpdater(&_Globalexitrootmanagerl2sovereignchain.TransactOpts, _globalExitRootUpdater)
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
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterInitialized(opts *bind.FilterOpts) (*Globalexitrootmanagerl2sovereignchainInitializedIterator, error) {

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainInitializedIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseInitialized(log types.Log) (*Globalexitrootmanagerl2sovereignchainInitialized, error) {
	event := new(Globalexitrootmanagerl2sovereignchainInitialized)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator is returned from FilterSetGlobalExitRootRemover and is used to iterate over the raw logs and unpacked data for SetGlobalExitRootRemover events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
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
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover represents a SetGlobalExitRootRemover event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover struct {
	NewGlobalExitRootRemover common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalExitRootRemover is a free log retrieval operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterSetGlobalExitRootRemover(opts *bind.FilterOpts, newGlobalExitRootRemover []common.Address) (*Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator, error) {

	var newGlobalExitRootRemoverRule []interface{}
	for _, newGlobalExitRootRemoverItem := range newGlobalExitRootRemover {
		newGlobalExitRootRemoverRule = append(newGlobalExitRootRemoverRule, newGlobalExitRootRemoverItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "SetGlobalExitRootRemover", newGlobalExitRootRemoverRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemoverIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "SetGlobalExitRootRemover", logs: logs, sub: sub}, nil
}

// WatchSetGlobalExitRootRemover is a free log subscription operation binding the contract event 0x00b4672b6135d1dfbd4e9520e01abb14ea5eac645990b0d24dfda00ae999b758.
//
// Solidity: event SetGlobalExitRootRemover(address indexed newGlobalExitRootRemover)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchSetGlobalExitRootRemover(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover, newGlobalExitRootRemover []common.Address) (event.Subscription, error) {

	var newGlobalExitRootRemoverRule []interface{}
	for _, newGlobalExitRootRemoverItem := range newGlobalExitRootRemover {
		newGlobalExitRootRemoverRule = append(newGlobalExitRootRemoverRule, newGlobalExitRootRemoverItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "SetGlobalExitRootRemover", newGlobalExitRootRemoverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootRemover", log); err != nil {
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
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseSetGlobalExitRootRemover(log types.Log) (*Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover, error) {
	event := new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootRemover)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootRemover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator is returned from FilterSetGlobalExitRootUpdater and is used to iterate over the raw logs and unpacked data for SetGlobalExitRootUpdater events raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator struct {
	Event *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater // Event containing the contract specifics and raw log

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
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
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
		it.Event = new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
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
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater represents a SetGlobalExitRootUpdater event raised by the Globalexitrootmanagerl2sovereignchain contract.
type Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater struct {
	NewGlobalExitRootUpdater common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSetGlobalExitRootUpdater is a free log retrieval operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) FilterSetGlobalExitRootUpdater(opts *bind.FilterOpts, newGlobalExitRootUpdater []common.Address) (*Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator, error) {

	var newGlobalExitRootUpdaterRule []interface{}
	for _, newGlobalExitRootUpdaterItem := range newGlobalExitRootUpdater {
		newGlobalExitRootUpdaterRule = append(newGlobalExitRootUpdaterRule, newGlobalExitRootUpdaterItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.FilterLogs(opts, "SetGlobalExitRootUpdater", newGlobalExitRootUpdaterRule)
	if err != nil {
		return nil, err
	}
	return &Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdaterIterator{contract: _Globalexitrootmanagerl2sovereignchain.contract, event: "SetGlobalExitRootUpdater", logs: logs, sub: sub}, nil
}

// WatchSetGlobalExitRootUpdater is a free log subscription operation binding the contract event 0x992b80814dbc3fba903486d81daddb07d1d5b20483742458c8b0540e3a37e37c.
//
// Solidity: event SetGlobalExitRootUpdater(address indexed newGlobalExitRootUpdater)
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) WatchSetGlobalExitRootUpdater(opts *bind.WatchOpts, sink chan<- *Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater, newGlobalExitRootUpdater []common.Address) (event.Subscription, error) {

	var newGlobalExitRootUpdaterRule []interface{}
	for _, newGlobalExitRootUpdaterItem := range newGlobalExitRootUpdater {
		newGlobalExitRootUpdaterRule = append(newGlobalExitRootUpdaterRule, newGlobalExitRootUpdaterItem)
	}

	logs, sub, err := _Globalexitrootmanagerl2sovereignchain.contract.WatchLogs(opts, "SetGlobalExitRootUpdater", newGlobalExitRootUpdaterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
				if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootUpdater", log); err != nil {
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
func (_Globalexitrootmanagerl2sovereignchain *Globalexitrootmanagerl2sovereignchainFilterer) ParseSetGlobalExitRootUpdater(log types.Log) (*Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater, error) {
	event := new(Globalexitrootmanagerl2sovereignchainSetGlobalExitRootUpdater)
	if err := _Globalexitrootmanagerl2sovereignchain.contract.UnpackLog(event, "SetGlobalExitRootUpdater", log); err != nil {
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
