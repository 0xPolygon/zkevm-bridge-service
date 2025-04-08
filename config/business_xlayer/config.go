package businessconfig

import (
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	StandardChainIds []uint64 `mapstructure:"StandardChainIds"`
	InnerChainIds    []uint64 `mapstructure:"InnerChainIds"`

	USDCContractAddresses   []common.Address `mapstructure:"USDCContractAddresses"`
	USDCTokenAddresses      []common.Address `mapstructure:"USDCTokenAddresses"`
	WstETHContractAddresses []common.Address `mapstructure:"WstETHContractAddresses"`
	WstETHTokenAddresses    []common.Address `mapstructure:"WstETHTokenAddresses"`
	EURCContractAddresses   []common.Address `mapstructure:"EURCContractAddresses"`
	EURCTokenAddresses      []common.Address `mapstructure:"EURCTokenAddresses"`
}
