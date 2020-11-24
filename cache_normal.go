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
func New(store Store, options ...CacheOption) Cache {
	appliedOptions := defaultCacheOptions()
	for _, apply := range options {
		apply(&appliedOptions)
	}

	return &cacheNormal{
		options: appliedOptions,
		store:   store,
	}
}

func (cn *cacheNormal) Get(key string) (obj interface{}, err error) {
	var data []byte

	if data, err = cn.store.Get(key); nil != err {
		return
	}
	obj = new(interface{})
	err = cn.options.Serializer.Unmarshal(data, obj)

	return
}

func (cn *cacheNormal) Set(key string, obj interface{}, options ...Option) (err error) {
	newOptions := cn.options.options
	for _, apply := range options {
		apply(&newOptions)
	}

	var data []byte
	if data, err = cn.options.Serializer.Marshal(obj); nil != err {
		return
	}
	err = cn.store.Set(key, data, options...)

	return
}

func (cn *cacheNormal) Delete(key string) (err error) {
	return cn.store.Delete(key)
}

func (cn *cacheNormal) Invalidate(options ...InvalidateOption) (err error) {
	return cn.store.Invalidate(options...)
}

func (cn *cacheNormal) Type() string {
	return NormalType
}

func (cn *cacheNormal) Clear() (err error) {
	return cn.store.Clear()
}
