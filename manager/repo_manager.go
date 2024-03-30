package manager

import "medioker-bank/repository"

type RepoManager interface {
	StockProductRepo() repository.StockProductRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) StockProductRepo() repository.StockProductRepository {
	return repository.NewStockProductRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
