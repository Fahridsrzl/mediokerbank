package usecase

import (
	"errors"
	"fmt"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	repository "medioker-bank/repository/transaction"
)

type TransferUseCase interface {
	CreateTransfer(transferDto dto.TransferDto) (model.TransferTransaction, error)
	GetTransferByTransferId(id string) (model.TransferTransaction, error)
	GetTransferBySenderId(senderId string) ([]dto.ResponseTransfer, error)
	GetAllTransfer() ([]dto.ResponseTransfer, error)
}

type transferUseCase struct {
	repo repository.TransferRepository
}

func (u *transferUseCase) CreateTransfer(transferDto dto.TransferDto) (model.TransferTransaction, error) {
	transfer := model.TransferTransaction{
		SenderID:   transferDto.SenderID,
		ReceiverID: transferDto.ReceiverID,
		Amount:     transferDto.Amount,
		Status:     transferDto.Status,
	}

	createdTransfer, err := u.repo.CreateTransfer(transfer)
	if err != nil {
		return model.TransferTransaction{}, errors.New("Error creating top up:" + err.Error())
	}
	return createdTransfer, nil
}

func (u *transferUseCase) GetTransferByTransferId(id string) (model.TransferTransaction, error) {
	transfer, err := u.repo.GetTransferByTransferId(id)
	if err != nil {
		return model.TransferTransaction{}, fmt.Errorf("there is no transfer of %s", id)
	}
	return transfer, nil
}

func (u *transferUseCase) GetTransferBySenderId(senderId string) ([]dto.ResponseTransfer, error) {
	transfers, err := u.repo.GetTransferBySenderId(senderId)
	if err != nil {
		return nil, fmt.Errorf("there is no transfer of %s", senderId)
	}
	return transfers, nil
}

func (u *transferUseCase) GetAllTransfer() ([]dto.ResponseTransfer, error) {
	transfers, err := u.repo.GetAllTransfer()
	if err != nil {
		return nil, fmt.Errorf("there is no transfer")
	}
	return transfers, nil
}

func NewTransferTransactionUseCase(repo repository.TransferRepository) TransferUseCase {
	return &transferUseCase{repo: repo}
}
