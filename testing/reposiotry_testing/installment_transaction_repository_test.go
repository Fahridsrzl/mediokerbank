package reposiotrytesting

import (
	"database/sql"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	repository "medioker-bank/repository/transaction"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InstallmentTransactionRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    repository.InstallmentTransactionRepository
}

func (suite *InstallmentTransactionRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.Nil(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = repository.NewInstallmentTransactionRepository(suite.mockDB)
}

func TestInstallmentTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(InstallmentTransactionRepositoryTestSuite))
}

func (suite *InstallmentTransactionRepositoryTestSuite) TestCreate_Success() {
	var tableTrx *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"trx_date",
		"user_id",
		"status",
		"created_at",
		"updated_at",
	})

	var tableTrxd *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"loan_id",
		"installment_amount",
		"payment_method",
		"trx_id",
		"created_at",
		"updated_at",
	})
	payloadMock := model.InstallmentTransaction{
		Id: "1",
	}
	expected := model.InstallmentTransaction{
		Id: "1",
	}

	suite.mockSql.ExpectBegin()

	tableTrx.AddRow(
		expected.Id, expected.TrxDate, expected.UserId, expected.Status, expected.CreatedAt, expected.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("INSERT INTO installment_transactions").WillReturnRows(tableTrx)

	tableTrxd.AddRow(
		expected.TrxDetail.Id, expected.TrxDetail.Loan.Id, expected.TrxDetail.InstallmentAmount, expected.TrxDetail.PaymentMethod, expected.TrxDetail.TrxId, expected.TrxDetail.CreatedAt, expected.TrxDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("INSERT INTO installment_transaction_details").WillReturnRows(tableTrxd)

	suite.mockSql.ExpectCommit()

	actual, err := suite.repo.Create(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual.Id)
}

func (suite *InstallmentTransactionRepositoryTestSuite) TestFindById_Success() {
	var tableTrx *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"trx_date",
		"user_id",
		"status",
		"created_at",
		"updated_at",
	})

	var tableTrxd *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"loan_id",
		"installment_amount",
		"payment_method",
		"trx_id",
		"created_at",
		"updated_at",
	})
	idMock := "1"
	expected := model.InstallmentTransaction{
		Id: "1",
	}

	tableTrx.AddRow(
		expected.Id, expected.TrxDate, expected.UserId, expected.Status, expected.CreatedAt, expected.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) FROM installment_transactions").WillReturnRows(tableTrx)

	tableTrxd.AddRow(
		expected.TrxDetail.Id, expected.TrxDetail.Loan.Id, expected.TrxDetail.InstallmentAmount, expected.TrxDetail.PaymentMethod, expected.TrxDetail.TrxId, expected.TrxDetail.CreatedAt, expected.TrxDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) installment_transaction_details").WillReturnRows(tableTrxd)

	actual, err := suite.repo.FindById(idMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.Id, actual.Id)
}

func (suite *InstallmentTransactionRepositoryTestSuite) TestFindAll_Success() {
	var tableTrx *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"trx_date",
		"user_id",
		"status",
		"created_at",
		"updated_at",
	})

	var tableTrxd *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"loan_id",
		"installment_amount",
		"payment_method",
		"trx_id",
		"created_at",
		"updated_at",
	})
	expected := model.InstallmentTransaction{}

	tableTrx.AddRow(
		expected.Id, expected.TrxDate, expected.UserId, expected.Status, expected.CreatedAt, expected.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) FROM installment_transactions").WillReturnRows(tableTrx)

	tableTrxd.AddRow(
		expected.TrxDetail.Id, expected.TrxDetail.Loan.Id, expected.TrxDetail.InstallmentAmount, expected.TrxDetail.PaymentMethod, expected.TrxDetail.TrxId, expected.TrxDetail.CreatedAt, expected.TrxDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) installment_transaction_details").WillReturnRows(tableTrxd)

	_, err := suite.repo.FindAll()
	assert.Nil(suite.T(), err)
}

func (suite *InstallmentTransactionRepositoryTestSuite) TestFindMany_Success() {
	var tableTrx *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"trx_date",
		"user_id",
		"status",
		"created_at",
		"updated_at",
	})

	var tableTrxd *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"loan_id",
		"installment_amount",
		"payment_method",
		"trx_id",
		"created_at",
		"updated_at",
	})
	payloadMock := dto.InstallmentTransactionSearchDto{
		TrxDate: time.Now().String(),
	}
	expected := model.InstallmentTransaction{
		TrxDate: time.Now(),
	}

	tableTrx.AddRow(
		expected.Id, expected.TrxDate, expected.UserId, expected.Status, expected.CreatedAt, expected.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) FROM installment_transactions").WillReturnRows(tableTrx)

	tableTrxd.AddRow(
		expected.TrxDetail.Id, expected.TrxDetail.Loan.Id, expected.TrxDetail.InstallmentAmount, expected.TrxDetail.PaymentMethod, expected.TrxDetail.TrxId, expected.TrxDetail.CreatedAt, expected.TrxDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) installment_transaction_details").WillReturnRows(tableTrxd)

	actual, err := suite.repo.FindMany(payloadMock)
	assert.Nil(suite.T(), err)
	notfound := false
	for _, item := range actual {
		if item.TrxDate == expected.TrxDate {
			notfound = false
			assert.Equal(suite.T(), expected.TrxDate, item.TrxDate)
			break
		} else {
			notfound = true
		}
	}
	if notfound == true {
		suite.T().Fatal()
	}
}

func (suite *InstallmentTransactionRepositoryTestSuite) TestFindByUserId_Success() {
	var tableTrx *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"trx_date",
		"user_id",
		"status",
		"created_at",
		"updated_at",
	})

	var tableTrxd *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"loan_id",
		"installment_amount",
		"payment_method",
		"trx_id",
		"created_at",
		"updated_at",
	})
	useridMock := "1"
	payloadMock := dto.InstallmentTransactionSearchDto{
		TrxDate: time.Now().String(),
	}
	expected := model.InstallmentTransaction{
		UserId:  useridMock,
		TrxDate: time.Now(),
	}

	tableTrx.AddRow(
		expected.Id, expected.TrxDate, expected.UserId, expected.Status, expected.CreatedAt, expected.UpdatedAt,
	)
	suite.mockSql.ExpectPrepare("SELECT (.+) FROM installment_transactions").ExpectQuery().WillReturnRows(tableTrx)

	tableTrxd.AddRow(
		expected.TrxDetail.Id, expected.TrxDetail.Loan.Id, expected.TrxDetail.InstallmentAmount, expected.TrxDetail.PaymentMethod, expected.TrxDetail.TrxId, expected.TrxDetail.CreatedAt, expected.TrxDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) installment_transaction_details").WillReturnRows(tableTrxd)

	actual, err := suite.repo.FindByUserId(useridMock, payloadMock)
	assert.Nil(suite.T(), err)
	notfound := false
	for _, item := range actual {
		if item.TrxDate == expected.TrxDate && item.UserId == useridMock {
			notfound = false
			assert.Equal(suite.T(), expected.TrxDate, item.TrxDate)
			assert.Equal(suite.T(), useridMock, item.UserId)
			break
		} else {
			notfound = true
		}
	}
	if notfound == true {
		suite.T().Fatal()
	}
}

func (suite *InstallmentTransactionRepositoryTestSuite) TestFindByUserIdAndTrxId_Success() {
	var tableTrx *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"trx_date",
		"user_id",
		"status",
		"created_at",
		"updated_at",
	})

	var tableTrxd *sqlmock.Rows = sqlmock.NewRows([]string{
		"id",
		"loan_id",
		"installment_amount",
		"payment_method",
		"trx_id",
		"created_at",
		"updated_at",
	})
	useridMock := "1"
	trxIdMock := "1"
	expected := model.InstallmentTransaction{
		Id:     trxIdMock,
		UserId: useridMock,
	}

	tableTrx.AddRow(
		expected.Id, expected.TrxDate, expected.UserId, expected.Status, expected.CreatedAt, expected.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) FROM installment_transactions").WillReturnRows(tableTrx)

	tableTrxd.AddRow(
		expected.TrxDetail.Id, expected.TrxDetail.Loan.Id, expected.TrxDetail.InstallmentAmount, expected.TrxDetail.PaymentMethod, expected.TrxDetail.TrxId, expected.TrxDetail.CreatedAt, expected.TrxDetail.UpdatedAt,
	)
	suite.mockSql.ExpectQuery("SELECT (.+) installment_transaction_details").WillReturnRows(tableTrxd)

	actual, err := suite.repo.FindByUserIdAndTrxId(useridMock, trxIdMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), trxIdMock, actual.Id)
	assert.Equal(suite.T(), useridMock, actual.UserId)
}

func (suite *InstallmentTransactionRepositoryTestSuite) TestUpdateById_Success() {

	loanIdRows := sqlmock.NewRows([]string{"id"})

	loanIdMock := "1"
	expected := loanIdMock

	suite.mockSql.ExpectExec("UPDATE installment_transactions").WillReturnResult(sqlmock.NewResult(1, 1))

	loanIdRows.AddRow(
		expected,
	)
	suite.mockSql.ExpectQuery("SELECT loan_id from installment_transaction_details").WillReturnRows(loanIdRows)

	actual, err := suite.repo.UpdateById(loanIdMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), loanIdMock, actual)
}

func (suite *InstallmentTransactionRepositoryTestSuite) TestDeleteById_Success() {
	idMock := "1"

	suite.mockSql.ExpectExec("DELETE FROM installment_transactions").WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.repo.DeleteById(idMock)
	assert.Nil(suite.T(), err)
}
