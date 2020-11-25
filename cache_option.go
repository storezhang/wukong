package wukong

type cacheOption interface {
	applyCache(options *cacheOptions)
}
