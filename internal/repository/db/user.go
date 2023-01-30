package db

import (
	"context"
	"go-otg/internal/repository/db/constants"
	"strconv"
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
func (repository *Repository) CreateUser(ctx context.Context, user User) (string, error) {
	ctxRepository, cancel := context.WithTimeout(ctx, time.Duration(4)*time.Second)
	defer cancel()

	res, err := repository.db.ExecContext(ctxRepository,
		`INSERT INTO user("id","username","email","password","address","age") 
		VALUES ($1,$2,$3,$4,$5,$6)`,
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.Address,
		user.Age.Int16,
	)
	if err != nil {
		return "", err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(rowCount, 10), err
}
