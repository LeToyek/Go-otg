package user

import (
	"context"
)

func (u *Usecase) GetUserDataByID(ctx context.Context, ID string) (User, error) {
	res, err := u.user.GetUserByID(ctx, ID)
	if err != nil {
		return User{}, err
	}
	return User(res), err
}
