package controller

import (
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/delivery/middleware"
	"medioker-bank/model"
	usecase "medioker-bank/usecase/master"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoanProductController struct {
	rg  *gin.RouterGroup
	ul  usecase.LoanProductUseCase
	jwt middleware.AuthMiddleware
}

func (l *LoanProductController) GetHandlerById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}
	response, err := l.ul.FindLoanProductById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succes", response)
}

func (l *LoanProductController) GetAllHandler(ctx *gin.Context) {
	response, err := l.ul.FindAllLoanProduct()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succes", response)
}

func (l *LoanProductController) CreateHandler(ctx *gin.Context) {
	var payload model.LoanProduct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdProduct, err := l.ul.CreateLoanProduct(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Created", createdProduct)
}
func (l *LoanProductController) UpdateHandler(ctx *gin.Context) {
	var payload model.LoanProduct
	id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	err = l.ul.UpdateLoanProduct(id, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response := "success update loan_products with id: " + id
	common.SendSingleResponse(ctx, "Updated successfully", response)
}

func (l *LoanProductController) DeleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		common.SendErrorResponse(ctx, http.StatusBadRequest, "id can't be empty")
		return
	}

	_, err := l.ul.DeleteLoanProduct(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response := "success delete loan_products with id: " + id
	common.SendSingleResponse(ctx, "Deleted successfully", response)
}

func (l *LoanProductController) Router() {
	spc := l.rg.Group(appconfig.LoanProductGroup)
	{
		spc.POST(appconfig.LoanProductCreate, l.jwt.RequireToken("admin"), l.CreateHandler)
		spc.GET(appconfig.LoanProductFindByid, l.jwt.RequireToken("admin", "user"), l.GetHandlerById)
		spc.GET(appconfig.LoanProductFindAll, l.jwt.RequireToken("admin", "user"), l.GetAllHandler)
		spc.PUT(appconfig.LoanProductUpdate, l.jwt.RequireToken("admin"), l.UpdateHandler)
		spc.DELETE(appconfig.LoanProductDelete, l.jwt.RequireToken("admin"), l.DeleteHandler)
	}
}

func NewLoanProductController(ul usecase.LoanProductUseCase, rg *gin.RouterGroup, jwt middleware.AuthMiddleware) *LoanProductController {
	return &LoanProductController{ul: ul, rg: rg, jwt: jwt}
}
