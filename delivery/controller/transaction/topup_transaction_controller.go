package controller

import (
	appconfig "medioker-bank/config/app_config"
	"medioker-bank/delivery/middleware"
	"medioker-bank/model/dto"
	usecase "medioker-bank/usecase/transaction"
	"medioker-bank/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TopupController struct {
	tc  usecase.TopupUseCase
	rg  *gin.RouterGroup
	jwt middleware.AuthMiddleware
}

func (t *TopupController) createHandler(ctx *gin.Context) {
	var payload dto.TopupDto
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

func (e *TopupController) getTopupIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := e.tc.GetTopUpByTopupId(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

func (e *TopupController) getTopupUserIdHandler(ctx *gin.Context) {
	userId := ctx.Param("userId")

	response, err := e.tc.GetTopupByUserId(userId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

func (u *TopupController) getAllTopupHandler(ctx *gin.Context) {
	topups, err := u.tc.GetAllTopUp()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Set response headers and send JSON data
	ctx.JSON(http.StatusOK, topups)
}

func (t *TopupController) Router() {
	ur := t.rg.Group(appconfig.TopupGroup)
	{
		ur.POST(appconfig.Topup, t.jwt.RequireToken("user"), t.createHandler)
		ur.GET(appconfig.Topup, t.jwt.RequireToken("admin"), t.getAllTopupHandler)
		ur.GET(appconfig.TopupId, t.jwt.RequireToken("admin"), t.getTopupIdHandler)
		ur.GET(appconfig.TopupUser, t.jwt.RequireToken("user"), t.getTopupUserIdHandler)
	}
}

func NewTopupController(tc usecase.TopupUseCase, rg *gin.RouterGroup, jwt middleware.AuthMiddleware) *TopupController {
	return &TopupController{tc: tc, rg: rg, jwt: jwt}
}
