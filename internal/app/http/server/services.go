package server

import (
	"go-otg/internal/infrastructure"
	"go-otg/internal/service/common"
	"go-otg/internal/service/user"
)

type Services struct {
	Common *common.Service
	User   *user.Service
}

// NewServices new services by given resources pointer of Resources.
//
// It returns pointer of Services when successful.
// Otherwise, nil pointer of Services will be returned.
func NewServices(infra *infrastructure.Infra, resources *Resources) *Services {
	return &Services{
		Common: common.NewService(resources.Common),
		User:   user.NewService(infra, resources.User),
	}
}
