package usecasetesting

import (
	repomock "medioker-bank/mock/repository_mock"
	utilsmock "medioker-bank/mock/utils_mock"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/other"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	arm  *repomock.AuthRepoMock
	jutm *utilsmock.JwtUtilsMock
	mutm *utilsmock.MailerUtilsMock
	butm *utilsmock.BcryptUtilsMock
	au   usecase.AuthUseCase
}

func (suite *AuthUseCaseTestSuite) SetupTest() {
	suite.arm = new(repomock.AuthRepoMock)
	suite.jutm = new(utilsmock.JwtUtilsMock)
	suite.mutm = new(utilsmock.MailerUtilsMock)
	suite.butm = new(utilsmock.BcryptUtilsMock)
	suite.au = usecase.NewAuthUseCase(suite.arm, suite.jutm, suite.mutm, suite.butm)
}

func TestAuthUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}

func (suite *AuthUseCaseTestSuite) TestingRegisterUser_Success() {
	payloadMock := dto.AuthRegisterDto{
		Username:        "maman",
		Email:           "mamanabdurrahman@gmail.com",
		Password:        "123456",
		ConfirmPassword: "123456",
	}
	expectedReturn := "register success"
	hashPasswordMock := payloadMock.Password + "hash"

	suite.butm.Mock.On("GeneratePasswordHash", payloadMock.Password).Return(hashPasswordMock, nil)
	suite.arm.Mock.On("CreateQueue", mock.Anything).Return(expectedReturn, nil)
	suite.mutm.Mock.On("SendEmail", mock.Anything).Return(nil)

	actual, err := suite.au.RegisterUser(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedReturn, actual)
}

func (suite *AuthUseCaseTestSuite) TestingVerifyUser_Success() {
	vCodeMock := 000000
	suite.arm.Mock.On("FindByVCode", mock.Anything).Return(dto.AuthVerifyDto{}, nil)
	suite.arm.Mock.On("CreateUser", mock.Anything).Return(model.User{}, nil)
	suite.arm.Mock.On("DeleteQueue", mock.Anything).Return(nil)

	actual, err := suite.au.VerifyUser(vCodeMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), model.User{}, actual)
}

func (suite *AuthUseCaseTestSuite) TestingLoginUser_Success() {
	payloadMock := dto.AuthLoginDto{
		Email:    "mamanabdurrahman@gmail.com",
		Password: "123456",
	}
	accessTokenMock := "access token"
	refreshTokenMock := "refresh token"
	userMock := model.User{
		Password: payloadMock.Password,
	}
	expected := dto.AuthResponseDto{
		AccessToken:  accessTokenMock,
		RefreshToken: refreshTokenMock,
	}

	suite.arm.Mock.On("FindUniqueUser", mock.Anything).Return(userMock, nil)
	suite.butm.Mock.On("ComparePasswordHash", userMock.Password, payloadMock.Password).Return(nil)
	suite.jutm.Mock.On("GenerateToken", mock.Anything).Return(accessTokenMock, nil)
	suite.jutm.Mock.On("GenerateRefreshToken", mock.Anything).Return(refreshTokenMock, nil)

	actual, err := suite.au.LoginUser(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}

func (suite *AuthUseCaseTestSuite) TestingLoginAdmin_Success() {
	payloadMock := dto.AuthLoginDto{
		Username: "admin",
		Password: "123456",
	}
	accessTokenMock := "access token"
	refreshTokenMock := "refresh token"
	adminMock := dto.Admin{
		Password: payloadMock.Password,
	}
	expected := dto.AuthResponseDto{
		AccessToken:  accessTokenMock,
		RefreshToken: refreshTokenMock,
	}

	suite.arm.Mock.On("FindUniqueAdmin", mock.Anything).Return(adminMock, nil)
	suite.jutm.Mock.On("GenerateToken", mock.Anything).Return(accessTokenMock, nil)
	suite.jutm.Mock.On("GenerateRefreshToken", mock.Anything).Return(refreshTokenMock, nil)

	actual, err := suite.au.LoginAdmin(payloadMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}
