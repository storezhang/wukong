package wukong

import (
	`testing`
)

var (
	sg *SerializerGob
)

func TestGobEncode(t *testing.T) {
	if _, err := sg.Encode(&user{
		Username: "storezhang",
		Password: "test",
		Age:      34,
		Money:    2000000000000,
	}); nil != err {
		t.Fatalf("序列化出错：%s", err)
	}
}
