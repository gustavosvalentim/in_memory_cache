package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.io/gustavosvalentim/in_memory_cache/common"
)

func GetItemHandler(w *http.ResponseWriter, r *http.Request, items *[][]common.CacheItem) {
	b, err := json.Marshal(*items)
	if err != nil {
		fmt.Fprintf(*w, "Error")
		return
	}
	fmt.Fprintf(*w, string(b))
}