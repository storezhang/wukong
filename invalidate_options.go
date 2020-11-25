package wukong

type invalidateOptions struct {
	// Tags 标签列表
	Tags []string
}

func defaultInvalidateOptions() invalidateOptions {
	return invalidateOptions{}
}
