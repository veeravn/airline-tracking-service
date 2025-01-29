package utils

import (
	"sync"
	"time"
)

type cacheItem struct {
	data      interface{}
	timestamp time.Time
}

var cache = make(map[string]cacheItem)
var mu sync.Mutex

func AddToCache(key string, value interface{}) {
	mu.Lock()
	defer mu.Unlock()
	cache[key] = cacheItem{data: value, timestamp: time.Now()}
}

func GetFromCache(key string) interface{} {
	mu.Lock()
	defer mu.Unlock()
	item, found := cache[key]
	if !found || time.Since(item.timestamp) > 5*time.Minute {
		return nil
	}
	return item.data
}
