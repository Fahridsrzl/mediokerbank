package controller

import (
	"fmt"
	"medioker-bank/config"
	"medioker-bank/model"
	"medioker-bank/usecase"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func (u *UserController) createHandler(ctx *gin.Context) {
	var payload model.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	respPayload, err := u.uc.RegisterNewUser(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "ok", respPayload)
}

func (u *UserController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}
	respPayload, err := u.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", respPayload)
}

func (u *UserController) getallhandler(ctx *gin.Context) {
	users, err := u.uc.ShowAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get users"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok", "data": users})
}

func (u *UserController) updateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	var updateUser model.User
	if err := ctx.BindJSON(&updateUser); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request payload")
		return
	}

	updateUser.Id = id

	err := u.uc.ModifyUser(id, updateUser)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	updatedUser, err := u.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", updatedUser)
}

func (u *UserController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}
	err := u.uc.RemoveUser(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("failed to delete user: %s", id))
		return
	}

	common.SendSingleResponse(ctx, "ok", nil)
}

func (u *UserController) Router() {
	ur := u.rg.Group(config.UserGroup)
	ur.POST(config.UserAll, u.createHandler)
	ur.GET(config.UserId, u.getHandler)
	ur.GET(config.UserAll, u.getallhandler)
	ur.DELETE(config.UserId, u.deleteHandler)
	ur.PUT(config.UserId, u.updateHandler)
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{uc: uc, rg: rg}
}
