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

type LoanTransactionControllerTestSuite struct {
	suite.Suite
	lum        *usecasemock.LoanTransactionMock
	rg         *gin.RouterGroup
	jutm       *middlewaremock.AuthMiddlewareMock
	controller *controller.LoanTransactionController
}

func (suite *LoanTransactionControllerTestSuite) SetupTest() {
	suite.lum = new(usecasemock.LoanTransactionMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.jutm = new(middlewaremock.AuthMiddlewareMock)
	suite.controller = controller.NewLoanTransactionController(suite.lum, suite.rg, suite.jutm)
}

func TestLoanTransactionControllerTestSuite(t *testing.T) {
	suite.Run(t, new(LoanTransactionControllerTestSuite))
}

func (suite *LoanTransactionControllerTestSuite) TestGetLoanTransacttionByUserIdAndTrxIdHandler_Success() {
	suite.lum.Mock.On("FIndLoanTransactionByUserIdAndTrxId", mock.Anything, mock.Anything).Return([]model.LoanTransaction{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/loans/users/:userId/:trxId", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.GetLoanTransacttionByUserIdAndTrxId(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *LoanTransactionControllerTestSuite) TestGetAllHandlerHandler_Success(){
	suite.lum.Mock.On("FindAllLoanTransaction").Return([]model.LoanTransaction{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/loans", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.GetAllHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *LoanTransactionControllerTestSuite) TestGetHandlerByUserId_Success(){
	suite.lum.Mock.On("FindByUserId", mock.Anything).Return(model.LoanTransaction{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/loans/users/:userId", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.GetHandlerByUserId(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *LoanTransactionControllerTestSuite) TestGetHandlerById_Success(){
	suite.lum.Mock.On("FindById", mock.Anything).Return(model.LoanTransaction{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/transactions/loans/:Id", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.GetHandlerById(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *LoanTransactionControllerTestSuite) TestCreateHandler_Success(){
	suite.lum.Mock.On("RegisterNewTransaction", mock.Anything).Return(model.LoanTransaction{}, nil)
	payloadMock := dto.LoanTransactionRequestDto{
		UserId: "1",
		LoanTransactionDetail: []dto.LoanTransactionDetailRequestDto{
			{
				ProductId: "1",
				Amount: 1,
				Purpose: "berhasil",
				InstallmentPeriod: 1,
			},
		},
	}
	suite.controller.Router()
	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(payloadMock)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/transactions/loans", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.CreateHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}
