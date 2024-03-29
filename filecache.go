package management

import (
	"io/ioutil"

	cache "github.com/gocacher/badger-cache"
	"github.com/gocacher/cacher"
)

// DefaultCachePath ...
var DefaultCachePath = "cache"

var _cache cacher.Cacher

// RegisterCache ...
func RegisterCache() {
	_cache = cache.NewBadgerCache(DefaultCachePath)
	cacher.Register(_cache)
}

// CacheFile ...
func CacheFile(hash, path string) error {
	bys, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
	e = _cache.Set(hash, bys)
	if e != nil {
		return e
	}
	return nil
}

// GetCache ...
func GetCache(hash string) ([]byte, error) {
	return _cache.Get(hash)
}
