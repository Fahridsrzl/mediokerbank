package usecase

import (
	"fmt"
	"medioker-bank/model"
	"medioker-bank/repository"
	"medioker-bank/utils/common"
)

type UserUseCase interface {
	RegisterNewUser(payload model.User) (model.User, error)
	FindById(id string) (model.User, error)
	FindByUsernameOrEmailAndPassword(usernameOrEmail, password string) (model.User, error)
	ShowAllUser() ([]model.User, error)
	ModifyUser(id string, payload model.User) error
	RemoveUser(id string) error
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) RegisterNewUser(payload model.User) (model.User, error) {
	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.User{}, err
	}
	payload.Password = newPassword
	return u.repo.CreateUser(payload)
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	user, err := u.repo.GetUserById(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (u *userUseCase) FindByUsernameOrEmailAndPassword(usernameOrEmail, password string) (model.User, error) {
	user, err := u.repo.GetUserByUsernameOrEmailAndPassword(usernameOrEmail, password)
	if err != nil {
		return model.User{}, fmt.Errorf("user with username or password %s and password %s not found", usernameOrEmail, password)
	}
	return user, nil
}

func (u *userUseCase) ShowAllUser() ([]model.User, error) {
	users, err := u.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUseCase) ModifyUser(id string, payload model.User) error {
	err := u.repo.UpdateUser(id, payload)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) RemoveUser(id string) error {
	err := u.repo.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %s", id)
	}

	return nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
