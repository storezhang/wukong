package wukong

import (
	`time`

	`github.com/bradfitz/gomemcache/memcache`
)

var _ Store = (*storeMemcache)(nil)

// storeMemcache 基于Memcache的缓存存储
type storeMemcache struct {
	storeBase

	client *memcache.Client
}

// Memcache 创建一个Memcache存储
func Memcache(client *memcache.Client) *storeMemcache {
	return &storeMemcache{
		client: client,
	}
}

func (sm *storeMemcache) Get(key string) (data []byte, err error) {
	var item *memcache.Item
	if item, err = sm.client.Get(key); nil != err {
		return
	}
	data = item.Value

	return
}

func (sm *storeMemcache) GetWithTTL(key string) (data []byte, ttl time.Duration, err error) {
	var item *memcache.Item
	if item, err = sm.client.Get(key); nil != err {
		return
	}
	data = item.Value
	ttl = time.Duration(item.Expiration)

	return
}

func (sm *storeMemcache) Set(key string, data []byte, expiration time.Duration, tags ...string) (err error) {
	if err = sm.client.Set(&memcache.Item{
		Key:        key,
		Value:      data,
		Flags:      0,
		Expiration: int32(expiration.Seconds()),
	}); nil != err {
		return
	}

	if len(tags) > 0 {
		err = sm.setTags(key, sm.Get, sm.Set, tags...)
	}

	return
}

func (sm *storeMemcache) Delete(key string) (err error) {
	return sm.client.Delete(key)
}

func (sm *storeMemcache) Invalidate(tags ...string) (err error) {
	return sm.invalidate(sm.Get, sm.Delete, tags...)
}

func (sm *storeMemcache) Type() Type {
	return TypeMemcache
}

func (sm *storeMemcache) Clear() (err error) {
	return sm.client.FlushAll()
}
