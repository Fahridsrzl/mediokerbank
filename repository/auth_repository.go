package repository

import (
	"database/sql"
	"medioker-bank/model"
	"medioker-bank/model/dto"
)

type AuthRepository interface {
	Create(payload dto.AuthVerifyDto) (int, error)
	FindByVCode(code int) (model.User, error)
	FindByUniqueAndPassword(payload dto.AuthLoginDto) (model.User, error)
}

type authRepository struct {
	db *sql.DB
}

func (a *authRepository) Create(payload dto.AuthVerifyDto) (int, error) {

}

func (a *authRepository) FindByVCode(code int) (model.StockProduct, error) {

}

func (a *authRepository) FindByUniqueAndPassword(payload dto.AuthLoginDto) (model.User, error) {

}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}
