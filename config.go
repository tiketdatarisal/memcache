package memcache

import "time"

type Config struct {
	// DefaultExpiration defines default expiration interval for local cache.
	//
	// Optional. Default: 5 minutes
	DefaultExpiration time.Duration

	// CleanupInterval defines cleanup interval for local cache to remove data permanently.
	//
	// Optional. Default: 10 minutes
	CleanupInterval time.Duration
}

var (
	ConfigDefault = Config{
		DefaultExpiration: defaultExpiration,
		CleanupInterval:   defaultCleanupInterval,
	}
)

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]
	if cfg.DefaultExpiration < 0 {
		cfg.DefaultExpiration = ConfigDefault.DefaultExpiration
	}

	if cfg.CleanupInterval < 0 {
		cfg.CleanupInterval = ConfigDefault.CleanupInterval
	}

	return cfg
}
