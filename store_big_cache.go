package wukong

import (
	`fmt`
	`strings`
	`time`

	`github.com/allegro/bigcache`
)

const (
	// BigCacheType 描述BigCache的存储类型的字符串
	BigCacheType = "bigCache"
	// BigCacheTagFormatter 用于生成标签
	BigCacheTagFormatter = "wukong-tag-%s"
)

var (
	_ Store = (*storeBigCache)(nil)
	_       = NewBigCache(bigcache.BigCache{})
)

// storeBigCache 基于BigCache的缓存存储
type storeBigCache struct {
	client  bigcache.BigCache
	options options
}

// NewBigCache 创建一个BigCache存储
func NewBigCache(client bigcache.BigCache, options ...option) Store {
	appliedOptions := defaultOptions()
	for _, option := range options {
		option.apply(&appliedOptions)
	}

	return &storeBigCache{
		client:  client,
		options: appliedOptions,
	}
}

func (sr *storeBigCache) Get(key string) ([]byte, error) {
	return sr.client.Get(key)
}

func (sr *storeBigCache) GetWithTTL(key string) (data []byte, ttl time.Duration, err error) {
	if data, err = sr.client.Get(key); nil != err {
		return
	}

	return
}

func (sr *storeBigCache) Set(key string, data []byte, options ...option) (err error) {
	newOptions := sr.options
	for _, option := range options {
		option.apply(&newOptions)
	}

	if err = sr.client.Set(key, data); nil != err {
		return
	}

	if tags := newOptions.Tags; len(tags) > 0 {
		err = sr.setTags(key, tags...)
	}

	return
}

func (sr *storeBigCache) setTags(key string, tags ...string) (err error) {
	for _, tag := range tags {
		var tagKey = fmt.Sprintf(BigCacheTagFormatter, tag)
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

func (sr *storeBigCache) getCacheKeysForTag(tagKey string) (keys []string) {
	if result, err := sr.Get(tagKey); nil != err && "" != string(result) {
		keys = strings.Split(string(result), ",")
	}

	return
}

func (sr *storeBigCache) Delete(key string) (err error) {
	return sr.client.Delete(key)
}

func (sr *storeBigCache) Invalidate(options ...invalidateOption) (err error) {
	appliedOptions := defaultInvalidateOptions()
	for _, option := range options {
		option.applyInvalidate(&appliedOptions)
	}

	for _, tag := range appliedOptions.Tags {
		var tagKey = fmt.Sprintf(BigCacheTagFormatter, tag)
		var cacheKeys = sr.getCacheKeysForTag(tagKey)

		for _, cacheKey := range cacheKeys {
			err = sr.Delete(cacheKey)
		}

		err = sr.Delete(tagKey)
	}

	return
}

func (sr *storeBigCache) Type() string {
	return BigCacheType
}

func (sr *storeBigCache) Clear() (err error) {
	return sr.client.Reset()
}
