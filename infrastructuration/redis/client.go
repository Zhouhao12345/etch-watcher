package redis

import "github.com/go-redis/redis"

func NewRedisClient(option *redis.Options) *redis.Client {
	return redis.NewClient(option)
}
