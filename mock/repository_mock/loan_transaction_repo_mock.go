package repomock

import (
	"medioker-bank/model"

	"github.com/stretchr/testify/mock"
)

type LoanTransactionRepoMock struct {
	mock.Mock
}


func (l *LoanTransactionRepoMock)GetAll() ([]model.LoanTransaction, error){
	args := l.Called()
	return  args.Get(0).([]model.LoanTransaction), args.Error(1)
}

func (l *LoanTransactionRepoMock) GetByID(id string) (model.LoanTransaction, error){
	args := l.Called(id)
	return args.Get(0).(model.LoanTransaction), args.Error(1)
}

func (l *LoanTransactionRepoMock) GetByUserID(userId string) (model.LoanTransaction, error){
	args := l.Called(userId)
	return args.Get(0).(model.LoanTransaction), args.Error(1)
}

func (l *LoanTransactionRepoMock) GetByUserIdAndTrxId(userId, trxId string) ([]model.LoanTransaction, error){
	args := l.Called(userId, trxId)
	return  args.Get(0).([]model.LoanTransaction), args.Error(1)
}

func (l *LoanTransactionRepoMock) Create(payload model.LoanTransaction) (model.LoanTransaction, error){
	args := l.Called(payload)
	return args.Get(0).(model.LoanTransaction), args.Error(1)
}
