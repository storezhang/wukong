package wukong

// Encoder 编码器
type Encoder interface {
	// Encode 将结构体编码成二进制数组
	Encode(obj interface{}) ([]byte, error)
}
