package pokecache

import ("sync"
		"time")

type Cache struct {
	mu sync.Mutex
	cache map[string]cacheEntry
	
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}


func NewCache (interval time.Duration) Cache {
	newCache := Cache{}
	go newCache.reapLoop(interval)
	return newCache
}

type Ticker struct {
	C <-chan time.Time

}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	t := time.Now()
	c.cache[key] = cacheEntry{t, val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.cache[key]
	if ok != true {
		return nil, ok
	} else {
		return value.val, ok
	}
}

func (c *Cache) reapLoop (interval time.Duration)  {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.cache {
			if time.Since(entry.createdAt) >= interval{
				delete(c.cache, key)
			}
		}
		c.mu.Unlock()
	}
}