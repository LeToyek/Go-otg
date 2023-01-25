package common

import "context"

// GetCommonMessageByID get common message by id by given id.
//
// It returns string, and nil error when successful.
// Otherwise, empty string, and error will be returned.
func (usecase *Usecase) GetCommonMessageByID(ctx context.Context, id int64) (string, error) {
	result, err := usecase.common.GetCommonByID(ctx, id)
	if err != nil {
		return "", err
	}

	return result.Message, nil
}
