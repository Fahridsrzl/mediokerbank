package usecase

import (
	"fmt"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	repository "medioker-bank/repository/master"
)

type UserUseCase interface {
	CreateProfileAndAddressThenUpdateUser(profileDto dto.ProfileCreateDto, addressDto dto.AddressCreateDto) (model.Profile, model.Address, error)
	FindByStatus(status string) ([]dto.ResponseStatus, error)
	UpdateStatus(id string) error
	GetUserByID(id string) (model.User, []model.Loan, error)
	RemoveUser(id string) (model.User, error)
	GetAllUser(page, limit int) ([]dto.UserDto, error)
	UpdateUserBalance(id string, amount int) (int, error)
}

type userUseCase struct {
	repo     repository.UserRepository
	loanrepo repository.LoanRepository
}

func (u *userUseCase) CreateProfileAndAddressThenUpdateUser(profileDto dto.ProfileCreateDto, addressDto dto.AddressCreateDto) (model.Profile, model.Address, error) {
	// Create Profile
	profile := model.Profile{
		FirstName:         profileDto.FirstName,
		LastName:          profileDto.LastName,
		Citizenship:       profileDto.Citizenship,
		NationalID:        profileDto.NationalID,
		BirthPlace:        profileDto.BirthPlace,
		BirthDate:         profileDto.BirthDate,
		Gender:            profileDto.Gender,
		MaritalStatus:     profileDto.MaritalStatus,
		Occupation:        profileDto.Occupation,
		MonthlyIncome:     profileDto.MonthlyIncome,
		PhoneNumber:       profileDto.PhoneNumber,
		UrgentPhoneNumber: profileDto.UrgentPhoneNumber,
		Photo:             profileDto.Photo,
		IDCard:            profileDto.IDCard,
		SalarySlip:        profileDto.SalarySlip,
		UserID:            profileDto.UserID,
	}
	createdProfile, err := u.repo.CreateProfile(profile)
	if err != nil {
		return model.Profile{}, model.Address{}, err
	}

	// create address
	address := model.Address{
		AddressLine: addressDto.AddressLine,
		City:        addressDto.City,
		Province:    addressDto.Province,
		PostalCode:  addressDto.PostalCode,
		Country:     addressDto.Country,
	}
	createdAddress, err := u.repo.CreateAddress(address, createdProfile)
	if err != nil {
		return createdProfile, model.Address{}, err
	}

	// update User
	user := model.User{
		Status: "pending",
		ID:     profileDto.UserID,
	}
	err = u.repo.UpdateUser(user)
	if err != nil {
		return model.Profile{}, model.Address{}, err
	}

	return createdProfile, createdAddress, nil
}

func (u *userUseCase) FindByStatus(status string) ([]dto.ResponseStatus, error) {
	user, err := u.repo.GetUserByStatus(status)
	if err != nil {
		return []dto.ResponseStatus{}, fmt.Errorf("no user with status %s", status)
	}

	return user, nil
}

func (u *userUseCase) UpdateStatus(id string) error {
	user := model.User{
		Status: "verified",
		ID:     id,
	}
	err := u.repo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) GetUserByID(id string) (model.User, []model.Loan, error) {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return model.User{}, nil, err
	}

	loans, err := u.loanrepo.FindByUserId(id)
	if err != nil {
		return model.User{}, nil, err
	}

	return user, loans, nil
}

func (u *userUseCase) RemoveUser(id string) (model.User, error) {
	user, err := u.repo.DeleteUser(id)
	if err != nil {
		return model.User{}, fmt.Errorf("no user with id %s", id)
	}
	return user, nil
}

func (u *userUseCase) GetAllUser(page, limit int) ([]dto.UserDto, error) {
	users, err := u.repo.GetAllUsers(page, limit)
	if err != nil {
		return nil, fmt.Errorf("there is no user")
	}
	return users, nil
}

func (u *userUseCase) UpdateUserBalance(id string, amount int) (int, error) {
	newBalance, err := u.repo.UpdateBalance(id, amount)
	if err != nil {
		return 0, err
	}
	return newBalance, nil
}

func NewUserUseCase(repo repository.UserRepository, loanrepo repository.LoanRepository) UserUseCase {
	return &userUseCase{repo: repo, loanrepo: loanrepo}
}
