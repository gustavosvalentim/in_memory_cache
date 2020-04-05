package handlers

import (
	"fmt"
	"net/http"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// InsertItemHandler will accept an CacheItem JSON object and after Unmarshal will add to items
func InsertItemHandler(w *http.ResponseWriter, r *http.Request, store *common.CacheStore) {
	cacheItem, err := common.NewCacheItem(r.Body)
	if err != nil {
		fmt.Fprintf(*w, err.Error())
		return
	}

	store.Add(cacheItem)

	marshalItem, err := cacheItem.Marshal()
	if err != nil {
		fmt.Fprintf(*w, err.Error())
		return
	}
	fmt.Fprintf(*w, marshalItem)
}