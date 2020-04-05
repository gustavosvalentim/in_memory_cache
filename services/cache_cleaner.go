package services

import (
	"fmt"
	"time"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// CacheCleaner tick every 1 second and receive values in channel
func CacheCleaner(store *common.CacheStore) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t, len(store.Metas))
				for i := range store.Metas {
					if val, ok := store.Metas[i]; ok {
						if (t.Unix() >= val.Expires) {
							go store.Slice(i)
						}
					}
				}
			}
		}
	}()
}