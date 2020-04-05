package common

import (
	"time"
)

type CacheStore struct {
	Store		[]CacheItem
	Metas		[]CacheItemMeta
}

func (store *CacheStore) Add(item CacheItem) {
	now := time.Now().
		Add(time.Second * time.Duration(15)).
		Unix()

	itemMeta := CacheItemMeta{Expires: now}
	store.Store = append(store.Store, item)
	store.Metas = append(store.Metas, itemMeta)
}

func (store *CacheStore) DeepCopy() *CacheStore {
	metas := store.Metas
	storeItems := store.Store

	return &CacheStore{
		Store:	storeItems,
		Metas:	metas,
	}
}

func (store *CacheStore) Populate(channel chan<- int) {
	for i := range store.Store {
		channel <- i
	}
}

func NewStore() *CacheStore {
	return &CacheStore{
		Store:	make([]CacheItem, 0),
		Metas:	make([]CacheItemMeta, 0),
	}
}