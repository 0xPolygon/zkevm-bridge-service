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

// AgglayerBridgeL2LeafData is an auto generated low-level Go binding around an user-defined struct.
type AgglayerBridgeL2LeafData struct {
	LeafType           uint8
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationNetwork uint32
	DestinationAddress common.Address
	Amount             *big.Int
	Metadata           []byte
}

// Bridgel2sovereignchainMetaData contains all meta data concerning the Bridgel2sovereignchain contract.
var Bridgel2sovereignchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyStateNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedProxyDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDepositCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExpectedLER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGlobalIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLBTLeaf\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeafType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLeavesLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidProxyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSovereignWETHAddressParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubtreeFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroNetworkID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidZeroProxyAdminOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeOverflow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"localBalanceTreeAmount\",\"type\":\"uint256\"}],\"name\":\"LocalBalanceTreeUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewDepositCountExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroValueForUnusedFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBridgeManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyDeployer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGlobalExitRootRemover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingEmergencyBridgeUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OriginNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubtreeFrontierMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAlreadyUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotMapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotRemapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WETHRemappingNotSupportedOnGasTokenNetworks\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"AcceptEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"AcceptProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"BackwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":false,\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"DetailedClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newLeaves\",\"type\":\"bytes\"}],\"name\":\"ForwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"updatedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MigrateLegacyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"RemoveLegacySovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgeManager\",\"type\":\"address\"}],\"name\":\"SetBridgeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceNetwork\",\"type\":\"uint32\"}],\"name\":\"SetClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"SetLocalBalanceTree\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignTokenAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"SetSovereignWETHAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgePauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgePauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentEmergencyBridgeUnpauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"TransferEmergencyBridgeUnpauserRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"TransferProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"claimedGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newClaimedGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedClaimedGlobalIndexHashChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unsetGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newUnsetGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedUnsetGlobalIndexHashChain\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_SOVEREIGN_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractITokenWrappedBridgeUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[32]\",\"name\":\"newFrontier\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32\",\"name\":\"nextLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"proof\",\"type\":\"bytes32[32]\"}],\"name\":\"backwardLET\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeLib\",\"outputs\":[{\"internalType\":\"contractBridgeLib\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"computeTokenProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"deployWrappedTokenAndRemap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structAgglayerBridgeL2.LeafData[]\",\"name\":\"newLeaves\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"expectedLER\",\"type\":\"bytes32\"}],\"name\":\"forwardLET\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadataHash\",\"type\":\"bytes32\"}],\"name\":\"getLeafValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBridgeImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sovereignWETHAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_sovereignWETHAddressIsNotMintable\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgePauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emergencyBridgeUnpauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxiedTokensManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenInfoHash\",\"type\":\"bytes32\"}],\"name\":\"localBalanceTree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"migrateLegacyToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgePauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingEmergencyBridgeUnpauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"legacySovereignTokenAddress\",\"type\":\"address\"}],\"name\":\"removeLegacySovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeManager\",\"type\":\"address\"}],\"name\":\"setBridgeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetwork\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddress\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"name\":\"setLocalBalanceTree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32[]\",\"name\":\"originNetworks\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"originTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"sovereignTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isNotMintable\",\"type\":\"bool[]\"}],\"name\":\"setMultipleSovereignTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sovereignWETHTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"name\":\"setSovereignWETHAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyBridgePauser\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgePauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyBridgeUnpauser\",\"type\":\"address\"}],\"name\":\"transferEmergencyBridgeUnpauserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"transferProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsetGlobalIndexHashChain\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"unsetMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leafHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedAddress\",\"type\":\"address\"}],\"name\":\"wrappedAddressIsNotMintable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNotMintable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedTokenBytecodeStorer\",\"outputs\":[{\"internalType\":\"contractIBytecodeStorer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610100604052348015610010575f5ffd5b5060405161001d9061017d565b604051809103905ff080158015610036573d5f5f3e3d5ffd5b506001600160a01b031660805260405161004f9061018a565b604051809103905ff080158015610068573d5f5f3e3d5ffd5b506001600160a01b031660c05260405161008190610197565b604051809103905ff08015801561009a573d5f5f3e3d5ffd5b506001600160a01b031660a0526100af6100c0565b3360e0526100bb6100c0565b6101a4565b5f54610100900460ff161561012b5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101561017b575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610fcf806161bd83390190565b6115988061718c83390190565b610cde8061872483390190565b60805160a05160c05160e051615fca6101f35f395f610c8c01525f6105ce01525f81816107bb015281816123840152818161288001526139dc01525f818161058101526123fe0152615fca5ff3fe608060405260043610610366575f3560e01c806383f24403116101c3578063cc461632116100f5578063ece93c6f11610099578063ece93c6f14610b25578063ee25560b14610b44578063f0e7080814610b6f578063f214e16114610b8e578063f5efcd7914610bad578063f67566e414610707578063f811bff714610bcc578063fb57083414610beb578063fd7640e814610c0a575f5ffd5b8063cc46163214610a3c578063ccaa2d1114610a5b578063cd58657914610a7a578063d02103ca14610a8d578063d9cb3aec14610ab3578063dbc1697614610ade578063e88f043614610af2578063eabd372a14610b06575f5ffd5b8063b0b3792011610167578063b0b379201461091b578063b45869621461093a578063b8b284d014610959578063bab161bf14610978578063be5831c714610999578063bf130d7f146109bc578063c00f14ab146109db578063c0f49163146109fa578063c514f24e14610a28575f5ffd5b806383f24403146108395780638b37b873146108585780638bd309c31461086c5780638c668f1c1461088b5780638d9420961461089f5780638ed7e3f2146108be5780638f9720bb146108dd578063ae24490a146108fc575f5ffd5b80633c351e101161029c57806365d6f6541161024057806365d6f6541461070757806369e3ab12146107385780636e4ecfed146107575780636e974cd4146107765780636ee84b23146107955780636f0bc3da146107aa57806379e2cf97146107dd5780638129fc1c146107f157806381b1c17414610805575f5ffd5b80633c351e10146105f25780633cbc795b146106115780633e197043146106495780634b2f336d1461066857806354fd4d501461068757806357cfbee3146106b55780635ca1e165146106d4578063606617ff146106e8575f5ffd5b806322e95f2c1161030e57806322e95f2c14610482578063240ff378146104a157806327aef4e8146104b45780632dfdf0b5146104d55780632f84c690146104ea578063318aee3d14610509578063381fef6d1461057057806338b8fbbb146105a35780633b2fee9a146105c0575f5ffd5b80626ee1711461036a57806303e6e1161461038b578063136a2c60146103c557806314cc01a0146103e457806315064c96146104035780631c2082291461042c5780631d081d8c1461044b5780632072f6c51461046e575b5f5ffd5b348015610375575f5ffd5b506103896103843660046149fe565b610c29565b005b348015610396575f5ffd5b5060a8546103af9061010090046001600160a01b031681565b6040516103bc9190614aee565b60405180910390f35b3480156103d0575f5ffd5b506103896103df366004614b89565b611105565b3480156103ef575f5ffd5b5060a3546103af906001600160a01b031681565b34801561040e575f5ffd5b5060685461041c9060ff1681565b60405190151581526020016103bc565b348015610437575f5ffd5b50610389610446366004614b89565b61123e565b348015610456575f5ffd5b5061046060a55481565b6040519081526020016103bc565b348015610479575f5ffd5b50610389611364565b34801561048d575f5ffd5b506103af61049c366004614bc2565b611399565b6103896104af366004614c3b565b6113e7565b3480156104bf575f5ffd5b506104c8611457565b6040516103bc9190614cfc565b3480156104e0575f5ffd5b5061046060535481565b3480156104f5575f5ffd5b5060a4546103af906001600160a01b031681565b348015610514575f5ffd5b5061054c610523366004614d0e565b606b6020525f908152604090205463ffffffff811690600160201b90046001600160a01b031682565b6040805163ffffffff90931683526001600160a01b039091166020830152016103bc565b34801561057b575f5ffd5b506103af7f000000000000000000000000000000000000000000000000000000000000000081565b3480156105ae575f5ffd5b506070546001600160a01b03166103af565b3480156105cb575f5ffd5b507f00000000000000000000000000000000000000000000000000000000000000006103af565b3480156105fd575f5ffd5b50606d546103af906001600160a01b031681565b34801561061c575f5ffd5b50606d5461063490600160a01b900463ffffffff1681565b60405163ffffffff90911681526020016103bc565b348015610654575f5ffd5b50610460610663366004614d42565b6114e3565b348015610673575f5ffd5b50606f546103af906001600160a01b031681565b348015610692575f5ffd5b50604080518082019091526006815265076312e312e360d41b60208201526104c8565b3480156106c0575f5ffd5b506103896106cf366004614e82565b61156f565b3480156106df575f5ffd5b5061046061165b565b3480156106f3575f5ffd5b5060aa546103af906001600160a01b031681565b348015610712575f5ffd5b506104c860405180604001604052806006815260200165076312e312e360d41b81525081565b348015610743575f5ffd5b50610389610752366004614d0e565b6116da565b348015610762575f5ffd5b506070546103af906001600160a01b031681565b348015610781575f5ffd5b50610389610790366004614f97565b61176c565b3480156107a0575f5ffd5b5061046060a65481565b3480156107b5575f5ffd5b506103af7f000000000000000000000000000000000000000000000000000000000000000081565b3480156107e8575f5ffd5b50610389611ac0565b3480156107fc575f5ffd5b50610389611ae1565b348015610810575f5ffd5b506103af61081f366004614fdd565b606a6020525f90815260409020546001600160a01b031681565b348015610844575f5ffd5b50610460610853366004615005565b611afa565b348015610863575f5ffd5b50610389611b89565b348015610877575f5ffd5b50610389610886366004614d0e565b611bff565b348015610896575f5ffd5b50610389611c7f565b3480156108aa575f5ffd5b506103896108b9366004614d0e565b611d08565b3480156108c9575f5ffd5b50606c546103af906001600160a01b031681565b3480156108e8575f5ffd5b506103896108f7366004615041565b611d88565b348015610907575f5ffd5b5060a9546103af906001600160a01b031681565b348015610926575f5ffd5b506103896109353660046150cc565b612002565b348015610945575f5ffd5b50610389610954366004614d0e565b612156565b348015610964575f5ffd5b50610389610973366004615123565b6122b7565b348015610983575f5ffd5b5060685461063490610100900463ffffffff1681565b3480156109a4575f5ffd5b5060685461063490600160c81b900463ffffffff1681565b3480156109c7575f5ffd5b506103896109d63660046151a0565b612335565b3480156109e6575f5ffd5b506104c86109f5366004614d0e565b61236a565b348015610a05575f5ffd5b5061041c610a14366004614d0e565b60a26020525f908152604090205460ff1681565b348015610a33575f5ffd5b506104c86123fa565b348015610a47575f5ffd5b5061041c610a563660046151cc565b612483565b348015610a66575f5ffd5b50610389610a753660046151fd565b6124d3565b610389610a883660046152db565b612549565b348015610a98575f5ffd5b506068546103af90600160281b90046001600160a01b031681565b348015610abe575f5ffd5b50610460610acd366004614fdd565b60a76020525f908152604090205481565b348015610ae9575f5ffd5b50610389612972565b348015610afd575f5ffd5b506103896129a5565b348015610b11575f5ffd5b50610389610b20366004614d0e565b612a2c565b348015610b30575f5ffd5b506071546103af906001600160a01b031681565b348015610b4f575f5ffd5b50610460610b5e366004614fdd565b60696020525f908152604090205481565b348015610b7a575f5ffd5b50610389610b8936600461536a565b612aca565b348015610b99575f5ffd5b506103af610ba8366004614bc2565b612d53565b348015610bb8575f5ffd5b50610389610bc73660046151fd565b612e19565b348015610bd7575f5ffd5b50610389610be63660046153dd565b613069565b348015610bf6575f5ffd5b5061041c610c0536600461546c565b6130f5565b348015610c15575f5ffd5b50610389610c243660046154b1565b61310c565b5f54600390610100900460ff16158015610c4957505f5460ff8083169116105b610c6e5760405162461bcd60e51b8152600401610c65906154f9565b60405180910390fd5b5f805461ffff191660ff831617610100179055336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cca5760405163618bbdd560e01b815260040160405180910390fd5b6001600160a01b038a16610cf15760405163f6b2911f60e01b815260040160405180910390fd5b8c63ffffffff165f03610d1757604051634e702fa560e01b815260040160405180910390fd5b8c606860016101000a81548163ffffffff021916908363ffffffff16021790555089606860056101000a8154816001600160a01b0302191690836001600160a01b0316021790555088606c5f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508660a35f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508360a45f6101000a8154816001600160a01b0302191690836001600160a01b031602179055505f516020615f755f395f51905f525f60a45f9054906101000a90046001600160a01b0316604051610e03929190615547565b60405180910390a160a980546001600160a01b0319166001600160a01b0385169081179091556040517f24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d91610e5a915f9190615547565b60405180910390a1306001600160a01b03831603610e8b57604051631ae0e03360e01b815260040160405180910390fd5b6001600160a01b038216610eb25760405163f6b2911f60e01b815260040160405180910390fd5b607080546001600160a01b0319166001600160a01b0384169081179091556040517fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f291610f01915f9190615547565b60405180910390a16001600160a01b038c16610f735763ffffffff8b1615610f3c57604051630d43a60960e11b815260040160405180910390fd5b6001600160a01b038616151580610f505750845b15610f6e57604051630e6e237560e11b815260040160405180910390fd5b6110ae565b606d805463ffffffff8d16600160a01b026001600160c01b03199091166001600160a01b038f1617179055606e610faa89826155dd565b506001600160a01b03861661107657841515600103610fdc57604051630e6e237560e11b815260040160405180910390fd5b6110515f5f1b601260405160200161103d91906060808252600d908201526c2bb930b83832b21022ba3432b960991b608082015260a060208201819052600490820152630ae8aa8960e31b60c082015260ff91909116604082015260e00190565b604051602081830303815290604052613323565b606f80546001600160a01b0319166001600160a01b03929092169190911790556110ae565b606f80546001600160a01b0319166001600160a01b0388169081179091555f90815260a260205260409020805460ff19168615151790555b6110b6613404565b5f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050505050505050565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa158015611153573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111779190615697565b6001600160a01b03161461119e5760405163a34ddeb160e01b815260040160405180910390fd5b5f5b815181101561123a575f8282815181106111bc576111bc6156b2565b602002602001015190505f5f6111d183613432565b92505091506111e082826134d6565b60a6545f90815260208490526040902060a68190556040805185815260208101929092527fc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03910160405180910390a15050506001016111a0565b5050565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa15801561128c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112b09190615697565b6001600160a01b0316146112d75760405163a34ddeb160e01b815260040160405180910390fd5b5f5b815181101561123a575f8282815181106112f5576112f56156b2565b602002602001015190505f5f61130a83613432565b92505091506113198282613547565b6040805163ffffffff8085168252831660208201527fab48452e173be3f004b3208cc344b14cc7f8d977fc1ca28171286a7abb5d7570910160405180910390a15050506001016112d9565b60a4546001600160a01b0316331461138f57604051631344c5df60e11b815260040160405180910390fd5b6113976135b9565b565b5f606a5f84846040516020016113b09291906156c6565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b031690505b92915050565b60685460ff161561140b57604051630bc011ff60e21b815260040160405180910390fd5b34158015906114245750606f546001600160a01b031615155b15611442576040516301bd897160e61b815260040160405180910390fd5b611450858534868686613614565b5050505050565b606e805461146490615561565b80601f016020809104026020016040519081016040528092919081815260200182805461149090615561565b80156114db5780601f106114b2576101008083540402835291602001916114db565b820191905f5260205f20905b8154815290600101906020018083116114be57829003601f168201915b505050505081565b6040516001600160f81b031960f889901b1660208201526001600160e01b031960e088811b821660218401526001600160601b0319606089811b821660258601529188901b909216603984015285901b16603d82015260518101839052607181018290525f90609101604051602081830303815290604052805190602001209050979650505050505050565b60a3546001600160a01b0316331461159a576040516357b738d160e11b815260040160405180910390fd5b825184511415806115ad57508151845114155b806115ba57508051845114155b156115d85760405163434f49f560e11b815260040160405180910390fd5b5f5b8251811015611450576116538582815181106115f8576115f86156b2565b6020026020010151858381518110611612576116126156b2565b602002602001015185848151811061162c5761162c6156b2565b6020026020010151858581518110611646576116466156b2565b60200260200101516136e8565b6001016115da565b6053545f90819081805b60208110156116d1578083901c6001166001036116aa576116a360338260208110611692576116926156b2565b0154855f9182526020526040902090565b93506116ba565b5f84815260208390526040902093505b5f8281526020839052604090209150600101611665565b50919392505050565b60a4546001600160a01b0316331461170557604051631344c5df60e11b815260040160405180910390fd5b60a880546001600160a01b0380841661010002610100600160a81b03199092169190911790915560a4546040517fb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a926117619216908490615547565b60405180910390a150565b60a3546001600160a01b03163314611797576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b0382161580156117b2575063ffffffff8316155b15611942575f6119305f5f1b606f5f9054906101000a90046001600160a01b03166001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa15801561180d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611834919081019061573c565b606f5f9054906101000a90046001600160a01b03166001600160a01b03166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611883573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526118aa919081019061573c565b606f5f9054906101000a90046001600160a01b03166001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118fa573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061191e919061576d565b60405160200161103d93929190615788565b905061193c8183613898565b50505050565b5f83836040516020016119569291906156c6565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b0316806119a45760405163828d566360e01b815260040160405180910390fd5b5f611aa983836001600160a01b03166306fdde036040518163ffffffff1660e01b81526004015f60405180830381865afa1580156119e4573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611a0b919081019061573c565b846001600160a01b03166395d89b416040518163ffffffff1660e01b81526004015f60405180830381865afa158015611a46573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611a6d919081019061573c565b856001600160a01b031663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118fa573d5f5f3e3d5ffd5b9050611ab7868683876136e8565b5050505b505050565b605354606854600160c81b900463ffffffff16101561139757611397613937565b60405163f57ac68360e01b815260040160405180910390fd5b5f83815b6020811015611b7e57600163ffffffff8516821c81169003611b4a57611b43858260208110611b2f57611b2f6156b2565b6020020135835f9182526020526040902090565b9150611b76565b611b7382868360208110611b6057611b606156b2565b60200201355f9182526020526040902090565b91505b600101611afe565b5090505b9392505050565b60aa546001600160a01b03163314611bb45760405163d491f0c160e01b815260040160405180910390fd5b60a9805460aa80546001600160a01b038082166001600160a01b03198086168217909655949091169091556040519116915f516020615f755f395f51905f5291611761918491615547565b6070546001600160a01b03163314611c2a57604051630866750360e01b815260040160405180910390fd5b607180546001600160a01b0319166001600160a01b03838116919091179091556070546040517f0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2926117619216908490615547565b6071546001600160a01b03163314611caa57604051630b59ef2760e21b815260040160405180910390fd5b60708054607180546001600160a01b038082166001600160a01b03198086168217909655949091169091556040519116917fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f291611761918491615547565b60a9546001600160a01b03163314611d3357604051638e9d821f60e01b815260040160405180910390fd5b60aa80546001600160a01b0319166001600160a01b038381169190911790915560a9546040517ff01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b926117619216908490615547565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa158015611dd6573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611dfa9190615697565b6001600160a01b031614611e215760405163a34ddeb160e01b815260040160405180910390fd5b60685460ff16611e4457604051635386698160e01b815260040160405180910390fd5b81518351141580611e5757508051835114155b15611e755760405163434f49f560e11b815260040160405180910390fd5b5f5b835181101561193c57606860019054906101000a900463ffffffff1663ffffffff16848281518110611eab57611eab6156b2565b602002602001015163ffffffff1603611ed75760405163b869a63f60e01b815260040160405180910390fd5b5f848281518110611eea57611eea6156b2565b6020026020010151848381518110611f0457611f046156b2565b6020026020010151604051602001611f1d9291906156c6565b604051602081830303815290604052805190602001209050828281518110611f4757611f476156b2565b602002602001015160a75f8381526020019081526020015f2081905550838281518110611f7657611f766156b2565b60200260200101516001600160a01b0316858381518110611f9957611f996156b2565b602002602001015163ffffffff167f73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2858581518110611fda57611fda6156b2565b6020026020010151604051611ff191815260200190565b60405180910390a350600101611e77565b8015612013576120138483836139c5565b6001600160a01b038085165f908152606b602090815260409182902082518084019093525463ffffffff81168352600160201b90049092169181018290529061206f5760405163828d566360e01b815260040160405180910390fd5b5f606a5f835f0151846020015160405160200161208d9291906156c6565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b039081169150861681036120e05760405163e273c4a160e01b815260040160405180910390fd5b5f6120eb8787613a59565b90506120f8823383613be2565b604080513381526001600160a01b03808a166020830152841691810191909152606081018290527fb7f8fd4d1faf9b2929dc269f59c53e3a2bccc44e9950f33a568fcbcb37eb69a9906080015b60405180910390a150505050505050565b60a3546001600160a01b03163314612181576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b038082165f908152606b6020908152604080832081518083018352905463ffffffff8116808352600160201b909104909516818401819052915190946121d193909291016156c6565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b0316158061222557505f818152606a60205260409020546001600160a01b038481169116145b156122435760405163e0c897a760e01b815260040160405180910390fd5b6001600160a01b0383165f908152606b6020908152604080832080546001600160c01b031916905560a290915290819020805460ff19169055517fc2ae0bd0ec0fd0352bfe5bacac49637af342c1e40f1b80a7f74440dc7fe3f063906122aa908590614aee565b60405180910390a1505050565b60685460ff16156122db57604051630bc011ff60e21b815260040160405180910390fd5b606f546001600160a01b03166123045760405163dde3cda760e01b815260040160405180910390fd5b606f545f9061231c906001600160a01b031686613a59565b905061232c878783878787613614565b50505050505050565b60a3546001600160a01b03163314612360576040516357b738d160e11b815260040160405180910390fd5b61123a8282613898565b60405163c00f14ab60e01b81526060906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906123b9908590600401614aee565b5f60405180830381865afa1580156123d3573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526113e191908101906157c0565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c514f24e6040518163ffffffff1660e01b81526004015f60405180830381865afa158015612457573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261247e91908101906157c0565b905090565b5f80612499600160201b63ffffffff8516615818565b6124a99063ffffffff861661582f565b600881901c5f90815260696020526040902054600160ff9092169190911b90811614949350505050565b6124e78c8c8c8c8c8c8c8c8c8c8c8c613c6d565b836001600160a01b03168a7f7fee792f28e46481b8f3ff0ce0983de269108f87dd43f7c7c10f7c922897a9448e8e8d8d8d8d8d8c8c8c6040516125339a9998979695949392919061586a565b60405180910390a3505050505050505050505050565b60685460ff161561256d57604051630bc011ff60e21b815260040160405180910390fd5b61257561405f565b60685463ffffffff6101009091048116908816036125a6576040516302caf51760e11b815260040160405180910390fd5b5f806060876001600160a01b038816612689578834146125d95760405163b89240f560e01b815260040160405180910390fd5b606d54606e80546001600160a01b0383169650600160a01b90920463ffffffff1694509061260690615561565b80601f016020809104026020016040519081016040528092919081815260200182805461263290615561565b801561267d5780601f106126545761010080835404028352916020019161267d565b820191905f5260205f20905b81548152906001019060200180831161266057829003601f168201915b505050505091506128fa565b34156126a85760405163798ee6f160e01b815260040160405180910390fd5b84156126b9576126b98887876139c5565b606f546001600160a01b03908116908916036126e0576126d9888a613a59565b90506128fa565b6001600160a01b038089165f908152606b602090815260409182902082518084019093525463ffffffff81168352600160201b9004909216918101829052901515806127325750805163ffffffff1615155b1561275457612741898b613a59565b6020820151825190965094509150612869565b6040516370a0823160e01b81525f906001600160a01b038b16906370a0823190612782903090600401614aee565b602060405180830381865afa15801561279d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127c191906158e1565b90506127d86001600160a01b038b1633308e6140b8565b6040516370a0823160e01b81525f906001600160a01b038c16906370a0823190612806903090600401614aee565b602060405180830381865afa158015612821573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061284591906158e1565b905061285182826158f8565b6068548c9850610100900463ffffffff169650935050505b60405163c00f14ab60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906128b5908c90600401614aee565b5f60405180830381865afa1580156128cf573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526128f691908101906157c0565b9250505b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b5f84868e8e868860535460405161293998979695949392919061590b565b60405180910390a16129575f84868e8e86888051906020012061411f565b861561296557612965613937565b5050505061232c60018055565b60a9546001600160a01b0316331461299d57604051638e9d821f60e01b815260040160405180910390fd5b61139761415d565b60a85461010090046001600160a01b031633146129d557604051637bb0100f60e01b815260040160405180910390fd5b60a4805460a880546001600160a01b03610100820481166001600160a01b031985168117909555610100600160a81b03199091169091556040519116915f516020615f755f395f51905f5291611761918491615547565b60a3546001600160a01b03163314612a57576040516357b738d160e11b815260040160405180910390fd5b6001600160a01b038116612a7e5760405163f6b2911f60e01b815260040160405180910390fd5b60a380546001600160a01b0319166001600160a01b0383169081179091556040517f32cf74f8a6d5f88593984d2cd52be5592bfa6884f5896175801a5069ef09cd679161176191614aee565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa158015612b18573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b3c9190615697565b6001600160a01b031614612b635760405163a34ddeb160e01b815260040160405180910390fd5b60685460ff16612b8657604051635386698160e01b815260040160405180910390fd5b5f829003612ba757604051638488056760e01b815260040160405180910390fd5b6053545f612bb361165b565b90505f5b84811015612c60575f868683818110612bd257612bd26156b2565b9050602002810190612be49190615977565b612bed90615995565b805190915060ff1615801590612c085750805160ff16600114155b15612c2657604051632f162e1960e11b815260040160405180910390fd5b612c57815f015182602001518360400151846060015185608001518660a001518760c00151805190602001206141b4565b50600101612bb7565b505f612c6a61165b565b9050838114612c8c5760405163277ffcf360e01b815260040160405180910390fd5b6068546040516333d6247d60e01b815260048101839052600160281b9091046001600160a01b0316906333d6247d906024015f604051808303815f87803b158015612cd5575f5ffd5b505af1158015612ce7573d5f5f3e3d5ffd5b505050507fe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c88383605354848a8a604051602001612d25929190615a77565b60408051601f1981840301815290829052612d439594939291615b85565b60405180910390a1505050505050565b5f5f8383604051602001612d689291906156c6565b6040516020818303038152906040528051906020012090505f60ff60f81b3083612d906123fa565b604051602001612da09190615bba565b60405160208183030381529060405280519060200120604051602001612df894939291906001600160f81b031994909416845260609290921b6001600160601b03191660018401526015830152603582015260550190565b60408051808303601f19018152919052805160209091012095945050505050565b60685460ff1615612e3d57604051630bc011ff60e21b815260040160405180910390fd5b60685463ffffffff8681166101009092041614612e6d576040516302caf51760e11b815260040160405180910390fd5b612e998c8c8c8c8c60018d8d8d8d8d8d8d604051612e8c929190615bcb565b60405180910390206141cb565b7f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d8a88888787604051612ed0959493929190615bda565b60405180910390a1606f545f906001600160a01b0316612f8b57846001600160a01b031684888a8686604051602401612f0c9493929190615c0c565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b17905251612f419190615bba565b5f6040518083038185875af1925050503d805f8114612f7b576040519150601f19603f3d011682016040523d82523d5f602084013e612f80565b606091505b50508091505061303c565b606f54612fa2906001600160a01b03168686613be2565b846001600160a01b031687898585604051602401612fc39493929190615c0c565b60408051601f198184030181529181526020820180516001600160e01b0316630c035af960e11b17905251612ff89190615bba565b5f604051808303815f865af19150503d805f8114613031576040519150601f19603f3d011682016040523d82523d5f602084013e613036565b606091505b50909150505b8061305a576040516337e391c360e01b815260040160405180910390fd5b50505050505050505050505050565b5f54610100900460ff161580801561308757505f54600160ff909116105b806130a05750303b1580156130a057505f5460ff166001145b6130bc5760405162461bcd60e51b8152600401610c65906154f9565b5f805460ff191660011790558015611ae1575f805461ff00191661010017905560405163f57ac68360e01b815260040160405180910390fd5b5f81613102868686611afa565b1495945050505050565b606854604080516391eb796d60e01b815290513392600160281b90046001600160a01b0316916391eb796d9160048083019260209291908290030181865afa15801561315a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061317e9190615697565b6001600160a01b0316146131a55760405163a34ddeb160e01b815260040160405180910390fd5b60685460ff166131c857604051635386698160e01b815260040160405180910390fd5b60535484106131ea57604051631abd431560e31b815260040160405180910390fd5b6131f8828286610c0561165b565b613215576040516338105f3b60e21b815260040160405180910390fd5b61322084848361427b565b6053545f61322c61165b565b90505f5b602081101561326d5785816020811061324b5761324b6156b2565b602002013560338260208110613263576132636156b2565b0155600101613230565b5060538690555f61327c61165b565b6068546040516333d6247d60e01b815260048101839052919250600160281b90046001600160a01b0316906333d6247d906024015f604051808303815f87803b1580156132c7575f5ffd5b505af11580156132d9573d5f5f3e3d5ffd5b505060408051868152602081018690529081018a9052606081018490527f2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f14955592506080019050612145565b5f5f61332d6123fa565b60405160200161333d9190615bba565b6040516020818303038152906040529050838151602083015ff591506001600160a01b038216613380576040516331682e8d60e11b815260040160405180910390fd5b5f5f5f858060200190518101906133979190615c3c565b925092509250846001600160a01b0316631624f6c68484846040518463ffffffff1660e01b81526004016133cd93929190615788565b5f604051808303815f87803b1580156133e4575f5ffd5b505af11580156133f6573d5f5f3e3d5ffd5b505050505050505092915050565b5f54610100900460ff1661342a5760405162461bcd60e51b8152600401610c6590615ca7565b61139761438e565b805f80600160401b83161561347e57505f9050808361345b63ffffffff8516600160401b61582f565b146134795760405163071389e960e01b815260040160405180910390fd5b6134cf565b602084901c9150613490826001615cf2565b9050836134b163ffffffff851667ffffffff00000000602086901b1661582f565b146134cf5760405163071389e960e01b815260040160405180910390fd5b9193909250565b5f6134eb600160201b63ffffffff8416615818565b6134fb9063ffffffff851661582f565b600881901c5f8181526069602052604090208054600160ff851690811b91821892839055939450919291908082161561232c57604051630631b5f760e31b815260040160405180910390fd5b5f61355c600160201b63ffffffff8416615818565b61356c9063ffffffff851661582f565b600881901c5f8181526069602052604081208054600160ff861690811b9182189283905594955092939291818316900361232c57604051630c8d9eab60e31b815260040160405180910390fd5b60685460ff16156135dd57604051630bc011ff60e21b815260040160405180910390fd5b6068805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b60685463ffffffff610100909104811690871603613645576040516302caf51760e11b815260040160405180910390fd5b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b6001606860019054906101000a900463ffffffff1633898989888860535460405161369999989796959493929190615d0e565b60405180910390a16136da6001606860019054906101000a900463ffffffff163389898988886040516136cd929190615bcb565b604051809103902061411f565b8215611ab757611ab7613937565b6001600160a01b038316158061370557506001600160a01b038216155b156137235760405163f6b2911f60e01b815260040160405180910390fd5b60685463ffffffff6101009091048116908516036137545760405163658b23ad60e01b815260040160405180910390fd5b6001600160a01b038281165f908152606b6020526040902054600160201b90041615613793576040516317abdeeb60e21b815260040160405180910390fd5b5f84846040516020016137a79291906156c6565b60408051808303601f1901815282825280516020918201205f818152606a835283812080546001600160a01b0319166001600160a01b038a8116918217909255868601865263ffffffff8c81168089528c8416878a01818152848752606b89528987209a518b54915194166001600160c01b031990911617600160201b93909516929092029390931790975560a2855291859020805460ff191689151590811790915585519182529381019590955292840192909252606083015291507fdbe8a5da6a7a916d9adfda9160167a0f8a3da415ee6610e810e753853597fce79060800160405180910390a15050505050565b606d546001600160a01b03166138c157604051634cb4711360e11b815260040160405180910390fd5b606f80546001600160a01b0319166001600160a01b0384169081179091555f81815260a26020908152604091829020805460ff19168515159081179091558251938452908301527fc7318b7ed6ba4f2908a3de396d8ab49b1dadb55db5b55123247a401f29ff8d82910160405180910390a15050565b6053546068805463ffffffff909216600160c81b0263ffffffff60c81b1990921691909117908190556001600160a01b03600160281b909104166333d6247d61397e61165b565b6040518263ffffffff1660e01b815260040161399c91815260200190565b5f604051808303815f87803b1580156139b3575f5ffd5b505af115801561193c573d5f5f3e3d5ffd5b60405163a28fa4a360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a28fa4a390613a199086908690869033903090600401615d7c565b6020604051808303815f875af1158015613a35573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061193c9190615dc2565b6001600160a01b0382165f90815260a2602052604081205460ff1615613b7f576040516370a0823160e01b81525f906001600160a01b038516906370a0823190613aa7903090600401614aee565b602060405180830381865afa158015613ac2573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ae691906158e1565b9050613afd6001600160a01b0385163330866140b8565b6040516370a0823160e01b81525f906001600160a01b038616906370a0823190613b2b903090600401614aee565b602060405180830381865afa158015613b46573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b6a91906158e1565b9050613b7682826158f8565b925050506113e1565b604051632770a7eb60e21b81526001600160a01b03841690639dc29fac90613bad9033908690600401615ddd565b5f604051808303815f87803b158015613bc4575f5ffd5b505af1158015613bd6573d5f5f3e3d5ffd5b505050508190506113e1565b6001600160a01b0383165f90815260a2602052604090205460ff1615613c1657611abb6001600160a01b03841683836143b4565b6040516340c10f1960e01b81526001600160a01b038416906340c10f1990613c449085908590600401615ddd565b5f604051808303815f87803b158015613c5b575f5ffd5b505af115801561232c573d5f5f3e3d5ffd5b60685460ff1615613c9157604051630bc011ff60e21b815260040160405180910390fd5b613c9961405f565b60685463ffffffff8681166101009092041614613cc9576040516302caf51760e11b815260040160405180910390fd5b613ce78c8c8c8c8c5f8d8d8d8d8d8d8d604051612e8c929190615bcb565b7f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d8a88888787604051613d1e959493929190615bda565b60405180910390a16001600160a01b038616158015613d41575063ffffffff8716155b15613e1f57606f546001600160a01b0316613e03575f6001600160a01b03851684825b6040519080825280601f01601f191660200182016040528015613d8e576020820181803683370190505b50604051613d9c9190615bba565b5f6040518083038185875af1925050503d805f8114613dd6576040519150601f19603f3d011682016040523d82523d5f602084013e613ddb565b606091505b5050905080613dfd57604051630ce8f45160e31b815260040160405180910390fd5b50614048565b606f54613e1a906001600160a01b03168585613be2565b614048565b606d546001600160a01b038781169116148015613e4d5750606d5463ffffffff888116600160a01b90920416145b15613e64575f6001600160a01b0385168482613d64565b60685463ffffffff610100909104811690881603613e9057613e1a6001600160a01b03871685856143b4565b5f8787604051602001613ea49291906156c6565b60408051601f1981840301815291815281516020928301205f818152606a9093529120549091506001600160a01b03168061403a575f613f198386868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061332392505050565b9050613f26818888613be2565b80606a5f8581526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555060405180604001604052808b63ffffffff1681526020018a6001600160a01b0316815250606b5f836001600160a01b03166001600160a01b031681526020019081526020015f205f820151815f015f6101000a81548163ffffffff021916908363ffffffff1602179055506020820151815f0160046101000a8154816001600160a01b0302191690836001600160a01b031602179055509050507f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a398a8a83888860405161402c959493929190615df6565b60405180910390a150614045565b614045818787613be2565b50505b61405160018055565b505050505050505050505050565b6002600154036140b15760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610c65565b6002600155565b6040516001600160a01b03848116602483015283811660448301526064820183905261193c9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b0383818316178352505050506143da565b61412e878787878787876141b4565b60ff871661414157614141868684614432565b5f1960ff88160161232c5761232c5f5f84614432565b60018055565b60685460ff1661418057604051635386698160e01b815260040160405180910390fd5b6068805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b61232c6141c6888888888888886114e3565b6144e7565b5f6141db888888888888886114e3565b90506141eb8d8d8d8d8d866145a6565b60a554614213906142058d845f9182526020526040902090565b5f9182526020526040902090565b60a5819055604080518d815260208101929092527f3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92910160405180910390a160ff8816614265576142658787856146df565b5f1960ff89160161305a5761305a5f5f856146df565b825f5b811580159061428d5750602081105b1561433d57816001166001036142ee578281602081106142af576142af6156b2565b60200201358482602081106142c6576142c66156b2565b6020020135146142e9576040516379c13f8b60e01b815260040160405180910390fd5b614324565b5f848260208110614301576143016156b2565b60200201351461432457604051630323f86f60e41b815260040160405180910390fd5b8061432e81615e2e565b915050600182901c915061427e565b6020811015611450575f848260208110614359576143596156b2565b60200201351461437c57604051630323f86f60e41b815260040160405180910390fd5b8061438681615e2e565b91505061433d565b5f54610100900460ff166141575760405162461bcd60e51b8152600401610c6590615ca7565b611abb83846001600160a01b031663a9059cbb85856040516024016140ed929190615ddd565b5f6143ee6001600160a01b03841683614794565b905080515f141580156144125750808060200190518101906144109190615dc2565b155b15611abb5782604051635274afe760e01b8152600401610c659190614aee565b60685463ffffffff61010090910481169084160361444f57505050565b5f83836040516020016144639291906156c6565b60408051601f1981840301815291815281516020928301205f81815260a79093529120549091508211156144bf575f81815260a76020526040908190205490516314603c0160e01b8152610c6591869186918691600401615e46565b5f81815260a76020526040812080548492906144dc9084906158f8565b909155505050505050565b8060016144f660206002615f55565b61450091906158f8565b60535410614521576040516377ae67b360e11b815260040160405180910390fd5b5f60535f815461453090615e2e565b918290555090505f5b602081101561459d578082901c60011660010361456c578260338260208110614564576145646156b2565b015550505050565b61459360338260208110614582576145826156b2565b0154845f9182526020526040902090565b9250600101614539565b50611abb615f60565b6068545f90600160281b90046001600160a01b031663257b36326145ca86866147a1565b6040518263ffffffff1660e01b81526004016145e891815260200190565b6020604051808303815f875af1158015614604573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061462891906158e1565b9050805f0361464957604051622f6fad60e01b815260040160405180910390fd5b5f5f5f61465588613432565b925092509250600160401b88165f1461469657614674858b858a6130f5565b614691576040516338105f3b60e21b815260040160405180910390fd5b6146c9565b6146ac6146a4868c86611afa565b8a84896130f5565b6146c9576040516338105f3b60e21b815260040160405180910390fd5b6146d38382613547565b50505050505050505050565b60685463ffffffff6101009091048116908416036146fc57505050565b5f83836040516020016147109291906156c6565b60408051601f1981840301815291815281516020928301205f81815260a7909352912054909150614742905f196158f8565b821115614777575f81815260a76020526040908190205490516323d7213360e01b8152610c6591869186918691600401615e46565b5f81815260a76020526040812080548492906144dc90849061582f565b6060611b8283835f6147b3565b5f828152602082905260408120611b82565b6060814710156147d8573060405163cd78605960e01b8152600401610c659190614aee565b5f5f856001600160a01b031684866040516147f39190615bba565b5f6040518083038185875af1925050503d805f811461482d576040519150601f19603f3d011682016040523d82523d5f602084013e614832565b606091505b509150915061484286838361484c565b9695505050505050565b6060826148615761485c8261489f565b611b82565b815115801561487857506001600160a01b0384163b155b156148985783604051639996b31560e01b8152600401610c659190614aee565b5080611b82565b8051156148af5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b50565b803563ffffffff811681146148de575f5ffd5b919050565b6001600160a01b03811681146148c8575f5ffd5b80356148de816148e3565b634e487b7160e01b5f52604160045260245ffd5b60405160e081016001600160401b038111828210171561493857614938614902565b60405290565b604051601f8201601f191681016001600160401b038111828210171561496657614966614902565b604052919050565b5f6001600160401b0382111561498657614986614902565b50601f01601f191660200190565b5f82601f8301126149a3575f5ffd5b81356149b66149b18261496e565b61493e565b8181528460208386010111156149ca575f5ffd5b816020850160208301375f918101602001919091529392505050565b80151581146148c8575f5ffd5b80356148de816149e6565b5f5f5f5f5f5f5f5f5f5f5f5f6101808d8f031215614a1a575f5ffd5b614a238d6148cb565b9b50614a3160208e016148f7565b9a50614a3f60408e016148cb565b9950614a4d60608e016148f7565b9850614a5b60808e016148f7565b97506001600160401b0360a08e01351115614a74575f5ffd5b614a848e60a08f01358f01614994565b9650614a9260c08e016148f7565b9550614aa060e08e016148f7565b9450614aaf6101008e016149f3565b9350614abe6101208e016148f7565b9250614acd6101408e016148f7565b9150614adc6101608e016148f7565b90509295989b509295989b509295989b565b6001600160a01b0391909116815260200190565b5f6001600160401b03821115614b1a57614b1a614902565b5060051b60200190565b5f82601f830112614b33575f5ffd5b8135614b416149b182614b02565b8082825260208201915060208360051b860101925085831115614b62575f5ffd5b602085015b83811015614b7f578035835260209283019201614b67565b5095945050505050565b5f60208284031215614b99575f5ffd5b81356001600160401b03811115614bae575f5ffd5b614bba84828501614b24565b949350505050565b5f5f60408385031215614bd3575f5ffd5b614bdc836148cb565b91506020830135614bec816148e3565b809150509250929050565b5f5f83601f840112614c07575f5ffd5b5081356001600160401b03811115614c1d575f5ffd5b602083019150836020828501011115614c34575f5ffd5b9250929050565b5f5f5f5f5f60808688031215614c4f575f5ffd5b614c58866148cb565b94506020860135614c68816148e3565b93506040860135614c78816149e6565b925060608601356001600160401b03811115614c92575f5ffd5b614c9e88828901614bf7565b969995985093965092949392505050565b5f5b83811015614cc9578181015183820152602001614cb1565b50505f910152565b5f8151808452614ce8816020860160208601614caf565b601f01601f19169290920160200192915050565b602081525f611b826020830184614cd1565b5f60208284031215614d1e575f5ffd5b8135611b82816148e3565b60ff811681146148c8575f5ffd5b80356148de81614d29565b5f5f5f5f5f5f5f60e0888a031215614d58575f5ffd5b8735614d6381614d29565b9650614d71602089016148cb565b95506040880135614d81816148e3565b9450614d8f606089016148cb565b93506080880135614d9f816148e3565b9699959850939692959460a0840135945060c09093013592915050565b5f82601f830112614dcb575f5ffd5b8135614dd96149b182614b02565b8082825260208201915060208360051b860101925085831115614dfa575f5ffd5b602085015b83811015614b7f57614e10816148cb565b835260209283019201614dff565b5f82601f830112614e2d575f5ffd5b8135614e3b6149b182614b02565b8082825260208201915060208360051b860101925085831115614e5c575f5ffd5b602085015b83811015614b7f578035614e74816148e3565b835260209283019201614e61565b5f5f5f5f60808587031215614e95575f5ffd5b84356001600160401b03811115614eaa575f5ffd5b614eb687828801614dbc565b94505060208501356001600160401b03811115614ed1575f5ffd5b614edd87828801614e1e565b93505060408501356001600160401b03811115614ef8575f5ffd5b614f0487828801614e1e565b92505060608501356001600160401b03811115614f1f575f5ffd5b8501601f81018713614f2f575f5ffd5b8035614f3d6149b182614b02565b8082825260208201915060208360051b850101925089831115614f5e575f5ffd5b6020840193505b82841015614f89578335614f78816149e6565b825260209384019390910190614f65565b969995985093965050505050565b5f5f5f60608486031215614fa9575f5ffd5b614fb2846148cb565b92506020840135614fc2816148e3565b91506040840135614fd2816149e6565b809150509250925092565b5f60208284031215614fed575f5ffd5b5035919050565b8061040081018310156113e1575f5ffd5b5f5f5f6104408486031215615018575f5ffd5b833592506150298560208601614ff4565b915061503861042085016148cb565b90509250925092565b5f5f5f60608486031215615053575f5ffd5b83356001600160401b03811115615068575f5ffd5b61507486828701614dbc565b93505060208401356001600160401b0381111561508f575f5ffd5b61509b86828701614e1e565b92505060408401356001600160401b038111156150b6575f5ffd5b6150c286828701614b24565b9150509250925092565b5f5f5f5f606085870312156150df575f5ffd5b84356150ea816148e3565b93506020850135925060408501356001600160401b0381111561510b575f5ffd5b61511787828801614bf7565b95989497509550505050565b5f5f5f5f5f5f60a08789031215615138575f5ffd5b615141876148cb565b95506020870135615151816148e3565b9450604087013593506060870135615168816149e6565b925060808701356001600160401b03811115615182575f5ffd5b61518e89828a01614bf7565b979a9699509497509295939492505050565b5f5f604083850312156151b1575f5ffd5b82356151bc816148e3565b91506020830135614bec816149e6565b5f5f604083850312156151dd575f5ffd5b6151e6836148cb565b91506151f4602084016148cb565b90509250929050565b5f5f5f5f5f5f5f5f5f5f5f5f6109208d8f031215615219575f5ffd5b6152238e8e614ff4565b9b506152338e6104008f01614ff4565b9a506108008d013599506108208d013598506108408d0135975061525a6108608e016148cb565b965061526a6108808e01356148e3565b6108808d0135955061527f6108a08e016148cb565b94506108c08d0135615290816148e3565b93506108e08d013592506001600160401b036109008e013511156152b2575f5ffd5b6152c38e6109008f01358f01614bf7565b81935080925050509295989b509295989b509295989b565b5f5f5f5f5f5f5f60c0888a0312156152f1575f5ffd5b6152fa886148cb565b9650602088013561530a816148e3565b9550604088013594506060880135615321816148e3565b93506080880135615331816149e6565b925060a08801356001600160401b0381111561534b575f5ffd5b6153578a828b01614bf7565b989b979a50959850939692959293505050565b5f5f5f6040848603121561537c575f5ffd5b83356001600160401b03811115615391575f5ffd5b8401601f810186136153a1575f5ffd5b80356001600160401b038111156153b6575f5ffd5b8660208260051b84010111156153ca575f5ffd5b6020918201979096509401359392505050565b5f5f5f5f5f5f60c087890312156153f2575f5ffd5b6153fb876148cb565b9550602087013561540b816148e3565b9450615419604088016148cb565b93506060870135615429816148e3565b92506080870135615439816148e3565b915060a08701356001600160401b03811115615453575f5ffd5b61545f89828a01614994565b9150509295509295509295565b5f5f5f5f6104608587031215615480575f5ffd5b843593506154918660208701614ff4565b92506154a061042086016148cb565b939692955092936104400135925050565b5f5f5f5f61084085870312156154c5575f5ffd5b843593506154d68660208701614ff4565b925061042085013591506154ee866104408701614ff4565b905092959194509250565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6001600160a01b0392831681529116602082015260400190565b600181811c9082168061557557607f821691505b60208210810361559357634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115611abb57805f5260205f20601f840160051c810160208510156155be5750805b601f840160051c820191505b81811015611450575f81556001016155ca565b81516001600160401b038111156155f6576155f6614902565b61560a816156048454615561565b84615599565b6020601f82116001811461563c575f83156156255750848201515b5f19600385901b1c1916600184901b178455611450565b5f84815260208120601f198516915b8281101561566b578785015182556020948501946001909201910161564b565b508482101561568857868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f602082840312156156a7575f5ffd5b8151611b82816148e3565b634e487b7160e01b5f52603260045260245ffd5b60e09290921b6001600160e01b031916825260601b6001600160601b031916600482015260180190565b5f6156fd6149b18461496e565b9050828152838383011115615710575f5ffd5b611b82836020830184614caf565b5f82601f83011261572d575f5ffd5b611b82838351602085016156f0565b5f6020828403121561574c575f5ffd5b81516001600160401b03811115615761575f5ffd5b614bba8482850161571e565b5f6020828403121561577d575f5ffd5b8151611b8281614d29565b606081525f61579a6060830186614cd1565b82810360208401526157ac8186614cd1565b91505060ff83166040830152949350505050565b5f602082840312156157d0575f5ffd5b81516001600160401b038111156157e5575f5ffd5b8201601f810184136157f5575f5ffd5b614bba848251602084016156f0565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176113e1576113e1615804565b808201808211156113e1576113e1615804565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6104008b82376104008a610400830137886108008201528761082082015263ffffffff871661084082015260018060a01b03861661086082015263ffffffff8516610880820152836108a08201526108e06108c08201525f6158d16108e083018486615842565b9c9b505050505050505050505050565b5f602082840312156158f1575f5ffd5b5051919050565b818103818111156113e1576113e1615804565b60ff8916815263ffffffff88811660208301526001600160a01b03888116604084015290871660608301528516608082015260a0810184905261010060c082018190525f9061595c90830185614cd1565b905063ffffffff831660e08301529998505050505050505050565b5f823560de1983360301811261598b575f5ffd5b9190910192915050565b5f60e082360312156159a5575f5ffd5b6159ad614916565b6159b683614d37565b81526159c4602084016148cb565b60208201526159d5604084016148f7565b60408201526159e6606084016148cb565b60608201526159f7608084016148f7565b608082015260a0838101359082015260c08301356001600160401b03811115615a1e575f5ffd5b615a2a36828601614994565b60c08301525092915050565b5f5f8335601e19843603018112615a4b575f5ffd5b83016020810192503590506001600160401b03811115615a69575f5ffd5b803603821315614c34575f5ffd5b602080825281018290525f6040600584901b83018101908301858360de1936839003015b87821015615b7857868503603f190184528235818112615ab9575f5ffd5b89018035615ac681614d29565b60ff16865263ffffffff615adc602083016148cb565b1660208701526040810135615af0816148e3565b6001600160a01b03166040870152615b0a606082016148cb565b63ffffffff166060870152615b21608082016148f7565b6001600160a01b0316608087015260a08181013590870152615b4660c0820182615a36565b915060e060c0880152615b5d60e088018383615842565b96505050602083019250602084019350600182019150615a9b565b5092979650505050505050565b85815284602082015283604082015282606082015260a060808201525f615baf60a0830184614cd1565b979650505050505050565b5f825161598b818460208701614caf565b818382375f9101908152919050565b94855263ffffffff9390931660208501526001600160a01b039182166040850152166060830152608082015260a00190565b6001600160a01b038516815263ffffffff841660208201526060604082018190525f906148429083018486615842565b5f5f5f60608486031215615c4e575f5ffd5b83516001600160401b03811115615c63575f5ffd5b615c6f8682870161571e565b93505060208401516001600160401b03811115615c8a575f5ffd5b615c968682870161571e565b9250506040840151614fd281614d29565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b63ffffffff81811683821601908111156113e1576113e1615804565b60ff8a16815263ffffffff89811660208301526001600160a01b03898116604084015290881660608301528616608082015260a0810185905261010060c082018190525f90615d609083018587615842565b905063ffffffff831660e08301529a9950505050505050505050565b6001600160a01b03861681526080602082018190525f90615da09083018688615842565b6001600160a01b03948516604084015292909316606090910152949350505050565b5f60208284031215615dd2575f5ffd5b8151611b82816149e6565b6001600160a01b03929092168252602082015260400190565b63ffffffff861681526001600160a01b038581166020830152841660408201526080606082018190525f90615baf9083018486615842565b5f60018201615e3f57615e3f615804565b5060010190565b63ffffffff9490941684526001600160a01b039290921660208401526040830152606082015260800190565b6001815b6001841115615ead57808504811115615e9157615e91615804565b6001841615615e9f57908102905b60019390931c928002615e76565b935093915050565b5f82615ec3575060016113e1565b81615ecf57505f6113e1565b8160018114615ee55760028114615eef57615f0b565b60019150506113e1565b60ff841115615f0057615f00615804565b50506001821b6113e1565b5060208310610133831016604e8410600b8410161715615f2e575081810a6113e1565b615f3a5f198484615e72565b805f1904821115615f4d57615f4d615804565b029392505050565b5f611b828383615eb5565b634e487b7160e01b5f52600160045260245ffdfe85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793a26469706673582212203bb24551b4426931dcaa107e9d105ece95994de559ac4967037369c5a51eaf7564736f6c634300081c00336080604052348015600e575f5ffd5b50610fb38061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c8063c514f24e1461002d575b5f5ffd5b61003561004b565b604051610042919061006a565b60405180910390f35b60405180610f000160405280610ec881526020016100b6610ec8913981565b602081525f82518060208401525f5b818110156100965760208186018101516040868401015201610079565b505f604082850101526040601f19601f8301168401019150509291505056fe60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d7847fa500b7c717aa6a7e73128e0f74956c8497543c956e303def2d68265f2564736f6c634300081c00336080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6114c2806100d65f395ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c806370a082311161008f57806370a08231146101a05780637ecebe00146101b357806384b0196e146101c657806395d89b41146101e15780639dc29fac146101e9578063a3c573eb146101fc578063a9059cbb14610211578063d505accf14610224578063dd62ed3e14610237575f5ffd5b806306fdde03146100ec578063095ea7b31461010a5780631624f6c61461012d57806318160ddd1461014257806323b872dd14610158578063313ce5671461016b5780633644e5151461018557806340c10f191461018d575b5f5ffd5b6100f461024a565b6040516101019190610ff3565b60405180910390f35b61011d610118366004611027565b6102e8565b6040519015158152602001610101565b61014061013b3660046110fc565b610301565b005b61014a61043e565b604051908152602001610101565b61011d61016636600461116e565b610452565b610173610475565b60405160ff9091168152602001610101565b61014a610489565b61014061019b366004611027565b610497565b61014a6101ae3660046111a8565b6104e1565b61014a6101c13660046111a8565b61050a565b6101ce610514565b60405161010197969594939291906111c1565b6100f46105bd565b6101406101f7366004611027565b6105d9565b61020461061e565b6040516101019190611257565b61011d61021f366004611027565b61063d565b61014061023236600461126b565b61064a565b61014a6102453660046112d1565b61079f565b60605f6102556107d9565b905080600301805461026690611302565b80601f016020809104026020016040519081016040528092919081815260200182805461029290611302565b80156102dd5780601f106102b4576101008083540402835291602001916102dd565b820191905f5260205f20905b8154815290600101906020018083116102c057829003601f168201915b505050505091505090565b5f336102f58185856107fd565b60019150505b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156103455750825b90505f826001600160401b031660011480156103605750303b155b90508115801561036e575080155b1561038c5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156103b657845460ff60401b1916600160401b1785555b6103c0888861080a565b6103c988610820565b5f6103d261084e565b805433610100026001600160a81b031990911660ff8a161717905550831561043457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f5f6104486107d9565b6002015492915050565b5f3361045f858285610872565b61046a8585856108c2565b506001949350505050565b5f5f61047f61084e565b5460ff1692915050565b5f61049261091f565b905090565b5f6104a061084e565b805490915061010090046001600160a01b031633146104d2576040516338da3b1560e01b815260040160405180910390fd5b6104dc8383610928565b505050565b5f5f6104eb6107d9565b6001600160a01b039093165f9081526020939093525050604090205490565b5f6102fb8261095c565b5f6060805f5f5f60605f610526610966565b805490915015801561053a57506001810154155b6105835760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b61058b61098a565b6105936109a6565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b60605f6105c86107d9565b905080600401805461026690611302565b5f6105e261084e565b805490915061010090046001600160a01b03163314610614576040516338da3b1560e01b815260040160405180910390fd5b6104dc83836109b1565b5f5f61062861084e565b5461010090046001600160a01b031692915050565b5f336102f58185856108c2565b8342111561066e5760405163313c898160e11b81526004810185905260240161057a565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886106d88c6001600160a01b03165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610732826109e5565b90505f61074182878787610a11565b9050896001600160a01b0316816001600160a01b031614610788576040516325c0072360e11b81526001600160a01b0380831660048301528b16602482015260440161057a565b6107938a8a8a6107fd565b50505050505050505050565b5f5f6107a96107d9565b6001600160a01b039485165f90815260019190910160209081526040808320959096168252939093525050205490565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0090565b6104dc8383836001610a3d565b610812610b1e565b61081c8282610b69565b5050565b610828610b1e565b61084b81604051806040016040528060018152602001603160f81b815250610b99565b50565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d0090565b5f61087d848461079f565b90505f1981146108bc57818110156108ae57828183604051637dc7a0d960e11b815260040161057a9392919061133a565b6108bc84848484035f610a3d565b50505050565b6001600160a01b0383166108eb575f604051634b637e8f60e11b815260040161057a9190611257565b6001600160a01b038216610914575f60405163ec442f0560e01b815260040161057a9190611257565b6104dc838383610bd8565b5f610492610cfb565b6001600160a01b038216610951575f60405163ec442f0560e01b815260040161057a9190611257565b61081c5f8383610bd8565b5f6102fb82610d6e565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10090565b60605f610995610966565b905080600201805461026690611302565b60605f610255610966565b6001600160a01b0382166109da575f604051634b637e8f60e11b815260040161057a9190611257565b61081c825f83610bd8565b5f6102fb6109f161091f565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610a2188888888610d96565b925092509250610a318282610e54565b50909695505050505050565b5f610a466107d9565b90506001600160a01b038516610a71575f60405163e602df0560e01b815260040161057a9190611257565b6001600160a01b038416610a9a575f604051634a1406b160e11b815260040161057a9190611257565b6001600160a01b038086165f90815260018301602090815260408083209388168352929052208390558115610b1757836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610b0e91815260200190565b60405180910390a35b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610b6757604051631afcd79f60e31b815260040160405180910390fd5b565b610b71610b1e565b5f610b7a6107d9565b905060038101610b8a848261139f565b50600481016108bc838261139f565b610ba1610b1e565b5f610baa610966565b905060028101610bba848261139f565b5060038101610bc9838261139f565b505f8082556001909101555050565b5f610be16107d9565b90506001600160a01b038416610c0f5781816002015f828254610c049190611459565b90915550610c6c9050565b6001600160a01b0384165f9081526020829052604090205482811015610c4e5784818460405163391434e360e21b815260040161057a9392919061133a565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b038316610c8a576002810180548390039055610ca8565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ced91815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f610d25610f0c565b610d2d610f71565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006104eb565b5f80806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03841115610dc557505f91506003905082610e4a565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610e16573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610e4157505f925060019150829050610e4a565b92505f91508190505b9450945094915050565b5f826003811115610e6757610e67611478565b03610e70575050565b6001826003811115610e8457610e84611478565b03610ea25760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610eb657610eb6611478565b03610ed75760405163fce698f760e01b81526004810182905260240161057a565b6003826003811115610eeb57610eeb611478565b0361081c576040516335e2f38360e21b81526004810182905260240161057a565b5f5f610f16610966565b90505f610f2161098a565b805190915015610f3957805160209091012092915050565b81548015610f48579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f610f7b610966565b90505f610f866109a6565b805190915015610f9e57805160209091012092915050565b60018201548015610f48579392505050565b5f81518084525f5b81811015610fd457602081850181015186830182015201610fb8565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f6110056020830184610fb0565b9392505050565b80356001600160a01b0381168114611022575f5ffd5b919050565b5f5f60408385031215611038575f5ffd5b6110418361100c565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112611072575f5ffd5b81356001600160401b0381111561108b5761108b61104f565b604051601f8201601f19908116603f011681016001600160401b03811182821017156110b9576110b961104f565b6040528181528382016020018510156110d0575f5ffd5b816020850160208301375f918101602001919091529392505050565b803560ff81168114611022575f5ffd5b5f5f5f6060848603121561110e575f5ffd5b83356001600160401b03811115611123575f5ffd5b61112f86828701611063565b93505060208401356001600160401b0381111561114a575f5ffd5b61115686828701611063565b925050611165604085016110ec565b90509250925092565b5f5f5f60608486031215611180575f5ffd5b6111898461100c565b92506111976020850161100c565b929592945050506040919091013590565b5f602082840312156111b8575f5ffd5b6110058261100c565b60ff60f81b8816815260e060208201525f6111df60e0830189610fb0565b82810360408401526111f18189610fb0565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611246578351835260209384019390920191600101611228565b50909b9a5050505050505050505050565b6001600160a01b0391909116815260200190565b5f5f5f5f5f5f5f60e0888a031215611281575f5ffd5b61128a8861100c565b96506112986020890161100c565b955060408801359450606088013593506112b4608089016110ec565b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156112e2575f5ffd5b6112eb8361100c565b91506112f96020840161100c565b90509250929050565b600181811c9082168061131657607f821691505b60208210810361133457634e487b7160e01b5f52602260045260245ffd5b50919050565b6001600160a01b039390931683526020830191909152604082015260600190565b601f8211156104dc57805f5260205f20601f840160051c810160208510156113805750805b601f840160051c820191505b81811015610b17575f815560010161138c565b81516001600160401b038111156113b8576113b861104f565b6113cc816113c68454611302565b8461135b565b6020601f8211600181146113fe575f83156113e75750848201515b5f19600385901b1c1916600184901b178455610b17565b5f84815260208120601f198516915b8281101561142d578785015182556020948501946001909201910161140d565b508482101561144a57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156102fb57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea2646970667358221220734889424b5f04b57ae9c0f182a620f22fa0c90954fa75a81eadce83f886d8bc64736f6c634300081c00336080604052348015600e575f5ffd5b50610cc28061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063737ce34114610059578063a28fa4a314610082578063be3dcf62146100a5578063c00f14ab146100ca578063cf825e55146100dd575b5f5ffd5b61006c6100673660046108ac565b6100f0565b604051610079919061091b565b60405180910390f35b61009561009036600461092d565b6101b5565b6040519015158152602001610079565b6100b86100b33660046108ac565b610559565b60405160ff9091168152602001610079565b61006c6100d83660046108ac565b61060d565b61006c6100eb3660046108ac565b610652565b60408051600481526024810182526020810180516001600160e01b03166395d89b4160e01b17905290516060915f9182916001600160a01b0386169161013691906109d0565b5f60405180830381855afa9150503d805f811461016e576040519150601f19603f3d011682016040523d82523d5f602084013e610173565b606091505b5091509150816101a457604051806040016040528060098152602001681393d7d4d6535093d360ba1b8152506101ad565b6101ad81610704565b949350505050565b5f806101c460048287896109eb565b6101cd91610a12565b9050632afa533160e01b6001600160e01b031982160161037c575f5f5f5f5f5f5f8c8c6004908092610201939291906109eb565b81019061020e9190610a58565b96509650965096509650965096508a6001600160a01b0316876001600160a01b03161461024e5760405163912ecce760e01b815260040160405180910390fd5b896001600160a01b0316866001600160a01b0316146102805760405163750643af60e01b815260040160405180910390fd5b5f8e6001600160a01b031663d505accf60e01b898989898989896040516024016102ec97969594939291906001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252905161032a91906109d0565b5f604051808303815f865af19150503d805f8114610363576040519150601f19603f3d011682016040523d82523d5f602084013e610368565b606091505b50909a506105509950505050505050505050565b631c0d143d60e21b6001600160e01b0319821601610537575f5f5f5f5f5f5f5f8d8d60049080926103af939291906109eb565b8101906103bc9190610ac4565b975097509750975097509750975097508b6001600160a01b0316886001600160a01b0316146103fe5760405163912ecce760e01b815260040160405180910390fd5b8a6001600160a01b0316876001600160a01b0316146104305760405163750643af60e01b815260040160405180910390fd5b5f8f6001600160a01b0316638fcbaf0c60e01b8a8a8a8a8a8a8a8a6040516024016104a69897969594939291906001600160a01b039889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516104e491906109d0565b5f604051808303815f865af19150503d805f811461051d576040519150601f19603f3d011682016040523d82523d5f602084013e610522565b606091505b50909b506105509a5050505050505050505050565b604051637141605d60e11b815260040160405180910390fd5b95945050505050565b60408051600481526024810182526020810180516001600160e01b031663313ce56760e01b17905290515f91829182916001600160a01b0386169161059e91906109d0565b5f60405180830381855afa9150503d805f81146105d6576040519150601f19603f3d011682016040523d82523d5f602084013e6105db565b606091505b50915091508180156105ee575080516020145b6105f95760126101ad565b808060200190518101906101ad9190610b46565b606061061882610652565b610621836100f0565b61062a84610559565b60405160200161063c93929190610b61565b6040516020818303038152906040529050919050565b60408051600481526024810182526020810180516001600160e01b03166306fdde0360e01b17905290516060915f9182916001600160a01b0386169161069891906109d0565b5f60405180830381855afa9150503d805f81146106d0576040519150601f19603f3d011682016040523d82523d5f602084013e6106d5565b606091505b5091509150816101a457604051806040016040528060078152602001664e4f5f4e414d4560c81b8152506101ad565b6060604082511061072957818060200190518101906107239190610bad565b92915050565b8151602003610857575f5b602081108015610763575082818151811061075157610751610c54565b01602001516001600160f81b03191615155b1561077a578061077281610c68565b915050610734565b805f036107b15750506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b6020820152919050565b5f8167ffffffffffffffff8111156107cb576107cb610b99565b6040519080825280601f01601f1916602001820160405280156107f5576020820181803683370190505b5090505f5b8281101561084f5784818151811061081457610814610c54565b602001015160f81c60f81b82828151811061083157610831610c54565b60200101906001600160f81b03191690815f1a9053506001016107fa565b509392505050565b50506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b602082015290565b919050565b6001600160a01b038116811461089e575f5ffd5b50565b80356108858161088a565b5f602082840312156108bc575f5ffd5b81356108c78161088a565b9392505050565b5f5b838110156108e85781810151838201526020016108d0565b50505f910152565b5f81518084526109078160208601602086016108ce565b601f01601f19169290920160200192915050565b602081525f6108c760208301846108f0565b5f5f5f5f5f60808688031215610941575f5ffd5b853561094c8161088a565b9450602086013567ffffffffffffffff811115610967575f5ffd5b8601601f81018813610977575f5ffd5b803567ffffffffffffffff81111561098d575f5ffd5b88602082840101111561099e575f5ffd5b6020919091019450925060408601356109b68161088a565b91506109c4606087016108a1565b90509295509295909350565b5f82516109e18184602087016108ce565b9190910192915050565b5f5f858511156109f9575f5ffd5b83861115610a05575f5ffd5b5050820193919092039150565b80356001600160e01b03198116906004841015610a43576001600160e01b0319600485900360031b81901b82161691505b5092915050565b60ff8116811461089e575f5ffd5b5f5f5f5f5f5f5f60e0888a031215610a6e575f5ffd5b8735610a798161088a565b96506020880135610a898161088a565b955060408801359450606088013593506080880135610aa781610a4a565b9699959850939692959460a0840135945060c09093013592915050565b5f5f5f5f5f5f5f5f610100898b031215610adc575f5ffd5b8835610ae78161088a565b97506020890135610af78161088a565b9650604089013595506060890135945060808901358015158114610b19575f5ffd5b935060a0890135610b2981610a4a565b979a969950949793969295929450505060c08201359160e0013590565b5f60208284031215610b56575f5ffd5b81516108c781610a4a565b606081525f610b7360608301866108f0565b8281036020840152610b8581866108f0565b91505060ff83166040830152949350505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610bbd575f5ffd5b815167ffffffffffffffff811115610bd3575f5ffd5b8201601f81018413610be3575f5ffd5b805167ffffffffffffffff811115610bfd57610bfd610b99565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610c2c57610c2c610b99565b604052818152828201602001861015610c43575f5ffd5b6105508260208301602086016108ce565b634e487b7160e01b5f52603260045260245ffd5b5f60018201610c8557634e487b7160e01b5f52601160045260245ffd5b506001019056fea26469706673582212206db9941992a6c13c8af2a735afe7360f7dfde05688ae4632f7209a0b5512f84d64736f6c634300081c0033",
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

// BRIDGESOVEREIGNVERSION is a free data retrieval call binding the contract method 0xf67566e4.
//
// Solidity: function BRIDGE_SOVEREIGN_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BRIDGESOVEREIGNVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "BRIDGE_SOVEREIGN_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BRIDGESOVEREIGNVERSION is a free data retrieval call binding the contract method 0xf67566e4.
//
// Solidity: function BRIDGE_SOVEREIGN_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BRIDGESOVEREIGNVERSION() (string, error) {
	return _Bridgel2sovereignchain.Contract.BRIDGESOVEREIGNVERSION(&_Bridgel2sovereignchain.CallOpts)
}

// BRIDGESOVEREIGNVERSION is a free data retrieval call binding the contract method 0xf67566e4.
//
// Solidity: function BRIDGE_SOVEREIGN_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BRIDGESOVEREIGNVERSION() (string, error) {
	return _Bridgel2sovereignchain.Contract.BRIDGESOVEREIGNVERSION(&_Bridgel2sovereignchain.CallOpts)
}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BRIDGEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "BRIDGE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BRIDGEVERSION() (string, error) {
	return _Bridgel2sovereignchain.Contract.BRIDGEVERSION(&_Bridgel2sovereignchain.CallOpts)
}

// BRIDGEVERSION is a free data retrieval call binding the contract method 0x65d6f654.
//
// Solidity: function BRIDGE_VERSION() view returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BRIDGEVERSION() (string, error) {
	return _Bridgel2sovereignchain.Contract.BRIDGEVERSION(&_Bridgel2sovereignchain.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.INITBYTECODETRANSPARENTPROXY(&_Bridgel2sovereignchain.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Bridgel2sovereignchain.Contract.INITBYTECODETRANSPARENTPROXY(&_Bridgel2sovereignchain.CallOpts)
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

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) BridgeLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "bridgeLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BridgeLib() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.BridgeLib(&_Bridgel2sovereignchain.CallOpts)
}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) BridgeLib() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.BridgeLib(&_Bridgel2sovereignchain.CallOpts)
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

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ClaimedGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "claimedGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.ClaimedGlobalIndexHashChain(&_Bridgel2sovereignchain.CallOpts)
}

// ClaimedGlobalIndexHashChain is a free data retrieval call binding the contract method 0x1d081d8c.
//
// Solidity: function claimedGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ClaimedGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.ClaimedGlobalIndexHashChain(&_Bridgel2sovereignchain.CallOpts)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ComputeTokenProxyAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "computeTokenProxyAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.ComputeTokenProxyAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.ComputeTokenProxyAddress(&_Bridgel2sovereignchain.CallOpts, originNetwork, originTokenAddress)
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

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) EmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "emergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) EmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.EmergencyBridgePauser(&_Bridgel2sovereignchain.CallOpts)
}

// EmergencyBridgePauser is a free data retrieval call binding the contract method 0x2f84c690.
//
// Solidity: function emergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) EmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.EmergencyBridgePauser(&_Bridgel2sovereignchain.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) EmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "emergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.EmergencyBridgeUnpauser(&_Bridgel2sovereignchain.CallOpts)
}

// EmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0xae24490a.
//
// Solidity: function emergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) EmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.EmergencyBridgeUnpauser(&_Bridgel2sovereignchain.CallOpts)
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

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
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

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) GetWrappedTokenBridgeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "getWrappedTokenBridgeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetWrappedTokenBridgeImplementation(&_Bridgel2sovereignchain.CallOpts)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.GetWrappedTokenBridgeImplementation(&_Bridgel2sovereignchain.CallOpts)
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

// Initialize0 is a free data retrieval call binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) Initialize0(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "initialize0")

	if err != nil {
		return err
	}

	return err

}

// Initialize0 is a free data retrieval call binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize0() error {
	return _Bridgel2sovereignchain.Contract.Initialize0(&_Bridgel2sovereignchain.CallOpts)
}

// Initialize0 is a free data retrieval call binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() pure returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) Initialize0() error {
	return _Bridgel2sovereignchain.Contract.Initialize0(&_Bridgel2sovereignchain.CallOpts)
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

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) LocalBalanceTree(opts *bind.CallOpts, tokenInfoHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "localBalanceTree", tokenInfoHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.LocalBalanceTree(&_Bridgel2sovereignchain.CallOpts, tokenInfoHash)
}

// LocalBalanceTree is a free data retrieval call binding the contract method 0xd9cb3aec.
//
// Solidity: function localBalanceTree(bytes32 tokenInfoHash) view returns(uint256 amount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) LocalBalanceTree(tokenInfoHash [32]byte) (*big.Int, error) {
	return _Bridgel2sovereignchain.Contract.LocalBalanceTree(&_Bridgel2sovereignchain.CallOpts, tokenInfoHash)
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

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PendingEmergencyBridgePauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "pendingEmergencyBridgePauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingEmergencyBridgePauser(&_Bridgel2sovereignchain.CallOpts)
}

// PendingEmergencyBridgePauser is a free data retrieval call binding the contract method 0x03e6e116.
//
// Solidity: function pendingEmergencyBridgePauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PendingEmergencyBridgePauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingEmergencyBridgePauser(&_Bridgel2sovereignchain.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PendingEmergencyBridgeUnpauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "pendingEmergencyBridgeUnpauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingEmergencyBridgeUnpauser(&_Bridgel2sovereignchain.CallOpts)
}

// PendingEmergencyBridgeUnpauser is a free data retrieval call binding the contract method 0x606617ff.
//
// Solidity: function pendingEmergencyBridgeUnpauser() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PendingEmergencyBridgeUnpauser() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingEmergencyBridgeUnpauser(&_Bridgel2sovereignchain.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) PendingProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "pendingProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.PendingProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
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

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) ProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "proxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.ProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
}

// ProxiedTokensManager is a free data retrieval call binding the contract method 0x6e4ecfed.
//
// Solidity: function proxiedTokensManager() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) ProxiedTokensManager() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.ProxiedTokensManager(&_Bridgel2sovereignchain.CallOpts)
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

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) UnsetGlobalIndexHashChain(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "unsetGlobalIndexHashChain")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.UnsetGlobalIndexHashChain(&_Bridgel2sovereignchain.CallOpts)
}

// UnsetGlobalIndexHashChain is a free data retrieval call binding the contract method 0x6ee84b23.
//
// Solidity: function unsetGlobalIndexHashChain() view returns(bytes32)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) UnsetGlobalIndexHashChain() ([32]byte, error) {
	return _Bridgel2sovereignchain.Contract.UnsetGlobalIndexHashChain(&_Bridgel2sovereignchain.CallOpts)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Version() (string, error) {
	return _Bridgel2sovereignchain.Contract.Version(&_Bridgel2sovereignchain.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) Version() (string, error) {
	return _Bridgel2sovereignchain.Contract.Version(&_Bridgel2sovereignchain.CallOpts)
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

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCaller) WrappedTokenBytecodeStorer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridgel2sovereignchain.contract.Call(opts, &out, "wrappedTokenBytecodeStorer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) WrappedTokenBytecodeStorer() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.WrappedTokenBytecodeStorer(&_Bridgel2sovereignchain.CallOpts)
}

// WrappedTokenBytecodeStorer is a free data retrieval call binding the contract method 0x381fef6d.
//
// Solidity: function wrappedTokenBytecodeStorer() view returns(address)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainCallerSession) WrappedTokenBytecodeStorer() (common.Address, error) {
	return _Bridgel2sovereignchain.Contract.WrappedTokenBytecodeStorer(&_Bridgel2sovereignchain.CallOpts)
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

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) AcceptEmergencyBridgePauserRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "acceptEmergencyBridgePauserRole")
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) AcceptEmergencyBridgePauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptEmergencyBridgePauserRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0xe88f0436.
//
// Solidity: function acceptEmergencyBridgePauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) AcceptEmergencyBridgePauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptEmergencyBridgePauserRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) AcceptEmergencyBridgeUnpauserRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "acceptEmergencyBridgeUnpauserRole")
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) AcceptEmergencyBridgeUnpauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8b37b873.
//
// Solidity: function acceptEmergencyBridgeUnpauserRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) AcceptEmergencyBridgeUnpauserRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) AcceptProxiedTokensManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "acceptProxiedTokensManagerRole")
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptProxiedTokensManagerRole(&_Bridgel2sovereignchain.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.AcceptProxiedTokensManagerRole(&_Bridgel2sovereignchain.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ActivateEmergencyState(&_Bridgel2sovereignchain.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ActivateEmergencyState(&_Bridgel2sovereignchain.TransactOpts)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) BackwardLET(opts *bind.TransactOpts, newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "backwardLET", newDepositCount, newFrontier, nextLeaf, proof)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) BackwardLET(newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BackwardLET(&_Bridgel2sovereignchain.TransactOpts, newDepositCount, newFrontier, nextLeaf, proof)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) BackwardLET(newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.BackwardLET(&_Bridgel2sovereignchain.TransactOpts, newDepositCount, newFrontier, nextLeaf, proof)
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

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchain.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.DeactivateEmergencyState(&_Bridgel2sovereignchain.TransactOpts)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) DeployWrappedTokenAndRemap(opts *bind.TransactOpts, originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "deployWrappedTokenAndRemap", originNetwork, originTokenAddress, isNotMintable)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) DeployWrappedTokenAndRemap(originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.DeployWrappedTokenAndRemap(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, isNotMintable)
}

// DeployWrappedTokenAndRemap is a paid mutator transaction binding the contract method 0x6e974cd4.
//
// Solidity: function deployWrappedTokenAndRemap(uint32 originNetwork, address originTokenAddress, bool isNotMintable) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) DeployWrappedTokenAndRemap(originNetwork uint32, originTokenAddress common.Address, isNotMintable bool) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.DeployWrappedTokenAndRemap(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, isNotMintable)
}

// ForwardLET is a paid mutator transaction binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] newLeaves, bytes32 expectedLER) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) ForwardLET(opts *bind.TransactOpts, newLeaves []AgglayerBridgeL2LeafData, expectedLER [32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "forwardLET", newLeaves, expectedLER)
}

// ForwardLET is a paid mutator transaction binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] newLeaves, bytes32 expectedLER) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) ForwardLET(newLeaves []AgglayerBridgeL2LeafData, expectedLER [32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ForwardLET(&_Bridgel2sovereignchain.TransactOpts, newLeaves, expectedLER)
}

// ForwardLET is a paid mutator transaction binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] newLeaves, bytes32 expectedLER) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) ForwardLET(newLeaves []AgglayerBridgeL2LeafData, expectedLER [32]byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.ForwardLET(&_Bridgel2sovereignchain.TransactOpts, newLeaves, expectedLER)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize(&_Bridgel2sovereignchain.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x006ee171.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata, address _bridgeManager, address _sovereignWETHAddress, bool _sovereignWETHAddressIsNotMintable, address _emergencyBridgePauser, address _emergencyBridgeUnpauser, address _proxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte, _bridgeManager common.Address, _sovereignWETHAddress common.Address, _sovereignWETHAddressIsNotMintable bool, _emergencyBridgePauser common.Address, _emergencyBridgeUnpauser common.Address, _proxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize(&_Bridgel2sovereignchain.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata, _bridgeManager, _sovereignWETHAddress, _sovereignWETHAddressIsNotMintable, _emergencyBridgePauser, _emergencyBridgeUnpauser, _proxiedTokensManager)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) Initialize1(opts *bind.TransactOpts, arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "initialize1", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) Initialize1(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize1(&_Bridgel2sovereignchain.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 , address , uint32 , address , address , bytes ) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) Initialize1(arg0 uint32, arg1 common.Address, arg2 uint32, arg3 common.Address, arg4 common.Address, arg5 []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.Initialize1(&_Bridgel2sovereignchain.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) MigrateLegacyToken(opts *bind.TransactOpts, legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "migrateLegacyToken", legacyTokenAddress, amount, permitData)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.MigrateLegacyToken(&_Bridgel2sovereignchain.TransactOpts, legacyTokenAddress, amount, permitData)
}

// MigrateLegacyToken is a paid mutator transaction binding the contract method 0xb0b37920.
//
// Solidity: function migrateLegacyToken(address legacyTokenAddress, uint256 amount, bytes permitData) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) MigrateLegacyToken(legacyTokenAddress common.Address, amount *big.Int, permitData []byte) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.MigrateLegacyToken(&_Bridgel2sovereignchain.TransactOpts, legacyTokenAddress, amount, permitData)
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

// SetLocalBalanceTree is a paid mutator transaction binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] originNetwork, address[] originTokenAddress, uint256[] amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetLocalBalanceTree(opts *bind.TransactOpts, originNetwork []uint32, originTokenAddress []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setLocalBalanceTree", originNetwork, originTokenAddress, amount)
}

// SetLocalBalanceTree is a paid mutator transaction binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] originNetwork, address[] originTokenAddress, uint256[] amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetLocalBalanceTree(originNetwork []uint32, originTokenAddress []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetLocalBalanceTree(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, amount)
}

// SetLocalBalanceTree is a paid mutator transaction binding the contract method 0x8f9720bb.
//
// Solidity: function setLocalBalanceTree(uint32[] originNetwork, address[] originTokenAddress, uint256[] amount) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetLocalBalanceTree(originNetwork []uint32, originTokenAddress []common.Address, amount []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetLocalBalanceTree(&_Bridgel2sovereignchain.TransactOpts, originNetwork, originTokenAddress, amount)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) SetMultipleClaims(opts *bind.TransactOpts, globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "setMultipleClaims", globalIndexes)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) SetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetMultipleClaims(&_Bridgel2sovereignchain.TransactOpts, globalIndexes)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) SetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.SetMultipleClaims(&_Bridgel2sovereignchain.TransactOpts, globalIndexes)
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

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) TransferEmergencyBridgePauserRole(opts *bind.TransactOpts, newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "transferEmergencyBridgePauserRole", newEmergencyBridgePauser)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) TransferEmergencyBridgePauserRole(newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferEmergencyBridgePauserRole(&_Bridgel2sovereignchain.TransactOpts, newEmergencyBridgePauser)
}

// TransferEmergencyBridgePauserRole is a paid mutator transaction binding the contract method 0x69e3ab12.
//
// Solidity: function transferEmergencyBridgePauserRole(address newEmergencyBridgePauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) TransferEmergencyBridgePauserRole(newEmergencyBridgePauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferEmergencyBridgePauserRole(&_Bridgel2sovereignchain.TransactOpts, newEmergencyBridgePauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) TransferEmergencyBridgeUnpauserRole(opts *bind.TransactOpts, newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "transferEmergencyBridgeUnpauserRole", newEmergencyBridgeUnpauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) TransferEmergencyBridgeUnpauserRole(newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchain.TransactOpts, newEmergencyBridgeUnpauser)
}

// TransferEmergencyBridgeUnpauserRole is a paid mutator transaction binding the contract method 0x8d942096.
//
// Solidity: function transferEmergencyBridgeUnpauserRole(address newEmergencyBridgeUnpauser) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) TransferEmergencyBridgeUnpauserRole(newEmergencyBridgeUnpauser common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferEmergencyBridgeUnpauserRole(&_Bridgel2sovereignchain.TransactOpts, newEmergencyBridgeUnpauser)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) TransferProxiedTokensManagerRole(opts *bind.TransactOpts, newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "transferProxiedTokensManagerRole", newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferProxiedTokensManagerRole(&_Bridgel2sovereignchain.TransactOpts, newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.TransferProxiedTokensManagerRole(&_Bridgel2sovereignchain.TransactOpts, newProxiedTokensManager)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactor) UnsetMultipleClaims(opts *bind.TransactOpts, globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.contract.Transact(opts, "unsetMultipleClaims", globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainSession) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UnsetMultipleClaims(&_Bridgel2sovereignchain.TransactOpts, globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Bridgel2sovereignchain *Bridgel2sovereignchainTransactorSession) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Bridgel2sovereignchain.Contract.UnsetMultipleClaims(&_Bridgel2sovereignchain.TransactOpts, globalIndexes)
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

// Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator is returned from FilterAcceptEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgePauserRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator struct {
	Event *Bridgel2sovereignchainAcceptEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainAcceptEmergencyBridgePauserRole)
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
		it.Event = new(Bridgel2sovereignchainAcceptEmergencyBridgePauserRole)
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
func (it *Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainAcceptEmergencyBridgePauserRole represents a AcceptEmergencyBridgePauserRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptEmergencyBridgePauserRole struct {
	OldEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterAcceptEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainAcceptEmergencyBridgePauserRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "AcceptEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchAcceptEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainAcceptEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "AcceptEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainAcceptEmergencyBridgePauserRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
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

// ParseAcceptEmergencyBridgePauserRole is a log parse operation binding the contract event 0x85d2bdfbe58cd81abf8199c13ce2509204be4aba8603b9d29f52c4e13e7bb793.
//
// Solidity: event AcceptEmergencyBridgePauserRole(address oldEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseAcceptEmergencyBridgePauserRole(log types.Log) (*Bridgel2sovereignchainAcceptEmergencyBridgePauserRole, error) {
	event := new(Bridgel2sovereignchainAcceptEmergencyBridgePauserRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator is returned from FilterAcceptEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for AcceptEmergencyBridgeUnpauserRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator struct {
	Event *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole)
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
		it.Event = new(Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole)
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
func (it *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole represents a AcceptEmergencyBridgeUnpauserRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole struct {
	OldEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterAcceptEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterAcceptEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "AcceptEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchAcceptEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchAcceptEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "AcceptEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
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

// ParseAcceptEmergencyBridgeUnpauserRole is a log parse operation binding the contract event 0x24cc8295aa5110cc216695db944ad2458c7795c6404449be980c3ce14aed752d.
//
// Solidity: event AcceptEmergencyBridgeUnpauserRole(address oldEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseAcceptEmergencyBridgeUnpauserRole(log types.Log) (*Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole, error) {
	event := new(Bridgel2sovereignchainAcceptEmergencyBridgeUnpauserRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator is returned from FilterAcceptProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for AcceptProxiedTokensManagerRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator struct {
	Event *Bridgel2sovereignchainAcceptProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainAcceptProxiedTokensManagerRole)
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
		it.Event = new(Bridgel2sovereignchainAcceptProxiedTokensManagerRole)
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
func (it *Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainAcceptProxiedTokensManagerRole represents a AcceptProxiedTokensManagerRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainAcceptProxiedTokensManagerRole struct {
	OldProxiedTokensManager common.Address
	NewProxiedTokensManager common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAcceptProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterAcceptProxiedTokensManagerRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainAcceptProxiedTokensManagerRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "AcceptProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptProxiedTokensManagerRole is a free log subscription operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchAcceptProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainAcceptProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainAcceptProxiedTokensManagerRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
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

// ParseAcceptProxiedTokensManagerRole is a log parse operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseAcceptProxiedTokensManagerRole(log types.Log) (*Bridgel2sovereignchainAcceptProxiedTokensManagerRole, error) {
	event := new(Bridgel2sovereignchainAcceptProxiedTokensManagerRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainBackwardLETIterator is returned from FilterBackwardLET and is used to iterate over the raw logs and unpacked data for BackwardLET events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainBackwardLETIterator struct {
	Event *Bridgel2sovereignchainBackwardLET // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainBackwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainBackwardLET)
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
		it.Event = new(Bridgel2sovereignchainBackwardLET)
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
func (it *Bridgel2sovereignchainBackwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainBackwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainBackwardLET represents a BackwardLET event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainBackwardLET struct {
	PreviousDepositCount *big.Int
	PreviousRoot         [32]byte
	NewDepositCount      *big.Int
	NewRoot              [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBackwardLET is a free log retrieval operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterBackwardLET(opts *bind.FilterOpts) (*Bridgel2sovereignchainBackwardLETIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainBackwardLETIterator{contract: _Bridgel2sovereignchain.contract, event: "BackwardLET", logs: logs, sub: sub}, nil
}

// WatchBackwardLET is a free log subscription operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchBackwardLET(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainBackwardLET) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainBackwardLET)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "BackwardLET", log); err != nil {
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

// ParseBackwardLET is a log parse operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseBackwardLET(log types.Log) (*Bridgel2sovereignchainBackwardLET, error) {
	event := new(Bridgel2sovereignchainBackwardLET)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "BackwardLET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// Bridgel2sovereignchainDetailedClaimEventIterator is returned from FilterDetailedClaimEvent and is used to iterate over the raw logs and unpacked data for DetailedClaimEvent events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainDetailedClaimEventIterator struct {
	Event *Bridgel2sovereignchainDetailedClaimEvent // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainDetailedClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainDetailedClaimEvent)
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
		it.Event = new(Bridgel2sovereignchainDetailedClaimEvent)
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
func (it *Bridgel2sovereignchainDetailedClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainDetailedClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainDetailedClaimEvent represents a DetailedClaimEvent event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainDetailedClaimEvent struct {
	SmtProofLocalExitRoot  [32][32]byte
	SmtProofRollupExitRoot [32][32]byte
	GlobalIndex            *big.Int
	MainnetExitRoot        [32]byte
	RollupExitRoot         [32]byte
	OriginNetwork          uint32
	OriginTokenAddress     common.Address
	DestinationNetwork     uint32
	DestinationAddress     common.Address
	Amount                 *big.Int
	Metadata               []byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterDetailedClaimEvent is a free log retrieval operation binding the contract event 0x7fee792f28e46481b8f3ff0ce0983de269108f87dd43f7c7c10f7c922897a944.
//
// Solidity: event DetailedClaimEvent(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 indexed globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address indexed destinationAddress, uint256 amount, bytes metadata)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterDetailedClaimEvent(opts *bind.FilterOpts, globalIndex []*big.Int, destinationAddress []common.Address) (*Bridgel2sovereignchainDetailedClaimEventIterator, error) {

	var globalIndexRule []interface{}
	for _, globalIndexItem := range globalIndex {
		globalIndexRule = append(globalIndexRule, globalIndexItem)
	}

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "DetailedClaimEvent", globalIndexRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainDetailedClaimEventIterator{contract: _Bridgel2sovereignchain.contract, event: "DetailedClaimEvent", logs: logs, sub: sub}, nil
}

// WatchDetailedClaimEvent is a free log subscription operation binding the contract event 0x7fee792f28e46481b8f3ff0ce0983de269108f87dd43f7c7c10f7c922897a944.
//
// Solidity: event DetailedClaimEvent(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 indexed globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address indexed destinationAddress, uint256 amount, bytes metadata)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchDetailedClaimEvent(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainDetailedClaimEvent, globalIndex []*big.Int, destinationAddress []common.Address) (event.Subscription, error) {

	var globalIndexRule []interface{}
	for _, globalIndexItem := range globalIndex {
		globalIndexRule = append(globalIndexRule, globalIndexItem)
	}

	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "DetailedClaimEvent", globalIndexRule, destinationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainDetailedClaimEvent)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "DetailedClaimEvent", log); err != nil {
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

// ParseDetailedClaimEvent is a log parse operation binding the contract event 0x7fee792f28e46481b8f3ff0ce0983de269108f87dd43f7c7c10f7c922897a944.
//
// Solidity: event DetailedClaimEvent(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 indexed globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address indexed destinationAddress, uint256 amount, bytes metadata)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseDetailedClaimEvent(log types.Log) (*Bridgel2sovereignchainDetailedClaimEvent, error) {
	event := new(Bridgel2sovereignchainDetailedClaimEvent)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "DetailedClaimEvent", log); err != nil {
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

// Bridgel2sovereignchainForwardLETIterator is returned from FilterForwardLET and is used to iterate over the raw logs and unpacked data for ForwardLET events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainForwardLETIterator struct {
	Event *Bridgel2sovereignchainForwardLET // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainForwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainForwardLET)
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
		it.Event = new(Bridgel2sovereignchainForwardLET)
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
func (it *Bridgel2sovereignchainForwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainForwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainForwardLET represents a ForwardLET event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainForwardLET struct {
	PreviousDepositCount *big.Int
	PreviousRoot         [32]byte
	NewDepositCount      *big.Int
	NewRoot              [32]byte
	NewLeaves            []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterForwardLET is a free log retrieval operation binding the contract event 0xe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c8.
//
// Solidity: event ForwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot, bytes newLeaves)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterForwardLET(opts *bind.FilterOpts) (*Bridgel2sovereignchainForwardLETIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainForwardLETIterator{contract: _Bridgel2sovereignchain.contract, event: "ForwardLET", logs: logs, sub: sub}, nil
}

// WatchForwardLET is a free log subscription operation binding the contract event 0xe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c8.
//
// Solidity: event ForwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot, bytes newLeaves)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchForwardLET(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainForwardLET) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainForwardLET)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "ForwardLET", log); err != nil {
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

// ParseForwardLET is a log parse operation binding the contract event 0xe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c8.
//
// Solidity: event ForwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot, bytes newLeaves)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseForwardLET(log types.Log) (*Bridgel2sovereignchainForwardLET, error) {
	event := new(Bridgel2sovereignchainForwardLET)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "ForwardLET", log); err != nil {
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

// Bridgel2sovereignchainSetClaimIterator is returned from FilterSetClaim and is used to iterate over the raw logs and unpacked data for SetClaim events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetClaimIterator struct {
	Event *Bridgel2sovereignchainSetClaim // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetClaim)
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
		it.Event = new(Bridgel2sovereignchainSetClaim)
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
func (it *Bridgel2sovereignchainSetClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetClaim represents a SetClaim event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetClaim struct {
	LeafIndex     uint32
	SourceNetwork uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetClaim is a free log retrieval operation binding the contract event 0xab48452e173be3f004b3208cc344b14cc7f8d977fc1ca28171286a7abb5d7570.
//
// Solidity: event SetClaim(uint32 leafIndex, uint32 sourceNetwork)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetClaim(opts *bind.FilterOpts) (*Bridgel2sovereignchainSetClaimIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetClaimIterator{contract: _Bridgel2sovereignchain.contract, event: "SetClaim", logs: logs, sub: sub}, nil
}

// WatchSetClaim is a free log subscription operation binding the contract event 0xab48452e173be3f004b3208cc344b14cc7f8d977fc1ca28171286a7abb5d7570.
//
// Solidity: event SetClaim(uint32 leafIndex, uint32 sourceNetwork)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetClaim(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetClaim) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetClaim)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetClaim", log); err != nil {
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

// ParseSetClaim is a log parse operation binding the contract event 0xab48452e173be3f004b3208cc344b14cc7f8d977fc1ca28171286a7abb5d7570.
//
// Solidity: event SetClaim(uint32 leafIndex, uint32 sourceNetwork)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetClaim(log types.Log) (*Bridgel2sovereignchainSetClaim, error) {
	event := new(Bridgel2sovereignchainSetClaim)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainSetLocalBalanceTreeIterator is returned from FilterSetLocalBalanceTree and is used to iterate over the raw logs and unpacked data for SetLocalBalanceTree events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetLocalBalanceTreeIterator struct {
	Event *Bridgel2sovereignchainSetLocalBalanceTree // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainSetLocalBalanceTreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainSetLocalBalanceTree)
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
		it.Event = new(Bridgel2sovereignchainSetLocalBalanceTree)
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
func (it *Bridgel2sovereignchainSetLocalBalanceTreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainSetLocalBalanceTreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainSetLocalBalanceTree represents a SetLocalBalanceTree event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainSetLocalBalanceTree struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
	NewAmount          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetLocalBalanceTree is a free log retrieval operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterSetLocalBalanceTree(opts *bind.FilterOpts, originNetwork []uint32, originTokenAddress []common.Address) (*Bridgel2sovereignchainSetLocalBalanceTreeIterator, error) {

	var originNetworkRule []interface{}
	for _, originNetworkItem := range originNetwork {
		originNetworkRule = append(originNetworkRule, originNetworkItem)
	}
	var originTokenAddressRule []interface{}
	for _, originTokenAddressItem := range originTokenAddress {
		originTokenAddressRule = append(originTokenAddressRule, originTokenAddressItem)
	}

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "SetLocalBalanceTree", originNetworkRule, originTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainSetLocalBalanceTreeIterator{contract: _Bridgel2sovereignchain.contract, event: "SetLocalBalanceTree", logs: logs, sub: sub}, nil
}

// WatchSetLocalBalanceTree is a free log subscription operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchSetLocalBalanceTree(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainSetLocalBalanceTree, originNetwork []uint32, originTokenAddress []common.Address) (event.Subscription, error) {

	var originNetworkRule []interface{}
	for _, originNetworkItem := range originNetwork {
		originNetworkRule = append(originNetworkRule, originNetworkItem)
	}
	var originTokenAddressRule []interface{}
	for _, originTokenAddressItem := range originTokenAddress {
		originTokenAddressRule = append(originTokenAddressRule, originTokenAddressItem)
	}

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "SetLocalBalanceTree", originNetworkRule, originTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainSetLocalBalanceTree)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetLocalBalanceTree", log); err != nil {
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

// ParseSetLocalBalanceTree is a log parse operation binding the contract event 0x73a03a6d9efc14b9f22a3af967e98580549eb76b1113b6a09a57ce1dae36ecd2.
//
// Solidity: event SetLocalBalanceTree(uint32 indexed originNetwork, address indexed originTokenAddress, uint256 newAmount)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseSetLocalBalanceTree(log types.Log) (*Bridgel2sovereignchainSetLocalBalanceTree, error) {
	event := new(Bridgel2sovereignchainSetLocalBalanceTree)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "SetLocalBalanceTree", log); err != nil {
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

// Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator is returned from FilterTransferEmergencyBridgePauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgePauserRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator struct {
	Event *Bridgel2sovereignchainTransferEmergencyBridgePauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainTransferEmergencyBridgePauserRole)
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
		it.Event = new(Bridgel2sovereignchainTransferEmergencyBridgePauserRole)
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
func (it *Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainTransferEmergencyBridgePauserRole represents a TransferEmergencyBridgePauserRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferEmergencyBridgePauserRole struct {
	CurrentEmergencyBridgePauser common.Address
	NewEmergencyBridgePauser     common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgePauserRole is a free log retrieval operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterTransferEmergencyBridgePauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainTransferEmergencyBridgePauserRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "TransferEmergencyBridgePauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgePauserRole is a free log subscription operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchTransferEmergencyBridgePauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainTransferEmergencyBridgePauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "TransferEmergencyBridgePauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainTransferEmergencyBridgePauserRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
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

// ParseTransferEmergencyBridgePauserRole is a log parse operation binding the contract event 0xb27de219766f47b82684842855ba6130b6dbf288ac66d1c3509e7bf17f4e925a.
//
// Solidity: event TransferEmergencyBridgePauserRole(address currentEmergencyBridgePauser, address newEmergencyBridgePauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseTransferEmergencyBridgePauserRole(log types.Log) (*Bridgel2sovereignchainTransferEmergencyBridgePauserRole, error) {
	event := new(Bridgel2sovereignchainTransferEmergencyBridgePauserRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferEmergencyBridgePauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator is returned from FilterTransferEmergencyBridgeUnpauserRole and is used to iterate over the raw logs and unpacked data for TransferEmergencyBridgeUnpauserRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator struct {
	Event *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole)
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
		it.Event = new(Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole)
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
func (it *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole represents a TransferEmergencyBridgeUnpauserRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole struct {
	CurrentEmergencyBridgeUnpauser common.Address
	NewEmergencyBridgeUnpauser     common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterTransferEmergencyBridgeUnpauserRole is a free log retrieval operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterTransferEmergencyBridgeUnpauserRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "TransferEmergencyBridgeUnpauserRole", logs: logs, sub: sub}, nil
}

// WatchTransferEmergencyBridgeUnpauserRole is a free log subscription operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchTransferEmergencyBridgeUnpauserRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "TransferEmergencyBridgeUnpauserRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
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

// ParseTransferEmergencyBridgeUnpauserRole is a log parse operation binding the contract event 0xf01a62a06940517bbc898dec8c75794b9feabcd2d263c8de823b36dbbeb8779b.
//
// Solidity: event TransferEmergencyBridgeUnpauserRole(address currentEmergencyBridgeUnpauser, address newEmergencyBridgeUnpauser)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseTransferEmergencyBridgeUnpauserRole(log types.Log) (*Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole, error) {
	event := new(Bridgel2sovereignchainTransferEmergencyBridgeUnpauserRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferEmergencyBridgeUnpauserRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator is returned from FilterTransferProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for TransferProxiedTokensManagerRole events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator struct {
	Event *Bridgel2sovereignchainTransferProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainTransferProxiedTokensManagerRole)
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
		it.Event = new(Bridgel2sovereignchainTransferProxiedTokensManagerRole)
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
func (it *Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainTransferProxiedTokensManagerRole represents a TransferProxiedTokensManagerRole event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainTransferProxiedTokensManagerRole struct {
	CurrentProxiedTokensManager common.Address
	NewProxiedTokensManager     common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTransferProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterTransferProxiedTokensManagerRole(opts *bind.FilterOpts) (*Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainTransferProxiedTokensManagerRoleIterator{contract: _Bridgel2sovereignchain.contract, event: "TransferProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferProxiedTokensManagerRole is a free log subscription operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchTransferProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainTransferProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainTransferProxiedTokensManagerRole)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
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

// ParseTransferProxiedTokensManagerRole is a log parse operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseTransferProxiedTokensManagerRole(log types.Log) (*Bridgel2sovereignchainTransferProxiedTokensManagerRole, error) {
	event := new(Bridgel2sovereignchainTransferProxiedTokensManagerRole)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator is returned from FilterUpdatedClaimedGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedClaimedGlobalIndexHashChain events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator struct {
	Event *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain)
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
		it.Event = new(Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain)
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
func (it *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain represents a UpdatedClaimedGlobalIndexHashChain event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain struct {
	ClaimedGlobalIndex             [32]byte
	NewClaimedGlobalIndexHashChain [32]byte
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedClaimedGlobalIndexHashChain is a free log retrieval operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterUpdatedClaimedGlobalIndexHashChain(opts *bind.FilterOpts) (*Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChainIterator{contract: _Bridgel2sovereignchain.contract, event: "UpdatedClaimedGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedClaimedGlobalIndexHashChain is a free log subscription operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchUpdatedClaimedGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "UpdatedClaimedGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
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

// ParseUpdatedClaimedGlobalIndexHashChain is a log parse operation binding the contract event 0x3e5936f910a78eb5181813a939c8d4c3e4d85f87943f659380d82ac6221b0e92.
//
// Solidity: event UpdatedClaimedGlobalIndexHashChain(bytes32 claimedGlobalIndex, bytes32 newClaimedGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseUpdatedClaimedGlobalIndexHashChain(log types.Log) (*Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain, error) {
	event := new(Bridgel2sovereignchainUpdatedClaimedGlobalIndexHashChain)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UpdatedClaimedGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator is returned from FilterUpdatedUnsetGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedUnsetGlobalIndexHashChain events raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator struct {
	Event *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain)
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
		it.Event = new(Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain)
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
func (it *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain represents a UpdatedUnsetGlobalIndexHashChain event raised by the Bridgel2sovereignchain contract.
type Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain struct {
	UnsetGlobalIndex             [32]byte
	NewUnsetGlobalIndexHashChain [32]byte
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedUnsetGlobalIndexHashChain is a free log retrieval operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) FilterUpdatedUnsetGlobalIndexHashChain(opts *bind.FilterOpts) (*Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.FilterLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChainIterator{contract: _Bridgel2sovereignchain.contract, event: "UpdatedUnsetGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedUnsetGlobalIndexHashChain is a free log subscription operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) WatchUpdatedUnsetGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Bridgel2sovereignchain.contract.WatchLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain)
				if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
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

// ParseUpdatedUnsetGlobalIndexHashChain is a log parse operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Bridgel2sovereignchain *Bridgel2sovereignchainFilterer) ParseUpdatedUnsetGlobalIndexHashChain(log types.Log) (*Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain, error) {
	event := new(Bridgel2sovereignchainUpdatedUnsetGlobalIndexHashChain)
	if err := _Bridgel2sovereignchain.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
