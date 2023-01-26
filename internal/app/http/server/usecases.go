package server

import (
	"go-otg/internal/usecase/common"
	"go-otg/internal/usecase/user"
)

type Usecases struct {
	Common *common.Usecase
	User   *user.Usecase
}

// NewUsecases new usecases by given services pointer of Services.
//
// It returns pointer of Usecases when successful.
// Otherwise, nil pointer of Usecases will be returned.
func NewUsecases(services *Services) *Usecases {
	return &Usecases{
		Common: common.New(services.Common),
		User:   user.New(services.User),
	}
}
