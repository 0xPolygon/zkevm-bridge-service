package main

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"

	"github.com/0xPolygonHermez/zkevm-bridge-service/bridgectrl"
	"github.com/0xPolygonHermez/zkevm-bridge-service/claimtxman"
	"github.com/0xPolygonHermez/zkevm-bridge-service/config"
	"github.com/0xPolygonHermez/zkevm-bridge-service/db"
	"github.com/0xPolygonHermez/zkevm-bridge-service/etherman"
	"github.com/0xPolygonHermez/zkevm-bridge-service/log"
	"github.com/0xPolygonHermez/zkevm-bridge-service/server"
	"github.com/0xPolygonHermez/zkevm-bridge-service/synchronizer"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils"
	"github.com/0xPolygonHermez/zkevm-bridge-service/utils/gerror"

	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/coinmiddleware"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/estimatetime"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/localcache"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/pushtask"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/redisstorage"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/sentinel"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/utils/messagebridge"

	client "github.com/0xPolygonHermez/zkevm-bridge-service/jsonrpcclient"
	xlayerUtils "github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/utils"
)

const (
	api  = "API"
	push = "PUSH"
	task = "TASK"
)

func run(ctx *cli.Context, choice string) error {
	var err error

	c, err := setupConfigAndLog(ctx)
	if err != nil {
		return err
	}

	go enableMetrics(c)

	switch choice {
	case api:
		if err = db.RunMigrations(c.UpstreamCfg.SyncDB); err != nil {
			return err
		}
		if err = runAPI(ctx.Context, c); err != nil {
			return err
		}
		waitUnlessInterrupt()
		return nil
	case push:
		if err = runPushTask(ctx.Context, c); err != nil {
			return err
		}
		waitUnlessInterrupt()
		return nil
	case task:
		return runTask(ctx.Context, c)
	default:
		if err = db.RunMigrations(c.UpstreamCfg.SyncDB); err != nil {
			return err
		}
		if err = runAPI(ctx.Context, c); err != nil {
			return err
		}
		if err = runPushTask(ctx.Context, c); err != nil {
			return err
		}
		return runTask(ctx.Context, c) // Blocking
	}
}

func runAPI(ctx context.Context, c *config.XLayerConfig) error {
	// Some API use this information to filter deposits.
	messagebridge.InitUSDCLxLyProcessor(c.BusinessConfig.USDCContractAddresses, c.BusinessConfig.USDCTokenAddresses)
	messagebridge.InitWstETHProcessor(c.BusinessConfig.WstETHContractAddresses, c.BusinessConfig.WstETHTokenAddresses)
	messagebridge.InitEURCProcessor(c.BusinessConfig.EURCContractAddresses, c.BusinessConfig.EURCTokenAddresses)

	redisStorage, err := redisstorage.NewRedisStorage(c.BridgeServer.Redis)
	if err != nil {
		return err
	}

	apiStorage, err := db.NewStorage(c.UpstreamCfg.BridgeServer.DB)
	if err != nil {
		return err
	}

	if err = localcache.InitDefaultCache(apiStorage); err != nil {
		return err
	}

	if err = estimatetime.InitDefaultCalculator(apiStorage, c.EstimateTime); err != nil {
		return err
	}

	l1Etherman, l2Ethermans, err := newEthermans(c.UpstreamCfg)
	if err != nil {
		return err
	}

	networkID := l1Etherman.GetNetworkID()
	networkIDs := setupNetworkIDs(networkID, l2Ethermans)

	l1ChainId := c.Etherman.L1ChainId
	l2ChainIds := c.Etherman.L2ChainIds
	var chainIDs = []uint{l1ChainId}
	chainIDs = append(chainIDs, l2ChainIds...)
	xlayerUtils.InitChainIdManager(networkIDs, chainIDs)

	l2NodeClients := make([]*utils.Client, len(c.UpstreamCfg.Etherman.L2URLs))
	l2Auths := make([]*bind.TransactOpts, len(c.UpstreamCfg.Etherman.L2URLs))
	for i := range c.UpstreamCfg.Etherman.L2URLs {
		nodeClient, err := utils.NewClient(
			ctx,
			c.UpstreamCfg.Etherman.L2URLs[i],
			c.UpstreamCfg.NetworkConfig.L2PolygonBridgeAddresses[i],
		)
		if err != nil {
			return err
		}
		auth, err := nodeClient.GetSignerFromKeystore(ctx, c.UpstreamCfg.ClaimTxManager.PrivateKey)
		if err != nil {
			return err
		}
		l2NodeClients[i] = nodeClient
		l2Auths[i] = auth
	}

	registerNacos(c.NacosConfig)

	if err := sentinel.InitFileDataSource(c.BridgeServer.SentinelConfigFilePath); err != nil {
		return err
	}

	bridgeService := server.NewBridgeService(c.UpstreamCfg.BridgeServer, c.UpstreamCfg.BridgeController.Height, networkIDs, apiStorage).
		WithRedisStorage(redisStorage).
		WithMainCoinsCache(localcache.GetDefaultCache()).
		WithL2Clients(l2NodeClients, l2Auths, networkIDs)
	bridgeService.LogConfig()

	return server.RunServer(c.UpstreamCfg.BridgeServer, bridgeService) // non-blocking
}

func runPushTask(ctx context.Context, c *config.XLayerConfig) error {
	apiStorage, err := db.NewStorage(c.UpstreamCfg.BridgeServer.DB)
	if err != nil {
		return err
	}

	redisStorage, err := redisstorage.NewRedisStorage(c.BridgeServer.Redis)
	if err != nil {
		return err
	}

	messagePushProducer, _, err := setupKafkaProducer(c.MessagePushProducer)
	if err != nil {
		return err
	}

	l1Etherman, err := etherman.NewClient(c.UpstreamCfg.Etherman,
		c.UpstreamCfg.NetworkConfig.PolygonBridgeAddress,
		c.UpstreamCfg.NetworkConfig.PolygonZkEVMGlobalExitRootAddress,
		c.UpstreamCfg.NetworkConfig.PolygonRollupManagerAddress)
	rollupID, err := l1Etherman.PolygonRollupManager.RollupAddressToID(
		&bind.CallOpts{Pending: false},
		c.UpstreamCfg.NetworkConfig.PolygonRollupManagerAddress,
	)
	log.Info("push service: RollupID =", rollupID)

	// Initialize the push task for L1 block num change
	l1BlockNumTask, err := pushtask.NewL1BlockNumTask(
		c.UpstreamCfg.Etherman.L1URL, apiStorage, redisStorage, messagePushProducer, uint(rollupID))
	if err != nil {
		return err
	}

	// Initialize the push task for sync l2 commit batch
	syncCommitBatchTask, err := pushtask.NewCommittedBatchHandler(
		c.UpstreamCfg.Etherman.L2URLs[0], apiStorage, redisStorage, messagePushProducer, uint(rollupID))
	if err != nil {
		return err
	}

	// Initialize the push task for sync verify batch
	syncVerifyBatchTask, err := pushtask.NewVerifiedBatchHandler(c.UpstreamCfg.Etherman.L2URLs[0], redisStorage)
	if err != nil {
		return err
	}

	go l1BlockNumTask.Start(ctx)
	go syncCommitBatchTask.Start(ctx)
	go syncVerifyBatchTask.Start(ctx)

	return nil
}

func runTask(ctx context.Context, c *config.XLayerConfig) error {
	// Use this to run Go routines
	errs, _ := errgroup.WithContext(ctx)

	messagebridge.InitUSDCLxLyProcessor(c.BusinessConfig.USDCContractAddresses, c.BusinessConfig.USDCTokenAddresses)
	messagebridge.InitWstETHProcessor(c.BusinessConfig.WstETHContractAddresses, c.BusinessConfig.WstETHTokenAddresses)
	messagebridge.InitEURCProcessor(c.BusinessConfig.EURCContractAddresses, c.BusinessConfig.EURCTokenAddresses)

	// Initialize all stores
	apiStorage, err := db.NewStorage(c.UpstreamCfg.BridgeServer.DB)
	if err != nil {
		return err
	}

	if err = estimatetime.InitDefaultCalculator(apiStorage, c.EstimateTime); err != nil {
		return err
	}

	storage, err := db.NewStorage(c.UpstreamCfg.SyncDB)
	if err != nil {
		return err
	}

	redisStorage, err := redisstorage.NewRedisStorage(c.BridgeServer.Redis)
	if err != nil {
		return err
	}

	l1Etherman, l2Ethermans, err := newEthermans(c.UpstreamCfg)
	if err != nil {
		return err
	}

	networkID := l1Etherman.GetNetworkID()
	networkIDs := setupNetworkIDs(networkID, l2Ethermans)

	l1ChainId := c.Etherman.L1ChainId
	l2ChainIds := c.Etherman.L2ChainIds
	var chainIDs = []uint{l1ChainId}
	chainIDs = append(chainIDs, l2ChainIds...)
	xlayerUtils.InitChainIdManager(networkIDs, chainIDs)

	var bridgeController *bridgectrl.BridgeController
	if c.UpstreamCfg.BridgeController.Store == "postgres" {
		bridgeController, err = bridgectrl.NewBridgeController(
			ctx, c.UpstreamCfg.BridgeController, networkIDs, storage)
		if err != nil {
			return err
		}
	} else {
		return gerror.ErrStorageNotRegister
	}

	messagePushProducer, closeKafka, err := setupKafkaProducer(c.MessagePushProducer)
	if err != nil {
		return err
	}
	defer closeKafka()

	bridgeService := server.NewBridgeService(c.UpstreamCfg.BridgeServer, c.UpstreamCfg.BridgeController.Height, networkIDs, apiStorage)
	bridgeService.LogConfig()

	chSynced := make(chan uint32)
	var chsExitRootEvent []chan *etherman.GlobalExitRoot
	var chsSyncedL2 []chan uint32
	for i, l2EthermanClient := range l2Ethermans {
		L2URL := c.UpstreamCfg.Etherman.L2URLs[i]
		zkEVMClient := client.NewClient(L2URL)
		rollupID := l2EthermanClient.GetNetworkID() // RollupID == networkID

		chExitRootEventL2 := make(chan *etherman.GlobalExitRoot)
		chsExitRootEvent = append(chsExitRootEvent, chExitRootEventL2)
		chSyncedL2 := make(chan uint32)
		chsSyncedL2 = append(chsSyncedL2, chSyncedL2)

		sync, err := synchronizer.NewSynchronizer(
			ctx, storage, bridgeController,
			l2EthermanClient, zkEVMClient, 0,
			chExitRootEventL2, nil, chSyncedL2,
			c.UpstreamCfg.Synchronizer, []uint32{}, c.UpstreamCfg.NetworkConfig.RequireSovereignChainSmcs[i],
		)
		if err != nil {
			return err
		}
		cliSyncL2, ok := sync.(*synchronizer.ClientSynchronizer)
		if !ok {
			log.Fatalf("Failed to cast synchronizer")
		}
		cliSyncL2 = cliSyncL2.
			SetProducer(messagePushProducer).
			SetRedis(redisStorage).
			SetRollupID(uint(rollupID)).
			SetLargeTxUsdLimit(c.Synchronizer.LargeTxUsdLimit)
		errs.Go(cliSyncL2.Sync)

		if c.UpstreamCfg.ClaimTxManager.Enabled {
			nodeClient, err := utils.NewClient(ctx, L2URL, c.UpstreamCfg.NetworkConfig.L2PolygonBridgeAddresses[i])
			if err != nil {
				return err
			}
			nonceCache, err := claimtxman.NewNonceCache(ctx, nodeClient)
			if err != nil {
				log.Fatalf("error creating nonceCache for L2 %s. Error: %v", L2URL, err)
			}
			auth, err := nodeClient.GetSignerFromKeystore(ctx, c.UpstreamCfg.ClaimTxManager.PrivateKey)
			if err != nil {
				return err
			}

			claimTxManager, err := claimtxman.NewClaimTxManager(
				ctx, c.UpstreamCfg.ClaimTxManager, chsExitRootEvent[i], chsSyncedL2[i],
				L2URL, networkIDs[i+1], c.UpstreamCfg.NetworkConfig.L2PolygonBridgeAddresses[i], bridgeService,
				storage, rollupID, l2Ethermans[i], nonceCache, auth,
			)
			claimTxManager.SetRedis(redisStorage)
			claimTxManager.SetProducer(messagePushProducer)
			claimTxManager.SetFreegas(c.ClaimTxManager.FreeGas)
			claimTxManager.SetOptClaim(c.ClaimTxManager.OptClaim)
			if err != nil {
				log.Fatalf("error creating claim tx manager for L2 %s. Error: %v", L2URL, err)
			}
			go claimTxManager.StartXLayer()
		}
	}

	sync, err := synchronizer.NewSynchronizer(
		ctx, storage, bridgeController, l1Etherman,
		nil, c.UpstreamCfg.NetworkConfig.GenBlockNumber, nil,
		chsExitRootEvent, chSynced, c.UpstreamCfg.Synchronizer, networkIDs, false,
	)
	if err != nil {
		return err
	}
	cliSyncL1, ok := sync.(*synchronizer.ClientSynchronizer)
	if ok {
		cliSyncL1 = cliSyncL1.
			SetProducer(messagePushProducer).
			SetRedis(redisStorage).
			SetRollupID(uint(networkID)).
			SetLargeTxUsdLimit(c.Synchronizer.LargeTxUsdLimit)
	}
	errs.Go(cliSyncL1.Sync)

	// Calls L1 to record some metrics to prometheus
	go cliSyncL1.RecordLatestBlockNum()

	go func() {
		for {
			select {
			case netID := <-chSynced:
				log.Debug("L1 Network synced. NetowrkID: ", netID)
				for _, ch := range chsSyncedL2 {
					ch <- netID
				}
			case <-ctx.Done():
				log.Debug("Stopping goroutine that listen new GER updates")
				return
			}
		}
	}()

	if !c.UpstreamCfg.ClaimTxManager.Enabled {
		log.Warn("ClaimTxManager not configured")
		for i := range chsExitRootEvent {
			monitorChannel(ctx, chsExitRootEvent[i], chsSyncedL2[i], networkIDs[i+1], storage)
		}
	}

	if len(c.CoinKafkaConsumer.Brokers) > 0 {
		coinKafkaConsumer, err := coinmiddleware.NewKafkaConsumer(c.CoinKafkaConsumer, redisStorage)
		if err != nil {
			return err
		}
		defer func() {
			if err := coinKafkaConsumer.Close(); err != nil {
				log.Errorf("close kafka consumer error: %v", err)
			}
		}()
		log.Debugf("finish initializing kafka consumer")
		go coinKafkaConsumer.Start(ctx)
	}

	return errs.Wait()
}
