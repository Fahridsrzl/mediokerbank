package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	rawquery "medioker-bank/utils/raw_query"
	"time"
)

type UserRepository interface {
	UpdateUser(payload model.User) (model.User, error)
	CreateProfile(payload model.Profile) (model.Profile, error)
	CreateAddress(payload model.Address, profileID model.Profile) (model.Address, error)
	GetUserByStatus(status string) ([]dto.ResponseStatus, error)
	GetUserByID(id string) (model.User, error)
	DeleteUser(id string) (model.User, error)
	GetAllUsers() ([]dto.UserDto, error)
}

type userRepository struct {
	db *sql.DB
}

// pake tx
// return id
// lanjut tx berikutnya buat profile
// query

func (u *userRepository) CreateProfile(payload model.Profile) (model.Profile, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return model.Profile{}, err
	}

	var profile model.Profile
	err = tx.QueryRow(rawquery.CreateProfile,
		payload.FirstName,
		payload.LastName,
		payload.Citizenship,
		payload.NationalID,
		payload.BirthPlace,
		payload.BirthDate,
		payload.Gender,
		payload.MaritalStatus,
		payload.Occupation,
		payload.MonthlyIncome,
		payload.PhoneNumber,
		payload.UrgentPhoneNumber,
		payload.Photo,
		payload.IDCard,
		payload.SalarySlip,
		payload.UserID,
		time.Now(),
		time.Now(),
	).Scan(
		&profile.ID,
		&profile.FirstName,
		&profile.LastName,
		&profile.Citizenship,
		&profile.NationalID,
		&profile.BirthPlace,
		&profile.BirthDate,
		&profile.Gender,
		&profile.MaritalStatus,
		&profile.Occupation,
		&profile.MonthlyIncome,
		&profile.PhoneNumber,
		&profile.UrgentPhoneNumber,
		&profile.Photo,
		&profile.IDCard,
		&profile.SalarySlip,
		&profile.UserID,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return model.Profile{}, errors.New("create profile: " + err.Error())
	}

	// Commit transaction if everything is successful
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return model.Profile{}, err
	}

	return profile, nil
}

func (u *userRepository) CreateAddress(payload model.Address, profileID model.Profile) (model.Address, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return model.Address{}, err
	}

	var address model.Address
	err = tx.QueryRow(rawquery.CreateAddress,
		payload.AddressLine,
		payload.City,
		payload.Province,
		payload.PostalCode,
		payload.Country,
		profileID.ID,
		time.Now(),
		time.Now(),
	).Scan(
		&address.ID,
		&address.AddressLine,
		&address.City,
		&address.Province,
		&address.PostalCode,
		&address.Country,
		&address.ProfileID,
		&address.CreatedAt,
		&address.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return model.Address{}, err
	}

	// Commit transaction if everything is successful
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return model.Address{}, errors.New("create address: " + err.Error())
	}

	return address, nil
}

func (u *userRepository) UpdateUser(payload model.User) (model.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	_, err = tx.Exec(rawquery.UpdateUser,
		payload.Status, // Update status field
		payload.ID,     // Where condition based on user ID
	)
	if err != nil {
		tx.Rollback()
		return model.User{}, errors.New("update user :" + err.Error())
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetUserByStatus(status string) ([]dto.ResponseStatus, error) {
	var users []dto.ResponseStatus

	rows, err := u.db.Query(rawquery.GetUserByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user dto.ResponseStatus
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Role,
			&user.Status,
			&user.CreditScore,
			&user.Balance,
			&user.LoanActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetUserByID(id string) (model.User, error) {
	var user model.User
	var profile model.Profile
	var address model.Address

	err := u.db.QueryRow(rawquery.GetUserById, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.Status,
		&user.CreditScore,
		&user.Balance,
		&user.LoanActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&profile.ID,
		&profile.FirstName,
		&profile.LastName,
		&profile.Citizenship,
		&profile.NationalID,
		&profile.BirthPlace,
		&profile.BirthDate,
		&profile.Gender,
		&profile.MaritalStatus,
		&profile.Occupation,
		&profile.MonthlyIncome,
		&profile.PhoneNumber,
		&profile.UrgentPhoneNumber,
		&profile.Photo,
		&profile.IDCard,
		&profile.SalarySlip,
		&profile.UserID,
		&profile.CreatedAt,
		&profile.UpdatedAt,
		&address.ID,
		&address.AddressLine,
		&address.City,
		&address.Province,
		&address.PostalCode,
		&address.Country,
		&address.ProfileID,
		&address.CreatedAt,
		&address.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}

	user.Profile = profile
	user.Profile.Address = address

	fmt.Println(user)
	return user, nil
}

func (u *userRepository) DeleteUser(id string) (model.User, error) {
	var user model.User
	_, err := u.db.Exec(rawquery.DeleteUser, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetAllUsers() ([]dto.UserDto, error) {
	var users []dto.UserDto
	rows, err := u.db.Query(rawquery.GetAllUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user dto.UserDto
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Role,
			&user.Status,
			&user.CreditScore,
			&user.Balance,
			&user.LoanActive,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
