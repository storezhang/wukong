package wukong

import (
	`time`
)

type options struct {
	// Expiration 过期时间
	Expiration time.Duration
	// Tags 标签列表
	Tags []string
}

func defaultOptions() options {
	return options{}
}
