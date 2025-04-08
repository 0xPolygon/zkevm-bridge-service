package estimatetime

type Config struct {
	// Time limits set for each L1, L2.
	// The average time calculated cannot exceed this limit.
	DefaultTime []uint32 `mapstructure:"DefaultTime"`

	// Number of DB records to calculate average time from.
	SampleLimit uint32 `mapstructure:"SampleLimit"`
}
