package user

import (
	"context"
	"go-otg/internal/repository/db"
)

func (resource *Resource) GetUserByIDFromDB(ctx context.Context, ID string) (db.User, error) {
	result, err := resource.db.GetUserByID(ctx, ID)
	if err != nil {
		return db.User{}, err
	}

	return db.User(result), nil
}
