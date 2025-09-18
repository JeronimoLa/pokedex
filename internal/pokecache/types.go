package cache

import (
	"sync"
	"time"	
)

type Cache struct {
	mu    		sync.Mutex
	entries 	map[string]cacheEntry
	interval	time.Duration
}

type cacheEntry struct {
	createdAt 	time.Time
	val			[]byte
}