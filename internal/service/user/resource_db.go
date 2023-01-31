package user

import (
	"context"
	"fmt"
)

func (resource *Resource) GetUserByIDFromDB(ctx context.Context, ID string) (User, error) {
	result, err := resource.db.GetUserByID(ctx, ID)
	fmt.Print("WORK HEREEE" + result.Email)
	if err != nil {
		return User{}, err
	}

	return User(result), nil
}
