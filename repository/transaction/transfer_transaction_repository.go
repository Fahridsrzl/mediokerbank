package repository

import (
	"database/sql"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	rawquery "medioker-bank/utils/raw_query"
	"time"
)

type TransferRepository interface {
	CreateTransfer(payload model.TransferTransaction) (model.TransferTransaction, error)
	GetTransferByTransferId(id string) (model.TransferTransaction, error)
	GetTransferBySenderId(senderId string) ([]dto.ResponseTransfer, error)
	GetAllTransfer() ([]dto.ResponseTransfer, error)
}

type transferRepository struct {
	db *sql.DB
}

func (t *transferRepository) CreateTransfer(payload model.TransferTransaction) (model.TransferTransaction, error) {
	var topup model.TransferTransaction
	err := t.db.QueryRow(rawquery.CreateTransfer,
		time.Now(),
		payload.SenderID,
		payload.ReceiverID,
		payload.Amount,
		payload.Status,
		time.Now(),
		time.Now(),
	).Scan(
		&topup.ID,
		&topup.TrxDate,
		&topup.SenderID,
		&topup.ReceiverID,
		&topup.Amount,
		&topup.Status,
		&topup.CreatedAt,
		&topup.UpdatedAt,
	)
	if err != nil {
		return model.TransferTransaction{}, err
	}
	return topup, nil
}

func (u *transferRepository) GetTransferByTransferId(id string) (model.TransferTransaction, error) {
	var transfer model.TransferTransaction

	err := u.db.QueryRow(rawquery.GetTansferByTransferId, id).Scan(
		&transfer.ID,
		&transfer.TrxDate,
		&transfer.SenderID,
		&transfer.ReceiverID,
		&transfer.Amount,
		&transfer.Status,
		&transfer.CreatedAt,
		&transfer.UpdatedAt,
	)
	if err != nil {
		return model.TransferTransaction{}, err
	}

	return transfer, nil
}

func (u *transferRepository) GetTransferBySenderId(senderId string) ([]dto.ResponseTransfer, error) {
	var transfers []dto.ResponseTransfer

	rows, err := u.db.Query(rawquery.GetTransferBySenderId, senderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transfer dto.ResponseTransfer
		err := rows.Scan(
			&transfer.ID,
			&transfer.TrxDate,
			&transfer.SenderID,
			&transfer.ReceiverID,
			&transfer.Amount,
			&transfer.Status,
			&transfer.CreatedAt,
			&transfer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}

func (u *transferRepository) GetAllTransfer() ([]dto.ResponseTransfer, error) {
	var transfers []dto.ResponseTransfer

	rows, err := u.db.Query(rawquery.GetAllTransfer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transfer dto.ResponseTransfer
		err := rows.Scan(
			&transfer.ID,
			&transfer.TrxDate,
			&transfer.SenderID,
			&transfer.ReceiverID,
			&transfer.Amount,
			&transfer.Status,
			&transfer.CreatedAt,
			&transfer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}

func NewTransferTransactionRepository(db *sql.DB) TransferRepository {
	return &transferRepository{db: db}
}
