package memcache

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

const (
	defaultExpiration      = 5 * time.Minute
	defaultCleanupInterval = 10 * time.Minute
)

var (
	c     *cache.Cache
	mutex sync.Mutex
)
