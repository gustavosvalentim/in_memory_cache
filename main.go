package main

import (
	"github.com/gustavosvalentim/in_memory_cache/common"
	"github.com/gustavosvalentim/in_memory_cache/server"
	"github.com/gustavosvalentim/in_memory_cache/services"
)

func main() {
	store := common.NewStore()

	services.CacheCleaner(store)

	server.Serve(store)
}
