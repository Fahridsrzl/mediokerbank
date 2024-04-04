package usecase

import (
	"errors"
	"fmt"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	repository "medioker-bank/repository/transaction"
	usecase "medioker-bank/usecase/master"
)

type TransferUseCase interface {
	CreateTransfer(transferDto dto.TransferDto) (model.TransferTransaction, error)
	GetTransferByTransferId(id string) (model.TransferTransaction, error)
	GetTransferBySenderId(senderId string) ([]dto.ResponseTransfer, error)
	GetAllTransfer(page, limit int) ([]dto.ResponseTransfer, error)
}

type transferUseCase struct {
	repo   repository.TransferRepository
	userUc usecase.UserUseCase
}

func (u *transferUseCase) CreateTransfer(transferDto dto.TransferDto) (model.TransferTransaction, error) {
	sender, _, err := u.userUc.GetUserByID(transferDto.SenderID)
	if err != nil {
		return model.TransferTransaction{}, errors.New("find sender: " + err.Error())
	}
	receiver, _, err := u.userUc.GetUserByID(transferDto.ReceiverID)
	if err != nil {
		return model.TransferTransaction{}, errors.New("find receiver: " + err.Error())
	}
	if sender.Status != "verified" || receiver.Status != "verified" {
		return model.TransferTransaction{}, errors.New("sender or receiver unverified")
	}
	if sender.Balance < transferDto.Amount {
		return model.TransferTransaction{}, errors.New("too low sender balance")
	}
	transfer := model.TransferTransaction{
		SenderID:   transferDto.SenderID,
		ReceiverID: transferDto.ReceiverID,
		Amount:     transferDto.Amount,
		Status:     "success",
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

func (u *transferUseCase) GetAllTransfer(page, limit int) ([]dto.ResponseTransfer, error) {
	transfers, err := u.repo.GetAllTransfer(page, limit)
	if err != nil {
		return nil, fmt.Errorf("there is no transfer")
	}
	return transfers, nil
}

func NewTransferTransactionUseCase(repo repository.TransferRepository, userUc usecase.UserUseCase) TransferUseCase {
	return &transferUseCase{repo: repo, userUc: userUc}
}
