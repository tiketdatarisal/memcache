package memcache

import "github.com/patrickmn/go-cache"

// Initialize create a new instance of cache with new configuration.
func Initialize(config ...Config) {
	mutex.Lock()
	defer mutex.Unlock()

	cfg := configDefault(config...)
	c = cache.New(cfg.DefaultExpiration, cfg.CleanupInterval)
}

// Singleton return a single instance of cache.
func Singleton() *cache.Cache {
	if c == nil {
		Initialize()
	}

	return c
}
