package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/anabiozz/yotunheim/backend/internal/models"
	"github.com/anabiozz/yotunheim/backend/metrics"

	"github.com/BurntSushi/toml"
)

// Config struct
type Config struct {
	Agent        *AgentConfig
	InputFilters map[string]interface{}

	Inputs []*models.RunningInput
}

type AgentConfig struct {
	// Interval at which to gather information
	Interval time.Duration `toml:"interval"`

	// RoundInterval rounds collection interval to 'interval'.
	//     ie, if Interval=10s then always collect on :00, :10, :20, etc.
	RoundInterval bool `toml:"round_interval"`

	// By default or when set to "0s", precision will be set to the same
	// timestamp order as the collection interval, with the maximum being 1s.
	//   ie, when interval = "10s", precision will be "1s"
	//       when interval = "250ms", precision will be "1ms"
	// Precision will NOT be used for service inputs. It is up to each individual
	// service input to set the timestamp at the appropriate precision.
	Precision int `toml:"precision"`

	// CollectionJitter is used to jitter the collection by a random amount.
	// Each plugin will sleep for a random time within jitter before collecting.
	// This can be used to avoid many plugins querying things like sysfs at the
	// same time, which can have a measurable effect on the system.
	CollectionJitter int `toml:"collection_jitter"`

	// FlushInterval is the Interval at which to flush data
	FlushInterval time.Duration `toml:"flush_interval"`

	// FlushJitter Jitters the flush interval by a random amount.
	// This is primarily to avoid large write spikes for users running a large
	// number of  instances.
	// ie, a jitter of 5s and interval 10s means flushes will happen every 10-15s
	FlushJitter int `toml:"flush_jitter"`

	// MetricBatchSize is the maximum number of metrics that is wrote to an
	// output plugin in one call.
	MetricBatchSize int `toml:"metric_batch_size"`

	// MetricBufferLimit is the max number of metrics that each output plugin
	// will cache. The buffer is cleared when a successful write occurs. When
	// full, the oldest metrics will be overwritten. This number should be a
	// multiple of MetricBatchSize. Due to current implementation, this could
	// not be less than 2 times MetricBatchSize.
	MetricBufferLimit int `toml:"metric_buffer_limit"`

	// FlushBufferWhenFull tells  to flush the metric buffer whenever
	// it fills up, regardless of FlushInterval. Setting this option to true
	// does _not_ deactivate FlushInterval.
	FlushBufferWhenFull bool `toml:"flush_buffer_when_dull"`

	// Debug is the option for running in debug mode
	Debug bool `toml:"debug"`

	// Logfile specifies the file to send logs to
	Logfile string `toml:"logfile"`

	// Quiet is the option for running in quiet mode
	Quiet        bool   `toml:"quiet"`
	Hostname     string `toml:"hostname"`
	OmitHostname bool   `toml:"omit_hostname"`
}

// NewConfig return new config
func NewConfig() *Config {
	c := &Config{
		Agent:        &AgentConfig{},
		InputFilters: make(map[string]interface{}, 0),
		// Inputs:       make([]*models.RunningInput, 0),
		// Outputs:      make([]*models.RunningOutput, 0),
	}
	return c
}

// Check the occurrence of the name in list array
func sliceContains(name string, list []interface{}) bool {
	for _, b := range list {
		if b == name {
			return true
		}
	}
	return false
}

// AddInput ...
func (c *Config) AddInput(name string) error {
	if len(c.InputFilters["inputs"].([]interface{})) > 0 && !sliceContains(name, c.InputFilters["inputs"].([]interface{})) {
		return nil
	}
	// Legacy support renaming io input to diskio
	if name == "io" {
		name = "diskio"
	}

	creator, ok := metrics.Metrics[name]
	if !ok {
		return fmt.Errorf("Undefined but requested input: %s", name)
	}
	input := creator()

	rp := models.NewRunningInput(input)
	c.Inputs = append(c.Inputs, rp)
	return nil
}

const (
	envConfigPath = "PG_API_CONFIG"
)

// LoadConfig ...
func (c *Config) LoadConfig() error {
	// utils.GetEnv(flag.Arg(0), flag.Arg(0))
	_, err := toml.DecodeFile(flag.Arg(0), c)
	if err != nil {
		return err
	}
	return nil
}
