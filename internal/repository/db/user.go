package db

import (
	"context"
	"time"
)

func (repository *Repository) GetUserByID(ctx context.Context, userID int64) (User, error) {
	ctxRepository, cancel := context.WithTimeout(ctx, time.Duration(4)*time.Second)
	defer cancel()

	var result User
	err := repository.db.GetContext(ctxRepository, &result, "select * from users where id = $1", userID)
	if err != nil {
		return User{}, err
	}

	return result, nil
}
