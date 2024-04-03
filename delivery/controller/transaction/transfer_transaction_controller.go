package controller

import (
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/delivery/middleware"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/transaction"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferController struct {
	tc  usecase.TransferUseCase
	rg  *gin.RouterGroup
	jwt middleware.AuthMiddleware
}

func (t *TransferController) CreateHandler(ctx *gin.Context) {
	var payload dto.TransferDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	respPayload, err := t.tc.CreateTransfer(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "ok", respPayload)
}

func (e *TransferController) GetSenderIdHandler(ctx *gin.Context) {
	id := ctx.Param("senderId")

	response, err := e.tc.GetTransferBySenderId(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

func (e *TransferController) GetTransferIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := e.tc.GetTransferByTransferId(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

func (u *TransferController) GetAllTransferHandler(ctx *gin.Context) {
	transfers, err := u.tc.GetAllTransfer()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Set response headers and send JSON data
	ctx.JSON(http.StatusOK, transfers)
}

func (t *TransferController) Router() {
	ur := t.rg.Group(appconfig.TransferGroup)
	{
		ur.POST(appconfig.Transfer, t.jwt.RequireToken("user"), t.CreateHandler)
		ur.GET(appconfig.Transfer, t.jwt.RequireToken("admin"), t.GetAllTransferHandler)
		ur.GET(appconfig.TransferSenderId, t.jwt.RequireToken("user"), t.GetSenderIdHandler)
		ur.GET(appconfig.TransferId, t.jwt.RequireToken("user"), t.GetTransferIdHandler)
	}
}

func NewTransferController(tc usecase.TransferUseCase, rg *gin.RouterGroup, jwt middleware.AuthMiddleware) *TransferController {
	return &TransferController{tc: tc, rg: rg, jwt: jwt}
}
