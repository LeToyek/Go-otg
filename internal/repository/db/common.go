package db

// GetCommonByID get common by id by given id.
//
// It returns Common, and nil error when successful.
// Otherwise, empty Common, and error will be returned.
func (repository *Repository) GetCommonByID(id int64) (Common, error) {
	return Common{
		Message: "Halo gan!",
	}, nil
}
