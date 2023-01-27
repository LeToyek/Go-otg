package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

func (r *Repository) GetUserByID(ctx context.Context, ID string) (User, error) {

	res, err := r.redis.Get(ctx, ID).Result()
	if err != nil {
		if err == redis.Nil {
			return User{}, nil
		}
		return User{}, err
	}

	var user User

	err = json.Unmarshal([]byte(res), &user)

	return user, err
}
func (r *Repository) SetRedisUser(ctx context.Context, user User) error {
	json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	ttl := 5 * time.Minute

	return r.redis.Set(ctx, user.ID, json, ttl).Err()
}
