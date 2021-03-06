package wukong

import (
	`time`
)

// Store 描述一个缓存真正的存储
type Store interface {
	// Get 取得缓存值
	Get(key string) ([]byte, error)

	// Set 设置缓存值
	Set(key string, data []byte, expiration time.Duration, tags ...string) error

	// Delete 删除缓存值
	Delete(key string) error

	// Invalidate 设置缓存失效
	Invalidate(tags ...string) error

	// Clear 清除缓存
	Clear() error

	// Type 缓存类型
	Type() Type
}
