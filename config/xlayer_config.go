package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	businessconfig "github.com/0xPolygonHermez/zkevm-bridge-service/config/business_xlayer"

	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/coinmiddleware"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/estimatetime"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/iprestriction"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/messagepush"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/nacos"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/redisstorage"
	"github.com/0xPolygonHermez/zkevm-bridge-service/xlayer/tokenlogoinfo"
)

type ethermanCfg struct {
	L1ChainId  uint   `mapstructure:"L1ChainId"`
	L2ChainIds []uint `mapstructure:"L2ChainIds"`
}

type serverCfg struct {
	// SentinelConfigFilePath is the file path to store the sentinel config
	SentinelConfigFilePath string              `mapstructure:"SentinelConfigFilePath"`
	Redis                  redisstorage.Config `mapstructure:"Redis"`
}

type metricsCfg struct {
	// Env is the environment label for the metrics, to separate mainnet and testnet metrics
	Env string `mapstructure:"Env"`
}

type claimTxManCfg struct {
	// FreeGas enabled whether gas price is 0
	FreeGas bool `mapstructure:"FreeGas"`
	// OptClaim enabled store claimTx into storage every deposit
	OptClaim bool `mapstructure:"OptClaim"`
	// Number of DB claim records to query
	MonitorTxsLimit uint `mapstructure:"MonitorTxsLimit"`
}

type syncCfg struct {
	LargeTxUsdLimit uint64 `mapstructure:"LargeTxUsdLimit"`
}

// We will use this config to wrap the upstream config and load
// our XLayer specfic values in this struct.
type XLayerConfig struct {
	// All upstream structs are preserved
	UpstreamCfg *Config

	// Configs that overlapped with upstream configs.
	// Add new fields here instead of adding them to the upstream structs.
	Etherman       ethermanCfg
	BridgeServer   serverCfg
	Metrics        metricsCfg
	ClaimTxManager claimTxManCfg
	Synchronizer   syncCfg

	// Pure XLayer configs
	L1TargetBlockConfirmations uint64
	NacosConfig                nacos.Config
	BusinessConfig             businessconfig.Config
	IPRestriction              iprestriction.Config
	TokenLogoServiceConfig     tokenlogoinfo.Config
	CoinKafkaConsumer          coinmiddleware.Config
	MessagePushProducer        messagepush.Config
	EstimateTime               estimatetime.Config
}

// Load the same config file, but this time load all X Layer config here.
// The config toml file structure must remains the same.
func LoadXLayerCfg(upstreamCfg *Config) (*XLayerConfig, error) {
	var cfg XLayerConfig
	viper.SetConfigType("toml")
	err := viper.Unmarshal(&cfg, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc()))
	if err != nil {
		return nil, err
	}
	cfg.UpstreamCfg = upstreamCfg
	return &cfg, nil
}
