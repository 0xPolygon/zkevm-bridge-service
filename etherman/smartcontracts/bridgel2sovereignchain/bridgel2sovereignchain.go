// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgel2sovereignchain

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

// Bridgel2sovereignchainMetaData contains all meta data concerning the Bridgel2sovereignchain contract.
var Bridgel2sovereignchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTokenWrappedDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MigrateLegacyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"RemoveLegacySovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"}],\"name\":\"SetBridgeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignWETHAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"UnsetClaim\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_INIT_BYTECODE_WRAPPED_TOKEN\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractTokenWrapped\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"calculateTokenWrapperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migrateLegacyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"precalculatedWrapperAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacySovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"removeLegacySovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"}],\"name\":\"setBridgeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetworks\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sovereignTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNotMintable\",\"type\":\"bool[]\"}],\"name\":\"setMultipleSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"setSovereignWETHAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"leafIndexes\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"sourceBridgeNetworks\",\"type\":\"uint32[]\"}],\"name\":\"unsetMultipleClaimedBitmap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedAddress\",\"type\":\"address\"}],\"name\":\"wrappedAddressIsNotMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506200001b6200002b565b620000256200002b565b620000e9565b5f54610100900460ff1615620000975760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614620000e7575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b615ffb80620000f75f395ff3fe60806040526004361061023e575f3560e01c80638c0dd47011610134578063c0f49163116100b3578063dbc1697611610078578063dbc16976146102a7578063eabd372a146106fc578063ee25560b1461071b578063f5efcd7914610746578063f811bff714610765578063fb57083414610784575f80fd5b8063c0f4916314610657578063cc46163214610685578063ccaa2d11146106a4578063cd586579146106c3578063d02103ca146106d6575f80fd5b8063b8b284d0116100f9578063b8b284d0146105b6578063bab161bf146105d5578063be5831c7146105f6578063bf130d7f14610619578063c00f14ab14610638575f80fd5b80638c0dd4701461051b5780638ed7e3f21461053a5780639e76158f14610559578063aaa13cc214610578578063b458696214610597575f80fd5b80633e197043116101c057806379e2cf971161018557806379e2cf971461048157806381b1c1741461049557806383c43a55146104c957806383f24403146104dd5780638781a5c5146104fc575f80fd5b80633e197043146103f15780634b2f336d1461041057806357cfbee31461042f5780635ca1e1651461044e5780637843298b14610462575f80fd5b806327aef4e81161020657806327aef4e8146102ef5780632dfdf0b514610310578063318aee3d146103335780633c351e101461039a5780633cbc795b146103b9575f80fd5b806314cc01a01461024257806315064c961461027e5780632072f6c5146102a757806322e95f2c146102bd578063240ff378146102dc575b5f80fd5b34801561024d575f80fd5b5060a354610261906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b348015610289575f80fd5b506068546102979060ff1681565b6040519015158152602001610275565b3480156102b2575f80fd5b506102bb6107a3565b005b3480156102c8575f80fd5b506102616102d736600461328d565b6107bc565b6102bb6102ea366004613313565b61080a565b3480156102fa575f80fd5b5061030361087a565b60405161027591906133d4565b34801561031b575f80fd5b5061032560535481565b604051908152602001610275565b34801561033e575f80fd5b5061037661034d3660046133ed565b606b6020525f908152604090205463ffffffff811690600160201b90046001600160a01b031682565b6040805163ffffffff90931683526001600160a01b03909116602083015201610275565b3480156103a5575f80fd5b50606d54610261906001600160a01b031681565b3480156103c4575f80fd5b50606d546103dc90600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610275565b3480156103fc575f80fd5b5061032561040b366004613416565b610906565b34801561041b575f80fd5b50606f54610261906001600160a01b031681565b34801561043a575f80fd5b506102bb6104493660046135c6565b610992565b348015610459575f80fd5b50610325610a88565b34801561046d575f80fd5b5061026161047c3660046136c5565b610b64565b34801561048c575f80fd5b506102bb610b8d565b3480156104a0575f80fd5b506102616104af36600461370b565b606a6020525f90815260409020546001600160a01b031681565b3480156104d4575f80fd5b50610303610bb0565b3480156104e8575f80fd5b506103256104f7366004613733565b610bcf565b348015610507575f80fd5b506102bb61051636600461376f565b610ca4565b348015610526575f80fd5b506102bb610535366004613841565b610d4f565b348015610545575f80fd5b50606c54610261906001600160a01b031681565b348015610564575f80fd5b506102bb61057336600461390a565b61102b565b348015610583575f80fd5b50610261610592366004613934565b611163565b3480156105a2575f80fd5b506102bb6105b13660046133ed565b611263565b3480156105c1575f80fd5b506102bb6105d03660046139c9565b6113bd565b3480156105e0575f80fd5b506068546103dc90610100900463ffffffff1681565b348015610601575f80fd5b506068546103dc90600160c81b900463ffffffff1681565b348015610624575f80fd5b506102bb610633366004613a46565b611436565b348015610643575f80fd5b506103036106523660046133ed565b611500565b348015610662575f80fd5b506102976106713660046133ed565b60a26020525f908152604090205460ff1681565b348015610690575f80fd5b5061029761069f366004613a72565b611545565b3480156106af575f80fd5b506102bb6106be366004613aa3565b611595565b6102bb6106d1366004613b86565b611996565b3480156106e1575f80fd5b5060685461026190600160281b90046001600160a01b031681565b348015610707575f80fd5b506102bb6107163660046133ed565b611d2c565b348015610726575f80fd5b5061032561073536600461370b565b60696020525f908152604090205481565b348015610751575f80fd5b506102bb610760366004613aa3565b611dab565b348015610770575f80fd5b506102bb61077f366004613c15565b61200e565b34801561078f575f80fd5b5061029761079e366004613ca4565b6120ac565b60405163441845b160e01b815260040160405180910390fd5b5f606a5f84846040516020016107d3929190613ce9565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b031690505b92915050565b60685460ff161561082e57604051630bc011ff60e21b815260040160405180910390fd5b34158015906108475750606f546001600160a01b031615155b15610865576040516301bd897160e61b815260040160405180910390fd5b6108738585348686866120c3565b5050505050565b606e805461088790613d13565b80601f01602080910402602001604051908101604052809291908181526020018280546108b390613d13565b80156108fe5780601f106108d5576101008083540402835291602001916108fe565b820191905f5260205f20905b8154815290600101906020018083116108e157829003601f168201915b505050505081565b6040516001600160f81b031960f889901b1660208201526001600160e01b031960e088811b821660218401526001600160601b0319606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b60a3546001600160a01b031633146109bd576040516357b738d160e11b815260040160405180910390fd5b825184511415806109d057508151845114155b806109dd57508051845114155b156109fb5760405163434f49f560e11b815260040160405180910390fd5b5f5b825181101561087357610a76858281518110610a1b57610a1b613d4b565b6020026020010151858381518110610a3557610a35613d4b565b6020026020010151858481518110610a4f57610a4f613d4b565b6020026020010151858581518110610a6957610a69613d4b565b602002602001015161218d565b80610a8081613d73565b9150506109fd565b6053545f90819081805b6020811015610b5b578083901c600116600103610aef5760338160208110610abc57610abc613d4b565b01546040805160208101929092528101859052606001604051602081830303815290604052805190602001209350610b1c565b60408051602081018690529081018390526060016040516020818303038152906040528051906020012093505b60408051602081018490529081018390526060016040516020818303038152906040528051906020012091508080610b5390613d73565b915050610a92565b50919392505050565b5f610b858484610b738561233d565b610b7c866123f8565b610592876124ac565b949350505050565b605354606854600160c81b900463ffffffff161015610bae57610bae612560565b565b60405180611ba00160405280611b668152602001614460611b66913981565b5f83815b6020811015610c9b57600163ffffffff8516821c81169003610c3e57848160208110610c0157610c01613d4b565b602002013582604051602001610c21929190918252602082015260400190565b604051602081830303815290604052805190602001209150610c89565b81858260208110610c5157610c51613d4b565b6020020135604051602001610c70929190918252602082015260400190565b6040516020818303038152906040528051906020012091505b80610c9381613d73565b915050610bd3565b50949350505050565b60a3546001600160a01b03163314610ccf576040516357b738d160e11b815260040160405180910390fd5b8051825114610cf15760405163434f49f560e11b815260040160405180910390fd5b5f5b8251811015610d4a57610d38838281518110610d1157610d11613d4b565b6020026020010151838381518110610d2b57610d2b613d4b565b60200260200101516125f4565b80610d4281613d73565b915050610cf3565b505050565b5f54610100900460ff1615808015610d6d57505f54600160ff909116105b80610d865750303b158015610d8657505f5460ff166001145b610dab5760405162461bcd60e51b8152600401610da290613d8b565b60405180910390fd5b5f805460ff191660011790558015610dcc575f805461ff0019166101001790555b60688054610100600160c81b03191661010063ffffffff8d160265010000000000600160c81b03191617600160281b6001600160a01b038a81169190910291909117909155606c80546001600160a01b03199081168984161790915560a380549091168683161790558916610e975763ffffffff881615610e6057604051630d43a60960e11b815260040160405180910390fd5b6001600160a01b038316151580610e745750815b15610e9257604051630e6e237560e11b815260040160405180910390fd5b610fd2565b606d805463ffffffff8a16600160a01b026001600160c01b03199091166001600160a01b038c1617179055606e610ece8682613e1e565b506001600160a01b038316610f9a57811515600103610f0057604051630e6e237560e11b815260040160405180910390fd5b610f755f801b6012604051602001610f6191906060808252600d908201526c2bb930b83832b21022ba3432b960991b608082015260a060208201819052600490820152630ae8aa8960e31b60c082015260ff91909116604082015260e00190565b6040516020818303038152906040526126a2565b606f80546001600160a01b0319166001600160a01b0392909216919091179055610fd2565b606f80546001600160a01b0319166001600160a01b0385169081179091555f90815260a260205260409020805460ff19168315151790555b610fda61271c565b801561101f575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b6001600160a01b038083165f908152606b602090815260409182902082518084019093525463ffffffff81168352600160201b9004909216918101829052906110875760405163828d566360e01b815260040160405180910390fd5b5f606a5f835f015184602001516040516020016110a5929190613ce9565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b039081169150841681036110f85760405163e273c4a160e01b815260040160405180910390fd5b611102848461274a565b61110d8133856127da565b604080513381526001600160a01b0386811660208301528316818301526060810185905290517fb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a99181900360800190a150505050565b5f808686604051602001611178929190613ce9565b6040516020818303038152906040528051906020012090505f60ff60f81b308360405180611ba00160405280611b668152602001614460611b6691398989896040516020016111c993929190613ed9565b60408051601f19818403018152908290526111e79291602001613f11565b6040516020818303038152906040528051906020012060405160200161123f94939291906001600160f81b031994909416845260609290921b6001600160601b03191660018401526015830152603582015260550190565b60408051808303601f19018152919052805160209091012098975050505050505050565b60a3546001600160a01b0316331461128e576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b038082165f908152606b6020908152604080832081518083018352905463ffffffff8116808352600160201b909104909516818401819052915190946112de9390929101613ce9565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b0316158061133257505f818152606a60205260409020546001600160a01b038481169116145b156113505760405163e0c897a760e01b815260040160405180910390fd5b6001600160a01b0383165f818152606b6020908152604080832080546001600160c01b031916905560a2825291829020805460ff1916905590519182527fc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063910160405180910390a1505050565b60685460ff16156113e157604051630bc011ff60e21b815260040160405180910390fd5b606f546001600160a01b031661140a5760405163dde3cda760e01b815260040160405180910390fd5b606f54611420906001600160a01b03168561274a565b61142e8686868686866120c3565b505050505050565b60a3546001600160a01b03163314611461576040516357b738d160e11b815260040160405180910390fd5b606d546001600160a01b031661148a57604051634cb4711360e11b815260040160405180910390fd5b606f80546001600160a01b0319166001600160a01b0384169081179091555f81815260a26020908152604091829020805460ff19168515159081179091558251938452908301527fc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82910160405180910390a15050565b606061150b8261233d565b611514836123f8565b61151d846124ac565b60405160200161152f93929190613ed9565b6040516020818303038152906040529050919050565b5f8061155b600160201b63ffffffff8516613f3f565b61156b9063ffffffff8616613f56565b600881901c5f90815260696020526040902054600160ff9092169190911b90811614949350505050565b60685460ff16156115b957604051630bc011ff60e21b815260040160405180910390fd5b60685463ffffffff86811661010090920416146115e9576040516302caf51760e11b815260040160405180910390fd5b61161c8c8c8c8c8c6116175f8e8e8e8e8e8e8e60405161160a929190613f69565b6040518091039020610906565b612867565b6001600160a01b03861661170357606f546001600160a01b03166116e7575f6001600160a01b03851684825b6040519080825280601f01601f191660200182016040528015611672576020820181803683370190505b506040516116809190613f78565b5f6040518083038185875af1925050503d805f81146116ba576040519150601f19603f3d011682016040523d82523d5f602084013e6116bf565b606091505b50509050806116e157604051630ce8f45160e31b815260040160405180910390fd5b5061192c565b606f546116fe906001600160a01b031685856127da565b61192c565b606d546001600160a01b0387811691161480156117315750606d5463ffffffff888116600160a01b90920416145b15611748575f6001600160a01b0385168482611648565b60685463ffffffff610100909104811690881603611774576116fe6001600160a01b03871685856129c5565b5f8787604051602001611788929190613ce9565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b03168061191e575f6117fd8386868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506126a292505050565b905061180a8188886127da565b80606a5f8581526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060405180604001604052808b63ffffffff1681526020018a6001600160a01b0316815250606b5f836001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a81548163ffffffff021916908363ffffffff1602179055506020820151815f0160046101000a8154816001600160a01b0302191690836001600160a01b031602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398a8a838888604051611910959493929190613fbb565b60405180910390a150611929565b6119298187876127da565b50505b604080518b815263ffffffff891660208201526001600160a01b0388811682840152861660608201526080810185905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a1505050505050505050505050565b60685460ff16156119ba57604051630bc011ff60e21b815260040160405180910390fd5b6119c2612a28565b60685463ffffffff6101009091048116908816036119f3576040516302caf51760e11b815260040160405180910390fd5b5f806060876001600160a01b038816611ad657883414611a265760405163b89240f560e01b815260040160405180910390fd5b606d54606e80546001600160a01b0383169650600160a01b90920463ffffffff16945090611a5390613d13565b80601f0160208091040260200160405190810160405280929190818152602001828054611a7f90613d13565b8015611aca5780601f10611aa157610100808354040283529160200191611aca565b820191905f5260205f20905b815481529060010190602001808311611aad57829003601f168201915b50505050509150611ca3565b3415611af55760405163798ee6f160e01b815260040160405180910390fd5b606f546001600160a01b0390811690891603611b1a57611b15888a61274a565b611ca3565b6001600160a01b038089165f908152606b602090815260409182902082518084019093525463ffffffff81168352600160201b90049092169181018290529015611b7957611b68898b61274a565b602081015181519095509350611c96565b8515611b8b57611b8b898b8989612a81565b6040516370a0823160e01b81523060048201525f906001600160a01b038b16906370a0823190602401602060405180830381865afa158015611bcf573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611bf39190613ff3565b9050611c0a6001600160a01b038b1633308e612d32565b6040516370a0823160e01b81523060048201525f906001600160a01b038c16906370a0823190602401602060405180830381865afa158015611c4e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c729190613ff3565b9050611c7e828261400a565b6068548c9850610100900463ffffffff169650935050505b611c9f89611500565b9250505b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b5f84868e8e8688605354604051611ce298979695949392919061401d565b60405180910390a1611d08611d035f85878f8f878980519060200120610906565b612d6a565b8615611d1657611d16612560565b50505050611d2360018055565b50505050505050565b60a3546001600160a01b03163314611d57576040516357b738d160e11b815260040160405180910390fd5b60a380546001600160a01b0319166001600160a01b0383169081179091556040519081527f32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd679060200160405180910390a150565b60685460ff1615611dcf57604051630bc011ff60e21b815260040160405180910390fd5b60685463ffffffff8681166101009092041614611dff576040516302caf51760e11b815260040160405180910390fd5b611e218c8c8c8c8c61161760018e8e8e8e8e8e8e60405161160a929190613f69565b606f545f906001600160a01b0316611ed457846001600160a01b031684888a8686604051602401611e559493929190614087565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b17905251611e8a9190613f78565b5f6040518083038185875af1925050503d805f8114611ec4576040519150601f19603f3d011682016040523d82523d5f602084013e611ec9565b606091505b505080915050611f85565b606f54611eeb906001600160a01b031686866127da565b846001600160a01b031687898585604051602401611f0c9493929190614087565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b17905251611f419190613f78565b5f604051808303815f865af19150503d805f8114611f7a576040519150601f19603f3d011682016040523d82523d5f602084013e611f7f565b606091505b50909150505b80611fa3576040516337e391c360e01b815260040160405180910390fd5b604080518c815263ffffffff8a1660208201526001600160a01b0389811682840152871660608201526080810186905290517f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d9181900360a00190a150505050505050505050505050565b5f54610100900460ff161580801561202c57505f54600160ff909116105b806120455750303b15801561204557505f5460ff166001145b6120615760405162461bcd60e51b8152600401610da290613d8b565b5f805460ff191660011790558015612082575f805461ff0019166101001790555b60405163f57ac68360e01b815260040160405180910390fd5b60405180910390a150505050505050565b5f816120b9868686610bcf565b1495945050505050565b60685463ffffffff6101009091048116908716036120f4576040516302caf51760e11b815260040160405180910390fd5b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b6001606860019054906101000a900463ffffffff16338989898888605354604051612148999897969594939291906140c1565b60405180910390a161217f611d036001606860019054906101000a900463ffffffff16338a8a8a898960405161160a929190613f69565b821561142e5761142e612560565b6001600160a01b03831615806121aa57506001600160a01b038216155b156121c85760405163f6b2911f60e01b815260040160405180910390fd5b60685463ffffffff6101009091048116908516036121f95760405163658b23ad60e01b815260040160405180910390fd5b6001600160a01b038281165f908152606b6020526040902054600160201b90041615612238576040516317abdeeb60e21b815260040160405180910390fd5b5f848460405160200161224c929190613ce9565b60408051808303601f1901815282825280516020918201205f818152606a835283812080546001600160a01b0319166001600160a01b038a8116918217909255868601865263ffffffff8c81168089528c8416878a01818152848752606b89528987209a518b54915194166001600160c01b031990911617600160201b93909516929092029390931790975560a2855291859020805460ff191689151590811790915585519182529381019590955292840192909252606083015291507fdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce79060800160405180910390a15050505050565b60408051600481526024810182526020810180516001600160e01b03166306fdde0360e01b17905290516060915f9182916001600160a01b038616916123839190613f78565b5f60405180830381855afa9150503d805f81146123bb576040519150601f19603f3d011682016040523d82523d5f602084013e6123c0565b606091505b5091509150816123ef57604051806040016040528060078152602001664e4f5f4e414d4560c81b815250610b85565b610b8581612e52565b60408051600481526024810182526020810180516001600160e01b03166395d89b4160e01b17905290516060915f9182916001600160a01b0386169161243e9190613f78565b5f60405180830381855afa9150503d805f8114612476576040519150601f19603f3d011682016040523d82523d5f602084013e61247b565b606091505b5091509150816123ef57604051806040016040528060098152602001681393d7d4d6535093d360ba1b815250610b85565b60408051600481526024810182526020810180516001600160e01b031663313ce56760e01b17905290515f91829182916001600160a01b038616916124f19190613f78565b5f60405180830381855afa9150503d805f8114612529576040519150601f19603f3d011682016040523d82523d5f602084013e61252e565b606091505b5091509150818015612541575080516020145b61254c576012610b85565b80806020019051810190610b85919061412d565b6053546068805463ffffffff909216600160c81b0263ffffffff60c81b1990921691909117908190556001600160a01b03600160281b909104166333d6247d6125a7610a88565b6040518263ffffffff1660e01b81526004016125c591815260200190565b5f604051808303815f87803b1580156125dc575f80fd5b505af11580156125ee573d5f803e3d5ffd5b50505050565b5f612609600160201b63ffffffff8416613f3f565b6126199063ffffffff8516613f56565b600881901c5f8181526069602052604090208054600160ff851690811b91821892839055939450919291908082161561266557604051630631b5f760e31b815260040160405180910390fd5b6040805163ffffffff808a168252881660208201527f4a402ac607e71571d0543be8fcccc358a4df62dcd019a1e7f7e98ca6175e8f2a910161209b565b5f8060405180611ba00160405280611b668152602001614460611b669139836040516020016126d2929190613f11565b6040516020818303038152906040529050838151602083015ff591506001600160a01b038216612715576040516305f7d84960e51b815260040160405180910390fd5b5092915050565b5f54610100900460ff166127425760405162461bcd60e51b8152600401610da290614148565b610bae612fdb565b6001600160a01b0382165f90815260a2602052604090205460ff16156127835761277f6001600160a01b038316333084612d32565b5050565b604051632770a7eb60e21b8152336004820152602481018290526001600160a01b03831690639dc29fac906044015f604051808303815f87803b1580156127c8575f80fd5b505af115801561142e573d5f803e3d5ffd5b6001600160a01b0383165f90815260a2602052604090205460ff161561280e57610d4a6001600160a01b03841683836129c5565b6040516340c10f1960e01b81526001600160a01b038381166004830152602482018390528416906340c10f19906044015f604051808303815f87803b158015612855575f80fd5b505af1158015611d23573d5f803e3d5ffd5b606854604080516020808201879052818301869052825180830384018152606083019384905280519101206312bd9b1960e11b90925260648101919091525f91600160281b90046001600160a01b03169063257b3632906084016020604051808303815f875af11580156128dd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129019190613ff3565b9050805f0361292257604051622f6fad60e01b815260040160405180910390fd5b5f806801000000000000000087161561296657869150612944848a84896120ac565b612961576040516338105f3b60e21b815260040160405180910390fd5b6129b0565b602087901c612976816001614193565b9150879250612991612989868c86610bcf565b8a83896120ac565b6129ae576040516338105f3b60e21b815260040160405180910390fd5b505b6129ba8282613001565b505050505050505050565b6040516001600160a01b038316602482015260448101829052610d4a90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613073565b600260015403612a7a5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610da2565b6002600155565b5f612a8f60048284866141b0565b612a98916141d7565b9050632afa533160e01b6001600160e01b0319821601612bcc575f808080808080612ac6896004818d6141b0565b810190612ad39190614207565b96509650965096509650965096508a8514612b01576040516303fffc4b60e01b815260040160405180910390fd5b604080516001600160a01b0389811660248301528881166044830152606482018890526084820187905260ff861660a483015260c4820185905260e48083018590528351808403909101815261010490920183526020820180516001600160e01b031663d505accf60e01b1790529151918e1691612b7f9190613f78565b5f604051808303815f865af19150503d805f8114612bb8576040519150601f19603f3d011682016040523d82523d5f602084013e612bbd565b606091505b50505050505050505050610873565b6001600160e01b031981166323f2ebc360e21b14612bfd57604051637141605d60e11b815260040160405180910390fd5b5f80808080808080612c128a6004818e6141b0565b810190612c1f9190614256565b975097509750975097509750975097508c6001600160a01b0316638fcbaf0c60e01b8989898989898989604051602401612ca49897969594939291906001600160a01b039889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612ce29190613f78565b5f604051808303815f865af19150503d805f8114612d1b576040519150601f19603f3d011682016040523d82523d5f602084013e612d20565b606091505b50505050505050505050505050505050565b6040516001600160a01b03808516602483015283166044820152606481018290526125ee9085906323b872dd60e01b906084016129f1565b806001612d79602060026143b4565b612d83919061400a565b60535410612da4576040516377ae67b360e11b815260040160405180910390fd5b5f60535f8154612db390613d73565b918290555090505f5b6020811015612e43578082901c600116600103612def578260338260208110612de757612de7613d4b565b015550505050565b60338160208110612e0257612e02613d4b565b015460408051602081019290925281018490526060016040516020818303038152906040528051906020012092508080612e3b90613d73565b915050612dbc565b50610d4a6143bf565b60018055565b60606040825110612e71578180602001905181019061080491906143d3565b8151602003612fa8575f5b602081108015612eab5750828181518110612e9957612e99613d4b565b01602001516001600160f81b03191615155b15612ec25780612eba81613d73565b915050612e7c565b805f03612ef95750506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b6020820152919050565b5f816001600160401b03811115612f1257612f12613490565b6040519080825280601f01601f191660200182016040528015612f3c576020820181803683370190505b5090505f5b82811015612fa057848181518110612f5b57612f5b613d4b565b602001015160f81c60f81b828281518110612f7857612f78613d4b565b60200101906001600160f81b03191690815f1a90535080612f9881613d73565b915050612f41565b509392505050565b50506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b602082015290565b919050565b5f54610100900460ff16612e4c5760405162461bcd60e51b8152600401610da290614148565b5f613016600160201b63ffffffff8416613f3f565b6130269063ffffffff8516613f56565b600881901c5f8181526069602052604081208054600160ff861690811b91821892839055949550929392918183169003611d2357604051630c8d9eab60e31b815260040160405180910390fd5b5f6130c7826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166131469092919063ffffffff16565b905080515f14806130e75750808060200190518101906130e79190614444565b610d4a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610da2565b6060610b8584845f85855f80866001600160a01b0316858760405161316b9190613f78565b5f6040518083038185875af1925050503d805f81146131a5576040519150601f19603f3d011682016040523d82523d5f602084013e6131aa565b606091505b50915091506131bb878383876131c6565b979650505050505050565b606083156132345782515f0361322d576001600160a01b0385163b61322d5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610da2565b5081610b85565b610b8583838151156132495781518083602001fd5b8060405162461bcd60e51b8152600401610da291906133d4565b803563ffffffff81168114612fd6575f80fd5b6001600160a01b038116811461328a575f80fd5b50565b5f806040838503121561329e575f80fd5b6132a783613263565b915060208301356132b781613276565b809150509250929050565b801515811461328a575f80fd5b5f8083601f8401126132df575f80fd5b5081356001600160401b038111156132f5575f80fd5b60208301915083602082850101111561330c575f80fd5b9250929050565b5f805f805f60808688031215613327575f80fd5b61333086613263565b9450602086013561334081613276565b93506040860135613350816132c2565b925060608601356001600160401b0381111561336a575f80fd5b613376888289016132cf565b969995985093965092949392505050565b5f5b838110156133a1578181015183820152602001613389565b50505f910152565b5f81518084526133c0816020860160208601613387565b601f01601f19169290920160200192915050565b602081525f6133e660208301846133a9565b9392505050565b5f602082840312156133fd575f80fd5b81356133e681613276565b60ff8116811461328a575f80fd5b5f805f805f805f60e0888a03121561342c575f80fd5b873561343781613408565b965061344560208901613263565b9550604088013561345581613276565b945061346360608901613263565b9350608088013561347381613276565b9699959850939692959460a0840135945060c09093013592915050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156134cc576134cc613490565b604052919050565b5f6001600160401b038211156134ec576134ec613490565b5060051b60200190565b5f82601f830112613505575f80fd5b8135602061351a613515836134d4565b6134a4565b82815260059290921b84018101918181019086841115613538575f80fd5b8286015b8481101561355a5761354d81613263565b835291830191830161353c565b509695505050505050565b5f82601f830112613574575f80fd5b81356020613584613515836134d4565b82815260059290921b840181019181810190868411156135a2575f80fd5b8286015b8481101561355a5780356135b981613276565b83529183019183016135a6565b5f805f80608085870312156135d9575f80fd5b84356001600160401b03808211156135ef575f80fd5b6135fb888389016134f6565b9550602091508187013581811115613611575f80fd5b61361d89828a01613565565b955050604087013581811115613631575f80fd5b61363d89828a01613565565b945050606087013581811115613651575f80fd5b87019050601f81018813613663575f80fd5b8035613671613515826134d4565b81815260059190911b8201830190838101908a83111561368f575f80fd5b928401925b828410156136b65783356136a7816132c2565b82529284019290840190613694565b979a9699509497505050505050565b5f805f606084860312156136d7575f80fd5b6136e084613263565b925060208401356136f081613276565b9150604084013561370081613276565b809150509250925092565b5f6020828403121561371b575f80fd5b5035919050565b806104008101831015610804575f80fd5b5f805f6104408486031215613746575f80fd5b833592506137578560208601613722565b91506137666104208501613263565b90509250925092565b5f8060408385031215613780575f80fd5b82356001600160401b0380821115613796575f80fd5b6137a2868387016134f6565b935060208501359150808211156137b7575f80fd5b506137c4858286016134f6565b9150509250929050565b5f6001600160401b038211156137e6576137e6613490565b50601f01601f191660200190565b5f82601f830112613803575f80fd5b8135613811613515826137ce565b818152846020838601011115613825575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f805f805f6101208a8c03121561385a575f80fd5b6138638a613263565b985060208a013561387381613276565b975061388160408b01613263565b965060608a013561389181613276565b955060808a01356138a181613276565b945060a08a01356001600160401b038111156138bb575f80fd5b6138c78c828d016137f4565b94505060c08a01356138d881613276565b925060e08a01356138e881613276565b91506101008a01356138f9816132c2565b809150509295985092959850929598565b5f806040838503121561391b575f80fd5b823561392681613276565b946020939093013593505050565b5f805f805f60a08688031215613948575f80fd5b61395186613263565b9450602086013561396181613276565b935060408601356001600160401b038082111561397c575f80fd5b61398889838a016137f4565b9450606088013591508082111561399d575f80fd5b506139aa888289016137f4565b92505060808601356139bb81613408565b809150509295509295909350565b5f805f805f8060a087890312156139de575f80fd5b6139e787613263565b955060208701356139f781613276565b9450604087013593506060870135613a0e816132c2565b925060808701356001600160401b03811115613a28575f80fd5b613a3489828a016132cf565b979a9699509497509295939492505050565b5f8060408385031215613a57575f80fd5b8235613a6281613276565b915060208301356132b7816132c2565b5f8060408385031215613a83575f80fd5b613a8c83613263565b9150613a9a60208401613263565b90509250929050565b5f805f805f805f805f805f806109208d8f031215613abf575f80fd5b613ac98e8e613722565b9b50613ad98e6104008f01613722565b9a506108008d013599506108208d013598506108408d01359750613b006108608e01613263565b9650613b106108808e0135613276565b6108808d01359550613b256108a08e01613263565b9450613b356108c08e0135613276565b6108c08d013593506108e08d013592506001600160401b036109008e01351115613b5d575f80fd5b613b6e8e6109008f01358f016132cf565b81935080925050509295989b509295989b509295989b565b5f805f805f805f60c0888a031215613b9c575f80fd5b613ba588613263565b96506020880135613bb581613276565b9550604088013594506060880135613bcc81613276565b93506080880135613bdc816132c2565b925060a08801356001600160401b03811115613bf6575f80fd5b613c028a828b016132cf565b989b979a50959850939692959293505050565b5f805f805f8060c08789031215613c2a575f80fd5b613c3387613263565b95506020870135613c4381613276565b9450613c5160408801613263565b93506060870135613c6181613276565b92506080870135613c7181613276565b915060a08701356001600160401b03811115613c8b575f80fd5b613c9789828a016137f4565b9150509295509295509295565b5f805f806104608587031215613cb8575f80fd5b84359350613cc98660208701613722565b9250613cd86104208601613263565b939692955092936104400135925050565b60e09290921b6001600160e01b031916825260601b6001600160601b031916600482015260180190565b600181811c90821680613d2757607f821691505b602082108103613d4557634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b5f60018201613d8457613d84613d5f565b5060010190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b601f821115610d4a575f81815260208120601f850160051c81016020861015613dff5750805b601f850160051c820191505b8181101561142e57828155600101613e0b565b81516001600160401b03811115613e3757613e37613490565b613e4b81613e458454613d13565b84613dd9565b602080601f831160018114613e7e575f8415613e675750858301515b5f19600386901b1c1916600185901b17855561142e565b5f85815260208120601f198616915b82811015613eac57888601518255948401946001909101908401613e8d565b5085821015613ec957878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b606081525f613eeb60608301866133a9565b8281036020840152613efd81866133a9565b91505060ff83166040830152949350505050565b5f8351613f22818460208801613387565b835190830190613f36818360208801613387565b01949350505050565b808202811582820484141761080457610804613d5f565b8082018082111561080457610804613d5f565b818382375f9101908152919050565b5f8251613f89818460208701613387565b9190910192915050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b63ffffffff861681526001600160a01b038581166020830152841660408201526080606082018190525f906131bb9083018486613f93565b5f60208284031215614003575f80fd5b5051919050565b8181038181111561080457610804613d5f565b60ff8916815263ffffffff88811660208301526001600160a01b03888116604084015287821660608401528616608083015260a0820185905261010060c083018190525f9161406e848301876133a9565b925080851660e085015250509998505050505050505050565b6001600160a01b038516815263ffffffff841660208201526060604082018190525f906140b79083018486613f93565b9695505050505050565b60ff8a16815263ffffffff89811660208301526001600160a01b03898116604084015288821660608401528716608083015260a0820186905261010060c083018190525f916141138483018789613f93565b925080851660e085015250509a9950505050505050505050565b5f6020828403121561413d575f80fd5b81516133e681613408565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b63ffffffff81811683821601908082111561271557612715613d5f565b5f80858511156141be575f80fd5b838611156141ca575f80fd5b5050820193919092039150565b6001600160e01b031981358181169160048510156141ff5780818660040360031b1b83161692505b505092915050565b5f805f805f805f60e0888a03121561421d575f80fd5b873561422881613276565b9650602088013561423881613276565b95506040880135945060608801359350608088013561347381613408565b5f805f805f805f80610100898b03121561426e575f80fd5b883561427981613276565b9750602089013561428981613276565b9650604089013595506060890135945060808901356142a7816132c2565b935060a08901356142b781613408565b979a969950949793969295929450505060c08201359160e0013590565b600181815b8085111561430e57815f19048211156142f4576142f4613d5f565b8085161561430157918102915b93841c93908002906142d9565b509250929050565b5f8261432457506001610804565b8161433057505f610804565b816001811461434657600281146143505761436c565b6001915050610804565b60ff84111561436157614361613d5f565b50506001821b610804565b5060208310610133831016604e8410600b841016171561438f575081810a610804565b61439983836142d4565b805f19048211156143ac576143ac613d5f565b029392505050565b5f6133e68383614316565b634e487b7160e01b5f52600160045260245ffd5b5f602082840312156143e3575f80fd5b81516001600160401b038111156143f8575f80fd5b8201601f81018413614408575f80fd5b8051614416613515826137ce565b81815285602083850101111561442a575f80fd5b61443b826020830160208601613387565b95945050505050565b5f60208284031215614454575f80fd5b81516133e6816132c256fe6101006040523480156200001257600080fd5b5060405162001b6638038062001b6683398101604081905262000035916200028d565b82826003620000458382620003a1565b506004620000548282620003a1565b50503360c0525060ff811660e052466080819052620000739062000080565b60a052506200046d915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f620000ad6200012e565b805160209182012060408051808201825260018152603160f81b90840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b6060600380546200013f9062000312565b80601f01602080910402602001604051908101604052809291908181526020018280546200016d9062000312565b8015620001be5780601f106200019257610100808354040283529160200191620001be565b820191906000526020600020905b815481529060010190602001808311620001a057829003601f168201915b5050505050905090565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001f057600080fd5b81516001600160401b03808211156200020d576200020d620001c8565b604051601f8301601f19908116603f01168101908282118183101715620002385762000238620001c8565b816040528381526020925086838588010111156200025557600080fd5b600091505b838210156200027957858201830151818301840152908201906200025a565b600093810190920192909252949350505050565b600080600060608486031215620002a357600080fd5b83516001600160401b0380821115620002bb57600080fd5b620002c987838801620001de565b94506020860151915080821115620002e057600080fd5b50620002ef86828701620001de565b925050604084015160ff811681146200030757600080fd5b809150509250925092565b600181811c908216806200032757607f821691505b6020821081036200034857634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200039c57600081815260208120601f850160051c81016020861015620003775750805b601f850160051c820191505b81811015620003985782815560010162000383565b5050505b505050565b81516001600160401b03811115620003bd57620003bd620001c8565b620003d581620003ce845462000312565b846200034e565b602080601f8311600181146200040d5760008415620003f45750858301515b600019600386901b1c1916600185901b17855562000398565b600085815260208120601f198616915b828110156200043e578886015182559484019460019091019084016200041d565b50858210156200045d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e0516116aa620004bc6000396000610237015260008181610307015281816105c001526106a70152600061053a015260008181610379015261050401526116aa6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063a457c2d71161008c578063d505accf11610066578063d505accf1461039b578063dd62ed3e146103ae578063ffa1ad74146103f457600080fd5b8063a457c2d71461034e578063a9059cbb14610361578063cd0d00961461037457600080fd5b806395d89b41116100bd57806395d89b41146102e75780639dc29fac146102ef578063a3c573eb1461030257600080fd5b806370a08231146102915780637ecebe00146102c757600080fd5b806330adf81f1161012f5780633644e515116101145780633644e51514610261578063395093511461026957806340c10f191461027c57600080fd5b806330adf81f14610209578063313ce5671461023057600080fd5b806318160ddd1161016057806318160ddd146101bd57806320606b70146101cf57806323b872dd146101f657600080fd5b806306fdde031461017c578063095ea7b31461019a575b600080fd5b610184610430565b60405161019191906113e4565b60405180910390f35b6101ad6101a8366004611479565b6104c2565b6040519015158152602001610191565b6002545b604051908152602001610191565b6101c17f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81565b6101ad6102043660046114a3565b6104dc565b6101c17f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b60405160ff7f0000000000000000000000000000000000000000000000000000000000000000168152602001610191565b6101c1610500565b6101ad610277366004611479565b61055c565b61028f61028a366004611479565b6105a8565b005b6101c161029f3660046114df565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101c16102d53660046114df565b60056020526000908152604090205481565b610184610680565b61028f6102fd366004611479565b61068f565b6103297f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610191565b6101ad61035c366004611479565b61075e565b6101ad61036f366004611479565b61082f565b6101c17f000000000000000000000000000000000000000000000000000000000000000081565b61028f6103a9366004611501565b61083d565b6101c16103bc366004611574565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101846040518060400160405280600181526020017f310000000000000000000000000000000000000000000000000000000000000081525081565b60606003805461043f906115a7565b80601f016020809104026020016040519081016040528092919081815260200182805461046b906115a7565b80156104b85780601f1061048d576101008083540402835291602001916104b8565b820191906000526020600020905b81548152906001019060200180831161049b57829003601f168201915b5050505050905090565b6000336104d0818585610b73565b60019150505b92915050565b6000336104ea858285610d27565b6104f5858585610dfe565b506001949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000004614610537576105324661106d565b905090565b507f000000000000000000000000000000000000000000000000000000000000000090565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906104d090829086906105a3908790611629565b610b73565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610672576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d4272696467650000000000000000000000000000000060648201526084015b60405180910390fd5b61067c8282611135565b5050565b60606004805461043f906115a7565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610754576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f546f6b656e577261707065643a3a6f6e6c794272696467653a204e6f7420506f60448201527f6c79676f6e5a6b45564d427269646765000000000000000000000000000000006064820152608401610669565b61067c8282611228565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919083811015610822576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610669565b6104f58286868403610b73565b6000336104d0818585610dfe565b834211156108cc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f546f6b656e577261707065643a3a7065726d69743a204578706972656420706560448201527f726d6974000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8716600090815260056020526040812080547f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9918a918a918a9190866109268361163c565b9091555060408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090506000610991610500565b6040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281019190915260428101839052606201604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120600080855291840180845281905260ff89169284019290925260608301879052608083018690529092509060019060a0016020604051602081039080840390855afa158015610a55573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811615801590610ad057508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b610b5c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f546f6b656e577261707065643a3a7065726d69743a20496e76616c696420736960448201527f676e6174757265000000000000000000000000000000000000000000000000006064820152608401610669565b610b678a8a8a610b73565b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610cb8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610df85781811015610deb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610669565b610df88484848403610b73565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610ea1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff8216610f44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610ffa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610df8565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f611098610430565b8051602091820120604080518082018252600181527f310000000000000000000000000000000000000000000000000000000000000090840152805192830193909352918101919091527fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc66060820152608081018390523060a082015260c001604051602081830303815290604052805190602001209050919050565b73ffffffffffffffffffffffffffffffffffffffff82166111b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610669565b80600260008282546111c49190611629565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b73ffffffffffffffffffffffffffffffffffffffff82166112cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090205481811015611381576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610669565b73ffffffffffffffffffffffffffffffffffffffff83166000818152602081815260408083208686039055600280548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610d1a565b600060208083528351808285015260005b81811015611411578581018301518582016040015282016113f5565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461147457600080fd5b919050565b6000806040838503121561148c57600080fd5b61149583611450565b946020939093013593505050565b6000806000606084860312156114b857600080fd5b6114c184611450565b92506114cf60208501611450565b9150604084013590509250925092565b6000602082840312156114f157600080fd5b6114fa82611450565b9392505050565b600080600080600080600060e0888a03121561151c57600080fd5b61152588611450565b965061153360208901611450565b95506040880135945060608801359350608088013560ff8116811461155757600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561158757600080fd5b61159083611450565b915061159e60208401611450565b90509250929050565b600181811c908216806115bb57607f821691505b6020821081036115f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156104d6576104d66115fa565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361166d5761166d6115fa565b506001019056fea26469706673582212208d88fee561cff7120d381c345cfc534cef8229a272dc5809d4bbb685ad67141164736f6c63430008110033a264697066735822122037f2d9d14450b30ae5085e3e49ec180f0b199039dce098c7f027c2222e9455b264736f6c63430008140033",
}

// Bridgel2sovereignchainABI is the input ABI used to generate the binding from.
// Deprecated: Use Bridgel2sovereignchainMetaData.ABI instead.
var Bridgel2sovereignchainABI = Bridgel2sovereignchainMetaData.ABI

// Bridgel2sovereignchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Bridgel2sovereignchainMetaData.Bin instead.
var Bridgel2sovereignchainBin = Bridgel2sovereignchainMetaData.Bin

// DeployBridgel2sovereignchain deploys a new Ethereum contract, binding an instance of Bridgel2sovereignchain to it.
func DeployBridgel2sovereignchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridgel2sovereignchain, error) {
	parsed, err := Bridgel2sovereignchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Bridgel2sovereignchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridgel2sovereignchain{Bridgel2sovereignchainCaller: Bridgel2sovereignchainCaller{contract: contract}, Bridgel2sovereignchainTransactor: Bridgel2sovereignchainTransactor{contract: contract}, Bridgel2sovereignchainFilterer: Bridgel2sovereignchainFilterer{contract: contract}}, nil
}

// Bridgel2sovereignchain is an auto generated Go binding around an Ethereum contract.
type Bridgel2sovereignchain struct {
	Bridgel2sovereignchainCaller     // Read-only binding to the contract
	Bridgel2sovereignchainTransactor // Write-only binding to the contract
	Bridgel2sovereignchainFilterer   // Log filterer for contract events
}

// Bridgel2sovereignchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Bridgel2sovereignchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Bridgel2sovereignchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Bridgel2sovereignchainSession struct {
	Contract     *Bridgel2sovereignchain // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Bridgel2sovereignchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Bridgel2sovereignchainCallerSession struct {
	Contract *Bridgel2sovereignchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// Bridgel2sovereignchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Bridgel2sovereignchainTransactorSession struct {
	Contract     *Bridgel2sovereignchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// Bridgel2sovereignchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type Bridgel2sovereignchainRaw struct {
	Contract *Bridgel2sovereignchain // Generic contract binding to access the raw methods on
}

// Bridgel2sovereignchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainCallerRaw struct {
	Contract *Bridgel2sovereignchainCaller // Generic read-only contract binding to access the raw methods on
}

// Bridgel2sovereignchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Bridgel2sovereignchainTransactorRaw struct {
	Contract *Bridgel2sovereignchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgel2sovereignchain creates a new instance of Bridgel2sovereignchain, bound to a specific deployed contract.
func NewBridgel2sovereignchain(address common.Address, backend bind.ContractBackend) (*Bridgel2sovereignchain, error) {
	contract, err := bindBridgel2sovereignchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchain{Bridgel2sovereignchainCaller: Bridgel2sovereignchainCaller{contract: contract}, Bridgel2sovereignchainTransactor: Bridgel2sovereignchainTransactor{contract: contract}, Bridgel2sovereignchainFilterer: Bridgel2sovereignchainFilterer{contract: contract}}, nil
}

// NewBridgel2sovereignchainCaller creates a new read-only instance of Bridgel2sovereignchain, bound to a specific deployed contract.
func NewBridgel2sovereignchainCaller(address common.Address, caller bind.ContractCaller) (*Bridgel2sovereignchainCaller, error) {
	contract, err := bindBridgel2sovereignchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainCaller{contract: contract}, nil
}

// NewBridgel2sovereignchainTransactor creates a new write-only instance of Bridgel2sovereignchain, bound to a specific deployed contract.
func NewBridgel2sovereignchainTransactor(address common.Address, transactor bind.ContractTransactor) (*Bridgel2sovereignchainTransactor, error) {
	contract, err := bindBridgel2sovereignchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainTransactor{contract: contract}, nil
}

// NewBridgel2sovereignchainFilterer creates a new log filterer instance of Bridgel2sovereignchain, bound to a specific deployed contract.
func NewBridgel2sovereignchainFilterer(address common.Address, filterer bind.ContractFilterer) (*Bridgel2sovereignchainFilterer, error) {
	contract, err := bindBridgel2sovereignchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainFilterer{contract: contract}, nil
}

// bindBridgel2sovereignchain binds a generic wrapper to an already deployed contract.
func bindBridgel2sovereignchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Bridgel2sovereignchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgel2sovereignchain.Contract.Bridgel2sovereignchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Bridgel2sovereignchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Bridgel2sovereignchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridgel2sovereignchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.contract.Transact(opts, method, params...)
}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BASEINITBYTECODEWRAPPEDTOKEN(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "BASE_INIT_BYTECODE_WRAPPED_TOKEN")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BASEINITBYTECODEWRAPPEDTOKEN() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.BASEINITBYTECODEWRAPPEDTOKEN(&_Bridgel2sovereignchain.CallOpts)
}

// BASEINITBYTECODEWRAPPEDTOKEN is a free data retrieval call binding the contract method 0x83c43a55.
//
// Solidity: function BASE_INIT_BYTECODE_WRAPPED_TOKEN() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BASEINITBYTECODEWRAPPEDTOKEN() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.BASEINITBYTECODEWRAPPEDTOKEN(&_Bridgel2sovereignchain.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) WETHToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "WETHToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) WETHToken() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.WETHToken(&_Bridgel2sovereignchain.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) WETHToken() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.WETHToken(&_Bridgel2sovereignchain.CallOpts)
}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ActivateEmergencyState(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "activateEmergencyState")

	if err != nil {
		return err
	}

	return err

}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ActivateEmergencyState() error {
	return _Bridgel2sovereignchain.Contract.ActivateEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// ActivateEmergencyState is a free data retrieval call binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ActivateEmergencyState() error {
	return _Bridgel2sovereignchain.Contract.ActivateEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BridgeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "bridgeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.BridgeManager(&_Bridgel2sovereignchain.CallOpts)
}

// BridgeManager is a free data retrieval call binding the contract method 0x14cc01a0.
//
// Solidity: function bridgeManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BridgeManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.BridgeManager(&_Bridgel2sovereignchain.CallOpts)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) CalculateRoot(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "calculateRoot", leafHash, smtProof, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.CalculateRoot(&_Bridgel2sovereignchain.CallOpts, leafHash, smtProof, index)
}

// CalculateRoot is a free data retrieval call binding the contract method 0x83f24403.
//
// Solidity: function calculateRoot(bytes32 leafHash, bytes32[32] smtProof, uint32 index) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) CalculateRoot(leafHash [32]byte, smtProof [32][32]byte, index uint32) ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.CalculateRoot(&_Bridgel2sovereignchain.CallOpts, leafHash, smtProof, index)
}

// CalculateTokenWrapperAddress is a free data retrieval call binding the contract method 0x7843298b.
//
// Solidity: function calculateTokenWrapperAddress(uint32 originNetwork, address originTokenAddress, address token) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) CalculateTokenWrapperAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "calculateTokenWrapperAddress", originNetwork, originTokenAddress, token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateTokenWrapperAddress is a free data retrieval call binding the contract method 0x7843298b.
//
// Solidity: function calculateTokenWrapperAddress(uint32 originNetwork, address originTokenAddress, address token) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) CalculateTokenWrapperAddress(originNetwork uint32, originTokenAddress common.Address, token common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.CalculateTokenWrapperAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress, token)
}

// CalculateTokenWrapperAddress is a free data retrieval call binding the contract method 0x7843298b.
//
// Solidity: function calculateTokenWrapperAddress(uint32 originNetwork, address originTokenAddress, address token) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) CalculateTokenWrapperAddress(originNetwork uint32, originTokenAddress common.Address, token common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.CalculateTokenWrapperAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress, token)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ClaimedBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "claimedBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.ClaimedBitMap(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.ClaimedBitMap(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) DeactivateEmergencyState(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "deactivateEmergencyState")

	if err != nil {
		return err
	}

	return err

}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) DeactivateEmergencyState() error {
	return _Bridgel2sovereignchain.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// DeactivateEmergencyState is a free data retrieval call binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) DeactivateEmergencyState() error {
	return _Bridgel2sovereignchain.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) DepositCount() (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.DepositCount(&_Bridgel2sovereignchain.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) DepositCount() (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.DepositCount(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GasTokenAddress() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenAddress(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GasTokenAddress() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenAddress(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GasTokenMetadata(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "gasTokenMetadata")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GasTokenMetadata() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenMetadata(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GasTokenMetadata() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenMetadata(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GasTokenNetwork() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenNetwork(&_Bridgel2sovereignchain.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GasTokenNetwork() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.GasTokenNetwork(&_Bridgel2sovereignchain.CallOpts)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetLeafValue(opts *bind.CallOpts, leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getLeafValue", leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetLeafValue(&_Bridgel2sovereignchain.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetLeafValue is a free data retrieval call binding the contract method 0x3e197043.
//
// Solidity: function getLeafValue(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes32 metadataHash) pure returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetLeafValue(leafType uint8, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadataHash [32]byte) ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetLeafValue(&_Bridgel2sovereignchain.CallOpts, leafType, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadataHash)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetRoot() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetRoot(&_Bridgel2sovereignchain.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetRoot() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetRoot(&_Bridgel2sovereignchain.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetTokenMetadata(&_Bridgel2sovereignchain.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.GetTokenMetadata(&_Bridgel2sovereignchain.CallOpts, token)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetTokenWrappedAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getTokenWrappedAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetTokenWrappedAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetTokenWrappedAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GlobalExitRootManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GlobalExitRootManager(&_Bridgel2sovereignchain.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GlobalExitRootManager(&_Bridgel2sovereignchain.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) IsClaimed(opts *bind.CallOpts, leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "isClaimed", leafIndex, sourceBridgeNetwork)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Bridgel2sovereignchain.Contract.IsClaimed(&_Bridgel2sovereignchain.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Bridgel2sovereignchain.Contract.IsClaimed(&_Bridgel2sovereignchain.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) IsEmergencyState() (bool, error) {
	return _Bridgel2sovereignchain.Contract.IsEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) IsEmergencyState() (bool, error) {
	return _Bridgel2sovereignchain.Contract.IsEmergencyState(&_Bridgel2sovereignchain.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) LastUpdatedDepositCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "lastUpdatedDepositCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) LastUpdatedDepositCount() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.LastUpdatedDepositCount(&_Bridgel2sovereignchain.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) LastUpdatedDepositCount() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.LastUpdatedDepositCount(&_Bridgel2sovereignchain.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) NetworkID() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.NetworkID(&_Bridgel2sovereignchain.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) NetworkID() (uint32, error) {
	return _Bridgel2sovereignchain.Contract.NetworkID(&_Bridgel2sovereignchain.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PolygonRollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "polygonRollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PolygonRollupManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PolygonRollupManager(&_Bridgel2sovereignchain.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PolygonRollupManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PolygonRollupManager(&_Bridgel2sovereignchain.CallOpts)
}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PrecalculatedWrapperAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "precalculatedWrapperAddress", originNetwork, originTokenAddress, name, symbol, decimals)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PrecalculatedWrapperAddress(originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PrecalculatedWrapperAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress, name, symbol, decimals)
}

// PrecalculatedWrapperAddress is a free data retrieval call binding the contract method 0xaaa13cc2.
//
// Solidity: function precalculatedWrapperAddress(uint32 originNetwork, address originTokenAddress, string name, string symbol, uint8 decimals) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PrecalculatedWrapperAddress(originNetwork uint32, originTokenAddress common.Address, name string, symbol string, decimals uint8) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PrecalculatedWrapperAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress, name, symbol, decimals)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) TokenInfoToWrappedToken(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "tokenInfoToWrappedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.TokenInfoToWrappedToken(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.TokenInfoToWrappedToken(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) VerifyMerkleProof(opts *bind.CallOpts, leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "verifyMerkleProof", leafHash, smtProof, index, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Bridgel2sovereignchain.Contract.VerifyMerkleProof(&_Bridgel2sovereignchain.CallOpts, leafHash, smtProof, index, root)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0xfb570834.
//
// Solidity: function verifyMerkleProof(bytes32 leafHash, bytes32[32] smtProof, uint32 index, bytes32 root) pure returns(bool)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) VerifyMerkleProof(leafHash [32]byte, smtProof [32][32]byte, index uint32, root [32]byte) (bool, error) {
	return _Bridgel2sovereignchain.Contract.VerifyMerkleProof(&_Bridgel2sovereignchain.CallOpts, leafHash, smtProof, index, root)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) WrappedAddressIsNotMintable(opts *bind.CallOpts, wrappedAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "wrappedAddressIsNotMintable", wrappedAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Bridgel2sovereignchain.Contract.WrappedAddressIsNotMintable(&_Bridgel2sovereignchain.CallOpts, wrappedAddress)
}

// WrappedAddressIsNotMintable is a free data retrieval call binding the contract method 0xc0f49163.
//
// Solidity: function wrappedAddressIsNotMintable(address wrappedAddress) view returns(bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) WrappedAddressIsNotMintable(wrappedAddress common.Address) (bool, error) {
	return _Bridgel2sovereignchain.Contract.WrappedAddressIsNotMintable(&_Bridgel2sovereignchain.CallOpts, wrappedAddress)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) WrappedTokenToTokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "wrappedTokenToTokenInfo", arg0)

	outstruct := new(struct {
		OriginNetwork      uint32
		OriginTokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OriginNetwork = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.OriginTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Bridgel2sovereignchain.Contract.WrappedTokenToTokenInfo(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Bridgel2sovereignchain.Contract.WrappedTokenToTokenInfo(&_Bridgel2sovereignchain.CallOpts, arg0)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeAsset(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeAsset(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeMessage(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeMessage(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeMessageWETH(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BridgeMessageWETH(&_Bridgel2sovereignchain.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ClaimAsset(&_Bridgel2sovereignchain.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ClaimAsset(&_Bridgel2sovereignchain.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ClaimMessage(&_Bridgel2sovereignchain.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ClaimMessage(&_Bridgel2sovereignchain.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize(&_Bridgel2sovereignchain.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize is a paid mutator transaction binding the contract method 0x8c0dd470.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize(&_Bridgel2sovereignchain.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) Initialize0(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "initialize0", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize0(&_Bridgel2sovereignchain.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) Initialize0(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize0(&_Bridgel2sovereignchain.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0x9e76158f.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) MigrateLegacyToken(opts *bind.TransactOpts, legacyTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "migrateLegacyToken", legacyTokenAddress, amount)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0x9e76158f.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.MigrateLegacyToken(&_Bridgel2sovereignchain.TransactOpts, legacyTokenAddress, amount)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0x9e76158f.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.MigrateLegacyToken(&_Bridgel2sovereignchain.TransactOpts, legacyTokenAddress, amount)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) RemoveLegacySovereignTokenAddress(opts *bind.TransactOpts, legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "removeLegacySovereignTokenAddress", legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, legacySovereignTokenAddress)
}

// RemoveLegacySovereignTokenAddress is a paid mutator transaction binding the contract method 0xb4586962.
//
// Solidity: function removeLegacySovereignTokenAddress(address legacySovereignTokenAddress) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) RemoveLegacySovereignTokenAddress(legacySovereignTokenAddress common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.RemoveLegacySovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, legacySovereignTokenAddress)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetBridgeManager(opts *bind.TransactOpts, _bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setBridgeManager", _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetBridgeManager(&_Bridgel2sovereignchain.TransactOpts, _bridgeManager)
}

// SetBridgeManager is a paid mutator transaction binding the contract method 0xeabd372a.
//
// Solidity: function setBridgeManager(address _bridgeManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetBridgeManager(_bridgeManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetBridgeManager(&_Bridgel2sovereignchain.TransactOpts, _bridgeManager)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetMultipleSovereignTokenAddress(opts *bind.TransactOpts, originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setMultipleSovereignTokenAddress", originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetMultipleSovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetMultipleSovereignTokenAddress is a paid mutator transaction binding the contract method 0x57cfbee3.
//
// Solidity: function setMultipleSovereignTokenAddress(uint32[] originNetworks, address[] originTokenAddresses, address[] sovereignTokenAddresses, bool[] isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetMultipleSovereignTokenAddress(originNetworks []uint32, originTokenAddresses []common.Address, sovereignTokenAddresses []common.Address, isNotMintable []bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetMultipleSovereignTokenAddress(&_Bridgel2sovereignchain.TransactOpts, originNetworks, originTokenAddresses, sovereignTokenAddresses, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetSovereignWETHAddress(opts *bind.TransactOpts, sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setSovereignWETHAddress", sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetSovereignWETHAddress(&_Bridgel2sovereignchain.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// SetSovereignWETHAddress is a paid mutator transaction binding the contract method 0xbf130d7f.
//
// Solidity: function setSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetSovereignWETHAddress(sovereignWETHTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetSovereignWETHAddress(&_Bridgel2sovereignchain.TransactOpts, sovereignWETHTokenAddress, isNotMintable)
}

// UnsetMultipleClaimedBitmap is a paid mutator transaction binding the contract method 0x8781a5c5.
//
// Solidity: function unsetMultipleClaimedBitmap(uint32[] leafIndexes, uint32[] sourceBridgeNetworks) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) UnsetMultipleClaimedBitmap(opts *bind.TransactOpts, leafIndexes []uint32, sourceBridgeNetworks []uint32) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "unsetMultipleClaimedBitmap", leafIndexes, sourceBridgeNetworks)
}

// UnsetMultipleClaimedBitmap is a paid mutator transaction binding the contract method 0x8781a5c5.
//
// Solidity: function unsetMultipleClaimedBitmap(uint32[] leafIndexes, uint32[] sourceBridgeNetworks) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) UnsetMultipleClaimedBitmap(leafIndexes []uint32, sourceBridgeNetworks []uint32) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UnsetMultipleClaimedBitmap(&_Bridgel2sovereignchain.TransactOpts, leafIndexes, sourceBridgeNetworks)
}

// UnsetMultipleClaimedBitmap is a paid mutator transaction binding the contract method 0x8781a5c5.
//
// Solidity: function unsetMultipleClaimedBitmap(uint32[] leafIndexes, uint32[] sourceBridgeNetworks) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) UnsetMultipleClaimedBitmap(leafIndexes []uint32, sourceBridgeNetworks []uint32) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UnsetMultipleClaimedBitmap(&_Bridgel2sovereignchain.TransactOpts, leafIndexes, sourceBridgeNetworks)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UpdateGlobalExitRoot(&_Bridgel2sovereignchain.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UpdateGlobalExitRoot(&_Bridgel2sovereignchain.TransactOpts)
}

// Bridgel2sovereignchainBridgeEventIterator is returned from FilterBridgeEvent and is used to iterate over the raw logs and unpacked data for BridgeEvent events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainBridgeEventIterator struct {
	Event *Bridgel2sovereignchainBridgeEvent // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainBridgeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainBridgeEvent)
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
		it.Event = new(Bridgel2sovereignchainBridgeEvent)
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
func (it *Bridgel2sovereignchainBridgeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainBridgeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainBridgeEvent represents a BridgeEvent event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainBridgeEvent struct {
	LeafType           uint8
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationNetwork uint32
	DestinationAddress common.Address
	Amount             *big.Int
	Metadata           []byte
	DepositCount       uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBridgeEvent is a free log retrieval operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterBridgeEvent(opts *bind.FilterOpts) (*Bridgel2sovereignchainBridgeEventIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainBridgeEventIterator{contract: _Bridgel2sovereignchain.contract, event: "BridgeEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeEvent is a free log subscription operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchBridgeEvent(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainBridgeEvent) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainBridgeEvent)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
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

// ParseBridgeEvent is a log parse operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseBridgeEvent(log types.Log) (*Bridgel2sovereignchainBridgeEvent, error) {
	event := new(Bridgel2sovereignchainBridgeEvent)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainClaimEventIterator struct {
	Event *Bridgel2sovereignchainClaimEvent // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainClaimEvent)
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
		it.Event = new(Bridgel2sovereignchainClaimEvent)
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
func (it *Bridgel2sovereignchainClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainClaimEvent represents a ClaimEvent event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainClaimEvent struct {
	GlobalIndex        *big.Int
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationAddress common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterClaimEvent is a free log retrieval operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterClaimEvent(opts *bind.FilterOpts) (*Bridgel2sovereignchainClaimEventIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainClaimEventIterator{contract: _Bridgel2sovereignchain.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainClaimEvent) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainClaimEvent)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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

// ParseClaimEvent is a log parse operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseClaimEvent(log types.Log) (*Bridgel2sovereignchainClaimEvent, error) {
	event := new(Bridgel2sovereignchainClaimEvent)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainEmergencyStateActivatedIterator struct {
	Event *Bridgel2sovereignchainEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainEmergencyStateActivated)
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
		it.Event = new(Bridgel2sovereignchainEmergencyStateActivated)
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
func (it *Bridgel2sovereignchainEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainEmergencyStateActivated represents a EmergencyStateActivated event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*Bridgel2sovereignchainEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainEmergencyStateActivatedIterator{contract: _Bridgel2sovereignchain.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainEmergencyStateActivated)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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

// ParseEmergencyStateActivated is a log parse operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseEmergencyStateActivated(log types.Log) (*Bridgel2sovereignchainEmergencyStateActivated, error) {
	event := new(Bridgel2sovereignchainEmergencyStateActivated)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainEmergencyStateDeactivatedIterator struct {
	Event *Bridgel2sovereignchainEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainEmergencyStateDeactivated)
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
		it.Event = new(Bridgel2sovereignchainEmergencyStateDeactivated)
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
func (it *Bridgel2sovereignchainEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*Bridgel2sovereignchainEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainEmergencyStateDeactivatedIterator{contract: _Bridgel2sovereignchain.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainEmergencyStateDeactivated)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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

// ParseEmergencyStateDeactivated is a log parse operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseEmergencyStateDeactivated(log types.Log) (*Bridgel2sovereignchainEmergencyStateDeactivated, error) {
	event := new(Bridgel2sovereignchainEmergencyStateDeactivated)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainInitializedIterator struct {
	Event *Bridgel2sovereignchainInitialized // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainInitialized)
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
		it.Event = new(Bridgel2sovereignchainInitialized)
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
func (it *Bridgel2sovereignchainInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainInitialized represents a Initialized event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterInitialized(opts *bind.FilterOpts) (*Bridgel2sovereignchainInitializedIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainInitializedIterator{contract: _Bridgel2sovereignchain.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainInitialized)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseInitialized(log types.Log) (*Bridgel2sovereignchainInitialized, error) {
	event := new(Bridgel2sovereignchainInitialized)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainMigrateLegacyTokenIterator is returned from FilterMigrateLegacyToken and is used to iterate over the raw logs and unpacked data for MigrateLegacyToken events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainMigrateLegacyTokenIterator struct {
	Event *Bridgel2sovereignchainMigrateLegacyToken // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainMigrateLegacyTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainMigrateLegacyToken)
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
		it.Event = new(Bridgel2sovereignchainMigrateLegacyToken)
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
func (it *Bridgel2sovereignchainMigrateLegacyTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainMigrateLegacyTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainMigrateLegacyToken represents a MigrateLegacyToken event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainMigrateLegacyToken struct {
	Sender              common.Address
	LegacyTokenAddress  common.Address
	UpdatedTokenAddress common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMigrateLegacyToken is a free log retrieval operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterMigrateLegacyToken(opts *bind.FilterOpts) (*Bridgel2sovereignchainMigrateLegacyTokenIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainMigrateLegacyTokenIterator{contract: _Bridgel2sovereignchain.contract, event: "MigrateLegacyToken", logs: logs, sub: sub}, nil
}

// WatchMigrateLegacyToken is a free log subscription operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchMigrateLegacyToken(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainMigrateLegacyToken) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "MigrateLegacyToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainMigrateLegacyToken)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
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

// ParseMigrateLegacyToken is a log parse operation binding the contract event 0xb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9.
//
// Solidity: event MigrateLegacyToken(address sender, address legacyTokenAddress, address updatedTokenAddress, uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseMigrateLegacyToken(log types.Log) (*Bridgel2sovereignchainMigrateLegacyToken, error) {
	event := new(Bridgel2sovereignchainMigrateLegacyToken)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "MigrateLegacyToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainNewWrappedTokenIterator is returned from FilterNewWrappedToken and is used to iterate over the raw logs and unpacked data for NewWrappedToken events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainNewWrappedTokenIterator struct {
	Event *Bridgel2sovereignchainNewWrappedToken // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainNewWrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainNewWrappedToken)
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
		it.Event = new(Bridgel2sovereignchainNewWrappedToken)
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
func (it *Bridgel2sovereignchainNewWrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainNewWrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainNewWrappedToken represents a NewWrappedToken event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainNewWrappedToken struct {
	OriginNetwork       uint32
	OriginTokenAddress  common.Address
	WrappedTokenAddress common.Address
	Metadata            []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWrappedToken is a free log retrieval operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterNewWrappedToken(opts *bind.FilterOpts) (*Bridgel2sovereignchainNewWrappedTokenIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainNewWrappedTokenIterator{contract: _Bridgel2sovereignchain.contract, event: "NewWrappedToken", logs: logs, sub: sub}, nil
}

// WatchNewWrappedToken is a free log subscription operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchNewWrappedToken(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainNewWrappedToken) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainNewWrappedToken)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
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

// ParseNewWrappedToken is a log parse operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseNewWrappedToken(log types.Log) (*Bridgel2sovereignchainNewWrappedToken, error) {
	event := new(Bridgel2sovereignchainNewWrappedToken)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator is returned from FilterRemoveLegacySovereignTokenAddress and is used to iterate over the raw logs and unpacked data for RemoveLegacySovereignTokenAddress events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator struct {
	Event *Bridgel2sovereignchainRemoveLegacySovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainRemoveLegacySovereignTokenAddress)
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
		it.Event = new(Bridgel2sovereignchainRemoveLegacySovereignTokenAddress)
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
func (it *Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainRemoveLegacySovereignTokenAddress represents a RemoveLegacySovereignTokenAddress event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainRemoveLegacySovereignTokenAddress struct {
	SovereignTokenAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRemoveLegacySovereignTokenAddress is a free log retrieval operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterRemoveLegacySovereignTokenAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainRemoveLegacySovereignTokenAddressIterator{contract: _Bridgel2sovereignchain.contract, event: "RemoveLegacySovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchRemoveLegacySovereignTokenAddress is a free log subscription operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchRemoveLegacySovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainRemoveLegacySovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "RemoveLegacySovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainRemoveLegacySovereignTokenAddress)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
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

// ParseRemoveLegacySovereignTokenAddress is a log parse operation binding the contract event 0xc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063.
//
// Solidity: event RemoveLegacySovereignTokenAddress(address sovereignTokenAddress)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseRemoveLegacySovereignTokenAddress(log types.Log) (*Bridgel2sovereignchainRemoveLegacySovereignTokenAddress, error) {
	event := new(Bridgel2sovereignchainRemoveLegacySovereignTokenAddress)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "RemoveLegacySovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainSetBridgeManagerIterator is returned from FilterSetBridgeManager and is used to iterate over the raw logs and unpacked data for SetBridgeManager events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetBridgeManagerIterator struct {
	Event *Bridgel2sovereignchainSetBridgeManager // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetBridgeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetBridgeManager)
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
		it.Event = new(Bridgel2sovereignchainSetBridgeManager)
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
func (it *Bridgel2sovereignchainSetBridgeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetBridgeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetBridgeManager represents a SetBridgeManager event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetBridgeManager struct {
	BridgeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetBridgeManager is a free log retrieval operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetBridgeManager(opts *bind.FilterOpts) (*Bridgel2sovereignchainSetBridgeManagerIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetBridgeManagerIterator{contract: _Bridgel2sovereignchain.contract, event: "SetBridgeManager", logs: logs, sub: sub}, nil
}

// WatchSetBridgeManager is a free log subscription operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetBridgeManager(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetBridgeManager) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetBridgeManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetBridgeManager)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
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

// ParseSetBridgeManager is a log parse operation binding the contract event 0x32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd67.
//
// Solidity: event SetBridgeManager(address bridgeManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetBridgeManager(log types.Log) (*Bridgel2sovereignchainSetBridgeManager, error) {
	event := new(Bridgel2sovereignchainSetBridgeManager)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetBridgeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainSetSovereignTokenAddressIterator is returned from FilterSetSovereignTokenAddress and is used to iterate over the raw logs and unpacked data for SetSovereignTokenAddress events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetSovereignTokenAddressIterator struct {
	Event *Bridgel2sovereignchainSetSovereignTokenAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetSovereignTokenAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetSovereignTokenAddress)
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
		it.Event = new(Bridgel2sovereignchainSetSovereignTokenAddress)
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
func (it *Bridgel2sovereignchainSetSovereignTokenAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetSovereignTokenAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetSovereignTokenAddress represents a SetSovereignTokenAddress event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetSovereignTokenAddress struct {
	OriginNetwork         uint32
	OriginTokenAddress    common.Address
	SovereignTokenAddress common.Address
	IsNotMintable         bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignTokenAddress is a free log retrieval operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetSovereignTokenAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainSetSovereignTokenAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetSovereignTokenAddressIterator{contract: _Bridgel2sovereignchain.contract, event: "SetSovereignTokenAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignTokenAddress is a free log subscription operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetSovereignTokenAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetSovereignTokenAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetSovereignTokenAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetSovereignTokenAddress)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
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

// ParseSetSovereignTokenAddress is a log parse operation binding the contract event 0xdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce7.
//
// Solidity: event SetSovereignTokenAddress(uint32 originNetwork, address originTokenAddress, address sovereignTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetSovereignTokenAddress(log types.Log) (*Bridgel2sovereignchainSetSovereignTokenAddress, error) {
	event := new(Bridgel2sovereignchainSetSovereignTokenAddress)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetSovereignTokenAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainSetSovereignWETHAddressIterator is returned from FilterSetSovereignWETHAddress and is used to iterate over the raw logs and unpacked data for SetSovereignWETHAddress events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetSovereignWETHAddressIterator struct {
	Event *Bridgel2sovereignchainSetSovereignWETHAddress // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetSovereignWETHAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetSovereignWETHAddress)
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
		it.Event = new(Bridgel2sovereignchainSetSovereignWETHAddress)
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
func (it *Bridgel2sovereignchainSetSovereignWETHAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetSovereignWETHAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetSovereignWETHAddress represents a SetSovereignWETHAddress event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetSovereignWETHAddress struct {
	SovereignWETHTokenAddress common.Address
	IsNotMintable             bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetSovereignWETHAddress is a free log retrieval operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetSovereignWETHAddress(opts *bind.FilterOpts) (*Bridgel2sovereignchainSetSovereignWETHAddressIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetSovereignWETHAddressIterator{contract: _Bridgel2sovereignchain.contract, event: "SetSovereignWETHAddress", logs: logs, sub: sub}, nil
}

// WatchSetSovereignWETHAddress is a free log subscription operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetSovereignWETHAddress(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetSovereignWETHAddress) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetSovereignWETHAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetSovereignWETHAddress)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
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

// ParseSetSovereignWETHAddress is a log parse operation binding the contract event 0xc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82.
//
// Solidity: event SetSovereignWETHAddress(address sovereignWETHTokenAddress, bool isNotMintable)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetSovereignWETHAddress(log types.Log) (*Bridgel2sovereignchainSetSovereignWETHAddress, error) {
	event := new(Bridgel2sovereignchainSetSovereignWETHAddress)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetSovereignWETHAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainUnsetClaimIterator is returned from FilterUnsetClaim and is used to iterate over the raw logs and unpacked data for UnsetClaim events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUnsetClaimIterator struct {
	Event *Bridgel2sovereignchainUnsetClaim // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainUnsetClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainUnsetClaim)
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
		it.Event = new(Bridgel2sovereignchainUnsetClaim)
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
func (it *Bridgel2sovereignchainUnsetClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainUnsetClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainUnsetClaim represents a UnsetClaim event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUnsetClaim struct {
	LeafIndex           uint32
	SourceBridgeNetwork uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUnsetClaim is a free log retrieval operation binding the contract event 0x4a402ac607e71571d0543be8fcccc358a4df62dcd019a1e7f7e98ca6175e8f2a.
//
// Solidity: event UnsetClaim(uint32 leafIndex, uint32 sourceBridgeNetwork)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterUnsetClaim(opts *bind.FilterOpts) (*Bridgel2sovereignchainUnsetClaimIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "UnsetClaim")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainUnsetClaimIterator{contract: _Bridgel2sovereignchain.contract, event: "UnsetClaim", logs: logs, sub: sub}, nil
}

// WatchUnsetClaim is a free log subscription operation binding the contract event 0x4a402ac607e71571d0543be8fcccc358a4df62dcd019a1e7f7e98ca6175e8f2a.
//
// Solidity: event UnsetClaim(uint32 leafIndex, uint32 sourceBridgeNetwork)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchUnsetClaim(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainUnsetClaim) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "UnsetClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainUnsetClaim)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UnsetClaim", log); err != nil {
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

// ParseUnsetClaim is a log parse operation binding the contract event 0x4a402ac607e71571d0543be8fcccc358a4df62dcd019a1e7f7e98ca6175e8f2a.
//
// Solidity: event UnsetClaim(uint32 leafIndex, uint32 sourceBridgeNetwork)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseUnsetClaim(log types.Log) (*Bridgel2sovereignchainUnsetClaim, error) {
	event := new(Bridgel2sovereignchainUnsetClaim)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UnsetClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
