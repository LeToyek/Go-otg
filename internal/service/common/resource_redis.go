package common

import (
	"context"
	"go-otg/internal/entity"
)

func (resource *Resource) GetCommonByIDFromRedis(ctx context.Context, id int64) (entity.Common, error) {
	result, err := resource.redis.GetCommonByID(ctx, id)
	if err != nil {
		return entity.Common{}, err
	}

	return entity.Common(result), nil
}
