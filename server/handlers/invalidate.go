package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.io/gustavosvalentim/in_memory_cache/common"
)

func InvalidateCacheHandler(w *http.ResponseWriter, r *http.Request, items *[][]common.CacheItem) {
	*items = make([][]common.CacheItem, 0)
	b, err := json.Marshal(*items)
	if err != nil {
		fmt.Fprintln(*w, "Error")
		return
	}
	fmt.Fprintf(*w, string(b))
}