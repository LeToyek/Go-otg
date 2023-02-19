package user

import (
	"context"
	"go-otg/internal/repository/db"
	"go-otg/internal/repository/redis"
)

type dbRepoProvider interface {
	GetUserByID(ctx context.Context, ID string) (db.User, error)
	CreateUser(ctx context.Context, user db.User) error
}
type redisRepoProvider interface {
	GetUserByID(ctx context.Context, ID string) (redis.User, error)
	SetRedisUser(ctx context.Context, user redis.User) error
}

type Resource struct {
	db    dbRepoProvider
	redis redisRepoProvider
}

func NewResource(db dbRepoProvider, redis redisRepoProvider) *Resource {
	return &Resource{
		db:    db,
		redis: redis,
	}
}
