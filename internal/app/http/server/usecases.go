package server

import "go-otg/internal/usecase/common"

type Usecases struct {
	Common *common.Usecase
}

// NewUsecases new usecases by given services pointer of Services.
//
// It returns pointer of Usecases when successful.
// Otherwise, nil pointer of Usecases will be returned.
func NewUsecases(services *Services) *Usecases {
	return &Usecases{
		Common: common.New(services.Common),
	}
}
