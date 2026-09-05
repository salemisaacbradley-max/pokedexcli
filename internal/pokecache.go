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

func NewCache (interval time.duration) {
	new := Cache{}
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
	
}