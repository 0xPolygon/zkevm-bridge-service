package main

import (
	"context"
	"math/big"

	"github.com/0xPolygon/zkevm-bridge-service/log"
	"github.com/0xPolygon/zkevm-bridge-service/utils"
	"github.com/0xPolygon/zkevm-bridge-service/etherman/smartcontracts/bridgel2sovereignchain"
	"github.com/ethereum/go-ethereum/common"
)

const (
	l2BridgeAddr = "0x71C95911E9a5D330f4D621842EC243EE1343292e"

	l1AccHexAddress    = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	l2AccHexPrivateKey = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	l2NetworkURL       = "http://localhost:8123"

	funds              = 1 // nolint
	destNetwork uint32 = 0
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
	amount := big.NewInt(0).SetInt64(funds)
	auth.Value = amount
	tx, err := br.BridgeAsset(auth, destNetwork, common.HexToAddress(l1AccHexAddress), amount, common.Address{}, true, []byte{})
	if err != nil {
		a, _ := bridgel2sovereignchain.Bridgel2sovereignchainMetaData.GetAbi()
		input, err3 := a.Pack("bridgeAsset", destNetwork, common.HexToAddress(l1AccHexAddress), amount, common.Address{}, true, []byte{})
		if err3 != nil {
			log.Error("error packing call. Error: ", err3)
		}
		log.Warnf(`Use the next command to debug it manually.
		curl --location --request POST 'http://localhost:8123' \
		--header 'Content-Type: application/json' \
		--data-raw '{
			"jsonrpc": "2.0",
			"method": "eth_call",
			"params": [{"from": "%s","to":"%s","value":"0x%s","data":"0x%s"},"latest"],
			"id": 1
		}'`, auth.From, l2BridgeAddr, amount.Text(16), common.Bytes2Hex(input)) //nolint:mnd
		log.Fatal("error: ", err)
	}
	log.Info("TxHash: ", tx.Hash())
	log.Info("Success!")
}
