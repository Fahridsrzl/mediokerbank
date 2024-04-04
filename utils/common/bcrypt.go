package common

import "golang.org/x/crypto/bcrypt"

type BcryptService interface {
	GeneratePasswordHash(password string) (string, error)
	ComparePasswordHash(hashPassword string, password string) error
}

type bcryptService struct {
}

func (b *bcryptService) GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (b *bcryptService) ComparePasswordHash(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func NewBcryptService() BcryptService {
	return &bcryptService{}
}
