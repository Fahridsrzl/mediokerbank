package controller

import (
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/model"
	usecase "medioker-bank/usecase/transaction"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TopupController struct {
	tc usecase.TopupUseCase
	rg *gin.RouterGroup
}

func (t *TopupController) createHandler (ctx *gin.Context) {
	var payload model.TopupTransaction
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	respPayload, err := t.tc.CreateTopup(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "ok", respPayload)
}

func (t *TopupController) Router() {
	ur := t.rg.Group(appconfig.TopupGroup)
	ur.POST(appconfig.Topup, t.createHandler)
}

func NewTopupController(tc usecase.TopupUseCase, rg *gin.RouterGroup) *TopupController {
	return &TopupController{tc: tc, rg: rg}
}