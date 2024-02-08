package global

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	Redis *redis.Client
	once  sync.Once
)

func init() {
	once.Do(func() {
		// 初始化Redis客户端
		rdb := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // Redis服务器地址
			Password: "",               // 如果没有设置密码，留空
			DB:       0,                // 默认数据库编号
		})
		Redis = rdb
	})
}
