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

type InstallmentTransactionController struct {
	uc  usecase.InstallmentTransactionUseCase
	rg  *gin.RouterGroup
	jwt middleware.AuthMiddleware
}

// createTrxHandler handles creating a new installment transaction
// @Summary Create a new installment transaction
// @Description Create a new installment transaction
// @Tags Installment Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param body body dto.InstallmentTransactionRequestDto true "Installment transaction data"
// @Success 201 {string} string "Installment transaction created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/installments [post]
func (i *InstallmentTransactionController) CreateTrxHandler(ctx *gin.Context) {
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
	common.SendCreateResponse(ctx, "Success", response)
}

// findTrxByIdHandler handles finding an installment transaction by ID
// @Summary Find installment transaction by ID
// @Description Retrieve installment transaction details by ID
// @Tags Installment Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} map[string]interface{} "Installment transaction details"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/installments/{id} [get]
func (i *InstallmentTransactionController) FindTrxByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := i.uc.FindTrxById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

// findTrxByIdHandler handles finding an installment transaction by ID
// @Summary Find installment transaction by ID
// @Description Retrieve installment transaction details by ID
// @Tags Installment Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} map[string]interface{} "Installment transaction details"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/installments/{id} [get]
func (i *InstallmentTransactionController) FindTrxManyHandler(ctx *gin.Context) {
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
	payload := dto.InstallmentTransactionSearchDto{
		TrxDate: ctx.Query("trxDate"),
	}
	response, err := i.uc.FindTrxMany(payload, page, limit)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

// findTrxByUserIdHandler handles finding installment transactions by user ID
// @Summary Find installment transactions by user ID
// @Description Retrieve installment transactions for a specific user
// @Tags Installment Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param trxDate query string false "Transaction Date"
// @Success 200 {object} map[string]interface{} "List of installment transactions"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/installments/users/{userId} [get]
func (i *InstallmentTransactionController) FindTrxByUserIdHandler(ctx *gin.Context) {
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

// findTrxByUserIdAndTrxIdHandler handles finding an installment transaction by user ID and transaction ID
// @Summary Find installment transaction by user ID and transaction ID
// @Description Retrieve installment transaction details by user ID and transaction ID
// @Tags Installment Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param trxId path string true "Transaction ID"
// @Success 200 {object} map[string]interface{} "Installment transaction details"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/installments/users/{userId}/{trxId} [get]
func (i *InstallmentTransactionController) FindTrxByUserIdAndTrxIdHandler(ctx *gin.Context) {
	userId := ctx.Param("userId")
	trxId := ctx.Param("trxId")
	response, err := i.uc.FindTrxByUserIdAndTrxId(userId, trxId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

// midtransHookHandler handles the Midtrans webhook for installment transactions
// @Summary Handle Midtrans webhook for installment transactions
// @Description Handle Midtrans webhook to update installment transaction status
// @Tags Installment Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param order_id query string true "Order ID"
// @Success 200 {string} string "Installment transaction status updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/installments/midtrans-hook [get]
func (i *InstallmentTransactionController) MidtransHookHandler(ctx *gin.Context) {
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
		installmentGroup.POST(appconfig.InstallmentCreate, i.jwt.RequireToken("user"), i.CreateTrxHandler)
		installmentGroup.GET(appconfig.InstallmentFindTrxById, i.jwt.RequireToken("admin", "user"), i.FindTrxByIdHandler)
		installmentGroup.GET(appconfig.InstallmentFindTrxMany, i.jwt.RequireToken("admin"), i.FindTrxManyHandler)
		installmentGroup.GET(appconfig.InstallmentFindTrxByUserId, i.jwt.RequireToken("user"), i.FindTrxByUserIdHandler)
		installmentGroup.GET(appconfig.InstallmentFindTrxByUserAndTrxId, i.jwt.RequireToken("user"), i.FindTrxByUserIdAndTrxIdHandler)
		installmentGroup.GET(appconfig.InstallmentMidtransHook, i.MidtransHookHandler)
	}
}

func NewInstallmentTransactionController(uc usecase.InstallmentTransactionUseCase, rg *gin.RouterGroup, jwt middleware.AuthMiddleware) *InstallmentTransactionController {
	return &InstallmentTransactionController{uc: uc, rg: rg, jwt: jwt}
}
