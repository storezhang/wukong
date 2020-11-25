package wukong

type invalidateOption interface {
	applyInvalidate(options *invalidateOptions)
}
