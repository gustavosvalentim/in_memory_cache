package common

import (
	"time"
	"sync"
	"math/rand"
)

type CacheStoreItems map[int]CacheItem
type CacheStoreMetas map[int]CacheItemMeta

type Shard struct {
	Items		CacheStoreItems
	Metas		CacheStoreMetas
	Lock		*sync.RWMutex
}

type CacheStore struct {
	Shards		[]Shard
}

func (shard *Shard) add(item CacheItem, meta CacheItemMeta) {
	randint := rand.Int()
	shard.Lock.Lock()
	shard.Items[randint] = item
	shard.Metas[randint] = meta
	shard.Lock.Unlock()
}

func (shard *Shard) remove(i int) {
	shard.Lock.Lock()
	delete(shard.Items, i)
	delete(shard.Metas, i)
	shard.Lock.Unlock()
}

func (store *CacheStore) Add(item CacheItem, expiration int64) {
	shard := store.Shards[len(store.Shards) - 1]
	itemMeta := CacheItemMeta{Expires: expiration}
	shard.add(item, itemMeta)
}

func (store *CacheStore) Populate(items []CacheItem) {
	now := time.Now().
		Add(time.Second * time.Duration(15)).
		Unix()

	for i := range items {
		store.Add(items[i], now)
	}
}

func (store *CacheStore) GetItems() *[]CacheItem {
	items := make([]CacheItem, 0)

	for _, shard := range store.Shards {
		for _, item := range shard.Items {
			items = append(items, item)
		}
	}

	return &items
}

func (store *CacheStore) Clean(t int64) {
	for i, v := range store.Shards {
		for mi, m := range v.Metas {
			if (t >= m.Expires) {
				store.Shards[i].remove(mi)
			}
		}
	}
}

func (store *CacheStore) Invalidate() {
	*store = *NewStore()
}

func NewStore() *CacheStore {
	shards := make([]Shard, 0)
	shards = append(shards, Shard{
		Items: make(CacheStoreItems, 0),
		Metas: make(CacheStoreMetas, 0),
		Lock:  &sync.RWMutex{},
	})
	return &CacheStore{
		Shards: shards,
	}
}