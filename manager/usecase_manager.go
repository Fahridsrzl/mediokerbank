package manager

import (
	master "medioker-bank/usecase/master"
)

type UseCaseManager interface {
	LoanProductUseCase() master.LoanProductUseCase
	UserUseCase() master.UserUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u useCaseManager) LoanProductUseCase() master.LoanProductUseCase {
	return master.NewLoanProductUseCase(u.repo.LoanProductRepo())
}

func (u *useCaseManager) UserUseCase() master.UserUseCase {
	return master.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
