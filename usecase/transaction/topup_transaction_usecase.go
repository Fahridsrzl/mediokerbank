package usecase

import (
	"fmt"
	"medioker-bank/model"
	repository "medioker-bank/repository/transaction"
)

type TopupUseCase interface {
	CreateTopup(payload model.TopupTransaction) (model.TopupTransaction, error)
}

type topupUseCase struct {
	repo repository.TopupRepository
}

func (u *topupUseCase) CreateTopup(payload model.TopupTransaction) (model.TopupTransaction, error) {
	//validasi pake usecase dari user find by id, nanti datany pake buat validasi
	user, err := u.repo.CreateTopup(payload)
	if err != nil {
		return model.TopupTransaction{}, fmt.Errorf("failed to Top Up")
	}
	return user, nil
}

func NewTopupTransactionUseCase(repo repository.TopupRepository) TopupUseCase {
	return &topupUseCase{repo: repo}
}
