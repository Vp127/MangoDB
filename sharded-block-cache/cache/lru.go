package cache

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	capacity int
	ll       *list.List
	cache    map[string]*list.Element
	mu       sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func (lru *LRUCache) Get(blockID string) (Block, bool) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if elem, ok := lru.cache[blockID]; ok {
		lru.ll.MoveToFront(elem)
		return elem.Value.(entry).block, true
	}
	return nil, false
}

func (lru *LRUCache) Put(blockID string, block Block) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if elem, ok := lru.cache[blockID]; ok {
		lru.ll.MoveToFront(elem)
		elem.Value = entry{blockID, block}
		return
	}

	if lru.ll.Len() == lru.capacity {
		oldest := lru.ll.Back()
		if oldest != nil {
			lru.ll.Remove(oldest)
			delete(lru.cache, oldest.Value.(entry).blockID)
		}
	}

	elem := lru.ll.PushFront(entry{blockID, block})
	lru.cache[blockID] = elem
}
