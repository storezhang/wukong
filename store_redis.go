package wukong

import (
	`context`
	`time`

	`github.com/go-redis/redis/v8`
)

var _ Store = (*storeRedis)(nil)

// storeRedis 基于Redis的缓存存储
type storeRedis struct {
	storeBase

	client *redis.Client
}

// Redis 创建一个Redis存储
func Redis(client *redis.Client) *storeRedis {
	return &storeRedis{
		client: client,
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

func (sr *storeRedis) Set(key string, data []byte, expiration time.Duration, tags ...string) (err error) {
	if err = sr.client.Set(context.Background(), key, data, expiration).Err(); nil != err {
		return
	}

	if len(tags) > 0 {
		err = sr.setTags(key, sr.Get, sr.Set, tags...)
	}

	return
}

func (sr *storeRedis) Delete(key string) (err error) {
	_, err = sr.client.Del(context.Background(), key).Result()

	return
}

func (sr *storeRedis) Invalidate(tags ...string) (err error) {
	return sr.invalidate(sr.Get, sr.Delete, tags...)
}

func (sr *storeRedis) Type() Type {
	return TypeRedis
}

func (sr *storeRedis) Clear() error {
	return sr.client.FlushAll(context.Background()).Err()
}
