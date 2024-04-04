package repository

import (
	"database/sql"
	"errors"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	rawquery "medioker-bank/utils/raw_query"
	"time"
)

type TransferRepository interface {
	CreateTransfer(payload model.TransferTransaction) (model.TransferTransaction, error)
	GetTransferByTransferId(id string) (model.TransferTransaction, error)
	GetTransferBySenderId(senderId string) ([]dto.ResponseTransfer, error)
	GetAllTransfer(page, limit int) ([]dto.ResponseTransfer, error)
}

type transferRepository struct {
	db *sql.DB
}

func (t *transferRepository) CreateTransfer(payload model.TransferTransaction) (model.TransferTransaction, error) {
	var topup model.TransferTransaction
	tx, err := t.db.Begin()
	if err != nil {
		return model.TransferTransaction{}, err
	}
	err = tx.QueryRow(rawquery.CreateTransfer,
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
		tx.Rollback()
		return model.TransferTransaction{}, errors.New("create trx: " + err.Error())
	}
	_, err = tx.Exec(rawquery.UpdateSenderBalance, payload.Amount, payload.SenderID)
	if err != nil {
		tx.Rollback()
		return model.TransferTransaction{}, errors.New("sender balance: " + err.Error())
	}
	_, err = tx.Exec(rawquery.UpdateReceiverBalance, payload.Amount, payload.ReceiverID)
	if err != nil {
		tx.Rollback()
		return model.TransferTransaction{}, errors.New("receiver balance: " + err.Error())
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return model.TransferTransaction{}, errors.New("commit: " + err.Error())
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

func (u *transferRepository) GetAllTransfer(page, limit int) ([]dto.ResponseTransfer, error) {
	var transfers []dto.ResponseTransfer

	offset := (page - 1) * limit

	rows, err := u.db.Query(rawquery.GetAllTransfer, limit, offset)
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
