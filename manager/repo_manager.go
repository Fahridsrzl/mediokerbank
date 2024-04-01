package manager

import repository "medioker-bank/repository/master"

type RepoManager interface {
	LoanTransactionRepo() repository.LoanTransactionRepository
	LoanProductRepo() repository.LoanProductRepository
	UserRepo() repository.UserRepository
	LoanRepo() repository.LoanRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) LoanTransactionRepo() repository.LoanTransactionRepository{
	return repository.NewLoanTransactionRepository(r.infra.Conn())
}

func (r *repoManager) LoanProductRepo() repository.LoanProductRepository {
	return repository.NewLoanProductRepository(r.infra.Conn())
}

func (r *repoManager) LoanRepo() repository.LoanRepository{
	return repository.NewLoanRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
