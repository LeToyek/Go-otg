package user

import (
	"context"
	"database/sql"
	"errors"
	"go-otg/internal/service/user"
)

func (u *Usecase) GetUserDataByID(ctx context.Context, ID string) (User, error) {
	res, err := u.user.GetUserByID(ctx, ID)
	if err != nil {
		return User{}, err
	}
	return User(res), err
}
func (u *Usecase) LoginUser(ctx context.Context, user LoginUserModel) error {
	return errors.New("not supported yet")
}

func (u *Usecase) RegisterUser(ctx context.Context, regisUser RegisterUserModel) error {
	var newUser user.User = user.User{
		Username: regisUser.Username,
		Email:    regisUser.Email,
		Password: regisUser.Password,
		Address:  regisUser.Address,
		Age: sql.NullInt16{
			Int16: int16(regisUser.Age),
			Valid: true,
		},
	}

	return u.user.CreateUser(ctx, newUser)
}
