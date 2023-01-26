package server

import (
	// golang package
	"go-otg/internal/repository/db"
	"go-otg/internal/repository/redis"
	"go-otg/internal/service/common"
	"go-otg/internal/service/user"
)

type Resources struct {
	Common *common.Resource
	User   *user.Resource
}

// NewResources new resources by given db pointer of db.Repository.
//
// It returns pointer of Resources when successful.
// Otherwise, nil pointer of Resources will be returned.
func NewResources(db *db.Repository, redis *redis.Repository) *Resources {
	return &Resources{
		Common: common.NewResource(db, redis),
		User:   user.NewResource(db, redis),
	}
}
