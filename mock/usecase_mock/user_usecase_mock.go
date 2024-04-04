package usecasemock

import (
	"medioker-bank/model"
	"medioker-bank/model/dto"

	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

func (u *UserUseCaseMock) CreateProfileAndAddressThenUpdateUser(profileDto dto.ProfileCreateDto, addressDto dto.AddressCreateDto) (model.Profile, model.Address, error) {
	args := u.Called(profileDto, addressDto)
	return args.Get(0).(model.Profile), args.Get(1).(model.Address), args.Error(2)
}

func (u *UserUseCaseMock) FindByStatus(status string) ([]dto.ResponseStatus, error) {
	args := u.Called(status)
	return args.Get(0).([]dto.ResponseStatus), args.Error(1)
}

func (u *UserUseCaseMock) UpdateStatus(id string) error {
	args := u.Called(id)
	return args.Error(0)
}

func (u *UserUseCaseMock) GetUserByID(id string) (model.User, []model.Loan, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Get(1).([]model.Loan), args.Error(2)
}

func (u *UserUseCaseMock) RemoveUser(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserUseCaseMock) GetAllUser(page, limit int) ([]dto.UserDto, error) {
	args := u.Called(page, limit)
	return args.Get(0).([]dto.UserDto), args.Error(1)
}

func (u *UserUseCaseMock) UpdateUserBalance(id string, amount int) (int, error) {
	args := u.Called(id, amount)
	return args.Int(0), args.Error(1)
}
