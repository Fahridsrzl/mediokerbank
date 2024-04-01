package repository

import (
	"database/sql"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	rawquery "medioker-bank/utils/raw_query"
)

type AuthRepository interface {
	CreateQueue(payload dto.AuthVerifyDto) (string, error)
	CreateUser(payload dto.AuthVerifyDto) (model.User, error)
	FindByVCode(code int) (dto.AuthVerifyDto, error)
	DeleteQueue(code int) error
	FindUniqueUser(payload dto.AuthLoginDto) (model.User, error)
	FindUniqueAdmin(payload dto.AuthLoginDto) (dto.Admin, error)
}

type authRepository struct {
	db *sql.DB
}

func (a *authRepository) CreateQueue(payload dto.AuthVerifyDto) (string, error) {
	_, err := a.db.Exec(rawquery.RegisterQueue, payload.Username, payload.Email, payload.Password, payload.VCode)
	if err != nil {
		return "", err
	}
	return "register success, please check your email for verification code", nil
}

func (a *authRepository) CreateUser(payload dto.AuthVerifyDto) (model.User, error) {
	var user model.User
	err := a.db.QueryRow(rawquery.RegisterUser, payload.Username, payload.Email, payload.Password, "user", "unverified", 100, 0, 0).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.Status, &user.CreditScore, &user.Balance, &user.LoanActive, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *authRepository) FindByVCode(code int) (dto.AuthVerifyDto, error) {
	var data dto.AuthVerifyDto
	var mockId string
	err := a.db.QueryRow(rawquery.FindByVCode, code).Scan(
		&mockId, &data.Username, &data.Email, &data.Password, &data.VCode,
	)
	if err != nil {
		return dto.AuthVerifyDto{}, err
	}
	return data, nil
}

func (a *authRepository) DeleteQueue(code int) error {
	_, err := a.db.Exec(rawquery.DeleteQueueByVCode, code)
	if err != nil {
		return err
	}
	return nil
}

func (a *authRepository) FindUniqueUser(payload dto.AuthLoginDto) (model.User, error) {
	var user model.User
	var stmt *sql.Stmt
	var err error
	var args []any
	if payload.Email == "" {
		stmt, err = a.db.Prepare(rawquery.FindByUsernameUser)
		if err != nil {
			return model.User{}, err
		}
		args = append(args, payload.Username)
	}
	if payload.Username == "" {
		stmt, err = a.db.Prepare(rawquery.FindByEmailUser)
		if err != nil {
			return model.User{}, err
		}
		args = append(args, payload.Email)
	}
	err = stmt.QueryRow(args...).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.Status, &user.CreditScore, &user.Balance, &user.LoanActive, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (a *authRepository) FindUniqueAdmin(payload dto.AuthLoginDto) (dto.Admin, error) {
	var admin dto.Admin
	var stmt *sql.Stmt
	var err error
	var args []any
	if payload.Email == "" {
		stmt, err = a.db.Prepare(rawquery.FindByUsernameAndPasswordAdmin)
		if err != nil {
			return dto.Admin{}, err
		}
		args = append(args, payload.Username)
	}
	if payload.Username == "" {
		stmt, err = a.db.Prepare(rawquery.FindByEmailAndPasswordAdmin)
		if err != nil {
			return dto.Admin{}, err
		}
		args = append(args, payload.Email)
	}
	args = append(args, payload.Password)
	err = stmt.QueryRow(args...).Scan(
		&admin.Id, &admin.Username, &admin.Email, &admin.Password, &admin.Role, &admin.CreatedAt, &admin.UpdatedAt,
	)
	if err != nil {
		return dto.Admin{}, err
	}
	return admin, nil
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}
