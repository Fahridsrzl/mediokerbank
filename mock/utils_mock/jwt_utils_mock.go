package utilsmock

import (
	"medioker-bank/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
)

type JwtUtilsMock struct {
	mock.Mock
}

func (j *JwtUtilsMock) GenerateToken(payload model.User) (string, error) {
	args := j.Called(payload)
	return args.String(0), args.Error(1)
}

func (j *JwtUtilsMock) GenerateRefreshToken(payload model.User) (string, error) {
	args := j.Called(payload)
	return args.String(0), args.Error(1)
}

func (j *JwtUtilsMock) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	args := j.Called(tokenString)
	return args.Get(0).(jwt.MapClaims), args.Error(1)
}

func (j *JwtUtilsMock) RefreshToken(oldTokenString string) (string, error) {
	args := j.Called(oldTokenString)
	return args.String(0), args.Error(1)
}
