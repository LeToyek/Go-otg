package redis

import (
	"context"
	"encoding/json"
	"time"
)

func (r *Repository) SetSingleUser(user User) {
	json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	ttl := time.Duration(4) * time.Second
	r.redis.Set(context.Background(), user.ID, json, ttl)
}

func (r *Repository) GetUserByID(userID string) (User, error) {
	res, err := r.redis.Get(context.Background(), userID).Result()
	if err != nil {
		panic(err)
	}

	var user User

	err = json.Unmarshal([]byte(res), &user)
	return user, err
}
