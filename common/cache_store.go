package common

import (
	"time"
	"math/rand"
)

type CacheStoreItems map[int]CacheItem
type CacheStoreMetas map[int]CacheItemMeta

type CacheStore struct {
	Items		CacheStoreItems
	Metas		CacheStoreMetas
}

func (store *CacheStore) Add(item CacheItem, expiration int64) {
	itemMeta := CacheItemMeta{Expires: expiration}
	randint := rand.Int()
	store.Items[randint] = item
	store.Metas[randint] = itemMeta
}

func (store *CacheStore) DeepCopy() *CacheStore {
	metas := store.Metas
	storeItems := store.Items

	return &CacheStore{
		Items:	storeItems,
		Metas:	metas,
	}
}

func (store *CacheStore) Slice(i int) {
	delete(store.Items, i)
	delete(store.Metas, i)
}

func (store *CacheStore) Populate(items []CacheItem) {
	now := time.Now().
		Add(time.Second * time.Duration(15)).
		Unix()

	for i := range items {
		store.Add(items[i], now)
	}
}

func NewStore() *CacheStore {
	return &CacheStore{
		Items:	make(CacheStoreItems, 0),
		Metas:	make(CacheStoreMetas, 0),
	}
}