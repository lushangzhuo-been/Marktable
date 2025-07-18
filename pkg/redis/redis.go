package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"mark3/global"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.GVA_CONFIG.Redis.RedisHost, global.GVA_CONFIG.Redis.RedisPort),
		Password: global.GVA_CONFIG.Redis.RedisPassword,
		DB:       global.GVA_CONFIG.Redis.RedisDbName,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
