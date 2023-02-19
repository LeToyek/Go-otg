package user

import (
	"context"
	"go-otg/internal/repository/db"
)

func (resource *Resource) GetUserByIDFromDB(ctx context.Context, ID string) (User, error) {
	result, err := resource.db.GetUserByID(ctx, ID)
	if err != nil {
		return User{}, err
	}

	return User(result), nil
}
func (resource *Resource) AddUserToDB(ctx context.Context, user User) error {
	err := resource.db.CreateUser(ctx, db.User(user))
	return err
}
