package redis

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

func (r *Repository) SetSingleUser(user User) {
	json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	ttl := time.Duration(4) * time.Second
	res := r.redis.Set(context.Background(), user.ID, json, ttl)
	if res.Err() != nil {
		panic(err.Error())
	}
}

func (r *Repository) GetUserByID(ctx context.Context, ID string) (User, error) {
	r.SetSingleUser(User{
		ID:       "tyq1",
		Username: "toyeqq",
		Email:    "toyeq@gmail.com",
		Password: "toyeqcleancode",
		Address:  "Jalan Cendana 13",
		Age: sql.NullInt16{
			Int16: 12,
			Valid: true,
		},
	})
	res, err := r.redis.Get(ctx, ID).Result()
	if err != nil {
		return User{}, err
	}

	var user User

	err = json.Unmarshal([]byte(res), &user)
	return user, err
}
