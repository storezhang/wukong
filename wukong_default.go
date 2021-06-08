package wukong

type wukongDefault struct {
	store      Store
	serializer Serializer
}

// New 创建一个普通缓存
func New(store Store) Wukong {
	return &wukongDefault{
		store:      store,
		serializer: &serializerGob{},
	}
}

func (wd *wukongDefault) Get(key string) (value interface{}, err error) {
	var data []byte
	if data, err = wd.store.Get(key); nil != err {
		return
	}
	value, err = wd.serializer.Decode(data)

	return
}

func (wd *wukongDefault) Set(key string, value interface{}, opts ...option) (err error) {
	options := defaultOptions()
	for _, option := range opts {
		option.apply(options)
	}

	var data []byte
	if data, err = options.serializer.Encode(value); nil != err {
		return
	}
	err = wd.store.Set(key, data, options.expiration, options.tags...)

	return
}

func (wd *wukongDefault) Delete(key string) (err error) {
	return wd.store.Delete(key)
}

func (wd *wukongDefault) Invalidate(opts ...option) (err error) {
	options := defaultOptions()
	for _, option := range opts {
		option.apply(options)
	}

	return wd.store.Invalidate(options.tags...)
}

func (wd *wukongDefault) Serializer(serializer Serializer) {
	wd.serializer = serializer
}

func (wd *wukongDefault) Type() Type {
	return wd.store.Type()
}

func (wd *wukongDefault) Clear() (err error) {
	return wd.store.Clear()
}
