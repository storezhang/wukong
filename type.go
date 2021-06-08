package wukong

const (
	// TypeBigCache 基于内存的缓存
	TypeBigCache Type = "bigCache"
	// TypeMemcache 基于Memcache分布式缓存
	TypeMemcache Type = "memcache"
	// TypeRedis 基于Redis的分布式缓存
	TypeRedis Type = "redis"
)

// Type 类型
type Type string
