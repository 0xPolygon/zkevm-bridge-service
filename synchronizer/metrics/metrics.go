package metrics

import (
	"fmt"
	"math/big"
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
)

var Prefix string

// Register the metrics for the synchronizer package.
func Register(networkID uint32) {
	// Prefix for the metrics of the synchronizer package.
	Prefix = "synchronizer_networkID_" + fmt.Sprintf("%d", networkID) + "_"

	gauges := []prometheus.GaugeOpts{
		{
			Name: Prefix + CurrentPendingBridgesToClaimName,
			Help: "[SYNCHRONIZER] current pending deposits to claim",
		},
		{
			Name: Prefix + LatestBlockSyncedName,
			Help: "[SYNCHRONIZER] latest block synced",
		},
	}
	counters := []prometheus.CounterOpts{
		{
			Name: Prefix + DepositCounterName,
			Help: "[SYNCHRONIZER] count processed deposit events",
		},
		{
			Name: Prefix + L1GERCounterName,
			Help: "[SYNCHRONIZER] count processed L1 GER events",
		},
		{
			Name: Prefix + L2GERCounterName,
			Help: "[SYNCHRONIZER] count processed L2 GER events",
		},
		{
			Name: Prefix + RemoveL2GERCounterName,
			Help: "[SYNCHRONIZER] count processed remove L2 GER events",
		},
		{
			Name: Prefix + VerifyBatchCounterName,
			Help: "[SYNCHRONIZER] count processed verify batch events",
		},
		{
			Name: Prefix + ClaimCounterName,
			Help: "[SYNCHRONIZER] count processed claim events",
		},
		{
			Name: Prefix + ReorgCounterName,
			Help: "[SYNCHRONIZER] count the number of reorgs",
		},
		{
			Name: Prefix + ReorgedBlockCounterName,
			Help: "[SYNCHRONIZER] count the reorged blocks",
		},
		{
			Name: Prefix + WrappedTokensCounterName,
			Help: "[SYNCHRONIZER] count the wrapped tokens",
		},
	}
	histograms := []prometheus.HistogramOpts{
		{
			Name: Prefix + InitializationTimeName,
			Help: "[SYNCHRONIZER] initialization time",
		},
		{
			Name: Prefix + FullTrustedSyncTimeName,
			Help: "[SYNCHRONIZER] full trusted state synchronization time",
		},
		{
			Name: Prefix + FullL1SyncTimeName,
			Help: "[SYNCHRONIZER] full L1 synchronization time",
		},
		{
			Name: Prefix + FullSyncIterationTimeName,
			Help: "[SYNCHRONIZER] full synchronization iteration time",
		},
		{
			Name: Prefix + ReadL1DataTimeName,
			Help: "[SYNCHRONIZER] read L1 data time",
		},
		{
			Name: Prefix + ProcessL1DataTimeName,
			Help: "[SYNCHRONIZER] process L1 data time",
		},
		{
			Name: Prefix + GetTrustedGerTimeName,
			Help: "[SYNCHRONIZER] get trusted GER time",
		},
		{
			Name: Prefix + GetTrustedExitRootsTimeName,
			Help: "[SYNCHRONIZER] get trusted ExitRoots time",
		},
		{
			Name: Prefix + DepositAmountName,
			Help: "[SYNCHRONIZER] deposit amount",
		},
		{
			Name: Prefix + ClaimAmountName,
			Help: "[SYNCHRONIZER] claim amount",
		},
	}

	metrics.RegisterHistograms(histograms...)
	metrics.RegisterCounters(counters...)
	metrics.RegisterGauges(gauges...)
}

// LatestBlockSynced sets the latest block synced on the gauge.
func LatestBlockSynced(blockNumber uint64) {
	// Be careful, this uint64 to float64 converion can overflow
	metrics.GaugeSet(Prefix+LatestBlockSyncedName, float64(blockNumber))
}

// IncrementsPendingBridgesToClaim increments the current pending bridges to claim on the gauge.
func IncrementsPendingBridgesToClaim(destNetwork uint32) {
	// This method modify the metric on the destination network.
	// It won't do nothing if the dest network is not synced and the metrics in the dest network are not registered.
	prefixKey := "synchronizer_networkID_" + fmt.Sprintf("%d", destNetwork) + "_"
	metrics.GaugeInc(prefixKey + CurrentPendingBridgesToClaimName)
}

// DecrementsPendingBridgesToClaim decrements the current pending bridges to claim on the gauge.
func DecrementsPendingBridgesToClaim() {
	// Decrements the current pending bridges to claim on the network that received the claim.
	metrics.GaugeDec(Prefix + CurrentPendingBridgesToClaimName)
}

// DepositAmount observes the value of the deposit amount on the histogram.
func DepositAmount(amount *big.Int) {
	a, _ := amount.Float64()
	metrics.HistogramObserve(Prefix+DepositAmountName, a)
}

// ClaimAmount observes the value of the Claim amount on the histogram.
func ClaimAmount(amount *big.Int) {
	a, _ := amount.Float64()
	metrics.HistogramObserve(Prefix+ClaimAmountName, a)
}

// InitializationTime observes the time initializing the synchronizer on the histogram.
func InitializationTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+InitializationTimeName, execTimeInSeconds)
}

// FullTrustedSyncTime observes the time for synchronize the trusted state on the histogram.
func FullTrustedSyncTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+FullTrustedSyncTimeName, execTimeInSeconds)
}

// FullL1SyncTime observes the time for synchronize the trusted state on the histogram.
func FullL1SyncTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+FullL1SyncTimeName, execTimeInSeconds)
}

// FullSyncIterationTime observes the time for synchronize the trusted state on the histogram.
func FullSyncIterationTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+FullSyncIterationTimeName, execTimeInSeconds)
}

// ReadL1DataTime observes the time for synchronize the trusted state on the histogram.
func ReadL1DataTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+ReadL1DataTimeName, execTimeInSeconds)
}

// ProcessL1DataTime observes the time for synchronize the trusted state on the histogram.
func ProcessL1DataTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+ProcessL1DataTimeName, execTimeInSeconds)
}

// GetTrustedGerTime observes the time for synchronize the trusted state on the histogram.
func GetTrustedGerTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+GetTrustedGerTimeName, execTimeInSeconds)
}

// GetTrustedExitRootsTime observes the time for synchronize the trusted state on the histogram.
func GetTrustedExitRootsTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+GetTrustedExitRootsTimeName, execTimeInSeconds)
}

// DepositCounter increases the counter for the processed deposits
func DepositCounter() {
	metrics.CounterInc(Prefix + DepositCounterName)
}

// ClaimCounter increases the counter for the processed claims
func ClaimCounter() {
	metrics.CounterInc(Prefix + ClaimCounterName)
}

// VerifyBatchCounter increases the counter for the processed verifyBatches
func VerifyBatchCounter() {
	metrics.CounterInc(Prefix + VerifyBatchCounterName)
}

// RemoveL2GERCounter increases the counter for the processed removeL2GERs
func RemoveL2GERCounter() {
	metrics.CounterInc(Prefix + RemoveL2GERCounterName)
}

// L2GERCounter increases the counter for the processed L2GER
func L2GERCounter() {
	metrics.CounterInc(Prefix + L2GERCounterName)
}

// L1GERCounter increases the counter for the processed L1GER
func L1GERCounter() {
	metrics.CounterInc(Prefix + L1GERCounterName)
}

// ReorgCounter increases the counter each reorg
func ReorgCounter() {
	metrics.CounterInc(Prefix + ReorgCounterName)
}

// ReorgedBlockCounter increases the counter for the processed Reorged Block
func ReorgedBlocksCounter() {
	metrics.CounterInc(Prefix + ReorgedBlockCounterName)
}

// WrappedTokensCounter increases the counter for the processed wrappedTokens
func WrappedTokensCounter() {
	metrics.CounterInc(Prefix + WrappedTokensCounterName)
}
