package cache

import (
	"hash/fnv"
)

type ShardedCache struct {
	shards    []*LRUCache
	numShards int
}

func NewShardedCache(shardCount int, shardCapacity int) *ShardedCache {
	shards := make([]*LRUCache, shardCount)
	for i := 0; i < shardCount; i++ {
		shards[i] = NewLRUCache(shardCapacity)
	}
	return &ShardedCache{
		shards:    shards,
		numShards: shardCount,
	}
}

func (sc *ShardedCache) getShard(blockID string) *LRUCache {
	h := fnv.New32a()
	h.Write([]byte(blockID))
	return sc.shards[int(h.Sum32())%sc.numShards]
}

func (sc *ShardedCache) Get(blockID string) (Block, bool) {
	return sc.getShard(blockID).Get(blockID)
}

func (sc *ShardedCache) Put(blockID string, block Block) {
	sc.getShard(blockID).Put(blockID, block)
}
