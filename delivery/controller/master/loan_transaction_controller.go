package controller

import (
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/master"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoanTransactionController struct {
	rg *gin.RouterGroup
	ul 	usecase.LoanTransactionUseCase
}

func (l *LoanTransactionController) GetLoanTransacttionByUserIdAndTrxId(ctx *gin.Context){
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
	rspPayload, err := l.ul.FIndLoanTransactionByUserIdAndTrxId(userId,trxId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succes", rspPayload)
}

func (l *LoanTransactionController) GetAllHandler(ctx *gin.Context){
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

func (l *LoanTransactionController) GetHandlerByUserId(ctx *gin.Context){
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

func (l *LoanTransactionController) GetHandlerById(ctx *gin.Context){
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
	br.POST(appconfig.LoanTransactionCreate, l.createHandler)
	br.GET(appconfig.LoanTransactionFindById, l.GetHandlerById)
	br.GET(appconfig.LoanTransactionFindByUserId, l.GetHandlerByUserId)
	br.GET(appconfig.LoanTransactionFindAll, l.GetAllHandler)
	br.GET(appconfig.LoanTransactionFindByUserIdAndTrxId, l.GetLoanTransacttionByUserIdAndTrxId)
}

func NewLoanTransactionController(ul usecase.LoanTransactionUseCase, rg *gin.RouterGroup) *LoanTransactionController {
	return &LoanTransactionController{ul: ul, rg: rg}
}