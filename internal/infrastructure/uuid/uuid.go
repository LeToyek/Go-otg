package uuid

import "github.com/gofrs/uuid"

type Service struct {
}

func (s Service) GenerateUUID() (uuid.UUID, error) {
	return uuid.NewV4()
}

func NewUUIDService() *Service {
	return &Service{}
}
