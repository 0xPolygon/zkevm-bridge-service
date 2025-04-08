package config

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// Assert XLayer configs can be parsed into config.
func TestXLayerConfig(t *testing.T) {
	c, err := Load("./configs_xlayer/xlayer-test-config.local.toml", "")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := LoadXLayerCfg(c)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("upstream config:\n%+v\n\n", cfg.UpstreamCfg)

	fmt.Printf("Etherman.L1ChainId = %+v\n", cfg.Etherman.L1ChainId)
	assert.True(t, viper.IsSet("Etherman.L1ChainId"))
	fmt.Printf("Etherman.L2ChainIds = %+v\n", cfg.Etherman.L2ChainIds)
	assert.True(t, viper.IsSet("Etherman.L2ChainIds"))
	fmt.Printf("BridgeServer.SentinelConfigFilePath = %+v\n", cfg.BridgeServer.SentinelConfigFilePath)
	assert.True(t, viper.IsSet("BridgeServer.SentinelConfigFilePath"))
	fmt.Printf("BridgeServer.Redis = %+v\n", cfg.BridgeServer.Redis)
	assert.True(t, viper.IsSet("BridgeServer.Redis"))
	fmt.Printf("Metrics.Env = %+v\n", cfg.Metrics.Env)
	assert.True(t, viper.IsSet("Metrics.Env"), "metrics env not set")
	fmt.Printf("NacosConfig = %+v\n", cfg.NacosConfig)
	assert.True(t, viper.IsSet("NacosConfig"))
	fmt.Printf("BusinessConfig = %+v\n", cfg.BusinessConfig)
	assert.True(t, viper.IsSet("BusinessConfig"), "business config not set")
	fmt.Printf("IPRestriction = %+v\n", cfg.IPRestriction)
	assert.True(t, viper.IsSet("IPRestriction"), "IP restriction not set")
	fmt.Printf("TokenLogoServiceConfig = %+v\n", cfg.TokenLogoServiceConfig)
	assert.True(t, viper.IsSet("TokenLogoServiceConfig"), "token Logo not set")
	fmt.Printf("CoinKafkaConsumer = %+v\n", cfg.CoinKafkaConsumer)
	assert.True(t, viper.IsSet("CoinKafkaConsumer"))
	fmt.Printf("MessagePushProducer = %+v\n", cfg.MessagePushProducer)
	assert.True(t, viper.IsSet("MessagePushProducer"))

	assert.True(t, viper.IsSet("Synchronizer.LargeTxUsdLimit"))
	assert.Equal(t, uint64(958), cfg.Synchronizer.LargeTxUsdLimit)
	assert.True(t, viper.IsSet("L1TargetBlockConfirmations"))
	assert.Equal(t, uint64(99), cfg.L1TargetBlockConfirmations)
	assert.True(t, viper.IsSet("ClaimTxManager.MonitorTxsLimit"))
	assert.Equal(t, uint(1234), cfg.ClaimTxManager.MonitorTxsLimit)
	assert.True(t, viper.IsSet("ClaimTxManager.FreeGas"))
	assert.Equal(t, true, cfg.ClaimTxManager.FreeGas)
	assert.True(t, viper.IsSet("ClaimTxManager.OptClaim"))
	assert.Equal(t, true, cfg.ClaimTxManager.OptClaim)
	assert.True(t, viper.IsSet("MessagePushProducer.BizCode"))
	assert.Equal(t, "hello", cfg.MessagePushProducer.BizCode)
	assert.True(t, viper.IsSet("MessagePushProducer.PushKey"))
	assert.Equal(t, "huh", cfg.MessagePushProducer.PushKey)

	assert.True(t, viper.IsSet("EstimateTime.SampleLimit"))
	assert.Equal(t, uint32(16), cfg.EstimateTime.SampleLimit)
	assert.True(t, viper.IsSet("EstimateTime.DefaultTime"))
	assert.Equal(t, []uint32{78, 12}, cfg.EstimateTime.DefaultTime)

	assert.True(t, viper.IsSet("BridgeServer.Redis.EnablePrice"))
	assert.Equal(t, true, cfg.BridgeServer.Redis.EnablePrice)

}
