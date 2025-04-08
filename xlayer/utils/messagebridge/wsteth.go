package messagebridge

import (
	"math/big"

	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/ethereum/go-ethereum/common"
)

func InitWstETHProcessor(wstETHContractAddresses, wstETHTokenAddresses []common.Address) {
	log.Debugf("WstETHMapping: contracts[%v] tokens[%v]", wstETHContractAddresses, wstETHTokenAddresses)
	if len(wstETHContractAddresses) != len(wstETHTokenAddresses) {
		log.Errorf("InitWstETHProcessor: contract addresses (%v) and token addresses (%v) have different length", len(wstETHContractAddresses), len(wstETHTokenAddresses))
	}

	contractToTokenMapping := make(map[common.Address]common.Address)
	l := min(len(wstETHContractAddresses), len(wstETHTokenAddresses))
	for i := 0; i < l; i++ {
		if wstETHTokenAddresses[i] == emptyAddress {
			continue
		}
		contractToTokenMapping[wstETHContractAddresses[i]] = wstETHTokenAddresses[i]
	}

	if len(contractToTokenMapping) > 0 {
		processorMap[WstETH] = &Processor{
			contractToTokenMapping: contractToTokenMapping,
			contractAddressList:    wstETHContractAddresses,
			tokenAddressList:       wstETHTokenAddresses,
			DecodeMetadataFn: func(metadata []byte) (common.Address, *big.Int) {
				// Metadata structure:
				// - Destination address: 32 bytes
				// - Bridging amount: 32 bytes
				// Maybe there's a more elegant way?
				return common.BytesToAddress(metadata[:32]), new(big.Int).SetBytes(metadata[32:]) //nolint:gomnd
			},
		}
	}
}
