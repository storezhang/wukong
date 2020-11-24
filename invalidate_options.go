package wukong

type (
	// InvalidateOption 配置选项
	InvalidateOption func(*invalidateOptions)

	invalidateOptions struct {
		// Tags 标签列表
		Tags []string
	}
)

func defaultInvalidateOptions() invalidateOptions {
	return invalidateOptions{}
}

// WithTags 配置标签
func WithInvalidateTags(tags ...string) InvalidateOption {
	return func(options *invalidateOptions) {
		options.Tags = tags
	}
}
