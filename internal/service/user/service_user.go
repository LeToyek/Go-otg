package user

import (
	// golang package
	"context"
	"go-otg/internal/entity"

	// external package
	"github.com/tokopedia/tdk/go/log"
	"github.com/tokopedia/tdk/go/tracer/v2"
	"golang.org/x/crypto/bcrypt"
)

// GetUserByID get user by id by given ID.
//
// It returns User, and nil error when successful.
// Otherwise, empty User, and error will be returned.
func (service *Service) GetUserByID(ctx context.Context, ID string) (User, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, "ServiceGetUserByID")
	defer span.Finish()
	result, err := service.resource.GetUserByIDFromRedis(ctx, ID)
	if err != nil {
		log.StdTrace(ctx, map[string]interface {
		}{
			"ID": ID,
		}, err, "service.resource.GetUserByIDFromRedis() got error - GetUserByID")
		resultFromDB, newErr := service.resource.GetUserByIDFromDB(ctx, ID)
		result = entity.User(resultFromDB)
		err = newErr
	}
	if err != nil {
		return User{}, err
	}

	//example hashing from service

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(result.Password), 10)
	if err != nil {
		log.StdTrace(ctx, map[string]interface {
		}{
			"result.Password": result.Password,
		}, err, "bcrypt.GenerateFromPassword() got error - GetUserByID")
		return User{}, err
	}
	result.Password = string(hashedPass)

	return User(result), nil
}
