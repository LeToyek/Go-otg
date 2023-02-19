package infrastructure

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h Hasher) HashPassword(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), 10)
}
