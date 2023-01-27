package user

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func (service *Service) GetUserByID(ctx context.Context, ID string) (User, error) {
	var user User
	result, err := service.resource.GetUserByIDFromRedis(ctx, ID)
	if err != nil {
		return User{}, err
	}
	user = User(result)
	if result.ID == "" {
		// resultFromDB, newErr := service.resource.GetUserByIDFromDB(ctx, ID)
		newUser := User{
			ID:       "tyq1",
			Username: "toyeqq",
			Email:    "toyeq@gmail.com",
			Password: "toyeqcleancode",
			Address:  "Jalan Cendana 13",
			Age: sql.NullInt16{
				Int16: 12,
				Valid: true,
			},
		}
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(result.Password), 6)
		if err != nil {
			return User{}, err
		}
		newUser.Password = string(hashedPass)
		err = service.resource.SetUserFromRedis(context.Background(), newUser)
		if err != nil {
			return User{}, err
		}
		user = newUser
	}
	//example hashing from service

	return user, nil
}
