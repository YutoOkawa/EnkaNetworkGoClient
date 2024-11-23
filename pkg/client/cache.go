package client

type Cache interface {
	Get(key string) []byte
	Set(key string, value []byte)
}

type UserDataCache struct {
	cache map[string][]byte
}

func (c *UserDataCache) Get(key string) []byte {
	return c.cache[key]
}

func (c *UserDataCache) Set(key string, value []byte) {
	c.cache[key] = value
}
