// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonzkevmbridge

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

// AgglayerBridgeLeafData is an auto generated low-level Go binding around an user-defined struct.
type AgglayerBridgeLeafData struct {
	LeafType           uint8
	OriginNetwork      uint32
	OriginAddress      common.Address
	DestinationNetwork uint32
	DestinationAddress common.Address
	Amount             *big.Int
	Metadata           []byte
}

// PolygonzkevmbridgeMetaData contains all meta data concerning the Polygonzkevmbridge contract.
var PolygonzkevmbridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountDoesNotMatchMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BridgeAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DestinationNetworkInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedProxyDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGlobalIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeFunction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidProxyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSmtProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyAdmin\",\"type\":\"address\"}],\"name\":\"InvalidZeroProxyAdminOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTokenIsEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewDepositCountExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoValueInMessagesOnGasTokenNetworks\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroValueForUnusedFrontier\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotValidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProxiedTokensManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubtreeFrontierMismatch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"AcceptProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"BackwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"depositCount\",\"type\":\"uint32\"}],\"name\":\"BridgeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmergencyStateDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"previousRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newLeaves\",\"type\":\"bytes\"}],\"name\":\"ForwardLET\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"NewWrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"globalIndex\",\"type\":\"bytes32\"}],\"name\":\"SetClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentProxiedTokensManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"TransferProxiedTokensManagerRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"unsetGlobalIndex\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newUnsetGlobalIndexHashChain\",\"type\":\"bytes32\"}],\"name\":\"UpdatedUnsetGlobalIndexHashChain\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INIT_BYTECODE_TRANSPARENT_PROXY\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETHToken\",\"outputs\":[{\"internalType\":\"contractITokenWrappedBridgeUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDepositCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[32]\",\"name\":\"newFrontier\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32\",\"name\":\"nextLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[32]\",\"name\":\"proof\",\"type\":\"bytes32[32]\"}],\"name\":\"backwardLET\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeAsset\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeLib\",\"outputs\":[{\"internalType\":\"contractBridgeLib\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountWETH\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceUpdateGlobalExitRoot\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"bridgeMessageWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofLocalExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"bytes32[32]\",\"name\":\"smtProofRollupExitRoot\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"globalIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"mainnetExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rollupExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimedBitMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"computeTokenProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateEmergencyState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"leafType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"destinationNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"}],\"internalType\":\"structAgglayerBridge.LeafData[]\",\"name\":\"newLeaves\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"expectedLER\",\"type\":\"bytes32\"}],\"name\":\"forwardLET\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenMetadata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBridgeImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"contractIBaseLegacyAgglayerGER\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_polygonRollupManager\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"sourceBridgeNetwork\",\"type\":\"uint32\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEmergencyState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedDepositCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkID\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingProxiedTokensManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polygonRollupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"setMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfoToWrappedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProxiedTokensManager\",\"type\":\"address\"}],\"name\":\"transferProxiedTokensManagerRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"globalIndexes\",\"type\":\"uint256[]\"}],\"name\":\"unsetMultipleClaims\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGlobalExitRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrappedTokenToTokenInfo\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"originNetwork\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161001c9061013e565b604051809103905ff080158015610035573d5f5f3e3d5ffd5b506001600160a01b031660a05260405161004e9061014b565b604051809103905ff080158015610067573d5f5f3e3d5ffd5b506001600160a01b031660805261007c610081565b610158565b5f54610100900460ff16156100ec5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101561013c575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6115988061249d83390190565b611bd880613a3583390190565b60805160a0516123166101875f395f6103c001525f81816104ad01528181610c770152610ce801526123165ff3fe608060405260043610610219575f3560e01c80638bd309c311610129578063cd586579116100a8578063f0e708081161006d578063f0e70808146106d4578063f214e161146106f3578063f5efcd7914610712578063f811bff714610731578063fd7640e8146107b0575f5ffd5b8063cd5865791461063d578063d02103ca14610650578063dbc1697614610676578063ece93c6f1461068a578063ee25560b146106a9575f5ffd5b8063be5831c7116100ee578063be5831c7146105a9578063c00f14ab146105cc578063c514f24e146105eb578063cc461632146105ff578063ccaa2d111461061e575f5ffd5b80638bd309c3146105175780638c668f1c146105365780638ed7e3f21461054a578063b8b284d014610569578063bab161bf14610588575f5ffd5b806338b8fbbb116101b557806354fd4d501161017a57806354fd4d501461045a5780635ca1e165146104885780636f0bc3da1461049c57806379e2cf97146104cf57806381b1c174146104e3575f5ffd5b806338b8fbbb146103955780633b2fee9a146103b25780633c351e10146103e45780633cbc795b146104035780634b2f336d1461043b575f5ffd5b8063136a2c601461021d57806315064c961461023e5780631c2082291461026c5780632072f6c51461028b57806322e95f2c1461029f578063240ff378146102d657806327aef4e8146102e95780632dfdf0b51461030a578063318aee3d1461032d575b5f5ffd5b348015610228575f5ffd5b5061023c61023736600461172e565b6107cf565b005b348015610249575f5ffd5b506068546102579060ff1681565b60405190151581526020015b60405180910390f35b348015610277575f5ffd5b5061023c61028636600461172e565b610858565b348015610296575f5ffd5b5061023c6108d0565b3480156102aa575f5ffd5b506102be6102b9366004611812565b610905565b6040516001600160a01b039091168152602001610263565b61023c6102e436600461189a565b610953565b3480156102f4575f5ffd5b506102fd61098c565b6040516102639190611959565b348015610315575f5ffd5b5061031f60535481565b604051908152602001610263565b348015610338575f5ffd5b50610371610347366004611972565b606b6020525f908152604090205463ffffffff81169064010000000090046001600160a01b031682565b6040805163ffffffff90931683526001600160a01b03909116602083015201610263565b3480156103a0575f5ffd5b506070546001600160a01b03166102be565b3480156103bd575f5ffd5b507f00000000000000000000000000000000000000000000000000000000000000006102be565b3480156103ef575f5ffd5b50606d546102be906001600160a01b031681565b34801561040e575f5ffd5b50606d5461042690600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610263565b348015610446575f5ffd5b50606f546102be906001600160a01b031681565b348015610465575f5ffd5b50604080518082019091526006815265076312e312e360d41b60208201526102fd565b348015610493575f5ffd5b5061031f610a18565b3480156104a7575f5ffd5b506102be7f000000000000000000000000000000000000000000000000000000000000000081565b3480156104da575f5ffd5b5061023c610a97565b3480156104ee575f5ffd5b506102be6104fd36600461198d565b606a6020525f90815260409020546001600160a01b031681565b348015610522575f5ffd5b5061023c610531366004611972565b610ab8565b348015610541575f5ffd5b5061023c610b46565b348015610555575f5ffd5b50606c546102be906001600160a01b031681565b348015610574575f5ffd5b5061023c6105833660046119a4565b610bd7565b348015610593575f5ffd5b5060685461042690610100900463ffffffff1681565b3480156105b4575f5ffd5b5060685461042690600160c81b900463ffffffff1681565b3480156105d7575f5ffd5b506102fd6105e6366004611972565b610c55565b3480156105f6575f5ffd5b506102fd610ce4565b34801561060a575f5ffd5b50610257610619366004611a1f565b610d6d565b348015610629575f5ffd5b5061023c610638366004611a61565b610df5565b61023c61064b366004611b3f565b610e81565b34801561065b575f5ffd5b506068546102be90600160281b90046001600160a01b031681565b348015610681575f5ffd5b5061023c610f5e565b348015610695575f5ffd5b506071546102be906001600160a01b031681565b3480156106b4575f5ffd5b5061031f6106c336600461198d565b60696020525f908152604090205481565b3480156106df575f5ffd5b5061023c6106ee366004611bcc565b610f91565b3480156106fe575f5ffd5b506102be61070d366004611812565b611038565b34801561071d575f5ffd5b5061023c61072c366004611a61565b6110dd565b34801561073c575f5ffd5b5061023c61074b366004611c65565b5060688054610100600160c81b03191661010063ffffffff979097169690960265010000000000600160c81b03191695909517600160281b6001600160a01b039384160217909455606c80546001600160a01b03191694909116939093179092555050565b3480156107bb575f5ffd5b5061023c6107ca366004611d34565b611138565b5f5b8151811015610854575f8282815181106107ed576107ed611d7c565b602002602001015190505f5f61080283611215565b604080518781526020810188905293955090935085927fc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b0392500160405180910390a15050600190920191506107d19050565b5050565b5f5b8151811015610854575f82828151811061087657610876611d7c565b602002602001015190505f5f61088b83611215565b60405186815292945092507f8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c91602001905060405180910390a150505060010161085a565b606c546001600160a01b031633146108fb57604051631736745960e31b815260040160405180910390fd5b6109036112b9565b565b5f606a5f848460405160200161091c929190611d90565b60408051601f198184030181529181528151602092830120835290820192909252015f20546001600160a01b031690505b92915050565b60685460ff161561097757604051630bc011ff60e21b815260040160405180910390fd5b610985858534868686611314565b5050505050565b606e805461099990611dba565b80601f01602080910402602001604051908101604052809291908181526020018280546109c590611dba565b8015610a105780601f106109e757610100808354040283529160200191610a10565b820191905f5260205f20905b8154815290600101906020018083116109f357829003601f168201915b505050505081565b6053545f90819081805b6020811015610a8e578083901c600116600103610a6757610a6060338260208110610a4f57610a4f611d7c565b0154855f9182526020526040902090565b9350610a77565b5f84815260208390526040902093505b5f8281526020839052604090209150600101610a22565b50919392505050565b605354606854600160c81b900463ffffffff161015610903576109036113f0565b6070546001600160a01b03163314610ae357604051630866750360e01b815260040160405180910390fd5b607180546001600160a01b0319166001600160a01b038381169182179092556070546040805191909316815260208101919091527f0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e291015b60405180910390a150565b6071546001600160a01b03163314610b7157604051630b59ef2760e21b815260040160405180910390fd5b60708054607180546001600160a01b038082166001600160a01b0319808616821790965594909116909155604080519190921680825260208201939093527fa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f29101610b3b565b60685460ff1615610bfb57604051630bc011ff60e21b815260040160405180910390fd5b606f546001600160a01b0316610c245760405163dde3cda760e01b815260040160405180910390fd5b606f545f90610c3c906001600160a01b031686611484565b9050610c4c878783878787611314565b50505050505050565b60405163c00f14ab60e01b81526001600160a01b0382811660048301526060917f00000000000000000000000000000000000000000000000000000000000000009091169063c00f14ab906024015f60405180830381865afa158015610cbd573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261094d9190810190611df2565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c514f24e6040518163ffffffff1660e01b81526004015f60405180830381865afa158015610d41573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d689190810190611df2565b905090565b6068545f908190610100900463ffffffff16158015610d92575063ffffffff83166001145b15610da4575063ffffffff8316610dcc565b610db964010000000063ffffffff8516611e77565b610dc99063ffffffff8616611e8e565b90505b600881901c5f90815260696020526040902054600160ff9092169190911b908116149392505050565b7f490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a3987875f8585604051610e2c959493929190611ec9565b60405180910390a17f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d8a88888787604051610e6b959493929190611f01565b60405180910390a1505050505050505050505050565b60685460ff1615610ea557604051630bc011ff60e21b815260040160405180910390fd5b610ead6114e8565b60685463ffffffff610100909104811690881603610ede576040516302caf51760e11b815260040160405180910390fd5b5f5f60605f8890507f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b5f84868e8e8688605354604051610f25989796959493929190611f33565b60405180910390a1610f435f84868e8e868880519060200120611545565b8615610f5157610f516113f0565b50505050610c4c60018055565b606c546001600160a01b03163314610f8957604051631736745960e31b815260040160405180910390fd5b6109036115cf565b5f829003610f9d575f5ffd5b6053545f610fa9610a18565b90508484905060535f828254610fbf9190611e8e565b909155505f9050610fce610a18565b90507fe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c88383605354848a8a60405160200161100a929190611fe0565b60408051601f198184030181529082905261102895949392916120f0565b60405180910390a1505050505050565b5f5f838360405160200161104d929190611d90565b6040516020818303038152906040528051906020012090505f60ff60f81b3083611075610ce4565b80516020918201206040516110bc95949392016001600160f81b031994909416845260609290921b6001600160601b03191660018401526015830152603582015260550190565b60408051808303601f19018152919052805160209091012095945050505050565b60685460ff161561110157604051630bc011ff60e21b815260040160405180910390fd5b7f1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d8a88888787604051610e6b959493929190611f01565b605354841015611146575f5ffd5b6053545f611152610a18565b605387905590505f611162610a18565b6068546040516333d6247d60e01b815260048101839052919250600160281b90046001600160a01b0316906333d6247d906024015f604051808303815f87803b1580156111ad575f5ffd5b505af11580156111bf573d5f5f3e3d5ffd5b505060408051868152602081018690529081018a9052606081018490527f2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f1495559250608001905060405180910390a150505050505050565b805f80600160401b83161561126157505f9050808361123e63ffffffff8516600160401b611e8e565b1461125c5760405163071389e960e01b815260040160405180910390fd5b6112b2565b602084901c915061127382600161211a565b90508361129463ffffffff851667ffffffff00000000602086901b16611e8e565b146112b25760405163071389e960e01b815260040160405180910390fd5b9193909250565b60685460ff16156112dd57604051630bc011ff60e21b815260040160405180910390fd5b6068805460ff191660011790556040517f2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497905f90a1565b60685463ffffffff610100909104811690871603611345576040516302caf51760e11b815260040160405180910390fd5b7f501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b6001606860019054906101000a900463ffffffff1633898989888860535460405161139999989796959493929190612136565b60405180910390a16113da6001606860019054906101000a900463ffffffff163389898988886040516113cd9291906121a4565b6040518091039020611545565b82156113e8576113e86113f0565b505050505050565b6053546068805463ffffffff909216600160c81b0263ffffffff60c81b1990921691909117908190556001600160a01b03600160281b909104166333d6247d611437610a18565b6040518263ffffffff1660e01b815260040161145591815260200190565b5f604051808303815f87803b15801561146c575f5ffd5b505af115801561147e573d5f5f3e3d5ffd5b50505050565b604051632770a7eb60e21b8152336004820152602481018290525f906001600160a01b03841690639dc29fac906044015f604051808303815f87803b1580156114cb575f5ffd5b505af11580156114dd573d5f5f3e3d5ffd5b509395945050505050565b60026001540361153e5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640160405180910390fd5b6002600155565b604080516001600160f81b031960f88a901b166020808301919091526001600160e01b031960e08a811b821660218501526001600160601b031960608b811b82166025870152918a901b909216603985015287901b16603d8301526051820185905260718083018590528351808403909101815260919092019092528051910120610c4c90611626565b60685460ff166115f257604051635386698160e01b815260040160405180910390fd5b6068805460ff191690556040517f1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3905f90a1565b80600161163560206002612296565b61163f91906122a1565b60535410611660576040516377ae67b360e11b815260040160405180910390fd5b5f60535f815461166f906122b4565b918290555090505f5b60208110156116dc578082901c6001166001036116ab5782603382602081106116a3576116a3611d7c565b015550505050565b6116d2603382602081106116c1576116c1611d7c565b0154845f9182526020526040902090565b9250600101611678565b506116e56122cc565b505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715611726576117266116ea565b604052919050565b5f6020828403121561173e575f5ffd5b81356001600160401b03811115611753575f5ffd5b8201601f81018413611763575f5ffd5b80356001600160401b0381111561177c5761177c6116ea565b8060051b61178c602082016116fe565b918252602081840181019290810190878411156117a7575f5ffd5b6020850194505b838510156117cd578435808352602095860195909350909101906117ae565b979650505050505050565b803563ffffffff811681146117eb575f5ffd5b919050565b6001600160a01b0381168114611804575f5ffd5b50565b80356117eb816117f0565b5f5f60408385031215611823575f5ffd5b61182c836117d8565b9150602083013561183c816117f0565b809150509250929050565b803580151581146117eb575f5ffd5b5f5f83601f840112611866575f5ffd5b5081356001600160401b0381111561187c575f5ffd5b602083019150836020828501011115611893575f5ffd5b9250929050565b5f5f5f5f5f608086880312156118ae575f5ffd5b6118b7866117d8565b945060208601356118c7816117f0565b93506118d560408701611847565b925060608601356001600160401b038111156118ef575f5ffd5b6118fb88828901611856565b969995985093965092949392505050565b5f5b8381101561192657818101518382015260200161190e565b50505f910152565b5f815180845261194581602086016020860161190c565b601f01601f19169290920160200192915050565b602081525f61196b602083018461192e565b9392505050565b5f60208284031215611982575f5ffd5b813561196b816117f0565b5f6020828403121561199d575f5ffd5b5035919050565b5f5f5f5f5f5f60a087890312156119b9575f5ffd5b6119c2876117d8565b955060208701356119d2816117f0565b9450604087013593506119e760608801611847565b925060808701356001600160401b03811115611a01575f5ffd5b611a0d89828a01611856565b979a9699509497509295939492505050565b5f5f60408385031215611a30575f5ffd5b611a39836117d8565b9150611a47602084016117d8565b90509250929050565b80610400810183101561094d575f5ffd5b5f5f5f5f5f5f5f5f5f5f5f5f6109208d8f031215611a7d575f5ffd5b611a878e8e611a50565b9b50611a978e6104008f01611a50565b9a506108008d013599506108208d013598506108408d01359750611abe6108608e016117d8565b9650611ace6108808e01356117f0565b6108808d01359550611ae36108a08e016117d8565b94506108c08d0135611af4816117f0565b93506108e08d013592506001600160401b036109008e01351115611b16575f5ffd5b611b278e6109008f01358f01611856565b81935080925050509295989b509295989b509295989b565b5f5f5f5f5f5f5f60c0888a031215611b55575f5ffd5b611b5e886117d8565b96506020880135611b6e816117f0565b9550604088013594506060880135611b85816117f0565b9350611b9360808901611847565b925060a08801356001600160401b03811115611bad575f5ffd5b611bb98a828b01611856565b989b979a50959850939692959293505050565b5f5f5f60408486031215611bde575f5ffd5b83356001600160401b03811115611bf3575f5ffd5b8401601f81018613611c03575f5ffd5b80356001600160401b03811115611c18575f5ffd5b8660208260051b8401011115611c2c575f5ffd5b6020918201979096509401359392505050565b5f6001600160401b03821115611c5757611c576116ea565b50601f01601f191660200190565b5f5f5f5f5f5f60c08789031215611c7a575f5ffd5b611c83876117d8565b95506020870135611c93816117f0565b9450611ca1604088016117d8565b93506060870135611cb1816117f0565b92506080870135611cc1816117f0565b915060a08701356001600160401b03811115611cdb575f5ffd5b8701601f81018913611ceb575f5ffd5b8035611cfe611cf982611c3f565b6116fe565b8181528a6020838501011115611d12575f5ffd5b816020840160208301375f602083830101528093505050509295509295509295565b5f5f5f5f6108408587031215611d48575f5ffd5b84359350611d598660208701611a50565b92506104208501359150611d71866104408701611a50565b905092959194509250565b634e487b7160e01b5f52603260045260245ffd5b60e09290921b6001600160e01b031916825260601b6001600160601b031916600482015260180190565b600181811c90821680611dce57607f821691505b602082108103611dec57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215611e02575f5ffd5b81516001600160401b03811115611e17575f5ffd5b8201601f81018413611e27575f5ffd5b8051611e35611cf982611c3f565b818152856020838501011115611e49575f5ffd5b611e5a82602083016020860161190c565b95945050505050565b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761094d5761094d611e63565b8082018082111561094d5761094d611e63565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b63ffffffff861681526001600160a01b038581166020830152841660408201526080606082018190525f906117cd9083018486611ea1565b94855263ffffffff9390931660208501526001600160a01b039182166040850152166060830152608082015260a00190565b60ff8916815263ffffffff88811660208301526001600160a01b03888116604084015290871660608301528516608082015260a0810184905261010060c082018190525f90611f849083018561192e565b905063ffffffff831660e08301529998505050505050505050565b5f5f8335601e19843603018112611fb4575f5ffd5b83016020810192503590506001600160401b03811115611fd2575f5ffd5b803603821315611893575f5ffd5b602080825281018290525f6040600584901b83018101908301858360de1936839003015b878210156120e357868503603f190184528235818112612022575f5ffd5b8901803560ff8116808214612035575f5ffd5b875250612044602082016117d8565b63ffffffff16602087015261205b60408201611807565b6001600160a01b03166040870152612075606082016117d8565b63ffffffff16606087015261208c60808201611807565b6001600160a01b0316608087015260a081810135908701526120b160c0820182611f9f565b915060e060c08801526120c860e088018383611ea1565b96505050602083019250602084019350600182019150612004565b5092979650505050505050565b85815284602082015283604082015282606082015260a060808201525f6117cd60a083018461192e565b63ffffffff818116838216019081111561094d5761094d611e63565b60ff8a16815263ffffffff89811660208301526001600160a01b03898116604084015290881660608301528616608082015260a0810185905261010060c082018190525f906121889083018587611ea1565b905063ffffffff831660e08301529a9950505050505050505050565b818382375f9101908152919050565b6001815b60018411156121ee578085048111156121d2576121d2611e63565b60018416156121e057908102905b60019390931c9280026121b7565b935093915050565b5f826122045750600161094d565b8161221057505f61094d565b816001811461222657600281146122305761224c565b600191505061094d565b60ff84111561224157612241611e63565b50506001821b61094d565b5060208310610133831016604e8410600b841016171561226f575081810a61094d565b61227b5f1984846121b3565b805f190482111561228e5761228e611e63565b029392505050565b5f61196b83836121f6565b8181038181111561094d5761094d611e63565b5f600182016122c5576122c5611e63565b5060010190565b634e487b7160e01b5f52600160045260245ffdfea2646970667358221220f2bb32bf0634523ff1c6bef4502710805fea68b7dad2519be738fdf5e80e549464736f6c634300081c00336080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6114c2806100d65f395ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c806370a082311161008f57806370a08231146101a05780637ecebe00146101b357806384b0196e146101c657806395d89b41146101e15780639dc29fac146101e9578063a3c573eb146101fc578063a9059cbb14610211578063d505accf14610224578063dd62ed3e14610237575f5ffd5b806306fdde03146100ec578063095ea7b31461010a5780631624f6c61461012d57806318160ddd1461014257806323b872dd14610158578063313ce5671461016b5780633644e5151461018557806340c10f191461018d575b5f5ffd5b6100f461024a565b6040516101019190610ff3565b60405180910390f35b61011d610118366004611027565b6102e8565b6040519015158152602001610101565b61014061013b3660046110fc565b610301565b005b61014a61043e565b604051908152602001610101565b61011d61016636600461116e565b610452565b610173610475565b60405160ff9091168152602001610101565b61014a610489565b61014061019b366004611027565b610497565b61014a6101ae3660046111a8565b6104e1565b61014a6101c13660046111a8565b61050a565b6101ce610514565b60405161010197969594939291906111c1565b6100f46105bd565b6101406101f7366004611027565b6105d9565b61020461061e565b6040516101019190611257565b61011d61021f366004611027565b61063d565b61014061023236600461126b565b61064a565b61014a6102453660046112d1565b61079f565b60605f6102556107d9565b905080600301805461026690611302565b80601f016020809104026020016040519081016040528092919081815260200182805461029290611302565b80156102dd5780601f106102b4576101008083540402835291602001916102dd565b820191905f5260205f20905b8154815290600101906020018083116102c057829003601f168201915b505050505091505090565b5f336102f58185856107fd565b60019150505b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156103455750825b90505f826001600160401b031660011480156103605750303b155b90508115801561036e575080155b1561038c5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156103b657845460ff60401b1916600160401b1785555b6103c0888861080a565b6103c988610820565b5f6103d261084e565b805433610100026001600160a81b031990911660ff8a161717905550831561043457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f5f6104486107d9565b6002015492915050565b5f3361045f858285610872565b61046a8585856108c2565b506001949350505050565b5f5f61047f61084e565b5460ff1692915050565b5f61049261091f565b905090565b5f6104a061084e565b805490915061010090046001600160a01b031633146104d2576040516338da3b1560e01b815260040160405180910390fd5b6104dc8383610928565b505050565b5f5f6104eb6107d9565b6001600160a01b039093165f9081526020939093525050604090205490565b5f6102fb8261095c565b5f6060805f5f5f60605f610526610966565b805490915015801561053a57506001810154155b6105835760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b61058b61098a565b6105936109a6565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b60605f6105c86107d9565b905080600401805461026690611302565b5f6105e261084e565b805490915061010090046001600160a01b03163314610614576040516338da3b1560e01b815260040160405180910390fd5b6104dc83836109b1565b5f5f61062861084e565b5461010090046001600160a01b031692915050565b5f336102f58185856108c2565b8342111561066e5760405163313c898160e11b81526004810185905260240161057a565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886106d88c6001600160a01b03165f9081527f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006020526040902080546001810190915590565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610732826109e5565b90505f61074182878787610a11565b9050896001600160a01b0316816001600160a01b031614610788576040516325c0072360e11b81526001600160a01b0380831660048301528b16602482015260440161057a565b6107938a8a8a6107fd565b50505050505050505050565b5f5f6107a96107d9565b6001600160a01b039485165f90815260019190910160209081526040808320959096168252939093525050205490565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0090565b6104dc8383836001610a3d565b610812610b1e565b61081c8282610b69565b5050565b610828610b1e565b61084b81604051806040016040528060018152602001603160f81b815250610b99565b50565b7f863b064fe9383d75d38f584f64f1aaba4520e9ebc98515fa15bdeae8c4274d0090565b5f61087d848461079f565b90505f1981146108bc57818110156108ae57828183604051637dc7a0d960e11b815260040161057a9392919061133a565b6108bc84848484035f610a3d565b50505050565b6001600160a01b0383166108eb575f604051634b637e8f60e11b815260040161057a9190611257565b6001600160a01b038216610914575f60405163ec442f0560e01b815260040161057a9190611257565b6104dc838383610bd8565b5f610492610cfb565b6001600160a01b038216610951575f60405163ec442f0560e01b815260040161057a9190611257565b61081c5f8383610bd8565b5f6102fb82610d6e565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10090565b60605f610995610966565b905080600201805461026690611302565b60605f610255610966565b6001600160a01b0382166109da575f604051634b637e8f60e11b815260040161057a9190611257565b61081c825f83610bd8565b5f6102fb6109f161091f565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f610a2188888888610d96565b925092509250610a318282610e54565b50909695505050505050565b5f610a466107d9565b90506001600160a01b038516610a71575f60405163e602df0560e01b815260040161057a9190611257565b6001600160a01b038416610a9a575f604051634a1406b160e11b815260040161057a9190611257565b6001600160a01b038086165f90815260018301602090815260408083209388168352929052208390558115610b1757836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92585604051610b0e91815260200190565b60405180910390a35b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610b6757604051631afcd79f60e31b815260040160405180910390fd5b565b610b71610b1e565b5f610b7a6107d9565b905060038101610b8a848261139f565b50600481016108bc838261139f565b610ba1610b1e565b5f610baa610966565b905060028101610bba848261139f565b5060038101610bc9838261139f565b505f8082556001909101555050565b5f610be16107d9565b90506001600160a01b038416610c0f5781816002015f828254610c049190611459565b90915550610c6c9050565b6001600160a01b0384165f9081526020829052604090205482811015610c4e5784818460405163391434e360e21b815260040161057a9392919061133a565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b038316610c8a576002810180548390039055610ca8565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ced91815260200190565b60405180910390a350505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f610d25610f0c565b610d2d610f71565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f807f5ab42ced628888259c08ac98db1eb0cf702fc1501344311d8b100cd1bfe4bb006104eb565b5f80806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03841115610dc557505f91506003905082610e4a565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610e16573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116610e4157505f925060019150829050610e4a565b92505f91508190505b9450945094915050565b5f826003811115610e6757610e67611478565b03610e70575050565b6001826003811115610e8457610e84611478565b03610ea25760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115610eb657610eb6611478565b03610ed75760405163fce698f760e01b81526004810182905260240161057a565b6003826003811115610eeb57610eeb611478565b0361081c576040516335e2f38360e21b81526004810182905260240161057a565b5f5f610f16610966565b90505f610f2161098a565b805190915015610f3957805160209091012092915050565b81548015610f48579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f610f7b610966565b90505f610f866109a6565b805190915015610f9e57805160209091012092915050565b60018201548015610f48579392505050565b5f81518084525f5b81811015610fd457602081850181015186830182015201610fb8565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f6110056020830184610fb0565b9392505050565b80356001600160a01b0381168114611022575f5ffd5b919050565b5f5f60408385031215611038575f5ffd5b6110418361100c565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112611072575f5ffd5b81356001600160401b0381111561108b5761108b61104f565b604051601f8201601f19908116603f011681016001600160401b03811182821017156110b9576110b961104f565b6040528181528382016020018510156110d0575f5ffd5b816020850160208301375f918101602001919091529392505050565b803560ff81168114611022575f5ffd5b5f5f5f6060848603121561110e575f5ffd5b83356001600160401b03811115611123575f5ffd5b61112f86828701611063565b93505060208401356001600160401b0381111561114a575f5ffd5b61115686828701611063565b925050611165604085016110ec565b90509250925092565b5f5f5f60608486031215611180575f5ffd5b6111898461100c565b92506111976020850161100c565b929592945050506040919091013590565b5f602082840312156111b8575f5ffd5b6110058261100c565b60ff60f81b8816815260e060208201525f6111df60e0830189610fb0565b82810360408401526111f18189610fb0565b606084018890526001600160a01b038716608085015260a0840186905283810360c0850152845180825260208087019350909101905f5b81811015611246578351835260209384019390920191600101611228565b50909b9a5050505050505050505050565b6001600160a01b0391909116815260200190565b5f5f5f5f5f5f5f60e0888a031215611281575f5ffd5b61128a8861100c565b96506112986020890161100c565b955060408801359450606088013593506112b4608089016110ec565b9699959850939692959460a0840135945060c09093013592915050565b5f5f604083850312156112e2575f5ffd5b6112eb8361100c565b91506112f96020840161100c565b90509250929050565b600181811c9082168061131657607f821691505b60208210810361133457634e487b7160e01b5f52602260045260245ffd5b50919050565b6001600160a01b039390931683526020830191909152604082015260600190565b601f8211156104dc57805f5260205f20601f840160051c810160208510156113805750805b601f840160051c820191505b81811015610b17575f815560010161138c565b81516001600160401b038111156113b8576113b861104f565b6113cc816113c68454611302565b8461135b565b6020601f8211600181146113fe575f83156113e75750848201515b5f19600385901b1c1916600184901b178455610b17565b5f84815260208120601f198516915b8281101561142d578785015182556020948501946001909201910161140d565b508482101561144a57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b808201808211156102fb57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea2646970667358221220734889424b5f04b57ae9c0f182a620f22fa0c90954fa75a81eadce83f886d8bc64736f6c634300081c00336080604052348015600e575f5ffd5b50611bbc8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c8063737ce34114610064578063a28fa4a31461008d578063be3dcf62146100b0578063c00f14ab146100d5578063c514f24e146100e8578063cf825e55146100f0575b5f5ffd5b6100776100723660046108de565b610103565b604051610084919061094d565b60405180910390f35b6100a061009b36600461095f565b6101c8565b6040519015158152602001610084565b6100c36100be3660046108de565b61056c565b60405160ff9091168152602001610084565b6100776100e33660046108de565b610620565b610077610665565b6100776100fe3660046108de565b610684565b60408051600481526024810182526020810180516001600160e01b03166395d89b4160e01b17905290516060915f9182916001600160a01b038616916101499190610a02565b5f60405180830381855afa9150503d805f8114610181576040519150601f19603f3d011682016040523d82523d5f602084013e610186565b606091505b5091509150816101b757604051806040016040528060098152602001681393d7d4d6535093d360ba1b8152506101c0565b6101c081610736565b949350505050565b5f806101d76004828789610a1d565b6101e091610a44565b9050632afa533160e01b6001600160e01b031982160161038f575f5f5f5f5f5f5f8c8c600490809261021493929190610a1d565b8101906102219190610a8a565b96509650965096509650965096508a6001600160a01b0316876001600160a01b0316146102615760405163912ecce760e01b815260040160405180910390fd5b896001600160a01b0316866001600160a01b0316146102935760405163750643af60e01b815260040160405180910390fd5b5f8e6001600160a01b031663d505accf60e01b898989898989896040516024016102ff97969594939291906001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252905161033d9190610a02565b5f604051808303815f865af19150503d805f8114610376576040519150601f19603f3d011682016040523d82523d5f602084013e61037b565b606091505b50909a506105639950505050505050505050565b631c0d143d60e21b6001600160e01b031982160161054a575f5f5f5f5f5f5f5f8d8d60049080926103c293929190610a1d565b8101906103cf9190610af6565b975097509750975097509750975097508b6001600160a01b0316886001600160a01b0316146104115760405163912ecce760e01b815260040160405180910390fd5b8a6001600160a01b0316876001600160a01b0316146104435760405163750643af60e01b815260040160405180910390fd5b5f8f6001600160a01b0316638fcbaf0c60e01b8a8a8a8a8a8a8a8a6040516024016104b99897969594939291906001600160a01b039889168152969097166020870152604086019490945260608501929092521515608084015260ff1660a083015260c082015260e08101919091526101000190565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516104f79190610a02565b5f604051808303815f865af19150503d805f8114610530576040519150601f19603f3d011682016040523d82523d5f602084013e610535565b606091505b50909b506105639a5050505050505050505050565b604051637141605d60e11b815260040160405180910390fd5b95945050505050565b60408051600481526024810182526020810180516001600160e01b031663313ce56760e01b17905290515f91829182916001600160a01b038616916105b19190610a02565b5f60405180830381855afa9150503d805f81146105e9576040519150601f19603f3d011682016040523d82523d5f602084013e6105ee565b606091505b5091509150818015610601575080516020145b61060c5760126101c0565b808060200190518101906101c09190610b78565b606061062b82610684565b61063483610103565b61063d8461056c565b60405160200161064f93929190610b93565b6040516020818303038152906040529050919050565b60405180610f000160405280610ec88152602001610cbf610ec8913981565b60408051600481526024810182526020810180516001600160e01b03166306fdde0360e01b17905290516060915f9182916001600160a01b038616916106ca9190610a02565b5f60405180830381855afa9150503d805f8114610702576040519150601f19603f3d011682016040523d82523d5f602084013e610707565b606091505b5091509150816101b757604051806040016040528060078152602001664e4f5f4e414d4560c81b8152506101c0565b6060604082511061075b57818060200190518101906107559190610bdf565b92915050565b8151602003610889575f5b602081108015610795575082818151811061078357610783610c86565b01602001516001600160f81b03191615155b156107ac57806107a481610c9a565b915050610766565b805f036107e35750506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b6020820152919050565b5f8167ffffffffffffffff8111156107fd576107fd610bcb565b6040519080825280601f01601f191660200182016040528015610827576020820181803683370190505b5090505f5b828110156108815784818151811061084657610846610c86565b602001015160f81c60f81b82828151811061086357610863610c86565b60200101906001600160f81b03191690815f1a90535060010161082c565b509392505050565b50506040805180820190915260128152714e4f545f56414c49445f454e434f44494e4760701b602082015290565b919050565b6001600160a01b03811681146108d0575f5ffd5b50565b80356108b7816108bc565b5f602082840312156108ee575f5ffd5b81356108f9816108bc565b9392505050565b5f5b8381101561091a578181015183820152602001610902565b50505f910152565b5f8151808452610939816020860160208601610900565b601f01601f19169290920160200192915050565b602081525f6108f96020830184610922565b5f5f5f5f5f60808688031215610973575f5ffd5b853561097e816108bc565b9450602086013567ffffffffffffffff811115610999575f5ffd5b8601601f810188136109a9575f5ffd5b803567ffffffffffffffff8111156109bf575f5ffd5b8860208284010111156109d0575f5ffd5b6020919091019450925060408601356109e8816108bc565b91506109f6606087016108d3565b90509295509295909350565b5f8251610a13818460208701610900565b9190910192915050565b5f5f85851115610a2b575f5ffd5b83861115610a37575f5ffd5b5050820193919092039150565b80356001600160e01b03198116906004841015610a75576001600160e01b0319600485900360031b81901b82161691505b5092915050565b60ff811681146108d0575f5ffd5b5f5f5f5f5f5f5f60e0888a031215610aa0575f5ffd5b8735610aab816108bc565b96506020880135610abb816108bc565b955060408801359450606088013593506080880135610ad981610a7c565b9699959850939692959460a0840135945060c09093013592915050565b5f5f5f5f5f5f5f5f610100898b031215610b0e575f5ffd5b8835610b19816108bc565b97506020890135610b29816108bc565b9650604089013595506060890135945060808901358015158114610b4b575f5ffd5b935060a0890135610b5b81610a7c565b979a969950949793969295929450505060c08201359160e0013590565b5f60208284031215610b88575f5ffd5b81516108f981610a7c565b606081525f610ba56060830186610922565b8281036020840152610bb78186610922565b91505060ff83166040830152949350505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610bef575f5ffd5b815167ffffffffffffffff811115610c05575f5ffd5b8201601f81018413610c15575f5ffd5b805167ffffffffffffffff811115610c2f57610c2f610bcb565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610c5e57610c5e610bcb565b604052818152828201602001861015610c75575f5ffd5b610563826020830160208601610900565b634e487b7160e01b5f52603260045260245ffd5b5f60018201610cb757634e487b7160e01b5f52601160045260245ffd5b506001019056fe60806040819052631d97f74d60e11b81523390633b2fee9a90608490602090600481865afa158015610033573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100579190610433565b604080515f80825260208201909252905061007382825f6100e2565b50506100dd336001600160a01b03166338b8fbbb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100d89190610433565b61010d565b6104c8565b6100eb8361017a565b5f825111806100f75750805b156101085761010683836101b9565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61014c5f516020610e815f395f51905f52546001600160a01b031690565b604080516001600160a01b03928316815291841660208301520160405180910390a1610177816101e5565b50565b61018381610280565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606101de8383604051806060016040528060278152602001610ea160279139610314565b9392505050565b6001600160a01b03811661024f5760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b805f516020610e815f395f51905f525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6102ed5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610246565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc61025f565b60605f5f856001600160a01b031685604051610330919061047b565b5f60405180830381855af49150503d805f8114610368576040519150601f19603f3d011682016040523d82523d5f602084013e61036d565b606091505b50909250905061037f86838387610389565b9695505050505050565b606083156103f75782515f036103f0576001600160a01b0385163b6103f05760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610246565b5081610401565b6104018383610409565b949350505050565b8151156104195781518083602001fd5b8060405162461bcd60e51b81526004016102469190610496565b5f60208284031215610443575f5ffd5b81516001600160a01b03811681146101de575f5ffd5b5f5b8381101561047357818101518382015260200161045b565b50505f910152565b5f825161048c818460208701610459565b9190910192915050565b602081525f82518060208401526104b4816040850160208701610459565b601f01601f19169190910160400192915050565b6109ac806104d55f395ff3fe60806040526004361061005d575f3560e01c80635c60da1b116100425780635c60da1b146100a65780638f283970146100e3578063f851a440146101025761006c565b80633659cfe6146100745780634f1ef286146100935761006c565b3661006c5761006a610116565b005b61006a610116565b34801561007f575f5ffd5b5061006a61008e366004610854565b610130565b61006a6100a136600461086d565b610178565b3480156100b1575f5ffd5b506100ba6101eb565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100ee575f5ffd5b5061006a6100fd366004610854565b610228565b34801561010d575f5ffd5b506100ba610255565b61011e610282565b61012e610129610359565b610362565b565b610138610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d8160405180602001604052805f8152505f6103bf565b50565b61016d610116565b610180610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101e3576101de8383838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250600192506103bf915050565b505050565b6101de610116565b5f6101f4610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610359565b905090565b610225610116565b90565b610230610380565b73ffffffffffffffffffffffffffffffffffffffff1633036101705761016d816103e9565b5f61025e610380565b73ffffffffffffffffffffffffffffffffffffffff16330361021d57610218610380565b61028a610380565b73ffffffffffffffffffffffffffffffffffffffff16330361012e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f7879207461726760648201527f6574000000000000000000000000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b5f61021861044a565b365f5f375f5f365f845af43d5f5f3e80801561037c573d5ff35b3d5ffd5b5f7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b6103c883610471565b5f825111806103d45750805b156101de576103e383836104bd565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f610412610380565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161016d816104e9565b5f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6103a3565b61047a816105f5565b60405173ffffffffffffffffffffffffffffffffffffffff8216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606104e28383604051806060016040528060278152602001610979602791396106c0565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff811661058c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610350565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b73ffffffffffffffffffffffffffffffffffffffff81163b610699576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e7472616374000000000000000000000000000000000000006064820152608401610350565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6105af565b60605f5f8573ffffffffffffffffffffffffffffffffffffffff16856040516106e9919061090d565b5f60405180830381855af49150503d805f8114610721576040519150601f19603f3d011682016040523d82523d5f602084013e610726565b606091505b509150915061073786838387610741565b9695505050505050565b606083156107d65782515f036107cf5773ffffffffffffffffffffffffffffffffffffffff85163b6107cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610350565b50816107e0565b6107e083836107e8565b949350505050565b8151156107f85781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103509190610928565b803573ffffffffffffffffffffffffffffffffffffffff8116811461084f575f5ffd5b919050565b5f60208284031215610864575f5ffd5b6104e28261082c565b5f5f5f6040848603121561087f575f5ffd5b6108888461082c565b9250602084013567ffffffffffffffff8111156108a3575f5ffd5b8401601f810186136108b3575f5ffd5b803567ffffffffffffffff8111156108c9575f5ffd5b8660208284010111156108da575f5ffd5b939660209190910195509293505050565b5f5b838110156109055781810151838201526020016108ed565b50505f910152565b5f825161091e8184602087016108eb565b9190910192915050565b602081525f82518060208401526109468160408501602087016108eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a164736f6c634300081c000ab53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e1009eb7638da0969c98fe82b7d95129c568a1d1da306b07407b7cfb2d95874b64736f6c634300081c0033",
}

// PolygonzkevmbridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonzkevmbridgeMetaData.ABI instead.
var PolygonzkevmbridgeABI = PolygonzkevmbridgeMetaData.ABI

// PolygonzkevmbridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonzkevmbridgeMetaData.Bin instead.
var PolygonzkevmbridgeBin = PolygonzkevmbridgeMetaData.Bin

// DeployPolygonzkevmbridge deploys a new Ethereum contract, binding an instance of Polygonzkevmbridge to it.
func DeployPolygonzkevmbridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Polygonzkevmbridge, error) {
	parsed, err := PolygonzkevmbridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonzkevmbridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonzkevmbridge{PolygonzkevmbridgeCaller: PolygonzkevmbridgeCaller{contract: contract}, PolygonzkevmbridgeTransactor: PolygonzkevmbridgeTransactor{contract: contract}, PolygonzkevmbridgeFilterer: PolygonzkevmbridgeFilterer{contract: contract}}, nil
}

// Polygonzkevmbridge is an auto generated Go binding around an Ethereum contract.
type Polygonzkevmbridge struct {
	PolygonzkevmbridgeCaller     // Read-only binding to the contract
	PolygonzkevmbridgeTransactor // Write-only binding to the contract
	PolygonzkevmbridgeFilterer   // Log filterer for contract events
}

// PolygonzkevmbridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonzkevmbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmbridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonzkevmbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmbridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonzkevmbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonzkevmbridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonzkevmbridgeSession struct {
	Contract     *Polygonzkevmbridge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PolygonzkevmbridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonzkevmbridgeCallerSession struct {
	Contract *PolygonzkevmbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PolygonzkevmbridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonzkevmbridgeTransactorSession struct {
	Contract     *PolygonzkevmbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PolygonzkevmbridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonzkevmbridgeRaw struct {
	Contract *Polygonzkevmbridge // Generic contract binding to access the raw methods on
}

// PolygonzkevmbridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonzkevmbridgeCallerRaw struct {
	Contract *PolygonzkevmbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonzkevmbridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonzkevmbridgeTransactorRaw struct {
	Contract *PolygonzkevmbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonzkevmbridge creates a new instance of Polygonzkevmbridge, bound to a specific deployed contract.
func NewPolygonzkevmbridge(address common.Address, backend bind.ContractBackend) (*Polygonzkevmbridge, error) {
	contract, err := bindPolygonzkevmbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonzkevmbridge{PolygonzkevmbridgeCaller: PolygonzkevmbridgeCaller{contract: contract}, PolygonzkevmbridgeTransactor: PolygonzkevmbridgeTransactor{contract: contract}, PolygonzkevmbridgeFilterer: PolygonzkevmbridgeFilterer{contract: contract}}, nil
}

// NewPolygonzkevmbridgeCaller creates a new read-only instance of Polygonzkevmbridge, bound to a specific deployed contract.
func NewPolygonzkevmbridgeCaller(address common.Address, caller bind.ContractCaller) (*PolygonzkevmbridgeCaller, error) {
	contract, err := bindPolygonzkevmbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeCaller{contract: contract}, nil
}

// NewPolygonzkevmbridgeTransactor creates a new write-only instance of Polygonzkevmbridge, bound to a specific deployed contract.
func NewPolygonzkevmbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonzkevmbridgeTransactor, error) {
	contract, err := bindPolygonzkevmbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeTransactor{contract: contract}, nil
}

// NewPolygonzkevmbridgeFilterer creates a new log filterer instance of Polygonzkevmbridge, bound to a specific deployed contract.
func NewPolygonzkevmbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonzkevmbridgeFilterer, error) {
	contract, err := bindPolygonzkevmbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeFilterer{contract: contract}, nil
}

// bindPolygonzkevmbridge binds a generic wrapper to an already deployed contract.
func bindPolygonzkevmbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonzkevmbridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmbridge *PolygonzkevmbridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmbridge.Contract.PolygonzkevmbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmbridge *PolygonzkevmbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.PolygonzkevmbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmbridge *PolygonzkevmbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.PolygonzkevmbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonzkevmbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.contract.Transact(opts, method, params...)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) INITBYTECODETRANSPARENTPROXY(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "INIT_BYTECODE_TRANSPARENT_PROXY")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Polygonzkevmbridge.Contract.INITBYTECODETRANSPARENTPROXY(&_Polygonzkevmbridge.CallOpts)
}

// INITBYTECODETRANSPARENTPROXY is a free data retrieval call binding the contract method 0xc514f24e.
//
// Solidity: function INIT_BYTECODE_TRANSPARENT_PROXY() view returns(bytes)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) INITBYTECODETRANSPARENTPROXY() ([]byte, error) {
	return _Polygonzkevmbridge.Contract.INITBYTECODETRANSPARENTPROXY(&_Polygonzkevmbridge.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) WETHToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "WETHToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) WETHToken() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.WETHToken(&_Polygonzkevmbridge.CallOpts)
}

// WETHToken is a free data retrieval call binding the contract method 0x4b2f336d.
//
// Solidity: function WETHToken() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) WETHToken() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.WETHToken(&_Polygonzkevmbridge.CallOpts)
}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) BridgeLib(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "bridgeLib")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeLib() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.BridgeLib(&_Polygonzkevmbridge.CallOpts)
}

// BridgeLib is a free data retrieval call binding the contract method 0x6f0bc3da.
//
// Solidity: function bridgeLib() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) BridgeLib() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.BridgeLib(&_Polygonzkevmbridge.CallOpts)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) ClaimedBitMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "claimedBitMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.ClaimedBitMap(&_Polygonzkevmbridge.CallOpts, arg0)
}

// ClaimedBitMap is a free data retrieval call binding the contract method 0xee25560b.
//
// Solidity: function claimedBitMap(uint256 ) view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) ClaimedBitMap(arg0 *big.Int) (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.ClaimedBitMap(&_Polygonzkevmbridge.CallOpts, arg0)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) ComputeTokenProxyAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "computeTokenProxyAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Polygonzkevmbridge.Contract.ComputeTokenProxyAddress(&_Polygonzkevmbridge.CallOpts, originNetwork, originTokenAddress)
}

// ComputeTokenProxyAddress is a free data retrieval call binding the contract method 0xf214e161.
//
// Solidity: function computeTokenProxyAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) ComputeTokenProxyAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Polygonzkevmbridge.Contract.ComputeTokenProxyAddress(&_Polygonzkevmbridge.CallOpts, originNetwork, originTokenAddress)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) DepositCount() (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.DepositCount(&_Polygonzkevmbridge.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) DepositCount() (*big.Int, error) {
	return _Polygonzkevmbridge.Contract.DepositCount(&_Polygonzkevmbridge.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GasTokenAddress(&_Polygonzkevmbridge.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GasTokenAddress(&_Polygonzkevmbridge.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GasTokenMetadata(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "gasTokenMetadata")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GasTokenMetadata() ([]byte, error) {
	return _Polygonzkevmbridge.Contract.GasTokenMetadata(&_Polygonzkevmbridge.CallOpts)
}

// GasTokenMetadata is a free data retrieval call binding the contract method 0x27aef4e8.
//
// Solidity: function gasTokenMetadata() view returns(bytes)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GasTokenMetadata() ([]byte, error) {
	return _Polygonzkevmbridge.Contract.GasTokenMetadata(&_Polygonzkevmbridge.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmbridge.Contract.GasTokenNetwork(&_Polygonzkevmbridge.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonzkevmbridge.Contract.GasTokenNetwork(&_Polygonzkevmbridge.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GetProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "getProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GetProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GetProxiedTokensManager(&_Polygonzkevmbridge.CallOpts)
}

// GetProxiedTokensManager is a free data retrieval call binding the contract method 0x38b8fbbb.
//
// Solidity: function getProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GetProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GetProxiedTokensManager(&_Polygonzkevmbridge.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GetRoot() ([32]byte, error) {
	return _Polygonzkevmbridge.Contract.GetRoot(&_Polygonzkevmbridge.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GetRoot() ([32]byte, error) {
	return _Polygonzkevmbridge.Contract.GetRoot(&_Polygonzkevmbridge.CallOpts)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GetTokenMetadata(opts *bind.CallOpts, token common.Address) ([]byte, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "getTokenMetadata", token)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Polygonzkevmbridge.Contract.GetTokenMetadata(&_Polygonzkevmbridge.CallOpts, token)
}

// GetTokenMetadata is a free data retrieval call binding the contract method 0xc00f14ab.
//
// Solidity: function getTokenMetadata(address token) view returns(bytes)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GetTokenMetadata(token common.Address) ([]byte, error) {
	return _Polygonzkevmbridge.Contract.GetTokenMetadata(&_Polygonzkevmbridge.CallOpts, token)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GetTokenWrappedAddress(opts *bind.CallOpts, originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "getTokenWrappedAddress", originNetwork, originTokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GetTokenWrappedAddress(&_Polygonzkevmbridge.CallOpts, originNetwork, originTokenAddress)
}

// GetTokenWrappedAddress is a free data retrieval call binding the contract method 0x22e95f2c.
//
// Solidity: function getTokenWrappedAddress(uint32 originNetwork, address originTokenAddress) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GetTokenWrappedAddress(originNetwork uint32, originTokenAddress common.Address) (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GetTokenWrappedAddress(&_Polygonzkevmbridge.CallOpts, originNetwork, originTokenAddress)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GetWrappedTokenBridgeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "getWrappedTokenBridgeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GetWrappedTokenBridgeImplementation(&_Polygonzkevmbridge.CallOpts)
}

// GetWrappedTokenBridgeImplementation is a free data retrieval call binding the contract method 0x3b2fee9a.
//
// Solidity: function getWrappedTokenBridgeImplementation() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GetWrappedTokenBridgeImplementation() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GetWrappedTokenBridgeImplementation(&_Polygonzkevmbridge.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GlobalExitRootManager(&_Polygonzkevmbridge.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.GlobalExitRootManager(&_Polygonzkevmbridge.CallOpts)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) IsClaimed(opts *bind.CallOpts, leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "isClaimed", leafIndex, sourceBridgeNetwork)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Polygonzkevmbridge.Contract.IsClaimed(&_Polygonzkevmbridge.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsClaimed is a free data retrieval call binding the contract method 0xcc461632.
//
// Solidity: function isClaimed(uint32 leafIndex, uint32 sourceBridgeNetwork) view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) IsClaimed(leafIndex uint32, sourceBridgeNetwork uint32) (bool, error) {
	return _Polygonzkevmbridge.Contract.IsClaimed(&_Polygonzkevmbridge.CallOpts, leafIndex, sourceBridgeNetwork)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) IsEmergencyState(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "isEmergencyState")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) IsEmergencyState() (bool, error) {
	return _Polygonzkevmbridge.Contract.IsEmergencyState(&_Polygonzkevmbridge.CallOpts)
}

// IsEmergencyState is a free data retrieval call binding the contract method 0x15064c96.
//
// Solidity: function isEmergencyState() view returns(bool)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) IsEmergencyState() (bool, error) {
	return _Polygonzkevmbridge.Contract.IsEmergencyState(&_Polygonzkevmbridge.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) LastUpdatedDepositCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "lastUpdatedDepositCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) LastUpdatedDepositCount() (uint32, error) {
	return _Polygonzkevmbridge.Contract.LastUpdatedDepositCount(&_Polygonzkevmbridge.CallOpts)
}

// LastUpdatedDepositCount is a free data retrieval call binding the contract method 0xbe5831c7.
//
// Solidity: function lastUpdatedDepositCount() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) LastUpdatedDepositCount() (uint32, error) {
	return _Polygonzkevmbridge.Contract.LastUpdatedDepositCount(&_Polygonzkevmbridge.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) NetworkID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "networkID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) NetworkID() (uint32, error) {
	return _Polygonzkevmbridge.Contract.NetworkID(&_Polygonzkevmbridge.CallOpts)
}

// NetworkID is a free data retrieval call binding the contract method 0xbab161bf.
//
// Solidity: function networkID() view returns(uint32)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) NetworkID() (uint32, error) {
	return _Polygonzkevmbridge.Contract.NetworkID(&_Polygonzkevmbridge.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) PendingProxiedTokensManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "pendingProxiedTokensManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.PendingProxiedTokensManager(&_Polygonzkevmbridge.CallOpts)
}

// PendingProxiedTokensManager is a free data retrieval call binding the contract method 0xece93c6f.
//
// Solidity: function pendingProxiedTokensManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) PendingProxiedTokensManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.PendingProxiedTokensManager(&_Polygonzkevmbridge.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) PolygonRollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "polygonRollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) PolygonRollupManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.PolygonRollupManager(&_Polygonzkevmbridge.CallOpts)
}

// PolygonRollupManager is a free data retrieval call binding the contract method 0x8ed7e3f2.
//
// Solidity: function polygonRollupManager() view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) PolygonRollupManager() (common.Address, error) {
	return _Polygonzkevmbridge.Contract.PolygonRollupManager(&_Polygonzkevmbridge.CallOpts)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) TokenInfoToWrappedToken(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "tokenInfoToWrappedToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Polygonzkevmbridge.Contract.TokenInfoToWrappedToken(&_Polygonzkevmbridge.CallOpts, arg0)
}

// TokenInfoToWrappedToken is a free data retrieval call binding the contract method 0x81b1c174.
//
// Solidity: function tokenInfoToWrappedToken(bytes32 ) view returns(address)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) TokenInfoToWrappedToken(arg0 [32]byte) (common.Address, error) {
	return _Polygonzkevmbridge.Contract.TokenInfoToWrappedToken(&_Polygonzkevmbridge.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) Version() (string, error) {
	return _Polygonzkevmbridge.Contract.Version(&_Polygonzkevmbridge.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) Version() (string, error) {
	return _Polygonzkevmbridge.Contract.Version(&_Polygonzkevmbridge.CallOpts)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCaller) WrappedTokenToTokenInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Polygonzkevmbridge.contract.Call(opts, &out, "wrappedTokenToTokenInfo", arg0)

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
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Polygonzkevmbridge.Contract.WrappedTokenToTokenInfo(&_Polygonzkevmbridge.CallOpts, arg0)
}

// WrappedTokenToTokenInfo is a free data retrieval call binding the contract method 0x318aee3d.
//
// Solidity: function wrappedTokenToTokenInfo(address ) view returns(uint32 originNetwork, address originTokenAddress)
func (_Polygonzkevmbridge *PolygonzkevmbridgeCallerSession) WrappedTokenToTokenInfo(arg0 common.Address) (struct {
	OriginNetwork      uint32
	OriginTokenAddress common.Address
}, error) {
	return _Polygonzkevmbridge.Contract.WrappedTokenToTokenInfo(&_Polygonzkevmbridge.CallOpts, arg0)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) AcceptProxiedTokensManagerRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "acceptProxiedTokensManagerRole")
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.AcceptProxiedTokensManagerRole(&_Polygonzkevmbridge.TransactOpts)
}

// AcceptProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8c668f1c.
//
// Solidity: function acceptProxiedTokensManagerRole() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) AcceptProxiedTokensManagerRole() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.AcceptProxiedTokensManagerRole(&_Polygonzkevmbridge.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) ActivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "activateEmergencyState")
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.ActivateEmergencyState(&_Polygonzkevmbridge.TransactOpts)
}

// ActivateEmergencyState is a paid mutator transaction binding the contract method 0x2072f6c5.
//
// Solidity: function activateEmergencyState() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) ActivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.ActivateEmergencyState(&_Polygonzkevmbridge.TransactOpts)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BackwardLET(opts *bind.TransactOpts, newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "backwardLET", newDepositCount, newFrontier, nextLeaf, proof)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BackwardLET(newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BackwardLET(&_Polygonzkevmbridge.TransactOpts, newDepositCount, newFrontier, nextLeaf, proof)
}

// BackwardLET is a paid mutator transaction binding the contract method 0xfd7640e8.
//
// Solidity: function backwardLET(uint256 newDepositCount, bytes32[32] newFrontier, bytes32 nextLeaf, bytes32[32] proof) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BackwardLET(newDepositCount *big.Int, newFrontier [32][32]byte, nextLeaf [32]byte, proof [32][32]byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BackwardLET(&_Polygonzkevmbridge.TransactOpts, newDepositCount, newFrontier, nextLeaf, proof)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BridgeAsset(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "bridgeAsset", destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeAsset(&_Polygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeAsset is a paid mutator transaction binding the contract method 0xcd586579.
//
// Solidity: function bridgeAsset(uint32 destinationNetwork, address destinationAddress, uint256 amount, address token, bool forceUpdateGlobalExitRoot, bytes permitData) payable returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BridgeAsset(destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, token common.Address, forceUpdateGlobalExitRoot bool, permitData []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeAsset(&_Polygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, amount, token, forceUpdateGlobalExitRoot, permitData)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BridgeMessage(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "bridgeMessage", destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeMessage(&_Polygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessage is a paid mutator transaction binding the contract method 0x240ff378.
//
// Solidity: function bridgeMessage(uint32 destinationNetwork, address destinationAddress, bool forceUpdateGlobalExitRoot, bytes metadata) payable returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BridgeMessage(destinationNetwork uint32, destinationAddress common.Address, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeMessage(&_Polygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) BridgeMessageWETH(opts *bind.TransactOpts, destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "bridgeMessageWETH", destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeMessageWETH(&_Polygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// BridgeMessageWETH is a paid mutator transaction binding the contract method 0xb8b284d0.
//
// Solidity: function bridgeMessageWETH(uint32 destinationNetwork, address destinationAddress, uint256 amountWETH, bool forceUpdateGlobalExitRoot, bytes metadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) BridgeMessageWETH(destinationNetwork uint32, destinationAddress common.Address, amountWETH *big.Int, forceUpdateGlobalExitRoot bool, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.BridgeMessageWETH(&_Polygonzkevmbridge.TransactOpts, destinationNetwork, destinationAddress, amountWETH, forceUpdateGlobalExitRoot, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) ClaimAsset(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "claimAsset", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.ClaimAsset(&_Polygonzkevmbridge.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0xccaa2d11.
//
// Solidity: function claimAsset(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originTokenAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) ClaimAsset(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originTokenAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.ClaimAsset(&_Polygonzkevmbridge.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originTokenAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) ClaimMessage(opts *bind.TransactOpts, smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "claimMessage", smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.ClaimMessage(&_Polygonzkevmbridge.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0xf5efcd79.
//
// Solidity: function claimMessage(bytes32[32] smtProofLocalExitRoot, bytes32[32] smtProofRollupExitRoot, uint256 globalIndex, bytes32 mainnetExitRoot, bytes32 rollupExitRoot, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) ClaimMessage(smtProofLocalExitRoot [32][32]byte, smtProofRollupExitRoot [32][32]byte, globalIndex *big.Int, mainnetExitRoot [32]byte, rollupExitRoot [32]byte, originNetwork uint32, originAddress common.Address, destinationNetwork uint32, destinationAddress common.Address, amount *big.Int, metadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.ClaimMessage(&_Polygonzkevmbridge.TransactOpts, smtProofLocalExitRoot, smtProofRollupExitRoot, globalIndex, mainnetExitRoot, rollupExitRoot, originNetwork, originAddress, destinationNetwork, destinationAddress, amount, metadata)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) DeactivateEmergencyState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "deactivateEmergencyState")
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.DeactivateEmergencyState(&_Polygonzkevmbridge.TransactOpts)
}

// DeactivateEmergencyState is a paid mutator transaction binding the contract method 0xdbc16976.
//
// Solidity: function deactivateEmergencyState() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) DeactivateEmergencyState() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.DeactivateEmergencyState(&_Polygonzkevmbridge.TransactOpts)
}

// ForwardLET is a paid mutator transaction binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] newLeaves, bytes32 expectedLER) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) ForwardLET(opts *bind.TransactOpts, newLeaves []AgglayerBridgeLeafData, expectedLER [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "forwardLET", newLeaves, expectedLER)
}

// ForwardLET is a paid mutator transaction binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] newLeaves, bytes32 expectedLER) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) ForwardLET(newLeaves []AgglayerBridgeLeafData, expectedLER [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.ForwardLET(&_Polygonzkevmbridge.TransactOpts, newLeaves, expectedLER)
}

// ForwardLET is a paid mutator transaction binding the contract method 0xf0e70808.
//
// Solidity: function forwardLET((uint8,uint32,address,uint32,address,uint256,bytes)[] newLeaves, bytes32 expectedLER) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) ForwardLET(newLeaves []AgglayerBridgeLeafData, expectedLER [32]byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.ForwardLET(&_Polygonzkevmbridge.TransactOpts, newLeaves, expectedLER)
}

// Initialize is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) Initialize(opts *bind.TransactOpts, _networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "initialize", _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.Initialize(&_Polygonzkevmbridge.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// Initialize is a paid mutator transaction binding the contract method 0xf811bff7.
//
// Solidity: function initialize(uint32 _networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, address _globalExitRootManager, address _polygonRollupManager, bytes _gasTokenMetadata) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) Initialize(_networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _globalExitRootManager common.Address, _polygonRollupManager common.Address, _gasTokenMetadata []byte) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.Initialize(&_Polygonzkevmbridge.TransactOpts, _networkID, _gasTokenAddress, _gasTokenNetwork, _globalExitRootManager, _polygonRollupManager, _gasTokenMetadata)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) SetMultipleClaims(opts *bind.TransactOpts, globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "setMultipleClaims", globalIndexes)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) SetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetMultipleClaims(&_Polygonzkevmbridge.TransactOpts, globalIndexes)
}

// SetMultipleClaims is a paid mutator transaction binding the contract method 0x1c208229.
//
// Solidity: function setMultipleClaims(uint256[] globalIndexes) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) SetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.SetMultipleClaims(&_Polygonzkevmbridge.TransactOpts, globalIndexes)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) TransferProxiedTokensManagerRole(opts *bind.TransactOpts, newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "transferProxiedTokensManagerRole", newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.TransferProxiedTokensManagerRole(&_Polygonzkevmbridge.TransactOpts, newProxiedTokensManager)
}

// TransferProxiedTokensManagerRole is a paid mutator transaction binding the contract method 0x8bd309c3.
//
// Solidity: function transferProxiedTokensManagerRole(address newProxiedTokensManager) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) TransferProxiedTokensManagerRole(newProxiedTokensManager common.Address) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.TransferProxiedTokensManagerRole(&_Polygonzkevmbridge.TransactOpts, newProxiedTokensManager)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) UnsetMultipleClaims(opts *bind.TransactOpts, globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "unsetMultipleClaims", globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.UnsetMultipleClaims(&_Polygonzkevmbridge.TransactOpts, globalIndexes)
}

// UnsetMultipleClaims is a paid mutator transaction binding the contract method 0x136a2c60.
//
// Solidity: function unsetMultipleClaims(uint256[] globalIndexes) returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) UnsetMultipleClaims(globalIndexes []*big.Int) (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.UnsetMultipleClaims(&_Polygonzkevmbridge.TransactOpts, globalIndexes)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactor) UpdateGlobalExitRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonzkevmbridge.contract.Transact(opts, "updateGlobalExitRoot")
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.UpdateGlobalExitRoot(&_Polygonzkevmbridge.TransactOpts)
}

// UpdateGlobalExitRoot is a paid mutator transaction binding the contract method 0x79e2cf97.
//
// Solidity: function updateGlobalExitRoot() returns()
func (_Polygonzkevmbridge *PolygonzkevmbridgeTransactorSession) UpdateGlobalExitRoot() (*types.Transaction, error) {
	return _Polygonzkevmbridge.Contract.UpdateGlobalExitRoot(&_Polygonzkevmbridge.TransactOpts)
}

// PolygonzkevmbridgeAcceptProxiedTokensManagerRoleIterator is returned from FilterAcceptProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for AcceptProxiedTokensManagerRole events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeAcceptProxiedTokensManagerRoleIterator struct {
	Event *PolygonzkevmbridgeAcceptProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeAcceptProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeAcceptProxiedTokensManagerRole)
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
		it.Event = new(PolygonzkevmbridgeAcceptProxiedTokensManagerRole)
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
func (it *PolygonzkevmbridgeAcceptProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeAcceptProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeAcceptProxiedTokensManagerRole represents a AcceptProxiedTokensManagerRole event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeAcceptProxiedTokensManagerRole struct {
	OldProxiedTokensManager common.Address
	NewProxiedTokensManager common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAcceptProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterAcceptProxiedTokensManagerRole(opts *bind.FilterOpts) (*PolygonzkevmbridgeAcceptProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeAcceptProxiedTokensManagerRoleIterator{contract: _Polygonzkevmbridge.contract, event: "AcceptProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchAcceptProxiedTokensManagerRole is a free log subscription operation binding the contract event 0xa9da6fb8c39e9c2fafda878eac316815987bdc948d241ba6d75ed035e0e829f2.
//
// Solidity: event AcceptProxiedTokensManagerRole(address oldProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchAcceptProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeAcceptProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "AcceptProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeAcceptProxiedTokensManagerRole)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseAcceptProxiedTokensManagerRole(log types.Log) (*PolygonzkevmbridgeAcceptProxiedTokensManagerRole, error) {
	event := new(PolygonzkevmbridgeAcceptProxiedTokensManagerRole)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "AcceptProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeBackwardLETIterator is returned from FilterBackwardLET and is used to iterate over the raw logs and unpacked data for BackwardLET events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeBackwardLETIterator struct {
	Event *PolygonzkevmbridgeBackwardLET // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeBackwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeBackwardLET)
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
		it.Event = new(PolygonzkevmbridgeBackwardLET)
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
func (it *PolygonzkevmbridgeBackwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeBackwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeBackwardLET represents a BackwardLET event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeBackwardLET struct {
	PreviousDepositCount *big.Int
	PreviousRoot         [32]byte
	NewDepositCount      *big.Int
	NewRoot              [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBackwardLET is a free log retrieval operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterBackwardLET(opts *bind.FilterOpts) (*PolygonzkevmbridgeBackwardLETIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeBackwardLETIterator{contract: _Polygonzkevmbridge.contract, event: "BackwardLET", logs: logs, sub: sub}, nil
}

// WatchBackwardLET is a free log subscription operation binding the contract event 0x2f8b2f36ca0d63af2efffe9d5e728a61c2bc273a3fc0574d12b558a41f149555.
//
// Solidity: event BackwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchBackwardLET(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeBackwardLET) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "BackwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeBackwardLET)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "BackwardLET", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseBackwardLET(log types.Log) (*PolygonzkevmbridgeBackwardLET, error) {
	event := new(PolygonzkevmbridgeBackwardLET)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "BackwardLET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeBridgeEventIterator is returned from FilterBridgeEvent and is used to iterate over the raw logs and unpacked data for BridgeEvent events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeBridgeEventIterator struct {
	Event *PolygonzkevmbridgeBridgeEvent // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeBridgeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeBridgeEvent)
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
		it.Event = new(PolygonzkevmbridgeBridgeEvent)
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
func (it *PolygonzkevmbridgeBridgeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeBridgeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeBridgeEvent represents a BridgeEvent event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeBridgeEvent struct {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterBridgeEvent(opts *bind.FilterOpts) (*PolygonzkevmbridgeBridgeEventIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeBridgeEventIterator{contract: _Polygonzkevmbridge.contract, event: "BridgeEvent", logs: logs, sub: sub}, nil
}

// WatchBridgeEvent is a free log subscription operation binding the contract event 0x501781209a1f8899323b96b4ef08b168df93e0a90c673d1e4cce39366cb62f9b.
//
// Solidity: event BridgeEvent(uint8 leafType, uint32 originNetwork, address originAddress, uint32 destinationNetwork, address destinationAddress, uint256 amount, bytes metadata, uint32 depositCount)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchBridgeEvent(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeBridgeEvent) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "BridgeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeBridgeEvent)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseBridgeEvent(log types.Log) (*PolygonzkevmbridgeBridgeEvent, error) {
	event := new(PolygonzkevmbridgeBridgeEvent)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "BridgeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeClaimEventIterator struct {
	Event *PolygonzkevmbridgeClaimEvent // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeClaimEvent)
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
		it.Event = new(PolygonzkevmbridgeClaimEvent)
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
func (it *PolygonzkevmbridgeClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeClaimEvent represents a ClaimEvent event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeClaimEvent struct {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterClaimEvent(opts *bind.FilterOpts) (*PolygonzkevmbridgeClaimEventIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeClaimEventIterator{contract: _Polygonzkevmbridge.contract, event: "ClaimEvent", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x1df3f2a973a00d6635911755c260704e95e8a5876997546798770f76396fda4d.
//
// Solidity: event ClaimEvent(uint256 globalIndex, uint32 originNetwork, address originAddress, address destinationAddress, uint256 amount)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeClaimEvent) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "ClaimEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeClaimEvent)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseClaimEvent(log types.Log) (*PolygonzkevmbridgeClaimEvent, error) {
	event := new(PolygonzkevmbridgeClaimEvent)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "ClaimEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeEmergencyStateActivatedIterator is returned from FilterEmergencyStateActivated and is used to iterate over the raw logs and unpacked data for EmergencyStateActivated events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeEmergencyStateActivatedIterator struct {
	Event *PolygonzkevmbridgeEmergencyStateActivated // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeEmergencyStateActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeEmergencyStateActivated)
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
		it.Event = new(PolygonzkevmbridgeEmergencyStateActivated)
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
func (it *PolygonzkevmbridgeEmergencyStateActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeEmergencyStateActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeEmergencyStateActivated represents a EmergencyStateActivated event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeEmergencyStateActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateActivated is a free log retrieval operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterEmergencyStateActivated(opts *bind.FilterOpts) (*PolygonzkevmbridgeEmergencyStateActivatedIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeEmergencyStateActivatedIterator{contract: _Polygonzkevmbridge.contract, event: "EmergencyStateActivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateActivated is a free log subscription operation binding the contract event 0x2261efe5aef6fedc1fd1550b25facc9181745623049c7901287030b9ad1a5497.
//
// Solidity: event EmergencyStateActivated()
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchEmergencyStateActivated(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeEmergencyStateActivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "EmergencyStateActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeEmergencyStateActivated)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseEmergencyStateActivated(log types.Log) (*PolygonzkevmbridgeEmergencyStateActivated, error) {
	event := new(PolygonzkevmbridgeEmergencyStateActivated)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "EmergencyStateActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeEmergencyStateDeactivatedIterator is returned from FilterEmergencyStateDeactivated and is used to iterate over the raw logs and unpacked data for EmergencyStateDeactivated events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeEmergencyStateDeactivatedIterator struct {
	Event *PolygonzkevmbridgeEmergencyStateDeactivated // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeEmergencyStateDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeEmergencyStateDeactivated)
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
		it.Event = new(PolygonzkevmbridgeEmergencyStateDeactivated)
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
func (it *PolygonzkevmbridgeEmergencyStateDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeEmergencyStateDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeEmergencyStateDeactivated represents a EmergencyStateDeactivated event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeEmergencyStateDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStateDeactivated is a free log retrieval operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterEmergencyStateDeactivated(opts *bind.FilterOpts) (*PolygonzkevmbridgeEmergencyStateDeactivatedIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeEmergencyStateDeactivatedIterator{contract: _Polygonzkevmbridge.contract, event: "EmergencyStateDeactivated", logs: logs, sub: sub}, nil
}

// WatchEmergencyStateDeactivated is a free log subscription operation binding the contract event 0x1e5e34eea33501aecf2ebec9fe0e884a40804275ea7fe10b2ba084c8374308b3.
//
// Solidity: event EmergencyStateDeactivated()
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchEmergencyStateDeactivated(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeEmergencyStateDeactivated) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "EmergencyStateDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeEmergencyStateDeactivated)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseEmergencyStateDeactivated(log types.Log) (*PolygonzkevmbridgeEmergencyStateDeactivated, error) {
	event := new(PolygonzkevmbridgeEmergencyStateDeactivated)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "EmergencyStateDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeForwardLETIterator is returned from FilterForwardLET and is used to iterate over the raw logs and unpacked data for ForwardLET events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeForwardLETIterator struct {
	Event *PolygonzkevmbridgeForwardLET // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeForwardLETIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeForwardLET)
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
		it.Event = new(PolygonzkevmbridgeForwardLET)
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
func (it *PolygonzkevmbridgeForwardLETIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeForwardLETIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeForwardLET represents a ForwardLET event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeForwardLET struct {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterForwardLET(opts *bind.FilterOpts) (*PolygonzkevmbridgeForwardLETIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeForwardLETIterator{contract: _Polygonzkevmbridge.contract, event: "ForwardLET", logs: logs, sub: sub}, nil
}

// WatchForwardLET is a free log subscription operation binding the contract event 0xe90f3fcbcd86a778012608da0cd7622521c8dac9919dddf25a3fb634adb3d2c8.
//
// Solidity: event ForwardLET(uint256 previousDepositCount, bytes32 previousRoot, uint256 newDepositCount, bytes32 newRoot, bytes newLeaves)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchForwardLET(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeForwardLET) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "ForwardLET")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeForwardLET)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "ForwardLET", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseForwardLET(log types.Log) (*PolygonzkevmbridgeForwardLET, error) {
	event := new(PolygonzkevmbridgeForwardLET)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "ForwardLET", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitializedIterator struct {
	Event *PolygonzkevmbridgeInitialized // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeInitialized)
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
		it.Event = new(PolygonzkevmbridgeInitialized)
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
func (it *PolygonzkevmbridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeInitialized represents a Initialized event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonzkevmbridgeInitializedIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeInitializedIterator{contract: _Polygonzkevmbridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeInitialized)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseInitialized(log types.Log) (*PolygonzkevmbridgeInitialized, error) {
	event := new(PolygonzkevmbridgeInitialized)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeNewWrappedTokenIterator is returned from FilterNewWrappedToken and is used to iterate over the raw logs and unpacked data for NewWrappedToken events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeNewWrappedTokenIterator struct {
	Event *PolygonzkevmbridgeNewWrappedToken // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeNewWrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeNewWrappedToken)
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
		it.Event = new(PolygonzkevmbridgeNewWrappedToken)
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
func (it *PolygonzkevmbridgeNewWrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeNewWrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeNewWrappedToken represents a NewWrappedToken event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeNewWrappedToken struct {
	OriginNetwork       uint32
	OriginTokenAddress  common.Address
	WrappedTokenAddress common.Address
	Metadata            []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWrappedToken is a free log retrieval operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterNewWrappedToken(opts *bind.FilterOpts) (*PolygonzkevmbridgeNewWrappedTokenIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeNewWrappedTokenIterator{contract: _Polygonzkevmbridge.contract, event: "NewWrappedToken", logs: logs, sub: sub}, nil
}

// WatchNewWrappedToken is a free log subscription operation binding the contract event 0x490e59a1701b938786ac72570a1efeac994a3dbe96e2e883e19e902ace6e6a39.
//
// Solidity: event NewWrappedToken(uint32 originNetwork, address originTokenAddress, address wrappedTokenAddress, bytes metadata)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchNewWrappedToken(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeNewWrappedToken) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "NewWrappedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeNewWrappedToken)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseNewWrappedToken(log types.Log) (*PolygonzkevmbridgeNewWrappedToken, error) {
	event := new(PolygonzkevmbridgeNewWrappedToken)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "NewWrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeSetClaimIterator is returned from FilterSetClaim and is used to iterate over the raw logs and unpacked data for SetClaim events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeSetClaimIterator struct {
	Event *PolygonzkevmbridgeSetClaim // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeSetClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeSetClaim)
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
		it.Event = new(PolygonzkevmbridgeSetClaim)
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
func (it *PolygonzkevmbridgeSetClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeSetClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeSetClaim represents a SetClaim event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeSetClaim struct {
	GlobalIndex [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetClaim is a free log retrieval operation binding the contract event 0x8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c.
//
// Solidity: event SetClaim(bytes32 globalIndex)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterSetClaim(opts *bind.FilterOpts) (*PolygonzkevmbridgeSetClaimIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeSetClaimIterator{contract: _Polygonzkevmbridge.contract, event: "SetClaim", logs: logs, sub: sub}, nil
}

// WatchSetClaim is a free log subscription operation binding the contract event 0x8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c.
//
// Solidity: event SetClaim(bytes32 globalIndex)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchSetClaim(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeSetClaim) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "SetClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeSetClaim)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "SetClaim", log); err != nil {
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

// ParseSetClaim is a log parse operation binding the contract event 0x8b65fb527138ff0fedc6b8e75e5a12268c71a1231b370b03209f7a286746823c.
//
// Solidity: event SetClaim(bytes32 globalIndex)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseSetClaim(log types.Log) (*PolygonzkevmbridgeSetClaim, error) {
	event := new(PolygonzkevmbridgeSetClaim)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "SetClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeTransferProxiedTokensManagerRoleIterator is returned from FilterTransferProxiedTokensManagerRole and is used to iterate over the raw logs and unpacked data for TransferProxiedTokensManagerRole events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeTransferProxiedTokensManagerRoleIterator struct {
	Event *PolygonzkevmbridgeTransferProxiedTokensManagerRole // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeTransferProxiedTokensManagerRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeTransferProxiedTokensManagerRole)
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
		it.Event = new(PolygonzkevmbridgeTransferProxiedTokensManagerRole)
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
func (it *PolygonzkevmbridgeTransferProxiedTokensManagerRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeTransferProxiedTokensManagerRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeTransferProxiedTokensManagerRole represents a TransferProxiedTokensManagerRole event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeTransferProxiedTokensManagerRole struct {
	CurrentProxiedTokensManager common.Address
	NewProxiedTokensManager     common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTransferProxiedTokensManagerRole is a free log retrieval operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterTransferProxiedTokensManagerRole(opts *bind.FilterOpts) (*PolygonzkevmbridgeTransferProxiedTokensManagerRoleIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeTransferProxiedTokensManagerRoleIterator{contract: _Polygonzkevmbridge.contract, event: "TransferProxiedTokensManagerRole", logs: logs, sub: sub}, nil
}

// WatchTransferProxiedTokensManagerRole is a free log subscription operation binding the contract event 0x0a34baa3feb299aef9c05cb59c6e0c8e7c0bcc65cbf0a647e7a7c8a2411591e2.
//
// Solidity: event TransferProxiedTokensManagerRole(address currentProxiedTokensManager, address newProxiedTokensManager)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchTransferProxiedTokensManagerRole(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeTransferProxiedTokensManagerRole) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "TransferProxiedTokensManagerRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeTransferProxiedTokensManagerRole)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseTransferProxiedTokensManagerRole(log types.Log) (*PolygonzkevmbridgeTransferProxiedTokensManagerRole, error) {
	event := new(PolygonzkevmbridgeTransferProxiedTokensManagerRole)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "TransferProxiedTokensManagerRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChainIterator is returned from FilterUpdatedUnsetGlobalIndexHashChain and is used to iterate over the raw logs and unpacked data for UpdatedUnsetGlobalIndexHashChain events raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChainIterator struct {
	Event *PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChain // Event containing the contract specifics and raw log

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
func (it *PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChain)
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
		it.Event = new(PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChain)
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
func (it *PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChain represents a UpdatedUnsetGlobalIndexHashChain event raised by the Polygonzkevmbridge contract.
type PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChain struct {
	UnsetGlobalIndex             [32]byte
	NewUnsetGlobalIndexHashChain [32]byte
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterUpdatedUnsetGlobalIndexHashChain is a free log retrieval operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) FilterUpdatedUnsetGlobalIndexHashChain(opts *bind.FilterOpts) (*PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChainIterator, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.FilterLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return &PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChainIterator{contract: _Polygonzkevmbridge.contract, event: "UpdatedUnsetGlobalIndexHashChain", logs: logs, sub: sub}, nil
}

// WatchUpdatedUnsetGlobalIndexHashChain is a free log subscription operation binding the contract event 0xc80e0aca446a59735359a7ae46124b57c47b892827642779bc6dafc84ba90b03.
//
// Solidity: event UpdatedUnsetGlobalIndexHashChain(bytes32 unsetGlobalIndex, bytes32 newUnsetGlobalIndexHashChain)
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) WatchUpdatedUnsetGlobalIndexHashChain(opts *bind.WatchOpts, sink chan<- *PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChain) (event.Subscription, error) {

	logs, sub, err := _Polygonzkevmbridge.contract.WatchLogs(opts, "UpdatedUnsetGlobalIndexHashChain")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChain)
				if err := _Polygonzkevmbridge.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
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
func (_Polygonzkevmbridge *PolygonzkevmbridgeFilterer) ParseUpdatedUnsetGlobalIndexHashChain(log types.Log) (*PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChain, error) {
	event := new(PolygonzkevmbridgeUpdatedUnsetGlobalIndexHashChain)
	if err := _Polygonzkevmbridge.contract.UnpackLog(event, "UpdatedUnsetGlobalIndexHashChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
