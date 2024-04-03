package controllertesting

import (
	"bytes"
	"encoding/json"
	controller "medioker-bank/delivery/controller/other"
	usecasemock "medioker-bank/mock/usecase_mock"
	utilsmock "medioker-bank/mock/utils_mock"
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

type AuthControllerTestSuite struct {
	suite.Suite
	aum        *usecasemock.AuthUseCaseMock
	rg         *gin.RouterGroup
	jutm       *utilsmock.JwtUtilsMock
	controller *controller.AuthController
}

func (suite *AuthControllerTestSuite) SetupTest() {
	suite.aum = new(usecasemock.AuthUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.jutm = new(utilsmock.JwtUtilsMock)
	suite.controller = controller.NewAuthController(suite.aum, suite.rg, suite.jutm)
}

func TestBillControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}

func (suite *AuthControllerTestSuite) TestRegisterHandler_Success() {
	suite.aum.Mock.On("RegisterUser", mock.Anything).Return("queue registered", nil)
	payloadMock := dto.AuthRegisterDto{
		Username:        "maman",
		Email:           "mamanabdurrahman@gmail.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	suite.controller.Router()
	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(payloadMock)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.RegisterHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

func (suite *AuthControllerTestSuite) TestVerifyHandler_Success() {
	suite.aum.Mock.On("VerifyUser", mock.Anything).Return(model.User{}, nil)
	payloadMock := dto.AuthVcodeDto{
		VCode: 111111,
	}
	suite.controller.Router()
	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(payloadMock)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/register/verify", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.VerifyHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}

func (suite *AuthControllerTestSuite) TestLoginUserHandler() {
	suite.aum.Mock.On("LoginUser", mock.Anything).Return(dto.AuthResponseDto{}, nil)
	payloadMock := dto.AuthLoginDto{
		Username: "maman",
		Password: "123456",
	}
	suite.controller.Router()
	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(payloadMock)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/users/login", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.LoginUserHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *AuthControllerTestSuite) TestLoginAdminHandler_Success() {
	suite.aum.Mock.On("LoginAdmin", mock.Anything).Return(dto.AuthResponseDto{}, nil)
	payloadMock := dto.AuthLoginDto{
		Email:    "mamanabdurrahman@gmail.com",
		Password: "123456",
	}
	suite.controller.Router()
	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(payloadMock)
	assert.NoError(suite.T(), err)
	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/admins/login", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.LoginAdminHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *AuthControllerTestSuite) TestRefreshTokenHandler_Success() {
	newTokenMock := "access token"
	suite.jutm.Mock.On("RefreshToken", mock.Anything).Return(newTokenMock, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/refresh-token", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	req.Header.Set("Authorization", "Bearer refresh token")
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.RefreshTokenHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, record.Code)
}
