package handlers

import (
	"fmt"
	"time"
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

	now := time.Now().
		Add(time.Second * time.Duration(15)).
		Unix()

	store.Add(cacheItem, now)

	marshalItem, err := cacheItem.Marshal()
	if err != nil {
		fmt.Fprintf(*w, err.Error())
		return
	}
	fmt.Fprintf(*w, marshalItem)
}