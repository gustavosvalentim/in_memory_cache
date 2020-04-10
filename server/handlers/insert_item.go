package handlers

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// Item : 
type Item struct {
	Key		string
	Data	interface{}
}

// InsertItemHandler will accept an CacheItem JSON object and after Unmarshal will add to items
func InsertItemHandler(w *http.ResponseWriter, r *http.Request, store *common.CacheStore) {
	b, err := ioutil.ReadAll(r.Body)

	if (err != nil) {
		fmt.Fprintf(*w, err.Error())
		return
	}

	var item Item

	if err := json.Unmarshal(b, &item); err != nil {
		fmt.Fprintf(*w, err.Error())
		return
	}

	store.Put(item.Key, item.Data)

	fmt.Fprintf(*w, string(b))
}