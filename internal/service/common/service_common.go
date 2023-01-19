package common

// GetCommonByID get common by id by given id.
//
// It returns Common, and nil error when successful.
// Otherwise, empty Common, and error will be returned.
func (service *Service) GetCommonByID(id int64) (Common, error) {
	result, err := service.resource.GetCommonByIDFromDB(id)

	if err != nil {
		return Common{}, err
	}

	return Common(result), nil
}
