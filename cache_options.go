package wukong

type (
	// CacheOption 配置选项
	CacheOption func(*cacheOptions)

	cacheOptions struct {
		options

		// Serializer 序列化器
		Serializer Serializer
	}
)

func defaultCacheOptions() cacheOptions {
	return cacheOptions{
		Serializer: &SerializerJson{},
	}
}

// WithSerializer 配置序列化器
func WithSerializer(serializer Serializer) CacheOption {
	return func(options *cacheOptions) {
		options.Serializer = serializer
	}
}
