package services

import (
	"fmt"
	"time"
	"sync"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// CacheCleaner tick every 1 second and receive values in channel
func CacheCleaner(store *common.CacheStore) {
	ticker := time.NewTicker(1 * time.Second)
	cacheCleanerStore := make(chan int, 1000000)
	mutex := &sync.Mutex{}

	go func() {
		for {
			select {
			case i := <-cacheCleanerStore:
				mutex.Lock()
				if (len(store.Metas) - 1 <= i) {
					expires := store.Metas[i].Expires
					now := time.Now().Unix()
		
					if now >= expires {
						store.Metas = append(store.Metas[:i], store.Metas[i+1:]...)
						store.Store = append(store.Store[:i], store.Store[i+1:]...)
					}
					fmt.Println(store.Metas)
				}
				mutex.Unlock()
			case <-ticker.C:
				store.Populate(cacheCleanerStore)
			}
		}
	}()
}