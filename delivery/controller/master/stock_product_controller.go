package controller

import (
	"medioker-bank/config"
	"medioker-bank/model"
	"medioker-bank/model/dto"
	"medioker-bank/usecase"
	"medioker-bank/utils/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StockProductController struct {
	uc usecase.StockProductUseCase
	rg *gin.RouterGroup
}

func (s *StockProductController) createHandler(ctx *gin.Context) {
	var payload dto.StockProductCreateDto
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := s.uc.CreateStockProduct(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, "Success", response)
}

func (s *StockProductController) findByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := s.uc.FindStockProductById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (s *StockProductController) findManyHandler(ctx *gin.Context) {
	rating := ctx.Query("rating")
	risk := ctx.Query("risk")
	var response []model.StockProduct
	var err error
	if rating == "" && risk == "" {
		response, err = s.uc.FindAllStockProducts()
		if err != nil {
			common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if rating == "" {
			rating = "0"
		}
		intRating, err := strconv.Atoi(rating)
		if err != nil {
			common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		payload := dto.StockProductSearchByQueryDto{
			Rating: intRating,
			Risk:   ctx.Query("risk"),
		}
		response, err = s.uc.FindStockProductsByQuery(payload)
		if err != nil {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (s *StockProductController) updateHandler(ctx *gin.Context) {
	var payload dto.StockProductUpdateDto
	id := ctx.Param("id")
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := s.uc.UpdateStockProducts(id, payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (s *StockProductController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := s.uc.DeleteStockProducts(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Success", response)
}

func (s *StockProductController) Router() {
	spc := s.rg.Group(config.StockGroup)
	{
		spc.POST(config.StockCreate, s.createHandler)
		spc.GET(config.StockFindById, s.findByIdHandler)
		spc.GET(config.StockFindMany, s.findManyHandler)
		spc.PUT(config.StockUpdate, s.updateHandler)
		spc.DELETE(config.StockDelete, s.deleteHandler)
	}
}

func NewStockProductController(uc usecase.StockProductUseCase, rg *gin.RouterGroup) *StockProductController {
	return &StockProductController{uc: uc, rg: rg}
}
