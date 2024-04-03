package controller

import (
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/other"
	"medioker-bank/utils/common"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	uc  usecase.AuthUseCase
	rg  *gin.RouterGroup
	jwt common.JwtToken
}

func (a *AuthController) RegisterHandler(ctx *gin.Context) {
	var payload dto.AuthRegisterDto
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := a.uc.RegisterUser(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "Success", response)
}

func (a *AuthController) VerifyHandler(ctx *gin.Context) {
	var vCode dto.AuthVcodeDto
	err := ctx.ShouldBindJSON(&vCode)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := a.uc.VerifyUser(vCode.VCode)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "Success", response)
}

func (a *AuthController) LoginUserHandler(ctx *gin.Context) {
	var payload dto.AuthLoginDto
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := a.uc.LoginUser(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (a *AuthController) LoginAdminHandler(ctx *gin.Context) {
	var payload dto.AuthLoginDto
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := a.uc.LoginAdmin(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (a *AuthController) RefreshTokenHandler(ctx *gin.Context) {
	refreshToken := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", -1)
	newAccessToken, err := a.jwt.RefreshToken(refreshToken)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "Success", newAccessToken)
}

func (a *AuthController) Router() {
	authGroup := a.rg.Group(appconfig.AuthGroup)
	{
		authGroup.POST(appconfig.AuthRegisterUser, a.RegisterHandler)
		authGroup.POST(appconfig.AuthVerifyUser, a.VerifyHandler)
		authGroup.POST(appconfig.AuthLoginUser, a.LoginUserHandler)
		authGroup.POST(appconfig.AuthLoginAdmin, a.LoginAdminHandler)
		authGroup.POST(appconfig.AuthRefreshToken, a.RefreshTokenHandler)
	}
}

func NewAuthController(uc usecase.AuthUseCase, rg *gin.RouterGroup, jwt common.JwtToken) *AuthController {
	return &AuthController{uc: uc, rg: rg, jwt: jwt}
}
