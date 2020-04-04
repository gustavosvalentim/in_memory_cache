package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.io/gustavosvalentim/in_memory_cache/common"
)

func InsertItemHandler(w *http.ResponseWriter, r *http.Request, items *[][]common.CacheItem) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(*w, "Error")
		return
	}
	unmarshal, err := common.UnmarshalCacheItem(body)
	if err != nil {
		fmt.Fprintf(*w, "Unmarshal Error")
		return
	}

	*items = append(*items, unmarshal)

	json.NewEncoder(*w).Encode(unmarshal)
}