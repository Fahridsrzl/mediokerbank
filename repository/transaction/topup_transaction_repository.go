package repository

import (
	"database/sql"
	"medioker-bank/model"
	rawquery "medioker-bank/utils/raw_query"
	"time"
)

type TopupRepository interface {
	CreateTopup(payload model.TopupTransaction) (model.TopupTransaction, error)
}

type topupRepository struct {
	db *sql.DB
}

func (t *topupRepository) CreateTopup(payload model.TopupTransaction) (model.TopupTransaction, error) {
	var topup model.TopupTransaction
	err := t.db.QueryRow(rawquery.CreateTopup,
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
		return model.TopupTransaction{}, err
	}

	return topup, nil

}

func NewTopupTransactionRepository(db *sql.DB) TopupRepository {
	return &topupRepository{db: db}
}
