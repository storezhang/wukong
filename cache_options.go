package wukong

type cacheOptions struct {
	options

	// Serializer 序列化器
	Serializer Serializer
}

func defaultCacheOptions() cacheOptions {
	return cacheOptions{
		Serializer: &SerializerMsgpack{},
	}
}
