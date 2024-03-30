package repository

import (
	"database/sql"
	"medioker-bank/model"
	"medioker-bank/utils/common"
	"time"
)

type UserRepository interface {
	CreateUser(payload model.User) (model.User, error)
	GetUserById(id string) (model.User, error)
	//TODO : check func usernamepassword
	GetUserByUsernameOrEmailAndPassword(usernameOrEmail, password string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	//TODO : get all users func
	UpdateUser(id string, payload model.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) CreateUser(payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.CreateUser, payload.Username, payload.Email, payload.Password, payload.Role, payload.CreditScore, time.Now(), time.Now()).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreditScore, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetUserById(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.GetUserById, id).Scan(&user.Id, &user.Username, &user.Email, &user.Role, &user.CreditScore, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetUserByUsernameOrEmailAndPassword(usernameOrEmail, password string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.GetUserByUsernameOrEmailAndPassword, usernameOrEmail, password).Scan(&user.Id, &user.Username, &user.Email, &user.Role, &user.CreditScore, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetAllUsers() ([]model.User, error) {
    rows, err := u.db.Query(common.GetAllUsers)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []model.User
    for rows.Next() {
        var user model.User
        if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Role, &user.CreditScore, &user.CreatedAt, &user.UpdatedAt); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (u *userRepository) UpdateUser(id string, payload model.User) error {
	_, err := u.db.Exec(common.UpdateUser, id, payload.Username, payload.Email, payload.Password, payload.Role, payload.CreditScore, time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) DeleteUser(id string) error {
	_, err := u.db.Exec(common.DeleteUser, id)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
