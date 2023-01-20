package redis

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}
