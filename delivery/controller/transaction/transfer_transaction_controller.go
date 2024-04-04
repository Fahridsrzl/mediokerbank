package controller

import (
	"errors"
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/delivery/middleware"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/transaction"
	"medioker-bank/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransferController struct {
	tc  usecase.TransferUseCase
	rg  *gin.RouterGroup
	jwt middleware.AuthMiddleware
}

// createHandler handles creating a new transfer
// @Summary Create a new transfer
// @Description Create a new transfer
// @Tags Transfer
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param body body dto.TransferDto true "Transfer data"
// @Success 201 {string} string "Transfer created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/transfers [post]
func (t *TransferController) CreateHandler(ctx *gin.Context) {
	var payload dto.TransferDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	userId := ctx.MustGet("id")
	if userId != payload.SenderID {
		common.SendErrorResponse(ctx, http.StatusBadRequest, errors.New("forbidden action, You should make transactions on your own account").Error())
		return
	}
	respPayload, err := t.tc.CreateTransfer(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "ok", respPayload)
}

// getSenderIdHandler handles getting a transfer by sender ID
// @Summary Get transfer by sender ID
// @Description Retrieve transfer details by sender ID
// @Tags Transfer
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param senderId path string true "Sender ID"
// @Success 200 {object} map[string]interface{} "Transfer details"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/transfers/users/{senderId} [get]
func (e *TransferController) GetSenderIdHandler(ctx *gin.Context) {
	id := ctx.Param("senderId")

	response, err := e.tc.GetTransferBySenderId(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

// getTransferIdHandler handles getting a transfer by ID
// @Summary Get transfer by ID
// @Description Retrieve transfer details by ID
// @Tags Transfer
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param id path string true "Transfer ID"
// @Success 200 {object} map[string]interface{} "Transfer details"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/transfers/{id} [get]
func (e *TransferController) GetTransferIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := e.tc.GetTransferByTransferId(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

// getAllTransferHandler handles getting all transfers
// @Summary Get all transfers
// @Description Retrieve all transfers
// @Tags Transfer
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of transfers"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/transfers [get]
func (u *TransferController) GetAllTransferHandler(ctx *gin.Context) {
	param1 := ctx.Query("page")
	param2 := ctx.Query("limit")
	if param1 == "" || param2 == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, errors.New("need query params 'page' & 'limit'").Error())
		return
	}
	page, err := strconv.Atoi(param1)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	limit, err := strconv.Atoi(param2)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	transfers, err := u.tc.GetAllTransfer(page, limit)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Set response headers and send JSON data
	common.SendSingleResponse(ctx, "ok", transfers)
}

func (t *TransferController) Router() {
	ur := t.rg.Group(appconfig.TransferGroup)
	{
		ur.POST(appconfig.Transfer, t.jwt.RequireToken("user"), t.CreateHandler)
		ur.GET(appconfig.Transfer, t.jwt.RequireToken("admin"), t.GetAllTransferHandler)
		ur.GET(appconfig.TransferSenderId, t.jwt.RequireToken("user"), t.GetSenderIdHandler)
		ur.GET(appconfig.TransferId, t.jwt.RequireToken("admin", "user"), t.GetTransferIdHandler)
	}
}

func NewTransferController(tc usecase.TransferUseCase, rg *gin.RouterGroup, jwt middleware.AuthMiddleware) *TransferController {
	return &TransferController{tc: tc, rg: rg, jwt: jwt}
}
