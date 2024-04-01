package manager

import (
	master "medioker-bank/usecase/master"
	transaction "medioker-bank/usecase/transaction"
)

type UseCaseManager interface {
	LoanProductUseCase() master.LoanProductUseCase
	UserUseCase() master.UserUseCase
	TopupUseCase() transaction.TopupUseCase
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

func (u *useCaseManager) TopupUseCase() transaction.TopupUseCase {
	return transaction.NewTopupTransactionUseCase(u.repo.TopupRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
