package wukong

import (
	`fmt`
	`testing`
)

var (
	sj *SerializerJson
)

func TestJsonMarshal(t *testing.T) {
	var (
		data []byte
		err  error
	)

	if data, err = sj.Encode(&user{
		Username: "storezhang",
		Password: "test",
		Age:      34,
		Money:    2000000000000,
	}); nil != err {
		t.Fatalf("序列化出错：%s", err)
	}

	var userPtr interface{}
	if userPtr, err = sj.Decode(data); nil != err {
		t.Fatalf("反序列化出错：%s", err)
	}
	fmt.Print(userPtr.(user))
}
