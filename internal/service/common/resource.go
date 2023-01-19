package common

import "go-otg/internal/repository/db"

type dbRepoProvider interface {
	// GetCommonByID get common by id by given id.
	//
	// It returns db.Common, and nil error when successful.
	// Otherwise, empty db.Common, and error will be returned.
	GetCommonByID(id int64) (db.Common, error)
}

type Resource struct {
	db dbRepoProvider
}

// NewResource new resource by given db dbRepoProvider.
//
// It returns pointer of Resource when successful.
// Otherwise, nil pointer of Resource will be returned.
func NewResource(db dbRepoProvider) *Resource {
	return &Resource{
		db: db,
	}
}
