package manager

import "medioker-bank/usecase"

type UseCaseManager interface {
	StockProductuseCase() usecase.StockProductUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) StockProductuseCase() usecase.StockProductUseCase {
	return usecase.NewStockProductUseCase(u.repo.StockProductRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
