package utils

import (
	"log"
	"time"

	"github.com/allegro/bigcache"
)

// BigCache interface of big cache object
type BigCache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

var config = bigcache.Config{
	// number of shards (must be a power of 2)
	Shards: 1024,

	// time after which entry can be evicted
	LifeWindow: 10 * time.Minute,

	// Interval between removing expired entries (clean up).
	// If set to <= 0 then no action is performed.
	// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
	CleanWindow: 5 * time.Minute,

	// rps * lifeWindow, used only in initial memory allocation
	MaxEntriesInWindow: 1000 * 10 * 60,

	// max entry size in bytes, used only in initial memory allocation
	MaxEntrySize: 500,

	// prints information about additional memory allocation
	Verbose: true,

	// cache will not allocate more memory than this limit, value in MB
	// if value is reached then the oldest entries can be overridden for the new ones
	// 0 value means no size limit
	HardMaxCacheSize: 8192,

	// callback fired when the oldest entry is removed because of its expiration time or no space left
	// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
	// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
	OnRemove: nil,

	// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
	// for the new entry, or because delete was called. A constant representing the reason will be passed through.
	// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
	// Ignored if OnRemove is specified.
	OnRemoveWithReason: nil,
}

type bigCache struct {
	cache *bigcache.BigCache
}

// NewBigCache func new big cache object
func NewBigCache() BigCache {
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		log.Fatal(err)
	}
	return &bigCache{cache: cache}
}

// Get func get the value of key
func (b *bigCache) Get(key string) (string, error) {
	value, err := b.cache.Get(key)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

// Set func set key to hold the string value
func (b *bigCache) Set(key string, value string) error {
	return b.cache.Set(key, []byte(value))
}
