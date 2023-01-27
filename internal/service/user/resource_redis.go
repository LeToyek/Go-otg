package user

import (
	"context"
	"go-otg/internal/entity"
	"go-otg/internal/repository/redis"
)

func (resource *Resource) GetUserByIDFromRedis(ctx context.Context, ID string) (entity.User, error) {
	result, err := resource.redis.GetUserByID(ctx, ID)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User(result), nil
}
func (resource *Resource) SetUserFromRedis(ctx context.Context, user User) error {
	return resource.redis.SetRedisUser(ctx, redis.User(user))
}
