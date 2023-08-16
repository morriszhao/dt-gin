package helper

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {
	ctx := context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := RedisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatal("redis连接失败：", err.Error())
	}
}

func RedisGet(key string) string {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	val := RedisClient.Get(timeoutCtx, key).Val()
	return val
}

func RedisSet(key string, val string, expire time.Duration) bool {

	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	err := RedisClient.Set(timeoutCtx, key, val, expire).Err()
	if err != nil {
		return false
	}

	return true
}
