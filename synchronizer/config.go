package synchronizer

import (
	"github.com/0xPolygonHermez/zkevm-bridge-service/config/types"
)

// Config represents the configuration of the synchronizer
type Config struct {
	// SyncInterval is the delay interval between reading new rollup information
	SyncInterval types.Duration `mapstructure:"SyncInterval"`

	// SyncChunkSize is the number of blocks to sync on each chunk
	SyncChunkSize uint64 `mapstructure:"SyncChunkSize"`

	// ForceL2SyncChunk is a flag to force the L2 synchronizer to sync a chunk. This will disable part of the reorg protection
	ForceL2SyncChunk bool `mapstructure:"ForceL2SyncChunk"`
}
