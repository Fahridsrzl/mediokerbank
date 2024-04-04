package repomock

import (
	"medioker-bank/model"
	"medioker-bank/model/dto"

	"github.com/stretchr/testify/mock"
)

type AuthRepoMock struct {
	mock.Mock
}

func (a *AuthRepoMock) CreateQueue(payload dto.AuthVerifyDto) (string, error) {
	args := a.Called(payload)
	return args.String(0), args.Error(1)
}

func (a *AuthRepoMock) CreateUser(payload dto.AuthVerifyDto) (model.User, error) {
	args := a.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (a *AuthRepoMock) FindByVCode(code int) (dto.AuthVerifyDto, error) {
	args := a.Called(code)
	return args.Get(0).(dto.AuthVerifyDto), args.Error(1)
}

func (a *AuthRepoMock) DeleteQueue(code int) error {
	args := a.Called(code)
	return args.Error(0)
}

func (a *AuthRepoMock) FindUniqueUser(payload dto.AuthLoginDto) (model.User, error) {
	args := a.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (a *AuthRepoMock) FindUniqueAdmin(payload dto.AuthLoginDto) (dto.Admin, error) {
	args := a.Called(payload)
	return args.Get(0).(dto.Admin), args.Error(1)
}
