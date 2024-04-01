package repository

import (
	"database/sql"
	"medioker-bank/model"
)

type TransferRepository interface {
	CreateTransfer(payload model.TransferTransaction) (model.TransferTransaction, error)
}

type transferRepository struct {
	db *sql.DB
}