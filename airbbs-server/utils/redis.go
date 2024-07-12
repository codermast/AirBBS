package utils

import (
	"codermast.com/airbbs/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var RedisClient *redis.Client

// RedisInit 初始化
func RedisInit() {
	redisConfig := config.GetRedisConfig()

	redisSocket := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	// 连接到 Redis 服务器
	client := redis.NewClient(&redis.Options{
		Addr:     redisSocket,          // Redis 服务器地址
		Password: redisConfig.Password, // Redis 访问密码，没有密码时留空
		DB:       redisConfig.DB,       // 使用的数据库，默认是 0
	})

	RedisClient = client
}

// Set 设置一个键值对
func Set(key string, value string, expireTime time.Duration) error {
	err := RedisClient.Set(context.Background(), key, value, expireTime).Err()
	return err
}

// Get 获取指定 Key 的值
func Get(key string) (string, error) {
	return RedisClient.Get(context.Background(), key).Result()
}

// Del 删除指定键值对
func Del(key string) error {
	return RedisClient.Del(context.Background(), key).Err()
}
