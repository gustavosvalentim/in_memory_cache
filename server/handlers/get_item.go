package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// GetItemHandler returns CacheItems stored in items
func GetItemHandler(w *http.ResponseWriter, r *http.Request, store *common.CacheStore) {
	b, err := json.Marshal(store.Store)
	if err != nil {
		fmt.Fprintf(*w, "Error")
		return
	}
	fmt.Fprintf(*w, string(b))
}