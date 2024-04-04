package usecasetesting

import (
	repomock "medioker-bank/mock/repository_mock"
	umock "medioker-bank/mock/usecase_mock"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	uTransaction "medioker-bank/usecase/transaction"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type LoanTransactionUseCaseTestSuite struct {
	suite.Suite
	trxrm  *repomock.LoanTransactionRepoMock
	uucm *umock.UserUseCaseMock
	lpucm *umock.LoanProductMock
	lrm *repomock.LoanRepoMock
	trxu uTransaction.LoanTransactionUseCase
}

func (suite *LoanTransactionUseCaseTestSuite) SetupTest(){
	suite.trxrm = new(repomock.LoanTransactionRepoMock)
	suite.uucm = new(umock.UserUseCaseMock)
	suite.lpucm = new(umock.LoanProductMock)
	suite.lrm = new(repomock.LoanRepoMock)
	suite.trxu = uTransaction.NewLoanTransactionUseCase(suite.trxrm, suite.uucm, suite.lpucm, suite.lrm)
}

func TestLoanTransactionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LoanTransactionUseCaseTestSuite))
}

func (suite *LoanTransactionUseCaseTestSuite) TestFIndLoanTransactionByUserIdAndTrxId_Success(){
	userId := "1"
	trxId := "1"
	suite.trxrm.Mock.On("GetByUserIdAndTrxId",mock.Anything, mock.Anything).Return([]model.LoanTransaction{}, nil)
	actual, err := suite.trxu.FIndLoanTransactionByUserIdAndTrxId(userId,trxId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []model.LoanTransaction{}, actual)
}

func (suite *LoanTransactionUseCaseTestSuite) TestFindAllLoanTransaction_Success(){
	suite.trxrm.Mock.On("GetAll",mock.Anything).Return([]model.LoanTransaction{}, nil)
	actual, err := suite.trxu.FindAllLoanTransaction()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), []model.LoanTransaction{}, actual)
}

func (suite *LoanTransactionUseCaseTestSuite) TestFindByUserId_Success(){
	userId := "1"
	suite.trxrm.Mock.On("GetByUserID",mock.Anything).Return(model.LoanTransaction{}, nil)
	actual, err := suite.trxu.FindByUserId(userId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), model.LoanTransaction{}, actual)
}

func (suite *LoanTransactionUseCaseTestSuite) TestFindById_Success(){
	id := "1"
	suite.trxrm.Mock.On("GetByID",mock.Anything).Return(model.LoanTransaction{}, nil)
	actual, err := suite.trxu.FindById(id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), model.LoanTransaction{}, actual)
}

func (suite *LoanTransactionUseCaseTestSuite) TestRegisterNewTransaction_Success() {
    userId := "1"
    userMock := model.User{
        ID:       userId,
        Username: "JohnDoe",
        Email:    "johndoe@example.com",
    }
    suite.uucm.Mock.On("GetUserByID", userId).Return(userMock,[]model.Loan{}, nil)

    productMock := model.LoanProduct{
        Name:                 "Product Name",
        MaxAmount:            1000,
        MinInstallmentPeriod: 1,
        MaxInstallmentPeriod: 12,
    }
    suite.lpucm.Mock.On("FindLoanProductById", mock.Anything).Return(productMock, nil)

	suite.trxrm.Mock.On("Create", mock.AnythingOfType("model.LoanTransaction")).Return(model.LoanTransaction{}, nil)

	suite.lrm.Mock.On("Create", mock.AnythingOfType("model.Loan")).Return(model.Loan{}, nil)

    actual, err := suite.trxu.RegisterNewTransaction(dto.LoanTransactionRequestDto{
        UserId: userId,
        LoanTransactionDetail: []dto.LoanTransactionDetailRequestDto{
            {
                ProductId:         "1",
                Amount:            500,
                Purpose:           "Purpose",
                InstallmentPeriod: 6,
            },
        },
    })
    assert.Nil(suite.T(), err)
    assert.Equal(suite.T(),model.LoanTransaction{}, actual)
}
