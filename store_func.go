package wukong

import (
	`time`
)

type (
	storeGetFunc    func(key string) ([]byte, error)
	storeSetFunc    func(key string, data []byte, expiration time.Duration, tags ...string) error
	storeDeleteFunc func(key string) error
)
