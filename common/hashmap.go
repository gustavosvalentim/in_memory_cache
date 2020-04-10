package common

import (
	"sync"
)

// HashMapShard : 
type HashMapShard struct {
	List	*LinkedList
	Lock	*sync.RWMutex
}

// HashMap : HashMap data structure using Shards and LinkedLists
type HashMap struct {
	Size	int
	Count	int
	Shards	[]*HashMapShard
}

func hash(key string) uint64 {
	var h uint64

	for _, c := range key {
		h += uint64(c)
	}

	return h
}

// Loc : Locate the index of an item by using its hash
func (hashmap *HashMap) Loc(h uint64) int {
	return int(h) % hashmap.Size
}

// Put : 
func (hashmap *HashMap) Put(key string, value interface{}) {
	loc := hashmap.Loc(hash(key))
	shard := hashmap.Shards[loc]
	shard.Lock.Lock()
	defer shard.Lock.Unlock()

	ll := shard.List
	ll.Append(key, value)
	hashmap.Count++
}

// Get : Get a key from cache, returns the value and ok
func (hashmap *HashMap) Get(key string) (*Node, bool) {
	h := hash(key)
	loc := hashmap.Loc(h)
	if len(hashmap.Shards) <= loc {
		return nil, false
	}
	return hashmap.Shards[loc].List.Get(key)
}

// Scan : 
func (hashmap *HashMap) Scan() []*Node {
	var items []*Node
	for _, shard := range hashmap.Shards {
		shard.Lock.RLock()
		if (shard.List.Count > 0) {
			shardItems := shard.List.Scan()
			for _, item := range shardItems {
				items = append(items, item)
			}
		}
		shard.Lock.RUnlock()
	}

	return items
}

// Delete : 
func (hashmap *HashMap) Delete(key string) bool {
	loc := hashmap.Loc(hash(key))

	if len(hashmap.Shards) <= loc {
		return false
	}
	shard := hashmap.Shards[loc]
	shard.Lock.Lock()
	defer shard.Lock.Unlock()

	shard.List.Pop(key)
	hashmap.Count--

	return true
}

// NewHashMap : 
func NewHashMap(numberOfShards int) *HashMap {
	hashmap := &HashMap{
		Count: 0,
		Size: numberOfShards,
		Shards: make([]*HashMapShard, numberOfShards),
	}
	for i := range hashmap.Shards {
		hashmap.Shards[i] = &HashMapShard{
			List: NewLinkedList(),
			Lock: &sync.RWMutex{},
		}
	}

	return hashmap
}