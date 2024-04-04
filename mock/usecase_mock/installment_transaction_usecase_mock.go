package usecasemock

import (
	"medioker-bank/model"
	"medioker-bank/model/dto"

	"github.com/stretchr/testify/mock"
)

type InstallmentTransactionUseCaseMock struct {
	mock.Mock
}

func (i *InstallmentTransactionUseCaseMock) CreateTrx(payload dto.InstallmentTransactionRequestDto) (dto.InstallmentTransactionResponseDto, error) {
	args := i.Called(payload)
	return args.Get(0).(dto.InstallmentTransactionResponseDto), args.Error(1)
}

func (i *InstallmentTransactionUseCaseMock) FindTrxById(id string) (model.InstallmentTransaction, error) {
	args := i.Called(id)
	return args.Get(0).(model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionUseCaseMock) FindTrxMany(payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error) {
	args := i.Called(payload)
	return args.Get(0).([]model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionUseCaseMock) FindTrxByUserId(userId string, payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error) {
	args := i.Called(payload)
	return args.Get(0).([]model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionUseCaseMock) FindTrxByUserIdAndTrxId(userId, trxId string) (model.InstallmentTransaction, error) {
	args := i.Called(userId, trxId)
	return args.Get(0).(model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionUseCaseMock) UpdateTrxById(id string) error {
	args := i.Called(id)
	return args.Error(0)
}
