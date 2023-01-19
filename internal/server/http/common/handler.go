package common

type commonUCManager interface {
	// GetCommonMessageByID get common message by id by given id.
	//
	// It returns string, and nil error when successful.
	// Otherwise, empty string, and error will be returned.
	GetCommonMessageByID(id int64) (string, error)
}

type Handler struct {
	common commonUCManager
}

// New instantiates a new handler.
//
// It returns pointer of Handler when successful.
// Otherwise, nil pointer of Handler will be returned.
func New(common commonUCManager) *Handler {
	return &Handler{
		common: common,
	}
}
