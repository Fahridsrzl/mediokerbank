package usecasemock

import (
	"medioker-bank/model"

	"github.com/stretchr/testify/mock"
)

type LoanProductMock struct {
	mock.Mock
}

func (l *LoanProductMock) FindLoanProductById(id string) (model.LoanProduct, error) {
	args := l.Called(id)
	return args.Get(0).(model.LoanProduct), args.Error(1)
}

func (l *LoanProductMock) FindAllLoanProduct() ([]model.LoanProduct, error) {
	args := l.Called()
	return args.Get(0).([]model.LoanProduct), args.Error(1)
}

func (l *LoanProductMock) CreateLoanProduct(payload model.LoanProduct) (model.LoanProduct, error) {
	args := l.Called(payload)
	return args.Get(0).(model.LoanProduct), args.Error(1)
}

func (l *LoanProductMock) UpdateLoanProduct(id string, payload model.LoanProduct) error {
	args := l.Called(id)
	return args.Error(1)
}

func (l *LoanProductMock) DeleteLoanProduct(id string) (model.LoanProduct, error) {
	args := l.Called(id)
	return args.Get(0).(model.LoanProduct), args.Error(1)
}
