package wukong

import (
	`encoding/xml`
)

type SerializerXml struct{}

func (sx *SerializerXml) Marshal(obj interface{}) ([]byte, error) {
	return xml.Marshal(obj)
}

func (sx *SerializerXml) Unmarshal(data []byte, obj interface{}) error {
	return xml.Unmarshal(data, obj)
}
