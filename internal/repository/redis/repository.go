package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	redis RedisProvider
}

type RedisProvider interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}

func NewRedis(redis RedisProvider) *Repository {
	return &Repository{
		redis: redis,
	}
}
