package user

import (
	"context"
	"go-otg/internal/service/user"
)

type userServiceProvider interface {
	GetUserByID(ctx context.Context, ID string) (user.User, error)
}

type Usecase struct {
	user userServiceProvider
}

func New(user userServiceProvider) *Usecase {
	return &Usecase{
		user: user,
	}
}
