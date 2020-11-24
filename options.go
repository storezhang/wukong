package wukong

import (
	`time`
)

type (
	// Option 配置选项
	Option func(*options)

	options struct {
		// Expiration 过期时间
		Expiration time.Duration
		// Tags 标签列表
		Tags []string
	}
)

func defaultOptions() options {
	return options{}
}

// WithExpiration 配置过期时间
func WithExpiration(expiration time.Duration) Option {
	return func(options *options) {
		options.Expiration = expiration
	}
}

// WithTags 配置标签
func WithTags(tags ...string) Option {
	return func(options *options) {
		options.Tags = tags
	}
}
