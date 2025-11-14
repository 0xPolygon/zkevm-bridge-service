package main

import (
	"context"
	"math/big"
	"time"

	"github.com/0xPolygon/zkevm-bridge-service/etherman/smartcontracts/bridgel2sovereignchain"
	"github.com/0xPolygon/zkevm-bridge-service/log"
	"github.com/0xPolygon/zkevm-bridge-service/bridgectrl"
	"github.com/0xPolygon/zkevm-bridge-service/test/operations"
	"github.com/0xPolygon/zkevm-bridge-service/db"
	"github.com/0xPolygon/zkevm-bridge-service/db/pgstorage"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/0xPolygon/zkevm-bridge-service/utils"
	"github.com/0xPolygon/zkevm-bridge-service/server"
	"github.com/ethereum/go-ethereum/common"
)

const (
	l2BridgeAddr = "0x71C95911E9a5D330f4D621842EC243EE1343292e"

	l2AccHexPrivateKey = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	l2NetworkURL       = "http://localhost:8123"

	index = 1
	l2NetworkID = 1
)

func main() {
	ctx := context.Background()
	l2BridgeAddress := common.HexToAddress(l2BridgeAddr)
	opsCfg := &operations.Config{
		L1NetworkURL: "http://localhost:8545",
		L2NetworkURL: "http://localhost:8123",
		L2NetworkID: l2NetworkID,
		Storage: db.Config{
			Database: "postgres",
			PgStorage: pgstorage.Config{
				Name:     "test_db",
				User:     "test_user",
				Password: "test_password",
				Host:     "localhost",
				Port:     "5435",
				MaxConns: 10, //nolint:mnd
			},
		},
		BT: bridgectrl.Config{
			Height: uint8(32), //nolint:mnd
		},
		BS: server.Config{
			GRPCPort:         "9090",
			HTTPPort:         "8080",
			CacheSize:        100000, //nolint:mnd
			DefaultPageLimit: 25, //nolint:mnd
			MaxPageLimit:     100, //nolint:mnd
			BridgeVersion:    "v1",
			DB: db.Config{
				Database: "postgres",
				PgStorage: pgstorage.Config{
					Name:     "test_db",
					User:     "test_user",
					Password: "test_password",
					Host:     "localhost",
					Port:     "5435",
					MaxConns: 10, //nolint:mnd
				},
			},
		},
	}
	manager, err := operations.NewManager(ctx, opsCfg)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	c, err := utils.NewClient(ctx, l2NetworkURL, l2BridgeAddress)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	auth, err := c.GetSigner(ctx, l2AccHexPrivateKey)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	br, err := bridgel2sovereignchain.NewBridgel2sovereignchain(l2BridgeAddress, c.Client)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	nextLeaf, deposit, err := manager.GetLeaf(ctx, index+1, l2NetworkID)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	rollupMerkleProof, err := manager.GetProof(ctx, deposit.NetworkID, deposit.DepositCount)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	frontier := manager.ComputeFrontierFromProof(index, rollupMerkleProof)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	// Enable emergencyState
	isEmergencyState, err := br.IsEmergencyState(&bind.CallOpts{Pending: false})
	if err != nil {
		log.Fatal("Error: ", err)
	}
	if !isEmergencyState {
		tx, err := br.ActivateEmergencyState(auth)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		log.Info("EmergencyState TxHash: ", tx.Hash())
		time.Sleep(5 * time.Second) //nolint:mnd
	}
	log.Info("Sending claim tx...")
	tx, err := br.BackwardLET(auth, big.NewInt(index), frontier, nextLeaf, rollupMerkleProof)
	if err != nil {
		a, _ := bridgel2sovereignchain.Bridgel2sovereignchainMetaData.GetAbi()
		input, err3 := a.Pack("backwardLET", big.NewInt(index), frontier, nextLeaf, rollupMerkleProof)
		if err3 != nil {
			log.Error("error packing call. Error: ", err3)
		}
		log.Warnf(`Use the next command to debug it manually.
		curl --location --request POST 'http://localhost:8123' \
		--header 'Content-Type: application/json' \
		--data-raw '{
			"jsonrpc": "2.0",
			"method": "eth_call",
			"params": [{"from": "%s","to":"%s","data":"0x%s"},"latest"],
			"id": 1
		}'`, auth.From, l2BridgeAddr, common.Bytes2Hex(input))
		log.Fatal("error: ", err)
	}
	log.Info("BackwardLET TxHash: ", tx.Hash())

	// Disable emergencyState
	tx, err = br.DeactivateEmergencyState(auth)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Info("Disable EmergencyState TxHash: ", tx.Hash())
}
