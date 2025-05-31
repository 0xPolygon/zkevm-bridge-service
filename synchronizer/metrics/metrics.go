package metrics

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/0xPolygonHermez/zkevm-bridge-service/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// InitializationTimeName is the name of the label for the initialization of the synchronizer.
	InitializationTimeName = "initialization_time"

	// FullTrustedSyncTimeName is the name of the label for the synchronization of the trusted state.
	FullTrustedSyncTimeName = "full_trusted_sync_time"

	// FullL1SyncTimeName is the name of the label for the synchronization of the L1 state.
	FullL1SyncTimeName = "full_L1_sync_time"

	// FullSyncIterationTimeName is the name of the label for a L1 synchronization.
	FullSyncIterationTimeName = "full_L1_sync_time"

	// ReadL1DataTimeName is the name of the label to read L1 data.
	ReadL1DataTimeName = "read_L1_data_time"

	// ProcessL1DataTimeName is the name of the label to process L1 data.
	ProcessL1DataTimeName = "process_L1_data_time"

	// GetTrustedGerTimeName is the name of the label to get trusted Ger.
	GetTrustedGerTimeName = "get_trusted_ger_time"

	// GetTrustedExitRootsTimeName is the name of the label to get trusted ExitRoots.
	GetTrustedExitRootsTimeName = "get_trusted_exitroots_time"

	// DepositAmountName is the name of the label that gets the deposit amount.
	DepositAmountName = "deposit_amount"

	// ClaimAmountName is the name of the label that gets the claim amount.
	ClaimAmountName = "claim_amount"

	// DepositCounterName is the name of the label to count the processed deposits.
	DepositCounterName = "processed_deposit_counter"

	// L1GERCounterName is the name of the label to count the processed L1GERs.
	L1GERCounterName = "processed_L1GER_counter"

	// L2GERCounterName is the name of the label to count the processed L2GERs.
	L2GERCounterName = "processed_L2GER_counter"

	// RemoveL2GERCounterName is the name of the label to count the processed remove L2GERs.
	RemoveL2GERCounterName = "processed_remove_L2GER_counter"

	// VerifyBatchCounterName is the name of the label to count the processed verifyBatches.
	VerifyBatchCounterName = "processed_verifyBatch_counter"

	// ClaimCounterName is the name of the label to count the processed claims.
	ClaimCounterName = "processed_claim_counter"

	// ReorgCounterName is the name of the label to count the reorgs.
	ReorgCounterName = "reorg_counter"

	// ReorgCounterName is the name of the label to count the block reorgs.
	ReorgedBlockCounterName = "reorged_block_counter"

	// WrappedTokensCounterName is the name of the label to count the wrapped tokens.
	WrappedTokensCounterName = "processed_wrapped_tokens_counter"

	// CurrentPendingBridgesToClaimName is the name of the label to count the current pending bridges to claim.
	CurrentPendingBridgesToClaimName = "current_pending_bridges_to_claim"

	// LatestBlockSyncedName is the name of the label to get the latest block synced.
	LatestBlockSyncedName = "latest_block_synced"

	// LatestBlockRollupInfoName is the name of the label to get the latest block rollup info.
	LatestBlockRollupInfoName = "latest_block_rollup_info"
)

var (
	registerMutex sync.Mutex
)

type MetricsInterface interface {
	LatestBlockSynced(blockNumber uint64)
	LatestBlockRollupInfo(blockNumber uint64)
	IncrementsPendingBridgesToClaim()
	DecrementsPendingBridgesToClaim()
	DepositAmount(amount *big.Int)
	ClaimAmount(amount *big.Int)
	InitializationTime(lastProcessTime time.Duration)
	FullTrustedSyncTime(lastProcessTime time.Duration)
	FullL1SyncTime(lastProcessTime time.Duration)
	FullSyncIterationTime(lastProcessTime time.Duration)
	ReadL1DataTime(lastProcessTime time.Duration)
	ProcessL1DataTime(lastProcessTime time.Duration)
	GetTrustedGerTime(lastProcessTime time.Duration)
	GetTrustedExitRootsTime(lastProcessTime time.Duration)
	DepositCounter()
	ClaimCounter()
	VerifyBatchCounter()
	RemoveL2GERCounter()
	L2GERCounter()
	L1GERCounter()
	ReorgCounter()
	ReorgedBlocksCounter()
	WrappedTokensCounter()
}

type Metrics struct {
	prefix string
}

// Register the metrics for the synchronizer package.
func Register(networkID uint32) MetricsInterface {
	registerMutex.Lock()
	defer registerMutex.Unlock()
	// Prefix for the metrics of the synchronizer package.
	prefix := "synchronizer_networkID_" + fmt.Sprintf("%d_", networkID)

	gauges := []prometheus.GaugeOpts{
		{
			Name: prefix + CurrentPendingBridgesToClaimName,
			Help: "[SYNCHRONIZER] current pending deposits to claim",
		},
		{
			Name: prefix + LatestBlockSyncedName,
			Help: "[SYNCHRONIZER] latest block synced",
		},
		{
			Name: prefix + LatestBlockRollupInfoName,
			Help: "[SYNCHRONIZER] latest block rollup info",
		},
	}
	counters := []prometheus.CounterOpts{
		{
			Name: prefix + DepositCounterName,
			Help: "[SYNCHRONIZER] count processed deposit events",
		},
		{
			Name: prefix + L1GERCounterName,
			Help: "[SYNCHRONIZER] count processed L1 GER events",
		},
		{
			Name: prefix + L2GERCounterName,
			Help: "[SYNCHRONIZER] count processed L2 GER events",
		},
		{
			Name: prefix + RemoveL2GERCounterName,
			Help: "[SYNCHRONIZER] count processed remove L2 GER events",
		},
		{
			Name: prefix + VerifyBatchCounterName,
			Help: "[SYNCHRONIZER] count processed verify batch events",
		},
		{
			Name: prefix + ClaimCounterName,
			Help: "[SYNCHRONIZER] count processed claim events",
		},
		{
			Name: prefix + ReorgCounterName,
			Help: "[SYNCHRONIZER] count the number of reorgs",
		},
		{
			Name: prefix + ReorgedBlockCounterName,
			Help: "[SYNCHRONIZER] count the reorged blocks",
		},
		{
			Name: prefix + WrappedTokensCounterName,
			Help: "[SYNCHRONIZER] count the wrapped tokens",
		},
	}
	histograms := []prometheus.HistogramOpts{
		{
			Name: prefix + InitializationTimeName,
			Help: "[SYNCHRONIZER] initialization time",
		},
		{
			Name: prefix + FullTrustedSyncTimeName,
			Help: "[SYNCHRONIZER] full trusted state synchronization time",
		},
		{
			Name: prefix + FullL1SyncTimeName,
			Help: "[SYNCHRONIZER] full L1 synchronization time",
		},
		{
			Name: prefix + FullSyncIterationTimeName,
			Help: "[SYNCHRONIZER] full synchronization iteration time",
		},
		{
			Name: prefix + ReadL1DataTimeName,
			Help: "[SYNCHRONIZER] read L1 data time",
		},
		{
			Name: prefix + ProcessL1DataTimeName,
			Help: "[SYNCHRONIZER] process L1 data time",
		},
		{
			Name: prefix + GetTrustedGerTimeName,
			Help: "[SYNCHRONIZER] get trusted GER time",
		},
		{
			Name: prefix + GetTrustedExitRootsTimeName,
			Help: "[SYNCHRONIZER] get trusted ExitRoots time",
		},
		{
			Name: prefix + DepositAmountName,
			Help: "[SYNCHRONIZER] deposit amount",
		},
		{
			Name: prefix + ClaimAmountName,
			Help: "[SYNCHRONIZER] claim amount",
		},
	}

	metrics.RegisterHistograms(histograms...)
	metrics.RegisterCounters(counters...)
	metrics.RegisterGauges(gauges...)

	return &Metrics{
		prefix: prefix,
	}
}

// LatestBlockSynced sets the latest block synced on the gauge.
func (m *Metrics) LatestBlockSynced(blockNumber uint64) {
	// Be careful, this uint64 to float64 converion can overflow
	metrics.GaugeSet(m.prefix+LatestBlockSyncedName, float64(blockNumber))
}

// LatestBlockRollupInfo sets the latest block rollup info on the gauge.
func (m *Metrics) LatestBlockRollupInfo(blockNumber uint64) {
	// Be careful, this uint64 to float64 converion can overflow
	metrics.GaugeSet(m.prefix+LatestBlockRollupInfoName, float64(blockNumber))
}

// IncrementsPendingBridgesToClaim increments the current pending bridges to claim on the gauge.
func (m *Metrics) IncrementsPendingBridgesToClaim() {
	// This method modify the metric on the destination network.
	// It won't do nothing if the dest network is not synced and the metrics in the dest network are not registered.
	metrics.GaugeInc(m.prefix + CurrentPendingBridgesToClaimName)
}

// DecrementsPendingBridgesToClaim decrements the current pending bridges to claim on the gauge.
func (m *Metrics) DecrementsPendingBridgesToClaim() {
	// Decrements the current pending bridges to claim on the network that received the claim.
	metrics.GaugeDec(m.prefix + CurrentPendingBridgesToClaimName)
}

// DepositAmount observes the value of the deposit amount on the histogram.
func (m *Metrics) DepositAmount(amount *big.Int) {
	a, _ := amount.Float64()
	metrics.HistogramObserve(m.prefix+DepositAmountName, a)
}

// ClaimAmount observes the value of the Claim amount on the histogram.
func (m *Metrics) ClaimAmount(amount *big.Int) {
	a, _ := amount.Float64()
	metrics.HistogramObserve(m.prefix+ClaimAmountName, a)
}

// InitializationTime observes the time initializing the synchronizer on the histogram.
func (m *Metrics) InitializationTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(m.prefix+InitializationTimeName, execTimeInSeconds)
}

// FullTrustedSyncTime observes the time for synchronize the trusted state on the histogram.
func (m *Metrics) FullTrustedSyncTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(m.prefix+FullTrustedSyncTimeName, execTimeInSeconds)
}

// FullL1SyncTime observes the time for synchronize the trusted state on the histogram.
func (m *Metrics) FullL1SyncTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(m.prefix+FullL1SyncTimeName, execTimeInSeconds)
}

// FullSyncIterationTime observes the time for synchronize the trusted state on the histogram.
func (m *Metrics) FullSyncIterationTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(m.prefix+FullSyncIterationTimeName, execTimeInSeconds)
}

// ReadL1DataTime observes the time for synchronize the trusted state on the histogram.
func (m *Metrics) ReadL1DataTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(m.prefix+ReadL1DataTimeName, execTimeInSeconds)
}

// ProcessL1DataTime observes the time for synchronize the trusted state on the histogram.
func (m *Metrics) ProcessL1DataTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(m.prefix+ProcessL1DataTimeName, execTimeInSeconds)
}

// GetTrustedGerTime observes the time for synchronize the trusted state on the histogram.
func (m *Metrics) GetTrustedGerTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(m.prefix+GetTrustedGerTimeName, execTimeInSeconds)
}

// GetTrustedExitRootsTime observes the time for synchronize the trusted state on the histogram.
func (m *Metrics) GetTrustedExitRootsTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(m.prefix+GetTrustedExitRootsTimeName, execTimeInSeconds)
}

// DepositCounter increases the counter for the processed deposits
func (m *Metrics) DepositCounter() {
	metrics.CounterInc(m.prefix + DepositCounterName)
}

// ClaimCounter increases the counter for the processed claims
func (m *Metrics) ClaimCounter() {
	metrics.CounterInc(m.prefix + ClaimCounterName)
}

// VerifyBatchCounter increases the counter for the processed verifyBatches
func (m *Metrics) VerifyBatchCounter() {
	metrics.CounterInc(m.prefix + VerifyBatchCounterName)
}

// RemoveL2GERCounter increases the counter for the processed removeL2GERs
func (m *Metrics) RemoveL2GERCounter() {
	metrics.CounterInc(m.prefix + RemoveL2GERCounterName)
}

// L2GERCounter increases the counter for the processed L2GER
func (m *Metrics) L2GERCounter() {
	metrics.CounterInc(m.prefix + L2GERCounterName)
}

// L1GERCounter increases the counter for the processed L1GER
func (m *Metrics) L1GERCounter() {
	metrics.CounterInc(m.prefix + L1GERCounterName)
}

// ReorgCounter increases the counter each reorg
func (m *Metrics) ReorgCounter() {
	metrics.CounterInc(m.prefix + ReorgCounterName)
}

// ReorgedBlockCounter increases the counter for the processed Reorged Block
func (m *Metrics) ReorgedBlocksCounter() {
	metrics.CounterInc(m.prefix + ReorgedBlockCounterName)
}

// WrappedTokensCounter increases the counter for the processed wrappedTokens
func (m *Metrics) WrappedTokensCounter() {
	metrics.CounterInc(m.prefix + WrappedTokensCounterName)
}
