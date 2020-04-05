package server

import (
	"log"
	"net/http"

	"github.com/gustavosvalentim/in_memory_cache/server/handlers"
	"github.com/gustavosvalentim/in_memory_cache/server/middlewares"
	"github.com/gustavosvalentim/in_memory_cache/common"

	"github.com/gorilla/mux"
)

// Handler represents a handler function that receives the items
type Handler = func(w *http.ResponseWriter, r *http.Request, items *common.CacheStore)

// WrappedHandler is a type Handler after being wrapped by a dependency injetor
type WrappedHandler = func(w http.ResponseWriter, r *http.Request)

// Serve starts the HTTP server
func Serve(store *common.CacheStore) {
	wrapHandler := func(handler Handler) WrappedHandler {
		return func(w http.ResponseWriter, r *http.Request) {
			handler(&w, r, store)
		}
	}

	r := mux.NewRouter()

	itemsRouter := r.PathPrefix("/items").Subrouter()
	itemsRouter.Methods("POST").HandlerFunc(wrapHandler(handlers.InsertItemHandler))
	itemsRouter.Methods("GET").HandlerFunc(wrapHandler(handlers.GetItemHandler))
	itemsRouter.Methods("DELETE").HandlerFunc(wrapHandler(handlers.InvalidateCacheHandler))

	populateRoute := r.PathPrefix("/populate").Subrouter()
	populateRoute.Methods("POST").HandlerFunc(wrapHandler(handlers.PopulateHandler))

	r.Use(middlewares.JSONMiddleware)

	log.Fatal(http.ListenAndServe(":8090", r))
}