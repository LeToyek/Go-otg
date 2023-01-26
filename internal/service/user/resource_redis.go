package user

import (
	"context"
	"go-otg/internal/entity"
)

func (resource *Resource) GetUserByIDFromRedis(ctx context.Context, ID string) (entity.User, error) {
	result, err := resource.redis.GetUserByID(ctx, ID)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User(result), nil
}
