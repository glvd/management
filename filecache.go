package management

import (
	"github.com/gocacher/badger-cache"
	"github.com/gocacher/cacher"
)

var DefaultCachePath = "cache"

var _cache = cache.NewBadgerCache(DefaultCachePath)

// RegisterCache ...
func RegisterCache() {
	cacher.Register(_cache)
}
