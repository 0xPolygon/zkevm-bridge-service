package metrics

import (
	"fmt"
	"time"

	"github.com/0xPolygonHermez/zkevm-bridge-service/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// ReadAndProcessAllEventsTimeName is the name of the label read and process all event.
	ReadAndProcessAllEventsTimeName = "read_and_process_all_event_time"

	// ProcessAllEventTimeName is the name of the label to process all event.
	ProcessAllEventTimeName = "process_all_event_time"

	// ProcessSingleEventTimeName is the name of the label to process a single event.
	ProcessSingleEventTimeName = "process_single_event_time"

	// GetEventsTimeName is the name of the label to get L1 events.
	GetEventsTimeName = "get_events_time"

	// EventCounterName is the name of the label to count the processed events.
	EventCounterName = "processed_events_counter"
)

var Prefix string

// Register the metrics for the etherman package.
func Register(networkID uint32) {
	// Prefix for the metrics of the etherman package.
	Prefix = "etherman_networkID_" + fmt.Sprintf("%d", networkID) + "_"
	var (
		counters   []prometheus.CounterOpts
		histograms []prometheus.HistogramOpts
	)

	counters = []prometheus.CounterOpts{
		{
			Name: Prefix + EventCounterName,
			Help: "[ETHERMAN] count processed events",
		},
	}

	histograms = []prometheus.HistogramOpts{
		{
			Name: Prefix + ReadAndProcessAllEventsTimeName,
			Help: "[ETHERMAN] read and process all event time",
		},
		{
			Name: Prefix + ProcessAllEventTimeName,
			Help: "[ETHERMAN] process all event time",
		},
		{
			Name: Prefix + ProcessSingleEventTimeName,
			Help: "[ETHERMAN] process single event time",
		},
		{
			Name: Prefix + GetEventsTimeName,
			Help: "[ETHERMAN] get L1 events time",
		},
	}

	metrics.RegisterCounters(counters...)
	metrics.RegisterHistograms(histograms...)
}

// ReadAndProcessAllEventsTime observes the time read and process all event on the histogram.
func ReadAndProcessAllEventsTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+ReadAndProcessAllEventsTimeName, execTimeInSeconds)
}

// ProcessAllEventTime observes the time to process all event on the histogram.
func ProcessAllEventTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+ProcessAllEventTimeName, execTimeInSeconds)
}

// ProcessSingleEventTime observes the time to process a single event on the histogram.
func ProcessSingleEventTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+ProcessSingleEventTimeName, execTimeInSeconds)
}

// GetEventsTime observes the time to get the events from L1 on the histogram.
func GetEventsTime(lastProcessTime time.Duration) {
	execTimeInSeconds := float64(lastProcessTime) / float64(time.Second)
	metrics.HistogramObserve(Prefix+GetEventsTimeName, execTimeInSeconds)
}

// EventCounter increases the counter for the processed events
func EventCounter() {
	metrics.CounterInc(Prefix + EventCounterName)
}
