package common

import (
	"io"
	"io/ioutil"
	"encoding/json"
)

// CacheItem represents an entire record
type CacheItem map[string]interface{}

// CacheItemMeta contains information about some item in CacheStore
type CacheItemMeta struct {
	Expires		int64
}

// Marshal will encode a CacheItem type
func (item *CacheItem) Marshal() (string, error) {
	b, err := json.Marshal(*item)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// NewCacheItem Unmarshal a valid http.Request.Body to a CacheItem
func NewCacheItem(requestBody io.ReadCloser) (CacheItem, error) {
	var cacheItem CacheItem
	bBody, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return cacheItem, err
	}
	if unmarshalErr := json.Unmarshal(bBody, &cacheItem); unmarshalErr != nil {
		return cacheItem, err
	}

	return cacheItem, nil
}