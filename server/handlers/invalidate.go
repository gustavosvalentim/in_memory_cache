package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// InvalidateCacheHandler will invalidate all cache, making items as an array of CacheItem
func InvalidateCacheHandler(w *http.ResponseWriter, r *http.Request, store *common.CacheStore) {
	store.Store = make([]common.CacheItem, 0)
	store.Metas = make([]common.CacheItemMeta, 0)
	b, err := json.Marshal(*store)
	if err != nil {
		fmt.Fprintln(*w, "Error")
		return
	}
	fmt.Fprintf(*w, string(b))
}