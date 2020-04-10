package services

import (
	"fmt"
	"time"

	"github.com/gustavosvalentim/in_memory_cache/common"
)

// CacheCleaner tick every 1 second and receive values in channel
func CacheCleaner(store *common.CacheStore) {
	ticker := time.NewTicker(store.DefaultTTL)
	
	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("[*] ", t)

				items := store.Scan()
				for i := range items {
					item := items[i]
					if (item.Key != "") {
						fmt.Println("[*] Deleting item ID: ", item.Key)

						store.Delete(item.Key)
					}
				}
			}
		}
	}()
}