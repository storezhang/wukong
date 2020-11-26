package wukong

const (
	// NormalType 普通缓存
	NormalType = "normal"
)

type cacheNormal struct {
	options cacheOptions
	store   Store
}

// New 创建一个普通缓存
func New(store Store, options ...cacheOption) Cache {
	appliedOptions := defaultCacheOptions()
	for _, option := range options {
		option.applyCache(&appliedOptions)
	}

	return &cacheNormal{
		options: appliedOptions,
		store:   store,
	}
}

func (cn *cacheNormal) Get(key string) (value interface{}, err error) {
	var data []byte

	if data, err = cn.store.Get(key); nil != err {
		return
	}
	value, err = cn.options.Serializer.Decode(data)

	return
}

func (cn *cacheNormal) Set(key string, value interface{}, options ...option) (err error) {
	newOptions := cn.options.options
	for _, option := range options {
		option.apply(&newOptions)
	}

	var data []byte
	if data, err = cn.options.Serializer.Encode(value); nil != err {
		return
	}
	err = cn.store.Set(key, data, options...)

	return
}

func (cn *cacheNormal) Delete(key string) (err error) {
	return cn.store.Delete(key)
}

func (cn *cacheNormal) Invalidate(options ...invalidateOption) (err error) {
	return cn.store.Invalidate(options...)
}

func (cn *cacheNormal) Type() string {
	return NormalType
}

func (cn *cacheNormal) Clear() (err error) {
	return cn.store.Clear()
}
