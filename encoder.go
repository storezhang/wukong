package wukong

// Encoder 编码器
type Encoder interface {
	// Marshal 将结构体编码成二进制数组
	Marshal(obj interface{}) ([]byte, error)
}
