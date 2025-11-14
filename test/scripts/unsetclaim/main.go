package main

import (
	"context"
	"math/big"

	"github.com/0xPolygon/zkevm-bridge-service/etherman/smartcontracts/bridgel2sovereignchain"
	"github.com/0xPolygon/zkevm-bridge-service/log"
	"github.com/0xPolygon/zkevm-bridge-service/utils"
	"github.com/ethereum/go-ethereum/common"
)

const (
	l2BridgeAddr = "0x71C95911E9a5D330f4D621842EC243EE1343292e"

	l2AccHexPrivateKey = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	l2NetworkURL       = "http://localhost:8123"

	globalIndex = "18446744073709551616"
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
	gI, _ := big.NewInt(0).SetString(globalIndex, 0)
	tx, err := br.UnsetMultipleClaims(auth, []*big.Int{gI})
	if err != nil {
		a, _ := bridgel2sovereignchain.Bridgel2sovereignchainMetaData.GetAbi()
		input, err3 := a.Pack("unsetMultipleClaims", []*big.Int{gI})
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
	log.Info("TxHash: ", tx.Hash())
}
