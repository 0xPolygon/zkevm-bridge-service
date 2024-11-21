package etherman

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Block struct
type Block struct {
	ID              uint64
	BlockNumber     uint64
	BlockHash       common.Hash
	ParentHash      common.Hash
	NetworkID       uint32
	GlobalExitRoots []GlobalExitRoot
	Deposits        []Deposit
	Claims          []Claim
	Tokens          []TokenWrapped
	VerifiedBatches []VerifiedBatch
	ActivateEtrog   []bool
	ReceivedAt      time.Time
}

// GlobalExitRoot struct
type GlobalExitRoot struct {
	BlockID        uint64
	BlockNumber    uint64
	ExitRoots      []common.Hash
	GlobalExitRoot common.Hash
	NetworkID      uint32
}

// Deposit struct
type Deposit struct {
	Id                 uint64
	LeafType           uint8
	OriginalNetwork    uint32
	OriginalAddress    common.Address
	Amount             *big.Int
	DestinationNetwork uint32
	DestinationAddress common.Address
	DepositCount       uint32
	BlockID            uint64
	BlockNumber        uint64
	NetworkID          uint32
	TxHash             common.Hash
	Metadata           []byte
	// it is only used for the bridge service
	ReadyForClaim bool
}

// Claim struct
type Claim struct {
	MainnetFlag        bool
	RollupIndex        uint32
	Index              uint32
	OriginalNetwork    uint32
	OriginalAddress    common.Address
	Amount             *big.Int
	DestinationAddress common.Address
	BlockID            uint64
	BlockNumber        uint64
	NetworkID          uint32
	TxHash             common.Hash
}

// TokenWrapped struct
type TokenWrapped struct {
	TokenMetadata
	OriginalNetwork      uint32
	OriginalTokenAddress common.Address
	WrappedTokenAddress  common.Address
	BlockID              uint64
	BlockNumber          uint64
	NetworkID            uint32
}

// TokenMetadata is a metadata of ERC20 token.
type TokenMetadata struct {
	Name     string
	Symbol   string
	Decimals uint8
}

type VerifiedBatch struct {
	BlockNumber   uint64
	BatchNumber   uint64
	RollupID      uint32
	LocalExitRoot common.Hash
	TxHash        common.Hash
	StateRoot     common.Hash
	Aggregator    common.Address
}

// RollupExitLeaf struct
type RollupExitLeaf struct {
	ID       uint64
	BlockID  uint64
	Leaf     common.Hash
	RollupId uint32
	Root     common.Hash
}
