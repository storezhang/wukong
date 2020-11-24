package wukong

// Serializer 序列化接口，用于将结构体编码成二进制数组或者将二进制数据解码成结构体
type Serializer interface {
	Encoder
	Decoder
}
