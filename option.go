package wukong

type option interface {
	apply(options *options)
}
