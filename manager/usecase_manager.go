package manager

import usecase "medioker-bank/usecase/master"

type UseCaseManager interface {
	LoanProductUseCase() usecase.LoanProductUseCase
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u useCaseManager) LoanProductUseCase() usecase.LoanProductUseCase {
	return usecase.NewLoanProductUseCase(u.repo.LoanProductRepo())
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
