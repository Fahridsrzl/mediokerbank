package manager

import (
	master "medioker-bank/repository/master"
	other "medioker-bank/repository/other"
)

type RepoManager interface {
	LoanProductRepo() master.LoanProductRepository
	UserRepo() master.UserRepository
	AuthRepo() other.AuthRepository
	LoanRepo() master.LoanRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) LoanProductRepo() master.LoanProductRepository {
	return master.NewLoanProductRepository(r.infra.Conn())
}

func (r *repoManager) UserRepo() master.UserRepository {
	return master.NewUserRepository(r.infra.Conn())
}

func (r *repoManager) AuthRepo() other.AuthRepository {
	return other.NewAuthRepository(r.infra.Conn())
}

func (r *repoManager) LoanRepo() master.LoanRepository {
	return master.NewLoanRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
