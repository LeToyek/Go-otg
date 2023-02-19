package db

import (
	"context"
	"go-otg/internal/repository/db/constants"
	"time"
)

func (repository *Repository) GetUserByID(ctx context.Context, ID string) (User, error) {
	ctxRepository, cancel := context.WithTimeout(ctx, time.Duration(4)*time.Second)
	defer cancel()

	var result User
	err := repository.db.GetContext(ctxRepository, &result, constants.GetUserById, ID)
	if err != nil {
		return User{}, err
	}

	return result, nil
}
func (repository *Repository) CreateUser(ctx context.Context, user User) error {
	ctxRepository, cancel := context.WithTimeout(ctx, time.Duration(4)*time.Second)
	defer cancel()

	_, err := repository.db.ExecContext(ctxRepository,
		constants.InsertOneUser,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.Address,
		user.Age.Int16,
	)
	if err != nil {
		return err
	}

	return err
}
