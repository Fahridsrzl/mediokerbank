package controller

import (
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/transaction"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InstallmentTransactionController struct {
	uc usecase.InstallmentTransactionUseCase
	rg *gin.RouterGroup
}

func (i *InstallmentTransactionController) createTrxHandler(ctx *gin.Context) {
	var payload dto.InstallmentTransactionRequestDto
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := i.uc.CreateTrx(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (i *InstallmentTransactionController) findTrxByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := i.uc.FindTrxById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (i *InstallmentTransactionController) findTrxManyHandler(ctx *gin.Context) {
	payload := dto.InstallmentTransactionSearchDto{
		TrxDate: ctx.Query("trxDate"),
	}
	response, err := i.uc.FindTrxMany(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (i *InstallmentTransactionController) findTrxByUserIdHandler(ctx *gin.Context) {
	userId := ctx.Param("userId")
	payload := dto.InstallmentTransactionSearchDto{
		TrxDate: ctx.Query("trxDate"),
	}
	response, err := i.uc.FindTrxByUserId(userId, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (i *InstallmentTransactionController) findTrxByUserIdAndTrxIdHandler(ctx *gin.Context) {
	userId := ctx.Param("userId")
	trxId := ctx.Param("trxId")
	response, err := i.uc.FindTrxByUserIdAndTrxId(userId, trxId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (i *InstallmentTransactionController) midtransHookHandler(ctx *gin.Context) {
	id := ctx.Query("order_id")
	err := i.uc.UpdateTrxById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", id)
}

func (i *InstallmentTransactionController) Router() {
	installmentGroup := i.rg.Group(appconfig.InstallmentGroup)
	{
		installmentGroup.POST(appconfig.InstallmentCreate, i.createTrxHandler)
		installmentGroup.GET(appconfig.InstallmentFindTrxById, i.findTrxByIdHandler)
		installmentGroup.GET(appconfig.InstallmentFindTrxMany, i.findTrxManyHandler)
		installmentGroup.GET(appconfig.InstallmentFindTrxByUserId, i.findTrxByUserIdHandler)
		installmentGroup.GET(appconfig.InstallmentFindTrxByUserAndTrxId, i.findTrxByUserIdAndTrxIdHandler)
		installmentGroup.GET(appconfig.InstallmentMidtransHook, i.midtransHookHandler)
	}
}

func NewInstallmentTransactionController(uc usecase.InstallmentTransactionUseCase, rg *gin.RouterGroup) *InstallmentTransactionController {
	return &InstallmentTransactionController{uc: uc, rg: rg}
}
