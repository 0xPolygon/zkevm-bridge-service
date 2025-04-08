package utils

type NetworkChainIdMapping struct {
	chainIDs map[uint]uint32
}

var networkChainIdMapping NetworkChainIdMapping

func InitChainIdManager(networks []uint32, chainIds []uint) {
	var chainIDs = make(map[uint]uint32)
	for i, network := range networks {
		chainIDs[uint(network)] = uint32(chainIds[i]) // nolint:gosec
	}
	networkChainIdMapping = NetworkChainIdMapping{
		chainIDs: chainIDs,
	}
}

func GetChainIdByNetworkId(networkId uint) uint32 {
	return networkChainIdMapping.chainIDs[networkId]
}
