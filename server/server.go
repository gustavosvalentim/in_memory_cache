package server

import (
	"net/http"

	"github.io/gustavosvalentim/in_memory_cache/server/handlers"
	"github.io/gustavosvalentim/in_memory_cache/common"
)

func Server() {
	items := make([][]common.CacheItem, 0)

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				handlers.GetItemHandler(&w, r, &items)
			case "POST":
				handlers.InsertItemHandler(&w, r, &items)
		}
	})

	http.ListenAndServe(":8090", nil)
}