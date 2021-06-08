package wukong

var _ option = (*optionTags)(nil)

type optionTags struct {
	tags []string
}

// Tags 配置标签
func Tags(tags ...string) *optionTags {
	return &optionTags{tags: tags}
}

func (ot *optionTags) apply(options *options) {
	options.tags = ot.tags
}
