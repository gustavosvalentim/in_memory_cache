package handlers

import (
	"fmt"
	"net/http"
	// "encoding/json"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// InvalidateCacheHandler will invalidate all cache, making items as an array of CacheItem
func InvalidateCacheHandler(w *http.ResponseWriter, r *http.Request, store *common.CacheStore) {
	fmt.Fprintf(*w, "Will get back soon")
}