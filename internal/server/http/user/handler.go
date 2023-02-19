package user

import (
	"context"
	"go-otg/internal/usecase/user"
)

type userUCManager interface {
	GetUserDataByID(ctx context.Context, ID string) (user.User, error)
	LoginUser(ctx context.Context, user user.LoginUserModel) error
	RegisterUser(ctx context.Context, regisUser user.RegisterUserModel) error
}

type Handler struct {
	user userUCManager
}

func New(user userUCManager) *Handler {
	return &Handler{
		user: user,
	}
}
