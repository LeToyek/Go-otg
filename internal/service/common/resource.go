package common

import (
	"context"
	"go-otg/internal/repository/db"
	"go-otg/internal/repository/redis"
)

type dbRepoProvider interface {
	// GetCommonByID get common by id by given id.
	//
	// It returns db.Common, and nil error when successful.
	// Otherwise, empty db.Common, and error will be returned.
	GetCommonByID(id int64) (db.Common, error)
}
type redisRepoProvider interface {
	GetCommonByID(ctx context.Context, id int64) (redis.Common, error)
}

type Resource struct {
	db    dbRepoProvider
	redis redisRepoProvider
}

// NewResource new resource by given db dbRepoProvider.
//
// It returns pointer of Resource when successful.
// Otherwise, nil pointer of Resource will be returned.
func NewResource(db dbRepoProvider, redis redisRepoProvider) *Resource {
	return &Resource{
		db:    db,
		redis: redis,
	}
}
