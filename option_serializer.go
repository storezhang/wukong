package wukong

type optionSerializer struct {
	serializer Serializer
}

// WithSerializer 配置序列化器
func WithSerializer(serializer Serializer) *optionSerializer {
	return &optionSerializer{serializer: serializer}
}

func (os *optionSerializer) applyCache(options *cacheOptions) {
	options.Serializer = os.serializer
}
