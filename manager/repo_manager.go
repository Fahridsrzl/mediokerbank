package manager

import "medioker-bank/repository"

type RepoManager interface {
	StockProductRepo() repository.StockProductRepository
	LoanProductRepo() repository.LoanProductRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) StockProductRepo() repository.StockProductRepository {
	return repository.NewStockProductRepository(r.infra.Conn())
}

func (r *repoManager) LoanProductRepo() repository.LoanProductRepository {
	return repository.NewLoanProductRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager{
	return &repoManager{infra: infra}
}
