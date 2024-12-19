//go:build autoclaiml2l2
// +build autoclaiml2l2

package e2e

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/0xPolygonHermez/zkevm-bridge-service/test/operations"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// TestAutoClaimL2L2 tests the flow of deposit and withdraw funds using the vector
func TestAutoClaimL2L2(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	err := os.Setenv("ZKEVM_BRIDGE_CLAIMTXMANAGER_ARECLAIMSBETWEENL2SENABLED", "true")
	require.NoError(t, err)
	require.NoError(t, operations.StartBridge3())
	ctx := context.Background()
	opsman1, err := operations.GetOpsman(ctx, "http://localhost:8123", "test_db", "8080", "9090", "5435", 1)
	require.NoError(t, err)
	opsman2, err := operations.GetOpsman(ctx, "http://localhost:8124", "test_db", "8080", "9090", "5435", 2)
	require.NoError(t, err)

	t.Run("AutoClaim L2-L2 eth bridge", func(t *testing.T) {
		// Check initial globalExitRoot. Must fail because at the beginning, no globalExitRoot event is thrown.
		globalExitRootSMC, err := opsman1.GetCurrentGlobalExitRootFromSmc(ctx)
		require.NoError(t, err)
		t.Logf("initial globalExitRootSMC: %+v,", globalExitRootSMC)
		// Send L2 deposit
		var destNetwork uint32 = 2
		amount := new(big.Int).SetUint64(10000000000000000001)
		tokenAddr := common.Address{} // This means is eth
		address := common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")

		l2Balance, err := opsman1.CheckAccountBalance(ctx, operations.L2, &address)
		require.NoError(t, err)
		t.Logf("Initial L2 Bridge Balance in origin network 1: %v", l2Balance)
		err = opsman1.SendL2Deposit(ctx, tokenAddr, amount, destNetwork, &address, operations.L22)
		require.NoError(t, err)
		l2Balance, err = opsman1.CheckAccountBalance(ctx, operations.L2, &address)
		require.NoError(t, err)
		t.Logf("Final L2 Bridge Balance in origin network 1: %v", l2Balance)

		// Check globalExitRoot
		globalExitRoot, err := opsman1.GetLatestGlobalExitRootFromL1(ctx)
		require.NoError(t, err)
		t.Logf("GlobalExitRoot %+v: ", globalExitRoot)
		require.NotEqual(t, globalExitRoot.ExitRoots[1], globalExitRootSMC.ExitRoots[1])
		require.Equal(t, globalExitRoot.ExitRoots[0], globalExitRootSMC.ExitRoots[0])
		// Check L2 destination funds
		balance, err := opsman2.CheckAccountBalance(ctx, operations.L2, &address)
		require.NoError(t, err)
		v, _ := big.NewInt(0).SetString("99999998433970000000000", 10)
		t.Log("balance: ", balance)
		require.Equal(t, 0,  v.Cmp(balance))
		// This deposit forces the update of the ger to process the previous ready for claim. It is
		// needed because of the race condition between both claimtxmanagers (network 1 and network 2). Both claimTxManagers
		// run at the same time and network 2 checks if there are some deposit ready for claim before the dbTx of
		// claimTxManager (network 1) is commited. With this second deposit we force another ger and claimTxManager rechecks
		// if there is some L2Deposit for claim.
		err = opsman1.SendL2Deposit(ctx, tokenAddr, amount, 0, &address, operations.L22)
		require.NoError(t, err)
		// Wait until the claimTxManager claims the first deposit.
		time.Sleep(30 * time.Second)

		// Check destination L2 funds to see if the amount has been increased
		balance, err = opsman2.CheckAccountBalance(ctx, operations.L2, &address)
		require.NoError(t, err)
		require.Equal(t, -1, v.Cmp(balance))
		// Check origin L2 funds to see that the amount has been reduced
		balance, err = opsman1.CheckAccountBalance(ctx, operations.L2, &address)
		require.NoError(t, err)
		require.Equal(t, 1, v.Cmp(balance))
	})
}
