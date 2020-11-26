package wukong

import (
	`log`
	`os`
	`testing`
	`time`

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
	Age      int    `json:"age"`
	Money    int64  `json:"money"`
}

func (u user) compare(ou user) bool {
	return u.Username == ou.Username && u.Password == ou.Password && u.Age == ou.Age && u.Money == ou.Money
}

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("启动Redis服务器出错：%s", err)
	}

	gob := &SerializerGob{}
	gob.RegisterGobConcreteType(user{})
	gob.RegisterGobConcreteType([]user{})

	cache = New(NewRedis(redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})), WithSerializer(gob), WithExpiration(time.Second))

	code := m.Run()
	os.Exit(code)
}

func TestSet(t *testing.T) {
	var setTests = []struct {
		key      string
		expected interface{}
	}{
		{xid.New().String(), user{
			Username: "storezhang",
			Password: "test",
			Age:      34,
			Money:    2000000000000,
		}},
		{xid.New().String(), []user{
			{
				Username: "storezhang",
				Password: "test",
				Age:      34,
				Money:    2000000000000,
			},
			{
				Username: "taoismzhang",
				Password: "test1",
				Age:      35,
				Money:    2000000000001,
			},
		}},
	}

	for _, st := range setTests {
		if err := cache.Set(st.key, &st.expected); nil != err {
			t.Fatalf("设置缓存出错：%s", err)
		}
	}
	for _, st := range setTests {
		switch st.expected.(type) {
		case user:
			if err := cache.Set(st.key, st.expected.(user)); nil != err {
				t.Fatalf("设置缓存出错：%s", err)
			}

			if cachedUser, err := cache.Get(st.key); nil != err {
				t.Fatalf("从缓存取出数据出错：%s", err)
			} else if !st.expected.(user).compare(*cachedUser.(*user)) {
				t.Fatalf("设置的缓存和从缓存取出来的值不匹配，缓存值：%v，期望值：%v", cachedUser, st.expected)
			}
		case []user:
			if err := cache.Set(st.key, st.expected.([]user)); nil != err {
				t.Fatalf("设置缓存出错：%s", err)
			}

			if cachedUsers, err := cache.Get(st.key); nil != err {
				t.Fatalf("从缓存取出数据出错：%s", err)
			} else if len(st.expected.([]user)) != len(cachedUsers.([]user)) || !st.expected.([]user)[0].compare(cachedUsers.([]user)[0]) {
				t.Fatalf("设置的缓存和从缓存取出来的值不匹配，缓存值：%v，期望值：%v", cachedUsers, st.expected)
			}
		}
	}
}
