package messagebridge

import (
	"math/big"

	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils"
	"github.com/ethereum/go-ethereum/common"
)

type ProcessorType int

const (
	USDC ProcessorType = iota
	WstETH
	EURC
)

var (
	emptyAddress = common.Address{}
	processorMap = make(map[ProcessorType]*Processor)
)

// Processor hosts the processing functions for an LxLy bridge using the message bridge feature
// Each Processor object should be used for one type of bridged token only
// Current supported tokens: USDC, wstETH
type Processor struct {
	contractToTokenMapping map[common.Address]common.Address
	contractAddressList    []common.Address
	tokenAddressList       []common.Address
	// DecodeMetadata decodes the metadata of the message bridge, returns the actual destination address and bridged amount
	DecodeMetadataFn func(metadata []byte) (common.Address, *big.Int)
}

// GetContractAddressList returns the list of contract addresses that need to be processed through this struct
func (u *Processor) GetContractAddressList() []common.Address {
	return u.contractAddressList
}

// GetTokenAddressList returns the list of original token addresses
func (u *Processor) GetTokenAddressList() []common.Address {
	return u.tokenAddressList
}

// CheckContractAddress returns true if the input address is in the contract address list of this bridge
func (u *Processor) CheckContractAddress(address common.Address) bool {
	if _, ok := u.contractToTokenMapping[address]; ok {
		return true
	}
	return false
}

// GetTokenFromContract return the token address from the bridge contract address, for displaying
func (u *Processor) GetTokenFromContract(contractAddress common.Address) (common.Address, bool) {
	if token, ok := u.contractToTokenMapping[contractAddress]; ok {
		return token, true
	}
	return common.Address{}, false
}

// ReplaceDepositInfo replaces the info of the deposit based on the address mapping
// Info to be replaced: amount, original token address
func (u *Processor) ReplaceDepositInfo(deposit *etherman.Deposit, overwriteOrigNetworkID bool) {
	token, ok := u.GetTokenFromContract(deposit.OriginalAddress)
	if !ok {
		return
	}
	deposit.OriginalAddress = token
	if overwriteOrigNetworkID {
		deposit.OriginalNetwork = 0 // Always use 0 for this case when reporting metrics
	}
	_, deposit.Amount = u.DecodeMetadataFn(deposit.Metadata)
}

// getProcessor returns the correct message bridge processor for the address
func getProcessor(origAddr, dstContractAddr common.Address) *Processor {
	for _, processor := range processorMap {
		if processor.CheckContractAddress(origAddr) {
			return processor
		}
	}
	return nil
}

func GetProcessorByType(t ProcessorType) *Processor {
	return processorMap[t]
}

func GetContractAddressList() []common.Address {
	result := make([]common.Address, 0)
	// Get all contract addresses from the lists
	for _, processor := range processorMap {
		result = append(result, processor.GetContractAddressList()...)
	}
	return result
}

func IsAllowedContractAddress(address common.Address) bool {
	for _, processor := range processorMap {
		if processor.CheckContractAddress(address) {
			return true
		}
	}
	return false
}

// ReplaceDepositDestAddresses swaps the actual dest address and the contract address so that we can query using the user's address from DB
func ReplaceDepositDestAddresses(deposit *etherman.Deposit) {
	if deposit.LeafType != uint8(utils.LeafTypeMessage) {
		// Only process message bridges
		return
	}
	processor := getProcessor(deposit.OriginalAddress, deposit.DestContractAddress)
	log.Infow("REPLACE DEPOSIT ADDR", "processor", processor, "origAddr", deposit.OriginalAddress, "DestContractAddress", deposit.DestContractAddress)
	if processor == nil {
		// Cannot find any valid processor
		return
	}
	deposit.DestContractAddress = deposit.DestinationAddress
	deposit.DestinationAddress, _ = processor.DecodeMetadataFn(deposit.Metadata)
}

func ReplaceDepositInfo(deposit *etherman.Deposit, overwriteOrigNetworkID bool) {
	processor := getProcessor(deposit.OriginalAddress, deposit.DestContractAddress)
	if processor == nil {
		return
	}
	processor.ReplaceDepositInfo(deposit, overwriteOrigNetworkID)
}
