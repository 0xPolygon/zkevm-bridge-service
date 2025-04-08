package messagebridge

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestUSDCLxLyMapping(t *testing.T) {
	contractAddr1 := common.HexToAddress("0xfe3240995c771f10D2583e8fa95F92ee40E15150")
	contractAddr2 := common.HexToAddress("0x1A8C4999D32F05B63A227517Be0824AeD47e4728")
	contractAddr3 := common.HexToAddress("0xfe3240995c771f10D2583e8fa95F92ee40E15151")
	tokenAddr1 := common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	tokenAddr2 := common.HexToAddress("0x00d69D72a429d4985b34A8E1A6C9e47997F0aFA3")

	InitUSDCLxLyProcessor([]common.Address{contractAddr1, contractAddr2}, []common.Address{tokenAddr1, tokenAddr2})
	require.Len(t, processorMap, 1)
	processor := GetProcessorByType(USDC)
	require.NotNil(t, processor)

	list := processor.GetContractAddressList()
	require.Len(t, list, 2)
	require.Contains(t, list, contractAddr1)
	require.Contains(t, list, contractAddr2)

	require.True(t, processor.CheckContractAddress(contractAddr1))
	require.True(t, processor.CheckContractAddress(contractAddr2))
	require.False(t, processor.CheckContractAddress(contractAddr3))

	token, ok := processor.GetTokenFromContract(contractAddr2)
	require.True(t, ok)
	require.Equal(t, tokenAddr2, token)

	_, ok = processor.GetTokenFromContract(contractAddr3)
	require.False(t, ok)
}
