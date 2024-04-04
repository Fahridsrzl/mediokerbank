package repomock

import (
	"medioker-bank/model"
	"medioker-bank/model/dto"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) UpdateUser(payload model.User) error {
	args := u.Called(payload)
	return args.Error(0)
}

func (u *UserRepoMock) CreateProfile(payload model.Profile) (model.Profile, error) {
	args := u.Called(payload)
	return args.Get(0).(model.Profile), args.Error(1)
}

func (u *UserRepoMock) CreateAddress(payload model.Address, profileID model.Profile) (model.Address, error) {
	args := u.Called(payload)
	return args.Get(0).(model.Address), args.Error(1)
}

func (u *UserRepoMock) GetUserByStatus(status string) ([]dto.ResponseStatus, error) {
	args := u.Called(status)
	return args.Get(0).([]dto.ResponseStatus), args.Error(1)
}

func (u *UserRepoMock) GetUserByID(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) DeleteUser(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) GetAllUsers(page, limit int) ([]dto.UserDto, error) {
	args := u.Called(page, limit)
	return args.Get(0).([]dto.UserDto), args.Error(1)
}

func (u *UserRepoMock) UpdateBalance(id string, amount int) (int, error) {
	args := u.Called(id, amount)
	return args.Int(0), args.Error(1)
}
