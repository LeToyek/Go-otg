package server

import "go-otg/internal/service/common"

type Services struct {
	Common *common.Service
}

// NewServices new services by given resources pointer of Resources.
//
// It returns pointer of Services when successful.
// Otherwise, nil pointer of Services will be returned.
func NewServices(resources *Resources) *Services {
	return &Services{
		Common: common.NewService(resources.Common),
	}
}
