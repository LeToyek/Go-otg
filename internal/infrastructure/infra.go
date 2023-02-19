package infrastructure

import "github.com/gofrs/uuid"

type HashManager interface {
	HashPassword(s string) ([]byte, error)
}

type UUIDManager interface {
	GenerateUUID() (uuid.UUID, error)
}

type Infra struct {
	hash HashManager
	uuid UUIDManager
}

func New(hash HashManager, uuid UUIDManager) *Infra {
	return &Infra{
		hash: hash,
		uuid: uuid,
	}
}

func (infra *Infra) HashPassword(s string) (string, error) {
	result, err := infra.hash.HashPassword(s)
	return string(result), err
}
func (infra *Infra) GenerateUniqueID() (string, error) {
	res, err := infra.uuid.GenerateUUID()
	return res.String(), err
}
