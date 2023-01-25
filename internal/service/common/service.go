package common

import (
	"context"
	"go-otg/internal/entity"
)

type resourceProvider interface {
	// GetCommonByIDFromDB get common by id from db by given id.
	//
	// It returns entity.Common, and nil error when successful.
	// Otherwise, empty entity.Common, and error will be returned.
	GetCommonByIDFromDB(id int64) (entity.Common, error)
	GetCommonByIDFromRedis(ctx context.Context, id int64) (entity.Common, error)
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
