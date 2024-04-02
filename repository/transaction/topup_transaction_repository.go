package repository

import (
	"database/sql"
	"errors"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	rawquery "medioker-bank/utils/raw_query"
	"time"
)

type TopupRepository interface {
	CreateTopup(payload model.TopupTransaction) (model.TopupTransaction, error)
	GetTopUpByTopupId(id string) (model.TopupTransaction, error)
	GetTopupByUserId(userId string) ([]dto.ResponseTopUp, error)
	GetAllTopUp() ([]dto.ResponseTopUp, error)
}

type topupRepository struct {
	db *sql.DB
}

func (t *topupRepository) CreateTopup(payload model.TopupTransaction) (model.TopupTransaction, error) {
	var topup model.TopupTransaction
	tx, err := t.db.Begin()
	if err != nil {
		return model.TopupTransaction{}, err
	}
	err = tx.QueryRow(rawquery.CreateTopup,
		time.Now(),
		payload.UserID,
		payload.Amount,
		payload.Status,
		time.Now(),
		time.Now(),
	).Scan(
		&topup.ID,
		&topup.TrxDate,
		&topup.UserID,
		&topup.Amount,
		&topup.Status,
		&topup.CreatedAt,
		&topup.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return model.TopupTransaction{}, errors.New("create trx: " + err.Error())
	}
	_, err = tx.Exec(rawquery.UpdateBalanceTopup, payload.Amount, payload.UserID)
	if err != nil {
		tx.Rollback()
		return model.TopupTransaction{}, errors.New("user balance: " + err.Error())
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return model.TopupTransaction{}, errors.New("commit: " + err.Error())
	}
	return topup, nil
}

func (u *topupRepository) GetTopUpByTopupId(id string) (model.TopupTransaction, error) {
	var topup model.TopupTransaction

	err := u.db.QueryRow(rawquery.GetTopUpByTopupId, id).Scan(
		&topup.ID,
		&topup.TrxDate,
		&topup.UserID,
		&topup.Amount,
		&topup.Status,
		&topup.CreatedAt,
		&topup.UpdatedAt,
	)
	if err != nil {
		return model.TopupTransaction{}, err
	}

	return topup, nil
}

func (u *topupRepository) GetTopupByUserId(userId string) ([]dto.ResponseTopUp, error) {
	var topups []dto.ResponseTopUp

	rows, err := u.db.Query(rawquery.GetTopupByUserId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var topup dto.ResponseTopUp
		err := rows.Scan(
			&topup.ID,
			&topup.TrxDate,
			&topup.UserID,
			&topup.Amount,
			&topup.Status,
			&topup.CreatedAt,
			&topup.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		topups = append(topups, topup)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return topups, nil
}

func (u *topupRepository) GetAllTopUp() ([]dto.ResponseTopUp, error) {
	var topups []dto.ResponseTopUp

	rows, err := u.db.Query(rawquery.GetAllTopUp)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var topup dto.ResponseTopUp
		err := rows.Scan(
			&topup.ID,
			&topup.TrxDate,
			&topup.UserID,
			&topup.Amount,
			&topup.Status,
			&topup.CreatedAt,
			&topup.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		topups = append(topups, topup)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return topups, nil
}

func NewTopupTransactionRepository(db *sql.DB) TopupRepository {
	return &topupRepository{db: db}
}
