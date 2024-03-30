package manager

import "medioker-bank/usecase"

type UseCaseManager interface {
	StockProductuseCase() usecase.StockProductUseCase
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) StockProductuseCase() usecase.StockProductUseCase {
	return usecase.NewStockProductUseCase(u.repo.StockProductRepo())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
