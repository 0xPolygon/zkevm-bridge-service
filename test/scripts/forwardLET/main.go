package main

import (
	"context"
	"math/big"
	"time"

	"github.com/0xPolygon/zkevm-bridge-service/etherman/smartcontracts/bridgel2sovereignchain"
	"github.com/0xPolygon/zkevm-bridge-service/log"
	"github.com/0xPolygon/zkevm-bridge-service/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const (
	l2BridgeAddr = "0x71C95911E9a5D330f4D621842EC243EE1343292e"

	destinationAddress = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	l2AccHexPrivateKey = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	l2NetworkURL       = "http://localhost:8123"

	expectedRoot = "0x1a343d64e1fec5f90636f6b2fa44858c4dd057c36c88956e2b00054546f8e804"
)

func main() {
	ctx := context.Background()
	l2BridgeAddress := common.HexToAddress(l2BridgeAddr)
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

	log.Info("Sending claim tx...")
	newLeaves := []bridgel2sovereignchain.AgglayerBridgeL2LeafData {
		{
			LeafType:           0,
			OriginNetwork:      0,      
			OriginAddress:      common.Address{},
			DestinationNetwork: 0,
			DestinationAddress: common.HexToAddress(destinationAddress),
			Amount:             big.NewInt(1),
			Metadata:           []byte{},
		},
		{
			LeafType:           0,
			OriginNetwork:      0,      
			OriginAddress:      common.Address{},
			DestinationNetwork: 0,
			DestinationAddress: common.HexToAddress(destinationAddress),
			Amount:             big.NewInt(1),
			Metadata:           []byte{},
		},
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
	tx, err := br.ForwardLET(auth, newLeaves, common.HexToHash(expectedRoot))
	if err != nil {
		a, _ := bridgel2sovereignchain.Bridgel2sovereignchainMetaData.GetAbi()
		input, err3 := a.Pack("forwardLET", newLeaves, common.HexToHash(expectedRoot))
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
	log.Info("ForwardLET TxHash: ", tx.Hash())

	// Disable emergencyState
	tx, err = br.DeactivateEmergencyState(auth)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Info("Disable EmergencyState TxHash: ", tx.Hash())
}
