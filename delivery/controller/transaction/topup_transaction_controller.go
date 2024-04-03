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

// createHandler handles creating a new topup
// @Summary Create a new topup
// @Description Create a new topup
// @Tags Topup
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param body body dto.TopupDto true "Topup data"
// @Success 201 {string} string "Topup created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/topups [post]
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

// getTopupIdHandler handles getting a topup by ID
// @Summary Get topup by ID
// @Description Retrieve topup details by ID
// @Tags Topup
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param id path string true "Topup ID"
// @Success 200 {object} map[string]interface{} "Topup details"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/topups/{id} [get]
func (e *TopupController) getTopupIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := e.tc.GetTopUpByTopupId(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

// getTopupUserIdHandler handles getting topups by user ID
// @Summary Get topups by user ID
// @Description Retrieve topups for a specific user
// @Tags Topup
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} map[string]interface{} "List of topups"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/topups/user/{userId} [get]
func (e *TopupController) getTopupUserIdHandler(ctx *gin.Context) {
	userId := ctx.Param("userId")

	response, err := e.tc.GetTopupByUserId(userId)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(ctx, "ok", response)
}

// getAllTopupHandler handles getting all topups
// @Summary Get all topups
// @Description Retrieve all topups
// @Tags Topup
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer"
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of topups"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /transactions/topups [get]
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
