package repomock

import (
	"medioker-bank/model"
	"medioker-bank/model/dto"

	"github.com/stretchr/testify/mock"
)

type InstallmentTransactionRepoMock struct {
	mock.Mock
}

func (i *InstallmentTransactionRepoMock) Create(payload model.InstallmentTransaction) (model.InstallmentTransaction, error) {
	args := i.Called(payload)
	return args.Get(0).(model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionRepoMock) FindById(id string) (model.InstallmentTransaction, error) {
	args := i.Called(id)
	return args.Get(0).(model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionRepoMock) FindAll() ([]model.InstallmentTransaction, error) {
	args := i.Called()
	return args.Get(0).([]model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionRepoMock) FindMany(payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error) {
	args := i.Called(payload)
	return args.Get(0).([]model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionRepoMock) FindByUserId(userId string, payload dto.InstallmentTransactionSearchDto) ([]model.InstallmentTransaction, error) {
	args := i.Called(userId, payload)
	return args.Get(0).([]model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionRepoMock) FindByUserIdAndTrxId(userId, trxId string) (model.InstallmentTransaction, error) {
	args := i.Called(userId, trxId)
	return args.Get(0).(model.InstallmentTransaction), args.Error(1)
}

func (i *InstallmentTransactionRepoMock) UpdateById(id string) (string, error) {
	args := i.Called(id)
	return args.String(0), args.Error(1)
}

func (i *InstallmentTransactionRepoMock) DeleteById(id string) error {
	args := i.Called(id)
	return args.Error(0)
}
