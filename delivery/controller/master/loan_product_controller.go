package controller

import (
	"errors"
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/delivery/middleware"
	"medioker-bank/model"
	usecase "medioker-bank/usecase/master"
	"medioker-bank/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type LoanProductController struct {
	rg  *gin.RouterGroup
	ul  usecase.LoanProductUseCase
	jwt middleware.AuthMiddleware
}

// @Summary Get loan product by ID
// @Description Get loan product details by ID
// @ID get-loan-product-by-id
// @Param id path string true "Loan Product ID"
// @Tags Loan Product
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Produce json
// @Success 200 {object} model.LoanProduct
// @Failure 400 {object} Status
// @Failure 500 {object} Status
// @Router /loan-products/{id} [get]
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

// @Summary Get all loan products
// @Description Get all loan products
// @ID get-all-loan-products
// @Tags Loan Product
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Produce json
// @Success 200 {array} model.LoanProduct
// @Failure 500 {object} Status
// @Router /loan-products [get]
func (l *LoanProductController) GetAllHandler(ctx *gin.Context) {
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
	response, err := l.ul.FindAllLoanProduct(page, limit)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "Succes", response)
}

// @Summary Create a new loan product
// @Description Create a new loan product
// @ID create-loan-product
// @Tags Loan Product
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param request body model.LoanProduct true "Loan Product Object"
// @Success 200 {object} model.LoanProduct
// @Failure 400 {object} Status
// @Failure 500 {object} Status
// @Router /loan-products [post]
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

// / @Summary Update a loan product
// @Description Update an existing loan product
// @ID update-loan-product
// @Tags Loan Product
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param id path string true "Loan Product ID"
// @Param request body model.LoanProduct true "Updated Loan Product Object"
// @Success 200 {object} model.LoanProduct
// @Failure 400 {object} Status
// @Failure 500 {object} Status
// @Router /loan-products/{id} [put]
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

// @Summary Delete a loan product
// @Description Delete an existing loan product
// @ID delete-loan-product
// @Param id path string true "Loan Product ID"
// @Tags Loan Product
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Produce json
// @Success 200 {object} model.LoanProduct
// @Failure 400 {object} Status
// @Failure 500 {object} Status
// @Router /loan-products/{id} [delete]
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
