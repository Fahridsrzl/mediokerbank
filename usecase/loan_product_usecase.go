package usecase

import (
	"fmt"
	"medioker-bank/model"
	"medioker-bank/repository"
)

type LoanProductUseCase interface {
	FindById(id string) (model.LoanProduct, error)
	GetAll() ([]model.LoanProduct, error)
	Create(payload model.LoanProduct) (model.LoanProduct, error)
	Update(payload model.LoanProduct) error
	Delete(id string) error
}

type loanProductUseCase struct{
	repo repository.LoanProductRepository
}

func (l *loanProductUseCase) FindById(id string)(model.LoanProduct, error){
	loanProduct, err := l.repo.Get(id)
	if err != nil {
		return model.LoanProduct{}, fmt.Errorf("loan product with ID %s not found", id)
	}
	return loanProduct, nil
}

func (l *loanProductUseCase) GetAll() ([]model.LoanProduct, error) {
	loanProducts, err := l.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return loanProducts, nil
}

func (l *loanProductUseCase)Create(payload model.LoanProduct) (model.LoanProduct, error){
	createdLoanProduct, err := l.repo.Create(payload)
	if err != nil {
		return model.LoanProduct{}, err
	}
	return createdLoanProduct, nil
}

func (l *loanProductUseCase)Update(payload model.LoanProduct) error{
	err := l.repo.Update(payload)
	if err != nil {
		return err
	}
	return nil
}

func (l *loanProductUseCase) Delete(id string) error {
	err := l.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewLoanProductUseCase(repo repository.LoanProductRepository) LoanProductUseCase{
	return &loanProductUseCase{repo: repo}
}