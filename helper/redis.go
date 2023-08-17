package helper

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

type redisConfig struct {
	host   string
	passwd string
	port   int
	db     int
}

var RedisClient *redis.Client

func InitRedis() {
	redisConfig := initConfig()
	dsn := redisConfig.host + ":" + strconv.Itoa(redisConfig.port)
	ctx := context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: redisConfig.passwd,
		DB:       redisConfig.db,
	})

	err := RedisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatal("redis连接失败：", err.Error())
	}
}

func initConfig() redisConfig {
	subViper := ViperConfig.Sub("db.redis")
	config := redisConfig{
		host:   subViper.GetString("host"),
		port:   subViper.GetInt("port"),
		passwd: subViper.GetString("passwd"),
		db:     subViper.GetInt("db"),
	}

	if config.host == "" {
		log.Fatal("请配置redis链接信息")
	}
	if config.port == 0 {
		config.port = 6379
	}
	return config
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
