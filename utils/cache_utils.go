package utils

import (
	"github.com/go-redis/redis"
	"haotian_main/config"
	"log"
)

var RedisClient *redis.Client

func InitUcache(config config.RedisConfig) {
	// timeout类型的配置单位为毫秒。需要做转换
	redisConfig := redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	}
	redisCache := redis.NewClient(&redisConfig)
	_, err := redisCache.Ping().Result()

	if err != nil {
		log.Fatalf("init redis fail, err: %v", err)
	}
	RedisClient = redisCache
}
