package manager

import "medioker-bank/usecase"

type UseCaseManager interface {
	StockProductuseCase() usecase.StockProductUseCase
	LoanProductUseCase() usecase.LoanProductUseCase

}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) StockProductuseCase() usecase.StockProductUseCase {
	return usecase.NewStockProductUseCase(u.repo.StockProductRepo())
}

func (u useCaseManager) LoanProductUseCase() usecase.LoanProductUseCase{
	return usecase.NewLoanProductUseCase(u.repo.LoanProductRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager{
	return &useCaseManager{repo:repo}
}

