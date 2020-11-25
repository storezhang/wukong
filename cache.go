package wukong

// Cache 描述一个缓存
type Cache interface {
	// Get 从缓存中取得对象
	Get(key string, value interface{}, ) error
	// Set 设置缓存值
	Set(key string, value interface{}, options ...option) error
	// Delete 从缓存中删除一个对象
	Delete(key string) error
	// Invalidate 让缓存失效
	Invalidate(options ...invalidateOption) error
	// Clear 清空缓存
	Clear() error
	// Type 缓存类型
	Type() string
}
