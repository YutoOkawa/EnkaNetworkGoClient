package client

import "sync"

type Cache interface {
	Get(key string) []byte
	Set(key string, value []byte)
}

type UserDataCache struct {
	mutex *sync.RWMutex
	cache map[string][]byte
}

func (c *UserDataCache) Get(key string) []byte {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	cache, ok := c.cache[key]
	if !ok {
		return nil
	}
	return cache
}

func (c *UserDataCache) Set(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = value
}
