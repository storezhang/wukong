package wukong

import (
	`time`

	`github.com/allegro/bigcache`
)

var _ Store = (*storeBigCache)(nil)

// storeBigCache 基于BigCache的缓存存储
type storeBigCache struct {
	storeBase

	client *bigcache.BigCache
}

// BigCache 创建一个BigCache存储
func BigCache(client *bigcache.BigCache) *storeBigCache {
	return &storeBigCache{
		client: client,
	}
}

func (sbc *storeBigCache) Get(key string) ([]byte, error) {
	return sbc.client.Get(key)
}

func (sbc *storeBigCache) GetWithTTL(key string) (data []byte, ttl time.Duration, err error) {
	if data, err = sbc.client.Get(key); nil != err {
		return
	}

	return
}

func (sbc *storeBigCache) Set(key string, data []byte, _ time.Duration, tags ...string) (err error) {
	if err = sbc.client.Set(key, data); nil != err {
		return
	}
	if len(tags) > 0 {
		err = sbc.setTags(key, sbc.Get, sbc.Set, tags...)
	}

	return
}

func (sbc *storeBigCache) Delete(key string) (err error) {
	return sbc.client.Delete(key)
}

func (sbc *storeBigCache) Invalidate(tags ...string) (err error) {
	return sbc.invalidate(sbc.Get, sbc.Delete, tags...)
}

func (sbc *storeBigCache) Type() Type {
	return TypeRedis
}

func (sbc *storeBigCache) Clear() (err error) {
	return sbc.client.Reset()
}
