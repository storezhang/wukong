package wukong

import (
	`fmt`
	`strings`
	`time`

	`github.com/bradfitz/gomemcache/memcache`
)

const (
	// MemcacheType 描述Memcache的存储类型的字符串
	MemcacheType = "memcache"
	// MemcacheTagFormatter 用于生成标签
	MemcacheTagFormatter = "wukong-tag-%s"
)

var _ Store = (*storeMemcache)(nil)

// storeMemcache 基于Memcache的缓存存储
type storeMemcache struct {
	client  *memcache.Client
	options options
}

// NewMemcache 创建一个Memcache存储
func NewMemcache(client *memcache.Client, options ...option) Store {
	appliedOptions := defaultOptions()
	for _, option := range options {
		option.apply(&appliedOptions)
	}

	return &storeMemcache{
		client:  client,
		options: appliedOptions,
	}
}

func (sr *storeMemcache) Get(key string) (data []byte, err error) {
	var item *memcache.Item
	if item, err = sr.client.Get(key); nil != err {
		return
	}
	data = item.Value

	return
}

func (sr *storeMemcache) GetWithTTL(key string) (data []byte, ttl time.Duration, err error) {
	var item *memcache.Item
	if item, err = sr.client.Get(key); nil != err {
		return
	}
	data = item.Value
	ttl = time.Duration(item.Expiration)

	return
}

func (sr *storeMemcache) Set(key string, data []byte, options ...option) (err error) {
	newOptions := sr.options
	for _, option := range options {
		option.apply(&newOptions)
	}

	if err = sr.client.Set(&memcache.Item{
		Key:        key,
		Value:      data,
		Flags:      0,
		Expiration: int32(newOptions.Expiration.Seconds()),
	}); nil != err {
		return
	}

	if tags := newOptions.Tags; len(tags) > 0 {
		err = sr.setTags(key, tags...)
	}

	return
}

func (sr *storeMemcache) setTags(key string, tags ...string) (err error) {
	for _, tag := range tags {
		var tagKey = fmt.Sprintf(MemcacheTagFormatter, tag)
		var cacheKeys = sr.getCacheKeysForTag(tagKey)

		var alreadyInserted = false
		for _, cacheKey := range cacheKeys {
			if cacheKey == key {
				alreadyInserted = true

				break
			}
		}

		if !alreadyInserted {
			cacheKeys = append(cacheKeys, key)
		}

		err = sr.Set(tagKey, []byte(strings.Join(cacheKeys, ",")), WithExpiration(720*time.Hour))
	}

	return
}

func (sr *storeMemcache) getCacheKeysForTag(tagKey string) (keys []string) {
	if result, err := sr.Get(tagKey); nil != err && "" != string(result) {
		keys = strings.Split(string(result), ",")
	}

	return
}

func (sr *storeMemcache) Delete(key string) (err error) {
	return sr.client.Delete(key)
}

func (sr *storeMemcache) Invalidate(options ...invalidateOption) (err error) {
	appliedOptions := defaultInvalidateOptions()
	for _, option := range options {
		option.applyInvalidate(&appliedOptions)
	}

	for _, tag := range appliedOptions.Tags {
		var tagKey = fmt.Sprintf(MemcacheTagFormatter, tag)
		var cacheKeys = sr.getCacheKeysForTag(tagKey)

		for _, cacheKey := range cacheKeys {
			err = sr.Delete(cacheKey)
		}

		err = sr.Delete(tagKey)
	}

	return
}

func (sr *storeMemcache) Type() string {
	return MemcacheType
}

func (sr *storeMemcache) Clear() (err error) {
	return sr.client.FlushAll()
}
