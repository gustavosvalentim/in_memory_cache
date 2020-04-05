package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

func PopulateHandler(w *http.ResponseWriter, r *http.Request, store *common.CacheStore) {
	var cacheItems []common.CacheItem
	body, err := ioutil.ReadAll(r.Body)
	if (err != nil) {
		fmt.Fprintf(*w, "Cannot read body")
		return
	}
	if unmarshalErr := json.Unmarshal(body, &cacheItems); unmarshalErr != nil {
		fmt.Fprintf(*w, "Unmarshal error")
		return
	}
	store.Populate(cacheItems)
	fmt.Fprintf(*w, string(body))
}