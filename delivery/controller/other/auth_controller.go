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

// Register endpoint
// @Summary Register a new user
// @Description Register a new user
// @ID register-user
// @Accept json
// @Produce json
// @Param body body dto.AuthRegisterDto true "User registration details"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/users/register [post]
func (a *AuthController) registerHandler(ctx *gin.Context) {
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

// Register endpoint
// @Summary Register a new user
// @Description Register a new user
// @ID verify-user
// @Accept json
// @Produce json
// @Param body body dto.AuthRegisterDto true "User registration details"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/users/register/verify [post]
func (a *AuthController) verifyHandler(ctx *gin.Context) {
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

// Login user endpoint
// @Summary Login user
// @Description Login user with username and password
// @ID login-user
// @Accept json
// @Produce json
// @Param body body dto.AuthLoginDto true "User login credentials"
// @Success 200 {string} string "User logged in successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/users/login [post]
func (a *AuthController) loginUserHandler(ctx *gin.Context) {
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

// Login admin endpoint
// @Summary Login admin
// @Description Login admin with username and password
// @ID login-admin
// @Accept json
// @Produce json
// @Param body body dto.AuthLoginDto true "Admin login credentials"
// @Success 200 {string} string "Admin logged in successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/admins/login [post]
func (a *AuthController) loginAdminHandler(ctx *gin.Context) {
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

// Refresh token endpoint
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @ID refresh-token
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Success 200 {string} string "Access token refreshed successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/refresh-token [post]
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
		authGroup.POST(appconfig.AuthRegisterUser, a.registerHandler)
		authGroup.POST(appconfig.AuthVerifyUser, a.verifyHandler)
		authGroup.POST(appconfig.AuthLoginUser, a.loginUserHandler)
		authGroup.POST(appconfig.AuthLoginAdmin, a.loginAdminHandler)
		authGroup.POST(appconfig.AuthRefreshToken, a.RefreshTokenHandler)
	}
}

func NewAuthController(uc usecase.AuthUseCase, rg *gin.RouterGroup, jwt common.JwtToken) *AuthController {
	return &AuthController{uc: uc, rg: rg, jwt: jwt}
}
