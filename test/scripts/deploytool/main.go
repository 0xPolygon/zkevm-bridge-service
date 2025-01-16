package main

import (
	"math/big"
	"os"
	"time"

	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman/smartcontracts/bridgel2sovereignchain"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman/smartcontracts/claimcompressor"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman/smartcontracts/globalexitrootmanagerl2sovereignchain"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils"
	"github.com/0xPolygonHermez/zkevm-node/config/types"
	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/proxy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/urfave/cli/v2"
)

const (
	flagURLName        = "url"
	flagBridgeAddrName = "bridgeAddress"
	flagDestAddrName   = "destAddress"
	flagWalletFileName = "walletFile"
	flagPasswordName   = "password"
)

var (
	flagURL = cli.StringFlag{
		Name:     flagURLName,
		Aliases:  []string{"u"},
		Usage:    "Node url",
		Required: true,
	}
	flagBridgeAddr = cli.StringFlag{
		Name:     flagBridgeAddrName,
		Aliases:  []string{"br"},
		Usage:    "Bridge smart contract address",
		Required: true,
	}
	flagWalletFile = cli.StringFlag{
		Name:     flagWalletFileName,
		Aliases:  []string{"f"},
		Usage:    "Wallet file",
		Required: true,
	}
	flagPassword = cli.StringFlag{
		Name:     flagPasswordName,
		Aliases:  []string{"pass"},
		Usage:    "Password",
		Required: true,
	}
	flagDestAddr = cli.StringFlag{
		Name:     flagDestAddrName,
		Aliases:  []string{"dest"},
		Usage:    "Destination address",
		Required: true,
	}
)

func main() {
	claimCompDeployer := cli.NewApp()
	claimCompDeployer.Name = "deploy smc tool"
	claimCompDeployer.Usage = "deploy smc for testing purposes"
	claimCompDeployer.DefaultCommand = "deployClaimCompressor"
	flags := []cli.Flag{&flagURL, &flagWalletFile, &flagPassword}
	claimCompDeployer.Commands = []*cli.Command{
		{
			Name:    "deployClaimCompressor",
			Aliases: []string{},
			Flags:   append(flags, &flagBridgeAddr),
			Action:  deployClaimCompressor,
		},
		{
			Name:    "deploySovereignChainSMC",
			Aliases: []string{},
			Flags:   flags,
			Action:  deploySovereignChainSMC,
		},
		{
			Name:    "sendETH",
			Aliases: []string{},
			Flags:   append(flags, &flagDestAddr),
			Action:  sendETH,
		},
	}

	err := claimCompDeployer.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
func deployClaimCompressor(ctx *cli.Context) error {
	l2NetworkURL := ctx.String(flagURLName)
	l2BridgeAddr := ctx.String(flagBridgeAddrName)
	c, err := utils.NewClient(ctx.Context, l2NetworkURL, common.HexToAddress(l2BridgeAddr))
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	privateKey := types.KeystoreFileConfig{
		Path:     ctx.String(flagWalletFileName),
		Password: ctx.String(flagPasswordName),
	}
	log.Debug(privateKey)
	auth, err := c.GetSignerFromKeystore(ctx.Context, privateKey)
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	networkID, err := c.Bridge.NetworkID(&bind.CallOpts{Pending: false})
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	log.Debug("networkID: ", networkID)
	log.Debug("auth.From: ", auth.From)
	balance, err := c.Client.BalanceAt(ctx.Context, auth.From, nil)
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	log.Debug("balance: ", balance)
	claimCompressorAddress, _, _, err := claimcompressor.DeployClaimcompressor(auth, c.Client, common.HexToAddress(l2BridgeAddr), networkID)
	if err != nil {
		log.Error("Error deploying claimCompressor contract: ", err)
		return err
	}
	log.Info("ClaimCompressorAddress: ", claimCompressorAddress)
	return nil
}

func deploySovereignChainSMC(ctx *cli.Context) error {
	l2NetworkURL := ctx.String(flagURLName)
	c, err := utils.NewClient(ctx.Context, l2NetworkURL, common.Address{})
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	privateKey := types.KeystoreFileConfig{
		Path:     ctx.String(flagWalletFileName),
		Password: ctx.String(flagPasswordName),
	}
	log.Debug(privateKey)
	auth, err := c.GetSignerFromKeystore(ctx.Context, privateKey)
	if err != nil {
		log.Error("Error: ", err)
		return err
	}

	log.Debug("auth.From: ", auth.From)
	balance, err := c.Client.BalanceAt(ctx.Context, auth.From, nil)
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	log.Debug("balance: ", balance)

	implementationBridgeAddr, _, _, err := bridgel2sovereignchain.DeployBridgel2sovereignchain(auth, c.Client)
	if err != nil {
		log.Error("error: ", err)
		return err
	}
	log.Info("ImplementationBridgeAddr: ", implementationBridgeAddr)
	time.Sleep(3 * time.Second) //nolint:gomnd

	bridgeAddr, _, _, err := proxy.DeployProxy(auth, c.Client, implementationBridgeAddr, implementationBridgeAddr, []byte{})
	if err != nil {
		log.Error("error: ", err)
		return err
	}
	log.Info("BridgeAddr: ", bridgeAddr)
	time.Sleep(3 * time.Second) //nolint:gomnd

	implementationGERManagerAddr, _, _, err := globalexitrootmanagerl2sovereignchain.DeployGlobalexitrootmanagerl2sovereignchain(auth, c.Client, bridgeAddr)
	if err != nil {
		log.Error("error: ", err)
		return err
	}
	log.Info("ImplementationGERManagerAddr: ", implementationGERManagerAddr)
	time.Sleep(3 * time.Second) //nolint:gomnd

	globalExitRootAddr, _, _, err := proxy.DeployProxy(auth, c.Client, implementationGERManagerAddr, implementationGERManagerAddr, []byte{})
	if err != nil {
		log.Error("error: ", err)
		return err
	}
	log.Info("GlobalExitRootAddr: ", globalExitRootAddr)
	time.Sleep(3 * time.Second) //nolint:gomnd

	br, err := bridgel2sovereignchain.NewBridgel2sovereignchain(bridgeAddr, c.Client)
	if err != nil {
		log.Error("error: ", err)
		return err
	}
	_, err = br.Initialize(auth, 1, common.Address{}, 0, globalExitRootAddr, auth.From, []byte{}, auth.From, common.Address{}, false)
	if err != nil {
		log.Error("error: ", err)
		return err
	}
	time.Sleep(3 * time.Second) //nolint:gomnd

	ger, err := globalexitrootmanagerl2sovereignchain.NewGlobalexitrootmanagerl2sovereignchain(globalExitRootAddr, c.Client)
	if err != nil {
		log.Error("error: ", err)
		return err
	}
	_, err = ger.Initialize(auth, auth.From, auth.From)
	if err != nil {
		log.Error("error: ", err)
		return err
	}
	log.Info("Success!! Contracts initiliazed")

	return nil
}

func sendETH(ctx *cli.Context) error {
	l2NetworkURL := ctx.String(flagURLName)
	destAddress := ctx.String(flagDestAddrName)
	c, err := utils.NewClient(ctx.Context, l2NetworkURL, common.Address{})
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	privateKey := types.KeystoreFileConfig{
		Path:     ctx.String(flagWalletFileName),
		Password: ctx.String(flagPasswordName),
	}
	log.Debug(privateKey)
	key, chainID, err := c.GetKeyFromKeystore(ctx.Context, privateKey)
	if err != nil {
		log.Error("Error: ", err)
		return err
	}

	log.Debug("destAddress: ", destAddress)
	destAddr := common.HexToAddress(destAddress)
	balance, err := c.Client.BalanceAt(ctx.Context, destAddr, nil)
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	log.Debug("init balance: ", balance)

	fromAddress := key.Address
	nonce, err := c.Client.PendingNonceAt(ctx.Context, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value, _ := big.NewInt(0).SetString("20000000000000000000", 0) // in wei (20 eth)
	gasLimit := uint64(21000)                                      //nolint:gomnd
	gasPrice, err := c.Client.SuggestGasPrice(ctx.Context)
	if err != nil {
		log.Fatal(err)
	}

	var data []byte
	tx := gethTypes.NewTransaction(nonce, destAddr, value, gasLimit, gasPrice, data)

	signedTx, err := gethTypes.SignTx(tx, gethTypes.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Client.SendTransaction(ctx.Context, signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("tx sent: %s", signedTx.Hash().Hex())

	time.Sleep(3 * time.Second) //nolint:gomnd
	balance, err = c.Client.BalanceAt(ctx.Context, destAddr, nil)
	if err != nil {
		log.Error("Error: ", err)
		return err
	}
	log.Debug("final balance: ", balance)
	return nil
}
