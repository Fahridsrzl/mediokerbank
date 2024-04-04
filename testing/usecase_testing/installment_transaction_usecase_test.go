package usecasetesting

import (
	repomock "medioker-bank/mock/repository_mock"
	usecasemock "medioker-bank/mock/usecase_mock"
	utilsmock "medioker-bank/mock/utils_mock"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/transaction"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type InstallmentTransactionUseCaseTestSuite struct {
	suite.Suite
	irm  *repomock.InstallmentTransactionRepoMock
	lrm  *repomock.LoanRepoMock
	uum  *usecasemock.UserUseCaseMock
	pum  *usecasemock.LoanProductMock
	mutm *utilsmock.MidtrasnUtilsMock
	iu   usecase.InstallmentTransactionUseCase
}

func (suite *InstallmentTransactionUseCaseTestSuite) SetupTest() {
	suite.irm = new(repomock.InstallmentTransactionRepoMock)
	suite.lrm = new(repomock.LoanRepoMock)
	suite.uum = new(usecasemock.UserUseCaseMock)
	suite.pum = new(usecasemock.LoanProductMock)
	suite.mutm = new(utilsmock.MidtrasnUtilsMock)
	suite.iu = usecase.NewInstallmentTransactionUseCase(suite.irm, suite.lrm, suite.uum, suite.pum, suite.mutm)
}

func TestInstallmentUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(InstallmentTransactionUseCaseTestSuite))
}

// func (suite *InstallmentTransactionUseCaseTestSuite) TestCreateTrx_Success() {
// 	payloadMock := dto.InstallmentTransactionRequestDto{
// 		UserId:        "1",
// 		PaymentMethod: "medioker balance",
// 	}
// 	loanMock := model.Loan{
// 		Id: "1",
// 		LoanProduct: model.LoanProduct{
// 			Id: "1",
// 		},
// 		PeriodLeft: 5,
// 	}
// 	userMock := model.User{
// 		ID: "1",
// 	}
// 	trxMock := model.InstallmentTransaction{}
// 	responseMock := dto.InstallmentTransactionResponseDto{
// 		Message:     "transaction success, your loan updated",
// 		PaymentLink: "-",
// 		Transaction: trxMock,
// 	}

// 	suite.lrm.Mock.On("FindByUserId", payloadMock.UserId).Return([]model.Loan{}, nil)
// 	suite.pum.Mock.On("FindLoanProductById", mock.Anything).Return(loanMock.LoanProduct, nil)
// 	suite.uum.Mock.On("GetUserByID", mock.Anything).Return(userMock, nil)
// 	suite.irm.Mock.On("Create", mock.Anything).Return(trxMock, nil)
// 	suite.lrm.Mock.On("UpdatePeriod", mock.Anything).Return(nil)
// 	suite.uum.Mock.On("UpdateUserBalance", mock.Anything).Return(0, nil)

// 	actual, err := suite.iu.CreateTrx(payloadMock)
// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), responseMock, actual)
// }

func (suite *InstallmentTransactionUseCaseTestSuite) TestFindTrxById_Success() {
	idMock := "1"
	trxMock := model.InstallmentTransaction{
		Id:     idMock,
		UserId: "1",
	}
	loansMock := []model.Loan{
		{LoanProduct: model.LoanProduct{
			Id: "1",
		}},
	}

	suite.irm.Mock.On("FindById", mock.Anything).Return(trxMock, nil)
	suite.lrm.Mock.On("FindByUserId", mock.Anything).Return(loansMock, nil)
	suite.pum.Mock.On("FindLoanProductById", mock.Anything).Return(model.LoanProduct{}, nil)

	actual, err := suite.iu.FindTrxById(idMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), trxMock.Id, actual.Id)
}

func (suite *InstallmentTransactionUseCaseTestSuite) TestFindTrxMany_Success() {
	payloadMock := dto.InstallmentTransactionSearchDto{}
	trxMock := model.InstallmentTransaction{
		Id:     "1",
		UserId: "1",
	}
	trxsMock := []model.InstallmentTransaction{}
	trxsMock = append(trxsMock, trxMock)
	loansMock := []model.Loan{
		{LoanProduct: model.LoanProduct{
			Id: "1",
		}},
	}

	suite.irm.Mock.On("FindAll", mock.Anything).Return(trxsMock, nil)
	suite.lrm.Mock.On("FindByUserId", mock.Anything).Return(loansMock, nil)
	suite.pum.Mock.On("FindLoanProductById", mock.Anything).Return(model.LoanProduct{}, nil)

	actual, err := suite.iu.FindTrxMany(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(trxsMock), len(actual))
}

func (suite *InstallmentTransactionUseCaseTestSuite) TestFindTrxByUserId_Success() {
	userIdMock := "1"
	payloadMock := dto.InstallmentTransactionSearchDto{}
	trxMock := model.InstallmentTransaction{
		Id:     "1",
		UserId: "1",
	}
	trxsMock := []model.InstallmentTransaction{}
	trxsMock = append(trxsMock, trxMock)
	loansMock := []model.Loan{
		{LoanProduct: model.LoanProduct{
			Id: "1",
		}},
	}

	suite.irm.Mock.On("FindByUserId", mock.Anything, mock.Anything).Return(trxsMock, nil)
	suite.lrm.Mock.On("FindByUserId", mock.Anything).Return(loansMock, nil)
	suite.pum.Mock.On("FindLoanProductById", mock.Anything).Return(model.LoanProduct{}, nil)

	actual, err := suite.iu.FindTrxByUserId(userIdMock, payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(trxsMock), len(actual))
}

func (suite *InstallmentTransactionUseCaseTestSuite) TestFindTrxByUserIdAndTrxId_Success() {
	userIdMock := "1"
	trxIdMock := "1"
	trxMock := model.InstallmentTransaction{
		Id:     "1",
		UserId: "1",
	}
	loansMock := []model.Loan{
		{LoanProduct: model.LoanProduct{
			Id: "1",
		}},
	}

	suite.irm.Mock.On("FindByUserIdAndTrxId", userIdMock, trxIdMock).Return(trxMock, nil)
	suite.lrm.Mock.On("FindByUserId", mock.Anything).Return(loansMock, nil)
	suite.pum.Mock.On("FindLoanProductById", mock.Anything).Return(model.LoanProduct{}, nil)

	actual, err := suite.iu.FindTrxByUserIdAndTrxId(userIdMock, trxIdMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), trxMock.Id, actual.Id)
}

func (suite *InstallmentTransactionUseCaseTestSuite) TestUpdateTrxById_Success() {
	trxIdMock := "1"
	loanIdMock := "1"

	suite.irm.Mock.On("UpdateById", mock.Anything).Return(loanIdMock, nil)
	suite.lrm.Mock.On("UpdatePeriod", mock.Anything).Return(nil)
	suite.pum.Mock.On("FindLoanProductById", mock.Anything).Return(model.LoanProduct{}, nil)

	err := suite.iu.UpdateTrxById(trxIdMock)
	assert.Nil(suite.T(), err)
}
