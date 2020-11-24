package wukong

type Decoder interface {
	// Unmarshal 将二进制数据解码成结构体
	Unmarshal(data []byte, obj interface{}) error
}
