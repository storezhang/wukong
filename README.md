# wukong（悟空）
悟空，Golang版本的缓存框架，存在的唯一目的就是加速系统性能


# 为什么要叫悟空
《西游记》里记载，孙悟空一个筋斗云十万八千里，和系统缓存加载异曲同工

## 功能
- 支持本地内存缓存BigCache
- 支持Redis
- 支持Memcache
- 统一的缓存接口
- 方便集成自己的缓存

## 为什么要写这个缓存
最近一直在寻找一个统一的Golang版本的缓存框架，无奈于Golang的生态确实不如Java，各自为政，许久寻得一个框架基本满足要求https://github.com/eko/gocache
但是其接口设计使用使用起来特别麻烦，每次存放数据，都得传Options参数，还不能省略，十分讨厌，所以在其基础之上，增加更易使用的方法形成此框架
