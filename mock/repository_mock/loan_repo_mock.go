package repomock

import (
	"medioker-bank/model"

	"github.com/stretchr/testify/mock"
)

type LoanRepoMock struct {
	mock.Mock
}

func (l *LoanRepoMock) Create(payload model.Loan) (model.Loan, error) {
	args := l.Called(payload)
	return args.Get(0).(model.Loan), args.Error(1)
}

func (l *LoanRepoMock) FindByUserId(userId string) ([]model.Loan, error) {
	args := l.Called(userId)
	return args.Get(0).([]model.Loan), args.Error(1)
}

func (l *LoanRepoMock) UpdatePeriod(loanId string) error {
	args := l.Called(loanId)
	return args.Error(0)
}

func (l *LoanRepoMock) Delete(id string) error {
	args := l.Called(id)
	return args.Error(0)
}
