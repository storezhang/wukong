package wukong

import (
	`bytes`
	`encoding/json`
	`reflect`
	`unsafe`
)

var _ Serializer = (*serializerJson)(nil)

type serializerJson struct{}

func (sj *serializerJson) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (sj *serializerJson) Decode(data []byte) (ptr interface{}, err error) {
	buffer := bytes.NewBuffer(data)
	decoder := json.NewDecoder(buffer)

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
