package wukong

import (
	`time`
)

type optionExpiration struct {
	expiration time.Duration
}

// WithExpiration 配置过期时间
func WithExpiration(expiration time.Duration) *optionExpiration {
	return &optionExpiration{expiration: expiration}
}

func (oe *optionExpiration) apply(options *options) {
	options.Expiration = oe.expiration
}

func (oe *optionExpiration) applyCache(options *cacheOptions) {
	options.Expiration = oe.expiration
}
