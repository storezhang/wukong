package wukong

import (
	`context`
	`fmt`
	`strings`
	`time`

	`github.com/go-redis/redis/v8`
)

const (
	// RedisType 描述Redis的存储类型的字符串
	RedisType = "redis"
	// RedisTagFormatter 用于生成标签
	RedisTagFormatter = "wukong-tag-%s"
)

var _ Store = (*storeRedis)(nil)

// storeRedis 基于Redis的缓存存储
type storeRedis struct {
	client  *redis.Client
	options options
}

// NewRedis 创建一个Redis存储
func NewRedis(client *redis.Client, options ...option) Store {
	appliedOptions := defaultOptions()
	for _, option := range options {
		option.apply(&appliedOptions)
	}

	return &storeRedis{
		client:  client,
		options: appliedOptions,
	}
}

func (sr *storeRedis) Get(key string) ([]byte, error) {
	return sr.client.Get(context.Background(), key).Bytes()
}

func (sr *storeRedis) GetWithTTL(key string) (data []byte, ttl time.Duration, err error) {
	if data, err = sr.client.Get(context.Background(), key).Bytes(); nil != err {
		return
	}
	ttl, err = sr.client.TTL(context.Background(), key).Result()

	return
}

func (sr *storeRedis) Set(key string, data []byte, options ...option) (err error) {
	newOptions := sr.options
	for _, option := range options {
		option.apply(&newOptions)
	}

	if err = sr.client.Set(context.Background(), key, data, newOptions.Expiration).Err(); nil != err {
		return
	}

	if tags := newOptions.Tags; len(tags) > 0 {
		err = sr.setTags(key, tags...)
	}

	return
}

func (sr *storeRedis) setTags(key string, tags ...string) (err error) {
	for _, tag := range tags {
		var tagKey = fmt.Sprintf(RedisTagFormatter, tag)
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

func (sr *storeRedis) getCacheKeysForTag(tagKey string) (keys []string) {
	if result, err := sr.Get(tagKey); nil != err && "" != string(result) {
		keys = strings.Split(string(result), ",")
	}

	return
}

func (sr *storeRedis) Delete(key string) (err error) {
	_, err = sr.client.Del(context.Background(), key).Result()

	return
}

func (sr *storeRedis) Invalidate(options ...invalidateOption) (err error) {
	appliedOptions := defaultInvalidateOptions()
	for _, option := range options {
		option.applyInvalidate(&appliedOptions)
	}

	for _, tag := range appliedOptions.Tags {
		var tagKey = fmt.Sprintf(RedisTagFormatter, tag)
		var cacheKeys = sr.getCacheKeysForTag(tagKey)

		for _, cacheKey := range cacheKeys {
			err = sr.Delete(cacheKey)
		}

		err = sr.Delete(tagKey)
	}

	return
}

func (sr *storeRedis) Type() string {
	return RedisType
}

func (sr *storeRedis) Clear() (err error) {
	err = sr.client.FlushAll(context.Background()).Err()

	return
}
