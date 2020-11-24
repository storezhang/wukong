package wukong

import (
	`github.com/vmihailenco/msgpack`
)

type SerializerMsgpack struct{}

func (sm *SerializerMsgpack) Marshal(obj interface{}) ([]byte, error) {
	return msgpack.Marshal(obj)
}

func (sm *SerializerMsgpack) Unmarshal(data []byte, obj interface{}) error {
	return msgpack.Unmarshal(data, obj)
}
