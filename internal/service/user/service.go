package user

import (
	"context"
	"go-otg/internal/entity"
)

type resourceProvider interface {
	GetUserByIDFromDB(ctx context.Context, ID string) (User, error)
	AddUserToDB(ctx context.Context, user User) error
	GetUserByIDFromRedis(ctx context.Context, ID string) (entity.User, error)
	SetUserFromRedis(ctx context.Context, user User) error
}

type infraProvider interface {
	HashPassword(s string) (string, error)
	GenerateUniqueID() (string, error)
}

type Service struct {
	infra    infraProvider
	resource resourceProvider
}

// NewService new service by given resource resourceProvider.
//
// It returns pointer of Service when successful.
// Otherwise, nil pointer of Service will be returned.
func NewService(infra infraProvider, resource resourceProvider) *Service {
	return &Service{
		infra:    infra,
		resource: resource,
	}
}
