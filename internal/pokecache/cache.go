package cache


import (
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{ 
		entries: make(map[string]cacheEntry),
		interval: interval,
	}	
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	// buf is now a separate slice with its own backing array containing the same data as val.
	// Storing buf in your cache ensures that later changes to val will NOT affect the cached copy.
	buf := make([]byte, len(val)) 
	copy(buf, val) 

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: buf}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
    	<-ticker.C
		c.mu.Lock()
		for key := range c.entries{
			elapsedTime := time.Since(c.entries[key].createdAt)
			if elapsedTime > c.interval {
				delete(c.entries, key)
			}
    	}
		c.mu.Unlock()
	}
}
