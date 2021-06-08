package wukong

import (
	`time`
)

type options struct {
	// 过期时间
	expiration time.Duration
	// 标签列表
	tags []string
	// 序列化
	serializer Serializer
}

func defaultOptions() *options {
	return &options{
		expiration: 30 * time.Minute,
		serializer: &serializerGob{},
	}
}
