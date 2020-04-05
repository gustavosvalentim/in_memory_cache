package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// GetItemHandler returns CacheItems stored in items
func GetItemHandler(w *http.ResponseWriter, r *http.Request, store *common.CacheStore) {
	itemList := make([]common.CacheItem, 0)

	for _, shard := range store.Shards {
		for _, item := range shard.Items {
			itemList = append(itemList, item)
		}
	}

	b, err := json.Marshal(itemList)
	if err != nil {
		fmt.Fprintf(*w, "Error")
		return
	}
	fmt.Fprintf(*w, string(b))
}