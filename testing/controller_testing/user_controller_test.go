package controllertesting

import (
	"bytes"
	controller "medioker-bank/delivery/controller/master"
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

type UserControllerTestSuite struct {
	suite.Suite
	uum        *usecasemock.UserUseCaseMock
	rg         *gin.RouterGroup
	amm        *middlewaremock.AuthMiddlewareMock
	controller *controller.UserController
}

func (suite *UserControllerTestSuite) SetupTest() {
	suite.uum = new(usecasemock.UserUseCaseMock)
	rg := gin.Default()
	suite.rg = rg.Group("/api/v1")
	suite.amm = new(middlewaremock.AuthMiddlewareMock)
	suite.controller = controller.NewUserController(suite.uum, suite.rg, suite.amm)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}

func (suite *UserControllerTestSuite) TestGetStatusHandler_Success() {
	suite.uum.Mock.On("FindByStatus", mock.Anything).Return([]dto.ResponseStatus{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/status/:status", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.GetStatusHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *UserControllerTestSuite) TestGetidHandler_Success() {

}

func (suite *UserControllerTestSuite) TestGetAllUserHandler_Success() {
	suite.uum.Mock.On("GetAllUser", mock.Anything, mock.Anything).Return([]dto.UserDto{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/api/v1/users?page=1&limit=1", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.GetAllUserHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *UserControllerTestSuite) TestUpdateHandler_Success() {
	suite.uum.Mock.On("UpdateStatus", mock.Anything).Return(nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPut, "/api/v1/users/:id", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.UpdateHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *UserControllerTestSuite) TestDeletehandler_Success() {
	suite.uum.Mock.On("RemoveUser", mock.Anything).Return(model.User{}, nil)
	suite.controller.Router()
	record := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/users/:id", bytes.NewBuffer([]byte{}))
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	suite.controller.Deletehandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}
