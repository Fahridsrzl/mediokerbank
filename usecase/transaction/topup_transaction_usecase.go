package usecase

import (
	"errors"
	"fmt"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	repository "medioker-bank/repository/transaction"
	usecase "medioker-bank/usecase/master"
)

type TopupUseCase interface {
	CreateTopup(topupDto dto.TopupDto) (model.TopupTransaction, error)
	GetTopUpByTopupId(id string) (model.TopupTransaction, error)
	GetTopupByUserId(userId string) ([]dto.ResponseTopUp, error)
	GetAllTopUp(page, limit int) ([]dto.ResponseTopUp, error)
}

type topupUseCase struct {
	repo   repository.TopupRepository
	userUc usecase.UserUseCase
}

func (u *topupUseCase) CreateTopup(topupDto dto.TopupDto) (model.TopupTransaction, error) {
	user, _, err := u.userUc.GetUserByID(topupDto.UserID)
	if err != nil {
		return model.TopupTransaction{}, errors.New("find user: " + err.Error())
	}
	if user.Status != "verified" {
		return model.TopupTransaction{}, errors.New("user unverified")
	}
	topup := model.TopupTransaction{
		UserID: topupDto.UserID,
		Amount: topupDto.Amount,
		Status: "success",
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

func (u *topupUseCase) GetAllTopUp(page, limit int) ([]dto.ResponseTopUp, error) {
	topups, err := u.repo.GetAllTopUp(page, limit)
	if err != nil {
		return nil, fmt.Errorf("there is no topup")
	}
	return topups, nil
}

func NewTopupTransactionUseCase(repo repository.TopupRepository, userUc usecase.UserUseCase) TopupUseCase {
	return &topupUseCase{repo: repo, userUc: userUc}
}
