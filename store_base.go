package wukong

import (
	`fmt`
	`strings`
	`time`
)

type storeBase struct{}

func (sb *storeBase) setTags(key string, get storeGetFunc, set storeSetFunc, tags ...string) (err error) {
	for _, tag := range tags {
		var tagKey = fmt.Sprintf(tagFormatter, tag)
		var cacheKeys = sb.getCacheKeysForTag(tagKey, get)

		var inserted = false
		for _, cacheKey := range cacheKeys {
			if cacheKey == key {
				inserted = true
				break
			}
		}

		if !inserted {
			cacheKeys = append(cacheKeys, key)
		}
		err = set(tagKey, []byte(strings.Join(cacheKeys, ",")), 720*time.Hour)
	}

	return
}

func (sb *storeBase) getCacheKeysForTag(tag string, get storeGetFunc) (keys []string) {
	if result, err := get(tag); nil != err && "" != string(result) {
		keys = strings.Split(string(result), ",")
	}

	return
}

func (sb *storeBase) invalidate(get storeGetFunc, delete storeDeleteFunc, tags ...string) (err error) {
	for _, tag := range tags {
		var tagKey = fmt.Sprintf(tagFormatter, tag)
		var cacheKeys = sb.getCacheKeysForTag(tagKey, get)

		for _, cacheKey := range cacheKeys {
			err = delete(cacheKey)
		}
		err = delete(tagKey)
	}

	return
}
