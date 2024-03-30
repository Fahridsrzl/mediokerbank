package manager

import repository "medioker-bank/repository/master"

type RepoManager interface {
	LoanProductRepo() repository.LoanProductRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) LoanProductRepo() repository.LoanProductRepository {
	return repository.NewLoanProductRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
