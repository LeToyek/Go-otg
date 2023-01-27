package user

import (
	"context"
	"go-otg/internal/entity"
	"go-otg/internal/repository/db"
)

type resourceProvider interface {
	GetUserByIDFromDB(ctx context.Context, ID string) (db.User, error)
	GetUserByIDFromRedis(ctx context.Context, ID string) (entity.User, error)
	SetUserFromRedis(ctx context.Context, user User) error
}

type Service struct {
	resource resourceProvider
}

// NewService new service by given resource resourceProvider.
//
// It returns pointer of Service when successful.
// Otherwise, nil pointer of Service will be returned.
func NewService(resource resourceProvider) *Service {
	return &Service{
		resource: resource,
	}
}
