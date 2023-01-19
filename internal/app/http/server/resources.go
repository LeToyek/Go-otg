package server

import (
	// golang package
	"go-otg/internal/repository/db"
	"go-otg/internal/service/common"
)

type Resources struct {
	Common *common.Resource
}

// NewResources new resources by given db pointer of db.Repository.
//
// It returns pointer of Resources when successful.
// Otherwise, nil pointer of Resources will be returned.
func NewResources(db *db.Repository) *Resources {
	return &Resources{
		Common: common.NewResource(db),
	}
}
