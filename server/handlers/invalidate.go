package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// InvalidateCacheHandler will invalidate all cache, making items as an array of CacheItem
func InvalidateCacheHandler(w *http.ResponseWriter, r *http.Request, store *common.CacheStore) {
	store.Invalidate()
	b, err := json.Marshal(store.GetItems())
	if err != nil {
		fmt.Fprintln(*w, "Error")
		return
	}
	fmt.Fprintf(*w, string(b))
}