package usecase

import (
	"medioker-bank/model"
	repository "medioker-bank/repository/master"
)

type LoanProductUseCase interface {
	FindLoanProductById(id string) (model.LoanProduct, error)
	FindAllLoanProduct() ([]model.LoanProduct, error)
	CreateLoanProduct(payload model.LoanProduct) (model.LoanProduct, error)
	UpdateLoanProduct(id string, payload model.LoanProduct) error
	DeleteLoanProduct(id string) (model.LoanProduct, error)
}

type loanProductUseCase struct {
	repo repository.LoanProductRepository
}

func (l *loanProductUseCase) FindLoanProductById(id string) (model.LoanProduct, error) {
	loanProduct, err := l.repo.GetById(id)
	if err != nil {
		return model.LoanProduct{}, err
	}
	return loanProduct, nil
}

func (l *loanProductUseCase) FindAllLoanProduct() ([]model.LoanProduct, error) {
	var loanProducts []model.LoanProduct
	var err error
	loanProducts, err = l.repo.GetAll()
	if err != nil {
		return []model.LoanProduct{}, err
	}
	return loanProducts, nil
}

func (l *loanProductUseCase) CreateLoanProduct(payload model.LoanProduct) (model.LoanProduct, error) {
	createdLoanProduct, err := l.repo.Create(payload)
	if err != nil {
		return model.LoanProduct{}, err
	}
	return createdLoanProduct, nil
}

func (l *loanProductUseCase) UpdateLoanProduct(id string, payload model.LoanProduct) error {
	err := l.repo.Update(id, payload)
	if err != nil {
		return err
	}
	return nil
}

func (l *loanProductUseCase) DeleteLoanProduct(id string) (model.LoanProduct, error) {
	loanProduct, err := l.repo.Delete(id)
	if err != nil {
		return model.LoanProduct{}, err
	}
	return loanProduct, nil
}

func NewLoanProductUseCase(repo repository.LoanProductRepository) LoanProductUseCase {
	return &loanProductUseCase{repo: repo}
}
