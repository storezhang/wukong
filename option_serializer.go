package wukong

var _ option = (*optionSerializer)(nil)

type optionSerializer struct {
	serializer Serializer
}

// Gob Golang内置序列化器
func Gob() *optionSerializer {
	return &optionSerializer{
		serializer: &serializerGob{},
	}
}

// Json Json列化器
func Json() *optionSerializer {
	return &optionSerializer{
		serializer: &serializerJson{},
	}
}

// Msgpack Msgpack列化器，比Json生成的序列化数据更简短
func Msgpack() *optionSerializer {
	return &optionSerializer{
		serializer: &serializerGob{},
	}
}

func (os *optionSerializer) apply(options *options) {
	options.serializer = os.serializer
}
