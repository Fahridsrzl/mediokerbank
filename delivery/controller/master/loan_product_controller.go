package controller

import (
	"medioker-bank/delivery/middleware"
	"medioker-bank/model"
	"medioker-bank/usecase"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoanProductController struct {
	rg *gin.RouterGroup
	ul usecase.LoanProductUseCase
	authMiddleware middleware.AuthMiddleware
}

func (l *LoanProductController) getHandler(ctx *gin.Context){
	id := ctx.Param("id")
	if id == ""{
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}
	loanProdutId := ctx.Param("Id")
	rspPayload, err := l.ul.FindById(loanProdutId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (l *LoanProductController) getAllHandler(ctx *gin.Context) {
	rspPayload, err := l.ul.GetAll()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Ok", rspPayload)
}

func (l *LoanProductController) createHandler(ctx *gin.Context) {
	var payload model.LoanProduct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request payload")
		return
	}

	createdProduct, err := l.ul.Create(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Created", createdProduct)
}
func (l *LoanProductController) updateHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	var payload model.LoanProduct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "invalid request payload")
		return
	}

	payload.Id = id // Set ID produk sesuai dengan parameter

	if err := l.ul.Update(payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Updated successfully", payload.Id)
}

func (l *LoanProductController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	if err := l.ul.Delete(id); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Deleted successfully", http.StatusOK,)
}

func NewLoanProductController(ul usecase.LoanProductUseCase, rg *gin.RouterGroup, authMiddleware middleware.AuthMiddleware) *LoanProductController {
	return &LoanProductController{ul: ul, rg: rg, authMiddleware: authMiddleware}
}