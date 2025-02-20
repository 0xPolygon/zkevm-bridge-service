package metrics

import (
	"time"

	"github.com/0xPolygonHermez/zkevm-bridge-service/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Prefix for the metrics of the server/api package.
	Prefix = "api_"

	// CheckAPILatencyName is the name of the label that measures the latency of the CheckAPI endpoint.
	CheckAPILatencyName = Prefix + "check_api_latency"
	// GetBridgesLatencyName is the name of the label that measures the latency of the GetBridges endpoint.
	GetBridgesLatencyName = Prefix + "get_bridges_latency"
	// GetClaimsLatencyName is the name of the label that measures the latency of the GetClaims endpoint.
	GetClaimsLatencyName = Prefix + "get_claims_latency"
	// GetProofLatencyName is the name of the label that measures the latency of the GetProof endpoint.
	GetProofLatencyName = Prefix + "get_proof_latency"
	// GetBridgeLatencyName is the name of the label that measures the latency of the GetBridge endpoint.
	GetBridgeLatencyName = Prefix + "get_bridge_latency"
	// GetTokenWrappedLatencyName is the name of the label that measures the latency of the GetTokenWrapped endpoint.
	GetTokenWrappedLatencyName = Prefix + "get_wrapped_token_latency"
	// GetProofByGERLatencyName is the name of the label that measures the latency of the GetProofByGER endpoint.
	GetProofByGERLatencyName = Prefix + "get_proof_by_ger_latency"
	// GetPendingBridgesToClaimLatencyName is the name of the label that measures the latency of the GetPendingBridgesToClaim endpoint.
	GetPendingBridgesToClaimLatencyName = Prefix + "get_pending_bridges_to_claim_latency"

	// CheckAPICounterName is the name of the label that counters the number of requests to CheckAPI endpoint.
	CheckAPICounterName = Prefix + "check_api_counter"
	// GetBridgesCounterName is the name of the label that counters the number of requests to GetBridges endpoint.
	GetBridgesCounterName = Prefix + "get_bridges_counter"
	// GetClaimsCounterName is the name of the label that counters the number of requests to GetClaims endpoint.
	GetClaimsCounterName = Prefix + "get_claims_counter"
	// GetProofCounterName is the name of the label that counters the number of requests to GetProof endpoint.
	GetProofCounterName = Prefix + "get_proof_counter"
	// GetBridgeCounterName is the name of the label that counters the number of requests to GetBridge endpoint.
	GetBridgeCounterName = Prefix + "get_bridge_counter"
	// GetTokenWrappedCounterName is the name of the label that counters the number of requests to GetTokenWrapped endpoint.
	GetTokenWrappedCounterName = Prefix + "get_wrapped_token_counter"
	// GetProofByGERCounterName is the name of the label that counters the number of requests to GetProofByGER endpoint.
	GetProofByGERCounterName = Prefix + "get_proof_by_ger_counter"
	// GetPendingBridgesToClaimCounterName is the name of the label that counters the number of requests to GetPendingBridgesToClaim endpoint.
	GetPendingBridgesToClaimCounterName = Prefix + "get_pending_bridges_to_claim_counter"
)

// Register the metrics for the etherman package.
func Register(networkID uint32) {
	var (
		counters   []prometheus.CounterOpts
		histograms []prometheus.HistogramOpts
	)

	counters = []prometheus.CounterOpts{
		{
			Name: CheckAPICounterName,
			Help: "[API] count the number of requests to CheckAPI endpoint",
		},
		{
			Name: GetBridgesCounterName,
			Help: "[API] count the number of requests to GetBridges endpoint",
		},
		{
			Name: GetClaimsCounterName,
			Help: "[API] count the number of requests to GetClaims endpoint",
		},
		{
			Name: GetProofCounterName,
			Help: "[API] count the number of requests to GetProof endpoint",
		},
		{
			Name: GetBridgeCounterName,
			Help: "[API] count the number of requests to GetBridge endpoint",
		},
		{
			Name: GetTokenWrappedCounterName,
			Help: "[API] count the number of requests to GetTokenWrapped endpoint",
		},
		{
			Name: GetProofByGERCounterName,
			Help: "[API] count the number of requests to GetProofByGER endpoint",
		},
		{
			Name: GetPendingBridgesToClaimCounterName,
			Help: "[API] count the number of requests to GetPendingBridgesToClaim endpoint",
		},
	}

	histograms = []prometheus.HistogramOpts{
		{
			Name: CheckAPILatencyName,
			Help: "[API] measures the latency of the CheckAPI endpoint",
		},
		{
			Name: GetBridgesLatencyName,
			Help: "[API] measures the latency of the GetBridges endpoint",
		},
		{
			Name: GetClaimsLatencyName,
			Help: "[API] measures the latency of the GetClaims endpoint",
		},
		{
			Name: GetProofLatencyName,
			Help: "[API] measures the latency of the GetProof endpoint",
		},
		{
			Name: GetBridgeLatencyName,
			Help: "[API] measures the latency of the GetBridge endpoint",
		},
		{
			Name: GetTokenWrappedLatencyName,
			Help: "[API] measures the latency of the GetTokenWrapped endpoint",
		},
		{
			Name: GetProofByGERLatencyName,
			Help: "[API] measures the latency of the GetProofByGER endpoint",
		},
		{
			Name: GetPendingBridgesToClaimLatencyName,
			Help: "[API] measures the latency of the GetPendingBridgesToClaim endpoint",
		},
	}
	metrics.RegisterCounters(counters...)
	metrics.RegisterHistograms(histograms...)
}

// CheckAPILatency observes the latency of the endpoint and shows it on the histogram.
func CheckAPILatency(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(CheckAPILatencyName, execTimeInSeconds)
}

// GetBridgesLatency observes the latency of the endpoint and shows it on the histogram.
func GetBridgesLatency(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(GetBridgesLatencyName, execTimeInSeconds)
}

// GetClaimsLatency observes the latency of the endpoint and shows it on the histogram.
func GetClaimsLatency(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(GetClaimsLatencyName, execTimeInSeconds)
}

// GetProofLatency observes the latency of the endpoint and shows it on the histogram.
func GetProofLatency(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(GetProofLatencyName, execTimeInSeconds)
}

// GetBridgeLatency observes the latency of the endpoint and shows it on the histogram.
func GetBridgeLatency(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(GetBridgeLatencyName, execTimeInSeconds)
}

// GetTokenWrappedLatency observes the latency of the endpoint and shows it on the histogram.
func GetTokenWrappedLatency(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(GetTokenWrappedLatencyName, execTimeInSeconds)
}

// GetProofByGERLatency observes the latency of the endpoint and shows it on the histogram.
func GetProofByGERLatency(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(GetProofByGERLatencyName, execTimeInSeconds)
}

// GetPendingBridgesToClaimLatency observes the latency of the endpoint and shows it on the histogram.
func GetPendingBridgesToClaimLatency(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(GetPendingBridgesToClaimLatencyName, execTimeInSeconds)
}

// CheckAPICounter increases the counter for the CheckAPI endpoint
func CheckAPICounter() {
	metrics.CounterInc(CheckAPICounterName)
}

// GetBridgesCounter increases the counter for the GetBridges endpoint
func GetBridgesCounter() {
	metrics.CounterInc(GetBridgesCounterName)
}

// GetClaimsCounter increases the counter for the GetClaims endpoint
func GetClaimsCounter() {
	metrics.CounterInc(GetClaimsCounterName)
}

// GetProofCounter increases the counter for the GetProof endpoint
func GetProofCounter() {
	metrics.CounterInc(GetProofCounterName)
}

// GetBridgeCounter increases the counter for the GetBridge endpoint
func GetBridgeCounter() {
	metrics.CounterInc(GetBridgeCounterName)
}

// GetTokenWrappedCounter increases the counter for the GetTokenWrapped endpoint
func GetTokenWrappedCounter() {
	metrics.CounterInc(GetTokenWrappedCounterName)
}

// GetProofByGERCounter increases the counter for the GetProofByGER endpoint
func GetProofByGERCounter() {
	metrics.CounterInc(GetProofByGERCounterName)
}

// GetPendingBridgesToClaimCounter increases the counter for the GetPendingBridgesToClaim endpoint
func GetPendingBridgesToClaimCounter() {
	metrics.CounterInc(GetPendingBridgesToClaimCounterName)
}
