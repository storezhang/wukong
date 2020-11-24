package wukong

import (
	`encoding/json`
)

type SerializerJson struct{}

func (sj *SerializerJson) Marshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (sj *SerializerJson) Unmarshal(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}
