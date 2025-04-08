package messagebridge

import (
	"math/big"

	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/ethereum/go-ethereum/common"
)

func InitEURCProcessor(eurcContractAddresses, eurcTokenAddresses []common.Address) {
	log.Debugf("EURCMapping: contracts[%v] tokens[%v]", eurcContractAddresses, eurcTokenAddresses)
	if len(eurcContractAddresses) != len(eurcTokenAddresses) {
		log.Errorf("InitEURCProcessor: contract addresses (%v) and token addresses (%v) have different length", len(eurcContractAddresses), len(eurcTokenAddresses))
	}

	contractToTokenMapping := make(map[common.Address]common.Address)
	l := min(len(eurcContractAddresses), len(eurcTokenAddresses))
	for i := 0; i < l; i++ {
		if eurcTokenAddresses[i] == emptyAddress {
			continue
		}
		contractToTokenMapping[eurcContractAddresses[i]] = eurcTokenAddresses[i]
	}

	if len(contractToTokenMapping) > 0 {
		processorMap[EURC] = &Processor{
			contractToTokenMapping: contractToTokenMapping,
			contractAddressList:    eurcContractAddresses,
			tokenAddressList:       eurcTokenAddresses,
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
