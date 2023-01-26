package server

import (
	"go-otg/internal/server/http/common"
	"go-otg/internal/server/http/user"
)

type Handlers struct {
	Common *common.Handler
	User   *user.Handler
}

// NewHandlers new handlers by given usecases pointer of Usecases.
//
// It returns pointer of Handlers when successful.
// Otherwise, nil pointer of Handlers will be returned.
func NewHandlers(usecases *Usecases) *Handlers {
	return &Handlers{
		Common: common.New(usecases.Common),
		User:   user.New(usecases.User),
	}
}
