package reposiotrytesting

import (
	"database/sql"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	repository "medioker-bank/repository/master"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = repository.NewUserRepository(suite.mockDB)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(InstallmentTransactionRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestUpdateUser_Success() {
	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectExec("UPDATE users").WillReturnResult(
		sqlmock.NewResult(1, 1),
	)
	suite.mockSql.ExpectCommit()

	payloadMock := model.User{}

	err := suite.repo.UpdateUser(payloadMock)
	assert.Nil(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestCreateProfile_Success() {
	table := sqlmock.NewRows([]string{
		"&profile.ID",
		"&profile.FirstName",
		"&profile.LastName",
		"&profile.Citizenship",
		"&profile.NationalID",
		"&profile.BirthPlace",
		"&profile.BirthDate",
		"&profile.Gender",
		"&profile.MaritalStatus",
		"&profile.Occupation",
		"&profile.MonthlyIncome",
		"&profile.PhoneNumber",
		"&profile.UrgentPhoneNumber",
		"&profile.Photo",
		"&profile.IDCard",
		"&profile.SalarySlip",
		"&profile.UserID",
		"&profile.CreatedAt",
		"&profile.UpdatedAt",
	})

	profile := model.Profile{}

	suite.mockSql.ExpectBegin()
	table.AddRow(
		profile.ID,
		profile.FirstName,
		profile.LastName,
		profile.Citizenship,
		profile.NationalID,
		profile.BirthPlace,
		profile.BirthDate,
		profile.Gender,
		profile.MaritalStatus,
		profile.Occupation,
		profile.MonthlyIncome,
		profile.PhoneNumber,
		profile.UrgentPhoneNumber,
		profile.Photo,
		profile.IDCard,
		profile.SalarySlip,
		profile.UserID,
		profile.CreatedAt,
	)
	suite.mockSql.ExpectQuery("INSERT INTO profiles").WillReturnRows(table)
	suite.mockSql.ExpectCommit()

	actual, err := suite.repo.CreateProfile(profile)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), profile, actual)
}

func (suite *UserRepositoryTestSuite) TestCreateAddress_Success() {
	table := sqlmock.NewRows([]string{
		"&address.ID",
		"&address.AddressLine",
		"&address.City",
		"&address.Province",
		"&address.PostalCode",
		"&address.Country",
		"&address.ProfileID",
		"&address.CreatedAt",
		"&address.UpdatedAt",
	})

	address := model.Address{}

	suite.mockSql.ExpectBegin()
	table.AddRow(
		address.ID,
		address.AddressLine,
		address.City,
		address.Province,
		address.PostalCode,
		address.Country,
		address.ProfileID,
		address.CreatedAt,
		address.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("INSERT INTO addresses").WillReturnRows(table)
	suite.mockSql.ExpectCommit()

	profileMock := model.Profile{}

	actual, err := suite.repo.CreateAddress(address, profileMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), address, actual)
}

func (suite *UserRepositoryTestSuite) TestGetUserByStatus_Success() {
	table := sqlmock.NewRows([]string{
		"&user.ID",
		"&user.Username",
		"&user.Email",
		"&user.Role",
		"&user.Status",
		"&user.CreditScore",
		"&user.Balance",
		"&user.CreatedAt",
		"&user.UpdatedAt",
	})

	statusMock := "1"
	user := dto.ResponseStatus{
		Status: statusMock,
	}
	users := []dto.ResponseStatus{}
	users = append(users, user)

	table.AddRow(
		user.ID,
		user.Username,
		user.Email,
		user.Role,
		user.Status,
		user.CreditScore,
		user.Balance,
		user.CreatedAt,
		user.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT id,username,email,role, status, credit_score, balance,created_at,updated_at FROM users").WillReturnRows(table)

	actual, err := suite.repo.GetUserByStatus(statusMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), users[0].Status, actual[0].Status)
}

func (suite *UserRepositoryTestSuite) TestGetUserByID_Success() {
	table := sqlmock.NewRows([]string{
		"user.ID",
		"user.Username",
		"user.Email",
		"user.Password",
		"user.Role",
		"user.Status",
		"user.CreditScore",
		"user.Balance",
		"user.CreatedAt",
		"user.UpdatedAt",
		"user.Profile.ID",
		"user.Profile.FirstName",
		"user.Profile.LastName",
		"user.Profile.Citizenship",
		"user.Profile.NationalID",
		"user.Profile.BirthPlace",
		"user.Profile.BirthDate",
		"user.Profile.Gender",
		"user.Profile.MaritalStatus",
		"user.Profile.Occupation",
		"user.Profile.MonthlyIncome",
		"user.Profile.PhoneNumber",
		"user.Profile.UrgentPhoneNumber",
		"user.Profile.Photo",
		"user.Profile.IDCard",
		"user.Profile.SalarySlip",
		"user.Profile.UserID",
		"user.Profile.CreatedAt",
		"user.Profile.UpdatedAt",
		"user.Profile.Address.ID",
		"user.Profile.Address.AddressLine",
		"user.Profile.Address.City",
		"user.Profile.Address.Province",
		"user.Profile.Address.PostalCode",
		"user.Profile.Address.Country",
		"user.Profile.Address.ProfileID",
		"user.Profile.Address.CreatedAt",
		"user.Profile.Address.UpdatedAt",
	})

	idMock := "1"
	user := model.User{
		ID: idMock,
	}

	table.AddRow(
		user.ID,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
		user.Status,
		user.CreditScore,
		user.Balance,
		user.CreatedAt,
		user.UpdatedAt,
		user.Profile.ID,
		user.Profile.FirstName,
		user.Profile.LastName,
		user.Profile.Citizenship,
		user.Profile.NationalID,
		user.Profile.BirthPlace,
		user.Profile.BirthDate,
		user.Profile.Gender,
		user.Profile.MaritalStatus,
		user.Profile.Occupation,
		user.Profile.MonthlyIncome,
		user.Profile.PhoneNumber,
		user.Profile.UrgentPhoneNumber,
		user.Profile.Photo,
		user.Profile.IDCard,
		user.Profile.SalarySlip,
		user.Profile.UserID,
		user.Profile.CreatedAt,
		user.Profile.UpdatedAt,
		user.Profile.Address.ID,
		user.Profile.Address.AddressLine,
		user.Profile.Address.City,
		user.Profile.Address.Province,
		user.Profile.Address.PostalCode,
		user.Profile.Address.Country,
		user.Profile.Address.ProfileID,
		user.Profile.Address.CreatedAt,
		user.Profile.Address.UpdatedAt,
	)
	suite.mockSql.ExpectQuery(`SELECT
    u.id,
    u.username,
    u.email,
    u.password,
    u.role,
    u.status,
    u.credit_score,
    u.balance,
    u.created_at,
    u.updated_at,
    p.id,
    p.first_name,
    p.last_name,
    p.citizenship,
    p.national_id,
    p.birth_place,
    p.birth_date,
    p.gender,
    p.marital_status,
    p.occupation,
    p.monthly_income,
    p.phone_number,
    p.urgent_phone_number,
    p.photo,
    p.id_card,
    p.salary_slip,
    p.user_id,
    p.created_at,
    p.updated_at,
    a.id,
    a.address_line,
    a.city,
    a.province,
    a.postal_code,
    a.country,
    a.profile_id,
    a.created_at,
    a.updated_at
FROM
    users`).WillReturnRows(table)

	actual, err := suite.repo.GetUserByID(idMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), user.ID, actual.ID)
}

func (suite *UserRepositoryTestSuite) TestDeleteUser_Success() {
	idMock := "1"
	expected := model.User{}

	suite.mockSql.ExpectExec("DELETE FROM users").WillReturnResult(sqlmock.NewResult(1, 1))

	actual, err := suite.repo.DeleteUser(idMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *UserRepositoryTestSuite) TestGetAllUsers_Success() {
	table := sqlmock.NewRows([]string{
		"&user.ID",
		"&user.Username",
		"&user.Email",
		"&user.Role",
		"&user.Status",
		"&user.CreditScore",
		"&user.Balance",
		"&user.CreatedAt",
		"&user.UpdatedAt",
	})

	user := dto.UserDto{}
	users := []dto.UserDto{}
	users = append(users, user)

	table.AddRow(
		user.ID,
		user.Username,
		user.Email,
		user.Role,
		user.Status,
		user.CreditScore,
		user.Balance,
		user.CreatedAt,
		user.UpdatedAt,
	)
	suite.mockSql.ExpectQuery(`SELECT id, username, email, role, status, credit_score, balance, created_at, updated_at
	FROM users`).WillReturnRows(table)

	pageMock := 1
	limitMock := 1
	actual, err := suite.repo.GetAllUsers(pageMock, limitMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), users, actual)
}

func (suite *UserRepositoryTestSuite) TestUpdateBalance_Success() {
	table := suite.mockSql.NewRows([]string{"balance"})

	var balance int
	table.AddRow(
		balance,
	)
	suite.mockSql.ExpectQuery("UPDATE users").WillReturnRows(table)

	idMock := "1"
	amountMock := 1

	actual, err := suite.repo.UpdateBalance(idMock, amountMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), balance, actual)
}
