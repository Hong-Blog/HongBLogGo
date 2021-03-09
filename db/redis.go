package db

import (
	"github.com/go-redis/redis/v8"
	"os"
)

var RedisClient *redis.Client

func getRedisUrl() string {
	url := os.Getenv("REDIS_URL")
	if url == "" {
		url = "redis://127.0.0.1:6379/1"
	}
	return url
}

func init() {
	opt, err := redis.ParseURL(getRedisUrl())
	if err != nil {
		panic(err)
	}

	RedisClient = redis.NewClient(opt)
}
