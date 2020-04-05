package services

import (
	"time"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// CacheCleaner tick every 1 second and receive values in channel
func CacheCleaner(store *common.CacheStore) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				storeCopy := store.DeepCopy()
				for i := 0; i < len(storeCopy.Metas); i++ {
					meta := storeCopy.Metas[i]
					timestamp := meta.Expires
					now := time.Now().Unix()
		
					if now - timestamp == 0 {
						store.Metas = append(store.Metas[:i], store.Metas[i+1:]...)
						store.Store = append(store.Store[:i], store.Store[i+1:]...)
					}
				}
			}
		}
	}()
}