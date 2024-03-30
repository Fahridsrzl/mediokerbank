package manager

import "medioker-bank/repository"

type RepoManager interface {
	LoanProductRepo() repository.LoanProductRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) LoanProductRepo() repository.LoanProductRepository {
	return repository.NewLoanProductRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager{
	return &repoManager{infra: infra}
}
