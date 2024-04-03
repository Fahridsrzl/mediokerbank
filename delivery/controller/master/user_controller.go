package controller

import (
	"encoding/json"
	"fmt"
	"math/rand"
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/delivery/middleware"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/master"
	"medioker-bank/utils/common"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc  usecase.UserUseCase
	rg  *gin.RouterGroup
	jwt middleware.AuthMiddleware
}

func (u *UserController) CreateHandler(ctx *gin.Context) {
	profile := ctx.PostForm("profile")
	address := ctx.PostForm("address")

	filePhoto, headerPhoto, err := ctx.Request.FormFile("photo")
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	defer filePhoto.Close()

	fileIdCard, headerIdCard, err := ctx.Request.FormFile("idCard")
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	defer fileIdCard.Close()

	fileSalarySlip, headerSalarySlip, err := ctx.Request.FormFile("salarySlip")
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	defer fileSalarySlip.Close()

	random := rand.New(rand.NewSource(time.Now().UnixNano())).Int() / 1e13
	photoName := fmt.Sprintf("photo_%v%s", random, filepath.Ext(headerPhoto.Filename))
	idCardName := fmt.Sprintf("id_%v%s", random, filepath.Ext(headerIdCard.Filename))
	salarySlipName := fmt.Sprintf("salary_%v%s", random, filepath.Ext(headerSalarySlip.Filename))
	randomString := strconv.Itoa(random)
	fileLocation := filepath.Join("uploads", "file_["+randomString+"]")
	photoLocation := filepath.Join(fileLocation, photoName)
	idCardLocation := filepath.Join(fileLocation, idCardName)
	salarySlipLocation := filepath.Join(fileLocation, salarySlipName)

	os.MkdirAll(filepath.Dir(photoLocation), os.ModePerm)
	os.MkdirAll(filepath.Dir(idCardLocation), os.ModePerm)
	os.MkdirAll(filepath.Dir(salarySlipLocation), os.ModePerm)

	if err := ctx.SaveUploadedFile(headerPhoto, photoLocation); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if err := ctx.SaveUploadedFile(headerIdCard, idCardLocation); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if err := ctx.SaveUploadedFile(headerSalarySlip, salarySlipLocation); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var profileDto dto.ProfileCreateDto
	var addressDto dto.AddressCreateDto

	json.Unmarshal([]byte(profile), &profileDto)
	json.Unmarshal([]byte(address), &addressDto)

	profileDto.Photo = photoLocation
	profileDto.IDCard = idCardLocation
	profileDto.SalarySlip = salarySlipLocation

	// Call use case to create user, profile, and address
	_, _, err = u.uc.CreateProfileAndAddressThenUpdateUser(profileDto, addressDto)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare response payload
	respPayload := "Success Create Profile, Please Wait For Verified"

	common.SendCreateResponse(ctx, "ok", respPayload)
}

func (e *UserController) GetStatusHandler(ctx *gin.Context) {
	status := ctx.Param("status")

	response, err := e.uc.FindByStatus(status)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "ok", response)
}

func (e *UserController) UpdateHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	err := e.uc.UpdateStatus(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := "success verifying user with id: " + id

	common.SendSingleResponse(ctx, "ok", response)
}

func (u *UserController) GetidHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	user, loans, err := u.uc.GetUserByID(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	role := ctx.MustGet("role")
	userId := ctx.MustGet("id")
	if role == "user" {
		if userId != user.ID {
			common.SendErrorResponse(ctx, http.StatusForbidden, "forbidden action")
			return
		}
	}

	response := struct {
		User  model.User   `json:"user"`
		Loans []model.Loan `json:"loans"`
	}{
		User:  user,
		Loans: loans,
	}

	common.SendSingleResponse(ctx, "ok", response)
}

func (e *UserController) Deletehandler(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := e.uc.RemoveUser(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response := "success delete user with id: " + id

	common.SendSingleResponse(ctx, "delete", response)
}

func (u *UserController) GetAllUserHandler(ctx *gin.Context) {
	users, err := u.uc.GetAllUser()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Set response headers and send JSON data
	common.SendSingleResponse(ctx, "ok", users)
}

func (u *UserController) Router() {
	ur := u.rg.Group(appconfig.UserGroup)
	{
		ur.POST(appconfig.UserAll, u.jwt.RequireToken("user"), u.CreateHandler)
		ur.GET(appconfig.UserStatus, u.jwt.RequireToken("admin"), u.GetStatusHandler)
		ur.GET(appconfig.UserId, u.jwt.RequireToken("admin", "user"), u.GetidHandler)
		ur.GET(appconfig.UserAll, u.jwt.RequireToken("admin"), u.GetAllUserHandler)
		ur.PUT(appconfig.UserId, u.jwt.RequireToken("admin"), u.UpdateHandler)
		ur.DELETE(appconfig.UserId, u.jwt.RequireToken("admin"), u.Deletehandler)
	}
}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup, jwt middleware.AuthMiddleware) *UserController {
	return &UserController{uc: uc, rg: rg, jwt: jwt}
}
