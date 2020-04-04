package server

import (
	"log"
	"net/http"

	"github.io/gustavosvalentim/in_memory_cache/server/handlers"
	"github.io/gustavosvalentim/in_memory_cache/common"

	"github.com/gorilla/mux"
)

type Handler = func(w *http.ResponseWriter, r *http.Request, items *[][]common.CacheItem)
type WrappedHandler = func(w http.ResponseWriter, r *http.Request)

func Server() {
	items := make([][]common.CacheItem, 0)

	wrapHandler := func(handler Handler) WrappedHandler {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			handler(&w, r, &items)
		}
	}

	r := mux.NewRouter().
		Headers("Content-Type", "application/json").
		Subrouter()

	r.HandleFunc("/items", wrapHandler(handlers.InsertItemHandler)).Methods("POST")
	r.HandleFunc("/items", wrapHandler(handlers.GetItemHandler)).Methods("GET")
	r.HandleFunc("/items", wrapHandler(handlers.InvalidateCacheHandler)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8090", r))
}