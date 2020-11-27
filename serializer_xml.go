package wukong

import (
	`bytes`
	`encoding/xml`
	`reflect`
	`unsafe`
)

var _ Serializer = (*SerializerXml)(nil)

type SerializerXml struct{}

func (sx *SerializerXml) Encode(obj interface{}) ([]byte, error) {
	return xml.Marshal(obj)
}

func (sx *SerializerXml) Decode(data []byte) (ptr interface{}, err error) {
	buffer := bytes.NewBuffer(data)
	decoder := xml.NewDecoder(buffer)

	var obj interface{}
	if err = decoder.Decode(&obj); nil != err {
		return
	}

	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Struct {
		var objPtr interface{} = &obj
		interfaceData := reflect.ValueOf(objPtr).Elem().InterfaceData()
		sp := reflect.NewAt(value.Type(), unsafe.Pointer(interfaceData[1])).Interface()
		ptr = sp
	} else {
		ptr = obj
	}

	return
}
