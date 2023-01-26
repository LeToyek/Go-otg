package user

import (
	"context"
	"go-otg/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

func (service *Service) GetUserByID(ctx context.Context, ID string) (User, error) {
	result, err := service.resource.GetUserByIDFromRedis(ctx, ID)
	if err != nil {
		resultFromDB, newErr := service.resource.GetUserByIDFromDB(ctx, ID)
		result = entity.User(resultFromDB)
		err = newErr
	}
	if err != nil {
		return User{}, err
	}

	//example hashing from service

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(result.Password), 14)
	if err != nil {
		return User{}, err
	}
	result.Password = string(hashedPass)

	return User(result), nil
}
