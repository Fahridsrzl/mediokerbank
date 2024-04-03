package reposiotrytesting

import (
	"database/sql"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	repository "medioker-bank/repository/other"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.AuthRepository
}

func (suite *AuthRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = repository.NewAuthRepository(suite.mockDB)
}

func TestAuthRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthRepositoryTestSuite))
}

var tableUsers *sqlmock.Rows = sqlmock.NewRows([]string{
	"id",
	"username",
	"email",
	"password",
	"role",
	"status",
	"credit_score",
	"balance",
	"created_at",
	"updated_at",
})

var tableAdmins *sqlmock.Rows = sqlmock.NewRows([]string{
	"id",
	"username",
	"email",
	"password",
	"role",
	"created_at",
	"updated_at",
})

var tableQueue *sqlmock.Rows = sqlmock.NewRows([]string{
	"id",
	"username",
	"email",
	"password",
	"v_code",
})

func (suite *AuthRepositoryTestSuite) TestCreateQueue_Success() {
	payloadMock := dto.AuthVerifyDto{}
	expected := "register success, please check your email for verification code"

	suite.mockSql.ExpectExec("INSERT INTO queue_register_users").WillReturnResult(
		sqlmock.NewResult(1, 1),
	)

	actual, err := suite.repo.CreateQueue(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *AuthRepositoryTestSuite) TestCreteUser_Success() {
	payloadMock := dto.AuthVerifyDto{
		Username: "maman",
	}
	expected := model.User{
		Username: "maman",
	}

	tableUsers.AddRow(
		expected.ID, expected.Username, expected.Email, expected.Password, expected.Role, expected.Status, expected.CreditScore, expected.Balance, expected.CreatedAt, expected.UpdatedAt)
	suite.mockSql.ExpectQuery("INSERT INTO users").WillReturnRows(tableUsers)

	actual, err := suite.repo.CreateUser(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.ID, actual.ID)
}

func (suite *AuthRepositoryTestSuite) TestFindByVCode_Success() {
	vCodeMock := 000000
	expected := dto.AuthVerifyDto{
		VCode: vCodeMock,
	}

	tableQueue.AddRow(
		"", expected.Username, expected.Email, expected.Password, expected.VCode)
	suite.mockSql.ExpectQuery("SELECT (.+) FROM queue_register_users").WillReturnRows(tableQueue)

	actual, err := suite.repo.FindByVCode(vCodeMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.VCode, actual.VCode)
}

func (suite *AuthRepositoryTestSuite) TestDeleteQueue_Success() {
	vCodeMock := 000000

	suite.mockSql.ExpectExec("DELETE FROM queue_register_users").WillReturnResult(
		sqlmock.NewResult(1, 1),
	)

	err := suite.repo.DeleteQueue(vCodeMock)
	assert.Nil(suite.T(), err)
}

func (suite *AuthRepositoryTestSuite) TestFindUniqueUser_Success() {
	payloadMock := dto.AuthLoginDto{
		Username: "maman",
	}
	expected := model.User{
		Username: "maman",
	}

	tableUsers.AddRow(
		expected.ID, expected.Username, expected.Email, expected.Password, expected.Role, expected.Status, expected.CreditScore, expected.Balance, expected.CreatedAt, expected.UpdatedAt,
	)
	suite.mockSql.ExpectPrepare("SELECT (.+) FROM users").ExpectQuery().WillReturnRows(tableUsers)

	actual, err := suite.repo.FindUniqueUser(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Username, actual.Username)
}

func (suite *AuthRepositoryTestSuite) TestFindUniqueAdmin_Success() {
	payloadMock := dto.AuthLoginDto{
		Username: "admin",
	}
	expected := model.User{
		Username: "admin",
	}

	tableAdmins.AddRow(
		expected.ID, expected.Username, expected.Email, expected.Password, expected.Role, expected.CreatedAt, expected.UpdatedAt,
	)
	suite.mockSql.ExpectPrepare("SELECT (.+) FROM admins").ExpectQuery().WillReturnRows(tableAdmins)

	actual, err := suite.repo.FindUniqueAdmin(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Username, actual.Username)
}
