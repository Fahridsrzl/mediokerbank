package usecasemock

import (
	"medioker-bank/model"
	"medioker-bank/model/dto"

	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (a *AuthUseCaseMock) RegisterUser(payload dto.AuthRegisterDto) (string, error) {
	args := a.Called(payload)
	return args.String(0), args.Error(1)
}

func (a *AuthUseCaseMock) VerifyUser(code int) (model.User, error) {
	args := a.Called(code)
	return args.Get(0).(model.User), args.Error(1)
}

func (a *AuthUseCaseMock) LoginUser(payload dto.AuthLoginDto) (dto.AuthResponseDto, error) {
	args := a.Called(payload)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}

func (a *AuthUseCaseMock) LoginAdmin(payload dto.AuthLoginDto) (dto.AuthResponseDto, error) {
	args := a.Called(payload)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}
