package wukong

import (
	`time`
)

var _ option = (*optionExpiration)(nil)

type optionExpiration struct {
	expiration time.Duration
}

// Expiration 配置过期时间
func Expiration(expiration time.Duration) *optionExpiration {
	return &optionExpiration{expiration: expiration}
}

func (oe *optionExpiration) apply(options *options) {
	options.expiration = oe.expiration
}
