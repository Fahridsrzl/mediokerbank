package controller

import (
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/master"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func (u *UserController) createHandler(ctx *gin.Context) {
	var inputData struct {
		ProfileDto dto.ProfileCreateDto `json:"profileDto"`
		AddressDto dto.AddressCreateDto `json:"addressDto"`
	}
	if err := ctx.ShouldBindJSON(&inputData); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Call use case to create user, profile, and address
	_, _, _, err := u.uc.CreateProfileAndAddressThenUpdateUser(inputData.ProfileDto, inputData.AddressDto)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare response payload
	respPayload := "Success Create Profile, Please Wait For Verified"

	common.SendCreateResponse(ctx, "ok", respPayload)
}

func (e *UserController) getStatusHandler(ctx *gin.Context) {
	status := ctx.Param("status")

	response, err := e.uc.FindByStatus(status)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "ok", response)
}

func (e *UserController) updateHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := e.uc.UpdateStatus(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

func (e *UserController) getIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := e.uc.GetUserByID(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

func (e *UserController) deletehandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := e.uc.RemoveUser(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "delete", response)
}

func (u *UserController) getAllUserHandler(ctx *gin.Context) {
	users, err := u.uc.GetAllUser()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Set response headers and send JSON data
	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) Router() {
	ur := u.rg.Group(appconfig.UserGroup)
	ur.POST(appconfig.UserAll, u.createHandler)
	ur.GET(appconfig.UserStatus, u.getStatusHandler)
	ur.GET(appconfig.UserId, u.getIdHandler)
	ur.GET(appconfig.UserAll, u.getAllUserHandler)
	ur.PUT(appconfig.UserId, u.updateHandler)
	ur.DELETE(appconfig.UserId, u.deletehandler)
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{uc: uc, rg: rg}
}
