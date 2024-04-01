package usecase

import (
	"errors"
	"fmt"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	repository "medioker-bank/repository/transaction"
)

type TopupUseCase interface {
	CreateTopup(topupDto dto.TopupDto) (model.TopupTransaction, error)
	GetTopUpByTopupId(id string) (model.TopupTransaction, error)
	GetTopupByUserId(userId string) ([]dto.ResponseTopUp, error)
	GetAllTopUp() ([]dto.ResponseTopUp, error)
}

type topupUseCase struct {
	repo repository.TopupRepository
}

func (u *topupUseCase) CreateTopup(topupDto dto.TopupDto) (model.TopupTransaction, error) {
	topup := model.TopupTransaction{
		UserID: topupDto.UserID,
		Amount: topupDto.Amount,
		Status: topupDto.Status,
	}

	createdtopup, err := u.repo.CreateTopup(topup)
	if err != nil {
		return model.TopupTransaction{}, errors.New("Error creating top up:" + err.Error())
	}
	return createdtopup, nil
}

func (u *topupUseCase) GetTopUpByTopupId(id string) (model.TopupTransaction, error) {
	topup, err := u.repo.GetTopUpByTopupId(id)
	if err != nil {
		return model.TopupTransaction{}, fmt.Errorf("there is no transfer of %s", id)
	}
	return topup, nil
}

func (u *topupUseCase) GetTopupByUserId(userId string) ([]dto.ResponseTopUp, error) {
	topups, err := u.repo.GetTopupByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("there is no topup of %s", userId)
	}
	return topups, nil
}

func (u *topupUseCase) GetAllTopUp() ([]dto.ResponseTopUp, error) {
	topups, err := u.repo.GetAllTopUp()
	if err != nil {
		return nil, fmt.Errorf("there is no topup")
	}
	return topups, nil
}

func NewTopupTransactionUseCase(repo repository.TopupRepository) TopupUseCase {
	return &topupUseCase{repo: repo}
}
