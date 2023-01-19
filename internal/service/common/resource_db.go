package common

import (
	"go-otg/internal/entity"
)

// GetCommonByIDFromDB get common by id from db by given id.
//
// It returns entity.Common, and nil error when successful.
// Otherwise, empty entity.Common, and error will be returned.
func (resource *Resource) GetCommonByIDFromDB(id int64) (entity.Common, error) {
	result, err := resource.db.GetCommonByID(21)
	if err != nil {
		return entity.Common{}, err
	}

	return entity.Common(result), nil
}
