package user

import (
	"context"
)

func (service *Service) GetUserByID(ctx context.Context, ID string) (User, error) {
	var user User
	result, err := service.resource.GetUserByIDFromRedis(ctx, ID)
	if err != nil {
		return User{}, err
	}
	user = User(result)
	if result.ID == "" {
		result, err := service.resource.GetUserByIDFromDB(ctx, ID)

		if err != nil || result.ID == "" {
			return User{}, err
		}
		user = result
		hashedPass, err := service.infra.HashPassword(user.Password)
		if err != nil {
			return User{}, err
		}

		user.Password = string(hashedPass)

		err = service.resource.SetUserFromRedis(context.Background(), user)
		if err != nil {
			return User{}, err
		}
	}

	return user, nil
}
func (s *Service) CreateUser(ctx context.Context, user User) error {
	id, err := s.infra.GenerateUniqueID()
	if err != nil {
		return err
	}

	hashedPass, err := s.infra.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.ID = id
	user.Password = hashedPass

	return s.resource.AddUserToDB(ctx, user)
}
