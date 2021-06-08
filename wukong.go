package wukong

// Wukong 描述一个缓存
type Wukong interface {
	// Get 从缓存中取得对象
	Get(key string) (value interface{}, err error)

	// Set 设置缓存值
	Set(key string, value interface{}, options ...option) (err error)

	// Delete 从缓存中删除一个对象
	Delete(key string) (err error)

	// Invalidate 让缓存失效
	Invalidate(opts ...option) (err error)

	// Clear 清空缓存
	Clear() (err error)

	// Serializer 设置序列化器
	Serializer(serializer Serializer)

	// Type 缓存类型
	Type() Type
}
