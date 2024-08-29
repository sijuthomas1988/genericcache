package genericcache

import (
	"sync"

	list "github.com/sijuthomas1988/genericcache/internal/list"
	lru "github.com/sijuthomas1988/genericcache/internal/lru"
)

var _ lru.LRU[int, int] = &Cache[int, int]{}

// Cache is a thread safe lru cache. It automatically removes elements as new elements are
// added if the capacity is reached. Items are removes based on how recently
// they were used where the oldest items are removed first.
type Cache[K comparable, V any] struct {
	ll      *list.List[entry[K, V]]
	items   map[K]*list.Node[entry[K, V]]
	options *options
	mu      sync.RWMutex
}

type entry[K comparable, V any] struct {
	key   K
	value V
}

// New initializes a new lru cache with the given capacity.
func New[K comparable, V any](cacheOptions ...CacheOption) *Cache[K, V] {
	c := &Cache[K, V]{
		ll:      list.NewList[entry[K, V]](),
		items:   make(map[K]*list.Node[entry[K, V]]),
		options: defaultOptions(),
		mu:      sync.RWMutex{},
	}

	for _, option := range cacheOptions {
		option.apply(c.options)
	}

	return c
}

// Len is the number of key value pairs in the cache.
func (c *Cache[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.ll.Len()
}

// Set the given key value pair.
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := entry[K, V]{
		key:   key,
		value: value,
	}

	e := c.ll.InsertNewNode(entry)
	if c.ll.Len() > c.options.capacity {
		c.deleteElement(c.ll.GetLastNode())
	}
	c.items[key] = e
}

// Get an item from the cache.
func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.items[key]
	if !ok {
		return
	}

	return e.Value.value, true
}

// Peek gets an item from the cache without updating the recent usage.
func (c *Cache[K, V]) Peek(key K) (value V, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.items[key]
	if !ok {
		return
	}

	return e.Value.value, true
}

// Delete an item from the cache.
func (c *Cache[K, V]) Delete(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.items[key]
	if !ok {
		return false
	}

	c.deleteElement(e)

	return true
}

func (c *Cache[K, V]) deleteElement(e *list.Node[entry[K, V]]) {
	delete(c.items, e.Value.key)
	c.ll.Remove(e)
}

// Flush deletes all items from the cache.
func (c *Cache[K, V]) Flush() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.ll.Init()
	c.items = make(map[K]*list.Node[entry[K, V]])
}
