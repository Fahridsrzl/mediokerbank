package controller

import (
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/delivery/middleware"
	"medioker-bank/model/dto"
	transaction "medioker-bank/usecase/transaction"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoanTransactionController struct {
	rg  *gin.RouterGroup
	ul  transaction.LoanTransactionUseCase
	jwt middleware.AuthMiddleware
}

// GetLoanTransacttionByUserIdAndTrxId handles finding a loan transaction by user ID and transaction ID
// @Summary Find loan transaction by user ID and transaction ID
// @Description Retrieve loan transaction details by user ID and transaction ID
// @Tags Loan Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Param trxId path string true "Transaction ID"
// @Success 200 {object} map[string]interface{} "Loan transaction details"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Loan transaction not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/loans/users/{userId}/{trxId} [get]
func (l *LoanTransactionController) GetLoanTransacttionByUserIdAndTrxId(ctx *gin.Context) {
	userId := ctx.Param("userId")
	if userId == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "user_id can't be empty")
		return
	}
	trxId := ctx.Param("trxId")
	if userId == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "user_id can't be empty")
		return
	}
	rspPayload, err := l.ul.FIndLoanTransactionByUserIdAndTrxId(userId, trxId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succes", rspPayload)
}

// GetAllHandler handles getting all loan transactions
// @Summary Get all loan transactions
// @Description Retrieve all loan transactions
// @Tags Loan Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of loan transactions"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/loans [get]
func (l *LoanTransactionController) GetAllHandler(ctx *gin.Context) {
	response, err := l.ul.FindAllLoanTransaction()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succes", response)
}

// createHandler handles creating a new loan transaction
// @Summary Create a new loan transaction
// @Description Create a new loan transaction
// @Tags Loan Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param body body dto.LoanTransactionRequestDto true "Loan transaction data"
// @Success 201 {string} string "Loan transaction created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/loans [post]
func (l *LoanTransactionController) CreateHandler(ctx *gin.Context) {
	var payload dto.LoanTransactionRequestDto
	if err := ctx.ShouldBind(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	rspPayload, err := l.ul.RegisterNewTransaction(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "ok", rspPayload)
}

// GetHandlerByUserId handles getting loan transactions by user ID
// @Summary Get loan transactions by user ID
// @Description Retrieve loan transactions for a specific user
// @Tags Loan Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} map[string]interface{} "List of loan transactions"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Loan transactions not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/loans/users/{userId} [get]
func (l *LoanTransactionController) GetHandlerByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")
	if userId == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "user_id can't be empty")
		return
	}
	rspPayload, err := l.ul.FindByUserId(userId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

// GetHandlerById handles getting a loan transaction by ID
// @Summary Get loan transaction by ID
// @Description Retrieve loan transaction details by ID
// @Tags Loan Transaction
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} map[string]interface{} "Loan transaction details"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Loan transaction not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/loans/{id} [get]
func (l *LoanTransactionController) GetHandlerById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}
	rspPayload, err := l.ul.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (l *LoanTransactionController) Router() {
	br := l.rg.Group(appconfig.LoanTransactionGroup)
	{
		br.POST(appconfig.LoanTransactionCreate, l.jwt.RequireToken("user"), l.CreateHandler)
		br.GET(appconfig.LoanTransactionFindById, l.jwt.RequireToken("admin"), l.GetHandlerById)
		br.GET(appconfig.LoanTransactionFindByUserId, l.jwt.RequireToken("user"), l.GetHandlerByUserId)
		br.GET(appconfig.LoanTransactionFindAll, l.jwt.RequireToken("admin"), l.GetAllHandler)
		br.GET(appconfig.LoanTransactionFindByUserIdAndTrxId, l.jwt.RequireToken("user"), l.GetLoanTransacttionByUserIdAndTrxId)
	}
}

func NewLoanTransactionController(ul transaction.LoanTransactionUseCase, rg *gin.RouterGroup, jwt middleware.AuthMiddleware) *LoanTransactionController {
	return &LoanTransactionController{ul: ul, rg: rg, jwt: jwt}
}
