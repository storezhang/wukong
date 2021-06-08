package wukong

import (
	`bytes`
	`reflect`
	`unsafe`

	`github.com/vmihailenco/msgpack`
)

var _ Serializer = (*serializerMsgpack)(nil)

type serializerMsgpack struct{}

func (sm *serializerMsgpack) Encode(obj interface{}) ([]byte, error) {
	return msgpack.Marshal(obj)
}

func (sm *serializerMsgpack) Decode(data []byte) (ptr interface{}, err error) {
	buffer := bytes.NewBuffer(data)
	decoder := msgpack.NewDecoder(buffer)

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
