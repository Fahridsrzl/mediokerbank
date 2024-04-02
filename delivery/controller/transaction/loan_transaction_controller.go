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

func (l *LoanTransactionController) GetAllHandler(ctx *gin.Context) {
	response, err := l.ul.FindAllLoanTransaction()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succes", response)
}

func (l *LoanTransactionController) createHandler(ctx *gin.Context) {
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
		br.POST(appconfig.LoanTransactionCreate, l.jwt.RequireToken("user"), l.createHandler)
		br.GET(appconfig.LoanTransactionFindById, l.jwt.RequireToken("admin"), l.GetHandlerById)
		br.GET(appconfig.LoanTransactionFindByUserId, l.jwt.RequireToken("user"), l.GetHandlerByUserId)
		br.GET(appconfig.LoanTransactionFindAll, l.jwt.RequireToken("admin"), l.GetAllHandler)
		br.GET(appconfig.LoanTransactionFindByUserIdAndTrxId, l.jwt.RequireToken("user"), l.GetLoanTransacttionByUserIdAndTrxId)
	}
}

func NewLoanTransactionController(ul transaction.LoanTransactionUseCase, rg *gin.RouterGroup, jwt middleware.AuthMiddleware) *LoanTransactionController {
	return &LoanTransactionController{ul: ul, rg: rg, jwt: jwt}
}
