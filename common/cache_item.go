package common

import (
	"encoding/json"
)

type CacheItem struct {
	PropertyName	string
	PropertyValue	interface{}
}

func UnmarshalCacheItem(itemString []byte) ([]CacheItem, error) {
	cacheItem := make([]CacheItem, 0)
	if err := json.Unmarshal(itemString, &cacheItem); err != nil {
		return cacheItem, err
	}

	return cacheItem, nil
}

func MarshalCacheItem(item []CacheItem) (string, error) {
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}

	return string(b), nil
}