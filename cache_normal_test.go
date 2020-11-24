package wukong

import (
	`log`
	`os`
	`testing`

	`github.com/alicebob/miniredis`
	`github.com/go-redis/redis/v8`
	`github.com/rs/xid`
)

var (
	cache Cache
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("启动Redis服务器出错：%s", err)
	}

	cache = New(NewRedis(redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})))

	code := m.Run()
	os.Exit(code)
}

func TestSet(t *testing.T) {
	var setTests = []struct {
		key      string
		expected interface{}
	}{
		{xid.New().String(), true},
		{xid.New().String(), false},
		{xid.New().String(), 1},
		{xid.New().String(), 1.0},
		{xid.New().String(), []int{1, 2, 3}},
		{xid.New().String(), []bool{true, false, true}},
		{xid.New().String(), []float64{1.0, 2.0}},
		{xid.New().String(), 13},
	}

	for _, st := range setTests {
		if err := cache.Set(st.key, &st.expected); nil != err {
			t.Fatalf("设置缓存出错：%s", err)
		}
	}
	for _, st := range setTests {
		obj, err := cache.Get(st.key)
		if nil != err {
			t.Fatalf("从缓存取出数据出错：%s", err)
		}
		if *obj != st.expected {
			t.Fatal("设置的缓存和从缓存取出来的值不匹配")
		}
	}
}
