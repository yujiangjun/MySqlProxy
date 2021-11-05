package config

import (
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

type RedisHelper struct {
	*redis.Client
}
var redisHelper *RedisHelper

var redisOnce sync.Once

func GetRedisHelper() *RedisHelper {
	return redisHelper
}

func NewRedisHelper() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client=rdb
		redisHelper=rdh
	})
	return rdb
}