package manager

import (
	master "medioker-bank/repository/master"
	other "medioker-bank/repository/other"
	transaction "medioker-bank/repository/transaction"
)

type RepoManager interface {
	LoanProductRepo() master.LoanProductRepository
	UserRepo() master.UserRepository
	AuthRepo() other.AuthRepository
	LoanRepo() master.LoanRepository
	InstallmentTransactionRepo() transaction.InstallmentTransactionRepository
	TopupRepo() transaction.TopupRepository
	TransferRepo() transaction.TransferRepository
	LoanTransactionRepo() transaction.LoanTransactionRepository
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

func (r *repoManager) InstallmentTransactionRepo() transaction.InstallmentTransactionRepository {
	return transaction.NewInstallmentTransactionRepository(r.infra.Conn())
}

func (r *repoManager) TopupRepo() transaction.TopupRepository {
	return transaction.NewTopupTransactionRepository(r.infra.Conn())
}

func (r *repoManager) TransferRepo() transaction.TransferRepository {
	return transaction.NewTransferTransactionRepository(r.infra.Conn())
}

func (r *repoManager) LoanTransactionRepo() transaction.LoanTransactionRepository {
	return transaction.NewLoanTransactionRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
