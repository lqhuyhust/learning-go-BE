package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	RedisAccessTokenClient *redis.Client
	RedisRateLimitClient   *redis.Client
	RedisPostClient        *redis.Client
	RedisUserClient        *redis.Client
)

func ConnectRedis() {
	// use db0 to store access token
	RedisAccessTokenClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if _, err := RedisAccessTokenClient.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	// use db1 to store rate limit info
	RedisRateLimitClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       1,
	})

	// use db2 to store posts cache
	RedisPostClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       2,
	})

	// use db3 to store users cache
	RedisUserClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       3,
	})

	if _, err := RedisRateLimitClient.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
}
