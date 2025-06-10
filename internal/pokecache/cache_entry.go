package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct{
	createdAt	 time.Time
	val		[]byte
}

type Cache struct{
	mu		sync.Mutex
	cacheEntries	map[string]cacheEntry
}


func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		cacheEntries: map[string]cacheEntry{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	cacheEntry, exists := c.cacheEntries[key]
	c.mu.Unlock()
	return cacheEntry.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker:= time.NewTicker(interval)
	for t := range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cacheEntries {
			timePlusInterval := entry.createdAt.Add(interval)
			if timePlusInterval.Before(t)  {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
