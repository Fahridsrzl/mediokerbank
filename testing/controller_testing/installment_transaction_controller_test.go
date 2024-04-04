package controllertesting

import (
	"bytes"
	"encoding/json"
	controller "medioker-bank/delivery/controller/transaction"
	middlewaremock "medioker-bank/mock/middleware_mock"
	usecasemock "medioker-bank/mock/usecase_mock"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type InstallmentTransactionControllerTestSuite struct {
	suite.Suite
	ium        *usecasemock.InstallmentTransactionUseCaseMock
	rg         *gin.RouterGroup
	amm        *middlewaremock.AuthMiddlewareMock
	controller *controller.InstallmentTransactionController
}

func (suite *InstallmentTransactionControllerTestSuite) SetupTest() {
	suite.ium = new(usecasemock.InstallmentTransactionUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.amm = new(middlewaremock.AuthMiddlewareMock)
	suite.controller = controller.NewInstallmentTransactionController(suite.ium, suite.rg, suite.amm)
}

func TestInstallmentTransactionControllerTestSuite(t *testing.T) {
	suite.Run(t, new(InstallmentTransactionControllerTestSuite))
}

func (suite *InstallmentTransactionControllerTestSuite) TestCreateTrxHandler_Success() {
	suite.ium.Mock.On("CreateTrx", mock.Anything).Return(dto.InstallmentTransactionResponseDto{}, nil)
	payloadMock := dto.InstallmentTransactionRequestDto{
		UserId:        "1",
		LoanId:        "1",
		PaymentMethod: "1",
	}
	suite.controller.Router()
	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(payloadMock)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/transactions/installments", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.CreateTrxHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

func (suite *InstallmentTransactionControllerTestSuite) TestFindTrxById_Success() {
	suite.ium.Mock.On("FindTrxById", mock.Anything).Return(model.InstallmentTransaction{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/installments/:id", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.FindTrxByIdHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *InstallmentTransactionControllerTestSuite) TestFindTrxManyhandler_Success() {
	suite.ium.Mock.On("FindTrxMany", mock.Anything).Return([]model.InstallmentTransaction{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/installments/", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.FindTrxManyHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *InstallmentTransactionControllerTestSuite) TestFindTrxByUserId_Success() {
	suite.ium.Mock.On("FindTrxByUserId", mock.Anything).Return([]model.InstallmentTransaction{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/installments/users/:userId", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.FindTrxByUserIdHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *InstallmentTransactionControllerTestSuite) TestFindTrxByUserIdAndTrxdId_Success() {
	suite.ium.Mock.On("FindTrxByUserIdAndTrxId", mock.Anything, mock.Anything).Return(model.InstallmentTransaction{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/installments/users/:userId/:trxId", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.FindTrxByUserIdAndTrxIdHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *InstallmentTransactionControllerTestSuite) TestMidtransHookHandler_Success() {
	suite.ium.Mock.On("UpdateTrxById", mock.Anything).Return(nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/installments/midtrans-hook", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.MidtransHookHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}
