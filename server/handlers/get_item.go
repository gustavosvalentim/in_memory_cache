package handlers

import (
	"fmt"
	// "strings"
	"encoding/json"
	"net/http"

	"github.com/gustavosvalentim/in_memory_cache/common"

	"github.com/gorilla/mux"
)

// GetItemError : 
type GetItemError struct {
	Message		string
}

// GetItemResponse : 
type GetItemResponse struct {
	Key			string
	Data		interface{}
}

// GetItemHandler returns CacheItems stored in items
func GetItemHandler(w *http.ResponseWriter, r *http.Request, store *common.CacheStore) {
	key := mux.Vars(r)["key"]
	item, ok := store.Get(key)
	if (ok == false) {
		b, _ := json.Marshal(GetItemError{
			Message: "item not found",
		})
		fmt.Fprintf(*w, string(b))
		return
	}

	b, _ := json.Marshal(GetItemResponse{
		Key: key,
		Data: item,
	})
	fmt.Fprintf(*w, string(b))
}