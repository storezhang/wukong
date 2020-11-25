package wukong

type optionTags struct {
	tags []string
}

// WithTags 配置标签
func WithTags(tags ...string) *optionTags {
	return &optionTags{tags: tags}
}

func (ot *optionTags) apply(options *options) {
	options.Tags = ot.tags
}

func (ot *optionTags) applyInvalidate(options *invalidateOptions) {
	options.Tags = ot.tags
}
