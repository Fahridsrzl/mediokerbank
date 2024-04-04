package usecasetesting

import (
	repomock "medioker-bank/mock/repository_mock"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/master"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	urm *repomock.UserRepoMock
	lrm *repomock.LoanRepoMock
	uu  usecase.UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.urm = new(repomock.UserRepoMock)
	suite.lrm = new(repomock.LoanRepoMock)
	suite.uu = usecase.NewUserUseCase(suite.urm, suite.lrm)
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}

func (suite *UserUseCaseTestSuite) TestCreateProfileAndAddressThenUpdateUser_Success() {
	payloadProfileMock := dto.ProfileCreateDto{
		FirstName: "maman",
	}
	payloadAddressMock := dto.AddressCreateDto{
		Country: "indonesia",
	}

	expectedProfile := model.Profile{
		FirstName: "maman",
	}
	expectedAddress := model.Address{
		Country: "indonesia",
	}

	suite.urm.Mock.On("CreateProfile", mock.Anything).Return(expectedProfile, nil)
	suite.urm.Mock.On("CreateAddress", mock.Anything).Return(expectedAddress, nil)
	suite.urm.Mock.On("UpdateUser", mock.Anything).Return(nil)

	actualProfile, actualAddress, err := suite.uu.CreateProfileAndAddressThenUpdateUser(payloadProfileMock, payloadAddressMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedProfile.FirstName, actualProfile.FirstName)
	assert.Equal(suite.T(), expectedAddress.Country, actualAddress.Country)
}

func (suite *UserUseCaseTestSuite) TestFindByStatus_Success() {
	statusMock := "1"
	expected := []dto.ResponseStatus{
		{Status: statusMock},
	}

	suite.urm.Mock.On("GetUserByStatus", mock.Anything).Return(expected, nil)
	actual, err := suite.uu.FindByStatus(statusMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected[0].Status, actual[0].Status)

}

func (suite *UserUseCaseTestSuite) TestUpdateStatus_Success() {
	idMock := "1"

	suite.urm.Mock.On("UpdateUser", mock.Anything).Return(nil)

	err := suite.uu.UpdateStatus(idMock)
	assert.Nil(suite.T(), err)
}

func (suite *UserUseCaseTestSuite) TestGetUserByID_Success() {
	idMock := "1"
	expectedUser := model.User{
		ID: idMock,
	}
	expectedLoans := []model.Loan{
		{UserId: idMock},
	}

	suite.urm.Mock.On("GetUserByID", mock.Anything).Return(expectedUser, nil)
	suite.lrm.Mock.On("FindByUserId", mock.Anything).Return(expectedLoans, nil)

	actualUser, actualLoans, err := suite.uu.GetUserByID(idMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedUser.ID, actualUser.ID)
	assert.Equal(suite.T(), expectedLoans[0].UserId, actualLoans[0].UserId)

}

func (suite *UserUseCaseTestSuite) TestRemoveUser_Success() {
	idMock := "1"
	expected := model.User{
		ID: idMock,
	}

	suite.urm.Mock.On("DeleteUser", mock.Anything).Return(expected, nil)

	actual, err := suite.uu.RemoveUser(idMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.ID, actual.ID)
}

func (suite *UserUseCaseTestSuite) TestGetAllUser_Success() {
	expected := []dto.UserDto{
		{ID: "1"},
	}

	suite.urm.Mock.On("GetAllUsers", mock.Anything, mock.Anything).Return(expected, nil)

	pageMock := 1
	limitMock := 1
	actual, err := suite.uu.GetAllUser(pageMock, limitMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), len(expected), len(actual))
}

func (suite *UserUseCaseTestSuite) TestUpdateUserBalance_Success() {
	idMock := "1"
	amountMock := 1
	expected := 1

	suite.urm.Mock.On("UpdateBalance", mock.Anything, mock.Anything).Return(expected, nil)

	actual, err := suite.uu.UpdateUserBalance(idMock, amountMock)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, actual)
}
