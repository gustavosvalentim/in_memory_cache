package common

import (
	"time"
)

// CacheStore : 
type CacheStore struct {
	*HashMap
	DefaultTTL			time.Duration
}

// NewCacheStore : 
func NewCacheStore(ttlInMinutes int) *CacheStore {
	hashmap := NewHashMap(100)

	return &CacheStore{
		HashMap: hashmap,
		DefaultTTL: time.Duration(ttlInMinutes) * time.Minute,
	}
}