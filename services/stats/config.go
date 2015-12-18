package stats

import (
	"time"

	"github.com/influxdata/config"
)

type Config struct {
	Enabled         bool            `toml:"enabled"`
	StatsInterval   config.Duration `toml:"stats-interval"`
	Database        string          `toml:"database"`
	RetentionPolicy string          `toml:"retention-policy"`
}

func NewConfig() Config {
	return Config{
		Enabled:         true,
		Database:        "_kapacitor",
		RetentionPolicy: "default",
		StatsInterval:   config.Duration(10 * time.Second),
	}
}
