package common

import "go-otg/internal/service/common"

type commonServiceProvider interface {
	// GetCommonByID get common by id by given id.
	//
	// It returns common.Common, and nil error when successful.
	// Otherwise, empty common.Common, and error will be returned.
	GetCommonByID(id int64) (common.Common, error)
}

type Usecase struct {
	common commonServiceProvider
}

// New new by given common commonServiceProvider.
//
// It returns pointer of Usecase when successful.
// Otherwise, nil pointer of Usecase will be returned.
func New(common commonServiceProvider) *Usecase {
	return &Usecase{
		common: common,
	}
}
