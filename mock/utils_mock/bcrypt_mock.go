package utilsmock

import "github.com/stretchr/testify/mock"

type BcryptUtilsMock struct {
	mock.Mock
}

func (b *BcryptUtilsMock) GeneratePasswordHash(password string) (string, error) {
	args := b.Called(password)
	return args.String(0), args.Error(1)
}

func (b *BcryptUtilsMock) ComparePasswordHash(hashPassword string, password string) error {
	args := b.Called(hashPassword, password)
	return args.Error(0)
}
