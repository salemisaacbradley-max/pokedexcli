package pokecache

import ("sync"
		"time")

type Cache struct {
	mu sync.Mutex
	cache map[string]cacheEntry
	cacheEntry struct{
		createdAt time.Time
		val []byte
	}
}

func NewCache () {
	if 
	newCache := Cache{}
	go reapLoop(5*time.Second)

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

func (c *Cache) Get(key string) []byte, bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cache[key]
	if ok != true {
		return nil, ok
	} else {
		return val, ok
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