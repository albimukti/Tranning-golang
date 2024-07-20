package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var C *cache.Cache

// InitCache initializes the cache with a default expiration time and cleanup interval.
func InitCache() {
	C = cache.New(5*time.Minute, 10*time.Minute)
}

// Set adds a key-value pair to the cache with default expiration.
func Set(key string, value interface{}) {
	C.Set(key, value, cache.DefaultExpiration) // You can also specify a custom duration here
}

// Get retrieves the value associated with the key from the cache.
func Get(key string) (interface{}, bool) {
	return C.Get(key)
}
