package reposiotrytesting

import (
	"database/sql"
	"medioker-bank/model"
	repository "medioker-bank/repository/transaction"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoanTransactionRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.LoanTransactionRepository
}

func (suite *LoanTransactionRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = repository.NewLoanTransactionRepository(suite.mockDB)
}

func TestLoanTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(LoanTransactionRepositoryTestSuite))
}

func (suite *LoanTransactionRepositoryTestSuite) TestCreateLoanTransactions_Success() {
	var tableLoanTransactions *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"trxDate",
		"userId",
		"status",
		"createdAt",
		"updatedAt",
	})

	var tableLoanTransactionDetails *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"amount",
		"loanProduct",
		"purpose",
		"interest",
		"installmentPeriod",
		"installmentUnit",
		"installmentAmount",
		"trxId",
		"createdAt",
		"updatedAt",
	})

	payloadMock := model.LoanTransaction{
		Id:   "1",
		User: model.User{ID: "1"},
		LoanTransactionDetaills: []model.LoanTransactionDetail{
			{Id: "1", LoanProduct: model.LoanProduct{Id: "1"}, Amount: 1000, Purpose: "Test"},
		},
	}
	expected := model.LoanTransaction{
		Id:   "1",
		User: model.User{ID: "1"},
		LoanTransactionDetaills: []model.LoanTransactionDetail{
			{Id: "1", LoanProduct: model.LoanProduct{Id: "1"}, Amount: 1000, Purpose: "Test"},
		},
	}
	expectedTrxDetails := model.LoanTransactionDetail{
		Id:          "1",
		LoanProduct: model.LoanProduct{Id: "1"},
		Amount:      1000,
		Purpose:     "Test",
	}

	suite.mockSql.ExpectBegin()
	tableLoanTransactions.AddRow(
		expected.Id, expected.TrxDate, expected.User.ID, expected.Status, expected.CreatedAt, expected.UpdatedAt)
	suite.mockSql.ExpectQuery("INSERT INTO loan_transactions").WillReturnRows(tableLoanTransactions)
	tableLoanTransactionDetails.AddRow(
		expectedTrxDetails.Id, expectedTrxDetails.LoanProduct.Id, expectedTrxDetails.Amount, expectedTrxDetails.Purpose, expectedTrxDetails.Interest, expectedTrxDetails.InstallmentPeriod, expectedTrxDetails.InstallmentUnit, expectedTrxDetails.InstallmentAmount, expectedTrxDetails.TrxId, expectedTrxDetails.CreatedAt, expectedTrxDetails.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("INSERT INTO loan_transaction_details").WillReturnRows(tableLoanTransactionDetails)
	suite.mockSql.ExpectCommit()

	actual, err := suite.repo.Create(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual.Id)
}

func (suite *LoanTransactionRepositoryTestSuite) TestGetAllLoanTransactions_Success() {
	var penyatuanDari2TableDiatas *sqlmock.Rows = sqlmock.NewRows([]string{
		"&loanTransaction.Id",
		"&loanTransaction.TrxDate",
		"&user.ID",
		"&user.Username",
		"&user.Email",
		"&loanTransaction.Status",
		"&loanTransaction.CreatedAt",
		"&loanTransaction.UpdatedAt",
		"&loanTransactionDetail.Id",
		"&loanProduct.Name",
		"&loanProduct.MaxAmount",
		"&loanProduct.MinInstallmentPeriod",
		"&loanProduct.MaxInstallmentPeriod",
		"&loanProduct.InstallmentPeriodUnit",
		"&loanProduct.AdminFee",
		"&loanProduct.MinCreditScore",
		"&loanProduct.MinMonthlyIncome",
		"&loanProduct.CreatedAt",
		"&loanProduct.UpdatedAt",
		"&loanTransactionDetail.Amount",
		"&loanTransactionDetail.Purpose",
		"&loanTransactionDetail.Interest",
		"&loanTransactionDetail.InstallmentPeriod",
		"&loanTransactionDetail.InstallmentUnit",
		"&loanTransactionDetail.InstallmentAmount",
		"&loanTransactionDetail.CreatedAt",
		"&loanTransactionDetail.UpdatedAt",
	})

	expected := []model.LoanTransaction{
		{
			Id: "1",
		},
	}
	var user model.User
	var loanTransactionDetail model.LoanTransactionDetail
	var loanProduct model.LoanProduct
	var loanTransaction model.LoanTransaction
	penyatuanDari2TableDiatas.AddRow(
		&loanTransaction.Id,
		&loanTransaction.TrxDate,
		&user.ID,
		&user.Username,
		&user.Email,
		&loanTransaction.Status,
		&loanTransaction.CreatedAt,
		&loanTransaction.UpdatedAt,
		&loanTransactionDetail.Id,
		&loanProduct.Name,
		&loanProduct.MaxAmount,
		&loanProduct.MinInstallmentPeriod,
		&loanProduct.MaxInstallmentPeriod,
		&loanProduct.InstallmentPeriodUnit,
		&loanProduct.AdminFee,
		&loanProduct.MinCreditScore,
		&loanProduct.MinMonthlyIncome,
		&loanProduct.CreatedAt,
		&loanProduct.UpdatedAt,
		&loanTransactionDetail.Amount,
		&loanTransactionDetail.Purpose,
		&loanTransactionDetail.Interest,
		&loanTransactionDetail.InstallmentPeriod,
		&loanTransactionDetail.InstallmentUnit,
		&loanTransactionDetail.InstallmentAmount,
		&loanTransactionDetail.CreatedAt,
		&loanTransactionDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) FROM loan_transactions").WillReturnRows(penyatuanDari2TableDiatas)
	pageMock := 1
	limitMock := 1
	actual, err := suite.repo.GetAll(pageMock, limitMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(expected), len(actual))
	err = suite.mockSql.ExpectationsWereMet()
	assert.Nil(suite.T(), err)
}

func (suite *LoanTransactionRepositoryTestSuite) TestGetByUserIdAndTrxIdLoanTransactions_Success() {
	var penyatuanDari2TableDiatas *sqlmock.Rows = sqlmock.NewRows([]string{
		"&loanTransaction.Id",
		"&loanTransaction.TrxDate",
		"&user.ID",
		"&user.Username",
		"&user.Email",
		"&loanTransaction.Status",
		"&loanTransaction.CreatedAt",
		"&loanTransaction.UpdatedAt",
		"&loanTransactionDetail.Id",
		"&loanProduct.Name",
		"&loanProduct.MaxAmount",
		"&loanProduct.MinInstallmentPeriod",
		"&loanProduct.MaxInstallmentPeriod",
		"&loanProduct.InstallmentPeriodUnit",
		"&loanProduct.AdminFee",
		"&loanProduct.MinCreditScore",
		"&loanProduct.MinMonthlyIncome",
		"&loanProduct.CreatedAt",
		"&loanProduct.UpdatedAt",
		"&loanTransactionDetail.Amount",
		"&loanTransactionDetail.Purpose",
		"&loanTransactionDetail.Interest",
		"&loanTransactionDetail.InstallmentPeriod",
		"&loanTransactionDetail.InstallmentUnit",
		"&loanTransactionDetail.InstallmentAmount",
		"&loanTransactionDetail.CreatedAt",
		"&loanTransactionDetail.UpdatedAt",
	})

	userId := "1"
	trxId := "1"
	expected := model.LoanTransaction{
		User: model.User{
			ID: "1",
		},
		Id: "1",
	}
	var user model.User
	var loanTransactionDetail model.LoanTransactionDetail
	var loanProduct model.LoanProduct
	var loanTransaction model.LoanTransaction
	penyatuanDari2TableDiatas.AddRow(
		&loanTransaction.Id,
		&loanTransaction.TrxDate,
		&user.ID,
		&user.Username,
		&user.Email,
		&loanTransaction.Status,
		&loanTransaction.CreatedAt,
		&loanTransaction.UpdatedAt,
		&loanTransactionDetail.Id,
		&loanProduct.Name,
		&loanProduct.MaxAmount,
		&loanProduct.MinInstallmentPeriod,
		&loanProduct.MaxInstallmentPeriod,
		&loanProduct.InstallmentPeriodUnit,
		&loanProduct.AdminFee,
		&loanProduct.MinCreditScore,
		&loanProduct.MinMonthlyIncome,
		&loanProduct.CreatedAt,
		&loanProduct.UpdatedAt,
		&loanTransactionDetail.Amount,
		&loanTransactionDetail.Purpose,
		&loanTransactionDetail.Interest,
		&loanTransactionDetail.InstallmentPeriod,
		&loanTransactionDetail.InstallmentUnit,
		&loanTransactionDetail.InstallmentAmount,
		&loanTransactionDetail.CreatedAt,
		&loanTransactionDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) FROM loan_transactions").WillReturnRows(penyatuanDari2TableDiatas)

	actual, err := suite.repo.GetByUserIdAndTrxId(userId, trxId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual.Id)
	assert.Equal(suite.T(), expected.User.ID, actual.UserId)
}

func (suite *LoanTransactionRepositoryTestSuite) TestGetByUserIdLoanTransactions_Success() {
	var penyatuanDari2TableDiatas *sqlmock.Rows = sqlmock.NewRows([]string{
		"&loanTransaction.Id",
		"&loanTransaction.TrxDate",
		"&user.ID",
		"&user.Username",
		"&user.Email",
		"&loanTransaction.Status",
		"&loanTransaction.CreatedAt",
		"&loanTransaction.UpdatedAt",
		"&loanTransactionDetail.Id",
		"&loanProduct.Name",
		"&loanProduct.MaxAmount",
		"&loanProduct.MinInstallmentPeriod",
		"&loanProduct.MaxInstallmentPeriod",
		"&loanProduct.InstallmentPeriodUnit",
		"&loanProduct.AdminFee",
		"&loanProduct.MinCreditScore",
		"&loanProduct.MinMonthlyIncome",
		"&loanProduct.CreatedAt",
		"&loanProduct.UpdatedAt",
		"&loanTransactionDetail.Amount",
		"&loanTransactionDetail.Purpose",
		"&loanTransactionDetail.Interest",
		"&loanTransactionDetail.InstallmentPeriod",
		"&loanTransactionDetail.InstallmentUnit",
		"&loanTransactionDetail.InstallmentAmount",
		"&loanTransactionDetail.CreatedAt",
		"&loanTransactionDetail.UpdatedAt",
	})

	userId := "1"
	expected := model.LoanTransaction{
		User: model.User{
			ID: "1",
		},
	}
	var user model.User
	var loanTransactionDetail model.LoanTransactionDetail
	var loanProduct model.LoanProduct
	var loanTransaction model.LoanTransaction
	penyatuanDari2TableDiatas.AddRow(
		&loanTransaction.Id,
		&loanTransaction.TrxDate,
		&user.ID,
		&user.Username,
		&user.Email,
		&loanTransaction.Status,
		&loanTransaction.CreatedAt,
		&loanTransaction.UpdatedAt,
		&loanTransactionDetail.Id,
		&loanProduct.Name,
		&loanProduct.MaxAmount,
		&loanProduct.MinInstallmentPeriod,
		&loanProduct.MaxInstallmentPeriod,
		&loanProduct.InstallmentPeriodUnit,
		&loanProduct.AdminFee,
		&loanProduct.MinCreditScore,
		&loanProduct.MinMonthlyIncome,
		&loanProduct.CreatedAt,
		&loanProduct.UpdatedAt,
		&loanTransactionDetail.Amount,
		&loanTransactionDetail.Purpose,
		&loanTransactionDetail.Interest,
		&loanTransactionDetail.InstallmentPeriod,
		&loanTransactionDetail.InstallmentUnit,
		&loanTransactionDetail.InstallmentAmount,
		&loanTransactionDetail.CreatedAt,
		&loanTransactionDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) FROM loan_transactions").WillReturnRows(penyatuanDari2TableDiatas)

	actual, err := suite.repo.GetByUserID(userId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual[0].UserId)
}

func (suite *LoanTransactionRepositoryTestSuite) TestGetByIdLoanTransactions_Success() {
	var penyatuanDari2TableDiatas *sqlmock.Rows = sqlmock.NewRows([]string{
		"&loanTransaction.Id",
		"&loanTransaction.TrxDate",
		"&user.ID",
		"&user.Username",
		"&user.Email",
		"&loanTransaction.Status",
		"&loanTransaction.CreatedAt",
		"&loanTransaction.UpdatedAt",
		"&loanTransactionDetail.Id",
		"&loanProduct.Name",
		"&loanProduct.MaxAmount",
		"&loanProduct.MinInstallmentPeriod",
		"&loanProduct.MaxInstallmentPeriod",
		"&loanProduct.InstallmentPeriodUnit",
		"&loanProduct.AdminFee",
		"&loanProduct.MinCreditScore",
		"&loanProduct.MinMonthlyIncome",
		"&loanProduct.CreatedAt",
		"&loanProduct.UpdatedAt",
		"&loanTransactionDetail.Amount",
		"&loanTransactionDetail.Purpose",
		"&loanTransactionDetail.Interest",
		"&loanTransactionDetail.InstallmentPeriod",
		"&loanTransactionDetail.InstallmentUnit",
		"&loanTransactionDetail.InstallmentAmount",
		"&loanTransactionDetail.CreatedAt",
		"&loanTransactionDetail.UpdatedAt",
	})

	Id := ""
	expected := model.LoanTransaction{
		Id: "",
	}
	var user model.User
	var loanTransactionDetail model.LoanTransactionDetail
	var loanProduct model.LoanProduct
	var loanTransaction model.LoanTransaction
	penyatuanDari2TableDiatas.AddRow(
		&loanTransaction.Id,
		&loanTransaction.TrxDate,
		&user.ID,
		&user.Username,
		&user.Email,
		&loanTransaction.Status,
		&loanTransaction.CreatedAt,
		&loanTransaction.UpdatedAt,
		&loanTransactionDetail.Id,
		&loanProduct.Name,
		&loanProduct.MaxAmount,
		&loanProduct.MinInstallmentPeriod,
		&loanProduct.MaxInstallmentPeriod,
		&loanProduct.InstallmentPeriodUnit,
		&loanProduct.AdminFee,
		&loanProduct.MinCreditScore,
		&loanProduct.MinMonthlyIncome,
		&loanProduct.CreatedAt,
		&loanProduct.UpdatedAt,
		&loanTransactionDetail.Amount,
		&loanTransactionDetail.Purpose,
		&loanTransactionDetail.Interest,
		&loanTransactionDetail.InstallmentPeriod,
		&loanTransactionDetail.InstallmentUnit,
		&loanTransactionDetail.InstallmentAmount,
		&loanTransactionDetail.CreatedAt,
		&loanTransactionDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) FROM loan_transactions").WillReturnRows(penyatuanDari2TableDiatas)

	actual, err := suite.repo.GetByID(Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual.Id)
}
