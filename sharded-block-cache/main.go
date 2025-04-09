package main

import (
	"fmt"
	"sharded-block-cache/cache"
)

func main() {
	sc := cache.NewShardedCache(4, 2)

	sc.Put("block1", cache.Block{"a": "1"})
	sc.Put("block2", cache.Block{"b": "2"})

	if b, ok := sc.Get("block1"); ok {
		fmt.Println("Found block1:", b)
	} else {
		fmt.Println("Block1 not found")
	}
}
