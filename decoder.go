package wukong

type Decoder interface {
	// Decode 将二进制数据解码成结构体
	Decode(data []byte) (interface{}, error)
}
