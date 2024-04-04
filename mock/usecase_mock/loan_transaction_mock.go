package usecasemock

import (
	"medioker-bank/model"
	"medioker-bank/model/dto"

	"github.com/stretchr/testify/mock"
)

type LoanTransactionMock struct {
	mock.Mock
}

func (l *LoanTransactionMock) FindAllLoanTransaction(page, limit int) ([]dto.LoanTransactionResponseDto, error) {
	args := l.Called(page, limit)
	return args.Get(0).([]dto.LoanTransactionResponseDto), args.Error(1)
}

func (l *LoanTransactionMock) FIndLoanTransactionByUserIdAndTrxId(userId, trxId string) (dto.LoanTransactionResponseDto, error) {
	args := l.Called(userId, trxId)
	return args.Get(0).(dto.LoanTransactionResponseDto), args.Error(1)
}

func (l *LoanTransactionMock) FindById(id string) (dto.LoanTransactionResponseDto, error) {
	args := l.Called(id)
	return args.Get(0).(dto.LoanTransactionResponseDto), args.Error(1)
}

func (l *LoanTransactionMock) FindByUserId(userId string) ([]dto.LoanTransactionResponseDto, error) {
	args := l.Called(userId)
	return args.Get(0).([]dto.LoanTransactionResponseDto), args.Error(1)
}

func (l *LoanTransactionMock) RegisterNewTransaction(payload dto.LoanTransactionRequestDto) (model.LoanTransaction, error) {
	args := l.Called(payload)
	return args.Get(0).(model.LoanTransaction), args.Error(1)
}
