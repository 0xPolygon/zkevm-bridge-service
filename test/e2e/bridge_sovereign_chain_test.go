//go:build sovereignchain
// +build sovereignchain

package e2e

import (
	"context"
	"math/big"
	"testing"

	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl"
	"github.com/0xPolygonHermez/zkevm-bridge-service/db"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/server"
	"github.com/0xPolygonHermez/zkevm-bridge-service/test/operations"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSovereignChainE2E tests the flow of deposit and withdraw funds in a sovereign chain
func TestSovereignChainE2E(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()
	// rpcURL := "http://localhost:8123"
	// bridgeAddress := common.HexToAddress("0x71C95911E9a5D330f4D621842EC243EE1343292e")
	l2PolygonZkEVMGlobalExitRootAddress := common.HexToAddress("0x712516e61C8B383dF4A63CFe83d7701Bce54B03e")
	opsCfg := &operations.Config{
		L1NetworkURL: "http://localhost:8545",
		L2NetworkURL: "http://localhost:8123",
		L2NetworkID:  1,
		Storage: db.Config{
			Database: "postgres",
			Name:     "test_db",
			User:     "test_user",
			Password: "test_password",
			Host:     "localhost",
			Port:     "5435",
			MaxConns: 10,
		},
		BT: bridgectrl.Config{
			Store:  "postgres",
			Height: uint8(32),
		},
		BS: server.Config{
			GRPCPort:         "9090",
			HTTPPort:         "8080",
			CacheSize:        100000,
			DefaultPageLimit: 25,
			MaxPageLimit:     100,
			BridgeVersion:    "v1",
			DB: db.Config{
				Database: "postgres",
				Name:     "test_db",
				User:     "test_user",
				Password: "test_password",
				Host:     "localhost",
				Port:     "5435",
				MaxConns: 10,
			},
		},
	}
	opsman, err := operations.NewManager(ctx, opsCfg)
	require.NoError(t, err)

	t.Run("L1-L2 eth bridge", func(t *testing.T) {
		// Check initial globalExitRoot. Must fail because at the beginning, no globalExitRoot event is thrown.
		globalExitRootSMC, err := opsman.GetCurrentGlobalExitRootFromSmc(ctx)
		require.NoError(t, err)
		log.Debugf("initial globalExitRootSMC: %+v,", globalExitRootSMC)
		// Send L1 deposit
		var destNetwork uint32 = 1
		amount := new(big.Int).SetUint64(10000000000000000000)
		tokenAddr := common.Address{} // This means is eth
		destAddr := common.HexToAddress("0xc949254d682d8c9ad5682521675b8f43b102aec4")

		// Check L2 funds
		balance, err := opsman.CheckAccountBalance(ctx, operations.L2, &destAddr)
		require.NoError(t, err)
		initL2Balance := big.NewInt(0)
		require.Equal(t, 0, balance.Cmp(initL2Balance))
		err = opsman.SendL1Deposit(ctx, tokenAddr, amount, destNetwork, &destAddr)
		require.NoError(t, err)

		// Check globalExitRoot
		globalExitRoot2, err := opsman.GetTrustedGlobalExitRootSynced(ctx, destNetwork)
		require.NoError(t, err)
		log.Debugf("Before deposit global exit root: %v", globalExitRootSMC)
		log.Debugf("After deposit global exit root: %v", globalExitRoot2)
		require.NotEqual(t, globalExitRootSMC.ExitRoots[0], globalExitRoot2.ExitRoots[0])
		require.Equal(t, globalExitRootSMC.ExitRoots[1], globalExitRoot2.ExitRoots[1])
		// Get Bridge Info By DestAddr
		deposits, err := opsman.GetBridgeInfoByDestAddr(ctx, &destAddr)
		require.NoError(t, err)
		t.Log("Deposit: ", deposits[0])
		// Check the claim tx
		err = opsman.CheckClaim(ctx, deposits[0])
		require.NoError(t, err)
		// Check L2 funds to see if the amount has been increased
		balance2, err := opsman.CheckAccountBalance(ctx, operations.L2, &destAddr)
		require.NoError(t, err)
		require.NotEqual(t, balance, balance2)

		//require.Equal(t, amount, balance2)
		require.Equal(t, 0, balance2.Sub(balance2, initL2Balance).Cmp(amount))

		// // Check globalExitRoot
		// globalExitRoot3, err := opsman.GetCurrentGlobalExitRootFromSmc(ctx)
		// require.NoError(t, err)
		// // Send L2 Deposit to withdraw the some funds
		// destNetwork = 0
		// amount = new(big.Int).SetUint64(1000000000000000000)
		// err = opsman.SendCustomDeposit(ctx, rpcURL, bridgeAddress, tokenAddr, amount, destNetwork, &destAddr, operations.L2)
		// require.NoError(t, err)

		// // Get Bridge Info By DestAddr
		// deposits, err = opsman.GetBridgeInfoByDestAddr(ctx, &destAddr)
		// require.NoError(t, err)
		// log.Debugf("Deposit 2: ", deposits[0])
		// // Check globalExitRoot
		// globalExitRoot4, err := opsman.GetLatestGlobalExitRootFromL1(ctx)
		// require.NoError(t, err)
		// log.Debugf("Global3 %+v: ", globalExitRoot3)
		// log.Debugf("Global4 %+v: ", globalExitRoot4)
		// require.NotEqual(t, globalExitRoot3.ExitRoots[1], globalExitRoot4.ExitRoots[1])
		// require.Equal(t, globalExitRoot3.ExitRoots[0], globalExitRoot4.ExitRoots[0])
		// // Check L1 funds
		// balance, err = opsman.CheckAccountBalance(ctx, operations.L1, &destAddr)
		// require.NoError(t, err)
		// require.Equal(t, 0, big.NewInt(0).Cmp(balance))
		// // Get the claim data
		// smtProof, smtRollupProof, globaExitRoot, err := opsman.GetClaimData(ctx, deposits[0].NetworkId, deposits[0].DepositCnt)
		// require.NoError(t, err)

		// // Claim funds in L1
		// err = opsman.SendL1Claim(ctx, deposits[0], smtProof, smtRollupProof, globaExitRoot)
		// require.NoError(t, err)

		// // Check L1 funds to see if the amount has been increased
		// balance, err = opsman.CheckAccountBalance(ctx, operations.L1, &destAddr)
		// require.NoError(t, err)
		// require.Equal(t, big.NewInt(1000000000000000000), balance)
		// // Check L2 funds to see that the amount has been reduced
		// balance, err = opsman.CheckAccountBalance(ctx, operations.L2, &destAddr)
		// require.NoError(t, err)
		// require.True(t, big.NewInt(9000000000000000000).Cmp(balance) > 0)
		log.Debug("L1-L2 eth bridge end")
	})
	t.Run("Remove GER bridge", func(t *testing.T) {
		// Check globalExitRoot
		var networkID uint32 = 1
		globalExitRoot, err := opsman.GetTrustedGlobalExitRootSynced(ctx, networkID)
		require.NoError(t, err)

		// Remove GER bridge transaction
		err = opsman.RemoveL2GER(ctx, l2PolygonZkEVMGlobalExitRootAddress, []common.Hash{globalExitRoot.GlobalExitRoot}, networkID, operations.L222)
		require.NoError(t, err)

		globalExitRoot2, err := opsman.GetTrustedGlobalExitRootSynced(ctx, networkID)
		require.NoError(t, err)
		require.NotEqual(t, globalExitRoot.ExitRoots[0], globalExitRoot2.ExitRoots[0])
		require.Equal(t, globalExitRoot.ExitRoots[1], globalExitRoot2.ExitRoots[1])
		require.NotEqual(t, globalExitRoot.GlobalExitRoot, globalExitRoot2.GlobalExitRoot)

		// Get Bridge Info By DestAddr
		destAddr := common.HexToAddress("0xc949254d682d8c9ad5682521675b8f43b102aec4")
		deposits, err := opsman.GetBridgeInfoByDestAddr(ctx, &destAddr)
		require.NoError(t, err)
		assert.Equal(t, 1, len(deposits))
		assert.Equal(t, false, deposits[0].ReadyForClaim)
		assert.Equal(t, "0xC949254d682D8c9ad5682521675b8F43b102aec4", deposits[0].DestAddr)

		log.Debug("Remove GER bridge end")
	})
}
