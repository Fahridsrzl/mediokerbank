package controller

import (
	"medioker-bank/config"
	"medioker-bank/usecase"
	"medioker-bank/utils/common"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	uc  usecase.AuthUseCase
	rg  *gin.RouterGroup
	jwt common.JwtToken
}

func (a *AuthController) registerHandler(*gin.Context) {

}

func (a *AuthController) loginUserHandler(*gin.Context) {

}

func (a *AuthController) verifyHandler(*gin.Context) {

}

func (a *AuthController) loginAdminHandler(*gin.Context) {

}

func (a *AuthController) RefreshTokenHandler(*gin.Context) {

}

func (a *AuthController) Router() {
	authGroup := a.rg.Group(config.AuthGroup)
	{
		authGroup.POST(config.AuthRegisterUser, a.registerHandler)
		authGroup.POST(config.AuthVerifyUser, a.verifyHandler)
		authGroup.POST(config.AuthLoginUser, a.loginUserHandler)
		authGroup.POST(config.AuthLoginAdmin, a.loginAdminHandler)
		authGroup.POST(config.AuthRefreshToken, a.RefreshTokenHandler)
	}
}

func NewAuthController(uc usecase.AuthUseCase, rg *gin.RouterGroup, jwt common.JwtToken) *AuthController {
	return &AuthController{uc: uc, rg: rg, jwt: jwt}
}
