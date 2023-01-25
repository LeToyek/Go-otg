package redis

import (
	"context"
	"encoding/json"
	"time"
)

// GetCommonByID get common by id by given id.
//
// It returns Common, and nil error when successful.
// Otherwise, empty Common, and error will be returned.
func (repository *Repository) GetCommonByID(ctx context.Context, id int64) (Common, error) {
	ctxRepository, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	res, err := repository.redis.Get(ctxRepository, "test").Result()
	if err != nil {
		return Common{}, err
	}

	var result Common
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		return Common{}, err
	}

	return result, nil
}
